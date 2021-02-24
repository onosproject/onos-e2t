// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package e2

import (
	"context"
	"github.com/onosproject/onos-api/go/onos/e2sub/subscription"
	subapi "github.com/onosproject/onos-api/go/onos/e2sub/subscription"
	sdksub "github.com/onosproject/onos-ric-sdk-go/pkg/e2/subscription"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/onosproject/onos-e2t/test/utils"
)

// getSubClient returns an SDK subscription client
func getSubClient(t *testing.T) sdksub.Client {
	conn, err := utils.ConnectSubscriptionServiceHost()
	assert.NoError(t, err)
	assert.NotNil(t, conn)

	return sdksub.NewClient(conn)
}

// checkSubscriptionList checks that the list of subscriptions has the correct length
func checkSubscriptionList(t *testing.T, expectedLen int) []subscription.Subscription {
	subClient := getSubClient(t)
	subList, err := subClient.List(context.Background())
	assert.NoError(t, err)
	assert.Len(t, subList, expectedLen)
	return subList
}

// checkSubscriptionGet make sure that the given subscription ID can be fetched via the subscription API
func checkSubscriptionGet(t *testing.T, expectedID subapi.ID) {
	subClient := getSubClient(t)

	fetched, err := subClient.Get(context.Background(), expectedID)
	assert.NoError(t, err)
	assert.NotNil(t, fetched)

	assert.Equal(t, expectedID, fetched.ID)
}

// checkSubscriptionIDInList makes sure that the give subscription ID appears once and only once in the subscription list
func checkSubscriptionIDInList(t *testing.T, expectedID subapi.ID, subList []subscription.Subscription) {
	found := 0
	for _, sub := range subList {
		if sub.ID == expectedID {
			found++
		}
	}
	assert.Equal(t, 1, found, "Subscription %s not found in subscription list", expectedID)
}
