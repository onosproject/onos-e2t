// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0
package pdubuilder

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap_go/encoder"
	"testing"

	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap_go/v2/e2ap-ies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap_go/types"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"gotest.tools/assert"
)

func TestE2NodeConfigurationUpdate(t *testing.T) {

	ge2nID, err := CreateGlobalE2nodeIDEnGnb([3]byte{0x00, 0x00, 0x01}, &asn1.BitString{
		Value: []byte{0x00, 0x00, 0x00, 0x80},
		Len:   25,
	})
	assert.NilError(t, err)
	ge2nID.GetEnGNb().SetGnbCuUpID(2).SetGnbDuID(13)

	e2ncID1 := CreateE2NodeComponentIDNg("NG-Component")
	e2ncID2 := CreateE2NodeComponentIDE1(13)

	e2nodeConfigurationUpdate, err := CreateE2NodeConfigurationUpdateE2apPdu(1)
	assert.NilError(t, err)
	assert.Assert(t, e2nodeConfigurationUpdate != nil)

	e2nodeConfigurationUpdate.GetInitiatingMessage().GetValue().GetE2NodeConfigurationUpdate().
		SetE2nodeComponentConfigUpdate([]*types.E2NodeComponentConfigUpdateItem{
			{E2NodeComponentType: e2ap_ies.E2NodeComponentInterfaceType_E2NODE_COMPONENT_INTERFACE_TYPE_NG,
				E2NodeComponentID: e2ncID1,
				E2NodeComponentConfiguration: e2ap_ies.E2NodeComponentConfiguration{
					E2NodeComponentResponsePart: []byte{0x01, 0x02, 0x03},
					E2NodeComponentRequestPart:  []byte{0x04, 0x05, 0x06},
				}},
			{E2NodeComponentType: e2ap_ies.E2NodeComponentInterfaceType_E2NODE_COMPONENT_INTERFACE_TYPE_E1,
				E2NodeComponentID: e2ncID2,
				E2NodeComponentConfiguration: e2ap_ies.E2NodeComponentConfiguration{
					E2NodeComponentResponsePart: []byte{0x07, 0x08, 0x09},
					E2NodeComponentRequestPart:  []byte{0x0A, 0x0B, 0x0C},
				}},
		}).SetGlobalE2nodeID(ge2nID)

	perNew, err := encoder.PerEncodeE2ApPdu(e2nodeConfigurationUpdate)
	assert.NilError(t, err)
	t.Logf("E2nodeConfigurationUpdate E2AP PDU PER with Go APER library\n%v", hex.Dump(perNew))

	//Comparing reference PER bytes with Go APER library produced
	//assert.DeepEqual(t, per, perNew)

	e2apPdu, err := encoder.PerDecodeE2ApPdu(perNew)
	assert.NilError(t, err)
	assert.DeepEqual(t, e2nodeConfigurationUpdate.String(), e2apPdu.String())

	//per, err := asn1cgo.PerEncodeE2apPdu(e2nodeConfigurationUpdate)
	//assert.NilError(t, err)
	//t.Logf("E2NodeConfigurationUpdate E2AP PDU PER\n%v", hex.Dump(per))
	//
	//resultPer, err := asn1cgo.PerDecodeE2apPdu(per)
	//assert.NilError(t, err)
	//t.Logf("E2NodeConfigurationUpdate E2AP PDU PER - decoded\n%v", resultPer)
	//assert.DeepEqual(t, e2nodeConfigurationUpdate.String(), resultPer.String())
}
