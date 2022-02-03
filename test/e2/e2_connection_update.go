// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2

import (
	"testing"
	"time"

	"github.com/onosproject/onos-e2t/test/utils"
	"github.com/stretchr/testify/assert"
)

// TestE2TConnectionUpdate checks that the control relations are correct
func (s *TestSuite) TestE2TConnectionUpdate(t *testing.T) {
	numberOfE2TNodes := int(s.E2TReplicaCount)
	numberOfE2Nodes := 2
	numberOfControlRelationships := numberOfE2TNodes * numberOfE2Nodes
	maxWaitForRelations := 30

	topoSdkClient, err := utils.NewTopoClient()
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
	utils.UninstallRanSimulatorOrDie(t, sim)

	// Check that there are no relations left
	assert.True(t, utils.Retry(maxWaitForRelations, time.Second,
		func() bool {
			relations, err := topoSdkClient.GetControlRelations()
			assert.NoError(t, err)
			return len(relations) == 0
		}))

}
