// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdudecoder

import (
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/pdubuilder"
	types1 "github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
	"gotest.tools/assert"
	"testing"
)

func Test_DecodeRicServiceQueryPdu(t *testing.T) {
	rfAccepted1 := make(types1.RanFunctionRevisions)
	rfAccepted1[100] = 2
	rfAccepted1[200] = 2

	rsq, err := pdubuilder.CreateRicServiceQueryE2apPdu(54)
	assert.NilError(t, err)
	assert.Assert(t, rsq != nil)
	rsq.GetInitiatingMessage().GetValue().GetRicServiceQuery().SetRanFunctionsAccepted(rfAccepted1)

	transactionID, ranFunctionsAccepted, err := DecodeRicServiceQueryPdu(rsq)
	assert.NilError(t, err)
	//assert.Assert(t, ricIdentity != nil) //Commented due to the Linters (v1.34.1) error - possible nil pointer dereference (https://staticcheck.io/docs/checks#SA5011) on lines 23, 24 & 25

	assert.Equal(t, 2, len(ranFunctionsAccepted))
	rfa100, ok := ranFunctionsAccepted[100]
	assert.Assert(t, ok, "expected a key '100'")
	assert.Equal(t, 2, int(rfa100))
	rfa200, ok := ranFunctionsAccepted[200]
	assert.Assert(t, ok, "expected a key '200'")
	assert.Equal(t, 2, int(rfa200))

	if transactionID != nil {
		assert.Equal(t, int32(54), *transactionID)
	}
}
