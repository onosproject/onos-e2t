// SPDX-FileCopyrightText: 2022-present Intel Corporation
// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"time"

	e2smkpmv2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	"github.com/onosproject/onos-e2t/test/e2utils"
	"github.com/onosproject/onos-e2t/test/utils"
	"google.golang.org/protobuf/proto"
)

// TestTopoNodeRestart checks that a subscription recovers after a topo node restart
func (s *TestSuite) TestTopoNodeRestart() {
	// Create a simulator
	sim := s.CreateRanSimulatorWithNameOrDie("topo-restart-subscription")
	s.NotNil(sim)

	nodeID := utils.GetTestNodeID(s.T())

	// Use one of the cell object IDs for action definition
	cellObjectID := e2utils.GetFirstCellObjectID(s.T(), nodeID)

	subName := "TestTopoNodeRestart"

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

	indicationReport := e2utils.CheckIndicationMessage(s.T(), 5*time.Minute, kpmv2Sub.Sub.Ch)
	indicationMessage := e2smkpmv2.E2SmKpmIndicationMessage{}
	indicationHeader := e2smkpmv2.E2SmKpmIndicationHeader{}

	err := proto.Unmarshal(indicationReport.Payload, &indicationMessage)
	s.NoError(err)
	indMsgFormat1 := indicationMessage.GetIndicationMessageFormats().GetIndicationMessageFormat1()
	s.Equal(indMsgFormat1.GetCellObjId().Value, cellObjectID)
	s.Equal(int(kpmv2Sub.ReportPeriod/kpmv2Sub.Granularity), len(indMsgFormat1.GetMeasData().GetValue()))

	err = proto.Unmarshal(indicationReport.Header, &indicationHeader)
	s.NoError(err)

	s.T().Log("Restarting topo node")
	pods, err := s.CoreV1().Pods(s.Namespace()).List(s.Context(), v1.ListOptions{
		LabelSelector: "app=onos,type=topo",
	})
	s.NoError(err)
	s.NotZero(len(pods.Items))
	pod := pods.Items[0]
	err = s.CoreV1().Pods(s.Namespace()).Delete(s.Context(), pod.Name, v1.DeleteOptions{})
	s.NoError(err)

	s.T().Log("Wait for topo deployment to be ready")
	// TODO - figure out how to do this with the new K8S API
	//topoDeployment, err := s.AppsV1().Deployments(s.Namespace()).Get(s.Context(), "onos-topo", v1.GetOptions{})
	//err = topoDeployment.Wait(ctx, 3*time.Minute)
	//s.NoError(err)

	s.T().Log("Checking indications")
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

	s.T().Logf("Unsubscribing %s", subName)
	kpmv2Sub.Sub.UnsubscribeOrFail(s.Context(), s.T())

	e2utils.CheckForEmptySubscriptionList(s.T())
	s.UninstallRanSimulatorOrDie(sim, "topo-restart-subscription")
}
