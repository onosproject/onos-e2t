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

func createE2nodeComponentConfigUpdateItemMsg() (*e2ap_pdu_contents.E2NodeComponentConfigUpdateItem, error) {

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

	if err := e2nodeComponentConfigUpdateItem.Validate(); err != nil {
		return nil, fmt.Errorf("error validating E2nodeComponentConfigUpdateItem %s", err.Error())
	}
	return &e2nodeComponentConfigUpdateItem, nil
}

func Test_xerEncodingE2nodeComponentConfigUpdateItem(t *testing.T) {

	e2nodeComponentConfigUpdateItem, err := createE2nodeComponentConfigUpdateItemMsg()
	assert.NilError(t, err, "Error creating E2nodeComponentConfigUpdateItem PDU")

	xer, err := xerEncodeE2nodeComponentConfigUpdateItem(e2nodeComponentConfigUpdateItem)
	assert.NilError(t, err)
	assert.Equal(t, 675, len(xer))
	t.Logf("E2nodeComponentConfigUpdateItem XER\n%s", string(xer))

	result, err := xerDecodeE2nodeComponentConfigUpdateItem(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentConfigUpdateItem XER - decoded\n%v", result)
	assert.Equal(t, int32(e2nodeComponentConfigUpdateItem.GetE2NodeComponentType()), int32(result.GetE2NodeComponentType()))
	assert.Equal(t, e2nodeComponentConfigUpdateItem.GetE2NodeComponentId().GetE2NodeComponentTypeGnbCuUp().GetGNbCuUpId().GetValue(), result.GetE2NodeComponentId().GetE2NodeComponentTypeGnbCuUp().GetGNbCuUpId().GetValue())
	assert.Equal(t, e2nodeComponentConfigUpdateItem.GetE2NodeComponentConfigUpdate().GetGNbconfigUpdate().GetE1ApconfigUpdate(), result.GetE2NodeComponentConfigUpdate().GetGNbconfigUpdate().GetE1ApconfigUpdate())
	assert.Equal(t, e2nodeComponentConfigUpdateItem.GetE2NodeComponentConfigUpdate().GetGNbconfigUpdate().GetF1ApconfigUpdate(), result.GetE2NodeComponentConfigUpdate().GetGNbconfigUpdate().GetF1ApconfigUpdate())
	assert.Equal(t, e2nodeComponentConfigUpdateItem.GetE2NodeComponentConfigUpdate().GetGNbconfigUpdate().GetXnApconfigUpdate(), result.GetE2NodeComponentConfigUpdate().GetGNbconfigUpdate().GetXnApconfigUpdate())
	assert.Equal(t, e2nodeComponentConfigUpdateItem.GetE2NodeComponentConfigUpdate().GetGNbconfigUpdate().GetNgApconfigUpdate(), result.GetE2NodeComponentConfigUpdate().GetGNbconfigUpdate().GetNgApconfigUpdate())
}

func Test_perEncodingE2nodeComponentConfigUpdateItem(t *testing.T) {

	e2nodeComponentConfigUpdateItem, err := createE2nodeComponentConfigUpdateItemMsg()
	assert.NilError(t, err, "Error creating E2nodeComponentConfigUpdateItem PDU")

	per, err := perEncodeE2nodeComponentConfigUpdateItem(e2nodeComponentConfigUpdateItem)
	assert.NilError(t, err)
	assert.Equal(t, 28, len(per))
	t.Logf("E2nodeComponentConfigUpdateItem PER\n%v", hex.Dump(per))

	result, err := perDecodeE2nodeComponentConfigUpdateItem(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentConfigUpdateItem PER - decoded\n%v", result)
	assert.Equal(t, int32(e2nodeComponentConfigUpdateItem.GetE2NodeComponentType()), int32(result.GetE2NodeComponentType()))
	assert.Equal(t, e2nodeComponentConfigUpdateItem.GetE2NodeComponentId().GetE2NodeComponentTypeGnbCuUp().GetGNbCuUpId().GetValue(), result.GetE2NodeComponentId().GetE2NodeComponentTypeGnbCuUp().GetGNbCuUpId().GetValue())
	assert.Equal(t, e2nodeComponentConfigUpdateItem.GetE2NodeComponentConfigUpdate().GetGNbconfigUpdate().GetE1ApconfigUpdate(), result.GetE2NodeComponentConfigUpdate().GetGNbconfigUpdate().GetE1ApconfigUpdate())
	assert.Equal(t, e2nodeComponentConfigUpdateItem.GetE2NodeComponentConfigUpdate().GetGNbconfigUpdate().GetF1ApconfigUpdate(), result.GetE2NodeComponentConfigUpdate().GetGNbconfigUpdate().GetF1ApconfigUpdate())
	assert.Equal(t, e2nodeComponentConfigUpdateItem.GetE2NodeComponentConfigUpdate().GetGNbconfigUpdate().GetXnApconfigUpdate(), result.GetE2NodeComponentConfigUpdate().GetGNbconfigUpdate().GetXnApconfigUpdate())
	assert.Equal(t, e2nodeComponentConfigUpdateItem.GetE2NodeComponentConfigUpdate().GetGNbconfigUpdate().GetNgApconfigUpdate(), result.GetE2NodeComponentConfigUpdate().GetGNbconfigUpdate().GetNgApconfigUpdate())
}
