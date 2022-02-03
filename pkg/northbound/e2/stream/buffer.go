// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package stream

import (
	"container/list"
	"fmt"
	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
	"github.com/prometheus/common/log"
	"sync"
)

const bufferMaxSize = 10000

type bufferID string

func getBufferID(meta e2api.ChannelMeta) bufferID {
	return bufferID(fmt.Sprintf("%s:%s:%s", meta.E2NodeID, meta.AppID, meta.TransactionID))
}

func newChannelBuffer(id bufferID) *channelBuffer {
	buffer := &channelBuffer{
		id:     id,
		in:     make(chan *e2appducontents.Ricindication),
		out:    make(chan *e2appducontents.Ricindication),
		buffer: list.New(),
		cond:   sync.NewCond(&sync.Mutex{}),
	}
	buffer.open()
	return buffer
}

type channelBuffer struct {
	id     bufferID
	in     chan *e2appducontents.Ricindication
	out    chan *e2appducontents.Ricindication
	buffer *list.List
	cond   *sync.Cond
	closed bool
}

func (s *channelBuffer) open() {
	go s.receive()
	go s.send()
}

func (s *channelBuffer) receive() {
	for ind := range s.in {
		s.write(ind)
	}
}

func (s *channelBuffer) write(ind *e2appducontents.Ricindication) {
	s.cond.L.Lock()
	defer s.cond.L.Unlock()
	if s.closed {
		return
	}
	if s.buffer.Len() == bufferMaxSize {
		log.Warnf("Discarded indication: maximum buffer size has been reached for the transaction")
	}
	s.buffer.PushBack(ind)
	s.cond.Signal()
}

func (s *channelBuffer) send() {
	defer close(s.out)
	for {
		ind, ok := s.read()
		if !ok {
			return
		}
		s.out <- ind
	}
}

func (s *channelBuffer) read() (*e2appducontents.Ricindication, bool) {
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

func (s *channelBuffer) Close() {
	close(s.in)
	s.cond.L.Lock()
	s.closed = true
	s.cond.Signal()
	s.cond.L.Unlock()
}
