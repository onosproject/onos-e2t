// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
package pdubuilder

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/asn1cgo"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/types"
	"gotest.tools/assert"
	"testing"
)

func TestE2SetupRequest(t *testing.T) {
	ranFunctionList := make(types.RanFunctions)
	ranFunctionList[100] = types.RanFunctionItem{
		Description: []byte("Type 1"),
		Revision:    1,
		OID:         []byte("oid1"),
	}

	ranFunctionList[200] = types.RanFunctionItem{
		Description: []byte("Type 2"),
		Revision:    2,
		OID:         []byte("oid2"),
	}

	gnbID, err := CreateGnbIDchoice(1, 22)
	assert.NilError(t, err)

	newE2apPdu, err := CreateE2SetupRequestPdu([3]byte{0x4F, 0x4E, 0x46}, gnbID, ranFunctionList)
	assert.NilError(t, err)
	assert.Assert(t, newE2apPdu != nil)

	xer, err := asn1cgo.XerEncodeE2apPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("E2SetupRequest XER\n%s", string(xer))

	e2apPdu, err := asn1cgo.XerDecodeE2apPdu(xer)
	assert.NilError(t, err)
	assert.DeepEqual(t, newE2apPdu, e2apPdu)

	per, err := asn1cgo.PerEncodeE2apPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("E2SetupRequest E2AP PDU PER\n%v", hex.Dump(per))

	e2apPdu, err = asn1cgo.PerDecodeE2apPdu(per)
	assert.NilError(t, err)
	assert.DeepEqual(t, newE2apPdu, e2apPdu)
}

func TestE2SetupRequestExcludeOptionalIE(t *testing.T) {
	gnbID, err := CreateGnbIDchoice(1, 22)
	assert.NilError(t, err)

	newE2apPdu, err := CreateE2SetupRequestPdu([3]byte{0x4F, 0x4E, 0x46}, gnbID, nil)
	assert.NilError(t, err)
	assert.Assert(t, newE2apPdu != nil)

	xer, err := asn1cgo.XerEncodeE2apPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("E2SetupRequest XER\n%s", string(xer))

	e2apPdu, err := asn1cgo.XerDecodeE2apPdu(xer)
	assert.NilError(t, err)
	assert.DeepEqual(t, newE2apPdu, e2apPdu)

	per, err := asn1cgo.PerEncodeE2apPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("E2SetupRequest E2AP PDU PER\n%v", hex.Dump(per))

	e2apPdu, err = asn1cgo.PerDecodeE2apPdu(per)
	assert.NilError(t, err)
	assert.DeepEqual(t, newE2apPdu, e2apPdu)
}
