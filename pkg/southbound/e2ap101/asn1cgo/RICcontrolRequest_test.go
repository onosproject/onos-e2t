// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

import (
	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/pdubuilder"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/types"
	"gotest.tools/assert"
	"testing"
)

func Test_RICcontrolRequest(t *testing.T) {
	ricRequestID := types.RicRequest{
		RequestorID: 21,
		InstanceID:  22,
	}
	var ranFuncID types.RanFunctionID = 9
	var ricCallPrID types.RicCallProcessID = []byte("123")
	var ricCtrlHdr types.RicControlHeader = []byte("456")
	var ricCtrlMsg types.RicControlMessage = []byte("789")
	ricCtrlAckRequest := e2apies.RiccontrolAckRequest_RICCONTROL_ACK_REQUEST_N_ACK
	e2ApPduRcr, err := pdubuilder.CreateRicControlRequestE2apPdu(ricRequestID,
		ranFuncID, ricCallPrID, ricCtrlHdr, ricCtrlMsg, ricCtrlAckRequest)
	assert.NilError(t, err)
	assert.Assert(t, e2ApPduRcr != nil)

	xer, err := xerEncodeRICcontrolRequest(
		e2ApPduRcr.GetInitiatingMessage().GetProcedureCode().GetRicControl().GetInitiatingMessage())
	assert.NilError(t, err)
	t.Logf("XER RICcontrolRequest\n%s", xer)

	e2apPdu, err := xerDecodeRICcontrolRequest(xer)
	assert.NilError(t, err)
	assert.DeepEqual(t, e2ApPduRcr.GetInitiatingMessage().GetProcedureCode().GetRicControl().GetInitiatingMessage(), e2apPdu)

	per, err := perEncodeRICcontrolRequest(
		e2ApPduRcr.GetInitiatingMessage().GetProcedureCode().GetRicControl().GetInitiatingMessage())
	assert.NilError(t, err)
	t.Logf("PER RICcontrolRequest\n%s", per)

	e2apPdu, err = perDecodeRICcontrolRequest(per)
	assert.NilError(t, err)
	assert.DeepEqual(t, e2ApPduRcr.GetInitiatingMessage().GetProcedureCode().GetRicControl().GetInitiatingMessage(), e2apPdu)
}
