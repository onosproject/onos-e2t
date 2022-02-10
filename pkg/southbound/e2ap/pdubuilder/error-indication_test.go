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

func TestErrorIndicationE2apPdu(t *testing.T) {
	//var ranFuncID1 types1.RanFunctionID = 9
	//procCode1 := v21.ProcedureCodeIDRICsubscription
	//criticality1 := e2apcommondatatypes.Criticality_CRITICALITY_IGNORE
	//ftg1 := e2apcommondatatypes.TriggeringMessage_TRIGGERING_MESSAGE_UNSUCCESSFUL_OUTCOME
	//
	//e2apPdu := pdubuilder.CreateErrorIndicationE2apPduEmpty()
	//e2apPdu.GetInitiatingMessage().GetProcedureCode().GetErrorIndication().GetInitiatingMessage().
	//	SetTransactionID(21).SetCause(&e2ap_ies.Cause{
	//	Cause: &e2ap_ies.Cause_Misc{
	//		Misc: e2ap_ies.CauseMisc_CAUSE_MISC_UNSPECIFIED,
	//	},
	//}).SetRicRequestID(&types1.RicRequest{
	//	RequestorID: 10,
	//	InstanceID:  20,
	//}).SetRanFunctionID(&ranFuncID1).SetCriticalityDiagnostics(&procCode1, &criticality1, &ftg1,
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
	//assert.Assert(t, e2apPdu != nil)
	//
	//per, err := asn1cgo.PerEncodeE2apPdu(e2apPdu)
	//assert.NilError(t, err)
	//t.Logf("E2SetupResponse E2AP PDU PER\n%v", hex.Dump(per))

	var ranFuncID types.RanFunctionID = 9
	procCode := v2.ProcedureCodeIDRICsubscription
	criticality := e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE
	ftg := e2ap_commondatatypes.TriggeringMessage_TRIGGERING_MESSAGE_UNSUCCESSFUL_OUTCOME

	newE2apPdu := CreateErrorIndicationE2apPduEmpty()
	newE2apPdu.GetInitiatingMessage().GetValue().GetErrorIndication().
		SetTransactionID(21).SetCause(&e2apies.Cause{
		Cause: &e2apies.Cause_Misc{
			Misc: e2apies.CauseMisc_CAUSE_MISC_UNSPECIFIED,
		},
	}).SetRicRequestID(&types.RicRequest{
		RequestorID: 10,
		InstanceID:  20,
	}).SetRanFunctionID(&ranFuncID).SetCriticalityDiagnostics(&procCode, &criticality, &ftg,
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
	assert.Assert(t, newE2apPdu != nil)

	perNew, err := encoder.PerEncodeE2ApPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("E2SetupResponse E2AP PDU PER with Go APER library\n%v", hex.Dump(perNew))

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
	//assert.DeepEqual(t, result1.String(), e2apPdu.String())
}

func TestErrorIndicationE2apPduExcludeSomeOptionalIEs(t *testing.T) {
	//var ranFuncID1 types1.RanFunctionID = 9
	//procCode1 := v21.ProcedureCodeIDRICsubscription
	//criticality1 := e2apcommondatatypes.Criticality_CRITICALITY_IGNORE
	//ftg1 := e2apcommondatatypes.TriggeringMessage_TRIGGERING_MESSAGE_UNSUCCESSFUL_OUTCOME
	//var trID1 int32 = 21
	//e2apPdu, err := pdubuilder.CreateErrorIndicationE2apPdu(&trID1, nil, &ranFuncID1,
	//	&e2ap_ies.Cause{
	//		Cause: &e2ap_ies.Cause_Misc{ // Probably, could be any other reason
	//			Misc: e2ap_ies.CauseMisc_CAUSE_MISC_UNSPECIFIED,
	//		},
	//	},
	//	&procCode1, &criticality1, &ftg1,
	//	&types1.RicRequest{
	//		RequestorID: 10,
	//		InstanceID:  20,
	//	}, []*types1.CritDiag{
	//		{
	//			TypeOfError:   e2ap_ies.TypeOfError_TYPE_OF_ERROR_MISSING,
	//			IECriticality: e2apcommondatatypes.Criticality_CRITICALITY_IGNORE,
	//			IEId:          v21.ProtocolIeIDRicsubscriptionDetails,
	//		},
	//	},
	//)
	//assert.NilError(t, err)
	//assert.Assert(t, e2apPdu != nil)
	//
	//per, err := asn1cgo.PerEncodeE2apPdu(e2apPdu)
	//assert.NilError(t, err)
	//t.Logf("ErrorIndication E2AP PDU PER\n%v", hex.Dump(per))

	var ranFuncID types.RanFunctionID = 9
	procCode := v2.ProcedureCodeIDRICsubscription
	criticality := e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE
	ftg := e2ap_commondatatypes.TriggeringMessage_TRIGGERING_MESSAGE_UNSUCCESSFUL_OUTCOME
	var trID int32 = 21
	newE2apPdu, err := CreateErrorIndicationE2apPdu(&trID, nil, &ranFuncID,
		&e2apies.Cause{
			Cause: &e2apies.Cause_Misc{ // Probably, could be any other reason
				Misc: e2apies.CauseMisc_CAUSE_MISC_UNSPECIFIED,
			},
		},
		&procCode, &criticality, &ftg,
		&types.RicRequest{
			RequestorID: 10,
			InstanceID:  20,
		}, []*types.CritDiag{
			{
				TypeOfError:   e2apies.TypeOfError_TYPE_OF_ERROR_MISSING,
				IECriticality: e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE,
				IEId:          v2.ProtocolIeIDRicsubscriptionDetails,
			},
		},
	)
	assert.NilError(t, err)
	assert.Assert(t, newE2apPdu != nil)

	perNew, err := encoder.PerEncodeE2ApPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("ErrorIndication E2AP PDU PER with Go APER library\n%v", hex.Dump(perNew))

	//Comparing reference PER bytes with Go APER library produced
	//assert.DeepEqual(t, per, perNew)

	result1, err := encoder.PerDecodeE2ApPdu(perNew)
	assert.NilError(t, err)
	assert.DeepEqual(t, newE2apPdu.String(), result1.String())

	// Decoding the message from the APER bytes produced by CGo
	//result11, err := encoder.PerDecodeE2ApPdu(per)
	//assert.NilError(t, err)
	//assert.DeepEqual(t, newE2apPdu.String(), result11.String())
	//
	//result, err := asn1cgo.PerDecodeE2apPdu(per)
	//assert.NilError(t, err)
	//assert.DeepEqual(t, result.String(), e2apPdu.String())
}
