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

func createCauseProtocolSyntaxError() e2ap_ies.CauseProtocol {
	return e2ap_ies.CauseProtocol_CAUSE_PROTOCOL_TRANSFER_SYNTAX_ERROR
}

func createCauseProtocolSyntaxErrorReject() e2ap_ies.CauseProtocol {
	return e2ap_ies.CauseProtocol_CAUSE_PROTOCOL_ABSTRACT_SYNTAX_ERROR_REJECT
}

func createCauseProtocolSyntaxErrorIgnore() e2ap_ies.CauseProtocol {
	return e2ap_ies.CauseProtocol_CAUSE_PROTOCOL_ABSTRACT_SYNTAX_ERROR_IGNORE_AND_NOTIFY
}

func createCauseProtocolMsgNotCompatible() e2ap_ies.CauseProtocol {
	return e2ap_ies.CauseProtocol_CAUSE_PROTOCOL_MESSAGE_NOT_COMPATIBLE_WITH_RECEIVER_STATE
}

func createCauseProtocolSemanticError() e2ap_ies.CauseProtocol {
	return e2ap_ies.CauseProtocol_CAUSE_PROTOCOL_SEMANTIC_ERROR
}

func createCauseProtocolFalselyMessage() e2ap_ies.CauseProtocol {
	return e2ap_ies.CauseProtocol_CAUSE_PROTOCOL_ABSTRACT_SYNTAX_ERROR_FALSELY_CONSTRUCTED_MESSAGE
}

func createCauseProtocolUspecified() e2ap_ies.CauseProtocol {
	return e2ap_ies.CauseProtocol_CAUSE_PROTOCOL_UNSPECIFIED
}

func Test_xerEncodingCauseProtocol(t *testing.T) {

	causeProtocol := createCauseProtocolSyntaxError()

	xer, err := xerEncodeCauseProtocol(&causeProtocol)
	assert.NilError(t, err)
	assert.Equal(t, 56, len(xer))
	t.Logf("CauseProtocol (Syntax Error) XER\n%s", string(xer))

	result, err := xerDecodeCauseProtocol(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("CauseProtocol (Syntax Error) XER - decoded\n%v", result)
	if result != nil {
		assert.Equal(t, causeProtocol, *result)
	}

	causeProtocol = createCauseProtocolSyntaxErrorReject()

	xer, err = xerEncodeCauseProtocol(&causeProtocol)
	assert.NilError(t, err)
	assert.Equal(t, 63, len(xer))
	t.Logf("CauseProtocol (Syntax Error Reject) XER\n%s", string(xer))

	result, err = xerDecodeCauseProtocol(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("CauseProtocol (Syntax Error Reject) XER - decoded\n%v", result)
	if result != nil {
		assert.Equal(t, causeProtocol, *result)
	}

	causeProtocol = createCauseProtocolSyntaxErrorIgnore()

	xer, err = xerEncodeCauseProtocol(&causeProtocol)
	assert.NilError(t, err)
	assert.Equal(t, 74, len(xer))
	t.Logf("CauseProtocol (Syntax Error Ignore) XER\n%s", string(xer))

	result, err = xerDecodeCauseProtocol(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("CauseProtocol (Syntax Error Ignore) XER - decoded\n%v", result)
	if result != nil {
		assert.Equal(t, causeProtocol, *result)
	}

	causeProtocol = createCauseProtocolMsgNotCompatible()

	xer, err = xerEncodeCauseProtocol(&causeProtocol)
	assert.NilError(t, err)
	assert.Equal(t, 77, len(xer))
	t.Logf("CauseProtocol (Message Not Compatible) XER\n%s", string(xer))

	result, err = xerDecodeCauseProtocol(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("CauseProtocol (Message Not Compatible) XER - decoded\n%v", result)
	if result != nil {
		assert.Equal(t, causeProtocol, *result)
	}

	causeProtocol = createCauseProtocolSemanticError()

	xer, err = xerEncodeCauseProtocol(&causeProtocol)
	assert.NilError(t, err)
	assert.Equal(t, 49, len(xer))
	t.Logf("CauseProtocol (Semantic Error) XER\n%s", string(xer))

	result, err = xerDecodeCauseProtocol(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("CauseProtocol (Semantic Error) XER - decoded\n%v", result)
	if result != nil {
		assert.Equal(t, causeProtocol, *result)
	}

	causeProtocol = createCauseProtocolFalselyMessage()

	xer, err = xerEncodeCauseProtocol(&causeProtocol)
	assert.NilError(t, err)
	assert.Equal(t, 84, len(xer))
	t.Logf("CauseProtocol (Falesely Constructed Message) XER\n%s", string(xer))

	result, err = xerDecodeCauseProtocol(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("CauseProtocol (Falesely Constructed Message) XER - decoded\n%v", result)
	if result != nil {
		assert.Equal(t, causeProtocol, *result)
	}

	causeProtocol = createCauseProtocolUspecified()

	xer, err = xerEncodeCauseProtocol(&causeProtocol)
	assert.NilError(t, err)
	assert.Equal(t, 46, len(xer))
	t.Logf("CauseProtocol (Unspecified) XER\n%s", string(xer))

	result, err = xerDecodeCauseProtocol(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("CauseProtocol (Unspecified) XER - decoded\n%v", result)
	if result != nil {
		assert.Equal(t, causeProtocol, *result)
	}
}

func Test_perEncodingCauseProtocol(t *testing.T) {

	causeProtocol := createCauseProtocolSyntaxError()

	per, err := perEncodeCauseProtocol(&causeProtocol)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(per))
	t.Logf("CauseProtocol (Syntax Error) PER\n%v", hex.Dump(per))

	result, err := perDecodeCauseProtocol(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("CauseProtocol (Syntax Error) PER - decoded\n%v", result)
	if result != nil {
		assert.Equal(t, causeProtocol, *result)
	}

	causeProtocol = createCauseProtocolSyntaxErrorReject()

	per, err = perEncodeCauseProtocol(&causeProtocol)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(per))
	t.Logf("CauseProtocol (Syntax Error Reject) PER\n%v", hex.Dump(per))

	result, err = perDecodeCauseProtocol(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("CauseProtocol (Syntax Error Reject) PER - decoded\n%v", result)
	if result != nil {
		assert.Equal(t, causeProtocol, *result)
	}

	causeProtocol = createCauseProtocolSyntaxErrorIgnore()

	per, err = perEncodeCauseProtocol(&causeProtocol)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(per))
	t.Logf("CauseProtocol (Syntax Error Ignore) PER\n%v", hex.Dump(per))

	result, err = perDecodeCauseProtocol(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("CauseProtocol (Syntax Error Ignore) PER - decoded\n%v", result)
	if result != nil {
		assert.Equal(t, causeProtocol, *result)
	}

	causeProtocol = createCauseProtocolMsgNotCompatible()

	per, err = perEncodeCauseProtocol(&causeProtocol)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(per))
	t.Logf("CauseProtocol (Message Not Compatible) PER\n%v", hex.Dump(per))

	result, err = perDecodeCauseProtocol(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("CauseProtocol (Message Not Compatible) PER - decoded\n%v", result)
	if result != nil {
		assert.Equal(t, causeProtocol, *result)
	}

	causeProtocol = createCauseProtocolSemanticError()

	per, err = perEncodeCauseProtocol(&causeProtocol)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(per))
	t.Logf("CauseProtocol (Semantic Error) PER\n%v", hex.Dump(per))

	result, err = perDecodeCauseProtocol(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("CauseProtocol (Semantic Error) PER - decoded\n%v", result)
	if result != nil {
		assert.Equal(t, causeProtocol, *result)
	}

	causeProtocol = createCauseProtocolFalselyMessage()

	per, err = perEncodeCauseProtocol(&causeProtocol)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(per))
	t.Logf("CauseProtocol (Falesely Constructed Message) PER\n%v", hex.Dump(per))

	result, err = perDecodeCauseProtocol(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("CauseProtocol (Falesely Constructed Message) PER - decoded\n%v", result)
	if result != nil {
		assert.Equal(t, causeProtocol, *result)
	}

	causeProtocol = createCauseProtocolUspecified()

	per, err = perEncodeCauseProtocol(&causeProtocol)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(per))
	t.Logf("CauseProtocol (Unspecified) PER\n%v", hex.Dump(per))

	result, err = perDecodeCauseProtocol(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("CauseProtocol (Unspecified) PER - decoded\n%v", result)
	if result != nil {
		assert.Equal(t, causeProtocol, *result)
	}
}
