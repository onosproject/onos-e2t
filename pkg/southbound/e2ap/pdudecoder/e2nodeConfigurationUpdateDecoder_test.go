// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdudecoder

//func Test_DecodeE2nodeConfigurationUpdatePdu(t *testing.T) {
//	e2ncuXer, err := ioutil.ReadFile("../test/E2nodeConfigurationUpdate.xml")
//	assert.NilError(t, err, "Unexpected error when loading file")
//	e2apPdu, err := asn1cgo.XerDecodeE2apPdu(e2ncuXer)
//	assert.NilError(t, err)
//
//	transactionID, nodeIdentity, e2nccul, err := DecodeE2nodeConfigurationUpdatePdu(e2apPdu)
//	assert.NilError(t, err)
//	//assert.Assert(t, ricIdentity != nil) //Commented due to the Linters (v1.34.1) error - possible nil pointer dereference (https://staticcheck.io/docs/checks#SA5011) on lines 23, 24 & 25
//
//	assert.Equal(t, e2nccul[0].E2NodeComponentType.Number(), e2apies.E2NodeComponentInterfaceType_E2NODE_COMPONENT_INTERFACE_TYPE_NG.Number())
//	assert.Equal(t, e2nccul[0].E2NodeComponentID.GetE2NodeComponentInterfaceTypeNg().GetAmfName().GetValue(), "NG-Component")
//	assert.DeepEqual(t, e2nccul[0].E2NodeComponentConfiguration.GetE2NodeComponentResponsePart(), []byte{0x01, 0x02, 0x03})
//	assert.DeepEqual(t, e2nccul[0].E2NodeComponentConfiguration.GetE2NodeComponentRequestPart(), []byte{0x04, 0x05, 0x06})
//	assert.Equal(t, e2nccul[1].E2NodeComponentType.Number(), e2apies.E2NodeComponentInterfaceType_E2NODE_COMPONENT_INTERFACE_TYPE_E1.Number())
//	assert.Equal(t, e2nccul[1].E2NodeComponentID.GetE2NodeComponentInterfaceTypeE1().GetGNbCuCpId().GetValue(), int64(13))
//	assert.DeepEqual(t, e2nccul[1].E2NodeComponentConfiguration.GetE2NodeComponentResponsePart(), []byte{0x07, 0x08, 0x09})
//	assert.DeepEqual(t, e2nccul[1].E2NodeComponentConfiguration.GetE2NodeComponentRequestPart(), []byte{0x0A, 0x0B, 0x0C})
//	if transactionID != nil {
//		assert.Equal(t, int32(1), *transactionID)
//	}
//	assert.Equal(t, nodeIdentity.NodeIDLength, 25)
//	assert.DeepEqual(t, nodeIdentity.NodeIdentifier, []byte{0x00, 0x00, 0x00, 0x80})
//	assert.DeepEqual(t, [3]byte(nodeIdentity.Plmn), [3]uint8{0x00, 0x00, 0x01})
//	assert.Equal(t, nodeIdentity.NodeType, types.E2NodeTypeEnGNB)
//	if nodeIdentity.DuID != nil {
//		assert.Equal(t, *nodeIdentity.DuID, int64(13))
//	}
//	if nodeIdentity.CuID != nil {
//		assert.Equal(t, *nodeIdentity.CuID, int64(2))
//	}
//}
