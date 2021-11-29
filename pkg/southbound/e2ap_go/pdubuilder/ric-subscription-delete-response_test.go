// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0
package pdubuilder

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap_go/encoder"
	"testing"

	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap_go/types"
	"gotest.tools/assert"
)

func TestRicSubscriptionDeleteResponse(t *testing.T) {
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
	//assert.DeepEqual(t, per, perNew)

	e2apPdu, err := encoder.PerDecodeE2ApPdu(perNew)
	assert.NilError(t, err)
	assert.DeepEqual(t, newE2apPdu.String(), e2apPdu.String())

	//per, err := asn1cgo.PerEncodeE2apPdu(newE2apPdu)
	//assert.NilError(t, err)
	//t.Logf("RicSubscriptionDeleteResponse E2AP PDU PER\n%v", hex.Dump(per))
	//
	//e2apPdu, err = asn1cgo.PerDecodeE2apPdu(per)
	//assert.NilError(t, err)
	//assert.DeepEqual(t, newE2apPdu.String(), e2apPdu.String())
}
