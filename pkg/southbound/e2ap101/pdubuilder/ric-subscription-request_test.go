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

	newE2apPdu, err := CreateRicSubscriptionRequestE2apPdu(
		types.RicRequest{RequestorID: 2, InstanceID: 1},
		3, []byte{0x08, 0x03, 0xe7},
		//ToDo - absence of RICactionsToBeSetup list is the reason why PER encoding is crushing
		nil) //ricActionsToBeSetup)
	//[]byte{0x5c, 0x30, 0x31, 0x30, 0x5c, 0x30, 0x30, 0x33, 0x5c, 0x33, 0x34, 0x37}

	assert.NilError(t, err)
	assert.Assert(t, newE2apPdu != nil)
	t.Logf("Hexdump of bytes is \n%v", hex.Dump([]byte("\010\003\347")))
	t.Logf("Message is \n%v", newE2apPdu.GetInitiatingMessage().GetProcedureCode().GetRicSubscription().GetInitiatingMessage())

	xer, err := asn1cgo.XerEncodeE2apPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("RicSubscriptionRequest E2AP PDU\n%s", xer)

	per, err := asn1cgo.PerEncodeE2apPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("RicSubscriptionRequest E2AP PDU\n%v", hex.Dump(per))
}
