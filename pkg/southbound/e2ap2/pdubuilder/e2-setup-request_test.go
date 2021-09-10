// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
package pdubuilder

import (
	"encoding/hex"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-ies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap2/asn1cgo"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap2/types"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	hexlib "github.com/onosproject/onos-lib-go/pkg/hex"
	"gotest.tools/assert"
	"testing"
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

func TestAdib(t *testing.T) {
	//bytes := "000100840d00000400030009002684133000051530000a0083d50300084080a00000057c20c04f52414e2d4532534d2d4d484f00001a312e332e362e312e342e312e35333134382e312e312e322e31303101004d484f0103600001010e00506572696f64696320616e64204f6e204368616e6765205265706f727401010001010c8050434920616e64204e52542075706461746520666f7220654e4201010101000100001a312e332e362e312e342e312e35333134382e312e312e322e31303100084080bb000001809820c04f52414e2d4532534d2d4b504d000018312e332e362e312e342e312e35333134382e312e312e322e3205004b504d204d6f6e69746f720101600001010700506572696f646963207265706f727401050001011e804f2d43552d4350204d6561737572656d656e7420436f6e7461696e657220666f72207468652035474320636f6e6e6563746564206465706c6f796d656e74010101010001000018312e332e362e312e342e312e35333134382e312e312e322e3200084080a5000003808020f04f52414e2d4532534d2d52432d50524500001a312e332e362e312e342e312e35333134382e312e322e322e313030028052432050524500036000010e00506572696f64696320616e64204f6e204368616e6765205265706f7274000100010c8050434920616e64204e52542075706461746520666f7220654e4200010001000100001a312e332e362e312e342e312e35333134382e312e322e322e31303000084081c0000004819d74184f52414e2d4532534d2d4b504d000018312e332e362e312e342e312e35333134382e312e322e322e3207004b504d20322e30204d6f6e69746f720001000043002684130001454c00000000000200000f3133383432363031343534633030310026841301454c0010000f3133383432363031343534633030320026841301454c0020000f3133383432363031343534633030330026841301454c003000010700506572696f646963205265706f7274000100010700506572696f646963205265706f72740001000742605252432e436f6e6e45737461624174742e53756d00000042805252432e436f6e6e4573746162537563632e53756d00000142a05252432e436f6e6e526545737461624174742e53756d00000243c05252432e436f6e6e526545737461624174742e7265636f6e6669674661696c00000343005252432e436f6e6e526545737461624174742e484f4661696c00000442e05252432e436f6e6e526545737461624174742e4f7468657200000541605252432e436f6e6e2e41766700000641605252432e436f6e6e2e4d6178000007000100010001000018312e332e362e312e342e312e35333134382e312e322e322e320021001900000022001340001516040000000000000000000025000000000000000180"
	//per, err := hexlib.Asn1BytesToByte(bytes)
	bytes := "00000000  00 01 00 84 19 00 00 04  00 03 00 09 00 26 84 13  |.............&..|"
	per, err := hexlib.DumpToByte(bytes)
	assert.NilError(t, err)

	e2apPdu, err := asn1cgo.PerDecodeE2apPdu(per)
	assert.NilError(t, err)
	t.Logf("Decoded message is %v\n", e2apPdu)
}

func TestAdibInverted(t *testing.T) {
	ge2nID, err := CreateGlobalE2nodeIDGnb([3]byte{0x26, 0x84, 0x13}, &asn1.BitString{
		//Value: []byte{0x00, 0x05, 0x15, 0x30}, // It is in APER bytes provided by Adib
		Value: []byte{0x00, 0x05, 0x15, 0x40}, // It is in E2AP message (protobuf notation) provided by Adib
		Len:   28,
	})
	assert.NilError(t, err)

	ranFunctionList := make(types.RanFunctions)
	ranFunctionList[5] = types.RanFunctionItem{
		Description: []byte{0x00, 0x00, 0x00},
		Revision:    1,
		OID:         "1.3.6.1.4.1.53148.1.1.2.101",
	}

	ranFunctionList[1] = types.RanFunctionItem{
		Description: []byte{0x10, 0x00, 0x00, 0x01},
		Revision:    1,
		OID:         "1.3.6.1.4.1.53148.1.1.2.2",
	}

	ranFunctionList[3] = types.RanFunctionItem{
		Description: []byte{},
		Revision:    1,
		OID:         "1.3.6.1.4.1.53148.1.2.2.100",
	}

	ranFunctionList[4] = types.RanFunctionItem{
		Description: []byte{0x00, 0x00},
		Revision:    1,
		OID:         "1.3.6.1.4.1.53148.1.2.2.2",
	}

	e2ncID := CreateE2NodeComponentIDGnbCuUp(21)
	ie33 := []*types.E2NodeComponentConfigUpdateItem{
		{E2NodeComponentType: e2ap_ies.E2NodeComponentType_E2NODE_COMPONENT_TYPE_G_NB_CU_UP,
			E2NodeComponentID:           &e2ncID,
			E2NodeComponentConfigUpdate: CreateE2NodeComponentConfigUpdateGnb([]byte("ngAp"), nil, []byte("e1Ap"), []byte("f1Ap"), nil)},
	}

	newE2apPdu, err := CreateE2SetupRequestPdu(2, ge2nID, ranFunctionList, ie33)
	assert.NilError(t, err)

	xer, err := asn1cgo.XerEncodeE2apPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("E2SetupRequest PDU in XER is \n%s", xer)

	per, err := asn1cgo.PerEncodeE2apPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("E2SetupRequest PDU in PER is \n%v", hex.Dump(per))

	result, err := asn1cgo.PerDecodeE2apPdu(per)
	assert.NilError(t, err)
	t.Logf("E2SetupRequest PDU decoded from PER is \n%v", result)
}
