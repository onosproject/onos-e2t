// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

import (
	"encoding/hex"
	"fmt"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta2"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-commondatatypes"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	e2ap_pdu_contents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
	"gotest.tools/assert"
	"testing"
)

func createE2nodeComponentConfigUpdateAckListMsg() (*e2ap_pdu_contents.E2NodeComponentConfigUpdateAckList, error) {

	e2nodeComponentConfigUpdateAckItem := e2ap_pdu_contents.E2NodeComponentConfigUpdateAckItem{
		E2NodeComponentType: e2ap_ies.E2NodeComponentType_E2NODE_COMPONENT_TYPE_G_NB,
		//E2NodeComponentId: &e2ap_ies.E2NodeComponentId{
		//	E2NodeComponentId: &e2ap_ies.E2NodeComponentId_E2NodeComponentTypeGnbCuUp{
		//		E2NodeComponentTypeGnbCuUp: &e2ap_ies.E2NodeComponentGnbCuUpId{
		//			GNbCuUpId: &e2ap_ies.GnbCuUpId{
		//				Value: 21,
		//			},
		//		},
		//	},
		//},
		E2NodeComponentConfigUpdateAck: &e2ap_ies.E2NodeComponentConfigUpdateAck{
			UpdateOutcome: 1,
			//FailureCause: &e2ap_ies.Cause{
			//	Cause: &e2ap_ies.Cause_Protocol{
			//		Protocol: e2ap_ies.CauseProtocol_CAUSE_PROTOCOL_TRANSFER_SYNTAX_ERROR,
			//	},
			//},
		},
	}

	item := &e2ap_pdu_contents.E2NodeComponentConfigUpdateAckItemIes{
		Id:          int32(v1beta2.ProtocolIeIDE2nodeComponentConfigUpdateAckItem),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value:       &e2nodeComponentConfigUpdateAckItem,
		Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
	}

	e2nodeComponentConfigUpdateAckList := e2ap_pdu_contents.E2NodeComponentConfigUpdateAckList{
		Value: make([]*e2ap_pdu_contents.E2NodeComponentConfigUpdateAckItemIes, 0),
	}
	e2nodeComponentConfigUpdateAckList.Value = append(e2nodeComponentConfigUpdateAckList.Value, item)

	if err := e2nodeComponentConfigUpdateAckList.Validate(); err != nil {
		return nil, fmt.Errorf("error validating E2nodeComponentConfigUpdateAckList %s", err.Error())
	}
	return &e2nodeComponentConfigUpdateAckList, nil
}

func Test_xerEncodingE2nodeComponentConfigUpdateAckList(t *testing.T) {

	e2nodeComponentConfigUpdateAckList, err := createE2nodeComponentConfigUpdateAckListMsg()
	assert.NilError(t, err, "Error creating E2nodeComponentConfigUpdateAckList PDU")

	xer, err := xerEncodeE2nodeComponentConfigUpdateAckList(e2nodeComponentConfigUpdateAckList)
	assert.NilError(t, err)
	//assert.Equal(t, 943, len(xer))
	t.Logf("E2nodeComponentConfigUpdateAckList XER\n%s", string(xer))

	result, err := xerDecodeE2nodeComponentConfigUpdateAckList(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentConfigUpdateAckList XER - decoded\n%v", result)
	assert.Equal(t, 1, len(result.GetValue()))
	assert.Equal(t, int32(e2nodeComponentConfigUpdateAckList.GetValue()[0].GetValue().GetE2NodeComponentType()), int32(result.GetValue()[0].GetValue().GetE2NodeComponentType()))
	//assert.Equal(t, e2nodeComponentConfigUpdateAckList.GetValue()[0].GetValue().GetE2NodeComponentId().GetE2NodeComponentTypeGnbCuUp().GetGNbCuUpId().GetValue(), result.GetValue()[0].GetValue().GetE2NodeComponentId().GetE2NodeComponentTypeGnbCuUp().GetGNbCuUpId().GetValue())
	assert.Equal(t, e2nodeComponentConfigUpdateAckList.GetValue()[0].GetValue().GetE2NodeComponentConfigUpdateAck().GetUpdateOutcome(), result.GetValue()[0].GetValue().GetE2NodeComponentConfigUpdateAck().GetUpdateOutcome())
}

func Test_perEncodingE2nodeComponentConfigUpdateAckList(t *testing.T) {

	e2nodeComponentConfigUpdateAckList, err := createE2nodeComponentConfigUpdateAckListMsg()
	assert.NilError(t, err, "Error creating E2nodeComponentConfigUpdateAckList PDU")

	per, err := perEncodeE2nodeComponentConfigUpdateAckList(e2nodeComponentConfigUpdateAckList)
	assert.NilError(t, err)
	//assert.Equal(t, 11, len(per))
	t.Logf("E2nodeComponentConfigUpdateAckList PER\n%v", hex.Dump(per))

	result, err := perDecodeE2nodeComponentConfigUpdateAckList(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentConfigUpdateAckList PER - decoded\n%v", result)
	assert.Equal(t, 1, len(result.GetValue()))
	assert.Equal(t, int32(e2nodeComponentConfigUpdateAckList.GetValue()[0].GetValue().GetE2NodeComponentType()), int32(result.GetValue()[0].GetValue().GetE2NodeComponentType()))
	//assert.Equal(t, e2nodeComponentConfigUpdateAckList.GetValue()[0].GetValue().GetE2NodeComponentId().GetE2NodeComponentTypeGnbCuUp().GetGNbCuUpId().GetValue(), result.GetValue()[0].GetValue().GetE2NodeComponentId().GetE2NodeComponentTypeGnbCuUp().GetGNbCuUpId().GetValue())
	assert.Equal(t, e2nodeComponentConfigUpdateAckList.GetValue()[0].GetValue().GetE2NodeComponentConfigUpdateAck().GetUpdateOutcome(), result.GetValue()[0].GetValue().GetE2NodeComponentConfigUpdateAck().GetUpdateOutcome())
}
