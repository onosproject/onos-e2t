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

func TestRicControlFailure(t *testing.T) {
	//ricRequestID1 := types1.RicRequest{
	//	RequestorID: 21,
	//	InstanceID:  22,
	//}
	//var ranFuncID1 types1.RanFunctionID = 9
	////var ricCallPrID types.RicCallProcessID = []byte("123")
	////var ricCtrlOut types.RicControlOutcome = []byte("456")
	//cause1 := e2ap_ies.Cause{
	//	Cause: &e2ap_ies.Cause_RicRequest{
	//		RicRequest: e2ap_ies.CauseRicrequest_CAUSE_RICREQUEST_REQUEST_ID_UNKNOWN,
	//	},
	//}
	//e2apPdu, err := pdubuilder.CreateRicControlFailureE2apPdu(ricRequestID1,
	//	ranFuncID1, &cause1)
	//assert.NilError(t, err)
	//assert.Assert(t, e2apPdu != nil)
	//e2apPdu.GetUnsuccessfulOutcome().GetProcedureCode().GetRicControl().GetUnsuccessfulOutcome().
	//	SetRicControlOutcome([]byte{0xFF, 0xFF, 0xDD, 0x4A}).SetRicCallProcessID([]byte{0xCC, 0x3D, 0x1F})
	//
	//per, err := asn1cgo.PerEncodeE2apPdu(e2apPdu)
	//assert.NilError(t, err)
	//t.Logf("RicControlFailure E2AP PDU PER\n%v", hex.Dump(per))

	ricRequestID := types.RicRequest{
		RequestorID: 21,
		InstanceID:  22,
	}
	var ranFuncID types.RanFunctionID = 9
	//var ricCallPrID types.RicCallProcessID = []byte("123")
	//var ricCtrlOut types.RicControlOutcome = []byte("456")
	cause := e2apies.Cause{
		Cause: &e2apies.Cause_RicRequest{
			RicRequest: e2apies.CauseRicrequest_CAUSE_RICREQUEST_REQUEST_ID_UNKNOWN,
		},
	}
	newE2apPdu, err := CreateRicControlFailureE2apPdu(ricRequestID,
		ranFuncID, &cause)
	assert.NilError(t, err)
	assert.Assert(t, newE2apPdu != nil)
	newE2apPdu.GetUnsuccessfulOutcome().GetValue().GetRicControl().
		SetRicControlOutcome([]byte{0xFF, 0xFF, 0xDD, 0x4A}).SetRicCallProcessID([]byte{0xCC, 0x3D, 0x1F})

	perNew, err := encoder.PerEncodeE2ApPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("RicControlFailure E2AP PDU PER with Go APER library\n%v", hex.Dump(perNew))

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
