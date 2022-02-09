// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdudecoder

//func Test_DecodeRicIndicationPdu(t *testing.T) {
//	e2setupRequestXer, err := ioutil.ReadFile("../test/RicIndication.xml")
//	assert.NilError(t, err, "Unexpected error when loading file")
//	e2apPdu, err := asn1cgo.XerDecodeE2apPdu(e2setupRequestXer)
//	assert.NilError(t, err)
//
//	ranFunctionID, ricActionID, ricCallProcessID, ricIndicationHeader, ricIndicationMessage, ricIndicationSn,
//		ricIndicationType, ricRequest, err := DecodeRicIndicationPdu(e2apPdu)
//	assert.NilError(t, err)
//	assert.Equal(t, 9, int(ranFunctionID), "unexpected ranFunctionID")
//	assert.Equal(t, 2, int(ricActionID), "unexpected ricActionID")
//	assert.DeepEqual(t, []byte{'1', '2', '3'}, []byte(*ricIndicationHeader))
//	assert.DeepEqual(t, []byte{'4', '5', '6'}, []byte(*ricIndicationMessage))
//	assert.DeepEqual(t, []byte{'7', '8', '9'}, []byte(*ricCallProcessID))
//	assert.Equal(t, 1, int(ricIndicationSn), "unexpected ricIndicationSn")
//	assert.Equal(t, e2apies.RicindicationType_RICINDICATION_TYPE_INSERT, ricIndicationType, "unexpected ricIndicationType")
//	assert.Assert(t, ricRequest != nil)
//}
//
//func Test_DecodeRicIndicationPdu2(t *testing.T) {
//	e2setupRequestXer, err := ioutil.ReadFile("../test/RicIndication2.xml")
//	assert.NilError(t, err, "Unexpected error when loading file")
//	e2apPdu, err := asn1cgo.XerDecodeE2apPdu(e2setupRequestXer)
//	assert.NilError(t, err)
//	ricInd := e2apPdu.GetInitiatingMessage().GetProcedureCode().GetRicIndication().GetInitiatingMessage()
//	assert.Assert(t, ricInd != nil)
//	assert.Equal(t, int32(0), ricInd.GetProtocolIes().GetE2ApProtocolIes29().GetValue().GetRicRequestorId())
//	assert.Equal(t, int32(0), ricInd.GetProtocolIes().GetE2ApProtocolIes29().GetValue().GetRicInstanceId())
//	t.Log(e2apPdu)
//
//	ranFunctionID, ricActionID, ricCallProcessID, ricIndicationHeader, ricIndicationMessage, ricIndicationSn,
//		ricIndicationType, ricRequest, err := DecodeRicIndicationPdu(e2apPdu)
//	assert.NilError(t, err)
//	assert.Equal(t, 1, int(ranFunctionID), "unexpected ranFunctionID")
//	assert.Equal(t, 0, int(ricActionID), "unexpected ricActionID")
//	assert.DeepEqual(t, []byte{0x3F, 0x08, 0x37, 0x34, 0x37, 0x38, 0xB5, 0xC6, 0x77,
//		0x88, 0x02, 0x37, 0x34, 0x37, 0x22, 0x5B, 0xD6, 0x00, 0x70, 0x37, 0x34, 0x37,
//		0x98, 0x80, 0x31, 0x30, 0x30, 0x09, 0x09}, []byte(*ricIndicationHeader))
//	assert.DeepEqual(t, []byte{0x40, 0x00, 0x00, 0x4C, 0x0C, 0x66, 0x6F, 0x6F, 0x2D, 0x67, 0x4E, 0x42, 0x80, 0x00, 0x00},
//		[]byte(*ricIndicationMessage))
//	assert.Assert(t, is.Nil(ricCallProcessID))
//	assert.Equal(t, 0, int(ricIndicationSn), "unexpected ricIndicationSn")
//	assert.Equal(t, e2apies.RicindicationType_RICINDICATION_TYPE_REPORT, ricIndicationType, "unexpected ricIndicationType")
//	assert.Assert(t, ricRequest != nil)
//}
