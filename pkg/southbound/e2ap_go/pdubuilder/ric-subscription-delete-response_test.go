// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
package pdubuilder

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/asn1cgo"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/pdubuilder"
	types1 "github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap_go/encoder"
	"testing"

	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap_go/types"
	"gotest.tools/assert"
)

func TestRicSubscriptionDeleteResponse(t *testing.T) {
	e2apPdu, err := pdubuilder.CreateRicSubscriptionDeleteResponseE2apPdu(&types1.RicRequest{
		RequestorID: 22,
		InstanceID:  6,
	}, 9)
	assert.NilError(t, err)
	assert.Assert(t, e2apPdu != nil)

	per, err := asn1cgo.PerEncodeE2apPdu(e2apPdu)
	assert.NilError(t, err)
	t.Logf("RicSubscriptionDeleteResponse E2AP PDU PER\n%v", hex.Dump(per))

	newE2apPdu, err := CreateRicSubscriptionDeleteResponseE2apPdu(&types.RicRequest{
		RequestorID: 22,
		InstanceID:  6,
	}, 9)
	assert.NilError(t, err)
	assert.Assert(t, newE2apPdu != nil)

	perNew, err := encoder.PerEncodeE2ApPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("RicSubscriptionDeleteResponse E2AP PDU PER with Go APER library\n%v", hex.Dump(perNew))

	//Comparing reference PER bytes with Go APER library produced
	assert.DeepEqual(t, per, perNew)

	result, err := encoder.PerDecodeE2ApPdu(perNew)
	assert.NilError(t, err)
	assert.DeepEqual(t, newE2apPdu.String(), result.String())

	// Decoding the message from the APER bytes produced by CGo
	result11, err := encoder.PerDecodeE2ApPdu(per)
	assert.NilError(t, err)
	assert.DeepEqual(t, newE2apPdu.String(), result11.String())

	result1, err := asn1cgo.PerDecodeE2apPdu(perNew)
	assert.NilError(t, err)
	assert.DeepEqual(t, result1.String(), e2apPdu.String())
}
