// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package broker

import (
	"context"
	"github.com/google/uuid"
	"sync"
)

type NorthboundStreamManager interface {
	StreamManager
	Get(id StreamID) (*NorthboundStream, bool)
	List() []*NorthboundStream
	New(ctx context.Context) *NorthboundStream
	Watch(ctx context.Context, ch chan<- StreamID)
	close(id StreamID)
}

type northboundStreamManager struct {
	instance   *AppInstanceStream
	streamID   StreamID
	streams    map[StreamID]*NorthboundStream
	streamsMu  sync.RWMutex
	watchers   map[uuid.UUID]chan<- StreamID
	watchersMu sync.RWMutex
}

func (s *northboundStreamManager) Get(id StreamID) (*NorthboundStream, bool) {
	s.streamsMu.RLock()
	defer s.streamsMu.RUnlock()
	stream, ok := s.streams[id]
	return stream, ok
}

func (s *northboundStreamManager) List() []*NorthboundStream {
	s.streamsMu.RLock()
	defer s.streamsMu.RUnlock()
	streams := make([]*NorthboundStream, 0, len(s.streams))
	for _, stream := range s.streams {
		streams = append(streams, stream)
	}
	return streams
}

func (s *northboundStreamManager) New(ctx context.Context) *NorthboundStream {
	s.streamsMu.Lock()
	defer s.streamsMu.Unlock()
	s.streamID++
	streamID := s.streamID
	ctx, cancel := context.WithCancel(ctx)
	stream := &NorthboundStream{
		AppInstanceStream: s.instance,
		streams:           s,
		StreamID:          streamID,
		ctx:               ctx,
		cancel:            cancel,
	}
	s.streams[streamID] = stream
	stream.open()
	go s.notify(streamID)
	return stream
}

func (s *northboundStreamManager) notify(requestID StreamID) {
	s.watchersMu.RLock()
	for _, watcher := range s.watchers {
		watcher <- requestID
	}
	s.watchersMu.RUnlock()
}

func (s *northboundStreamManager) Watch(ctx context.Context, ch chan<- StreamID) {
	s.watchersMu.Lock()
	id := uuid.New()
	s.watchers[id] = ch
	s.watchersMu.Unlock()

	s.streamsMu.RLock()
	requestIDs := make([]StreamID, 0, len(s.streams))
	for _, request := range s.streams {
		requestIDs = append(requestIDs, request.StreamID)
	}
	s.streamsMu.RUnlock()

	go func() {
		for _, requestID := range requestIDs {
			ch <- requestID
		}
		<-ctx.Done()
		s.watchersMu.Lock()
		delete(s.watchers, id)
		s.watchersMu.Unlock()
	}()
}

func (s *northboundStreamManager) close(id StreamID) {
	s.streamsMu.Lock()
	defer s.streamsMu.Unlock()
	if _, ok := s.streams[id]; ok {
		delete(s.streams, id)
		go s.notify(id)
	}
}

type NorthboundStream struct {
	*AppInstanceStream
	streams  NorthboundStreamManager
	StreamID StreamID
	ctx      context.Context
	cancel   context.CancelFunc
}

func (s *NorthboundStream) Context() context.Context {
	return s.ctx
}

func (s *NorthboundStream) open() {
	go func() {
		<-s.ctx.Done()
		s.streams.close(s.StreamID)
	}()
}

func (s *NorthboundStream) Close() {
	s.cancel()
}

var _ Stream = &NorthboundStream{}
