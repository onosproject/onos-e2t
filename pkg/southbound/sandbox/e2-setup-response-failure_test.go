// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
package sandbox

import (
	"gotest.tools/assert"
	"testing"
)

func TestE2SetupResponseFailure(t *testing.T) {
	newE2apPdu, err := CreateResponseFailureE2apPdu(21)
	assert.NilError(t, err)
	assert.Assert(t, newE2apPdu != nil)
}
