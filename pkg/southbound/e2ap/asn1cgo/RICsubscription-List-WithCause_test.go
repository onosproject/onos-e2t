// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package asn1cgo

import (
	"encoding/hex"
	v2 "github.com/onosproject/onos-e2t/api/e2ap/v2"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-commondatatypes"
	"testing"

	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
	"gotest.tools/assert"
)

func createRicSubscriptionWithCauseList() (*e2appducontents.RicsubscriptionListWithCause, error) {

	rswcl := e2appducontents.RicsubscriptionListWithCause{
		Value: make([]*e2appducontents.RicsubscriptionWithCauseItemIes, 0),
	}

	rswci := &e2appducontents.RicsubscriptionWithCauseItemIes{
		Id:          int32(v2.ProtocolIeIDRICsubscriptionWithCauseItem),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
		Value: &e2appducontents.RicsubscriptionWithCauseItem{
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
		},
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
	}

	rswcl.Value = append(rswcl.Value, rswci)

	return &rswcl, nil
}

func Test_RicSubscriptionWithCauseList(t *testing.T) {

	rswci, err := createRicSubscriptionWithCauseList()
	assert.NilError(t, err)

	xer, err := xerEncodeRicSubscriptionListWithCause(rswci)
	assert.NilError(t, err)
	t.Logf("RicSubscriptionWithCauseList XER\n%s", xer)

	per, err := perEncodeRicSubscriptionListWithCause(rswci)
	assert.NilError(t, err)
	t.Logf("RicSubscriptionWithCauseList PER\n%v", hex.Dump(per))

	// Now reverse the XER
	rswciReversed, err := xerDecodeRicSubscriptionListWithCause(xer)
	assert.NilError(t, err)
	assert.Assert(t, rswciReversed != nil)
	t.Logf("RicSubscriptionWithCauseList decoded from XER is \n%v", rswciReversed)
	assert.Equal(t, rswci.String(), rswciReversed.String())

	// Now reverse the PER
	rswciReversedFromPer, err := perDecodeRicSubscriptionListWithCause(per)
	assert.NilError(t, err)
	assert.Assert(t, rswciReversedFromPer != nil)
	t.Logf("RicSubscriptionWithCauseList decoded from PER is \n%v", rswciReversedFromPer)
	assert.Equal(t, rswci.String(), rswciReversedFromPer.String())
}
