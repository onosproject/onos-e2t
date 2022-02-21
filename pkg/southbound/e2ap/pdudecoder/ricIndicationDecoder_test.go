// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdudecoder

import (
	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/pdubuilder"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
	"gotest.tools/assert"
	"testing"
)

func Test_DecodeRicIndicationPdu(t *testing.T) {
	ricRequestID := types.RicRequest{
		RequestorID: 21,
		InstanceID:  22,
	}
	var ranFuncID types.RanFunctionID = 9
	var ricAction int32 = 2
	var ricIndicationType = e2apies.RicindicationType_RICINDICATION_TYPE_INSERT
	var ricSn types.RicIndicationSn = 1
	var ricIndHd types.RicIndicationHeader = []byte("123")
	var ricIndMsg types.RicIndicationMessage = []byte("456")
	var ricCallPrID types.RicCallProcessID = []byte("789")
	e2apPdu, err := pdubuilder.RicIndicationE2apPdu(ricRequestID,
		ranFuncID, ricAction, ricIndicationType, ricIndHd, ricIndMsg)
	assert.NilError(t, err)
	assert.Assert(t, e2apPdu != nil)
	e2apPdu.GetInitiatingMessage().GetValue().GetRicIndication().
		SetRicCallProcessID(ricCallPrID).SetRicIndicationSN(ricSn)

	ranFunctionID, ricActionID, ricCallProcessID, ricIndicationHeader, ricIndicationMessage, ricIndicationSn,
		ricIndicationType, ricRequest, err := DecodeRicIndicationPdu(e2apPdu)
	assert.NilError(t, err)
	assert.Equal(t, 9, int(ranFunctionID), "unexpected ranFunctionID")
	assert.Equal(t, 2, int(ricActionID), "unexpected ricActionID")
	assert.DeepEqual(t, []byte{'1', '2', '3'}, []byte(*ricIndicationHeader))
	assert.DeepEqual(t, []byte{'4', '5', '6'}, []byte(*ricIndicationMessage))
	assert.DeepEqual(t, []byte{'7', '8', '9'}, []byte(*ricCallProcessID))
	assert.Equal(t, 1, int(ricIndicationSn), "unexpected ricIndicationSn")
	assert.Equal(t, e2apies.RicindicationType_RICINDICATION_TYPE_INSERT, ricIndicationType, "unexpected ricIndicationType")
	assert.Assert(t, ricRequest != nil)
}

func Test_DecodeRicIndicationPdu2(t *testing.T) {
	ricRequestID := types.RicRequest{
		RequestorID: 21,
		InstanceID:  22,
	}
	var ranFuncID types.RanFunctionID = 1
	var ricAction int32 = 10
	var ricIndicationType = e2apies.RicindicationType_RICINDICATION_TYPE_REPORT
	var ricIndHd types.RicIndicationHeader = []byte("123")
	var ricIndMsg types.RicIndicationMessage = []byte("456")
	e2apPdu, err := pdubuilder.RicIndicationE2apPdu(ricRequestID,
		ranFuncID, ricAction, ricIndicationType, ricIndHd, ricIndMsg)
	assert.NilError(t, err)
	assert.Assert(t, e2apPdu != nil)
	e2apPdu.GetInitiatingMessage().GetValue().GetRicIndication().SetRicIndicationSN(0).
		SetRicIndicationHeader([]byte{0x3F, 0x08, 0x37, 0x34, 0x37, 0x38, 0xB5, 0xC6, 0x77,
			0x88, 0x02, 0x37, 0x34, 0x37, 0x22, 0x5B, 0xD6, 0x00, 0x70, 0x37, 0x34, 0x37,
			0x98, 0x80, 0x31, 0x30, 0x30, 0x09, 0x09}).
		SetRicIndicationMessage([]byte{0x40, 0x00, 0x00, 0x4C, 0x0C, 0x66, 0x6F, 0x6F, 0x2D, 0x67, 0x4E, 0x42, 0x80, 0x00, 0x00})

	ranFunctionID, ricActionID, ricCallProcessID, ricIndicationHeader, ricIndicationMessage, ricIndicationSn,
		ricIndicationType, ricRequest, err := DecodeRicIndicationPdu(e2apPdu)
	assert.NilError(t, err)
	assert.Equal(t, 1, int(ranFunctionID), "unexpected ranFunctionID")
	assert.Equal(t, 10, int(ricActionID), "unexpected ricActionID")
	assert.DeepEqual(t, []byte{0x3F, 0x08, 0x37, 0x34, 0x37, 0x38, 0xB5, 0xC6, 0x77,
		0x88, 0x02, 0x37, 0x34, 0x37, 0x22, 0x5B, 0xD6, 0x00, 0x70, 0x37, 0x34, 0x37,
		0x98, 0x80, 0x31, 0x30, 0x30, 0x09, 0x09}, []byte(*ricIndicationHeader))
	assert.DeepEqual(t, []byte{0x40, 0x00, 0x00, 0x4C, 0x0C, 0x66, 0x6F, 0x6F, 0x2D, 0x67, 0x4E, 0x42, 0x80, 0x00, 0x00},
		[]byte(*ricIndicationMessage))
	assert.Assert(t, ricCallProcessID == nil)
	assert.Equal(t, 0, int(ricIndicationSn), "unexpected ricIndicationSn")
	assert.Equal(t, e2apies.RicindicationType_RICINDICATION_TYPE_REPORT, ricIndicationType, "unexpected ricIndicationType")
	assert.Assert(t, ricRequest != nil)
}
