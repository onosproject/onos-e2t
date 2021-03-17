// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package subscription

import (
	"context"
	"io"
	"time"

	"github.com/atomix/go-client/pkg/client/util/net"
	"github.com/onosproject/onos-lib-go/pkg/errors"
	"github.com/onosproject/onos-lib-go/pkg/logging"

	_map "github.com/atomix/go-client/pkg/client/map"
	"github.com/atomix/go-client/pkg/client/primitive"
	"github.com/gogo/protobuf/proto"
	subapi "github.com/onosproject/onos-api/go/onos/e2sub/subscription"
	"github.com/onosproject/onos-e2sub/pkg/config"
	"github.com/onosproject/onos-lib-go/pkg/atomix"
)

var log = logging.GetLogger("store", "subscription")

// NewAtomixStore returns a new persistent Store
func NewAtomixStore() (Store, error) {
	ricConfig, err := config.GetConfig()
	if err != nil {
		return nil, err
	}

	database, err := atomix.GetDatabase(ricConfig.Atomix, ricConfig.Atomix.GetDatabase(atomix.DatabaseTypeConsensus))
	if err != nil {
		return nil, err
	}

	subscriptions, err := database.GetMap(context.Background(), "subscriptions")
	if err != nil {
		return nil, err
	}

	return &atomixStore{
		subscriptions: subscriptions,
	}, nil
}

// NewLocalStore returns a new local subscription store
func NewLocalStore() (Store, error) {
	_, address := atomix.StartLocalNode()
	return newLocalStore(address)
}

// newLocalStore creates a new local subscription store
func newLocalStore(address net.Address) (Store, error) {
	name := primitive.Name{
		Namespace: "local",
		Name:      "subscriptions",
	}

	session, err := primitive.NewSession(context.TODO(), primitive.Partition{ID: 1, Address: address})
	if err != nil {
		return nil, err
	}

	subscriptions, err := _map.New(context.Background(), name, []*primitive.Session{session})
	if err != nil {
		return nil, err
	}

	return &atomixStore{
		subscriptions: subscriptions,
	}, nil
}

// Store stores end-point registry information
type Store interface {
	io.Closer

	// Create creates a subscription in the store
	Create(ctx context.Context, sub *subapi.Subscription) error

	// Update updates a subscription in the store
	Update(ctx context.Context, sub *subapi.Subscription) error

	// Delete deletes an subscription from the store
	Get(ctx context.Context, id subapi.ID) (*subapi.Subscription, error)

	// Delete deletes an subscription from the store
	Delete(ctx context.Context, id subapi.ID) error

	// List streams subscriptions to the given channel
	List(ctx context.Context) ([]subapi.Subscription, error)

	// Watch streams subscription events to the given channel
	Watch(ctx context.Context, ch chan<- subapi.Event, opts ...WatchOption) error
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

// atomixStore is the implementation of the subscription Store
type atomixStore struct {
	subscriptions _map.Map
}

func (s *atomixStore) Create(ctx context.Context, sub *subapi.Subscription) error {
	if sub.ID == "" {
		return errors.NewInvalid("ID cannot be empty")
	}

	log.Infof("Creating Subscription %+v", sub)
	bytes, err := proto.Marshal(sub)
	if err != nil {
		log.Errorf("Failed to create Subscription %+v: %s", sub, err)
		return errors.NewInvalid(err.Error())
	}

	// Create the subscription in the map only if it does not already exist
	entry, err := s.subscriptions.Put(ctx, string(sub.ID), bytes, _map.IfNotSet())
	if err != nil {
		log.Errorf("Failed to create Subscription %+v: %s", sub, err)
		return errors.FromAtomix(err)
	}
	sub.Revision = subapi.Revision(entry.Version)
	return nil
}

func (s *atomixStore) Update(ctx context.Context, sub *subapi.Subscription) error {
	if sub.ID == "" {
		return errors.NewInvalid("ID cannot be empty")
	}
	if sub.Revision == 0 {
		return errors.NewInvalid("object must contain a revision on update")
	}

	log.Infof("Updating Subscription %+v", sub)
	bytes, err := proto.Marshal(sub)
	if err != nil {
		log.Errorf("Failed to update Subscription %+v: %s", sub, err)
		return errors.NewInvalid(err.Error())
	}

	// Update the subscription in the map
	entry, err := s.subscriptions.Put(ctx, string(sub.ID), bytes, _map.IfVersion(_map.Version(sub.Revision)))
	if err != nil {
		log.Errorf("Failed to update Subscription %+v: %s", sub, err)
		return errors.FromAtomix(err)
	}
	sub.Revision = subapi.Revision(entry.Version)
	return nil
}

func (s *atomixStore) Get(ctx context.Context, id subapi.ID) (*subapi.Subscription, error) {
	if id == "" {
		return nil, errors.NewInvalid("ID cannot be empty")
	}

	entry, err := s.subscriptions.Get(ctx, string(id))
	if err != nil {
		return nil, errors.FromAtomix(err)
	}
	return decodeObject(entry)
}

func (s *atomixStore) Delete(ctx context.Context, id subapi.ID) error {
	if id == "" {
		return errors.NewInvalid("ID cannot be empty")
	}
	log.Infof("Deleting Subscription %s", id)
	_, err := s.subscriptions.Remove(ctx, string(id))
	if err != nil {
		log.Errorf("Failed to delete Subscription %s: %s", id, err)
		return errors.FromAtomix(err)
	}
	return nil
}

func (s *atomixStore) List(ctx context.Context) ([]subapi.Subscription, error) {
	mapCh := make(chan *_map.Entry)
	if err := s.subscriptions.Entries(context.Background(), mapCh); err != nil {
		return nil, errors.FromAtomix(err)
	}

	subs := make([]subapi.Subscription, 0)
	for entry := range mapCh {
		if sub, err := decodeObject(entry); err == nil {
			subs = append(subs, *sub)
		}
	}
	return subs, nil
}

func (s *atomixStore) Watch(ctx context.Context, ch chan<- subapi.Event, opts ...WatchOption) error {
	watchOpts := make([]_map.WatchOption, 0)
	for _, opt := range opts {
		watchOpts = opt.apply(watchOpts)
	}

	mapCh := make(chan *_map.Event)
	if err := s.subscriptions.Watch(context.Background(), mapCh, watchOpts...); err != nil {
		return errors.FromAtomix(err)
	}

	go func() {
		defer close(ch)
		for event := range mapCh {
			if sub, err := decodeObject(event.Entry); err == nil {
				var eventType subapi.EventType
				switch event.Type {
				case _map.EventNone:
					eventType = subapi.EventType_NONE
				case _map.EventInserted:
					eventType = subapi.EventType_ADDED
				case _map.EventUpdated:
					eventType = subapi.EventType_UPDATED
				case _map.EventRemoved:
					eventType = subapi.EventType_REMOVED
				}
				ch <- subapi.Event{
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
	defer cancel()
	return s.subscriptions.Close(ctx)
}

func decodeObject(entry *_map.Entry) (*subapi.Subscription, error) {
	sub := &subapi.Subscription{}
	if err := proto.Unmarshal(entry.Value, sub); err != nil {
		return nil, errors.NewInvalid(err.Error())
	}
	sub.ID = subapi.ID(entry.Key)
	sub.Revision = subapi.Revision(entry.Version)
	return sub, nil
}
