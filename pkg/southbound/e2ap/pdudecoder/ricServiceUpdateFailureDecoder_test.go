// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdudecoder

import (
	v2 "github.com/onosproject/onos-e2t/api/e2ap/v2"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-commondatatypes"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/pdubuilder"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
	"gotest.tools/assert"
	"testing"
)

func Test_DecodeRicServiceUpdateFailurePdu(t *testing.T) {
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

	procCode := v2.ProcedureCodeIDRICsubscription
	criticality := e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE
	ftg := e2ap_commondatatypes.TriggeringMessage_TRIGGERING_MESSAGE_UNSUCCESSFUL_OUTCOME

	e2apPdu, err := pdubuilder.CreateRicServiceUpdateFailureE2apPdu(1, &e2ap_ies.Cause{
		Cause: &e2ap_ies.Cause_RicService{
			RicService: e2ap_ies.CauseRicservice_CAUSE_RICSERVICE_RIC_RESOURCE_LIMIT,
		},
	})
	assert.NilError(t, err)
	assert.Assert(t, e2apPdu != nil)
	e2apPdu.GetUnsuccessfulOutcome().GetValue().GetRicServiceUpdate().
		SetTimeToWait(e2ap_ies.TimeToWait_TIME_TO_WAIT_V2S).SetCriticalityDiagnostics(&procCode, &criticality, &ftg,
		&types.RicRequest{
			RequestorID: 10,
			InstanceID:  20,
		}, []*types.CritDiag{
			{
				TypeOfError:   e2ap_ies.TypeOfError_TYPE_OF_ERROR_MISSING,
				IECriticality: e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE,
				IEId:          v2.ProtocolIeIDRicsubscriptionDetails,
			},
		})

	transactionID, cause, ttw, pr, crit, tm, cdrID, diags, err := DecodeRicServiceUpdateFailurePdu(e2apPdu)
	assert.NilError(t, err)
	//assert.Assert(t, ricIdentity != nil) //Commented due to the Linters (v1.34.1) error - possible nil pointer dereference (https://staticcheck.io/docs/checks#SA5011) on lines 23, 24 & 25

	assert.Equal(t, cause.GetRicService(), e2ap_ies.CauseRicservice_CAUSE_RICSERVICE_RIC_RESOURCE_LIMIT)
	assert.Equal(t, int32(*ttw), int32(e2ap_ies.TimeToWait_TIME_TO_WAIT_V2S))
	assert.Equal(t, int32(*pr), int32(8))
	assert.Equal(t, int32(*crit), int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE))
	assert.Equal(t, int32(*tm), int32(e2ap_commondatatypes.TriggeringMessage_TRIGGERING_MESSAGE_UNSUCCESSFUL_OUTCOME))
	assert.Equal(t, int32(cdrID.InstanceID), int32(20))
	assert.Equal(t, int32(cdrID.RequestorID), int32(10))
	assert.Equal(t, int32(diags[0].IEId), int32(30))
	assert.Equal(t, int32(diags[0].IECriticality), int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE))
	assert.Equal(t, int32(diags[0].TypeOfError), int32(e2ap_ies.TypeOfError_TYPE_OF_ERROR_MISSING))
	if transactionID != nil {
		assert.Equal(t, int32(1), *transactionID)
	}
}
