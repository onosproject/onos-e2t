// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

import (
	"encoding/hex"
	"fmt"
	e2ap_pdu_contents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/pdubuilder"
	"gotest.tools/assert"
	"testing"
)

func createRicServiceUpdateMsg() (*e2ap_pdu_contents.RicserviceUpdate, error) {

	rsu, err := pdubuilder.CreateRicServiceUpdateE2apPdu()
	if err != nil {
		return nil, err
	}

	if err := rsu.Validate(); err != nil {
		return nil, fmt.Errorf("error validating RicServiceUpdate %s", err.Error())
	}
	return rsu.GetInitiatingMessage().GetProcedureCode().GetRicServiceUpdate().GetInitiatingMessage(), nil
}

func Test_xerEncodingRicServiceUpdate(t *testing.T) {

	rsu, err := createRicServiceUpdateMsg()
	assert.NilError(t, err, "Error creating RicServiceUpdate PDU")

	xer, err := xerEncodeRicServiceUpdate(rsu)
	assert.NilError(t, err)
	//assert.Equal(t, 2646, len(xer))
	t.Logf("RicServiceUpdate XER\n%s", string(xer))

	//result, err := xerDecodeRicServiceUpdate(xer)
	//assert.NilError(t, err)
	//assert.Assert(t, result != nil)
	//t.Logf("RicServiceUpdate XER - decoded\n%v", result)
	//assert.Equal(t, rsu.GetProtocolIes().GetE2ApProtocolIes10().GetRanFunctionsAddedList().GetValue()[0].GetE2ApProtocolIes10().GetValue().GetRanFunctionOid().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes10().GetRanFunctionsAddedList().GetValue()[0].GetE2ApProtocolIes10().GetValue().GetRanFunctionOid().GetValue())
	//assert.Equal(t, rsu.GetProtocolIes().GetE2ApProtocolIes10().GetRanFunctionsAddedList().GetValue()[0].GetE2ApProtocolIes10().GetValue().GetRanFunctionRevision().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes10().GetRanFunctionsAddedList().GetValue()[0].GetE2ApProtocolIes10().GetValue().GetRanFunctionRevision().GetValue())
	//assert.Equal(t, rsu.GetProtocolIes().GetE2ApProtocolIes11().GetRanFunctionsDeletedList().GetValue()[0].GetRanFunctionIdItemIes6().GetValue().GetRanFunctionId().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes11().GetRanFunctionsDeletedList().GetValue()[0].GetRanFunctionIdItemIes6().GetValue().GetRanFunctionId().GetValue())
	//assert.Equal(t, rsu.GetProtocolIes().GetE2ApProtocolIes12().GetRanFunctionsModifiedList().GetValue()[0].GetE2ApProtocolIes10().GetValue().GetRanFunctionId().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes12().GetRanFunctionsModifiedList().GetValue()[0].GetE2ApProtocolIes10().GetValue().GetRanFunctionId().GetValue())
	//assert.Equal(t, rsu.GetProtocolIes().GetE2ApProtocolIes12().GetRanFunctionsModifiedList().GetValue()[0].GetE2ApProtocolIes10().GetValue().GetRanFunctionOid().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes12().GetRanFunctionsModifiedList().GetValue()[0].GetE2ApProtocolIes10().GetValue().GetRanFunctionOid().GetValue())
	//assert.Equal(t, rsu.GetProtocolIes().GetE2ApProtocolIes12().GetRanFunctionsModifiedList().GetValue()[0].GetE2ApProtocolIes10().GetValue().GetRanFunctionRevision().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes12().GetRanFunctionsModifiedList().GetValue()[0].GetE2ApProtocolIes10().GetValue().GetRanFunctionRevision().GetValue())
}

func Test_perEncodingRicServiceUpdate(t *testing.T) {

	rsu, err := createRicServiceUpdateMsg()
	assert.NilError(t, err, "Error creating RicServiceUpdate PDU")

	per, err := perEncodeRicServiceUpdate(rsu)
	assert.NilError(t, err)
	//assert.Equal(t, 72, len(per))
	t.Logf("RicServiceUpdate PER\n%v", hex.Dump(per))

	//result, err := perDecodeRicServiceUpdate(per)
	//assert.NilError(t, err)
	//assert.Assert(t, result != nil)
	//t.Logf("RicServiceUpdate PER - decoded\n%v", result)
	//assert.Equal(t, rsu.GetProtocolIes().GetE2ApProtocolIes10().GetRanFunctionsAddedList().GetValue()[0].GetE2ApProtocolIes10().GetValue().GetRanFunctionOid().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes10().GetRanFunctionsAddedList().GetValue()[0].GetE2ApProtocolIes10().GetValue().GetRanFunctionOid().GetValue())
	//assert.Equal(t, rsu.GetProtocolIes().GetE2ApProtocolIes10().GetRanFunctionsAddedList().GetValue()[0].GetE2ApProtocolIes10().GetValue().GetRanFunctionRevision().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes10().GetRanFunctionsAddedList().GetValue()[0].GetE2ApProtocolIes10().GetValue().GetRanFunctionRevision().GetValue())
	//assert.Equal(t, rsu.GetProtocolIes().GetE2ApProtocolIes11().GetRanFunctionsDeletedList().GetValue()[0].GetRanFunctionIdItemIes6().GetValue().GetRanFunctionId().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes11().GetRanFunctionsDeletedList().GetValue()[0].GetRanFunctionIdItemIes6().GetValue().GetRanFunctionId().GetValue())
	//assert.Equal(t, rsu.GetProtocolIes().GetE2ApProtocolIes12().GetRanFunctionsModifiedList().GetValue()[0].GetE2ApProtocolIes10().GetValue().GetRanFunctionId().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes12().GetRanFunctionsModifiedList().GetValue()[0].GetE2ApProtocolIes10().GetValue().GetRanFunctionId().GetValue())
	//assert.Equal(t, rsu.GetProtocolIes().GetE2ApProtocolIes12().GetRanFunctionsModifiedList().GetValue()[0].GetE2ApProtocolIes10().GetValue().GetRanFunctionOid().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes12().GetRanFunctionsModifiedList().GetValue()[0].GetE2ApProtocolIes10().GetValue().GetRanFunctionOid().GetValue())
	//assert.Equal(t, rsu.GetProtocolIes().GetE2ApProtocolIes12().GetRanFunctionsModifiedList().GetValue()[0].GetE2ApProtocolIes10().GetValue().GetRanFunctionRevision().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes12().GetRanFunctionsModifiedList().GetValue()[0].GetE2ApProtocolIes10().GetValue().GetRanFunctionRevision().GetValue())
}
