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

func (s *TestSuite) verifyIndicationMessages(cellObjectID string, sub e2utils.KPMV2Sub) {
	indicationReport := e2utils.CheckIndicationMessage(s.T(), e2utils.DefaultIndicationTimeout, sub.Sub.Ch)
	indicationMessage := e2smkpmv2.E2SmKpmIndicationMessage{}
	indicationHeader := e2smkpmv2.E2SmKpmIndicationHeader{}

	err := proto.Unmarshal(indicationReport.Payload, &indicationMessage)
	s.NoError(err)
	indMsgFormat1 := indicationMessage.GetIndicationMessageFormats().GetIndicationMessageFormat1()
	s.Equal(indMsgFormat1.GetCellObjId().Value, cellObjectID)
	s.Equal(int(sub.ReportPeriod/sub.Granularity), len(indMsgFormat1.GetMeasData().GetValue()))

	err = proto.Unmarshal(indicationReport.Header, &indicationHeader)
	s.NoError(err)
}

// TestIdenticalSubscriptionSingleApp tests identical subscriptions are absorbed by E2T in a single xApp
func (s *TestSuite) TestIdenticalSubscriptionSingleApp() {
	sim := s.CreateRanSimulatorWithNameOrDie("identical-subscriptions-single-app")
	s.NotNil(sim)

	ctx, cancel := context.WithTimeout(s.Context(), subscriptionTimeout)
	defer cancel()

	nodeID := utils.GetTestNodeID(s.T())

	// Use one of the cell object IDs for action definition
	cellObjectID := e2utils.GetFirstCellObjectID(s.T(), nodeID)

	subName1 := "identical-sub1"
	subName2 := "identical-sub2"

	kpmv2Sub1 := e2utils.KPMV2Sub{
		Sub: e2utils.Sub{
			Name:   subName1,
			NodeID: nodeID,
		},
		CellObjectID: cellObjectID,
	}
	s.NoError(kpmv2Sub1.UseDefaultReportAction())
	channelID1, err := kpmv2Sub1.Subscribe(ctx)
	s.NoError(err)

	kpmv2Sub2 := e2utils.KPMV2Sub{
		Sub: e2utils.Sub{
			Name:   subName2,
			NodeID: nodeID,
		},
		CellObjectID: cellObjectID,
	}
	s.NoError(kpmv2Sub2.UseDefaultReportAction())
	channelID2, err := kpmv2Sub2.Subscribe(ctx)
	s.NoError(err)

	s.True(channelID1 != channelID2)

	// Should be able to receive indication messages on both channels
	s.verifyIndicationMessages(cellObjectID, kpmv2Sub1)
	s.verifyIndicationMessages(cellObjectID, kpmv2Sub2)

	subList := e2utils.GetSubscriptionList(s.T())
	s.Equal(1, len(subList))

	s.NoError(kpmv2Sub1.Sub.Unsubscribe(ctx))

	subList = e2utils.GetSubscriptionList(s.T())
	s.T().Logf("Subscription List after deleting subscription %s is %v:", subName1, subList)

	s.NoError(kpmv2Sub2.Sub.Unsubscribe(ctx))

	subList = e2utils.GetSubscriptionList(s.T())
	s.T().Logf("Subscription List after deleting subscription %s is %v:", subName2, subList)

	e2utils.CheckForEmptySubscriptionList(s.T())
	s.UninstallRanSimulatorOrDie(sim, "identical-subscriptions-single-app")
}
