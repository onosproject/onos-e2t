// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package e2

import (
	"context"
	"testing"

	e2smkpmv2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"google.golang.org/protobuf/proto"

	subapi "github.com/onosproject/onos-api/go/onos/e2sub/subscription"

	"github.com/onosproject/onos-ric-sdk-go/pkg/e2/indication"
	"github.com/stretchr/testify/assert"

	"github.com/onosproject/onos-e2t/test/utils"
)

// TestSubscriptionKpmV2 tests e2 subscription and subscription delete procedures using kpm version 2
func (s *TestSuite) TestSubscriptionKpmV2(t *testing.T) {
	sim := utils.CreateRanSimulatorWithNameOrDie(t, "subscription-kpm-v2")
	assert.NotNil(t, sim)

	e2Client := getE2Client(t, "subscription-kpm-v2-test")

	ch := make(chan indication.Indication)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	nodeIDs, err := utils.GetNodeIDs()
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
	// Kpm v2 interval is defined in ms
	eventTriggerBytes, err := utils.CreateKpmV2EventTrigger(5000)
	assert.NoError(t, err)

	// Use one of the cell object IDs for action definition
	cellObjectID := ranFunctionDescription.GetRicKpmNodeList()[0].GetCellMeasurementObjectList()[0].CellObjectId.Value
	actionDefinitionBytes, err := utils.CreateKpmV2ActionDefinition(cellObjectID)
	assert.NoError(t, err)

	subRequest := utils.Subscription{
		NodeID:               nodeIDs[0],
		EncodingType:         subapi.Encoding_ENCODING_PROTO,
		ActionType:           subapi.ActionType_ACTION_TYPE_REPORT,
		EventTrigger:         eventTriggerBytes,
		ServiceModelName:     utils.KpmServiceModelName,
		ServiceModelVersion:  utils.KpmServiceModelVersion2,
		ActionID:             100,
		SubSequentActionType: subapi.SubsequentActionType_SUBSEQUENT_ACTION_TYPE_CONTINUE,
		TimeToWait:           subapi.TimeToWait_TIME_TO_WAIT_ZERO,
		ActionDefinition:     actionDefinitionBytes,
	}

	subReq, err := subRequest.CreateWithActionDefinition()
	assert.NoError(t, err)

	sub, err := e2Client.Subscribe(ctx, subReq, ch)
	assert.NoError(t, err)

	indicationReport := checkIndicationMessage(t, defaultIndicationTimeout, ch)
	indicationMessage := e2smkpmv2.E2SmKpmIndicationMessage{}
	indicationHeader := e2smkpmv2.E2SmKpmIndicationHeader{}

	err = proto.Unmarshal(indicationReport.Payload.Message, &indicationMessage)
	assert.NoError(t, err)
	assert.Equal(t, indicationMessage.GetIndicationMessageFormat1().GetCellObjId().Value, cellObjectID)

	err = proto.Unmarshal(indicationReport.Payload.Header, &indicationHeader)
	assert.NoError(t, err)

	err = sub.Close()
	assert.NoError(t, err)

	err = sim.Uninstall()
	assert.NoError(t, err)

}
