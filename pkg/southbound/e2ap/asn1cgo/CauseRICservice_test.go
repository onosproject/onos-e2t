// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

import (
	"encoding/hex"
	"testing"

	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	"gotest.tools/assert"
)

func createCauseRicServiceFunctionNotRequeired() e2ap_ies.CauseRicservice {
	return e2ap_ies.CauseRicservice_CAUSE_RICSERVICE_RAN_FUNCTION_NOT_SUPPORTED
}

func createCauseRicServiceExcessiveFunctions() e2ap_ies.CauseRicservice {
	return e2ap_ies.CauseRicservice_CAUSE_RICSERVICE_EXCESSIVE_FUNCTIONS
}

func createCauseRicServiceResourceLimit() e2ap_ies.CauseRicservice {
	return e2ap_ies.CauseRicservice_CAUSE_RICSERVICE_RIC_RESOURCE_LIMIT
}

func Test_xerEncodingCauseRicService(t *testing.T) {

	causeRicService := createCauseRicServiceFunctionNotRequeired()

	xer, err := xerEncodeCauseRicservice(&causeRicService)
	assert.NilError(t, err)
	t.Logf("CauseRICservice (Function Not Required) XER\n%s", string(xer))

	result, err := xerDecodeCauseRicservice(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("CauseRICservice (Function Not Required) XER - decoded\n%v", result)
	assert.Equal(t, causeRicService.Number(), result.Number())

	causeRicService = createCauseRicServiceExcessiveFunctions()

	xer, err = xerEncodeCauseRicservice(&causeRicService)
	assert.NilError(t, err)
	t.Logf("CauseRICservice (Excessive Functions) XER\n%s", string(xer))

	result, err = xerDecodeCauseRicservice(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("CauseRICservice (Excessive Functions) XER - decoded\n%v", result)
	assert.Equal(t, causeRicService.Number(), result.Number())

	causeRicService = createCauseRicServiceResourceLimit()

	xer, err = xerEncodeCauseRicservice(&causeRicService)
	assert.NilError(t, err)
	t.Logf("CauseRICservice (Resource Limit) XER\n%s", string(xer))

	result, err = xerDecodeCauseRicservice(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("CauseRICservice (Resource Limit) XER - decoded\n%v", result)
	assert.Equal(t, causeRicService.Number(), result.Number())
}

func Test_perEncodingCauseRicService(t *testing.T) {

	causeRicService := createCauseRicServiceFunctionNotRequeired()

	per, err := perEncodeCauseRicservice(&causeRicService)
	assert.NilError(t, err)
	t.Logf("CauseRICservice (Function Not Required) PER\n%v", hex.Dump(per))

	result, err := perDecodeCauseRicservice(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("CauseRICservice (Function Not Required) PER - decoded\n%v", result)
	assert.Equal(t, causeRicService.Number(), result.Number())

	causeRicService = createCauseRicServiceExcessiveFunctions()

	per, err = perEncodeCauseRicservice(&causeRicService)
	assert.NilError(t, err)
	t.Logf("CauseRICservice (Excessive Functions) PER\n%v", hex.Dump(per))

	result, err = perDecodeCauseRicservice(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("CauseRICservice (Excessive Functions) PER - decoded\n%v", result)
	assert.Equal(t, causeRicService.Number(), result.Number())
	causeRicService = createCauseRicServiceResourceLimit()

	per, err = perEncodeCauseRicservice(&causeRicService)
	assert.NilError(t, err)
	t.Logf("CauseRICservice (Resource Limit) PER\n%v", hex.Dump(per))

	result, err = perDecodeCauseRicservice(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("CauseRICservice (Resource Limit) PER - decoded\n%v", result)
	assert.Equal(t, causeRicService.Number(), result.Number())
}
