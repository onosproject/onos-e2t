// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package subscription

import (
	subapi "github.com/onosproject/onos-e2sub/api/e2/subscription/v1beta1"
	"sync"
)

type CatalogEventType int

const (
	// CatalogEventAdded is a catalog add event
	CatalogEventAdded CatalogEventType = iota
	// CatalogEventRemoved is a catalog remove event
	CatalogEventRemoved
)

// CatalogEvent is a catalog event
type CatalogEvent struct {
	// Type is the catalog event type
	Type CatalogEventType
	// CatalogRecord is the associated record
	Record CatalogRecord
}

// RequestID is a subscription request identifier
type RequestID int32

// CatalogRecord is a record in the catalog
type CatalogRecord struct {
	// RequestID is the record request ID
	RequestID RequestID
	// Subscription is the subscription
	Subscription subapi.Subscription
}

// NewCatalog creates a new subscription catalog
func NewCatalog() *Catalog {
	catalog := &Catalog{
		records: make(map[subapi.ID]CatalogRecord),
	}
	catalog.open()
	return catalog
}

// Catalog is a subscription catalog
type Catalog struct {
	records    map[subapi.ID]CatalogRecord
	recordsMu  sync.RWMutex
	watchers   []chan<- CatalogEvent
	watchersMu sync.RWMutex
	eventCh    chan CatalogEvent
}

func (c *Catalog) open() {
	c.eventCh = make(chan CatalogEvent)
	go func() {
		for event := range c.eventCh {
			log.Infof("Notifying CatalogEvent %v", event)
			c.watchersMu.RLock()
			for _, watcher := range c.watchers {
				watcher <- event
			}
			c.watchersMu.RUnlock()
		}
	}()
}

func (c *Catalog) Add(id subapi.ID, record CatalogRecord) {
	c.recordsMu.Lock()
	defer c.recordsMu.Unlock()
	if _, ok := c.records[id]; !ok {
		c.records[id] = record
		c.eventCh <- CatalogEvent{
			Type:   CatalogEventAdded,
			Record: record,
		}
	}
}

func (c *Catalog) Remove(id subapi.ID) {
	c.recordsMu.Lock()
	defer c.recordsMu.Unlock()
	record, ok := c.records[id]
	if ok {
		delete(c.records, id)
		c.eventCh <- CatalogEvent{
			Type:   CatalogEventAdded,
			Record: record,
		}
	}
}

func (c *Catalog) Get(id subapi.ID) CatalogRecord {
	c.recordsMu.RLock()
	defer c.recordsMu.RUnlock()
	return c.records[id]
}

func (c *Catalog) Watch(ch chan<- CatalogEvent) func() {
	c.watchersMu.Lock()
	defer c.watchersMu.Unlock()
	c.watchers = append(c.watchers, ch)
	return func() {
		c.watchersMu.Lock()
		watchers := make([]chan<- CatalogEvent, 0, len(c.watchers)-1)
		for _, watcher := range watchers {
			if watcher != ch {
				watchers = append(watchers, watcher)
			}
		}
		c.watchers = watchers
		c.watchersMu.Unlock()
	}
}

func (c *Catalog) Close() error {
	close(c.eventCh)
	return nil
}
