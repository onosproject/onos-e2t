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

func createE2nodeConfigurationUpdateFailureMsg() (*e2ap_pdu_contents.E2NodeConfigurationUpdateFailure, error) {

	e2nodeConfigurationUpdateFailure, err := pdubuilder.CreateE2NodeConfigurationUpdateFailureE2apPdu(
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
	if e2nodeConfigurationUpdateFailure == nil {
		return nil, fmt.Errorf("returned structure is nil")
	}

	fmt.Printf("Cause is \n%v\n", e2nodeConfigurationUpdateFailure.GetUnsuccessfulOutcome().GetProcedureCode().GetE2NodeConfigurationUpdate().GetUnsuccessfulOutcome().GetProtocolIes().GetE2ApProtocolIes1())
	fmt.Printf("TimeToWait is \n%v\n", e2nodeConfigurationUpdateFailure.GetUnsuccessfulOutcome().GetProcedureCode().GetE2NodeConfigurationUpdate().GetUnsuccessfulOutcome().GetProtocolIes().GetE2ApProtocolIes31())

	if err := e2nodeConfigurationUpdateFailure.Validate(); err != nil {
		return nil, fmt.Errorf("error validating E2nodeConfigurationUpdateFailure %s", err.Error())
	}
	return e2nodeConfigurationUpdateFailure.GetUnsuccessfulOutcome().GetProcedureCode().GetE2NodeConfigurationUpdate().GetUnsuccessfulOutcome(), nil
}

func Test_xerEncodingE2nodeConfigurationUpdateFailure(t *testing.T) {

	e2nodeConfigurationUpdateFailure, err := createE2nodeConfigurationUpdateFailureMsg()
	assert.NilError(t, err, "Error creating E2nodeConfigurationUpdateFailure PDU")
	t.Logf("Created message is \n%v", e2nodeConfigurationUpdateFailure.GetProtocolIes().GetE2ApProtocolIes31())

	xer, err := xerEncodeE2nodeConfigurationUpdateFailure(e2nodeConfigurationUpdateFailure)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(xer)) //ToDo - adjust length of the XER encoded message
	t.Logf("E2nodeConfigurationUpdateFailure XER\n%s", string(xer))

	result, err := xerDecodeE2nodeConfigurationUpdateFailure(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeConfigurationUpdateFailure XER - decoded\n%v", result)
	//ToDo - adjust field's verification
	assert.Equal(t, e2nodeConfigurationUpdateFailure.GetProtocolIes().GetE2ApProtocolIes1().GetValue().GetProtocol(), result.GetProtocolIes().GetE2ApProtocolIes1().GetValue().GetProtocol())
	assert.Equal(t, e2nodeConfigurationUpdateFailure.GetProtocolIes().GetE2ApProtocolIes2().GetValue().GetRicRequestorId().GetRicInstanceId(), result.GetProtocolIes().GetE2ApProtocolIes2().GetValue().GetRicRequestorId().GetRicInstanceId())
	assert.Equal(t, e2nodeConfigurationUpdateFailure.GetProtocolIes().GetE2ApProtocolIes2().GetValue().GetRicRequestorId().GetRicRequestorId(), result.GetProtocolIes().GetE2ApProtocolIes2().GetValue().GetRicRequestorId().GetRicRequestorId())
	assert.Equal(t, e2nodeConfigurationUpdateFailure.GetProtocolIes().GetE2ApProtocolIes31().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes31().GetValue())
}

func Test_perEncodingE2nodeConfigurationUpdateFailure(t *testing.T) {

	e2nodeConfigurationUpdateFailure, err := createE2nodeConfigurationUpdateFailureMsg()
	assert.NilError(t, err, "Error creating E2nodeConfigurationUpdateFailure PDU")

	per, err := perEncodeE2nodeConfigurationUpdateFailure(e2nodeConfigurationUpdateFailure)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(per)) // ToDo - adjust length of the PER encoded message
	t.Logf("E2nodeConfigurationUpdateFailure PER\n%v", hex.Dump(per))

	result, err := perDecodeE2nodeConfigurationUpdateFailure(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeConfigurationUpdateFailure PER - decoded\n%v", result)
	//ToDo - adjust field's verification
	assert.Equal(t, e2nodeConfigurationUpdateFailure.GetProtocolIes().GetE2ApProtocolIes1().GetValue().GetProtocol(), result.GetProtocolIes().GetE2ApProtocolIes1().GetValue().GetProtocol())
	assert.Equal(t, e2nodeConfigurationUpdateFailure.GetProtocolIes().GetE2ApProtocolIes2().GetValue().GetRicRequestorId().GetRicInstanceId(), result.GetProtocolIes().GetE2ApProtocolIes2().GetValue().GetRicRequestorId().GetRicInstanceId())
	assert.Equal(t, e2nodeConfigurationUpdateFailure.GetProtocolIes().GetE2ApProtocolIes2().GetValue().GetRicRequestorId().GetRicRequestorId(), result.GetProtocolIes().GetE2ApProtocolIes2().GetValue().GetRicRequestorId().GetRicRequestorId())
	assert.Equal(t, e2nodeConfigurationUpdateFailure.GetProtocolIes().GetE2ApProtocolIes31().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes31().GetValue())
}
