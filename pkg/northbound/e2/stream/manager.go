// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package stream

import (
	"context"
	"github.com/google/uuid"
	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/stream"
	"sync"
)

type Manager interface {
	Get(channelID e2api.ChannelID) (Channel, bool)
	Open(channelID e2api.ChannelID, meta e2api.ChannelMeta) Channel
	Watch(ctx context.Context, ch chan<- Channel) error
}

func NewManager(streams stream.Manager) (Manager, error) {
	manager := &channelManager{
		streams:     streams,
		channels:    make(map[e2api.ChannelID]Channel),
		buffers:     make(map[bufferID]*channelBuffer),
		subBuffers:  make(map[e2api.SubscriptionID]map[bufferID]bool),
		bufferChans: make(map[bufferID]map[e2api.ChannelID]bool),
		watchers:    make(map[uuid.UUID]chan<- Channel),
	}
	if err := manager.open(); err != nil {
		return nil, err
	}
	return manager, nil
}

type channelManager struct {
	streams     stream.Manager
	channels    map[e2api.ChannelID]Channel
	channelsMu  sync.RWMutex
	buffers     map[bufferID]*channelBuffer
	subBuffers  map[e2api.SubscriptionID]map[bufferID]bool
	bufferChans map[bufferID]map[e2api.ChannelID]bool
	watchers    map[uuid.UUID]chan<- Channel
	watchersMu  sync.RWMutex
	cancel      context.CancelFunc
}

func (m *channelManager) open() error {
	ctx, cancel := context.WithCancel(context.Background())
	m.cancel = cancel
	subsCh := make(chan stream.Subscription)
	if err := m.streams.Watch(ctx, subsCh); err != nil {
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

func (m *channelManager) propagateStream(subStream stream.Subscription) {
	for ind := range subStream.Out() {
		m.propagateIndication(subStream, ind)
	}
}

func (m *channelManager) propagateIndication(sub stream.Subscription, ind *e2appducontents.Ricindication) {
	m.channelsMu.RLock()
	defer m.channelsMu.RUnlock()
	if subBuffers, ok := m.subBuffers[sub.ID()]; ok {
		for bufferID := range subBuffers {
			if buffer, ok := m.buffers[bufferID]; ok {
				buffer.in <- ind
			}
		}
	}
}

func (m *channelManager) Get(channelID e2api.ChannelID) (Channel, bool) {
	m.channelsMu.RLock()
	defer m.channelsMu.RUnlock()
	c, ok := m.channels[channelID]
	return c, ok
}

func (m *channelManager) Open(channelID e2api.ChannelID, meta e2api.ChannelMeta) Channel {
	m.channelsMu.Lock()
	defer m.channelsMu.Unlock()

	// Get or create a buffer for the transaction
	bufID := getBufferID(meta)
	buf, ok := m.buffers[bufID]
	if !ok {
		buf = newChannelBuffer(bufID)
		m.buffers[bufID] = buf
	}

	ch, ok := m.channels[channelID]
	if !ok {
		ch = newChannel(channelID, meta, buf, m)
		m.channels[channelID] = ch
		go m.notify(ch)
	}

	subBuffers, ok := m.subBuffers[meta.SubscriptionID]
	if !ok {
		subBuffers = make(map[bufferID]bool)
		m.subBuffers[meta.SubscriptionID] = subBuffers
	}
	subBuffers[bufID] = true

	bufferChans, ok := m.bufferChans[bufID]
	if !ok {
		bufferChans = make(map[e2api.ChannelID]bool)
		m.bufferChans[bufID] = bufferChans
	}
	bufferChans[channelID] = true
	return ch
}

func (m *channelManager) close(c Channel) {
	m.channelsMu.Lock()
	defer m.channelsMu.Unlock()

	if _, ok := m.channels[c.ID()]; !ok {
		return
	}

	delete(m.channels, c.ID())

	// Delete the channel reference for the buffer
	// If there are no more channels open for the buffer, delete the buffer references
	bufID := getBufferID(c.Meta())
	if bufChans, ok := m.bufferChans[bufID]; ok {
		delete(bufChans, c.ID())
		if len(bufChans) == 0 {
			delete(m.bufferChans, bufID)

			// Close the underlying buffer
			if buf, ok := m.buffers[bufID]; ok {
				delete(m.buffers, bufID)
				buf.Close()
			}

			// Delete the buffer reference from the subscription
			if subBufs, ok := m.subBuffers[c.Meta().SubscriptionID]; ok {
				delete(subBufs, bufID)
				if len(subBufs) == 0 {
					delete(m.subBuffers, c.Meta().SubscriptionID)
				}
			}
		}
	}
	go m.notify(c)
}

func (m *channelManager) Watch(ctx context.Context, ch chan<- Channel) error {
	m.watchersMu.Lock()
	id := uuid.New()
	m.watchers[id] = ch
	m.watchersMu.Unlock()

	m.channelsMu.RLock()
	channels := make([]Channel, 0, len(m.channels))
	for _, channel := range m.channels {
		channels = append(channels, channel)
	}
	m.channelsMu.RUnlock()

	go func() {
		for _, channel := range channels {
			ch <- channel
		}
		<-ctx.Done()
		m.watchersMu.Lock()
		delete(m.watchers, id)
		m.watchersMu.Unlock()
	}()
	return nil
}

func (m *channelManager) notify(channel Channel) {
	m.watchersMu.RLock()
	for _, watcher := range m.watchers {
		watcher <- channel
	}
	m.watchersMu.RUnlock()
}

func (m *channelManager) Close() error {
	m.cancel()
	return nil
}
