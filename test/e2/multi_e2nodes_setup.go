// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package e2

import (
	"context"
	"math/rand"
	"testing"

	topoapi "github.com/onosproject/onos-api/go/onos/topo"
	"github.com/onosproject/onos-ric-sdk-go/pkg/topo/options"

	"github.com/onosproject/onos-e2t/test/utils"

	"github.com/stretchr/testify/assert"

	modelapi "github.com/onosproject/onos-api/go/onos/ransim/model"
	ransimtypes "github.com/onosproject/onos-api/go/onos/ransim/types"
)

const (
	numRequestedE2Nodes       = 50
	numRequestedCells         = 150
	expectedDefaultNumE2Nodes = 2 // Based on default helm chart
)

var (
	initialEnbID  = 155000
	serviceModels = []string{"kpm", "rc"}
	controllers   = []string{"e2t-1", "e2t-2"}
)

func (s *TestSuite) TestMultiE2Nodes(t *testing.T) {
	topoClient, err := utils.GetTopoClient()
	assert.NoError(t, err)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ch := make(chan topoapi.Event)
	err = topoClient.Watch(ctx, ch, options.WithWatchFilters(utils.GetControlRelationFilter()))
	assert.NoError(t, err)

	sim := utils.CreateRanSimulatorWithNameOrDie(t, s.c, "multi-e2nodes")
	assert.NotNil(t, sim)

	nodeClient := utils.GetRansimNodeClient(t, sim)
	assert.NotNil(t, nodeClient)
	cellClient := utils.GetRansimCellClient(t, sim)
	assert.NotNil(t, cellClient)
	defaultNumCells := utils.GetNumCells(t, cellClient)

	for i := 1; i < numRequestedCells+1; i++ {
		_, err := cellClient.CreateCell(ctx, &modelapi.CreateCellRequest{
			Cell: &ransimtypes.Cell{
				ECGI:      ransimtypes.ECGI(i),
				Color:     "red",
				Neighbors: []ransimtypes.ECGI{ransimtypes.ECGI(i + 1)},
				Sector: &ransimtypes.Sector{Arc: 180,
					Centroid: &ransimtypes.Point{
						Lat: 56.0,
						Lng: 78.9,
					}},
				Location:      &ransimtypes.Point{Lat: 42.0, Lng: 54.23},
				MaxUEs:        12,
				EventA3Params: &ransimtypes.EventA3Params{},
			},
		})
		assert.NoError(t, err)
	}

	numCells := utils.GetNumCells(t, cellClient)
	assert.Equal(t, numRequestedCells, numCells-defaultNumCells)

	for i := 0; i < expectedDefaultNumE2Nodes; i++ {
		<-ch
	}

	defaultNumNodes := utils.GetNumNodes(t, nodeClient)
	assert.Equal(t, defaultNumNodes, expectedDefaultNumE2Nodes)
	connections, err := utils.GetTopoObjects(utils.GetControlRelationFilter())
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
				EnbID:         ransimtypes.EnbID(enbID),
				ServiceModels: serviceModels,
				Controllers:   controllers,
				CellECGIs:     []ransimtypes.ECGI{cell1.ECGI, cell2.ECGI, cell3.ECGI},
			},
		}
		e2node, err := nodeClient.CreateNode(ctx, createNodeRequest)
		assert.NoError(t, err)
		assert.NotNil(t, e2node)
	}
	for i := 0; i < numRequestedE2Nodes; i++ {
		<-ch
	}

	numNodes := utils.GetNumNodes(t, nodeClient)
	connections, err = utils.GetTopoObjects(utils.GetControlRelationFilter())
	assert.NoError(t, err)
	assert.Equal(t, numRequestedE2Nodes, numNodes-defaultNumNodes)
	assert.Equal(t, len(connections), numNodes)
	e2nodes := utils.GetNodes(t, nodeClient)
	for _, e2node := range e2nodes {
		_, err = nodeClient.DeleteNode(ctx, &modelapi.DeleteNodeRequest{
			EnbID: e2node.EnbID,
		})
		assert.NoError(t, err)
	}
	for i := 0; i < numRequestedE2Nodes; i++ {
		<-ch
	}
	numNodes = utils.GetNumNodes(t, nodeClient)
	connections, err = utils.GetTopoObjects(utils.GetControlRelationFilter())
	assert.NoError(t, err)
	assert.Equal(t, 0, numNodes)
	assert.Equal(t, 0, len(connections))
	err = sim.Uninstall()
	assert.NoError(t, err)

}
