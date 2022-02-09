// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdudecoder

//func Test_DecodeRicServiceUpdatePdu(t *testing.T) {
//	rsuXer, err := ioutil.ReadFile("../test/RICserviceUpdate.xml")
//	assert.NilError(t, err, "Unexpected error when loading file")
//	e2apPdu, err := asn1cgo.XerDecodeE2apPdu(rsuXer)
//	assert.NilError(t, err)
//
//	transactionID, rfal, rfdl, rfml, err := DecodeRicServiceUpdatePdu(e2apPdu)
//	assert.NilError(t, err)
//	//assert.Assert(t, ricIdentity != nil) //Commented due to the Linters (v1.34.1) error - possible nil pointer dereference (https://staticcheck.io/docs/checks#SA5011) on lines 23, 24 & 25
//
//	assert.DeepEqual(t, []byte(rfal[100].Description), []byte("Type 1"))
//	assert.DeepEqual(t, []byte(rfal[100].OID), []byte("oid1"))
//	assert.Equal(t, int(rfal[100].Revision), 1)
//	assert.DeepEqual(t, []byte(rfal[200].Description), []byte("Type 2"))
//	assert.DeepEqual(t, []byte(rfal[200].OID), []byte("oid2"))
//	assert.Equal(t, int(rfal[200].Revision), 2)
//	rfd100, ok := rfdl[100]
//	assert.Assert(t, ok, "expected a key '100'")
//	assert.Equal(t, 2, int(rfd100))
//	rfd200, ok := rfdl[200]
//	assert.Assert(t, ok, "expected a key '200'")
//	assert.Equal(t, 2, int(rfd200))
//	assert.DeepEqual(t, []byte(rfml[100].Description), []byte("Type 3"))
//	assert.DeepEqual(t, []byte(rfml[100].OID), []byte("oid3"))
//	assert.Equal(t, int(rfml[100].Revision), 3)
//	assert.DeepEqual(t, []byte(rfml[200].Description), []byte("Type 4"))
//	assert.DeepEqual(t, []byte(rfml[200].OID), []byte("oid4"))
//	assert.Equal(t, int(rfml[200].Revision), 4)
//
//	if transactionID != nil {
//		assert.Equal(t, int32(1), *transactionID)
//	}
//}
