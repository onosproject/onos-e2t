// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package e2

import (
	"context"
	"testing"

	"github.com/onosproject/onos-e2t/test/e2utils"

	e2smkpmv2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"google.golang.org/protobuf/proto"

	subapi "github.com/onosproject/onos-api/go/onos/e2sub/subscription"

	"github.com/onosproject/onos-ric-sdk-go/pkg/e2/indication"
	"github.com/stretchr/testify/assert"

	"github.com/onosproject/onos-e2t/test/utils"
)

// TestSubscriptionKpmV2AdminAPI tests e2 subscription and subscription delete procedures using kpm version 2
func (s *TestSuite) TestSubscriptionKpmV2AdminAPI(t *testing.T) {
	sim := utils.CreateRanSimulatorWithNameOrDie(t, s.c, "subscription-kpm-v2-admin")
	assert.NotNil(t, sim)

	e2Client := utils.GetE2Client(t, "subscription-kpm-v2-test-admin")

	ch := make(chan indication.Indication)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	nodeIDs, err := utils.NodeIDs()
	assert.NoError(t, err)

	ranFunctions, err := utils.GetRANFunctions(nodeIDs[0])
	assert.NoError(t, err)

	ranFunctionDescription := &e2smkpmv2.E2SmKpmRanfunctionDescription{}
	ranFunctionFound := false
	for _, ranFunction := range ranFunctions {
		if ranFunction.Oid == utils.KpmServiceModelOIDV2 {
			err = proto.Unmarshal(ranFunction.Description, ranFunctionDescription)
			assert.NoError(t, err)
			ranFunctionFound = true
		}
	}

	assert.Equal(t, ranFunctionFound, true)
	reportPeriod := uint32(5000)
	granularity := uint32(500)
	// Kpm v2 interval is defined in ms
	eventTriggerBytes, err := utils.CreateKpmV2EventTrigger(reportPeriod)
	assert.NoError(t, err)

	// Use one of the cell object IDs for action definition
	cellObjectID := ranFunctionDescription.GetRicKpmNodeList()[0].GetCellMeasurementObjectList()[0].CellObjectId.Value
	actionDefinitionBytes, err := utils.CreateKpmV2ActionDefinition(cellObjectID, granularity)
	assert.NoError(t, err)

	var actions []subapi.Action
	action := subapi.Action{
		ID:   100,
		Type: subapi.ActionType_ACTION_TYPE_REPORT,
		SubsequentAction: &subapi.SubsequentAction{
			Type:       subapi.SubsequentActionType_SUBSEQUENT_ACTION_TYPE_CONTINUE,
			TimeToWait: subapi.TimeToWait_TIME_TO_WAIT_ZERO,
		},
		Payload: subapi.Payload{
			Encoding: subapi.Encoding_ENCODING_PROTO,
			Data:     actionDefinitionBytes,
		},
	}

	actions = append(actions, action)

	subRequest := utils.Subscription{
		NodeID:              nodeIDs[0],
		EncodingType:        subapi.Encoding_ENCODING_PROTO,
		EventTrigger:        eventTriggerBytes,
		ServiceModelName:    utils.KpmServiceModelName,
		ServiceModelVersion: utils.Version2,
		Actions:             actions,
	}

	subReq, err := subRequest.CreateWithActionDefinition()
	assert.NoError(t, err)

	sub, err := e2Client.Subscribe(ctx, subReq, ch)
	assert.NoError(t, err)

	indicationReport := e2utils.CheckIndicationMessage(t, e2utils.DefaultIndicationTimeout, ch)
	indicationMessage := e2smkpmv2.E2SmKpmIndicationMessage{}
	indicationHeader := e2smkpmv2.E2SmKpmIndicationHeader{}

	err = proto.Unmarshal(indicationReport.Payload.Message, &indicationMessage)
	assert.NoError(t, err)
	assert.Equal(t, indicationMessage.GetIndicationMessageFormat1().GetCellObjId().Value, cellObjectID)
	assert.Equal(t, int(reportPeriod/granularity), len(indicationMessage.GetIndicationMessageFormat1().GetMeasData().GetValue()))

	err = proto.Unmarshal(indicationReport.Payload.Header, &indicationHeader)
	assert.NoError(t, err)

	err = sub.Close()
	assert.NoError(t, err)

	err = sim.Uninstall()
	assert.NoError(t, err)

}
