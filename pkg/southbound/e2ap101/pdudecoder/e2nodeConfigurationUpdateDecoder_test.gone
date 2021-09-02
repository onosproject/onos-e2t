// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdudecoder

import (
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/asn1cgo"
	"gotest.tools/assert"
	"io/ioutil"
	"testing"
)

func Test_DecodeE2nodeConfigurationUpdatePdu(t *testing.T) {
	e2ncuXer, err := ioutil.ReadFile("../test/E2nodeConfigurationUpdate.xml")
	assert.NilError(t, err, "Unexpected error when loading file")
	e2apPdu, err := asn1cgo.XerDecodeE2apPdu(e2ncuXer)
	assert.NilError(t, err)

	e2nccul, err := DecodeE2nodeConfigurationUpdatePdu(e2apPdu)
	assert.NilError(t, err)
	//assert.Assert(t, ricIdentity != nil) //Commented due to the Linters (v1.34.1) error - possible nil pointer dereference (https://staticcheck.io/docs/checks#SA5011) on lines 23, 24 & 25

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
}
