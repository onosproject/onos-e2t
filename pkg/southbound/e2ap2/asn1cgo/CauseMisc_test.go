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

func createCauseMiscOverload() e2ap_ies.CauseMisc {
	return e2ap_ies.CauseMisc_CAUSE_MISC_CONTROL_PROCESSING_OVERLOAD
}

func createCauseMiscFailure() e2ap_ies.CauseMisc {
	return e2ap_ies.CauseMisc_CAUSE_MISC_HARDWARE_FAILURE
}

func createCauseMiscIntervention() e2ap_ies.CauseMisc {
	return e2ap_ies.CauseMisc_CAUSE_MISC_OM_INTERVENTION
}

func createCauseMiscUnspecified() e2ap_ies.CauseMisc {
	return e2ap_ies.CauseMisc_CAUSE_MISC_UNSPECIFIED
}

func Test_xerEncodingCauseMisc(t *testing.T) {

	causeMisc := createCauseMiscOverload()

	xer, err := xerEncodeCauseMisc(&causeMisc)
	assert.NilError(t, err)
	assert.Equal(t, 54, len(xer))
	t.Logf("CauseMisc (OVERLOAD) XER\n%s", string(xer))

	result, err := xerDecodeCauseMisc(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("CauseMisc (OVERLOAD) XER - decoded\n%v", result)
	if result != nil {
		assert.Equal(t, causeMisc, *result)
	}

	causeMisc = createCauseMiscFailure()

	xer, err = xerEncodeCauseMisc(&causeMisc)
	assert.NilError(t, err)
	assert.Equal(t, 43, len(xer))
	t.Logf("CauseMisc (FAILURE) XER\n%s", string(xer))

	result, err = xerDecodeCauseMisc(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("CauseMisc (FAILURE) XER - decoded\n%v", result)
	if result != nil {
		assert.Equal(t, causeMisc, *result)
	}

	causeMisc = createCauseMiscIntervention()

	xer, err = xerEncodeCauseMisc(&causeMisc)
	assert.NilError(t, err)
	assert.Equal(t, 42, len(xer))
	t.Logf("CauseMisc (INTERVENTION) XER\n%s", string(xer))

	result, err = xerDecodeCauseMisc(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("CauseMisc (INTERVENTION) XER - decoded\n%v", result)
	if result != nil {
		assert.Equal(t, causeMisc, *result)
	}

	causeMisc = createCauseMiscUnspecified()

	xer, err = xerEncodeCauseMisc(&causeMisc)
	assert.NilError(t, err)
	assert.Equal(t, 38, len(xer))
	t.Logf("CauseMisc (UNSPECIFIED) XER\n%s", string(xer))

	result, err = xerDecodeCauseMisc(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("CauseMisc (UNSPECIFIED) XER - decoded\n%v", result)
	if result != nil {
		assert.Equal(t, causeMisc, *result)
	}
}

func Test_perEncodingCauseMisc(t *testing.T) {

	causeMisc := createCauseMiscOverload()

	per, err := perEncodeCauseMisc(&causeMisc)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(per))
	t.Logf("CauseMisc (OVERLOAD) PER\n%v", hex.Dump(per))

	result, err := perDecodeCauseMisc(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("CauseMisc (OVERLOAD) PER - decoded\n%v", result)
	if result != nil {
		assert.Equal(t, causeMisc, *result)
	}

	causeMisc = createCauseMiscFailure()

	per, err = perEncodeCauseMisc(&causeMisc)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(per))
	t.Logf("CauseMisc (FAILURE) PER\n%v", hex.Dump(per))

	result, err = perDecodeCauseMisc(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("CauseMisc (FAILURE) PER - decoded\n%v", result)
	if result != nil {
		assert.Equal(t, causeMisc, *result)
	}

	causeMisc = createCauseMiscIntervention()

	per, err = perEncodeCauseMisc(&causeMisc)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(per))
	t.Logf("CauseMisc (INTERVENTION) PER\n%v", hex.Dump(per))

	result, err = perDecodeCauseMisc(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("CauseMisc (INTERVENTION) PER - decoded\n%v", result)
	if result != nil {
		assert.Equal(t, causeMisc, *result)
	}

	causeMisc = createCauseMiscUnspecified()

	per, err = perEncodeCauseMisc(&causeMisc)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(per))
	t.Logf("CauseMisc (UNSPECIFIED) PER\n%v", hex.Dump(per))

	result, err = perDecodeCauseMisc(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("CauseMisc (UNSPECIFIED) PER - decoded\n%v", result)
	if result != nil {
		assert.Equal(t, causeMisc, *result)
	}
}
