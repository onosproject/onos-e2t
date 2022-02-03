// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package server

import (
	"context"
	"sync"

	"github.com/onosproject/onos-lib-go/pkg/errors"
	"github.com/onosproject/onos-lib-go/pkg/logging"
)

var log = logging.GetLogger("southbound", "e2ap", "server")

type E2APConnManager interface {
	Get(ctx context.Context, id ConnID) (*E2APConn, error)
	List(ctx context.Context) ([]*E2APConn, error)
	Watch(ctx context.Context, ch chan<- *E2APConn) error
	open(conn *E2APConn)
}

// NewE2APConnManager creates a new E2AP connection manager
func NewE2APConnManager() E2APConnManager {
	mgr := &e2apConnManager{
		conns:   make(map[ConnID]*E2APConn),
		eventCh: make(chan *E2APConn),
	}
	go mgr.processEvents()
	return mgr
}

type e2apConnManager struct {
	conns      map[ConnID]*E2APConn
	connsMu    sync.RWMutex
	watchers   []chan<- *E2APConn
	watchersMu sync.RWMutex
	eventCh    chan *E2APConn
}

func (m *e2apConnManager) processEvents() {
	for conn := range m.eventCh {
		m.processEvent(conn)
	}
}

func (m *e2apConnManager) processEvent(conn *E2APConn) {
	log.Infof("Notifying data connection: %s", conn.ID)
	m.watchersMu.RLock()
	for _, watcher := range m.watchers {
		watcher <- conn
	}
	m.watchersMu.RUnlock()
}

func (m *e2apConnManager) open(conn *E2APConn) {
	log.Infof("Opened data connection %s", conn.ID)
	m.connsMu.Lock()
	m.conns[conn.ID] = conn
	m.connsMu.Unlock()
	m.eventCh <- conn
	go func() {
		<-conn.Context().Done()
		log.Infof("Closing data connection for e2 node %s:%s", conn.ID, conn.E2NodeID)
		m.connsMu.Lock()
		delete(m.conns, conn.ID)
		m.connsMu.Unlock()
		m.eventCh <- conn
	}()
}

// Get gets a connection by ID
func (m *e2apConnManager) Get(ctx context.Context, connID ConnID) (*E2APConn, error) {
	m.connsMu.RLock()
	defer m.connsMu.RUnlock()
	conn, ok := m.conns[connID]
	if !ok {
		return nil, errors.NewNotFound("data connection '%s' not found", connID)
	}
	return conn, nil
}

// List lists connections
func (m *e2apConnManager) List(ctx context.Context) ([]*E2APConn, error) {
	m.connsMu.RLock()
	defer m.connsMu.RUnlock()
	conns := make([]*E2APConn, 0, len(m.conns))
	for _, conn := range m.conns {
		conns = append(conns, conn)
	}
	return conns, nil
}

// Watch watches for new connections
func (m *e2apConnManager) Watch(ctx context.Context, ch chan<- *E2APConn) error {
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
		watchers := make([]chan<- *E2APConn, 0, len(m.watchers)-1)
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
func (m *e2apConnManager) Close() error {
	log.Infof("Closing data connection manager")
	close(m.eventCh)
	return nil
}

var _ E2APConnManager = &e2apConnManager{}
