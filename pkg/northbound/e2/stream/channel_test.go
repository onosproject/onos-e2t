// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package stream

import (
	"context"
	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	v2 "github.com/onosproject/onos-e2t/api/e2ap/v2"
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/stream"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
	"github.com/onosproject/onos-lib-go/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestChannelStreams(t *testing.T) {
	subs, err := stream.NewManager()
	assert.NoError(t, err)

	channels, err := NewManager(subs)
	assert.NoError(t, err)

	channel1 := channels.Open("chan-1", e2api.ChannelMeta{
		AppID:          "app-1",
		AppInstanceID:  "instance-1",
		E2NodeID:       "node-1",
		TransactionID:  "trans-1",
		SubscriptionID: "sub-1",
	})

	_, ok := channels.Get("chan-1")
	assert.True(t, ok)

	ctx1, cancel1 := context.WithCancel(context.Background())
	stream1 := channel1.Output().Open(ctx1)

	select {
	case <-stream1.Ready():
		t.Error("stream ready prematurely")
	case <-time.After(time.Second):
		break
	}

	sub := subs.Open("sub-1")

	channel1.Input().Open()

	select {
	case <-stream1.Ready():
		assert.Nil(t, stream1.Err())
		break
	case <-time.After(time.Second):
		t.Error("timed out waiting for stream ready")
	}

	sub.In() <- newIndication(1)

	select {
	case ind := <-stream1.Indications():
		var rrID int32
		for _, v := range ind.GetProtocolIes() {
			if v.Id == int32(v2.ProtocolIeIDRicrequestID) {
				rrID = v.GetValue().GetRrId().GetRicRequestorId()
				break
			}
		}
		assert.Equal(t, int32(1), rrID)
	case <-time.After(time.Second):
		t.Error("Timed out waiting for indication 1")
	}

	sub.In() <- newIndication(2)

	channel2 := channels.Open("chan-2", e2api.ChannelMeta{
		AppID:          "app-1",
		AppInstanceID:  "instance-2",
		E2NodeID:       "node-1",
		TransactionID:  "trans-1",
		SubscriptionID: "sub-1",
	})

	channel2.Input().Close(errors.NewUnknown("something bad happened"))

	ctx2, cancel2 := context.WithCancel(context.Background())
	stream2 := channel2.Output().Open(ctx2)

	select {
	case <-stream2.Done():
		assert.NotNil(t, stream2.Err())
		assert.True(t, errors.IsUnknown(stream2.Err()))
		break
	case <-time.After(time.Second):
		t.Error("timed out waiting for stream done")
	}

	_, ok = channels.Get("chan-2")
	assert.False(t, ok)

	cancel2()

	channel2 = channels.Open("chan-2", e2api.ChannelMeta{
		AppID:          "app-1",
		AppInstanceID:  "instance-2",
		E2NodeID:       "node-1",
		TransactionID:  "trans-1",
		SubscriptionID: "sub-1",
	})

	_, ok = channels.Get("chan-2")
	assert.True(t, ok)

	channel2.Input().Open()

	ctx2, cancel2 = context.WithCancel(context.Background())
	stream2 = channel2.Output().Open(ctx2)

	select {
	case <-stream2.Ready():
		assert.Nil(t, stream2.Err())
		break
	case <-time.After(time.Second):
		t.Error("timed out waiting for stream ready")
	}

	select {
	case ind := <-stream2.Indications():
		var rrID int32
		for _, v := range ind.GetProtocolIes() {
			if v.Id == int32(v2.ProtocolIeIDRicrequestID) {
				rrID = v.GetValue().GetRrId().GetRicRequestorId()
				break
			}
		}
		assert.Equal(t, int32(2), rrID)
	case <-time.After(time.Second):
		t.Error("Timed out waiting for indication 2")
	}

	select {
	case <-stream1.Done():
		t.Error("stream done prematurely")
	case <-time.After(time.Second):
		break
	}

	channel1.Input().Close(nil)

	_, ok = channels.Get("chan-1")
	assert.False(t, ok)
	cancel1()

	select {
	case <-stream1.Done():
		assert.Nil(t, stream1.Err())
		break
	case <-time.After(time.Second):
		t.Error("timed out waiting for stream done")
	}

	channel2.Input().Close(errors.NewUnavailable("stream closed"))

	select {
	case <-stream2.Done():
		assert.NotNil(t, stream2.Err())
		assert.True(t, errors.IsUnavailable(stream2.Err()))
		break
	case <-time.After(time.Second):
		t.Error("timed out waiting for stream done")
	}

	_, ok = channels.Get("chan-2")
	assert.False(t, ok)
	cancel2()

	channel3 := channels.Open("chan-3", e2api.ChannelMeta{
		AppID:          "app-1",
		AppInstanceID:  "instance-3",
		E2NodeID:       "node-1",
		TransactionID:  "trans-1",
		SubscriptionID: "sub-1",
	})

	_, ok = channels.Get("chan-3")
	assert.True(t, ok)

	ctx3, cancel3 := context.WithCancel(context.Background())
	channel3.Output().Open(ctx3)
	assert.Len(t, channel3.Output().Streams(), 1)

	cancel3()

	time.Sleep(time.Second)
	assert.Len(t, channel3.Output().Streams(), 0)
}

func newIndication(requestID int32) *e2appducontents.Ricindication {
	ri := &e2appducontents.Ricindication{
		ProtocolIes: make([]*e2appducontents.RicindicationIes, 0),
	}
	ri.SetRicRequestID(types.RicRequest{
		RequestorID: types.RicRequestorID(requestID),
		InstanceID:  types.RicInstanceID(0),
	})

	return ri
}
