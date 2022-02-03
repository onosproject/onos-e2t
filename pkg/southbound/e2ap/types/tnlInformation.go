// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
)

type TnlInformation struct {
	TnlPort    *asn1.BitString // optional structure
	TnlAddress asn1.BitString
}

type TnlAssociationRemovalItem struct {
	TnlInformation    TnlInformation
	TnlInformationRic TnlInformation
}
