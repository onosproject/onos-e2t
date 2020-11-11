// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package stream

import (
	"context"
	"fmt"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"io"
	"sync"
)

var log = logging.GetLogger("northbound", "stream")

// NewManager creates a new stream manager
func NewManager() *Manager {
	mgr := &Manager{
		streams: make(map[ID]Stream),
		eventCh: make(chan Stream),
	}
	go mgr.processEvents()
	return mgr
}

// Manager is a stream manager
type Manager struct {
	streams    map[ID]Stream
	streamsMu  sync.RWMutex
	watchers   []chan<- Stream
	watchersMu sync.RWMutex
	eventCh    chan Stream
	streamID   ID
}

func (m *Manager) processEvents() {
	for stream := range m.eventCh {
		m.processEvent(stream)
	}
}

func (m *Manager) processEvent(stream Stream) {
	log.Infof("Notifying stream %s", stream.ID())
	m.watchersMu.RLock()
	for _, watcher := range m.watchers {
		watcher <- stream
	}
	m.watchersMu.RUnlock()
}

// Open opens a new stream
func (m *Manager) Open(ctx context.Context, meta Metadata, ch chan Message) (ReadStream, error) {
	m.streamsMu.Lock()
	defer m.streamsMu.Unlock()
	m.streamID++
	stream := newChannelStream(ctx, m.streamID, meta, ch)
	m.streams[m.streamID] = stream
	m.eventCh <- stream
	go func() {
		<-ctx.Done()
		m.streamsMu.Lock()
		delete(m.streams, m.streamID)
		m.streamsMu.Unlock()
	}()
	return stream, nil
}

// Get gets a stream by ID
func (m *Manager) Get(ctx context.Context, id ID) (WriteStream, error) {
	m.streamsMu.RLock()
	defer m.streamsMu.RUnlock()
	stream, ok := m.streams[id]
	if !ok {
		return nil, fmt.Errorf("unknown stream %s", id)
	}
	return stream, nil
}

// Watch watches for streams
func (m *Manager) Watch(ctx context.Context, ch chan<- Stream) error {
	m.watchersMu.Lock()
	m.streamsMu.Lock()
	m.watchers = append(m.watchers, ch)
	m.watchersMu.Unlock()

	go func() {
		for _, stream := range m.streams {
			ch <- stream
		}
		m.streamsMu.Unlock()

		<-ctx.Done()
		m.watchersMu.Lock()
		watchers := make([]chan<- Stream, 0, len(m.watchers)-1)
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
func (m *Manager) Close() error {
	close(m.eventCh)
	return nil
}

var _ io.Closer = &Manager{}
