// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2

import (
	"context"
	"math/rand"
	"testing"
	"time"

	topoapi "github.com/onosproject/onos-api/go/onos/topo"

	sdkclient "github.com/onosproject/onos-ric-sdk-go/pkg/e2/v1beta1"

	"github.com/onosproject/onos-e2t/test/e2utils"

	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	e2smrcpreies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v2/e2sm-rc-pre-v2"
	"google.golang.org/protobuf/proto"

	ransimtypes "github.com/onosproject/onos-api/go/onos/ransim/types"

	modelapi "github.com/onosproject/onos-api/go/onos/ransim/model"

	"github.com/onosproject/onos-e2t/test/utils"
	"github.com/stretchr/testify/assert"
)

// TestSubscriptionOnChange tests E2 subscription on change using ransim, SDK
func (s *TestSuite) TestSubscriptionOnChange(t *testing.T) {

	sim := utils.CreateRanSimulatorWithNameOrDie(t, s.c, "subscription-on-change")
	assert.NotNil(t, sim)

	ctx, cancel := context.WithTimeout(context.Background(), subscriptionTimeout)
	defer cancel()
	topoSdkClient, err := utils.NewTopoClient()
	assert.NoError(t, err)
	topoEventChan := make(chan topoapi.Event)
	err = topoSdkClient.WatchE2Connections(ctx, topoEventChan)
	assert.NoError(t, err)

	nodeClient := utils.GetRansimNodeClient(t, sim)
	assert.NotNil(t, nodeClient)
	cellClient := utils.GetRansimCellClient(t, sim)
	assert.NotNil(t, cellClient)

	defaultNumNodes := utils.GetNumNodes(t, nodeClient)
	utils.CountTopoAddedOrNoneEvent(topoEventChan, defaultNumNodes)

	// Get list of e2 nodes using RAN simulator API
	e2nodes := utils.GetNodes(t, nodeClient)
	numNodes := utils.GetNumNodes(t, nodeClient)

	// Delete all of the available nodes
	for _, e2node := range e2nodes {
		_, err := nodeClient.DeleteNode(ctx, &modelapi.DeleteNodeRequest{
			GnbID: e2node.GnbID,
		})
		assert.NoError(t, err)
	}

	utils.CountTopoRemovedEvent(topoEventChan, numNodes)

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

	numNodes = utils.GetNumNodes(t, nodeClient)
	utils.CountTopoAddedOrNoneEvent(topoEventChan, numNodes)

	testNodeID := utils.GetTestNodeID(t)

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

	subRequest := utils.Subscription{
		NodeID:              string(testNodeID),
		Actions:             actions,
		EventTrigger:        eventTriggerBytes,
		ServiceModelName:    utils.RcServiceModelName,
		ServiceModelVersion: utils.Version2,
	}

	subSpec, err := subRequest.Create()
	assert.NoError(t, err)

	subName := "TestSubscriptionOnChange"

	sdkClient := utils.GetE2Client(t, utils.RcServiceModelName, utils.Version2, sdkclient.ProtoEncoding)
	node := sdkClient.Node(sdkclient.NodeID(testNodeID))
	ch := make(chan e2api.Indication)
	_, err = node.Subscribe(ctx, subName, subSpec, ch)
	assert.NoError(t, err)

	var indMessage e2api.Indication
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

	header := indMessage.Header
	ricIndicationHeader := e2smrcpreies.E2SmRcPreIndicationHeader{}

	err = proto.Unmarshal(header, &ricIndicationHeader)
	assert.NoError(t, err)
	plmnID := ricIndicationHeader.GetIndicationHeaderFormat1().GetCgi().GetNrCgi().GetPLmnIdentity().Value
	nrcid := utils.BitStringToUint64(ricIndicationHeader.GetIndicationHeaderFormat1().GetCgi().GetNrCgi().GetNRcellIdentity().Value.Value, 36)
	plmnIDValue := ransimtypes.Uint24ToUint32(plmnID)
	ncgi := ransimtypes.ToNCGI(ransimtypes.PlmnID(plmnIDValue), ransimtypes.NCI(nrcid))

	testCell, err := cellClient.GetCell(ctx, &modelapi.GetCellRequest{
		NCGI: ncgi,
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

	err = node.Unsubscribe(ctx, subName)
	assert.NoError(t, err)

	e2utils.CheckForEmptySubscriptionList(t)
	utils.UninstallRanSimulatorOrDie(t, sim)
}
