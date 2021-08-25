// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package e2node

import (
	"context"
	"sync"

	"github.com/onosproject/onos-e2t/pkg/southbound/e2conn"

	"github.com/onosproject/onos-e2t/pkg/controller/utils"

	topoapi "github.com/onosproject/onos-api/go/onos/topo"
	"github.com/onosproject/onos-e2t/pkg/store/rnib"

	"github.com/onosproject/onos-lib-go/pkg/controller"
)

const queueSize = 100

// Watcher is a connection watcher
type Watcher struct {
	connections e2conn.ConnManager
	cancel      context.CancelFunc
	mu          sync.Mutex
	connCh      chan e2conn.E2BaseConn
}

// Start starts the connection watcher
func (w *Watcher) Start(ch chan<- controller.ID) error {
	w.mu.Lock()
	defer w.mu.Unlock()
	if w.cancel != nil {
		return nil
	}

	w.connCh = make(chan e2conn.E2BaseConn, queueSize)
	ctx, cancel := context.WithCancel(context.Background())
	err := w.connections.Watch(ctx, w.connCh)
	if err != nil {
		cancel()
		return err
	}
	w.cancel = cancel

	go func() {
		for conn := range w.connCh {
			log.Debugf("Received connection event '%s'", conn.GetID())
			ch <- controller.NewID(conn.GetID())
		}
		close(ch)
	}()
	return nil
}

// Stop stops the channel watcher
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
	topo        rnib.Store
	connections e2conn.ConnManager
	cancel      context.CancelFunc
	mu          sync.Mutex
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
			connections, err := w.connections.List(ctx)
			if err != nil {
				log.Warnf("cannot retrieve the list connections  %v:%v", connections, err)
				continue
			}
			if entity, ok := event.Object.Obj.(*topoapi.Object_Entity); ok {

				// Enqueue the connection with matching e2node
				if entity.Entity.KindID == topoapi.E2NODE {
					for _, conn := range connections {
						if conn.GetE2NodeID() == event.Object.GetID() {
							ch <- controller.NewID(conn.GetID())
						}
					}
				}
				// Enqueue the connections with matching cells
				if entity.Entity.KindID == topoapi.E2CELL {
					for _, conn := range connections {
						e2Cells := conn.GetE2Cells()
						for _, e2Cell := range e2Cells {
							cellID := utils.GetCellID(conn, e2Cell)
							if cellID == event.Object.GetID() {
								ch <- controller.NewID(conn.GetID())
							}
						}
					}
				}
			}
			if relation, ok := event.Object.Obj.(*topoapi.Object_Relation); ok {
				if relation.Relation.KindID == topoapi.CONTAINS {
					for _, conn := range connections {
						e2Cells := conn.GetE2Cells()
						for _, e2Cell := range e2Cells {
							cellID := utils.GetCellID(conn, e2Cell)
							if cellID == relation.Relation.TgtEntityID {
								ch <- controller.NewID(conn.GetID())
							}
						}
					}
				}
				if relation.Relation.KindID == topoapi.CONTROLS {
					for _, conn := range connections {
						relationID := utils.GetE2ControlRelationID(conn.GetID())
						if relationID == event.Object.GetID() {
							ch <- controller.NewID(conn.GetID())
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
