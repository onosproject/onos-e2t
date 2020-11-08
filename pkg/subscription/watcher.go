// Copyright 2019-present Open Networking Foundation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package subscription

import (
	"context"
	"github.com/onosproject/onos-e2t/pkg/controller"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2/channel"
	"github.com/onosproject/onos-e2t/pkg/store/subscription"
	"sync"
)

const queueSize = 100

// Watcher is a subscription watcher
type Watcher struct {
	subs   subscription.Store
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

	subCh := make(chan subscription.Event, queueSize)
	ctx, cancel := context.WithCancel(context.Background())
	err := w.subs.Watch(ctx, subCh)
	if err != nil {
		cancel()
		return err
	}
	w.cancel = cancel

	go func() {
		for request := range subCh {
			ch <- controller.NewID(request.Subscription.ID)
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

// ChannelWatcher is a channel watcher
type ChannelWatcher struct {
	subs     subscription.Store
	channels *channel.Manager
	cancel   context.CancelFunc
	mu       sync.Mutex
}

// Start starts the channel watcher
func (w *ChannelWatcher) Start(ch chan<- controller.ID) error {
	w.mu.Lock()
	defer w.mu.Unlock()
	if w.cancel != nil {
		return nil
	}

	channelCh := make(chan channel.Channel, queueSize)
	ctx, cancel := context.WithCancel(context.Background())
	err := w.channels.Watch(ctx, channelCh)
	if err != nil {
		cancel()
		return err
	}
	w.cancel = cancel

	go func() {
		for c := range channelCh {
			subs, err := w.subs.List(ctx)
			if err == nil {
				for _, sub := range subs {
					if channel.ID(sub.E2NodeID) == c.ID() {
						ch <- controller.NewID(sub.ID)
					}
				}
			}
		}
		close(ch)
	}()
	return nil
}

// Stop stops the channel watcher
func (w *ChannelWatcher) Stop() {
	w.mu.Lock()
	if w.cancel != nil {
		w.cancel()
		w.cancel = nil
	}
	w.mu.Unlock()
}
