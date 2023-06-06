// SPDX-FileCopyrightText: 2022-present Intel Corporation
// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2

import (
	e2smkpmv2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	"github.com/onosproject/onos-e2t/test/e2utils"
	"github.com/onosproject/onos-e2t/test/utils"
	"google.golang.org/protobuf/proto"
	"time"
)

// TestE2NodeRestart checks that a subscription channel read times out if
// the e2 node is down.
func (s *TestSuite) TestE2NodeRestart() {
	// Create a simulator
	sim := s.CreateRanSimulatorWithNameOrDie("e2node-restart-subscription")

	nodeID := utils.GetTestNodeID(s.T())
	subName := "TestE2NodeRestart"
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

	kpmv2Sub.SubscribeOrFail(s.Context(), s.T())

	indicationReport := e2utils.CheckIndicationMessage(s.T(), e2utils.DefaultIndicationTimeout, kpmv2Sub.Sub.Ch)
	indicationMessage := e2smkpmv2.E2SmKpmIndicationMessage{}
	indicationHeader := e2smkpmv2.E2SmKpmIndicationHeader{}

	err := proto.Unmarshal(indicationReport.Payload, &indicationMessage)
	s.NoError(err)
	indMsgFormat1 := indicationMessage.GetIndicationMessageFormats().GetIndicationMessageFormat1()
	s.Equal(indMsgFormat1.GetCellObjId().Value, cellObjectID)
	s.Equal(int(kpmv2Sub.ReportPeriod/kpmv2Sub.Granularity), len(indMsgFormat1.GetMeasData().GetValue()))

	err = proto.Unmarshal(indicationReport.Header, &indicationHeader)
	s.NoError(err)

	s.T().Log("Restart e2 node")
	s.CrashSimulatorPodOrDie("e2node-restart-subscription")

	s.T().Log("Check indications")
	indicationReport = e2utils.CheckIndicationMessage(s.T(), 5*time.Minute, kpmv2Sub.Sub.Ch)
	indicationMessage = e2smkpmv2.E2SmKpmIndicationMessage{}
	indicationHeader = e2smkpmv2.E2SmKpmIndicationHeader{}

	err = proto.Unmarshal(indicationReport.Payload, &indicationMessage)
	s.NoError(err)
	indMsgFormat1 = indicationMessage.GetIndicationMessageFormats().GetIndicationMessageFormat1()
	s.Equal(indMsgFormat1.GetCellObjId().Value, cellObjectID)
	s.Equal(int(kpmv2Sub.ReportPeriod/kpmv2Sub.Granularity), len(indMsgFormat1.GetMeasData().GetValue()))

	err = proto.Unmarshal(indicationReport.Header, &indicationHeader)
	s.NoError(err)

	s.T().Log("Unsubscribe")
	kpmv2Sub.Sub.UnsubscribeOrFail(s.Context(), s.T())
	s.NoError(err)

	e2utils.CheckForEmptySubscriptionList(s.T())

	// Tear down the simulator
	s.UninstallRanSimulatorOrDie(sim, "e2node-restart-subscription")
}
