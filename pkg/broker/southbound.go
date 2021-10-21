// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package broker

import (
	"context"
	"github.com/google/uuid"
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
	"sync"
)

type SouthboundStreamManager interface {
	StreamManager
	Get(id StreamID) (*SouthboundStream, bool)
	List() []*SouthboundStream
	Watch(ctx context.Context, ch chan<- StreamID)
	create(ch chan<- *e2appducontents.Ricindication) *SouthboundStream
	close(id StreamID)
}

type southboundStreamManager struct {
	streamID   StreamID
	streams    map[StreamID]*SouthboundStream
	streamsMu  sync.RWMutex
	watchers   map[uuid.UUID]chan<- StreamID
	watchersMu sync.RWMutex
}

func (s *southboundStreamManager) Get(id StreamID) (*SouthboundStream, bool) {
	s.streamsMu.RLock()
	defer s.streamsMu.RUnlock()
	stream, ok := s.streams[id]
	return stream, ok
}

func (s *southboundStreamManager) List() []*SouthboundStream {
	s.streamsMu.RLock()
	defer s.streamsMu.RUnlock()
	streams := make([]*SouthboundStream, 0, len(s.streams))
	for _, stream := range s.streams {
		streams = append(streams, stream)
	}
	return streams
}

func (s *southboundStreamManager) notify(subID StreamID) {
	s.watchersMu.RLock()
	for _, watcher := range s.watchers {
		watcher <- subID
	}
	s.watchersMu.RUnlock()
}

func (s *southboundStreamManager) Watch(ctx context.Context, ch chan<- StreamID) {
	s.watchersMu.Lock()
	id := uuid.New()
	s.watchers[id] = ch
	s.watchersMu.Unlock()

	s.streamsMu.RLock()
	streamIDs := make([]StreamID, 0, len(s.streams))
	for _, stream := range s.streams {
		streamIDs = append(streamIDs, stream.StreamID)
	}
	s.streamsMu.RUnlock()

	go func() {
		for _, streamID := range streamIDs {
			ch <- streamID
		}
		<-ctx.Done()
		s.watchersMu.Lock()
		delete(s.watchers, id)
		s.watchersMu.Unlock()
	}()
}

func (s *southboundStreamManager) create(ch chan<- *e2appducontents.Ricindication) *SouthboundStream {
	s.streamsMu.Lock()
	defer s.streamsMu.Unlock()
	s.streamID++
	streamID := s.streamID
	stream := &SouthboundStream{
		manager:  s,
		StreamID: streamID,
		C:        ch,
	}
	s.streams[streamID] = stream
	go s.notify(streamID)
	return stream
}

func (s *southboundStreamManager) close(id StreamID) {
	s.streamsMu.Lock()
	defer s.streamsMu.Unlock()
	if _, ok := s.streams[id]; ok {
		delete(s.streams, id)
		go s.notify(id)
	}
}

type SouthboundStream struct {
	manager  SouthboundStreamManager
	StreamID StreamID
	C        chan<- *e2appducontents.Ricindication
}

func (s *SouthboundStream) Close() {
	s.manager.close(s.StreamID)
	close(s.C)
}

var _ Stream = &SouthboundStream{}
