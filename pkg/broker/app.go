// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package broker

import (
	"context"
	"github.com/google/uuid"
	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
	"sync"
)

type AppManager struct {
	sub        *SubscriptionStream
	apps       map[e2api.AppID]*AppStream
	appsMu     sync.RWMutex
	watchers   map[uuid.UUID]chan<- e2api.AppID
	watchersMu sync.RWMutex
}

func (s *AppManager) Get(id e2api.AppID) (*AppStream, bool) {
	s.appsMu.RLock()
	defer s.appsMu.RUnlock()
	sub, ok := s.apps[id]
	return sub, ok
}

func (s *AppManager) List() []*AppStream {
	s.appsMu.RLock()
	defer s.appsMu.RUnlock()
	subs := make([]*AppStream, 0, len(s.apps))
	for _, sub := range s.apps {
		subs = append(subs, sub)
	}
	return subs
}

func (s *AppManager) Create(id e2api.AppID) *AppStream {
	s.appsMu.RLock()
	app, ok := s.apps[id]
	s.appsMu.RUnlock()
	if ok {
		return app
	}
	s.appsMu.Lock()
	defer s.appsMu.Unlock()
	app, ok = s.apps[id]
	if ok {
		return app
	}
	ch := make(chan *e2appducontents.Ricindication)
	app = &AppStream{
		SubscriptionStream: s.sub,
		AppID:              id,
		ch:                 ch,
	}
	app.transactions = &TransactionManager{
		app:          app,
		transactions: make(map[e2api.TransactionID]*TransactionStream),
		watchers:     make(map[uuid.UUID]chan<- e2api.TransactionID),
	}
	s.apps[id] = app
	app.open()
	go s.notify(id)
	return app
}

func (s *AppManager) Close(id e2api.AppID) {
	s.appsMu.Lock()
	app, ok := s.apps[id]
	delete(s.apps, id)
	s.appsMu.Unlock()
	if ok {
		close(app.ch)
		go s.notify(id)
	}
}

func (s *AppManager) notify(appID e2api.AppID) {
	s.watchersMu.RLock()
	for _, watcher := range s.watchers {
		watcher <- appID
	}
	s.watchersMu.RUnlock()
}

func (s *AppManager) Watch(ctx context.Context, ch chan<- e2api.AppID) {
	s.watchersMu.Lock()
	id := uuid.New()
	s.watchers[id] = ch
	s.watchersMu.Unlock()

	s.appsMu.RLock()
	appIDs := make([]e2api.AppID, 0, len(s.apps))
	for _, app := range s.apps {
		appIDs = append(appIDs, app.AppID)
	}
	s.appsMu.RUnlock()

	go func() {
		for _, appID := range appIDs {
			ch <- appID
		}
		<-ctx.Done()
		s.watchersMu.Lock()
		delete(s.watchers, id)
		s.watchersMu.Unlock()
	}()
}

func (s *AppManager) send(ind *e2appducontents.Ricindication) {
	s.appsMu.RLock()
	defer s.appsMu.RUnlock()
	for _, app := range s.apps {
		app.send(ind)
	}
}

func (s *AppManager) close() {
	s.appsMu.RLock()
	defer s.appsMu.RUnlock()
	for _, app := range s.apps {
		app.close()
	}
}

type AppStream struct {
	*SubscriptionStream
	AppID        e2api.AppID
	transactions *TransactionManager
	ch           chan *e2appducontents.Ricindication
}

func (s *AppStream) open() {
	go s.receive()
}

func (s *AppStream) send(ind *e2appducontents.Ricindication) {
	s.ch <- ind
}

func (s *AppStream) receive() {
	defer s.transactions.close()
	for ind := range s.ch {
		s.transactions.send(ind)
	}
}

func (s *AppStream) Transactions() *TransactionManager {
	return s.transactions
}

func (s *AppStream) close() {
	close(s.ch)
}
