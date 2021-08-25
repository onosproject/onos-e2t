// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
package pdubuilder

import (
	"encoding/hex"
	"testing"

	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap201/asn1cgo"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap201/types"
	"gotest.tools/assert"
)

func TestRicServiceUpdate(t *testing.T) {
	ranFunctionAddedList := make(map[types.RanFunctionID]types.RanFunctionItem)
	ranFunctionAddedList[100] = types.RanFunctionItem{
		Description: []byte("Type 1"),
		Revision:    1,
		OID:         "oid1",
	}

	ranFunctionAddedList[200] = types.RanFunctionItem{
		Description: []byte("Type 2"),
		Revision:    2,
		OID:         "oid2",
	}

	rfDeleted := make(types.RanFunctionRevisions)
	rfDeleted[100] = 2
	rfDeleted[200] = 2

	ranFunctionModifiedList := make(map[types.RanFunctionID]types.RanFunctionItem)
	ranFunctionModifiedList[100] = types.RanFunctionItem{
		Description: []byte("Type 3"),
		Revision:    3,
		OID:         "oid3",
	}

	ranFunctionModifiedList[200] = types.RanFunctionItem{
		Description: []byte("Type 4"),
		Revision:    4,
		OID:         "oid4",
	}

	newE2apPdu, err := CreateRicServiceUpdateE2apPdu(1)
	assert.NilError(t, err)
	assert.Assert(t, newE2apPdu != nil)
	newE2apPdu.GetInitiatingMessage().GetProcedureCode().GetRicServiceUpdate().GetInitiatingMessage().
		SetRanFunctionsAdded(ranFunctionAddedList).SetRanFunctionsModified(ranFunctionModifiedList).SetRanFunctionsDeleted(rfDeleted)

	xer, err := asn1cgo.XerEncodeE2apPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("RicServiceUpdate E2AP PDU XER\n%s", string(xer))

	result, err := asn1cgo.XerDecodeE2apPdu(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("RicServiceUpdate E2AP PDU XER - decoded is \n%v", result)
	assert.DeepEqual(t, newE2apPdu.String(), result.String())

	per, err := asn1cgo.PerEncodeE2apPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("RicServiceUpdate E2AP PDU PER\n%v", hex.Dump(per))

	result1, err := asn1cgo.PerDecodeE2apPdu(per)
	assert.NilError(t, err)
	assert.Assert(t, result1 != nil)
	t.Logf("RicServiceUpdate E2AP PDU PER - decoded is \n%v", result1)
	assert.DeepEqual(t, newE2apPdu.String(), result1.String())
}

func TestRicServiceUpdateExcludeOptionalIEs(t *testing.T) {
	ranFunctionAddedList := make(map[types.RanFunctionID]types.RanFunctionItem)
	ranFunctionAddedList[100] = types.RanFunctionItem{
		Description: []byte("Type 1"),
		Revision:    1,
		OID:         "oid1",
	}

	ranFunctionAddedList[200] = types.RanFunctionItem{
		Description: []byte("Type 2"),
		Revision:    2,
		OID:         "oid2",
	}

	newE2apPdu, err := CreateRicServiceUpdateE2apPdu(1)
	assert.NilError(t, err)
	assert.Assert(t, newE2apPdu != nil)
	newE2apPdu.GetInitiatingMessage().GetProcedureCode().GetRicServiceUpdate().GetInitiatingMessage().SetRanFunctionsAdded(ranFunctionAddedList)

	xer, err := asn1cgo.XerEncodeE2apPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("RicServiceUpdate E2AP PDU XER\n%s", string(xer))

	result, err := asn1cgo.XerDecodeE2apPdu(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("RicServiceUpdate E2AP PDU XER - decoded is \n%v", result)
	assert.DeepEqual(t, newE2apPdu.String(), result.String())

	per, err := asn1cgo.PerEncodeE2apPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("RicServiceUpdate E2AP PDU PER\n%v", hex.Dump(per))

	result1, err := asn1cgo.PerDecodeE2apPdu(per)
	assert.NilError(t, err)
	assert.Assert(t, result1 != nil)
	t.Logf("RicServiceUpdate E2AP PDU PER - decoded is \n%v", result1)
	assert.DeepEqual(t, newE2apPdu.String(), result1.String())
}
