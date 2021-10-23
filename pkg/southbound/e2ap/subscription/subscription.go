// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package subscription

import (
	"context"
	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
)

type StreamID int

type Stream interface {
	ID() StreamID
	Subscription() *e2api.Subscription
	Context() context.Context
	Reader() Reader
	Writer() Writer
}

type Reader interface {
	Indications() <-chan *e2appducontents.Ricindication
	Err() error
}

type Writer interface {
	Indications() chan<- *e2appducontents.Ricindication
	Close(err error)
}
