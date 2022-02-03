// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package controlrelation

import (
	"context"
	"sync"

	topoapi "github.com/onosproject/onos-api/go/onos/topo"
	"github.com/onosproject/onos-e2t/pkg/controller/utils"
	"github.com/onosproject/onos-e2t/pkg/store/rnib"

	e2server "github.com/onosproject/onos-e2t/pkg/southbound/e2ap/server"
	"github.com/onosproject/onos-lib-go/pkg/controller"
)

const queueSize = 100

// E2APConnWatcher is an e2ap conn connection watcher
type E2APConnWatcher struct {
	e2apConns e2server.E2APConnManager
	cancel    context.CancelFunc
	mu        sync.Mutex
	connCh    chan *e2server.E2APConn
}

// Start starts the connection watcher
func (w *E2APConnWatcher) Start(ch chan<- controller.ID) error {
	w.mu.Lock()
	defer w.mu.Unlock()
	if w.cancel != nil {
		return nil
	}

	w.connCh = make(chan *e2server.E2APConn, queueSize)
	ctx, cancel := context.WithCancel(context.Background())
	err := w.e2apConns.Watch(ctx, w.connCh)
	if err != nil {
		cancel()
		return err
	}
	w.cancel = cancel

	go func() {
		for conn := range w.connCh {
			log.Debugf("Received  E2AP Connection event '%s'", conn.ID)
			ch <- controller.NewID(conn.ID)
		}
		close(ch)
	}()
	return nil
}

// Stop stops the connection watcher
func (w *E2APConnWatcher) Stop() {
	w.mu.Lock()
	if w.cancel != nil {
		w.cancel()
		w.cancel = nil
	}
	w.mu.Unlock()
}

// TopoWatcher is a topology watcher
type TopoWatcher struct {
	rnib   rnib.Store
	cancel context.CancelFunc
	mu     sync.Mutex
}

// Start starts the rnib store watcher
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
			if entity, ok := event.Object.Obj.(*topoapi.Object_Entity); ok &&
				entity.Entity.KindID == topoapi.E2T && event.Type == topoapi.EventType_REMOVED {
				log.Debugf("Received E2T topo event '%s'", event.Object.ID)
				controlRelationSrcIDFilter := &topoapi.Filters{
					RelationFilter: &topoapi.RelationFilter{
						RelationKind: topoapi.CONTROLS,
						SrcId:        string(utils.GetE2TID()),
					},
				}

				relations, err := w.rnib.List(ctx, controlRelationSrcIDFilter)
				if err != nil {
					log.Warn(err)
					continue
				}

				for _, relation := range relations {
					ch <- controller.NewID(e2server.ConnID(relation.ID))
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
