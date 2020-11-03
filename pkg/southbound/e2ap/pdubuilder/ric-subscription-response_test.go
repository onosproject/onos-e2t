// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
package pdubuilder

import (
	"gotest.tools/assert"
	"testing"
)

func TestRicSubscriptionResponse(t *testing.T) {
	newE2apPdu, err := CreateRicSubscriptionResponseE2apPdu(21, 22, 9)
	assert.NilError(t, err)
	assert.Assert(t, newE2apPdu != nil)
}
