// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package e2relation

import (
	"context"
	"sync"

	e2server "github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/server"
	"github.com/onosproject/onos-lib-go/pkg/controller"
)

const queueSize = 100

// ChannelWatcher is a channel watcher
type ChannelWatcher struct {
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
		for channel := range w.channelCh {
			log.Debugf("Received Channel event '%s'", channel.ID)
			ch <- controller.NewID(channel.ID)
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
