// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package controlrelation

import (
	"context"
	"sync"

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
			log.Debugf("Received Connection event '%s'", conn.ID)
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
