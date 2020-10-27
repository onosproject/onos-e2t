// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
package sandbox

import (
	"gotest.tools/assert"
	"testing"
)

func TestE2SetupResponse(t *testing.T) {
	newE2apPdu, err := CreateResponseE2apPdu("ONF")
	assert.NilError(t, err)
	assert.Assert(t, newE2apPdu != nil)
}
