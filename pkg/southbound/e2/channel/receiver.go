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
	// Done indicates whether the receiver is done
	Done() bool
}

// newStreamReceiver creates a new stream receiver
func newStreamReceiver(ch chan<- *e2appdudescriptions.E2ApPdu, filter filter.Filter, codec codec.Codec) Receiver {
	return &streamReceiver{
		id:     ReceiverID(uuid.New().String()),
		ch:     ch,
		filter: filter,
		codec:  codec,
	}
}

// streamReceiver is a channel receiver
type streamReceiver struct {
	id     ReceiverID
	ch     chan<- *e2appdudescriptions.E2ApPdu
	filter filter.Filter
	codec  codec.Codec
}

func (r *streamReceiver) ID() ReceiverID {
	return r.id
}

func (r *streamReceiver) Decode(bytes []byte) (*e2appdudescriptions.E2ApPdu, error) {
	return r.codec.Decode(bytes)
}

func (r *streamReceiver) Match(message *e2appdudescriptions.E2ApPdu) bool {
	return r.filter(message)
}

func (r *streamReceiver) Receive(message *e2appdudescriptions.E2ApPdu) error {
	r.ch <- message
	return nil
}

func (r *streamReceiver) Done() bool {
	return false
}

func (r *streamReceiver) Close() error {
	close(r.ch)
	return nil
}

var _ io.Closer = &streamReceiver{}

// newUnaryReceiver creates a new unary receiver
func newUnaryReceiver(ch chan<- *e2appdudescriptions.E2ApPdu, filter filter.Filter, codec codec.Codec) Receiver {
	return &unaryReceiver{
		streamReceiver: &streamReceiver{
			ch:     ch,
			filter: filter,
			codec:  codec,
		},
	}
}

// unaryReceiver is a single message receiver
type unaryReceiver struct {
	*streamReceiver
	done bool
}

func (r *unaryReceiver) Receive(message *e2appdudescriptions.E2ApPdu) error {
	err := r.streamReceiver.Receive(message)
	r.done = true
	return err
}

func (r *unaryReceiver) Done() bool {
	return r.done
}
