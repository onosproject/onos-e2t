// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//func createE2connectionUpdateListMsg() (*e2ap_pdu_contents.E2ConnectionUpdateList, error) {
//
//	// e2connectionUpdateList := pdubuilder.CreateE2connectionUpdateList() //ToDo - fill in arguments here(if this function exists
//
//	e2connectionUpdateList := e2ap_pdu_contents.E2ConnectionUpdateList{}
//
//	if err := e2connectionUpdateList.Validate(); err != nil {
//		return nil, fmt.Errorf("error validating E2connectionUpdateList %s", err.Error())
//	}
//	return &e2connectionUpdateList, nil
//}
//
//func Test_xerEncodingE2connectionUpdateList(t *testing.T) {
//
//	e2connectionUpdateList, err := createE2connectionUpdateListMsg()
//	assert.NilError(t, err, "Error creating E2connectionUpdateList PDU")
//
//	xer, err := xerEncodeE2connectionUpdateList(e2connectionUpdateList)
//	assert.NilError(t, err)
//	assert.Equal(t, 1, len(xer)) //ToDo - adjust length of the XER encoded message
//	t.Logf("E2connectionUpdateList XER\n%s", string(xer))
//
//	result, err := xerDecodeE2connectionUpdateList(xer)
//	assert.NilError(t, err)
//	assert.Assert(t, result != nil)
//	t.Logf("E2connectionUpdateList XER - decoded\n%v", result)
//	//ToDo - adjust field's verification
//
//	assert.Equal(t, 1, len(result.GetValue())) //ToDo - adjust length of a list
//	assert.DeepEqual(t, e2connectionUpdateList.GetValue(), result.GetValue())
//
//}
//
//func Test_perEncodingE2connectionUpdateList(t *testing.T) {
//
//	e2connectionUpdateList, err := createE2connectionUpdateListMsg()
//	assert.NilError(t, err, "Error creating E2connectionUpdateList PDU")
//
//	per, err := perEncodeE2connectionUpdateList(e2connectionUpdateList)
//	assert.NilError(t, err)
//	assert.Equal(t, 1, len(per)) // ToDo - adjust length of the PER encoded message
//	t.Logf("E2connectionUpdateList PER\n%v", hex.Dump(per))
//
//	result, err := perDecodeE2connectionUpdateList(per)
//	assert.NilError(t, err)
//	assert.Assert(t, result != nil)
//	t.Logf("E2connectionUpdateList PER - decoded\n%v", result)
//	//ToDo - adjust field's verification
//
//	assert.Equal(t, 1, len(result.GetValue())) //ToDo - adjust length of a list
//	assert.DeepEqual(t, e2connectionUpdateList.GetValue(), result.GetValue())
//
//}
