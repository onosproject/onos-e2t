// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
package pdubuilder

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/encoder"
	"testing"

	"github.com/onosproject/onos-e2t/api/e2ap/v2"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-commondatatypes"
	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
	"gotest.tools/assert"
)

func TestRicServiceUpdateFailure(t *testing.T) {
	rfRejected := make(types.RanFunctionCauses)
	rfRejected[101] = &e2apies.Cause{
		Cause: &e2apies.Cause_Misc{
			Misc: e2apies.CauseMisc_CAUSE_MISC_HARDWARE_FAILURE,
		},
	}
	rfRejected[102] = &e2apies.Cause{
		Cause: &e2apies.Cause_Protocol{
			Protocol: e2apies.CauseProtocol_CAUSE_PROTOCOL_SEMANTIC_ERROR,
		},
	}

	ttw := e2apies.TimeToWait_TIME_TO_WAIT_V2S
	procCode := v2.ProcedureCodeIDRICsubscription
	criticality := e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE
	ftg := e2ap_commondatatypes.TriggeringMessage_TRIGGERING_MESSAGE_UNSUCCESSFUL_OUTCOME

	newE2apPdu, err := CreateRicServiceUpdateFailureE2apPdu(1, &e2apies.Cause{
		Cause: &e2apies.Cause_RicService{
			RicService: e2apies.CauseRicservice_CAUSE_RICSERVICE_RIC_RESOURCE_LIMIT,
		},
	})
	assert.NilError(t, err)
	assert.Assert(t, newE2apPdu != nil)
	newE2apPdu.GetUnsuccessfulOutcome().GetValue().GetRicServiceUpdate().
		SetTimeToWait(ttw).SetCriticalityDiagnostics(&procCode, &criticality, &ftg,
		&types.RicRequest{
			RequestorID: 10,
			InstanceID:  20,
		}, []*types.CritDiag{
			{
				TypeOfError:   e2apies.TypeOfError_TYPE_OF_ERROR_MISSING,
				IECriticality: e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE,
				IEId:          v2.ProtocolIeIDRicsubscriptionDetails,
			},
		})

	assert.NilError(t, err)
	assert.Assert(t, newE2apPdu != nil)

	perNew, err := encoder.PerEncodeE2ApPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("RicServiceUpdateFailure E2AP PDU PER with Go APER library\n%v", hex.Dump(perNew))

	//Comparing reference PER bytes with Go APER library produced
	//assert.DeepEqual(t, per, perNew)

	result, err := encoder.PerDecodeE2ApPdu(perNew)
	assert.NilError(t, err)
	assert.DeepEqual(t, newE2apPdu.String(), result.String())
}

func TestRicServiceUpdateFailureExcludeOptionalIE(t *testing.T) {
	ttw := e2apies.TimeToWait_TIME_TO_WAIT_V2S

	newE2apPdu, err := CreateRicServiceUpdateFailureE2apPdu(1, &e2apies.Cause{
		Cause: &e2apies.Cause_RicService{
			RicService: e2apies.CauseRicservice_CAUSE_RICSERVICE_RIC_RESOURCE_LIMIT,
		},
	})
	assert.NilError(t, err)
	assert.Assert(t, newE2apPdu != nil)
	newE2apPdu.GetUnsuccessfulOutcome().GetValue().GetRicServiceUpdate().SetTimeToWait(ttw)

	perNew, err := encoder.PerEncodeE2ApPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("RicServiceUpdateFailure E2AP PDU PER with Go APER library\n%v", hex.Dump(perNew))

	//Comparing reference PER bytes with Go APER library produced
	//assert.DeepEqual(t, per, perNew)

	result, err := encoder.PerDecodeE2ApPdu(perNew)
	assert.NilError(t, err)
	assert.DeepEqual(t, newE2apPdu.String(), result.String())
}
