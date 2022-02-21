// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdudecoder

import (
	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/pdubuilder"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"gotest.tools/assert"
	"testing"
)

func Test_DecodeE2nodeConfigurationUpdatePdu(t *testing.T) {
	ge2nID, err := pdubuilder.CreateGlobalE2nodeIDEnGnb([3]byte{0x00, 0x00, 0x01}, &asn1.BitString{
		Value: []byte{0x00, 0x00, 0x00, 0x80},
		Len:   25,
	})
	assert.NilError(t, err)
	ge2nID.GetEnGNb().SetGnbCuUpID(2).SetGnbDuID(13)

	e2ncID1 := pdubuilder.CreateE2NodeComponentIDNg("NG-Component")
	e2ncID2 := pdubuilder.CreateE2NodeComponentIDE1(13)

	e2nodeConfigurationUpdate, err := pdubuilder.CreateE2NodeConfigurationUpdateE2apPdu(1)
	assert.NilError(t, err)
	assert.Assert(t, e2nodeConfigurationUpdate != nil)

	e2nodeConfigurationUpdate.GetInitiatingMessage().GetValue().GetE2NodeConfigurationUpdate().SetGlobalE2nodeID(ge2nID).
		SetE2nodeComponentConfigUpdate([]*types.E2NodeComponentConfigUpdateItem{
			{E2NodeComponentType: e2apies.E2NodeComponentInterfaceType_E2NODE_COMPONENT_INTERFACE_TYPE_NG,
				E2NodeComponentID: e2ncID1,
				E2NodeComponentConfiguration: e2apies.E2NodeComponentConfiguration{
					E2NodeComponentResponsePart: []byte{0x01, 0x02, 0x03},
					E2NodeComponentRequestPart:  []byte{0x04, 0x05, 0x06},
				}},
			{E2NodeComponentType: e2apies.E2NodeComponentInterfaceType_E2NODE_COMPONENT_INTERFACE_TYPE_E1,
				E2NodeComponentID: e2ncID2,
				E2NodeComponentConfiguration: e2apies.E2NodeComponentConfiguration{
					E2NodeComponentResponsePart: []byte{0x07, 0x08, 0x09},
					E2NodeComponentRequestPart:  []byte{0x0A, 0x0B, 0x0C},
				}},
		}).SetE2nodeComponentConfigAddition([]*types.E2NodeComponentConfigAdditionItem{
		{E2NodeComponentType: e2apies.E2NodeComponentInterfaceType_E2NODE_COMPONENT_INTERFACE_TYPE_NG,
			E2NodeComponentID: e2ncID1,
			E2NodeComponentConfiguration: e2apies.E2NodeComponentConfiguration{
				E2NodeComponentResponsePart: []byte{0x01, 0x02, 0x03},
				E2NodeComponentRequestPart:  []byte{0x04, 0x05, 0x06},
			}},
		{E2NodeComponentType: e2apies.E2NodeComponentInterfaceType_E2NODE_COMPONENT_INTERFACE_TYPE_E1,
			E2NodeComponentID: e2ncID2,
			E2NodeComponentConfiguration: e2apies.E2NodeComponentConfiguration{
				E2NodeComponentResponsePart: []byte{0x07, 0x08, 0x09},
				E2NodeComponentRequestPart:  []byte{0x0A, 0x0B, 0x0C},
			}},
	}).SetE2nodeComponentConfigRemoval([]*types.E2NodeComponentConfigRemovalItem{
		{E2NodeComponentType: e2apies.E2NodeComponentInterfaceType_E2NODE_COMPONENT_INTERFACE_TYPE_NG,
			E2NodeComponentID: e2ncID1,
		},
		{E2NodeComponentType: e2apies.E2NodeComponentInterfaceType_E2NODE_COMPONENT_INTERFACE_TYPE_E1,
			E2NodeComponentID: e2ncID2,
		},
	}).SetE2nodeTnlAssociationRemoval([]*types.TnlAssociationRemovalItem{
		{TnlInformation: types.TnlInformation{
			TnlAddress: asn1.BitString{
				Value: []byte{0xF0, 0xAB, 0x34, 0x9F},
				Len:   32,
			},
			TnlPort: &asn1.BitString{
				Value: []byte{0x00, 0x02},
				Len:   16,
			},
		},
			TnlInformationRic: types.TnlInformation{
				TnlAddress: asn1.BitString{
					Value: []byte{0xF0, 0xAB, 0x34, 0x9F},
					Len:   32,
				},
				TnlPort: &asn1.BitString{
					Value: []byte{0x00, 0x02},
					Len:   16,
				},
			},
		},
	})

	transactionID, nodeIdentity, e2nccul, err := DecodeE2nodeConfigurationUpdatePdu(e2nodeConfigurationUpdate)
	assert.NilError(t, err)
	//assert.Assert(t, ricIdentity != nil) //Commented due to the Linters (v1.34.1) error - possible nil pointer dereference (https://staticcheck.io/docs/checks#SA5011) on lines 23, 24 & 25

	assert.Equal(t, e2nccul[0].E2NodeComponentType.Number(), e2apies.E2NodeComponentInterfaceType_E2NODE_COMPONENT_INTERFACE_TYPE_NG.Number())
	assert.Equal(t, e2nccul[0].E2NodeComponentID.GetE2NodeComponentInterfaceTypeNg().GetAmfName().GetValue(), "NG-Component")
	assert.DeepEqual(t, e2nccul[0].E2NodeComponentConfiguration.GetE2NodeComponentResponsePart(), []byte{0x01, 0x02, 0x03})
	assert.DeepEqual(t, e2nccul[0].E2NodeComponentConfiguration.GetE2NodeComponentRequestPart(), []byte{0x04, 0x05, 0x06})
	assert.Equal(t, e2nccul[1].E2NodeComponentType.Number(), e2apies.E2NodeComponentInterfaceType_E2NODE_COMPONENT_INTERFACE_TYPE_E1.Number())
	assert.Equal(t, e2nccul[1].E2NodeComponentID.GetE2NodeComponentInterfaceTypeE1().GetGNbCuCpId().GetValue(), int64(13))
	assert.DeepEqual(t, e2nccul[1].E2NodeComponentConfiguration.GetE2NodeComponentResponsePart(), []byte{0x07, 0x08, 0x09})
	assert.DeepEqual(t, e2nccul[1].E2NodeComponentConfiguration.GetE2NodeComponentRequestPart(), []byte{0x0A, 0x0B, 0x0C})
	if transactionID != nil {
		assert.Equal(t, int32(1), *transactionID)
	}
	assert.Equal(t, nodeIdentity.NodeIDLength, 25)
	assert.DeepEqual(t, nodeIdentity.NodeIdentifier, []byte{0x00, 0x00, 0x00, 0x80})
	assert.DeepEqual(t, [3]byte(nodeIdentity.Plmn), [3]uint8{0x00, 0x00, 0x01})
	assert.Equal(t, nodeIdentity.NodeType, types.E2NodeTypeEnGNB)
	if nodeIdentity.DuID != nil {
		assert.Equal(t, *nodeIdentity.DuID, int64(13))
	}
	if nodeIdentity.CuID != nil {
		assert.Equal(t, *nodeIdentity.CuID, int64(2))
	}
}
