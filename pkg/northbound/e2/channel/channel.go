// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package channel

import (
	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
	"sync"
)

type Channel interface {
	ID() e2api.ChannelID
	BufferID() BufferID
	Meta() e2api.ChannelMeta
	Writer() Writer
	Reader() Reader
}

type Writer interface {
	Ack()
	Fail(err error)
	Close(err error)
}

type Reader interface {
	Open() <-chan error
	Indications() <-chan *e2appducontents.Ricindication
	Done() <-chan error
}

func newChannelStream(id e2api.ChannelID, meta e2api.ChannelMeta, buffer Buffer, manager *channelManager) Channel {
	return &channelStream{
		manager: manager,
		id:      id,
		meta:    meta,
		buffer:  buffer,
		openCh:  make(chan error, 1),
		doneCh:  make(chan error, 1),
	}
}

type channelStream struct {
	manager *channelManager
	id      e2api.ChannelID
	meta    e2api.ChannelMeta
	buffer  Buffer
	openCh  chan error
	open    bool
	doneCh  chan error
	done    bool
	mu      sync.RWMutex
}

func (s *channelStream) ID() e2api.ChannelID {
	return s.id
}

func (s *channelStream) Meta() e2api.ChannelMeta {
	return s.meta
}

func (s *channelStream) BufferID() BufferID {
	return s.buffer.ID()
}

func (s *channelStream) Writer() Writer {
	return s
}

func (s *channelStream) Reader() Reader {
	return s
}

func (s *channelStream) Open() <-chan error {
	return s.openCh
}

func (s *channelStream) Ack() {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.open {
		return
	}
	close(s.openCh)
	s.open = true
}

func (s *channelStream) Fail(err error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.open {
		return
	}
	s.manager.close(s)
	s.openCh <- err
	close(s.openCh)
	s.open = true
}

func (s *channelStream) Indications() <-chan *e2appducontents.Ricindication {
	return s.buffer.Out()
}

func (s *channelStream) Done() <-chan error {
	return s.doneCh
}

func (s *channelStream) Close(err error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.done {
		return
	}
	s.manager.close(s)
	if err != nil {
		s.doneCh <- err
	}
	close(s.doneCh)
	s.done = true
}
