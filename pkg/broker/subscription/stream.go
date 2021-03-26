// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package subscription

import (
	"container/list"
	"context"
	"github.com/onosproject/onos-api/go/onos/e2sub/subscription"
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
	"io"
	"sync"
)

// StreamReader is a read stream
type StreamReader interface {
	StreamIO

	// Recv reads a message from the stream
	Recv(context.Context) (*e2appducontents.Ricindication, error)
}

// StreamWriter is a write stream
type StreamWriter interface {
	StreamIO

	// Send sends a message on the stream
	Send(indication *e2appducontents.Ricindication) error
}

// StreamID is a stream identifier
type StreamID int

type StreamIO interface {
	io.Closer
	SubscriptionID() subscription.ID
	StreamID() StreamID
}

// Stream is a read/write stream
type Stream interface {
	StreamIO
	StreamReader
	StreamWriter
}

func newBufferedStream(id subscription.ID, streamID StreamID) Stream {
	ch := make(chan e2appducontents.Ricindication)
	return &bufferedStream{
		bufferedIO: &bufferedIO{
			subID:    id,
			streamID: streamID,
		},
		bufferedReader: newBufferedReader(ch),
		bufferedWriter: newBufferedWriter(ch),
	}
}

type bufferedIO struct {
	subID    subscription.ID
	streamID StreamID
}

func (s *bufferedIO) SubscriptionID() subscription.ID {
	return s.subID
}

func (s *bufferedIO) StreamID() StreamID {
	return s.streamID
}

type bufferedStream struct {
	*bufferedIO
	*bufferedReader
	*bufferedWriter
}

var _ Stream = &bufferedStream{}

func newBufferedReader(ch <-chan e2appducontents.Ricindication) *bufferedReader {
	return &bufferedReader{
		ch: ch,
	}
}

type bufferedReader struct {
	ch <-chan e2appducontents.Ricindication
}

func (s *bufferedReader) Recv(ctx context.Context) (*e2appducontents.Ricindication, error) {
	select {
	case ind, ok := <-s.ch:
		if !ok {
			return nil, io.EOF
		}
		return &ind, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func newBufferedWriter(ch chan<- e2appducontents.Ricindication) *bufferedWriter {
	writer := &bufferedWriter{
		ch:     ch,
		buffer: list.New(),
		cond:   sync.NewCond(&sync.Mutex{}),
	}
	go writer.open()
	return writer
}

type bufferedWriter struct {
	ch     chan<- e2appducontents.Ricindication
	buffer *list.List
	cond   *sync.Cond
	closed bool
}

func (s *bufferedWriter) open() {
	for {
		message, ok := s.recv()
		if !ok {
			break
		}
		s.ch <- *message
	}
}

func (s *bufferedWriter) recv() (*e2appducontents.Ricindication, bool) {
	s.cond.L.Lock()
	defer s.cond.L.Unlock()
	for s.buffer.Len() == 0 {
		if s.closed {
			return nil, false
		}
		s.cond.Wait()
	}
	result := s.buffer.Front().Value.(*e2appducontents.Ricindication)
	s.buffer.Remove(s.buffer.Front())
	return result, true
}

func (s *bufferedWriter) Send(message *e2appducontents.Ricindication) error {
	s.cond.L.Lock()
	defer s.cond.L.Unlock()
	if s.closed {
		return io.EOF
	}
	s.buffer.PushBack(message)
	s.cond.Signal()
	return nil
}

func (s *bufferedWriter) Close() error {
	s.cond.L.Lock()
	defer s.cond.L.Unlock()
	s.closed = true
	close(s.ch)
	return nil
}
