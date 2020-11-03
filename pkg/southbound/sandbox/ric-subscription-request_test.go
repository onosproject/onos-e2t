// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
package sandbox

import (
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2apies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2proxy/orane2"
	"gotest.tools/assert"
	"testing"
)

func TestRicSubscriptionRequest(t *testing.T) {
	var ricAction = e2apies.RicactionType_RICACTION_TYPE_POLICY
	var ricSubsequentAction = e2apies.RicsubsequentActionType_RICSUBSEQUENT_ACTION_TYPE_CONTINUE
	var ricttw = e2apies.RictimeToWait_RICTIME_TO_WAIT_ZERO
	newE2apPdu, err := CreateRicSubscriptionRequestE2apPdu(21, 22,
		9, 15, ricAction, ricSubsequentAction, ricttw, 29, 17)
	assert.NilError(t, err)
	assert.Assert(t, newE2apPdu != nil)

	xer, err := orane2.XerEncodeE2apPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("RicSubscriptionRequest E2AP PDU\n%s", xer)

}
