// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package task

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

var log = logging.GetLogger("store", "task")

// NewAtomixStore returns a new persistent Store
func NewAtomixStore(client atomix.Client) (Store, error) {
	tasks, err := client.GetMap(context.Background(), "onos-e2t-tasks")
	if err != nil {
		return nil, err
	}
	return &atomixStore{
		tasks: tasks,
	}, nil
}

// Store stores task information
type Store interface {
	io.Closer

	// Get retrieves an task from the store
	Get(ctx context.Context, id api.TaskID) (*api.Task, error)

	// Create creates an task in the store
	Create(ctx context.Context, sub *api.Task) error

	// Update updates an existing task in the store
	Update(ctx context.Context, sub *api.Task) error

	// Delete deletes a task from the store
	Delete(ctx context.Context, sub *api.Task) error

	// List streams tasks to the given channel
	List(ctx context.Context) ([]api.Task, error)

	// Watch streams task events to the given channel
	Watch(ctx context.Context, ch chan<- api.TaskEvent, opts ...WatchOption) error
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

// atomixStore is the task implementation of the Store
type atomixStore struct {
	tasks _map.Map
}

func (s *atomixStore) Get(ctx context.Context, id api.TaskID) (*api.Task, error) {
	if err := id.Validate(); err != nil {
		return nil, err
	}

	entry, err := s.tasks.Get(ctx, id.Key())
	if err != nil {
		return nil, errors.FromAtomix(err)
	}
	return decodeTask(*entry)
}

func (s *atomixStore) Create(ctx context.Context, task *api.Task) error {
	if err := task.ID.Validate(); err != nil {
		return err
	}

	log.Infof("Creating subscription task %+v", task)
	bytes, err := proto.Marshal(task)
	if err != nil {
		log.Errorf("Failed to create subscription task %+v: %s", task, err)
		return errors.NewInvalid(err.Error())
	}

	// Put the subscription task in the map using an optimistic lock if this is an update
	entry, err := s.tasks.Put(ctx, task.ID.Key(), bytes, _map.IfNotSet())
	if err != nil {
		log.Errorf("Failed to create subscription task %+v: %s", task, err)
		return errors.FromAtomix(err)
	}

	task.Revision = api.Revision(entry.Revision)
	return nil
}

func (s *atomixStore) Update(ctx context.Context, task *api.Task) error {
	if err := task.ID.Validate(); err != nil {
		return err
	}
	if task.Revision == 0 {
		return errors.NewInvalid("subscription task must contain a revision on update")
	}

	log.Infof("Updating subscription task %+v", task)
	bytes, err := proto.Marshal(task)
	if err != nil {
		log.Errorf("Failed to update subscription task %+v: %s", task, err)
		return errors.NewInvalid(err.Error())
	}

	// Update the subscription task in the map
	entry, err := s.tasks.Put(ctx, task.ID.Key(), bytes, _map.IfMatch(meta.NewRevision(meta.Revision(task.Revision))))
	if err != nil {
		log.Errorf("Failed to update subscription %+v: %s", task, err)
		return errors.FromAtomix(err)
	}
	task.Revision = api.Revision(entry.Revision)
	return nil
}

func (s *atomixStore) Delete(ctx context.Context, task *api.Task) error {
	if err := task.ID.Validate(); err != nil {
		return err
	}
	if task.Revision == 0 {
		return errors.NewInvalid("subscription task must contain a revision on update")
	}

	log.Infof("Deleting subscription task %s", task.ID)
	_, err := s.tasks.Remove(ctx, task.ID.Key(), _map.IfMatch(meta.NewRevision(meta.Revision(task.Revision))))
	if err != nil {
		log.Errorf("Failed to delete subscription task %s: %s", task.ID, err)
		return errors.FromAtomix(err)
	}
	return nil
}

func (s *atomixStore) List(ctx context.Context) ([]api.Task, error) {
	mapCh := make(chan _map.Entry)
	if err := s.tasks.Entries(ctx, mapCh); err != nil {
		return nil, errors.FromAtomix(err)
	}

	tasks := make([]api.Task, 0)
	for entry := range mapCh {
		if ep, err := decodeTask(entry); err == nil {
			tasks = append(tasks, *ep)
		}
	}
	return tasks, nil
}

func (s *atomixStore) Watch(ctx context.Context, ch chan<- api.TaskEvent, opts ...WatchOption) error {
	watchOpts := make([]_map.WatchOption, 0)
	for _, opt := range opts {
		watchOpts = opt.apply(watchOpts)
	}

	mapCh := make(chan _map.Event)
	if err := s.tasks.Watch(ctx, mapCh, watchOpts...); err != nil {
		return errors.FromAtomix(err)
	}

	go func() {
		defer close(ch)
		for event := range mapCh {
			if task, err := decodeTask(event.Entry); err == nil {
				var eventType api.TaskEventType
				switch event.Type {
				case _map.EventReplay:
					eventType = api.TaskEventType_TASK_EVENT_UNKNOWN
				case _map.EventInsert:
					eventType = api.TaskEventType_TASK_CREATED
				case _map.EventRemove:
					eventType = api.TaskEventType_TASK_DELETED
				case _map.EventUpdate:
					eventType = api.TaskEventType_TASK_UPDATED
				default:
					eventType = api.TaskEventType_TASK_UPDATED
				}
				ch <- api.TaskEvent{
					Type: eventType,
					Task: *task,
				}
			}
		}
	}()
	return nil
}

func (s *atomixStore) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	_ = s.tasks.Close(ctx)
	defer cancel()
	return s.tasks.Close(ctx)
}

func decodeTask(entry _map.Entry) (*api.Task, error) {
	task := &api.Task{}
	if err := proto.Unmarshal(entry.Value, task); err != nil {
		return nil, errors.NewInvalid(err.Error())
	}
	task.Revision = api.Revision(entry.Revision)
	return task, nil
}
