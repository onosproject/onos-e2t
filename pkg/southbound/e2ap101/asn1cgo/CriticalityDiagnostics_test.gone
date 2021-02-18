// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

import (
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta2"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-commondatatypes"
	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/pdubuilder"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
	"gotest.tools/assert"
	"testing"
)

func Test_CriticalityDiagnostics(t *testing.T) {
	newE2apPdu, err := pdubuilder.CreateRicSubscriptionDeleteFailureE2apPdu(&types.RicRequest{
		RequestorID: 22,
		InstanceID:  6,
	}, 9,
		&e2apies.Cause{
			Cause: &e2apies.Cause_Transport{
				Transport: e2apies.CauseTransport_CAUSE_TRANSPORT_TRANSPORT_RESOURCE_UNAVAILABLE,
			},
		}, v1beta2.ProcedureCodeIDRICsubscription, e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE,
		e2ap_commondatatypes.TriggeringMessage_TRIGGERING_MESSAGE_UNSUCCESSFULL_OUTCOME,
		&types.RicRequest{
			RequestorID: 10,
			InstanceID:  20,
		}, []*types.CritDiag{
			{
				TypeOfError:   e2apies.TypeOfError_TYPE_OF_ERROR_MISSING,
				IECriticality: e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE,
				IEId:          v1beta2.ProtocolIeIDRicsubscriptionDetails,
			},
		},
	)
	assert.NilError(t, err)
	assert.Assert(t, newE2apPdu != nil)

	critDiagsTestC, err := newCriticalityDiagnostics(newE2apPdu.GetUnsuccessfulOutcome().GetProcedureCode().GetRicSubscriptionDelete().GetUnsuccessfulOutcome().GetProtocolIes().GetE2ApProtocolIes2().GetValue())
	assert.NilError(t, err)
	assert.Assert(t, critDiagsTestC != nil)

	critDiagsReversed, err := decodeCriticalityDiagnostics(critDiagsTestC)
	assert.NilError(t, err)
	assert.Assert(t, critDiagsReversed != nil)
}
