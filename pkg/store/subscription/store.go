// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package subscription

import (
	"context"
	"fmt"
	api "github.com/onosproject/onos-e2t/api/ricapi/e2/subscription/v1beta1"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"io"
	"sync"
)

var log = logging.GetLogger("store", "subscription")

// NewStore creates a new subscription store
func NewStore() (Store, error) {
	store := &localStore{
		subscriptions: make(map[api.ID]api.Subscription),
	}
	if err := store.open(); err != nil {
		return nil, err
	}
	return store, nil
}

// Event is a subscription store event
type Event struct {
	// Type is the event type
	Type api.EventType

	// Subscription is the changed subscription
	Subscription api.Subscription
}

// Store is a subscription store
type Store interface {
	io.Closer

	// Add adds a subscription to the store
	Add(ctx context.Context, subscription *api.Subscription) error

	// Update updates a subscription in the store
	Update(ctx context.Context, subscription *api.Subscription) error

	// Remove removes a subscription from the store
	Remove(ctx context.Context, subscription *api.Subscription) error

	// Get gets a subscription by ID
	Get(ctx context.Context, id api.ID) (*api.Subscription, error)

	// List lists the subscriptions in the store
	List(ctx context.Context) ([]api.Subscription, error)

	// Watch watches the store for changes
	Watch(ctx context.Context, ch chan<- Event) error
}

// localStore is a local implementation of the subscription store
type localStore struct {
	subscriptions  map[api.ID]api.Subscription
	subscriptionID api.ID
	mu             sync.RWMutex
	watchers       []chan<- Event
	eventCh        chan Event
	watchMu        sync.RWMutex
}

func (s *localStore) open() error {
	s.eventCh = make(chan Event)
	for event := range s.eventCh {
		log.Infof("Notifying Subscription event %v", event)
		s.mu.RLock()
		for _, watcher := range s.watchers {
			watcher <- event
		}
		s.mu.RUnlock()
	}
	return nil
}

func (s *localStore) Add(ctx context.Context, subscription *api.Subscription) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.subscriptionID++
	id := s.subscriptionID
	subscription.ID = id
	subscription.Revision = 1
	s.subscriptions[id] = *subscription
	log.Infof("Added Subscription %+v", subscription)
	s.eventCh <- Event{
		Type:         api.EventType_ADDED,
		Subscription: *subscription,
	}
	return nil
}

func (s *localStore) Update(ctx context.Context, subscription *api.Subscription) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	stored, ok := s.subscriptions[subscription.ID]
	if !ok {
		return fmt.Errorf("unknown subscription %d", subscription.ID)
	}
	if stored.Revision != subscription.Revision {
		return fmt.Errorf("concurrent update detected")
	}
	subscription.Revision++
	s.subscriptions[subscription.ID] = *subscription
	log.Infof("Updated Subscription %+v", subscription)
	s.eventCh <- Event{
		Type:         api.EventType_UPDATED,
		Subscription: *subscription,
	}
	return nil
}

func (s *localStore) Remove(ctx context.Context, subscription *api.Subscription) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	stored, ok := s.subscriptions[subscription.ID]
	if !ok {
		return fmt.Errorf("unknown subscription %d", subscription.ID)
	}
	if stored.Revision != subscription.Revision {
		return fmt.Errorf("concurrent update detected")
	}
	delete(s.subscriptions, subscription.ID)
	log.Infof("Removed Subscription %v", stored)
	s.eventCh <- Event{
		Type:         api.EventType_REMOVED,
		Subscription: stored,
	}
	return nil
}

func (s *localStore) Get(ctx context.Context, id api.ID) (*api.Subscription, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	subscription, ok := s.subscriptions[id]
	if !ok {
		return nil, fmt.Errorf("unknown subscription %d", id)
	}
	return &subscription, nil
}

func (s *localStore) List(ctx context.Context) ([]api.Subscription, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	subscriptions := make([]api.Subscription, 0, len(s.subscriptions))
	for _, subscription := range s.subscriptions {
		subscriptions = append(subscriptions, subscription)
	}
	return subscriptions, nil
}

func (s *localStore) Watch(ctx context.Context, ch chan<- Event) error {
	s.watchMu.Lock()
	defer s.watchMu.Unlock()
	s.watchers = append(s.watchers, ch)

	go func() {
		<-ctx.Done()
		s.watchMu.Lock()
		watchers := make([]chan<- Event, 0, len(s.watchers)-1)
		for _, watcher := range watchers {
			if watcher != ch {
				watchers = append(watchers, watcher)
			}
		}
		s.watchers = watchers
		s.watchMu.Unlock()
	}()
	return nil
}

func (s *localStore) Close() error {
	close(s.eventCh)
	return nil
}

var _ Store = &localStore{}
