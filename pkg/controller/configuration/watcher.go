// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package configuration

import (
	"context"
	"sync"

	topoapi "github.com/onosproject/onos-api/go/onos/topo"
	e2server "github.com/onosproject/onos-e2t/pkg/southbound/e2ap/server"
	"github.com/onosproject/onos-e2t/pkg/store/rnib"
	"github.com/onosproject/onos-lib-go/pkg/controller"
)

const queueSize = 100

// MgmtConnWatcher  is a management connection watcher
type MgmtConnWatcher struct {
	mgmtConns e2server.MgmtConnManager
	cancel    context.CancelFunc
	mu        sync.Mutex
	connCh    chan *e2server.ManagementConn
}

// Start starts the management connection watcher
func (w *MgmtConnWatcher) Start(ch chan<- controller.ID) error {
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
			log.Debugf("Received management Connection event'%s'", conn.ID)
			ch <- controller.NewID(conn.ID)
		}
		close(ch)
	}()
	return nil
}

// Stop stops the management connection watcher
func (w *MgmtConnWatcher) Stop() {
	w.mu.Lock()
	if w.cancel != nil {
		w.cancel()
		w.cancel = nil
	}
	w.mu.Unlock()
}

// TopoWatcher  is a topology watcher
type TopoWatcher struct {
	rnib      rnib.Store
	mgmtConns e2server.MgmtConnManager
	e2apConns e2server.E2APConnManager
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

	err := w.rnib.Watch(ctx, eventCh, nil)
	if err != nil {
		cancel()
		return err
	}
	w.cancel = cancel

	go func() {
		for event := range eventCh {
			conns, err := w.mgmtConns.List(ctx)
			if err != nil {
				log.Warnf("cannot retrieve the list conns %s", err)
				continue
			}

			if relation, ok := event.Object.Obj.(*topoapi.Object_Relation); ok {
				if relation.Relation.KindID == topoapi.CONTROLS {
					log.Debugf("Received control relation event: %+v", event.Object.ID)
					for _, conn := range conns {
						if conn.E2NodeID == relation.Relation.GetTgtEntityID() {
							ch <- controller.NewID(conn.ID)
						}
					}
				}
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

			if entity, ok := event.Object.Obj.(*topoapi.Object_Entity); ok {
				if entity.Entity.KindID == topoapi.E2T {
					log.Debugf("Received E2T node event: %+v", event.Object.ID)
					for _, conn := range conns {
						ch <- controller.NewID(conn.ID)
					}
				}
				// Enqueue the management connection with matching e2node
				if entity.Entity.KindID == topoapi.E2NODE {
					log.Debugf("Received E2 node event: %+v", event.Object.ID)
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
