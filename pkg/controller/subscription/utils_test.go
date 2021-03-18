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
	substore "github.com/onosproject/onos-e2sub/pkg/store/subscription"
	taskstore "github.com/onosproject/onos-e2sub/pkg/store/task"

	"github.com/onosproject/onos-lib-go/pkg/northbound"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"net"
	"testing"
)

const E2NodeID = "e2node"

var subTask = &subtaskapi.SubscriptionTask{
	ID:             "1",
	Revision:       0,
	SubscriptionID: "1",
	EndpointID:     E2NodeID,
	Lifecycle:      subtaskapi.Lifecycle{},
}

type serverScaffolding struct {
	lis       *bufconn.Listener
	server    *grpc.Server
	taskStore taskstore.Store
	subStore  substore.Store
}

var scaffold *serverScaffolding

func createServerScaffolding(t *testing.T) {
	taskStore, err := taskstore.NewLocalStore()
	assert.NoError(t, err)

	subStore, err := substore.NewLocalStore()
	assert.NoError(t, err)

	scaffold = &serverScaffolding{
		taskStore: taskStore,
		subStore:  subStore,
	}
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return scaffold.lis.Dial()
}

func newTestSubService(t *testing.T) northbound.Service {
	return e2sub.NewService(scaffold.subStore)
}

func newTestTaskService(t *testing.T) northbound.Service {
	return e2task.NewService(scaffold.taskStore)
}

func createSubServerConnection(t *testing.T) *grpc.ClientConn {
	scaffold.lis = bufconn.Listen(1024 * 1024)
	s := newTestSubService(t)
	assert.NotNil(t, s)

	s.Register(scaffold.server)

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

	s.Register(scaffold.server)

	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	return conn
}

func serve(t *testing.T) {
	go func() {
		if err := scaffold.server.Serve(scaffold.lis); err != nil {
			assert.NoError(t, err, "Server exited with error: %v", err)
		}
	}()
}

func createServers(t *testing.T) (*grpc.ClientConn, *grpc.ClientConn) {
	scaffold.server = grpc.NewServer()
	subConn := createSubServerConnection(t)
	taskConn := createTaskServerConnection(t)
	serve(t)
	return subConn, taskConn
}

func createClients(t *testing.T) (subapi.E2SubscriptionServiceClient, subtaskapi.E2SubscriptionTaskServiceClient) {
	subConn, taskConn := createServers(t)

	subscriptionClient := subapi.NewE2SubscriptionServiceClient(subConn)
	assert.NotNil(t, subscriptionClient)

	subscriptionTaskClient := subtaskapi.NewE2SubscriptionTaskServiceClient(taskConn)
	assert.NotNil(t, subscriptionTaskClient)
	return subscriptionClient, subscriptionTaskClient
}
