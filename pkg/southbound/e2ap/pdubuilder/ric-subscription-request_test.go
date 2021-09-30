// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0
package pdubuilder

import (
	"encoding/hex"
	"testing"

	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-ies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/asn1cgo"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
	"gotest.tools/assert"
)

func TestRicSubscriptionRequest(t *testing.T) {
	ricActionsToBeSetup := make(map[types.RicActionID]types.RicActionDef)
	ricActionsToBeSetup[100] = types.RicActionDef{
		RicActionID:         100,
		RicActionType:       e2apies.RicactionType_RICACTION_TYPE_INSERT,
		RicSubsequentAction: e2apies.RicsubsequentActionType_RICSUBSEQUENT_ACTION_TYPE_CONTINUE,
		Ricttw:              e2apies.RictimeToWait_RICTIME_TO_WAIT_W5MS,
		RicActionDefinition: []byte{0x11, 0x22},
	}

	ricActionsToBeSetup[200] = types.RicActionDef{
		RicActionID:         200,
		RicActionType:       e2apies.RicactionType_RICACTION_TYPE_INSERT,
		RicSubsequentAction: e2apies.RicsubsequentActionType_RICSUBSEQUENT_ACTION_TYPE_CONTINUE,
		Ricttw:              e2apies.RictimeToWait_RICTIME_TO_WAIT_W10MS,
		RicActionDefinition: []byte{0x33, 0x44},
	}

	rsr, err := NewRicSubscriptionRequest(
		types.RicRequest{RequestorID: 1, InstanceID: 2},
		3, []byte{0x55, 0x66}, ricActionsToBeSetup)
	assert.NilError(t, err)
	assert.Assert(t, rsr != nil)

	newE2apPdu, err := CreateRicSubscriptionRequestE2apPdu(rsr)
	assert.NilError(t, err)
	assert.Assert(t, rsr != nil)

	xer, err := asn1cgo.XerEncodeE2apPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("RicSubscriptionRequest E2AP PDU\n%s", xer)

	result, err := asn1cgo.XerDecodeE2apPdu(xer)
	assert.NilError(t, err)
	t.Logf("RicSubscriptionRequest E2AP PDU XER - decoded\n%v\n", result)
	assert.DeepEqual(t, newE2apPdu.String(), result.String())

	per, err := asn1cgo.PerEncodeE2apPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("RicSubscriptionRequest E2AP PDU PER\n%v", hex.Dump(per))

	result1, err := asn1cgo.PerDecodeE2apPdu(per)
	assert.NilError(t, err)
	t.Logf("RicSubscriptionRequest E2AP PDU PER - decoded\n%v\n", result1)
	assert.DeepEqual(t, newE2apPdu.String(), result1.String())
}
