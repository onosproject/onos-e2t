// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package codec

import (
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2appdudescriptions"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/asn1cgo"
)

// Codec is a channel codec
type Codec interface {
	// Encode encodes the given message
	Encode(message *e2appdudescriptions.E2ApPdu) ([]byte, error)
	// Decode decodes the given message
	Decode(bytes []byte) (*e2appdudescriptions.E2ApPdu, error)
}

// PER is a codec that uses PER encoding
var PER = &PERCodec{}

// PERCodec is a codec that uses PER encoding
type PERCodec struct{}

func (c *PERCodec) Encode(message *e2appdudescriptions.E2ApPdu) ([]byte, error) {
	return asn1cgo.PerEncodeE2apPdu(message)
}

func (c *PERCodec) Decode(bytes []byte) (*e2appdudescriptions.E2ApPdu, error) {
	return asn1cgo.PerDecodeE2apPdu(bytes)
}

var _ Codec = &PERCodec{}

// XER is a codec that uses XER encoding
var XER = &XERCodec{}

// XERCodec is a codec that uses XER encoding
type XERCodec struct{}

func (c *XERCodec) Encode(message *e2appdudescriptions.E2ApPdu) ([]byte, error) {
	return asn1cgo.XerEncodeE2apPdu(message)
}

func (c *XERCodec) Decode(bytes []byte) (*e2appdudescriptions.E2ApPdu, error) {
	return asn1cgo.XerDecodeE2apPdu(bytes)
}

var _ Codec = &XERCodec{}
