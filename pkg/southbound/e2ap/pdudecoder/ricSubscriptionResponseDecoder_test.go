// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdudecoder

import (
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/pdubuilder"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
	"gotest.tools/assert"
	"testing"
)

func Test_DecodeRicSubscriptionResponsePdu(t *testing.T) {
	ricActionsNotAdmittedList := make(map[types.RicActionID]*e2ap_ies.Cause)
	ricActionsNotAdmittedList[100] = &e2ap_ies.Cause{
		Cause: &e2ap_ies.Cause_Transport{
			Transport: e2ap_ies.CauseTransport_CAUSE_TRANSPORT_TRANSPORT_RESOURCE_UNAVAILABLE,
		},
	}
	ricActionsNotAdmittedList[200] = &e2ap_ies.Cause{
		Cause: &e2ap_ies.Cause_Misc{
			Misc: e2ap_ies.CauseMisc_CAUSE_MISC_HARDWARE_FAILURE,
		},
	}

	var ricActionAdmitted10 types.RicActionID = 10
	var ricActionAdmitted20 types.RicActionID = 20
	e2apPdu, err := pdubuilder.CreateRicSubscriptionResponseE2apPdu(&types.RicRequest{
		RequestorID: 22,
		InstanceID:  6,
	}, 9, []*types.RicActionID{&ricActionAdmitted10, &ricActionAdmitted20})
	assert.NilError(t, err)
	assert.Assert(t, e2apPdu != nil)
	e2apPdu.GetSuccessfulOutcome().GetValue().GetRicSubscription().
		SetRicActionNotAdmitted(ricActionsNotAdmittedList)

	rfID, rrID, ricActionIDs, causes, err := DecodeRicSubscriptionResponsePdu(e2apPdu)
	assert.NilError(t, err)
	assert.Equal(t, 9, int(*rfID))

	assert.Assert(t, causes != nil)
	for id, cause := range causes {
		switch id {
		case 100:
			assert.Equal(t, e2ap_ies.CauseTransport_CAUSE_TRANSPORT_TRANSPORT_RESOURCE_UNAVAILABLE.Number(), cause.GetTransport().Number())
		case 200:
			assert.Equal(t, e2ap_ies.CauseMisc_CAUSE_MISC_HARDWARE_FAILURE.Number(), cause.GetMisc().Number())
		default:
			assert.Assert(t, false, "unexpected cause %d", id)
		}
	}

	assert.Equal(t, 22, int(rrID.RequestorID))
	assert.Equal(t, 6, int(rrID.InstanceID))

	assert.Equal(t, 2, len(ricActionIDs))
	assert.Equal(t, 10, int(ricActionIDs[0]))
	assert.Equal(t, 20, int(ricActionIDs[1]))
}
