// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

import (
	"encoding/hex"
	"testing"

	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-ies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/pdubuilder"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
	"gotest.tools/assert"
)

func Test_RICcontrolAcknowledge(t *testing.T) {
	ricRequestID := types.RicRequest{
		RequestorID: 21,
		InstanceID:  22,
	}
	var ranFuncID types.RanFunctionID = 9
	var ricCallPrID types.RicCallProcessID = []byte("123")
	ricControlStatus := e2apies.RiccontrolStatus_RICCONTROL_STATUS_SUCCESS
	var ricCtrlOut types.RicControlOutcome = []byte("456")
	e2ApPduRca, err := pdubuilder.CreateRicControlAcknowledgeE2apPdu(ricRequestID,
		ranFuncID, ricControlStatus)
	assert.NilError(t, err)
	assert.Assert(t, e2ApPduRca != nil)
	e2ApPduRca.GetSuccessfulOutcome().GetProcedureCode().GetRicControl().GetSuccessfulOutcome().
		SetRicControlOutcome(ricCtrlOut).SetRicCallProcessID(ricCallPrID)
	//t.Logf("That's what we're going to encode: \n %v \n", e2ApPduRca.GetSuccessfulOutcome().GetProcedureCode().GetRicControl().GetSuccessfulOutcome())

	xer, err := xerEncodeRICcontrolAcknowledge(
		e2ApPduRca.GetSuccessfulOutcome().GetProcedureCode().GetRicControl().GetSuccessfulOutcome())
	assert.NilError(t, err)
	t.Logf("XER RICcontrolAcknowledge\n%s", xer)

	e2apPdu, err := xerDecodeRICcontrolAcknowledge(xer)
	assert.NilError(t, err)
	assert.DeepEqual(t, e2ApPduRca.GetSuccessfulOutcome().GetProcedureCode().GetRicControl().GetSuccessfulOutcome().String(), e2apPdu.String())

	per, err := perEncodeRICcontrolAcknowledge(
		e2ApPduRca.GetSuccessfulOutcome().GetProcedureCode().GetRicControl().GetSuccessfulOutcome())
	assert.NilError(t, err)
	t.Logf("PER RICcontrolAcknowledge\n%v", hex.Dump(per))

	e2apPdu, err = perDecodeRICcontrolAcknowledge(per)
	assert.NilError(t, err)
	assert.DeepEqual(t, e2ApPduRca.GetSuccessfulOutcome().GetProcedureCode().GetRicControl().GetSuccessfulOutcome().String(), e2apPdu.String())
}
