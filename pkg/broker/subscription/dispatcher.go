// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package subscription

import (
	"context"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2apies"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2appducontents"
	"github.com/onosproject/onos-e2t/pkg/config"
	subctrl "github.com/onosproject/onos-e2t/pkg/controller/subscription"
	"github.com/onosproject/onos-e2t/pkg/northbound/stream"
	e2server "github.com/onosproject/onos-e2t/pkg/southbound/e2ap/server"
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
	closer := d.requests.Watch(eventCh)
	go d.processCatalogEvents(eventCh)

	ricRequestID := e2apies.RicrequestId{
		RicInstanceId: config.InstanceID,
	}
	indCh := make(chan e2appducontents.Ricindication)
	ctx, cancel := context.WithCancel(context.Background())
	d.channel.WatchRICIndications(ctx, ricRequestID, indCh)
	d.closeFunc = func() {
		closer()
		cancel()
	}
	go d.processIndications(indCh)
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
		l, err := newListener(listenerID, record.Subscription, d.streams)
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

// readIndications reads indications from the connection
func (d *Dispatcher) processIndications(ch <-chan e2appducontents.Ricindication) {
	for indication := range ch {
		requestID := indication.ProtocolIes.E2ApProtocolIes29.Value
		listenerID := ListenerID(requestID.RicRequestorId)
		d.mu.RLock()
		listener, ok := d.listeners[listenerID]
		d.mu.RUnlock()
		if ok {
			err := listener.Notify(indication)
			if err != nil {
				log.Errorf("Failed to process indication %+v : %v", indication, err)
			}
		}
	}
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
