// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package stream

import (
	"context"
	"github.com/google/uuid"
	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	"sync"
)

type Manager interface {
	Get(streamID ID) (Subscription, bool)
	Open(id e2api.SubscriptionID) Subscription
	Watch(ctx context.Context, ch chan<- Subscription) error
}

func NewManager() (Manager, error) {
	manager := &subscriptionManager{
		streams:  make(map[ID]Subscription),
		subs:     make(map[e2api.SubscriptionID]Subscription),
		watchers: make(map[uuid.UUID]chan<- Subscription),
	}
	if err := manager.open(); err != nil {
		return nil, err
	}
	return manager, nil
}

type subscriptionManager struct {
	streamID   ID
	streams    map[ID]Subscription
	subs       map[e2api.SubscriptionID]Subscription
	subsMu     sync.RWMutex
	watchers   map[uuid.UUID]chan<- Subscription
	watchersMu sync.RWMutex
}

func (m *subscriptionManager) open() error {
	return nil
}

func (m *subscriptionManager) Get(streamID ID) (Subscription, bool) {
	m.subsMu.RLock()
	defer m.subsMu.RUnlock()
	channel, ok := m.streams[streamID]
	return channel, ok
}

func (m *subscriptionManager) Open(id e2api.SubscriptionID) Subscription {
	m.subsMu.Lock()
	defer m.subsMu.Unlock()

	stream, ok := m.subs[id]
	if !ok {
		m.streamID++
		stream = newSubscriptionStream(id, m.streamID, m)
		m.subs[id] = stream
		m.streams[m.streamID] = stream
		go m.notify(stream)
	}
	return stream
}

func (m *subscriptionManager) close(stream Subscription) {
	m.subsMu.Lock()
	defer m.subsMu.Unlock()
	delete(m.subs, stream.ID())
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
