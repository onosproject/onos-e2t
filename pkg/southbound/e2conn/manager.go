// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package e2conn

import (
	"context"
	"sync"

	"github.com/onosproject/onos-lib-go/pkg/errors"
)

type ConnManager interface {
	Get(ctx context.Context, id ID) (E2BaseConn, error)
	List(ctx context.Context) ([]E2BaseConn, error)
	Watch(ctx context.Context, ch chan<- E2BaseConn) error
	Open(conn E2BaseConn)
}

// NewConnManager creates a new connection manager
func NewConnManager() ConnManager {
	mgr := connManager{
		connections: make(map[ID]E2BaseConn),
		eventCh:     make(chan E2BaseConn),
	}
	go mgr.processEvents()
	return &mgr
}

type connManager struct {
	connections map[ID]E2BaseConn
	channelsMu  sync.RWMutex
	watchers    []chan<- E2BaseConn
	watchersMu  sync.RWMutex
	eventCh     chan E2BaseConn
}

func (m *connManager) processEvents() {
	for channel := range m.eventCh {
		m.processEvent(channel)
	}
}

func (m *connManager) processEvent(conn E2BaseConn) {
	log.Info("Notifying connection")
	m.watchersMu.RLock()
	for _, watcher := range m.watchers {
		watcher <- conn
	}
	m.watchersMu.RUnlock()
}

func (m *connManager) Open(conn E2BaseConn) {
	log.Infof("Opened connection %s", conn.GetID())
	m.channelsMu.Lock()
	defer m.channelsMu.Unlock()
	m.connections[conn.GetID()] = conn
	m.eventCh <- conn
	go func() {
		<-conn.Context().Done()
		log.Infof("Closing connection %s", conn.GetID())
		m.channelsMu.Lock()
		delete(m.connections, conn.GetID())
		m.channelsMu.Unlock()
		m.eventCh <- conn
	}()
}

// Get gets a connection by ID
func (m *connManager) Get(ctx context.Context, connID ID) (E2BaseConn, error) {
	m.channelsMu.RLock()
	defer m.channelsMu.RUnlock()
	conn, ok := m.connections[connID]
	if !ok {
		return nil, errors.NewNotFound("connection '%s' not found", connID)
	}
	return conn, nil
}

// List lists channels
func (m *connManager) List(ctx context.Context) ([]E2BaseConn, error) {
	m.channelsMu.RLock()
	defer m.channelsMu.RUnlock()
	connections := make([]E2BaseConn, 0, len(m.connections))
	for _, conn := range m.connections {
		connections = append(connections, conn)
	}
	return connections, nil
}

// Watch watches for new channels
func (m *connManager) Watch(ctx context.Context, ch chan<- E2BaseConn) error {
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
		watchers := make([]chan<- E2BaseConn, 0, len(m.watchers)-1)
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
