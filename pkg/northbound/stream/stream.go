// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package stream

import (
	"context"
	"errors"
	"io"
	"sync"
)

// Value creates a new value message
func Value(id MessageID, payload []byte) Message {
	return Message{
		ID:      id,
		Payload: payload,
	}
}

// Error creates a new error message
func Error(id MessageID, err error) Message {
	return Message{
		ID:    id,
		Error: err,
	}
}

// MessageID is a message identifier
type MessageID uint64

// Message is a stream message
type Message struct {
	// ID is the message identifier
	ID MessageID

	// Payload is the message payload
	Payload []byte

	// Error is the message error
	Error error
}

// ReadStream is a read stream
type ReadStream interface {
	// Recv reads a message from the stream
	Recv() (Message, error)

	// Context returns the stream context
	Context() context.Context
}

// WriteStream is a write stream
type WriteStream interface {
	io.Closer

	// Send sends a message on the stream
	Send(message Message) error
}

// ID is a stream identifier
type ID string

// Stream is a read/write stream
type Stream interface {
	ReadStream
	WriteStream

	// ID returns the stream identifier
	ID() ID
}

// channelReadStream is a channel-based read stream
type channelReadStream struct {
	readCh <-chan Message
	ctx    context.Context
}

func (s *channelReadStream) Recv() (Message, error) {
	m, ok := <-s.readCh
	if !ok {
		return Message{}, io.EOF
	}
	return m, nil
}

func (s *channelReadStream) Context() context.Context {
	return s.ctx
}

var _ ReadStream = &channelReadStream{}

// channelWriteStream is a channel-based write stream
type channelWriteStream struct {
	writeCh chan<- Message
	mu      sync.Mutex
	closed  bool
}

func (s *channelWriteStream) Send(message Message) error {
	failed := false
	defer func() {
		if e := recover(); e != nil {
			failed = true
		}
	}()
	s.writeCh <- message
	if failed {
		return errors.New("WriteStream is closed")
	}
	return nil
}

func (s *channelWriteStream) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.closed {
		return errors.New("WriteStream is closed")
	}
	close(s.writeCh)
	s.closed = true
	return nil
}

var _ WriteStream = &channelWriteStream{}

func newChannelStream(ctx context.Context, id ID, ch chan Message) Stream {
	stream := &channelStream{
		channelReadStream: &channelReadStream{
			readCh: ch,
			ctx:    ctx,
		},
		channelWriteStream: &channelWriteStream{
			writeCh: ch,
		},
		id: id,
	}
	go func() {
		<-ctx.Done()
		stream.Close()
	}()
	return stream
}

// channelStream is a channel-based stream
type channelStream struct {
	*channelReadStream
	*channelWriteStream
	id ID
}

func (s *channelStream) ID() ID {
	return s.id
}

var _ Stream = &channelStream{}
