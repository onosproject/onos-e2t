// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package channel

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/subscription"
	"sync"
)

type Manager interface {
	Get(channelID e2api.ChannelID) (Channel, bool)
	Open(channel *e2api.Channel) Channel
	Watch(ctx context.Context, ch chan<- Channel) error
}

func NewManager(subs subscription.Manager) (Manager, error) {
	manager := &channelManager{
		subs:        subs,
		chanStreams: make(map[e2api.ChannelID]Channel),
		buffers:     make(map[BufferID]Buffer),
		bufferChans: make(map[BufferID]map[e2api.ChannelID]bool),
		subBuffers:  make(map[e2api.SubscriptionID]map[BufferID]bool),
		watchers:    make(map[uuid.UUID]chan<- Channel),
	}
	if err := manager.open(); err != nil {
		return nil, err
	}
	return manager, nil
}

type channelManager struct {
	subs        subscription.Manager
	chanStreams map[e2api.ChannelID]Channel
	buffers     map[BufferID]Buffer
	bufferChans map[BufferID]map[e2api.ChannelID]bool
	subBuffers  map[e2api.SubscriptionID]map[BufferID]bool
	streamsMu   sync.RWMutex
	watchers    map[uuid.UUID]chan<- Channel
	watchersMu  sync.RWMutex
	cancel      context.CancelFunc
}

func (m *channelManager) open() error {
	ctx, cancel := context.WithCancel(context.Background())
	m.cancel = cancel
	subsCh := make(chan subscription.Subscription)
	if err := m.subs.Watch(ctx, subsCh); err != nil {
		return err
	}
	go func() {
		for subStream := range subsCh {
			go func(subStream subscription.Subscription) {
				m.propagateStream(subStream)
			}(subStream)
		}
	}()
	return nil
}

func (m *channelManager) propagateStream(subStream subscription.Subscription) {
	for ind := range subStream.Out() {
		m.propagateIndication(subStream, ind)
	}
}

func (m *channelManager) propagateIndication(subStream subscription.Subscription, ind *e2appducontents.Ricindication) {
	m.streamsMu.RLock()
	defer m.streamsMu.RUnlock()
	buffers := m.subBuffers[subStream.Subscription().ID]
	for bufferID := range buffers {
		if buffer, ok := m.buffers[bufferID]; ok {
			buffer.In() <- ind
		}
	}
}

func (m *channelManager) Get(channelID e2api.ChannelID) (Channel, bool) {
	m.streamsMu.RLock()
	defer m.streamsMu.RUnlock()
	channel, ok := m.chanStreams[channelID]
	return channel, ok
}

func (m *channelManager) Open(channel *e2api.Channel) Channel {
	m.streamsMu.Lock()
	defer m.streamsMu.Unlock()

	stream, ok := m.chanStreams[channel.ID]
	if ok {
		return stream
	}

	bufferID := BufferID(fmt.Sprintf("%s:%s:%s", channel.E2NodeID, channel.AppID, channel.TransactionID))
	buffer, ok := m.buffers[bufferID]
	if !ok {
		buffer = newChannelBuffer(bufferID)
		m.buffers[bufferID] = buffer
	}

	stream = newChannelStream(channel, buffer, m)
	m.chanStreams[channel.ID] = stream

	bufferChans, ok := m.bufferChans[bufferID]
	if !ok {
		bufferChans = make(map[e2api.ChannelID]bool)
		m.bufferChans[bufferID] = bufferChans
	}
	bufferChans[channel.ID] = true

	subBuffers, ok := m.subBuffers[channel.SubscriptionID]
	if !ok {
		subBuffers = make(map[BufferID]bool)
		m.subBuffers[channel.SubscriptionID] = subBuffers
	}
	subBuffers[bufferID] = true

	go m.notify(stream)
	return stream
}

func (m *channelManager) close(stream Channel) {
	m.streamsMu.Lock()
	defer m.streamsMu.Unlock()

	delete(m.chanStreams, stream.Channel().ID)

	bufferID := BufferID(fmt.Sprintf("%s:%s:%s", stream.Channel().E2NodeID, stream.Channel().AppID, stream.Channel().TransactionID))
	bufferChans := m.bufferChans[bufferID]
	delete(bufferChans, stream.Channel().ID)
	if len(bufferChans) == 0 {
		if buffer, ok := m.buffers[bufferID]; ok {
			buffer.Close()
			delete(m.buffers, bufferID)
		}
		delete(m.bufferChans, bufferID)
		for _, subBuffers := range m.subBuffers {
			delete(subBuffers, bufferID)
		}
	}
	go m.notify(stream)
}

func (m *channelManager) Watch(ctx context.Context, ch chan<- Channel) error {
	m.watchersMu.Lock()
	id := uuid.New()
	m.watchers[id] = ch
	m.watchersMu.Unlock()

	m.streamsMu.RLock()
	streams := make([]Channel, 0, len(m.chanStreams))
	for _, stream := range m.chanStreams {
		streams = append(streams, stream)
	}
	m.streamsMu.RUnlock()

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

func (m *channelManager) notify(stream Channel) {
	m.watchersMu.RLock()
	for _, watcher := range m.watchers {
		watcher <- stream
	}
	m.watchersMu.RUnlock()
}

func (m *channelManager) Close() error {
	m.cancel()
	return nil
}
