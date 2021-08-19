// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package types

import (
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
)

type TnlInformation struct {
	TnlPort    asn1.BitString
	TnlAddress asn1.BitString
}
