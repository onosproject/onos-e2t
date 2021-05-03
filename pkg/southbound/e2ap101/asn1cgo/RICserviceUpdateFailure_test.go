// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//
//func createRicserviceUpdateFailureMsg() (*e2ap_pdu_contents.RicserviceUpdateFailure, error) {
//
//	// ricserviceUpdateFailure := pdubuilder.CreateRicserviceUpdateFailure() //ToDo - fill in arguments here(if this function exists
//
//	ricserviceUpdateFailure := e2ap_pdu_contents.RicserviceUpdateFailure{
//		ProtocolIes: nil,
//	}
//
//	if err := ricserviceUpdateFailure.Validate(); err != nil {
//		return nil, fmt.Errorf("error validating RicserviceUpdateFailure %s", err.Error())
//	}
//	return &ricserviceUpdateFailure, nil
//}
//
//func Test_xerEncodingRicserviceUpdateFailure(t *testing.T) {
//
//	ricserviceUpdateFailure, err := createRicserviceUpdateFailureMsg()
//	assert.NilError(t, err, "Error creating RicserviceUpdateFailure PDU")
//
//	xer, err := xerEncodeRicserviceUpdateFailure(ricserviceUpdateFailure)
//	assert.NilError(t, err)
//	assert.Equal(t, 1, len(xer)) //ToDo - adjust length of the XER encoded message
//	t.Logf("RicserviceUpdateFailure XER\n%s", string(xer))
//
//	result, err := xerDecodeRicserviceUpdateFailure(xer)
//	assert.NilError(t, err)
//	assert.Assert(t, result != nil)
//	t.Logf("RicserviceUpdateFailure XER - decoded\n%v", result)
//	//ToDo - adjust field's verification
//	assert.Equal(t, ricserviceUpdateFailure.GetProtocolIes(), result.GetProtocolIes())
//
//}
//
//func Test_perEncodingRicserviceUpdateFailure(t *testing.T) {
//
//	ricserviceUpdateFailure, err := createRicserviceUpdateFailureMsg()
//	assert.NilError(t, err, "Error creating RicserviceUpdateFailure PDU")
//
//	per, err := perEncodeRicserviceUpdateFailure(ricserviceUpdateFailure)
//	assert.NilError(t, err)
//	assert.Equal(t, 1, len(per)) // ToDo - adjust length of the PER encoded message
//	t.Logf("RicserviceUpdateFailure PER\n%v", hex.Dump(per))
//
//	result, err := perDecodeRicserviceUpdateFailure(per)
//	assert.NilError(t, err)
//	assert.Assert(t, result != nil)
//	t.Logf("RicserviceUpdateFailure PER - decoded\n%v", result)
//	//ToDo - adjust field's verification
//	assert.Equal(t, ricserviceUpdateFailure.GetProtocolIes(), result.GetProtocolIes())
//
//}
