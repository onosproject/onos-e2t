// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
package pdubuilder

import (
	"encoding/hex"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	"testing"

	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/asn1cgo"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
	"gotest.tools/assert"
)

func TestRicSubscriptionDeleteRequired(t *testing.T) {
	rswcl := make(types.RicSubscriptionWithCauseList)
	rswcl[100] = &types.RicSubscriptionWithCauseItem{
		RicRequestID: types.RicRequest{
			RequestorID: 1,
			InstanceID:  1,
		},
		Cause: &e2ap_ies.Cause{
			Cause: &e2ap_ies.Cause_E2Node{
				E2Node: e2ap_ies.CauseE2Node_CAUSE_E2NODE_E2NODE_COMPONENT_UNKNOWN,
			},
		},
	}
	rswcl[200] = &types.RicSubscriptionWithCauseItem{
		RicRequestID: types.RicRequest{
			RequestorID: 2,
			InstanceID:  12,
		},
		Cause: &e2ap_ies.Cause{
			Cause: &e2ap_ies.Cause_E2Node{
				E2Node: e2ap_ies.CauseE2Node_CAUSE_E2NODE_E2NODE_COMPONENT_UNKNOWN,
			},
		},
	}

	newE2apPdu, err := CreateRicSubscriptionDeleteRequiredE2apPdu(rswcl)
	assert.NilError(t, err)
	assert.Assert(t, newE2apPdu != nil)

	xer, err := asn1cgo.XerEncodeE2apPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("RicSubscriptionDeleteRequired E2AP PDU XER\n%s", string(xer))

	e2apPdu, err := asn1cgo.XerDecodeE2apPdu(xer)
	assert.NilError(t, err)
	assert.DeepEqual(t, newE2apPdu.String(), e2apPdu.String())

	per, err := asn1cgo.PerEncodeE2apPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("RicSubscriptionDeleteRequired E2AP PDU PER\n%v", hex.Dump(per))

	e2apPdu1, err := asn1cgo.PerDecodeE2apPdu(per)
	assert.NilError(t, err)
	assert.DeepEqual(t, newE2apPdu.String(), e2apPdu1.String())
}
