// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdudecoder

//func Test_DecodeResetResponsePdu(t *testing.T) {
//	rrXer, err := ioutil.ReadFile("../test/ResetResponse.xml")
//	assert.NilError(t, err, "Unexpected error when loading file")
//	e2apPdu, err := asn1cgo.XerDecodeE2apPdu(rrXer)
//	assert.NilError(t, err)
//
//	transactionID, pr, crit, tm, cdrID, diags, err := DecodeResetResponsePdu(e2apPdu)
//	assert.NilError(t, err)
//	//assert.Assert(t, ricIdentity != nil) //Commented due to the Linters (v1.34.1) error - possible nil pointer dereference (https://staticcheck.io/docs/checks#SA5011) on lines 23, 24 & 25
//
//	assert.Equal(t, int32(*pr), int32(3))
//	assert.Equal(t, int32(*crit), int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE))
//	assert.Equal(t, int32(*tm), int32(e2ap_commondatatypes.TriggeringMessage_TRIGGERING_MESSAGE_UNSUCCESSFUL_OUTCOME))
//	assert.Equal(t, int32(cdrID.InstanceID), int32(20))
//	assert.Equal(t, int32(cdrID.RequestorID), int32(10))
//	assert.Equal(t, int32(diags[0].IEId), int32(30))
//	assert.Equal(t, int32(diags[0].IECriticality), int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE))
//	assert.Equal(t, int32(diags[0].TypeOfError), int32(e2apies.TypeOfError_TYPE_OF_ERROR_MISSING))
//	if transactionID != nil {
//		assert.Equal(t, int32(1), *transactionID)
//	}
//}
