// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package asn1cgo

import (
	"encoding/hex"
	"testing"

	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-ies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/pdubuilder"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
	"gotest.tools/assert"
)

func Test_RICcontrolRequest(t *testing.T) {
	ricRequestID := types.RicRequest{
		RequestorID: 21,
		InstanceID:  22,
	}
	var ranFuncID types.RanFunctionID = 9
	//var ricCallPrID types.RicCallProcessID = []byte("123")
	var ricCtrlHdr types.RicControlHeader = []byte("456")
	var ricCtrlMsg types.RicControlMessage = []byte("789")
	//ricCtrlAckRequest := e2apies.RiccontrolAckRequest_RICCONTROL_ACK_REQUEST_N_ACK
	e2ApPduRcr, err := pdubuilder.NewControlRequest(ricRequestID,
		ranFuncID, ricCtrlHdr, ricCtrlMsg)
	assert.NilError(t, err)
	assert.Assert(t, e2ApPduRcr != nil)
	e2ApPduRcr.SetRicCallProcessID([]byte{0xCF, 0xFF}).SetRicControlAckRequest(e2ap_ies.RiccontrolAckRequest_RICCONTROL_ACK_REQUEST_N_ACK)

	t.Logf("That's what we're going to encode: \n %v \n", e2ApPduRcr)

	xer, err := xerEncodeRICcontrolRequest(e2ApPduRcr)
	assert.NilError(t, err)
	t.Logf("XER RICcontrolRequest\n%s", xer)

	e2apPdu, err := xerDecodeRICcontrolRequest(xer)
	assert.NilError(t, err)
	//assert.DeepEqual(t, e2ApPduRcr, e2apPdu)
	t.Logf("XER RICcontrolRequest - decoded\n%v", e2apPdu)
	assert.Equal(t, e2ApPduRcr.String(), e2apPdu.String())

	per, err := perEncodeRICcontrolRequest(e2ApPduRcr)
	assert.NilError(t, err)
	t.Logf("PER RICcontrolRequest\n%v", hex.Dump(per))

	e2apPdu, err = perDecodeRICcontrolRequest(per)
	assert.NilError(t, err)
	//assert.DeepEqual(t, e2ApPduRcr, e2apPdu)
	t.Logf("PER RICcontrolRequest - decoded\n%v", e2apPdu)
	assert.Equal(t, e2ApPduRcr.String(), e2apPdu.String())
}
