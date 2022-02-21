// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdudecoder

import (
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/pdubuilder"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"gotest.tools/assert"
	"testing"
)

func Test_DecodeE2connectionUpdatePdu(t *testing.T) {
	e2apPdu, err := pdubuilder.CreateE2connectionUpdateE2apPdu(1)
	assert.NilError(t, err)
	assert.Assert(t, e2apPdu != nil)

	e2apPdu.GetInitiatingMessage().GetValue().GetE2ConnectionUpdate().
		SetE2ConnectionUpdateAdd([]*types.E2ConnectionUpdateItem{{TnlInformation: types.TnlInformation{
			TnlPort: &asn1.BitString{
				Value: []byte{0xae, 0x89},
				Len:   16,
			},
			TnlAddress: asn1.BitString{
				Value: []byte{0x89, 0xab, 0xdc, 0xdf, 0x01, 0x23, 0x45, 0x67},
				Len:   64,
			}},
			TnlUsage: e2ap_ies.Tnlusage_TNLUSAGE_BOTH}}).SetE2ConnectionUpdateModify([]*types.E2ConnectionUpdateItem{{TnlInformation: types.TnlInformation{
		TnlPort: &asn1.BitString{
			Value: []byte{0xba, 0x91},
			Len:   16,
		},
		TnlAddress: asn1.BitString{
			Value: []byte{0x89, 0xab, 0xdc, 0xdf, 0x01, 0x23, 0x45, 0x62},
			Len:   64,
		}},
		TnlUsage: e2ap_ies.Tnlusage_TNLUSAGE_RIC_SERVICE}}).SetE2ConnectionUpdateRemove([]*types.TnlInformation{
		{TnlPort: &asn1.BitString{
			Value: []byte{0xba, 0x98},
			Len:   16,
		},
			TnlAddress: asn1.BitString{
				Value: []byte{0x89, 0xab, 0xdc, 0xdf, 0x01, 0x23, 0x45, 0x76},
				Len:   64,
			}},
		{TnlPort: &asn1.BitString{
			Value: []byte{0xdc, 0x98},
			Len:   16,
		},
			TnlAddress: asn1.BitString{
				Value: []byte{0x89, 0xab, 0xdc, 0xdf, 0x01, 0x23, 0x45, 0x78},
				Len:   64,
			}},
	})

	transactionID, connSetup, connModify, connRemove, err := DecodeE2connectionUpdatePdu(e2apPdu)
	assert.NilError(t, err)

	assert.Equal(t, int32(connSetup[0].TnlInformation.TnlAddress.GetLen()), int32(64))
	assert.DeepEqual(t, connSetup[0].TnlInformation.TnlAddress.GetValue(), []byte{0x89, 0xab, 0xdc, 0xdf, 0x01, 0x23, 0x45, 0x67})
	assert.Equal(t, int32(connSetup[0].TnlInformation.TnlPort.GetLen()), int32(16))
	assert.DeepEqual(t, connSetup[0].TnlInformation.TnlPort.GetValue(), []byte{0xae, 0x89})
	assert.Equal(t, int32(connSetup[0].TnlUsage), int32(e2ap_ies.Tnlusage_TNLUSAGE_BOTH))
	assert.Equal(t, int32(connModify[0].TnlInformation.TnlAddress.GetLen()), int32(64))
	assert.DeepEqual(t, connModify[0].TnlInformation.TnlAddress.GetValue(), []byte{0x89, 0xab, 0xdc, 0xdf, 0x01, 0x23, 0x45, 0x62})
	assert.Equal(t, int32(connModify[0].TnlInformation.TnlPort.GetLen()), int32(16))
	assert.DeepEqual(t, connModify[0].TnlInformation.TnlPort.GetValue(), []byte{0xba, 0x91})
	assert.Equal(t, int32(connModify[0].TnlUsage), int32(e2ap_ies.Tnlusage_TNLUSAGE_RIC_SERVICE))
	assert.Equal(t, int32(connRemove[0].TnlAddress.GetLen()), int32(64))
	assert.DeepEqual(t, connRemove[0].TnlAddress.GetValue(), []byte{0x89, 0xab, 0xdc, 0xdf, 0x01, 0x23, 0x45, 0x76})
	assert.Equal(t, int32(connRemove[0].TnlPort.GetLen()), int32(16))
	assert.DeepEqual(t, connRemove[0].TnlPort.GetValue(), []byte{0xba, 0x98})
	assert.Equal(t, int32(connRemove[1].TnlAddress.GetLen()), int32(64))
	assert.DeepEqual(t, connRemove[1].TnlAddress.GetValue(), []byte{0x89, 0xab, 0xdc, 0xdf, 0x01, 0x23, 0x45, 0x78})
	assert.Equal(t, int32(connRemove[1].TnlPort.GetLen()), int32(16))
	assert.DeepEqual(t, connRemove[1].TnlPort.GetValue(), []byte{0xdc, 0x98})
	if transactionID != nil {
		assert.Equal(t, int32(1), *transactionID)
	}
}
