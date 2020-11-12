// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package filter

import (
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2apies"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2appdudescriptions"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/pdubuilder"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFilterRicSubscriptionResponse(t *testing.T) {
	assert.True(t, RicSubscription(newRequestID())(newSubscriptionResponse()))
}

func TestFilterRicIndication(t *testing.T) {
	assert.True(t, RicIndication(newRequestID())(newIndication()))
}

func newRequestID() *e2apies.RicrequestId {
	return &e2apies.RicrequestId{
		RicRequestorId: 1,
		RicInstanceId:  2,
	}
}

func newSubscriptionResponse() *e2appdudescriptions.E2ApPdu {
	res, _ := pdubuilder.CreateRicSubscriptionResponseE2apPdu(&types.RicRequest{RequestorID: 1, InstanceID: 2}, 2, []*types.RicActionID{})
	return res
}

func newIndication() *e2appdudescriptions.E2ApPdu {
	res, _ := pdubuilder.RicIndicationE2apPdu(1, 2, 3, 4, 1, e2apies.RicindicationType_RICINDICATION_TYPE_REPORT, "foo", "bar", "baz")
	return res
}
