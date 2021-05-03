// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//
//func createE2connectionUpdateMsg() (*e2ap_pdu_contents.E2ConnectionUpdate, error) {
//
//	// e2connectionUpdate := pdubuilder.CreateE2connectionUpdate() //ToDo - fill in arguments here(if this function exists
//
//	e2connectionUpdate := e2ap_pdu_contents.E2ConnectionUpdate{
//		ProtocolIes: nil,
//	}
//
//	if err := e2connectionUpdate.Validate(); err != nil {
//		return nil, fmt.Errorf("error validating E2connectionUpdate %s", err.Error())
//	}
//	return &e2connectionUpdate, nil
//}
//
//func Test_xerEncodingE2connectionUpdate(t *testing.T) {
//
//	e2connectionUpdate, err := createE2connectionUpdateMsg()
//	assert.NilError(t, err, "Error creating E2connectionUpdate PDU")
//
//	xer, err := xerEncodeE2connectionUpdate(e2connectionUpdate)
//	assert.NilError(t, err)
//	assert.Equal(t, 1, len(xer)) //ToDo - adjust length of the XER encoded message
//	t.Logf("E2connectionUpdate XER\n%s", string(xer))
//
//	result, err := xerDecodeE2connectionUpdate(xer)
//	assert.NilError(t, err)
//	assert.Assert(t, result != nil)
//	t.Logf("E2connectionUpdate XER - decoded\n%v", result)
//	//ToDo - adjust field's verification
//	assert.Equal(t, e2connectionUpdate.GetProtocolIes(), result.GetProtocolIes())
//
//}
//
//func Test_perEncodingE2connectionUpdate(t *testing.T) {
//
//	e2connectionUpdate, err := createE2connectionUpdateMsg()
//	assert.NilError(t, err, "Error creating E2connectionUpdate PDU")
//
//	per, err := perEncodeE2connectionUpdate(e2connectionUpdate)
//	assert.NilError(t, err)
//	assert.Equal(t, 1, len(per)) // ToDo - adjust length of the PER encoded message
//	t.Logf("E2connectionUpdate PER\n%v", hex.Dump(per))
//
//	result, err := perDecodeE2connectionUpdate(per)
//	assert.NilError(t, err)
//	assert.Assert(t, result != nil)
//	t.Logf("E2connectionUpdate PER - decoded\n%v", result)
//	//ToDo - adjust field's verification
//	assert.Equal(t, e2connectionUpdate.GetProtocolIes(), result.GetProtocolIes())
//
//}
