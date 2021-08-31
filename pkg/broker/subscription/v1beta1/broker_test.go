// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package v1beta1

import (
	"context"
	"io"
	"testing"
	"time"

	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"

	"github.com/stretchr/testify/assert"
)

func TestBroker(t *testing.T) {
	broker := NewBroker()

	writer, ok := broker.GetWriter(1)
	assert.Nil(t, writer)
	assert.False(t, ok)

	reader1, ok := broker.GetReader("sub-1")
	assert.Nil(t, reader1)
	assert.False(t, ok)

	reader2, ok := broker.GetReader("sub-1")
	assert.Nil(t, reader2)
	assert.False(t, ok)

	reader3, ok := broker.GetReader("sub-1")
	assert.Nil(t, reader3)
	assert.False(t, ok)

	reader4, ok := broker.GetReader("sub-2")
	assert.Nil(t, reader4)
	assert.False(t, ok)

	reader1 = broker.OpenReader("sub-1", "app-1", "instance-1", "transaction-1")
	reader2 = broker.OpenReader("sub-1", "app-1", "instance-2", "transaction-1")
	reader3 = broker.OpenReader("sub-1", "app-2", "instance-1", "transaction-1")
	reader4 = broker.OpenReader("sub-2", "app-2", "instance-1", "transaction-1")

	assert.Equal(t, reader1.ID(), reader2.ID())
	assert.Equal(t, reader2.ID(), reader3.ID())
	assert.NotEqual(t, reader3.ID(), reader4.ID())

	reader1, ok = broker.GetReader("sub-1")
	assert.NotNil(t, reader1)
	assert.True(t, ok)

	reader2, ok = broker.GetReader("sub-1")
	assert.NotNil(t, reader2)
	assert.True(t, ok)

	reader3, ok = broker.GetReader("sub-1")
	assert.NotNil(t, reader3)
	assert.True(t, ok)

	reader4, ok = broker.GetReader("sub-2")
	assert.NotNil(t, reader4)
	assert.True(t, ok)

	reader1 = broker.OpenReader("sub-1", "app-1", "instance-1", "transaction-1")
	reader2 = broker.OpenReader("sub-1", "app-1", "instance-2", "transaction-1")
	reader3 = broker.OpenReader("sub-1", "app-2", "instance-1", "transaction-1")
	reader4 = broker.OpenReader("sub-2", "app-2", "instance-1", "transaction-1")

	writer1, ok := broker.GetWriter(reader1.ID())
	assert.NotNil(t, writer1)
	assert.True(t, ok)

	err := writer1.Send(&e2api.Indication{})
	assert.NoError(t, err)
	err = writer1.Send(&e2api.Indication{})
	assert.NoError(t, err)

	writer4, ok := broker.GetWriter(reader4.ID())
	assert.NotNil(t, writer4)
	assert.True(t, ok)

	err = writer4.Send(&e2api.Indication{})
	assert.NoError(t, err)
	err = writer4.Send(&e2api.Indication{})
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
		err := writer4.Send(&e2api.Indication{})
		assert.NoError(t, err)
	}()

	ind, err = reader4.Recv(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, ind)

	broker.CloseReader("sub-2", "app-2", "instance-1", "transaction-1")

	ind, err = reader4.Recv(context.Background())
	assert.Nil(t, ind)
	assert.ErrorIs(t, err, io.EOF)
}
