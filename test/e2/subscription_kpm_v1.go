// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package e2

import (
	"context"
	"testing"
	"time"

	subapi "github.com/onosproject/onos-api/go/onos/e2sub/subscription"

	"github.com/onosproject/onos-ric-sdk-go/pkg/e2/indication"
	"github.com/stretchr/testify/assert"

	"github.com/onosproject/onos-e2t/test/utils"
)

// TestSubscriptionKpmV1 tests e2 subscription and subscription delete procedures
func (s *TestSuite) TestSubscriptionKpmV1(t *testing.T) {
	sim := utils.CreateRanSimulatorWithNameOrDie(t, "subscription-kpm-v1")

	e2Client := getE2Client(t, "subscription-kpm-v1-test")

	ch := make(chan indication.Indication)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	nodeIDs, err := utils.GetNodeIDs()
	assert.NoError(t, err)

	eventTriggerBytes, err := utils.CreateKpmV1EventTrigger(12)
	assert.NoError(t, err)
	var actions []subapi.Action
	action := subapi.Action{
		ID:   100,
		Type: subapi.ActionType_ACTION_TYPE_REPORT,
		SubsequentAction: &subapi.SubsequentAction{
			Type:       subapi.SubsequentActionType_SUBSEQUENT_ACTION_TYPE_CONTINUE,
			TimeToWait: subapi.TimeToWait_TIME_TO_WAIT_ZERO,
		},
	}
	actions = append(actions, action)

	subRequest := utils.Subscription{
		NodeID:              nodeIDs[0],
		EncodingType:        subapi.Encoding_ENCODING_PROTO,
		Actions:             actions,
		EventTrigger:        eventTriggerBytes,
		ServiceModelName:    utils.KpmServiceModelName,
		ServiceModelVersion: utils.KpmServiceModelVersion1,
	}

	subReq, err := subRequest.Create()
	assert.NoError(t, err)

	sub, err := e2Client.Subscribe(ctx, subReq, ch)
	assert.NoError(t, err)

	checkIndicationMessage(t, defaultIndicationTimeout, ch)

	err = sub.Close()
	assert.NoError(t, err)

	select {
	case <-ch:
		t.Fatal("received an extraneous indication")

	case <-time.After(10 * time.Second):
		t.Log("Subscription test is PASSED")
	}

	err = sim.Uninstall()
	assert.NoError(t, err)

}
