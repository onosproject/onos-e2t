// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//
//func createE2nodeConfigurationUpdateFailureMsg() (*e2ap_pdu_contents.E2NodeConfigurationUpdateFailure, error) {
//
//	// e2nodeConfigurationUpdateFailure := pdubuilder.CreateE2nodeConfigurationUpdateFailure() //ToDo - fill in arguments here(if this function exists
//
//	e2nodeConfigurationUpdateFailure := e2ap_pdu_contents.E2NodeConfigurationUpdateFailure{
//		ProtocolIes: nil,
//	}
//
//	if err := e2nodeConfigurationUpdateFailure.Validate(); err != nil {
//		return nil, fmt.Errorf("error validating E2nodeConfigurationUpdateFailure %s", err.Error())
//	}
//	return &e2nodeConfigurationUpdateFailure, nil
//}
//
//func Test_xerEncodingE2nodeConfigurationUpdateFailure(t *testing.T) {
//
//	e2nodeConfigurationUpdateFailure, err := createE2nodeConfigurationUpdateFailureMsg()
//	assert.NilError(t, err, "Error creating E2nodeConfigurationUpdateFailure PDU")
//
//	xer, err := xerEncodeE2nodeConfigurationUpdateFailure(e2nodeConfigurationUpdateFailure)
//	assert.NilError(t, err)
//	assert.Equal(t, 1, len(xer)) //ToDo - adjust length of the XER encoded message
//	t.Logf("E2nodeConfigurationUpdateFailure XER\n%s", string(xer))
//
//	result, err := xerDecodeE2nodeConfigurationUpdateFailure(xer)
//	assert.NilError(t, err)
//	assert.Assert(t, result != nil)
//	t.Logf("E2nodeConfigurationUpdateFailure XER - decoded\n%v", result)
//	//ToDo - adjust field's verification
//	assert.Equal(t, e2nodeConfigurationUpdateFailure.GetProtocolIes(), result.GetProtocolIes())
//
//}
//
//func Test_perEncodingE2nodeConfigurationUpdateFailure(t *testing.T) {
//
//	e2nodeConfigurationUpdateFailure, err := createE2nodeConfigurationUpdateFailureMsg()
//	assert.NilError(t, err, "Error creating E2nodeConfigurationUpdateFailure PDU")
//
//	per, err := perEncodeE2nodeConfigurationUpdateFailure(e2nodeConfigurationUpdateFailure)
//	assert.NilError(t, err)
//	assert.Equal(t, 1, len(per)) // ToDo - adjust length of the PER encoded message
//	t.Logf("E2nodeConfigurationUpdateFailure PER\n%v", hex.Dump(per))
//
//	result, err := perDecodeE2nodeConfigurationUpdateFailure(per)
//	assert.NilError(t, err)
//	assert.Assert(t, result != nil)
//	t.Logf("E2nodeConfigurationUpdateFailure PER - decoded\n%v", result)
//	//ToDo - adjust field's verification
//	assert.Equal(t, e2nodeConfigurationUpdateFailure.GetProtocolIes(), result.GetProtocolIes())
//
//}
