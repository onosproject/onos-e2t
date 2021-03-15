// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package e2

import (
	"context"
	"testing"

	modelapi "github.com/onosproject/onos-api/go/onos/ransim/model"

	ransimtypes "github.com/onosproject/onos-api/go/onos/ransim/types"

	subapi "github.com/onosproject/onos-api/go/onos/e2sub/subscription"
	e2sm_rc_pre_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v1/e2sm-rc-pre-ies"
	"github.com/onosproject/onos-e2t/test/utils"
	e2client "github.com/onosproject/onos-ric-sdk-go/pkg/e2"
	"github.com/onosproject/onos-ric-sdk-go/pkg/e2/indication"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
)

// TestSubscriptionOnChange tests E2 subscription on change using ransim and SDK
func (s *TestSuite) TestSubscriptionOnChange(t *testing.T) {
	sim := utils.CreateRanSimulatorWithNameOrDie(t, "subscription-on-change")
	assert.NotNil(t, sim)
	ch := make(chan indication.Indication)
	ctx := context.Background()

	clientConfig := e2client.Config{
		AppID: "subscription-on-change-test",
		E2TService: e2client.ServiceConfig{
			Host: utils.E2TServiceHost,
			Port: utils.E2TServicePort,
		},
		SubscriptionService: e2client.ServiceConfig{
			Host: utils.SubscriptionServiceHost,
			Port: utils.SubscriptionServicePort,
		},
	}
	nodeClient := utils.GetRansimNodeClient(t, sim)
	assert.NotNil(t, nodeClient)
	cellClient := utils.GetRansimCellClient(t, sim)
	assert.NotNil(t, cellClient)
	client, err := e2client.NewClient(clientConfig)
	assert.NoError(t, err)

	nodeIDs, err := utils.GetNodeIDs()
	assert.NoError(t, err)
	testNodeID := nodeIDs[0]

	// Subscription
	eventTriggerBytes, err := utils.CreateRcEventTrigger()
	assert.NoError(t, err)

	subRequest := utils.Subscription{
		NodeID:               testNodeID,
		EncodingType:         subapi.Encoding_ENCODING_PROTO,
		ActionType:           subapi.ActionType_ACTION_TYPE_REPORT,
		EventTrigger:         eventTriggerBytes,
		ServiceModelID:       utils.RcServiceModelID,
		ActionID:             100,
		SubSequentActionType: subapi.SubsequentActionType_SUBSEQUENT_ACTION_TYPE_CONTINUE,
		TimeToWait:           subapi.TimeToWait_TIME_TO_WAIT_ZERO,
	}

	subReq, err := subRequest.Create()
	assert.NoError(t, err)

	sub, err := client.Subscribe(ctx, subReq, ch)
	assert.NoError(t, err)

	indMessage := checkIndicationMessage(t, defaultIndicationTimeout, ch)
	indMessage = checkIndicationMessage(t, defaultIndicationTimeout, ch)
	indMessage = checkIndicationMessage(t, defaultIndicationTimeout, ch)
	header := indMessage.Payload.Header
	ricIndicationHeader := e2sm_rc_pre_ies.E2SmRcPreIndicationHeader{}

	err = proto.Unmarshal(header, &ricIndicationHeader)
	assert.NoError(t, err)
	plmnID := ricIndicationHeader.GetIndicationHeaderFormat1().GetCgi().GetEUtraCgi().GetPLmnIdentity().Value
	testEci := ricIndicationHeader.GetIndicationHeaderFormat1().GetCgi().GetEUtraCgi().GetEUtracellIdentity().Value.Value

	plmnIDValue := ransimtypes.Uint24ToUint32(plmnID)
	ecgi := ransimtypes.ToECGI(ransimtypes.PlmnID(plmnIDValue), ransimtypes.GetECI(testEci))

	testCell, err := cellClient.GetCell(ctx, &modelapi.GetCellRequest{
		ECGI: ecgi,
	})
	assert.NoError(t, err)
	neighborsList := testCell.GetCell().Neighbors
	// Update the list of neighbors
	neighborsList = append(neighborsList[:1], neighborsList[2:]...)
	testCell.Cell.Neighbors = neighborsList
	_, err = cellClient.UpdateCell(ctx, &modelapi.UpdateCellRequest{
		Cell: testCell.Cell,
	})
	assert.NoError(t, err)
	// Expect to receive indication message on neighbor list change
	indMessage = checkIndicationMessage(t, defaultIndicationTimeout, ch)
	err = sub.Close()
	assert.NoError(t, err)
	//err = sim.Uninstall()
	//assert.NoError(t, err)

}
