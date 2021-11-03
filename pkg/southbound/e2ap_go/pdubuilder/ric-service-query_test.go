// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0
package pdubuilder

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/asn1cgo"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
	"gotest.tools/assert"
	"testing"
)

func TestRicServiceQuery(t *testing.T) {
	rfAccepted := make(types.RanFunctionRevisions)
	rfAccepted[100] = 2
	rfAccepted[200] = 2

	newE2apPdu, err := CreateRicServiceQueryE2apPdu(54)
	assert.NilError(t, err)
	assert.Assert(t, newE2apPdu != nil)
	newE2apPdu.GetInitiatingMessage().GetProcedureCode().GetRicServiceQuery().GetInitiatingMessage().
		SetRanFunctionsAccepted(rfAccepted)

	//xer, err := asn1cgo.XerEncodeE2apPdu(newE2apPdu)
	//assert.NilError(t, err)
	//t.Logf("RicServiceQuery E2AP PDU XER\n%s", string(xer))
	//
	//e2apPdu, err := asn1cgo.XerDecodeE2apPdu(xer)
	//assert.NilError(t, err)
	//assert.DeepEqual(t, newE2apPdu.String(), e2apPdu.String())

	per, err := asn1cgo.PerEncodeE2apPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("RicServiceQuery E2AP PDU PER\n%v", hex.Dump(per))

	//perNew, err := encoder.PerEncodeE2ApPdu(newE2apPdu)
	//assert.NilError(t, err)
	//t.Logf("RicServiceQuery E2AP PDU PER with Go APER library\n%v", hex.Dump(perNew))
	//
	////Comparing reference PER bytes with Go APER library produced
	//assert.DeepEqual(t, per, perNew)

	//Comparing reference PER bytes with Go APER library produced
	//assert.DeepEqual(t, per, perNew)

	e2apPdu, err := asn1cgo.PerDecodeE2apPdu(per)
	assert.NilError(t, err)
	assert.DeepEqual(t, newE2apPdu.String(), e2apPdu.String())
}
