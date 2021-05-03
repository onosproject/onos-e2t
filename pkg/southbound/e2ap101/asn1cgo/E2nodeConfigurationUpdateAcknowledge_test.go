// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//
//func createE2nodeConfigurationUpdateAcknowledgeMsg() (*e2ap_pdu_contents.E2NodeConfigurationUpdateAcknowledge, error) {
//
//	// e2nodeConfigurationUpdateAcknowledge := pdubuilder.CreateE2nodeConfigurationUpdateAcknowledge() //ToDo - fill in arguments here(if this function exists
//
//	e2nodeConfigurationUpdateAcknowledge := e2ap_pdu_contents.E2NodeConfigurationUpdateAcknowledge{
//		ProtocolIes: nil,
//	}
//
//	if err := e2nodeConfigurationUpdateAcknowledge.Validate(); err != nil {
//		return nil, fmt.Errorf("error validating E2nodeConfigurationUpdateAcknowledge %s", err.Error())
//	}
//	return &e2nodeConfigurationUpdateAcknowledge, nil
//}
//
//func Test_xerEncodingE2nodeConfigurationUpdateAcknowledge(t *testing.T) {
//
//	e2nodeConfigurationUpdateAcknowledge, err := createE2nodeConfigurationUpdateAcknowledgeMsg()
//	assert.NilError(t, err, "Error creating E2nodeConfigurationUpdateAcknowledge PDU")
//
//	xer, err := xerEncodeE2nodeConfigurationUpdateAcknowledge(e2nodeConfigurationUpdateAcknowledge)
//	assert.NilError(t, err)
//	assert.Equal(t, 1, len(xer)) //ToDo - adjust length of the XER encoded message
//	t.Logf("E2nodeConfigurationUpdateAcknowledge XER\n%s", string(xer))
//
//	result, err := xerDecodeE2nodeConfigurationUpdateAcknowledge(xer)
//	assert.NilError(t, err)
//	assert.Assert(t, result != nil)
//	t.Logf("E2nodeConfigurationUpdateAcknowledge XER - decoded\n%v", result)
//	//ToDo - adjust field's verification
//	assert.Equal(t, e2nodeConfigurationUpdateAcknowledge.GetProtocolIes(), result.GetProtocolIes())
//
//}
//
//func Test_perEncodingE2nodeConfigurationUpdateAcknowledge(t *testing.T) {
//
//	e2nodeConfigurationUpdateAcknowledge, err := createE2nodeConfigurationUpdateAcknowledgeMsg()
//	assert.NilError(t, err, "Error creating E2nodeConfigurationUpdateAcknowledge PDU")
//
//	per, err := perEncodeE2nodeConfigurationUpdateAcknowledge(e2nodeConfigurationUpdateAcknowledge)
//	assert.NilError(t, err)
//	assert.Equal(t, 1, len(per)) // ToDo - adjust length of the PER encoded message
//	t.Logf("E2nodeConfigurationUpdateAcknowledge PER\n%v", hex.Dump(per))
//
//	result, err := perDecodeE2nodeConfigurationUpdateAcknowledge(per)
//	assert.NilError(t, err)
//	assert.Assert(t, result != nil)
//	t.Logf("E2nodeConfigurationUpdateAcknowledge PER - decoded\n%v", result)
//	//ToDo - adjust field's verification
//	assert.Equal(t, e2nodeConfigurationUpdateAcknowledge.GetProtocolIes(), result.GetProtocolIes())
//
//}
