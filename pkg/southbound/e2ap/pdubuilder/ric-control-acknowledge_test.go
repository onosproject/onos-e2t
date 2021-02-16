// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
package pdubuilder

import (
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2apies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
	"gotest.tools/assert"
	"testing"
)

func TestRicControlAcknowledge(t *testing.T) {
	ricRequestID := types.RicRequest{
		RequestorID: 21,
		InstanceID:  22,
	}
	var ranFuncID types.RanFunctionID = 9
	var ricCallPrID types.RicCallProcessID = []byte("123")
	ricControlStatus := e2apies.RiccontrolStatus_RICCONTROL_STATUS_SUCCESS
	var ricCtrlOut types.RicControlOutcome = []byte("456")
	newE2apPdu, err := CreateRicControlAcknowledgeE2apPdu(ricRequestID,
		ranFuncID, ricCallPrID, ricControlStatus, ricCtrlOut)
	assert.NilError(t, err)
	assert.Assert(t, newE2apPdu != nil)

	//xer, err := asn1cgo.XerEncodeE2apPdu(newE2apPdu)
	//assert.NilError(t, err)
	//t.Logf("RIC Control Request XER\n%s", string(xer))
}
