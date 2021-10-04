// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

import (
	"encoding/hex"
	"testing"

	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
	"gotest.tools/assert"
)

func createRicSubscriptionWithCauseItem() (*e2appducontents.RicsubscriptionWithCauseItem, error) {

	rswci := &e2appducontents.RicsubscriptionWithCauseItem{
		RanFunctionId: &e2apies.RanfunctionId{
			Value: 4094,
		},
		RicRequestId: &e2apies.RicrequestId{
			RicInstanceId:  1,
			RicRequestorId: 2,
		},
		Cause: &e2apies.Cause{
			Cause: &e2apies.Cause_E2Node{
				E2Node: e2apies.CauseE2Node_CAUSE_E2NODE_E2NODE_COMPONENT_UNKNOWN,
			},
		},
	}

	return rswci, nil
}

func Test_RicSubscriptionWithCauseItem(t *testing.T) {

	rswci, err := createRicSubscriptionWithCauseItem()
	assert.NilError(t, err)

	xer, err := xerEncodeRicSubscriptionWithCauseItem(rswci)
	assert.NilError(t, err)
	t.Logf("RicSubscriptionWithCauseItem XER\n%s", xer)

	per, err := perEncodeRicSubscriptionWithCauseItem(rswci)
	assert.NilError(t, err)
	t.Logf("RicSubscriptionWithCauseItem PER\n%v", hex.Dump(per))

	// Now reverse the XER
	rswciReversed, err := xerDecodeRicSubscriptionWithCauseItem(xer)
	assert.NilError(t, err)
	assert.Assert(t, rswciReversed != nil)
	t.Logf("RicSubscriptionWithCauseItem decoded from XER is \n%v", rswciReversed)
	assert.Equal(t, rswci.String(), rswciReversed.String())

	// Now reverse the PER
	rswciReversedFromPer, err := perDecodeRicSubscriptionWithCauseItem(per)
	assert.NilError(t, err)
	assert.Assert(t, rswciReversedFromPer != nil)
	t.Logf("RicSubscriptionWithCauseItem decoded from PER is \n%v", rswciReversedFromPer)
	assert.Equal(t, rswci.String(), rswciReversedFromPer.String())
}
