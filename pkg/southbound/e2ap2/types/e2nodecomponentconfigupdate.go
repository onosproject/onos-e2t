// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package types

import (
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
)

type E2NodeComponentConfigUpdateItem struct {
	E2NodeComponentType         e2ap_ies.E2NodeComponentType
	E2NodeComponentID           *e2ap_ies.E2NodeComponentId
	E2NodeComponentConfigUpdate e2ap_ies.E2NodeComponentConfigUpdate
}

type E2NodeComponentConfigUpdateAckItem struct {
	E2NodeComponentType            e2ap_ies.E2NodeComponentType
	E2NodeComponentID              *e2ap_ies.E2NodeComponentId
	E2NodeComponentConfigUpdateAck E2NodeComponentConfigUpdateAck
}

type E2NodeComponentConfigUpdateAck struct {
	UpdateOutcome int32
	FailureCause  *e2ap_ies.Cause
}
