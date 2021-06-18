// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package v1beta1

import (
	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"io"
	"sync"
)

var log = logging.GetLogger("broker", "subscription", "v1beta1")

// NewBroker creates a new subscription stream broker
func NewBroker() Broker {
	return &streamBroker{
		streams: newStreamRegistry(),
	}
}

// Broker is a subscription stream broker
// The Broker is responsible for managing Streams for propagating indications from the southbound API
// to the northbound API.
type Broker interface {
	io.Closer

	// OpenReader opens a subscription StreamReader
	// If a stream already exists for the subscription, the existing stream will be returned.
	// If no stream exists, a new stream will be allocated with a unique StreamID.
	OpenReader(subID e2api.SubscriptionID, appID e2api.AppID, instanceID e2api.AppInstanceID) StreamReader

	// CloseReader closes a subscription StreamReader
	CloseReader(subID e2api.SubscriptionID, appID e2api.AppID, instanceID e2api.AppInstanceID)

	// GetReader gets a read stream by its StreamID
	// If no StreamReader exists for the given StreamID, ok will be false
	GetReader(subID e2api.SubscriptionID) (reader StreamReader, ok bool)

	// GetWriter gets a write stream by its StreamID
	// If no StreamWriter exists for the given StreamID, ok will be false
	GetWriter(streamID StreamID) (writer StreamWriter, ok bool)
}

type streamBroker struct {
	streams *streamRegistry
	mu      sync.RWMutex
}

func (b *streamBroker) OpenReader(subID e2api.SubscriptionID, appID e2api.AppID, instanceID e2api.AppInstanceID) StreamReader {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.streams.openSubStream(subID).
		openAppStream(appID).
		openInstanceStream(instanceID)
}

func (b *streamBroker) CloseReader(subID e2api.SubscriptionID, appID e2api.AppID, instanceID e2api.AppInstanceID) {
	b.mu.Lock()
	defer b.mu.Unlock()
	subStream, ok := b.streams.getSubStream(subID)
	if !ok {
		return
	}
	appStream, ok := subStream.getAppStream(appID)
	if !ok {
		return
	}
	instanceStream, ok := appStream.getInstanceStream(instanceID)
	if !ok {
		return
	}
	log.Infof("Closed reader for subscription %s", subID)
	instanceStream.Close()
}

func (b *streamBroker) GetReader(subID e2api.SubscriptionID) (StreamReader, bool) {
	b.mu.RLock()
	defer b.mu.RUnlock()
	stream, ok := b.streams.getSubStream(subID)
	return stream, ok
}

func (b *streamBroker) GetWriter(streamID StreamID) (StreamWriter, bool) {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return b.streams.getStream(streamID)
}

func (b *streamBroker) Close() error {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.streams.Close()
}

var _ Broker = &streamBroker{}
