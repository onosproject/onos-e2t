// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdudecoder

import (
	"io/ioutil"
	"testing"

	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-ies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/asn1cgo"
	"gotest.tools/assert"
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
	assert.Equal(t, 3, int(ranFuncID))
	assert.Equal(t, 2, len(ricEventDef))
	assert.Equal(t, 2, len(ricActionsToBeSetup))

	assert.Equal(t, 100, int(ricActionsToBeSetup[0].RicActionID))
	assert.Equal(t, e2apies.RictimeToWait_RICTIME_TO_WAIT_W5MS, ricActionsToBeSetup[0].Ricttw)
	assert.Equal(t, e2apies.RicsubsequentActionType_RICSUBSEQUENT_ACTION_TYPE_CONTINUE, ricActionsToBeSetup[0].RicSubsequentAction)
}
