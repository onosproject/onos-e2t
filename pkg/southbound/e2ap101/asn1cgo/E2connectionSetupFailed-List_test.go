// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//func createE2connectionSetupFailedListMsg() (*e2ap_pdu_contents.E2ConnectionSetupFailedList, error) {
//
//	// e2connectionSetupFailedList := pdubuilder.CreateE2connectionSetupFailedList() //ToDo - fill in arguments here(if this function exists
//
//	e2connectionSetupFailedList := e2ap_pdu_contents.E2ConnectionSetupFailedList{}
//
//	if err := e2connectionSetupFailedList.Validate(); err != nil {
//		return nil, fmt.Errorf("error validating E2connectionSetupFailedList %s", err.Error())
//	}
//	return &e2connectionSetupFailedList, nil
//}
//
//func Test_xerEncodingE2connectionSetupFailedList(t *testing.T) {
//
//	e2connectionSetupFailedList, err := createE2connectionSetupFailedListMsg()
//	assert.NilError(t, err, "Error creating E2connectionSetupFailedList PDU")
//
//	xer, err := xerEncodeE2connectionSetupFailedList(e2connectionSetupFailedList)
//	assert.NilError(t, err)
//	assert.Equal(t, 1, len(xer)) //ToDo - adjust length of the XER encoded message
//	t.Logf("E2connectionSetupFailedList XER\n%s", string(xer))
//
//	result, err := xerDecodeE2connectionSetupFailedList(xer)
//	assert.NilError(t, err)
//	assert.Assert(t, result != nil)
//	t.Logf("E2connectionSetupFailedList XER - decoded\n%v", result)
//	//ToDo - adjust field's verification
//
//	assert.Equal(t, 1, len(result.GetValue())) //ToDo - adjust length of a list
//	assert.DeepEqual(t, e2connectionSetupFailedList.GetValue(), result.GetValue())
//
//}
//
//func Test_perEncodingE2connectionSetupFailedList(t *testing.T) {
//
//	e2connectionSetupFailedList, err := createE2connectionSetupFailedListMsg()
//	assert.NilError(t, err, "Error creating E2connectionSetupFailedList PDU")
//
//	per, err := perEncodeE2connectionSetupFailedList(e2connectionSetupFailedList)
//	assert.NilError(t, err)
//	assert.Equal(t, 1, len(per)) // ToDo - adjust length of the PER encoded message
//	t.Logf("E2connectionSetupFailedList PER\n%v", hex.Dump(per))
//
//	result, err := perDecodeE2connectionSetupFailedList(per)
//	assert.NilError(t, err)
//	assert.Assert(t, result != nil)
//	t.Logf("E2connectionSetupFailedList PER - decoded\n%v", result)
//	//ToDo - adjust field's verification
//
//	assert.Equal(t, 1, len(result.GetValue())) //ToDo - adjust length of a list
//	assert.DeepEqual(t, e2connectionSetupFailedList.GetValue(), result.GetValue())
//
//}
