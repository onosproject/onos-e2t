// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

import (
	"encoding/hex"
	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-ies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap2/pdubuilder"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap2/types"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"gotest.tools/assert"
	"testing"
)

func Test_E2setupRequest(t *testing.T) {

	//e2ncID1 := CreateE2NodeComponentIDGnbCuUp(21)
	//e2ncID2 := CreateE2NodeComponentIDGnbDu(13)
	e2nccu1 := pdubuilder.CreateE2NodeComponentConfigUpdateGnb([]byte("ngAp"), nil, []byte("e1Ap"), []byte("f1Ap"), nil)
	e2nccu2 := pdubuilder.CreateE2NodeComponentConfigUpdateEnb(nil, nil, nil, []byte("s1"), nil)
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

	e2srPdu, err := pdubuilder.CreateE2SetupRequestPdu(1, globale2nID, ranFunctionList, []*types.E2NodeComponentConfigUpdateItem{
		{E2NodeComponentType: e2apies.E2NodeComponentType_E2NODE_COMPONENT_TYPE_G_NB,
			//E2NodeComponentID:           &e2ncID1,
			E2NodeComponentConfigUpdate: e2nccu1},
		{E2NodeComponentType: e2apies.E2NodeComponentType_E2NODE_COMPONENT_TYPE_E_NB,
			//E2NodeComponentID:           &e2ncID2,
			E2NodeComponentConfigUpdate: e2nccu2},
	})
	assert.NilError(t, err)
	assert.Assert(t, e2srPdu != nil)

	e2sr := e2srPdu.GetInitiatingMessage().GetProcedureCode().GetE2Setup().GetInitiatingMessage()

	// Convert this Go struct in to a C struct
	e2srC, err := newE2SetupRequest(e2sr)
	assert.NilError(t, err)
	assert.Assert(t, e2srC != nil)

	// Now reverse it and decode the other way round to a Go struct
	e2srFedback, err := decodeE2setupRequest(e2srC)
	assert.NilError(t, err)
	//assert.Assert(t, e2srFedback != nil) //Commented due to the Linters (v1.34.1) error - possible nil pointer dereference (https://staticcheck.io/docs/checks#SA5011) on line 54
	ge2nID := e2srFedback.ProtocolIes.E2ApProtocolIes3.Value.GlobalE2NodeId.(*e2apies.GlobalE2NodeId_GNb)
	assert.Equal(t, "ONF", string(ge2nID.GNb.GlobalGNbId.PlmnId.Value))
	gnbID := ge2nID.GNb.GlobalGNbId.GnbId.GnbIdChoice.(*e2apies.GnbIdChoice_GnbId)
	assert.DeepEqual(t, []byte{0xd4, 0xcb, 0x8c}, gnbID.GnbId.Value)
	assert.Equal(t, uint32(22), gnbID.GnbId.Len)

	xer, err := xerEncodeE2SetupRequest(e2srFedback)
	assert.NilError(t, err)
	t.Logf("XER E2SetupRequest: \n%s", string(xer))

	per, err := perEncodeE2SetupRequest(e2srFedback)
	assert.NilError(t, err)
	t.Logf("PER E2SetupRequest: \n%v", hex.Dump(per))
}
