// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package asn1cgo

import (
	"encoding/hex"
	"testing"

	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	"gotest.tools/assert"
)

func createCauseRicRfIDInvalid() e2ap_ies.CauseRicrequest {
	return e2ap_ies.CauseRicrequest_CAUSE_RICREQUEST_RAN_FUNCTION_ID_INVALID
}

func createCauseRicActionNotSupported() e2ap_ies.CauseRicrequest {
	return e2ap_ies.CauseRicrequest_CAUSE_RICREQUEST_ACTION_NOT_SUPPORTED
}

func createCauseRicExcessiveActions() e2ap_ies.CauseRicrequest {
	return e2ap_ies.CauseRicrequest_CAUSE_RICREQUEST_EXCESSIVE_ACTIONS
}

func createCauseRicDuplicateAction() e2ap_ies.CauseRicrequest {
	return e2ap_ies.CauseRicrequest_CAUSE_RICREQUEST_DUPLICATE_ACTION
}

func createCauseRicDuplicateEvent() e2ap_ies.CauseRicrequest {
	return e2ap_ies.CauseRicrequest_CAUSE_RICREQUEST_DUPLICATE_EVENT_TRIGGER
}

func createCauseRicFunctionResourceLimit() e2ap_ies.CauseRicrequest {
	return e2ap_ies.CauseRicrequest_CAUSE_RICREQUEST_FUNCTION_RESOURCE_LIMIT
}

func createCauseRicRequestIDUnknown() e2ap_ies.CauseRicrequest {
	return e2ap_ies.CauseRicrequest_CAUSE_RICREQUEST_REQUEST_ID_UNKNOWN
}

func createCauseRicActionSubsequentActionSequence() e2ap_ies.CauseRicrequest {
	return e2ap_ies.CauseRicrequest_CAUSE_RICREQUEST_INCONSISTENT_ACTION_SUBSEQUENT_ACTION_SEQUENCE
}

func createCauseRicControlMessageInvalid() e2ap_ies.CauseRicrequest {
	return e2ap_ies.CauseRicrequest_CAUSE_RICREQUEST_CONTROL_MESSAGE_INVALID
}

func createCauseRicProcessIDInvalid() e2ap_ies.CauseRicrequest {
	return e2ap_ies.CauseRicrequest_CAUSE_RICREQUEST_RIC_CALL_PROCESS_ID_INVALID
}

func createCauseRicUnspecified() e2ap_ies.CauseRicrequest {
	return e2ap_ies.CauseRicrequest_CAUSE_RICREQUEST_UNSPECIFIED
}

func Test_xerEncodingCauseRic(t *testing.T) {

	causeRic := createCauseRicRfIDInvalid()

	xer, err := xerEncodeCauseRic(&causeRic)
	assert.NilError(t, err)
	t.Logf("CauseRIC (RanFunction ID Invalid) XER\n%s", string(xer))

	result, err := xerDecodeCauseRic(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("CauseRIC (RanFunction ID Invalid) XER - decoded\n%v", result)
	assert.Equal(t, causeRic.Number(), result.Number())

	causeRic = createCauseRicActionNotSupported()

	xer, err = xerEncodeCauseRic(&causeRic)
	assert.NilError(t, err)
	t.Logf("CauseRIC (Action Not Supported) XER\n%s", string(xer))

	result, err = xerDecodeCauseRic(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("CauseRIC (Action Not Supported) XER - decoded\n%v", result)
	assert.Equal(t, causeRic.Number(), result.Number())

	causeRic = createCauseRicExcessiveActions()

	xer, err = xerEncodeCauseRic(&causeRic)
	assert.NilError(t, err)
	t.Logf("CauseRIC (Excessive Actions) XER\n%s", string(xer))

	result, err = xerDecodeCauseRic(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("CauseRIC (Excessive Actions) XER - decoded\n%v", result)
	assert.Equal(t, causeRic.Number(), result.Number())

	causeRic = createCauseRicDuplicateAction()

	xer, err = xerEncodeCauseRic(&causeRic)
	assert.NilError(t, err)
	t.Logf("CauseRIC (Duplicate Action) XER\n%s", string(xer))

	result, err = xerDecodeCauseRic(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("CauseRIC (Duplicate Action) XER - decoded\n%v", result)
	assert.Equal(t, causeRic.Number(), result.Number())

	causeRic = createCauseRicDuplicateEvent()

	xer, err = xerEncodeCauseRic(&causeRic)
	assert.NilError(t, err)
	t.Logf("CauseRIC (Duplicate Event) XER\n%s", string(xer))

	result, err = xerDecodeCauseRic(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("CauseRIC (Duplicate Event) XER - decoded\n%v", result)
	assert.Equal(t, causeRic.Number(), result.Number())

	causeRic = createCauseRicFunctionResourceLimit()

	xer, err = xerEncodeCauseRic(&causeRic)
	assert.NilError(t, err)
	t.Logf("CauseRIC (Function Resource Limit) XER\n%s", string(xer))

	result, err = xerDecodeCauseRic(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("CauseRIC (Function Resource Limit) XER - decoded\n%v", result)
	assert.Equal(t, causeRic.Number(), result.Number())

	causeRic = createCauseRicRequestIDUnknown()

	xer, err = xerEncodeCauseRic(&causeRic)
	assert.NilError(t, err)
	t.Logf("CauseRIC (Request ID Unknown) XER\n%s", string(xer))

	result, err = xerDecodeCauseRic(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("CauseRIC (Request ID Unknown) XER - decoded\n%v", result)
	assert.Equal(t, causeRic.Number(), result.Number())

	causeRic = createCauseRicActionSubsequentActionSequence()

	xer, err = xerEncodeCauseRic(&causeRic)
	assert.NilError(t, err)
	t.Logf("CauseRIC (Action Subsequent Action Sequence) XER\n%s", string(xer))

	result, err = xerDecodeCauseRic(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("CauseRIC (Action Subsequent Action Sequence) XER - decoded\n%v", result)
	assert.Equal(t, causeRic.Number(), result.Number())

	causeRic = createCauseRicControlMessageInvalid()

	xer, err = xerEncodeCauseRic(&causeRic)
	assert.NilError(t, err)
	t.Logf("CauseRIC (Control Message Invalid) XER\n%s", string(xer))

	result, err = xerDecodeCauseRic(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("CauseRIC (Control Message Invalid) XER - decoded\n%v", result)
	assert.Equal(t, causeRic.Number(), result.Number())

	causeRic = createCauseRicProcessIDInvalid()

	xer, err = xerEncodeCauseRic(&causeRic)
	assert.NilError(t, err)
	t.Logf("CauseRIC (Process ID Invalid) XER\n%s", string(xer))

	result, err = xerDecodeCauseRic(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("CauseRIC (Process ID Invalid) XER - decoded\n%v", result)
	assert.Equal(t, causeRic.Number(), result.Number())

	causeRic = createCauseRicUnspecified()

	xer, err = xerEncodeCauseRic(&causeRic)
	assert.NilError(t, err)
	t.Logf("CauseRIC (Unspecified) XER\n%s", string(xer))

	result, err = xerDecodeCauseRic(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("CauseRIC (Unspecified) XER - decoded\n%v", result)
	assert.Equal(t, causeRic.Number(), result.Number())
}

func Test_perEncodingCauseRic(t *testing.T) {

	causeRic := createCauseRicRfIDInvalid()

	per, err := perEncodeCauseRic(&causeRic)
	assert.NilError(t, err)
	t.Logf("CauseRIC (RanFunction ID Invalid) PER\n%v", hex.Dump(per))

	result, err := perDecodeCauseRic(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("CauseRIC (RanFunction ID Invalid) PER - decoded\n%v", result)
	assert.Equal(t, causeRic.Number(), result.Number())

	causeRic = createCauseRicActionNotSupported()

	per, err = perEncodeCauseRic(&causeRic)
	assert.NilError(t, err)
	t.Logf("CauseRIC (Action Not Supported) PER\n%v", hex.Dump(per))

	result, err = perDecodeCauseRic(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("CauseRIC (Action Not Supported) PER - decoded\n%v", result)
	assert.Equal(t, causeRic.Number(), result.Number())

	causeRic = createCauseRicExcessiveActions()

	per, err = perEncodeCauseRic(&causeRic)
	assert.NilError(t, err)
	t.Logf("CauseRIC (Excessive Actions) PER\n%v", hex.Dump(per))

	result, err = perDecodeCauseRic(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("CauseRIC (Excessive Actions) PER - decoded\n%v", result)
	assert.Equal(t, causeRic.Number(), result.Number())

	causeRic = createCauseRicDuplicateAction()

	per, err = perEncodeCauseRic(&causeRic)
	assert.NilError(t, err)
	t.Logf("CauseRIC (Duplicate Action) PER\n%v", hex.Dump(per))

	result, err = perDecodeCauseRic(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("CauseRIC (Duplicate Action) PER - decoded\n%v", result)
	assert.Equal(t, causeRic.Number(), result.Number())

	causeRic = createCauseRicDuplicateEvent()

	per, err = perEncodeCauseRic(&causeRic)
	assert.NilError(t, err)
	t.Logf("CauseRIC (Duplicate Event) PER\n%v", hex.Dump(per))

	result, err = perDecodeCauseRic(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("CauseRIC (Duplicate Event) PER - decoded\n%v", result)
	assert.Equal(t, causeRic.Number(), result.Number())

	causeRic = createCauseRicFunctionResourceLimit()

	per, err = perEncodeCauseRic(&causeRic)
	assert.NilError(t, err)
	t.Logf("CauseRIC (Function Resource Limit) PER\n%v", hex.Dump(per))

	result, err = perDecodeCauseRic(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("CauseRIC (Function Resource Limit) PER - decoded\n%v", result)
	assert.Equal(t, causeRic.Number(), result.Number())

	causeRic = createCauseRicRequestIDUnknown()

	per, err = perEncodeCauseRic(&causeRic)
	assert.NilError(t, err)
	t.Logf("CauseRIC (Request ID Unknown) PER\n%v", hex.Dump(per))

	result, err = perDecodeCauseRic(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("CauseRIC (Request ID Unknown) PER - decoded\n%v", result)
	assert.Equal(t, causeRic.Number(), result.Number())

	causeRic = createCauseRicActionSubsequentActionSequence()

	per, err = perEncodeCauseRic(&causeRic)
	assert.NilError(t, err)
	t.Logf("CauseRIC (Action Subsequent Action Sequence) PER\n%v", hex.Dump(per))

	result, err = perDecodeCauseRic(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("CauseRIC (Action Subsequent Action Sequence) PER - decoded\n%v", result)
	assert.Equal(t, causeRic.Number(), result.Number())

	causeRic = createCauseRicControlMessageInvalid()

	per, err = perEncodeCauseRic(&causeRic)
	assert.NilError(t, err)
	t.Logf("CauseRIC (Control Message Invalid) PER\n%v", hex.Dump(per))

	result, err = perDecodeCauseRic(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("CauseRIC (Control Message Invalid) PER - decoded\n%v", result)
	assert.Equal(t, causeRic.Number(), result.Number())

	causeRic = createCauseRicProcessIDInvalid()

	per, err = perEncodeCauseRic(&causeRic)
	assert.NilError(t, err)
	t.Logf("CauseRIC (Process ID Invalid) PER\n%s", hex.Dump(per))

	result, err = perDecodeCauseRic(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("CauseRIC (Process ID Invalid) PER - decoded\n%v", result)
	assert.Equal(t, causeRic.Number(), result.Number())

	causeRic = createCauseRicUnspecified()

	per, err = perEncodeCauseRic(&causeRic)
	assert.NilError(t, err)
	t.Logf("CauseRIC (Unspecified) PER\n%s", hex.Dump(per))

	result, err = perDecodeCauseRic(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("CauseRIC (Unspecified) PER - decoded\n%v", result)
	assert.Equal(t, causeRic.Number(), result.Number())
}
