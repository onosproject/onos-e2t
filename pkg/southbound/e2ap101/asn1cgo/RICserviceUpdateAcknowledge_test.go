// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

import (
	"encoding/hex"
	"fmt"
	e2ap_pdu_contents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/pdubuilder"
	"gotest.tools/assert"
	"testing"
)

func createRicServiceUpdateAcknowledgeMsg() (*e2ap_pdu_contents.RicserviceUpdateAcknowledge, error) {

	ricserviceUpdateAcknowledge, err := pdubuilder.CreateRicServiceUpdateAcknowledgeE2apPdu()
	if err != nil {
		return nil, err
	}

	if err := ricserviceUpdateAcknowledge.Validate(); err != nil {
		return nil, fmt.Errorf("error validating RicserviceUpdateAcknowledge %s", err.Error())
	}
	return ricserviceUpdateAcknowledge.GetSuccessfulOutcome().GetProcedureCode().GetRicServiceUpdate().GetSuccessfulOutcome(), nil
}

func Test_xerEncodingRicserviceUpdateAcknowledge(t *testing.T) {

	ricserviceUpdateAcknowledge, err := createRicServiceUpdateAcknowledgeMsg()
	assert.NilError(t, err, "Error creating RicServiceUpdateAcknowledge PDU")

	xer, err := xerEncodeRicServiceUpdateAcknowledge(ricserviceUpdateAcknowledge)
	assert.NilError(t, err)
	assert.Equal(t, 1762, len(xer))
	t.Logf("RicServiceUpdateAcknowledge XER\n%s", string(xer))

	result, err := xerDecodeRicServiceUpdateAcknowledge(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("RicServiceUpdateAcknowledge XER - decoded\n%v", result)
	assert.Equal(t, ricserviceUpdateAcknowledge.GetProtocolIes().GetE2ApProtocolIes9().GetValue().GetValue()[0].GetRanFunctionIdItemIes6().GetValue().GetRanFunctionRevision().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes9().GetValue().GetValue()[0].GetRanFunctionIdItemIes6().GetValue().GetRanFunctionRevision().GetValue())
	assert.Equal(t, ricserviceUpdateAcknowledge.GetProtocolIes().GetE2ApProtocolIes9().GetValue().GetValue()[0].GetRanFunctionIdItemIes6().GetValue().GetRanFunctionId().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes9().GetValue().GetValue()[0].GetRanFunctionIdItemIes6().GetValue().GetRanFunctionId().GetValue())
	assert.Equal(t, ricserviceUpdateAcknowledge.GetProtocolIes().GetE2ApProtocolIes13().GetValue().GetValue()[0].GetRanFunctionIdcauseItemIes7().GetValue().GetRanFunctionId().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes13().GetValue().GetValue()[0].GetRanFunctionIdcauseItemIes7().GetValue().GetRanFunctionId().GetValue())
	assert.Equal(t, ricserviceUpdateAcknowledge.GetProtocolIes().GetE2ApProtocolIes13().GetValue().GetValue()[0].GetRanFunctionIdcauseItemIes7().GetValue().GetCause().GetRicService(), result.GetProtocolIes().GetE2ApProtocolIes13().GetValue().GetValue()[0].GetRanFunctionIdcauseItemIes7().GetValue().GetCause().GetRicService())
}

func Test_perEncodingRicServiceUpdateAcknowledge(t *testing.T) {

	ricserviceUpdateAcknowledge, err := createRicServiceUpdateAcknowledgeMsg()
	assert.NilError(t, err, "Error creating RicServiceUpdateAcknowledge PDU")

	per, err := perEncodeRicServiceUpdateAcknowledge(ricserviceUpdateAcknowledge)
	assert.NilError(t, err)
	assert.Equal(t, 32, len(per))
	t.Logf("RicServiceUpdateAcknowledge PER\n%v", hex.Dump(per))

	result, err := perDecodeRicServiceUpdateAcknowledge(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("RicServiceUpdateAcknowledge PER - decoded\n%v", result)
	assert.Equal(t, ricserviceUpdateAcknowledge.GetProtocolIes().GetE2ApProtocolIes9().GetValue().GetValue()[0].GetRanFunctionIdItemIes6().GetValue().GetRanFunctionRevision().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes9().GetValue().GetValue()[0].GetRanFunctionIdItemIes6().GetValue().GetRanFunctionRevision().GetValue())
	assert.Equal(t, ricserviceUpdateAcknowledge.GetProtocolIes().GetE2ApProtocolIes9().GetValue().GetValue()[0].GetRanFunctionIdItemIes6().GetValue().GetRanFunctionId().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes9().GetValue().GetValue()[0].GetRanFunctionIdItemIes6().GetValue().GetRanFunctionId().GetValue())
	assert.Equal(t, ricserviceUpdateAcknowledge.GetProtocolIes().GetE2ApProtocolIes13().GetValue().GetValue()[0].GetRanFunctionIdcauseItemIes7().GetValue().GetRanFunctionId().GetValue(), result.GetProtocolIes().GetE2ApProtocolIes13().GetValue().GetValue()[0].GetRanFunctionIdcauseItemIes7().GetValue().GetRanFunctionId().GetValue())
	assert.Equal(t, ricserviceUpdateAcknowledge.GetProtocolIes().GetE2ApProtocolIes13().GetValue().GetValue()[0].GetRanFunctionIdcauseItemIes7().GetValue().GetCause().GetRicService(), result.GetProtocolIes().GetE2ApProtocolIes13().GetValue().GetValue()[0].GetRanFunctionIdcauseItemIes7().GetValue().GetCause().GetRicService())
}
