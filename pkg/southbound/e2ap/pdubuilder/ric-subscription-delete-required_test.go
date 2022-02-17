// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
package pdubuilder

import (
	"encoding/hex"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/encoder"
	"testing"

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

	perNew, err := encoder.PerEncodeE2ApPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("RicSubscriptionDeleteRequired E2AP PDU PER with Go APER library\n%v", hex.Dump(perNew))

	//Comparing reference PER bytes with Go APER library produced
	//assert.DeepEqual(t, per, perNew)

	result, err := encoder.PerDecodeE2ApPdu(perNew)
	assert.NilError(t, err)
	assert.DeepEqual(t, newE2apPdu.String(), result.String())
}
