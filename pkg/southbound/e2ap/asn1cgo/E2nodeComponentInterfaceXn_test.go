// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package asn1cgo

import (
	"encoding/hex"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-commondatatypes"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"testing"

	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	"gotest.tools/assert"
)

func createE2nodeComponentInterfaceXnMsg() (*e2ap_ies.E2NodeComponentInterfaceXn, error) {

	e2ncc := e2ap_ies.E2NodeComponentInterfaceXn{
		GlobalNgRanNodeId: &e2ap_ies.GlobalNgRannodeId{
			GlobalNgRannodeId: &e2ap_ies.GlobalNgRannodeId_NgENb{
				NgENb: &e2ap_ies.GlobalngeNbId{
					PlmnId: &e2ap_commondatatypes.PlmnIdentity{
						Value: []byte{0x01, 0x02, 0x03},
					},
					EnbId: &e2ap_ies.EnbIdChoice{
						EnbIdChoice: &e2ap_ies.EnbIdChoice_EnbIdMacro{
							EnbIdMacro: &asn1.BitString{
								Value: []byte{0x00, 0x00, 0x10},
								Len:   20,
							},
						},
					},
				},
			},
		},
	}

	//if err := e2nodeComponentConfigUpdateGnb.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validatiXn E2nodeComponentConfigUpdateGnb %s", err.Error())
	//}
	return &e2ncc, nil
}

func Test_xerEncodingXnE2nodeComponentInterfaceXn(t *testing.T) {

	e2ncc, err := createE2nodeComponentInterfaceXnMsg()
	assert.NilError(t, err, "Error creatiXn E2nodeComponentInterfaceXn PDU")

	xer, err := xerEncodeE2nodeComponentInterfaceXn(e2ncc)
	assert.NilError(t, err)
	t.Logf("E2nodeComponentInterfaceXn XER\n%s", string(xer))

	result, err := xerDecodeE2nodeComponentInterfaceXn(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentInterfaceXn XER - decoded\n%v", result)
	assert.DeepEqual(t, e2ncc.GetGlobalNgRanNodeId().GetNgENb().GetPlmnId().GetValue(), result.GetGlobalNgRanNodeId().GetNgENb().GetPlmnId().GetValue())
	assert.DeepEqual(t, e2ncc.GetGlobalNgRanNodeId().GetNgENb().GetEnbId().GetEnbIdMacro().GetValue(), result.GetGlobalNgRanNodeId().GetNgENb().GetEnbId().GetEnbIdMacro().GetValue())
}

func Test_perEncodingXnE2nodeComponentInterfaceXn(t *testing.T) {

	e2ncc, err := createE2nodeComponentInterfaceXnMsg()
	assert.NilError(t, err, "Error creatiXn E2nodeComponentInterfaceXn PDU")

	per, err := perEncodeE2nodeComponentInterfaceXn(e2ncc)
	assert.NilError(t, err)
	t.Logf("E2nodeComponentInterfaceXn PER\n%v", hex.Dump(per))

	result, err := perDecodeE2nodeComponentInterfaceXn(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentInterfaceXn PER - decoded\n%v", result)
	assert.DeepEqual(t, e2ncc.GetGlobalNgRanNodeId().GetNgENb().GetPlmnId().GetValue(), result.GetGlobalNgRanNodeId().GetNgENb().GetPlmnId().GetValue())
	assert.DeepEqual(t, e2ncc.GetGlobalNgRanNodeId().GetNgENb().GetEnbId().GetEnbIdMacro().GetValue(), result.GetGlobalNgRanNodeId().GetNgENb().GetEnbId().GetEnbIdMacro().GetValue())
}
