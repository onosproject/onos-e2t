// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package stream

import (
	"container/list"
	"context"
	"fmt"
	"github.com/google/uuid"
	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/stream"
	"github.com/prometheus/common/log"
	"sync"
)

const bufferMaxSize = 10000

type TransactionManager interface {
	Get(meta e2api.ChannelMeta) (Transaction, bool)
	Open(meta e2api.ChannelMeta) Transaction
	Watch(ctx context.Context, ch chan<- Transaction) error
}

func newTransactionManager(subs stream.Manager) (TransactionManager, error) {
	manager := &transactionManager{
		subs:            subs,
		transactions:    make(map[TransactionID]Transaction),
		subTransactions: make(map[e2api.SubscriptionID]map[TransactionID]bool),
		watchers:        make(map[uuid.UUID]chan<- Transaction),
	}
	if err := manager.open(); err != nil {
		return nil, err
	}
	return manager, nil
}

type transactionManager struct {
	subs         stream.Manager
	transactions map[TransactionID]Transaction
	subTransactions map[e2api.SubscriptionID]map[TransactionID]bool
	transactionsMu  sync.RWMutex
	watchers        map[uuid.UUID]chan<- Transaction
	watchersMu      sync.RWMutex
	cancel          context.CancelFunc
}

func (m *transactionManager) open() error {
	ctx, cancel := context.WithCancel(context.Background())
	m.cancel = cancel
	subsCh := make(chan stream.Subscription)
	if err := m.subs.Watch(ctx, subsCh); err != nil {
		return err
	}
	go func() {
		for subStream := range subsCh {
			go func(subStream stream.Subscription) {
				m.propagateStream(subStream)
			}(subStream)
		}
	}()
	return nil
}

func (m *transactionManager) propagateStream(subStream stream.Subscription) {
	for ind := range subStream.Out() {
		m.propagateIndication(subStream, ind)
	}
}

func (m *transactionManager) propagateIndication(sub stream.Subscription, ind *e2appducontents.Ricindication) {
	m.transactionsMu.RLock()
	defer m.transactionsMu.RUnlock()
	subTransactions := m.subTransactions[sub.ID()]
	for transactionID := range subTransactions {
		if transaction, ok := m.transactions[transactionID]; ok {
			transaction.In() <- ind
		}
	}
}

func getTransactionID(meta e2api.ChannelMeta) TransactionID {
	return TransactionID(fmt.Sprintf("%s:%s:%s", meta.E2NodeID, meta.AppID, meta.TransactionID))
}

func (m *transactionManager) Get(meta e2api.ChannelMeta) (Transaction, bool) {
	m.transactionsMu.RLock()
	defer m.transactionsMu.RUnlock()
	transaction, ok := m.transactions[getTransactionID(meta)]
	return transaction, ok
}

func (m *transactionManager) Open(meta e2api.ChannelMeta) Transaction {
	m.transactionsMu.Lock()
	defer m.transactionsMu.Unlock()

	transactionID := getTransactionID(meta)
	transaction, ok := m.transactions[transactionID]
	if !ok {
		transaction = newTransactionStream(transactionID, m)
		m.transactions[transactionID] = transaction
		go m.notify(transaction)
	}

	subTransactions, ok := m.subTransactions[meta.SubscriptionID]
	if !ok {
		subTransactions = make(map[TransactionID]bool)
		m.subTransactions[meta.SubscriptionID] = subTransactions
	}
	subTransactions[transactionID] = true
	return transaction
}

func (m *transactionManager) close(transaction Transaction) {
	m.transactionsMu.Lock()
	defer m.transactionsMu.Unlock()
	delete(m.transactions, transaction.ID())
	for subID, subTransactions := range m.subTransactions {
		delete(subTransactions, transaction.ID())
		if len(subTransactions) == 0 {
			delete(m.subTransactions, subID)
		}
	}
	go m.notify(transaction)
}

func (m *transactionManager) Watch(ctx context.Context, ch chan<- Transaction) error {
	m.watchersMu.Lock()
	id := uuid.New()
	m.watchers[id] = ch
	m.watchersMu.Unlock()

	m.transactionsMu.RLock()
	transactions := make([]Transaction, 0, len(m.transactions))
	for _, stream := range m.transactions {
		transactions = append(transactions, stream)
	}
	m.transactionsMu.RUnlock()

	go func() {
		for _, transaction := range transactions {
			ch <- transaction
		}
		<-ctx.Done()
		m.watchersMu.Lock()
		delete(m.watchers, id)
		m.watchersMu.Unlock()
	}()
	return nil
}

func (m *transactionManager) notify(transaction Transaction) {
	m.watchersMu.RLock()
	for _, watcher := range m.watchers {
		watcher <- transaction
	}
	m.watchersMu.RUnlock()
}

func (m *transactionManager) Close() error {
	m.cancel()
	return nil
}

type TransactionID string

type Transaction interface {
	ID() TransactionID
	In() chan<- *e2appducontents.Ricindication
	Out() <-chan *e2appducontents.Ricindication
	Close()
}

func newTransactionStream(id TransactionID, manager *transactionManager) Transaction {
	buffer := &transactionStream{
		manager: manager,
		id:      id,
		in:      make(chan *e2appducontents.Ricindication),
		out:     make(chan *e2appducontents.Ricindication),
		buffer:  list.New(),
		cond:    sync.NewCond(&sync.Mutex{}),
	}
	buffer.open()
	return buffer
}

type transactionStream struct {
	manager *transactionManager
	id      TransactionID
	in      chan *e2appducontents.Ricindication
	out     chan *e2appducontents.Ricindication
	buffer  *list.List
	cond    *sync.Cond
	closed  bool
}

func (s *transactionStream) open() {
	go s.receive()
	go s.send()
}

func (s *transactionStream) ID() TransactionID {
	return s.id
}

func (s *transactionStream) In() chan<- *e2appducontents.Ricindication {
	return s.in
}

func (s *transactionStream) Out() <-chan *e2appducontents.Ricindication {
	return s.out
}

func (s *transactionStream) receive() {
	for ind := range s.in {
		s.write(ind)
	}
}

func (s *transactionStream) write(ind *e2appducontents.Ricindication) {
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

func (s *transactionStream) send() {
	defer close(s.out)
	for {
		ind, ok := s.read()
		if !ok {
			return
		}
		s.out <- ind
	}
}

func (s *transactionStream) read() (*e2appducontents.Ricindication, bool) {
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

func (s *transactionStream) Close() {
	s.manager.close(s)
	close(s.in)
	s.cond.L.Lock()
	s.closed = true
	s.cond.Signal()
	s.cond.L.Unlock()
}
