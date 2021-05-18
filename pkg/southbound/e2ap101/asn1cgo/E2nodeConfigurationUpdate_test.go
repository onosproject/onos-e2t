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

func createE2nodeConfigurationUpdateMsg() (*e2ap_pdu_contents.E2NodeConfigurationUpdate, error) {

	e2nodeConfigurationUpdate, err := pdubuilder.CreateE2NodeConfigurationUpdateE2apPdu() //ToDo - fill in arguments here(if this function exists
	if err != nil {
		return nil, err
	}

	if err := e2nodeConfigurationUpdate.Validate(); err != nil {
		return nil, fmt.Errorf("error validating E2nodeConfigurationUpdate %s", err.Error())
	}
	return e2nodeConfigurationUpdate.GetInitiatingMessage().GetProcedureCode().GetE2NodeConfigurationUpdate().GetInitiatingMessage(), nil
}

func Test_xerEncodingE2nodeConfigurationUpdate(t *testing.T) {

	e2nodeConfigurationUpdate, err := createE2nodeConfigurationUpdateMsg()
	assert.NilError(t, err, "Error creating E2nodeConfigurationUpdate PDU")

	xer, err := xerEncodeE2nodeConfigurationUpdate(e2nodeConfigurationUpdate)
	assert.NilError(t, err)
	assert.Equal(t, 1776, len(xer))
	t.Logf("E2nodeConfigurationUpdate XER\n%s", string(xer))

	result, err := xerDecodeE2nodeConfigurationUpdate(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeConfigurationUpdate XER - decoded\n%v", result)
	assert.Equal(t, e2nodeConfigurationUpdate.GetProtocolIes().GetValue().GetValue()[0].GetValue().GetE2NodeComponentType(), result.GetProtocolIes().GetValue().GetValue()[0].GetValue().GetE2NodeComponentType())
	assert.Equal(t, e2nodeConfigurationUpdate.GetProtocolIes().GetValue().GetValue()[0].GetValue().GetE2NodeComponentConfigUpdate().GetGNbconfigUpdate().GetXnApconfigUpdate(), result.GetProtocolIes().GetValue().GetValue()[0].GetValue().GetE2NodeComponentConfigUpdate().GetGNbconfigUpdate().GetXnApconfigUpdate())
}

func Test_perEncodingE2nodeConfigurationUpdate(t *testing.T) {

	e2nodeConfigurationUpdate, err := createE2nodeConfigurationUpdateMsg()
	assert.NilError(t, err, "Error creating E2nodeConfigurationUpdate PDU")

	per, err := perEncodeE2nodeConfigurationUpdate(e2nodeConfigurationUpdate)
	assert.NilError(t, err)
	assert.Equal(t, 41, len(per))
	t.Logf("E2nodeConfigurationUpdate PER\n%v", hex.Dump(per))

	result, err := perDecodeE2nodeConfigurationUpdate(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeConfigurationUpdate PER - decoded\n%v", result)
	assert.Equal(t, e2nodeConfigurationUpdate.GetProtocolIes().GetValue().GetValue()[0].GetValue().GetE2NodeComponentType(), result.GetProtocolIes().GetValue().GetValue()[0].GetValue().GetE2NodeComponentType())
	assert.Equal(t, e2nodeConfigurationUpdate.GetProtocolIes().GetValue().GetValue()[0].GetValue().GetE2NodeComponentConfigUpdate().GetGNbconfigUpdate().GetXnApconfigUpdate(), result.GetProtocolIes().GetValue().GetValue()[0].GetValue().GetE2NodeComponentConfigUpdate().GetGNbconfigUpdate().GetXnApconfigUpdate())
}
