// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdudecoder

import (
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2apies"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-commondatatypes"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/asn1cgo"
	"gotest.tools/assert"
	"io/ioutil"
	"testing"
)

func Test_DecodeRicServiceUpdateFailurePdu(t *testing.T) {
	rsufXer, err := ioutil.ReadFile("../test/RICserviceUpdateFailure.xml")
	assert.NilError(t, err, "Unexpected error when loading file")
	e2apPdu, err := asn1cgo.XerDecodeE2apPdu(rsufXer)
	assert.NilError(t, err)

	causes, ttw, pr, crit, tm, cdrID, diags, err := DecodeRicServiceUpdateFailurePdu(e2apPdu)
	assert.NilError(t, err)
	//assert.Assert(t, ricIdentity != nil) //Commented due to the Linters (v1.34.1) error - possible nil pointer dereference (https://staticcheck.io/docs/checks#SA5011) on lines 23, 24 & 25

	assert.Assert(t, causes != nil)
	if causes != nil {
		for id, cause := range *causes {
			switch id {
			case 101:
				assert.Equal(t, e2ap_ies.CauseMisc_CAUSE_MISC_HARDWARE_FAILURE, cause.GetMisc())
			case 102:
				assert.Equal(t, e2ap_ies.CauseProtocol_CAUSE_PROTOCOL_SEMANTIC_ERROR, cause.GetProtocol())
			default:
				assert.Assert(t, false, "unexpected cause %d", id)
			}
		}
	}
	assert.Equal(t, int32(*ttw), int32(e2ap_ies.TimeToWait_TIME_TO_WAIT_V2S))
	assert.Equal(t, int32(*pr), int32(8))
	assert.Equal(t, int32(*crit), int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE))
	assert.Equal(t, int32(*tm), int32(e2ap_commondatatypes.TriggeringMessage_TRIGGERING_MESSAGE_UNSUCCESSFULL_OUTCOME))
	assert.Equal(t, int32(cdrID.InstanceID), int32(20))
	assert.Equal(t, int32(cdrID.RequestorID), int32(10))
	assert.Equal(t, int32(diags[0].IEId), int32(30))
	assert.Equal(t, int32(diags[0].IECriticality), int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE))
	assert.Equal(t, int32(diags[0].TypeOfError), int32(e2apies.TypeOfError_TYPE_OF_ERROR_MISSING))
}
