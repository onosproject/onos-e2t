// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package e2

import (
	"context"
	"testing"
	"time"

	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"

	"github.com/onosproject/onos-e2t/test/e2utils"
	sdkclient "github.com/onosproject/onos-ric-sdk-go/pkg/e2/v1beta1"

	"github.com/onosproject/onos-e2t/test/utils"
	"github.com/stretchr/testify/assert"
)

type invalidSubscriptionTestCase struct {
	description         string
	encodingType        sdkclient.Encoding
	actionType          e2api.ActionType
	serviceModelName    e2api.ServiceModelName
	serviceModelVersion e2api.ServiceModelVersion
	eventTrigger        []byte
	actionID            int32
	expectedError       e2api.Error_Cause_Ric_Type
	enabled             bool
}

func runTestCase(t *testing.T, testCase invalidSubscriptionTestCase) {
	if !testCase.enabled {
		t.Skip()
		return
	}
	sdkClient := utils.GetE2Client2(t, string(testCase.serviceModelName), string(testCase.serviceModelVersion), testCase.encodingType)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	nodeID := utils.GetTestNodeID(t)
	var actions []e2api.Action
	action := e2api.Action{
		ID:   testCase.actionID,
		Type: testCase.actionType,
		SubsequentAction: &e2api.SubsequentAction{
			Type:       e2api.SubsequentActionType_SUBSEQUENT_ACTION_TYPE_CONTINUE,
			TimeToWait: e2api.TimeToWait_TIME_TO_WAIT_ZERO,
		},
	}
	actions = append(actions, action)

	subRequest := utils.Subscription{
		NodeID:              string(nodeID),
		Actions:             actions,
		EventTrigger:        testCase.eventTrigger,
		ServiceModelName:    testCase.serviceModelName,
		ServiceModelVersion: testCase.serviceModelVersion,
	}

	subSpec, err := subRequest.CreateWithActionDefinition2()
	assert.NoError(t, err)

	ch := make(chan e2api.Indication)
	node := sdkClient.Node(sdkclient.NodeID(nodeID))
	_, err = node.Subscribe(ctx, testCase.description, subSpec, ch)
	assert.Error(t, err)
	cancel()

	t.Log("Error:", err, testCase.expectedError)
	//assert.Error(t, testCase.expectedError, err.Error())

	/*for _, detail := range st.Details() {
		assert.Nil(t, detail)
		switch typeDetail := detail.(type) {
		case *e2api.Error:
			// This is currently not working
			t.Log(typeDetail.String())
			assert.Equal(t, testCase.expectedError, typeDetail.String())
		}
	}*/

	err = node.Unsubscribe(ctx, testCase.description)
	t.Log(err)
}

// TestInvalidSubscriptions tests invalid inputs into the SDK
func (s *TestSuite) TestInvalidSubscriptions(t *testing.T) {
	const actionID = 11
	eventTriggerBytes, err := utils.CreateKpmV1EventTrigger(12)
	assert.NoError(t, err)

	testCases := []invalidSubscriptionTestCase{
		{
			description:         "Non-existent Service Model ID",
			enabled:             false,
			encodingType:        sdkclient.ProtoEncoding,
			actionType:          e2api.ActionType_ACTION_TYPE_REPORT,
			serviceModelName:    "no-such-service-model",
			serviceModelVersion: "v1",
			eventTrigger:        eventTriggerBytes,
			actionID:            actionID,
			//expectedError:       e2api.Error_Cause_Ric_RAN_FUNCTION_ID_INVALID,
		},
		{
			description:         "Invalid action type",
			enabled:             false,
			encodingType:        sdkclient.ProtoEncoding,
			actionType:          e2api.ActionType_ACTION_TYPE_INSERT,
			serviceModelName:    utils.KpmServiceModelName,
			serviceModelVersion: utils.Version1,
			eventTrigger:        eventTriggerBytes,
			actionID:            actionID,
			expectedError:       e2api.Error_Cause_Ric_ACTION_NOT_SUPPORTED,
		},
		{
			description:         "Invalid encoding type",
			enabled:             false,
			encodingType:        77,
			actionType:          e2api.ActionType_ACTION_TYPE_REPORT,
			serviceModelName:    utils.KpmServiceModelName,
			serviceModelVersion: utils.Version1,
			eventTrigger:        eventTriggerBytes,
			actionID:            actionID,
			//expectedError:       subtaskapi.Cause_CAUSE_PROTOCOL_MESSAGE_NOT_COMPATIBLE_WITH_RECEIVER_STATE,
		},
		{
			description:         "Invalid action ID",
			enabled:             false,
			encodingType:        sdkclient.ProtoEncoding,
			actionType:          e2api.ActionType_ACTION_TYPE_REPORT,
			serviceModelName:    utils.KpmServiceModelName,
			serviceModelVersion: utils.Version1,
			eventTrigger:        eventTriggerBytes,
			actionID:            100000,
			//expectedError:       subtaskapi.Cause_CAUSE_PROTOCOL_ABSTRACT_SYNTAX_ERROR_FALSELY_CONSTRUCTED_MESSAGE,
		},
		{
			description:         "Invalid event trigger",
			enabled:             false,
			encodingType:        sdkclient.ProtoEncoding,
			actionType:          e2api.ActionType_ACTION_TYPE_REPORT,
			serviceModelName:    utils.KpmServiceModelName,
			serviceModelVersion: utils.Version1,
			eventTrigger:        make([]byte, 50),
			actionID:            actionID,
			//expectedError:       subtaskapi.Cause_CAUSE_PROTOCOL_ABSTRACT_SYNTAX_ERROR_FALSELY_CONSTRUCTED_MESSAGE,
		},
	}

	sim := utils.CreateRanSimulatorWithNameOrDie(t, s.c, "invalid-subscriptions")

	for _, testCase := range testCases {
		pinTestCase := testCase
		t.Run(pinTestCase.description, func(t *testing.T) {
			runTestCase(t, pinTestCase)
		})
	}
	err = sim.Uninstall()
	assert.NoError(t, err)

	e2utils.CheckForEmptySubscriptionList(t)
}
