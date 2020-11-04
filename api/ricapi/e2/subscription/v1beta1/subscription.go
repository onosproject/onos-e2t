// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package v1beta1

// ID is a subscription identifier
type ID uint64

// Revision is a subscription revision
type Revision uint64

// AppID is an application identifier
type AppID string

// ServiceModelID is a service model identifier
type ServiceModelID string

// Encoding is an encoding type
type Encoding string

const (
	EncodingASN1     Encoding = "asn1"
	EncodingProtobuf Encoding = "proto"
)

// GetPayloadEncoding returns the payload encoding
func (s *Subscription) GetPayloadEncoding() Encoding {
	return Encoding(s.GetPayload().TypeUrl)
}

// GetPayloadBytes returns the payload bytes
func (s *Subscription) GetPayloadBytes() []byte {
	return s.GetPayload().GetValue()
}
