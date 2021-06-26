// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package e2

import (
	"context"
	"math/rand"
	"testing"

	topoapi "github.com/onosproject/onos-api/go/onos/topo"

	modelapi "github.com/onosproject/onos-api/go/onos/ransim/model"
	ransimtypes "github.com/onosproject/onos-api/go/onos/ransim/types"
	"github.com/onosproject/onos-e2t/test/utils"
	"github.com/stretchr/testify/assert"
)

const (
	numRequestedE2Nodes = 50
	numRequestedCells   = 150
)

var (
	initialEnbID  = 155000
	serviceModels = []string{"kpm2", "rcpre2"}
	controllers   = []string{"e2t-1", "e2t-2"}
)

func createCellRequest(index int) *modelapi.CreateCellRequest {
	return &modelapi.CreateCellRequest{
		Cell: &ransimtypes.Cell{
			NCGI:      ransimtypes.NCGI(index),
			Color:     "red",
			Neighbors: []ransimtypes.NCGI{ransimtypes.NCGI(index + 1)},
			Sector: &ransimtypes.Sector{Arc: 180,
				Centroid: &ransimtypes.Point{
					Lat: 56.0,
					Lng: 78.9,
				}},
			Location:  &ransimtypes.Point{Lat: 42.0, Lng: 54.23},
			MaxUEs:    12,
			TxPowerdB: 10,
			MeasurementParams: &ransimtypes.MeasurementParams{
				EventA3Params: &ransimtypes.EventA3Params{},
			},
		},
	}
}

func (s *TestSuite) TestMultiE2Nodes(t *testing.T) {
	sim := utils.CreateRanSimulatorWithNameOrDie(t, s.c, "multi-e2nodes")
	assert.NotNil(t, sim)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	nodeClient := utils.GetRansimNodeClient(t, sim)
	assert.NotNil(t, nodeClient)
	cellClient := utils.GetRansimCellClient(t, sim)
	assert.NotNil(t, cellClient)

	topoSdkClient, err := utils.NewTopoClient()
	assert.NoError(t, err)

	defaultNumCells := utils.GetNumCells(t, cellClient)
	for i := 1; i < numRequestedCells+1; i++ {
		_, err := cellClient.CreateCell(ctx, createCellRequest(i))
		assert.NoError(t, err)
	}
	numCells := utils.GetNumCells(t, cellClient)
	assert.Equal(t, numRequestedCells, numCells-defaultNumCells)

	topoEventChan := make(chan topoapi.Event)
	err = topoSdkClient.WatchE2Connections(ctx, topoEventChan)
	assert.NoError(t, err)
	defaultNumNodes := utils.GetNumNodes(t, nodeClient)
	for i := 0; i < defaultNumNodes; i++ {
		topoEvent := <-topoEventChan
		if topoEvent.Type.String() != topoapi.EventType_NONE.String() &&
			topoEvent.Type.String() != topoapi.EventType_ADDED.String() {
			assert.Fail(t, "topo event type does not match", topoEvent.Type)
		}
	}

	cells := utils.GetCells(t, cellClient)
	for i := 0; i < numRequestedE2Nodes; i++ {
		cell1Index := rand.Intn(len(cells))
		cell2Index := rand.Intn(len(cells))
		cell3Index := rand.Intn(len(cells))
		cell1 := cells[cell1Index]
		cell2 := cells[cell2Index]
		cell3 := cells[cell3Index]
		enbID := i + initialEnbID
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
	}
	numNodes := utils.GetNumNodes(t, nodeClient)

	for i := 0; i < numNodes-defaultNumNodes; i++ {
		topoEvent := <-topoEventChan
		assert.Equal(t, topoEvent.Type.String(), topoapi.EventType_ADDED.String())
	}

	e2nodes := utils.GetNodes(t, nodeClient)
	numNodes = utils.GetNumNodes(t, nodeClient)
	for _, e2node := range e2nodes {
		_, err = nodeClient.DeleteNode(ctx, &modelapi.DeleteNodeRequest{
			GnbID: e2node.GnbID,
		})
		assert.NoError(t, err)
	}

	for i := 0; i < numNodes; i++ {
		topoEvent := <-topoEventChan
		assert.Equal(t, topoEvent.Type.String(), topoapi.EventType_REMOVED.String())
	}

	numNodes = utils.GetNumNodes(t, nodeClient)
	assert.Equal(t, 0, numNodes)
	err = sim.Uninstall()
	assert.NoError(t, err)
}
