// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package ha

import (
	"context"
	"testing"
	"time"

	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"

	sdkclient "github.com/onosproject/onos-ric-sdk-go/pkg/e2/v1beta1"

	"github.com/onosproject/onos-e2t/test/e2utils"

	"github.com/stretchr/testify/assert"

	"github.com/onosproject/onos-e2t/test/utils"
)

// TestSubscriptionRestart :
func (s *TestSuite) TestSubscriptionRestart(t *testing.T) {

	// Test currently does not work
	t.Skip()

	sim := utils.CreateRanSimulatorWithNameOrDie(t, s.c, "ha-subscription-kpm-v2")
	assert.NotNil(t, sim)

	e2Client := utils.GetE2Client(t, utils.KpmServiceModelName, utils.Version2, sdkclient.ProtoEncoding)

	ch := make(chan e2api.Indication)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	nodeID := utils.GetTestNodeID(t)
	node := e2Client.Node(sdkclient.NodeID(nodeID))

	eventTriggerBytes, err := utils.CreateKpmV2EventTrigger(5000)
	assert.NoError(t, err)
	var actions []e2api.Action
	action := e2api.Action{
		ID:   100,
		Type: e2api.ActionType_ACTION_TYPE_REPORT,
		SubsequentAction: &e2api.SubsequentAction{
			Type:       e2api.SubsequentActionType_SUBSEQUENT_ACTION_TYPE_CONTINUE,
			TimeToWait: e2api.TimeToWait_TIME_TO_WAIT_ZERO,
		},
	}

	actions = append(actions, action)

	subRequest := utils.Subscription{
		NodeID:              string(nodeID),
		Actions:             actions,
		EventTrigger:        eventTriggerBytes,
		ServiceModelName:    utils.KpmServiceModelName,
		ServiceModelVersion: utils.Version1,
	}

	const ITERATIONS = 1

	for i := 1; i <= ITERATIONS; i++ {
		subReq, err := subRequest.Create()
		assert.NoError(t, err)

		sub, err := node.Subscribe(ctx, "subscription-restart", subReq, ch)
		assert.NoError(t, err)
		assert.NotNil(t, sub)

		e2utils.CheckIndicationMessage(t, 60*time.Second, ch)
		e2utils.CheckIndicationMessage(t, 120*time.Second, ch)
		e2utils.CheckIndicationMessage(t, 20*time.Second, ch)

		e2tPod := FindPodWithPrefix(t, "onos-e2t")
		CrashPodOrFail(t, e2tPod)

		time.Sleep(15 * time.Second)
		e2tPodReboot := FindPodWithPrefix(t, "onos-e2t")
		err = e2tPodReboot.Wait(context.Background(), 45*time.Second)
		assert.NoError(t, err)
		time.Sleep(30 * time.Second)

		//err = sub.Close()
		assert.NoError(t, err)
	}
	//err = sim.Uninstall()
	assert.NoError(t, err)

}
