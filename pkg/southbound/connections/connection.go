// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package connections

// SctpConnection represents an ASN1 SCTP connection
type SctpConnection struct {
	ID              uint32
	PlmnID          string
	RemoteIPAddress string
	RemotePort      uint32
}
