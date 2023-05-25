// SPDX-FileCopyrightText: 2022-present Intel Corporation
// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2

import (
	"context"
	"math/rand"
	"time"

	topoapi "github.com/onosproject/onos-api/go/onos/topo"

	"github.com/onosproject/onos-e2t/test/e2utils"

	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	e2smrcpreies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre_go/v2/e2sm-rc-pre-v2-go"
	"google.golang.org/protobuf/proto"

	ransimtypes "github.com/onosproject/onos-api/go/onos/ransim/types"

	modelapi "github.com/onosproject/onos-api/go/onos/ransim/model"

	"github.com/onosproject/onos-e2t/test/utils"
)

// TestSubscriptionOnChange tests E2 subscription on change using ransim, SDK
func (s *TestSuite) TestSubscriptionOnChange() {

	sim := s.CreateRanSimulatorWithNameOrDie("subscription-on-change")
	s.NotNil(sim)

	ctx, cancel := context.WithTimeout(s.Context(), subscriptionTimeout)
	defer cancel()
	topoSdkClient, err := utils.NewTopoClient()
	s.NoError(err)
	topoEventChan := make(chan topoapi.Event)
	err = topoSdkClient.WatchE2Connections(ctx, topoEventChan)
	s.NoError(err)

	nodeClient := s.GetRansimNodeClient()
	s.NotNil(nodeClient)
	cellClient := s.GetRansimCellClient()
	s.NotNil(cellClient)

	defaultNumNodes := s.GetNumNodes(nodeClient)
	utils.CountTopoAddedOrNoneEvent(topoEventChan, defaultNumNodes)

	// Get list of e2 nodes using RAN simulator API
	e2nodes := s.GetNodes(nodeClient)
	numNodes := s.GetNumNodes(nodeClient)

	// Delete all the available nodes
	for _, e2node := range e2nodes {
		_, err := nodeClient.DeleteNode(ctx, &modelapi.DeleteNodeRequest{
			GnbID: e2node.GnbID,
		})
		s.NoError(err)
	}

	utils.CountTopoRemovedEvent(topoEventChan, numNodes)

	// Create an e2 node with 3 cells from list of available cells.
	cells := s.GetCells(cellClient)
	s.Greater(len(cells), 2)

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
	s.NoError(err)
	s.NotNil(e2node)

	numNodes = s.GetNumNodes(nodeClient)
	utils.CountTopoAddedOrNoneEvent(topoEventChan, numNodes)

	testNodeID := utils.GetTestNodeID(s.T())

	// Creates a subscription using RC service model
	subName := "TestSubscriptionOnChange"
	rcPreSub := e2utils.RCPreSub{
		Sub: e2utils.Sub{
			Name:   subName,
			NodeID: testNodeID,
		},
	}
	s.NoError(rcPreSub.UseDefaultReportAction())
	rcPreSub.SubscribeOrFail(s.Context(), s.T())

	var indMessage e2api.Indication
	// expects three indication messages since we have three cells for that node
	for i := 0; i < 3; i++ {
		indMessage = e2utils.CheckIndicationMessage(s.T(), e2utils.DefaultIndicationTimeout, rcPreSub.Sub.Ch)

	}

	// Make sure that reads on the subscription channel time out. There should be no
	// indication messages available
	var gotIndication bool
	select {
	case indicationMsg := <-rcPreSub.Sub.Ch:
		// We got an indication. This is an error, as there is no E2 node to send one
		gotIndication = true
		s.T().Log(indicationMsg)

	case <-time.After(10 * time.Second):
		// The read timed out. This is the expected behavior.
		gotIndication = false
	}
	s.False(gotIndication, "received an extraneous indication")

	header := indMessage.Header
	ricIndicationHeader := e2smrcpreies.E2SmRcPreIndicationHeader{}

	err = proto.Unmarshal(header, &ricIndicationHeader)
	s.NoError(err)

	plmnID := ricIndicationHeader.GetIndicationHeaderFormat1().GetCgi().GetNrCgi().GetPLmnIdentity().Value
	nrcid := utils.BitStringToUint64(ricIndicationHeader.GetIndicationHeaderFormat1().GetCgi().GetNrCgi().GetNRcellIdentity().Value.Value, 36)
	plmnIDValue := ransimtypes.Uint24ToUint32(plmnID)
	ncgi := ransimtypes.ToNCGI(ransimtypes.PlmnID(plmnIDValue), ransimtypes.NCI(nrcid))

	testCell, err := cellClient.GetCell(ctx, &modelapi.GetCellRequest{
		NCGI: ncgi,
	})
	s.NoError(err)
	neighborsList := testCell.GetCell().Neighbors
	// Update the list of neighbors
	neighborsList = append(neighborsList[:1], neighborsList[2:]...)
	testCell.Cell.Neighbors = neighborsList
	_, err = cellClient.UpdateCell(ctx, &modelapi.UpdateCellRequest{
		Cell: testCell.Cell,
	})
	s.NoError(err)
	// Expect to receive indication message on neighbor list change
	indMessage = e2utils.CheckIndicationMessage(s.T(), e2utils.DefaultIndicationTimeout, rcPreSub.Sub.Ch)

	rcPreSub.Sub.UnsubscribeOrFail(ctx, s.T())

	e2utils.CheckForEmptySubscriptionList(s.T())
	s.UninstallRanSimulatorOrDie(sim, "subscription-on-change")
}
