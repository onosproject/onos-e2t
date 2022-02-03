// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2

import (
	"context"
	"testing"
	"time"

	topoapi "github.com/onosproject/onos-api/go/onos/topo"

	modelapi "github.com/onosproject/onos-api/go/onos/ransim/model"
	ransimtypes "github.com/onosproject/onos-api/go/onos/ransim/types"
	"github.com/onosproject/onos-e2t/test/utils"
	"github.com/stretchr/testify/assert"
)

const (
	numRequestedE2Nodes = 50
)

var (
	initialEnbID  = 155000
	serviceModels = []string{"kpm2", "rcpre2"}
	controllers   = []string{"e2t-1"}
)

func (s *TestSuite) TestMultiE2Nodes(t *testing.T) {
	sim := utils.CreateRanSimulatorWithNameOrDie(t, s.c, "multi-e2nodes")
	assert.NotNil(t, sim)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	topoSdkClient, err := utils.NewTopoClient()
	assert.NoError(t, err)
	topoEventChan := make(chan topoapi.Event)
	err = topoSdkClient.WatchE2Connections(ctx, topoEventChan)
	assert.NoError(t, err)

	nodeClient := utils.GetRansimNodeClient(t, sim)
	assert.NotNil(t, nodeClient)

	defaultNumNodes := utils.GetNumNodes(t, nodeClient)

	for i := 0; i < numRequestedE2Nodes; i++ {
		enbID := i + initialEnbID
		createNodeRequest := &modelapi.CreateNodeRequest{
			Node: &ransimtypes.Node{
				GnbID:         ransimtypes.GnbID(enbID),
				ServiceModels: serviceModels,
				Controllers:   controllers,
				CellNCGIs:     []ransimtypes.NCGI{},
			},
		}
		e2node, err := nodeClient.CreateNode(ctx, createNodeRequest)
		assert.NoError(t, err)
		assert.NotNil(t, e2node)
	}
	numNodes := utils.GetNumNodes(t, nodeClient)
	assert.Equal(t, numRequestedE2Nodes+defaultNumNodes, numNodes)

	utils.CountTopoAddedOrNoneEvent(topoEventChan, numNodes)
	e2nodes := utils.GetNodes(t, nodeClient)
	for _, e2node := range e2nodes {
		_, err = nodeClient.DeleteNode(ctx, &modelapi.DeleteNodeRequest{
			GnbID: e2node.GnbID,
		})
		assert.NoError(t, err)
	}

	utils.CountTopoRemovedEvent(topoEventChan, numNodes)

	numNodes = utils.GetNumNodes(t, nodeClient)
	assert.Equal(t, 0, numNodes)
	utils.UninstallRanSimulatorOrDie(t, sim)
}
