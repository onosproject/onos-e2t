// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package asn1cgo

import (
	"encoding/hex"
	"testing"

	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/pdubuilder"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
	"gotest.tools/assert"
)

func Test_newE2setupResponse(t *testing.T) {
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
		E2NodeComponentID:   e2ncID3,
		E2NodeComponentType: e2ap_ies.E2NodeComponentInterfaceType_E2NODE_COMPONENT_INTERFACE_TYPE_S1,
	}
	e2nccaal = append(e2nccaal, &ie1)

	e2SetupResponse, err := pdubuilder.NewE2SetupResponse(1, plmnID, ricID, e2nccaal)
	assert.NilError(t, err)
	assert.Assert(t, e2SetupResponse != nil)

	e2srC, err := newE2setupResponse(e2SetupResponse)
	assert.NilError(t, err)
	assert.Assert(t, e2srC != nil)

	xer, err := xerEncodeE2setupResponse(e2SetupResponse)
	assert.NilError(t, err)
	t.Logf("E2SetupResponse XER\n%s", string(xer))

	per, err := perEncodeE2setupResponse(e2SetupResponse)
	assert.NilError(t, err)
	t.Logf("E2SetupResponse PER\n%v", hex.Dump(per))
}
