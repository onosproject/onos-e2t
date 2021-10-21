// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package broker

import (
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
	"sync"
)

type StreamManager interface {
	Get(id StreamID) (*Stream, bool)
	create(ch chan<- *e2appducontents.Ricindication) *Stream
	close(id StreamID)
}

type streamManager struct {
	streams  map[StreamID]*Stream
	streamID StreamID
	mu       sync.RWMutex
}

func (s *streamManager) Get(id StreamID) (*Stream, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	stream, ok := s.streams[id]
	return stream, ok
}

func (s *streamManager) create(ch chan<- *e2appducontents.Ricindication) *Stream {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.streamID++
	stream := &Stream{
		StreamID: s.streamID,
		C:        ch,
	}
	s.streams[s.streamID] = stream
	return stream
}

func (s *streamManager) close(id StreamID) {
	s.mu.Lock()
	stream, ok := s.streams[id]
	delete(s.streams, id)
	s.mu.Unlock()
	if ok {
		close(stream.C)
	}
}

type Stream struct {
	StreamID StreamID
	C        chan<- *e2appducontents.Ricindication
}
