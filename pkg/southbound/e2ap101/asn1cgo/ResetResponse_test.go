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

func createResetResponseMsg() (*e2ap_pdu_contents.ResetResponse, error) {

	resetResponse, err := pdubuilder.CreateResetResponseE2apPdu(
		v1beta2.ProcedureCodeIDRICsubscription, e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE,
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
	if err != nil {
		return nil, err
	}

	if err := resetResponse.Validate(); err != nil {
		return nil, fmt.Errorf("error validating ResetResponse %s", err.Error())
	}
	return resetResponse.GetSuccessfulOutcome().GetProcedureCode().GetReset_().GetSuccessfulOutcome(), nil
}

func Test_xerEncodingResetResponse(t *testing.T) {

	resetResponse, err := createResetResponseMsg()
	assert.NilError(t, err, "Error creating ResetResponse PDU")

	xer, err := xerEncodeResetResponse(resetResponse)
	assert.NilError(t, err)
	assert.Equal(t, 1119, len(xer))
	t.Logf("ResetResponse XER\n%s", string(xer))

	result, err := xerDecodeResetResponse(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("ResetResponse XER - decoded\n%v", result)
	assert.Equal(t, resetResponse.GetProtocolIes().GetResetResponseIes2().GetValue().GetRicRequestorId().GetRicInstanceId(), result.GetProtocolIes().GetResetResponseIes2().GetValue().GetRicRequestorId().GetRicInstanceId())
	assert.Equal(t, resetResponse.GetProtocolIes().GetResetResponseIes2().GetValue().GetRicRequestorId().GetRicRequestorId(), result.GetProtocolIes().GetResetResponseIes2().GetValue().GetRicRequestorId().GetRicRequestorId())
	assert.Equal(t, resetResponse.GetProtocolIes().GetResetResponseIes2().GetValue().GetTriggeringMessage(), result.GetProtocolIes().GetResetResponseIes2().GetValue().GetTriggeringMessage())
	assert.Equal(t, resetResponse.GetProtocolIes().GetResetResponseIes2().GetValue().GetProcedureCriticality(), result.GetProtocolIes().GetResetResponseIes2().GetValue().GetProcedureCriticality())
	assert.Equal(t, resetResponse.GetProtocolIes().GetResetResponseIes2().GetValue().GetIEsCriticalityDiagnostics().GetValue()[0].GetIEcriticality(), result.GetProtocolIes().GetResetResponseIes2().GetValue().GetIEsCriticalityDiagnostics().GetValue()[0].GetIEcriticality())
	assert.Equal(t, resetResponse.GetProtocolIes().GetResetResponseIes2().GetValue().GetIEsCriticalityDiagnostics().GetValue()[0].GetTypeOfError(), result.GetProtocolIes().GetResetResponseIes2().GetValue().GetIEsCriticalityDiagnostics().GetValue()[0].GetTypeOfError())

}

func Test_perEncodingResetResponse(t *testing.T) {

	resetResponse, err := createResetResponseMsg()
	assert.NilError(t, err, "Error creating ResetResponse PDU")

	per, err := perEncodeResetResponse(resetResponse)
	assert.NilError(t, err)
	assert.Equal(t, 19, len(per))
	t.Logf("ResetResponse PER\n%v", hex.Dump(per))

	result, err := perDecodeResetResponse(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("ResetResponse PER - decoded\n%v", result)
	assert.Equal(t, resetResponse.GetProtocolIes().GetResetResponseIes2().GetValue().GetRicRequestorId().GetRicInstanceId(), result.GetProtocolIes().GetResetResponseIes2().GetValue().GetRicRequestorId().GetRicInstanceId())
	assert.Equal(t, resetResponse.GetProtocolIes().GetResetResponseIes2().GetValue().GetRicRequestorId().GetRicRequestorId(), result.GetProtocolIes().GetResetResponseIes2().GetValue().GetRicRequestorId().GetRicRequestorId())
	assert.Equal(t, resetResponse.GetProtocolIes().GetResetResponseIes2().GetValue().GetTriggeringMessage(), result.GetProtocolIes().GetResetResponseIes2().GetValue().GetTriggeringMessage())
	assert.Equal(t, resetResponse.GetProtocolIes().GetResetResponseIes2().GetValue().GetProcedureCriticality(), result.GetProtocolIes().GetResetResponseIes2().GetValue().GetProcedureCriticality())
	assert.Equal(t, resetResponse.GetProtocolIes().GetResetResponseIes2().GetValue().GetIEsCriticalityDiagnostics().GetValue()[0].GetIEcriticality(), result.GetProtocolIes().GetResetResponseIes2().GetValue().GetIEsCriticalityDiagnostics().GetValue()[0].GetIEcriticality())
	assert.Equal(t, resetResponse.GetProtocolIes().GetResetResponseIes2().GetValue().GetIEsCriticalityDiagnostics().GetValue()[0].GetTypeOfError(), result.GetProtocolIes().GetResetResponseIes2().GetValue().GetIEsCriticalityDiagnostics().GetValue()[0].GetTypeOfError())
}
