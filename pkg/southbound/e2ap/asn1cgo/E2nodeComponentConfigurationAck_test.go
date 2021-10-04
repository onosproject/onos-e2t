// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package asn1cgo

import (
	"encoding/hex"
	"testing"

	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	"gotest.tools/assert"
)

func createE2nodeComponentConfigurationAckMsg() (*e2ap_ies.E2NodeComponentConfigurationAck, error) {

	e2ncca := e2ap_ies.E2NodeComponentConfigurationAck{
		UpdateOutcome: e2ap_ies.UpdateOutcome_UPDATE_OUTCOME_SUCCESS,
		FailureCause: &e2ap_ies.Cause{
			Cause: &e2ap_ies.Cause_E2Node{
				E2Node: e2ap_ies.CauseE2Node_CAUSE_E2NODE_E2NODE_COMPONENT_UNKNOWN,
			},
		},
	}

	//if err := e2nodeComponentConfigUpdateGnb.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating E2nodeComponentConfigUpdateGnb %s", err.Error())
	//}
	return &e2ncca, nil
}

func Test_xerEncodingE2nodeComponentConfigurationAck(t *testing.T) {

	e2ncca, err := createE2nodeComponentConfigurationAckMsg()
	assert.NilError(t, err, "Error creating E2nodeComponentConfigurationAck PDU")

	xer, err := xerEncodeE2nodeComponentConfigurationAck(e2ncca)
	assert.NilError(t, err)
	t.Logf("E2nodeComponentConfigurationAck XER\n%s", string(xer))

	result, err := xerDecodeE2nodeComponentConfigurationAck(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentConfigurationAck XER - decoded\n%v", result)
	assert.Equal(t, e2ncca.GetUpdateOutcome().Number(), result.GetUpdateOutcome().Number())
	assert.Equal(t, e2ncca.GetFailureCause().GetE2Node().Number(), result.GetFailureCause().GetE2Node().Number())
}

func Test_perEncodingE2nodeComponentConfigurationAck(t *testing.T) {

	e2ncca, err := createE2nodeComponentConfigurationAckMsg()
	assert.NilError(t, err, "Error creating E2nodeComponentConfigurationAck PDU")

	per, err := perEncodeE2nodeComponentConfigurationAck(e2ncca)
	assert.NilError(t, err)
	t.Logf("E2nodeComponentConfigurationAck PER\n%v", hex.Dump(per))

	result, err := perDecodeE2nodeComponentConfigurationAck(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentConfigurationAck PER - decoded\n%v", result)
	assert.Equal(t, e2ncca.GetUpdateOutcome().Number(), result.GetUpdateOutcome().Number())
	assert.Equal(t, e2ncca.GetFailureCause().GetE2Node().Number(), result.GetFailureCause().GetE2Node().Number())
}
