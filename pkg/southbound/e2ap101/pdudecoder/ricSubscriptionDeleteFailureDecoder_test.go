// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdudecoder

import (
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/asn1cgo"
	"gotest.tools/assert"
	"io/ioutil"
	"testing"
)

func Test_DecodeRicSubscriptionDeleteFailurePdu(t *testing.T) {
	ricSubscriptionDeleteFailureXer, err := ioutil.ReadFile("../test/RICsubscriptionDeleteFailure.xml")
	assert.NilError(t, err, "Unexpected error when loading file")
	e2apPdu, err := asn1cgo.XerDecodeE2apPdu(ricSubscriptionDeleteFailureXer)
	assert.NilError(t, err)

	rfID, rrID, err := DecodeRicSubscriptionDeleteFailurePdu(e2apPdu)
	assert.NilError(t, err)
	//assert.Assert(t, rfID != nil) //Commented due to the Linters (v1.34.1) error - possible nil pointer dereference (https://staticcheck.io/docs/checks#SA5011) on line 23
	assert.Equal(t, 9, int(*rfID))

	//assert.Assert(t, rrID != nil) //Commented due to the Linters (v1.34.1) error - possible nil pointer dereference (https://staticcheck.io/docs/checks#SA5011) on lines 26 & 27
	assert.Equal(t, 22, int(rrID.RequestorID))
	assert.Equal(t, 6, int(rrID.InstanceID))

}
