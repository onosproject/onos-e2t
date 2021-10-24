// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package stream

import (
	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/subscription"
	"github.com/onosproject/onos-lib-go/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestChannelManager(t *testing.T) {
	subs, err := subscription.NewManager()
	assert.NoError(t, err)

	streams, err := NewBroker(subs)
	assert.NoError(t, err)

	streams.Transactions().Open(e2api.ChannelMeta{
		AppID:          "app-1",
		E2NodeID:       "node-1",
		TransactionID:  "trans-1",
	})

	_, ok := streams.Transactions().Get(e2api.ChannelMeta{
		AppID:          "app-1",
		E2NodeID:       "node-1",
		TransactionID:  "trans-1",
	})
	assert.True(t, ok)

	channel1 := streams.Channels().Open("chan-1", e2api.ChannelMeta{
		AppID:          "app-1",
		AppInstanceID:  "instance-1",
		E2NodeID:       "node-1",
		TransactionID:  "trans-1",
		SubscriptionID: "sub-1",
	})

	_, ok = streams.Channels().Get("chan-1")
	assert.True(t, ok)

	select {
	case <-channel1.Reader().Open():
		t.Error("stream opened prematurely")
	case <-time.After(time.Second):
		break
	}

	sub := subs.Open("sub-1", e2api.SubscriptionMeta{
		E2NodeID: "node-1",
	})

	channel1.Writer().Ack()

	select {
	case err := <-channel1.Reader().Open():
		assert.Nil(t, err)
		break
	case <-time.After(time.Second):
		t.Error("timed out waiting for stream open")
	}

	sub.In() <- newIndication(1)

	select {
	case ind := <-channel1.Reader().Indications():
		assert.Equal(t, int32(1), ind.ProtocolIes.E2ApProtocolIes29.Value.RicRequestorId)
	case <-time.After(time.Second):
		t.Error("Timed out waiting for indication 1")
	}

	sub.In() <- newIndication(2)

	channel2 := streams.Channels().Open("chan-2", e2api.ChannelMeta{
		AppID:          "app-1",
		AppInstanceID:  "instance-2",
		E2NodeID:       "node-1",
		TransactionID:  "trans-1",
		SubscriptionID: "sub-1",
	})

	channel2.Writer().Fail(errors.NewUnknown("something bad happened"))

	select {
	case err := <-channel2.Reader().Open():
		assert.NotNil(t, err)
		assert.True(t, errors.IsUnknown(err))
		break
	case <-time.After(time.Second):
		t.Error("timed out waiting for stream open")
	}

	_, ok = streams.Channels().Get("chan-2")
	assert.False(t, ok)

	channel2 = streams.Channels().Open("chan-2", e2api.ChannelMeta{
		AppID:          "app-1",
		AppInstanceID:  "instance-2",
		E2NodeID:       "node-1",
		TransactionID:  "trans-1",
		SubscriptionID: "sub-1",
	})

	_, ok = streams.Channels().Get("chan-2")
	assert.True(t, ok)

	channel2.Writer().Ack()

	select {
	case err := <-channel2.Reader().Open():
		assert.Nil(t, err)
		break
	case <-time.After(time.Second):
		t.Error("timed out waiting for stream open")
	}

	select {
	case ind := <-channel2.Reader().Indications():
		assert.Equal(t, int32(2), ind.ProtocolIes.E2ApProtocolIes29.Value.RicRequestorId)
	case <-time.After(time.Second):
		t.Error("Timed out waiting for indication 2")
	}

	select {
	case <-channel1.Reader().Done():
		t.Error("stream closed prematurely")
	case <-time.After(time.Second):
		break
	}

	channel1.Writer().Close(nil)

	_, ok = streams.Channels().Get("chan-1")
	assert.False(t, ok)

	select {
	case err := <-channel1.Reader().Done():
		assert.Nil(t, err)
		break
	case <-time.After(time.Second):
		t.Error("timed out waiting for stream close")
	}

	channel2.Writer().Close(errors.NewUnavailable("stream closed"))

	select {
	case err := <-channel2.Reader().Done():
		assert.NotNil(t, err)
		assert.True(t, errors.IsUnavailable(err))
		break
	case <-time.After(time.Second):
		t.Error("timed out waiting for stream close")
	}

	_, ok = streams.Channels().Get("chan-2")
	assert.False(t, ok)

	_, ok = streams.Transactions().Get(e2api.ChannelMeta{
		AppID:          "app-1",
		E2NodeID:       "node-1",
		TransactionID:  "trans-1",
	})
	assert.False(t, ok)
}

func newIndication(requestID int32) *e2appducontents.Ricindication {
	return &e2appducontents.Ricindication{
		ProtocolIes: &e2appducontents.RicindicationIes{
			E2ApProtocolIes29: &e2appducontents.RicindicationIes_RicindicationIes29{
				Value: &e2apies.RicrequestId{
					RicRequestorId: requestID,
				},
			},
		},
	}
}
