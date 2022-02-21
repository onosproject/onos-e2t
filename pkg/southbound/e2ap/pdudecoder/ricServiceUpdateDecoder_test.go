// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdudecoder

import (
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/pdubuilder"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
	"gotest.tools/assert"
	"testing"
)

func Test_DecodeRicServiceUpdatePdu(t *testing.T) {
	ranFunctionAddedList := make(map[types.RanFunctionID]types.RanFunctionItem)
	ranFunctionAddedList[100] = types.RanFunctionItem{
		Description: []byte("Type 1"),
		Revision:    1,
		OID:         "oid1",
	}

	ranFunctionAddedList[200] = types.RanFunctionItem{
		Description: []byte("Type 2"),
		Revision:    2,
		OID:         "oid2",
	}

	rfDeleted := make(types.RanFunctionRevisions)
	rfDeleted[100] = 2
	rfDeleted[200] = 2

	ranFunctionModifiedList := make(map[types.RanFunctionID]types.RanFunctionItem)
	ranFunctionModifiedList[100] = types.RanFunctionItem{
		Description: []byte("Type 3"),
		Revision:    3,
		OID:         "oid3",
	}

	ranFunctionModifiedList[200] = types.RanFunctionItem{
		Description: []byte("Type 4"),
		Revision:    4,
		OID:         "oid4",
	}

	e2apPdu, err := pdubuilder.CreateRicServiceUpdateE2apPdu(1)
	assert.NilError(t, err)
	assert.Assert(t, e2apPdu != nil)
	e2apPdu.GetInitiatingMessage().GetValue().GetRicServiceUpdate().
		SetRanFunctionsAdded(ranFunctionAddedList).SetRanFunctionsModified(ranFunctionModifiedList).SetRanFunctionsDeleted(rfDeleted)

	transactionID, rfal, rfdl, rfml, err := DecodeRicServiceUpdatePdu(e2apPdu)
	assert.NilError(t, err)

	assert.DeepEqual(t, []byte(rfal[100].Description), []byte("Type 1"))
	assert.DeepEqual(t, []byte(rfal[100].OID), []byte("oid1"))
	assert.Equal(t, int(rfal[100].Revision), 1)
	assert.DeepEqual(t, []byte(rfal[200].Description), []byte("Type 2"))
	assert.DeepEqual(t, []byte(rfal[200].OID), []byte("oid2"))
	assert.Equal(t, int(rfal[200].Revision), 2)
	rfd100, ok := rfdl[100]
	assert.Assert(t, ok, "expected a key '100'")
	assert.Equal(t, 2, int(rfd100))
	rfd200, ok := rfdl[200]
	assert.Assert(t, ok, "expected a key '200'")
	assert.Equal(t, 2, int(rfd200))
	assert.DeepEqual(t, []byte(rfml[100].Description), []byte("Type 3"))
	assert.DeepEqual(t, []byte(rfml[100].OID), []byte("oid3"))
	assert.Equal(t, int(rfml[100].Revision), 3)
	assert.DeepEqual(t, []byte(rfml[200].Description), []byte("Type 4"))
	assert.DeepEqual(t, []byte(rfml[200].OID), []byte("oid4"))
	assert.Equal(t, int(rfml[200].Revision), 4)

	if transactionID != nil {
		assert.Equal(t, int32(1), *transactionID)
	}
}
