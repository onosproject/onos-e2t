// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//
//func createE2connectionUpdateFailureMsg() (*e2ap_pdu_contents.E2ConnectionUpdateFailure, error) {
//
//	// e2connectionUpdateFailure := pdubuilder.CreateE2connectionUpdateFailure() //ToDo - fill in arguments here(if this function exists
//
//	e2connectionUpdateFailure := e2ap_pdu_contents.E2ConnectionUpdateFailure{
//		ProtocolIes: nil,
//	}
//
//	if err := e2connectionUpdateFailure.Validate(); err != nil {
//		return nil, fmt.Errorf("error validating E2connectionUpdateFailure %s", err.Error())
//	}
//	return &e2connectionUpdateFailure, nil
//}
//
//func Test_xerEncodingE2connectionUpdateFailure(t *testing.T) {
//
//	e2connectionUpdateFailure, err := createE2connectionUpdateFailureMsg()
//	assert.NilError(t, err, "Error creating E2connectionUpdateFailure PDU")
//
//	xer, err := xerEncodeE2connectionUpdateFailure(e2connectionUpdateFailure)
//	assert.NilError(t, err)
//	assert.Equal(t, 1, len(xer)) //ToDo - adjust length of the XER encoded message
//	t.Logf("E2connectionUpdateFailure XER\n%s", string(xer))
//
//	result, err := xerDecodeE2connectionUpdateFailure(xer)
//	assert.NilError(t, err)
//	assert.Assert(t, result != nil)
//	t.Logf("E2connectionUpdateFailure XER - decoded\n%v", result)
//	//ToDo - adjust field's verification
//	assert.Equal(t, e2connectionUpdateFailure.GetProtocolIes(), result.GetProtocolIes())
//
//}
//
//func Test_perEncodingE2connectionUpdateFailure(t *testing.T) {
//
//	e2connectionUpdateFailure, err := createE2connectionUpdateFailureMsg()
//	assert.NilError(t, err, "Error creating E2connectionUpdateFailure PDU")
//
//	per, err := perEncodeE2connectionUpdateFailure(e2connectionUpdateFailure)
//	assert.NilError(t, err)
//	assert.Equal(t, 1, len(per)) // ToDo - adjust length of the PER encoded message
//	t.Logf("E2connectionUpdateFailure PER\n%v", hex.Dump(per))
//
//	result, err := perDecodeE2connectionUpdateFailure(per)
//	assert.NilError(t, err)
//	assert.Assert(t, result != nil)
//	t.Logf("E2connectionUpdateFailure PER - decoded\n%v", result)
//	//ToDo - adjust field's verification
//	assert.Equal(t, e2connectionUpdateFailure.GetProtocolIes(), result.GetProtocolIes())
//
//}
