// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package subscription

import (
	subapi "github.com/onosproject/onos-api/go/onos/e2sub/subscription"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

// Test public APIs on the Request Journal
func TestRequestJournal(t *testing.T) {
	requestJournal := NewRequestJournal()
	assert.NotNil(t, requestJournal)

	events := make([]RequestEvent, 0)
	subID := subapi.ID("1")
	req1 := RequestEntry{
		RequestID:    1,
		Subscription: subapi.Subscription{ID: subID},
	}

	reqBeforeAdd := requestJournal.Get(subID)
	assert.NotEqual(t, req1.RequestID, reqBeforeAdd.RequestID)

	ch := make(chan RequestEvent)

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		for e := range ch {
			events = append(events, e)
			wg.Done()
		}
	}()

	closeFunc := requestJournal.Watch(ch)

	requestJournal.Add(subID, req1)
	reqAfterAdd := requestJournal.Get(subID)
	assert.Equal(t, req1.RequestID, reqAfterAdd.RequestID)

	requestJournal.Remove(subID)
	reqAfterDelete := requestJournal.Get(subID)
	assert.NotEqual(t, req1.RequestID, reqAfterDelete.RequestID)

	wg.Done()
	wg.Wait()

	assert.Equal(t, 2, len(events))
	assert.Equal(t, RequestEventAdded, events[0].Type)
	assert.Equal(t, subID, events[0].Record.Subscription.ID)
	assert.Equal(t, req1.RequestID, events[0].Record.RequestID)

	assert.Equal(t, RequestEventRemoved, events[1].Type)
	assert.Equal(t, subID, events[1].Record.Subscription.ID)
	assert.Equal(t, req1.RequestID, events[1].Record.RequestID)

	err := requestJournal.Close()
	assert.NoError(t, err)
	closeFunc()
	close(ch)
}
