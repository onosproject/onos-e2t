// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2

import (
	"context"
	"testing"
	"time"

	"github.com/onosproject/onos-e2t/test/e2utils"

	"github.com/onosproject/helmit/pkg/kubernetes"

	"github.com/onosproject/onos-e2t/test/utils"
	"github.com/stretchr/testify/assert"
)

// TestE2NodeDownSubscription checks that a subscription channel read times out if
// the e2 node is down.
func (s *TestSuite) TestE2NodeDownSubscription(t *testing.T) {
	// Create a simulator
	sim := utils.CreateRanSimulatorWithNameOrDie(t, s.c, "e2node-down-subscription")
	nodeID := utils.GetTestNodeID(t)

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// Use one of the cell object IDs for action definition
	cellObjectID := utils.GetFirstCell(t, nodeID)

	// Create a KPM V2 subscription
	subName := "TestE2NodeDownSubscription"
	kpmv2Sub := e2utils.KPMV2Sub{
		Sub: e2utils.Sub{
			Name:   subName,
			NodeID: nodeID,
		},
		CellObjectID: cellObjectID,
	}
	kpmv2Sub.SubscribeOrFail(ctx, t)

	kube, err := kubernetes.NewForRelease(sim)
	assert.NoError(t, err)

	// Cause the simulator to crash
	utils.UninstallRanSimulatorOrDie(t, sim)

	for {
		pods, err := kube.CoreV1().Pods().List(context.Background())
		assert.NoError(t, err)
		if len(pods) > 0 {
			time.Sleep(time.Second)
		} else {
			t.Log("no ransim pod")
			break
		}
	}

	//  Create the subscription
	_, err = kpmv2Sub.Subscribe(ctx)

	//  Subscribe should have failed because of a timeout
	assert.Error(t, err)
	cancel()

	// Delete the subscription and ran simulator
	sim = utils.CreateRanSimulatorWithNameOrDie(t, s.c, "e2node-down-subscription")
	kpmv2Sub.UnsubscribeOrFail(context.Background(), t)

	e2utils.CheckForEmptySubscriptionList(t)
	err = sim.Uninstall()
	assert.NoError(t, err)

}
