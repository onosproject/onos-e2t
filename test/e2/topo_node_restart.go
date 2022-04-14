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
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
)

// TestTopoNodeRestart checks that a subscription recovers after a topo node restart
func (s *TestSuite) TestTopoNodeRestart(t *testing.T) {
	// Create a simulator
	sim := utils.CreateRanSimulatorWithNameOrDie(t, s.c, "topo-restart-subscription")
	assert.NotNil(t, sim)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	nodeID := utils.GetTestNodeID(t)

	// Use one of the cell object IDs for action definition
	cellObjectID := e2utils.GetFirstCellObjectID(t, nodeID)

	subName := "TestTopoNodeRestart"

	// Create a KPM V2 subscription
	kpmv2Sub := e2utils.KPMV2Sub{
		Sub: e2utils.Sub{
			Name:   subName,
			NodeID: nodeID,
		},
		CellObjectID: cellObjectID,
	}
	kpmv2Sub.SubscribeOrFail(ctx, t)

	indicationReport := e2utils.CheckIndicationMessage(t, 5*time.Minute, kpmv2Sub.Sub.Ch)
	indicationMessage := e2smkpmv2.E2SmKpmIndicationMessage{}
	indicationHeader := e2smkpmv2.E2SmKpmIndicationHeader{}

	err := proto.Unmarshal(indicationReport.Payload, &indicationMessage)
	assert.NoError(t, err)
	indMsgFormat1 := indicationMessage.GetIndicationMessageFormats().GetIndicationMessageFormat1()
	assert.Equal(t, indMsgFormat1.GetCellObjId().Value, cellObjectID)
	assert.Equal(t, int(kpmv2Sub.ReportPeriod/kpmv2Sub.Granularity), len(indMsgFormat1.GetMeasData().GetValue()))

	err = proto.Unmarshal(indicationReport.Header, &indicationHeader)
	assert.NoError(t, err)

	t.Log("Restarting topo node")
	client, err := kubernetes.NewForRelease(s.release)
	assert.NoError(t, err)
	topoDeployment, err := client.AppsV1().Deployments().Get(ctx, "onos-topo")
	assert.NoError(t, err)
	pods, err := topoDeployment.Pods().List(ctx)
	assert.NoError(t, err)
	assert.NotZero(t, len(pods))
	pod := pods[0]
	assert.NoError(t, pod.Delete(ctx))

	t.Log("Wait for topo deployment to be ready")
	err = topoDeployment.Wait(ctx, 3*time.Minute)
	assert.NoError(t, err)

	t.Log("Checking indications")
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

	t.Logf("Unsubscribing %s", subName)
	kpmv2Sub.UnsubscribeOrFail(ctx, t)

	e2utils.CheckForEmptySubscriptionList(t)
	utils.UninstallRanSimulatorOrDie(t, sim)
}
