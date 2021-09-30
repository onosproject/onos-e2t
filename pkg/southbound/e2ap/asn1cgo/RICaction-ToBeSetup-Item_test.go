// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package asn1cgo

import (
	"encoding/hex"
	"testing"

	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-ies"
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-pdu-contents"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/pdubuilder"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
	"gotest.tools/assert"
)

func createRicActionToBeSetupItem() (*e2appducontents.RicactionToBeSetupItem, error) {

	ricActionsToBeSetup := make(map[types.RicActionID]types.RicActionDef)
	ricActionsToBeSetup[200] = types.RicActionDef{
		RicActionID:   200,
		RicActionType: e2apies.RicactionType_RICACTION_TYPE_INSERT,
		//RicSubsequentAction: e2apies.RicsubsequentActionType_RICSUBSEQUENT_ACTION_TYPE_CONTINUE,
		//Ricttw:              e2apies.RictimeToWait_RICTIME_TO_WAIT_W10MS,
		//RicActionDefinition: []byte{0x33, 0x44},
	}
	rsr, err := pdubuilder.NewRicSubscriptionRequest(
		types.RicRequest{RequestorID: 1, InstanceID: 2},
		3, []byte{0x55, 0x66}, ricActionsToBeSetup)
	if err != nil {
		return nil, err
	}

	res := rsr.GetProtocolIes().GetE2ApProtocolIes30().GetValue().GetRicActionToBeSetupList().GetValue()[0].GetValue()
	//fmt.Printf("Returning following structure: \n %v \n", res)

	return res, nil
}

func Test_RICactionToBeSetupItem(t *testing.T) {

	ratbsi, err := createRicActionToBeSetupItem()
	assert.NilError(t, err)

	xer, err := xerEncodeRicActionToBeSetupItem(ratbsi)
	assert.NilError(t, err)
	t.Logf("RanFunctionList XER\n%s", xer)

	per, err := perEncodeRicActionToBeSetupItem(ratbsi)
	assert.NilError(t, err)
	t.Logf("RanFunctionList PER\n%v", hex.Dump(per))

	// Now reverse the XER
	ratbsiReversed, err := xerDecodeRicActionToBeSetupItem(xer)
	assert.NilError(t, err)
	assert.Assert(t, ratbsiReversed != nil)
	t.Logf("RanFunctionList decoded from XER is \n%v", ratbsiReversed)
	//assert.Equal(t, 2, len(rflReversed.GetValue()))

	// Now reverse the PER
	ratbsiReversedFromPer, err := perDecodeRicActionToBeSetupItem(per)
	assert.NilError(t, err)
	assert.Assert(t, ratbsiReversedFromPer != nil)
	t.Logf("RanFunctionList decoded from PER is \n%v", ratbsiReversedFromPer)
	//assert.Equal(t, 2, len(rflReversedFromPer.GetValue()))

}
