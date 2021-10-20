// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package broker

import (
	"container/list"
	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
	"github.com/onosproject/onos-lib-go/pkg/errors"
	"io"
	"sync"
)

type AppManager struct {
	sub  *SubscriptionStream
	apps map[e2api.AppID]*AppStream
	mu   sync.RWMutex
}

func (s *AppManager) Get(id e2api.AppID) (*AppStream, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	app, ok := s.apps[id]
	return app, ok
}

func (s *AppManager) List() []*AppStream {
	s.mu.RLock()
	defer s.mu.RUnlock()
	apps := make([]*AppStream, 0, len(s.apps))
	for _, app := range s.apps {
		apps = append(apps, app)
	}
	return apps
}

func (s *AppManager) Create(id e2api.AppID) *AppStream {
	s.mu.RLock()
	app, ok := s.apps[id]
	s.mu.RUnlock()
	if ok {
		return app
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	app, ok = s.apps[id]
	if ok {
		return app
	}
	app = &AppStream{
		SubscriptionStream: s.sub,
		AppID:              id,
		buffer:             list.New(),
		cond:               sync.NewCond(&sync.Mutex{}),
	}
	app.transactions = &TransactionManager{
		app: app,
		chs: make(map[e2api.TransactionID]chan *e2appducontents.Ricindication),
	}
	s.apps[id] = app
	go app.receive()
	return app
}

func (s *AppManager) Close(id e2api.AppID) {
	s.mu.Lock()
	app, ok := s.apps[id]
	delete(s.apps, id)
	s.mu.Unlock()
	if ok {
		app.close()
	}
}

type AppStream struct {
	*SubscriptionStream
	AppID        e2api.AppID
	buffer       *list.List
	cond         *sync.Cond
	closed       bool
	transactions *TransactionManager
}

func (s *AppStream) Transactions() *TransactionManager {
	return s.transactions
}

func (s *AppStream) send(ind *e2appducontents.Ricindication) error {
	s.cond.L.Lock()
	defer s.cond.L.Unlock()
	if s.closed {
		return io.EOF
	}
	if s.buffer.Len() == bufferMaxSize {
		return errors.NewUnavailable("cannot append indication to stream: maximum buffer size has been reached")
	}
	s.buffer.PushBack(ind)
	s.cond.Signal()
	return nil
}

func (s *AppStream) receive() {
	for {
		ind, ok := s.next()
		if !ok {
			s.transactions.mu.RLock()
			for _, ch := range s.transactions.chs {
				close(ch)
			}
			s.transactions.mu.RUnlock()
			break
		}
		s.transactions.mu.RLock()
		for _, ch := range s.transactions.chs {
			ch <- ind
		}
		s.transactions.mu.RUnlock()
	}
}

func (s *AppStream) next() (*e2appducontents.Ricindication, bool) {
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

func (s *AppStream) close() {
	s.cond.L.Lock()
	s.closed = true
	s.cond.Signal()
	s.cond.L.Unlock()
}
