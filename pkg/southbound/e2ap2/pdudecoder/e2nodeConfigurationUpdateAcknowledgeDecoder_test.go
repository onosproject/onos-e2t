// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdudecoder

import (
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2apies"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/asn1cgo"
	"gotest.tools/assert"
	"io/ioutil"
	"testing"
)

func Test_DecodeE2nodeConfigurationUpdateAcknowledgePdu(t *testing.T) {
	e2ncuaXer, err := ioutil.ReadFile("../test/E2nodeConfigurationUpdateAcknowledge.xml")
	assert.NilError(t, err, "Unexpected error when loading file")
	e2apPdu, err := asn1cgo.XerDecodeE2apPdu(e2ncuaXer)
	assert.NilError(t, err)

	e2nccual, err := DecodeE2nodeConfigurationUpdateAcknowledgePdu(e2apPdu)
	assert.NilError(t, err)
	//assert.Assert(t, ricIdentity != nil) //Commented due to the Linters (v1.34.1) error - possible nil pointer dereference (https://staticcheck.io/docs/checks#SA5011) on lines 23, 24 & 25

	assert.Equal(t, int32(e2nccual[0].E2NodeComponentType), int32(e2ap_ies.E2NodeComponentType_E2NODE_COMPONENT_TYPE_G_NB))
	assert.Equal(t, int32(e2nccual[0].E2NodeComponentID.GetE2NodeComponentTypeGnbCuUp().GetGNbCuUpId().GetValue()), int32(21))
	assert.Equal(t, e2nccual[0].E2NodeComponentConfigUpdateAck.UpdateOutcome, int32(1))
	assert.Equal(t, int32(e2nccual[0].E2NodeComponentConfigUpdateAck.FailureCause.GetProtocol()), int32(e2apies.CauseProtocol_CAUSE_PROTOCOL_TRANSFER_SYNTAX_ERROR))
	assert.Equal(t, int32(e2nccual[1].E2NodeComponentType), int32(e2ap_ies.E2NodeComponentType_E2NODE_COMPONENT_TYPE_E_NB))
	assert.Equal(t, int32(e2nccual[1].E2NodeComponentID.GetE2NodeComponentTypeGnbDu().GetGNbDuId().GetValue()), int32(13))
	assert.Equal(t, e2nccual[1].E2NodeComponentConfigUpdateAck.UpdateOutcome, int32(1))
	assert.Equal(t, int32(e2nccual[1].E2NodeComponentConfigUpdateAck.FailureCause.GetProtocol()), int32(e2apies.CauseProtocol_CAUSE_PROTOCOL_ABSTRACT_SYNTAX_ERROR_FALSELY_CONSTRUCTED_MESSAGE))
}
