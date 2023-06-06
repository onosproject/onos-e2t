// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2

import (
	"context"
	"time"

	topoapi "github.com/onosproject/onos-api/go/onos/topo"

	modelapi "github.com/onosproject/onos-api/go/onos/ransim/model"
	ransimtypes "github.com/onosproject/onos-api/go/onos/ransim/types"
	"github.com/onosproject/onos-e2t/test/utils"
)

const (
	numRequestedE2Nodes = 2
)

var (
	initialEnbID  = 155000
	serviceModels = []string{"kpm2", "rcpre2"}
	controllers   = []string{"e2t-1"}
)

func (s *TestSuite) TestMultiE2Nodes() {
	ctx, cancel := context.WithTimeout(s.Context(), 1*time.Minute)
	defer cancel()

	topoSdkClient, err := utils.NewTopoClient()
	s.NoError(err)
	topoEventChan := make(chan topoapi.Event)

	err = topoSdkClient.WatchE2Connections(ctx, topoEventChan)
	s.NoError(err)
	sim := s.CreateRanSimulatorWithNameOrDie("multi-e2nodes")
	s.NotNil(sim)

	nodeClient := s.GetRansimNodeClient()
	s.NotNil(nodeClient)
	defaultNumNodes := s.GetNumNodes(nodeClient)

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
		s.NoError(err)
		s.NotNil(e2node)
	}
	numNodes := s.GetNumNodes(nodeClient)
	s.Equal(numRequestedE2Nodes+defaultNumNodes, numNodes)

	expectedControlRelations := int(s.E2TReplicaCount) * numNodes
	utils.CountTopoAddedOrNoneEvent(topoEventChan, expectedControlRelations)
	e2nodes := s.GetNodes(nodeClient)
	for _, e2node := range e2nodes {
		_, err := nodeClient.DeleteNode(ctx, &modelapi.DeleteNodeRequest{
			GnbID: e2node.GnbID,
		})
		s.NoError(err)
	}

	utils.CountTopoRemovedEvent(topoEventChan, expectedControlRelations)
	numNodes = s.GetNumNodes(nodeClient)
	s.Equal(0, numNodes)
	//s.UninstallRanSimulatorOrDie(sim, "multi-e2nodes")
}
