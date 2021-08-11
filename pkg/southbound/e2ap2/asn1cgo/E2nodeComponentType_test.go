// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

import (
	"encoding/hex"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	"gotest.tools/assert"
	"testing"
)

func createE2nodeComponentTypeGNb() e2ap_ies.E2NodeComponentType {
	return e2ap_ies.E2NodeComponentType_E2NODE_COMPONENT_TYPE_G_NB
}

func createE2nodeComponentTypeGNbCUUP() e2ap_ies.E2NodeComponentType {
	return e2ap_ies.E2NodeComponentType_E2NODE_COMPONENT_TYPE_G_NB_CU_UP
}

func createE2nodeComponentTypeGNbDU() e2ap_ies.E2NodeComponentType {
	return e2ap_ies.E2NodeComponentType_E2NODE_COMPONENT_TYPE_G_NB_DU
}

func createE2nodeComponentTypeEnGNb() e2ap_ies.E2NodeComponentType {
	return e2ap_ies.E2NodeComponentType_E2NODE_COMPONENT_TYPE_EN_G_NB
}

func createE2nodeComponentTypeENb() e2ap_ies.E2NodeComponentType {
	return e2ap_ies.E2NodeComponentType_E2NODE_COMPONENT_TYPE_E_NB
}

func createE2nodeComponentTypeNgENb() e2ap_ies.E2NodeComponentType {
	return e2ap_ies.E2NodeComponentType_E2NODE_COMPONENT_TYPE_NG_E_NB
}

func Test_xerEncodingE2nodeComponentType(t *testing.T) {

	componentType := createE2nodeComponentTypeGNb()

	xer, err := xerEncodeE2nodeComponentType(&componentType)
	assert.NilError(t, err)
	assert.Equal(t, 50, len(xer))
	t.Logf("E2nodeComponentType (GNb) XER\n%s", string(xer))

	result, err := xerDecodeE2nodeComponentType(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentType (GNb) XER - decoded\n%v", result)
	if result != nil {
		assert.Equal(t, componentType, *result)
	}

	componentType = createE2nodeComponentTypeGNbCUUP()

	xer, err = xerEncodeE2nodeComponentType(&componentType)
	assert.NilError(t, err)
	assert.Equal(t, 56, len(xer))
	t.Logf("E2nodeComponentType (GNb-CU-UP) XER\n%s", string(xer))

	result, err = xerDecodeE2nodeComponentType(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentType (GNb-CU-UP) XER - decoded\n%v", result)
	if result != nil {
		assert.Equal(t, componentType, *result)
	}

	componentType = createE2nodeComponentTypeGNbDU()

	xer, err = xerEncodeE2nodeComponentType(&componentType)
	assert.NilError(t, err)
	assert.Equal(t, 53, len(xer))
	t.Logf("E2nodeComponentType (GNb-DU) XER\n%s", string(xer))

	result, err = xerDecodeE2nodeComponentType(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentType (GNb-DU) XER - decoded\n%v", result)
	if result != nil {
		assert.Equal(t, componentType, *result)
	}

	componentType = createE2nodeComponentTypeEnGNb()

	xer, err = xerEncodeE2nodeComponentType(&componentType)
	assert.NilError(t, err)
	assert.Equal(t, 53, len(xer))
	t.Logf("E2nodeComponentType (en-GNb) XER\n%s", string(xer))

	result, err = xerDecodeE2nodeComponentType(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentType (en-GNb) XER - decoded\n%v", result)
	if result != nil {
		assert.Equal(t, componentType, *result)
	}

	componentType = createE2nodeComponentTypeENb()

	xer, err = xerEncodeE2nodeComponentType(&componentType)
	assert.NilError(t, err)
	assert.Equal(t, 50, len(xer))
	t.Logf("E2nodeComponentType (ENb) XER\n%s", string(xer))

	result, err = xerDecodeE2nodeComponentType(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentType (ENb) XER - decoded\n%v", result)
	if result != nil {
		assert.Equal(t, componentType, *result)
	}

	componentType = createE2nodeComponentTypeNgENb()

	xer, err = xerEncodeE2nodeComponentType(&componentType)
	assert.NilError(t, err)
	assert.Equal(t, 53, len(xer))
	t.Logf("E2nodeComponentType (ng-ENb) XER\n%s", string(xer))

	result, err = xerDecodeE2nodeComponentType(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentType (ng-ENb) XER - decoded\n%v", result)
	if result != nil {
		assert.Equal(t, componentType, *result)
	}
}

func Test_perEncodingE2nodeComponentType(t *testing.T) {

	componentType := createE2nodeComponentTypeGNb()

	per, err := perEncodeE2nodeComponentType(&componentType)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(per))
	t.Logf("E2nodeComponentType (GNb) PER\n%v", hex.Dump(per))

	result, err := perDecodeE2nodeComponentType(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentType (GNb) PER - decoded\n%v", result)
	if result != nil {
		assert.Equal(t, componentType, *result)
	}

	componentType = createE2nodeComponentTypeGNbCUUP()

	per, err = perEncodeE2nodeComponentType(&componentType)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(per))
	t.Logf("E2nodeComponentType (GNb-CU-UP) PER\n%v", hex.Dump(per))

	result, err = perDecodeE2nodeComponentType(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentType (GNb-CU-UP) PER - decoded\n%v", result)
	if result != nil {
		assert.Equal(t, componentType, *result)
	}

	componentType = createE2nodeComponentTypeGNbDU()

	per, err = perEncodeE2nodeComponentType(&componentType)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(per))
	t.Logf("E2nodeComponentType (GNb-DU) PER\n%v", hex.Dump(per))

	result, err = perDecodeE2nodeComponentType(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentType (GNb-DU) PER - decoded\n%v", result)
	if result != nil {
		assert.Equal(t, componentType, *result)
	}

	componentType = createE2nodeComponentTypeEnGNb()

	per, err = perEncodeE2nodeComponentType(&componentType)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(per))
	t.Logf("E2nodeComponentType (en-GNb) PER\n%v", hex.Dump(per))

	result, err = perDecodeE2nodeComponentType(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentType (en-GNb) PER - decoded\n%v", result)
	if result != nil {
		assert.Equal(t, componentType, *result)
	}

	componentType = createE2nodeComponentTypeENb()

	per, err = perEncodeE2nodeComponentType(&componentType)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(per))
	t.Logf("E2nodeComponentType (ENb) PER\n%v", hex.Dump(per))

	result, err = perDecodeE2nodeComponentType(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentType (ENb) PER - decoded\n%v", result)
	if result != nil {
		assert.Equal(t, componentType, *result)
	}

	componentType = createE2nodeComponentTypeNgENb()

	per, err = perEncodeE2nodeComponentType(&componentType)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(per))
	t.Logf("E2nodeComponentType (ng-ENb) PER\n%v", hex.Dump(per))

	result, err = perDecodeE2nodeComponentType(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentType (ng-ENb) PER - decoded\n%v", result)
	if result != nil {
		assert.Equal(t, componentType, *result)
	}
}
