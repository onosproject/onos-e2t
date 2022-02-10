// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdudecoder

//func Test_DecodeErrorIndicationPdu(t *testing.T) {
//	rrXer, err := ioutil.ReadFile("../test/ErrorIndication.xml")
//	assert.NilError(t, err, "Unexpected error when loading file")
//	e2apPdu, err := asn1cgo.XerDecodeE2apPdu(rrXer)
//	assert.NilError(t, err)
//
//	transactionID, cause, rfID, rrID, pr, crit, tm, cdrID, diags, err := DecodeErrorIndicationPdu(e2apPdu)
//	assert.NilError(t, err)
//	if rfID != nil {
//		assert.Equal(t, 9, int(*rfID))
//	}
//	assert.Equal(t, e2apies.CauseMisc_CAUSE_MISC_UNSPECIFIED, cause.GetMisc())
//	if transactionID != nil {
//		assert.Equal(t, int32(21), *transactionID)
//	}
//
//	assert.Equal(t, int32(*pr), int32(8))
//	assert.Equal(t, int32(*crit), int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE))
//	assert.Equal(t, int32(*tm), int32(e2ap_commondatatypes.TriggeringMessage_TRIGGERING_MESSAGE_UNSUCCESSFUL_OUTCOME))
//	assert.Equal(t, int32(cdrID.InstanceID), int32(20))
//	assert.Equal(t, int32(cdrID.RequestorID), int32(10))
//	assert.Equal(t, int32(diags[0].IEId), int32(30))
//	assert.Equal(t, int32(diags[0].IECriticality), int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE))
//	assert.Equal(t, int32(diags[0].TypeOfError), int32(e2apies.TypeOfError_TYPE_OF_ERROR_MISSING))
//
//	assert.Equal(t, 22, int(rrID.RequestorID))
//	assert.Equal(t, 6, int(rrID.InstanceID))
//}
