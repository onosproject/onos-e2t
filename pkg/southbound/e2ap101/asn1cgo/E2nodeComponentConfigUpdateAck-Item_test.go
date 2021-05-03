// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

import (
	"encoding/hex"
	"fmt"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	e2ap_pdu_contents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
	"gotest.tools/assert"
	"testing"
)

func createE2nodeComponentConfigUpdateAckItemMsg() (*e2ap_pdu_contents.E2NodeComponentConfigUpdateAckItem, error) {

	e2nodeComponentConfigUpdateAckItem := e2ap_pdu_contents.E2NodeComponentConfigUpdateAckItem{
		E2NodeComponentType: e2ap_ies.E2NodeComponentType_E2NODE_COMPONENT_TYPE_G_NB,
		E2NodeComponentId: &e2ap_ies.E2NodeComponentId{
			E2NodeComponentId: &e2ap_ies.E2NodeComponentId_E2NodeComponentTypeGnbCuUp{
				E2NodeComponentTypeGnbCuUp: &e2ap_ies.E2NodeComponentGnbCuUpId{
					GNbCuUpId: &e2ap_ies.GnbCuUpId{
						Value: 21,
					},
				},
			},
		},
		E2NodeComponentConfigUpdateAck: &e2ap_ies.E2NodeComponentConfigUpdateAck{
			UpdateOutcome: 1,
			FailureCause: &e2ap_ies.Cause{
				Cause: &e2ap_ies.Cause_Protocol{
					Protocol: e2ap_ies.CauseProtocol_CAUSE_PROTOCOL_TRANSFER_SYNTAX_ERROR,
				},
			},
		},
	}

	if err := e2nodeComponentConfigUpdateAckItem.Validate(); err != nil {
		return nil, fmt.Errorf("error validating E2nodeComponentConfigUpdateAckItem %s", err.Error())
	}
	return &e2nodeComponentConfigUpdateAckItem, nil
}

func Test_xerEncodingE2nodeComponentConfigUpdateAckItem(t *testing.T) {

	e2nodeComponentConfigUpdateAckItem, err := createE2nodeComponentConfigUpdateAckItemMsg()
	assert.NilError(t, err, "Error creating E2nodeComponentConfigUpdateAckItem PDU")

	xer, err := xerEncodeE2nodeComponentConfigUpdateAckItem(e2nodeComponentConfigUpdateAckItem)
	assert.NilError(t, err)
	assert.Equal(t, 533, len(xer))
	t.Logf("E2nodeComponentConfigUpdateAckItem XER\n%s", string(xer))

	result, err := xerDecodeE2nodeComponentConfigUpdateAckItem(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentConfigUpdateAckItem XER - decoded\n%v", result)
	assert.Equal(t, int32(e2nodeComponentConfigUpdateAckItem.GetE2NodeComponentType()), int32(result.GetE2NodeComponentType()))
	assert.Equal(t, e2nodeComponentConfigUpdateAckItem.GetE2NodeComponentId().GetE2NodeComponentTypeGnbCuUp().GetGNbCuUpId().GetValue(), result.GetE2NodeComponentId().GetE2NodeComponentTypeGnbCuUp().GetGNbCuUpId().GetValue())
	assert.Equal(t, e2nodeComponentConfigUpdateAckItem.GetE2NodeComponentConfigUpdateAck().GetUpdateOutcome(), result.GetE2NodeComponentConfigUpdateAck().GetUpdateOutcome())
	assert.Equal(t, int32(e2nodeComponentConfigUpdateAckItem.GetE2NodeComponentConfigUpdateAck().GetFailureCause().GetProtocol()), int32(result.GetE2NodeComponentConfigUpdateAck().GetFailureCause().GetProtocol()))
}

func Test_perEncodingE2nodeComponentConfigUpdateAckItem(t *testing.T) {

	e2nodeComponentConfigUpdateAckItem, err := createE2nodeComponentConfigUpdateAckItemMsg()
	assert.NilError(t, err, "Error creating E2nodeComponentConfigUpdateAckItem PDU")

	per, err := perEncodeE2nodeComponentConfigUpdateAckItem(e2nodeComponentConfigUpdateAckItem)
	assert.NilError(t, err)
	assert.Equal(t, 5, len(per))
	t.Logf("E2nodeComponentConfigUpdateAckItem PER\n%v", hex.Dump(per))

	result, err := perDecodeE2nodeComponentConfigUpdateAckItem(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentConfigUpdateAckItem PER - decoded\n%v", result)
	assert.Equal(t, int32(e2nodeComponentConfigUpdateAckItem.GetE2NodeComponentType()), int32(result.GetE2NodeComponentType()))
	assert.Equal(t, e2nodeComponentConfigUpdateAckItem.GetE2NodeComponentId().GetE2NodeComponentTypeGnbCuUp().GetGNbCuUpId().GetValue(), result.GetE2NodeComponentId().GetE2NodeComponentTypeGnbCuUp().GetGNbCuUpId().GetValue())
	assert.Equal(t, e2nodeComponentConfigUpdateAckItem.GetE2NodeComponentConfigUpdateAck().GetUpdateOutcome(), result.GetE2NodeComponentConfigUpdateAck().GetUpdateOutcome())
	assert.Equal(t, int32(e2nodeComponentConfigUpdateAckItem.GetE2NodeComponentConfigUpdateAck().GetFailureCause().GetProtocol()), int32(result.GetE2NodeComponentConfigUpdateAck().GetFailureCause().GetProtocol()))
}
