// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package subscription

import (
	"sync"

	subapi "github.com/onosproject/onos-api/go/onos/e2sub/subscription"
)

// RequestEventType is a request event type
type RequestEventType int

const (
	// RequestEventAdded is a request add event
	RequestEventAdded RequestEventType = iota
	// RequestEventRemoved is a request remove event
	RequestEventRemoved
)

// RequestEvent is a request event
type RequestEvent struct {
	// Type is the request event type
	Type RequestEventType
	// RequestEntry is the associated record
	Record RequestEntry
}

// RequestID is a subscription request identifier
type RequestID int32

// RequestEntry is an entry in the request journal
type RequestEntry struct {
	// RequestID is the record request ID
	RequestID RequestID
	// Subscription is the subscription
	Subscription subapi.Subscription
}

// NewRequestJournal creates a new subscription request journal
func NewRequestJournal() *RequestJournal {
	catalog := &RequestJournal{
		records: make(map[subapi.ID]RequestEntry),
	}
	catalog.open()
	return catalog
}

// RequestJournal is a subscription request journal
type RequestJournal struct {
	records    map[subapi.ID]RequestEntry
	recordsMu  sync.RWMutex
	watchers   []chan<- RequestEvent
	watchersMu sync.RWMutex
	eventCh    chan RequestEvent
}

func (c *RequestJournal) open() {
	c.eventCh = make(chan RequestEvent)
	go func() {
		for event := range c.eventCh {
			log.Infof("Notifying RequestEvent %v", event)
			c.watchersMu.RLock()
			for _, watcher := range c.watchers {
				watcher <- event
			}
			c.watchersMu.RUnlock()
		}
	}()
}

func (c *RequestJournal) Add(id subapi.ID, record RequestEntry) {
	c.recordsMu.Lock()
	defer c.recordsMu.Unlock()
	if _, ok := c.records[id]; !ok {
		log.Infof("Added RequestEntry %v", record)
		c.records[id] = record
		c.eventCh <- RequestEvent{
			Type:   RequestEventAdded,
			Record: record,
		}
	}
}

func (c *RequestJournal) Remove(id subapi.ID) {
	c.recordsMu.Lock()
	defer c.recordsMu.Unlock()
	record, ok := c.records[id]
	if ok {
		log.Infof("Removed RequestEntry %v", record)
		delete(c.records, id)
		c.eventCh <- RequestEvent{
			Type:   RequestEventRemoved,
			Record: record,
		}
	}
}

func (c *RequestJournal) Get(id subapi.ID) RequestEntry {
	c.recordsMu.RLock()
	defer c.recordsMu.RUnlock()
	return c.records[id]
}

func (c *RequestJournal) Watch(ch chan<- RequestEvent) func() {
	c.watchersMu.Lock()
	defer c.watchersMu.Unlock()
	c.watchers = append(c.watchers, ch)
	return func() {
		c.watchersMu.Lock()
		if len(c.watchers) > 0 {
			watchers := make([]chan<- RequestEvent, 0, len(c.watchers)-1)
			for _, watcher := range watchers {
				if watcher != ch {
					watchers = append(watchers, watcher)
				}
			}
			c.watchers = watchers
		}
		c.watchersMu.Unlock()
	}
}

func (c *RequestJournal) Close() error {
	close(c.eventCh)
	return nil
}
