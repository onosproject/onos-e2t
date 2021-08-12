// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

import (
	"gotest.tools/assert"
	"testing"
)

const trasnactionIDxer = `<TransactionID>2</TransactionID>`

func Test_TransactionID(t *testing.T) {
	result, err := xerDecodeTransactionID([]byte(trasnactionIDxer))
	assert.NilError(t, err, "Unexpected error when decoding XER payload")
	assert.Equal(t, int32(2), result.Value)
}
