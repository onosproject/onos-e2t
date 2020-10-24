// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
package sandbox

import (
	"gotest.tools/assert"
	"testing"
)

func TestE2SetupRequest(t *testing.T) {
	newE2apPdu, err := CreateE2apPdu("ONF", 100, 200)
	assert.NilError(t, err)
	assert.Assert(t, newE2apPdu != nil)
}
