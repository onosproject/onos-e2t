// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//
//func createResetResponseMsg() (*e2ap_pdu_contents.ResetResponse, error) {
//
//	// resetResponse := pdubuilder.CreateResetResponse() //ToDo - fill in arguments here(if this function exists
//
//	resetResponse := e2ap_pdu_contents.ResetResponse{
//		ProtocolIes: nil,
//	}
//
//	if err := resetResponse.Validate(); err != nil {
//		return nil, fmt.Errorf("error validating ResetResponse %s", err.Error())
//	}
//	return &resetResponse, nil
//}
//
//func Test_xerEncodingResetResponse(t *testing.T) {
//
//	resetResponse, err := createResetResponseMsg()
//	assert.NilError(t, err, "Error creating ResetResponse PDU")
//
//	xer, err := xerEncodeResetResponse(resetResponse)
//	assert.NilError(t, err)
//	assert.Equal(t, 1, len(xer)) //ToDo - adjust length of the XER encoded message
//	t.Logf("ResetResponse XER\n%s", string(xer))
//
//	result, err := xerDecodeResetResponse(xer)
//	assert.NilError(t, err)
//	assert.Assert(t, result != nil)
//	t.Logf("ResetResponse XER - decoded\n%v", result)
//	//ToDo - adjust field's verification
//	assert.Equal(t, resetResponse.GetProtocolIes(), result.GetProtocolIes())
//
//}
//
//func Test_perEncodingResetResponse(t *testing.T) {
//
//	resetResponse, err := createResetResponseMsg()
//	assert.NilError(t, err, "Error creating ResetResponse PDU")
//
//	per, err := perEncodeResetResponse(resetResponse)
//	assert.NilError(t, err)
//	assert.Equal(t, 1, len(per)) // ToDo - adjust length of the PER encoded message
//	t.Logf("ResetResponse PER\n%v", hex.Dump(per))
//
//	result, err := perDecodeResetResponse(per)
//	assert.NilError(t, err)
//	assert.Assert(t, result != nil)
//	t.Logf("ResetResponse PER - decoded\n%v", result)
//	//ToDo - adjust field's verification
//	assert.Equal(t, resetResponse.GetProtocolIes(), result.GetProtocolIes())
//
//}
