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

func createE2connectionUpdateAcknowledgeMsg() (*e2ap_pdu_contents.E2ConnectionUpdateAcknowledge, error) {

	e2connectionUpdateAcknowledge, err := pdubuilder.CreateE2connectionUpdateAcknowledgeE2apPdu()
	if err != nil {
		return nil, err
	}

	if err := e2connectionUpdateAcknowledge.Validate(); err != nil {
		return nil, fmt.Errorf("error validating E2connectionUpdateAcknowledge %s", err.Error())
	}
	return e2connectionUpdateAcknowledge.GetSuccessfulOutcome().GetProcedureCode().GetE2ConnectionUpdate().GetSuccessfulOutcome(), nil
}

func Test_xerEncodingE2connectionUpdateAcknowledge(t *testing.T) {

	e2connectionUpdateAcknowledge, err := createE2connectionUpdateAcknowledgeMsg()
	assert.NilError(t, err, "Error creating E2connectionUpdateAcknowledge PDU")

	xer, err := xerEncodeE2connectionUpdateAcknowledge(e2connectionUpdateAcknowledge)
	assert.NilError(t, err)
	assert.Equal(t, 2536, len(xer))
	t.Logf("E2connectionUpdateAcknowledge XER\n%s", string(xer))

	result, err := xerDecodeE2connectionUpdateAcknowledge(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2connectionUpdateAcknowledge XER - decoded\n%v", result)
	//ToDo - adjust field's verification
	assert.Equal(t, e2connectionUpdateAcknowledge.GetProtocolIes().GetE2ApProtocolIes39().GetConnectionSetup().GetValue()[0].GetValue().GetTnlUsage(), result.GetProtocolIes().GetE2ApProtocolIes39().GetConnectionSetup().GetValue()[0].GetValue().GetTnlUsage())
	assert.Equal(t, e2connectionUpdateAcknowledge.GetProtocolIes().GetE2ApProtocolIes39().GetConnectionSetup().GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetLen(), result.GetProtocolIes().GetE2ApProtocolIes39().GetConnectionSetup().GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetLen())
	assert.Equal(t, e2connectionUpdateAcknowledge.GetProtocolIes().GetE2ApProtocolIes39().GetConnectionSetup().GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetLen(), result.GetProtocolIes().GetE2ApProtocolIes39().GetConnectionSetup().GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetLen())
	assert.Equal(t, e2connectionUpdateAcknowledge.GetProtocolIes().GetE2ApProtocolIes40().GetConnectionSetupFailed().GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetLen(), result.GetProtocolIes().GetE2ApProtocolIes40().GetConnectionSetupFailed().GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetLen())
	assert.Equal(t, e2connectionUpdateAcknowledge.GetProtocolIes().GetE2ApProtocolIes40().GetConnectionSetupFailed().GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetLen(), result.GetProtocolIes().GetE2ApProtocolIes40().GetConnectionSetupFailed().GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetLen())
	assert.Equal(t, e2connectionUpdateAcknowledge.GetProtocolIes().GetE2ApProtocolIes40().GetConnectionSetupFailed().GetValue()[0].GetValue().GetCause().GetProtocol(), result.GetProtocolIes().GetE2ApProtocolIes40().GetConnectionSetupFailed().GetValue()[0].GetValue().GetCause().GetProtocol())
}

func Test_perEncodingE2connectionUpdateAcknowledge(t *testing.T) {

	e2connectionUpdateAcknowledge, err := createE2connectionUpdateAcknowledgeMsg()
	assert.NilError(t, err, "Error creating E2connectionUpdateAcknowledge PDU")

	per, err := perEncodeE2connectionUpdateAcknowledge(e2connectionUpdateAcknowledge)
	assert.NilError(t, err)
	assert.Equal(t, 47, len(per))
	t.Logf("E2connectionUpdateAcknowledge PER\n%v", hex.Dump(per))

	result, err := perDecodeE2connectionUpdateAcknowledge(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2connectionUpdateAcknowledge PER - decoded\n%v", result)
	//ToDo - adjust field's verification
	assert.Equal(t, e2connectionUpdateAcknowledge.GetProtocolIes().GetE2ApProtocolIes39().GetConnectionSetup().GetValue()[0].GetValue().GetTnlUsage(), result.GetProtocolIes().GetE2ApProtocolIes39().GetConnectionSetup().GetValue()[0].GetValue().GetTnlUsage())
	assert.Equal(t, e2connectionUpdateAcknowledge.GetProtocolIes().GetE2ApProtocolIes39().GetConnectionSetup().GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetLen(), result.GetProtocolIes().GetE2ApProtocolIes39().GetConnectionSetup().GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetLen())
	assert.Equal(t, e2connectionUpdateAcknowledge.GetProtocolIes().GetE2ApProtocolIes39().GetConnectionSetup().GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetLen(), result.GetProtocolIes().GetE2ApProtocolIes39().GetConnectionSetup().GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetLen())
	assert.Equal(t, e2connectionUpdateAcknowledge.GetProtocolIes().GetE2ApProtocolIes40().GetConnectionSetupFailed().GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetLen(), result.GetProtocolIes().GetE2ApProtocolIes40().GetConnectionSetupFailed().GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetLen())
	assert.Equal(t, e2connectionUpdateAcknowledge.GetProtocolIes().GetE2ApProtocolIes40().GetConnectionSetupFailed().GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetLen(), result.GetProtocolIes().GetE2ApProtocolIes40().GetConnectionSetupFailed().GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetLen())
	assert.Equal(t, e2connectionUpdateAcknowledge.GetProtocolIes().GetE2ApProtocolIes40().GetConnectionSetupFailed().GetValue()[0].GetValue().GetCause().GetProtocol(), result.GetProtocolIes().GetE2ApProtocolIes40().GetConnectionSetupFailed().GetValue()[0].GetValue().GetCause().GetProtocol())

}
