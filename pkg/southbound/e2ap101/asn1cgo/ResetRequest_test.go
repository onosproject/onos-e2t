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

func createResetRequestMsg() (*e2ap_pdu_contents.ResetRequest, error) {

	resetRequest, err := pdubuilder.CreateResetRequestE2apPdu()
	if err != nil {
		return nil, err
	}

	if err := resetRequest.Validate(); err != nil {
		return nil, fmt.Errorf("error validating ResetRequest %s", err.Error())
	}
	return resetRequest.GetInitiatingMessage().GetProcedureCode().GetReset_().GetInitiatingMessage(), nil
}

func Test_xerEncodingResetRequest(t *testing.T) {

	resetRequest, err := createResetRequestMsg()
	assert.NilError(t, err, "Error creating ResetRequest PDU")

	xer, err := xerEncodeResetRequest(resetRequest)
	assert.NilError(t, err)
	assert.Equal(t, 349, len(xer)) //ToDo - adjust length of the XER encoded message
	t.Logf("ResetRequest XER\n%s", string(xer))

	result, err := xerDecodeResetRequest(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("ResetRequest XER - decoded\n%v", result)
	//ToDo - adjust field's verification
	assert.Equal(t, resetRequest.GetProtocolIes(), result.GetProtocolIes())

}

func Test_perEncodingResetRequest(t *testing.T) {

	resetRequest, err := createResetRequestMsg()
	assert.NilError(t, err, "Error creating ResetRequest PDU")

	per, err := perEncodeResetRequest(resetRequest)
	assert.NilError(t, err)
	assert.Equal(t, 8, len(per)) // ToDo - adjust length of the PER encoded message
	t.Logf("ResetRequest PER\n%v", hex.Dump(per))

	result, err := perDecodeResetRequest(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("ResetRequest PER - decoded\n%v", result)
	//ToDo - adjust field's verification
	assert.Equal(t, resetRequest.GetProtocolIes(), result.GetProtocolIes())

}
