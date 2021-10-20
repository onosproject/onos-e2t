// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package broker

import (
	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
	"sync"
)

type TransactionManager struct {
	app *AppStream
	chs map[e2api.TransactionID]chan *e2appducontents.Ricindication
	mu  sync.RWMutex
}

func (s *TransactionManager) Get(id e2api.TransactionID) (*TransactionStream, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	ch, ok := s.chs[id]
	if !ok {
		return nil, false
	}
	return &TransactionStream{
		AppStream:     s.app,
		TransactionID: id,
		C:             ch,
	}, true
}

func (s *TransactionManager) List() []*TransactionStream {
	s.mu.RLock()
	defer s.mu.RUnlock()
	transactions := make([]*TransactionStream, 0, len(s.chs))
	for id, ch := range s.chs {
		transactions = append(transactions, &TransactionStream{
			AppStream:     s.app,
			TransactionID: id,
			C:             ch,
		})
	}
	return transactions
}

func (s *TransactionManager) Create(id e2api.TransactionID) *TransactionStream {
	s.mu.RLock()
	ch, ok := s.chs[id]
	s.mu.RUnlock()
	if ok {
		return &TransactionStream{
			AppStream:     s.app,
			TransactionID: id,
			C:             ch,
		}
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	ch, ok = s.chs[id]
	if ok {
		return &TransactionStream{
			AppStream:     s.app,
			TransactionID: id,
			C:             ch,
		}
	}
	ch = make(chan *e2appducontents.Ricindication)
	s.chs[id] = ch
	return &TransactionStream{
		AppStream:     s.app,
		TransactionID: id,
		C:             ch,
	}
}

func (s *TransactionManager) Close(id e2api.TransactionID) {
	s.mu.Lock()
	transaction, ok := s.chs[id]
	delete(s.chs, id)
	s.mu.Unlock()
	if ok {
		close(transaction)
	}
}

type TransactionStream struct {
	*AppStream
	TransactionID e2api.TransactionID
	C             <-chan *e2appducontents.Ricindication
}
