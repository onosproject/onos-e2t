// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//
//func createRicserviceUpdateAcknowledgeMsg() (*e2ap_pdu_contents.RicserviceUpdateAcknowledge, error) {
//
//	// ricserviceUpdateAcknowledge := pdubuilder.CreateRicserviceUpdateAcknowledge() //ToDo - fill in arguments here(if this function exists
//
//	ricserviceUpdateAcknowledge := e2ap_pdu_contents.RicserviceUpdateAcknowledge{
//		ProtocolIes: nil,
//	}
//
//	if err := ricserviceUpdateAcknowledge.Validate(); err != nil {
//		return nil, fmt.Errorf("error validating RicserviceUpdateAcknowledge %s", err.Error())
//	}
//	return &ricserviceUpdateAcknowledge, nil
//}
//
//func Test_xerEncodingRicserviceUpdateAcknowledge(t *testing.T) {
//
//	ricserviceUpdateAcknowledge, err := createRicserviceUpdateAcknowledgeMsg()
//	assert.NilError(t, err, "Error creating RicserviceUpdateAcknowledge PDU")
//
//	xer, err := xerEncodeRicserviceUpdateAcknowledge(ricserviceUpdateAcknowledge)
//	assert.NilError(t, err)
//	assert.Equal(t, 1, len(xer)) //ToDo - adjust length of the XER encoded message
//	t.Logf("RicserviceUpdateAcknowledge XER\n%s", string(xer))
//
//	result, err := xerDecodeRicserviceUpdateAcknowledge(xer)
//	assert.NilError(t, err)
//	assert.Assert(t, result != nil)
//	t.Logf("RicserviceUpdateAcknowledge XER - decoded\n%v", result)
//	//ToDo - adjust field's verification
//	assert.Equal(t, ricserviceUpdateAcknowledge.GetProtocolIes(), result.GetProtocolIes())
//
//}
//
//func Test_perEncodingRicserviceUpdateAcknowledge(t *testing.T) {
//
//	ricserviceUpdateAcknowledge, err := createRicserviceUpdateAcknowledgeMsg()
//	assert.NilError(t, err, "Error creating RicserviceUpdateAcknowledge PDU")
//
//	per, err := perEncodeRicserviceUpdateAcknowledge(ricserviceUpdateAcknowledge)
//	assert.NilError(t, err)
//	assert.Equal(t, 1, len(per)) // ToDo - adjust length of the PER encoded message
//	t.Logf("RicserviceUpdateAcknowledge PER\n%v", hex.Dump(per))
//
//	result, err := perDecodeRicserviceUpdateAcknowledge(per)
//	assert.NilError(t, err)
//	assert.Assert(t, result != nil)
//	t.Logf("RicserviceUpdateAcknowledge PER - decoded\n%v", result)
//	//ToDo - adjust field's verification
//	assert.Equal(t, ricserviceUpdateAcknowledge.GetProtocolIes(), result.GetProtocolIes())
//
//}
