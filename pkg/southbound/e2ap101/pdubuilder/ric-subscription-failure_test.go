// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
package pdubuilder

import (
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta2"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-commondatatypes"
	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/asn1cgo"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/types"
	"gotest.tools/assert"
	"testing"
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

	newE2apPdu, err := CreateRicSubscriptionFailureE2apPdu(&types.RicRequest{
		RequestorID: 22,
		InstanceID:  6,
	}, 9,
		v1beta2.ProcedureCodeIDRICsubscription, e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE,
		e2ap_commondatatypes.TriggeringMessage_TRIGGERING_MESSAGE_UNSUCCESSFULL_OUTCOME,
		&types.RicRequest{
			RequestorID: 10,
			InstanceID:  20,
		}, ricActionsNotAdmittedList,
		[]*types.CritDiag{
			{
				TypeOfError:   e2apies.TypeOfError_TYPE_OF_ERROR_MISSING,
				IECriticality: e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE,
				IEId:          v1beta2.ProtocolIeIDRicsubscriptionDetails,
			},
		},
	)
	assert.NilError(t, err)
	assert.Assert(t, newE2apPdu != nil)

	xer, err := asn1cgo.XerEncodeE2apPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("RicSubscriptionDeleteFailure E2AP PDU XER\n%s", string(xer))

	per, err := asn1cgo.PerEncodeE2apPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("RicSubscriptionDeleteFailure E2AP PDU PER\n%v", per)
}
