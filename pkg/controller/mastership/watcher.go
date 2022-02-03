// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package mastership

import (
	"context"
	"sync"

	topoapi "github.com/onosproject/onos-api/go/onos/topo"
	"github.com/onosproject/onos-e2t/pkg/store/rnib"

	"github.com/onosproject/onos-lib-go/pkg/controller"
)

const queueSize = 100

// Watcher is a topology watcher
type TopoWatcher struct {
	topo   rnib.Store
	cancel context.CancelFunc
	mu     sync.Mutex
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
			if relation, ok := event.Object.Obj.(*topoapi.Object_Relation); ok &&
				relation.Relation.KindID == topoapi.CONTROLS {
				ch <- controller.NewID(relation.Relation.TgtEntityID)
			}
			if entity, ok := event.Object.Obj.(*topoapi.Object_Entity); ok &&
				entity.Entity.KindID == topoapi.E2NODE {
				ch <- controller.NewID(event.Object.ID)
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
