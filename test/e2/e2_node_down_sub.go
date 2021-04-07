// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package e2

import (
	"context"
	"testing"
	"time"

	"github.com/onosproject/helmit/pkg/kubernetes"

	subapi "github.com/onosproject/onos-api/go/onos/e2sub/subscription"
	"github.com/onosproject/onos-e2t/test/utils"
	"github.com/onosproject/onos-ric-sdk-go/pkg/e2/indication"
	"github.com/stretchr/testify/assert"
)

// TestE2NodeDownSubscription checks that a subscription channel read times out if
// the e2 node is down.
func (s *TestSuite) TestE2NodeDownSubscription(t *testing.T) {
	// Create a simulator
	sim := utils.CreateRanSimulatorWithNameOrDie(t, "e2node-down-subscription")

	// Create an e2 client
	e2Client := getE2Client(t, "subscription-e2node-down-test")

	ch := make(chan indication.Indication)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	nodeIDs, err := utils.GetNodeIDs()
	assert.NoError(t, err)

	eventTriggerBytes, err := utils.CreateKpmV1EventTrigger(12)
	assert.NoError(t, err)

	subRequest := utils.Subscription{
		NodeID:               nodeIDs[0],
		EncodingType:         subapi.Encoding_ENCODING_PROTO,
		ActionType:           subapi.ActionType_ACTION_TYPE_REPORT,
		EventTrigger:         eventTriggerBytes,
		ServiceModelName:     utils.KpmServiceModelName,
		ServiceModelVersion:  utils.KpmServiceModelVersion1,
		ActionID:             100,
		SubSequentActionType: subapi.SubsequentActionType_SUBSEQUENT_ACTION_TYPE_CONTINUE,
		TimeToWait:           subapi.TimeToWait_TIME_TO_WAIT_ZERO,
	}

	// Create a subscription request to indication messages from the client
	subReq, err := subRequest.Create()
	assert.NoError(t, err)

	kube, err := kubernetes.NewForRelease(sim)
	assert.NoError(t, err)

	// Cause the simulator to crash
	err = sim.Uninstall()
	assert.NoError(t, err)

	for {
		pods, err := kube.CoreV1().Pods().List()
		assert.NoError(t, err)
		if len(pods) > 0 {
			time.Sleep(time.Second)
		} else {
			break
		}
	}

	//  Create the subscription
	sub, err := e2Client.Subscribe(ctx, subReq, ch)
	assert.NoError(t, err)

	// Make sure that reads on the subscription channel time out. There should be no
	// indication messages available
	var gotIndication bool
	select {
	case indicationMsg := <-ch:
		// We got an indication. This is an error, as there is no E2 node to send one
		gotIndication = true
		t.Log(indicationMsg)

	case <-time.After(10 * time.Second):
		// The read timed out. This is the expected behavior.
		gotIndication = false
	}

	assert.False(t, gotIndication, "Indication message was delivered for a node that is down")
	err = sub.Close()
	assert.NoError(t, err)
}
