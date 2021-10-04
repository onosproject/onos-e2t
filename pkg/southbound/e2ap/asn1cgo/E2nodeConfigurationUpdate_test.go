// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

import (
	"encoding/hex"
	"testing"

	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	e2ap_pdu_contents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/pdubuilder"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"gotest.tools/assert"
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

	e2ncID1 := pdubuilder.CreateE2NodeComponentIDW1(21)
	e2ncID2 := pdubuilder.CreateE2NodeComponentIDNg("ONF")

	e2nodeConfigurationUpdate, err := pdubuilder.CreateE2NodeConfigurationUpdateE2apPdu(1)
	if err != nil {
		return nil, err
	}

	e2nodeConfigurationUpdate.GetInitiatingMessage().GetProcedureCode().GetE2NodeConfigurationUpdate().GetInitiatingMessage().
		SetE2nodeComponentConfigUpdate([]*types.E2NodeComponentConfigUpdateItem{
			{E2NodeComponentType: e2ap_ies.E2NodeComponentInterfaceType_E2NODE_COMPONENT_INTERFACE_TYPE_W1,
				E2NodeComponentID: e2ncID1,
				E2NodeComponentConfiguration: e2ap_ies.E2NodeComponentConfiguration{
					E2NodeComponentResponsePart: []byte{0x00, 0x01},
					E2NodeComponentRequestPart:  []byte{0x02, 0x03},
				}},
			{E2NodeComponentType: e2ap_ies.E2NodeComponentInterfaceType_E2NODE_COMPONENT_INTERFACE_TYPE_NG,
				E2NodeComponentID: e2ncID2,
				E2NodeComponentConfiguration: e2ap_ies.E2NodeComponentConfiguration{
					E2NodeComponentResponsePart: []byte{0x04, 0x05},
					E2NodeComponentRequestPart:  []byte{0x06, 0x07},
				}},
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
	assert.Equal(t, e2nodeConfigurationUpdate.GetProtocolIes().GetE2ApProtocolIes33().GetValue().GetValue()[0].GetValue().GetE2NodeComponentInterfaceType().Number(), result.GetProtocolIes().GetE2ApProtocolIes33().GetValue().GetValue()[0].GetValue().GetE2NodeComponentInterfaceType().Number())
	assert.DeepEqual(t, e2nodeConfigurationUpdate.GetProtocolIes().GetE2ApProtocolIes33().GetValue().GetValue()[0].GetValue().GetE2NodeComponentConfiguration().GetE2NodeComponentRequestPart(), result.GetProtocolIes().GetE2ApProtocolIes33().GetValue().GetValue()[0].GetValue().GetE2NodeComponentConfiguration().GetE2NodeComponentRequestPart())
	assert.DeepEqual(t, e2nodeConfigurationUpdate.GetProtocolIes().GetE2ApProtocolIes33().GetValue().GetValue()[0].GetValue().GetE2NodeComponentConfiguration().GetE2NodeComponentResponsePart(), result.GetProtocolIes().GetE2ApProtocolIes33().GetValue().GetValue()[0].GetValue().GetE2NodeComponentConfiguration().GetE2NodeComponentResponsePart())
	assert.Equal(t, e2nodeConfigurationUpdate.GetProtocolIes().GetE2ApProtocolIes33().GetValue().GetValue()[1].GetValue().GetE2NodeComponentInterfaceType().Number(), result.GetProtocolIes().GetE2ApProtocolIes33().GetValue().GetValue()[1].GetValue().GetE2NodeComponentInterfaceType().Number())
	assert.DeepEqual(t, e2nodeConfigurationUpdate.GetProtocolIes().GetE2ApProtocolIes33().GetValue().GetValue()[1].GetValue().GetE2NodeComponentConfiguration().GetE2NodeComponentRequestPart(), result.GetProtocolIes().GetE2ApProtocolIes33().GetValue().GetValue()[1].GetValue().GetE2NodeComponentConfiguration().GetE2NodeComponentRequestPart())
	assert.DeepEqual(t, e2nodeConfigurationUpdate.GetProtocolIes().GetE2ApProtocolIes33().GetValue().GetValue()[1].GetValue().GetE2NodeComponentConfiguration().GetE2NodeComponentResponsePart(), result.GetProtocolIes().GetE2ApProtocolIes33().GetValue().GetValue()[1].GetValue().GetE2NodeComponentConfiguration().GetE2NodeComponentResponsePart())
	assert.Equal(t, e2nodeConfigurationUpdate.GetProtocolIes().GetE2ApProtocolIes49().GetValue().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes49().GetValue().GetValue())
	assert.DeepEqual(t, e2nodeConfigurationUpdate.GetProtocolIes().GetE2ApProtocolIes3().GetValue().GetEnGNb().GetGlobalEnGNbId().GetPLmnIdentity().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes3().GetValue().GetEnGNb().GetGlobalEnGNbId().GetPLmnIdentity().GetValue())
	assert.DeepEqual(t, e2nodeConfigurationUpdate.GetProtocolIes().GetE2ApProtocolIes3().GetValue().GetEnGNb().GetGlobalEnGNbId().GetGNbId().GetGNbId().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes3().GetValue().GetEnGNb().GetGlobalEnGNbId().GetGNbId().GetGNbId().GetValue())
	assert.Equal(t, e2nodeConfigurationUpdate.GetProtocolIes().GetE2ApProtocolIes3().GetValue().GetEnGNb().GetGlobalEnGNbId().GetGNbId().GetGNbId().GetLen(), result.GetProtocolIes().GetE2ApProtocolIes3().GetValue().GetEnGNb().GetGlobalEnGNbId().GetGNbId().GetGNbId().GetLen())
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
	assert.Equal(t, e2nodeConfigurationUpdate.GetProtocolIes().GetE2ApProtocolIes33().GetValue().GetValue()[0].GetValue().GetE2NodeComponentInterfaceType().Number(), result.GetProtocolIes().GetE2ApProtocolIes33().GetValue().GetValue()[0].GetValue().GetE2NodeComponentInterfaceType().Number())
	assert.DeepEqual(t, e2nodeConfigurationUpdate.GetProtocolIes().GetE2ApProtocolIes33().GetValue().GetValue()[0].GetValue().GetE2NodeComponentConfiguration().GetE2NodeComponentRequestPart(), result.GetProtocolIes().GetE2ApProtocolIes33().GetValue().GetValue()[0].GetValue().GetE2NodeComponentConfiguration().GetE2NodeComponentRequestPart())
	assert.DeepEqual(t, e2nodeConfigurationUpdate.GetProtocolIes().GetE2ApProtocolIes33().GetValue().GetValue()[0].GetValue().GetE2NodeComponentConfiguration().GetE2NodeComponentResponsePart(), result.GetProtocolIes().GetE2ApProtocolIes33().GetValue().GetValue()[0].GetValue().GetE2NodeComponentConfiguration().GetE2NodeComponentResponsePart())
	assert.Equal(t, e2nodeConfigurationUpdate.GetProtocolIes().GetE2ApProtocolIes33().GetValue().GetValue()[1].GetValue().GetE2NodeComponentInterfaceType().Number(), result.GetProtocolIes().GetE2ApProtocolIes33().GetValue().GetValue()[1].GetValue().GetE2NodeComponentInterfaceType().Number())
	assert.DeepEqual(t, e2nodeConfigurationUpdate.GetProtocolIes().GetE2ApProtocolIes33().GetValue().GetValue()[1].GetValue().GetE2NodeComponentConfiguration().GetE2NodeComponentRequestPart(), result.GetProtocolIes().GetE2ApProtocolIes33().GetValue().GetValue()[1].GetValue().GetE2NodeComponentConfiguration().GetE2NodeComponentRequestPart())
	assert.DeepEqual(t, e2nodeConfigurationUpdate.GetProtocolIes().GetE2ApProtocolIes33().GetValue().GetValue()[1].GetValue().GetE2NodeComponentConfiguration().GetE2NodeComponentResponsePart(), result.GetProtocolIes().GetE2ApProtocolIes33().GetValue().GetValue()[1].GetValue().GetE2NodeComponentConfiguration().GetE2NodeComponentResponsePart())
	assert.Equal(t, e2nodeConfigurationUpdate.GetProtocolIes().GetE2ApProtocolIes49().GetValue().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes49().GetValue().GetValue())
	assert.DeepEqual(t, e2nodeConfigurationUpdate.GetProtocolIes().GetE2ApProtocolIes3().GetValue().GetEnGNb().GetGlobalEnGNbId().GetPLmnIdentity().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes3().GetValue().GetEnGNb().GetGlobalEnGNbId().GetPLmnIdentity().GetValue())
	assert.DeepEqual(t, e2nodeConfigurationUpdate.GetProtocolIes().GetE2ApProtocolIes3().GetValue().GetEnGNb().GetGlobalEnGNbId().GetGNbId().GetGNbId().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes3().GetValue().GetEnGNb().GetGlobalEnGNbId().GetGNbId().GetGNbId().GetValue())
	assert.Equal(t, e2nodeConfigurationUpdate.GetProtocolIes().GetE2ApProtocolIes3().GetValue().GetEnGNb().GetGlobalEnGNbId().GetGNbId().GetGNbId().GetLen(), result.GetProtocolIes().GetE2ApProtocolIes3().GetValue().GetEnGNb().GetGlobalEnGNbId().GetGNbId().GetGNbId().GetLen())
	assert.Equal(t, e2nodeConfigurationUpdate.GetProtocolIes().GetE2ApProtocolIes3().GetValue().GetEnGNb().GetEnGNbCuUpId().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes3().GetValue().GetEnGNb().GetEnGNbCuUpId().GetValue())
	assert.Equal(t, e2nodeConfigurationUpdate.GetProtocolIes().GetE2ApProtocolIes3().GetValue().GetEnGNb().GetEnGNbDuId().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes3().GetValue().GetEnGNb().GetEnGNbDuId().GetValue())
	assert.Equal(t, e2nodeConfigurationUpdate.String(), result.String())
}
