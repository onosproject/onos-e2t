// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
package sandbox

import (
	"gotest.tools/assert"
	"testing"
)

func TestRicSubscriptionRequest(t *testing.T) {
	newE2apPdu, err := CreateRicSubscriptionRequestE2apPdu(21, 22, 9, 15, 29, 17 )
	assert.NilError(t, err)
	assert.Assert(t, newE2apPdu != nil)
} 

