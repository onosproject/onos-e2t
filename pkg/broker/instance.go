// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package broker

import (
	"context"
	"github.com/google/uuid"
	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	"sync"
)

type AppInstanceStreamManager interface {
	StreamManager
	Get(id e2api.AppInstanceID) (*AppInstanceStream, bool)
	List() []*AppInstanceStream
	Open(id e2api.AppInstanceID) *AppInstanceStream
	Watch(ctx context.Context, ch chan<- e2api.AppInstanceID)
	close(id e2api.AppInstanceID)
}

type appInstanceStreamManager struct {
	transaction *TransactionStream
	instances   map[e2api.AppInstanceID]*AppInstanceStream
	instancesMu sync.RWMutex
	watchers    map[uuid.UUID]chan<- e2api.AppInstanceID
	watchersMu  sync.RWMutex
}

func (s *appInstanceStreamManager) Get(id e2api.AppInstanceID) (*AppInstanceStream, bool) {
	s.instancesMu.RLock()
	defer s.instancesMu.RUnlock()
	sub, ok := s.instances[id]
	return sub, ok
}

func (s *appInstanceStreamManager) List() []*AppInstanceStream {
	s.instancesMu.RLock()
	defer s.instancesMu.RUnlock()
	instances := make([]*AppInstanceStream, 0, len(s.instances))
	for _, instance := range s.instances {
		instances = append(instances, instance)
	}
	return instances
}

func (s *appInstanceStreamManager) Open(id e2api.AppInstanceID) *AppInstanceStream {
	s.instancesMu.RLock()
	instance, ok := s.instances[id]
	s.instancesMu.RUnlock()
	if ok {
		return instance
	}
	s.instancesMu.Lock()
	defer s.instancesMu.Unlock()
	instance, ok = s.instances[id]
	if ok {
		return instance
	}
	instance = &AppInstanceStream{
		TransactionStream: s.transaction,
		AppInstanceID:     id,
	}
	instance.streams = &northboundStreamManager{
		instance: instance,
		streams:  make(map[StreamID]*NorthboundStream),
		watchers: make(map[uuid.UUID]chan<- StreamID),
	}
	s.instances[id] = instance
	instance.open()
	go s.notify(id)
	return instance
}

func (s *appInstanceStreamManager) notify(instanceID e2api.AppInstanceID) {
	s.watchersMu.RLock()
	for _, watcher := range s.watchers {
		watcher <- instanceID
	}
	s.watchersMu.RUnlock()
}

func (s *appInstanceStreamManager) Watch(ctx context.Context, ch chan<- e2api.AppInstanceID) {
	s.watchersMu.Lock()
	id := uuid.New()
	s.watchers[id] = ch
	s.watchersMu.Unlock()

	s.instancesMu.RLock()
	instanceIDs := make([]e2api.AppInstanceID, 0, len(s.instances))
	for _, instance := range s.instances {
		instanceIDs = append(instanceIDs, instance.AppInstanceID)
	}
	s.instancesMu.RUnlock()

	go func() {
		for _, instanceID := range instanceIDs {
			ch <- instanceID
		}
		<-ctx.Done()
		s.watchersMu.Lock()
		delete(s.watchers, id)
		s.watchersMu.Unlock()
	}()
}

func (s *appInstanceStreamManager) close(id e2api.AppInstanceID) {
	s.instancesMu.Lock()
	defer s.instancesMu.Unlock()
	if _, ok := s.instances[id]; ok {
		delete(s.instances, id)
		go s.notify(id)
	}
}

type AppInstanceStream struct {
	*TransactionStream
	manager       AppInstanceStreamManager
	streams       NorthboundStreamManager
	AppInstanceID e2api.AppInstanceID
}

func (s *AppInstanceStream) Streams() NorthboundStreamManager {
	return s.streams
}

func (s *AppInstanceStream) open() {

}

func (s *AppInstanceStream) Close() {
	s.manager.close(s.AppInstanceID)
}

var _ Stream = &AppInstanceStream{}
