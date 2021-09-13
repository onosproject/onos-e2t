// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

import (
	"encoding/hex"
	"testing"

	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-ies"
	"gotest.tools/assert"
)

const trasnactionIDxer = `<TransactionID>2</TransactionID>`

func Test_TransactionID(t *testing.T) {
	result, err := xerDecodeTransactionID([]byte(trasnactionIDxer))
	assert.NilError(t, err, "Unexpected error when decoding XER payload")
	t.Logf("Transaction ID decoded from XER is \n%v", result)
	assert.Equal(t, int32(2), result.Value)
}

func TestTransactionIDlb(t *testing.T) {
	trID := &e2apies.TransactionId{
		Value: 0,
	}

	xer, err := xerEncodeTransactionID(trID)
	assert.NilError(t, err)
	assert.Assert(t, xer != nil)
	t.Logf("Transaction ID XER is \n%s", xer)

	result, err := xerDecodeTransactionID(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("Transaction ID decoded from XER is \n%v", result)
	assert.Equal(t, trID.GetValue(), result.GetValue())

	per, err := perEncodeTransactionID(trID)
	assert.NilError(t, err)
	assert.Assert(t, xer != nil)
	t.Logf("Transaction ID PER is \n%v", hex.Dump(per))

	result1, err := perDecodeTransactionID(per)
	assert.NilError(t, err)
	assert.Assert(t, result1 != nil)
	t.Logf("Transaction ID decoded from PER is \n%v", result1)
	assert.Equal(t, trID.GetValue(), result1.GetValue())
}

func TestTransactionIDub(t *testing.T) {
	trID := &e2apies.TransactionId{
		Value: 255,
	}

	xer, err := xerEncodeTransactionID(trID)
	assert.NilError(t, err)
	assert.Assert(t, xer != nil)
	t.Logf("Transaction ID XER is \n%s", xer)

	result, err := xerDecodeTransactionID(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("Transaction ID decoded from XER is \n%v", result)
	assert.Equal(t, trID.GetValue(), result.GetValue())

	per, err := perEncodeTransactionID(trID)
	assert.NilError(t, err)
	assert.Assert(t, xer != nil)
	t.Logf("Transaction ID PER is \n%v", hex.Dump(per))

	result1, err := perDecodeTransactionID(per)
	assert.NilError(t, err)
	assert.Assert(t, result1 != nil)
	t.Logf("Transaction ID decoded from PER is \n%v", result1)
	assert.Equal(t, trID.GetValue(), result1.GetValue())
}

func TestTransactionIDmb(t *testing.T) {
	trID := &e2apies.TransactionId{
		Value: 127,
	}

	xer, err := xerEncodeTransactionID(trID)
	assert.NilError(t, err)
	assert.Assert(t, xer != nil)
	t.Logf("Transaction ID XER is \n%s", xer)

	result, err := xerDecodeTransactionID(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("Transaction ID decoded from XER is \n%v", result)
	assert.Equal(t, trID.GetValue(), result.GetValue())

	per, err := perEncodeTransactionID(trID)
	assert.NilError(t, err)
	assert.Assert(t, xer != nil)
	t.Logf("Transaction ID PER is \n%v", hex.Dump(per))

	result1, err := perDecodeTransactionID(per)
	assert.NilError(t, err)
	assert.Assert(t, result1 != nil)
	t.Logf("Transaction ID decoded from PER is \n%v", result1)
	assert.Equal(t, trID.GetValue(), result1.GetValue())
}
