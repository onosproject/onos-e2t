// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package e2

import (
	"context"
	"fmt"
	"github.com/onosproject/onos-api/go/onos/e2t/e2"
	"github.com/onosproject/onos-ric-sdk-go/pkg/e2/connection"
	"github.com/onosproject/onos-ric-sdk-go/pkg/e2/subscription"
	"github.com/onosproject/onos-ric-sdk-go/pkg/e2/termination"
	"math/rand"
	"testing"
	"time"

	subapi "github.com/onosproject/onos-api/go/onos/e2sub/subscription"
	ransimtypes "github.com/onosproject/onos-api/go/onos/ransim/types"

	modelapi "github.com/onosproject/onos-api/go/onos/ransim/model"

	"github.com/onosproject/onos-e2t/test/utils"
	"github.com/stretchr/testify/assert"
)

// TestSubscriptionIndicationBuffering tests E2 indication buffering
func (s *TestSuite) TestSubscriptionIndicationBuffering(t *testing.T) {
	sim := utils.CreateRanSimulatorWithNameOrDie(t, "subscription-indication-buffering")
	assert.NotNil(t, sim)
	ctx := context.Background()

	conns := connection.NewManager()

	subConn, err := conns.Connect(fmt.Sprintf("%s:%d", utils.SubscriptionServiceHost, utils.SubscriptionServicePort))
	assert.NoError(t, err)
	subClient := subscription.NewClient(subConn)

	e2tConn, err := conns.Connect(fmt.Sprintf("%s:%d", utils.E2TServiceHost, utils.E2TServicePort))
	assert.NoError(t, err)
	e2tClient := termination.NewClient(e2tConn)

	nodeClient := utils.GetRansimNodeClient(t, sim)
	assert.NotNil(t, nodeClient)
	cellClient := utils.GetRansimCellClient(t, sim)
	assert.NotNil(t, cellClient)

	// Get list of e2 nodes using RAN simulator API
	e2nodes := utils.GetNodes(t, nodeClient)
	// Delete all of the available nodes
	for _, e2node := range e2nodes {
		_, err := nodeClient.DeleteNode(ctx, &modelapi.DeleteNodeRequest{
			EnbID: e2node.EnbID,
		})
		assert.NoError(t, err)
	}
	// Get list of all available e2 nodes and make sure no node is connected
	nodeIDs, err := utils.GetNodeIDs()
	assert.NoError(t, err)
	assert.Equal(t, 0, len(nodeIDs))

	// Create an e2 node with 3 cells from list of available cells.
	cells := utils.GetCells(t, cellClient)
	assert.Greater(t, len(cells), 2)

	cell1Index := rand.Intn(len(cells))
	cell2Index := rand.Intn(len(cells))
	cell3Index := rand.Intn(len(cells))
	cell1 := cells[cell1Index]
	cell2 := cells[cell2Index]
	cell3 := cells[cell3Index]
	enbID := 157000
	createNodeRequest := &modelapi.CreateNodeRequest{
		Node: &ransimtypes.Node{
			EnbID:         ransimtypes.EnbID(enbID),
			ServiceModels: serviceModels,
			Controllers:   controllers,
			CellECGIs:     []ransimtypes.ECGI{cell1.ECGI, cell2.ECGI, cell3.ECGI},
		},
	}
	e2node, err := nodeClient.CreateNode(ctx, createNodeRequest)
	assert.NoError(t, err)
	assert.NotNil(t, e2node)

	// Waits until the connection gets established and make sure there is just one node connected
	// TODO this should be replaced with a mechanism to make sure all of the nodes are gone before asking
	// for the number of nodes
	time.Sleep(10 * time.Second)
	nodeIDs, err = utils.GetNodeIDs()
	assert.NoError(t, err)
	assert.Equal(t, 1, len(nodeIDs))
	testNodeID := nodeIDs[0]

	// Creates a subscription using RC service model
	eventTriggerBytes, err := utils.CreateRcEventTrigger()
	assert.NoError(t, err)

	subBuilder := utils.Subscription{
		NodeID:               testNodeID,
		EncodingType:         subapi.Encoding_ENCODING_PROTO,
		ActionType:           subapi.ActionType_ACTION_TYPE_REPORT,
		EventTrigger:         eventTriggerBytes,
		ServiceModelName:     utils.RcServiceModelName,
		ServiceModelVersion:  utils.RcServiceModelVersion1,
		ActionID:             100,
		SubSequentActionType: subapi.SubsequentActionType_SUBSEQUENT_ACTION_TYPE_CONTINUE,
		TimeToWait:           subapi.TimeToWait_TIME_TO_WAIT_ZERO,
	}

	subDetails, err := subBuilder.Create()
	assert.NoError(t, err)

	subscription := &subapi.Subscription{
		ID:      "test-subscription",
		AppID:   "subscription-indication-buffering-test",
		Details: &subDetails,
	}
	err = subClient.Add(context.Background(), subscription)
	assert.NoError(t, err)

	// Sleep for ten seconds to ensure indications are sent before opening a stream
	time.Sleep(10 * time.Second)

	// Open the subscription stream
	responseCh := make(chan e2.StreamResponse)
	requestCh, err := e2tClient.Stream(context.Background(), responseCh)
	assert.NoError(t, err)
	requestCh <- e2.StreamRequest{
		Header: &e2.RequestHeader{
			EncodingType: e2.EncodingType_PROTO,
			ServiceModel: &e2.ServiceModel{
				Name:    utils.RcServiceModelName,
				Version: utils.RcServiceModelVersion1,
			},
		},
		AppID:          "subscription-indication-buffering-test",
		InstanceID:     "subscription-indication-buffering-test-1",
		SubscriptionID: e2.SubscriptionID(subscription.ID),
	}

	// expects three indication messages since we have three cells for that node
	for i := 0; i < 3; i++ {
		select {
		case response, ok := <-responseCh:
			t.Log(response)
			assert.True(t, ok)
		case <-time.After(10 * time.Second):
			t.Fatal("failed to receive indication message")
		}
	}

	// Make sure that reads on the subscription channel time out. There should be no
	// indication messages available
	var gotResponse bool
	select {
	case response := <-responseCh:
		// We got an indication. This is an error, as there is no E2 node to send one
		gotResponse = true
		t.Log(response)
	case <-time.After(10 * time.Second):
		// The read timed out. This is the expected behavior.
		gotResponse = false
	}
	assert.False(t, gotResponse, "received an extraneous indication")

	err = subClient.Close()
	assert.NoError(t, err)
	err = e2tClient.Close()
	assert.NoError(t, err)
	err = sim.Uninstall()
	assert.NoError(t, err)
}
