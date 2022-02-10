// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
package pdubuilder

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/encoder"
	"testing"

	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
	"gotest.tools/assert"
)

func TestRicServiceUpdate(t *testing.T) {
	//ranFunctionAddedList1 := make(map[types1.RanFunctionID]types1.RanFunctionItem)
	//ranFunctionAddedList1[100] = types1.RanFunctionItem{
	//	Description: []byte("Type 1"),
	//	Revision:    1,
	//	OID:         "oid1",
	//}
	//
	////ranFunctionAddedList1[200] = types1.RanFunctionItem{
	////	Description: []byte("Type 2"),
	////	Revision:    2,
	////	OID:         "oid2",
	////}
	//
	//rfDeleted1 := make(types1.RanFunctionRevisions)
	//rfDeleted1[100] = 2
	////rfDeleted1[200] = 2
	//
	//ranFunctionModifiedList1 := make(map[types1.RanFunctionID]types1.RanFunctionItem)
	//ranFunctionModifiedList1[100] = types1.RanFunctionItem{
	//	Description: []byte("Type 3"),
	//	Revision:    3,
	//	OID:         "oid3",
	//}
	//
	////ranFunctionModifiedList1[200] = types1.RanFunctionItem{
	////	Description: []byte("Type 4"),
	////	Revision:    4,
	////	OID:         "oid4",
	////}
	//
	//e2apPdu, err := pdubuilder.CreateRicServiceUpdateE2apPdu(1)
	//assert.NilError(t, err)
	//assert.Assert(t, e2apPdu != nil)
	//e2apPdu.GetInitiatingMessage().GetProcedureCode().GetRicServiceUpdate().GetInitiatingMessage().
	//	SetRanFunctionsAdded(ranFunctionAddedList1).SetRanFunctionsModified(ranFunctionModifiedList1).SetRanFunctionsDeleted(rfDeleted1)
	//
	//per, err := asn1cgo.PerEncodeE2apPdu(e2apPdu)
	//assert.NilError(t, err)
	//t.Logf("RicServiceUpdate E2AP PDU PER\n%v", hex.Dump(per))

	ranFunctionAddedList := make(map[types.RanFunctionID]types.RanFunctionItem)
	ranFunctionAddedList[100] = types.RanFunctionItem{
		Description: []byte("Type 1"),
		Revision:    1,
		OID:         "oid1",
	}

	//ranFunctionAddedList[200] = types.RanFunctionItem{
	//	Description: []byte("Type 2"),
	//	Revision:    2,
	//	OID:         "oid2",
	//}

	rfDeleted := make(types.RanFunctionRevisions)
	rfDeleted[100] = 2
	//rfDeleted[200] = 2

	ranFunctionModifiedList := make(map[types.RanFunctionID]types.RanFunctionItem)
	ranFunctionModifiedList[100] = types.RanFunctionItem{
		Description: []byte("Type 3"),
		Revision:    3,
		OID:         "oid3",
	}

	//ranFunctionModifiedList[200] = types.RanFunctionItem{
	//	Description: []byte("Type 4"),
	//	Revision:    4,
	//	OID:         "oid4",
	//}

	newE2apPdu, err := CreateRicServiceUpdateE2apPdu(1)
	assert.NilError(t, err)
	assert.Assert(t, newE2apPdu != nil)
	newE2apPdu.GetInitiatingMessage().GetValue().GetRicServiceUpdate().
		SetRanFunctionsAdded(ranFunctionAddedList).SetRanFunctionsModified(ranFunctionModifiedList).SetRanFunctionsDeleted(rfDeleted)

	perNew, err := encoder.PerEncodeE2ApPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("RicServiceUpdate E2AP PDU PER with Go APER library\n%v", hex.Dump(perNew))

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
	//t.Logf("RicServiceUpdate E2AP PDU PER - decoded is \n%v", result1)
	//assert.DeepEqual(t, e2apPdu.String(), result1.String())
}

func TestRicServiceUpdateExcludeOptionalIEs(t *testing.T) {
	//ranFunctionAddedList1 := make(map[types1.RanFunctionID]types1.RanFunctionItem)
	//ranFunctionAddedList1[100] = types1.RanFunctionItem{
	//	Description: []byte("Type 1"),
	//	Revision:    1,
	//	OID:         "oid1",
	//}
	//
	////ranFunctionAddedList1[200] = types1.RanFunctionItem{
	////	Description: []byte("Type 2"),
	////	Revision:    2,
	////	OID:         "oid2",
	////}
	//
	//e2apPdu, err := pdubuilder.CreateRicServiceUpdateE2apPdu(1)
	//assert.NilError(t, err)
	//assert.Assert(t, e2apPdu != nil)
	//e2apPdu.GetInitiatingMessage().GetProcedureCode().GetRicServiceUpdate().GetInitiatingMessage().SetRanFunctionsAdded(ranFunctionAddedList1)
	//
	//per, err := asn1cgo.PerEncodeE2apPdu(e2apPdu)
	//assert.NilError(t, err)
	//t.Logf("RicServiceUpdate E2AP PDU PER\n%v", hex.Dump(per))

	ranFunctionAddedList := make(map[types.RanFunctionID]types.RanFunctionItem)
	ranFunctionAddedList[100] = types.RanFunctionItem{
		Description: []byte("Type 1"),
		Revision:    1,
		OID:         "oid1",
	}

	//ranFunctionAddedList[200] = types.RanFunctionItem{
	//	Description: []byte("Type 2"),
	//	Revision:    2,
	//	OID:         "oid2",
	//}

	newE2apPdu, err := CreateRicServiceUpdateE2apPdu(1)
	assert.NilError(t, err)
	assert.Assert(t, newE2apPdu != nil)
	newE2apPdu.GetInitiatingMessage().GetValue().GetRicServiceUpdate().SetRanFunctionsAdded(ranFunctionAddedList)

	perNew, err := encoder.PerEncodeE2ApPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("RicServiceUpdate E2AP PDU PER with Go APER library\n%v", hex.Dump(perNew))

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
	//t.Logf("RicServiceUpdate E2AP PDU PER - decoded is \n%v", result1)
	//assert.DeepEqual(t, e2apPdu.String(), result1.String())
}
