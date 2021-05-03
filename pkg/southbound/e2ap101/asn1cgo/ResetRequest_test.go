// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//
//func createResetRequestMsg() (*e2ap_pdu_contents.ResetRequest, error) {
//
//	// resetRequest := pdubuilder.CreateResetRequest() //ToDo - fill in arguments here(if this function exists
//
//	resetRequest := e2ap_pdu_contents.ResetRequest{
//		ProtocolIes: nil,
//	}
//
//	if err := resetRequest.Validate(); err != nil {
//		return nil, fmt.Errorf("error validating ResetRequest %s", err.Error())
//	}
//	return &resetRequest, nil
//}
//
//func Test_xerEncodingResetRequest(t *testing.T) {
//
//	resetRequest, err := createResetRequestMsg()
//	assert.NilError(t, err, "Error creating ResetRequest PDU")
//
//	xer, err := xerEncodeResetRequest(resetRequest)
//	assert.NilError(t, err)
//	assert.Equal(t, 1, len(xer)) //ToDo - adjust length of the XER encoded message
//	t.Logf("ResetRequest XER\n%s", string(xer))
//
//	result, err := xerDecodeResetRequest(xer)
//	assert.NilError(t, err)
//	assert.Assert(t, result != nil)
//	t.Logf("ResetRequest XER - decoded\n%v", result)
//	//ToDo - adjust field's verification
//	assert.Equal(t, resetRequest.GetProtocolIes(), result.GetProtocolIes())
//
//}
//
//func Test_perEncodingResetRequest(t *testing.T) {
//
//	resetRequest, err := createResetRequestMsg()
//	assert.NilError(t, err, "Error creating ResetRequest PDU")
//
//	per, err := perEncodeResetRequest(resetRequest)
//	assert.NilError(t, err)
//	assert.Equal(t, 1, len(per)) // ToDo - adjust length of the PER encoded message
//	t.Logf("ResetRequest PER\n%v", hex.Dump(per))
//
//	result, err := perDecodeResetRequest(per)
//	assert.NilError(t, err)
//	assert.Assert(t, result != nil)
//	t.Logf("ResetRequest PER - decoded\n%v", result)
//	//ToDo - adjust field's verification
//	assert.Equal(t, resetRequest.GetProtocolIes(), result.GetProtocolIes())
//
//}
