// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

import (
	"encoding/hex"
	"testing"

	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	e2appdudescriptions "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-descriptions"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/pdubuilder"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"gotest.tools/assert"
)

func Test_RanFunctionsList(t *testing.T) {

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
	assert.NilError(t, err)

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
	assert.NilError(t, err)

	im := e2apSetupRequest.GetE2ApPdu().(*e2appdudescriptions.E2ApPdu_InitiatingMessage)
	rflist := im.InitiatingMessage.GetProcedureCode().GetE2Setup().GetInitiatingMessage().GetProtocolIes().GetE2ApProtocolIes10().GetValue()
	xer, err := xerEncodeRanFunctionsList(rflist)
	assert.NilError(t, err)
	t.Logf("RanFunctionList XER\n%s", xer)

	per, err := perEncodeRanFunctionsList(rflist)
	assert.NilError(t, err)
	t.Logf("RanFunctionList PER\n%v", hex.Dump(per))

	// Now reverse the XER
	rflReversed, err := xerDecodeRanFunctionList(xer)
	assert.NilError(t, err)
	assert.Assert(t, rflReversed != nil)

	assert.Equal(t, 2, len(rflReversed.GetValue()))

	// Now reverse the PER
	rflReversedFromPer, err := perDecodeRanFunctionList(per)
	assert.NilError(t, err)
	assert.Assert(t, rflReversedFromPer != nil)
	assert.Equal(t, 2, len(rflReversedFromPer.GetValue()))

}
