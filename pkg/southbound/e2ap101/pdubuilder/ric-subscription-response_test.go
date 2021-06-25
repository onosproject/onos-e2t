// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
package pdubuilder

import (
	"encoding/hex"
	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/asn1cgo"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/types"
	"gotest.tools/assert"
	"testing"
)

func TestRicSubscriptionResponse(t *testing.T) {
	ricActionsNotAdmittedList := make(map[types.RicActionID]*e2apies.Cause)
	ricActionsNotAdmittedList[100] = &e2apies.Cause{
		Cause: &e2apies.Cause_Transport{
			Transport: e2apies.CauseTransport_CAUSE_TRANSPORT_TRANSPORT_RESOURCE_UNAVAILABLE,
		},
	}
	ricActionsNotAdmittedList[200] = &e2apies.Cause{
		Cause: &e2apies.Cause_Misc{
			Misc: e2apies.CauseMisc_CAUSE_MISC_HARDWARE_FAILURE,
		},
	}

	var ricActionAdmitted10 types.RicActionID = 10
	var ricActionAdmitted20 types.RicActionID = 20
	newE2apPdu, err := CreateRicSubscriptionResponseE2apPdu(&types.RicRequest{
		RequestorID: 22,
		InstanceID:  6,
	}, 9, []*types.RicActionID{&ricActionAdmitted10, &ricActionAdmitted20}, ricActionsNotAdmittedList)
	assert.NilError(t, err)
	assert.Assert(t, newE2apPdu != nil)

	xer, err := asn1cgo.XerEncodeE2apPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("RicSubscriptionResponse E2AP PDU XER\n%s", string(xer))

	result, err := asn1cgo.XerDecodeE2apPdu(xer)
	assert.NilError(t, err)
	t.Logf("RicSubscriptionResponse E2AP PDU XER - decoded\n%v\n", result)
	assert.DeepEqual(t, newE2apPdu, result)

	per, err := asn1cgo.PerEncodeE2apPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("RicSubscriptionResponse E2AP PDU PER\n%v", hex.Dump(per))

	result1, err := asn1cgo.PerDecodeE2apPdu(per)
	assert.NilError(t, err)
	t.Logf("RicSubscriptionResponse E2AP PDU PER - decoded\n%v\n", result1)
	assert.DeepEqual(t, newE2apPdu, result1)
}

func TestRicSubscriptionResponseExceptOptional(t *testing.T) {
	var ricActionAdmitted10 types.RicActionID = 10
	var ricActionAdmitted20 types.RicActionID = 20
	newE2apPdu, err := CreateRicSubscriptionResponseE2apPdu(&types.RicRequest{
		RequestorID: 22,
		InstanceID:  6,
	}, 9, []*types.RicActionID{&ricActionAdmitted10, &ricActionAdmitted20}, nil)
	assert.NilError(t, err)
	assert.Assert(t, newE2apPdu != nil)

	xer, err := asn1cgo.XerEncodeE2apPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("RicSubscriptionResponse E2AP PDU XER\n%s", string(xer))

	result, err := asn1cgo.XerDecodeE2apPdu(xer)
	assert.NilError(t, err)
	t.Logf("RicSubscriptionResponse E2AP PDU XER - decoded\n%v\n", result)
	assert.DeepEqual(t, newE2apPdu, result)

	per, err := asn1cgo.PerEncodeE2apPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("RicSubscriptionResponse E2AP PDU PER\n%v", hex.Dump(per))

	result1, err := asn1cgo.PerDecodeE2apPdu(per)
	assert.NilError(t, err)
	t.Logf("RicSubscriptionResponse E2AP PDU PER - decoded\n%v\n", result1)
	assert.DeepEqual(t, newE2apPdu, result1)
}
