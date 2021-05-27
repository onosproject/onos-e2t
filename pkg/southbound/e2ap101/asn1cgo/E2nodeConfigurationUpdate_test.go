// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

import (
	"encoding/hex"
	"fmt"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	e2ap_pdu_contents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/pdubuilder"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/types"
	"gotest.tools/assert"
	"testing"
)

func createE2nodeConfigurationUpdateMsg() (*e2ap_pdu_contents.E2NodeConfigurationUpdate, error) {

	e2ncID1 := pdubuilder.CreateE2NodeComponentIDGnbCuUp(21)
	e2ncID2 := pdubuilder.CreateE2NodeComponentIDGnbDu(13)
	e2nccu1 := pdubuilder.CreateE2NodeComponentConfigUpdateGnb("ngAp", "xnAp", "e1Ap", "f1Ap")
	e2nccu2 := pdubuilder.CreateE2NodeComponentConfigUpdateEnb("s1", "x2")

	e2nodeConfigurationUpdate, err := pdubuilder.CreateE2NodeConfigurationUpdateE2apPdu([]*types.E2NodeComponentConfigUpdateItem{
		{E2NodeComponentType: e2ap_ies.E2NodeComponentType_E2NODE_COMPONENT_TYPE_G_NB,
			E2NodeComponentID:           e2ncID1,
			E2NodeComponentConfigUpdate: e2nccu1},
		{E2NodeComponentType: e2ap_ies.E2NodeComponentType_E2NODE_COMPONENT_TYPE_E_NB,
			E2NodeComponentID:           e2ncID2,
			E2NodeComponentConfigUpdate: e2nccu2},
	})
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
	assert.Equal(t, 2934, len(xer))
	t.Logf("E2nodeConfigurationUpdate XER\n%s", string(xer))

	result, err := xerDecodeE2nodeConfigurationUpdate(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeConfigurationUpdate XER - decoded\n%v", result)
	assert.Equal(t, e2nodeConfigurationUpdate.GetProtocolIes().GetValue().GetValue()[0].GetValue().GetE2NodeComponentType(), result.GetProtocolIes().GetValue().GetValue()[0].GetValue().GetE2NodeComponentType())
	assert.Equal(t, e2nodeConfigurationUpdate.GetProtocolIes().GetValue().GetValue()[0].GetValue().GetE2NodeComponentConfigUpdate().GetGNbconfigUpdate().GetXnApconfigUpdate(), result.GetProtocolIes().GetValue().GetValue()[0].GetValue().GetE2NodeComponentConfigUpdate().GetGNbconfigUpdate().GetXnApconfigUpdate())
	assert.Equal(t, e2nodeConfigurationUpdate.GetProtocolIes().GetValue().GetValue()[1].GetValue().GetE2NodeComponentType(), result.GetProtocolIes().GetValue().GetValue()[1].GetValue().GetE2NodeComponentType())
	assert.Equal(t, e2nodeConfigurationUpdate.GetProtocolIes().GetValue().GetValue()[1].GetValue().GetE2NodeComponentConfigUpdate().GetENbconfigUpdate().GetX2ApconfigUpdate(), result.GetProtocolIes().GetValue().GetValue()[1].GetValue().GetE2NodeComponentConfigUpdate().GetENbconfigUpdate().GetX2ApconfigUpdate())
}

func Test_perEncodingE2nodeConfigurationUpdate(t *testing.T) {

	e2nodeConfigurationUpdate, err := createE2nodeConfigurationUpdateMsg()
	assert.NilError(t, err, "Error creating E2nodeConfigurationUpdate PDU")

	per, err := perEncodeE2nodeConfigurationUpdate(e2nodeConfigurationUpdate)
	assert.NilError(t, err)
	assert.Equal(t, 51, len(per))
	t.Logf("E2nodeConfigurationUpdate PER\n%v", hex.Dump(per))

	result, err := perDecodeE2nodeConfigurationUpdate(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeConfigurationUpdate PER - decoded\n%v", result)
	assert.Equal(t, e2nodeConfigurationUpdate.GetProtocolIes().GetValue().GetValue()[0].GetValue().GetE2NodeComponentType(), result.GetProtocolIes().GetValue().GetValue()[0].GetValue().GetE2NodeComponentType())
	assert.Equal(t, e2nodeConfigurationUpdate.GetProtocolIes().GetValue().GetValue()[0].GetValue().GetE2NodeComponentConfigUpdate().GetGNbconfigUpdate().GetXnApconfigUpdate(), result.GetProtocolIes().GetValue().GetValue()[0].GetValue().GetE2NodeComponentConfigUpdate().GetGNbconfigUpdate().GetXnApconfigUpdate())
	assert.Equal(t, e2nodeConfigurationUpdate.GetProtocolIes().GetValue().GetValue()[1].GetValue().GetE2NodeComponentType(), result.GetProtocolIes().GetValue().GetValue()[1].GetValue().GetE2NodeComponentType())
	assert.Equal(t, e2nodeConfigurationUpdate.GetProtocolIes().GetValue().GetValue()[1].GetValue().GetE2NodeComponentConfigUpdate().GetENbconfigUpdate().GetX2ApconfigUpdate(), result.GetProtocolIes().GetValue().GetValue()[1].GetValue().GetE2NodeComponentConfigUpdate().GetENbconfigUpdate().GetX2ApconfigUpdate())
}
