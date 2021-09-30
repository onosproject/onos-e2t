// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package pdudecoder

import (
	"io/ioutil"
	"testing"

	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-ies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/asn1cgo"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
	"gotest.tools/assert"
)

func Test_DecodeE2nodeConfigurationUpdatePdu(t *testing.T) {
	e2ncuXer, err := ioutil.ReadFile("../test/E2nodeConfigurationUpdate.xml")
	assert.NilError(t, err, "Unexpected error when loading file")
	e2apPdu, err := asn1cgo.XerDecodeE2apPdu(e2ncuXer)
	assert.NilError(t, err)

	transactionID, nodeIdentity, e2nccul, err := DecodeE2nodeConfigurationUpdatePdu(e2apPdu)
	assert.NilError(t, err)
	//assert.Assert(t, ricIdentity != nil) //Commented due to the Linters (v1.34.1) error - possible nil pointer dereference (https://staticcheck.io/docs/checks#SA5011) on lines 23, 24 & 25

	assert.Equal(t, int32(e2nccul[0].E2NodeComponentType), int32(e2apies.E2NodeComponentType_E2NODE_COMPONENT_TYPE_G_NB))
	assert.Equal(t, int32(e2nccul[0].E2NodeComponentID.GetE2NodeComponentTypeGnbCuUp().GetGNbCuUpId().GetValue()), int32(21))
	assert.DeepEqual(t, e2nccul[0].E2NodeComponentConfigUpdate.GetGNbconfigUpdate().GetNgApconfigUpdate(), []byte("ngAp"))
	assert.DeepEqual(t, e2nccul[0].E2NodeComponentConfigUpdate.GetGNbconfigUpdate().GetXnApconfigUpdate(), []byte("xnAp"))
	assert.DeepEqual(t, e2nccul[0].E2NodeComponentConfigUpdate.GetGNbconfigUpdate().GetF1ApconfigUpdate(), []byte("f1Ap"))
	assert.DeepEqual(t, e2nccul[0].E2NodeComponentConfigUpdate.GetGNbconfigUpdate().GetE1ApconfigUpdate(), []byte("e1Ap"))
	assert.Equal(t, int32(e2nccul[1].E2NodeComponentType), int32(e2apies.E2NodeComponentType_E2NODE_COMPONENT_TYPE_E_NB))
	assert.Equal(t, int32(e2nccul[1].E2NodeComponentID.GetE2NodeComponentTypeGnbDu().GetGNbDuId().GetValue()), int32(13))
	assert.DeepEqual(t, e2nccul[1].E2NodeComponentConfigUpdate.GetENbconfigUpdate().GetX2ApconfigUpdate(), []byte("x2"))
	assert.DeepEqual(t, e2nccul[1].E2NodeComponentConfigUpdate.GetENbconfigUpdate().GetS1ApconfigUpdate(), []byte("s1"))
	if transactionID != nil {
		assert.Equal(t, int32(1), *transactionID)
	}
	assert.Equal(t, nodeIdentity.NodeIDLength, 25)
	assert.DeepEqual(t, nodeIdentity.NodeIdentifier, []byte{0x00, 0x00, 0x00, 0x80})
	assert.DeepEqual(t, [3]byte(nodeIdentity.Plmn), [3]uint8{0x00, 0x00, 0x01})
	assert.Equal(t, nodeIdentity.NodeType, types.E2NodeTypeEnGNB)
	if nodeIdentity.DuID != nil {
		assert.Equal(t, *nodeIdentity.DuID, int64(13))
	}
	if nodeIdentity.CuID != nil {
		assert.Equal(t, *nodeIdentity.CuID, int64(2))
	}
}
