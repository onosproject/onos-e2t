// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package e2node

import (
	"context"
	"sync"

	topoapi "github.com/onosproject/onos-api/go/onos/topo"
	"github.com/onosproject/onos-e2t/pkg/store/rnib"

	e2server "github.com/onosproject/onos-e2t/pkg/southbound/e2ap/server"
	"github.com/onosproject/onos-lib-go/pkg/controller"
)

const queueSize = 100

// Watcher is a connection managements watcher
type Watcher struct {
	mgmtConns e2server.MgmtConnManager
	cancel    context.CancelFunc
	mu        sync.Mutex
	connCh    chan *e2server.ManagementConn
}

// Start starts the connection watcher
func (w *Watcher) Start(ch chan<- controller.ID) error {
	w.mu.Lock()
	defer w.mu.Unlock()
	if w.cancel != nil {
		return nil
	}

	w.connCh = make(chan *e2server.ManagementConn, queueSize)
	ctx, cancel := context.WithCancel(context.Background())
	err := w.mgmtConns.Watch(ctx, w.connCh)
	if err != nil {
		cancel()
		return err
	}
	w.cancel = cancel

	go func() {
		for conn := range w.connCh {
			log.Debugf("Received Connection event '%s'", conn.ID)
			ch <- controller.NewID(conn.ID)
		}
		close(ch)
	}()
	return nil
}

// Stop stops the connection watcher
func (w *Watcher) Stop() {
	w.mu.Lock()
	if w.cancel != nil {
		w.cancel()
		w.cancel = nil
	}
	w.mu.Unlock()
}

//  TopoWatcher is a topology watcher
type TopoWatcher struct {
	topo      rnib.Store
	mgmtConns e2server.MgmtConnManager
	cancel    context.CancelFunc
	mu        sync.Mutex
}

// Start starts the topology watcher
func (w *TopoWatcher) Start(ch chan<- controller.ID) error {
	w.mu.Lock()
	defer w.mu.Unlock()
	if w.cancel != nil {
		return nil
	}

	eventCh := make(chan topoapi.Event, queueSize)
	ctx, cancel := context.WithCancel(context.Background())
	err := w.topo.Watch(ctx, eventCh, nil)
	if err != nil {
		cancel()
		return err
	}
	w.cancel = cancel

	go func() {
		for event := range eventCh {
			log.Debugf("Received topo event '%s'", event.Object.ID)
			conns, err := w.mgmtConns.List(ctx)
			if err != nil {
				log.Warnf("cannot retrieve the list conns %v:%v", conns, err)
				continue
			}
			if entity, ok := event.Object.Obj.(*topoapi.Object_Entity); ok {

				// Enqueue the connection with matching e2node
				if entity.Entity.KindID == topoapi.E2NODE {
					for _, conn := range conns {
						if conn.E2NodeID == event.Object.GetID() {
							ch <- controller.NewID(conn.ID)
						}
					}
				}
				// Enqueue the conns with matching cells
				if entity.Entity.KindID == topoapi.E2CELL {
					for _, conn := range conns {
						e2Cells := conn.E2Cells
						for _, e2Cell := range e2Cells {
							cellID := e2server.GetCellID(conn, e2Cell)
							if cellID == event.Object.GetID() {
								ch <- controller.NewID(conn.ID)
							}
						}
					}
				}
			}
			if relation, ok := event.Object.Obj.(*topoapi.Object_Relation); ok {
				if relation.Relation.KindID == topoapi.CONTAINS {
					for _, conn := range conns {
						e2Cells := conn.E2Cells
						for _, e2Cell := range e2Cells {
							cellID := e2server.GetCellID(conn, e2Cell)
							if cellID == relation.Relation.TgtEntityID {
								ch <- controller.NewID(conn.ID)
							}
						}
					}
				}
			}
		}
		close(ch)
	}()
	return nil
}

// Stop stops the topology watcher
func (w *TopoWatcher) Stop() {
	w.mu.Lock()
	if w.cancel != nil {
		w.cancel()
		w.cancel = nil
	}
	w.mu.Unlock()
}

var _ controller.Watcher = &TopoWatcher{}
