// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package channel

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

	chans, err := NewManager(subs)
	assert.NoError(t, err)

	channel1 := chans.Open(&e2api.Channel{
		ID: "chan-1",
		ChannelMeta: e2api.ChannelMeta{
			AppID:          "app-1",
			AppInstanceID:  "instance-1",
			E2NodeID:       "node-1",
			TransactionID:  "trans-1",
			SubscriptionID: "sub-1",
		},
	})

	select {
	case <-channel1.Reader().Open():
		t.Error("stream opened prematurely")
	case <-time.After(time.Second):
		break
	}

	sub := subs.Open(&e2api.Subscription{
		ID: "sub-1",
		SubscriptionMeta: e2api.SubscriptionMeta{
			E2NodeID: "node-1",
		},
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

	channel2 := chans.Open(&e2api.Channel{
		ID: "chan-2",
		ChannelMeta: e2api.ChannelMeta{
			AppID:          "app-1",
			AppInstanceID:  "instance-2",
			E2NodeID:       "node-1",
			TransactionID:  "trans-1",
			SubscriptionID: "sub-1",
		},
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

	channel2 = chans.Open(&e2api.Channel{
		ID: "chan-2",
		ChannelMeta: e2api.ChannelMeta{
			AppID:          "app-1",
			AppInstanceID:  "instance-2",
			E2NodeID:       "node-1",
			TransactionID:  "trans-1",
			SubscriptionID: "sub-1",
		},
	})

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
