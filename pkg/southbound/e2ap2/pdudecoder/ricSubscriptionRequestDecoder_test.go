// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdudecoder

import (
	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/asn1cgo"
	"gotest.tools/assert"
	"io/ioutil"
	"testing"
)

func Test_DecodeRicSubscriptionRequestPdu(t *testing.T) {
	ricSubscriptionRequestXer, err := ioutil.ReadFile("../test/RICsubscriptionRequest.xml")
	assert.NilError(t, err, "Unexpected error when loading file")
	e2apPdu, err := asn1cgo.XerDecodeE2apPdu(ricSubscriptionRequestXer)
	assert.NilError(t, err)

	ricReq, ranFuncID, ricEventDef, ricActionsToBeSetup, err := DecodeRicSubscriptionRequestPdu(e2apPdu)
	assert.NilError(t, err)
	assert.Equal(t, 1, int(ricReq.RequestorID))
	assert.Equal(t, 2, int(ricReq.InstanceID))
	assert.Equal(t, 1, int(ranFuncID))
	assert.Equal(t, 20, len(ricEventDef))
	assert.Equal(t, 1, len(ricActionsToBeSetup))

	ra1 := ricActionsToBeSetup[0]
	assert.Equal(t, 1, int(ra1.RicActionID))
	assert.Equal(t, e2apies.RictimeToWait_RICTIME_TO_WAIT_W50MS, ra1.Ricttw)
}
