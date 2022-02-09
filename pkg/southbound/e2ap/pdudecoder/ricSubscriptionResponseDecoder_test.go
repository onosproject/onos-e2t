// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdudecoder

//func Test_DecodeRicSubscriptionResponsePdu(t *testing.T) {
//	ricSubscriptionResponseXer, err := ioutil.ReadFile("../test/RICsubscriptionResponse.xml")
//	assert.NilError(t, err, "Unexpected error when loading file")
//	e2apPdu, err := asn1cgo.XerDecodeE2apPdu(ricSubscriptionResponseXer)
//	assert.NilError(t, err)
//
//	rfID, rrID, ricActionIDs, causes, err := DecodeRicSubscriptionResponsePdu(e2apPdu)
//	assert.NilError(t, err)
//	//assert.Assert(t, rfID != nil) //Commented due to the Linters (v1.34.1) error - possible nil pointer dereference (https://staticcheck.io/docs/checks#SA5011) on lines 23
//	assert.Equal(t, 9, int(*rfID))
//
//	//ToDo - adjust verification of RicActionsNotAdmittedList
//	assert.Assert(t, causes != nil)
//	for id, cause := range causes {
//		switch id {
//		case 100:
//			assert.Equal(t, e2ap_ies.CauseTransport_CAUSE_TRANSPORT_TRANSPORT_RESOURCE_UNAVAILABLE, cause.GetTransport())
//		case 200:
//			assert.Equal(t, e2ap_ies.CauseMisc_CAUSE_MISC_HARDWARE_FAILURE, cause.GetMisc())
//		default:
//			assert.Assert(t, false, "unexpected cause %d", id)
//		}
//	}
//
//	//assert.Assert(t, rrID != nil) //Commented due to the Linters (v1.34.1) error - possible nil pointer dereference (https://staticcheck.io/docs/checks#SA5011) on lines 26 & 27
//	assert.Equal(t, 22, int(rrID.RequestorID))
//	assert.Equal(t, 6, int(rrID.InstanceID))
//
//	assert.Equal(t, 2, len(ricActionIDs))
//	assert.Equal(t, 10, int(ricActionIDs[0]))
//	assert.Equal(t, 20, int(ricActionIDs[1]))
//}
