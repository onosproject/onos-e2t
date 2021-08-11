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

func createE2nodeComponentConfigUpdateListMsg() (*e2ap_pdu_contents.E2NodeComponentConfigUpdateList, error) {

	e2nodeComponentConfigUpdateItem := e2ap_pdu_contents.E2NodeComponentConfigUpdateItem{
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
		E2NodeComponentConfigUpdate: &e2ap_ies.E2NodeComponentConfigUpdate{
			E2NodeComponentConfigUpdate: &e2ap_ies.E2NodeComponentConfigUpdate_GNbconfigUpdate{
				GNbconfigUpdate: &e2ap_ies.E2NodeComponentConfigUpdateGnb{
					NgApconfigUpdate: "ng_AP",
					XnApconfigUpdate: "xn_AP",
					E1ApconfigUpdate: "e1_AP",
					F1ApconfigUpdate: "f1_AP",
				},
			},
		},
	}

	item := e2ap_pdu_contents.E2NodeComponentConfigUpdateItemIes{
		Id:          int32(v1beta2.ProtocolIeIDE2nodeComponentConfigUpdateItem),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value:       &e2nodeComponentConfigUpdateItem,
		Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
	}

	e2nodeComponentConfigUpdateList := e2ap_pdu_contents.E2NodeComponentConfigUpdateList{
		Value: make([]*e2ap_pdu_contents.E2NodeComponentConfigUpdateItemIes, 0),
	}
	e2nodeComponentConfigUpdateList.Value = append(e2nodeComponentConfigUpdateList.Value, &item)

	if err := e2nodeComponentConfigUpdateList.Validate(); err != nil {
		return nil, fmt.Errorf("error validating E2nodeComponentConfigUpdateList %s", err.Error())
	}
	return &e2nodeComponentConfigUpdateList, nil
}

func Test_xerEncodingE2nodeComponentConfigUpdateList(t *testing.T) {

	e2nodeComponentConfigUpdateList, err := createE2nodeComponentConfigUpdateListMsg()
	assert.NilError(t, err, "Error creating E2nodeComponentConfigUpdateList PDU")

	xer, err := xerEncodeE2nodeComponentConfigUpdateList(e2nodeComponentConfigUpdateList)
	assert.NilError(t, err)
	assert.Equal(t, 1103, len(xer))
	t.Logf("E2nodeComponentConfigUpdateList XER\n%s", string(xer))

	result, err := xerDecodeE2nodeComponentConfigUpdateList(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentConfigUpdateList XER - decoded\n%v", result)
	assert.Equal(t, 1, len(result.GetValue()))
	assert.Equal(t, int32(e2nodeComponentConfigUpdateList.GetValue()[0].GetValue().GetE2NodeComponentType()), int32(result.GetValue()[0].GetValue().GetE2NodeComponentType()))
	assert.Equal(t, e2nodeComponentConfigUpdateList.GetValue()[0].GetValue().GetE2NodeComponentId().GetE2NodeComponentTypeGnbCuUp().GetGNbCuUpId().GetValue(), result.GetValue()[0].GetValue().GetE2NodeComponentId().GetE2NodeComponentTypeGnbCuUp().GetGNbCuUpId().GetValue())
	assert.Equal(t, e2nodeComponentConfigUpdateList.GetValue()[0].GetValue().GetE2NodeComponentConfigUpdate().GetGNbconfigUpdate().GetE1ApconfigUpdate(), result.GetValue()[0].GetValue().GetE2NodeComponentConfigUpdate().GetGNbconfigUpdate().GetE1ApconfigUpdate())
	assert.Equal(t, e2nodeComponentConfigUpdateList.GetValue()[0].GetValue().GetE2NodeComponentConfigUpdate().GetGNbconfigUpdate().GetF1ApconfigUpdate(), result.GetValue()[0].GetValue().GetE2NodeComponentConfigUpdate().GetGNbconfigUpdate().GetF1ApconfigUpdate())
	assert.Equal(t, e2nodeComponentConfigUpdateList.GetValue()[0].GetValue().GetE2NodeComponentConfigUpdate().GetGNbconfigUpdate().GetXnApconfigUpdate(), result.GetValue()[0].GetValue().GetE2NodeComponentConfigUpdate().GetGNbconfigUpdate().GetXnApconfigUpdate())
	assert.Equal(t, e2nodeComponentConfigUpdateList.GetValue()[0].GetValue().GetE2NodeComponentConfigUpdate().GetGNbconfigUpdate().GetNgApconfigUpdate(), result.GetValue()[0].GetValue().GetE2NodeComponentConfigUpdate().GetGNbconfigUpdate().GetNgApconfigUpdate())
}

func Test_perEncodingE2nodeComponentConfigUpdateList(t *testing.T) {

	e2nodeComponentConfigUpdateList, err := createE2nodeComponentConfigUpdateListMsg()
	assert.NilError(t, err, "Error creating E2nodeComponentConfigUpdateList PDU")

	per, err := perEncodeE2nodeComponentConfigUpdateList(e2nodeComponentConfigUpdateList)
	assert.NilError(t, err)
	assert.Equal(t, 34, len(per))
	t.Logf("E2nodeComponentConfigUpdateList PER\n%v", hex.Dump(per))

	result, err := perDecodeE2nodeComponentConfigUpdateList(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentConfigUpdateList PER - decoded\n%v", result)
	assert.Equal(t, 1, len(result.GetValue()))
	assert.Equal(t, int32(e2nodeComponentConfigUpdateList.GetValue()[0].GetValue().GetE2NodeComponentType()), int32(result.GetValue()[0].GetValue().GetE2NodeComponentType()))
	assert.Equal(t, e2nodeComponentConfigUpdateList.GetValue()[0].GetValue().GetE2NodeComponentId().GetE2NodeComponentTypeGnbCuUp().GetGNbCuUpId().GetValue(), result.GetValue()[0].GetValue().GetE2NodeComponentId().GetE2NodeComponentTypeGnbCuUp().GetGNbCuUpId().GetValue())
	assert.Equal(t, e2nodeComponentConfigUpdateList.GetValue()[0].GetValue().GetE2NodeComponentConfigUpdate().GetGNbconfigUpdate().GetE1ApconfigUpdate(), result.GetValue()[0].GetValue().GetE2NodeComponentConfigUpdate().GetGNbconfigUpdate().GetE1ApconfigUpdate())
	assert.Equal(t, e2nodeComponentConfigUpdateList.GetValue()[0].GetValue().GetE2NodeComponentConfigUpdate().GetGNbconfigUpdate().GetF1ApconfigUpdate(), result.GetValue()[0].GetValue().GetE2NodeComponentConfigUpdate().GetGNbconfigUpdate().GetF1ApconfigUpdate())
	assert.Equal(t, e2nodeComponentConfigUpdateList.GetValue()[0].GetValue().GetE2NodeComponentConfigUpdate().GetGNbconfigUpdate().GetXnApconfigUpdate(), result.GetValue()[0].GetValue().GetE2NodeComponentConfigUpdate().GetGNbconfigUpdate().GetXnApconfigUpdate())
	assert.Equal(t, e2nodeComponentConfigUpdateList.GetValue()[0].GetValue().GetE2NodeComponentConfigUpdate().GetGNbconfigUpdate().GetNgApconfigUpdate(), result.GetValue()[0].GetValue().GetE2NodeComponentConfigUpdate().GetGNbconfigUpdate().GetNgApconfigUpdate())
}
