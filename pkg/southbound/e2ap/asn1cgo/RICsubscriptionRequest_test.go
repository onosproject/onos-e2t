// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

import (
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2apies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/pdubuilder"
	"gotest.tools/assert"
	"testing"
)

func Test_xerEncodeRICsubscriptionRequest(t *testing.T) {
	var ricAction = e2apies.RicactionType_RICACTION_TYPE_POLICY
	var ricSubsequentAction = e2apies.RicsubsequentActionType_RICSUBSEQUENT_ACTION_TYPE_CONTINUE
	var ricttw = e2apies.RictimeToWait_RICTIME_TO_WAIT_ZERO

	e2ApPduRsr, err := pdubuilder.CreateRicSubscriptionRequestE2apPdu(21, 22,
		9, 15, ricAction, ricSubsequentAction, ricttw, []byte{29}, []byte{17})
	assert.NilError(t, err)
	xer, err := xerEncodeRICsubscriptionRequest(
		e2ApPduRsr.GetInitiatingMessage().GetProcedureCode().GetRicSubscription().GetInitiatingMessage())
	assert.NilError(t, err)
	t.Logf("XER RICsubscriptionRequest\n%s", xer)

	per, err := perEncodeRICsubscriptionRequest(
		e2ApPduRsr.GetInitiatingMessage().GetProcedureCode().GetRicSubscription().GetInitiatingMessage())
	assert.NilError(t, err)
	t.Logf("PER RICsubscriptionRequest\n%s", per)
}
