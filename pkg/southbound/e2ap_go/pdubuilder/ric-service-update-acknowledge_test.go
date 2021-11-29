// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0
package pdubuilder

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap_go/encoder"
	"testing"

	e2apies "github.com/onosproject/onos-e2t/api/e2ap_go/v2/e2ap-ies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap_go/types"
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
	newE2apPdu.GetSuccessfulOutcome().GetValue().GetRicServiceUpdate().
		SetRanFunctionsRejected(rfRejected)

	perNew, err := encoder.PerEncodeE2ApPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("RicIndication E2AP PDU PER with Go APER library\n%v", hex.Dump(perNew))

	//Comparing reference PER bytes with Go APER library produced
	//assert.DeepEqual(t, per, perNew)

	e2apPdu, err := encoder.PerDecodeE2ApPdu(perNew)
	assert.NilError(t, err)
	assert.DeepEqual(t, newE2apPdu.String(), e2apPdu.String())

	//per, err := asn1cgo.PerEncodeE2apPdu(newE2apPdu)
	//assert.NilError(t, err)
	//t.Logf("RicServiceUpdateAcknowledge E2AP PDU PER\n%v", hex.Dump(per))
	//
	//result1, err := asn1cgo.PerDecodeE2apPdu(per)
	//assert.NilError(t, err)
	//assert.Assert(t, result1 != nil)
	//t.Logf("RicServiceUpdateAcknowledge E2AP PDU PER - decoded is \n%v", result1)
	//assert.DeepEqual(t, newE2apPdu.String(), result1.String())
}

func TestRicServiceUpdateAcknowledgeExcludeOptionalIE(t *testing.T) {
	rfAccepted := make(types.RanFunctionRevisions)
	rfAccepted[100] = 2
	rfAccepted[200] = 2

	newE2apPdu, err := CreateRicServiceUpdateAcknowledgeE2apPdu(3, rfAccepted)
	assert.NilError(t, err)
	assert.Assert(t, newE2apPdu != nil)

	perNew, err := encoder.PerEncodeE2ApPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("RicIndication E2AP PDU PER with Go APER library\n%v", hex.Dump(perNew))

	//Comparing reference PER bytes with Go APER library produced
	//assert.DeepEqual(t, per, perNew)

	e2apPdu, err := encoder.PerDecodeE2ApPdu(perNew)
	assert.NilError(t, err)
	assert.DeepEqual(t, newE2apPdu.String(), e2apPdu.String())

	//per, err := asn1cgo.PerEncodeE2apPdu(newE2apPdu)
	//assert.NilError(t, err)
	//t.Logf("RicServiceUpdateAcknowledge E2AP PDU PER\n%v", hex.Dump(per))
	//
	//result1, err := asn1cgo.PerDecodeE2apPdu(per)
	//assert.NilError(t, err)
	//assert.Assert(t, result1 != nil)
	//t.Logf("RicServiceUpdateAcknowledge E2AP PDU PER - decoded is \n%v", result1)
	//assert.DeepEqual(t, newE2apPdu.String(), result1.String())
}
