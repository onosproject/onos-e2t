// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package broker

import (
	"context"
	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestStreamBroker(t *testing.T) {
	broker := NewBroker()
	stream, ok := broker.Streams().Get(1)
	assert.False(t, ok)
	assert.Nil(t, stream)

	sub, ok := broker.Subscriptions().Get("foo")
	assert.False(t, ok)
	assert.Nil(t, sub)

	sub = broker.Subscriptions().Open("foo")
	assert.Equal(t, e2api.SubscriptionID("foo"), sub.SubscriptionID)
	assert.Equal(t, StreamID(1), sub.StreamID)

	sub, ok = broker.Subscriptions().Get("foo")
	assert.True(t, ok)
	assert.Equal(t, e2api.SubscriptionID("foo"), sub.SubscriptionID)
	assert.Equal(t, StreamID(1), sub.StreamID)

	subs := broker.Subscriptions().List()
	assert.Len(t, subs, 1)

	stream, ok = broker.Streams().Get(1)
	assert.True(t, ok)
	assert.Equal(t, StreamID(1), stream.StreamID)
	stream.C <- &e2appducontents.Ricindication{
		ProtocolIes: &e2appducontents.RicindicationIes{
			E2ApProtocolIes29: &e2appducontents.RicindicationIes_RicindicationIes29{
				Value: &e2apies.RicrequestId{
					RicInstanceId:  1,
					RicRequestorId: 1,
				},
			},
		},
	}

	time.Sleep(time.Second)

	app, ok := sub.Apps().Get("bar")
	assert.False(t, ok)
	assert.Nil(t, app)

	app = sub.Apps().Open("bar")
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

	stream.C <- &e2appducontents.Ricindication{
		ProtocolIes: &e2appducontents.RicindicationIes{
			E2ApProtocolIes29: &e2appducontents.RicindicationIes_RicindicationIes29{
				Value: &e2apies.RicrequestId{
					RicInstanceId:  1,
					RicRequestorId: 2,
				},
			},
		},
	}

	time.Sleep(time.Second)

	transaction, ok := app.Transactions().Get("baz")
	assert.False(t, ok)
	assert.Nil(t, transaction)

	transaction = app.Transactions().Open("baz")
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

	stream.C <- &e2appducontents.Ricindication{
		ProtocolIes: &e2appducontents.RicindicationIes{
			E2ApProtocolIes29: &e2appducontents.RicindicationIes_RicindicationIes29{
				Value: &e2apies.RicrequestId{
					RicInstanceId:  1,
					RicRequestorId: 3,
				},
			},
		},
	}

	time.Sleep(time.Second)

	instance, ok := transaction.Instances().Get("foo")
	assert.False(t, ok)
	assert.Nil(t, instance)

	instance = transaction.Instances().Open("foo")
	assert.Equal(t, e2api.AppInstanceID("foo"), instance.AppInstanceID)
	assert.Equal(t, e2api.TransactionID("baz"), instance.TransactionID)
	assert.Equal(t, e2api.AppID("bar"), instance.AppID)
	assert.Equal(t, e2api.SubscriptionID("foo"), instance.SubscriptionID)
	assert.Equal(t, StreamID(1), instance.StreamID)

	instance, ok = transaction.Instances().Get("foo")
	assert.True(t, ok)
	assert.Equal(t, e2api.AppInstanceID("foo"), instance.AppInstanceID)
	assert.Equal(t, e2api.TransactionID("baz"), instance.TransactionID)
	assert.Equal(t, e2api.AppID("bar"), instance.AppID)
	assert.Equal(t, e2api.SubscriptionID("foo"), instance.SubscriptionID)
	assert.Equal(t, StreamID(1), instance.StreamID)

	instances := transaction.Instances().List()
	assert.Len(t, instances, 1)

	stream, ok = broker.Streams().Get(1)
	assert.True(t, ok)
	assert.Equal(t, StreamID(1), stream.StreamID)

	stream.C <- &e2appducontents.Ricindication{
		ProtocolIes: &e2appducontents.RicindicationIes{
			E2ApProtocolIes29: &e2appducontents.RicindicationIes_RicindicationIes29{
				Value: &e2apies.RicrequestId{
					RicInstanceId:  1,
					RicRequestorId: 4,
				},
			},
		},
	}

	time.Sleep(time.Second)

	ctx, cancel := context.WithCancel(context.Background())
	nbStream := broker.Subscriptions().
		Open("foo").
		Apps().
		Open("bar").
		Transactions().
		Open("baz").
		Instances().
		Open("foo").
		Streams().
		New(ctx)
	assert.Equal(t, StreamID(1), nbStream.StreamID)
	assert.Equal(t, e2api.AppInstanceID("foo"), instance.AppInstanceID)
	assert.Equal(t, e2api.TransactionID("baz"), instance.TransactionID)
	assert.Equal(t, e2api.AppID("bar"), instance.AppID)
	assert.Equal(t, e2api.SubscriptionID("foo"), instance.SubscriptionID)

	ind, ok := <-nbStream.C
	if !ok {
		t.FailNow()
	} else {
		assert.Equal(t, int32(3), ind.ProtocolIes.E2ApProtocolIes29.Value.RicRequestorId)
		go cancel()
	}

	received := false
	select {
	case ind, ok := <-nbStream.C:
		if !ok {
			t.FailNow()
		} else {
			assert.Equal(t, int32(4), ind.ProtocolIes.E2ApProtocolIes29.Value.RicRequestorId)
			received = true
		}
	case <-nbStream.Context().Done():
		assert.NotNil(t, nbStream.Context().Err())
	}

	nbStream = broker.Subscriptions().
		Open("foo").
		Apps().
		Open("bar").
		Transactions().
		Open("baz").
		Instances().
		Open("foo").
		Streams().
		New(context.Background())
	assert.Equal(t, StreamID(2), nbStream.StreamID)
	assert.Equal(t, e2api.AppInstanceID("foo"), instance.AppInstanceID)
	assert.Equal(t, e2api.TransactionID("baz"), instance.TransactionID)
	assert.Equal(t, e2api.AppID("bar"), instance.AppID)
	assert.Equal(t, e2api.SubscriptionID("foo"), instance.SubscriptionID)

	sub.Close()
	_, ok = broker.Subscriptions().Get("foo")
	assert.False(t, ok)
	subs = broker.Subscriptions().List()
	assert.Len(t, subs, 0)

	if !received {
		ind, ok = <-nbStream.C
		assert.True(t, ok)
		assert.Equal(t, int32(4), ind.ProtocolIes.E2ApProtocolIes29.Value.RicRequestorId)
	}

	_, ok = <-nbStream.C
	assert.False(t, ok)
}
