// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package channel

import (
	"context"
	"github.com/atomix/go-sdk/pkg/generic"
	"github.com/google/uuid"
	"io"
	"sync"
	"time"

	"github.com/atomix/go-sdk/pkg/primitive"
	api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	"github.com/onosproject/onos-lib-go/pkg/errors"
	"github.com/onosproject/onos-lib-go/pkg/logging"

	_map "github.com/atomix/go-sdk/pkg/primitive/map"
)

var log = logging.GetLogger()

// NewAtomixStore returns a new persistent Store
func NewAtomixStore(client primitive.Client) (Store, error) {
	channels, err := _map.NewBuilder[api.ChannelID, *api.Channel](client, "onos-e2t-objects").
		Tag("onos-e2t", "channels").Codec(generic.Proto[*api.Channel](&api.Channel{})).Get(context.Background())
	if err != nil {
		return nil, errors.FromAtomix(err)
	}

	store := &atomixStore{
		channels: channels,
		cache:    make(map[api.ChannelID]api.Channel),
		watchers: make(map[uuid.UUID]chan<- api.ChannelEvent),
	}

	events, err := channels.Events(context.Background())
	if err != nil {
		return nil, errors.FromAtomix(err)
	}
	entries, err := channels.List(context.Background())
	if err != nil {
		return nil, errors.FromAtomix(err)
	}
	go store.watchStoreEvents(entries, events)
	return store, nil
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

// atomixStore is the channel implementation of the Store
type atomixStore struct {
	channels   _map.Map[api.ChannelID, *api.Channel]
	cache      map[api.ChannelID]api.Channel
	cacheMu    sync.RWMutex
	watchers   map[uuid.UUID]chan<- api.ChannelEvent
	watchersMu sync.RWMutex
}

func (s *atomixStore) watchStoreEvents(entries _map.EntryStream[api.ChannelID, *api.Channel], events _map.EventStream[api.ChannelID, *api.Channel]) {
	for {
		entry, err := entries.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Error(err)
			continue
		}

		channel := entry.Value
		channel.Revision = api.Revision(entry.Version)

		s.cacheMu.Lock()
		s.cache[channel.ID] = *channel
		s.cacheMu.Unlock()

		s.watchersMu.RLock()
		for _, watcher := range s.watchers {
			watcher <- api.ChannelEvent{
				Type:    api.ChannelEventType_CHANNEL_REPLAYED,
				Channel: *channel,
			}
		}
		s.watchersMu.RUnlock()
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

		var eventType api.ChannelEventType
		var channel *api.Channel
		switch e := event.(type) {
		case *_map.Inserted[api.ChannelID, *api.Channel]:
			channel = e.Entry.Value
			channel.Revision = api.Revision(e.Entry.Version)
			eventType = api.ChannelEventType_CHANNEL_CREATED
			s.cacheMu.Lock()
			s.cache[channel.ID] = *channel
			s.cacheMu.Unlock()
		case *_map.Updated[api.ChannelID, *api.Channel]:
			channel = e.NewEntry.Value
			channel.Revision = api.Revision(e.NewEntry.Version)
			eventType = api.ChannelEventType_CHANNEL_UPDATED
			s.cacheMu.Lock()
			s.cache[channel.ID] = *channel
			s.cacheMu.Unlock()
		case *_map.Removed[api.ChannelID, *api.Channel]:
			channel = e.Entry.Value
			channel.Revision = api.Revision(e.Entry.Version)
			eventType = api.ChannelEventType_CHANNEL_DELETED
			s.cacheMu.Lock()
			delete(s.cache, e.Entry.Key)
			s.cacheMu.Unlock()
		}

		s.watchersMu.RLock()
		for _, watcher := range s.watchers {
			watcher <- api.ChannelEvent{
				Type:    eventType,
				Channel: *channel,
			}
		}
		s.watchersMu.RUnlock()
	}
}

func (s *atomixStore) Get(ctx context.Context, id api.ChannelID) (*api.Channel, error) {
	if id == "" {
		return nil, errors.NewInvalid("ID cannot be empty")
	}

	entry, err := s.channels.Get(ctx, id)
	if err != nil {
		err = errors.FromAtomix(err)
		if !errors.IsNotFound(err) {
			log.Errorf("Failed to get channel %+v: %+v", id, err)
		} else {
			log.Warnf("Failed to get channel %+v: %+v", id, err)
		}
		return nil, err
	}
	channel := entry.Value
	channel.Revision = api.Revision(entry.Version)
	return channel, nil
}

func (s *atomixStore) Create(ctx context.Context, channel *api.Channel) error {
	log.Infof("Creating channel %+v", channel)

	// Put the channel in the map using an optimistic lock if this is an update
	entry, err := s.channels.Insert(ctx, channel.ID, channel)
	if err != nil {
		err = errors.FromAtomix(err)
		if !errors.IsAlreadyExists(err) {
			log.Errorf("Failed to create channel %+v: %+v", channel, err)
		} else {
			log.Errorf("Failed to create channel %+v: %+v", channel, err)
		}
		return err
	}

	channel.Revision = api.Revision(entry.Version)
	return nil
}

func (s *atomixStore) Update(ctx context.Context, channel *api.Channel) error {
	if channel.ID == "" {
		return errors.NewInvalid("ID cannot be empty")
	}

	if channel.Revision == 0 {
		return errors.NewInvalid("channel must contain a revision on update")
	}

	log.Infof("Updating channel %+v", channel)

	// Update the channel in the map
	entry, err := s.channels.Update(ctx, channel.ID, channel, _map.IfVersion(primitive.Version(channel.Revision)))
	if err != nil {
		err = errors.FromAtomix(err)
		if !errors.IsNotFound(err) && !errors.IsConflict(err) {
			log.Errorf("Failed to update channel %+v: %+v", channel, err)
		} else {
			log.Warnf("Failed to update channel %+v: %+v", channel, err)
		}
		return err
	}
	channel.Revision = api.Revision(entry.Version)
	return nil
}

func (s *atomixStore) Delete(ctx context.Context, channel *api.Channel) error {
	var err error
	if channel.ID == "" {
		return errors.NewInvalid("ID cannot be empty")
	}

	log.Infof("Deleting channel %s", channel.ID)

	if channel.Revision == 0 {
		_, err = s.channels.Remove(ctx, channel.ID)
	} else {
		_, err = s.channels.Remove(ctx, channel.ID, _map.IfVersion(primitive.Version(channel.Revision)))
	}

	if err != nil {
		err = errors.FromAtomix(err)
		if !errors.IsNotFound(err) && !errors.IsConflict(err) {
			log.Errorf("Failed to delete channel %+v: %+v", channel.ID, err)
		} else {
			log.Warnf("Failed to delete channel %+v: %+v", channel.ID, err)
		}
		return err
	}
	return nil
}

func (s *atomixStore) List(ctx context.Context) ([]api.Channel, error) {
	list, err := s.channels.List(ctx)
	if err != nil {
		return nil, errors.FromAtomix(err)
	}

	eps := make([]api.Channel, 0)

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

func (s *atomixStore) Watch(ctx context.Context, ch chan<- api.ChannelEvent, opts ...WatchOption) error {
	var watchOpts watchOptions
	for _, opt := range opts {
		opt.apply(&watchOpts)
	}

	// create separate channels for replay and watch events
	replayCh := make(chan api.Channel)
	eventCh := make(chan api.ChannelEvent)

	go func() {
		defer close(ch)

	replayLoop:
		// process the replay channel first
		for {
			select {
			case channel, ok := <-replayCh:
				// if the replay channel is closed, break out of the replay loop
				if !ok {
					break replayLoop
				}
				// if a channel is received on the replay channel, write it to the watch channel
				ch <- api.ChannelEvent{
					Type:    api.ChannelEventType_CHANNEL_REPLAYED,
					Channel: channel,
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
				break eventLoop
			}
		}
	}()

	// add watcher's event channel
	watcherID := uuid.New()
	s.watchersMu.Lock()
	s.watchers[watcherID] = eventCh
	s.watchersMu.Unlock()

	// get the objects to replay
	var channels []api.Channel
	if watchOpts.replay {
		s.cacheMu.RLock()
		channels = make([]api.Channel, 0, len(s.cache))
		for _, channel := range s.cache {
			channels = append(channels, channel)
		}
		s.cacheMu.RUnlock()
	}

	// replay existing channels in the cache and then close the replay channel
	go func() {
		defer close(replayCh)
		for _, channel := range channels {
			replayCh <- channel
		}
	}()

	// remove the watcher and close the event channel once the watcher context is done
	go func() {
		<-ctx.Done()
		s.watchersMu.Lock()
		delete(s.watchers, watcherID)
		s.watchersMu.Unlock()
		close(eventCh)
	}()
	return nil
}

func (s *atomixStore) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := s.channels.Close(ctx)
	if err != nil {
		return errors.FromAtomix(err)
	}
	return nil
}
