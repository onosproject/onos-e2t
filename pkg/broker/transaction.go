// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package broker

import (
	"container/list"
	"context"
	"github.com/google/uuid"
	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
	"sync"
)

type TransactionManager interface {
	Get(id e2api.TransactionID) (*TransactionStream, bool)
	List() []*TransactionStream
	Create(id e2api.TransactionID) *TransactionStream
	Close(id e2api.TransactionID)
	Watch(ctx context.Context, ch chan<- e2api.TransactionID)
	send(ind *e2appducontents.Ricindication)
	close()
}

type transactionManager struct {
	app            *AppStream
	transactions   map[e2api.TransactionID]*TransactionStream
	transactionsMu sync.RWMutex
	watchers       map[uuid.UUID]chan<- e2api.TransactionID
	watchersMu     sync.RWMutex
}

func (s *transactionManager) Get(id e2api.TransactionID) (*TransactionStream, bool) {
	s.transactionsMu.RLock()
	defer s.transactionsMu.RUnlock()
	app, ok := s.transactions[id]
	return app, ok
}

func (s *transactionManager) List() []*TransactionStream {
	s.transactionsMu.RLock()
	defer s.transactionsMu.RUnlock()
	apps := make([]*TransactionStream, 0, len(s.transactions))
	for _, app := range s.transactions {
		apps = append(apps, app)
	}
	return apps
}

func (s *transactionManager) Create(id e2api.TransactionID) *TransactionStream {
	s.transactionsMu.RLock()
	transaction, ok := s.transactions[id]
	s.transactionsMu.RUnlock()
	if ok {
		return transaction
	}
	s.transactionsMu.Lock()
	defer s.transactionsMu.Unlock()
	transaction, ok = s.transactions[id]
	if ok {
		return transaction
	}
	ch := make(chan *e2appducontents.Ricindication)
	transaction = &TransactionStream{
		AppStream:     s.app,
		TransactionID: id,
		C:             ch,
		ch:            ch,
		buffer:        list.New(),
		cond:          sync.NewCond(&sync.Mutex{}),
	}
	s.transactions[id] = transaction
	transaction.open()
	go s.notify(id)
	return transaction
}

func (s *transactionManager) Close(id e2api.TransactionID) {
	s.transactionsMu.Lock()
	app, ok := s.transactions[id]
	delete(s.transactions, id)
	s.transactionsMu.Unlock()
	if ok {
		app.close()
		go s.notify(id)
	}
}

func (s *transactionManager) notify(transactionID e2api.TransactionID) {
	s.watchersMu.RLock()
	for _, watcher := range s.watchers {
		watcher <- transactionID
	}
	s.watchersMu.RUnlock()
}

func (s *transactionManager) Watch(ctx context.Context, ch chan<- e2api.TransactionID) {
	s.watchersMu.Lock()
	id := uuid.New()
	s.watchers[id] = ch
	s.watchersMu.Unlock()

	s.transactionsMu.RLock()
	transactionIDs := make([]e2api.TransactionID, 0, len(s.transactions))
	for _, transaction := range s.transactions {
		transactionIDs = append(transactionIDs, transaction.TransactionID)
	}
	s.transactionsMu.RUnlock()

	go func() {
		for _, transactionID := range transactionIDs {
			ch <- transactionID
		}
		<-ctx.Done()
		s.watchersMu.Lock()
		delete(s.watchers, id)
		s.watchersMu.Unlock()
	}()
}

func (s *transactionManager) send(ind *e2appducontents.Ricindication) {
	s.transactionsMu.RLock()
	defer s.transactionsMu.RUnlock()
	for _, transaction := range s.transactions {
		transaction.send(ind)
	}
}

func (s *transactionManager) close() {
	s.transactionsMu.RLock()
	defer s.transactionsMu.RUnlock()
	for _, transaction := range s.transactions {
		transaction.close()
	}
}

type TransactionStream struct {
	*AppStream
	TransactionID e2api.TransactionID
	C             <-chan *e2appducontents.Ricindication
	ch            chan *e2appducontents.Ricindication
	buffer        *list.List
	cond          *sync.Cond
	closed        bool
}

func (s *TransactionStream) open() {
	go s.receive()
}

func (s *TransactionStream) receive() {
	defer close(s.ch)
	for {
		ind, ok := s.next()
		if !ok {
			return
		}
		s.ch <- ind
	}
}

func (s *TransactionStream) send(ind *e2appducontents.Ricindication) {
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

func (s *TransactionStream) next() (*e2appducontents.Ricindication, bool) {
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

func (s *TransactionStream) close() {
	s.cond.L.Lock()
	s.closed = true
	s.cond.Signal()
	s.cond.L.Unlock()
}
