// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdudecoder

import (
	"io/ioutil"
	"testing"

	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-ies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap201/asn1cgo"
	"gotest.tools/assert"
)

func Test_DecodeE2SetupResponsePdu(t *testing.T) {
	e2setupResponseXer, err := ioutil.ReadFile("../test/E2setupResponse.xml")
	assert.NilError(t, err, "Unexpected error when loading file")
	e2apPdu, err := asn1cgo.XerDecodeE2apPdu(e2setupResponseXer)
	assert.NilError(t, err)

	transactionID, ricIdentity, ranFunctionsAccepted, ranFunctionsRejected, e2nccual, err := DecodeE2SetupResponsePdu(e2apPdu)
	assert.NilError(t, err)
	//assert.Assert(t, ricIdentity != nil) //Commented due to the Linters (v1.34.1) error - possible nil pointer dereference (https://staticcheck.io/docs/checks#SA5011) on lines 23, 24 & 25
	assert.DeepEqual(t, []byte{0x79, 0x78, 0x70}, []byte{ricIdentity.PlmnID[0], ricIdentity.PlmnID[1], ricIdentity.PlmnID[2]})
	assert.Equal(t, 20, int(ricIdentity.RicIdentifier.RicIdentifierLen))
	assert.DeepEqual(t, []byte{0x4d, 0x20, 0x00}, []byte(ricIdentity.RicIdentifier.RicIdentifierValue))

	assert.Equal(t, 2, len(ranFunctionsAccepted))
	rfa100, ok := ranFunctionsAccepted[100]
	assert.Assert(t, ok, "expected a key '100'")
	assert.Equal(t, 2, int(rfa100))
	rfa200, ok := ranFunctionsAccepted[200]
	assert.Assert(t, ok, "expected a key '200'")
	assert.Equal(t, 2, int(rfa200))

	assert.Equal(t, 2, len(ranFunctionsRejected))
	rfr101, ok := ranFunctionsRejected[101]
	assert.Assert(t, ok, "expected a key '101'")
	assert.Equal(t, "CAUSE_MISC_OM_INTERVENTION", rfr101.GetMisc().String())
	rfr102, ok := ranFunctionsRejected[102]
	assert.Assert(t, ok, "expected a key '102'")
	assert.Equal(t, "CAUSE_PROTOCOL_TRANSFER_SYNTAX_ERROR", rfr102.GetProtocol().String())

	if transactionID != nil {
		assert.Equal(t, int32(1), *transactionID)
	}
	assert.Equal(t, int32(e2nccual[0].E2NodeComponentType), int32(e2ap_ies.E2NodeComponentType_E2NODE_COMPONENT_TYPE_G_NB))
	assert.Equal(t, int32(e2nccual[0].E2NodeComponentID.GetE2NodeComponentTypeGnbCuUp().GetGNbCuUpId().GetValue()), int32(21))
	assert.Equal(t, e2nccual[0].E2NodeComponentConfigUpdateAck.UpdateOutcome, int32(1))
	assert.Equal(t, int32(e2nccual[0].E2NodeComponentConfigUpdateAck.FailureCause.GetProtocol()), int32(e2ap_ies.CauseProtocol_CAUSE_PROTOCOL_TRANSFER_SYNTAX_ERROR))
	assert.Equal(t, int32(e2nccual[1].E2NodeComponentType), int32(e2ap_ies.E2NodeComponentType_E2NODE_COMPONENT_TYPE_E_NB))
	assert.Equal(t, int32(e2nccual[1].E2NodeComponentID.GetE2NodeComponentTypeGnbDu().GetGNbDuId().GetValue()), int32(13))
	assert.Equal(t, e2nccual[1].E2NodeComponentConfigUpdateAck.UpdateOutcome, int32(1))
	assert.Equal(t, int32(e2nccual[1].E2NodeComponentConfigUpdateAck.FailureCause.GetProtocol()), int32(e2ap_ies.CauseProtocol_CAUSE_PROTOCOL_ABSTRACT_SYNTAX_ERROR_FALSELY_CONSTRUCTED_MESSAGE))
}

func Test_DecodeE2SetupResponsePduNoOptional(t *testing.T) {
	e2setupResponseXer, err := ioutil.ReadFile("../test/E2setupResponse2.xml")
	assert.NilError(t, err, "Unexpected error when loading file")
	e2apPdu, err := asn1cgo.XerDecodeE2apPdu(e2setupResponseXer)
	assert.NilError(t, err)

	transactionID, ricIdentity, ranFunctionsAccepted, ranFunctionsRejected, e2nccual, err := DecodeE2SetupResponsePdu(e2apPdu)
	assert.NilError(t, err)
	//assert.Assert(t, ricIdentity != nil) //Commented due to the Linters (v1.34.1) error - possible nil pointer dereference (https://staticcheck.io/docs/checks#SA5011) on lines 54, 55 & 56
	assert.DeepEqual(t, []byte{0x00, 0x02, 0x10}, []byte{ricIdentity.PlmnID[0], ricIdentity.PlmnID[1], ricIdentity.PlmnID[2]})
	assert.Equal(t, 20, int(ricIdentity.RicIdentifier.RicIdentifierLen))
	assert.DeepEqual(t, []byte{0x01, 0x00, 0x00}, []byte(ricIdentity.RicIdentifier.RicIdentifierValue))

	assert.Equal(t, 0, len(ranFunctionsAccepted))
	assert.Equal(t, 0, len(ranFunctionsRejected))
	assert.Equal(t, 0, len(e2nccual))
	if transactionID != nil {
		assert.Equal(t, int32(11), *transactionID)
	}
}
