// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdudecoder

//func Test_DecodeE2SetupResponsePdu(t *testing.T) {
//	e2setupResponseXer, err := ioutil.ReadFile("../test/E2setupResponse.xml")
//	assert.NilError(t, err, "Unexpected error when loading file")
//	e2apPdu, err := asn1cgo.XerDecodeE2apPdu(e2setupResponseXer)
//	assert.NilError(t, err)
//
//	transactionID, ricIdentity, ranFunctionsAccepted, ranFunctionsRejected, additionAckList, err := DecodeE2SetupResponsePdu(e2apPdu)
//	assert.NilError(t, err)
//	//assert.Assert(t, ricIdentity != nil) //Commented due to the Linters (v1.34.1) error - possible nil pointer dereference (https://staticcheck.io/docs/checks#SA5011) on lines 23, 24 & 25
//	assert.DeepEqual(t, []byte{0x79, 0x78, 0x70}, []byte{ricIdentity.PlmnID[0], ricIdentity.PlmnID[1], ricIdentity.PlmnID[2]})
//	assert.Equal(t, 20, int(ricIdentity.RicIdentifier.RicIdentifierLen))
//	assert.DeepEqual(t, []byte{0x4d, 0x20, 0x00}, []byte(ricIdentity.RicIdentifier.RicIdentifierValue))
//
//	assert.Equal(t, 2, len(ranFunctionsAccepted))
//	rfa100, ok := ranFunctionsAccepted[100]
//	assert.Assert(t, ok, "expected a key '100'")
//	assert.Equal(t, 2, int(rfa100))
//	rfa200, ok := ranFunctionsAccepted[200]
//	assert.Assert(t, ok, "expected a key '200'")
//	assert.Equal(t, 2, int(rfa200))
//
//	assert.Equal(t, 2, len(ranFunctionsRejected))
//	rfr101, ok := ranFunctionsRejected[101]
//	assert.Assert(t, ok, "expected a key '101'")
//	assert.Equal(t, "CAUSE_MISC_OM_INTERVENTION", rfr101.GetMisc().String())
//	rfr102, ok := ranFunctionsRejected[102]
//	assert.Assert(t, ok, "expected a key '102'")
//	assert.Equal(t, "CAUSE_PROTOCOL_TRANSFER_SYNTAX_ERROR", rfr102.GetProtocol().String())
//
//	if transactionID != nil {
//		assert.Equal(t, int32(1), *transactionID)
//	}
//	assert.Equal(t, additionAckList[0].E2NodeComponentType.Number(), e2ap_ies.E2NodeComponentInterfaceType_E2NODE_COMPONENT_INTERFACE_TYPE_S1.Number())
//	assert.Equal(t, additionAckList[0].E2NodeComponentID.GetE2NodeComponentInterfaceTypeS1().GetMmeName().GetValue(), "S1-component")
//	assert.Equal(t, additionAckList[0].E2NodeComponentConfigurationAck.UpdateOutcome.Number(), e2ap_ies.UpdateOutcome_UPDATE_OUTCOME_SUCCESS.Number())
//	//assert.Equal(t, int32(additionAckList[0].E2NodeComponentConfigurationAck.FailureCause.GetProtocol()), int32(e2ap_ies.CauseProtocol_CAUSE_PROTOCOL_TRANSFER_SYNTAX_ERROR))
//}

//func Test_DecodeE2SetupResponsePduNoOptional(t *testing.T) {
//	e2setupResponseXer, err := ioutil.ReadFile("../test/E2setupResponse2.xml")
//	assert.NilError(t, err, "Unexpected error when loading file")
//	e2apPdu, err := asn1cgo.XerDecodeE2apPdu(e2setupResponseXer)
//	assert.NilError(t, err)
//
//	transactionID, ricIdentity, ranFunctionsAccepted, ranFunctionsRejected, additionAckList, err := DecodeE2SetupResponsePdu(e2apPdu)
//	assert.NilError(t, err)
//	//assert.Assert(t, ricIdentity != nil) //Commented due to the Linters (v1.34.1) error - possible nil pointer dereference (https://staticcheck.io/docs/checks#SA5011) on lines 54, 55 & 56
//	assert.DeepEqual(t, []byte{0x79, 0x78, 0x70}, []byte{ricIdentity.PlmnID[0], ricIdentity.PlmnID[1], ricIdentity.PlmnID[2]})
//	assert.Equal(t, 20, int(ricIdentity.RicIdentifier.RicIdentifierLen))
//	assert.DeepEqual(t, []byte{0x4d, 0x20, 0x00}, []byte(ricIdentity.RicIdentifier.RicIdentifierValue))
//
//	assert.Equal(t, 0, len(ranFunctionsAccepted))
//	assert.Equal(t, 0, len(ranFunctionsRejected))
//
//	assert.Equal(t, additionAckList[0].E2NodeComponentType.Number(), e2ap_ies.E2NodeComponentInterfaceType_E2NODE_COMPONENT_INTERFACE_TYPE_S1.Number())
//	assert.Equal(t, additionAckList[0].E2NodeComponentID.GetE2NodeComponentInterfaceTypeS1().GetMmeName().GetValue(), "S1-component")
//	assert.Equal(t, additionAckList[0].E2NodeComponentConfigurationAck.UpdateOutcome.Number(), e2ap_ies.UpdateOutcome_UPDATE_OUTCOME_SUCCESS.Number())
//
//	if transactionID != nil {
//		assert.Equal(t, int32(11), *transactionID)
//	}
//}
