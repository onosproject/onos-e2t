// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package stream

import (
	"context"
	api "github.com/onosproject/onos-e2t/api/ricapi/e2/v1beta1"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

func TestStream(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	ch := make(chan Message, 1)
	meta := Metadata{
		AppID:          "2",
		InstanceID:     "3",
		SubscriptionID: "4",
	}
	stream := newChannelStream(ctx, ID(1), meta, ch)
	assert.Equal(t, ID(1), stream.ID())
	assert.Equal(t, api.AppID("2"), stream.Metadata().AppID)
	assert.Equal(t, api.InstanceID("3"), stream.Metadata().InstanceID)
	assert.Equal(t, api.SubscriptionID("4"), stream.Metadata().SubscriptionID)

	err := stream.Send(Value(MessageID(1), []byte("foo")))
	assert.NoError(t, err)

	msg, err := stream.Recv()
	assert.NoError(t, err)
	assert.Equal(t, MessageID(1), msg.ID)
	assert.Equal(t, "foo", string(msg.Payload))

	cancel()
	msg, err = stream.Recv()
	assert.Equal(t, io.EOF, err)
	assert.Nil(t, msg.Payload)
}
