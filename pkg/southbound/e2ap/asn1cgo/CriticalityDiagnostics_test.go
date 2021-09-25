// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

import (
	"encoding/hex"
	"testing"

	"github.com/onosproject/onos-e2t/api/e2ap/v2beta1"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-commondatatypes"
	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/pdubuilder"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
	"gotest.tools/assert"
)

func Test_CriticalityDiagnostics(t *testing.T) {
	procCode := v2beta1.ProcedureCodeIDRICsubscription
	criticality := e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE
	ftg := e2ap_commondatatypes.TriggeringMessage_TRIGGERING_MESSAGE_UNSUCCESSFUL_OUTCOME
	newE2apPdu, err := pdubuilder.CreateRicSubscriptionDeleteFailureE2apPdu(&types.RicRequest{
		RequestorID: 22,
		InstanceID:  6,
	}, 9,
		&e2apies.Cause{
			Cause: &e2apies.Cause_Transport{
				Transport: e2apies.CauseTransport_CAUSE_TRANSPORT_TRANSPORT_RESOURCE_UNAVAILABLE,
			},
		})
	assert.NilError(t, err)
	assert.Assert(t, newE2apPdu != nil)

	newE2apPdu.GetUnsuccessfulOutcome().GetProcedureCode().GetRicSubscriptionDelete().GetUnsuccessfulOutcome().SetCriticalityDiagnostics(&procCode, &criticality, &ftg,
		&types.RicRequest{
			RequestorID: 10,
			InstanceID:  20,
		}, []*types.CritDiag{
			{
				TypeOfError:   e2apies.TypeOfError_TYPE_OF_ERROR_MISSING,
				IECriticality: e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE,
				IEId:          v2beta1.ProtocolIeIDRicsubscriptionDetails,
			},
		})

	critDiagsTestC, err := newCriticalityDiagnostics(newE2apPdu.GetUnsuccessfulOutcome().GetProcedureCode().GetRicSubscriptionDelete().GetUnsuccessfulOutcome().GetProtocolIes().GetE2ApProtocolIes2().GetValue())
	assert.NilError(t, err)
	assert.Assert(t, critDiagsTestC != nil)

	critDiagsReversed, err := decodeCriticalityDiagnostics(critDiagsTestC)
	assert.NilError(t, err)
	assert.Assert(t, critDiagsReversed != nil)

	xer, err := xerEncodeCriticalityDiagnostics(newE2apPdu.GetUnsuccessfulOutcome().GetProcedureCode().GetRicSubscriptionDelete().GetUnsuccessfulOutcome().GetProtocolIes().GetE2ApProtocolIes2().GetValue())
	assert.NilError(t, err)
	t.Logf("CriticalityDiagnostics XER\n%s", xer)

	per, err := perEncodeCriticalityDiagnostics(newE2apPdu.GetUnsuccessfulOutcome().GetProcedureCode().GetRicSubscriptionDelete().GetUnsuccessfulOutcome().GetProtocolIes().GetE2ApProtocolIes2().GetValue())
	assert.NilError(t, err)
	t.Logf("CriticalityDiagnostics PER\n%s", hex.Dump(per))

	// Now reverse the XER
	cdReversed, err := xerDecodeCriticalityDiagnostics(xer)
	assert.NilError(t, err)
	assert.Assert(t, cdReversed != nil)
	t.Logf("CriticalityDiagnostics decoded from XER is \n%v", cdReversed)
	//assert.Equal(t, 2, len(rflReversed.GetValue()))
	assert.Equal(t, newE2apPdu.GetUnsuccessfulOutcome().GetProcedureCode().GetRicSubscriptionDelete().GetUnsuccessfulOutcome().GetProtocolIes().GetE2ApProtocolIes2().GetValue().GetRicRequestorId().GetRicInstanceId(), cdReversed.GetRicRequestorId().GetRicInstanceId())
	assert.Equal(t, newE2apPdu.GetUnsuccessfulOutcome().GetProcedureCode().GetRicSubscriptionDelete().GetUnsuccessfulOutcome().GetProtocolIes().GetE2ApProtocolIes2().GetValue().GetRicRequestorId().GetRicRequestorId(), cdReversed.GetRicRequestorId().GetRicRequestorId())
	assert.Equal(t, newE2apPdu.GetUnsuccessfulOutcome().GetProcedureCode().GetRicSubscriptionDelete().GetUnsuccessfulOutcome().GetProtocolIes().GetE2ApProtocolIes2().GetValue().GetIEsCriticalityDiagnostics().GetValue()[0].GetIEId().GetValue(), cdReversed.GetIEsCriticalityDiagnostics().GetValue()[0].GetIEId().GetValue())
	assert.Equal(t, newE2apPdu.GetUnsuccessfulOutcome().GetProcedureCode().GetRicSubscriptionDelete().GetUnsuccessfulOutcome().GetProtocolIes().GetE2ApProtocolIes2().GetValue().GetIEsCriticalityDiagnostics().GetValue()[0].GetIEcriticality().Number(), cdReversed.GetIEsCriticalityDiagnostics().GetValue()[0].GetIEcriticality().Number())
	assert.Equal(t, newE2apPdu.GetUnsuccessfulOutcome().GetProcedureCode().GetRicSubscriptionDelete().GetUnsuccessfulOutcome().GetProtocolIes().GetE2ApProtocolIes2().GetValue().GetIEsCriticalityDiagnostics().GetValue()[0].GetTypeOfError().Number(), cdReversed.GetIEsCriticalityDiagnostics().GetValue()[0].GetTypeOfError().Number())

	// Now reverse the PER
	cdReversedFromPer, err := perDecodeCriticalityDiagnostics(per)
	assert.NilError(t, err)
	assert.Assert(t, cdReversedFromPer != nil)
	t.Logf("CriticalityDiagnostics decoded from PER is \n%v", cdReversedFromPer)
	//assert.Equal(t, 2, len(rflReversedFromPer.GetValue()))
	assert.Equal(t, newE2apPdu.GetUnsuccessfulOutcome().GetProcedureCode().GetRicSubscriptionDelete().GetUnsuccessfulOutcome().GetProtocolIes().GetE2ApProtocolIes2().GetValue().GetRicRequestorId().GetRicInstanceId(), cdReversed.GetRicRequestorId().GetRicInstanceId())
	assert.Equal(t, newE2apPdu.GetUnsuccessfulOutcome().GetProcedureCode().GetRicSubscriptionDelete().GetUnsuccessfulOutcome().GetProtocolIes().GetE2ApProtocolIes2().GetValue().GetRicRequestorId().GetRicRequestorId(), cdReversed.GetRicRequestorId().GetRicRequestorId())
	assert.Equal(t, newE2apPdu.GetUnsuccessfulOutcome().GetProcedureCode().GetRicSubscriptionDelete().GetUnsuccessfulOutcome().GetProtocolIes().GetE2ApProtocolIes2().GetValue().GetIEsCriticalityDiagnostics().GetValue()[0].GetIEId().GetValue(), cdReversed.GetIEsCriticalityDiagnostics().GetValue()[0].GetIEId().GetValue())
	assert.Equal(t, newE2apPdu.GetUnsuccessfulOutcome().GetProcedureCode().GetRicSubscriptionDelete().GetUnsuccessfulOutcome().GetProtocolIes().GetE2ApProtocolIes2().GetValue().GetIEsCriticalityDiagnostics().GetValue()[0].GetIEcriticality().Number(), cdReversed.GetIEsCriticalityDiagnostics().GetValue()[0].GetIEcriticality().Number())
	assert.Equal(t, newE2apPdu.GetUnsuccessfulOutcome().GetProcedureCode().GetRicSubscriptionDelete().GetUnsuccessfulOutcome().GetProtocolIes().GetE2ApProtocolIes2().GetValue().GetIEsCriticalityDiagnostics().GetValue()[0].GetTypeOfError().Number(), cdReversed.GetIEsCriticalityDiagnostics().GetValue()[0].GetTypeOfError().Number())
}
