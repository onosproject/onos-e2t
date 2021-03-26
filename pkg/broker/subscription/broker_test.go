// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package subscription

import (
	"context"
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
	"github.com/onosproject/onos-lib-go/pkg/errors"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
	"time"
)

func TestStreamBroker(t *testing.T) {
	broker := NewBroker()
	writer, err := broker.GetStream(1)
	assert.True(t, errors.IsNotFound(err))
	assert.Nil(t, writer)
	reader, err := broker.OpenStream("test")
	assert.NoError(t, err)
	writer, err = broker.GetStream(reader.StreamID())
	assert.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		time.Sleep(time.Second)
		cancel()
	}()
	ind, err := reader.Recv(ctx)
	assert.ErrorIs(t, err, context.Canceled)
	assert.Nil(t, ind)

	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = writer.Send(&e2appducontents.Ricindication{})
	assert.NoError(t, err)

	err = writer.Send(&e2appducontents.Ricindication{})
	assert.NoError(t, err)

	ind, err = reader.Recv(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, ind)

	ind, err = reader.Recv(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, ind)

	go func() {
		time.Sleep(time.Second)
		err := writer.Send(&e2appducontents.Ricindication{})
		assert.NoError(t, err)
	}()
	ind, err = reader.Recv(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, ind)

	err = writer.Send(&e2appducontents.Ricindication{})
	assert.NoError(t, err)

	reader, err = broker.CloseStream("test")
	assert.NoError(t, err)

	ind, err = reader.Recv(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, ind)

	ind, err = reader.Recv(ctx)
	assert.ErrorIs(t, err, io.EOF)
	assert.Nil(t, ind)
}
