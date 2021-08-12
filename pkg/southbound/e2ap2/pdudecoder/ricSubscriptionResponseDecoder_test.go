// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdudecoder

import (
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-ies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap2/asn1cgo"
	"gotest.tools/assert"
	"io/ioutil"
	"testing"
)

func Test_DecodeRicSubscriptionResponsePdu(t *testing.T) {
	ricSubscriptionResponseXer, err := ioutil.ReadFile("../test/RICsubscriptionResponse.xml")
	assert.NilError(t, err, "Unexpected error when loading file")
	e2apPdu, err := asn1cgo.XerDecodeE2apPdu(ricSubscriptionResponseXer)
	assert.NilError(t, err)

	rfID, rrID, ricActionIDs, causes, err := DecodeRicSubscriptionResponsePdu(e2apPdu)
	assert.NilError(t, err)
	//assert.Assert(t, rfID != nil) //Commented due to the Linters (v1.34.1) error - possible nil pointer dereference (https://staticcheck.io/docs/checks#SA5011) on lines 23
	assert.Equal(t, 20, int(*rfID))

	//ToDo - adjust verification of RicActionsNotAdmittedList
	assert.Assert(t, causes != nil)
	if causes != nil {
		for id, cause := range causes {
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

	//assert.Assert(t, rrID != nil) //Commented due to the Linters (v1.34.1) error - possible nil pointer dereference (https://staticcheck.io/docs/checks#SA5011) on lines 26 & 27
	assert.Equal(t, 22, int(rrID.RequestorID))
	assert.Equal(t, 6, int(rrID.InstanceID))

	assert.Equal(t, 1, len(ricActionIDs))
	assert.Equal(t, 5, int(ricActionIDs[0]))
}
