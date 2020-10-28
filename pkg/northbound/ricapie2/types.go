// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package ricapie2

// EncodingType encoding types
type EncodingType int

const (
	// ASN.1 encoding
	ASN1 EncodingType = iota
	// Protobuf encoding
	PROTO
)

func (e EncodingType) String() string {
	return [...]string{"ASN1", "PROTO"}[e]
}

type ServiceModelID int

const (
	KPM ServiceModelID = iota
	IN
)

func (m ServiceModelID) String() string {
	return [...]string{"KPM", "IN"}[m]
}

type RequestInfo struct {
	EncodingType   EncodingType
	ServiceModelID ServiceModelID
	Payload        []byte
}
