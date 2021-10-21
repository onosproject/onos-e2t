// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package broker

import (
	"context"
	"github.com/google/uuid"
	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
	"sync"
)

type AppStreamManager interface {
	StreamManager
	Get(id e2api.AppID) (*AppStream, bool)
	List() []*AppStream
	Open(id e2api.AppID) *AppStream
	Watch(ctx context.Context, ch chan<- e2api.AppID)
	close(id e2api.AppID)
}

type appStreamManager struct {
	sub        *SubscriptionStream
	apps       map[e2api.AppID]*AppStream
	appsMu     sync.RWMutex
	watchers   map[uuid.UUID]chan<- e2api.AppID
	watchersMu sync.RWMutex
}

func (s *appStreamManager) Get(id e2api.AppID) (*AppStream, bool) {
	s.appsMu.RLock()
	defer s.appsMu.RUnlock()
	app, ok := s.apps[id]
	return app, ok
}

func (s *appStreamManager) List() []*AppStream {
	s.appsMu.RLock()
	defer s.appsMu.RUnlock()
	apps := make([]*AppStream, 0, len(s.apps))
	for _, app := range s.apps {
		apps = append(apps, app)
	}
	return apps
}

func (s *appStreamManager) Open(id e2api.AppID) *AppStream {
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
		manager:            s,
		AppID:              id,
		ch:                 ch,
	}
	app.transactions = &transactionStreamManager{
		app:          app,
		transactions: make(map[e2api.TransactionID]*TransactionStream),
		watchers:     make(map[uuid.UUID]chan<- e2api.TransactionID),
	}
	s.apps[id] = app
	app.open()
	go s.notify(id)
	return app
}

func (s *appStreamManager) notify(appID e2api.AppID) {
	s.watchersMu.RLock()
	for _, watcher := range s.watchers {
		watcher <- appID
	}
	s.watchersMu.RUnlock()
}

func (s *appStreamManager) Watch(ctx context.Context, ch chan<- e2api.AppID) {
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

func (s *appStreamManager) close(id e2api.AppID) {
	s.appsMu.Lock()
	defer s.appsMu.Unlock()
	if _, ok := s.apps[id]; ok {
		delete(s.apps, id)
		go s.notify(id)
	}
}

type AppStream struct {
	*SubscriptionStream
	manager      AppStreamManager
	AppID        e2api.AppID
	transactions TransactionStreamManager
	ch           chan *e2appducontents.Ricindication
}

func (s *AppStream) Transactions() TransactionStreamManager {
	return s.transactions
}

func (s *AppStream) open() {
	go s.receive()
}

func (s *AppStream) send(ind *e2appducontents.Ricindication) {
	s.ch <- ind
}

func (s *AppStream) receive() {
	for ind := range s.ch {
		for _, transaction := range s.transactions.List() {
			transaction.send(ind)
		}
	}
	for _, transaction := range s.transactions.List() {
		transaction.Close()
	}
}

func (s *AppStream) Close() {
	s.manager.close(s.AppID)
	close(s.ch)
}

var _ Stream = &AppStream{}
