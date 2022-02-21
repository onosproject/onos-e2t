// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdudecoder

import (
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/pdubuilder"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"testing"

	"gotest.tools/assert"
)

func Test_GetE2NodeID(t *testing.T) {
	nodeID := []byte{0, 0, 0, 0, 0, 0x01, 0x45, 0x4c}
	id := GetE2NodeID(nodeID, 22)
	assert.Equal(t, id, "5153")

	nodeID2 := []byte{0x01, 0x45, 0x4c}
	id = GetE2NodeID(nodeID2, 22)
	assert.Equal(t, id, "5153")
}

func Test_DecodeE2SetupRequestPduCuDuIDs(t *testing.T) {
	enbID, err := pdubuilder.CreateEnbIDHome(&asn1.BitString{
		Value: []byte{0x00, 0xA7, 0xDD, 0xF0},
		Len:   28,
	})
	assert.NilError(t, err)
	gEnbID, err := pdubuilder.CreateGlobalEnbID([]byte{0xAA, 0xBB, 0xCC}, enbID)
	assert.NilError(t, err)

	gEnGnbID, err := pdubuilder.CreateGlobalEnGnbID([]byte{0xFF, 0xCD, 0xBF}, &asn1.BitString{
		Value: []byte{0xFA, 0x2C, 0xD4, 0xF8},
		Len:   29,
	})
	assert.NilError(t, err)

	e2ncID1 := pdubuilder.CreateE2NodeComponentIDX2(gEnbID, gEnGnbID)
	e2ncID2 := pdubuilder.CreateE2NodeComponentIDNg("NG-Component")

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
	ge2nID.GetGNb().SetGnbCuUpID(2).SetGnbDuID(13)

	e2apPdu, err := pdubuilder.CreateE2SetupRequestPdu(1, ge2nID, ranFunctionList, []*types.E2NodeComponentConfigAdditionItem{
		{E2NodeComponentType: e2ap_ies.E2NodeComponentInterfaceType_E2NODE_COMPONENT_INTERFACE_TYPE_X2,
			E2NodeComponentID: e2ncID1,
			E2NodeComponentConfiguration: e2ap_ies.E2NodeComponentConfiguration{
				E2NodeComponentResponsePart: []byte{0x01, 0x02, 0x03},
				E2NodeComponentRequestPart:  []byte{0x04, 0x05, 0x06},
			}},
		{E2NodeComponentType: e2ap_ies.E2NodeComponentInterfaceType_E2NODE_COMPONENT_INTERFACE_TYPE_NG,
			E2NodeComponentID: e2ncID2,
			E2NodeComponentConfiguration: e2ap_ies.E2NodeComponentConfiguration{
				E2NodeComponentResponsePart: []byte{0x07, 0x08, 0x09},
				E2NodeComponentRequestPart:  []byte{0x0A, 0x0B, 0x0C},
			}},
	})
	assert.NilError(t, err)
	assert.Assert(t, e2apPdu != nil)

	transactionID, identifier, ranFunctions, e2nccul, err := DecodeE2SetupRequestPdu(e2apPdu)
	assert.NilError(t, err)
	//assert.Assert(t, identifier != nil) //Commented due to the Linters (v1.34.1) error - possible nil pointer dereference (https://staticcheck.io/docs/checks#SA5011) on lines 24, 25 & 26
	assert.DeepEqual(t, []byte{0x4f, 0x4e, 0x46}, []byte{identifier.Plmn[0], identifier.Plmn[1], identifier.Plmn[2]})
	assert.Equal(t, types.E2NodeTypeGNB, identifier.NodeType)
	assert.DeepEqual(t, []byte{0x00, 0x00, 0x04}, identifier.NodeIdentifier)
	if identifier.CuID != nil {
		assert.Equal(t, int64(2), *identifier.CuID)
	}
	if identifier.DuID != nil {
		assert.Equal(t, int64(13), *identifier.DuID)
	}

	nodeID := GetE2NodeID(identifier.NodeIdentifier, 30)
	t.Logf("Node ID is %s\n", nodeID)

	assert.Equal(t, 2, len(*ranFunctions))
	rf0 := (*ranFunctions)[100]
	assert.Equal(t, 1, int(rf0.Revision))
	assert.Equal(t, "oid1", string(rf0.OID))
	assert.DeepEqual(t, []byte{0x54, 0x79, 0x70, 0x65, 0x20, 0x31}, []byte(rf0.Description))
	rf1 := (*ranFunctions)[200]
	assert.Equal(t, 2, int(rf1.Revision))
	assert.Equal(t, "oid2", string(rf1.OID))
	assert.DeepEqual(t, []byte{0x54, 0x79, 0x70, 0x65, 0x20, 0x32}, []byte(rf1.Description))

	assert.Equal(t, e2nccul[0].E2NodeComponentType.Number(), e2ap_ies.E2NodeComponentInterfaceType_E2NODE_COMPONENT_INTERFACE_TYPE_X2.Number())
	assert.DeepEqual(t, e2nccul[0].E2NodeComponentID.GetE2NodeComponentInterfaceTypeX2().GetGlobalENbId().GetPLmnIdentity().GetValue(), []byte{0xAA, 0xBB, 0xCC})
	assert.DeepEqual(t, e2nccul[0].E2NodeComponentID.GetE2NodeComponentInterfaceTypeX2().GetGlobalENbId().GetENbId().GetHomeENbId().GetValue(), []byte{0x00, 0xA7, 0xDD, 0xF0})
	assert.DeepEqual(t, e2nccul[0].E2NodeComponentID.GetE2NodeComponentInterfaceTypeX2().GetGlobalEnGNbId().GetPLmnIdentity().GetValue(), []byte{0xFF, 0xCD, 0xBF})
	assert.DeepEqual(t, e2nccul[0].E2NodeComponentID.GetE2NodeComponentInterfaceTypeX2().GetGlobalEnGNbId().GetGNbId().GetGNbId().GetValue(), []byte{0xFA, 0x2C, 0xD4, 0xF8})
	assert.DeepEqual(t, e2nccul[0].E2NodeComponentConfiguration.GetE2NodeComponentResponsePart(), []byte{0x01, 0x02, 0x03})
	assert.DeepEqual(t, e2nccul[0].E2NodeComponentConfiguration.GetE2NodeComponentRequestPart(), []byte{0x04, 0x05, 0x06})
	assert.Equal(t, e2nccul[1].E2NodeComponentType.Number(), e2ap_ies.E2NodeComponentInterfaceType_E2NODE_COMPONENT_INTERFACE_TYPE_NG.Number())
	assert.Equal(t, e2nccul[1].E2NodeComponentID.GetE2NodeComponentInterfaceTypeNg().GetAmfName().GetValue(), "NG-Component")
	assert.DeepEqual(t, e2nccul[1].E2NodeComponentConfiguration.GetE2NodeComponentRequestPart(), []byte{0x0A, 0x0B, 0x0C})
	assert.DeepEqual(t, e2nccul[1].E2NodeComponentConfiguration.GetE2NodeComponentResponsePart(), []byte{0x07, 0x08, 0x09})
	if transactionID != nil {
		assert.Equal(t, int32(1), *transactionID)
	}
}
