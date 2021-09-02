// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdudecoder

import (
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta2"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-commondatatypes"
	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/asn1cgo"
	"gotest.tools/assert"
	"io/ioutil"
	"testing"
)

func Test_DecodeRicSubscriptionFailurePdu(t *testing.T) {
	ricSubscriptionFailureXer, err := ioutil.ReadFile("../test/RICsubscriptionFailure.xml")
	assert.NilError(t, err, "Unexpected error when loading file")
	e2apPdu, err := asn1cgo.XerDecodeE2apPdu(ricSubscriptionFailureXer)
	assert.NilError(t, err)

	rrID, rfID, pc, crit, tm, critReq, causes, diags, err := DecodeRicSubscriptionFailurePdu(e2apPdu)
	assert.NilError(t, err)
	//assert.Assert(t, rfID != nil) //Commented due to the Linters (v1.34.1) error - possible nil pointer dereference (https://staticcheck.io/docs/checks#SA5011) on line 26
	assert.Equal(t, 9, int(*rfID))

	//assert.Assert(t, rrID != nil) //Commented due to the Linters (v1.34.1) error - possible nil pointer dereference (https://staticcheck.io/docs/checks#SA5011) on lines 29 & 30
	assert.Equal(t, 22, int(rrID.RequestorID))
	assert.Equal(t, 6, int(rrID.InstanceID))

	assert.Equal(t, v1beta2.ProcedureCodeIDRICsubscription, pc)
	assert.Equal(t, e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE, crit)
	assert.Equal(t, e2ap_commondatatypes.TriggeringMessage_TRIGGERING_MESSAGE_UNSUCCESSFULL_OUTCOME, tm)

	// TODO: Should be 10
	assert.Equal(t, 10, int(critReq.RequestorID))
	// TODO: Should be 20
	assert.Equal(t, 20, int(critReq.InstanceID))

	assert.Assert(t, causes != nil)
	for id, cause := range causes {
		switch id {
		case 100:
			assert.Equal(t, e2apies.CauseTransport_CAUSE_TRANSPORT_TRANSPORT_RESOURCE_UNAVAILABLE, cause.GetTransport())
		case 200:
			assert.Equal(t, e2apies.CauseMisc_CAUSE_MISC_HARDWARE_FAILURE, cause.GetMisc())
		default:
			assert.Assert(t, false, "unexpected cause %d", id)
		}
	}
	assert.Assert(t, diags == nil)

}
