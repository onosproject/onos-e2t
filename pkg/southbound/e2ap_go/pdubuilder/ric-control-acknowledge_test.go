// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0
package pdubuilder

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/asn1cgo"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/pdubuilder"
	types1 "github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap_go/encoder"
	"testing"

	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap_go/types"
	"gotest.tools/assert"
)

func TestRicControlAcknowledge(t *testing.T) {
	ricRequestID1 := types1.RicRequest{
		RequestorID: 21,
		InstanceID:  22,
	}
	var ranFuncID1 types1.RanFunctionID = 9
	var ricCallPrID1 types1.RicCallProcessID = []byte("123")
	var ricCtrlOut1 types1.RicControlOutcome = []byte("456")
	e2apPdu, err := pdubuilder.CreateRicControlAcknowledgeE2apPdu(ricRequestID1,
		ranFuncID1)
	assert.NilError(t, err)
	assert.Assert(t, e2apPdu != nil)
	e2apPdu.GetSuccessfulOutcome().GetProcedureCode().GetRicControl().GetSuccessfulOutcome().
		SetRicControlOutcome(ricCtrlOut1).SetRicCallProcessID(ricCallPrID1)

	per, err := asn1cgo.PerEncodeE2apPdu(e2apPdu)
	assert.NilError(t, err)
	t.Logf("RicControlAcknowledge E2AP PDU PER\n%v", hex.Dump(per))

	ricRequestID := types.RicRequest{
		RequestorID: 21,
		InstanceID:  22,
	}
	var ranFuncID types.RanFunctionID = 9
	var ricCallPrID types.RicCallProcessID = []byte("123")
	var ricCtrlOut types.RicControlOutcome = []byte("456")
	newE2apPdu, err := CreateRicControlAcknowledgeE2apPdu(ricRequestID,
		ranFuncID)
	assert.NilError(t, err)
	assert.Assert(t, newE2apPdu != nil)
	newE2apPdu.GetSuccessfulOutcome().GetValue().GetRicControl().
		SetRicControlOutcome(ricCtrlOut).SetRicCallProcessID(ricCallPrID)

	perNew, err := encoder.PerEncodeE2ApPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("RicControlAcknowledge E2AP PDU PER with Go APER library\n%v", hex.Dump(perNew))

	//Comparing reference PER bytes with Go APER library produced
	assert.DeepEqual(t, per, perNew)

	result, err := encoder.PerDecodeE2ApPdu(perNew)
	assert.NilError(t, err)
	assert.DeepEqual(t, newE2apPdu.String(), result.String())

	// Decoding the message from the APER bytes produced by CGo
	result11, err := encoder.PerDecodeE2ApPdu(per)
	assert.NilError(t, err)
	assert.DeepEqual(t, newE2apPdu.String(), result11.String())

	result1, err := asn1cgo.PerDecodeE2apPdu(perNew)
	assert.NilError(t, err)
	assert.DeepEqual(t, result1.String(), e2apPdu.String())
}
