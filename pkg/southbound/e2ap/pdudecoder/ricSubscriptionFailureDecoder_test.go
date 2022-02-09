// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdudecoder

//func Test_DecodeRicSubscriptionFailurePdu(t *testing.T) {
//	ricSubscriptionFailureXer, err := ioutil.ReadFile("../test/RICsubscriptionFailure.xml")
//	assert.NilError(t, err, "Unexpected error when loading file")
//	e2apPdu, err := asn1cgo.XerDecodeE2apPdu(ricSubscriptionFailureXer)
//	assert.NilError(t, err)
//
//	rrID, rfID, pc, crit, tm, critReq, cause, diags, err := DecodeRicSubscriptionFailurePdu(e2apPdu)
//	assert.NilError(t, err)
//	//assert.Assert(t, rfID != nil) //Commented due to the Linters (v1.34.1) error - possible nil pointer dereference (https://staticcheck.io/docs/checks#SA5011) on line 26
//	assert.Equal(t, 9, int(*rfID))
//
//	//assert.Assert(t, rrID != nil) //Commented due to the Linters (v1.34.1) error - possible nil pointer dereference (https://staticcheck.io/docs/checks#SA5011) on lines 29 & 30
//	assert.Equal(t, 22, int(rrID.RequestorID))
//	assert.Equal(t, 6, int(rrID.InstanceID))
//
//	assert.Equal(t, v2.ProcedureCodeIDRICsubscription, pc)
//	assert.Equal(t, e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE, crit)
//	assert.Equal(t, e2ap_commondatatypes.TriggeringMessage_TRIGGERING_MESSAGE_UNSUCCESSFUL_OUTCOME, tm)
//
//	assert.Equal(t, 10, int(critReq.RequestorID))
//	assert.Equal(t, 20, int(critReq.InstanceID))
//
//	//ToDo - adjust Cause verification
//	assert.Equal(t, e2apies.CauseMisc_CAUSE_MISC_CONTROL_PROCESSING_OVERLOAD, cause.GetMisc())
//	assert.Assert(t, diags != nil)
//}
