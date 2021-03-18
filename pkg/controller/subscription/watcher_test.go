// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package subscription

import (
	"context"
	subapi "github.com/onosproject/onos-api/go/onos/e2sub/subscription"
	server2 "github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/server"
	"github.com/onosproject/onos-lib-go/pkg/controller"
	"sync"

	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWatcher(t *testing.T) {
	createServerScaffolding(t)
	_, subscriptionTaskClient := createClients(t)

	watch := Watcher{
		endpointID: E2NodeID,
		tasks:      subscriptionTaskClient,
		cancel:     nil,
		mu:         sync.Mutex{},
	}
	ch := make(chan controller.ID)
	err := watch.Start(ch)
	assert.NoError(t, err)

	err = scaffold.taskStore.Create(context.Background(), subTask)
	assert.NoError(t, err)

	newID := <-ch
	assert.Equal(t, controller.NewID(subTask.ID), newID)

	watch.Stop()
}

func TestChannelWatcher(t *testing.T) {
	createServerScaffolding(t)
	subscriptionClient, subscriptionTaskClient := createClients(t)

	watch := ChannelWatcher{
		endpointID: E2NodeID,
		tasks:      subscriptionTaskClient,
		subs:       subscriptionClient,
		channels:   server2.NewChannelManager(),
		cancel:     nil,
		mu:         sync.Mutex{},
	}

	ch := make(chan controller.ID)
	err := watch.Start(ch)
	assert.NoError(t, err)

	err = scaffold.taskStore.Create(context.Background(), subTask)
	assert.NoError(t, err)

	subscription := &subapi.Subscription{
		ID: "1", AppID: "foo", Details: &subapi.SubscriptionDetails{E2NodeID: E2NodeID, ServiceModel: subapi.ServiceModel{ID: "sm1"}},
	}
	_, err = subscriptionClient.AddSubscription(context.Background(), &subapi.AddSubscriptionRequest{
		Subscription: subscription,
	})
	assert.NoError(t, err)

	watch.channelCh <- &server2.E2Channel{ID: E2NodeID}
	newID := <-ch
	assert.Equal(t, controller.NewID(subTask.ID), newID)

	watch.Stop()
}
