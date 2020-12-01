// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package stream

import (
	"context"
	"io"
	"testing"

	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2"
	"github.com/stretchr/testify/assert"
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
	assert.Equal(t, e2api.AppID("2"), stream.Metadata().AppID)
	assert.Equal(t, e2api.InstanceID("3"), stream.Metadata().InstanceID)
	assert.Equal(t, e2api.SubscriptionID("4"), stream.Metadata().SubscriptionID)

	err := stream.Send(Value(MessageID(1), "foo"))
	assert.NoError(t, err)

	msg, err := stream.Recv()
	assert.NoError(t, err)
	assert.Equal(t, MessageID(1), msg.ID)
	assert.Equal(t, "foo", msg.Payload.(string))

	cancel()
	msg, err = stream.Recv()
	assert.Equal(t, io.EOF, err)
	assert.Nil(t, msg.Payload)
}
