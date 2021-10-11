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
	numberOfE2Nodes := 2
	numberOfControlRelationships := numberOfE2TNodes * numberOfE2Nodes
	maxWaitForRelations := 15

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

	// Check that there are the correct number of relations
	assert.True(t, utils.Retry(maxWaitForRelations, time.Second,
		func() bool {
			relations, err := topoSdkClient.GetControlRelations()
			assert.NoError(t, err)
			return len(relations) == numberOfControlRelationships
		}))

	// tear down the simulator
	assert.NoError(t, sim.Uninstall())

	// Wait for the nodes to shut down
	utils.CountTopoRemovedEvent(topoNodeEventChan, numberOfControlRelationships)

	// Check that there are no relations left
	assert.True(t, utils.Retry(maxWaitForRelations, time.Second,
		func() bool {
			relations, err := topoSdkClient.GetControlRelations()
			assert.NoError(t, err)
			return len(relations) == 0
		}))
}
