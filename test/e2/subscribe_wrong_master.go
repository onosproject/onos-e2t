// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package e2

import (
	"context"
	"testing"

	"github.com/onosproject/onos-e2t/test/e2utils"

	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	"github.com/stretchr/testify/assert"

	"github.com/onosproject/onos-e2t/test/utils"
)

// TestSubscribeWrongMaster tests e2 subscription to a non-master node
func (s *TestSuite) TestSubscribeWrongMaster(t *testing.T) {
	if s.E2TReplicaCount == 1 {
		// Test is not applicable - no non-master nodes
		t.Skip("Test not applicable to single node")
		return
	}
	sim := utils.CreateRanSimulatorWithNameOrDie(t, s.c, "subscription-wrong-master")
	assert.NotNil(t, sim)

	ctx, cancel := context.WithTimeout(context.Background(), subscriptionTimeout)
	defer cancel()

	e2NodeID := utils.GetTestNodeID(t)

	nonMasters := utils.GetE2NodeNonMasterNodes(t, e2NodeID)

	client := utils.GetSubClientForIP(t, nonMasters[0].IP, nonMasters[0].Port)
	assert.NotNil(t, client)

	spec := utils.CreateKpmV2Sub(t, e2NodeID)

	req := &e2api.SubscribeRequest{
		Headers: e2api.RequestHeaders{
			AppID:         "app",
			AppInstanceID: "",
			E2NodeID:      e2api.E2NodeID(e2NodeID),
			ServiceModel: e2api.ServiceModel{
				Name:    utils.KpmServiceModelName,
				Version: utils.Version2,
			},
		},
		TransactionID: "sub1",
		Subscription:  spec,
	}

	c, err := client.Subscribe(ctx, req)
	assert.NoError(t, err)

	resp, err := c.Recv()
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "Unavailable")

	e2utils.CheckForEmptySubscriptionList(t)
	utils.UninstallRanSimulatorOrDie(t, sim)
}
