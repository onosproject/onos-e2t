// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
package pdubuilder

import (
	"encoding/hex"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-commondatatypes"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/asn1cgo"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/types"
	"gotest.tools/assert"
	"testing"
)

func TestE2connectionUpdateAcknowledge(t *testing.T) {
	newE2apPdu, err := CreateE2connectionUpdateAcknowledgeE2apPdu([]*types.E2ConnectionUpdateItem{{TnlInformation: types.TnlInformation{
		TnlPort: e2ap_commondatatypes.BitString{
			Value: []byte{0xae, 0x89},
			Len:   16,
		},
		TnlAddress: e2ap_commondatatypes.BitString{
			Value: []byte{0x89, 0xab, 0xdc, 0xdf, 0x01, 0x23, 0x45, 0x67},
			Len:   64,
		}},
		TnlUsage: e2ap_ies.Tnlusage_TNLUSAGE_BOTH}},
		[]*types.E2ConnectionSetupFailedItem{{TnlInformation: types.TnlInformation{
			TnlPort: e2ap_commondatatypes.BitString{
				Value: []byte{0xae, 0x89},
				Len:   16,
			},
			TnlAddress: e2ap_commondatatypes.BitString{
				Value: []byte{0x89, 0xab, 0xdc, 0xdf, 0x01, 0x23, 0x45, 0x67},
				Len:   64,
			}},
			Cause: e2ap_ies.Cause{
				Cause: &e2ap_ies.Cause_Protocol{
					Protocol: e2ap_ies.CauseProtocol_CAUSE_PROTOCOL_SEMANTIC_ERROR,
				}}}})
	assert.NilError(t, err)
	assert.Assert(t, newE2apPdu != nil)

	xer, err := asn1cgo.XerEncodeE2apPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("E2connectionUpdateAcknowledge E2AP PDU XER\n%s", string(xer))

	result, err := asn1cgo.XerDecodeE2apPdu(xer)
	assert.NilError(t, err)
	t.Logf("E2connectionUpdateAcknowledge E2AP PDU XER - decoded\n%v", result)
	assert.DeepEqual(t, newE2apPdu, result)

	per, err := asn1cgo.PerEncodeE2apPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("E2connectionUpdateAcknowledge E2AP PDU PER\n%v", hex.Dump(per))

	result1, err := asn1cgo.PerDecodeE2apPdu(per)
	assert.NilError(t, err)
	t.Logf("E2connectionUpdateAcknowledge E2AP PDU PER - decoded\n%v", result1)
	assert.DeepEqual(t, newE2apPdu, result1)
}

func TestE2connectionUpdateAcknowledgeExcludeOptionalIE(t *testing.T) {
	newE2apPdu, err := CreateE2connectionUpdateAcknowledgeE2apPdu(nil,
		[]*types.E2ConnectionSetupFailedItem{{TnlInformation: types.TnlInformation{
			TnlPort: e2ap_commondatatypes.BitString{
				Value: []byte{0xae, 0x89},
				Len:   16,
			},
			TnlAddress: e2ap_commondatatypes.BitString{
				Value: []byte{0x89, 0xab, 0xdc, 0xdf, 0x01, 0x23, 0x45, 0x67},
				Len:   64,
			}},
			Cause: e2ap_ies.Cause{
				Cause: &e2ap_ies.Cause_Protocol{
					Protocol: e2ap_ies.CauseProtocol_CAUSE_PROTOCOL_SEMANTIC_ERROR,
				}}}})
	assert.NilError(t, err)
	assert.Assert(t, newE2apPdu != nil)

	xer, err := asn1cgo.XerEncodeE2apPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("E2connectionUpdateAcknowledge E2AP PDU XER\n%s", string(xer))

	result, err := asn1cgo.XerDecodeE2apPdu(xer)
	assert.NilError(t, err)
	t.Logf("E2connectionUpdateAcknowledge E2AP PDU XER - decoded\n%v", result)
	assert.DeepEqual(t, newE2apPdu, result)

	per, err := asn1cgo.PerEncodeE2apPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("E2connectionUpdateAcknowledge E2AP PDU PER\n%v", hex.Dump(per))

	result1, err := asn1cgo.PerDecodeE2apPdu(per)
	assert.NilError(t, err)
	t.Logf("E2connectionUpdateAcknowledge E2AP PDU PER - decoded\n%v", result1)
	assert.DeepEqual(t, newE2apPdu, result1)
}
