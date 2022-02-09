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

func TestRicSubscriptionRequest(t *testing.T) {
	//ricActionsToBeSetup1 := make(map[types1.RicActionID]types1.RicActionDef)
	//ricActionsToBeSetup1[100] = types1.RicActionDef{
	//	RicActionID:         100,
	//	RicActionType:       e2ap_ies.RicactionType_RICACTION_TYPE_INSERT,
	//	RicSubsequentAction: e2ap_ies.RicsubsequentActionType_RICSUBSEQUENT_ACTION_TYPE_CONTINUE,
	//	Ricttw:              e2ap_ies.RictimeToWait_RICTIME_TO_WAIT_W5MS,
	//	RicActionDefinition: []byte{0x11, 0x22},
	//}
	//
	////ricActionsToBeSetup1[200] = types1.RicActionDef{
	////	RicActionID:         200,
	////	RicActionType:       e2ap_ies.RicactionType_RICACTION_TYPE_INSERT,
	////	RicSubsequentAction: e2ap_ies.RicsubsequentActionType_RICSUBSEQUENT_ACTION_TYPE_CONTINUE,
	////	Ricttw:              e2ap_ies.RictimeToWait_RICTIME_TO_WAIT_W10MS,
	////	RicActionDefinition: []byte{0x33, 0x44},
	////}
	//
	//rsr1, err := pdubuilder.NewRicSubscriptionRequest(
	//	types1.RicRequest{RequestorID: 1, InstanceID: 2},
	//	3, []byte{0x55, 0x66}, ricActionsToBeSetup1)
	//assert.NilError(t, err)
	//assert.Assert(t, rsr1 != nil)
	//
	//e2apPdu, err := pdubuilder.CreateRicSubscriptionRequestE2apPdu(rsr1)
	//assert.NilError(t, err)
	//assert.Assert(t, rsr1 != nil)
	//
	//per, err := asn1cgo.PerEncodeE2apPdu(e2apPdu)
	//assert.NilError(t, err)
	//t.Logf("RicSubscriptionRequest E2AP PDU PER\n%v", hex.Dump(per))

	ricActionsToBeSetup := make(map[types.RicActionID]types.RicActionDef)
	ricActionsToBeSetup[100] = types.RicActionDef{
		RicActionID:         100,
		RicActionType:       e2apies.RicactionType_RICACTION_TYPE_INSERT,
		RicSubsequentAction: e2apies.RicsubsequentActionType_RICSUBSEQUENT_ACTION_TYPE_CONTINUE,
		Ricttw:              e2apies.RictimeToWait_RICTIME_TO_WAIT_W5MS,
		RicActionDefinition: []byte{0x11, 0x22},
	}

	//ricActionsToBeSetup[200] = types.RicActionDef{
	//	RicActionID:         200,
	//	RicActionType:       e2apies.RicactionType_RICACTION_TYPE_INSERT,
	//	RicSubsequentAction: e2apies.RicsubsequentActionType_RICSUBSEQUENT_ACTION_TYPE_CONTINUE,
	//	Ricttw:              e2apies.RictimeToWait_RICTIME_TO_WAIT_W10MS,
	//	RicActionDefinition: []byte{0x33, 0x44},
	//}

	rsr, err := NewRicSubscriptionRequest(
		types.RicRequest{RequestorID: 1, InstanceID: 2},
		3, []byte{0x55, 0x66}, ricActionsToBeSetup)
	assert.NilError(t, err)
	assert.Assert(t, rsr != nil)

	newE2apPdu, err := CreateRicSubscriptionRequestE2apPdu(rsr)
	assert.NilError(t, err)
	assert.Assert(t, rsr != nil)

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
	//t.Logf("RicSubscriptionRequest E2AP PDU PER - decoded\n%v\n", result1)
	//assert.DeepEqual(t, e2apPdu.String(), result1.String())
}
