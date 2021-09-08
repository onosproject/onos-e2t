// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

import (
	"encoding/hex"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-commondatatypes"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	"gotest.tools/assert"
	"testing"
)

func createEnbIDChoiceMacro() *e2ap_ies.EnbIdChoice {

	return &e2ap_ies.EnbIdChoice{
		EnbIdChoice: &e2ap_ies.EnbIdChoice_EnbIdMacro{
			EnbIdMacro: &e2ap_commondatatypes.BitString{
				Value: []byte{0xd4, 0xcb, 0x90},
				Len:   20,
			},
		},
	}
}

func createEnbIDChoiceShortMacro() *e2ap_ies.EnbIdChoice {

	return &e2ap_ies.EnbIdChoice{
		EnbIdChoice: &e2ap_ies.EnbIdChoice_EnbIdShortmacro{
			EnbIdShortmacro: &e2ap_commondatatypes.BitString{
				Value: []byte{0xd4, 0xcb, 0xc0},
				Len:   18,
			},
		},
	}
}

func createEnbIDChoiceLongMacro() *e2ap_ies.EnbIdChoice {

	return &e2ap_ies.EnbIdChoice{
		EnbIdChoice: &e2ap_ies.EnbIdChoice_EnbIdLongmacro{
			EnbIdLongmacro: &e2ap_commondatatypes.BitString{
				Value: []byte{0xd4, 0xcb, 0xf8},
				Len:   21,
			},
		},
	}
}

func Test_xerEncodeEnbIDChoice(t *testing.T) {

	enbIDchoice := createEnbIDChoiceMacro()

	xer, err := xerEncodeEnbIDChoice(enbIDchoice)
	assert.NilError(t, err)
	assert.Equal(t, 101, len(xer))
	t.Logf("EnbIDChoice (Macro) XER\n%s", string(xer))

	enbIDchoice = createEnbIDChoiceShortMacro()

	xer, err = xerEncodeEnbIDChoice(enbIDchoice)
	assert.NilError(t, err)
	assert.Equal(t, 109, len(xer))
	t.Logf("EnbIDChoice (ShortMacro) XER\n%s", string(xer))

	enbIDchoice = createEnbIDChoiceLongMacro()

	xer, err = xerEncodeEnbIDChoice(enbIDchoice)
	assert.NilError(t, err)
	assert.Equal(t, 110, len(xer))
	t.Logf("EnbIDChoice (LongMacro) XER\n%s", string(xer))
}

func Test_xerDecodeEnbIDChoice(t *testing.T) {

	enbIDchoice := createEnbIDChoiceMacro()

	xer, err := xerEncodeEnbIDChoice(enbIDchoice)
	assert.NilError(t, err)
	assert.Equal(t, 101, len(xer))
	t.Logf("EnbIDChoice (Macro) XER\n%s", string(xer))

	result, err := xerDecodeEnbIDChoice(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("EnbIDChoice (Macro) XER - decoded\n%s", result)
	assert.Equal(t, enbIDchoice.GetEnbIdMacro().GetLen(), result.GetEnbIdMacro().GetLen())
	assert.DeepEqual(t, enbIDchoice.GetEnbIdMacro().GetValue(), result.GetEnbIdMacro().GetValue())

	enbIDchoice = createEnbIDChoiceShortMacro()

	xer, err = xerEncodeEnbIDChoice(enbIDchoice)
	assert.NilError(t, err)
	assert.Equal(t, 109, len(xer))
	t.Logf("EnbIDChoice (ShortMacro) XER\n%s", string(xer))

	result, err = xerDecodeEnbIDChoice(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("EnbIDChoice (ShortMacro) XER - decoded\n%s", result)
	assert.Equal(t, enbIDchoice.GetEnbIdShortmacro().GetLen(), result.GetEnbIdShortmacro().GetLen())
	assert.DeepEqual(t, enbIDchoice.GetEnbIdShortmacro().GetValue(), result.GetEnbIdShortmacro().GetValue())

	enbIDchoice = createEnbIDChoiceLongMacro()

	xer, err = xerEncodeEnbIDChoice(enbIDchoice)
	assert.NilError(t, err)
	assert.Equal(t, 110, len(xer))
	t.Logf("EnbIDChoice (LongMacro) XER\n%s", string(xer))

	result, err = xerDecodeEnbIDChoice(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("EnbIDChoice (LongMacro) XER - decoded\n%s", result)
	assert.Equal(t, enbIDchoice.GetEnbIdLongmacro().GetLen(), result.GetEnbIdLongmacro().GetLen())
	assert.DeepEqual(t, enbIDchoice.GetEnbIdLongmacro().GetValue(), result.GetEnbIdLongmacro().GetValue())
}

func Test_perEncodeEnbIDChoice(t *testing.T) {

	enbIDchoice := createEnbIDChoiceMacro()

	per, err := perEncodeEnbIDChoice(enbIDchoice)
	assert.NilError(t, err)
	assert.Equal(t, 4, len(per))
	t.Logf("EnbIDChoice (Macro) PER\n%v", hex.Dump(per))

	enbIDchoice = createEnbIDChoiceShortMacro()

	per, err = perEncodeEnbIDChoice(enbIDchoice)
	assert.NilError(t, err)
	assert.Equal(t, 4, len(per))
	t.Logf("EnbIDChoice (ShortMacro)PER\n%v", hex.Dump(per))

	enbIDchoice = createEnbIDChoiceLongMacro()

	per, err = perEncodeEnbIDChoice(enbIDchoice)
	assert.NilError(t, err)
	assert.Equal(t, 4, len(per))
	t.Logf("EnbIDChoice (LongMacro) PER\n%v", hex.Dump(per))
}

func Test_perDecodeEnbIDChoice(t *testing.T) {

	enbIDchoice := createEnbIDChoiceMacro()

	per, err := perEncodeEnbIDChoice(enbIDchoice)
	assert.NilError(t, err)
	assert.Equal(t, 4, len(per))
	t.Logf("EnbIDChoice (Macro) PER\n%v", hex.Dump(per))

	result, err := perDecodeEnbIDChoice(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("EnbIDChoice (Macro) PER - decoded\n%s", result)
	assert.Equal(t, enbIDchoice.GetEnbIdMacro().GetLen(), result.GetEnbIdMacro().GetLen())
	assert.DeepEqual(t, enbIDchoice.GetEnbIdMacro().GetValue(), result.GetEnbIdMacro().GetValue())

	enbIDchoice = createEnbIDChoiceShortMacro()

	per, err = perEncodeEnbIDChoice(enbIDchoice)
	assert.NilError(t, err)
	assert.Equal(t, 4, len(per))
	t.Logf("EnbIDChoice (ShortMacro) PER\n%v", hex.Dump(per))

	result, err = perDecodeEnbIDChoice(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("EnbIDChoice (ShortMacro) PER - decoded\n%s", result)
	assert.Equal(t, enbIDchoice.GetEnbIdShortmacro().GetLen(), result.GetEnbIdShortmacro().GetLen())
	assert.DeepEqual(t, enbIDchoice.GetEnbIdShortmacro().GetValue(), result.GetEnbIdShortmacro().GetValue())

	enbIDchoice = createEnbIDChoiceLongMacro()

	per, err = perEncodeEnbIDChoice(enbIDchoice)
	assert.NilError(t, err)
	assert.Equal(t, 4, len(per))
	t.Logf("EnbIDChoice (LongMacro) PER\n%v", hex.Dump(per))

	result, err = perDecodeEnbIDChoice(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("EnbIDChoice (LongMacro) PER - decoded\n%s", result)
	assert.Equal(t, enbIDchoice.GetEnbIdLongmacro().GetLen(), result.GetEnbIdLongmacro().GetLen())
	assert.DeepEqual(t, enbIDchoice.GetEnbIdLongmacro().GetValue(), result.GetEnbIdLongmacro().GetValue())
}
