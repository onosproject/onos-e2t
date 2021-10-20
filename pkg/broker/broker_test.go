// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package broker

import (
	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStreamBroker(t *testing.T) {
	broker := NewBroker()
	stream, ok := broker.Streams().Get(1)
	assert.False(t, ok)
	assert.Nil(t, stream)

	sub, ok := broker.Subscriptions().Get("foo")
	assert.False(t, ok)
	assert.Nil(t, sub)

	sub = broker.Subscriptions().Create("foo")
	assert.Equal(t, e2api.SubscriptionID("foo"), sub.SubscriptionID)
	assert.Equal(t, StreamID(1), sub.StreamID)

	sub, ok = broker.Subscriptions().Get("foo")
	assert.True(t, ok)
	assert.Equal(t, e2api.SubscriptionID("foo"), sub.SubscriptionID)
	assert.Equal(t, StreamID(1), sub.StreamID)

	subs := broker.Subscriptions().List()
	assert.Len(t, subs, 1)

	app, ok := sub.Apps().Get("bar")
	assert.False(t, ok)
	assert.Nil(t, app)

	app = sub.Apps().Create("bar")
	assert.Equal(t, e2api.AppID("bar"), app.AppID)
	assert.Equal(t, e2api.SubscriptionID("foo"), app.SubscriptionID)
	assert.Equal(t, StreamID(1), app.StreamID)

	app, ok = sub.Apps().Get("bar")
	assert.True(t, ok)
	assert.Equal(t, e2api.AppID("bar"), app.AppID)
	assert.Equal(t, e2api.SubscriptionID("foo"), app.SubscriptionID)
	assert.Equal(t, StreamID(1), app.StreamID)

	apps := sub.Apps().List()
	assert.Len(t, apps, 1)

	transaction, ok := app.Transactions().Get("baz")
	assert.False(t, ok)
	assert.Nil(t, transaction)

	transaction = app.Transactions().Create("baz")
	assert.Equal(t, e2api.TransactionID("baz"), transaction.TransactionID)
	assert.Equal(t, e2api.AppID("bar"), transaction.AppID)
	assert.Equal(t, e2api.SubscriptionID("foo"), transaction.SubscriptionID)
	assert.Equal(t, StreamID(1), transaction.StreamID)

	transaction, ok = app.Transactions().Get("baz")
	assert.True(t, ok)
	assert.Equal(t, e2api.TransactionID("baz"), transaction.TransactionID)
	assert.Equal(t, e2api.AppID("bar"), transaction.AppID)
	assert.Equal(t, e2api.SubscriptionID("foo"), transaction.SubscriptionID)
	assert.Equal(t, StreamID(1), transaction.StreamID)

	transactions := app.Transactions().List()
	assert.Len(t, transactions, 1)

	stream, ok = broker.Streams().Get(1)
	assert.True(t, ok)
	assert.Equal(t, StreamID(1), stream.StreamID)

	broker.Subscriptions().Close("foo")
	_, ok = broker.Subscriptions().Get("foo")
	assert.False(t, ok)
	subs = broker.Subscriptions().List()
	assert.Len(t, subs, 0)
	_, ok = <-transaction.C
	assert.False(t, ok)
}
