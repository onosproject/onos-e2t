// SPDX-FileCopyrightText: 2022-present Intel Corporation
// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2

import (
	"context"
	"time"

	e2smkpmies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	e2smrcpreies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre_go/v2/e2sm-rc-pre-v2-go"
	"github.com/onosproject/onos-e2t/test/e2utils"
	"github.com/onosproject/onos-e2t/test/utils"
	"google.golang.org/protobuf/proto"
)

// TestMultiSmSubscription tests multiple subscription to different service models on different nodes
func (s *TestSuite) TestMultiSmSubscription() {
	sim := s.CreateRanSimulatorWithNameOrDie("multi-sm-subscription")
	s.NotNil(sim)

	nodeIDs := utils.GetTestNodeIDs(s.T(), 2)
	s.True(len(nodeIDs) > 0)

	kpmNodeID := nodeIDs[0]
	rcPreNodeID := nodeIDs[1]

	KPMSubName := "TestSubscriptionKpmV2"
	RCSubName := "TestSubscriptionRCPreV2"

	KPMCtx, KPMCancel := context.WithTimeout(s.Context(), 30*time.Second)

	nodeID := utils.GetTestNodeID(s.T())

	cellObjectID := e2utils.GetFirstCellObjectID(s.T(), nodeID)

	// Create a KPM V2 subscription
	kpmv2Sub := e2utils.KPMV2Sub{
		Sub: e2utils.Sub{
			Name:   KPMSubName,
			NodeID: kpmNodeID,
		},
		CellObjectID: cellObjectID,
	}
	s.NoError(kpmv2Sub.UseDefaultReportAction())

	kpmv2Sub.SubscribeOrFail(KPMCtx, s.T())

	// Subscribe to RC service model
	rcPreSub := e2utils.RCPreSub{
		Sub: e2utils.Sub{
			Name:   RCSubName,
			NodeID: rcPreNodeID,
		},
	}
	s.NoError(rcPreSub.UseDefaultReportAction())

	rcPreSub.SubscribeOrFail(KPMCtx, s.T())

	// Check that indications can be received
	KPMMsg := e2utils.CheckIndicationMessage(s.T(), e2utils.DefaultIndicationTimeout, kpmv2Sub.Sub.Ch)
	RCMsg := e2utils.CheckIndicationMessage(s.T(), e2utils.DefaultIndicationTimeout, rcPreSub.Sub.Ch)

	kpmIndicationHeader := &e2smkpmies.E2SmKpmIndicationHeader{}
	rcIndicationHeader := &e2smrcpreies.E2SmRcPreIndicationHeader{}

	err := proto.Unmarshal(KPMMsg.Header, kpmIndicationHeader)
	s.NoError(err)

	err = proto.Unmarshal(RCMsg.Header, rcIndicationHeader)
	s.NoError(err)

	// Clean up subscriptions
	kpmv2Sub.Sub.UnsubscribeOrFail(s.Context(), s.T())

	rcPreSub.Sub.UnsubscribeOrFail(s.Context(), s.T())
	s.NoError(err)

	KPMCancel()

	e2utils.CheckForEmptySubscriptionList(s.T())
	s.UninstallRanSimulatorOrDie(sim, "multi-sm-subscription")
}
