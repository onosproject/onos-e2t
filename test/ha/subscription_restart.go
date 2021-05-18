// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package ha

import (
	"context"
	"github.com/onosproject/onos-e2t/test/e2utils"
	"testing"
	"time"

	subapi "github.com/onosproject/onos-api/go/onos/e2sub/subscription"

	"github.com/onosproject/onos-ric-sdk-go/pkg/e2/indication"
	"github.com/stretchr/testify/assert"

	"github.com/onosproject/onos-e2t/test/utils"
)

// TestSubscriptionRestart :
func (s *TestSuite) TestSubscriptionRestart(t *testing.T) {

	// Test currently does not work
	t.Skip()

	sim := utils.CreateRanSimulatorWithNameOrDie(t, s.c, "subscription-kpm-v1")
	assert.NotNil(t, sim)

	e2Client := utils.GetE2Client(t, "subscription-kpm-v1-test")

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
		ServiceModelVersion: utils.Version1,
	}

	const ITERATIONS = 1

	for i := 1; i <= ITERATIONS; i++ {
		subReq, err := subRequest.Create()
		assert.NoError(t, err)

		sub, err := e2Client.Subscribe(ctx, subReq, ch)
		assert.NoError(t, err)
		assert.NotNil(t, sub)

		e2utils.CheckIndicationMessage(t, 60*time.Second, ch)
		e2utils.CheckIndicationMessage(t, 120*time.Second, ch)
		e2utils.CheckIndicationMessage(t, 20*time.Second, ch)

		e2tPod := FindPodWithPrefix(t, "onos-e2t")
		CrashPodOrFail(t, e2tPod)

		time.Sleep(15 * time.Second)
		e2tPodReboot := FindPodWithPrefix(t, "onos-e2t")
		err = e2tPodReboot.Wait(45 * time.Second)
		assert.NoError(t, err)
		time.Sleep(30 * time.Second)

		err = sub.Close()
		assert.NoError(t, err)
	}
	//err = sim.Uninstall()
	assert.NoError(t, err)

}
