// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

import (
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2apies"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2appdudescriptions"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/pdubuilder"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
	"gotest.tools/assert"
	"testing"
)

func Test_RicActionsToBeSetupList(t *testing.T) {
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

	ricSubscriptionRequest, err := pdubuilder.CreateRicSubscriptionRequestE2apPdu(types.RicRequest{RequestorID: 1, InstanceID: 2},
		3, []byte{0x55, 0x66}, ricActionsToBeSetup)
	assert.NilError(t, err)
	assert.Assert(t, ricSubscriptionRequest != nil)

	im := ricSubscriptionRequest.GetE2ApPdu().(*e2appdudescriptions.E2ApPdu_InitiatingMessage)
	ricSubDetails := im.InitiatingMessage.GetProcedureCode().GetRicSubscription().GetInitiatingMessage().GetProtocolIes().GetE2ApProtocolIes30().GetValue()

	xer, err := xerEncodeRicActionsToBeSetupList(ricSubDetails.GetRicActionToBeSetupList())
	assert.NilError(t, err)
	t.Logf("RicActionToBeSetupList XER\n%s", xer)

	per, err := perEncodeRicActionsToBeSetupList(ricSubDetails.GetRicActionToBeSetupList())
	assert.NilError(t, err)
	t.Logf("RicActionToBeSetupList PER\n%s", per)

	// Now reverse it
	ratbsL, err := xerDecodeRicActionsToBeSetupList(xer)
	assert.NilError(t, err)
	assert.Equal(t, 2, len(ratbsL.GetValue()))
	raTbsItem0 := ratbsL.GetValue()[0]
	assert.Equal(t, int32(v1beta1.ProtocolIeIDRicactionToBeSetupItem), raTbsItem0.GetId())
	assert.Equal(t, 100, int(raTbsItem0.GetValue().GetRicActionId().GetValue()))
	assert.Equal(t, e2apies.RicactionType_RICACTION_TYPE_INSERT, raTbsItem0.GetValue().GetRicActionType())
	assert.Equal(t, e2apies.RicsubsequentActionType_RICSUBSEQUENT_ACTION_TYPE_CONTINUE, raTbsItem0.GetValue().GetRicSubsequentAction().GetRicSubsequentActionType())
	assert.Equal(t, e2apies.RictimeToWait_RICTIME_TO_WAIT_W5MS, raTbsItem0.GetValue().GetRicSubsequentAction().GetRicTimeToWait())

	raTbsItem1 := ratbsL.GetValue()[1]
	assert.Equal(t, int32(v1beta1.ProtocolIeIDRicactionToBeSetupItem), raTbsItem1.GetId())
	assert.Equal(t, 200, int(raTbsItem1.GetValue().GetRicActionId().GetValue()))
	assert.Equal(t, e2apies.RicactionType_RICACTION_TYPE_INSERT, raTbsItem1.GetValue().GetRicActionType())
	assert.Equal(t, e2apies.RicsubsequentActionType_RICSUBSEQUENT_ACTION_TYPE_CONTINUE, raTbsItem1.GetValue().GetRicSubsequentAction().GetRicSubsequentActionType())
	assert.Equal(t, e2apies.RictimeToWait_RICTIME_TO_WAIT_W10MS, raTbsItem1.GetValue().GetRicSubsequentAction().GetRicTimeToWait())
}
