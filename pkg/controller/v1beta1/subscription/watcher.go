// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package subscription

import (
	"context"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/stream"
	"sync"

	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	topoapi "github.com/onosproject/onos-api/go/onos/topo"
	e2server "github.com/onosproject/onos-e2t/pkg/southbound/e2ap/server"
	"github.com/onosproject/onos-e2t/pkg/store/rnib"

	substore "github.com/onosproject/onos-e2t/pkg/store/subscription"
	"github.com/onosproject/onos-lib-go/pkg/controller"
)

const queueSize = 100

// Watcher is a subscription watcher
type Watcher struct {
	subs   substore.Store
	cancel context.CancelFunc
	mu     sync.Mutex
}

// Start starts the subscription watcher
func (w *Watcher) Start(ch chan<- controller.ID) error {
	w.mu.Lock()
	defer w.mu.Unlock()
	if w.cancel != nil {
		return nil
	}

	eventCh := make(chan e2api.SubscriptionEvent, queueSize)
	ctx, cancel := context.WithCancel(context.Background())
	err := w.subs.Watch(ctx, eventCh, substore.WithReplay())
	if err != nil {
		cancel()
		return err
	}
	w.cancel = cancel

	go func() {
		for event := range eventCh {
			ch <- controller.NewID(event.Subscription.ID)
		}
		close(ch)
	}()
	return nil
}

// Stop stops the subscription watcher
func (w *Watcher) Stop() {
	w.mu.Lock()
	if w.cancel != nil {
		w.cancel()
		w.cancel = nil
	}
	w.mu.Unlock()
}

var _ controller.Watcher = &Watcher{}

// ConnWatcher is a connection watcher
type ConnWatcher struct {
	subs   substore.Store
	conns  e2server.E2APConnManager
	cancel context.CancelFunc
	mu     sync.Mutex
	connCh chan *e2server.E2APConn
}

// Start starts the connection watcher
func (w *ConnWatcher) Start(ch chan<- controller.ID) error {
	w.mu.Lock()
	defer w.mu.Unlock()
	if w.cancel != nil {
		return nil
	}

	w.connCh = make(chan *e2server.E2APConn, queueSize)
	ctx, cancel := context.WithCancel(context.Background())
	err := w.conns.Watch(ctx, w.connCh)
	if err != nil {
		cancel()
		return err
	}
	w.cancel = cancel

	go func() {
		for conn := range w.connCh {
			subs, err := w.subs.List(ctx)
			if err != nil {
				log.Error(err)
			} else {
				for _, sub := range subs {
					if topoapi.ID(sub.E2NodeID) == conn.E2NodeID {
						ch <- controller.NewID(sub.ID)
					}
				}
			}
		}
		close(ch)
	}()
	return nil
}

// Stop stops the connection watcher
func (w *ConnWatcher) Stop() {
	w.mu.Lock()
	if w.cancel != nil {
		w.cancel()
		w.cancel = nil
	}
	w.mu.Unlock()
}

// TopoWatcher is a topo watcher
type TopoWatcher struct {
	topo   rnib.Store
	subs   substore.Store
	cancel context.CancelFunc
	mu     sync.Mutex
}

// Start starts the subscription watcher
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
			if entity, ok := event.Object.Obj.(*topoapi.Object_Entity); ok &&
				entity.Entity.KindID == topoapi.E2NODE {
				subs, err := w.subs.List(ctx)
				if err != nil {
					log.Error(err)
				} else {
					for _, sub := range subs {
						if topoapi.ID(sub.E2NodeID) == event.Object.ID {
							ch <- controller.NewID(sub.ID)
						}
					}
				}
			}
		}
		close(ch)
	}()
	return nil
}

// Stop stops the subscription watcher
func (w *TopoWatcher) Stop() {
	w.mu.Lock()
	if w.cancel != nil {
		w.cancel()
		w.cancel = nil
	}
	w.mu.Unlock()
}

var _ controller.Watcher = &TopoWatcher{}

// StreamWatcher is a stream broker watcher
type StreamWatcher struct {
	streams stream.Manager
	cancel  context.CancelFunc
	mu      sync.Mutex
}

// Start starts the subscription watcher
func (w *StreamWatcher) Start(ch chan<- controller.ID) error {
	w.mu.Lock()
	defer w.mu.Unlock()
	if w.cancel != nil {
		return nil
	}

	streamCh := make(chan stream.Subscription, queueSize)
	ctx, cancel := context.WithCancel(context.Background())
	if err := w.streams.Watch(ctx, streamCh); err != nil {
		cancel()
		return err
	}
	w.cancel = cancel

	go func() {
		for stream := range streamCh {
			ch <- controller.NewID(stream.ID())
		}
		close(ch)
	}()
	return nil
}

// Stop stops the subscription watcher
func (w *StreamWatcher) Stop() {
	w.mu.Lock()
	if w.cancel != nil {
		w.cancel()
		w.cancel = nil
	}
	w.mu.Unlock()
}

var _ controller.Watcher = &StreamWatcher{}
