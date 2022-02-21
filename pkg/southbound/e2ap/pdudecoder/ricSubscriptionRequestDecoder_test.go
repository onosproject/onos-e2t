// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdudecoder

import (
	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/pdubuilder"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
	"gotest.tools/assert"
	"testing"
)

func Test_DecodeRicSubscriptionRequestPdu(t *testing.T) {
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

	rsr, err := pdubuilder.NewRicSubscriptionRequest(
		types.RicRequest{RequestorID: 1, InstanceID: 2},
		3, []byte{0x55, 0x66}, ricActionsToBeSetup)
	assert.NilError(t, err)
	assert.Assert(t, rsr != nil)

	e2apPdu, err := pdubuilder.CreateRicSubscriptionRequestE2apPdu(rsr)
	assert.NilError(t, err)
	assert.Assert(t, rsr != nil)

	ricReq, ranFuncID, ricEventDef, ricActionsToBeSetup, err := DecodeRicSubscriptionRequestPdu(e2apPdu)
	assert.NilError(t, err)
	assert.Equal(t, 1, int(ricReq.RequestorID))
	assert.Equal(t, 2, int(ricReq.InstanceID))
	assert.Equal(t, 3, int(ranFuncID))
	assert.Equal(t, 2, len(ricEventDef))
	assert.Equal(t, 2, len(ricActionsToBeSetup))

	assert.Equal(t, 100, int(ricActionsToBeSetup[0].RicActionID))
	assert.Equal(t, e2apies.RictimeToWait_RICTIME_TO_WAIT_W5MS, ricActionsToBeSetup[0].Ricttw)
	assert.Equal(t, e2apies.RicsubsequentActionType_RICSUBSEQUENT_ACTION_TYPE_CONTINUE, ricActionsToBeSetup[0].RicSubsequentAction)
}
