// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package subscription

import (
	"context"
	"github.com/golang/mock/gomock"
	subapi "github.com/onosproject/onos-api/go/onos/e2sub/subscription"
	subtaskapi "github.com/onosproject/onos-api/go/onos/e2sub/task"
	"github.com/onosproject/onos-e2t/pkg/modelregistry"
	e2server "github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/server"
	"github.com/onosproject/onos-lib-go/pkg/controller"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOpenNoPlugin(t *testing.T) {
	ctrl := gomock.NewController(t)
	channelManager := NewMockChannelManager(ctrl)
	var channel *e2server.E2Channel
	channelManager.EXPECT().Get(gomock.Any(), gomock.Any()).Return(channel, nil)

	createServerScaffolding(t)
	subscriptionClient, subscriptionTaskClient := createClients(t)
	requestJournal := NewRequestJournal()

	modelRegistry := modelregistry.NewModelRegistry()
	c := NewController(requestJournal, subscriptionClient, subscriptionTaskClient, channelManager, &modelRegistry)
	assert.NotNil(t, c)

	subscription := &subapi.Subscription{
		ID: "1", AppID: "foo", Details: &subapi.SubscriptionDetails{E2NodeID: E2NodeID, ServiceModel: subapi.ServiceModel{ID: "sm1"}},
	}
	_, err := subscriptionClient.AddSubscription(context.Background(), &subapi.AddSubscriptionRequest{
		Subscription: subscription,
	})
	assert.NoError(t, err)

	reconciler := Reconciler{
		catalog:   requestJournal,
		subs:      subscriptionClient,
		tasks:     subscriptionTaskClient,
		channels:  channelManager,
		models:    &modelRegistry,
		requestID: 0,
	}
	err = scaffold.taskStore.Create(context.Background(), subTask)
	assert.NoError(t, err)

	result, err := reconciler.Reconcile(controller.ID{Value: subTask.ID})

	assert.NotNil(t, result)
	assert.NotNil(t, err)

	updatedTask, err := scaffold.taskStore.Get(context.Background(), subTask.ID)
	assert.Nil(t, err)
	assert.NotNil(t, updatedTask)
	assert.Equal(t, subtaskapi.Status_FAILED, updatedTask.Lifecycle.Status)
	assert.Equal(t, subtaskapi.Cause_CAUSE_RIC_RAN_FUNCTION_ID_INVALID, updatedTask.Lifecycle.Failure.Cause)
}

func TestOpenValidPlugin(t *testing.T) {
	ctrl := gomock.NewController(t)
	channelManager := NewMockChannelManager(ctrl)
	var channel *e2server.E2Channel
	channelManager.EXPECT().Get(gomock.Any(), gomock.Any()).Return(channel, nil)
	modelRegistry := modelregistry.NewModelRegistry()
	sm := NewMockServiceModel(ctrl)
	modelRegistry.RegisterModelPlugin()

	createServerScaffolding(t)
	subscriptionClient, subscriptionTaskClient := createClients(t)
	requestJournal := NewRequestJournal()

	c := NewController(requestJournal, subscriptionClient, subscriptionTaskClient, channelManager, &modelRegistry)
	assert.NotNil(t, c)

	subscription := &subapi.Subscription{
		ID: "1", AppID: "foo", Details: &subapi.SubscriptionDetails{E2NodeID: E2NodeID, ServiceModel: subapi.ServiceModel{ID: "sm1"}},
	}
	_, err := subscriptionClient.AddSubscription(context.Background(), &subapi.AddSubscriptionRequest{
		Subscription: subscription,
	})
	assert.NoError(t, err)

	reconciler := Reconciler{
		catalog:   requestJournal,
		subs:      subscriptionClient,
		tasks:     subscriptionTaskClient,
		channels:  channelManager,
		models:    &modelRegistry,
		requestID: 0,
	}
	err = scaffold.taskStore.Create(context.Background(), subTask)
	assert.NoError(t, err)

	result, err := reconciler.Reconcile(controller.ID{Value: subTask.ID})

	assert.NotNil(t, result)
	assert.NotNil(t, err)

	updatedTask, err := scaffold.taskStore.Get(context.Background(), subTask.ID)
	assert.Nil(t, err)
	assert.NotNil(t, updatedTask)
	assert.Equal(t, subtaskapi.Status_FAILED, updatedTask.Lifecycle.Status)
	assert.Equal(t, subtaskapi.Cause_CAUSE_RIC_RAN_FUNCTION_ID_INVALID, updatedTask.Lifecycle.Failure.Cause)
}
