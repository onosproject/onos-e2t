// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2

import (
	"github.com/onosproject/onos-e2t/test/utils"
	"time"
)

// TestE2TConnectionUpdate checks that the control relations are correct
func (s *TestSuite) TestE2TConnectionUpdate() {
	numberOfE2TNodes := int(s.E2TReplicaCount)
	numberOfE2Nodes := 2
	numberOfControlRelationships := numberOfE2TNodes * numberOfE2Nodes
	maxWaitForRelations := 30

	topoSdkClient, err := utils.NewTopoClient()
	s.NoError(err)

	// create a simulator
	sim := s.CreateRanSimulatorWithNameOrDie("e2t-connection-update")
	s.NotNil(sim)

	// Check that there are the correct number of relations
	s.True(utils.Retry(maxWaitForRelations, time.Second,
		func() bool {
			relations, err := topoSdkClient.GetControlRelations()
			s.NoError(err)
			return len(relations) == numberOfControlRelationships
		}))

	// tear down the simulator
	s.UninstallRanSimulatorOrDie(sim, "e2t-connection-update")

	// Check that there are no relations left
	s.True(utils.Retry(maxWaitForRelations, time.Second,
		func() bool {
			relations, err := topoSdkClient.GetControlRelations()
			s.NoError(err)
			return len(relations) == 0
		}))

}
