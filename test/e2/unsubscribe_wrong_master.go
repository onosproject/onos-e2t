// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2

import (
	"context"
	"testing"

	"github.com/onosproject/onos-e2t/test/e2utils"

	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	"github.com/stretchr/testify/assert"

	"github.com/onosproject/onos-e2t/test/utils"
)

// TestUnsubscribeWrongMaster tests e2 subscription to a non-master node
func (s *TestSuite) TestUnsubscribeWrongMaster(t *testing.T) {
	if s.E2TReplicaCount == 1 {
		// Test is not applicable - no non-master nodes
		t.Skip("Test not applicable to single node")
		return
	}

	sim := utils.CreateRanSimulatorWithNameOrDie(t, s.c, "unsubscribe-non-master")
	assert.NotNil(t, sim)

	ctx, cancel := context.WithTimeout(context.Background(), subscriptionTimeout)
	defer cancel()

	e2NodeID := utils.GetTestNodeID(t)

	nonMasters := utils.GetE2NodeNonMasterNodes(t, e2NodeID)
	master := utils.GetE2NodeMaster(t, e2NodeID)

	nonMasterClient := utils.GetSubClientForIP(t, nonMasters[0].IP, nonMasters[0].Port)
	assert.NotNil(t, nonMasterClient)
	masterClient := utils.GetSubClientForIP(t, master.Interface.IP, master.Interface.Port)
	assert.NotNil(t, masterClient)

	// Create a KPM V2 subscription
	kpmv2Sub := e2utils.KPMV2Sub{
		Sub: e2utils.Sub{
			NodeID: e2NodeID,
		},
	}
	spec := kpmv2Sub.SubscriptionSpecOrFail(t)

	headers := e2api.RequestHeaders{
		AppID:         "app",
		AppInstanceID: "",
		E2NodeID:      e2api.E2NodeID(e2NodeID),
		ServiceModel: e2api.ServiceModel{
			Name:    utils.KpmServiceModelName,
			Version: utils.Version2,
		},
	}

	req := &e2api.SubscribeRequest{
		Headers:       headers,
		TransactionID: "sub1",
		Subscription:  spec,
	}

	c, err := masterClient.Subscribe(ctx, req)
	assert.NoError(t, err)

	msg, err := c.Recv()
	assert.NotNil(t, msg)
	assert.NoError(t, err)

	unsubscribeRequest := &e2api.UnsubscribeRequest{
		Headers:       headers,
		TransactionID: "sub1",
	}
	unsubscribeResponse, err := nonMasterClient.Unsubscribe(ctx, unsubscribeRequest)
	assert.NoError(t, err)
	assert.NotNil(t, unsubscribeResponse)

	e2utils.CheckForEmptySubscriptionList(t)
	utils.UninstallRanSimulatorOrDie(t, sim)
}
