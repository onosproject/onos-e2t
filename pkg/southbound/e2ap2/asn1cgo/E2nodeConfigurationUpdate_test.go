// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

import (
	"encoding/hex"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-ies"
	e2ap_pdu_contents "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-pdu-contents"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap2/pdubuilder"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap2/types"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"gotest.tools/assert"
	"testing"
)

func createE2nodeConfigurationUpdateMsg() (*e2ap_pdu_contents.E2NodeConfigurationUpdate, error) {

	ge2nID, err := pdubuilder.CreateGlobalE2nodeIDEnGnb([3]byte{0x00, 0x00, 0x01}, &asn1.BitString{
		Value: []byte{0x00, 0x00, 0x00, 0x80},
		Len:   25,
	})
	if err != nil {
		return nil, err
	}
	ge2nID.GetEnGNb().SetGnbCuUpID(2).SetGnbDuID(13)

	e2ncID1 := pdubuilder.CreateE2NodeComponentIDGnbCuUp(21)
	e2ncID2 := pdubuilder.CreateE2NodeComponentIDGnbDu(13)

	e2nodeConfigurationUpdate, err := pdubuilder.CreateE2NodeConfigurationUpdateE2apPdu(1)
	if err != nil {
		return nil, err
	}

	e2nodeConfigurationUpdate.GetInitiatingMessage().GetProcedureCode().GetE2NodeConfigurationUpdate().GetInitiatingMessage().
		SetE2nodeComponentConfigUpdate([]*types.E2NodeComponentConfigUpdateItem{
			{E2NodeComponentType: e2ap_ies.E2NodeComponentType_E2NODE_COMPONENT_TYPE_G_NB,
				E2NodeComponentID:           &e2ncID1,
				E2NodeComponentConfigUpdate: pdubuilder.CreateE2NodeComponentConfigUpdateGnb([]byte("ngAp"), []byte("xnAp"), []byte("e1Ap"), []byte("f1Ap"), nil)},
			{E2NodeComponentType: e2ap_ies.E2NodeComponentType_E2NODE_COMPONENT_TYPE_E_NB,
				E2NodeComponentID:           &e2ncID2,
				E2NodeComponentConfigUpdate: pdubuilder.CreateE2NodeComponentConfigUpdateEnb(nil, nil, nil, []byte("s1"), []byte("x2"))},
		}).SetGlobalE2nodeID(ge2nID)

	//if err := e2nodeConfigurationUpdate.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating E2nodeConfigurationUpdate %s", err.Error())
	//}
	return e2nodeConfigurationUpdate.GetInitiatingMessage().GetProcedureCode().GetE2NodeConfigurationUpdate().GetInitiatingMessage(), nil
}

func Test_xerEncodingE2nodeConfigurationUpdate(t *testing.T) {

	e2nodeConfigurationUpdate, err := createE2nodeConfigurationUpdateMsg()
	assert.NilError(t, err, "Error creating E2nodeConfigurationUpdate PDU")

	xer, err := xerEncodeE2nodeConfigurationUpdate(e2nodeConfigurationUpdate)
	assert.NilError(t, err)
	t.Logf("E2nodeConfigurationUpdate XER\n%s", string(xer))

	result, err := xerDecodeE2nodeConfigurationUpdate(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeConfigurationUpdate XER - decoded\n%v", result)
	assert.Equal(t, e2nodeConfigurationUpdate.GetProtocolIes().GetE2ApProtocolIes33().GetValue().GetValue()[0].GetValue().GetE2NodeComponentType().Number(), result.GetProtocolIes().GetE2ApProtocolIes33().GetValue().GetValue()[0].GetValue().GetE2NodeComponentType().Number())
	assert.DeepEqual(t, e2nodeConfigurationUpdate.GetProtocolIes().GetE2ApProtocolIes33().GetValue().GetValue()[0].GetValue().GetE2NodeComponentConfigUpdate().GetGNbconfigUpdate().GetXnApconfigUpdate(), result.GetProtocolIes().GetE2ApProtocolIes33().GetValue().GetValue()[0].GetValue().GetE2NodeComponentConfigUpdate().GetGNbconfigUpdate().GetXnApconfigUpdate())
	assert.Equal(t, e2nodeConfigurationUpdate.GetProtocolIes().GetE2ApProtocolIes33().GetValue().GetValue()[1].GetValue().GetE2NodeComponentType().Number(), result.GetProtocolIes().GetE2ApProtocolIes33().GetValue().GetValue()[1].GetValue().GetE2NodeComponentType().Number())
	assert.DeepEqual(t, e2nodeConfigurationUpdate.GetProtocolIes().GetE2ApProtocolIes33().GetValue().GetValue()[1].GetValue().GetE2NodeComponentConfigUpdate().GetENbconfigUpdate().GetX2ApconfigUpdate(), result.GetProtocolIes().GetE2ApProtocolIes33().GetValue().GetValue()[1].GetValue().GetE2NodeComponentConfigUpdate().GetENbconfigUpdate().GetX2ApconfigUpdate())
	assert.Equal(t, e2nodeConfigurationUpdate.GetProtocolIes().GetE2ApProtocolIes49().GetValue().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes49().GetValue().GetValue())
	assert.DeepEqual(t, e2nodeConfigurationUpdate.GetProtocolIes().GetE2ApProtocolIes3().GetValue().GetEnGNb().GetGlobalGNbId().GetPLmnIdentity().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes3().GetValue().GetEnGNb().GetGlobalGNbId().GetPLmnIdentity().GetValue())
	assert.DeepEqual(t, e2nodeConfigurationUpdate.GetProtocolIes().GetE2ApProtocolIes3().GetValue().GetEnGNb().GetGlobalGNbId().GetGNbId().GetGNbId().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes3().GetValue().GetEnGNb().GetGlobalGNbId().GetGNbId().GetGNbId().GetValue())
	assert.Equal(t, e2nodeConfigurationUpdate.GetProtocolIes().GetE2ApProtocolIes3().GetValue().GetEnGNb().GetGlobalGNbId().GetGNbId().GetGNbId().GetLen(), result.GetProtocolIes().GetE2ApProtocolIes3().GetValue().GetEnGNb().GetGlobalGNbId().GetGNbId().GetGNbId().GetLen())
	assert.Equal(t, e2nodeConfigurationUpdate.GetProtocolIes().GetE2ApProtocolIes3().GetValue().GetEnGNb().GetEnGNbCuUpId().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes3().GetValue().GetEnGNb().GetEnGNbCuUpId().GetValue())
	assert.Equal(t, e2nodeConfigurationUpdate.GetProtocolIes().GetE2ApProtocolIes3().GetValue().GetEnGNb().GetEnGNbDuId().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes3().GetValue().GetEnGNb().GetEnGNbDuId().GetValue())
	assert.Equal(t, e2nodeConfigurationUpdate.String(), result.String())
}

func Test_perEncodingE2nodeConfigurationUpdate(t *testing.T) {

	e2nodeConfigurationUpdate, err := createE2nodeConfigurationUpdateMsg()
	assert.NilError(t, err, "Error creating E2nodeConfigurationUpdate PDU")

	per, err := perEncodeE2nodeConfigurationUpdate(e2nodeConfigurationUpdate)
	assert.NilError(t, err)
	t.Logf("E2nodeConfigurationUpdate PER\n%v", hex.Dump(per))

	result, err := perDecodeE2nodeConfigurationUpdate(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeConfigurationUpdate PER - decoded\n%v", result)
	assert.Equal(t, e2nodeConfigurationUpdate.GetProtocolIes().GetE2ApProtocolIes33().GetValue().GetValue()[0].GetValue().GetE2NodeComponentType().Number(), result.GetProtocolIes().GetE2ApProtocolIes33().GetValue().GetValue()[0].GetValue().GetE2NodeComponentType().Number())
	assert.DeepEqual(t, e2nodeConfigurationUpdate.GetProtocolIes().GetE2ApProtocolIes33().GetValue().GetValue()[0].GetValue().GetE2NodeComponentConfigUpdate().GetGNbconfigUpdate().GetXnApconfigUpdate(), result.GetProtocolIes().GetE2ApProtocolIes33().GetValue().GetValue()[0].GetValue().GetE2NodeComponentConfigUpdate().GetGNbconfigUpdate().GetXnApconfigUpdate())
	assert.Equal(t, e2nodeConfigurationUpdate.GetProtocolIes().GetE2ApProtocolIes33().GetValue().GetValue()[1].GetValue().GetE2NodeComponentType().Number(), result.GetProtocolIes().GetE2ApProtocolIes33().GetValue().GetValue()[1].GetValue().GetE2NodeComponentType().Number())
	assert.DeepEqual(t, e2nodeConfigurationUpdate.GetProtocolIes().GetE2ApProtocolIes33().GetValue().GetValue()[1].GetValue().GetE2NodeComponentConfigUpdate().GetENbconfigUpdate().GetX2ApconfigUpdate(), result.GetProtocolIes().GetE2ApProtocolIes33().GetValue().GetValue()[1].GetValue().GetE2NodeComponentConfigUpdate().GetENbconfigUpdate().GetX2ApconfigUpdate())
	assert.Equal(t, e2nodeConfigurationUpdate.GetProtocolIes().GetE2ApProtocolIes49().GetValue().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes49().GetValue().GetValue())
	assert.DeepEqual(t, e2nodeConfigurationUpdate.GetProtocolIes().GetE2ApProtocolIes3().GetValue().GetEnGNb().GetGlobalGNbId().GetPLmnIdentity().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes3().GetValue().GetEnGNb().GetGlobalGNbId().GetPLmnIdentity().GetValue())
	assert.DeepEqual(t, e2nodeConfigurationUpdate.GetProtocolIes().GetE2ApProtocolIes3().GetValue().GetEnGNb().GetGlobalGNbId().GetGNbId().GetGNbId().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes3().GetValue().GetEnGNb().GetGlobalGNbId().GetGNbId().GetGNbId().GetValue())
	assert.Equal(t, e2nodeConfigurationUpdate.GetProtocolIes().GetE2ApProtocolIes3().GetValue().GetEnGNb().GetGlobalGNbId().GetGNbId().GetGNbId().GetLen(), result.GetProtocolIes().GetE2ApProtocolIes3().GetValue().GetEnGNb().GetGlobalGNbId().GetGNbId().GetGNbId().GetLen())
	assert.Equal(t, e2nodeConfigurationUpdate.GetProtocolIes().GetE2ApProtocolIes3().GetValue().GetEnGNb().GetEnGNbCuUpId().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes3().GetValue().GetEnGNb().GetEnGNbCuUpId().GetValue())
	assert.Equal(t, e2nodeConfigurationUpdate.GetProtocolIes().GetE2ApProtocolIes3().GetValue().GetEnGNb().GetEnGNbDuId().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes3().GetValue().GetEnGNb().GetEnGNbDuId().GetValue())
	assert.Equal(t, e2nodeConfigurationUpdate.String(), result.String())
	assert.Equal(t, e2nodeConfigurationUpdate.String(), result.String())
}
