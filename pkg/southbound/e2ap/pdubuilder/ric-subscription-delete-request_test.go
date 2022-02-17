// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdubuilder

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/encoder"
	"testing"

	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
	"gotest.tools/assert"
)

func TestRicSubscriptionDeleteRequest(t *testing.T) {
	newE2apPdu, err := CreateRicSubscriptionDeleteRequestE2apPdu(
		types.RicRequest{RequestorID: 1, InstanceID: 2},
		3)

	assert.NilError(t, err)
	assert.Assert(t, newE2apPdu != nil)

	perNew, err := encoder.PerEncodeE2ApPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("RicSubscriptionDeleteRequest E2AP PDU PER with Go APER library\n%v", hex.Dump(perNew))

	//Comparing reference PER bytes with Go APER library produced
	//assert.DeepEqual(t, per, perNew)

	result, err := encoder.PerDecodeE2ApPdu(perNew)
	assert.NilError(t, err)
	assert.DeepEqual(t, newE2apPdu.String(), result.String())
}
