// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package subscription

import (
	"context"
	"github.com/google/uuid"
	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	"sync"
)

type Manager interface {
	Get(streamID StreamID) (Subscription, bool)
	Open(sub *e2api.Subscription) Subscription
	Watch(ctx context.Context, ch chan<- Subscription) error
}

func NewManager() (Manager, error) {
	manager := &subscriptionManager{
		streams:  make(map[StreamID]Subscription),
		subs:     make(map[e2api.SubscriptionID]Subscription),
		watchers: make(map[uuid.UUID]chan<- Subscription),
	}
	if err := manager.open(); err != nil {
		return nil, err
	}
	return manager, nil
}

type subscriptionManager struct {
	streamID   StreamID
	streams    map[StreamID]Subscription
	subs       map[e2api.SubscriptionID]Subscription
	subsMu     sync.RWMutex
	watchers   map[uuid.UUID]chan<- Subscription
	watchersMu sync.RWMutex
}

func (m *subscriptionManager) open() error {
	return nil
}

func (m *subscriptionManager) Get(streamID StreamID) (Subscription, bool) {
	m.subsMu.RLock()
	defer m.subsMu.RUnlock()
	channel, ok := m.streams[streamID]
	return channel, ok
}

func (m *subscriptionManager) Open(sub *e2api.Subscription) Subscription {
	m.subsMu.Lock()
	defer m.subsMu.Unlock()

	stream, ok := m.subs[sub.ID]
	if ok {
		return stream
	}

	m.streamID++
	stream = newSubscriptionStream(m.streamID, sub, m)
	m.subs[sub.ID] = stream
	m.streams[m.streamID] = stream

	go m.notify(stream)
	return stream
}

func (m *subscriptionManager) close(stream Subscription) {
	m.subsMu.Lock()
	defer m.subsMu.Unlock()
	delete(m.subs, stream.Subscription().ID)
	delete(m.streams, stream.StreamID())
	go m.notify(stream)
}

func (m *subscriptionManager) Watch(ctx context.Context, ch chan<- Subscription) error {
	m.watchersMu.Lock()
	id := uuid.New()
	m.watchers[id] = ch
	m.watchersMu.Unlock()

	m.subsMu.RLock()
	streams := make([]Subscription, 0, len(m.subs))
	for _, stream := range m.subs {
		streams = append(streams, stream)
	}
	m.subsMu.RUnlock()

	go func() {
		for _, stream := range streams {
			ch <- stream
		}
		<-ctx.Done()
		m.watchersMu.Lock()
		delete(m.watchers, id)
		m.watchersMu.Unlock()
	}()
	return nil
}

func (m *subscriptionManager) notify(stream Subscription) {
	m.watchersMu.RLock()
	for _, watcher := range m.watchers {
		watcher <- stream
	}
	m.watchersMu.RUnlock()
}
