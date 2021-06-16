// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package channel

import (
	"context"
	"github.com/atomix/atomix-go-client/pkg/atomix"
	"github.com/atomix/atomix-go-framework/pkg/atomix/meta"
	api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	"github.com/onosproject/onos-lib-go/pkg/errors"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"io"
	"time"

	_map "github.com/atomix/atomix-go-client/pkg/atomix/map"
	"github.com/gogo/protobuf/proto"
)

var log = logging.GetLogger("store", "channel")

// NewAtomixStore returns a new persistent Store
func NewAtomixStore(client atomix.Client) (Store, error) {
	channels, err := client.GetMap(context.Background(), "onos-e2t-channels")
	if err != nil {
		return nil, err
	}
	return &atomixStore{
		channels: channels,
	}, nil
}

// Store stores channel information
type Store interface {
	io.Closer

	// Get retrieves an channel from the store
	Get(ctx context.Context, id api.ChannelID) (*api.Channel, error)

	// Create creates an channel in the store
	Create(ctx context.Context, channel *api.Channel) error

	// Update updates an existing channel in the store
	Update(ctx context.Context, channel *api.Channel) error

	// Delete deletes a channel from the store
	Delete(ctx context.Context, channel *api.Channel) error

	// List streams channels to the given channel
	List(ctx context.Context) ([]api.Channel, error)

	// Watch streams channel events to the given channel
	Watch(ctx context.Context, ch chan<- api.ChannelEvent, opts ...WatchOption) error
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

// atomixStore is the channel implementation of the Store
type atomixStore struct {
	channels _map.Map
}

func (s *atomixStore) Get(ctx context.Context, id api.ChannelID) (*api.Channel, error) {
	entry, err := s.channels.Get(ctx, string(id))
	if err != nil {
		return nil, errors.FromAtomix(err)
	}
	return decodeChannel(*entry)
}

func (s *atomixStore) Create(ctx context.Context, channel *api.Channel) error {
	log.Infof("Creating channel %+v", channel)
	bytes, err := proto.Marshal(channel)
	if err != nil {
		log.Errorf("Failed to create channel %+v: %s", channel, err)
		return errors.NewInvalid(err.Error())
	}

	// Put the channel in the map using an optimistic lock if this is an update
	entry, err := s.channels.Put(ctx, string(channel.ID), bytes, _map.IfNotSet())
	if err != nil {
		log.Errorf("Failed to create channel %+v: %s", channel, err)
		return errors.FromAtomix(err)
	}

	channel.Revision = api.Revision(entry.Revision)
	return nil
}

func (s *atomixStore) Update(ctx context.Context, channel *api.Channel) error {
	if channel.Revision == 0 {
		return errors.NewInvalid("channel must contain a revision on update")
	}

	log.Infof("Updating channel %+v", channel)
	bytes, err := proto.Marshal(channel)
	if err != nil {
		log.Errorf("Failed to update channel %+v: %s", channel, err)
		return errors.NewInvalid(err.Error())
	}

	// Update the channel in the map
	entry, err := s.channels.Put(ctx, string(channel.ID), bytes, _map.IfMatch(meta.NewRevision(meta.Revision(channel.Revision))))
	if err != nil {
		log.Errorf("Failed to update channel %+v: %s", channel, err)
		return errors.FromAtomix(err)
	}
	channel.Revision = api.Revision(entry.Revision)
	return nil
}

func (s *atomixStore) Delete(ctx context.Context, channel *api.Channel) error {
	if channel.Revision == 0 {
		return errors.NewInvalid("channel must contain a revision on update")
	}

	log.Infof("Deleting channel %s", channel.ID)
	_, err := s.channels.Remove(ctx, string(channel.ID), _map.IfMatch(meta.NewRevision(meta.Revision(channel.Revision))))
	if err != nil {
		log.Errorf("Failed to delete channel %s: %s", channel.ID, err)
		return errors.FromAtomix(err)
	}
	return nil
}

func (s *atomixStore) List(ctx context.Context) ([]api.Channel, error) {
	mapCh := make(chan _map.Entry)
	if err := s.channels.Entries(ctx, mapCh); err != nil {
		return nil, errors.FromAtomix(err)
	}

	eps := make([]api.Channel, 0)

	for entry := range mapCh {
		if ep, err := decodeChannel(entry); err == nil {
			eps = append(eps, *ep)
		}
	}
	return eps, nil
}

func (s *atomixStore) Watch(ctx context.Context, ch chan<- api.ChannelEvent, opts ...WatchOption) error {
	watchOpts := make([]_map.WatchOption, 0)
	for _, opt := range opts {
		watchOpts = opt.apply(watchOpts)
	}

	eventCh := make(chan _map.Event)
	if err := s.channels.Watch(ctx, eventCh, watchOpts...); err != nil {
		return errors.FromAtomix(err)
	}

	go func() {
		defer close(ch)
		for event := range eventCh {
			if channel, err := decodeChannel(event.Entry); err == nil {
				var eventType api.ChannelEventType
				switch event.Type {
				case _map.EventReplay:
					eventType = api.ChannelEventType_CHANNEL_EVENT_UNKNOWN
				case _map.EventInsert:
					eventType = api.ChannelEventType_CHANNEL_CREATED
				case _map.EventRemove:
					eventType = api.ChannelEventType_CHANNEL_DELETED
				case _map.EventUpdate:
					eventType = api.ChannelEventType_CHANNEL_UPDATED
				default:
					eventType = api.ChannelEventType_CHANNEL_UPDATED
				}
				ch <- api.ChannelEvent{
					Type:    eventType,
					Channel: *channel,
				}
			}
		}
	}()
	return nil
}

func (s *atomixStore) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	_ = s.channels.Close(ctx)
	defer cancel()
	return s.channels.Close(ctx)
}

func decodeChannel(entry _map.Entry) (*api.Channel, error) {
	channel := &api.Channel{}
	if err := proto.Unmarshal(entry.Value, channel); err != nil {
		return nil, errors.NewInvalid(err.Error())
	}
	channel.Revision = api.Revision(entry.Revision)
	return channel, nil
}
