// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package asn1cgo

import (
	"encoding/hex"
	"testing"

	"github.com/onosproject/onos-e2t/api/e2ap/v2beta1"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-commondatatypes"
	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-ies"
	e2ap_pdu_contents "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-pdu-contents"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/pdubuilder"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
	"gotest.tools/assert"
)

func createRicServiceUpdateFailureMsg() (*e2ap_pdu_contents.RicserviceUpdateFailure, error) {
	rfRejected := make(types.RanFunctionCauses)
	rfRejected[101] = &e2apies.Cause{
		Cause: &e2apies.Cause_Misc{
			Misc: e2apies.CauseMisc_CAUSE_MISC_HARDWARE_FAILURE,
		},
	}
	rfRejected[102] = &e2apies.Cause{
		Cause: &e2apies.Cause_Protocol{
			Protocol: e2apies.CauseProtocol_CAUSE_PROTOCOL_SEMANTIC_ERROR,
		},
	}

	ttw := e2apies.TimeToWait_TIME_TO_WAIT_V2S
	procCode := v2beta1.ProcedureCodeIDRICsubscription
	criticality := e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE
	ftg := e2ap_commondatatypes.TriggeringMessage_TRIGGERING_MESSAGE_UNSUCCESSFUL_OUTCOME

	rsuf, err := pdubuilder.CreateRicServiceUpdateFailureE2apPdu(1, &e2apies.Cause{
		Cause: &e2apies.Cause_RicService{
			RicService: e2apies.CauseRicservice_CAUSE_RICSERVICE_RIC_RESOURCE_LIMIT,
		},
	})
	if err != nil {
		return nil, err
	}

	rsuf.GetUnsuccessfulOutcome().GetProcedureCode().GetRicServiceUpdate().GetUnsuccessfulOutcome().
		SetTimeToWait(ttw).SetCriticalityDiagnostics(&procCode, &criticality, &ftg,
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

	//if err := rsuf.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating RicServiceUpdateFailure %s", err.Error())
	//}
	return rsuf.GetUnsuccessfulOutcome().GetProcedureCode().GetRicServiceUpdate().GetUnsuccessfulOutcome(), nil
}

func Test_xerEncodingRicserviceUpdateFailure(t *testing.T) {

	rsuf, err := createRicServiceUpdateFailureMsg()
	assert.NilError(t, err, "Error creating RicServiceUpdateFailure PDU")

	xer, err := xerEncodeRicServiceUpdateFailure(rsuf)
	assert.NilError(t, err)
	t.Logf("RicServiceUpdateFailure XER\n%s", string(xer))

	result, err := xerDecodeRicServiceUpdateFailure(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("RicServiceUpdateFailure XER - decoded\n%v", result)
	assert.Equal(t, rsuf.GetProtocolIes().GetE2ApProtocolIes31().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes31().GetValue())
	assert.Equal(t, rsuf.GetProtocolIes().GetE2ApProtocolIes2().GetValue().GetRicRequestorId().GetRicInstanceId(), result.GetProtocolIes().GetE2ApProtocolIes2().GetValue().GetRicRequestorId().GetRicInstanceId())
	assert.Equal(t, rsuf.GetProtocolIes().GetE2ApProtocolIes2().GetValue().GetRicRequestorId().GetRicRequestorId(), result.GetProtocolIes().GetE2ApProtocolIes2().GetValue().GetRicRequestorId().GetRicRequestorId())
	assert.Equal(t, rsuf.GetProtocolIes().GetE2ApProtocolIes2().GetValue().GetTriggeringMessage(), result.GetProtocolIes().GetE2ApProtocolIes2().GetValue().GetTriggeringMessage())
	assert.Equal(t, rsuf.GetProtocolIes().GetE2ApProtocolIes2().GetValue().GetIEsCriticalityDiagnostics().GetValue()[0].GetIEcriticality(), result.GetProtocolIes().GetE2ApProtocolIes2().GetValue().GetIEsCriticalityDiagnostics().GetValue()[0].GetIEcriticality())
	assert.Equal(t, rsuf.GetProtocolIes().GetE2ApProtocolIes2().GetValue().GetIEsCriticalityDiagnostics().GetValue()[0].GetIEId().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes2().GetValue().GetIEsCriticalityDiagnostics().GetValue()[0].GetIEId().GetValue())
	assert.Equal(t, rsuf.GetProtocolIes().GetE2ApProtocolIes2().GetValue().GetIEsCriticalityDiagnostics().GetValue()[0].GetTypeOfError(), result.GetProtocolIes().GetE2ApProtocolIes2().GetValue().GetIEsCriticalityDiagnostics().GetValue()[0].GetTypeOfError())
	assert.Equal(t, rsuf.GetProtocolIes().GetE2ApProtocolIes49().GetValue().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes49().GetValue().GetValue())
	assert.Equal(t, rsuf.GetProtocolIes().GetE2ApProtocolIes1().GetValue().GetRicService().Number(), result.GetProtocolIes().GetE2ApProtocolIes1().GetValue().GetRicService().Number())
}

func Test_perEncodingRicserviceUpdateFailure(t *testing.T) {

	rsuf, err := createRicServiceUpdateFailureMsg()
	assert.NilError(t, err, "Error creating RicServiceUpdateFailure PDU")

	per, err := perEncodeRicServiceUpdateFailure(rsuf)
	assert.NilError(t, err)
	t.Logf("RicServiceUpdateFailure PER\n%v", hex.Dump(per))

	result, err := perDecodeRicServiceUpdateFailure(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("RicServiceUpdateFailure PER - decoded\n%v", result)
	assert.Equal(t, rsuf.GetProtocolIes().GetE2ApProtocolIes31().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes31().GetValue())
	assert.Equal(t, rsuf.GetProtocolIes().GetE2ApProtocolIes2().GetValue().GetRicRequestorId().GetRicInstanceId(), result.GetProtocolIes().GetE2ApProtocolIes2().GetValue().GetRicRequestorId().GetRicInstanceId())
	assert.Equal(t, rsuf.GetProtocolIes().GetE2ApProtocolIes2().GetValue().GetRicRequestorId().GetRicRequestorId(), result.GetProtocolIes().GetE2ApProtocolIes2().GetValue().GetRicRequestorId().GetRicRequestorId())
	assert.Equal(t, rsuf.GetProtocolIes().GetE2ApProtocolIes2().GetValue().GetTriggeringMessage(), result.GetProtocolIes().GetE2ApProtocolIes2().GetValue().GetTriggeringMessage())
	assert.Equal(t, rsuf.GetProtocolIes().GetE2ApProtocolIes2().GetValue().GetIEsCriticalityDiagnostics().GetValue()[0].GetIEcriticality(), result.GetProtocolIes().GetE2ApProtocolIes2().GetValue().GetIEsCriticalityDiagnostics().GetValue()[0].GetIEcriticality())
	assert.Equal(t, rsuf.GetProtocolIes().GetE2ApProtocolIes2().GetValue().GetIEsCriticalityDiagnostics().GetValue()[0].GetIEId().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes2().GetValue().GetIEsCriticalityDiagnostics().GetValue()[0].GetIEId().GetValue())
	assert.Equal(t, rsuf.GetProtocolIes().GetE2ApProtocolIes2().GetValue().GetIEsCriticalityDiagnostics().GetValue()[0].GetTypeOfError(), result.GetProtocolIes().GetE2ApProtocolIes2().GetValue().GetIEsCriticalityDiagnostics().GetValue()[0].GetTypeOfError())
	assert.Equal(t, rsuf.GetProtocolIes().GetE2ApProtocolIes49().GetValue().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes49().GetValue().GetValue())
	assert.Equal(t, rsuf.GetProtocolIes().GetE2ApProtocolIes1().GetValue().GetRicService().Number(), result.GetProtocolIes().GetE2ApProtocolIes1().GetValue().GetRicService().Number())
}
