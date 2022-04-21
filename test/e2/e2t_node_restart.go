// SPDX-FileCopyrightText: 2022-present Intel Corporation
// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2

import (
	"context"
	"strings"
	"testing"
	"time"

	topoapi "github.com/onosproject/onos-api/go/onos/topo"

	"github.com/onosproject/helmit/pkg/kubernetes"

	"github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	e2smkpmv2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	"github.com/onosproject/onos-e2t/test/e2utils"
	"github.com/onosproject/onos-e2t/test/utils"
	sdkclient "github.com/onosproject/onos-ric-sdk-go/pkg/e2/v1beta1"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
)

// TestE2TNodeRestart checks that a subscription recovers after an E2T node restart
func (s *TestSuite) TestE2TNodeRestart(t *testing.T) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	topoSdkClient, err := utils.NewTopoClient()
	assert.NoError(t, err)

	topoE2NodeEventChan := make(chan topoapi.Event)
	err = topoSdkClient.WatchE2Nodes(ctx, topoE2NodeEventChan)
	assert.NoError(t, err)
	// Create a simulator
	sim := utils.CreateRanSimulatorWithNameOrDie(t, s.c, "e2t-restart-subscription")

	nodeID := utils.GetTestNodeID(t)

	mastershipState, err := topoSdkClient.GetE2NodeMastershipState(ctx, nodeID)
	assert.NoError(t, err)
	currentMastershipTerm := mastershipState.Term
	assert.Greater(t, currentMastershipTerm, uint64(0))
	t.Logf("Current mastership term: %d", currentMastershipTerm)

	cellObjectID := e2utils.GetFirstCellObjectID(t, nodeID)
	subName := "TestE2TNodeRestart"

	// Create a KPM V2 subscription
	kpmv2Sub := e2utils.KPMV2Sub{
		Sub: e2utils.Sub{
			Name:   subName,
			NodeID: nodeID,
		},
		CellObjectID: cellObjectID,
	}
	assert.NoError(t, kpmv2Sub.UseDefaultReportAction())

	subSpec, err := kpmv2Sub.CreateSubscriptionSpec()
	assert.NoError(t, err)
	reportPeriod := kpmv2Sub.ReportPeriod
	granularity := kpmv2Sub.Granularity

	sdkClient := utils.GetE2Client(t, utils.KpmServiceModelName, utils.Version2, sdkclient.ProtoEncoding)
	node := sdkClient.Node(sdkclient.NodeID(nodeID))
	ch := make(chan v1beta1.Indication)
	_, err = node.Subscribe(ctx, subName, subSpec, ch)
	assert.NoError(t, err)

	indicationReport := e2utils.CheckIndicationMessage(t, e2utils.DefaultIndicationTimeout, ch)
	indicationMessage := e2smkpmv2.E2SmKpmIndicationMessage{}
	indicationHeader := e2smkpmv2.E2SmKpmIndicationHeader{}

	err = proto.Unmarshal(indicationReport.Payload, &indicationMessage)
	assert.NoError(t, err)
	indMsgFormat1 := indicationMessage.GetIndicationMessageFormats().GetIndicationMessageFormat1()
	assert.Equal(t, indMsgFormat1.GetCellObjId().Value, cellObjectID)
	assert.Equal(t, int(reportPeriod/granularity), len(indMsgFormat1.GetMeasData().GetValue()))

	err = proto.Unmarshal(indicationReport.Header, &indicationHeader)
	assert.NoError(t, err)

	master := utils.GetE2NodeMaster(t, nodeID)

	t.Logf("Deleting e2t master with ID %s for e2 node: %s", master.ID, nodeID)
	sdranClient, err := kubernetes.NewForRelease(s.release)
	assert.NoError(t, err)
	sdranDep, err := sdranClient.AppsV1().
		Deployments().
		Get(ctx, "onos-e2t")
	assert.NoError(t, err)
	masterPod, err := sdranDep.Pods().Get(ctx, strings.TrimPrefix(string(master.ID), "e2:"))
	assert.NoError(t, err)
	err = masterPod.Delete(ctx)
	assert.NoError(t, err)

	t.Log("Checking indications")
	indicationReport = e2utils.CheckIndicationMessage(t, 2*time.Minute, ch)
	indicationMessage = e2smkpmv2.E2SmKpmIndicationMessage{}
	indicationHeader = e2smkpmv2.E2SmKpmIndicationHeader{}

	err = proto.Unmarshal(indicationReport.Payload, &indicationMessage)
	assert.NoError(t, err)
	indMsgFormat1 = indicationMessage.GetIndicationMessageFormats().GetIndicationMessageFormat1()
	assert.Equal(t, indMsgFormat1.GetCellObjId().Value, cellObjectID)
	assert.Equal(t, int(reportPeriod/granularity), len(indMsgFormat1.GetMeasData().GetValue()))

	err = proto.Unmarshal(indicationReport.Header, &indicationHeader)
	assert.NoError(t, err)

	mastershipState, err = topoSdkClient.GetE2NodeMastershipState(ctx, nodeID)
	assert.NoError(t, err)
	t.Logf("Mastership term after restarting E2T: %d", mastershipState.GetTerm())
	assert.Equal(t, currentMastershipTerm+1, mastershipState.GetTerm())

	t.Logf("Unsubscribing %s", subName)
	err = node.Unsubscribe(ctx, subName)
	assert.NoError(t, err)

	e2utils.CheckForEmptySubscriptionList(t)
	utils.UninstallRanSimulatorOrDie(t, sim)
}
