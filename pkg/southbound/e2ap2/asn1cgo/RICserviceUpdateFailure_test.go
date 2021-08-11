// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

import (
	"encoding/hex"
	"fmt"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta2"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-commondatatypes"
	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	e2ap_pdu_contents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/pdubuilder"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/types"
	"gotest.tools/assert"
	"testing"
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
	procCode := v1beta2.ProcedureCodeIDRICsubscription
	criticality := e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE
	ftg := e2ap_commondatatypes.TriggeringMessage_TRIGGERING_MESSAGE_UNSUCCESSFULL_OUTCOME

	rsuf, err := pdubuilder.CreateRicServiceUpdateFailureE2apPdu(rfRejected, &ttw,
		&procCode, &criticality, &ftg,
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
	if err != nil {
		return nil, err
	}

	if err := rsuf.Validate(); err != nil {
		return nil, fmt.Errorf("error validating RicServiceUpdateFailure %s", err.Error())
	}
	return rsuf.GetUnsuccessfulOutcome().GetProcedureCode().GetRicServiceUpdate().GetUnsuccessfulOutcome(), nil
}

func Test_xerEncodingRicserviceUpdateFailure(t *testing.T) {

	rsuf, err := createRicServiceUpdateFailureMsg()
	assert.NilError(t, err, "Error creating RicServiceUpdateFailure PDU")

	xer, err := xerEncodeRicServiceUpdateFailure(rsuf)
	assert.NilError(t, err)
	assert.Equal(t, 2856, len(xer))
	t.Logf("RicServiceUpdateFailure XER\n%s", string(xer))

	result, err := xerDecodeRicServiceUpdateFailure(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("RicServiceUpdateFailure XER - decoded\n%v", result)
	assert.Equal(t, rsuf.GetProtocolIes().GetE2ApProtocolIes31().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes31().GetValue())
	assert.Equal(t, rsuf.GetProtocolIes().GetE2ApProtocolIes13().GetValue().GetValue()[0].GetRanFunctionIdcauseItemIes7().GetValue().GetCause().GetRicService(), result.GetProtocolIes().GetE2ApProtocolIes13().GetValue().GetValue()[0].GetRanFunctionIdcauseItemIes7().GetValue().GetCause().GetRicService())
	assert.Equal(t, rsuf.GetProtocolIes().GetE2ApProtocolIes13().GetValue().GetValue()[0].GetRanFunctionIdcauseItemIes7().GetValue().GetRanFunctionId().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes13().GetValue().GetValue()[0].GetRanFunctionIdcauseItemIes7().GetValue().GetRanFunctionId().GetValue())
	assert.Equal(t, rsuf.GetProtocolIes().GetE2ApProtocolIes2().GetValue().GetRicRequestorId().GetRicInstanceId(), result.GetProtocolIes().GetE2ApProtocolIes2().GetValue().GetRicRequestorId().GetRicInstanceId())
	assert.Equal(t, rsuf.GetProtocolIes().GetE2ApProtocolIes2().GetValue().GetRicRequestorId().GetRicRequestorId(), result.GetProtocolIes().GetE2ApProtocolIes2().GetValue().GetRicRequestorId().GetRicRequestorId())
	assert.Equal(t, rsuf.GetProtocolIes().GetE2ApProtocolIes2().GetValue().GetTriggeringMessage(), result.GetProtocolIes().GetE2ApProtocolIes2().GetValue().GetTriggeringMessage())
	assert.Equal(t, rsuf.GetProtocolIes().GetE2ApProtocolIes2().GetValue().GetIEsCriticalityDiagnostics().GetValue()[0].GetIEcriticality(), result.GetProtocolIes().GetE2ApProtocolIes2().GetValue().GetIEsCriticalityDiagnostics().GetValue()[0].GetIEcriticality())
	assert.Equal(t, rsuf.GetProtocolIes().GetE2ApProtocolIes2().GetValue().GetIEsCriticalityDiagnostics().GetValue()[0].GetIEId().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes2().GetValue().GetIEsCriticalityDiagnostics().GetValue()[0].GetIEId().GetValue())
	assert.Equal(t, rsuf.GetProtocolIes().GetE2ApProtocolIes2().GetValue().GetIEsCriticalityDiagnostics().GetValue()[0].GetTypeOfError(), result.GetProtocolIes().GetE2ApProtocolIes2().GetValue().GetIEsCriticalityDiagnostics().GetValue()[0].GetTypeOfError())
}

func Test_perEncodingRicserviceUpdateFailure(t *testing.T) {

	rsuf, err := createRicServiceUpdateFailureMsg()
	assert.NilError(t, err, "Error creating RicServiceUpdateFailure PDU")

	per, err := perEncodeRicServiceUpdateFailure(rsuf)
	assert.NilError(t, err)
	assert.Equal(t, 46, len(per))
	t.Logf("RicServiceUpdateFailure PER\n%v", hex.Dump(per))

	result, err := perDecodeRicServiceUpdateFailure(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("RicServiceUpdateFailure PER - decoded\n%v", result)
	assert.Equal(t, rsuf.GetProtocolIes().GetE2ApProtocolIes31().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes31().GetValue())
	assert.Equal(t, rsuf.GetProtocolIes().GetE2ApProtocolIes13().GetValue().GetValue()[0].GetRanFunctionIdcauseItemIes7().GetValue().GetCause().GetRicService(), result.GetProtocolIes().GetE2ApProtocolIes13().GetValue().GetValue()[0].GetRanFunctionIdcauseItemIes7().GetValue().GetCause().GetRicService())
	assert.Equal(t, rsuf.GetProtocolIes().GetE2ApProtocolIes13().GetValue().GetValue()[0].GetRanFunctionIdcauseItemIes7().GetValue().GetRanFunctionId().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes13().GetValue().GetValue()[0].GetRanFunctionIdcauseItemIes7().GetValue().GetRanFunctionId().GetValue())
	assert.Equal(t, rsuf.GetProtocolIes().GetE2ApProtocolIes2().GetValue().GetRicRequestorId().GetRicInstanceId(), result.GetProtocolIes().GetE2ApProtocolIes2().GetValue().GetRicRequestorId().GetRicInstanceId())
	assert.Equal(t, rsuf.GetProtocolIes().GetE2ApProtocolIes2().GetValue().GetRicRequestorId().GetRicRequestorId(), result.GetProtocolIes().GetE2ApProtocolIes2().GetValue().GetRicRequestorId().GetRicRequestorId())
	assert.Equal(t, rsuf.GetProtocolIes().GetE2ApProtocolIes2().GetValue().GetTriggeringMessage(), result.GetProtocolIes().GetE2ApProtocolIes2().GetValue().GetTriggeringMessage())
	assert.Equal(t, rsuf.GetProtocolIes().GetE2ApProtocolIes2().GetValue().GetIEsCriticalityDiagnostics().GetValue()[0].GetIEcriticality(), result.GetProtocolIes().GetE2ApProtocolIes2().GetValue().GetIEsCriticalityDiagnostics().GetValue()[0].GetIEcriticality())
	assert.Equal(t, rsuf.GetProtocolIes().GetE2ApProtocolIes2().GetValue().GetIEsCriticalityDiagnostics().GetValue()[0].GetIEId().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes2().GetValue().GetIEsCriticalityDiagnostics().GetValue()[0].GetIEId().GetValue())
	assert.Equal(t, rsuf.GetProtocolIes().GetE2ApProtocolIes2().GetValue().GetIEsCriticalityDiagnostics().GetValue()[0].GetTypeOfError(), result.GetProtocolIes().GetE2ApProtocolIes2().GetValue().GetIEsCriticalityDiagnostics().GetValue()[0].GetTypeOfError())
}
