// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package stream

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestStreamManager(t *testing.T) {
	mgr := NewManager()

	ch := make(chan Stream)
	err := mgr.Watch(context.Background(), ch)
	assert.NoError(t, err)

	streamCh := make(chan Message)
	stream, err := mgr.Open(context.Background(), ID("1"), streamCh)
	assert.NoError(t, err)
	assert.NotNil(t, stream)

	select {
	case event := <-ch:
		assert.NotNil(t, event)
		assert.Equal(t, ID("1"), event.ID())
	case <-time.After(time.Second):
		t.FailNow()
	}
}
