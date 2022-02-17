// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
package pdubuilder

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/encoder"
	types1 "github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
	"gotest.tools/assert"
	"testing"
)

func TestRicServiceQuery(t *testing.T) {
	rfAccepted1 := make(types1.RanFunctionRevisions)
	rfAccepted1[100] = 2

	rsq, err := CreateRicServiceQueryE2apPdu(54)
	assert.NilError(t, err)
	assert.Assert(t, rsq != nil)
	rsq.GetInitiatingMessage().GetValue().GetRicServiceQuery().SetRanFunctionsAccepted(rfAccepted1)

	perNew, err := encoder.PerEncodeE2ApPdu(rsq)
	assert.NilError(t, err)
	t.Logf("RicServiceQuery E2AP PDU PER with Go APER library\n%v", hex.Dump(perNew))

	//Comparing reference PER bytes with Go APER library produced
	//assert.DeepEqual(t, per, perNew)

	e2apPdu, err := encoder.PerDecodeE2ApPdu(perNew)
	assert.NilError(t, err)
	assert.DeepEqual(t, rsq.String(), e2apPdu.String())
}
