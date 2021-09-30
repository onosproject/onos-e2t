// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0
package pdubuilder

import (
	"encoding/hex"
	"testing"

	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-ies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/asn1cgo"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
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

	e2ncID1 := CreateE2NodeComponentIDGnbCuUp(21)
	e2ncID2 := CreateE2NodeComponentIDGnbDu(13)

	e2nodeConfigurationUpdate, err := CreateE2NodeConfigurationUpdateE2apPdu(1)
	assert.NilError(t, err)
	assert.Assert(t, e2nodeConfigurationUpdate != nil)

	e2nodeConfigurationUpdate.GetInitiatingMessage().GetProcedureCode().GetE2NodeConfigurationUpdate().GetInitiatingMessage().
		SetE2nodeComponentConfigUpdate([]*types.E2NodeComponentConfigUpdateItem{
			{E2NodeComponentType: e2ap_ies.E2NodeComponentType_E2NODE_COMPONENT_TYPE_G_NB,
				E2NodeComponentID:           &e2ncID1,
				E2NodeComponentConfigUpdate: CreateE2NodeComponentConfigUpdateGnb([]byte("ngAp"), []byte("xnAp"), []byte("e1Ap"), []byte("f1Ap"), nil)},
			{E2NodeComponentType: e2ap_ies.E2NodeComponentType_E2NODE_COMPONENT_TYPE_E_NB,
				E2NodeComponentID:           &e2ncID2,
				E2NodeComponentConfigUpdate: CreateE2NodeComponentConfigUpdateEnb(nil, nil, nil, []byte("s1"), []byte("x2"))},
		}).SetGlobalE2nodeID(ge2nID)

	xer, err := asn1cgo.XerEncodeE2apPdu(e2nodeConfigurationUpdate)
	assert.NilError(t, err)
	t.Logf("E2NodeConfigurationUpdate E2AP PDU XER\n%s", string(xer))

	result1, err := asn1cgo.XerDecodeE2apPdu(xer)
	assert.NilError(t, err)
	t.Logf("E2NodeConfigurationUpdate E2AP PDU XER - decoded\n%v", result1)
	assert.DeepEqual(t, e2nodeConfigurationUpdate.String(), result1.String())

	per, err := asn1cgo.PerEncodeE2apPdu(e2nodeConfigurationUpdate)
	assert.NilError(t, err)
	t.Logf("E2NodeConfigurationUpdate E2AP PDU PER\n%v", hex.Dump(per))

	resultPer, err := asn1cgo.PerDecodeE2apPdu(per)
	assert.NilError(t, err)
	t.Logf("E2NodeConfigurationUpdate E2AP PDU PER - decoded\n%v", resultPer)
	assert.DeepEqual(t, e2nodeConfigurationUpdate.String(), resultPer.String())
}
