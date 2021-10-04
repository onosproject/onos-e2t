// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package asn1cgo

import (
	"encoding/hex"
	"testing"

	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/pdubuilder"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"gotest.tools/assert"
)

func Test_E2setupRequest(t *testing.T) {

	//e2ncID1 := CreateE2NodeComponentIDGnbCuUp(21)
	//e2ncID2 := CreateE2NodeComponentIDGnbDu(13)
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

	globale2nID, err := pdubuilder.CreateGlobalE2nodeIDGnb([3]byte{0x4F, 0x4E, 0x46}, &asn1.BitString{
		Value: []byte{0xd4, 0xcb, 0x8c},
		Len:   22,
	})
	assert.NilError(t, err)

	e2srPdu, err := pdubuilder.CreateE2SetupRequestPdu(1, globale2nID, ranFunctionList, []*types.E2NodeComponentConfigAdditionItem{
		{E2NodeComponentType: e2apies.E2NodeComponentInterfaceType_E2NODE_COMPONENT_INTERFACE_TYPE_W1,
			E2NodeComponentID: pdubuilder.CreateE2NodeComponentIDW1(3),
			E2NodeComponentConfiguration: e2apies.E2NodeComponentConfiguration{
				E2NodeComponentResponsePart: []byte{0x00, 0x01, 0x02},
				E2NodeComponentRequestPart:  []byte{0xAB, 0xCD, 0xEF},
			}},
		{E2NodeComponentType: e2apies.E2NodeComponentInterfaceType_E2NODE_COMPONENT_INTERFACE_TYPE_E1,
			E2NodeComponentID: pdubuilder.CreateE2NodeComponentIDE1(2),
			E2NodeComponentConfiguration: e2apies.E2NodeComponentConfiguration{
				E2NodeComponentResponsePart: []byte{0x00, 0x01, 0x02},
				E2NodeComponentRequestPart:  []byte{0xAB, 0xCD, 0xEF},
			}},
	})
	assert.NilError(t, err)
	assert.Assert(t, e2srPdu != nil)

	e2sr := e2srPdu.GetInitiatingMessage().GetProcedureCode().GetE2Setup().GetInitiatingMessage()

	xer, err := xerEncodeE2SetupRequest(e2sr)
	assert.NilError(t, err)
	t.Logf("XER E2SetupRequest: \n%s", string(xer))

	e2srReversed, err := xerDecodeE2SetupRequest(xer)
	assert.NilError(t, err)
	assert.Assert(t, e2srReversed != nil)
	t.Logf("E2SetupRequest decoded from XER is \n%v", e2srReversed)
	assert.Equal(t, e2sr.String(), e2srReversed.String())

	per, err := perEncodeE2SetupRequest(e2sr)
	assert.NilError(t, err)
	t.Logf("PER E2SetupRequest: \n%v", hex.Dump(per))

	e2srReversedPer, err := perDecodeE2SetupRequest(per)
	assert.NilError(t, err)
	assert.Assert(t, e2srReversedPer != nil)
	t.Logf("E2SetupRequest decoded from PER is \n%v", e2srReversedPer)
	assert.Equal(t, e2sr.String(), e2srReversedPer.String())
}
