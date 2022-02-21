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

func Test_DecodeE2connectionUpdateAcknowledgePdu(t *testing.T) {
	e2apPdu, err := pdubuilder.CreateE2connectionUpdateAcknowledgeE2apPdu(1)
	assert.NilError(t, err)
	assert.Assert(t, e2apPdu != nil)

	e2apPdu.GetSuccessfulOutcome().GetValue().GetE2ConnectionUpdate().
		SetE2ConnectionSetup([]*types.E2ConnectionUpdateItem{{TnlInformation: types.TnlInformation{
			TnlPort: &asn1.BitString{
				Value: []byte{0xae, 0x89},
				Len:   16,
			},
			TnlAddress: asn1.BitString{
				Value: []byte{0x89, 0xab, 0xdc, 0xdf, 0x01, 0x23, 0x45, 0x67},
				Len:   64,
			}},
			TnlUsage: e2ap_ies.Tnlusage_TNLUSAGE_BOTH}}).SetE2ConnectionSetupFailed([]*types.E2ConnectionSetupFailedItem{{TnlInformation: types.TnlInformation{
		TnlPort: &asn1.BitString{
			Value: []byte{0xae, 0x8A},
			Len:   16,
		},
		TnlAddress: asn1.BitString{
			Value: []byte{0x89, 0xab, 0xdc, 0xdf, 0x01, 0x23, 0x45, 0x68},
			Len:   64,
		}},
		Cause: e2ap_ies.Cause{
			Cause: &e2ap_ies.Cause_Protocol{
				Protocol: e2ap_ies.CauseProtocol_CAUSE_PROTOCOL_SEMANTIC_ERROR,
			}}}})

	transactionID, connSetup, connSetupFailed, err := DecodeE2connectionUpdateAcknowledgePdu(e2apPdu)
	assert.NilError(t, err)

	assert.Equal(t, int32(connSetup[0].TnlInformation.TnlAddress.GetLen()), int32(64))
	assert.DeepEqual(t, connSetup[0].TnlInformation.TnlAddress.GetValue(), []byte{0x89, 0xab, 0xdc, 0xdf, 0x01, 0x23, 0x45, 0x67})
	assert.Equal(t, int32(connSetup[0].TnlInformation.TnlPort.GetLen()), int32(16))
	assert.DeepEqual(t, connSetup[0].TnlInformation.TnlPort.GetValue(), []byte{0xae, 0x89})
	assert.Equal(t, int32(connSetup[0].TnlUsage), int32(e2ap_ies.Tnlusage_TNLUSAGE_BOTH))
	assert.Equal(t, int32(connSetupFailed[0].TnlInformation.TnlAddress.GetLen()), int32(64))
	assert.DeepEqual(t, connSetupFailed[0].TnlInformation.TnlAddress.GetValue(), []byte{0x89, 0xab, 0xdc, 0xdf, 0x01, 0x23, 0x45, 0x68})
	assert.Equal(t, int32(connSetupFailed[0].TnlInformation.TnlPort.GetLen()), int32(16))
	assert.DeepEqual(t, connSetupFailed[0].TnlInformation.TnlPort.GetValue(), []byte{0xae, 0x8A})
	assert.Equal(t, int32(connSetupFailed[0].Cause.GetProtocol()), int32(e2ap_ies.CauseProtocol_CAUSE_PROTOCOL_SEMANTIC_ERROR))
	if transactionID != nil {
		assert.Equal(t, int32(1), *transactionID)
	}
}
