// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package asn1cgo

import (
	"encoding/hex"
	"testing"

	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/pdubuilder"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"gotest.tools/assert"
)

func createRanFunctionItem() (*e2appducontents.RanfunctionItem, error) {

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

	ge2nID, err := pdubuilder.CreateGlobalE2nodeIDGnb([3]byte{0x4F, 0x4E, 0x46}, &asn1.BitString{
		Value: []byte{0x00, 0x00, 0x04},
		Len:   22,
	})
	if err != nil {
		return nil, err
	}

	e2apSetupRequest, err := pdubuilder.CreateE2SetupRequestPdu(1, ge2nID, ranFunctionList, []*types.E2NodeComponentConfigAdditionItem{
		{E2NodeComponentType: e2ap_ies.E2NodeComponentInterfaceType_E2NODE_COMPONENT_INTERFACE_TYPE_W1,
			E2NodeComponentID: pdubuilder.CreateE2NodeComponentIDW1(1),
			E2NodeComponentConfiguration: e2ap_ies.E2NodeComponentConfiguration{
				E2NodeComponentResponsePart: []byte{0x00, 0x01, 0x02},
				E2NodeComponentRequestPart:  []byte{0xAB, 0xCD, 0xEF},
			}},
		{E2NodeComponentType: e2ap_ies.E2NodeComponentInterfaceType_E2NODE_COMPONENT_INTERFACE_TYPE_E1,
			E2NodeComponentID: pdubuilder.CreateE2NodeComponentIDE1(2),
			E2NodeComponentConfiguration: e2ap_ies.E2NodeComponentConfiguration{
				E2NodeComponentResponsePart: []byte{0x00, 0x01, 0x02},
				E2NodeComponentRequestPart:  []byte{0xAB, 0xCD, 0xEF},
			}},
	})
	if err != nil {
		return nil, err
	}
	res := e2apSetupRequest.GetInitiatingMessage().GetProcedureCode().GetE2Setup().GetInitiatingMessage().
		GetProtocolIes().GetE2ApProtocolIes10().GetValue().GetValue()[0].GetE2ApProtocolIes8().GetValue()
	//fmt.Printf("Returning following structure: \n %v \n", res)

	return res, nil
}

func Test_RanFunctionItem(t *testing.T) {

	rfi, err := createRanFunctionItem()
	assert.NilError(t, err)

	xer, err := xerEncodeRanFunctionItem(rfi)
	assert.NilError(t, err)
	t.Logf("RanFunctionList XER\n%s", xer)

	per, err := perEncodeRanFunctionItem(rfi)
	assert.NilError(t, err)
	t.Logf("RanFunctionList PER\n%v", hex.Dump(per))

	// Now reverse the XER
	rfiReversed, err := xerDecodeRanFunctionItem(xer)
	assert.NilError(t, err)
	assert.Assert(t, rfiReversed != nil)
	t.Logf("RanFunctionList decoded from XER is \n%v", rfiReversed)
	//assert.Equal(t, 2, len(rflReversed.GetValue()))

	// Now reverse the PER
	rfiReversedFromPer, err := perDecodeRanFunctionItem(per)
	assert.NilError(t, err)
	assert.Assert(t, rfiReversedFromPer != nil)
	t.Logf("RanFunctionList decoded from PER is \n%v", rfiReversedFromPer)
	//assert.Equal(t, 2, len(rflReversedFromPer.GetValue()))

}
