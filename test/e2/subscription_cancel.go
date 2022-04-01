// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2

import (
	"context"
	"testing"
	"time"

	"github.com/onosproject/onos-e2t/test/e2utils"
	"github.com/onosproject/onos-e2t/test/utils"
)

// TestSubscriptionCancel : tests cancelling the context for a subscription then subscribing again
func (s *TestSuite) TestSubscriptionCancel(t *testing.T) {
	const (
		simName    = "subscription-cancel"
		subName    = "TestSubscriptionCancelKpmV2"
		iterations = 10
	)
	// make a simulator
	sim := utils.CreateRanSimulatorWithNameOrDie(t, s.c, simName)

	// Get a node and a cell to use for KPM
	nodeID := utils.GetTestNodeID(t)
	cellObjectID := utils.GetFirstCell(t, nodeID)

	// Create a KPM V2 subscription
	kpmv2Sub := e2utils.KPMV2Sub{
		Sub: e2utils.Sub{
			Name:   subName,
			NodeID: nodeID,
		},
		CellObjectID: cellObjectID,
	}

	var ctx context.Context
	var cancel context.CancelFunc

	for i := 1; i <= iterations; i++ {
		// Create the subscription
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		kpmv2Sub.SubscribeOrFail(ctx, t)

		// Check that there is a message available
		_ = e2utils.CheckIndicationMessage(t, e2utils.DefaultIndicationTimeout, kpmv2Sub.Sub.Ch)

		// Cancel the subscription
		cancel()
		_ = utils.ReadToEndOfChannel(kpmv2Sub.Sub.Ch)
	}

	ctx, cancel = context.WithTimeout(context.Background(), 15*time.Second)
	kpmv2Sub.UnsubscribeOrFail(ctx, t)
	e2utils.CheckForEmptySubscriptionList(t)
	utils.UninstallRanSimulatorOrDie(t, sim)
	cancel()
}
