// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package channel

import (
	"github.com/google/uuid"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2appdudescriptions"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2/channel/codec"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2/channel/filter"
	"io"
)

// ReceiverID is a receiver identifier
type ReceiverID string

// Receiver is a channel receiver
type Receiver interface {
	io.Closer
	// ID returns the receiver identifier
	ID() ReceiverID
	// Decode decodes a message using the receiver's codec
	Decode(bytes []byte) (*e2appdudescriptions.E2ApPdu, error)
	// Match returns a boolean indicating whether the given message matches this receiver
	Match(message *e2appdudescriptions.E2ApPdu) bool
	// Receive receives a message
	Receive(message *e2appdudescriptions.E2ApPdu) error
}

// newChannelReceiver creates a new channel receiver
func newChannelReceiver(ch chan<- *e2appdudescriptions.E2ApPdu, filter filter.Filter, codec codec.Codec) Receiver {
	return &channelReceiver{
		id:     ReceiverID(uuid.New().String()),
		ch:     ch,
		filter: filter,
		codec:  codec,
	}
}

// channelReceiver is a channel receiver
type channelReceiver struct {
	id     ReceiverID
	ch     chan<- *e2appdudescriptions.E2ApPdu
	filter filter.Filter
	codec  codec.Codec
}

func (r *channelReceiver) ID() ReceiverID {
	return r.id
}

func (r *channelReceiver) Decode(bytes []byte) (*e2appdudescriptions.E2ApPdu, error) {
	return r.codec.Decode(bytes)
}

func (r *channelReceiver) Match(message *e2appdudescriptions.E2ApPdu) bool {
	return r.filter(message)
}

func (r *channelReceiver) Receive(message *e2appdudescriptions.E2ApPdu) error {
	r.ch <- message
	return nil
}

func (r *channelReceiver) Done() bool {
	return false
}

func (r *channelReceiver) Close() error {
	close(r.ch)
	return nil
}

var _ Receiver = &channelReceiver{}
