// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package e2

import (
	"context"
	e2v1beta1api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	"github.com/onosproject/onos-e2t/test/e2utils"
	e2v1beta1client "github.com/onosproject/onos-ric-sdk-go/pkg/e2/v1beta1"
	"math/rand"
	"testing"
	"time"

	subapi "github.com/onosproject/onos-api/go/onos/e2sub/subscription"
	e2smrcpreies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v2/e2sm-rc-pre-v2"
	"google.golang.org/protobuf/proto"

	ransimtypes "github.com/onosproject/onos-api/go/onos/ransim/types"

	modelapi "github.com/onosproject/onos-api/go/onos/ransim/model"

	"github.com/onosproject/onos-e2t/test/utils"
	"github.com/onosproject/onos-ric-sdk-go/pkg/e2/indication"
	"github.com/stretchr/testify/assert"
)

// TestSubscriptionOnChangeV1beta1 tests E2 subscription on change using ransim, SDK
func (s *TestSuite) TestSubscriptionOnChangeV1beta1(t *testing.T) {
	sim := utils.CreateRanSimulatorWithNameOrDie(t, s.c, "subscription-on-change")
	assert.NotNil(t, sim)
	ch := make(chan indication.Indication)
	ctx := context.Background()

	e2Client := utils.GetE2V1beta1Client(t,
		"subscription-on-change-v1beta1-test",
		"subscription-on-change-v1beta1-test-1",
		utils.RcServiceModelName,
		utils.Version2)

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
	connections, err = utils.GetAllE2Connections(t)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(connections))
	nodeIDs, err := utils.GetNodeIDs(t)
	assert.NoError(t, err)
	testNodeID := nodeIDs[0]

	// Creates a subscription using RC service model
	eventTriggerBytes, err := utils.CreateRcEventTrigger()
	assert.NoError(t, err)
	var actions []subapi.Action
	action := subapi.Action{
		ID:   100,
		Type: subapi.ActionType_ACTION_TYPE_REPORT,
		SubsequentAction: &subapi.SubsequentAction{
			Type:       subapi.SubsequentActionType_SUBSEQUENT_ACTION_TYPE_CONTINUE,
			TimeToWait: subapi.TimeToWait_TIME_TO_WAIT_ZERO,
		},
	}
	actions = append(actions, action)

	subscription := &e2v1beta1api.Subscription{
		ID: "subscription-on-change-v1beta1",
		EventTrigger: e2v1beta1api.EventTrigger{
			Payload: eventTriggerBytes,
		},
		Actions: []e2v1beta1api.Action{
			{
				ID:   100,
				Type: e2v1beta1api.ActionType_ACTION_TYPE_REPORT,
				SubsequentAction: &e2v1beta1api.SubsequentAction{
					Type:       e2v1beta1api.SubsequentActionType_SUBSEQUENT_ACTION_TYPE_CONTINUE,
					TimeToWait: e2v1beta1api.TimeToWait_TIME_TO_WAIT_ZERO,
				},
			},
		},
	}

	indCh := make(chan e2v1beta1api.Indication)
	errCh := make(chan error)
	err = e2Client.Node(e2v1beta1client.NodeID(testNodeID)).Subscribe(ctx, subscription, indCh, errCh)
	assert.NoError(t, err)

	var indMessage indication.Indication
	// expects three indication messages since we have three cells for that node
	for i := 0; i < 3; i++ {
		indMessage = e2utils.CheckIndicationMessage(t, e2utils.DefaultIndicationTimeout, ch)
	}

	// Make sure that reads on the subscription channel time out. There should be no
	// indication messages available
	var gotIndication bool
	select {
	case indicationMsg := <-ch:
		// We got an indication. This is an error, as there is no E2 node to send one
		gotIndication = true
		t.Log(indicationMsg)

	case <-time.After(10 * time.Second):
		// The read timed out. This is the expected behavior.
		gotIndication = false
	}
	assert.False(t, gotIndication, "received an extraneous indication")

	header := indMessage.Payload.Header
	ricIndicationHeader := e2smrcpreies.E2SmRcPreIndicationHeader{}

	err = proto.Unmarshal(header, &ricIndicationHeader)
	assert.NoError(t, err)
	plmnID := ricIndicationHeader.GetIndicationHeaderFormat1().GetCgi().GetNrCgi().GetPLmnIdentity().Value
	nrcid := ricIndicationHeader.GetIndicationHeaderFormat1().GetCgi().GetNrCgi().GetNRcellIdentity().Value.Value
	plmnIDValue := ransimtypes.Uint24ToUint32(plmnID)
	ecgi := ransimtypes.ToECGI(ransimtypes.PlmnID(plmnIDValue), ransimtypes.GetECI(nrcid))

	testCell, err := cellClient.GetCell(ctx, &modelapi.GetCellRequest{
		ECGI: ecgi,
	})
	assert.NoError(t, err)
	neighborsList := testCell.GetCell().Neighbors
	// Update the list of neighbors
	neighborsList = append(neighborsList[:1], neighborsList[2:]...)
	testCell.Cell.Neighbors = neighborsList
	_, err = cellClient.UpdateCell(ctx, &modelapi.UpdateCellRequest{
		Cell: testCell.Cell,
	})
	assert.NoError(t, err)
	// Expect to receive indication message on neighbor list change
	indMessage = e2utils.CheckIndicationMessage(t, e2utils.DefaultIndicationTimeout, ch)
	err = sim.Uninstall()
	assert.NoError(t, err)
}
