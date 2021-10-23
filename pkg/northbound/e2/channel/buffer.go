// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package channel

import (
	"container/list"
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
	"github.com/prometheus/common/log"
	"sync"
)

const bufferMaxSize = 10000

type BufferID string

type Buffer interface {
	ID() BufferID
	In() chan<- *e2appducontents.Ricindication
	Out() <-chan *e2appducontents.Ricindication
	Close()
}

func newChannelBuffer(id BufferID) Buffer {
	buffer := &channelBuffer{
		id:  id,
		in:  make(chan *e2appducontents.Ricindication),
		out: make(chan *e2appducontents.Ricindication),
	}
	buffer.open()
	return buffer
}

type channelBuffer struct {
	id     BufferID
	in     chan *e2appducontents.Ricindication
	out    chan *e2appducontents.Ricindication
	buffer *list.List
	cond   *sync.Cond
	closed bool
}

func (b *channelBuffer) open() {
	go b.receive()
	go b.send()
}

func (b *channelBuffer) ID() BufferID {
	return b.id
}

func (b *channelBuffer) In() chan<- *e2appducontents.Ricindication {
	return b.in
}

func (b *channelBuffer) Out() <-chan *e2appducontents.Ricindication {
	return b.out
}

func (b *channelBuffer) receive() {
	for ind := range b.in {
		b.write(ind)
	}
}

func (b *channelBuffer) write(ind *e2appducontents.Ricindication) {
	b.cond.L.Lock()
	defer b.cond.L.Unlock()
	if b.closed {
		return
	}
	if b.buffer.Len() == bufferMaxSize {
		log.Warnf("Discarded indication: maximum buffer size has been reached for the transaction")
	}
	b.buffer.PushBack(ind)
	b.cond.Signal()
}

func (b *channelBuffer) send() {
	defer close(b.out)
	for {
		ind, ok := b.read()
		if !ok {
			return
		}
		b.out <- ind
	}
}

func (b *channelBuffer) read() (*e2appducontents.Ricindication, bool) {
	b.cond.L.Lock()
	defer b.cond.L.Unlock()
	for b.buffer.Len() == 0 {
		if b.closed {
			return nil, false
		}
		b.cond.Wait()
	}
	result := b.buffer.Front().Value.(*e2appducontents.Ricindication)
	b.buffer.Remove(b.buffer.Front())
	return result, true
}

func (b *channelBuffer) Close() {
	close(b.in)
	b.cond.L.Lock()
	b.closed = true
	b.cond.Signal()
	b.cond.L.Unlock()
}
