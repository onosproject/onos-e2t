// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

import (
	"encoding/hex"
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/pdubuilder"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/types"
	"gotest.tools/assert"
	"testing"
)

func createRanFunctionItem() (*e2appducontents.RanfunctionItem, error) {

	ranFunctionList := make(map[types.RanFunctionID]types.RanFunctionItem)
	ranFunctionList[200] = types.RanFunctionItem{
		Description: []byte("Type 2"),
		Revision:    2,
		OID:         []byte("oid2"),
	}

	gnbID, err := pdubuilder.CreateGnbIDchoice(1, 22)
	if err != nil {
		return nil, err
	}

	newE2apPdu, err := pdubuilder.CreateE2SetupRequestPdu([3]byte{0x4F, 0x4E, 0x46}, gnbID, ranFunctionList)
	if err != nil {
		return nil, err
	}
	res := newE2apPdu.GetInitiatingMessage().GetProcedureCode().GetE2Setup().GetInitiatingMessage().
		GetProtocolIes().GetE2ApProtocolIes10().GetValue().GetValue()[0].GetE2ApProtocolIes10().GetValue()
	//fmt.Printf("Returning following structure: \n %v \n", res)

	return res, nil
}

func Test_RanFunctionItem(t *testing.T) {

	rfi, err := createRanFunctionItem()
	assert.NilError(t, err)

	xer, err := xerEncodeRanFunctionItem(rfi)
	assert.NilError(t, err)
	t.Logf("RanFunctionList XER\n%s", xer)

	per, err := perEncodeRanFunctionItem(rfi)
	assert.NilError(t, err)
	t.Logf("RanFunctionList PER\n%v", hex.Dump(per))

	// Now reverse the XER
	rfiReversed, err := xerDecodeRanFunctionItem(xer)
	assert.NilError(t, err)
	assert.Assert(t, rfiReversed != nil)
	t.Logf("RanFunctionList decoded from XER is \n%v", rfiReversed)
	//assert.Equal(t, 2, len(rflReversed.GetValue()))

	// Now reverse the PER
	rfiReversedFromPer, err := perDecodeRanFunctionItem(per)
	assert.NilError(t, err)
	assert.Assert(t, rfiReversedFromPer != nil)
	t.Logf("RanFunctionList decoded from PER is \n%v", rfiReversedFromPer)
	//assert.Equal(t, 2, len(rflReversedFromPer.GetValue()))

}
