// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
package pdubuilder

import (
	"encoding/hex"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/asn1cgo"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/types"
	"gotest.tools/assert"
	"testing"
)

func TestE2NodeConfigurationUpdate(t *testing.T) {

	//e2ncID1 := CreateE2NodeComponentIDGnbCuUp(21)
	//e2ncID2 := CreateE2NodeComponentIDGnbDu(13)
	e2nccu1 := CreateE2NodeComponentConfigUpdateGnb("ngAp", "", "e1Ap", "f1Ap")
	e2nccu2 := CreateE2NodeComponentConfigUpdateEnb("s1", "")

	newE2apPdu, err := CreateE2NodeConfigurationUpdateE2apPdu([]*types.E2NodeComponentConfigUpdateItem{
		{E2NodeComponentType: e2ap_ies.E2NodeComponentType_E2NODE_COMPONENT_TYPE_G_NB,
			//E2NodeComponentID:           &e2ncID1,
			E2NodeComponentConfigUpdate: e2nccu1},
		{E2NodeComponentType: e2ap_ies.E2NodeComponentType_E2NODE_COMPONENT_TYPE_E_NB,
			//E2NodeComponentID:           &e2ncID2,
			E2NodeComponentConfigUpdate: e2nccu2},
	})
	assert.NilError(t, err)
	assert.Assert(t, newE2apPdu != nil)

	xer, err := asn1cgo.XerEncodeE2apPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("E2NodeConfigurationUpdate E2AP PDU XER\n%s", string(xer))

	result1, err := asn1cgo.XerDecodeE2apPdu(xer)
	assert.NilError(t, err)
	t.Logf("E2NodeConfigurationUpdate E2AP PDU XER - decoded\n%v", result1)
	assert.DeepEqual(t, newE2apPdu, result1)

	per, err := asn1cgo.PerEncodeE2apPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("E2NodeConfigurationUpdate E2AP PDU PER\n%v", hex.Dump(per))

	resultPer, err := asn1cgo.PerDecodeE2apPdu(per)
	assert.NilError(t, err)
	t.Logf("E2NodeConfigurationUpdate E2AP PDU PER - decoded\n%v", resultPer)
	assert.DeepEqual(t, newE2apPdu, resultPer)
}
