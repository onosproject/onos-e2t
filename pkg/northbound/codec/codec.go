// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package codec

import (
	"github.com/gogo/protobuf/proto"
	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2"
	e2appdudescriptions "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-descriptions"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/asn1cgo"
)

// getCodec gets the codec for the given encoding
func getCodec(encoding e2api.EncodingType) Codec {
	switch encoding {
	case e2api.EncodingType_PROTO:
		return Proto
	case e2api.EncodingType_ASN1_PER:
		return PER
	case e2api.EncodingType_ASN1_XER:
		return XER
	default:
		panic("encountered unexpected encoding")
	}
}

// Encode encodes the given message using the given encoding
func Encode(message *e2appdudescriptions.E2ApPdu, encoding e2api.EncodingType) ([]byte, error) {
	return getCodec(encoding).Encode(message)
}

// Decode decodes the given message using the given encoding
func Decode(bytes []byte, encoding e2api.EncodingType) (*e2appdudescriptions.E2ApPdu, error) {
	return getCodec(encoding).Decode(bytes)
}

// Codec is a channel codec
type Codec interface {
	// Encode encodes the given message
	Encode(message *e2appdudescriptions.E2ApPdu) ([]byte, error)
	// Decode decodes the given message
	Decode(bytes []byte) (*e2appdudescriptions.E2ApPdu, error)
}

// Proto is a codec that uses Protobuf encoding
var Proto = &ProtoCodec{}

// ProtoCodec is a codec that uses Protobuf encoding
type ProtoCodec struct{}

func (c *ProtoCodec) Encode(message *e2appdudescriptions.E2ApPdu) ([]byte, error) {
	return proto.Marshal(message)
}

func (c *ProtoCodec) Decode(bytes []byte) (*e2appdudescriptions.E2ApPdu, error) {
	pdu := &e2appdudescriptions.E2ApPdu{}
	if err := proto.Unmarshal(bytes, pdu); err != nil {
		return nil, err
	}
	return pdu, nil
}

var _ Codec = &ProtoCodec{}

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
