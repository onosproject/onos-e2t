// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package subscription

import (
	"context"
	"github.com/atomix/go-sdk/pkg/generic"
	"github.com/google/uuid"
	"io"
	"sync"

	"github.com/atomix/go-sdk/pkg/primitive"
	api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	"github.com/onosproject/onos-lib-go/pkg/errors"
	"github.com/onosproject/onos-lib-go/pkg/logging"

	_map "github.com/atomix/go-sdk/pkg/primitive/map"
)

var log = logging.GetLogger()

// NewAtomixStore returns a new persistent Store
func NewAtomixStore(client primitive.Client) (Store, error) {
	subs, err := _map.NewBuilder[api.SubscriptionID, *api.Subscription](client, "onos-e2t-subscriptions").
		Tag("onos-e2t", "subs").Codec(generic.Proto[*api.Subscription](&api.Subscription{})).Get(context.Background())
	if err != nil {
		return nil, errors.FromAtomix(err)
	}

	store := &atomixStore{
		subs:     subs,
		cache:    make(map[api.SubscriptionID]api.Subscription),
		watchers: make(map[uuid.UUID]chan<- api.SubscriptionEvent),
	}

	events, err := subs.Events(context.Background())
	if err != nil {
		return nil, errors.FromAtomix(err)
	}
	entries, err := subs.List(context.Background())
	if err != nil {
		return nil, errors.FromAtomix(err)
	}

	go store.watchStoreEvents(entries, events)
	return store, nil
}

// Store stores subscription information
type Store interface {
	// Get retrieves an subscription from the store
	Get(ctx context.Context, id api.SubscriptionID) (*api.Subscription, error)

	// Create creates an subscription in the store
	Create(ctx context.Context, sub *api.Subscription) error

	// Update updates an existing subscription in the store
	Update(ctx context.Context, sub *api.Subscription) error

	// Delete deletes a subscription from the store
	Delete(ctx context.Context, sub *api.Subscription) error

	// List streams subscriptions to the given channel
	List(ctx context.Context) ([]api.Subscription, error)

	// Watch streams subscription events to the given channel
	Watch(ctx context.Context, ch chan<- api.SubscriptionEvent, opts ...WatchOption) error

	// Close closes the subscription store
	Close(ctx context.Context) error
}

// WatchOption is a configuration option for Watch calls
type WatchOption interface {
	apply(*watchOptions)
}

// watchReplyOption is an option to replay events on watch
type watchReplayOption struct {
	replay bool
}

func (o watchReplayOption) apply(opts *watchOptions) {
	opts.replay = o.replay
}

// WithReplay returns a WatchOption that replays past changes
func WithReplay() WatchOption {
	return watchReplayOption{true}
}

type watchOptions struct {
	replay bool
}

// atomixStore is the subscription implementation of the Store
type atomixStore struct {
	subs      _map.Map[api.SubscriptionID, *api.Subscription]
	cache     map[api.SubscriptionID]api.Subscription
	cacheMu   sync.RWMutex
	watchers  map[uuid.UUID]chan<- api.SubscriptionEvent
	watcherMu sync.RWMutex
}

func (s *atomixStore) watchStoreEvents(entries _map.EntryStream[api.SubscriptionID, *api.Subscription], events _map.EventStream[api.SubscriptionID, *api.Subscription]) {
	for {
		entry, err := entries.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Error(err)
			continue
		}

		sub := entry.Value
		sub.Revision = api.Revision(entry.Version)

		s.cacheMu.Lock()
		s.cache[sub.ID] = *sub
		s.cacheMu.Unlock()

		s.watcherMu.RLock()
		for _, watcher := range s.watchers {
			watcher <- api.SubscriptionEvent{
				Type:         api.SubscriptionEventType_SUBSCRIPTION_REPLAYED,
				Subscription: *sub,
			}
		}
		s.watcherMu.RUnlock()
	}

	for {
		event, err := events.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Error(err)
			continue
		}

		var eventType api.SubscriptionEventType
		var sub *api.Subscription
		switch e := event.(type) {
		case *_map.Inserted[api.SubscriptionID, *api.Subscription]:
			sub = e.Entry.Value
			sub.Revision = api.Revision(e.Entry.Version)
			eventType = api.SubscriptionEventType_SUBSCRIPTION_CREATED
			s.cacheMu.Lock()
			s.cache[sub.ID] = *sub
			s.cacheMu.Unlock()
		case *_map.Updated[api.SubscriptionID, *api.Subscription]:
			sub = e.NewEntry.Value
			sub.Revision = api.Revision(e.NewEntry.Version)
			eventType = api.SubscriptionEventType_SUBSCRIPTION_UPDATED
			s.cacheMu.Lock()
			s.cache[sub.ID] = *sub
			s.cacheMu.Unlock()
		case *_map.Removed[api.SubscriptionID, *api.Subscription]:
			sub = e.Entry.Value
			sub.Revision = api.Revision(e.Entry.Version)
			eventType = api.SubscriptionEventType_SUBSCRIPTION_DELETED
			s.cacheMu.Lock()
			delete(s.cache, e.Entry.Key)
			s.cacheMu.Unlock()
		}

		s.watcherMu.RLock()
		for _, watcher := range s.watchers {
			watcher <- api.SubscriptionEvent{
				Type:         eventType,
				Subscription: *sub,
			}
		}
		s.watcherMu.RUnlock()
	}
}

func (s *atomixStore) Get(ctx context.Context, id api.SubscriptionID) (*api.Subscription, error) {
	if id == "" {
		return nil, errors.NewInvalid("ID cannot be empty")
	}

	entry, err := s.subs.Get(ctx, id)
	if err != nil {
		err = errors.FromAtomix(err)
		if !errors.IsNotFound(err) {
			log.Errorf("Failed to get subscription %+v: %+v", id, err)
		} else {
			log.Warnf("Failed to get subscription %+v: %+v", id, err)
		}
		return nil, err
	}
	sub := entry.Value
	sub.Revision = api.Revision(entry.Version)
	return sub, nil
}

func (s *atomixStore) Create(ctx context.Context, sub *api.Subscription) error {
	log.Infof("Creating subscription %+v", sub)

	// Put the subscription in the map using an optimistic lock if this is an update
	entry, err := s.subs.Insert(ctx, sub.ID, sub)
	if err != nil {
		err = errors.FromAtomix(err)
		if !errors.IsAlreadyExists(err) {
			log.Errorf("Failed to create subscription %+v: %+v", sub, err)
		} else {
			log.Warnf("Failed to create subscription %+v: %+v", sub, err)
		}
		return err
	}

	sub.Revision = api.Revision(entry.Version)
	return nil
}

func (s *atomixStore) Update(ctx context.Context, sub *api.Subscription) error {
	if sub.ID == "" {
		return errors.NewInvalid("ID cannot be empty")
	}

	if sub.Revision == 0 {
		return errors.NewInvalid("subscription must contain a revision on update")
	}

	log.Infof("Updating subscription %+v", sub)

	// Update the subscription in the map
	entry, err := s.subs.Update(ctx, sub.ID, sub, _map.IfVersion(primitive.Version(sub.Revision)))
	if err != nil {
		err = errors.FromAtomix(err)
		if !errors.IsNotFound(err) && !errors.IsConflict(err) {
			log.Errorf("Failed to update sub %+v: %+v", sub, err)
		} else {
			log.Warnf("Failed to update sub %+v: %+v", sub, err)
		}
		return err
	}

	sub.Revision = api.Revision(entry.Version)
	return nil
}

func (s *atomixStore) Delete(ctx context.Context, sub *api.Subscription) error {
	var err error
	if sub.ID == "" {
		return errors.NewInvalid("ID cannot be empty")
	}

	if sub.Revision == 0 {
		_, err = s.subs.Remove(ctx, sub.ID)
	} else {
		_, err = s.subs.Remove(ctx, sub.ID, _map.IfVersion(primitive.Version(sub.Revision)))
	}

	if err != nil {
		err = errors.FromAtomix(err)
		if !errors.IsNotFound(err) && !errors.IsConflict(err) {
			log.Errorf("Failed to delete sub %+v: %+v", sub.ID, err)
		} else {
			log.Warnf("Failed to delete sub %+v: %+v", sub.ID, err)
		}
		return err
	}
	return nil
}

func (s *atomixStore) List(ctx context.Context) ([]api.Subscription, error) {
	list, err := s.subs.List(ctx)
	if err != nil {
		return nil, errors.FromAtomix(err)
	}

	eps := make([]api.Subscription, 0)

	for {
		entry, err := list.Next()
		if err == io.EOF {
			return eps, nil
		}
		if err != nil {
			return nil, errors.FromAtomix(err)
		}
		eps = append(eps, *entry.Value)
	}
}

func (s *atomixStore) Watch(ctx context.Context, ch chan<- api.SubscriptionEvent, opts ...WatchOption) error {
	var watchOpts watchOptions
	for _, opt := range opts {
		opt.apply(&watchOpts)
	}

	// create separate channels for replay and watch events
	replayCh := make(chan api.Subscription)
	eventCh := make(chan api.SubscriptionEvent)

	go func() {
		defer close(ch)

	replayLoop:
		// process the replay channel first
		for {
			select {
			case sub, ok := <-replayCh:
				// if the replay channel is closed, break out of the replay loop
				if !ok {
					break replayLoop
				}
				// if a channel is received on the replay channel, write it to the watch channel
				ch <- api.SubscriptionEvent{
					Type:         api.SubscriptionEventType_SUBSCRIPTION_REPLAYED,
					Subscription: sub,
				}
			case <-ctx.Done():
				// if the watch context is closed, drain the replay channel and break out the replay loop
				go func() {
					for range replayCh {
					}
				}()
				break replayLoop
			}
		}
	eventLoop:
		// once the replay channel is processed, process event channel
		for {
			select {
			case event, ok := <-eventCh:
				// if the event channel is closed, break out the event loop
				if !ok {
					break eventLoop
				}
				ch <- event
			case <-ctx.Done():
				// if the watch context is closed, drain the event channel and break out of the event loop
				go func() {
					for range eventCh {
					}
				}()
			}
		}
	}()

	// add watcher's event channel
	watcherID := uuid.New()
	s.watcherMu.Lock()
	s.watchers[watcherID] = eventCh
	s.watcherMu.Unlock()

	// get the objects to replay
	var subs []api.Subscription
	if watchOpts.replay {
		s.cacheMu.RLock()
		subs = make([]api.Subscription, 0, len(s.cache))
		for _, sub := range s.cache {
			subs = append(subs, sub)
		}
		s.cacheMu.RUnlock()
	}

	// replay existing subs in the cache and then close the replay channel
	go func() {
		defer close(replayCh)
		for _, sub := range subs {
			replayCh <- sub
		}
	}()

	// remove the watcher and close the event channel once the watcher context is done
	go func() {
		<-ctx.Done()
		s.watcherMu.Lock()
		delete(s.watchers, watcherID)
		s.watcherMu.Unlock()
		close(eventCh)
	}()
	return nil
}

func (s *atomixStore) Close(ctx context.Context) error {
	err := s.subs.Close(ctx)
	if err != nil {
		return errors.FromAtomix(err)
	}
	return nil
}
