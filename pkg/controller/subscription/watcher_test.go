// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package subscription

import (
	"context"
	"sync"

	"github.com/onosproject/onos-e2t/pkg/topo"

	"github.com/onosproject/onos-e2t/pkg/store/rnib"

	subapi "github.com/onosproject/onos-api/go/onos/e2sub/subscription"
	server2 "github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/server"
	"github.com/onosproject/onos-lib-go/pkg/controller"

	"testing"

	"github.com/stretchr/testify/assert"
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
	// TODO this should be changed to mock the store but the test is not using it
	store, _ := rnib.NewStore("")
	topoManager := topo.NewManager(store)

	watch := ChannelWatcher{
		endpointID: E2NodeID,
		tasks:      subscriptionTaskClient,
		subs:       subscriptionClient,
		channels:   server2.NewChannelManager(topoManager),
		cancel:     nil,
		mu:         sync.Mutex{},
	}

	ch := make(chan controller.ID)
	err := watch.Start(ch)
	assert.NoError(t, err)

	err = scaffold.taskStore.Create(context.Background(), subTask)
	assert.NoError(t, err)

	subscription := &subapi.Subscription{
		ID: "1", AppID: "foo", Details: &subapi.SubscriptionDetails{E2NodeID: E2NodeID, ServiceModel: subapi.ServiceModel{
			Name:    "sm1",
			Version: "v1",
		}},
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
