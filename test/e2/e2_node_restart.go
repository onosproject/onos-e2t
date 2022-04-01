// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2

import (
	"context"
	"testing"
	"time"

	"github.com/onosproject/helmit/pkg/kubernetes"

	e2smkpmv2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	"github.com/onosproject/onos-e2t/test/e2utils"
	"github.com/onosproject/onos-e2t/test/utils"
	sdkclient "github.com/onosproject/onos-ric-sdk-go/pkg/e2/v1beta1"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
)

// TestE2NodeRestart checks that a subscription channel read times out if
// the e2 node is down.
func (s *TestSuite) TestE2NodeRestart(t *testing.T) {
	// Create a simulator
	sim := utils.CreateRanSimulatorWithNameOrDie(t, s.c, "e2node-restart-subscription")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Get a node and a cell to use for KPM
	nodeID := utils.GetTestNodeID(t)
	cellObjectID := utils.GetFirstCell(t, nodeID)

	// Create a KPM V2 subscription
	subName := "TestE2NodeRestart"
	kpmv2Sub := e2utils.KPMV2Sub{
		Sub: e2utils.Sub{
			Name:   subName,
			NodeID: nodeID,
		},
		CellObjectID: cellObjectID,
	}
	kpmv2Sub.SubscribeOrFail(ctx, t)

	sdkClient := utils.GetE2Client(t, utils.KpmServiceModelName, utils.Version2, sdkclient.ProtoEncoding)
	node := sdkClient.Node(sdkclient.NodeID(nodeID))

	indicationReport := e2utils.CheckIndicationMessage(t, e2utils.DefaultIndicationTimeout, kpmv2Sub.Sub.Ch)
	indicationMessage := e2smkpmv2.E2SmKpmIndicationMessage{}
	indicationHeader := e2smkpmv2.E2SmKpmIndicationHeader{}

	err := proto.Unmarshal(indicationReport.Payload, &indicationMessage)
	assert.NoError(t, err)
	indMsgFormat1 := indicationMessage.GetIndicationMessageFormats().GetIndicationMessageFormat1()
	assert.Equal(t, indMsgFormat1.GetCellObjId().Value, cellObjectID)
	assert.Equal(t, int(kpmv2Sub.ReportPeriod/kpmv2Sub.Granularity), len(indMsgFormat1.GetMeasData().GetValue()))

	err = proto.Unmarshal(indicationReport.Header, &indicationHeader)
	assert.NoError(t, err)

	t.Log("Restart e2 node")
	ransimClient, err := kubernetes.NewForRelease(sim)
	assert.NoError(t, err)
	ransimDep, err := ransimClient.AppsV1().
		Deployments().
		Get(ctx, "e2node-restart-subscription-ran-simulator")
	assert.NoError(t, err)
	ransimPods, err := ransimDep.Pods().List(ctx)
	assert.NoError(t, err)
	assert.NotZero(t, len(ransimPods))
	ransimPod := ransimPods[0]
	err = ransimPod.Delete(ctx)
	assert.NoError(t, err)

	t.Log("Check indications")
	indicationReport = e2utils.CheckIndicationMessage(t, 5*time.Minute, kpmv2Sub.Sub.Ch)
	indicationMessage = e2smkpmv2.E2SmKpmIndicationMessage{}
	indicationHeader = e2smkpmv2.E2SmKpmIndicationHeader{}

	err = proto.Unmarshal(indicationReport.Payload, &indicationMessage)
	assert.NoError(t, err)
	indMsgFormat1 = indicationMessage.GetIndicationMessageFormats().GetIndicationMessageFormat1()
	assert.Equal(t, indMsgFormat1.GetCellObjId().Value, cellObjectID)
	assert.Equal(t, int(kpmv2Sub.ReportPeriod/kpmv2Sub.Granularity), len(indMsgFormat1.GetMeasData().GetValue()))

	err = proto.Unmarshal(indicationReport.Header, &indicationHeader)
	assert.NoError(t, err)

	t.Log("Unsubscribe")
	err = node.Unsubscribe(context.Background(), subName)
	assert.NoError(t, err)

	e2utils.CheckForEmptySubscriptionList(t)

	// Tear down the simulator
	utils.UninstallRanSimulatorOrDie(t, sim)
}
