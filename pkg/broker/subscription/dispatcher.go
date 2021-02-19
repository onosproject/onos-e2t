// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package subscription

import (
	subctrl "github.com/onosproject/onos-e2t/pkg/controller/subscription"
	"github.com/onosproject/onos-e2t/pkg/northbound/stream"
	e2server "github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/server"
	"sync"
)

// newDispatcher creates a new subscription dispatcher
func newDispatcher(requests *subctrl.RequestJournal, channel *e2server.E2Channel, streams *stream.Manager) (*Dispatcher, error) {
	dispatcher := &Dispatcher{
		requests:  requests,
		channel:   channel,
		streams:   streams,
		listeners: make(map[ListenerID]*Listener),
	}
	if err := dispatcher.open(); err != nil {
		return nil, err
	}
	return dispatcher, nil
}

// Dispatcher is a subscription dispatcher
type Dispatcher struct {
	requests  *subctrl.RequestJournal
	channel   *e2server.E2Channel
	streams   *stream.Manager
	listeners map[ListenerID]*Listener
	mu        sync.RWMutex
	closeFunc func()
}

// open opens the dispatcher
func (d *Dispatcher) open() error {
	eventCh := make(chan subctrl.RequestEvent)
	d.closeFunc = d.requests.Watch(eventCh)
	go d.processCatalogEvents(eventCh)
	return nil
}

func (d *Dispatcher) processCatalogEvents(eventCh <-chan subctrl.RequestEvent) {
	for event := range eventCh {
		d.processCatalogEvent(event)
	}
}

func (d *Dispatcher) processCatalogEvent(event subctrl.RequestEvent) {
	log.Infof("Received RequestEvent %v", event)
	switch event.Type {
	case subctrl.RequestEventAdded:
		if err := d.processSubscriptionAdded(event.Record); err != nil {
			log.Errorf("Failed to process RequestEvent %v: %v", event, err)
		}
	case subctrl.RequestEventRemoved:
		if err := d.processSubscriptionRemoved(event.Record); err != nil {
			log.Errorf("Failed to process RequestEvent %v: %v", event, err)
		}
	}
}

func (d *Dispatcher) processSubscriptionAdded(record subctrl.RequestEntry) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	listenerID := ListenerID(record.RequestID)
	_, ok := d.listeners[listenerID]
	if !ok {
		log.Infof("Opening request %d listener", record.RequestID)
		l, err := newListener(listenerID, record.Subscription, d.channel, d.streams)
		if err != nil {
			return err
		}
		d.listeners[listenerID] = l
	}
	return nil
}

func (d *Dispatcher) processSubscriptionRemoved(record subctrl.RequestEntry) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	listenerID := ListenerID(record.RequestID)
	listener, ok := d.listeners[listenerID]
	if ok {
		log.Infof("Closing request %d listener", record.RequestID)
		delete(d.listeners, listenerID)
		return listener.Close()
	}
	return nil
}

// Close closes the dispatcher
func (d *Dispatcher) Close() error {
	d.mu.Lock()
	var err error
	for _, listener := range d.listeners {
		if e := listener.Close(); e != nil {
			err = e
		}
	}
	if d.closeFunc != nil {
		d.closeFunc()
	}
	d.mu.Unlock()
	return err
}
