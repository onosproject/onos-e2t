// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdudecoder

import (
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-ies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap2/asn1cgo"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap2/types"
	"gotest.tools/assert"
	"io/ioutil"
	"testing"
)

func Test_DecodeE2SetupRequestPdu(t *testing.T) {
	e2setupRequestXer, err := ioutil.ReadFile("../test/E2setupRequest-eNB.xml")
	assert.NilError(t, err, "Unexpected error when loading file")
	e2apPdu, err := asn1cgo.XerDecodeE2apPdu(e2setupRequestXer)
	assert.NilError(t, err)

	transactionID, identifier, ranFunctions, e2nccul, err := DecodeE2SetupRequestPdu(e2apPdu)
	assert.NilError(t, err)
	//assert.Assert(t, identifier != nil) //Commented due to the Linters (v1.34.1) error - possible nil pointer dereference (https://staticcheck.io/docs/checks#SA5011) on lines 24, 25 & 26
	assert.DeepEqual(t, []byte{0x00, 0x02, 0x10}, []byte{identifier.Plmn[0], identifier.Plmn[1], identifier.Plmn[2]})
	assert.Equal(t, types.E2NodeTypeENB, identifier.NodeType)
	assert.DeepEqual(t, []byte{0x00, 0xE0, 0x00}, identifier.NodeIdentifier)

	//t.Logf("Node ID is %x\n", identifier.NodeIdentifier)
	nodeID := GetE2NodeID(identifier.NodeIdentifier, 22)
	t.Logf("Node ID is %s\n", nodeID)

	//assert.Assert(t, ranFunctions != nil) //Commented due to the Linters (v1.34.1) error - possible nil pointer dereference (https://staticcheck.io/docs/checks#SA5011) on line 29
	assert.Equal(t, 1, len(*ranFunctions))
	rf0 := (*ranFunctions)[20]
	assert.Equal(t, 10, int(rf0.Revision))
	assert.DeepEqual(t, []byte("abc"), []byte(rf0.OID))

	assert.Equal(t, int32(e2nccul[0].E2NodeComponentType), int32(e2ap_ies.E2NodeComponentType_E2NODE_COMPONENT_TYPE_G_NB))
	assert.Equal(t, int32(e2nccul[0].E2NodeComponentID.GetE2NodeComponentTypeGnbCuUp().GetGNbCuUpId().GetValue()), int32(21))
	assert.Equal(t, e2nccul[0].E2NodeComponentConfigUpdate.GetGNbconfigUpdate().GetNgApconfigUpdate(), "ngAp")
	assert.Equal(t, e2nccul[0].E2NodeComponentConfigUpdate.GetGNbconfigUpdate().GetXnApconfigUpdate(), "xnAp")
	assert.Equal(t, e2nccul[0].E2NodeComponentConfigUpdate.GetGNbconfigUpdate().GetF1ApconfigUpdate(), "f1Ap")
	assert.Equal(t, e2nccul[0].E2NodeComponentConfigUpdate.GetGNbconfigUpdate().GetE1ApconfigUpdate(), "e1Ap")
	assert.Equal(t, int32(e2nccul[1].E2NodeComponentType), int32(e2ap_ies.E2NodeComponentType_E2NODE_COMPONENT_TYPE_E_NB))
	assert.Equal(t, int32(e2nccul[1].E2NodeComponentID.GetE2NodeComponentTypeGnbDu().GetGNbDuId().GetValue()), int32(13))
	assert.Equal(t, e2nccul[1].E2NodeComponentConfigUpdate.GetENbconfigUpdate().GetX2ApconfigUpdate(), "x2")
	assert.Equal(t, e2nccul[1].E2NodeComponentConfigUpdate.GetENbconfigUpdate().GetS1ApconfigUpdate(), "s1")

	//ToDo - change Transaction ID to real one
	assert.Equal(t, int32(0), transactionID)
}

func Test_GetE2NodeID(t *testing.T) {
	nodeID := []byte{0, 0, 0, 0, 0, 0x01, 0x45, 0x4c}
	id := GetE2NodeID(nodeID, 22)
	assert.Equal(t, id, "5153")

	nodeID2 := []byte{0x01, 0x45, 0x4c}
	id = GetE2NodeID(nodeID2, 22)
	assert.Equal(t, id, "5153")
}

func Test_DecodeE2SetupRequestPduCuDuIDs(t *testing.T) {
	e2setupRequestXer, err := ioutil.ReadFile("../test/E2setupRequest-gNB.xml")
	assert.NilError(t, err, "Unexpected error when loading file")
	e2apPdu, err := asn1cgo.XerDecodeE2apPdu(e2setupRequestXer)
	assert.NilError(t, err)

	transactionID, identifier, ranFunctions, e2nccul, err := DecodeE2SetupRequestPdu(e2apPdu)
	assert.NilError(t, err)
	//assert.Assert(t, identifier != nil) //Commented due to the Linters (v1.34.1) error - possible nil pointer dereference (https://staticcheck.io/docs/checks#SA5011) on lines 24, 25 & 26
	assert.DeepEqual(t, []byte{0x4f, 0x4e, 0x46}, []byte{identifier.Plmn[0], identifier.Plmn[1], identifier.Plmn[2]})
	assert.Equal(t, types.E2NodeTypeGNB, identifier.NodeType)
	assert.DeepEqual(t, []byte{0xb5, 0xc6, 0x77, 0x88}, identifier.NodeIdentifier)
	assert.Equal(t, int64(15), *identifier.CuID)
	assert.Equal(t, int64(21), *identifier.DuID)

	//t.Logf("Node ID is %x\n", identifier.NodeIdentifier)
	nodeID := GetE2NodeID(identifier.NodeIdentifier, 30)
	t.Logf("Node ID is %s\n", nodeID)

	//assert.Assert(t, ranFunctions != nil) //Commented due to the Linters (v1.34.1) error - possible nil pointer dereference (https://staticcheck.io/docs/checks#SA5011) on line 29
	assert.Equal(t, 2, len(*ranFunctions))
	rf0 := (*ranFunctions)[100]
	assert.Equal(t, 2, int(rf0.Revision))
	rf1 := (*ranFunctions)[200]
	assert.Equal(t, 3, int(rf1.Revision))

	assert.Equal(t, int32(e2nccul[0].E2NodeComponentType), int32(e2ap_ies.E2NodeComponentType_E2NODE_COMPONENT_TYPE_G_NB))
	assert.Equal(t, int32(e2nccul[0].E2NodeComponentID.GetE2NodeComponentTypeGnbCuUp().GetGNbCuUpId().GetValue()), int32(21))
	assert.Equal(t, e2nccul[0].E2NodeComponentConfigUpdate.GetGNbconfigUpdate().GetNgApconfigUpdate(), "ngAp")
	assert.Equal(t, e2nccul[0].E2NodeComponentConfigUpdate.GetGNbconfigUpdate().GetXnApconfigUpdate(), "xnAp")
	assert.Equal(t, e2nccul[0].E2NodeComponentConfigUpdate.GetGNbconfigUpdate().GetF1ApconfigUpdate(), "f1Ap")
	assert.Equal(t, e2nccul[0].E2NodeComponentConfigUpdate.GetGNbconfigUpdate().GetE1ApconfigUpdate(), "e1Ap")
	assert.Equal(t, int32(e2nccul[1].E2NodeComponentType), int32(e2ap_ies.E2NodeComponentType_E2NODE_COMPONENT_TYPE_E_NB))
	assert.Equal(t, int32(e2nccul[1].E2NodeComponentID.GetE2NodeComponentTypeGnbDu().GetGNbDuId().GetValue()), int32(13))
	assert.Equal(t, e2nccul[1].E2NodeComponentConfigUpdate.GetENbconfigUpdate().GetX2ApconfigUpdate(), "x2")
	assert.Equal(t, e2nccul[1].E2NodeComponentConfigUpdate.GetENbconfigUpdate().GetS1ApconfigUpdate(), "s1")

	//ToDo - change Transaction ID to real one
	assert.Equal(t, int32(0), transactionID)
}
