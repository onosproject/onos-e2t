// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package e2

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"

	sdkclient "github.com/onosproject/onos-ric-sdk-go/pkg/e2/v1beta1"

	"github.com/onosproject/onos-ric-sdk-go/pkg/e2/connection"
	"github.com/onosproject/onos-ric-sdk-go/pkg/e2/termination"

	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	ransimtypes "github.com/onosproject/onos-api/go/onos/ransim/types"

	modelapi "github.com/onosproject/onos-api/go/onos/ransim/model"

	"github.com/onosproject/onos-e2t/test/utils"
	"github.com/stretchr/testify/assert"
)

// TestSubscriptionIndicationBuffering tests E2 indication buffering
func (s *TestSuite) TestSubscriptionIndicationBuffering(t *testing.T) {
	sim := utils.CreateRanSimulatorWithNameOrDie(t, s.c, "subscription-indication-buffering")
	assert.NotNil(t, sim)
	ctx, cancel := context.WithCancel(context.Background())

	conns := connection.NewManager()

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
			GnbID: e2node.GnbID,
		})
		assert.NoError(t, err)
	}
	// Get list of all available e2 nodes and make sure no node is connected
	connections, err := utils.GetAllE2Connections(t)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(connections))

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
			GnbID:         ransimtypes.GnbID(enbID),
			ServiceModels: serviceModels,
			Controllers:   controllers,
			CellNCGIs:     []ransimtypes.NCGI{cell1.NCGI, cell2.NCGI, cell3.NCGI},
		},
	}
	e2node, err := nodeClient.CreateNode(ctx, createNodeRequest)
	assert.NoError(t, err)
	assert.NotNil(t, e2node)

	// Waits until the connection gets established and make sure there is just one node connected
	// TODO this should be replaced with a mechanism to make sure all of the nodes are gone before asking
	// for the number of nodes
	time.Sleep(10 * time.Second)
	connections, err = utils.GetAllE2Connections(t)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(connections))
	nodeIDs, err := utils.GetNodeIDs(t)
	assert.NoError(t, err)
	testNodeID := nodeIDs[0]

	// Creates a subscription using RC service model
	eventTriggerBytes, err := utils.CreateRcEventTrigger()
	assert.NoError(t, err)
	var actions []e2api.Action
	action := e2api.Action{
		ID:   100,
		Type: e2api.ActionType_ACTION_TYPE_REPORT,
		SubsequentAction: &e2api.SubsequentAction{
			Type:       e2api.SubsequentActionType_SUBSEQUENT_ACTION_TYPE_CONTINUE,
			TimeToWait: e2api.TimeToWait_TIME_TO_WAIT_ZERO,
		},
	}
	actions = append(actions, action)

	subBuilder := utils.Subscription2{
		NodeID:              string(testNodeID),
		Actions:             actions,
		EventTrigger:        eventTriggerBytes,
		ServiceModelName:    utils.RcServiceModelName,
		ServiceModelVersion: utils.Version2,
	}

	subSpec, err := subBuilder.Create()
	assert.NoError(t, err)

	// Sleep for ten seconds to ensure indications are sent before opening a stream
	time.Sleep(10 * time.Second)

	sdkClient := utils.GetE2Client2(t, utils.RcServiceModelName, utils.Version2, sdkclient.ProtoEncoding)
	node := sdkClient.Node(sdkclient.NodeID(testNodeID))
	responseCh := make(chan e2api.Indication)
	_, err = node.Subscribe(ctx, "test-subscription", subSpec, responseCh)
	assert.NoError(t, err)

	// expects three indication messages since we have three cells for that node
	for i := 0; i < 3; i++ {
		select {
		case response, ok := <-responseCh:
			t.Log(response)
			assert.True(t, ok)
		case <-time.After(10 * time.Second):
			assert.Equal(t, false, "failed to receive indication message")
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

	err = e2tClient.Close()
	assert.NoError(t, err)
	err = sim.Uninstall()
	assert.NoError(t, err)

	cancel()
}
