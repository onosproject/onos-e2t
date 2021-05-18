// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

import (
	"encoding/hex"
	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/pdubuilder"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/types"
	"gotest.tools/assert"
	"testing"
)

func Test_RICcontrolFailure(t *testing.T) {
	ricRequestID := types.RicRequest{
		RequestorID: 21,
		InstanceID:  22,
	}
	var ranFuncID types.RanFunctionID = 9
	var ricCallPrID types.RicCallProcessID = []byte("123")
	var ricCtrlOut types.RicControlOutcome = []byte("456")
	cause := e2apies.Cause{
		Cause: &e2apies.Cause_RicRequest{
			RicRequest: e2apies.CauseRic_CAUSE_RIC_CONTROL_MESSAGE_INVALID,
		},
	}
	e2ApPduRcf, err := pdubuilder.CreateRicControlFailureE2apPdu(ricRequestID,
		ranFuncID, ricCallPrID, cause, ricCtrlOut)
	assert.NilError(t, err)
	assert.Assert(t, e2ApPduRcf != nil)
	//fmt.Printf("Message we're going to encode is following: \n %v \n", e2ApPduRcf)

	xer, err := xerEncodeRICcontrolFailure(
		e2ApPduRcf.GetUnsuccessfulOutcome().GetProcedureCode().GetRicControl().GetUnsuccessfulOutcome())
	assert.NilError(t, err)
	t.Logf("XER RICcontrolFailure\n%s", xer)

	e2apPdu, err := xerDecodeRICcontrolFailure(xer)
	assert.NilError(t, err)
	assert.DeepEqual(t, e2ApPduRcf.GetUnsuccessfulOutcome().GetProcedureCode().GetRicControl().GetUnsuccessfulOutcome(), e2apPdu)

	per, err := perEncodeRICcontrolFailure(
		e2ApPduRcf.GetUnsuccessfulOutcome().GetProcedureCode().GetRicControl().GetUnsuccessfulOutcome())
	assert.NilError(t, err)
	t.Logf("PER RICcontrolFailure\n%v", hex.Dump(per))

	e2apPdu, err = perDecodeRICcontrolFailure(per)
	assert.NilError(t, err)
	assert.DeepEqual(t, e2ApPduRcf.GetUnsuccessfulOutcome().GetProcedureCode().GetRicControl().GetUnsuccessfulOutcome(), e2apPdu)
}
