// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
package pdubuilder

import (
	"encoding/hex"
	"testing"

	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-ies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2apv201/asn1cgo"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2apv201/types"
	"gotest.tools/assert"
)

func TestRicServiceUpdateAcknowledge(t *testing.T) {
	rfAccepted := make(types.RanFunctionRevisions)
	rfAccepted[100] = 2
	rfAccepted[200] = 2

	rfRejected := make(types.RanFunctionCauses)
	rfRejected[101] = &e2apies.Cause{
		Cause: &e2apies.Cause_Misc{
			Misc: e2apies.CauseMisc_CAUSE_MISC_HARDWARE_FAILURE,
		},
	}
	rfRejected[102] = &e2apies.Cause{
		Cause: &e2apies.Cause_Protocol{
			Protocol: e2apies.CauseProtocol_CAUSE_PROTOCOL_SEMANTIC_ERROR,
		},
	}

	newE2apPdu, err := CreateRicServiceUpdateAcknowledgeE2apPdu(1, rfAccepted)
	assert.NilError(t, err)
	assert.Assert(t, newE2apPdu != nil)
	newE2apPdu.GetSuccessfulOutcome().GetProcedureCode().GetRicServiceUpdate().GetSuccessfulOutcome().
		SetRanFunctionsRejected(rfRejected)

	xer, err := asn1cgo.XerEncodeE2apPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("RicServiceUpdateAcknowledge E2AP PDU XER\n%s", string(xer))

	result, err := asn1cgo.XerDecodeE2apPdu(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("RicServiceUpdateAcknowledge E2AP PDU XER - decoded is \n%v", result)
	assert.DeepEqual(t, newE2apPdu.String(), result.String())

	per, err := asn1cgo.PerEncodeE2apPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("RicServiceUpdateAcknowledge E2AP PDU PER\n%v", hex.Dump(per))

	result1, err := asn1cgo.PerDecodeE2apPdu(per)
	assert.NilError(t, err)
	assert.Assert(t, result1 != nil)
	t.Logf("RicServiceUpdateAcknowledge E2AP PDU PER - decoded is \n%v", result1)
	assert.DeepEqual(t, newE2apPdu.String(), result1.String())
}

func TestRicServiceUpdateAcknowledgeExcludeOptionalIE(t *testing.T) {
	rfAccepted := make(types.RanFunctionRevisions)
	rfAccepted[100] = 2
	rfAccepted[200] = 2

	newE2apPdu, err := CreateRicServiceUpdateAcknowledgeE2apPdu(3, rfAccepted)
	assert.NilError(t, err)
	assert.Assert(t, newE2apPdu != nil)

	xer, err := asn1cgo.XerEncodeE2apPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("RicServiceUpdateAcknowledge E2AP PDU XER\n%s", string(xer))

	result, err := asn1cgo.XerDecodeE2apPdu(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("RicServiceUpdateAcknowledge E2AP PDU XER - decoded is \n%v", result)
	assert.DeepEqual(t, newE2apPdu.String(), result.String())

	per, err := asn1cgo.PerEncodeE2apPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("RicServiceUpdateAcknowledge E2AP PDU PER\n%v", hex.Dump(per))

	result1, err := asn1cgo.PerDecodeE2apPdu(per)
	assert.NilError(t, err)
	assert.Assert(t, result1 != nil)
	t.Logf("RicServiceUpdateAcknowledge E2AP PDU PER - decoded is \n%v", result1)
	assert.DeepEqual(t, newE2apPdu.String(), result1.String())
}
