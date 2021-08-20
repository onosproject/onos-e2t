// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package channel

import (
	"context"
	"sync"

	"github.com/onosproject/onos-e2t/pkg/controller/utils"

	topoapi "github.com/onosproject/onos-api/go/onos/topo"
	"github.com/onosproject/onos-e2t/pkg/store/rnib"

	e2server "github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/server"
	"github.com/onosproject/onos-lib-go/pkg/controller"
)

const queueSize = 100

// ChannelWatcher is a channel watcher
type Watcher struct {
	channels  e2server.ChannelManager
	cancel    context.CancelFunc
	mu        sync.Mutex
	channelCh chan *e2server.E2Channel
}

// Start starts the channel watcher
func (w *Watcher) Start(ch chan<- controller.ID) error {
	w.mu.Lock()
	defer w.mu.Unlock()
	if w.cancel != nil {
		return nil
	}

	w.channelCh = make(chan *e2server.E2Channel, queueSize)
	ctx, cancel := context.WithCancel(context.Background())
	err := w.channels.Watch(ctx, w.channelCh)
	if err != nil {
		cancel()
		return err
	}
	w.cancel = cancel

	go func() {
		for channel := range w.channelCh {
			log.Debugf("Received Channel event '%s'", channel.ID)
			ch <- controller.NewID(channel.ID)
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
	topo     rnib.Store
	channels e2server.ChannelManager
	cancel   context.CancelFunc
	mu       sync.Mutex
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
			channels, err := w.channels.List(ctx)
			if err != nil {
				log.Warnf("cannot retrieve the list channels %v:%v", channels, err)
				continue
			}
			if entity, ok := event.Object.Obj.(*topoapi.Object_Entity); ok {

				// Enqueue the channel with matching e2node
				if entity.Entity.KindID == topoapi.E2NODE {
					for _, channel := range channels {
						if channel.E2NodeID == event.Object.GetID() {
							ch <- controller.NewID(channel.ID)
						}
					}
				}
				// Enqueue the channels with matching cells
				if entity.Entity.KindID == topoapi.E2CELL {
					for _, channel := range channels {
						e2Cells := channel.E2Cells
						for _, e2Cell := range e2Cells {
							cellID := utils.GetCellID(channel, e2Cell)
							if cellID == event.Object.GetID() {
								ch <- controller.NewID(channel.ID)
							}
						}
					}
				}
			}
			if relation, ok := event.Object.Obj.(*topoapi.Object_Relation); ok {
				if relation.Relation.KindID == topoapi.CONTAINS {
					for _, channel := range channels {
						e2Cells := channel.E2Cells
						for _, e2Cell := range e2Cells {
							cellID := utils.GetCellID(channel, e2Cell)
							if cellID == relation.Relation.TgtEntityID {
								ch <- controller.NewID(channel.ID)
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
