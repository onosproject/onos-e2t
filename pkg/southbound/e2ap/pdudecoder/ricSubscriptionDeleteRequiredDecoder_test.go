// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdudecoder

import (
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/pdubuilder"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
	"gotest.tools/assert"
	"testing"
)

func Test_DecodeRicSubscriptionDeleteRequiredPdu(t *testing.T) {
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

	e2apPdu, err := pdubuilder.CreateRicSubscriptionDeleteRequiredE2apPdu(rswcl)
	assert.NilError(t, err)
	assert.Assert(t, e2apPdu != nil)

	rswclBack, err := DecodeRicSubscriptionDeleteRequiredPdu(e2apPdu)
	assert.NilError(t, err)
	assert.Assert(t, rswclBack != nil) //Commented due to the Linters (v1.34.1) error - possible nil pointer dereference (https://staticcheck.io/docs/checks#SA5011) on line 23

	assert.Assert(t, rswclBack != nil)
	assert.Equal(t, 2, len(rswclBack))

	for id, item := range rswclBack {
		switch id {
		case 100:
			assert.Equal(t, e2ap_ies.CauseE2Node_CAUSE_E2NODE_E2NODE_COMPONENT_UNKNOWN.Number(), item.Cause.GetE2Node().Number())
			assert.Equal(t, 1, int(item.RicRequestID.RequestorID))
			assert.Equal(t, 1, int(item.RicRequestID.InstanceID))
		case 200:
			assert.Equal(t, e2ap_ies.CauseE2Node_CAUSE_E2NODE_E2NODE_COMPONENT_UNKNOWN.Number(), item.Cause.GetE2Node().Number())
			assert.Equal(t, 2, int(item.RicRequestID.RequestorID))
			assert.Equal(t, 12, int(item.RicRequestID.InstanceID))
		default:
			assert.Assert(t, false, "unexpected cause %d", id)
		}
	}
}
