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

func TestRicIndication(t *testing.T) {
	//ricRequestID1 := types1.RicRequest{
	//	RequestorID: 21,
	//	InstanceID:  22,
	//}
	//var ranFuncID1 types1.RanFunctionID = 9
	//var ricAction1 = e2ap_ies.RicactionType_RICACTION_TYPE_POLICY
	//var ricIndicationType1 = e2ap_ies.RicindicationType_RICINDICATION_TYPE_INSERT
	//var ricSn1 types1.RicIndicationSn = 1
	//var ricIndHd1 types1.RicIndicationHeader = []byte("123")
	//var ricIndMsg1 types1.RicIndicationMessage = []byte("456")
	//var ricCallPrID1 types1.RicCallProcessID = []byte("789")
	//e2apPdu, err := pdubuilder.RicIndicationE2apPdu(ricRequestID1,
	//	ranFuncID1, ricAction1, ricIndicationType1, ricIndHd1, ricIndMsg1)
	//assert.NilError(t, err)
	//assert.Assert(t, e2apPdu != nil)
	//e2apPdu.GetInitiatingMessage().GetProcedureCode().GetRicIndication().GetInitiatingMessage().
	//	SetRicCallProcessID(ricCallPrID1).SetRicIndicationSN(ricSn1)
	//
	//per, err := asn1cgo.PerEncodeE2apPdu(e2apPdu)
	//assert.NilError(t, err)
	//t.Logf("RicIndication E2AP PDU PER\n%v", hex.Dump(per))

	ricRequestID := types.RicRequest{
		RequestorID: 21,
		InstanceID:  22,
	}
	var ranFuncID types.RanFunctionID = 9
	var ricAction int32 = 2
	var ricIndicationType = e2apies.RicindicationType_RICINDICATION_TYPE_INSERT
	var ricSn types.RicIndicationSn = 1
	var ricIndHd types.RicIndicationHeader = []byte("123")
	var ricIndMsg types.RicIndicationMessage = []byte("456")
	var ricCallPrID types.RicCallProcessID = []byte("789")
	newE2apPdu, err := RicIndicationE2apPdu(ricRequestID,
		ranFuncID, ricAction, ricIndicationType, ricIndHd, ricIndMsg)
	assert.NilError(t, err)
	assert.Assert(t, newE2apPdu != nil)
	newE2apPdu.GetInitiatingMessage().GetValue().GetRicIndication().
		SetRicCallProcessID(ricCallPrID).SetRicIndicationSN(ricSn)

	perNew, err := encoder.PerEncodeE2ApPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("RicIndication E2AP PDU PER with Go APER library\n%v", hex.Dump(perNew))

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
