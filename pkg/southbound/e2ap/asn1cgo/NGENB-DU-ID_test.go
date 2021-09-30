// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package asn1cgo

import (
	"encoding/hex"
	"testing"

	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-ies"
	"gotest.tools/assert"
)

func createNgEnbDuID() *e2apies.NgenbDuId {

	return &e2apies.NgenbDuId{
		Value: 1234,
	}
}

func Test_xerEncodeNgEnbDuID(t *testing.T) {

	ngEnbDuID := createNgEnbDuID()

	xer, err := xerEncodeNgEnbDuID(ngEnbDuID)
	assert.NilError(t, err)
	t.Logf("NgEnbDuID XER\n%s", string(xer))
}

func Test_xerDecodeNgEnbDuID(t *testing.T) {

	ngEnbDuID := createNgEnbDuID()

	xer, err := xerEncodeNgEnbDuID(ngEnbDuID)
	assert.NilError(t, err)
	t.Logf("NgEnbDuID XER\n%s", string(xer))

	result, err := xerDecodeNgEnbDuID(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("NgEnbDuID XER - decoded\n%s", result)
}

func Test_perEncodeNgEnbDuID(t *testing.T) {

	ngEnbDuID := createNgEnbDuID()

	per, err := perEncodeNgEnbDuID(ngEnbDuID)
	assert.NilError(t, err)
	t.Logf("NgEnbDuID PER\n%v", hex.Dump(per))
}

func Test_perDecodeNgEnbDuID(t *testing.T) {

	ngEnbDuID := createNgEnbDuID()

	per, err := perEncodeNgEnbDuID(ngEnbDuID)
	assert.NilError(t, err)
	t.Logf("NgEnbDuID PER\n%v", hex.Dump(per))

	result, err := perDecodeNgEnbDuID(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("NgEnbDuID PER - decoded\n%s", result)
}
