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

func createE2connectionUpdateFailureMsg() (*e2ap_pdu_contents.E2ConnectionUpdateFailure, error) {

	ttw := e2apies.TimeToWait_TIME_TO_WAIT_V5S
	procCode := v2beta1.ProcedureCodeIDRICsubscription
	criticality := e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE
	ftg := e2ap_commondatatypes.TriggeringMessage_TRIGGERING_MESSAGE_UNSUCCESSFUL_OUTCOME

	e2connectionUpdateFailure, err := pdubuilder.CreateE2connectionUpdateFailureE2apPdu(1)
	if err != nil {
		return nil, err
	}

	e2connectionUpdateFailure.GetUnsuccessfulOutcome().GetProcedureCode().GetE2ConnectionUpdate().GetUnsuccessfulOutcome().
		SetCause(&e2apies.Cause{
			Cause: &e2apies.Cause_Protocol{
				Protocol: e2apies.CauseProtocol_CAUSE_PROTOCOL_TRANSFER_SYNTAX_ERROR,
			},
		}).SetTimeToWait(ttw).SetCriticalityDiagnostics(&procCode, &criticality, &ftg,
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

	//if err := e2connectionUpdateFailure.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating E2connectionUpdateFailure %s", err.Error())
	//}
	return e2connectionUpdateFailure.GetUnsuccessfulOutcome().GetProcedureCode().GetE2ConnectionUpdate().GetUnsuccessfulOutcome(), nil
}

func Test_xerEncodingE2connectionUpdateFailure(t *testing.T) {

	e2connectionUpdateFailure, err := createE2connectionUpdateFailureMsg()
	assert.NilError(t, err, "Error creating E2connectionUpdateFailure PDU")

	xer, err := xerEncodeE2connectionUpdateFailure(e2connectionUpdateFailure)
	assert.NilError(t, err)
	t.Logf("E2connectionUpdateFailure XER\n%s", string(xer))

	result, err := xerDecodeE2connectionUpdateFailure(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2connectionUpdateFailure XER - decoded\n%v", result)
	assert.Equal(t, e2connectionUpdateFailure.GetProtocolIes().GetE2ApProtocolIes1().GetValue().GetProtocol(), result.GetProtocolIes().GetE2ApProtocolIes1().GetValue().GetProtocol())
	assert.Equal(t, e2connectionUpdateFailure.GetProtocolIes().GetE2ApProtocolIes2().GetValue().GetRicRequestorId().GetRicInstanceId(), result.GetProtocolIes().GetE2ApProtocolIes2().GetValue().GetRicRequestorId().GetRicInstanceId())
	assert.Equal(t, e2connectionUpdateFailure.GetProtocolIes().GetE2ApProtocolIes2().GetValue().GetRicRequestorId().GetRicRequestorId(), result.GetProtocolIes().GetE2ApProtocolIes2().GetValue().GetRicRequestorId().GetRicRequestorId())
	assert.Equal(t, e2connectionUpdateFailure.GetProtocolIes().GetE2ApProtocolIes31().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes31().GetValue())
	assert.Equal(t, e2connectionUpdateFailure.GetProtocolIes().GetE2ApProtocolIes49().GetValue().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes49().GetValue().GetValue())
}

func Test_perEncodingE2connectionUpdateFailure(t *testing.T) {

	e2connectionUpdateFailure, err := createE2connectionUpdateFailureMsg()
	assert.NilError(t, err, "Error creating E2connectionUpdateFailure PDU")

	per, err := perEncodeE2connectionUpdateFailure(e2connectionUpdateFailure)
	assert.NilError(t, err)
	t.Logf("E2connectionUpdateFailure PER\n%v", hex.Dump(per))

	result, err := perDecodeE2connectionUpdateFailure(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2connectionUpdateFailure PER - decoded\n%v", result)
	assert.Equal(t, e2connectionUpdateFailure.GetProtocolIes().GetE2ApProtocolIes1().GetValue().GetProtocol(), result.GetProtocolIes().GetE2ApProtocolIes1().GetValue().GetProtocol())
	assert.Equal(t, e2connectionUpdateFailure.GetProtocolIes().GetE2ApProtocolIes2().GetValue().GetRicRequestorId().GetRicInstanceId(), result.GetProtocolIes().GetE2ApProtocolIes2().GetValue().GetRicRequestorId().GetRicInstanceId())
	assert.Equal(t, e2connectionUpdateFailure.GetProtocolIes().GetE2ApProtocolIes2().GetValue().GetRicRequestorId().GetRicRequestorId(), result.GetProtocolIes().GetE2ApProtocolIes2().GetValue().GetRicRequestorId().GetRicRequestorId())
	assert.Equal(t, e2connectionUpdateFailure.GetProtocolIes().GetE2ApProtocolIes31().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes31().GetValue())
	assert.Equal(t, e2connectionUpdateFailure.GetProtocolIes().GetE2ApProtocolIes49().GetValue().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes49().GetValue().GetValue())
}
