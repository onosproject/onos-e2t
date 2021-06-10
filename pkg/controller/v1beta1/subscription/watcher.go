// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package subscription

import (
	"context"
	"sync"

	api "github.com/onosproject/onos-e2t/api/onos/e2t/store/subscription"
	"github.com/onosproject/onos-e2t/pkg/store/subscription"
	"github.com/onosproject/onos-e2t/pkg/store/task"
	"github.com/onosproject/onos-lib-go/pkg/controller"
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

	subCh := make(chan api.SubscriptionEvent, queueSize)
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

// TaskWatcher is a termination endpoint watcher
type TaskWatcher struct {
	subs   subscription.Store
	tasks  task.Store
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
			subs, err := w.subs.List(ctx)
			if err != nil {
				log.Error(err)
			} else {
				for _, sub := range subs {
					if sub.ID.NodeID == event.Task.ID.NodeID && sub.ID.RequestID == event.Task.ID.RequestID {
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
func (w *TaskWatcher) Stop() {
	w.mu.Lock()
	if w.cancel != nil {
		w.cancel()
		w.cancel = nil
	}
	w.mu.Unlock()
}

var _ controller.Watcher = &TaskWatcher{}
