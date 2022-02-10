// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdudecoder

//func Test_DecodeE2connectionUpdatePdu(t *testing.T) {
//	e2ncuXer, err := ioutil.ReadFile("../test/E2connectionUpdate.xml")
//	assert.NilError(t, err, "Unexpected error when loading file")
//	e2apPdu, err := asn1cgo.XerDecodeE2apPdu(e2ncuXer)
//	assert.NilError(t, err)
//
//	transactionID, connSetup, connModify, connRemove, err := DecodeE2connectionUpdatePdu(e2apPdu)
//	assert.NilError(t, err)
//	//assert.Assert(t, ricIdentity != nil) //Commented due to the Linters (v1.34.1) error - possible nil pointer dereference (https://staticcheck.io/docs/checks#SA5011) on lines 23, 24 & 25
//
//	assert.Equal(t, int32(connSetup[0].TnlInformation.TnlAddress.GetLen()), int32(64))
//	assert.DeepEqual(t, connSetup[0].TnlInformation.TnlAddress.GetValue(), []byte{0x89, 0xab, 0xdc, 0xdf, 0x01, 0x23, 0x45, 0x67})
//	assert.Equal(t, int32(connSetup[0].TnlInformation.TnlPort.GetLen()), int32(16))
//	assert.DeepEqual(t, connSetup[0].TnlInformation.TnlPort.GetValue(), []byte{0xae, 0x89})
//	assert.Equal(t, int32(connSetup[0].TnlUsage), int32(e2ap_ies.Tnlusage_TNLUSAGE_BOTH))
//	assert.Equal(t, int32(connModify[0].TnlInformation.TnlAddress.GetLen()), int32(64))
//	assert.DeepEqual(t, connModify[0].TnlInformation.TnlAddress.GetValue(), []byte{0x89, 0xab, 0xdc, 0xdf, 0x01, 0x23, 0x45, 0x62})
//	assert.Equal(t, int32(connModify[0].TnlInformation.TnlPort.GetLen()), int32(16))
//	assert.DeepEqual(t, connModify[0].TnlInformation.TnlPort.GetValue(), []byte{0xba, 0x91})
//	assert.Equal(t, int32(connModify[0].TnlUsage), int32(e2ap_ies.Tnlusage_TNLUSAGE_RIC_SERVICE))
//	assert.Equal(t, int32(connRemove[0].TnlAddress.GetLen()), int32(64))
//	assert.DeepEqual(t, connRemove[0].TnlAddress.GetValue(), []byte{0x89, 0xab, 0xdc, 0xdf, 0x01, 0x23, 0x45, 0x76})
//	assert.Equal(t, int32(connRemove[0].TnlPort.GetLen()), int32(16))
//	assert.DeepEqual(t, connRemove[0].TnlPort.GetValue(), []byte{0xba, 0x98})
//	assert.Equal(t, int32(connRemove[1].TnlAddress.GetLen()), int32(64))
//	assert.DeepEqual(t, connRemove[1].TnlAddress.GetValue(), []byte{0x89, 0xab, 0xdc, 0xdf, 0x01, 0x23, 0x45, 0x78})
//	assert.Equal(t, int32(connRemove[1].TnlPort.GetLen()), int32(16))
//	assert.DeepEqual(t, connRemove[1].TnlPort.GetValue(), []byte{0xdc, 0x98})
//	if transactionID != nil {
//		assert.Equal(t, int32(1), *transactionID)
//	}
//}
