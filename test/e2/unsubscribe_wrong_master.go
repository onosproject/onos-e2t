// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
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

// TestUnsubscribeWrongMaster tests e2 subscription to a non-master node
func (s *TestSuite) TestUnsubscribeWrongMaster(t *testing.T) {
	sim := utils.CreateRanSimulatorWithNameOrDie(t, s.c, "subscription-kpm-v2")
	assert.NotNil(t, sim)

	ctx, cancel := context.WithTimeout(context.Background(), subscriptionTimeout)
	defer cancel()

	e2NodeID := utils.GetTestNodeID(t)

	master, nonMasters := utils.GetE2Masters(t, e2NodeID)

	fmt.Fprintf(os.Stderr, "master ip is %s:%d\n", master.IP, master.Port)
	fmt.Fprintf(os.Stderr, "non master ip is %s:%d\n", nonMasters[0].IP, nonMasters[0].Port)
	nonMasterClient := utils.GetSubClientForIP(t, nonMasters[0].IP, nonMasters[0].Port)
	assert.NotNil(t, nonMasterClient)
	masterClient := utils.GetSubClientForIP(t, master.IP, master.Port)
	assert.NotNil(t, masterClient)

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

	c, err := masterClient.Subscribe(ctx, req)
	assert.NoError(t, err)

	msg, err := c.Recv()
	assert.NotNil(t, msg)

	unsubscribeRequest := &e2api.UnsubscribeRequest{
		Headers:       e2api.RequestHeaders{},
		TransactionID: "sub1",
	}
	unsubscribeResponse, err := nonMasterClient.Unsubscribe(ctx, unsubscribeRequest)
	assert.NoError(t, err)
	assert.NotNil(t, unsubscribeResponse)

	//assert.NoError(t, sim.Uninstall())
	//e2utils.CheckForEmptySubscriptionList(t)
}
