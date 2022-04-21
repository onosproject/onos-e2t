// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2

import (
	"context"
	"testing"
	"time"

	"github.com/onosproject/onos-e2t/test/e2utils"
	"github.com/stretchr/testify/assert"

	"github.com/onosproject/onos-e2t/test/utils"
)

// TestSubscriptionCancel : tests cancelling the context for a subscription then subscribing again
func (s *TestSuite) TestSubscriptionCancel(t *testing.T) {
	const (
		simName    = "subscription-cancel"
		subName    = "TestSubscriptionCancelKpmV2"
		iterations = 10
	)
	var err error
	// make a simulator
	sim := utils.CreateRanSimulatorWithNameOrDie(t, s.c, simName)

	// Create a KPM V2 subscription
	nodeID := utils.GetTestNodeID(t)
	kpmv2Sub := e2utils.KPMV2Sub{
		Sub: e2utils.Sub{
			Name:   subName,
			NodeID: nodeID,
		},
		CellObjectID: e2utils.GetFirstCellObjectID(t, nodeID),
	}

	for i := 1; i <= iterations; i++ {
		// Create the subscription
		subCtx, subCancel := context.WithTimeout(context.Background(), 15*time.Second)
		kpmv2Sub.SubscribeOrFail(subCtx, t)
		// Check that there is a message available
		_ = e2utils.CheckIndicationMessage(t, e2utils.DefaultIndicationTimeout, kpmv2Sub.Sub.Ch)

		// Cancel the subscription
		subCancel()
		_ = utils.ReadToEndOfChannel(kpmv2Sub.Sub.Ch)
	}

	unsubCtx, unsubCancel := context.WithTimeout(context.Background(), 15*time.Second)
	kpmv2Sub.Sub.UnsubscribeOrFail(unsubCtx, t)
	assert.NoError(t, err)
	unsubCancel()

	e2utils.CheckForEmptySubscriptionList(t)
	utils.UninstallRanSimulatorOrDie(t, sim)
}
