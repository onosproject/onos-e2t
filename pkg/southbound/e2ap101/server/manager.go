// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package server

import (
	"context"
	"sync"

	"github.com/onosproject/onos-lib-go/pkg/errors"
	"github.com/onosproject/onos-lib-go/pkg/logging"
)

var log = logging.GetLogger("southbound", "e2ap", "server")

type ChannelManager interface {
	Get(ctx context.Context, id ChannelID) (*E2Channel, error)
	List(ctx context.Context) ([]*E2Channel, error)
	Watch(ctx context.Context, ch chan<- *E2Channel) error
	Open(id ChannelID, channel *E2Channel)
}

// NewChannelManager creates a new channel manager
func NewChannelManager() ChannelManager {
	mgr := channelManager{
		channels: make(map[ChannelID]*E2Channel),
		eventCh:  make(chan *E2Channel),
	}
	go mgr.processEvents()

	return &mgr
}

type ChannelID string

type channelManager struct {
	channels   map[ChannelID]*E2Channel
	channelsMu sync.RWMutex
	watchers   []chan<- *E2Channel
	watchersMu sync.RWMutex
	eventCh    chan *E2Channel
}

func (m *channelManager) processEvents() {
	for channel := range m.eventCh {
		m.processEvent(channel)
	}
}

func (m *channelManager) processEvent(channel *E2Channel) {
	log.Info("Notifying channel")
	m.watchersMu.RLock()
	for _, watcher := range m.watchers {
		watcher <- channel
	}
	m.watchersMu.RUnlock()
}

func (m *channelManager) Open(id ChannelID, channel *E2Channel) {
	log.Infof("Opened channel %s", id)
	m.channelsMu.Lock()
	defer m.channelsMu.Unlock()
	m.channels[id] = channel
	m.eventCh <- channel
	go func() {
		<-channel.Context().Done()
		log.Infof("Closing channel %s", id)
		m.channelsMu.Lock()
		delete(m.channels, id)
		m.channelsMu.Unlock()
	}()
}

// Get gets a channel by ID
func (m *channelManager) Get(ctx context.Context, id ChannelID) (*E2Channel, error) {
	m.channelsMu.RLock()
	defer m.channelsMu.RUnlock()
	channel, ok := m.channels[id]
	if !ok {
		return nil, errors.NewNotFound("channel '%s' not found", id)
	}
	return channel, nil
}

// List lists channels
func (m *channelManager) List(ctx context.Context) ([]*E2Channel, error) {
	m.channelsMu.RLock()
	defer m.channelsMu.RUnlock()
	channels := make([]*E2Channel, 0, len(m.channels))
	for _, channel := range m.channels {
		channels = append(channels, channel)
	}
	return channels, nil
}

// Watch watches for new channels
func (m *channelManager) Watch(ctx context.Context, ch chan<- *E2Channel) error {
	m.watchersMu.Lock()
	m.channelsMu.Lock()
	m.watchers = append(m.watchers, ch)
	m.watchersMu.Unlock()

	go func() {
		for _, stream := range m.channels {
			ch <- stream
		}
		m.channelsMu.Unlock()

		<-ctx.Done()
		m.watchersMu.Lock()
		watchers := make([]chan<- *E2Channel, 0, len(m.watchers)-1)
		for _, watcher := range watchers {
			if watcher != ch {
				watchers = append(watchers, watcher)
			}
		}
		m.watchers = watchers
		m.watchersMu.Unlock()
	}()
	return nil
}

// Close closes the manager
func (m *channelManager) Close() error {
	close(m.eventCh)
	return nil
}

var _ ChannelManager = &channelManager{}
