// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package v1beta1

import (
	"context"
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
	"github.com/onosproject/onos-e2t/api/onos/e2t/store/subscription"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestBroker(t *testing.T) {
	broker := NewBroker()

	sub1ID := subscription.SubscriptionID{
		NodeID:     "node-1",
		RequestID:  "request-1",
		AppID:      "app-1",
		InstanceID: "instance-1",
	}
	sub2ID := subscription.SubscriptionID{
		NodeID:     "node-1",
		RequestID:  "request-1",
		AppID:      "app-1",
		InstanceID: "instance-2",
	}
	sub3ID := subscription.SubscriptionID{
		NodeID:     "node-1",
		RequestID:  "request-1",
		AppID:      "app-2",
		InstanceID: "instance-1",
	}
	sub4ID := subscription.SubscriptionID{
		NodeID:     "node-1",
		RequestID:  "request-2",
		AppID:      "app-2",
		InstanceID: "instance-1",
	}

	writer, ok := broker.GetWriter(1)
	assert.Nil(t, writer)
	assert.False(t, ok)

	reader1, ok := broker.GetReader(sub1ID.TaskID())
	assert.Nil(t, reader1)
	assert.False(t, ok)

	reader2, ok := broker.GetReader(sub2ID.TaskID())
	assert.Nil(t, reader2)
	assert.False(t, ok)

	reader3, ok := broker.GetReader(sub3ID.TaskID())
	assert.Nil(t, reader3)
	assert.False(t, ok)

	reader4, ok := broker.GetReader(sub4ID.TaskID())
	assert.Nil(t, reader4)
	assert.False(t, ok)

	reader1 = broker.OpenReader(sub1ID)
	reader2 = broker.OpenReader(sub2ID)
	reader3 = broker.OpenReader(sub3ID)
	reader4 = broker.OpenReader(sub4ID)

	assert.Equal(t, reader1.ID(), reader2.ID())
	assert.Equal(t, reader2.ID(), reader3.ID())
	assert.NotEqual(t, reader3.ID(), reader4.ID())

	reader1, ok = broker.GetReader(sub1ID.TaskID())
	assert.NotNil(t, reader1)
	assert.True(t, ok)

	reader2, ok = broker.GetReader(sub2ID.TaskID())
	assert.NotNil(t, reader2)
	assert.True(t, ok)

	reader3, ok = broker.GetReader(sub3ID.TaskID())
	assert.NotNil(t, reader3)
	assert.True(t, ok)

	reader4, ok = broker.GetReader(sub4ID.TaskID())
	assert.NotNil(t, reader4)
	assert.True(t, ok)

	reader1 = broker.OpenReader(sub1ID)
	reader2 = broker.OpenReader(sub2ID)
	reader3 = broker.OpenReader(sub3ID)
	reader4 = broker.OpenReader(sub4ID)

	writer1, ok := broker.GetWriter(reader1.ID())
	assert.NotNil(t, writer1)
	assert.True(t, ok)

	err := writer1.Send(&e2appducontents.Ricindication{})
	assert.NoError(t, err)
	err = writer1.Send(&e2appducontents.Ricindication{})
	assert.NoError(t, err)

	writer4, ok := broker.GetWriter(reader4.ID())
	assert.NotNil(t, writer4)
	assert.True(t, ok)

	err = writer4.Send(&e2appducontents.Ricindication{})
	assert.NoError(t, err)
	err = writer4.Send(&e2appducontents.Ricindication{})
	assert.NoError(t, err)

	ind, err := reader1.Recv(context.TODO())
	assert.NoError(t, err)
	assert.NotNil(t, ind)

	ind, err = reader2.Recv(context.TODO())
	assert.NoError(t, err)
	assert.NotNil(t, ind)

	ind, err = reader3.Recv(context.TODO())
	assert.NoError(t, err)
	assert.NotNil(t, ind)

	ind, err = reader3.Recv(context.TODO())
	assert.NoError(t, err)
	assert.NotNil(t, ind)

	ind, err = reader4.Recv(context.TODO())
	assert.NoError(t, err)
	assert.NotNil(t, ind)

	ind, err = reader4.Recv(context.TODO())
	assert.NoError(t, err)
	assert.NotNil(t, ind)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		time.Sleep(time.Second)
		cancel()
	}()
	ind, err = reader2.Recv(ctx)
	assert.ErrorIs(t, err, context.Canceled)
	assert.Nil(t, ind)

	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	go func() {
		time.Sleep(time.Second)
		err := writer4.Send(&e2appducontents.Ricindication{})
		assert.NoError(t, err)
	}()
	ind, err = reader4.Recv(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, ind)

	err = writer4.Send(&e2appducontents.Ricindication{})
	assert.NoError(t, err)

	reader4.Close()
}
