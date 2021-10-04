// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package asn1cgo

import (
	"encoding/hex"
	"testing"

	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	"gotest.tools/assert"
)

func createE2nodeComponentInterfaceTypeNg() e2ap_ies.E2NodeComponentInterfaceType {
	return e2ap_ies.E2NodeComponentInterfaceType_E2NODE_COMPONENT_INTERFACE_TYPE_NG
}

func createE2nodeComponentInterfaceTypeXn() e2ap_ies.E2NodeComponentInterfaceType {
	return e2ap_ies.E2NodeComponentInterfaceType_E2NODE_COMPONENT_INTERFACE_TYPE_XN
}

func createE2nodeComponentInterfaceTypeE1() e2ap_ies.E2NodeComponentInterfaceType {
	return e2ap_ies.E2NodeComponentInterfaceType_E2NODE_COMPONENT_INTERFACE_TYPE_E1
}

func createE2nodeComponentInterfaceTypeF1() e2ap_ies.E2NodeComponentInterfaceType {
	return e2ap_ies.E2NodeComponentInterfaceType_E2NODE_COMPONENT_INTERFACE_TYPE_F1
}

func createE2nodeComponentInterfaceTypeW1() e2ap_ies.E2NodeComponentInterfaceType {
	return e2ap_ies.E2NodeComponentInterfaceType_E2NODE_COMPONENT_INTERFACE_TYPE_W1
}

func createE2nodeComponentInterfaceTypeS1() e2ap_ies.E2NodeComponentInterfaceType {
	return e2ap_ies.E2NodeComponentInterfaceType_E2NODE_COMPONENT_INTERFACE_TYPE_S1
}

func createE2nodeComponentInterfaceTypeX2() e2ap_ies.E2NodeComponentInterfaceType {
	return e2ap_ies.E2NodeComponentInterfaceType_E2NODE_COMPONENT_INTERFACE_TYPE_X2
}

func Test_xerEncodingTypeE2nodeComponentInterfaceType(t *testing.T) {

	it := createE2nodeComponentInterfaceTypeNg()

	xer, err := xerEncodeE2nodeComponentInterfaceType(&it)
	assert.NilError(t, err)
	t.Logf("E2nodeComponentInterfaceType (NG) XER\n%s", string(xer))

	result, err := xerDecodeE2nodeComponentInterfaceType(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentInterfaceType (NG) XER - decoded\n%v", result)
	assert.Equal(t, it.Number(), result.Number())

	it1 := createE2nodeComponentInterfaceTypeXn()

	xer1, err := xerEncodeE2nodeComponentInterfaceType(&it1)
	assert.NilError(t, err)
	t.Logf("E2nodeComponentInterfaceType (Xn) XER\n%s", string(xer1))

	result1, err := xerDecodeE2nodeComponentInterfaceType(xer1)
	assert.NilError(t, err)
	assert.Assert(t, result1 != nil)
	t.Logf("E2nodeComponentInterfaceType (Xn) XER - decoded\n%v", result1)
	assert.Equal(t, it1.Number(), result1.Number())

	it2 := createE2nodeComponentInterfaceTypeE1()

	xer2, err := xerEncodeE2nodeComponentInterfaceType(&it2)
	assert.NilError(t, err)
	t.Logf("E2nodeComponentInterfaceType (E1) XER\n%s", string(xer2))

	result2, err := xerDecodeE2nodeComponentInterfaceType(xer2)
	assert.NilError(t, err)
	assert.Assert(t, result2 != nil)
	t.Logf("E2nodeComponentInterfaceType (E1) XER - decoded\n%v", result2)
	assert.Equal(t, it2.Number(), result2.Number())

	it3 := createE2nodeComponentInterfaceTypeF1()

	xer3, err := xerEncodeE2nodeComponentInterfaceType(&it3)
	assert.NilError(t, err)
	t.Logf("E2nodeComponentInterfaceType (F1) XER\n%s", string(xer3))

	result3, err := xerDecodeE2nodeComponentInterfaceType(xer3)
	assert.NilError(t, err)
	assert.Assert(t, result3 != nil)
	t.Logf("E2nodeComponentInterfaceType (F1) XER - decoded\n%v", result3)
	assert.Equal(t, it3.Number(), result3.Number())

	it4 := createE2nodeComponentInterfaceTypeW1()

	xer4, err := xerEncodeE2nodeComponentInterfaceType(&it4)
	assert.NilError(t, err)
	t.Logf("E2nodeComponentInterfaceType (W1) XER\n%s", string(xer4))

	result4, err := xerDecodeE2nodeComponentInterfaceType(xer4)
	assert.NilError(t, err)
	assert.Assert(t, result4 != nil)
	t.Logf("E2nodeComponentInterfaceType (W1) XER - decoded\n%v", result4)
	assert.Equal(t, it4.Number(), result4.Number())

	it5 := createE2nodeComponentInterfaceTypeS1()

	xer5, err := xerEncodeE2nodeComponentInterfaceType(&it5)
	assert.NilError(t, err)
	t.Logf("E2nodeComponentInterfaceType (S1) XER\n%s", string(xer5))

	result5, err := xerDecodeE2nodeComponentInterfaceType(xer5)
	assert.NilError(t, err)
	assert.Assert(t, result5 != nil)
	t.Logf("E2nodeComponentInterfaceType (S1) XER - decoded\n%v", result5)
	assert.Equal(t, it5.Number(), result5.Number())

	it6 := createE2nodeComponentInterfaceTypeX2()

	xer6, err := xerEncodeE2nodeComponentInterfaceType(&it6)
	assert.NilError(t, err)
	t.Logf("E2nodeComponentInterfaceType (X2) XER\n%s", string(xer6))

	result6, err := xerDecodeE2nodeComponentInterfaceType(xer6)
	assert.NilError(t, err)
	assert.Assert(t, result6 != nil)
	t.Logf("E2nodeComponentInterfaceType (X2) XER - decoded\n%v", result6)
	assert.Equal(t, it6.Number(), result6.Number())
}

func Test_perEncodingTypeE2nodeComponentInterfaceType(t *testing.T) {

	it := createE2nodeComponentInterfaceTypeNg()

	per, err := perEncodeE2nodeComponentInterfaceType(&it)
	assert.NilError(t, err)
	t.Logf("E2nodeComponentInterfaceType (NG) PER\n%v", hex.Dump(per))

	result, err := perDecodeE2nodeComponentInterfaceType(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentInterfaceType (NG) PER - decoded\n%v", result)
	assert.Equal(t, it.Number(), result.Number())

	it1 := createE2nodeComponentInterfaceTypeXn()

	per1, err := perEncodeE2nodeComponentInterfaceType(&it1)
	assert.NilError(t, err)
	t.Logf("E2nodeComponentInterfaceType (Xn) PER\n%v", hex.Dump(per1))

	result1, err := perDecodeE2nodeComponentInterfaceType(per1)
	assert.NilError(t, err)
	assert.Assert(t, result1 != nil)
	t.Logf("E2nodeComponentInterfaceType (Xn) PER - decoded\n%v", result1)
	assert.Equal(t, it1.Number(), result1.Number())

	it2 := createE2nodeComponentInterfaceTypeE1()

	per2, err := perEncodeE2nodeComponentInterfaceType(&it2)
	assert.NilError(t, err)
	t.Logf("E2nodeComponentInterfaceType (E1) PER\n%v", hex.Dump(per2))

	result2, err := perDecodeE2nodeComponentInterfaceType(per2)
	assert.NilError(t, err)
	assert.Assert(t, result2 != nil)
	t.Logf("E2nodeComponentInterfaceType (E1) PER - decoded\n%v", result2)
	assert.Equal(t, it2.Number(), result2.Number())

	it3 := createE2nodeComponentInterfaceTypeF1()

	per3, err := perEncodeE2nodeComponentInterfaceType(&it3)
	assert.NilError(t, err)
	t.Logf("E2nodeComponentInterfaceType (F1) PER\n%v", hex.Dump(per3))

	result3, err := perDecodeE2nodeComponentInterfaceType(per3)
	assert.NilError(t, err)
	assert.Assert(t, result3 != nil)
	t.Logf("E2nodeComponentInterfaceType (F1) PER - decoded\n%v", result3)
	assert.Equal(t, it3.Number(), result3.Number())

	it4 := createE2nodeComponentInterfaceTypeW1()

	per4, err := perEncodeE2nodeComponentInterfaceType(&it4)
	assert.NilError(t, err)
	t.Logf("E2nodeComponentInterfaceType (W1) PER\n%v", hex.Dump(per4))

	result4, err := perDecodeE2nodeComponentInterfaceType(per4)
	assert.NilError(t, err)
	assert.Assert(t, result4 != nil)
	t.Logf("E2nodeComponentInterfaceType (W1) PER - decoded\n%v", result4)
	assert.Equal(t, it4.Number(), result4.Number())

	it5 := createE2nodeComponentInterfaceTypeS1()

	per5, err := perEncodeE2nodeComponentInterfaceType(&it5)
	assert.NilError(t, err)
	t.Logf("E2nodeComponentInterfaceType (S1) PER\n%v", hex.Dump(per5))

	result5, err := perDecodeE2nodeComponentInterfaceType(per5)
	assert.NilError(t, err)
	assert.Assert(t, result5 != nil)
	t.Logf("E2nodeComponentInterfaceType (S1) PER - decoded\n%v", result5)
	assert.Equal(t, it5.Number(), result5.Number())

	it6 := createE2nodeComponentInterfaceTypeX2()

	per6, err := perEncodeE2nodeComponentInterfaceType(&it6)
	assert.NilError(t, err)
	t.Logf("E2nodeComponentInterfaceType (X2) PER\n%v", hex.Dump(per6))

	result6, err := perDecodeE2nodeComponentInterfaceType(per6)
	assert.NilError(t, err)
	assert.Assert(t, result6 != nil)
	t.Logf("E2nodeComponentInterfaceType (X2) PER - decoded\n%v", result6)
	assert.Equal(t, it6.Number(), result6.Number())
}
