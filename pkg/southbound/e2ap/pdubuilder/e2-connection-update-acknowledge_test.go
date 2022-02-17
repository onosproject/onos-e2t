// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
package pdubuilder

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/encoder"
	"testing"

	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"gotest.tools/assert"
)

func TestE2connectionUpdateAcknowledge(t *testing.T) {
	newE2apPdu, err := CreateE2connectionUpdateAcknowledgeE2apPdu(1)
	assert.NilError(t, err)
	assert.Assert(t, newE2apPdu != nil)

	newE2apPdu.GetSuccessfulOutcome().GetValue().GetE2ConnectionUpdate().
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
			Value: []byte{0xae, 0x89},
			Len:   16,
		},
		TnlAddress: asn1.BitString{
			Value: []byte{0x89, 0xab, 0xdc, 0xdf, 0x01, 0x23, 0x45, 0x67},
			Len:   64,
		}},
		Cause: e2ap_ies.Cause{
			Cause: &e2ap_ies.Cause_Protocol{
				Protocol: e2ap_ies.CauseProtocol_CAUSE_PROTOCOL_SEMANTIC_ERROR,
			}}}})

	perNew, err := encoder.PerEncodeE2ApPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("E2connectionUpdateAcknowledge E2AP PDU PER with Go APER library\n%v", hex.Dump(perNew))

	//Comparing reference PER bytes with Go APER library produced
	//assert.DeepEqual(t, per, perNew)

	result, err := encoder.PerDecodeE2ApPdu(perNew)
	assert.NilError(t, err)
	assert.DeepEqual(t, newE2apPdu.String(), result.String())
}

func TestE2connectionUpdateAcknowledgeExcludeOptionalIE(t *testing.T) {
	newE2apPdu, err := CreateE2connectionUpdateAcknowledgeE2apPdu(1)
	assert.NilError(t, err)
	assert.Assert(t, newE2apPdu != nil)
	newE2apPdu.GetSuccessfulOutcome().GetValue().GetE2ConnectionUpdate().
		SetE2ConnectionSetupFailed([]*types.E2ConnectionSetupFailedItem{{TnlInformation: types.TnlInformation{
			TnlPort: &asn1.BitString{
				Value: []byte{0xae, 0x89},
				Len:   16,
			},
			TnlAddress: asn1.BitString{
				Value: []byte{0x89, 0xab, 0xdc, 0xdf, 0x01, 0x23, 0x45, 0x67},
				Len:   64,
			}},
			Cause: e2ap_ies.Cause{
				Cause: &e2ap_ies.Cause_Protocol{
					Protocol: e2ap_ies.CauseProtocol_CAUSE_PROTOCOL_SEMANTIC_ERROR,
				}}}})

	perNew, err := encoder.PerEncodeE2ApPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("E2connectionUpdateAcknowledge E2AP PDU PER with Go APER library\n%v", hex.Dump(perNew))

	//Comparing reference PER bytes with Go APER library produced
	//assert.DeepEqual(t, per, perNew)

	result, err := encoder.PerDecodeE2ApPdu(perNew)
	assert.NilError(t, err)
	assert.DeepEqual(t, newE2apPdu.String(), result.String())
}
