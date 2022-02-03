// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package channel

import (
	"context"
	topoapi "github.com/onosproject/onos-api/go/onos/topo"
	"github.com/onosproject/onos-e2t/pkg/northbound/e2/stream"
	"github.com/onosproject/onos-e2t/pkg/store/rnib"
	"sync"

	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	chanstore "github.com/onosproject/onos-e2t/pkg/store/channel"
	substore "github.com/onosproject/onos-e2t/pkg/store/subscription"
	"github.com/onosproject/onos-lib-go/pkg/controller"
)

const queueSize = 100

// Watcher is a channel watcher
type Watcher struct {
	chans  chanstore.Store
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

	eventCh := make(chan e2api.ChannelEvent, queueSize)
	ctx, cancel := context.WithCancel(context.Background())
	err := w.chans.Watch(ctx, eventCh, chanstore.WithReplay())
	if err != nil {
		cancel()
		return err
	}
	w.cancel = cancel

	go func() {
		for event := range eventCh {
			ch <- controller.NewID(event.Channel.ID)
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

// SubscriptionWatcher is a subscription watcher
type SubscriptionWatcher struct {
	chans  chanstore.Store
	subs   substore.Store
	cancel context.CancelFunc
	mu     sync.Mutex
}

// Start starts the subscription watcher
func (w *SubscriptionWatcher) Start(ch chan<- controller.ID) error {
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
			chans, err := w.chans.List(ctx)
			if err != nil {
				log.Error(err)
			} else {
				for _, channel := range chans {
					if channel.ChannelMeta.E2NodeID == event.Subscription.SubscriptionMeta.E2NodeID {
						ch <- controller.NewID(channel.ID)
					}
				}
			}
		}
		close(ch)
	}()
	return nil
}

// Stop stops the channel watcher
func (w *SubscriptionWatcher) Stop() {
	w.mu.Lock()
	if w.cancel != nil {
		w.cancel()
		w.cancel = nil
	}
	w.mu.Unlock()
}

var _ controller.Watcher = &SubscriptionWatcher{}

// TopoWatcher is a topo watcher
type TopoWatcher struct {
	topo   rnib.Store
	chans  chanstore.Store
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
				entity.Entity.KindID == topoapi.E2T {
				channels, err := w.chans.List(ctx)
				if err != nil {
					log.Error(err)
				} else {
					for _, channel := range channels {
						ch <- controller.NewID(channel.ID)
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

	streamCh := make(chan stream.Channel, queueSize)
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
