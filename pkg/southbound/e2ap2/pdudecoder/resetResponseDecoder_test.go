// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdudecoder

import (
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2ap-commondatatypes"
	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/asn1cgo"
	"gotest.tools/assert"
	"io/ioutil"
	"testing"
)

func Test_DecodeResetResponsePdu(t *testing.T) {
	rrXer, err := ioutil.ReadFile("../test/ResetResponse.xml")
	assert.NilError(t, err, "Unexpected error when loading file")
	e2apPdu, err := asn1cgo.XerDecodeE2apPdu(rrXer)
	assert.NilError(t, err)

	pr, crit, tm, cdrID, diags, err := DecodeResetResponsePdu(e2apPdu)
	assert.NilError(t, err)
	//assert.Assert(t, ricIdentity != nil) //Commented due to the Linters (v1.34.1) error - possible nil pointer dereference (https://staticcheck.io/docs/checks#SA5011) on lines 23, 24 & 25

	assert.Equal(t, int32(*pr), int32(8))
	assert.Equal(t, int32(*crit), int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE))
	assert.Equal(t, int32(*tm), int32(e2ap_commondatatypes.TriggeringMessage_TRIGGERING_MESSAGE_UNSUCCESSFULL_OUTCOME))
	assert.Equal(t, int32(cdrID.InstanceID), int32(20))
	assert.Equal(t, int32(cdrID.RequestorID), int32(10))
	assert.Equal(t, int32(diags[0].IEId), int32(30))
	assert.Equal(t, int32(diags[0].IECriticality), int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE))
	assert.Equal(t, int32(diags[0].TypeOfError), int32(e2apies.TypeOfError_TYPE_OF_ERROR_MISSING))
}
