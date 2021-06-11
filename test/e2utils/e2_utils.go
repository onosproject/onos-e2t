// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package e2utils

import (
	"context"
	"github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	"github.com/onosproject/onos-ric-sdk-go/pkg/e2/indication"
	"testing"
	"time"

	"github.com/onosproject/onos-api/go/onos/e2sub/subscription"
	subapi "github.com/onosproject/onos-api/go/onos/e2sub/subscription"
	"github.com/stretchr/testify/assert"

	"github.com/onosproject/onos-e2t/test/utils"
)

const (
	DefaultIndicationTimeout = 10 * time.Second
)

// GetSubscriptionList get  list of subscriptions
func GetSubscriptionList(t *testing.T) []subscription.Subscription {
	subClient := utils.GetSubClient(t)
	subList, err := subClient.List(context.Background())
	assert.NoError(t, err)
	return subList
}

// CheckSubscriptionGet make sure that the given subscription ID can be fetched via the subscription API
func CheckSubscriptionGet(t *testing.T, expectedID subapi.ID) {
	subClient := utils.GetSubClient(t)

	fetched, err := subClient.Get(context.Background(), expectedID)
	assert.NoError(t, err)
	assert.NotNil(t, fetched)

	assert.Equal(t, expectedID, fetched.ID)
}

// CheckSubscriptionIDInList makes sure that the give subscription ID appears once and only once in the subscription list
func CheckSubscriptionIDInList(t *testing.T, expectedID subapi.ID, subList []subscription.Subscription) {
	found := 0
	for _, sub := range subList {
		if sub.ID == expectedID {
			found++
		}
	}
	assert.Equal(t, 1, found, "Subscription %s not found in subscription list", expectedID)
}

// CheckIndicationMessage makes sure that a valid indication message can be read from the channel
func CheckIndicationMessage(t *testing.T, timeout time.Duration, ch chan indication.Indication) indication.Indication {
	select {
	case indicationMsg := <-ch:
		t.Log(indicationMsg)
		return indicationMsg
	case <-time.After(timeout):
		assert.Equal(t, false, "failed to receive indication message")
	}
	return indication.Indication{}
}

////////////////////////////////////////////////////////////////////////
//  New API Stuff
///////////////////////////////////////////////////////////////////////

// CheckIndicationMessage2 makes sure that a valid indication message can be read from the channel
func CheckIndicationMessage2(t *testing.T, timeout time.Duration, ch chan v1beta1.Indication) v1beta1.Indication {
	select {
	case indicationMsg := <-ch:
		t.Log(indicationMsg)
		return indicationMsg
	case <-time.After(timeout):
		assert.Equal(t, false, "failed to receive indication message")
	}
	return v1beta1.Indication{}
}
