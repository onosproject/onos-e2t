// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
package pdubuilder

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/encoder"
	"testing"

	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
	"gotest.tools/assert"
)

func TestRicSubscriptionResponse(t *testing.T) {
	//ricActionsNotAdmittedList1 := make(map[types1.RicActionID]*e2ap_ies.Cause)
	//ricActionsNotAdmittedList1[100] = &e2ap_ies.Cause{
	//	Cause: &e2ap_ies.Cause_Transport{
	//		Transport: e2ap_ies.CauseTransport_CAUSE_TRANSPORT_TRANSPORT_RESOURCE_UNAVAILABLE,
	//	},
	//}
	////ricActionsNotAdmittedList1[200] = &e2ap_ies.Cause{
	////	Cause: &e2ap_ies.Cause_Misc{
	////		Misc: e2ap_ies.CauseMisc_CAUSE_MISC_HARDWARE_FAILURE,
	////	},
	////}
	//
	//var ricActionAdmitted110 types1.RicActionID = 10
	//var ricActionAdmitted120 types1.RicActionID = 20
	//e2apPdu, err := pdubuilder.CreateRicSubscriptionResponseE2apPdu(&types1.RicRequest{
	//	RequestorID: 22,
	//	InstanceID:  6,
	//}, 9, []*types1.RicActionID{&ricActionAdmitted110, &ricActionAdmitted120})
	//assert.NilError(t, err)
	//assert.Assert(t, e2apPdu != nil)
	//e2apPdu.GetSuccessfulOutcome().GetProcedureCode().GetRicSubscription().GetSuccessfulOutcome().
	//	SetRicActionNotAdmitted(ricActionsNotAdmittedList1)
	//
	//per, err := asn1cgo.PerEncodeE2apPdu(e2apPdu)
	//assert.NilError(t, err)
	//t.Logf("RicSubscriptionRequest E2AP PDU PER\n%v", hex.Dump(per))

	ricActionsNotAdmittedList := make(map[types.RicActionID]*e2apies.Cause)
	ricActionsNotAdmittedList[100] = &e2apies.Cause{
		Cause: &e2apies.Cause_Transport{
			Transport: e2apies.CauseTransport_CAUSE_TRANSPORT_TRANSPORT_RESOURCE_UNAVAILABLE,
		},
	}
	//ricActionsNotAdmittedList[200] = &e2apies.Cause{
	//	Cause: &e2apies.Cause_Misc{
	//		Misc: e2apies.CauseMisc_CAUSE_MISC_HARDWARE_FAILURE,
	//	},
	//}

	var ricActionAdmitted10 types.RicActionID = 10
	var ricActionAdmitted20 types.RicActionID = 20
	newE2apPdu, err := CreateRicSubscriptionResponseE2apPdu(&types.RicRequest{
		RequestorID: 22,
		InstanceID:  6,
	}, 9, []*types.RicActionID{&ricActionAdmitted10, &ricActionAdmitted20})
	assert.NilError(t, err)
	assert.Assert(t, newE2apPdu != nil)
	newE2apPdu.GetSuccessfulOutcome().GetValue().GetRicSubscription().
		SetRicActionNotAdmitted(ricActionsNotAdmittedList)

	perNew, err := encoder.PerEncodeE2ApPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("RicSubscriptionRequest E2AP PDU PER with Go APER library\n%v", hex.Dump(perNew))

	//Comparing reference PER bytes with Go APER library produced
	//assert.DeepEqual(t, per, perNew)

	result, err := encoder.PerDecodeE2ApPdu(perNew)
	assert.NilError(t, err)
	assert.DeepEqual(t, newE2apPdu.String(), result.String())

	// Decoding the message from the APER bytes produced by CGo
	//result11, err := encoder.PerDecodeE2ApPdu(per)
	//assert.NilError(t, err)
	//assert.DeepEqual(t, newE2apPdu.String(), result11.String())
	//
	//result1, err := asn1cgo.PerDecodeE2apPdu(perNew)
	//assert.NilError(t, err)
	//t.Logf("RicSubscriptionResponse E2AP PDU PER - decoded\n%v\n", result1)
	//assert.DeepEqual(t, e2apPdu.String(), result1.String())
}

func TestRicSubscriptionResponseExceptOptionalIE(t *testing.T) {
	//var ricActionAdmitted110 types1.RicActionID = 10
	//var ricActionAdmitted120 types1.RicActionID = 20
	//e2apPdu, err := pdubuilder.CreateRicSubscriptionResponseE2apPdu(&types1.RicRequest{
	//	RequestorID: 22,
	//	InstanceID:  6,
	//}, 9, []*types1.RicActionID{&ricActionAdmitted110, &ricActionAdmitted120})
	//assert.NilError(t, err)
	//assert.Assert(t, e2apPdu != nil)
	//
	//per, err := asn1cgo.PerEncodeE2apPdu(e2apPdu)
	//assert.NilError(t, err)
	//t.Logf("RicSubscriptionRequest E2AP PDU PER\n%v", hex.Dump(per))

	var ricActionAdmitted10 types.RicActionID = 10
	var ricActionAdmitted20 types.RicActionID = 20
	newE2apPdu, err := CreateRicSubscriptionResponseE2apPdu(&types.RicRequest{
		RequestorID: 22,
		InstanceID:  6,
	}, 9, []*types.RicActionID{&ricActionAdmitted10, &ricActionAdmitted20})
	assert.NilError(t, err)
	assert.Assert(t, newE2apPdu != nil)

	perNew, err := encoder.PerEncodeE2ApPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("RicSubscriptionRequest E2AP PDU PER with Go APER library\n%v", hex.Dump(perNew))

	//Comparing reference PER bytes with Go APER library produced
	//assert.DeepEqual(t, per, perNew)

	result, err := encoder.PerDecodeE2ApPdu(perNew)
	assert.NilError(t, err)
	assert.DeepEqual(t, newE2apPdu.String(), result.String())

	// Decoding the message from the APER bytes produced by CGo
	//result11, err := encoder.PerDecodeE2ApPdu(per)
	//assert.NilError(t, err)
	//assert.DeepEqual(t, newE2apPdu.String(), result11.String())
	//
	//result1, err := asn1cgo.PerDecodeE2apPdu(perNew)
	//assert.NilError(t, err)
	//t.Logf("RicSubscriptionResponse E2AP PDU PER - decoded\n%v\n", result1)
	//assert.DeepEqual(t, e2apPdu.String(), result1.String())
}
