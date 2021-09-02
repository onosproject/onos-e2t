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

func Test_DecodeRicSubscriptionDeleteRequestPdu(t *testing.T) {
	ricSubscriptionDeleteRequestXer, err := ioutil.ReadFile("../test/RICsubscriptionDeleteRequest.xml")
	assert.NilError(t, err, "Unexpected error when loading file")
	e2apPdu, err := asn1cgo.XerDecodeE2apPdu(ricSubscriptionDeleteRequestXer)
	assert.NilError(t, err)

	ricReq, ranFuncID, err := DecodeRicSubscriptionDeleteRequestPdu(e2apPdu)
	assert.NilError(t, err)
	assert.Equal(t, 1, int(ricReq.RequestorID))
	assert.Equal(t, 2, int(ricReq.InstanceID))
	assert.Equal(t, 1, int(ranFuncID))
}
