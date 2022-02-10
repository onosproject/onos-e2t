// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2utils

import (
	"context"
	"encoding/hex"
	"testing"
	"time"

	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	"github.com/stretchr/testify/assert"

	"github.com/onosproject/onos-e2t/test/utils"
)

const (
	DefaultIndicationTimeout = 10 * time.Second
)

// GetSubscriptionList get list of subscriptions
func GetSubscriptionList(t *testing.T) []e2api.Subscription {
	subClient := utils.GetSubAdminClient(t)
	req := e2api.ListSubscriptionsRequest{}
	subList, err := subClient.ListSubscriptions(context.Background(), &req)
	assert.NoError(t, err)
	return subList.GetSubscriptions()
}

// GetChannelList get list of channels
func GetChannelList(t *testing.T) []e2api.Channel {
	subClient := utils.GetSubAdminClient(t)
	req := e2api.ListChannelsRequest{}
	chanList, err := subClient.ListChannels(context.Background(), &req)
	assert.NoError(t, err)
	return chanList.Channels
}

// CheckSubscriptionGet make sure that the given subscription ID can be fetched via the subscription API
func CheckSubscriptionGet(t *testing.T, expectedID e2api.SubscriptionID) {
	subClient := utils.GetSubAdminClient(t)
	req := e2api.GetSubscriptionRequest{
		SubscriptionID: expectedID,
	}
	fetched, err := subClient.GetSubscription(context.Background(), &req)
	assert.NoError(t, err)

	assert.Equal(t, expectedID, fetched.GetSubscription().ID)
}

// CheckSubscriptionIDInList makes sure that the give subscription ID appears once and only once in the subscription list
func CheckSubscriptionIDInList(t *testing.T, expectedID e2api.SubscriptionID, subList []e2api.Subscription) {
	found := 0
	for _, sub := range subList {
		if sub.ID == expectedID {
			found++
		}
	}
	assert.Equal(t, 1, found, "Subscription %s not found in subscription list", expectedID)
}

// CheckIndicationMessage makes sure that a valid indication message can be read from the channel
func CheckIndicationMessage(t *testing.T, timeout time.Duration, ch chan e2api.Indication) e2api.Indication {
	select {
	case indicationMsg := <-ch:
		t.Logf("Indication Message Header is\n%v\nIndication Message Payload is\n%v", hex.Dump(indicationMsg.Header), hex.Dump(indicationMsg.Payload))
		return indicationMsg
	case <-time.After(timeout):
		assert.Equal(t, false, "failed to receive indication message")
	}
	return e2api.Indication{}
}

func CheckForEmptySubscriptionList(t *testing.T) {
	iterations := 10
	for i := 1; i <= iterations; i++ {
		subList := GetSubscriptionList(t)
		if len(subList) == 0 {
			return
		}
		time.Sleep(2 * time.Second)
	}
	subList := GetSubscriptionList(t)
	assert.Fail(t, "subscription list is not empty:", len(subList))
}

// GetCtx returns a context to use in gRPC calls
func GetCtx() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 2*time.Minute)
}
