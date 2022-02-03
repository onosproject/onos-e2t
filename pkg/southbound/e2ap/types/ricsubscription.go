// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package types

import (
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
)

type RicSubscriptionWithCauseItem struct {
	RicRequestID RicRequest
	Cause        *e2ap_ies.Cause
}

type RicSubscriptionWithCauseList map[RanFunctionID]*RicSubscriptionWithCauseItem
