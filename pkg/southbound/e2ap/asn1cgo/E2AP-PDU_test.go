// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

import (
	"encoding/hex"
	"testing"

	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/pdubuilder"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
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

	e2ncID3 := pdubuilder.CreateE2NodeComponentIDS1("S1-component")
	e2nccaal := make([]*types.E2NodeComponentConfigAdditionAckItem, 0)
	ie1 := types.E2NodeComponentConfigAdditionAckItem{
		E2NodeComponentConfigurationAck: e2ap_ies.E2NodeComponentConfigurationAck{
			UpdateOutcome: e2ap_ies.UpdateOutcome_UPDATE_OUTCOME_SUCCESS,
		},
		E2NodeComponentID: e2ncID3,
		E2NodeComponentType: e2ap_ies.E2NodeComponentInterfaceType_E2NODE_COMPONENT_INTERFACE_TYPE_S1,
	}
	e2nccaal = append(e2nccaal, &ie1)

	response, err := pdubuilder.NewE2SetupResponse(1, plmnID, ricID, e2nccaal)
	assert.NilError(t, err)
	assert.Assert(t, response != nil)
	response.SetRanFunctionAccepted(rfAccepted)

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
