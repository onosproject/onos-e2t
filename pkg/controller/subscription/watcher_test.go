// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package subscription

import (
	"context"
	subapi "github.com/onosproject/onos-api/go/onos/e2sub/subscription"
	subtaskapi "github.com/onosproject/onos-api/go/onos/e2sub/task"
	e2sub "github.com/onosproject/onos-e2sub/pkg/northbound/subscription"
	e2task "github.com/onosproject/onos-e2sub/pkg/northbound/task"
	server2 "github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/server"
	"github.com/onosproject/onos-lib-go/pkg/controller"
	"sync"

	substore "github.com/onosproject/onos-e2sub/pkg/store/subscription"
	taskstore "github.com/onosproject/onos-e2sub/pkg/store/task"

	"github.com/onosproject/onos-lib-go/pkg/northbound"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"net"
	"testing"
)

var (
	lis       = bufconn.Listen(1024 * 1024)
	server    = grpc.NewServer()
	taskStore taskstore.Store
	subStore  substore.Store
)

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func newTestSubService(t *testing.T) northbound.Service {
	var err error
	subStore, err = substore.NewLocalStore()
	assert.NoError(t, err)

	return e2sub.NewService(subStore)
}

func newTestTaskService(t *testing.T) northbound.Service {
	var err error
	taskStore, err = taskstore.NewLocalStore()
	assert.NoError(t, err)

	return e2task.NewService(taskStore)
}

func createSubServerConnection(t *testing.T) *grpc.ClientConn {
	lis = bufconn.Listen(1024 * 1024)
	s := newTestSubService(t)
	assert.NotNil(t, s)

	s.Register(server)

	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	return conn
}

func createTaskServerConnection(t *testing.T) *grpc.ClientConn {
	s := newTestTaskService(t)
	assert.NotNil(t, s)

	s.Register(server)

	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	return conn
}

func serve(t *testing.T) {
	go func() {
		if err := server.Serve(lis); err != nil {
			assert.NoError(t, err, "Server exited with error: %v", err)
		}
	}()
}

func createServers(t *testing.T) (*grpc.ClientConn, *grpc.ClientConn) {
	subConn := createSubServerConnection(t)
	taskConn := createTaskServerConnection(t)
	serve(t)
	return subConn, taskConn
}

func TestWatcher(t *testing.T) {
	subConn, taskConn := createServers(t)

	subscriptionClient := subapi.NewE2SubscriptionServiceClient(subConn)
	assert.NotNil(t, subscriptionClient)

	subscriptionTaskClient := subtaskapi.NewE2SubscriptionTaskServiceClient(taskConn)
	assert.NotNil(t, subscriptionTaskClient)

	watch := Watcher{
		endpointID: "1",
		tasks:      subscriptionTaskClient,
		cancel:     nil,
		mu:         sync.Mutex{},
	}
	ch := make(chan controller.ID)
	err := watch.Start(ch)
	assert.NoError(t, err)

	subTask := &subtaskapi.SubscriptionTask{
		ID:             "1",
		Revision:       0,
		SubscriptionID: "1",
		EndpointID:     "1",
		Lifecycle:      subtaskapi.Lifecycle{},
	}

	err = taskStore.Create(context.Background(), subTask)
	assert.NoError(t, err)

	newID := <-ch
	assert.Equal(t, controller.NewID(subTask.ID), newID)

	watch.Stop()
}

func TestChannelWatcher(t *testing.T) {
	subConn, taskConn := createServers(t)

	subscriptionClient := subapi.NewE2SubscriptionServiceClient(subConn)
	assert.NotNil(t, subscriptionClient)

	subscriptionTaskClient := subtaskapi.NewE2SubscriptionTaskServiceClient(taskConn)
	assert.NotNil(t, subscriptionTaskClient)

	watch := ChannelWatcher{
		endpointID: "1",
		tasks:      subscriptionTaskClient,
		subs:       subscriptionClient,
		channels:   server2.NewChannelManager(),
		cancel:     nil,
		mu:         sync.Mutex{},
	}

	ch := make(chan controller.ID)
	err := watch.Start(ch)
	assert.NoError(t, err)

	subTask := &subtaskapi.SubscriptionTask{
		ID:             "1",
		Revision:       0,
		SubscriptionID: "1",
		EndpointID:     "1",
		Lifecycle:      subtaskapi.Lifecycle{},
	}

	err = taskStore.Create(context.Background(), subTask)
	assert.NoError(t, err)

	subscription := &subapi.Subscription{
		ID: "1", AppID: "foo", Details: &subapi.SubscriptionDetails{E2NodeID: "bar", ServiceModel: subapi.ServiceModel{ID: "sm1"}},
	}
	_, err = subscriptionClient.AddSubscription(context.Background(), &subapi.AddSubscriptionRequest{
		Subscription: subscription,
	})
	assert.NoError(t, err)

	watch.channelCh <- nil
	newID := <-ch
	assert.Equal(t, controller.NewID(subTask.ID), newID)

	watch.Stop()
}

//subscription := &subapi.Subscription{
//ID: "1", AppID: "foo", Details: &subapi.SubscriptionDetails{E2NodeID: "bar", ServiceModel: subapi.ServiceModel{ID: "sm1"}},
//}
//_, err = subscriptionClient.AddSubscription(context.Background(), &subapi.AddSubscriptionRequest{
//	Subscription: subscription,
//})
//assert.NoError(t, err)
