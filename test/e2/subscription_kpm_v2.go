// SPDX-FileCopyrightText: 2022-present Intel Corporation
// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2

import (
	"context"
	e2smkpmv2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	"github.com/onosproject/onos-e2t/test/e2utils"
	"google.golang.org/protobuf/proto"

	"github.com/onosproject/onos-e2t/test/utils"
)

// TestSubscriptionKpmV2 tests e2 subscription and subscription delete procedures using kpm version 2
func (s *TestSuite) TestSubscriptionKpmV2() {
	sim := s.CreateRanSimulatorWithNameOrDie("subscription-kpm-v2")
	s.NotNil(sim)

	ctx, cancel := context.WithTimeout(s.Context(), subscriptionTimeout)
	defer cancel()

	nodeID := utils.GetTestNodeID(s.T())

	subName := "TestSubscriptionKpmV2"

	cellObjectID := e2utils.GetFirstCellObjectID(s.T(), nodeID)

	// Create a KPM V2 subscription
	kpmv2Sub := e2utils.KPMV2Sub{
		Sub: e2utils.Sub{
			Name:   subName,
			NodeID: nodeID,
		},
		CellObjectID: cellObjectID,
	}
	s.NoError(kpmv2Sub.UseDefaultReportAction())

	kpmv2Sub.SubscribeOrFail(ctx, s.T())

	// Read an indication
	indicationReport := e2utils.CheckIndicationMessage(s.T(), e2utils.DefaultIndicationTimeout, kpmv2Sub.Sub.Ch)

	// Check the format of the indiction message
	indicationMessage := e2smkpmv2.E2SmKpmIndicationMessage{}
	s.NoError(proto.Unmarshal(indicationReport.Payload, &indicationMessage))
	indMsgFormat1 := indicationMessage.GetIndicationMessageFormats().GetIndicationMessageFormat1()
	s.Equal(indMsgFormat1.GetCellObjId().Value, cellObjectID)
	s.Equal(int(kpmv2Sub.ReportPeriod/kpmv2Sub.Granularity), len(indMsgFormat1.GetMeasData().GetValue()))

	// Check the format of the indication header
	indicationHeader := e2smkpmv2.E2SmKpmIndicationHeader{}
	s.NoError(proto.Unmarshal(indicationReport.Header, &indicationHeader))
	format1 := indicationHeader.IndicationHeaderFormats.E2SmKpmIndicationHeader.(*e2smkpmv2.IndicationHeaderFormats_IndicationHeaderFormat1)
	s.NotNil(format1)
	s.Equal("ONF", *format1.IndicationHeaderFormat1.VendorName)
	s.Equal("RAN Simulator", *format1.IndicationHeaderFormat1.SenderName)

	// Clean up
	s.NoError(kpmv2Sub.Sub.Unsubscribe(ctx))
	e2utils.CheckForEmptySubscriptionList(s.T())
	s.UninstallRanSimulatorOrDie(sim, "subscription-kpm-v2")
}
