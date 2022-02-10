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
	//ricRequestID1 := types1.RicRequest{
	//	RequestorID: 21,
	//	InstanceID:  22,
	//}
	//var ranFuncID1 types1.RanFunctionID = 9
	////var ricCallPrID types.RicCallProcessID = []byte("123")
	//var ricCtrlHdr1 types1.RicControlHeader = []byte("456")
	//var ricCtrlMsg1 types1.RicControlMessage = []byte("789")
	////ricCtrlAckRequest := e2apies.RiccontrolAckRequest_RICCONTROL_ACK_REQUEST_N_ACK
	//request1, err := pdubuilder.NewControlRequest(ricRequestID1,
	//	ranFuncID1, ricCtrlHdr1, ricCtrlMsg1)
	//assert.NilError(t, err)
	//assert.Assert(t, request1 != nil)
	//
	//request1.SetRicCallProcessID([]byte{0xCF, 0xFF}).SetRicControlAckRequest(e2apies.RiccontrolAckRequest_RICCONTROL_ACK_REQUEST_NO_ACK)
	//e2apPdu, err := pdubuilder.CreateRicControlRequestE2apPdu(request1)
	//assert.NilError(t, err)
	//assert.Assert(t, e2apPdu != nil)
	//
	//per, err := asn1cgo.PerEncodeE2apPdu(e2apPdu)
	//assert.NilError(t, err)
	//t.Logf("RicControlRequest E2AP PDU PER\n%v", hex.Dump(per))

	ricRequestID := types.RicRequest{
		RequestorID: 21,
		InstanceID:  22,
	}
	var ranFuncID types.RanFunctionID = 9
	//var ricCallPrID types.RicCallProcessID = []byte("123")
	var ricCtrlHdr types.RicControlHeader = []byte("456")
	var ricCtrlMsg types.RicControlMessage = []byte("789")
	//ricCtrlAckRequest := e2apies.RiccontrolAckRequest_RICCONTROL_ACK_REQUEST_N_ACK
	request, err := NewControlRequest(ricRequestID,
		ranFuncID, ricCtrlHdr, ricCtrlMsg)
	assert.NilError(t, err)
	assert.Assert(t, request != nil)

	request.SetRicCallProcessID([]byte{0xCF, 0xFF}).SetRicControlAckRequest(e2ap_ies.RiccontrolAckRequest_RICCONTROL_ACK_REQUEST_NO_ACK)
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

	// Decoding the message from the APER bytes produced by CGo
	//result11, err := encoder.PerDecodeE2ApPdu(per)
	//assert.NilError(t, err)
	//assert.DeepEqual(t, newE2apPdu.String(), result11.String())
	//
	//result1, err := asn1cgo.PerDecodeE2apPdu(perNew)
	//assert.NilError(t, err)
	//assert.DeepEqual(t, result1.String(), e2apPdu.String())
}
