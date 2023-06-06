// SPDX-FileCopyrightText: 2022-present Intel Corporation
// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2

import (
	"context"
	"time"

	"github.com/onosproject/onos-e2t/test/e2utils"
	"github.com/onosproject/onos-e2t/test/utils"
)

// TestSubscriptionCancel : tests cancelling the context for a subscription then subscribing again
func (s *TestSuite) TestSubscriptionCancel() {
	const (
		simName    = "subscription-cancel"
		subName    = "TestSubscriptionCancelKpmV2"
		iterations = 10
	)
	var err error
	// make a simulator
	sim := s.CreateRanSimulatorWithNameOrDie(simName)

	// Create a KPM V2 subscription
	nodeID := utils.GetTestNodeID(s.T())
	kpmv2Sub := e2utils.KPMV2Sub{
		Sub: e2utils.Sub{
			Name:   subName,
			NodeID: nodeID,
		},
		CellObjectID: e2utils.GetFirstCellObjectID(s.T(), nodeID),
	}
	s.NoError(kpmv2Sub.UseDefaultReportAction())

	for i := 1; i <= iterations; i++ {
		// Create the subscription
		subCtx, subCancel := context.WithTimeout(s.Context(), 15*time.Second)
		kpmv2Sub.SubscribeOrFail(subCtx, s.T())
		// Check that there is a message available
		_ = e2utils.CheckIndicationMessage(s.T(), e2utils.DefaultIndicationTimeout, kpmv2Sub.Sub.Ch)

		// Cancel the subscription
		subCancel()
		_ = utils.ReadToEndOfChannel(kpmv2Sub.Sub.Ch)
	}

	unsubCtx, unsubCancel := context.WithTimeout(s.Context(), 15*time.Second)
	kpmv2Sub.Sub.UnsubscribeOrFail(unsubCtx, s.T())
	s.NoError(err)
	unsubCancel()

	e2utils.CheckForEmptySubscriptionList(s.T())
	s.UninstallRanSimulatorOrDie(sim, simName)
}
