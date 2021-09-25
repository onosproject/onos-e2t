// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
package pdubuilder

import (
	"encoding/hex"
	"testing"

	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/asn1cgo"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"gotest.tools/assert"
)

func TestE2SetupRequest(t *testing.T) {
	e2ncID1 := CreateE2NodeComponentIDGnbCuUp(21)
	e2ncID2 := CreateE2NodeComponentIDGnbDu(13)
	ranFunctionList := make(types.RanFunctions)
	ranFunctionList[100] = types.RanFunctionItem{
		Description: []byte("Type 1"),
		Revision:    1,
		OID:         "oid1",
	}

	ranFunctionList[200] = types.RanFunctionItem{
		Description: []byte("Type 2"),
		Revision:    2,
		OID:         "oid2",
	}

	ge2nID, err := CreateGlobalE2nodeIDGnb([3]byte{0x4F, 0x4E, 0x46}, &asn1.BitString{
		Value: []byte{0x00, 0x00, 0x04},
		Len:   22,
	})
	assert.NilError(t, err)

	newE2apPdu, err := CreateE2SetupRequestPdu(1, ge2nID, ranFunctionList, []*types.E2NodeComponentConfigUpdateItem{
		{E2NodeComponentType: e2ap_ies.E2NodeComponentType_E2NODE_COMPONENT_TYPE_G_NB,
			E2NodeComponentID:           &e2ncID1,
			E2NodeComponentConfigUpdate: CreateE2NodeComponentConfigUpdateGnb([]byte("ngAp"), nil, []byte("e1Ap"), []byte("f1Ap"), nil)},
		{E2NodeComponentType: e2ap_ies.E2NodeComponentType_E2NODE_COMPONENT_TYPE_E_NB,
			E2NodeComponentID:           &e2ncID2,
			E2NodeComponentConfigUpdate: CreateE2NodeComponentConfigUpdateEnb(nil, nil, nil, []byte("s1"), nil)},
	})
	assert.NilError(t, err)
	assert.Assert(t, newE2apPdu != nil)

	xer, err := asn1cgo.XerEncodeE2apPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("E2SetupRequest XER\n%s", string(xer))

	e2apPdu, err := asn1cgo.XerDecodeE2apPdu(xer)
	assert.NilError(t, err)
	assert.DeepEqual(t, newE2apPdu.String(), e2apPdu.String())

	per, err := asn1cgo.PerEncodeE2apPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("E2SetupRequest E2AP PDU PER\n%v", hex.Dump(per))

	e2apPdu, err = asn1cgo.PerDecodeE2apPdu(per)
	assert.NilError(t, err)
	assert.DeepEqual(t, newE2apPdu.String(), e2apPdu.String())
}

func TestE2SetupRequestCuDuIDs(t *testing.T) {
	e2ncID1 := CreateE2NodeComponentIDGnbCuUp(21)
	e2ncID2 := CreateE2NodeComponentIDGnbDu(13)
	ranFunctionList := make(types.RanFunctions)
	ranFunctionList[100] = types.RanFunctionItem{
		Description: []byte("Type 1"),
		Revision:    1,
		OID:         "oid1",
	}

	ranFunctionList[200] = types.RanFunctionItem{
		Description: []byte("Type 2"),
		Revision:    2,
		OID:         "oid2",
	}

	ge2nID, err := CreateGlobalE2nodeIDGnb([3]byte{0x4F, 0x4E, 0x46}, &asn1.BitString{
		Value: []byte{0x00, 0x00, 0x04},
		Len:   22,
	})
	assert.NilError(t, err)

	newE2apPdu, err := CreateE2SetupRequestPdu(1, ge2nID, ranFunctionList, []*types.E2NodeComponentConfigUpdateItem{
		{E2NodeComponentType: e2ap_ies.E2NodeComponentType_E2NODE_COMPONENT_TYPE_G_NB,
			E2NodeComponentID:           &e2ncID1,
			E2NodeComponentConfigUpdate: CreateE2NodeComponentConfigUpdateGnb([]byte("ngAp"), nil, []byte("e1Ap"), []byte("f1Ap"), nil)},
		{E2NodeComponentType: e2ap_ies.E2NodeComponentType_E2NODE_COMPONENT_TYPE_E_NB,
			E2NodeComponentID:           &e2ncID2,
			E2NodeComponentConfigUpdate: CreateE2NodeComponentConfigUpdateEnb(nil, nil, nil, []byte("s1"), nil)},
	})
	assert.NilError(t, err)
	assert.Assert(t, newE2apPdu != nil)
	newE2apPdu.GetInitiatingMessage().GetProcedureCode().GetE2Setup().GetInitiatingMessage().GetProtocolIes().GetE2ApProtocolIes3().GetValue().GetGNb().GNbCuUpId = &e2ap_ies.GnbCuUpId{
		Value: 15,
	}
	newE2apPdu.GetInitiatingMessage().GetProcedureCode().GetE2Setup().GetInitiatingMessage().GetProtocolIes().GetE2ApProtocolIes3().GetValue().GetGNb().GNbDuId = &e2ap_ies.GnbDuId{
		Value: 21,
	}

	xer, err := asn1cgo.XerEncodeE2apPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("E2SetupRequest XER\n%s", string(xer))

	e2apPdu, err := asn1cgo.XerDecodeE2apPdu(xer)
	assert.NilError(t, err)
	assert.DeepEqual(t, newE2apPdu.String(), e2apPdu.String())

	per, err := asn1cgo.PerEncodeE2apPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("E2SetupRequest E2AP PDU PER\n%v", hex.Dump(per))

	e2apPdu, err = asn1cgo.PerDecodeE2apPdu(per)
	assert.NilError(t, err)
	assert.DeepEqual(t, newE2apPdu.String(), e2apPdu.String())
}
