// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

import (
	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/pdubuilder"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/types"
	"gotest.tools/assert"
	"testing"
)

func Test_xerEncodeRICsubscriptionRequest(t *testing.T) {
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

	e2ApPduRsr, err := pdubuilder.CreateRicSubscriptionRequestE2apPdu(
		types.RicRequest{RequestorID: 1, InstanceID: 2},
		3, []byte{0x55, 0x66}, ricActionsToBeSetup)

	assert.NilError(t, err)
	xer, err := xerEncodeRICsubscriptionRequest(
		e2ApPduRsr.GetInitiatingMessage().GetProcedureCode().GetRicSubscription().GetInitiatingMessage())
	assert.NilError(t, err)
	t.Logf("XER RICsubscriptionRequest\n%s", xer)

	per, err := perEncodeRICsubscriptionRequest(
		e2ApPduRsr.GetInitiatingMessage().GetProcedureCode().GetRicSubscription().GetInitiatingMessage())
	assert.NilError(t, err)
	t.Logf("PER RICsubscriptionRequest\n%s", per)
}
