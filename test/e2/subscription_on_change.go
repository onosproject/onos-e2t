// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package e2

import (
	"context"
	"math/rand"
	"testing"
	"time"

	subapi "github.com/onosproject/onos-api/go/onos/e2sub/subscription"
	e2sm_rc_pre_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v1/e2sm-rc-pre-ies"
	"google.golang.org/protobuf/proto"

	ransimtypes "github.com/onosproject/onos-api/go/onos/ransim/types"

	modelapi "github.com/onosproject/onos-api/go/onos/ransim/model"

	"github.com/onosproject/onos-e2t/test/utils"
	"github.com/onosproject/onos-ric-sdk-go/pkg/e2/indication"
	"github.com/stretchr/testify/assert"
)

// TestSubscriptionOnChange tests E2 subscription on change using ransim, SDK
func (s *TestSuite) TestSubscriptionOnChange(t *testing.T) {
	sim := utils.CreateRanSimulatorWithNameOrDie(t, "subscription-on-change")
	assert.NotNil(t, sim)
	ch := make(chan indication.Indication)
	ctx := context.Background()

	e2Client := getE2Client(t, "subscription-on-change-test")

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

	subRequest := utils.Subscription{
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

	subReq, err := subRequest.Create()
	assert.NoError(t, err)

	sub, err := e2Client.Subscribe(ctx, subReq, ch)
	assert.NoError(t, err)

	var indMessage indication.Indication
	// expects three indication messages since we have three cells for that node
	for i := 0; i < 3; i++ {
		indMessage = checkIndicationMessage(t, defaultIndicationTimeout, ch)
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
	ricIndicationHeader := e2sm_rc_pre_ies.E2SmRcPreIndicationHeader{}

	err = proto.Unmarshal(header, &ricIndicationHeader)
	assert.NoError(t, err)
	plmnID := ricIndicationHeader.GetIndicationHeaderFormat1().GetCgi().GetEUtraCgi().GetPLmnIdentity().Value
	testEci := ricIndicationHeader.GetIndicationHeaderFormat1().GetCgi().GetEUtraCgi().GetEUtracellIdentity().Value.Value
	plmnIDValue := ransimtypes.Uint24ToUint32(plmnID)
	ecgi := ransimtypes.ToECGI(ransimtypes.PlmnID(plmnIDValue), ransimtypes.GetECI(testEci))

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
	indMessage = checkIndicationMessage(t, defaultIndicationTimeout, ch)
	err = sub.Close()
	assert.NoError(t, err)
	err = sim.Uninstall()
	assert.NoError(t, err)

}
