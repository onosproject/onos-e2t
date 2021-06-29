// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package e2utils

import (
	"context"
	"testing"
	"time"

	"github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	"github.com/stretchr/testify/assert"

	"github.com/onosproject/onos-e2t/test/utils"
)

const (
	DefaultIndicationTimeout = 10 * time.Second
)

// GetSubscriptionList get  list of subscriptions
func GetSubscriptionList(t *testing.T) []v1beta1.Subscription {
	subClient := utils.GetSubAdminClient(t)
	req := v1beta1.ListSubscriptionsRequest{}
	subList, err := subClient.ListSubscriptions(context.Background(), &req)
	assert.NoError(t, err)
	return subList.GetSubscriptions()
}

// CheckSubscriptionGet make sure that the given subscription ID can be fetched via the subscription API
func CheckSubscriptionGet(t *testing.T, expectedID v1beta1.SubscriptionID) {
	subClient := utils.GetSubAdminClient(t)
	req := v1beta1.GetSubscriptionRequest{
		SubscriptionID: expectedID,
	}
	fetched, err := subClient.GetSubscription(context.Background(), &req)
	assert.NoError(t, err)

	assert.Equal(t, expectedID, fetched.GetSubscription().ID)
}

// CheckSubscriptionIDInList makes sure that the give subscription ID appears once and only once in the subscription list
func CheckSubscriptionIDInList(t *testing.T, expectedID v1beta1.SubscriptionID, subList []v1beta1.Subscription) {
	found := 0
	for _, sub := range subList {
		if sub.ID == expectedID {
			found++
		}
	}
	assert.Equal(t, 1, found, "Subscription %s not found in subscription list", expectedID)
}

// CheckIndicationMessage makes sure that a valid indication message can be read from the channel
func CheckIndicationMessage(t *testing.T, timeout time.Duration, ch chan v1beta1.Indication) v1beta1.Indication {
	select {
	case indicationMsg := <-ch:
		t.Log(indicationMsg)
		return indicationMsg
	case <-time.After(timeout):
		assert.Equal(t, false, "failed to receive indication message")
	}
	return v1beta1.Indication{}
}

func CheckForEmptySubscriptionList(t *testing.T) {
	subList := GetSubscriptionList(t)
	assert.Equal(t, 0, len(subList))
}
