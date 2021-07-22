// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package channel

import (
	"context"
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

// TaskWatcher is a termination endpoint watcher
type TaskWatcher struct {
	chans  chanstore.Store
	subs   substore.Store
	cancel context.CancelFunc
	mu     sync.Mutex
}

// Start starts the channel watcher
func (w *TaskWatcher) Start(ch chan<- controller.ID) error {
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
func (w *TaskWatcher) Stop() {
	w.mu.Lock()
	if w.cancel != nil {
		w.cancel()
		w.cancel = nil
	}
	w.mu.Unlock()
}

var _ controller.Watcher = &TaskWatcher{}
