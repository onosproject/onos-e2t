// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
package pdubuilder

import (
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2apies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/asn1cgo"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
	"gotest.tools/assert"
	"testing"
)

func TestRicSubscriptionDeleteFailure(t *testing.T) {
	newE2apPdu, err := CreateRicSubscriptionDeleteFailureE2apPdu(&types.RicRequest{
		RequestorID: 22,
		InstanceID:  6,
	}, 9,
		&e2apies.Cause{
			Cause: &e2apies.Cause_Transport{
				Transport: e2apies.CauseTransport_CAUSE_TRANSPORT_TRANSPORT_RESOURCE_UNAVAILABLE,
			},
		}, 200, 0)
	assert.NilError(t, err)
	assert.Assert(t, newE2apPdu != nil)

	xer, err := asn1cgo.XerEncodeE2apPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("RicSubscriptionDeleteFailure E2AP PDU XER\n%s", string(xer))

	per, err := asn1cgo.PerEncodeE2apPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("RicSubscriptionDeleteFailure E2AP PDU PER\n%v", per)
}
