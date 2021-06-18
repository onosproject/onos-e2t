// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package v1beta1

import (
	"container/list"
	"context"
	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
	"github.com/onosproject/onos-lib-go/pkg/errors"
	"io"
	"sync"
)

const bufferMaxSize = 10000

// StreamReader defines methods for reading indications from a Stream
type StreamReader interface {
	StreamIO

	// Recv reads an indication from the stream
	// This method is thread-safe. If multiple goroutines are receiving from the stream, indications will be
	// distributed randomly between them. If no indications are available, the goroutine will be blocked until
	// an indication is received or the Context is canceled. If the context is canceled, a context.Canceled error
	// will be returned. If the stream has been closed, an io.EOF error will be returned.
	Recv(context.Context) (*e2appducontents.Ricindication, error)
}

// StreamWriter is a write stream
type StreamWriter interface {
	StreamIO

	// Send sends an indication on the stream
	// The Send method is non-blocking. If no StreamReader is immediately available to consume the indication
	// it will be placed in a bounded memory buffer. If the buffer is full, an Unavailable error will be returned.
	// This method is thread-safe.
	Send(indication *e2appducontents.Ricindication) error
}

// StreamID is a stream identifier
type StreamID int

// StreamIO is a base interface for Stream information
type StreamIO interface {
	io.Closer
	ID() StreamID
}

// Stream is a read/write stream
type Stream interface {
	StreamIO
	StreamReader
	StreamWriter
}

type streamIO struct {
	streamID StreamID
}

func (s *streamIO) ID() StreamID {
	return s.streamID
}

func newStreamRegistry() *streamRegistry {
	return &streamRegistry{
		subs:    make(map[e2api.SubscriptionID]*subStream),
		streams: make(map[StreamID]*subStream),
	}
}

type streamRegistry struct {
	subs     map[e2api.SubscriptionID]*subStream
	streams  map[StreamID]*subStream
	streamID StreamID
	mu       sync.RWMutex
}

func (s *streamRegistry) getStream(streamID StreamID) (*subStream, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	stream, ok := s.streams[streamID]
	return stream, ok
}

func (s *streamRegistry) getSubStream(subID e2api.SubscriptionID) (*subStream, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	stream, ok := s.subs[subID]
	return stream, ok
}

func (s *streamRegistry) openSubStream(subID e2api.SubscriptionID) *subStream {
	s.mu.Lock()
	defer s.mu.Unlock()
	stream, ok := s.subs[subID]
	if !ok {
		s.streamID++
		stream = newSubStream(s, subID, s.streamID)
		s.subs[subID] = stream
		s.streams[s.streamID] = stream
	}
	return stream
}

func (s *streamRegistry) closeSubStream(subID e2api.SubscriptionID) {
	s.mu.Lock()
	defer s.mu.Unlock()
	stream, ok := s.subs[subID]
	if ok {
		delete(s.subs, subID)
		delete(s.streams, stream.ID())
	}
}

func (s *streamRegistry) Close() error {
	for _, stream := range s.subs {
		stream.Close()
	}
	return nil
}

func newSubStream(registry *streamRegistry, subID e2api.SubscriptionID, streamID StreamID) *subStream {
	stream := &subStream{
		streamIO: &streamIO{
			streamID: streamID,
		},
		streams: registry,
		subID:   subID,
		ch:      make(chan e2appducontents.Ricindication),
		closer:  make(chan struct{}),
		apps:    make(map[e2api.AppID]*appStream),
	}
	go stream.open()
	return stream
}

type subStream struct {
	*streamIO
	streams *streamRegistry
	subID   e2api.SubscriptionID
	ch      chan e2appducontents.Ricindication
	apps    map[e2api.AppID]*appStream
	mu      sync.RWMutex
	closer  chan struct{}
}

func (s *subStream) open() {
	for {
		select {
		case ind := <-s.ch:
			s.mu.RLock()
			for _, appStream := range s.apps {
				_ = appStream.Send(&ind)
			}
			s.mu.RUnlock()
		case <-s.closer:
			return
		}
	}
}

func (s *subStream) openAppStream(appID e2api.AppID) *appStream {
	s.mu.Lock()
	defer s.mu.Unlock()
	stream, ok := s.apps[appID]
	if !ok {
		stream = newAppStream(s, appID)
		s.apps[appID] = stream
	}
	return stream
}

func (s *subStream) getAppStream(appID e2api.AppID) (*appStream, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	stream, ok := s.apps[appID]
	return stream, ok
}

func (s *subStream) closeAppStream(appID e2api.AppID) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.apps, appID)
	if len(s.apps) == 0 {
		s.streams.closeSubStream(s.subID)
	}
}

func (s *subStream) Send(indication *e2appducontents.Ricindication) error {
	s.ch <- *indication
	return nil
}

func (s *subStream) Recv(ctx context.Context) (*e2appducontents.Ricindication, error) {
	return nil, errors.NewNotSupported("Recv not supported")
}

func (s *subStream) Close() error {
	close(s.ch)
	close(s.closer)
	return nil
}

func newAppStream(subStream *subStream, appID e2api.AppID) *appStream {
	ch := make(chan e2appducontents.Ricindication)
	return &appStream{
		subStream:       subStream,
		appID:           appID,
		streamIO:        subStream.streamIO,
		appStreamReader: newAppStreamReader(ch),
		appStreamWriter: newAppStreamWriter(ch),
		instances:       make(map[e2api.AppInstanceID]*instanceStreamReader),
	}
}

type appStream struct {
	*streamIO
	*appStreamReader
	*appStreamWriter
	subStream *subStream
	appID     e2api.AppID
	instances map[e2api.AppInstanceID]*instanceStreamReader
	mu        sync.RWMutex
}

func (s *appStream) openInstanceStream(instanceID e2api.AppInstanceID) StreamReader {
	s.mu.Lock()
	defer s.mu.Unlock()
	stream, ok := s.instances[instanceID]
	if !ok {
		stream = newInstanceStreamReader(instanceID, s)
		s.instances[instanceID] = stream
	}
	return stream
}

func (s *appStream) getInstanceStream(instanceID e2api.AppInstanceID) (StreamReader, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	stream, ok := s.instances[instanceID]
	return stream, ok
}

func (s *appStream) closeInstanceStream(instanceID e2api.AppInstanceID) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.instances, instanceID)
	if len(s.instances) == 0 {
		s.subStream.closeAppStream(s.appID)
	}
}

var _ Stream = &appStream{}

func newAppStreamReader(ch <-chan e2appducontents.Ricindication) *appStreamReader {
	return &appStreamReader{
		ch: ch,
	}
}

type appStreamReader struct {
	ch <-chan e2appducontents.Ricindication
}

func (s *appStreamReader) Recv(ctx context.Context) (*e2appducontents.Ricindication, error) {
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

func newAppStreamWriter(ch chan<- e2appducontents.Ricindication) *appStreamWriter {
	writer := &appStreamWriter{
		ch:     ch,
		buffer: list.New(),
		cond:   sync.NewCond(&sync.Mutex{}),
	}
	writer.open()
	return writer
}

type appStreamWriter struct {
	ch     chan<- e2appducontents.Ricindication
	buffer *list.List
	cond   *sync.Cond
	closed bool
}

// open starts the goroutine propagating indications from the writer to the reader
func (s *appStreamWriter) open() {
	go s.drain()
}

// drain dequeues indications and writes them to the read channel
func (s *appStreamWriter) drain() {
	for {
		ind, ok := s.next()
		if !ok {
			close(s.ch)
			break
		}
		s.ch <- *ind
	}
}

// next reads the next indication from the buffer or blocks until one becomes available
func (s *appStreamWriter) next() (*e2appducontents.Ricindication, bool) {
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

// Send appends the indication to the buffer and notifies the reader
func (s *appStreamWriter) Send(ind *e2appducontents.Ricindication) error {
	s.cond.L.Lock()
	defer s.cond.L.Unlock()
	if s.closed {
		return io.EOF
	}
	if s.buffer.Len() == bufferMaxSize {
		return errors.NewUnavailable("cannot append indication to stream: maximum buffer size has been reached")
	}
	s.buffer.PushBack(ind)
	s.cond.Signal()
	return nil
}

func (s *appStreamWriter) Close() error {
	s.cond.L.Lock()
	defer s.cond.L.Unlock()
	s.closed = true
	s.cond.Signal()
	return nil
}

func newInstanceStreamReader(instanceID e2api.AppInstanceID, appStream *appStream) *instanceStreamReader {
	return &instanceStreamReader{
		instanceID: instanceID,
		appStream:  appStream,
	}
}

type instanceStreamReader struct {
	instanceID e2api.AppInstanceID
	appStream  *appStream
}

func (s *instanceStreamReader) ID() StreamID {
	return s.appStream.ID()
}

func (s *instanceStreamReader) Recv(ctx context.Context) (*e2appducontents.Ricindication, error) {
	return s.appStream.Recv(ctx)
}

func (s *instanceStreamReader) Close() error {
	s.appStream.closeInstanceStream(s.instanceID)
	return nil
}

var _ StreamReader = &instanceStreamReader{}
