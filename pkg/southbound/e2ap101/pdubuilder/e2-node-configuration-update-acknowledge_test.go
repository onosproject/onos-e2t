// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
package pdubuilder

import (
	"encoding/hex"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/asn1cgo"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/types"
	"gotest.tools/assert"
	"testing"
)

func TestE2NodeConfigurationUpdateAck(t *testing.T) {

	e2ncID1 := CreateE2NodeComponentIDGnbCuUp(21)
	e2ncID2 := CreateE2NodeComponentIDGnbDu(13)

	newE2apPdu, err := CreateE2NodeConfigurationUpdateAcknowledgeE2apPdu([]*types.E2NodeComponentConfigUpdateAckItem{
		{E2NodeComponentType: e2ap_ies.E2NodeComponentType_E2NODE_COMPONENT_TYPE_G_NB,
			E2NodeComponentID: e2ncID1,
			E2NodeComponentConfigUpdateAck: types.E2NodeComponentConfigUpdateAck{
				UpdateOutcome: 1,
				FailureCause: e2ap_ies.Cause{
					Cause: &e2ap_ies.Cause_Protocol{
						Protocol: e2ap_ies.CauseProtocol_CAUSE_PROTOCOL_TRANSFER_SYNTAX_ERROR,
					},
				},
			}},
		{E2NodeComponentType: e2ap_ies.E2NodeComponentType_E2NODE_COMPONENT_TYPE_E_NB,
			E2NodeComponentID: e2ncID2,
			E2NodeComponentConfigUpdateAck: types.E2NodeComponentConfigUpdateAck{
				UpdateOutcome: 1,
				FailureCause: e2ap_ies.Cause{
					Cause: &e2ap_ies.Cause_Protocol{
						Protocol: e2ap_ies.CauseProtocol_CAUSE_PROTOCOL_ABSTRACT_SYNTAX_ERROR_FALSELY_CONSTRUCTED_MESSAGE,
					},
				},
			}}})
	assert.NilError(t, err)
	assert.Assert(t, newE2apPdu != nil)

	xer, err := asn1cgo.XerEncodeE2apPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("E2NodeConfigurationUpdateAck E2AP PDU XER\n%s", string(xer))

	per, err := asn1cgo.PerEncodeE2apPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("E2NodeConfigurationUpdateAck E2AP PDU PER\n%v", hex.Dump(per))
}
