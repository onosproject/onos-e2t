// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdudecoder

//func Test_DecodeRicSubscriptionDeleteRequiredPdu(t *testing.T) {
//	ricSubscriptionDeleteRequiredXer, err := ioutil.ReadFile("../test/RICsubscriptionDeleteRequired.xml")
//	assert.NilError(t, err, "Unexpected error when loading file")
//	e2apPdu, err := asn1cgo.XerDecodeE2apPdu(ricSubscriptionDeleteRequiredXer)
//	assert.NilError(t, err)
//
//	rswcl, err := DecodeRicSubscriptionDeleteRequiredPdu(e2apPdu)
//	assert.NilError(t, err)
//	assert.Assert(t, rswcl != nil) //Commented due to the Linters (v1.34.1) error - possible nil pointer dereference (https://staticcheck.io/docs/checks#SA5011) on line 23
//
//	assert.Assert(t, rswcl != nil)
//	assert.Equal(t, 2, len(rswcl))
//
//	for id, item := range rswcl {
//		switch id {
//		case 100:
//			assert.Equal(t, e2ap_ies.CauseE2Node_CAUSE_E2NODE_E2NODE_COMPONENT_UNKNOWN.Number(), item.Cause.GetE2Node().Number())
//			assert.Equal(t, 1, int(item.RicRequestID.RequestorID))
//			assert.Equal(t, 1, int(item.RicRequestID.InstanceID))
//		case 200:
//			assert.Equal(t, e2ap_ies.CauseE2Node_CAUSE_E2NODE_E2NODE_COMPONENT_UNKNOWN.Number(), item.Cause.GetE2Node().Number())
//			assert.Equal(t, 2, int(item.RicRequestID.RequestorID))
//			assert.Equal(t, 12, int(item.RicRequestID.InstanceID))
//		default:
//			assert.Assert(t, false, "unexpected cause %d", id)
//		}
//	}
//}
