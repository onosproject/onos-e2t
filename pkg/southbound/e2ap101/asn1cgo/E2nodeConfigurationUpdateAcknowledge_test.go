// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

import (
	"encoding/hex"
	"fmt"
	e2ap_pdu_contents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/pdubuilder"
	"gotest.tools/assert"
	"testing"
)

func createE2nodeConfigurationUpdateAcknowledgeMsg() (*e2ap_pdu_contents.E2NodeConfigurationUpdateAcknowledge, error) {

	e2nodeConfigurationUpdateAcknowledge, err := pdubuilder.CreateE2NodeConfigurationUpdateAcknowledgeE2apPdu() //ToDo - fill in arguments here(if this function exists
	if err != nil {
		return nil, err
	}

	if err := e2nodeConfigurationUpdateAcknowledge.Validate(); err != nil {
		return nil, fmt.Errorf("error validating E2nodeConfigurationUpdateAcknowledge %s", err.Error())
	}
	return e2nodeConfigurationUpdateAcknowledge.GetSuccessfulOutcome().GetProcedureCode().GetE2NodeConfigurationUpdate().GetSuccessfulOutcome(), nil
}

func Test_xerEncodingE2nodeConfigurationUpdateAcknowledge(t *testing.T) {

	e2nodeConfigurationUpdateAcknowledge, err := createE2nodeConfigurationUpdateAcknowledgeMsg()
	assert.NilError(t, err, "Error creating E2nodeConfigurationUpdateAcknowledge PDU")

	xer, err := xerEncodeE2nodeConfigurationUpdateAcknowledge(e2nodeConfigurationUpdateAcknowledge)
	assert.NilError(t, err)
	assert.Equal(t, 1628, len(xer))
	t.Logf("E2nodeConfigurationUpdateAcknowledge XER\n%s", string(xer))

	result, err := xerDecodeE2nodeConfigurationUpdateAcknowledge(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeConfigurationUpdateAcknowledge XER - decoded\n%v", result)
	//ToDo - adjust field's verification
	assert.Equal(t, e2nodeConfigurationUpdateAcknowledge.GetProtocolIes().GetValue().GetValue()[0].GetId(), result.GetProtocolIes().GetValue().GetValue()[0].GetId())
	assert.Equal(t, e2nodeConfigurationUpdateAcknowledge.GetProtocolIes().GetValue().GetValue()[0].GetValue().GetE2NodeComponentType(), result.GetProtocolIes().GetValue().GetValue()[0].GetValue().GetE2NodeComponentType())
	assert.Equal(t, e2nodeConfigurationUpdateAcknowledge.GetProtocolIes().GetValue().GetValue()[0].GetValue().GetE2NodeComponentConfigUpdateAck().GetUpdateOutcome(), result.GetProtocolIes().GetValue().GetValue()[0].GetValue().GetE2NodeComponentConfigUpdateAck().GetUpdateOutcome())
}

func Test_perEncodingE2nodeConfigurationUpdateAcknowledge(t *testing.T) {

	e2nodeConfigurationUpdateAcknowledge, err := createE2nodeConfigurationUpdateAcknowledgeMsg()
	assert.NilError(t, err, "Error creating E2nodeConfigurationUpdateAcknowledge PDU")

	per, err := perEncodeE2nodeConfigurationUpdateAcknowledge(e2nodeConfigurationUpdateAcknowledge)
	assert.NilError(t, err)
	assert.Equal(t, 18, len(per))
	t.Logf("E2nodeConfigurationUpdateAcknowledge PER\n%v", hex.Dump(per))

	result, err := perDecodeE2nodeConfigurationUpdateAcknowledge(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeConfigurationUpdateAcknowledge PER - decoded\n%v", result)
	//ToDo - adjust field's verification
	assert.Equal(t, e2nodeConfigurationUpdateAcknowledge.GetProtocolIes().GetValue().GetValue()[0].GetId(), result.GetProtocolIes().GetValue().GetValue()[0].GetId())
	assert.Equal(t, e2nodeConfigurationUpdateAcknowledge.GetProtocolIes().GetValue().GetValue()[0].GetValue().GetE2NodeComponentType(), result.GetProtocolIes().GetValue().GetValue()[0].GetValue().GetE2NodeComponentType())
	assert.Equal(t, e2nodeConfigurationUpdateAcknowledge.GetProtocolIes().GetValue().GetValue()[0].GetValue().GetE2NodeComponentConfigUpdateAck().GetUpdateOutcome(), result.GetProtocolIes().GetValue().GetValue()[0].GetValue().GetE2NodeComponentConfigUpdateAck().GetUpdateOutcome())
}
