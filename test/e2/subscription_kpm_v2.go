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

// TestSubscriptionKpmV2 tests e2 subscription and subscription delete procedures using kpm version 2
func (s *TestSuite) TestSubscriptionKpmV2(t *testing.T) {
	sim := utils.CreateRanSimulatorWithNameOrDie(t, s.c, "subscription-kpm-v2")
	assert.NotNil(t, sim)

	ctx, cancel := context.WithTimeout(context.Background(), subscriptionTimeout)
	defer cancel()

	nodeID := utils.GetTestNodeID(t)

	topoSdkClient, err := utils.NewTopoClient()
	assert.NoError(t, err)

	cells, err := topoSdkClient.GetCells(ctx, nodeID)
	assert.NoError(t, err)

	subName := "TestSubscriptionKpmV2"

	// Use one of the cell object IDs for action definition
	cellObjectID := cells[0].CellObjectID

	// Create a KPM V2 subscription
	kpmv2Sub := e2utils.KPMV2Sub{
		Ctx:          ctx,
		SubName:      subName,
		NodeID:       nodeID,
		CellObjectID: cellObjectID,
	}
	_, err = kpmv2Sub.Subscribe()
	assert.NoError(t, err)

	// Read an indication
	indicationReport := e2utils.CheckIndicationMessage(t, e2utils.DefaultIndicationTimeout, kpmv2Sub.Ch)

	// Check the format of the indiction message
	indicationMessage := e2smkpmv2.E2SmKpmIndicationMessage{}
	assert.NoError(t, proto.Unmarshal(indicationReport.Payload, &indicationMessage))
	indMsgFormat1 := indicationMessage.GetIndicationMessageFormats().GetIndicationMessageFormat1()
	assert.Equal(t, indMsgFormat1.GetCellObjId().Value, cellObjectID)
	assert.Equal(t, int(kpmv2Sub.ReportPeriod/kpmv2Sub.Granularity), len(indMsgFormat1.GetMeasData().GetValue()))

	// Check the format of the indication header
	indicationHeader := e2smkpmv2.E2SmKpmIndicationHeader{}
	assert.NoError(t, proto.Unmarshal(indicationReport.Header, &indicationHeader))
	format1 := indicationHeader.IndicationHeaderFormats.E2SmKpmIndicationHeader.(*e2smkpmv2.IndicationHeaderFormats_IndicationHeaderFormat1)
	assert.NotNil(t, format1)
	assert.Equal(t, "ONF", *format1.IndicationHeaderFormat1.VendorName)
	assert.Equal(t, "RAN Simulator", *format1.IndicationHeaderFormat1.SenderName)

	// Clean up
	assert.NoError(t, kpmv2Sub.Unsubscribe())
	e2utils.CheckForEmptySubscriptionList(t)
	utils.UninstallRanSimulatorOrDie(t, sim)
}
