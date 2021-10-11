// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package e2

import (
	topoapi "github.com/onosproject/onos-api/go/onos/topo"
	"github.com/onosproject/onos-e2t/test/e2utils"
	"github.com/onosproject/onos-e2t/test/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestE2TConnectionUpdate checks that the control relations are correct
func (s *TestSuite) TestE2TConnectionUpdate(t *testing.T) {
	numberOfE2TNodes := int(s.E2TReplicaCount)
	numberOfE2Nodes := 2
	numberOfControlRelationships := numberOfE2TNodes * numberOfE2Nodes

	ctx, cancel := e2utils.GetCtx()
	defer cancel()
	topoSdkClient, err := utils.NewTopoClient()
	assert.NoError(t, err)

	// create a simulator
	sim := utils.CreateRanSimulatorWithNameOrDie(t, s.c, "e2t-connection-update")
	assert.NotNil(t, sim)

	// Wait for the E2 nodes to connect
	topoNodeEventChan := make(chan topoapi.Event)
	err = topoSdkClient.WatchE2Nodes(ctx, topoNodeEventChan)
	assert.NoError(t, err)
	utils.CountTopoAddedOrNoneEvent(topoNodeEventChan, numberOfE2Nodes)

	// Check that the E2T Nodes are correct in topology
	e2TNodes, err := topoSdkClient.E2TNodes(ctx)
	assert.NoError(t, err)
	assert.Len(t, e2TNodes, numberOfE2TNodes)

	// Check that the E2 Nodes are correct in topology
	e2Nodes, err := topoSdkClient.E2Nodes(ctx)
	assert.NoError(t, err)
	assert.Len(t, e2TNodes, numberOfE2Nodes)

	// Check that there are the correct number of relations
	relations, err := topoSdkClient.GetControlRelationsForTarget()
	assert.NoError(t, err)
	assert.Len(t, relations, numberOfControlRelationships)

	// Check that each E2 Node has the correct relations
	for _, e2Node := range e2Nodes {
		relCount := 0
		for _, o := range relations {
			rel := o.GetRelation()
			if rel.TgtEntityID == e2Node.ID {
				relCount++
			}
		}
		assert.Equal(t, numberOfE2Nodes, relCount)
	}

	// tear down the simulator
	assert.NoError(t, sim.Uninstall())
}
