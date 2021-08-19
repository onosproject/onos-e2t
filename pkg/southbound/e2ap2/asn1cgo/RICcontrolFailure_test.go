// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

import (
	"encoding/hex"
	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-ies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap2/pdubuilder"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap2/types"
	"gotest.tools/assert"
	"reflect"
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
	cause := &e2apies.Cause{
		Cause: &e2apies.Cause_RicRequest{
			RicRequest: e2apies.CauseRic_CAUSE_RIC_CONTROL_MESSAGE_INVALID,
		},
	}
	e2ApPduRcf, err := pdubuilder.CreateRicControlFailureE2apPdu(ricRequestID,
		ranFuncID, cause)
	assert.NilError(t, err)
	assert.Assert(t, e2ApPduRcf != nil)
	e2ApPduRcf.GetUnsuccessfulOutcome().GetProcedureCode().GetRicControl().GetUnsuccessfulOutcome().
		SetRicControlOutcome(ricCtrlOut).SetRicCallProcessID(ricCallPrID)
	t.Logf("Message we're going to encode is following: \n %v \n", e2ApPduRcf)

	xer, err := xerEncodeRICcontrolFailure(
		e2ApPduRcf.GetUnsuccessfulOutcome().GetProcedureCode().GetRicControl().GetUnsuccessfulOutcome())
	assert.NilError(t, err)
	t.Logf("XER RICcontrolFailure\n%s", xer)

	e2apPdu, err := xerDecodeRICcontrolFailure(xer)
	assert.NilError(t, err)
	t.Logf("RICcontrolFailureMessage decoded from XER is \n%v", e2apPdu)
	//assert.DeepEqual(t, e2ApPduRcf.GetUnsuccessfulOutcome().GetProcedureCode().GetRicControl().GetUnsuccessfulOutcome(), e2apPdu)
	out := reflect.DeepEqual(e2ApPduRcf.GetUnsuccessfulOutcome().GetProcedureCode().GetRicControl().GetUnsuccessfulOutcome(), e2apPdu)
	assert.Assert(t, out != false)

	per, err := perEncodeRICcontrolFailure(
		e2ApPduRcf.GetUnsuccessfulOutcome().GetProcedureCode().GetRicControl().GetUnsuccessfulOutcome())
	assert.NilError(t, err)
	t.Logf("PER RICcontrolFailure\n%v", hex.Dump(per))

	e2apPdu, err = perDecodeRICcontrolFailure(per)
	assert.NilError(t, err)
	t.Logf("RICcontrolFailureMessage is \n%v", e2apPdu)
	//assert.DeepEqual(t, e2ApPduRcf.GetUnsuccessfulOutcome().GetProcedureCode().GetRicControl().GetUnsuccessfulOutcome(), e2apPdu)
	out1 := reflect.DeepEqual(e2ApPduRcf.GetUnsuccessfulOutcome().GetProcedureCode().GetRicControl().GetUnsuccessfulOutcome(), e2apPdu)
	assert.Assert(t, out1 != false)
}
