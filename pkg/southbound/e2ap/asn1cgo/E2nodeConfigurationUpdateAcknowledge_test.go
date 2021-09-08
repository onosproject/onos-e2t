// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

import (
	"encoding/hex"
	"testing"

	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-ies"
	e2ap_pdu_contents "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-pdu-contents"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/pdubuilder"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
	"gotest.tools/assert"
)

func createE2nodeConfigurationUpdateAcknowledgeMsg() (*e2ap_pdu_contents.E2NodeConfigurationUpdateAcknowledge, error) {

	//e2ncID1 := pdubuilder.CreateE2NodeComponentIDGnbCuUp(21)
	e2ncID2 := pdubuilder.CreateE2NodeComponentIDGnbDu(13)

	e2nodeConfigurationUpdateAcknowledge, err := pdubuilder.CreateE2NodeConfigurationUpdateAcknowledgeE2apPdu(1)
	if err != nil {
		return nil, err
	}

	e2nodeConfigurationUpdateAcknowledge.GetSuccessfulOutcome().GetProcedureCode().GetE2NodeConfigurationUpdate().GetSuccessfulOutcome().
		SetE2nodeComponentConfigUpdate([]*types.E2NodeComponentConfigUpdateAckItem{
			{E2NodeComponentType: e2ap_ies.E2NodeComponentType_E2NODE_COMPONENT_TYPE_G_NB,
				//E2NodeComponentID: e2ncID1,
				E2NodeComponentConfigUpdateAck: types.E2NodeComponentConfigUpdateAck{
					UpdateOutcome: 1,
					FailureCause: &e2ap_ies.Cause{
						Cause: &e2ap_ies.Cause_Protocol{
							Protocol: e2ap_ies.CauseProtocol_CAUSE_PROTOCOL_TRANSFER_SYNTAX_ERROR,
						},
					},
				}},
			{E2NodeComponentType: e2ap_ies.E2NodeComponentType_E2NODE_COMPONENT_TYPE_E_NB,
				E2NodeComponentID: &e2ncID2,
				E2NodeComponentConfigUpdateAck: types.E2NodeComponentConfigUpdateAck{
					UpdateOutcome: 1,
					//FailureCause: e2ap_ies.Cause{
					//	Cause: &e2ap_ies.Cause_Protocol{
					//		Protocol: e2ap_ies.CauseProtocol_CAUSE_PROTOCOL_ABSTRACT_SYNTAX_ERROR_FALSELY_CONSTRUCTED_MESSAGE,
					//	},
					//},
				}}})

	//if err := e2nodeConfigurationUpdateAcknowledge.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating E2nodeConfigurationUpdateAcknowledge %s", err.Error())
	//}
	return e2nodeConfigurationUpdateAcknowledge.GetSuccessfulOutcome().GetProcedureCode().GetE2NodeConfigurationUpdate().GetSuccessfulOutcome(), nil
}

func Test_xerEncodingE2nodeConfigurationUpdateAcknowledge(t *testing.T) {

	e2nodeConfigurationUpdateAcknowledge, err := createE2nodeConfigurationUpdateAcknowledgeMsg()
	assert.NilError(t, err, "Error creating E2nodeConfigurationUpdateAcknowledge PDU")

	xer, err := xerEncodeE2nodeConfigurationUpdateAcknowledge(e2nodeConfigurationUpdateAcknowledge)
	assert.NilError(t, err)
	t.Logf("E2nodeConfigurationUpdateAcknowledge XER\n%s", string(xer))

	result, err := xerDecodeE2nodeConfigurationUpdateAcknowledge(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeConfigurationUpdateAcknowledge XER - decoded\n%v", result)
	assert.Equal(t, e2nodeConfigurationUpdateAcknowledge.GetProtocolIes().GetE2ApProtocolIes35().GetValue().GetValue()[0].GetId(), result.GetProtocolIes().GetE2ApProtocolIes35().GetValue().GetValue()[0].GetId())
	assert.Equal(t, e2nodeConfigurationUpdateAcknowledge.GetProtocolIes().GetE2ApProtocolIes35().GetValue().GetValue()[0].GetValue().GetE2NodeComponentType(), result.GetProtocolIes().GetE2ApProtocolIes35().GetValue().GetValue()[0].GetValue().GetE2NodeComponentType())
	assert.Equal(t, e2nodeConfigurationUpdateAcknowledge.GetProtocolIes().GetE2ApProtocolIes35().GetValue().GetValue()[0].GetValue().GetE2NodeComponentConfigUpdateAck().GetUpdateOutcome(), result.GetProtocolIes().GetE2ApProtocolIes35().GetValue().GetValue()[0].GetValue().GetE2NodeComponentConfigUpdateAck().GetUpdateOutcome())
	assert.Equal(t, e2nodeConfigurationUpdateAcknowledge.GetProtocolIes().GetE2ApProtocolIes35().GetValue().GetValue()[1].GetValue().GetE2NodeComponentType(), result.GetProtocolIes().GetE2ApProtocolIes35().GetValue().GetValue()[1].GetValue().GetE2NodeComponentType())
	assert.Equal(t, e2nodeConfigurationUpdateAcknowledge.GetProtocolIes().GetE2ApProtocolIes35().GetValue().GetValue()[1].GetValue().GetE2NodeComponentConfigUpdateAck().GetUpdateOutcome(), result.GetProtocolIes().GetE2ApProtocolIes35().GetValue().GetValue()[1].GetValue().GetE2NodeComponentConfigUpdateAck().GetUpdateOutcome())
	assert.Equal(t, e2nodeConfigurationUpdateAcknowledge.GetProtocolIes().GetE2ApProtocolIes49().GetValue().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes49().GetValue().GetValue())
}

func Test_perEncodingE2nodeConfigurationUpdateAcknowledge(t *testing.T) {

	e2nodeConfigurationUpdateAcknowledge, err := createE2nodeConfigurationUpdateAcknowledgeMsg()
	assert.NilError(t, err, "Error creating E2nodeConfigurationUpdateAcknowledge PDU")

	per, err := perEncodeE2nodeConfigurationUpdateAcknowledge(e2nodeConfigurationUpdateAcknowledge)
	assert.NilError(t, err)
	t.Logf("E2nodeConfigurationUpdateAcknowledge PER\n%v", hex.Dump(per))

	result, err := perDecodeE2nodeConfigurationUpdateAcknowledge(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeConfigurationUpdateAcknowledge PER - decoded\n%v", result)
	assert.Equal(t, e2nodeConfigurationUpdateAcknowledge.GetProtocolIes().GetE2ApProtocolIes35().GetValue().GetValue()[0].GetId(), result.GetProtocolIes().GetE2ApProtocolIes35().GetValue().GetValue()[0].GetId())
	assert.Equal(t, e2nodeConfigurationUpdateAcknowledge.GetProtocolIes().GetE2ApProtocolIes35().GetValue().GetValue()[0].GetValue().GetE2NodeComponentType(), result.GetProtocolIes().GetE2ApProtocolIes35().GetValue().GetValue()[0].GetValue().GetE2NodeComponentType())
	assert.Equal(t, e2nodeConfigurationUpdateAcknowledge.GetProtocolIes().GetE2ApProtocolIes35().GetValue().GetValue()[0].GetValue().GetE2NodeComponentConfigUpdateAck().GetUpdateOutcome(), result.GetProtocolIes().GetE2ApProtocolIes35().GetValue().GetValue()[0].GetValue().GetE2NodeComponentConfigUpdateAck().GetUpdateOutcome())
	assert.Equal(t, e2nodeConfigurationUpdateAcknowledge.GetProtocolIes().GetE2ApProtocolIes35().GetValue().GetValue()[1].GetValue().GetE2NodeComponentType(), result.GetProtocolIes().GetE2ApProtocolIes35().GetValue().GetValue()[1].GetValue().GetE2NodeComponentType())
	assert.Equal(t, e2nodeConfigurationUpdateAcknowledge.GetProtocolIes().GetE2ApProtocolIes35().GetValue().GetValue()[1].GetValue().GetE2NodeComponentConfigUpdateAck().GetUpdateOutcome(), result.GetProtocolIes().GetE2ApProtocolIes35().GetValue().GetValue()[1].GetValue().GetE2NodeComponentConfigUpdateAck().GetUpdateOutcome())
	assert.Equal(t, e2nodeConfigurationUpdateAcknowledge.GetProtocolIes().GetE2ApProtocolIes49().GetValue().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes49().GetValue().GetValue())
}
