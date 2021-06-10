// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package subscription

import (
	"context"
	api "github.com/onosproject/onos-e2t/api/onos/e2t/store/subscription"
	e2server "github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/server"
	"sync"

	substore "github.com/onosproject/onos-e2t/pkg/store/subscription"
	taskstore "github.com/onosproject/onos-e2t/pkg/store/task"
	"github.com/onosproject/onos-lib-go/pkg/controller"
)

const queueSize = 100

// Watcher is a subscription watcher
type Watcher struct {
	tasks  taskstore.Store
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

	taskCh := make(chan api.TaskEvent, queueSize)
	ctx, cancel := context.WithCancel(context.Background())
	err := w.tasks.Watch(ctx, taskCh)
	if err != nil {
		cancel()
		return err
	}
	w.cancel = cancel

	go func() {
		for event := range taskCh {
			ch <- controller.NewID(event.Task.ID)
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
	tasks     taskstore.Store
	subs      substore.Store
	channels  e2server.ChannelManager
	cancel    context.CancelFunc
	mu        sync.Mutex
	channelCh chan *e2server.E2Channel
}

// Start starts the channel watcher
func (w *ChannelWatcher) Start(ch chan<- controller.ID) error {
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
		for range w.channelCh {
			tasks, err := w.tasks.List(ctx)
			if err != nil {
				log.Error(err)
			} else {
				for _, task := range tasks {
					ch <- controller.NewID(task.ID)
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
