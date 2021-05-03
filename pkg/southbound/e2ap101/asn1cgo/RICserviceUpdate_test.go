// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//
//func createRicserviceUpdateMsg() (*e2ap_pdu_contents.RicserviceUpdate, error) {
//
//	// ricserviceUpdate := pdubuilder.CreateRicserviceUpdate() //ToDo - fill in arguments here(if this function exists
//
//	ricserviceUpdate := e2ap_pdu_contents.RicserviceUpdate{
//		ProtocolIes: nil,
//	}
//
//	if err := ricserviceUpdate.Validate(); err != nil {
//		return nil, fmt.Errorf("error validating RicserviceUpdate %s", err.Error())
//	}
//	return &ricserviceUpdate, nil
//}
//
//func Test_xerEncodingRicserviceUpdate(t *testing.T) {
//
//	ricserviceUpdate, err := createRicserviceUpdateMsg()
//	assert.NilError(t, err, "Error creating RicserviceUpdate PDU")
//
//	xer, err := xerEncodeRicserviceUpdate(ricserviceUpdate)
//	assert.NilError(t, err)
//	assert.Equal(t, 1, len(xer)) //ToDo - adjust length of the XER encoded message
//	t.Logf("RicserviceUpdate XER\n%s", string(xer))
//
//	result, err := xerDecodeRicserviceUpdate(xer)
//	assert.NilError(t, err)
//	assert.Assert(t, result != nil)
//	t.Logf("RicserviceUpdate XER - decoded\n%v", result)
//	//ToDo - adjust field's verification
//	assert.Equal(t, ricserviceUpdate.GetProtocolIes(), result.GetProtocolIes())
//
//}
//
//func Test_perEncodingRicserviceUpdate(t *testing.T) {
//
//	ricserviceUpdate, err := createRicserviceUpdateMsg()
//	assert.NilError(t, err, "Error creating RicserviceUpdate PDU")
//
//	per, err := perEncodeRicserviceUpdate(ricserviceUpdate)
//	assert.NilError(t, err)
//	assert.Equal(t, 1, len(per)) // ToDo - adjust length of the PER encoded message
//	t.Logf("RicserviceUpdate PER\n%v", hex.Dump(per))
//
//	result, err := perDecodeRicserviceUpdate(per)
//	assert.NilError(t, err)
//	assert.Assert(t, result != nil)
//	t.Logf("RicserviceUpdate PER - decoded\n%v", result)
//	//ToDo - adjust field's verification
//	assert.Equal(t, ricserviceUpdate.GetProtocolIes(), result.GetProtocolIes())
//
//}
