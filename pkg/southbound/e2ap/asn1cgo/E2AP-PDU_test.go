// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

import (
	"encoding/hex"
	"testing"

	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-ies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap2/pdubuilder"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap2/types"
	"gotest.tools/assert"
)

func Test_newE2setupResponseE2APpdu(t *testing.T) {
	rfAccepted := make(types.RanFunctionRevisions)
	rfAccepted[100] = 2
	rfAccepted[200] = 2

	plmnID := [3]byte{0x79, 0x78, 0x70}
	ricID := types.RicIdentifier{
		RicIdentifierValue: []byte{0xDE, 0xBC, 0xA0},
		RicIdentifierLen:   20,
	}

	response, err := pdubuilder.NewE2SetupResponse(1, plmnID, ricID)
	assert.NilError(t, err)
	assert.Assert(t, response != nil)
	response.SetRanFunctionAccepted(rfAccepted).
		SetE2nodeComponentConfigUpdateAck([]*types.E2NodeComponentConfigUpdateAckItem{
			{E2NodeComponentType: e2ap_ies.E2NodeComponentType_E2NODE_COMPONENT_TYPE_G_NB,
				//E2NodeComponentID: e2ncID1,
				E2NodeComponentConfigUpdateAck: types.E2NodeComponentConfigUpdateAck{
					UpdateOutcome: 1,
					//FailureCause: e2ap_ies.Cause{
					//	Cause: &e2ap_ies.Cause_Protocol{
					//		Protocol: e2ap_ies.CauseProtocol_CAUSE_PROTOCOL_TRANSFER_SYNTAX_ERROR,
					//	},
					//},
				}},
			{E2NodeComponentType: e2ap_ies.E2NodeComponentType_E2NODE_COMPONENT_TYPE_E_NB,
				//E2NodeComponentID: e2ncID2,
				E2NodeComponentConfigUpdateAck: types.E2NodeComponentConfigUpdateAck{
					UpdateOutcome: 1,
					//FailureCause: e2ap_ies.Cause{
					//	Cause: &e2ap_ies.Cause_Protocol{
					//		Protocol: e2ap_ies.CauseProtocol_CAUSE_PROTOCOL_ABSTRACT_SYNTAX_ERROR_FALSELY_CONSTRUCTED_MESSAGE,
					//	},
					//},
				}}})

	e2SetupResponseE2APpdu, err := pdubuilder.CreateResponseE2apPdu(response)
	assert.NilError(t, err)

	e2SetupResponseE2APpduC, err := newE2apPdu(e2SetupResponseE2APpdu)
	assert.NilError(t, err)
	assert.Assert(t, e2SetupResponseE2APpduC != nil)

	xer, err := XerEncodeE2apPdu(e2SetupResponseE2APpdu)
	assert.NilError(t, err)
	t.Logf("XER of E2AP: %s\n", string(xer))

	per, err := PerEncodeE2apPdu(e2SetupResponseE2APpdu)
	assert.NilError(t, err)
	t.Logf("PER of E2AP: %v\n", hex.Dump(per))

}
