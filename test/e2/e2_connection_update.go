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
	"time"
)

// TestE2TConnectionUpdate checks that the control relations are correct
func (s *TestSuite) TestE2TConnectionUpdate(t *testing.T) {
	numberOfE2TNodes := int(s.E2TReplicaCount)
	numberOfSimulatorNodes := 1
	numberOfE2Nodes := numberOfSimulatorNodes * 2
	//numberOfControlRelationsPerNode := 2 // Will change
	//numberOfControlRelationships := numberOfE2TNodes * numberOfControlRelationsPerNode
	numberOfControlRelationships := numberOfE2Nodes * 1

	ctx, cancel := e2utils.GetCtx()
	defer cancel()
	topoSdkClient, err := utils.NewTopoClient()
	assert.NoError(t, err)
	topoNodeEventChan := make(chan topoapi.Event)
	err = topoSdkClient.WatchE2Nodes(ctx, topoNodeEventChan)
	assert.NoError(t, err)

	// create a simulator
	sim := utils.CreateRanSimulatorWithNameOrDie(t, s.c, "e2t-connection-update")
	assert.NotNil(t, sim)

	time.Sleep(15 * time.Second)

	// Check that the e2T Nodes are correct in topology
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
		assert.Equal(t, 1, relCount) // will change
	}

	// tear down the simulator
	assert.NoError(t, sim.Uninstall())
}
