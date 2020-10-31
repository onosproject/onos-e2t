// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package orane2

import (
	"github.com/onosproject/onos-e2t/pkg/southbound/sandbox"
	"gotest.tools/assert"
	"testing"
)

func Test_newE2setupResponse(t *testing.T) {
	e2apPduE2SetupResponse, err := sandbox.CreateResponseE2apPdu("ONF", 0xABCDE)
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
