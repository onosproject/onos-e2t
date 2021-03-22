// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package subscription

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	subapi "github.com/onosproject/onos-api/go/onos/e2sub/subscription"
	subtaskapi "github.com/onosproject/onos-api/go/onos/e2sub/task"
	e2ap_pdu_contents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
	"github.com/onosproject/onos-e2t/pkg/modelregistry"
	e2server "github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/server"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/types"
	"github.com/onosproject/onos-lib-go/pkg/controller"
	"github.com/stretchr/testify/assert"
	"testing"
)

type controllerTestContext struct {
	ctrl                   *gomock.Controller
	channelManager         *MockChannelManager
	subscriptionClient     subapi.E2SubscriptionServiceClient
	subscriptionTaskClient subtaskapi.E2SubscriptionTaskServiceClient
	requestJournal         *RequestJournal
	modelRegistry          *MockModelRegistry
	reconciler             Reconciler
	controller             *controller.Controller
}

func initControllerTest(t *testing.T, testContext *controllerTestContext) {
	// subscription and task clients
	testContext.subscriptionClient, testContext.subscriptionTaskClient = createClients(t)

	// Controller for mocks
	ctrl := gomock.NewController(t)

	// Mock channel manager
	testContext.channelManager = NewMockChannelManager(ctrl)

	// Mock request journal
	testContext.requestJournal = NewRequestJournal()

	// Function ID map for mocked service models
	modelFuncIDs := make(map[modelregistry.ModelFullName]types.RanFunctionID)
	modelFuncIDs["sm1"] = 2

	// Mocked RIC channel
	serverChannel := NewMockRICChannel(ctrl)
	resp := e2ap_pdu_contents.RicsubscriptionResponse{}
	serverChannel.EXPECT().RICSubscription(gomock.Any(), gomock.Any()).Return(&resp, nil, nil).AnyTimes()
	channel := e2server.NewE2Channel("channel", "plmnid", serverChannel, modelFuncIDs)
	testContext.channelManager = NewMockChannelManager(ctrl)
	testContext.channelManager.EXPECT().Get(gomock.Any(), gomock.Any()).Return(channel, nil)

	// Mocked model registry
	testContext.modelRegistry = NewMockModelRegistry(ctrl)
	testContext.ctrl = ctrl

	// Controller
	testContext.controller = NewController(testContext.requestJournal, testContext.subscriptionClient, testContext.subscriptionTaskClient, testContext.channelManager, testContext.modelRegistry)


		// reconciler to test
	testContext.reconciler = Reconciler{
		catalog:  testContext.requestJournal,
		subs:     testContext.subscriptionClient,
		tasks:    testContext.subscriptionTaskClient,
		channels: testContext.channelManager,
		models:   testContext.modelRegistry,
	}
}

func TestOpenNoPlugin(t *testing.T) {
	var testContext controllerTestContext
	createServerScaffolding(t)
	initControllerTest(t, &testContext)

	testContext.modelRegistry.EXPECT().GetPlugin(modelregistry.ModelFullName("sm1")).Return(nil, errors.New("no such model"))

	subscription := &subapi.Subscription{
		ID: "1", AppID: "foo", Details: &subapi.SubscriptionDetails{E2NodeID: E2NodeID, ServiceModel: subapi.ServiceModel{ID: "sm1"}},
	}
	_, err := testContext.subscriptionClient.AddSubscription(context.Background(), &subapi.AddSubscriptionRequest{
		Subscription: subscription,
	})
	assert.NoError(t, err)

	err = scaffold.taskStore.Create(context.Background(), subTask)
	assert.NoError(t, err)

	result, err := testContext.reconciler.Reconcile(controller.ID{Value: subTask.ID})
	assert.NotNil(t, result)
	assert.NotNil(t, err)

	updatedTask, err := scaffold.taskStore.Get(context.Background(), subTask.ID)
	assert.Nil(t, err)
	assert.NotNil(t, updatedTask)
	assert.Equal(t, subtaskapi.Status_FAILED, updatedTask.Lifecycle.Status)
	assert.Equal(t, subtaskapi.Cause_CAUSE_RIC_RAN_FUNCTION_ID_INVALID, updatedTask.Lifecycle.Failure.Cause)

	testContext.ctrl.Finish()
}

func TestOpenProtoToASNError(t *testing.T) {
	var testContext controllerTestContext
	createServerScaffolding(t)
	initControllerTest(t, &testContext)

	sm := NewMockServiceModel(testContext.ctrl)
	sm.EXPECT().ServiceModelData().Return("D1", "D2", "D3")
	sm.EXPECT().EventTriggerDefinitionProtoToASN1(gomock.Any()).Return(nil, errors.New("this should fail"))
	testContext.modelRegistry.EXPECT().GetPlugin(gomock.Any()).Return(sm, nil)

	subscription := &subapi.Subscription{
		ID: "1", AppID: "foo", Details: &subapi.SubscriptionDetails{E2NodeID: E2NodeID, ServiceModel: subapi.ServiceModel{ID: "sm1"}},
	}
	subscription.Details.EventTrigger.Payload.Encoding = subapi.Encoding_ENCODING_PROTO

	_, err := testContext.subscriptionClient.AddSubscription(context.Background(), &subapi.AddSubscriptionRequest{
		Subscription: subscription,
	})
	assert.NoError(t, err)

	err = scaffold.taskStore.Create(context.Background(), subTask)
	assert.NoError(t, err)

	result, err := testContext.reconciler.Reconcile(controller.ID{Value: subTask.ID})
	assert.NotNil(t, result)
	assert.NoError(t, err)

	updatedTask, err := scaffold.taskStore.Get(context.Background(), subTask.ID)
	assert.NoError(t, err)
	assert.NotNil(t, updatedTask)
	assert.Equal(t, subtaskapi.Status_FAILED, updatedTask.Lifecycle.Status)
	assert.Equal(t, subtaskapi.Cause_CAUSE_PROTOCOL_ABSTRACT_SYNTAX_ERROR_FALSELY_CONSTRUCTED_MESSAGE, updatedTask.Lifecycle.Failure.Cause)

	testContext.ctrl.Finish()
}

func TestOpenBadProtocolError(t *testing.T) {
	var testContext controllerTestContext
	createServerScaffolding(t)
	initControllerTest(t, &testContext)

	sm := NewMockServiceModel(testContext.ctrl)
	sm.EXPECT().ServiceModelData().Return("D1", "D2", "D3")
	testContext.modelRegistry.EXPECT().GetPlugin(gomock.Any()).Return(sm, nil)

	subscription := &subapi.Subscription{
		ID: "1", AppID: "foo", Details: &subapi.SubscriptionDetails{E2NodeID: E2NodeID, ServiceModel: subapi.ServiceModel{ID: "sm1"}},
	}
	subscription.Details.EventTrigger.Payload.Encoding = 123

	_, err := testContext.subscriptionClient.AddSubscription(context.Background(), &subapi.AddSubscriptionRequest{
		Subscription: subscription,
	})
	assert.NoError(t, err)

	err = scaffold.taskStore.Create(context.Background(), subTask)
	assert.NoError(t, err)

	result, err := testContext.reconciler.Reconcile(controller.ID{Value: subTask.ID})
	assert.NotNil(t, result)
	assert.Error(t, err)

	updatedTask, err := scaffold.taskStore.Get(context.Background(), subTask.ID)
	assert.NoError(t, err)
	assert.NotNil(t, updatedTask)
	assert.Equal(t, subtaskapi.Status_FAILED, updatedTask.Lifecycle.Status)
	assert.Equal(t, subtaskapi.Cause_CAUSE_PROTOCOL_MESSAGE_NOT_COMPATIBLE_WITH_RECEIVER_STATE, updatedTask.Lifecycle.Failure.Cause)

	testContext.ctrl.Finish()
}

func TestOpenValidPlugin(t *testing.T) {
	var testContext controllerTestContext
	createServerScaffolding(t)
	initControllerTest(t, &testContext)

	sm := NewMockServiceModel(testContext.ctrl)
	sm.EXPECT().ServiceModelData().Return("D1", "D2", "D3")
	testContext.modelRegistry.EXPECT().GetPlugin(gomock.Any()).Return(sm, nil)

	subscription := &subapi.Subscription{
		ID: "1", AppID: "foo", Details: &subapi.SubscriptionDetails{E2NodeID: E2NodeID, ServiceModel: subapi.ServiceModel{ID: "sm1"}},
	}
	_, err := testContext.subscriptionClient.AddSubscription(context.Background(), &subapi.AddSubscriptionRequest{
		Subscription: subscription,
	})
	assert.NoError(t, err)

	err = scaffold.taskStore.Create(context.Background(), subTask)
	assert.NoError(t, err)

	result, err := testContext.reconciler.Reconcile(controller.ID{Value: subTask.ID})
	assert.NotNil(t, result)
	assert.NoError(t, err)

	updatedTask, err := scaffold.taskStore.Get(context.Background(), subTask.ID)
	assert.NoError(t, err)
	assert.NotNil(t, updatedTask)
	assert.Equal(t, subtaskapi.Status_COMPLETE, updatedTask.Lifecycle.Status)

	testContext.ctrl.Finish()
}
