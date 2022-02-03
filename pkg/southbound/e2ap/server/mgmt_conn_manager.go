// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package server

import (
	"context"
	"sync"

	"github.com/onosproject/onos-lib-go/pkg/errors"
)

// MgmtConnManager management connection manager interface
type MgmtConnManager interface {
	Get(ctx context.Context, id ConnID) (*ManagementConn, error)
	List(ctx context.Context) ([]*ManagementConn, error)
	Watch(ctx context.Context, ch chan<- *ManagementConn) error
	open(conn *ManagementConn)
}

// NewMgmtConnManager creates a new management connection manager
func NewMgmtConnManager() MgmtConnManager {
	mgr := &mgmtConnManager{
		conns:   make(map[ConnID]*ManagementConn),
		eventCh: make(chan *ManagementConn),
	}
	go mgr.processEvents()
	return mgr
}

type mgmtConnManager struct {
	conns      map[ConnID]*ManagementConn
	connsMu    sync.RWMutex
	watchers   []chan<- *ManagementConn
	watchersMu sync.RWMutex
	eventCh    chan *ManagementConn
}

func (m *mgmtConnManager) processEvents() {
	for conn := range m.eventCh {
		m.processEvent(conn)
	}
}

func (m *mgmtConnManager) processEvent(conn *ManagementConn) {
	log.Infof("Notifying management connection:", conn.ID)
	m.watchersMu.RLock()
	for _, watcher := range m.watchers {
		watcher <- conn
	}
	m.watchersMu.RUnlock()
}

func (m *mgmtConnManager) open(conn *ManagementConn) {
	log.Infof("Opened management connection %s", conn.ID)
	m.connsMu.Lock()
	m.conns[conn.ID] = conn
	m.connsMu.Unlock()
	m.eventCh <- conn
	go func() {
		<-conn.Context().Done()
		log.Infof("Closing management connection %s", conn.ID)
		m.connsMu.Lock()
		delete(m.conns, conn.ID)
		m.connsMu.Unlock()
		m.eventCh <- conn
	}()
}

// Get gets a connection by ID
func (m *mgmtConnManager) Get(ctx context.Context, connID ConnID) (*ManagementConn, error) {
	m.connsMu.RLock()
	defer m.connsMu.RUnlock()
	conn, ok := m.conns[connID]
	if !ok {
		return nil, errors.NewNotFound("management connection '%s' not found", connID)
	}
	return conn, nil
}

// List lists connections
func (m *mgmtConnManager) List(ctx context.Context) ([]*ManagementConn, error) {
	m.connsMu.RLock()
	defer m.connsMu.RUnlock()
	conns := make([]*ManagementConn, 0, len(m.conns))
	for _, conn := range m.conns {
		conns = append(conns, conn)
	}
	return conns, nil
}

// Watch watches for new connections
func (m *mgmtConnManager) Watch(ctx context.Context, ch chan<- *ManagementConn) error {
	m.watchersMu.Lock()
	m.connsMu.Lock()
	m.watchers = append(m.watchers, ch)
	m.watchersMu.Unlock()

	go func() {
		for _, stream := range m.conns {
			ch <- stream
		}
		m.connsMu.Unlock()

		<-ctx.Done()
		m.watchersMu.Lock()
		watchers := make([]chan<- *ManagementConn, 0, len(m.watchers)-1)
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
func (m *mgmtConnManager) Close() error {
	close(m.eventCh)
	return nil
}

var _ E2APConnManager = &e2apConnManager{}
