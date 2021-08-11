// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

import (
	"gotest.tools/assert"
	"testing"
)

const RicreqidXer = `<RICrequestID>
	<ricRequestorID>543210</ricRequestorID>
	<ricInstanceID>6789</ricInstanceID>
</RICrequestID>`

func Test_RicRequestID(t *testing.T) {
	result, err := xerDecodeRicRequestID([]byte(RicreqidXer))
	assert.NilError(t, err, "Unexpected error when decoding XER payload")
	assert.Equal(t, int32(543210), result.RicRequestorId)
	assert.Equal(t, int32(6789), result.RicInstanceId)
}
