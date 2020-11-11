// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package stream

import (
	"context"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

func TestStream(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	ch := make(chan Message, 1)
	stream := newChannelStream(ctx, ID("1"), ch)

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
