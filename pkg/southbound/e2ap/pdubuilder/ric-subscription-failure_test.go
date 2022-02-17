// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
package pdubuilder

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/encoder"
	"testing"

	"github.com/onosproject/onos-e2t/api/e2ap/v2"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-commondatatypes"
	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
	"gotest.tools/assert"
)

func TestRicSubscriptionFailure(t *testing.T) {
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

	newE2apPdu, err := CreateRicSubscriptionFailureE2apPdu(&types.RicRequest{
		RequestorID: 22,
		InstanceID:  6,
	}, 9, &e2apies.Cause{
		Cause: &e2apies.Cause_Misc{
			Misc: e2apies.CauseMisc_CAUSE_MISC_CONTROL_PROCESSING_OVERLOAD,
		},
	})
	assert.NilError(t, err)
	assert.Assert(t, newE2apPdu != nil)
	newE2apPdu.GetUnsuccessfulOutcome().GetValue().GetRicSubscription().
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

	perNew, err := encoder.PerEncodeE2ApPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("RicSubscriptionFailure E2AP PDU PER with Go APER library\n%v", hex.Dump(perNew))

	//Comparing reference PER bytes with Go APER library produced
	//assert.DeepEqual(t, per, perNew)

	result, err := encoder.PerDecodeE2ApPdu(perNew)
	assert.NilError(t, err)
	assert.DeepEqual(t, newE2apPdu.String(), result.String())
}

func TestRicSubscriptionFailureExcludeOptionalIE(t *testing.T) {
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

	newE2apPdu, err := CreateRicSubscriptionFailureE2apPdu(&types.RicRequest{
		RequestorID: 22,
		InstanceID:  6,
	}, 9,
		&e2apies.Cause{
			Cause: &e2apies.Cause_Misc{
				Misc: e2apies.CauseMisc_CAUSE_MISC_CONTROL_PROCESSING_OVERLOAD,
			},
		})
	assert.NilError(t, err)
	assert.Assert(t, newE2apPdu != nil)

	perNew, err := encoder.PerEncodeE2ApPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("RicSubscriptionFailure E2AP PDU PER with Go APER library\n%v", hex.Dump(perNew))

	//Comparing reference PER bytes with Go APER library produced
	//assert.DeepEqual(t, per, perNew)

	result, err := encoder.PerDecodeE2ApPdu(perNew)
	assert.NilError(t, err)
	assert.DeepEqual(t, newE2apPdu.String(), result.String())
}
