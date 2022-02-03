// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package stream

import (
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
)

type Stream interface {
	Ready() <-chan struct{}
	Indications() <-chan *e2appducontents.Ricindication
	Done() <-chan struct{}
	Err() error
}

func newStream(c *channelOutput) Stream {
	return &channelStream{
		channelOutput: c,
	}
}

type channelStream struct {
	*channelOutput
}

func (c *channelStream) Indications() <-chan *e2appducontents.Ricindication {
	return c.buffer.out
}
