// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package subscription

import (
	"context"
	"github.com/atomix/atomix-go-client/pkg/atomix/test"
	"github.com/atomix/atomix-go-client/pkg/atomix/test/rsm"
	api "github.com/onosproject/onos-e2t/api/onos/e2t/store/subscription"
	"github.com/onosproject/onos-lib-go/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestSubscriptionStore(t *testing.T) {
	test := test.NewTest(
		rsm.NewProtocol(),
		test.WithReplicas(1),
		test.WithPartitions(1))
	assert.NoError(t, test.Start())
	defer test.Stop()

	client1, err := test.NewClient("node-1")
	assert.NoError(t, err)

	client2, err := test.NewClient("node-2")
	assert.NoError(t, err)

	store1, err := NewAtomixStore(client1)
	assert.NoError(t, err)

	store2, err := NewAtomixStore(client2)
	assert.NoError(t, err)

	ch := make(chan api.SubscriptionEvent)
	err = store2.Watch(context.Background(), ch)
	assert.NoError(t, err)

	obj1 := &api.Subscription{
		SubscriptionMeta: api.SubscriptionMeta{
			ID: api.SubscriptionID{
				NodeID:     "node-1",
				RequestID:  "request-1",
				AppID:      "app-1",
				InstanceID: "instance-1",
				Hash:       "foo",
			},
		},
	}
	obj2 := &api.Subscription{
		SubscriptionMeta: api.SubscriptionMeta{
			ID: api.SubscriptionID{
				NodeID:     "node-1",
				RequestID:  "request-1",
				AppID:      "app-2",
				InstanceID: "instance-1",
				Hash:       "bar",
			},
		},
	}

	// Create a new object
	err = store1.Create(context.TODO(), obj1)
	assert.NoError(t, err)
	assert.NotEqual(t, api.Revision(0), obj1.Revision)

	// Get the object
	obj, err := store2.Get(context.TODO(), obj1.ID)
	assert.NoError(t, err)
	assert.NotNil(t, obj)
	assert.Equal(t, obj1.ID, obj.ID)
	assert.NotEqual(t, api.Revision(0), obj.Revision)

	// Create another object
	err = store2.Create(context.TODO(), obj2)
	assert.NoError(t, err)
	assert.NotEqual(t, api.Revision(0), obj2.Revision)

	// Verify events were received for the objects
	event := nextEvent(t, ch)
	assert.Equal(t, obj1.ID, event.ID)
	event = nextEvent(t, ch)
	assert.Equal(t, obj2.ID, event.ID)

	// Delete the object
	err = store1.Delete(context.TODO(), obj)
	assert.NoError(t, err)

	// Verify the object was deleted
	obj, err = store2.Get(context.TODO(), obj1.ID)
	assert.Nil(t, obj)
	assert.True(t, errors.IsNotFound(err))
}

func nextEvent(t *testing.T, ch chan api.SubscriptionEvent) *api.Subscription {
	select {
	case c := <-ch:
		return &c.Subscription
	case <-time.After(5 * time.Second):
		t.FailNow()
	}
	return nil
}
