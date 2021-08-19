// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
package pdubuilder

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap2/asn1cgo"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap2/types"
	"gotest.tools/assert"
	"testing"
)

func TestRicServiceUpdateAcknowledge(t *testing.T) {
	rfAccepted := make(types.RanFunctionRevisions)
	rfAccepted[100] = 2
	rfAccepted[200] = 2

	ranFunctionAddedList := make(types.RanFunctions)
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

	ranFunctionModifiedList := make(types.RanFunctions)
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

	rfDeleted := make(types.RanFunctionRevisions)
	rfDeleted[100] = 2
	rfDeleted[200] = 2

	newE2apPdu, err := CreateRicServiceUpdateAcknowledgeE2apPdu(1, rfAccepted)
	assert.NilError(t, err)
	assert.Assert(t, newE2apPdu != nil)
	newE2apPdu.GetInitiatingMessage().GetProcedureCode().GetRicServiceUpdate().GetInitiatingMessage().
		SetRanFunctionsAdded(ranFunctionAddedList).SetRanFunctionsModified(ranFunctionModifiedList).SetRanFunctionsDeleted(rfDeleted)

	xer, err := asn1cgo.XerEncodeE2apPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("RicServiceUpdateAcknowledge E2AP PDU XER\n%s", string(xer))

	result, err := asn1cgo.XerDecodeE2apPdu(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("RicServiceUpdateAcknowledge E2AP PDU XER - decoded is \n%v", result)
	assert.DeepEqual(t, newE2apPdu, result)

	per, err := asn1cgo.PerEncodeE2apPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("RicServiceUpdateAcknowledge E2AP PDU PER\n%v", hex.Dump(per))

	result1, err := asn1cgo.PerDecodeE2apPdu(per)
	assert.NilError(t, err)
	assert.Assert(t, result1 != nil)
	t.Logf("RicServiceUpdateAcknowledge E2AP PDU PER - decoded is \n%v", result1)
	assert.DeepEqual(t, newE2apPdu, result1)
}

func TestRicServiceUpdateAcknowledgeExcludeOptionalIE(t *testing.T) {
	rfAccepted := make(types.RanFunctionRevisions)
	rfAccepted[100] = 2
	rfAccepted[200] = 2

	newE2apPdu, err := CreateRicServiceUpdateAcknowledgeE2apPdu(3, rfAccepted)
	assert.NilError(t, err)
	assert.Assert(t, newE2apPdu != nil)

	xer, err := asn1cgo.XerEncodeE2apPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("RicServiceUpdateAcknowledge E2AP PDU XER\n%s", string(xer))

	result, err := asn1cgo.XerDecodeE2apPdu(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("RicServiceUpdateAcknowledge E2AP PDU XER - decoded is \n%v", result)
	assert.DeepEqual(t, newE2apPdu, result)

	per, err := asn1cgo.PerEncodeE2apPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("RicServiceUpdateAcknowledge E2AP PDU PER\n%v", hex.Dump(per))

	result1, err := asn1cgo.PerDecodeE2apPdu(per)
	assert.NilError(t, err)
	assert.Assert(t, result1 != nil)
	t.Logf("RicServiceUpdateAcknowledge E2AP PDU PER - decoded is \n%v", result1)
	assert.DeepEqual(t, newE2apPdu, result1)
}
