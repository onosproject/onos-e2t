// SPDX-FileCopyrightText: 2022-present Intel Corporation
// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2

import (
	"fmt"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strings"
	"time"

	topoapi "github.com/onosproject/onos-api/go/onos/topo"

	"github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	e2smkpmv2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	"github.com/onosproject/onos-e2t/test/e2utils"
	"github.com/onosproject/onos-e2t/test/utils"
	sdkclient "github.com/onosproject/onos-ric-sdk-go/pkg/e2/v1beta1"
	"google.golang.org/protobuf/proto"
)

// TestE2TNodeRestart checks that a subscription recovers after an E2T node restart
func (s *TestSuite) TestE2TNodeRestart() {

	topoSdkClient, err := utils.NewTopoClient()
	s.NoError(err)

	topoE2NodeEventChan := make(chan topoapi.Event)
	err = topoSdkClient.WatchE2Nodes(s.Context(), topoE2NodeEventChan)
	s.NoError(err)
	// Create a simulator
	sim := s.CreateRanSimulatorWithNameOrDie("e2t-restart-subscription")

	nodeID := utils.GetTestNodeID(s.T())

	mastershipState, err := topoSdkClient.GetE2NodeMastershipState(s.Context(), nodeID)
	s.NoError(err)
	currentMastershipTerm := mastershipState.Term
	s.Greater(currentMastershipTerm, uint64(0))
	s.T().Logf("Current mastership term: %d", currentMastershipTerm)

	cellObjectID := e2utils.GetFirstCellObjectID(s.T(), nodeID)
	subName := "TestE2TNodeRestart"

	// Create a KPM V2 subscription
	kpmv2Sub := e2utils.KPMV2Sub{
		Sub: e2utils.Sub{
			Name:   subName,
			NodeID: nodeID,
		},
		CellObjectID: cellObjectID,
	}
	s.NoError(kpmv2Sub.UseDefaultReportAction())

	subSpec, err := kpmv2Sub.CreateSubscriptionSpec()
	s.NoError(err)
	reportPeriod := kpmv2Sub.ReportPeriod
	granularity := kpmv2Sub.Granularity

	sdkClient := utils.GetE2Client(s.T(), utils.KpmServiceModelName, utils.Version2, sdkclient.ProtoEncoding)
	node := sdkClient.Node(sdkclient.NodeID(nodeID))
	ch := make(chan v1beta1.Indication)
	_, err = node.Subscribe(s.Context(), subName, subSpec, ch)
	s.NoError(err)

	indicationReport := e2utils.CheckIndicationMessage(s.T(), e2utils.DefaultIndicationTimeout, ch)
	indicationMessage := e2smkpmv2.E2SmKpmIndicationMessage{}
	indicationHeader := e2smkpmv2.E2SmKpmIndicationHeader{}

	err = proto.Unmarshal(indicationReport.Payload, &indicationMessage)
	s.NoError(err)
	indMsgFormat1 := indicationMessage.GetIndicationMessageFormats().GetIndicationMessageFormat1()
	s.Equal(indMsgFormat1.GetCellObjId().Value, cellObjectID)
	s.Equal(int(reportPeriod/granularity), len(indMsgFormat1.GetMeasData().GetValue()))

	err = proto.Unmarshal(indicationReport.Header, &indicationHeader)
	s.NoError(err)

	master := utils.GetE2NodeMaster(s.T(), nodeID)

	s.T().Logf("Deleting e2t master with ID %s for e2 node: %s", master.ID, nodeID)
	s.NoError(err)
	pods, err := s.CoreV1().Pods(s.Namespace()).List(s.Context(), v1.ListOptions{
		LabelSelector: fmt.Sprintf("name=%s", strings.TrimPrefix(string(master.ID), "e2:")),
	})

	s.NoError(err)
	s.Len(pods.Items, 1)
	masterPod := pods.Items[0]
	err = s.CoreV1().Pods(s.Namespace()).Delete(s.Context(), masterPod.Name, v1.DeleteOptions{})
	s.NoError(err)

	s.T().Log("Checking indications")
	indicationReport = e2utils.CheckIndicationMessage(s.T(), 2*time.Minute, ch)
	indicationMessage = e2smkpmv2.E2SmKpmIndicationMessage{}
	indicationHeader = e2smkpmv2.E2SmKpmIndicationHeader{}

	err = proto.Unmarshal(indicationReport.Payload, &indicationMessage)
	s.NoError(err)
	indMsgFormat1 = indicationMessage.GetIndicationMessageFormats().GetIndicationMessageFormat1()
	s.Equal(indMsgFormat1.GetCellObjId().Value, cellObjectID)
	s.Equal(int(reportPeriod/granularity), len(indMsgFormat1.GetMeasData().GetValue()))

	err = proto.Unmarshal(indicationReport.Header, &indicationHeader)
	s.NoError(err)

	mastershipState, err = topoSdkClient.GetE2NodeMastershipState(s.Context(), nodeID)
	s.NoError(err)
	s.T().Logf("Mastership term after restarting E2T: %d", mastershipState.GetTerm())
	s.Equal(currentMastershipTerm+1, mastershipState.GetTerm())

	s.T().Logf("Unsubscribing %s", subName)
	err = node.Unsubscribe(s.Context(), subName)
	s.NoError(err)

	e2utils.CheckForEmptySubscriptionList(s.T())
	s.UninstallRanSimulatorOrDie(sim, "e2t-restart-subscription")
}
