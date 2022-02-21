// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdudecoder

import (
	v2 "github.com/onosproject/onos-e2t/api/e2ap/v2"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-commondatatypes"
	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/pdubuilder"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
	"gotest.tools/assert"
	"testing"
)

func Test_DecodeRicSubscriptionFailurePdu(t *testing.T) {
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

	procCode := v2.ProcedureCodeIDRICsubscription
	criticality := e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE
	ftg := e2ap_commondatatypes.TriggeringMessage_TRIGGERING_MESSAGE_UNSUCCESSFUL_OUTCOME

	e2apPdu, err := pdubuilder.CreateRicSubscriptionFailureE2apPdu(&types.RicRequest{
		RequestorID: 22,
		InstanceID:  6,
	}, 9, &e2apies.Cause{
		Cause: &e2apies.Cause_Misc{
			Misc: e2apies.CauseMisc_CAUSE_MISC_CONTROL_PROCESSING_OVERLOAD,
		},
	})
	assert.NilError(t, err)
	assert.Assert(t, e2apPdu != nil)
	e2apPdu.GetUnsuccessfulOutcome().GetValue().GetRicSubscription().
		SetCriticalityDiagnostics(procCode, &criticality, &ftg,
			&types.RicRequest{
				RequestorID: 10,
				InstanceID:  20,
			}, []*types.CritDiag{
				{
					TypeOfError:   e2apies.TypeOfError_TYPE_OF_ERROR_MISSING,
					IECriticality: e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE,
					IEId:          v2.ProtocolIeIDRicsubscriptionDetails,
				},
			})

	rrID, rfID, pc, crit, tm, critReq, cause, diags, err := DecodeRicSubscriptionFailurePdu(e2apPdu)
	assert.NilError(t, err)
	//assert.Assert(t, rfID != nil) //Commented due to the Linters (v1.34.1) error - possible nil pointer dereference (https://staticcheck.io/docs/checks#SA5011) on line 26
	assert.Equal(t, 9, int(*rfID))

	//assert.Assert(t, rrID != nil) //Commented due to the Linters (v1.34.1) error - possible nil pointer dereference (https://staticcheck.io/docs/checks#SA5011) on lines 29 & 30
	assert.Equal(t, 22, int(rrID.RequestorID))
	assert.Equal(t, 6, int(rrID.InstanceID))

	assert.Equal(t, v2.ProcedureCodeIDRICsubscription, pc)
	assert.Equal(t, e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE, crit)
	assert.Equal(t, e2ap_commondatatypes.TriggeringMessage_TRIGGERING_MESSAGE_UNSUCCESSFUL_OUTCOME, tm)

	assert.Equal(t, 10, int(critReq.RequestorID))
	assert.Equal(t, 20, int(critReq.InstanceID))

	assert.Equal(t, e2apies.CauseMisc_CAUSE_MISC_CONTROL_PROCESSING_OVERLOAD.Number(), cause.GetMisc().Number())
	assert.Assert(t, diags != nil)
}
