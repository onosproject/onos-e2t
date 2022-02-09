// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdudecoder

//func Test_DecodeRicSubscriptionDeleteResponsePdu(t *testing.T) {
//	ricSubscriptionDeleteResponseXer, err := ioutil.ReadFile("../test/RICsubscriptionDeleteResponse.xml")
//	assert.NilError(t, err, "Unexpected error when loading file")
//	e2apPdu, err := asn1cgo.XerDecodeE2apPdu(ricSubscriptionDeleteResponseXer)
//	assert.NilError(t, err)
//
//	rfID, rrID, err := DecodeRicSubscriptionDeleteResponsePdu(e2apPdu)
//	assert.NilError(t, err)
//	//assert.Assert(t, rfID != nil) //Commented due to the Linters (v1.34.1) error - possible nil pointer dereference (https://staticcheck.io/docs/checks#SA5011) on line 23
//	assert.Equal(t, 9, int(*rfID))
//
//	//assert.Assert(t, rrID != nil) //Commented due to the Linters (v1.34.1) error - possible nil pointer dereference (https://staticcheck.io/docs/checks#SA5011) on lines 26 & 27
//	assert.Equal(t, 22, int(rrID.RequestorID))
//	assert.Equal(t, 6, int(rrID.InstanceID))
//}
