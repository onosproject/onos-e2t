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

func createTnlUsageRicService() e2ap_ies.Tnlusage {
	return e2ap_ies.Tnlusage_TNLUSAGE_RIC_SERVICE
}

func createTnlUsageSupportFunction() e2ap_ies.Tnlusage {
	return e2ap_ies.Tnlusage_TNLUSAGE_SUPPORT_FUNCTION
}

func createTnlUsageBoth() e2ap_ies.Tnlusage {
	return e2ap_ies.Tnlusage_TNLUSAGE_BOTH
}

func Test_xerEncodingTnlUsage(t *testing.T) {

	tnlUsage := createTnlUsageRicService()
	t.Logf("TNLusage (message)\n%v", tnlUsage)

	xer, err := xerEncodeTnlusage(&tnlUsage)
	assert.NilError(t, err)
	assert.Equal(t, 36, len(xer))
	t.Logf("TNLusage (RICservice) XER\n%s", string(xer))

	result, err := xerDecodeTnlusage(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TNLusage (RICservice) XER - decoded\n%v", result)
	//assert.Equal(t, tnlUsage, result)

	tnlUsage = createTnlUsageSupportFunction()
	t.Logf("TNLusage (message)\n%v", tnlUsage)

	xer, err = xerEncodeTnlusage(&tnlUsage)
	assert.NilError(t, err)
	assert.Equal(t, 41, len(xer))
	t.Logf("TNLusage (SupportFunction) XER\n%s", string(xer))

	result, err = xerDecodeTnlusage(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TNLusage (SupportFunction) XER - decoded\n%v", result)
	//assert.Equal(t, tnlUsage, result)

	tnlUsage = createTnlUsageBoth()
	t.Logf("TNLusage (message)\n%v", tnlUsage)

	xer, err = xerEncodeTnlusage(&tnlUsage)
	assert.NilError(t, err)
	assert.Equal(t, 29, len(xer))
	t.Logf("TNLusage (Both) XER\n%s", string(xer))

	result, err = xerDecodeTnlusage(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TNLusage (Both) XER - decoded\n%v", result)
	//assert.Equal(t, tnlUsage, result)
}

func Test_perEncodingTnlUsage(t *testing.T) {

	tnlUsage := createTnlUsageRicService()
	t.Logf("TNLusage (message)\n%v", tnlUsage)

	per, err := perEncodeTnlusage(&tnlUsage)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(per))
	t.Logf("TNLusage (RICservice) PER\n%v", hex.Dump(per))

	result, err := perDecodeTnlusage(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TNLusage (RICservice) PER - decoded\n%v", result)
	//assert.Equal(t, tnlUsage, result)

	tnlUsage = createTnlUsageSupportFunction()
	t.Logf("TNLusage (message)\n%v", tnlUsage)

	per, err = perEncodeTnlusage(&tnlUsage)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(per))
	t.Logf("TNLusage (SupportFunction) PER\n%v", hex.Dump(per))

	result, err = perDecodeTnlusage(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TNLusage (SupportFunction)  PER - decoded\n%v", result)
	//assert.Equal(t, tnlUsage, result)

	tnlUsage = createTnlUsageBoth()
	t.Logf("TNLusage (message)\n%v", tnlUsage)

	per, err = perEncodeTnlusage(&tnlUsage)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(per))
	t.Logf("TNLusage (Both) PER\n%v", hex.Dump(per))

	result, err = perDecodeTnlusage(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TNLusage (Both)  PER - decoded\n%v", result)
	//assert.Equal(t, tnlUsage, result)
}
