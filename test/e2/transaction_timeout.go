// SPDX-FileCopyrightText: 2022-present Intel Corporation
// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2

import (
	"time"

	"github.com/onosproject/onos-e2t/test/e2utils"

	"github.com/onosproject/onos-e2t/test/utils"
)

// TestTransactionTimeout tests that channels and subscriptions are removed if their subscription timeout period expires
func (s *TestSuite) TestTransactionTimeout() {
	// spin up a ransim instance
	sim := s.CreateRanSimulatorWithNameOrDie("transaction-timeout")
	s.NotNil(sim)

	// create a KPM V2 subscription
	const baseTimeout = 10 * time.Second
	nodeID := utils.GetTestNodeID(s.T())
	subName := "TestTransactionTimeout"
	ctx, cancel := e2utils.GetCtx()

	cellObjectID := e2utils.GetFirstCellObjectID(s.T(), nodeID)

	// Create a KPM V2 subscription
	kpmv2Sub := e2utils.KPMV2Sub{
		Sub: e2utils.Sub{
			Name:    subName,
			NodeID:  nodeID,
			Timeout: baseTimeout,
		},
		CellObjectID: cellObjectID,
	}
	s.NoError(kpmv2Sub.UseDefaultReportAction())
	kpmv2Sub.SubscribeOrFail(ctx, s.T())

	// make sure the subscription channel is working by reading an indication from it
	indication := e2utils.CheckIndicationMessage(s.T(), e2utils.DefaultIndicationTimeout, kpmv2Sub.Sub.Ch)
	s.NotNil(indication)

	// check that the number of subscriptions and channels is now 1
	s.Equal(1, len(e2utils.GetSubscriptionList(s.T())))
	s.Equal(1, len(e2utils.GetChannelList(s.T())))

	// Cause the subscription to time out and wait for it to happen
	cancel()
	time.Sleep(baseTimeout + (2 * time.Second))

	// Make sure that the subscription and the channel were removed
	e2utils.CheckForEmptySubscriptionList(s.T())
	s.Equal(0, len(e2utils.GetChannelList(s.T())))

	// clean up the simulator
	s.UninstallRanSimulatorOrDie(sim, "transaction-timeout")
}
