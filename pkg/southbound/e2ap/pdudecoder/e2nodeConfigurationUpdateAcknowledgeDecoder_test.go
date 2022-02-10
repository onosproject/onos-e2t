// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdudecoder

//func Test_DecodeE2nodeConfigurationUpdateAcknowledgePdu(t *testing.T) {
//	e2ncuaXer, err := ioutil.ReadFile("../test/E2nodeConfigurationUpdateAcknowledge.xml")
//	assert.NilError(t, err, "Unexpected error when loading file")
//	e2apPdu, err := asn1cgo.XerDecodeE2apPdu(e2ncuaXer)
//	assert.NilError(t, err)
//
//	transactionID, additionAckList, err := DecodeE2nodeConfigurationUpdateAcknowledgePdu(e2apPdu)
//	assert.NilError(t, err)
//	//assert.Assert(t, ricIdentity != nil) //Commented due to the Linters (v1.34.1) error - possible nil pointer dereference (https://staticcheck.io/docs/checks#SA5011) on lines 23, 24 & 25
//
//	assert.Equal(t, additionAckList[0].E2NodeComponentType.Number(), e2ap_ies.E2NodeComponentInterfaceType_E2NODE_COMPONENT_INTERFACE_TYPE_F1.Number())
//	assert.Equal(t, additionAckList[0].E2NodeComponentID.GetE2NodeComponentInterfaceTypeF1().GetGNbDuId().GetValue(), int64(21))
//	assert.Equal(t, additionAckList[0].E2NodeComponentConfigurationAck.UpdateOutcome.Number(), e2ap_ies.UpdateOutcome_UPDATE_OUTCOME_FAILURE.Number())
//
//	assert.Equal(t, additionAckList[1].E2NodeComponentType.Number(), e2ap_ies.E2NodeComponentInterfaceType_E2NODE_COMPONENT_INTERFACE_TYPE_XN.Number())
//	assert.DeepEqual(t, additionAckList[1].E2NodeComponentID.GetE2NodeComponentInterfaceTypeXn().GetGlobalNgRanNodeId().GetGNb().GetPlmnId().GetValue(), []byte{0x01, 0x02, 0x03})
//	assert.DeepEqual(t, additionAckList[1].E2NodeComponentID.GetE2NodeComponentInterfaceTypeXn().GetGlobalNgRanNodeId().GetGNb().GetGnbId().GetGnbId().GetValue(), []byte{0xAB, 0xCD, 0xEF, 0xFF})
//	assert.Equal(t, additionAckList[1].E2NodeComponentConfigurationAck.UpdateOutcome.Number(), e2ap_ies.UpdateOutcome_UPDATE_OUTCOME_SUCCESS.Number())
//	assert.Equal(t, additionAckList[1].E2NodeComponentConfigurationAck.FailureCause.GetProtocol().Number(), e2ap_ies.CauseProtocol_CAUSE_PROTOCOL_ABSTRACT_SYNTAX_ERROR_FALSELY_CONSTRUCTED_MESSAGE.Number())
//
//	if transactionID != nil {
//		assert.Equal(t, int32(1), *transactionID)
//	}
//}
