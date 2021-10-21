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

type SubscriptionStreamManager interface {
	StreamManager
	Get(id e2api.SubscriptionID) (*SubscriptionStream, bool)
	List() []*SubscriptionStream
	Open(id e2api.SubscriptionID) *SubscriptionStream
	Watch(ctx context.Context, ch chan<- e2api.SubscriptionID)
	close(id e2api.SubscriptionID)
}

type subscriptionStreamManager struct {
	streams    SouthboundStreamManager
	subs       map[e2api.SubscriptionID]*SubscriptionStream
	subsMu     sync.RWMutex
	watchers   map[uuid.UUID]chan<- e2api.SubscriptionID
	watchersMu sync.RWMutex
}

func (s *subscriptionStreamManager) Get(id e2api.SubscriptionID) (*SubscriptionStream, bool) {
	s.subsMu.RLock()
	defer s.subsMu.RUnlock()
	sub, ok := s.subs[id]
	return sub, ok
}

func (s *subscriptionStreamManager) List() []*SubscriptionStream {
	s.subsMu.RLock()
	defer s.subsMu.RUnlock()
	subs := make([]*SubscriptionStream, 0, len(s.subs))
	for _, sub := range s.subs {
		subs = append(subs, sub)
	}
	return subs
}

func (s *subscriptionStreamManager) Open(id e2api.SubscriptionID) *SubscriptionStream {
	s.subsMu.RLock()
	sub, ok := s.subs[id]
	s.subsMu.RUnlock()
	if ok {
		return sub
	}
	s.subsMu.Lock()
	defer s.subsMu.Unlock()
	sub, ok = s.subs[id]
	if ok {
		return sub
	}
	ch := make(chan *e2appducontents.Ricindication)
	sub = &SubscriptionStream{
		SouthboundStream: s.streams.create(ch),
		manager:          s,
		SubscriptionID:   id,
		ch:               ch,
	}
	sub.apps = &appStreamManager{
		sub:      sub,
		apps:     make(map[e2api.AppID]*AppStream),
		watchers: make(map[uuid.UUID]chan<- e2api.AppID),
	}
	s.subs[id] = sub
	sub.open()
	go s.notify(id)
	return sub
}

func (s *subscriptionStreamManager) notify(subID e2api.SubscriptionID) {
	s.watchersMu.RLock()
	for _, watcher := range s.watchers {
		watcher <- subID
	}
	s.watchersMu.RUnlock()
}

func (s *subscriptionStreamManager) Watch(ctx context.Context, ch chan<- e2api.SubscriptionID) {
	s.watchersMu.Lock()
	id := uuid.New()
	s.watchers[id] = ch
	s.watchersMu.Unlock()

	s.subsMu.RLock()
	subIDs := make([]e2api.SubscriptionID, 0, len(s.subs))
	for _, sub := range s.subs {
		subIDs = append(subIDs, sub.SubscriptionID)
	}
	s.subsMu.RUnlock()

	go func() {
		for _, subID := range subIDs {
			ch <- subID
		}
		<-ctx.Done()
		s.watchersMu.Lock()
		delete(s.watchers, id)
		s.watchersMu.Unlock()
	}()
}

func (s *subscriptionStreamManager) close(id e2api.SubscriptionID) {
	s.subsMu.Lock()
	defer s.subsMu.Unlock()
	if _, ok := s.subs[id]; ok {
		delete(s.subs, id)
		go s.notify(id)
	}
}

type SubscriptionStream struct {
	*SouthboundStream
	manager        SubscriptionStreamManager
	SubscriptionID e2api.SubscriptionID
	apps           AppStreamManager
	ch             <-chan *e2appducontents.Ricindication
}

func (s *SubscriptionStream) open() {
	go s.receive()
}

func (s *SubscriptionStream) receive() {
	for ind := range s.ch {
		for _, app := range s.apps.List() {
			app.send(ind)
		}
	}
	for _, app := range s.apps.List() {
		app.Close()
	}
}

func (s *SubscriptionStream) Apps() AppStreamManager {
	return s.apps
}

func (s *SubscriptionStream) Close() {
	s.manager.close(s.SubscriptionID)
	s.SouthboundStream.Close()
}

var _ Stream = &SubscriptionStream{}
