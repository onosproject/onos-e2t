// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package asn1cgo

import (
	"encoding/hex"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-commondatatypes"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"testing"

	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	"gotest.tools/assert"
)

func createE2nodeComponentInterfaceX2Msg() (*e2ap_ies.E2NodeComponentInterfaceX2, error) {

	e2ncc := e2ap_ies.E2NodeComponentInterfaceX2{
		GlobalENbId: &e2ap_ies.GlobalEnbId{
			PLmnIdentity: &e2ap_commondatatypes.PlmnIdentity{
				Value: []byte{0x01, 0x02, 0x03},
			},
			ENbId: &e2ap_ies.EnbId{
				EnbId: &e2ap_ies.EnbId_MacroENbId{
					MacroENbId: &asn1.BitString{
						Value: []byte{0x00, 0x00, 0x10},
						Len:   20,
					},
				},
			},
		},
		GlobalEnGNbId: &e2ap_ies.GlobalenGnbId{
			PLmnIdentity: &e2ap_commondatatypes.PlmnIdentity{
				Value: []byte{0x01, 0x02, 0x03},
			},
			GNbId: &e2ap_ies.EngnbId{
				EngnbId: &e2ap_ies.EngnbId_GNbId{
					GNbId: &asn1.BitString{
						Value: []byte{0x00, 0x00, 0x01},
						Len:   24,
					},
				},
			},
		},
	}

	//if err := e2nodeComponentConfigUpdateGnb.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validatiX2 E2nodeComponentConfigUpdateGnb %s", err.Error())
	//}
	return &e2ncc, nil
}

func Test_xerEncodingX2E2nodeComponentInterfaceX2(t *testing.T) {

	e2ncc, err := createE2nodeComponentInterfaceX2Msg()
	assert.NilError(t, err, "Error creatiX2 E2nodeComponentInterfaceX2 PDU")

	xer, err := xerEncodeE2nodeComponentInterfaceX2(e2ncc)
	assert.NilError(t, err)
	t.Logf("E2nodeComponentInterfaceX2 XER\n%s", string(xer))

	result, err := xerDecodeE2nodeComponentInterfaceX2(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentInterfaceX2 XER - decoded\n%v", result)
	assert.DeepEqual(t, e2ncc.GetGlobalENbId().GetPLmnIdentity().GetValue(), result.GetGlobalENbId().GetPLmnIdentity().GetValue())
	assert.DeepEqual(t, e2ncc.GetGlobalENbId().GetENbId().GetMacroENbId().GetValue(), result.GetGlobalENbId().GetENbId().GetMacroENbId().GetValue())
	assert.DeepEqual(t, e2ncc.GetGlobalEnGNbId().GetPLmnIdentity().GetValue(), result.GetGlobalEnGNbId().GetPLmnIdentity().GetValue())
	assert.DeepEqual(t, e2ncc.GetGlobalEnGNbId().GetGNbId().GetGNbId().GetValue(), result.GetGlobalEnGNbId().GetGNbId().GetGNbId().GetValue())
}

func Test_perEncodingX2E2nodeComponentInterfaceX2(t *testing.T) {

	e2ncc, err := createE2nodeComponentInterfaceX2Msg()
	assert.NilError(t, err, "Error creatiX2 E2nodeComponentInterfaceX2 PDU")

	per, err := perEncodeE2nodeComponentInterfaceX2(e2ncc)
	assert.NilError(t, err)
	t.Logf("E2nodeComponentInterfaceX2 PER\n%v", hex.Dump(per))

	result, err := perDecodeE2nodeComponentInterfaceX2(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentInterfaceX2 PER - decoded\n%v", result)
	assert.DeepEqual(t, e2ncc.GetGlobalENbId().GetPLmnIdentity().GetValue(), result.GetGlobalENbId().GetPLmnIdentity().GetValue())
	assert.DeepEqual(t, e2ncc.GetGlobalENbId().GetENbId().GetMacroENbId().GetValue(), result.GetGlobalENbId().GetENbId().GetMacroENbId().GetValue())
	assert.DeepEqual(t, e2ncc.GetGlobalEnGNbId().GetPLmnIdentity().GetValue(), result.GetGlobalEnGNbId().GetPLmnIdentity().GetValue())
	assert.DeepEqual(t, e2ncc.GetGlobalEnGNbId().GetGNbId().GetGNbId().GetValue(), result.GetGlobalEnGNbId().GetGNbId().GetGNbId().GetValue())
}
