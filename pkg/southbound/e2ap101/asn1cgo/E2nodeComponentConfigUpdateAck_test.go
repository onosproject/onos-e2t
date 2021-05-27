// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

import (
	"encoding/hex"
	"fmt"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	"gotest.tools/assert"
	"testing"
)

func createE2nodeComponentConfigUpdateAckMsg() (*e2ap_ies.E2NodeComponentConfigUpdateAck, error) {

	e2nodeComponentConfigUpdateAck := e2ap_ies.E2NodeComponentConfigUpdateAck{
		UpdateOutcome: 1,
		FailureCause: &e2ap_ies.Cause{
			Cause: &e2ap_ies.Cause_Protocol{
				Protocol: e2ap_ies.CauseProtocol_CAUSE_PROTOCOL_TRANSFER_SYNTAX_ERROR,
			},
		},
	}

	if err := e2nodeComponentConfigUpdateAck.Validate(); err != nil {
		return nil, fmt.Errorf("error validating E2nodeComponentConfigUpdateAck %s", err.Error())
	}
	return &e2nodeComponentConfigUpdateAck, nil
}

func Test_xerEncodingE2nodeComponentConfigUpdateAck(t *testing.T) {

	e2nodeComponentConfigUpdateAck, err := createE2nodeComponentConfigUpdateAckMsg()
	assert.NilError(t, err, "Error creating E2nodeComponentConfigUpdateAck PDU")

	xer, err := xerEncodeE2nodeComponentConfigUpdateAck(e2nodeComponentConfigUpdateAck)
	assert.NilError(t, err)
	assert.Equal(t, 206, len(xer))
	t.Logf("E2nodeComponentConfigUpdateAck XER\n%s", string(xer))

	result, err := xerDecodeE2nodeComponentConfigUpdateAck(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentConfigUpdateAck XER - decoded\n%v", result)
	assert.Equal(t, e2nodeComponentConfigUpdateAck.GetUpdateOutcome(), result.GetUpdateOutcome())
	assert.Equal(t, e2nodeComponentConfigUpdateAck.GetFailureCause().GetProtocol(), result.GetFailureCause().GetProtocol())
}

func Test_perEncodingE2nodeComponentConfigUpdateAck(t *testing.T) {

	e2nodeComponentConfigUpdateAck, err := createE2nodeComponentConfigUpdateAckMsg()
	assert.NilError(t, err, "Error creating E2nodeComponentConfigUpdateAck PDU")

	per, err := perEncodeE2nodeComponentConfigUpdateAck(e2nodeComponentConfigUpdateAck)
	assert.NilError(t, err)
	assert.Equal(t, 2, len(per))
	t.Logf("E2nodeComponentConfigUpdateAck PER\n%v", hex.Dump(per))

	result, err := perDecodeE2nodeComponentConfigUpdateAck(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentConfigUpdateAck PER - decoded\n%v", result)
	assert.Equal(t, e2nodeComponentConfigUpdateAck.GetUpdateOutcome(), result.GetUpdateOutcome())
	assert.Equal(t, e2nodeComponentConfigUpdateAck.GetFailureCause().GetProtocol(), result.GetFailureCause().GetProtocol())
}
