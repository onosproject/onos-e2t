// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//
//func createE2nodeConfigurationUpdateMsg() (*e2ap_pdu_contents.E2NodeConfigurationUpdate, error) {
//
//	// e2nodeConfigurationUpdate := pdubuilder.CreateE2nodeConfigurationUpdate() //ToDo - fill in arguments here(if this function exists
//
//	e2nodeConfigurationUpdate := e2ap_pdu_contents.E2NodeConfigurationUpdate{
//		ProtocolIes: nil,
//	}
//
//	if err := e2nodeConfigurationUpdate.Validate(); err != nil {
//		return nil, fmt.Errorf("error validating E2nodeConfigurationUpdate %s", err.Error())
//	}
//	return &e2nodeConfigurationUpdate, nil
//}
//
//func Test_xerEncodingE2nodeConfigurationUpdate(t *testing.T) {
//
//	e2nodeConfigurationUpdate, err := createE2nodeConfigurationUpdateMsg()
//	assert.NilError(t, err, "Error creating E2nodeConfigurationUpdate PDU")
//
//	xer, err := xerEncodeE2nodeConfigurationUpdate(e2nodeConfigurationUpdate)
//	assert.NilError(t, err)
//	assert.Equal(t, 1, len(xer)) //ToDo - adjust length of the XER encoded message
//	t.Logf("E2nodeConfigurationUpdate XER\n%s", string(xer))
//
//	result, err := xerDecodeE2nodeConfigurationUpdate(xer)
//	assert.NilError(t, err)
//	assert.Assert(t, result != nil)
//	t.Logf("E2nodeConfigurationUpdate XER - decoded\n%v", result)
//	//ToDo - adjust field's verification
//	assert.Equal(t, e2nodeConfigurationUpdate.GetProtocolIes(), result.GetProtocolIes())
//
//}
//
//func Test_perEncodingE2nodeConfigurationUpdate(t *testing.T) {
//
//	e2nodeConfigurationUpdate, err := createE2nodeConfigurationUpdateMsg()
//	assert.NilError(t, err, "Error creating E2nodeConfigurationUpdate PDU")
//
//	per, err := perEncodeE2nodeConfigurationUpdate(e2nodeConfigurationUpdate)
//	assert.NilError(t, err)
//	assert.Equal(t, 1, len(per)) // ToDo - adjust length of the PER encoded message
//	t.Logf("E2nodeConfigurationUpdate PER\n%v", hex.Dump(per))
//
//	result, err := perDecodeE2nodeConfigurationUpdate(per)
//	assert.NilError(t, err)
//	assert.Assert(t, result != nil)
//	t.Logf("E2nodeConfigurationUpdate PER - decoded\n%v", result)
//	//ToDo - adjust field's verification
//	assert.Equal(t, e2nodeConfigurationUpdate.GetProtocolIes(), result.GetProtocolIes())
//
//}
