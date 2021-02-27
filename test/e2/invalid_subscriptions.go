// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package e2

import (
	"context"
	"testing"
	"time"

	subtaskapi "github.com/onosproject/onos-api/go/onos/e2sub/task"

	subapi "github.com/onosproject/onos-api/go/onos/e2sub/subscription"
	"github.com/onosproject/onos-e2t/test/utils"
	e2client "github.com/onosproject/onos-ric-sdk-go/pkg/e2"
	"github.com/onosproject/onos-ric-sdk-go/pkg/e2/indication"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	description   string
	encodingType  subapi.Encoding
	actionType    subapi.ActionType
	serviceModeID string
	expectedError subtaskapi.Cause
}

func runTestCase(t *testing.T, testCase testCase) {
	clientConfig := e2client.Config{
		AppID: "invalid-action-id",
		SubscriptionService: e2client.ServiceConfig{
			Host: utils.SubscriptionServiceHost,
			Port: utils.SubscriptionServicePort,
		},
	}
	client, err := e2client.NewClient(clientConfig)
	assert.NoError(t, err)

	ch := make(chan indication.Indication)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	nodeIDs, err := utils.GetNodeIDs()
	assert.NoError(t, err)

	eventTriggerBytes, err := utils.CreateKpmEventTrigger(12)
	assert.NoError(t, err)

	subRequest := utils.Subscription{
		NodeID:               nodeIDs[0],
		EncodingType:         testCase.encodingType,
		ActionType:           testCase.actionType,
		EventTrigger:         eventTriggerBytes,
		ServiceModelID:       testCase.serviceModeID,
		ActionID:             100,
		SubSequentActionType: subapi.SubsequentActionType_SUBSEQUENT_ACTION_TYPE_CONTINUE,
		TimeToWait:           subapi.TimeToWait_TIME_TO_WAIT_ZERO,
	}

	subReq, err := subRequest.Create()
	assert.NoError(t, err)

	sub, err := client.Subscribe(ctx, subReq, ch)
	assert.NoError(t, err)

	select {
	case err = <-sub.Err():
		t.Log(err.Error())
		assert.Equal(t, testCase.expectedError.String(), err.Error())
	case <-time.After(10 * time.Second):
		t.Fatal("test is failed because of timeout")

	}
}

// TestInvalidActionID tests invalid action ID (i.e. INSERT action) for kpm service model that
// supports just REPORT action
func (s *TestSuite) TestInvalidSubscriptions(t *testing.T) {
	testCases := []struct {
		description   string
		encodingType  subapi.Encoding
		actionType    subapi.ActionType
		serviceModeID string
		expectedError subtaskapi.Cause
	}{
		{description: "Non-existent Service Model ID",
			encodingType:  subapi.Encoding_ENCODING_PROTO,
			actionType:    subapi.ActionType_ACTION_TYPE_REPORT,
			serviceModeID: "no-such-service-model",
			expectedError: subtaskapi.Cause_CAUSE_RIC_RAN_FUNCTION_ID_INVALID,
		},
		{description: "Invalid Action ID",
			encodingType:  subapi.Encoding_ENCODING_PROTO,
			actionType:    subapi.ActionType_ACTION_TYPE_INSERT,
			serviceModeID: utils.KpmServiceModelID,
			expectedError: subtaskapi.Cause_CAUSE_RIC_ACTION_NOT_SUPPORTED,
		},
		{description: "Invalid encoding type",
			encodingType:  77,
			actionType:    subapi.ActionType_ACTION_TYPE_REPORT,
			serviceModeID: utils.KpmServiceModelID,
			expectedError: subtaskapi.Cause_CAUSE_RIC_ACTION_NOT_SUPPORTED,
		},
	}

	sim := utils.CreateRanSimulatorWithNameOrDie(t, "ran-simulator")

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			runTestCase(t, testCase)
		})
	}
	err := sim.Uninstall()
	assert.NoError(t, err)
}
