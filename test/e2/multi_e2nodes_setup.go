// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package e2

import (
	"context"
	modelapi "github.com/onosproject/onos-api/go/onos/ransim/model"
	ransimtypes "github.com/onosproject/onos-api/go/onos/ransim/types"
	"github.com/onosproject/onos-e2t/test/utils"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
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

func (s *TestSuite) TestMultiE2Nodes(t *testing.T) {
	sim := utils.CreateRanSimulatorWithNameOrDie(t, s.c, "multi-e2nodes")
	assert.NotNil(t, sim)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	nodeClient := utils.GetRansimNodeClient(t, sim)
	assert.NotNil(t, nodeClient)
	cellClient := utils.GetRansimCellClient(t, sim)
	assert.NotNil(t, cellClient)
	defaultNumCells := utils.GetNumCells(t, cellClient)
	for i := 1; i < numRequestedCells+1; i++ {
		_, err := cellClient.CreateCell(ctx, &modelapi.CreateCellRequest{
			Cell: &ransimtypes.Cell{
				NCGI:      ransimtypes.NCGI(i),
				Color:     "red",
				Neighbors: []ransimtypes.NCGI{ransimtypes.NCGI(i + 1)},
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
		})
		assert.NoError(t, err)
	}
	numCells := utils.GetNumCells(t, cellClient)
	assert.Equal(t, numRequestedCells, numCells-defaultNumCells)
	// TODO this should be replaced with a mechanism to make sure all of the nodes are connected before asking
	time.Sleep(20 * time.Second)
	defaultNumNodes := utils.GetNumNodes(t, nodeClient)
	connections, err := utils.GetAllE2Connections(t)
	assert.Equal(t, len(connections), defaultNumNodes)
	assert.NoError(t, err)
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
	// Wait for a few seconds to make sure all of the connections are established
	// TODO this should be replaced with a mechanism to make sure all of the nodes are gone before asking
	// for the number of nodes
	time.Sleep(20 * time.Second)
	numNodes := utils.GetNumNodes(t, nodeClient)
	connections, err = utils.GetAllE2Connections(t)
	assert.NoError(t, err)
	assert.Equal(t, numRequestedE2Nodes, numNodes-defaultNumNodes)
	assert.Equal(t, len(connections), numNodes)
	e2nodes := utils.GetNodes(t, nodeClient)
	for _, e2node := range e2nodes {
		_, err = nodeClient.DeleteNode(ctx, &modelapi.DeleteNodeRequest{
			GnbID: e2node.GnbID,
		})
		assert.NoError(t, err)
	}
	time.Sleep(10 * time.Second)
	numNodes = utils.GetNumNodes(t, nodeClient)
	connections, err = utils.GetAllE2Connections(t)
	assert.NoError(t, err)
	assert.Equal(t, 0, numNodes)
	assert.Equal(t, 0, len(connections))
	err = sim.Uninstall()
	assert.NoError(t, err)
}
