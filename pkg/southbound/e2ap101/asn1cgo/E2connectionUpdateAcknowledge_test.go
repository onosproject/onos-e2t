// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//
//func createE2connectionUpdateAcknowledgeMsg() (*e2ap_pdu_contents.E2ConnectionUpdateAcknowledge, error) {
//
//	// e2connectionUpdateAcknowledge := pdubuilder.CreateE2connectionUpdateAcknowledge() //ToDo - fill in arguments here(if this function exists
//
//	e2connectionUpdateAcknowledge := e2ap_pdu_contents.E2ConnectionUpdateAcknowledge{
//		ProtocolIes: nil,
//	}
//
//	if err := e2connectionUpdateAcknowledge.Validate(); err != nil {
//		return nil, fmt.Errorf("error validating E2connectionUpdateAcknowledge %s", err.Error())
//	}
//	return &e2connectionUpdateAcknowledge, nil
//}
//
//func Test_xerEncodingE2connectionUpdateAcknowledge(t *testing.T) {
//
//	e2connectionUpdateAcknowledge, err := createE2connectionUpdateAcknowledgeMsg()
//	assert.NilError(t, err, "Error creating E2connectionUpdateAcknowledge PDU")
//
//	xer, err := xerEncodeE2connectionUpdateAcknowledge(e2connectionUpdateAcknowledge)
//	assert.NilError(t, err)
//	assert.Equal(t, 1, len(xer)) //ToDo - adjust length of the XER encoded message
//	t.Logf("E2connectionUpdateAcknowledge XER\n%s", string(xer))
//
//	result, err := xerDecodeE2connectionUpdateAcknowledge(xer)
//	assert.NilError(t, err)
//	assert.Assert(t, result != nil)
//	t.Logf("E2connectionUpdateAcknowledge XER - decoded\n%v", result)
//	//ToDo - adjust field's verification
//	assert.Equal(t, e2connectionUpdateAcknowledge.GetProtocolIes(), result.GetProtocolIes())
//
//}
//
//func Test_perEncodingE2connectionUpdateAcknowledge(t *testing.T) {
//
//	e2connectionUpdateAcknowledge, err := createE2connectionUpdateAcknowledgeMsg()
//	assert.NilError(t, err, "Error creating E2connectionUpdateAcknowledge PDU")
//
//	per, err := perEncodeE2connectionUpdateAcknowledge(e2connectionUpdateAcknowledge)
//	assert.NilError(t, err)
//	assert.Equal(t, 1, len(per)) // ToDo - adjust length of the PER encoded message
//	t.Logf("E2connectionUpdateAcknowledge PER\n%v", hex.Dump(per))
//
//	result, err := perDecodeE2connectionUpdateAcknowledge(per)
//	assert.NilError(t, err)
//	assert.Assert(t, result != nil)
//	t.Logf("E2connectionUpdateAcknowledge PER - decoded\n%v", result)
//	//ToDo - adjust field's verification
//	assert.Equal(t, e2connectionUpdateAcknowledge.GetProtocolIes(), result.GetProtocolIes())
//
//}
