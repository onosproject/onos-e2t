// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package types

import (
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
)

type E2ConnectionUpdateItem struct {
	TnlInformation TnlInformation
	TnlUsage       e2ap_ies.Tnlusage
}

type E2ConnectionSetupFailedItem struct {
	TnlInformation TnlInformation
	Cause          e2ap_ies.Cause
}
