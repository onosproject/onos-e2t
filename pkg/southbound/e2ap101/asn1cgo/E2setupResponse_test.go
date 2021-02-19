// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

import (
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/pdubuilder"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/types"
	"gotest.tools/assert"
	"testing"
)

func Test_newE2setupResponse(t *testing.T) {
	plmnID := [3]byte{0x79, 0x78, 0x70}
	ricID := types.RicIdentifier{
		RicIdentifierValue: 0xABCDE,
		RicIdentifierLen:   20,
	}

	e2apPduE2SetupResponse, err := pdubuilder.CreateResponseE2apPdu(plmnID, ricID, nil, nil)
	assert.NilError(t, err)

	e2SetupResponse := e2apPduE2SetupResponse.GetSuccessfulOutcome().GetProcedureCode().GetE2Setup().GetSuccessfulOutcome()

	e2srC, err := newE2setupResponse(e2SetupResponse)
	assert.NilError(t, err)
	assert.Assert(t, e2srC != nil)

	xer, err := xerEncodeE2setupResponse(e2SetupResponse)
	assert.NilError(t, err)
	t.Logf("E2SetupResponse XER\n%s", string(xer))

	per, err := perEncodeE2setupResponse(e2SetupResponse)
	assert.NilError(t, err)
	t.Logf("E2SetupResponse PER\n%v", per)
}
