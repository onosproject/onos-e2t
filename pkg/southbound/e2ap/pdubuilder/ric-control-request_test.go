// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
package pdubuilder

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/encoder"
	"testing"

	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
	"gotest.tools/assert"
)

func TestRicControlRequest(t *testing.T) {
	ricRequestID := types.RicRequest{
		RequestorID: 21,
		InstanceID:  22,
	}
	var ranFuncID types.RanFunctionID = 9
	var ricCallPrID types.RicCallProcessID = []byte("123")
	var ricCtrlHdr types.RicControlHeader = []byte("456")
	var ricCtrlMsg types.RicControlMessage = []byte("789")
	request, err := NewControlRequest(ricRequestID,
		ranFuncID, ricCtrlHdr, ricCtrlMsg)
	assert.NilError(t, err)
	assert.Assert(t, request != nil)

	request.SetRicCallProcessID([]byte{0xCF, 0xFF}).SetRicControlAckRequest(e2ap_ies.RiccontrolAckRequest_RICCONTROL_ACK_REQUEST_NO_ACK).
		SetRicCallProcessID(ricCallPrID)
	newE2apPdu, err := CreateRicControlRequestE2apPdu(request)
	assert.NilError(t, err)
	assert.Assert(t, request != nil)

	perNew, err := encoder.PerEncodeE2ApPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("RicControlRequest E2AP PDU PER with Go APER library\n%v", hex.Dump(perNew))

	//Comparing reference PER bytes with Go APER library produced
	//assert.DeepEqual(t, per, perNew)

	result, err := encoder.PerDecodeE2ApPdu(perNew)
	assert.NilError(t, err)
	assert.DeepEqual(t, newE2apPdu.String(), result.String())
}
