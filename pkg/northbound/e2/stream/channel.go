// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package stream

import (
	"context"
	"github.com/google/uuid"
	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
	"sync"
)

type ChannelManager interface {
	Get(channelID e2api.ChannelID) (Channel, bool)
	Open(channelID e2api.ChannelID, meta e2api.ChannelMeta) Channel
	Watch(ctx context.Context, ch chan<- Channel) error
}

func newChannelManager(transactions TransactionManager) ChannelManager {
	return &channelManager{
		transactions:     transactions,
		chans:            make(map[e2api.ChannelID]Channel),
		transactionChans: make(map[TransactionID]map[e2api.ChannelID]bool),
		watchers:         make(map[uuid.UUID]chan<- Channel),
	}
}

type channelManager struct {
	transactions     TransactionManager
	chans            map[e2api.ChannelID]Channel
	chansMu          sync.RWMutex
	transactionChans map[TransactionID]map[e2api.ChannelID]bool
	watchers         map[uuid.UUID]chan<- Channel
	watchersMu       sync.RWMutex
}

func (m *channelManager) Get(channelID e2api.ChannelID) (Channel, bool) {
	m.chansMu.RLock()
	defer m.chansMu.RUnlock()
	channel, ok := m.chans[channelID]
	return channel, ok
}

func (m *channelManager) Open(channelID e2api.ChannelID, meta e2api.ChannelMeta) Channel {
	m.chansMu.Lock()
	defer m.chansMu.Unlock()

	channel, ok := m.chans[channelID]
	if ok {
		return channel
	}

	channel = newChannelStream(channelID, meta, m.transactions.Open(meta), m)
	m.chans[channelID] = channel

	transactionChans, ok := m.transactionChans[channel.TransactionID()]
	if !ok {
		transactionChans = make(map[e2api.ChannelID]bool)
		m.transactionChans[channel.TransactionID()] = transactionChans
	}
	transactionChans[channelID] = true

	go m.notify(channel)
	return channel
}

func (m *channelManager) close(channel Channel) {
	m.chansMu.Lock()
	defer m.chansMu.Unlock()
	delete(m.chans, channel.ID())
	transactionChans := m.transactionChans[channel.TransactionID()]
	delete(transactionChans, channel.ID())
	if len(transactionChans) == 0 {
		if transaction, ok := m.transactions.Get(channel.Meta()); ok {
			transaction.Close()
		}
	}
	go m.notify(channel)
}

func (m *channelManager) Watch(ctx context.Context, ch chan<- Channel) error {
	m.watchersMu.Lock()
	id := uuid.New()
	m.watchers[id] = ch
	m.watchersMu.Unlock()

	m.chansMu.RLock()
	channels := make([]Channel, 0, len(m.chans))
	for _, stream := range m.chans {
		channels = append(channels, stream)
	}
	m.chansMu.RUnlock()

	go func() {
		for _, channel := range channels {
			ch <- channel
		}
		<-ctx.Done()
		m.watchersMu.Lock()
		delete(m.watchers, id)
		m.watchersMu.Unlock()
	}()
	return nil
}

func (m *channelManager) notify(channel Channel) {
	m.watchersMu.RLock()
	for _, watcher := range m.watchers {
		watcher <- channel
	}
	m.watchersMu.RUnlock()
}

type Channel interface {
	ID() e2api.ChannelID
	TransactionID() TransactionID
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

func newChannelStream(id e2api.ChannelID, meta e2api.ChannelMeta, transaction Transaction, manager *channelManager) Channel {
	return &channelStream{
		manager:     manager,
		id:          id,
		meta:        meta,
		transaction: transaction,
		openCh:      make(chan error, 1),
		doneCh:      make(chan error, 1),
	}
}

type channelStream struct {
	manager     *channelManager
	id          e2api.ChannelID
	meta        e2api.ChannelMeta
	transaction Transaction
	openCh      chan error
	open        bool
	doneCh      chan error
	done        bool
	mu          sync.RWMutex
}

func (s *channelStream) ID() e2api.ChannelID {
	return s.id
}

func (s *channelStream) Meta() e2api.ChannelMeta {
	return s.meta
}

func (s *channelStream) TransactionID() TransactionID {
	return s.transaction.ID()
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
	return s.transaction.Out()
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
