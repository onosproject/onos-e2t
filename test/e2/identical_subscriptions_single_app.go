// SPDX-FileCopyrightText: 2022-present Intel Corporation
// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2

import (
	"context"
	"testing"

	e2smkpmv2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	"github.com/onosproject/onos-e2t/test/e2utils"
	"google.golang.org/protobuf/proto"

	"github.com/stretchr/testify/assert"

	"github.com/onosproject/onos-e2t/test/utils"
)

func verifyIndicationMessages(t *testing.T, cellObjectID string, sub e2utils.KPMV2Sub) {
	indicationReport := e2utils.CheckIndicationMessage(t, e2utils.DefaultIndicationTimeout, sub.Sub.Ch)
	indicationMessage := e2smkpmv2.E2SmKpmIndicationMessage{}
	indicationHeader := e2smkpmv2.E2SmKpmIndicationHeader{}

	err := proto.Unmarshal(indicationReport.Payload, &indicationMessage)
	assert.NoError(t, err)
	indMsgFormat1 := indicationMessage.GetIndicationMessageFormats().GetIndicationMessageFormat1()
	assert.Equal(t, indMsgFormat1.GetCellObjId().Value, cellObjectID)
	assert.Equal(t, int(sub.ReportPeriod/sub.Granularity), len(indMsgFormat1.GetMeasData().GetValue()))

	err = proto.Unmarshal(indicationReport.Header, &indicationHeader)
	assert.NoError(t, err)
}

// TestIdenticalSubscriptionSingleApp tests identical subscriptions are absorbed by E2T in a single xApp
func (s *TestSuite) TestIdenticalSubscriptionSingleApp(t *testing.T) {
	sim := utils.CreateRanSimulatorWithNameOrDie(t, s.c, "identical-subscriptions-single-app")
	assert.NotNil(t, sim)

	ctx, cancel := context.WithTimeout(context.Background(), subscriptionTimeout)
	defer cancel()

	nodeID := utils.GetTestNodeID(t)

	// Use one of the cell object IDs for action definition
	cellObjectID := e2utils.GetFirstCellObjectID(t, nodeID)

	subName1 := "identical-sub1"
	subName2 := "identical-sub2"

	kpmv2Sub1 := e2utils.KPMV2Sub{
		Sub: e2utils.Sub{
			Name:   subName1,
			NodeID: nodeID,
		},
		CellObjectID: cellObjectID,
	}
	channelID1, err := kpmv2Sub1.Subscribe(ctx)
	assert.NoError(t, err)

	kpmv2Sub2 := e2utils.KPMV2Sub{
		Sub: e2utils.Sub{
			Name:   subName2,
			NodeID: nodeID,
		},
		CellObjectID: cellObjectID,
	}
	channelID2, err := kpmv2Sub2.Subscribe(ctx)
	assert.NoError(t, err)

	assert.True(t, channelID1 != channelID2)

	// Should be able to receive indication messages on both channels
	verifyIndicationMessages(t, cellObjectID, kpmv2Sub1)
	verifyIndicationMessages(t, cellObjectID, kpmv2Sub2)

	subList := e2utils.GetSubscriptionList(t)
	assert.Equal(t, 1, len(subList))

	assert.NoError(t, kpmv2Sub1.Sub.Unsubscribe(ctx))

	subList = e2utils.GetSubscriptionList(t)
	t.Logf("Subscription List after deleting subscription %s is %v:", subName1, subList)

	assert.NoError(t, kpmv2Sub2.Sub.Unsubscribe(ctx))

	subList = e2utils.GetSubscriptionList(t)
	t.Logf("Subscription List after deleting subscription %s is %v:", subName2, subList)

	e2utils.CheckForEmptySubscriptionList(t)
	utils.UninstallRanSimulatorOrDie(t, sim)
}
