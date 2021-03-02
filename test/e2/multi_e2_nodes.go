// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package e2

import (
	"context"
	"testing"
	"time"

	"github.com/onosproject/onos-e2t/test/utils"

	"github.com/stretchr/testify/assert"

	modelapi "github.com/onosproject/onos-api/go/onos/ransim/model"
	ransimtypes "github.com/onosproject/onos-api/go/onos/ransim/types"
)

func (s *TestSuite) TestMultiE2Nodes(t *testing.T) {
	sim := utils.CreateRanSimulatorWithNameOrDie(t, "ran-simulator")
	assert.NotNil(t, sim)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	nodeClient := getRansimNodeClient(t)
	numNodes := getNumNodes(t, nodeClient)
	nodeIDs, err := utils.GetNodeIDs()
	assert.Equal(t, len(nodeIDs), numNodes)
	assert.NoError(t, err)
	initialEnbID := 155000
	for i := 0; i < 50; i++ {
		enbID := i + initialEnbID
		_, err = nodeClient.CreateNode(ctx, &modelapi.CreateNodeRequest{
			Node: &ransimtypes.Node{
				EnbID:         ransimtypes.EnbID(enbID),
				ServiceModels: []string{"kpm", "rc"},
				Controllers:   []string{"e2t-1", "e2t-2"},
			},
		})
		assert.NoError(t, err)
	}

	// To make sure all of the connections are established
	time.Sleep(10 * time.Second)
	numNodes = getNumNodes(t, nodeClient)
	nodeIDs, err = utils.GetNodeIDs()
	assert.NoError(t, err)
	assert.Equal(t, len(nodeIDs), numNodes)

}
