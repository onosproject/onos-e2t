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

func TestE2connectionUpdateFailure(t *testing.T) {
	//ttw1 := e2ap_ies.TimeToWait_TIME_TO_WAIT_V5S
	//procCode1 := v21.ProcedureCodeIDRICsubscription
	//criticality1 := e2apcommondatatypes.Criticality_CRITICALITY_IGNORE
	//ftg1 := e2apcommondatatypes.TriggeringMessage_TRIGGERING_MESSAGE_UNSUCCESSFUL_OUTCOME
	//
	//e2apPdu, err := pdubuilder.CreateE2connectionUpdateFailureE2apPdu(1)
	//assert.NilError(t, err)
	//assert.Assert(t, e2apPdu != nil)
	//e2apPdu.GetUnsuccessfulOutcome().GetProcedureCode().GetE2ConnectionUpdate().GetUnsuccessfulOutcome().
	//	SetCause(&e2ap_ies.Cause{
	//		Cause: &e2ap_ies.Cause_Protocol{
	//			Protocol: e2ap_ies.CauseProtocol_CAUSE_PROTOCOL_TRANSFER_SYNTAX_ERROR,
	//		},
	//	}).SetTimeToWait(ttw1).SetCriticalityDiagnostics(&procCode1, &criticality1, &ftg1,
	//	&types1.RicRequest{
	//		RequestorID: 10,
	//		InstanceID:  20,
	//	}, []*types1.CritDiag{
	//		{
	//			TypeOfError:   e2ap_ies.TypeOfError_TYPE_OF_ERROR_MISSING,
	//			IECriticality: e2apcommondatatypes.Criticality_CRITICALITY_IGNORE,
	//			IEId:          v21.ProtocolIeIDRicsubscriptionDetails,
	//		},
	//	})
	//
	//per, err := asn1cgo.PerEncodeE2apPdu(e2apPdu)
	//assert.NilError(t, err)
	//t.Logf("E2connectionUpdateFailure E2AP PDU PER\n%v", hex.Dump(per))

	ttw := e2apies.TimeToWait_TIME_TO_WAIT_V5S
	procCode := v2.ProcedureCodeIDRICsubscription
	criticality := e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE
	ftg := e2ap_commondatatypes.TriggeringMessage_TRIGGERING_MESSAGE_UNSUCCESSFUL_OUTCOME

	newE2apPdu, err := CreateE2connectionUpdateFailureE2apPdu(1)
	assert.NilError(t, err)
	assert.Assert(t, newE2apPdu != nil)
	newE2apPdu.GetUnsuccessfulOutcome().GetValue().GetE2ConnectionUpdate().
		SetCause(&e2apies.Cause{
			Cause: &e2apies.Cause_Protocol{
				Protocol: e2apies.CauseProtocol_CAUSE_PROTOCOL_TRANSFER_SYNTAX_ERROR,
			},
		}).SetTimeToWait(ttw).SetCriticalityDiagnostics(&procCode, &criticality, &ftg,
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

	perNew, err := encoder.PerEncodeE2ApPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("E2connectionUpdateFailure E2AP PDU PER with Go APER library\n%v", hex.Dump(perNew))

	//Comparing reference PER bytes with Go APER library produced
	//assert.DeepEqual(t, per, perNew)

	result, err := encoder.PerDecodeE2ApPdu(perNew)
	assert.NilError(t, err)
	assert.DeepEqual(t, newE2apPdu.String(), result.String())

	//// Decoding the message from the APER bytes produced by CGo
	//result11, err := encoder.PerDecodeE2ApPdu(per)
	//assert.NilError(t, err)
	//assert.DeepEqual(t, newE2apPdu.String(), result11.String())
	//
	//result1, err := asn1cgo.PerDecodeE2apPdu(perNew)
	//assert.NilError(t, err)
	//assert.Assert(t, result1 != nil)
	//t.Logf("E2connectionUpdateFailure E2AP PDU PER - decoded is \n%v", result1)
	//assert.DeepEqual(t, e2apPdu.String(), result1.String())
}

func TestE2connectionUpdateFailureExcludeOptionalIE(t *testing.T) {
	//ttw1 := e2ap_ies.TimeToWait_TIME_TO_WAIT_V5S
	//e2apPdu, err := pdubuilder.CreateE2connectionUpdateFailureE2apPdu(1)
	//assert.NilError(t, err)
	//assert.Assert(t, e2apPdu != nil)
	//e2apPdu.GetUnsuccessfulOutcome().GetProcedureCode().GetE2ConnectionUpdate().GetUnsuccessfulOutcome().
	//	SetCause(&e2ap_ies.Cause{
	//		Cause: &e2ap_ies.Cause_Protocol{
	//			Protocol: e2ap_ies.CauseProtocol_CAUSE_PROTOCOL_TRANSFER_SYNTAX_ERROR,
	//		},
	//	}).SetTimeToWait(ttw1)
	//
	//per, err := asn1cgo.PerEncodeE2apPdu(e2apPdu)
	//assert.NilError(t, err)
	//t.Logf("E2connectionUpdateFailure E2AP PDU PER\n%v", hex.Dump(per))

	ttw := e2apies.TimeToWait_TIME_TO_WAIT_V5S
	newE2apPdu, err := CreateE2connectionUpdateFailureE2apPdu(1)
	assert.NilError(t, err)
	assert.Assert(t, newE2apPdu != nil)
	newE2apPdu.GetUnsuccessfulOutcome().GetValue().GetE2ConnectionUpdate().
		SetCause(&e2apies.Cause{
			Cause: &e2apies.Cause_Protocol{
				Protocol: e2apies.CauseProtocol_CAUSE_PROTOCOL_TRANSFER_SYNTAX_ERROR,
			},
		}).SetTimeToWait(ttw)

	perNew, err := encoder.PerEncodeE2ApPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("E2connectionUpdateFailure E2AP PDU PER with Go APER library\n%v", hex.Dump(perNew))

	//Comparing reference PER bytes with Go APER library produced
	//assert.DeepEqual(t, per, perNew)

	result, err := encoder.PerDecodeE2ApPdu(perNew)
	assert.NilError(t, err)
	assert.DeepEqual(t, newE2apPdu.String(), result.String())

	// Decoding the message from the APER bytes produced by CGo
	//result11, err := encoder.PerDecodeE2ApPdu(per)
	//assert.NilError(t, err)
	//assert.DeepEqual(t, newE2apPdu.String(), result11.String())
	//
	//result1, err := asn1cgo.PerDecodeE2apPdu(perNew)
	//assert.NilError(t, err)
	//assert.Assert(t, result1 != nil)
	//t.Logf("E2connectionUpdateFailure E2AP PDU PER - decoded is \n%v", result1)
	//assert.DeepEqual(t, e2apPdu.String(), result1.String())
}
