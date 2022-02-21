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

func Test_DecodeRicServiceUpdateAcknowledgePdu(t *testing.T) {
	rfAccepted := make(types.RanFunctionRevisions)
	rfAccepted[100] = 2
	rfAccepted[200] = 2

	rfRejected := make(types.RanFunctionCauses)
	rfRejected[101] = &e2ap_ies.Cause{
		Cause: &e2ap_ies.Cause_Misc{
			Misc: e2ap_ies.CauseMisc_CAUSE_MISC_HARDWARE_FAILURE,
		},
	}
	rfRejected[102] = &e2ap_ies.Cause{
		Cause: &e2ap_ies.Cause_Protocol{
			Protocol: e2ap_ies.CauseProtocol_CAUSE_PROTOCOL_SEMANTIC_ERROR,
		},
	}

	e2apPdu, err := pdubuilder.CreateRicServiceUpdateAcknowledgeE2apPdu(1, rfAccepted)
	assert.NilError(t, err)
	assert.Assert(t, e2apPdu != nil)
	e2apPdu.GetSuccessfulOutcome().GetValue().GetRicServiceUpdate().
		SetRanFunctionsRejected(rfRejected)

	transactionID, ranFunctionsAccepted, causes, err := DecodeRicServiceUpdateAcknowledgePdu(e2apPdu)
	assert.NilError(t, err)

	assert.Equal(t, 2, len(ranFunctionsAccepted))
	rfa100, ok := ranFunctionsAccepted[100]
	assert.Assert(t, ok, "expected a key '100'")
	assert.Equal(t, 2, int(rfa100))
	rfa200, ok := ranFunctionsAccepted[200]
	assert.Assert(t, ok, "expected a key '200'")
	assert.Equal(t, 2, int(rfa200))

	assert.Assert(t, causes != nil)
	for id, cause := range causes {
		switch id {
		case 101:
			assert.Equal(t, e2ap_ies.CauseMisc_CAUSE_MISC_HARDWARE_FAILURE, cause.GetMisc())
		case 102:
			assert.Equal(t, e2ap_ies.CauseProtocol_CAUSE_PROTOCOL_SEMANTIC_ERROR, cause.GetProtocol())
		default:
			assert.Assert(t, false, "unexpected cause %d", id)
		}
	}

	if transactionID != nil {
		assert.Equal(t, int32(1), *transactionID)
	}
}
