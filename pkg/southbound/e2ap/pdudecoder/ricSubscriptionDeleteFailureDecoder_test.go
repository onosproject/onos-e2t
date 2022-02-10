// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdudecoder

//func Test_DecodeRicSubscriptionDeleteFailurePdu(t *testing.T) {
//	ricSubscriptionDeleteFailureXer, err := ioutil.ReadFile("../test/RICsubscriptionDeleteFailure.xml")
//	assert.NilError(t, err, "Unexpected error when loading file")
//	e2apPdu, err := asn1cgo.XerDecodeE2apPdu(ricSubscriptionDeleteFailureXer)
//	assert.NilError(t, err)
//
//	rfID, rrID, cause, pr, crit, tm, cdrID, diags, err := DecodeRicSubscriptionDeleteFailurePdu(e2apPdu)
//	assert.NilError(t, err)
//	//assert.Assert(t, rfID != nil) //Commented due to the Linters (v1.34.1) error - possible nil pointer dereference (https://staticcheck.io/docs/checks#SA5011) on line 23
//	assert.Equal(t, 9, int(*rfID))
//
//	assert.Equal(t, cause.GetTransport(), e2ap_ies.CauseTransport_CAUSE_TRANSPORT_TRANSPORT_RESOURCE_UNAVAILABLE)
//	assert.Equal(t, int32(*pr), int32(8))
//	assert.Equal(t, int32(*crit), int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE))
//	assert.Equal(t, int32(*tm), int32(e2ap_commondatatypes.TriggeringMessage_TRIGGERING_MESSAGE_UNSUCCESSFUL_OUTCOME))
//	assert.Equal(t, int32(cdrID.InstanceID), int32(20))
//	assert.Equal(t, int32(cdrID.RequestorID), int32(10))
//	assert.Equal(t, int32(diags[0].IEId), int32(30))
//	assert.Equal(t, int32(diags[0].IECriticality), int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE))
//	assert.Equal(t, int32(diags[0].TypeOfError), int32(e2ap_ies.TypeOfError_TYPE_OF_ERROR_MISSING))
//
//	//assert.Assert(t, rrID != nil) //Commented due to the Linters (v1.34.1) error - possible nil pointer dereference (https://staticcheck.io/docs/checks#SA5011) on lines 26 & 27
//	assert.Equal(t, 22, int(rrID.RequestorID))
//	assert.Equal(t, 6, int(rrID.InstanceID))
//}
