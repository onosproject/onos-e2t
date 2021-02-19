// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

import (
	e2appdudescriptions "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-descriptions"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/pdubuilder"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/types"
	"gotest.tools/assert"
	"testing"
)

func Test_RanFunctionsList(t *testing.T) {
	ranFunctionList := make(types.RanFunctions)
	ranFunctionList[100] = types.RanFunctionItem{
		Description: []byte("Type 1"),
		Revision:    1,
	}

	ranFunctionList[200] = types.RanFunctionItem{
		Description: []byte("Type 2"),
		Revision:    1,
	}

	e2apSetupRequest, err := pdubuilder.CreateE2SetupRequestPdu([3]byte{0x4F, 0x4E, 0x46}, ranFunctionList)
	assert.NilError(t, err)

	im := e2apSetupRequest.GetE2ApPdu().(*e2appdudescriptions.E2ApPdu_InitiatingMessage)
	rflist := im.InitiatingMessage.GetProcedureCode().GetE2Setup().GetInitiatingMessage().GetProtocolIes().GetE2ApProtocolIes10().GetValue()
	xer, err := xerEncodeRanFunctionsList(rflist)
	assert.NilError(t, err)
	t.Logf("RanFunctionList XER\n%s", xer)

	per, err := perEncodeRanFunctionsList(rflist)
	assert.NilError(t, err)
	t.Logf("RanFunctionList PER\n%s", per)

	// Now reverse the XER
	rflReversed, err := xerDecodeRanFunctionList(xer)
	assert.NilError(t, err)
	assert.Assert(t, rflReversed != nil)

	assert.Equal(t, 2, len(rflReversed.GetValue()))

	// Now reverse the PER
	rflReversedFromPer, err := perDecodeRanFunctionList(per)
	assert.NilError(t, err)
	assert.Assert(t, rflReversedFromPer != nil)
	assert.Equal(t, 2, len(rflReversedFromPer.GetValue()))

}
