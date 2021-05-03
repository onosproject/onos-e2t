// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//
//func createRicserviceQueryMsg() (*e2ap_pdu_contents.RicserviceQuery, error) {
//
//	// ricserviceQuery := pdubuilder.CreateRicserviceQuery() //ToDo - fill in arguments here(if this function exists
//
//	ricserviceQuery := e2ap_pdu_contents.RicserviceQuery{
//		ProtocolIes: nil,
//	}
//
//	if err := ricserviceQuery.Validate(); err != nil {
//		return nil, fmt.Errorf("error validating RicserviceQuery %s", err.Error())
//	}
//	return &ricserviceQuery, nil
//}
//
//func Test_xerEncodingRicserviceQuery(t *testing.T) {
//
//	ricserviceQuery, err := createRicserviceQueryMsg()
//	assert.NilError(t, err, "Error creating RicserviceQuery PDU")
//
//	xer, err := xerEncodeRicserviceQuery(ricserviceQuery)
//	assert.NilError(t, err)
//	assert.Equal(t, 1, len(xer)) //ToDo - adjust length of the XER encoded message
//	t.Logf("RicserviceQuery XER\n%s", string(xer))
//
//	result, err := xerDecodeRicserviceQuery(xer)
//	assert.NilError(t, err)
//	assert.Assert(t, result != nil)
//	t.Logf("RicserviceQuery XER - decoded\n%v", result)
//	//ToDo - adjust field's verification
//	assert.Equal(t, ricserviceQuery.GetProtocolIes(), result.GetProtocolIes())
//
//}
//
//func Test_perEncodingRicserviceQuery(t *testing.T) {
//
//	ricserviceQuery, err := createRicserviceQueryMsg()
//	assert.NilError(t, err, "Error creating RicserviceQuery PDU")
//
//	per, err := perEncodeRicserviceQuery(ricserviceQuery)
//	assert.NilError(t, err)
//	assert.Equal(t, 1, len(per)) // ToDo - adjust length of the PER encoded message
//	t.Logf("RicserviceQuery PER\n%v", hex.Dump(per))
//
//	result, err := perDecodeRicserviceQuery(per)
//	assert.NilError(t, err)
//	assert.Assert(t, result != nil)
//	t.Logf("RicserviceQuery PER - decoded\n%v", result)
//	//ToDo - adjust field's verification
//	assert.Equal(t, ricserviceQuery.GetProtocolIes(), result.GetProtocolIes())
//
//}
