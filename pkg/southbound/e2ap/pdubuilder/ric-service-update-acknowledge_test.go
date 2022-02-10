// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
package pdubuilder

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/encoder"
	"testing"

	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
	"gotest.tools/assert"
)

func TestRicServiceUpdateAcknowledge(t *testing.T) {
	//rfAccepted1 := make(types1.RanFunctionRevisions)
	//rfAccepted1[100] = 2
	////rfAccepted1[200] = 2
	//
	//rfRejected1 := make(types1.RanFunctionCauses)
	//rfRejected1[101] = &e2ap_ies.Cause{
	//	Cause: &e2ap_ies.Cause_Misc{
	//		Misc: e2ap_ies.CauseMisc_CAUSE_MISC_HARDWARE_FAILURE,
	//	},
	//}
	////rfRejected1[102] = &e2ap_ies.Cause{
	////	Cause: &e2ap_ies.Cause_Protocol{
	////		Protocol: e2ap_ies.CauseProtocol_CAUSE_PROTOCOL_SEMANTIC_ERROR,
	////	},
	////}
	//
	//e2apPdu, err := pdubuilder.CreateRicServiceUpdateAcknowledgeE2apPdu(1, rfAccepted1)
	//assert.NilError(t, err)
	//assert.Assert(t, e2apPdu != nil)
	//e2apPdu.GetSuccessfulOutcome().GetProcedureCode().GetRicServiceUpdate().GetSuccessfulOutcome().
	//	SetRanFunctionsRejected(rfRejected1)
	//
	//per, err := asn1cgo.PerEncodeE2apPdu(e2apPdu)
	//assert.NilError(t, err)
	//t.Logf("RicServiceUpdateAcknowledge E2AP PDU PER\n%v", hex.Dump(per))

	rfAccepted := make(types.RanFunctionRevisions)
	rfAccepted[100] = 2
	//rfAccepted[200] = 2

	rfRejected := make(types.RanFunctionCauses)
	rfRejected[101] = &e2apies.Cause{
		Cause: &e2apies.Cause_Misc{
			Misc: e2apies.CauseMisc_CAUSE_MISC_HARDWARE_FAILURE,
		},
	}
	//rfRejected[102] = &e2apies.Cause{
	//	Cause: &e2apies.Cause_Protocol{
	//		Protocol: e2apies.CauseProtocol_CAUSE_PROTOCOL_SEMANTIC_ERROR,
	//	},
	//}

	newE2apPdu, err := CreateRicServiceUpdateAcknowledgeE2apPdu(1, rfAccepted)
	assert.NilError(t, err)
	assert.Assert(t, newE2apPdu != nil)
	newE2apPdu.GetSuccessfulOutcome().GetValue().GetRicServiceUpdate().
		SetRanFunctionsRejected(rfRejected)

	perNew, err := encoder.PerEncodeE2ApPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("RicServiceUpdateAcknowledge E2AP PDU PER with Go APER library\n%v", hex.Dump(perNew))

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
	//t.Logf("RicServiceUpdateAcknowledge E2AP PDU PER - decoded is \n%v", result1)
	//assert.DeepEqual(t, e2apPdu.String(), result1.String())
}

func TestRicServiceUpdateAcknowledgeExcludeOptionalIE(t *testing.T) {
	//rfAccepted1 := make(types1.RanFunctionRevisions)
	//rfAccepted1[100] = 2
	////rfAccepted1[200] = 2
	//
	//e2apPdu, err := pdubuilder.CreateRicServiceUpdateAcknowledgeE2apPdu(3, rfAccepted1)
	//assert.NilError(t, err)
	//assert.Assert(t, e2apPdu != nil)
	//
	//per, err := asn1cgo.PerEncodeE2apPdu(e2apPdu)
	//assert.NilError(t, err)
	//t.Logf("RicServiceUpdateAcknowledge E2AP PDU PER\n%v", hex.Dump(per))

	rfAccepted := make(types.RanFunctionRevisions)
	rfAccepted[100] = 2
	//rfAccepted[200] = 2

	newE2apPdu, err := CreateRicServiceUpdateAcknowledgeE2apPdu(3, rfAccepted)
	assert.NilError(t, err)
	assert.Assert(t, newE2apPdu != nil)

	perNew, err := encoder.PerEncodeE2ApPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("RicServiceUpdateAcknowledge E2AP PDU PER with Go APER library\n%v", hex.Dump(perNew))

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
	//t.Logf("RicServiceUpdateAcknowledge E2AP PDU PER - decoded is \n%v", result1)
	//assert.DeepEqual(t, e2apPdu.String(), result1.String())
}
