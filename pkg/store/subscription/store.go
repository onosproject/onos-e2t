// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package subscription

import (
	"context"
	"github.com/atomix/atomix-go-client/pkg/atomix"
	"github.com/atomix/atomix-go-framework/pkg/atomix/meta"
	"github.com/onosproject/onos-lib-go/pkg/errors"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"io"
	"time"

	_map "github.com/atomix/atomix-go-client/pkg/atomix/map"
	"github.com/gogo/protobuf/proto"
	api "github.com/onosproject/onos-e2t/api/onos/e2t/store/subscription"
)

var log = logging.GetLogger("store", "subscription")

// NewAtomixStore returns a new persistent Store
func NewAtomixStore(client atomix.Client) (Store, error) {
	subs, err := client.GetMap(context.Background(), "onos-e2t-subscriptions")
	if err != nil {
		return nil, err
	}
	return &atomixStore{
		subs: subs,
	}, nil
}

// Store stores subscription information
type Store interface {
	io.Closer

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
}

// WatchOption is a configuration option for Watch calls
type WatchOption interface {
	apply([]_map.WatchOption) []_map.WatchOption
}

// watchReplyOption is an option to replay events on watch
type watchReplayOption struct {
}

func (o watchReplayOption) apply(opts []_map.WatchOption) []_map.WatchOption {
	return append(opts, _map.WithReplay())
}

// WithReplay returns a WatchOption that replays past changes
func WithReplay() WatchOption {
	return watchReplayOption{}
}

// atomixStore is the subscription implementation of the Store
type atomixStore struct {
	subs _map.Map
}

func (s *atomixStore) Get(ctx context.Context, id api.SubscriptionID) (*api.Subscription, error) {
	if err := id.Validate(); err != nil {
		return nil, err
	}

	entry, err := s.subs.Get(ctx, id.Key())
	if err != nil {
		return nil, errors.FromAtomix(err)
	}
	return decodeSub(*entry)
}

func (s *atomixStore) Create(ctx context.Context, sub *api.Subscription) error {
	if err := sub.ID.Validate(); err != nil {
		return err
	}

	log.Infof("Creating subscription %+v", sub)
	bytes, err := proto.Marshal(sub)
	if err != nil {
		log.Errorf("Failed to create subscription %+v: %s", sub, err)
		return errors.NewInvalid(err.Error())
	}

	// Put the subscription in the map using an optimistic lock if this is an update
	entry, err := s.subs.Put(ctx, sub.ID.Key(), bytes, _map.IfNotSet())
	if err != nil {
		log.Errorf("Failed to create subscription %+v: %s", sub, err)
		return errors.FromAtomix(err)
	}

	sub.Revision = api.Revision(entry.Revision)
	return nil
}

func (s *atomixStore) Update(ctx context.Context, sub *api.Subscription) error {
	if err := sub.ID.Validate(); err != nil {
		return err
	}
	if sub.Revision == 0 {
		return errors.NewInvalid("subscription must contain a revision on update")
	}

	log.Infof("Updating subscription %+v", sub)
	bytes, err := proto.Marshal(sub)
	if err != nil {
		log.Errorf("Failed to update subscription %+v: %s", sub, err)
		return errors.NewInvalid(err.Error())
	}

	// Update the subscription in the map
	entry, err := s.subs.Put(ctx, sub.ID.Key(), bytes, _map.IfMatch(meta.NewRevision(meta.Revision(sub.Revision))))
	if err != nil {
		log.Errorf("Failed to update subscription %+v: %s", sub, err)
		return errors.FromAtomix(err)
	}
	sub.Revision = api.Revision(entry.Revision)
	return nil
}

func (s *atomixStore) Delete(ctx context.Context, sub *api.Subscription) error {
	if err := sub.ID.Validate(); err != nil {
		return err
	}
	if sub.Revision == 0 {
		return errors.NewInvalid("subscription must contain a revision on update")
	}

	log.Infof("Deleting subscription %s", sub.ID)
	_, err := s.subs.Remove(ctx, sub.ID.Key(), _map.IfMatch(meta.NewRevision(meta.Revision(sub.Revision))))
	if err != nil {
		log.Errorf("Failed to delete subscription %s: %s", sub.ID, err)
		return errors.FromAtomix(err)
	}
	return nil
}

func (s *atomixStore) List(ctx context.Context) ([]api.Subscription, error) {
	mapCh := make(chan _map.Entry)
	if err := s.subs.Entries(ctx, mapCh); err != nil {
		return nil, errors.FromAtomix(err)
	}

	eps := make([]api.Subscription, 0)

	for entry := range mapCh {
		if ep, err := decodeSub(entry); err == nil {
			eps = append(eps, *ep)
		}
	}
	return eps, nil
}

func (s *atomixStore) Watch(ctx context.Context, ch chan<- api.SubscriptionEvent, opts ...WatchOption) error {
	watchOpts := make([]_map.WatchOption, 0)
	for _, opt := range opts {
		watchOpts = opt.apply(watchOpts)
	}

	mapCh := make(chan _map.Event)
	if err := s.subs.Watch(ctx, mapCh, watchOpts...); err != nil {
		return errors.FromAtomix(err)
	}

	go func() {
		defer close(ch)
		for event := range mapCh {
			if sub, err := decodeSub(event.Entry); err == nil {
				var eventType api.SubscriptionEventType
				switch event.Type {
				case _map.EventReplay:
					eventType = api.SubscriptionEventType_SUBSCRIPTION_EVENT_UNKNOWN
				case _map.EventInsert:
					eventType = api.SubscriptionEventType_SUBSCRIPTION_CREATED
				case _map.EventRemove:
					eventType = api.SubscriptionEventType_SUBSCRIPTION_DELETED
				case _map.EventUpdate:
					eventType = api.SubscriptionEventType_SUBSCRIPTION_UPDATED
				default:
					eventType = api.SubscriptionEventType_SUBSCRIPTION_UPDATED
				}
				ch <- api.SubscriptionEvent{
					Type:         eventType,
					Subscription: *sub,
				}
			}
		}
	}()
	return nil
}

func (s *atomixStore) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	_ = s.subs.Close(ctx)
	defer cancel()
	return s.subs.Close(ctx)
}

func decodeSub(entry _map.Entry) (*api.Subscription, error) {
	sub := &api.Subscription{}
	if err := proto.Unmarshal(entry.Value, sub); err != nil {
		return nil, errors.NewInvalid(err.Error())
	}
	sub.Revision = api.Revision(entry.Revision)
	return sub, nil
}
