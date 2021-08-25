// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

import (
	"encoding/hex"
	"fmt"
	"testing"

	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-commondatatypes"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	e2ap_pdu_contents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/pdubuilder"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/types"
	"gotest.tools/assert"
)

func createE2connectionUpdateMsg() (*e2ap_pdu_contents.E2ConnectionUpdate, error) {

	e2connectionUpdate, err := pdubuilder.CreateE2connectionUpdateE2apPdu([]*types.E2ConnectionUpdateItem{{TnlInformation: types.TnlInformation{
		TnlPort: e2ap_commondatatypes.BitString{
			Value: []byte{0xae, 0x89},
			Len:   16,
		},
		TnlAddress: e2ap_commondatatypes.BitString{
			Value: []byte{0x89, 0xab, 0xdc, 0xdf, 0x01, 0x23, 0x45, 0x67},
			Len:   64,
		}},
		TnlUsage: e2ap_ies.Tnlusage_TNLUSAGE_BOTH}},
		[]*types.E2ConnectionUpdateItem{{TnlInformation: types.TnlInformation{
			TnlPort: e2ap_commondatatypes.BitString{
				Value: []byte{0xae, 0x87},
				Len:   16,
			},
			TnlAddress: e2ap_commondatatypes.BitString{
				Value: []byte{0x89, 0xab, 0xdc, 0xdf, 0x01, 0x23, 0x45, 0x65},
				Len:   64,
			}},
			TnlUsage: e2ap_ies.Tnlusage_TNLUSAGE_RIC_SERVICE}},
		[]*types.TnlInformation{
			{TnlPort: e2ap_commondatatypes.BitString{
				Value: []byte{0xab, 0x89},
				Len:   16,
			},
				TnlAddress: e2ap_commondatatypes.BitString{
					Value: []byte{0x89, 0xab, 0xdc, 0xdf, 0x01, 0x23, 0x45, 0x61},
					Len:   64,
				}},
			{TnlPort: e2ap_commondatatypes.BitString{
				Value: []byte{0xcd, 0x89},
				Len:   16,
			},
				TnlAddress: e2ap_commondatatypes.BitString{
					Value: []byte{0x89, 0xab, 0xdc, 0xdf, 0x01, 0x23, 0x45, 0x76},
					Len:   64,
				}},
		})
	if err != nil {
		return nil, err
	}

	if err := e2connectionUpdate.Validate(); err != nil {
		return nil, fmt.Errorf("error validating E2connectionUpdate %s", err.Error())
	}
	return e2connectionUpdate.GetInitiatingMessage().GetProcedureCode().GetE2ConnectionUpdate().GetInitiatingMessage(), nil
}

func Test_xerEncodingE2connectionUpdate(t *testing.T) {

	e2connectionUpdate, err := createE2connectionUpdateMsg()
	assert.NilError(t, err, "Error creating E2connectionUpdate PDU")

	xer, err := xerEncodeE2connectionUpdate(e2connectionUpdate)
	assert.NilError(t, err)
	assert.Equal(t, 4335, len(xer))
	t.Logf("E2connectionUpdate XER\n%s", string(xer))

	result, err := xerDecodeE2connectionUpdate(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2connectionUpdate XER - decoded\n%v", result)
	assert.Equal(t, e2connectionUpdate.GetProtocolIes().GetE2ApProtocolIes44().GetConnectionAdd().GetValue()[0].GetValue().GetTnlUsage(), result.GetProtocolIes().GetE2ApProtocolIes44().GetConnectionAdd().GetValue()[0].GetValue().GetTnlUsage())
	assert.Equal(t, e2connectionUpdate.GetProtocolIes().GetE2ApProtocolIes44().GetConnectionAdd().GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetLen(), result.GetProtocolIes().GetE2ApProtocolIes44().GetConnectionAdd().GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetLen())
	assert.DeepEqual(t, e2connectionUpdate.GetProtocolIes().GetE2ApProtocolIes44().GetConnectionAdd().GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes44().GetConnectionAdd().GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetValue())
	assert.Equal(t, e2connectionUpdate.GetProtocolIes().GetE2ApProtocolIes44().GetConnectionAdd().GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetLen(), result.GetProtocolIes().GetE2ApProtocolIes44().GetConnectionAdd().GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetLen())
	assert.DeepEqual(t, e2connectionUpdate.GetProtocolIes().GetE2ApProtocolIes44().GetConnectionAdd().GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes44().GetConnectionAdd().GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetValue())
	assert.Equal(t, e2connectionUpdate.GetProtocolIes().GetE2ApProtocolIes44().GetConnectionAdd().GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetLen(), result.GetProtocolIes().GetE2ApProtocolIes44().GetConnectionAdd().GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetLen())
	assert.DeepEqual(t, e2connectionUpdate.GetProtocolIes().GetE2ApProtocolIes44().GetConnectionAdd().GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes44().GetConnectionAdd().GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetValue())
	assert.Equal(t, e2connectionUpdate.GetProtocolIes().GetE2ApProtocolIes45().GetConnectionModify().GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetLen(), result.GetProtocolIes().GetE2ApProtocolIes45().GetConnectionModify().GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetLen())
	assert.DeepEqual(t, e2connectionUpdate.GetProtocolIes().GetE2ApProtocolIes45().GetConnectionModify().GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes45().GetConnectionModify().GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetValue())
	assert.Equal(t, e2connectionUpdate.GetProtocolIes().GetE2ApProtocolIes45().GetConnectionModify().GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetLen(), result.GetProtocolIes().GetE2ApProtocolIes45().GetConnectionModify().GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetLen())
	assert.DeepEqual(t, e2connectionUpdate.GetProtocolIes().GetE2ApProtocolIes45().GetConnectionModify().GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes45().GetConnectionModify().GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetValue())
	assert.Equal(t, e2connectionUpdate.GetProtocolIes().GetE2ApProtocolIes45().GetConnectionModify().GetValue()[0].GetValue().GetTnlUsage(), result.GetProtocolIes().GetE2ApProtocolIes45().GetConnectionModify().GetValue()[0].GetValue().GetTnlUsage())
	assert.Equal(t, e2connectionUpdate.GetProtocolIes().GetE2ApProtocolIes46().GetConnectionRemove().GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetLen(), result.GetProtocolIes().GetE2ApProtocolIes46().GetConnectionRemove().GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetLen())
	assert.DeepEqual(t, e2connectionUpdate.GetProtocolIes().GetE2ApProtocolIes46().GetConnectionRemove().GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes46().GetConnectionRemove().GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetValue())
	assert.Equal(t, e2connectionUpdate.GetProtocolIes().GetE2ApProtocolIes46().GetConnectionRemove().GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetLen(), result.GetProtocolIes().GetE2ApProtocolIes46().GetConnectionRemove().GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetLen())
	assert.DeepEqual(t, e2connectionUpdate.GetProtocolIes().GetE2ApProtocolIes46().GetConnectionRemove().GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes46().GetConnectionRemove().GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetValue())
}

func Test_perEncodingE2connectionUpdate(t *testing.T) {

	e2connectionUpdate, err := createE2connectionUpdateMsg()
	assert.NilError(t, err, "Error creating E2connectionUpdate PDU")

	per, err := perEncodeE2connectionUpdate(e2connectionUpdate)
	assert.NilError(t, err)
	assert.Equal(t, 84, len(per))
	t.Logf("E2connectionUpdate PER\n%v", hex.Dump(per))

	result, err := perDecodeE2connectionUpdate(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2connectionUpdate PER - decoded\n%v", result)
	assert.Equal(t, e2connectionUpdate.GetProtocolIes().GetE2ApProtocolIes44().GetConnectionAdd().GetValue()[0].GetValue().GetTnlUsage(), result.GetProtocolIes().GetE2ApProtocolIes44().GetConnectionAdd().GetValue()[0].GetValue().GetTnlUsage())
	assert.Equal(t, e2connectionUpdate.GetProtocolIes().GetE2ApProtocolIes44().GetConnectionAdd().GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetLen(), result.GetProtocolIes().GetE2ApProtocolIes44().GetConnectionAdd().GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetLen())
	assert.DeepEqual(t, e2connectionUpdate.GetProtocolIes().GetE2ApProtocolIes44().GetConnectionAdd().GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes44().GetConnectionAdd().GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetValue())
	assert.Equal(t, e2connectionUpdate.GetProtocolIes().GetE2ApProtocolIes44().GetConnectionAdd().GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetLen(), result.GetProtocolIes().GetE2ApProtocolIes44().GetConnectionAdd().GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetLen())
	assert.DeepEqual(t, e2connectionUpdate.GetProtocolIes().GetE2ApProtocolIes44().GetConnectionAdd().GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes44().GetConnectionAdd().GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetValue())
	assert.Equal(t, e2connectionUpdate.GetProtocolIes().GetE2ApProtocolIes44().GetConnectionAdd().GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetLen(), result.GetProtocolIes().GetE2ApProtocolIes44().GetConnectionAdd().GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetLen())
	assert.DeepEqual(t, e2connectionUpdate.GetProtocolIes().GetE2ApProtocolIes44().GetConnectionAdd().GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes44().GetConnectionAdd().GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetValue())
	assert.Equal(t, e2connectionUpdate.GetProtocolIes().GetE2ApProtocolIes45().GetConnectionModify().GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetLen(), result.GetProtocolIes().GetE2ApProtocolIes45().GetConnectionModify().GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetLen())
	assert.DeepEqual(t, e2connectionUpdate.GetProtocolIes().GetE2ApProtocolIes45().GetConnectionModify().GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes45().GetConnectionModify().GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetValue())
	assert.Equal(t, e2connectionUpdate.GetProtocolIes().GetE2ApProtocolIes45().GetConnectionModify().GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetLen(), result.GetProtocolIes().GetE2ApProtocolIes45().GetConnectionModify().GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetLen())
	assert.DeepEqual(t, e2connectionUpdate.GetProtocolIes().GetE2ApProtocolIes45().GetConnectionModify().GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes45().GetConnectionModify().GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetValue())
	assert.Equal(t, e2connectionUpdate.GetProtocolIes().GetE2ApProtocolIes45().GetConnectionModify().GetValue()[0].GetValue().GetTnlUsage(), result.GetProtocolIes().GetE2ApProtocolIes45().GetConnectionModify().GetValue()[0].GetValue().GetTnlUsage())
	assert.Equal(t, e2connectionUpdate.GetProtocolIes().GetE2ApProtocolIes46().GetConnectionRemove().GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetLen(), result.GetProtocolIes().GetE2ApProtocolIes46().GetConnectionRemove().GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetLen())
	assert.DeepEqual(t, e2connectionUpdate.GetProtocolIes().GetE2ApProtocolIes46().GetConnectionRemove().GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes46().GetConnectionRemove().GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetValue())
	assert.Equal(t, e2connectionUpdate.GetProtocolIes().GetE2ApProtocolIes46().GetConnectionRemove().GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetLen(), result.GetProtocolIes().GetE2ApProtocolIes46().GetConnectionRemove().GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetLen())
	assert.DeepEqual(t, e2connectionUpdate.GetProtocolIes().GetE2ApProtocolIes46().GetConnectionRemove().GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes46().GetConnectionRemove().GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetValue())
}
