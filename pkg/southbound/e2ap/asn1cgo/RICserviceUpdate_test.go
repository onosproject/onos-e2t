// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

import (
	"encoding/hex"
	"testing"

	e2ap_pdu_contents "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-pdu-contents"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/pdubuilder"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
	"gotest.tools/assert"
)

func createRicServiceUpdateMsg() (*e2ap_pdu_contents.RicserviceUpdate, error) {
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

	rsu, err := pdubuilder.CreateRicServiceUpdateE2apPdu(1)
	if err != nil {
		return nil, err
	}
	rsu.GetInitiatingMessage().GetProcedureCode().GetRicServiceUpdate().GetInitiatingMessage().
		SetRanFunctionsAdded(ranFunctionAddedList).SetRanFunctionsModified(ranFunctionModifiedList).SetRanFunctionsDeleted(rfDeleted)

	//if err := rsu.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating RicServiceUpdate %s", err.Error())
	//}
	return rsu.GetInitiatingMessage().GetProcedureCode().GetRicServiceUpdate().GetInitiatingMessage(), nil
}

func Test_xerEncodingRicServiceUpdate(t *testing.T) {

	rsu, err := createRicServiceUpdateMsg()
	assert.NilError(t, err, "Error creating RicServiceUpdate PDU")

	xer, err := xerEncodeRicServiceUpdate(rsu)
	assert.NilError(t, err)
	t.Logf("RicServiceUpdate XER\n%s", string(xer))

	result, err := xerDecodeRicServiceUpdate(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("RicServiceUpdate XER - decoded\n%v", result)
	assert.DeepEqual(t, rsu.GetProtocolIes().GetE2ApProtocolIes10().GetValue().GetValue()[0].GetE2ApProtocolIes10().GetValue().GetRanFunctionOid().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes10().GetValue().GetValue()[0].GetE2ApProtocolIes10().GetValue().GetRanFunctionOid().GetValue())
	assert.Equal(t, rsu.GetProtocolIes().GetE2ApProtocolIes10().GetValue().GetValue()[0].GetE2ApProtocolIes10().GetValue().GetRanFunctionRevision().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes10().GetValue().GetValue()[0].GetE2ApProtocolIes10().GetValue().GetRanFunctionRevision().GetValue())
	assert.Equal(t, rsu.GetProtocolIes().GetE2ApProtocolIes11().GetValue().GetValue()[0].GetRanFunctionIdItemIes6().GetValue().GetRanFunctionId().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes11().GetValue().GetValue()[0].GetRanFunctionIdItemIes6().GetValue().GetRanFunctionId().GetValue())
	assert.Equal(t, rsu.GetProtocolIes().GetE2ApProtocolIes12().GetValue().GetValue()[0].GetE2ApProtocolIes10().GetValue().GetRanFunctionId().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes12().GetValue().GetValue()[0].GetE2ApProtocolIes10().GetValue().GetRanFunctionId().GetValue())
	assert.DeepEqual(t, rsu.GetProtocolIes().GetE2ApProtocolIes12().GetValue().GetValue()[0].GetE2ApProtocolIes10().GetValue().GetRanFunctionOid().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes12().GetValue().GetValue()[0].GetE2ApProtocolIes10().GetValue().GetRanFunctionOid().GetValue())
	assert.Equal(t, rsu.GetProtocolIes().GetE2ApProtocolIes12().GetValue().GetValue()[0].GetE2ApProtocolIes10().GetValue().GetRanFunctionRevision().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes12().GetValue().GetValue()[0].GetE2ApProtocolIes10().GetValue().GetRanFunctionRevision().GetValue())
	assert.Equal(t, rsu.GetProtocolIes().GetE2ApProtocolIes49().GetValue().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes49().GetValue().GetValue())
}

func Test_perEncodingRicServiceUpdate(t *testing.T) {

	rsu, err := createRicServiceUpdateMsg()
	assert.NilError(t, err, "Error creating RicServiceUpdate PDU")

	per, err := perEncodeRicServiceUpdate(rsu)
	assert.NilError(t, err)
	t.Logf("RicServiceUpdate PER\n%v", hex.Dump(per))

	result, err := perDecodeRicServiceUpdate(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("RicServiceUpdate PER - decoded\n%v", result)
	assert.DeepEqual(t, rsu.GetProtocolIes().GetE2ApProtocolIes10().GetValue().GetValue()[0].GetE2ApProtocolIes10().GetValue().GetRanFunctionOid().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes10().GetValue().GetValue()[0].GetE2ApProtocolIes10().GetValue().GetRanFunctionOid().GetValue())
	assert.Equal(t, rsu.GetProtocolIes().GetE2ApProtocolIes10().GetValue().GetValue()[0].GetE2ApProtocolIes10().GetValue().GetRanFunctionRevision().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes10().GetValue().GetValue()[0].GetE2ApProtocolIes10().GetValue().GetRanFunctionRevision().GetValue())
	assert.Equal(t, rsu.GetProtocolIes().GetE2ApProtocolIes11().GetValue().GetValue()[0].GetRanFunctionIdItemIes6().GetValue().GetRanFunctionId().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes11().GetValue().GetValue()[0].GetRanFunctionIdItemIes6().GetValue().GetRanFunctionId().GetValue())
	assert.Equal(t, rsu.GetProtocolIes().GetE2ApProtocolIes12().GetValue().GetValue()[0].GetE2ApProtocolIes10().GetValue().GetRanFunctionId().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes12().GetValue().GetValue()[0].GetE2ApProtocolIes10().GetValue().GetRanFunctionId().GetValue())
	assert.DeepEqual(t, rsu.GetProtocolIes().GetE2ApProtocolIes12().GetValue().GetValue()[0].GetE2ApProtocolIes10().GetValue().GetRanFunctionOid().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes12().GetValue().GetValue()[0].GetE2ApProtocolIes10().GetValue().GetRanFunctionOid().GetValue())
	assert.Equal(t, rsu.GetProtocolIes().GetE2ApProtocolIes12().GetValue().GetValue()[0].GetE2ApProtocolIes10().GetValue().GetRanFunctionRevision().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes12().GetValue().GetValue()[0].GetE2ApProtocolIes10().GetValue().GetRanFunctionRevision().GetValue())
	assert.Equal(t, rsu.GetProtocolIes().GetE2ApProtocolIes49().GetValue().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes49().GetValue().GetValue())
}
