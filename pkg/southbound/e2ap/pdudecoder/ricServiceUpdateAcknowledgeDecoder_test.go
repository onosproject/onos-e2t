// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdudecoder

//func Test_DecodeRicServiceUpdateAcknowledgePdu(t *testing.T) {
//	rsuaXer, err := ioutil.ReadFile("../test/RICserviceUpdateAcknowledge.xml")
//	assert.NilError(t, err, "Unexpected error when loading file")
//	e2apPdu, err := asn1cgo.XerDecodeE2apPdu(rsuaXer)
//	assert.NilError(t, err)
//
//	transactionID, ranFunctionsAccepted, causes, err := DecodeRicServiceUpdateAcknowledgePdu(e2apPdu)
//	assert.NilError(t, err)
//	//assert.Assert(t, ricIdentity != nil) //Commented due to the Linters (v1.34.1) error - possible nil pointer dereference (https://staticcheck.io/docs/checks#SA5011) on lines 23, 24 & 25
//
//	assert.Equal(t, 2, len(ranFunctionsAccepted))
//	rfa100, ok := ranFunctionsAccepted[100]
//	assert.Assert(t, ok, "expected a key '100'")
//	assert.Equal(t, 2, int(rfa100))
//	rfa200, ok := ranFunctionsAccepted[200]
//	assert.Assert(t, ok, "expected a key '200'")
//	assert.Equal(t, 2, int(rfa200))
//
//	assert.Assert(t, causes != nil)
//	for id, cause := range causes {
//		switch id {
//		case 101:
//			assert.Equal(t, e2ap_ies.CauseMisc_CAUSE_MISC_HARDWARE_FAILURE, cause.GetMisc())
//		case 102:
//			assert.Equal(t, e2ap_ies.CauseProtocol_CAUSE_PROTOCOL_SEMANTIC_ERROR, cause.GetProtocol())
//		default:
//			assert.Assert(t, false, "unexpected cause %d", id)
//		}
//	}
//
//	if transactionID != nil {
//		assert.Equal(t, int32(1), *transactionID)
//	}
//}
