// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package e2

import (
	"context"
	"fmt"
	"os"
	"testing"

	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	"github.com/stretchr/testify/assert"

	"github.com/onosproject/onos-e2t/test/utils"
)

// TestSubscriptionWrongMaster tests e2 subscription to a non-master node
func (s *TestSuite) TestSubscriptionWrongMaster(t *testing.T) {
	sim := utils.CreateRanSimulatorWithNameOrDie(t, s.c, "subscription-kpm-v2")
	assert.NotNil(t, sim)

	ctx, cancel := context.WithTimeout(context.Background(), subscriptionTimeout)
	defer cancel()

	client := utils.GetSubClient(t)
	assert.NotNil(t, client)

	nodeIDs := utils.GetTestNodeIDs(t, 2)
	goodNodeID := nodeIDs[0]
	
	topoSdkClient, err := utils.NewTopoClient()
	assert.NoError(t, err)

	cells, err := topoSdkClient.GetCells(context.Background(), goodNodeID)
	assert.NoError(t, err)

	cellObjID := cells[0].CellObjectID
	fmt.Fprintf(os.Stderr, "get cells returns %v\n", cells[0].CellObjectID)
	fmt.Fprintf(os.Stderr, "node %v good node %v\n", utils.GetTestNodeID(t), goodNodeID)

	spec := utils.CreateKpmV2SubWithCell(t, goodNodeID, cellObjID)

	req := &e2api.SubscribeRequest{
		Headers:            e2api.RequestHeaders{
			AppID:         "app",
			AppInstanceID: "",
			E2NodeID:      e2api.E2NodeID(goodNodeID),
			ServiceModel:  e2api.ServiceModel{
				Name:    utils.KpmServiceModelName,
				Version: utils.Version2,
			},
		},
		TransactionID:      "sub1",
		Subscription:       spec,
	}

	c, err := client.Subscribe(ctx, req)
	assert.NoError(t, err)
	assert.NotNil(t, c)

	response, err := c.Recv()
	assert.NoError(t, err)
	assert.NotNil(t, response)

	fmt.Fprintf(os.Stderr, "Recv of:\n%v\n", response)

	//assert.NoError(t, sim.Uninstall())
	//e2utils.CheckForEmptySubscriptionList(t)
}
