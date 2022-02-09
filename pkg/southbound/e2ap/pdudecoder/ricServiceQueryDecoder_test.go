// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdudecoder

//func Test_DecodeRicServiceQueryPdu(t *testing.T) {
//	rsqXer, err := ioutil.ReadFile("../test/RICserviceQuery.xml")
//	assert.NilError(t, err, "Unexpected error when loading file")
//	e2apPdu, err := asn1cgo.XerDecodeE2apPdu(rsqXer)
//	assert.NilError(t, err)
//
//	transactionID, ranFunctionsAccepted, err := DecodeRicServiceQueryPdu(e2apPdu)
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
//	if transactionID != nil {
//		assert.Equal(t, int32(54), *transactionID)
//	}
//}
