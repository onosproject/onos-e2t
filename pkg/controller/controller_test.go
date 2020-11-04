// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package controller

import (
	"errors"
	"strconv"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/stretchr/testify/assert"
)

func TestController(t *testing.T) {
	ctrl := gomock.NewController(t)

	activatorValue := &atomic.Value{}
	activator := NewMockActivator(ctrl)
	activator.EXPECT().
		Start(gomock.Any()).
		DoAndReturn(func(ch chan<- bool) error {
			activatorValue.Store(ch)
			return nil
		})
	activator.EXPECT().Stop()

	filter := NewMockFilter(ctrl)
	filter.EXPECT().
		Accept(gomock.Any()).
		DoAndReturn(func(id ID) bool {
			i := id.Int()
			return i%2 == 0
		}).
		AnyTimes()

	wg := &sync.WaitGroup{}
	wg.Add(1)

	watcherValue := &atomic.Value{}
	watcher := NewMockWatcher(ctrl)
	watcher.EXPECT().
		Start(gomock.Any()).
		DoAndReturn(func(ch chan<- ID) error {
			watcherValue.Store(ch)
			wg.Done()
			return nil
		})

	partitions := 3
	partitioner := NewMockWorkPartitioner(ctrl)
	partitioner.EXPECT().
		Partition(gomock.Any()).
		DoAndReturn(func(id ID) (PartitionKey, error) {
			i := id.Int()
			partition := i % partitions
			return PartitionKey(strconv.Itoa(partition)), nil
		}).
		AnyTimes()

	reconciler := NewMockReconciler(ctrl)

	controller := NewController("Test").
		Activate(activator).
		Filter(filter).
		Watch(watcher).
		Partition(partitioner).
		Reconcile(reconciler)
	defer controller.Stop()

	err := controller.Start()
	assert.NoError(t, err)

	activatorCh := activatorValue.Load().(chan<- bool)
	activatorCh <- true

	reconciler.EXPECT().
		Reconcile(gomock.Eq(NewID(2))).
		Return(Result{}, nil)
	reconciler.EXPECT().
		Reconcile(gomock.Eq(NewID(2))).
		Return(Result{}, errors.New("some error"))
	reconciler.EXPECT().
		Reconcile(gomock.Eq(NewID(2))).
		Return(Result{}, errors.New("some error"))
	reconciler.EXPECT().
		Reconcile(gomock.Eq(NewID(2))).
		Return(Result{}, errors.New("some error"))
	reconciler.EXPECT().
		Reconcile(gomock.Eq(NewID(2))).
		Return(Result{}, nil)

	reconciler.EXPECT().
		Reconcile(gomock.Eq(NewID(4))).
		Return(Result{}, errors.New("some error"))
	reconciler.EXPECT().
		Reconcile(gomock.Eq(NewID(4))).
		Return(Result{}, nil)

	wg.Wait()
	watcherCh := watcherValue.Load().(chan<- ID)
	watcherCh <- NewID(1)
	watcherCh <- NewID(2)
	watcherCh <- NewID(3)
	watcherCh <- NewID(4)

}
