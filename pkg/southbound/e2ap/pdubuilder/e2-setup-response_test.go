// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
package pdubuilder

import (
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/asn1cgo"
	"gotest.tools/assert"
	"testing"
)

func TestE2SetupResponse(t *testing.T) {
	newE2apPdu, err := CreateResponseE2apPdu("ONF", 0xABCDE, nil, nil)
	assert.NilError(t, err)
	assert.Assert(t, newE2apPdu != nil)

	xer, err := asn1cgo.XerEncodeE2apPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("E2SetupResponse\n%s", xer)
}
