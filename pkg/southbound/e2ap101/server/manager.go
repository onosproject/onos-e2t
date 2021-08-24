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

type ConnManager interface {
	Get(ctx context.Context, id ConnID) (*E2Conn, error)
	List(ctx context.Context) ([]*E2Conn, error)
	Watch(ctx context.Context, ch chan<- *E2Conn) error
	open(channel *E2Conn)
}

// NewConnManager creates a new connection manager
func NewConnManager() ConnManager {
	mgr := connManager{
		connections: make(map[ConnID]*E2Conn),
		eventCh:     make(chan *E2Conn),
	}
	go mgr.processEvents()
	return &mgr
}

type connManager struct {
	connections map[ConnID]*E2Conn
	channelsMu  sync.RWMutex
	watchers    []chan<- *E2Conn
	watchersMu  sync.RWMutex
	eventCh     chan *E2Conn
}

func (m *connManager) processEvents() {
	for channel := range m.eventCh {
		m.processEvent(channel)
	}
}

func (m *connManager) processEvent(conn *E2Conn) {
	log.Info("Notifying connection")
	m.watchersMu.RLock()
	for _, watcher := range m.watchers {
		watcher <- conn
	}
	m.watchersMu.RUnlock()
}

func (m *connManager) open(conn *E2Conn) {
	log.Infof("Opened connection %s", conn.ID)
	m.channelsMu.Lock()
	defer m.channelsMu.Unlock()
	m.connections[conn.ID] = conn
	m.eventCh <- conn
	go func() {
		<-conn.Context().Done()
		log.Infof("Closing connection %s", conn.ID)
		m.channelsMu.Lock()
		delete(m.connections, conn.ID)
		m.channelsMu.Unlock()
		m.eventCh <- conn
	}()
}

// Get gets a connection by ID
func (m *connManager) Get(ctx context.Context, connID ConnID) (*E2Conn, error) {
	m.channelsMu.RLock()
	defer m.channelsMu.RUnlock()
	conn, ok := m.connections[connID]
	if !ok {
		return nil, errors.NewNotFound("connection '%s' not found", connID)
	}
	return conn, nil
}

// List lists channels
func (m *connManager) List(ctx context.Context) ([]*E2Conn, error) {
	m.channelsMu.RLock()
	defer m.channelsMu.RUnlock()
	connections := make([]*E2Conn, 0, len(m.connections))
	for _, conn := range m.connections {
		connections = append(connections, conn)
	}
	return connections, nil
}

// Watch watches for new channels
func (m *connManager) Watch(ctx context.Context, ch chan<- *E2Conn) error {
	m.watchersMu.Lock()
	m.channelsMu.Lock()
	m.watchers = append(m.watchers, ch)
	m.watchersMu.Unlock()

	go func() {
		for _, stream := range m.connections {
			ch <- stream
		}
		m.channelsMu.Unlock()

		<-ctx.Done()
		m.watchersMu.Lock()
		watchers := make([]chan<- *E2Conn, 0, len(m.watchers)-1)
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
func (m *connManager) Close() error {
	close(m.eventCh)
	return nil
}

var _ ConnManager = &connManager{}
