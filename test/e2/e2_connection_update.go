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
	topoNodeEventChan := make(chan topoapi.Event)
	err = topoSdkClient.WatchE2Nodes(ctx, topoNodeEventChan)

	// create a simulator
	sim := utils.CreateRanSimulatorWithNameOrDie(t, s.c, "e2t-connection-update")
	assert.NotNil(t, sim)

	// Wait for the E2 nodes to connect
	assert.NoError(t, err)
	utils.CountTopoAddedOrNoneEvent(topoNodeEventChan, numberOfE2Nodes)

	// Check that there are the correct number of relations
	relations, err := topoSdkClient.GetControlRelations()
	assert.NoError(t, err)
	assert.Len(t, relations, numberOfControlRelationships)

	// tear down the simulator
	assert.NoError(t, sim.Uninstall())

	// Wait for the nodes to shut down
	utils.CountTopoRemovedEvent(topoNodeEventChan, numberOfE2Nodes)

	// Check that there are no relations left
	relationsAfter, err := topoSdkClient.GetControlRelations()
	assert.NoError(t, err)
	assert.Len(t, relationsAfter, 0)
}
