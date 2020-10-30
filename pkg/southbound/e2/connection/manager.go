// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package connection

import "sync"

// NewManager creates a new connection manager
func NewManager() *Manager {
	return &Manager{
		connections: make(map[ID]*Connection),
	}
}

// Manager is a connection manager
type Manager struct {
	connections map[ID]*Connection
	mu          sync.RWMutex
}

// Get gets a connection by ID
func (m *Manager) Get(id ID) *Connection {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.connections[id]
}

// List lists connections
func (m *Manager) List() []*Connection {
	m.mu.RLock()
	defer m.mu.RUnlock()
	conns := make([]*Connection, 0, len(m.connections))
	for _, conn := range m.connections {
		conns = append(conns, conn)
	}
	return conns
}

// Register registers a connection
func (m *Manager) Register(conn *Connection) {
	m.mu.Lock()
	m.connections[conn.ID] = conn
	m.mu.Unlock()
}

// Unregister unregisters a connection
func (m *Manager) Unregister(conn *Connection) {
	m.mu.Lock()
	delete(m.connections, conn.ID)
	m.mu.Unlock()
}
