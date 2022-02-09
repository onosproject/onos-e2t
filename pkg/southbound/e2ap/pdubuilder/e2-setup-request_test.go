// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
package pdubuilder

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/encoder"
	"testing"

	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"gotest.tools/assert"
)

func TestE2SetupRequest(t *testing.T) {
	//e2ncID11 := pdubuilder.CreateE2NodeComponentIDW1(21)
	////e2ncID12 := pdubuilder.CreateE2NodeComponentIDS1("S1-Component")
	//ranFunctionList1 := make(types1.RanFunctions)
	//ranFunctionList1[100] = types1.RanFunctionItem{
	//	Description: []byte("Type 1"),
	//	Revision:    1,
	//	OID:         "oid1",
	//}
	//
	////ranFunctionList1[200] = types1.RanFunctionItem{
	////	Description: []byte("Type 2"),
	////	Revision:    2,
	////	OID:         "oid2",
	////}
	//
	//ge2nID1, err := pdubuilder.CreateGlobalE2nodeIDGnb([3]byte{0x4F, 0x4E, 0x46}, &asn1.BitString{
	//	Value: []byte{0x00, 0x00, 0x04},
	//	Len:   22,
	//})
	//assert.NilError(t, err)
	//
	//e2apPdu, err := pdubuilder.CreateE2SetupRequestPdu(1, ge2nID1, ranFunctionList1, []*types1.E2NodeComponentConfigAdditionItem{
	//	{E2NodeComponentType: e2apies.E2NodeComponentInterfaceType_E2NODE_COMPONENT_INTERFACE_TYPE_W1,
	//		E2NodeComponentID: e2ncID11,
	//		E2NodeComponentConfiguration: e2apies.E2NodeComponentConfiguration{
	//			E2NodeComponentResponsePart: []byte{0x01, 0x02, 0x03},
	//			E2NodeComponentRequestPart:  []byte{0x04, 0x05, 0x06},
	//		}},
	//	//{E2NodeComponentType: e2apies.E2NodeComponentInterfaceType_E2NODE_COMPONENT_INTERFACE_TYPE_S1,
	//	//	E2NodeComponentID: e2ncID12,
	//	//	E2NodeComponentConfiguration: e2apies.E2NodeComponentConfiguration{
	//	//		E2NodeComponentResponsePart: []byte{0x07, 0x08, 0x09},
	//	//		E2NodeComponentRequestPart:  []byte{0x0A, 0x0B, 0x0C},
	//	//	}},
	//})
	//assert.NilError(t, err)
	//assert.Assert(t, e2apPdu != nil)
	//
	//per, err := asn1cgo.PerEncodeE2apPdu(e2apPdu)
	//assert.NilError(t, err)
	//t.Logf("E2connectionUpdate E2AP PDU PER\n%v", hex.Dump(per))

	e2ncID1 := CreateE2NodeComponentIDW1(21)
	//e2ncID2 := CreateE2NodeComponentIDS1("S1-Component")
	ranFunctionList := make(types.RanFunctions)
	ranFunctionList[100] = types.RanFunctionItem{
		Description: []byte("Type 1"),
		Revision:    1,
		OID:         "oid1",
	}

	//ranFunctionList[200] = types.RanFunctionItem{
	//	Description: []byte("Type 2"),
	//	Revision:    2,
	//	OID:         "oid2",
	//}

	ge2nID, err := CreateGlobalE2nodeIDGnb([3]byte{0x4F, 0x4E, 0x46}, &asn1.BitString{
		Value: []byte{0x00, 0x00, 0x04},
		Len:   22,
	})
	assert.NilError(t, err)

	newE2apPdu, err := CreateE2SetupRequestPdu(1, ge2nID, ranFunctionList, []*types.E2NodeComponentConfigAdditionItem{
		{E2NodeComponentType: e2ap_ies.E2NodeComponentInterfaceType_E2NODE_COMPONENT_INTERFACE_TYPE_W1,
			E2NodeComponentID: e2ncID1,
			E2NodeComponentConfiguration: e2ap_ies.E2NodeComponentConfiguration{
				E2NodeComponentResponsePart: []byte{0x01, 0x02, 0x03},
				E2NodeComponentRequestPart:  []byte{0x04, 0x05, 0x06},
			}},
		//{E2NodeComponentType: e2ap_ies.E2NodeComponentInterfaceType_E2NODE_COMPONENT_INTERFACE_TYPE_S1,
		//	E2NodeComponentID: e2ncID2,
		//	E2NodeComponentConfiguration: e2ap_ies.E2NodeComponentConfiguration{
		//		E2NodeComponentResponsePart: []byte{0x07, 0x08, 0x09},
		//		E2NodeComponentRequestPart:  []byte{0x0A, 0x0B, 0x0C},
		//	}},
	})
	assert.NilError(t, err)
	assert.Assert(t, newE2apPdu != nil)

	perNew, err := encoder.PerEncodeE2ApPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("E2connectionUpdate E2AP PDU PER with Go APER library\n%v", hex.Dump(perNew))

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

func TestE2SetupRequestCuDuIDs(t *testing.T) {
	//enbID1, err := pdubuilder.CreateEnbIDHome(&asn1.BitString{
	//	Value: []byte{0x00, 0xA7, 0xDD, 0xF0},
	//	Len:   28,
	//})
	//assert.NilError(t, err)
	//gEnbID1, err := pdubuilder.CreateGlobalEnbID([]byte{0xAA, 0xBB, 0xCC}, enbID1)
	//assert.NilError(t, err)
	//
	//gEnGnbID1, err := pdubuilder.CreateGlobalEnGnbID([]byte{0xFF, 0xCD, 0xBF}, &asn1.BitString{
	//	Value: []byte{0xFA, 0x2C, 0xD4, 0xF8},
	//	Len:   29,
	//})
	//assert.NilError(t, err)
	//
	//e2ncID11 := pdubuilder.CreateE2NodeComponentIDX2(gEnbID1, gEnGnbID1)
	////e2ncID12 := pdubuilder.CreateE2NodeComponentIDNg("NG-Component")
	//
	//ranFunctionList1 := make(types1.RanFunctions)
	//ranFunctionList1[100] = types1.RanFunctionItem{
	//	Description: []byte("Type 1"),
	//	Revision:    1,
	//	OID:         "oid1",
	//}
	//
	////ranFunctionList1[200] = types1.RanFunctionItem{
	////	Description: []byte("Type 2"),
	////	Revision:    2,
	////	OID:         "oid2",
	////}
	//
	//ge2nID1, err := pdubuilder.CreateGlobalE2nodeIDGnb([3]byte{0x4F, 0x4E, 0x46}, &asn1.BitString{
	//	Value: []byte{0x00, 0x00, 0x04},
	//	Len:   22,
	//})
	//assert.NilError(t, err)
	//ge2nID1.GetGNb().SetGnbCuUpID(2).SetGnbDuID(13)
	//
	//e2apPdu, err := pdubuilder.CreateE2SetupRequestPdu(1, ge2nID1, ranFunctionList1, []*types1.E2NodeComponentConfigAdditionItem{
	//	{E2NodeComponentType: e2apies.E2NodeComponentInterfaceType_E2NODE_COMPONENT_INTERFACE_TYPE_X2,
	//		E2NodeComponentID: e2ncID11,
	//		E2NodeComponentConfiguration: e2apies.E2NodeComponentConfiguration{
	//			E2NodeComponentResponsePart: []byte{0x01, 0x02, 0x03},
	//			E2NodeComponentRequestPart:  []byte{0x04, 0x05, 0x06},
	//		}},
	//	//{E2NodeComponentType: e2apies.E2NodeComponentInterfaceType_E2NODE_COMPONENT_INTERFACE_TYPE_NG,
	//	//	E2NodeComponentID: e2ncID12,
	//	//	E2NodeComponentConfiguration: e2apies.E2NodeComponentConfiguration{
	//	//		E2NodeComponentResponsePart: []byte{0x07, 0x08, 0x09},
	//	//		E2NodeComponentRequestPart:  []byte{0x0A, 0x0B, 0x0C},
	//	//	}},
	//})
	//assert.NilError(t, err)
	//assert.Assert(t, e2apPdu != nil)
	//
	//per, err := asn1cgo.PerEncodeE2apPdu(e2apPdu)
	//assert.NilError(t, err)
	//t.Logf("E2SetupRequest E2AP PDU PER\n%v", hex.Dump(per))

	enbID, err := CreateEnbIDHome(&asn1.BitString{
		Value: []byte{0x00, 0xA7, 0xDD, 0xF0},
		Len:   28,
	})
	assert.NilError(t, err)
	gEnbID, err := CreateGlobalEnbID([]byte{0xAA, 0xBB, 0xCC}, enbID)
	assert.NilError(t, err)

	gEnGnbID, err := CreateGlobalEnGnbID([]byte{0xFF, 0xCD, 0xBF}, &asn1.BitString{
		Value: []byte{0xFA, 0x2C, 0xD4, 0xF8},
		Len:   29,
	})
	assert.NilError(t, err)

	e2ncID1 := CreateE2NodeComponentIDX2(gEnbID, gEnGnbID)
	//e2ncID2 := CreateE2NodeComponentIDNg("NG-Component")

	ranFunctionList := make(types.RanFunctions)
	ranFunctionList[100] = types.RanFunctionItem{
		Description: []byte("Type 1"),
		Revision:    1,
		OID:         "oid1",
	}

	//ranFunctionList[200] = types.RanFunctionItem{
	//	Description: []byte("Type 2"),
	//	Revision:    2,
	//	OID:         "oid2",
	//}

	ge2nID, err := CreateGlobalE2nodeIDGnb([3]byte{0x4F, 0x4E, 0x46}, &asn1.BitString{
		Value: []byte{0x00, 0x00, 0x04},
		Len:   22,
	})
	assert.NilError(t, err)
	ge2nID.GetGNb().SetGnbCuUpID(2).SetGnbDuID(13)

	newE2apPdu, err := CreateE2SetupRequestPdu(1, ge2nID, ranFunctionList, []*types.E2NodeComponentConfigAdditionItem{
		{E2NodeComponentType: e2ap_ies.E2NodeComponentInterfaceType_E2NODE_COMPONENT_INTERFACE_TYPE_X2,
			E2NodeComponentID: e2ncID1,
			E2NodeComponentConfiguration: e2ap_ies.E2NodeComponentConfiguration{
				E2NodeComponentResponsePart: []byte{0x01, 0x02, 0x03},
				E2NodeComponentRequestPart:  []byte{0x04, 0x05, 0x06},
			}},
		//{E2NodeComponentType: e2ap_ies.E2NodeComponentInterfaceType_E2NODE_COMPONENT_INTERFACE_TYPE_NG,
		//	E2NodeComponentID: e2ncID2,
		//	E2NodeComponentConfiguration: e2ap_ies.E2NodeComponentConfiguration{
		//		E2NodeComponentResponsePart: []byte{0x07, 0x08, 0x09},
		//		E2NodeComponentRequestPart:  []byte{0x0A, 0x0B, 0x0C},
		//	}},
	})
	assert.NilError(t, err)
	assert.Assert(t, newE2apPdu != nil)

	perNew, err := encoder.PerEncodeE2ApPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("E2SetupRequest E2AP PDU PER with Go APER library\n%v", hex.Dump(perNew))

	//Comparing reference PER bytes with Go APER library produced
	//assert.DeepEqual(t, per, perNew)

	//e2apPdu, err := encoder.PerDecodeE2ApPdu(perNew)
	//assert.NilError(t, err)
	//assert.DeepEqual(t, newE2apPdu.String(), e2apPdu.String())

	//result, err := asn1cgo.PerDecodeE2apPdu(perNew)
	//assert.NilError(t, err)
	//assert.DeepEqual(t, result.String(), e2apPdu.String())
}
