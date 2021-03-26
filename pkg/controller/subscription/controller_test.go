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
	e2smtypes "github.com/onosproject/onos-api/go/onos/e2t/e2sm"
	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
	"github.com/onosproject/onos-e2t/pkg/oid"
	e2server "github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/server"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/types"
	"github.com/onosproject/onos-lib-go/pkg/controller"
	"github.com/stretchr/testify/assert"
	"testing"
)

type controllerTestContext struct {
	ctrl                   *gomock.Controller
	serverChannel          *MockRICChannel
	channelManager         *MockChannelManager
	subscriptionClient     subapi.E2SubscriptionServiceClient
	subscriptionTaskClient subtaskapi.E2SubscriptionTaskServiceClient
	requestJournal         *RequestJournal
	modelRegistry          *MockModelRegistry
	oidRegistry            *MockRegistry
	reconciler             Reconciler
	controller             *controller.Controller
}

func initControllerTest(t *testing.T, testContext *controllerTestContext) {
	initControllerTestNoRICSubscription(t, testContext)
	resp := e2appducontents.RicsubscriptionResponse{}
	testContext.serverChannel.EXPECT().RICSubscription(gomock.Any(), gomock.Any()).Return(&resp, nil, nil).AnyTimes()
}

func initControllerTestNoRICSubscription(t *testing.T, testContext *controllerTestContext) {
	// subscription and task clients
	testContext.subscriptionClient, testContext.subscriptionTaskClient = createClients(t)

	// Controller for mocks
	ctrl := gomock.NewController(t)

	// Mock channel manager
	testContext.channelManager = NewMockChannelManager(ctrl)

	// Mock request journal
	testContext.requestJournal = NewRequestJournal()

	// Function ID map for mocked service models
	modelFuncIDs := make(map[e2smtypes.OID]types.RanFunctionID)
	modelFuncIDs["sm1"] = 2

	// Mocked RIC channel
	serverChannel := NewMockRICChannel(ctrl)
	testContext.serverChannel = serverChannel
	channel := e2server.NewE2Channel("channel", "plmnid", serverChannel, modelFuncIDs)
	testContext.channelManager = NewMockChannelManager(ctrl)
	testContext.channelManager.EXPECT().Get(gomock.Any(), gomock.Any()).Return(channel, nil)

	// Mocked model registry
	testContext.modelRegistry = NewMockModelRegistry(ctrl)
	testContext.ctrl = ctrl

	// Controller
	testContext.controller = NewController(testContext.requestJournal, testContext.subscriptionClient,
		testContext.subscriptionTaskClient, testContext.channelManager, testContext.modelRegistry, testContext.oidRegistry)

	// OID registry
	testContext.oidRegistry = NewMockRegistry(ctrl)
	testContext.oidRegistry.EXPECT().GetOid(gomock.Any()).Return(oid.Oid(12)).AnyTimes()

	// reconciler to test
	testContext.reconciler = Reconciler{
		catalog:     testContext.requestJournal,
		subs:        testContext.subscriptionClient,
		tasks:       testContext.subscriptionTaskClient,
		channels:    testContext.channelManager,
		models:      testContext.modelRegistry,
		oidRegistry: testContext.oidRegistry,
	}
}

func TestOpenNoPlugin(t *testing.T) {
	var testContext controllerTestContext
	createServerScaffolding(t)
	initControllerTest(t, &testContext)

	testContext.modelRegistry.EXPECT().GetPlugin(gomock.Any()).Return(nil, errors.New("no such model"))

	subscription := &subapi.Subscription{
		ID: "1", AppID: "foo", Details: &subapi.SubscriptionDetails{E2NodeID: E2NodeID, ServiceModel: subapi.ServiceModel{Name: "sm1"}},
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
	smd := e2smtypes.ServiceModelData{}
	sm.EXPECT().ServiceModelData().Return(smd)
	sm.EXPECT().EventTriggerDefinitionProtoToASN1(gomock.Any()).Return(nil, errors.New("this should fail"))
	testContext.modelRegistry.EXPECT().GetPlugin(gomock.Any()).Return(sm, nil)

	subscription := &subapi.Subscription{
		ID: "1", AppID: "foo", Details: &subapi.SubscriptionDetails{E2NodeID: E2NodeID, ServiceModel: subapi.ServiceModel{Name: "sm1"}},
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
	smd := e2smtypes.ServiceModelData{}
	sm.EXPECT().ServiceModelData().Return(smd)
	testContext.modelRegistry.EXPECT().GetPlugin(gomock.Any()).Return(sm, nil)

	subscription := &subapi.Subscription{
		ID: "1", AppID: "foo", Details: &subapi.SubscriptionDetails{E2NodeID: E2NodeID, ServiceModel: subapi.ServiceModel{Name: "sm1"}},
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
	smd := e2smtypes.ServiceModelData{}
	sm.EXPECT().ServiceModelData().Return(smd)
	testContext.modelRegistry.EXPECT().GetPlugin(gomock.Any()).Return(sm, nil)

	subscription := &subapi.Subscription{
		ID: "1", AppID: "foo", Details: &subapi.SubscriptionDetails{E2NodeID: E2NodeID, ServiceModel: subapi.ServiceModel{Name: "sm1"}},
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

func TestOpenActionBadProtocolError(t *testing.T) {
	var testContext controllerTestContext
	createServerScaffolding(t)
	initControllerTest(t, &testContext)

	sm := NewMockServiceModel(testContext.ctrl)
	smd := e2smtypes.ServiceModelData{}
	sm.EXPECT().ServiceModelData().Return(smd)
	sm.EXPECT().EventTriggerDefinitionProtoToASN1(gomock.Any()).Return(make([]byte, 1), nil)
	testContext.modelRegistry.EXPECT().GetPlugin(gomock.Any()).Return(sm, nil)

	action := []subapi.Action{{
		Payload: subapi.Payload{Encoding: 55},
	}}
	subscription := &subapi.Subscription{
		ID: "1", AppID: "foo",
		Details: &subapi.SubscriptionDetails{
			E2NodeID:     E2NodeID,
			ServiceModel: subapi.ServiceModel{Name: "sm1"},
			Actions:      action,
		},
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
	assert.Error(t, err)

	updatedTask, err := scaffold.taskStore.Get(context.Background(), subTask.ID)
	assert.NoError(t, err)
	assert.NotNil(t, updatedTask)
	assert.Equal(t, subtaskapi.Status_FAILED, updatedTask.Lifecycle.Status)
	assert.Equal(t, subtaskapi.Cause_CAUSE_PROTOCOL_MESSAGE_NOT_COMPATIBLE_WITH_RECEIVER_STATE, updatedTask.Lifecycle.Failure.Cause)

	testContext.ctrl.Finish()
}

func TestOpenActionBadProtoPayload(t *testing.T) {
	var testContext controllerTestContext
	createServerScaffolding(t)
	initControllerTest(t, &testContext)

	sm := NewMockServiceModel(testContext.ctrl)
	smd := e2smtypes.ServiceModelData{}
	sm.EXPECT().ServiceModelData().Return(smd)
	sm.EXPECT().EventTriggerDefinitionProtoToASN1(gomock.Any()).Return(make([]byte, 1), nil)
	sm.EXPECT().ActionDefinitionProtoToASN1(gomock.Any()).Return(nil, errors.New("bad proto payload"))
	testContext.modelRegistry.EXPECT().GetPlugin(gomock.Any()).Return(sm, nil)

	action := []subapi.Action{{
		Payload: subapi.Payload{Encoding: subapi.Encoding_ENCODING_PROTO},
	}}
	subscription := &subapi.Subscription{
		ID: "1", AppID: "foo",
		Details: &subapi.SubscriptionDetails{
			E2NodeID:     E2NodeID,
			ServiceModel: subapi.ServiceModel{Name: "sm1"},
			Actions:      action,
		},
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
	assert.Error(t, err)

	updatedTask, err := scaffold.taskStore.Get(context.Background(), subTask.ID)
	assert.NoError(t, err)
	assert.NotNil(t, updatedTask)
	assert.Equal(t, subtaskapi.Status_FAILED, updatedTask.Lifecycle.Status)
	assert.Equal(t, subtaskapi.Cause_CAUSE_PROTOCOL_ABSTRACT_SYNTAX_ERROR_FALSELY_CONSTRUCTED_MESSAGE, updatedTask.Lifecycle.Failure.Cause)

	testContext.ctrl.Finish()
}

func TestOpenBadChannelResponse(t *testing.T) {
	var testContext controllerTestContext
	createServerScaffolding(t)
	initControllerTestNoRICSubscription(t, &testContext)
	failure := e2appducontents.RicsubscriptionFailure{
		ProtocolIes: &e2appducontents.RicsubscriptionFailureIes{
			E2ApProtocolIes29: nil,
			E2ApProtocolIes5:  nil,
			E2ApProtocolIes18: &e2appducontents.RicsubscriptionFailureIes_RicsubscriptionFailureIes18{
				Id:          0,
				Criticality: 0,
				Value: &e2appducontents.RicactionNotAdmittedList{
					Value: []*e2appducontents.RicactionNotAdmittedItemIes{
						{Value: &e2appducontents.RicactionNotAdmittedItem{
							Cause: &e2apies.Cause{
								Cause: &e2apies.Cause_RicRequest{RicRequest: e2apies.CauseRic_CAUSE_RIC_RAN_FUNCTION_ID_INVALID},
							},
						}},
					},
				},
				Presence: 0,
			},
			E2ApProtocolIes2: nil,
		},
	}
	testContext.serverChannel.EXPECT().RICSubscription(gomock.Any(), gomock.Any()).Return(nil, &failure, nil).AnyTimes()

	sm := NewMockServiceModel(testContext.ctrl)
	smd := e2smtypes.ServiceModelData{}
	sm.EXPECT().ServiceModelData().Return(smd)
	testContext.modelRegistry.EXPECT().GetPlugin(gomock.Any()).Return(sm, nil)

	subscription := &subapi.Subscription{
		ID: "1", AppID: "foo", Details: &subapi.SubscriptionDetails{E2NodeID: E2NodeID, ServiceModel: subapi.ServiceModel{Name: "sm1"}},
	}
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
	assert.Equal(t, subtaskapi.Cause_CAUSE_RIC_RAN_FUNCTION_ID_INVALID, updatedTask.Lifecycle.Failure.Cause)

	testContext.ctrl.Finish()
}

func TestOpenAction(t *testing.T) {
	var testContext controllerTestContext
	createServerScaffolding(t)
	initControllerTest(t, &testContext)

	sm := NewMockServiceModel(testContext.ctrl)
	smd := e2smtypes.ServiceModelData{}
	sm.EXPECT().ServiceModelData().Return(smd)
	sm.EXPECT().EventTriggerDefinitionProtoToASN1(gomock.Any()).Return(make([]byte, 1), nil)
	sm.EXPECT().ActionDefinitionProtoToASN1(gomock.Any()).Return(make([]byte, 1), nil)
	testContext.modelRegistry.EXPECT().GetPlugin(gomock.Any()).Return(sm, nil)

	action := []subapi.Action{{
		Payload: subapi.Payload{Encoding: subapi.Encoding_ENCODING_PROTO},
		SubsequentAction: &subapi.SubsequentAction{
			Type:       subapi.SubsequentActionType_SUBSEQUENT_ACTION_TYPE_CONTINUE,
			TimeToWait: 5,
		},
	}}
	subscription := &subapi.Subscription{
		ID: "1", AppID: "foo",
		Details: &subapi.SubscriptionDetails{
			E2NodeID:     E2NodeID,
			ServiceModel: subapi.ServiceModel{Name: "sm1"},
			Actions:      action,
		},
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
	assert.Equal(t, subtaskapi.Status_COMPLETE, updatedTask.Lifecycle.Status)

	testContext.ctrl.Finish()
}

func TestClose(t *testing.T) {
	var testContext controllerTestContext
	createServerScaffolding(t)
	initControllerTest(t, &testContext)
	response := e2appducontents.RicsubscriptionDeleteResponse{}
	testContext.serverChannel.EXPECT().RICSubscriptionDelete(gomock.Any(), gomock.Any()).Return(&response, nil, nil).AnyTimes()

	subscription := &subapi.Subscription{
		ID: "1", AppID: "foo", Details: &subapi.SubscriptionDetails{E2NodeID: E2NodeID, ServiceModel: subapi.ServiceModel{Name: "sm1"}},
	}
	_, err := testContext.subscriptionClient.AddSubscription(context.Background(), &subapi.AddSubscriptionRequest{
		Subscription: subscription,
	})
	assert.NoError(t, err)

	_, err = testContext.subscriptionClient.RemoveSubscription(context.Background(), &subapi.RemoveSubscriptionRequest{
		ID: "1",
	})
	assert.NoError(t, err)

	subTask.Lifecycle.Phase = subtaskapi.Phase_CLOSE
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
