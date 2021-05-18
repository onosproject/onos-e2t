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

func createE2connectionUpdateListMsg() (*e2ap_pdu_contents.E2ConnectionUpdateList, error) {

	e2connectionUpdateList := e2ap_pdu_contents.E2ConnectionUpdateList{
		Value: make([]*e2ap_pdu_contents.E2ConnectionUpdateItemIes, 0),
	}

	item := &e2ap_pdu_contents.E2ConnectionUpdateItemIes{
		Id:          int32(v1beta2.ProtocolIeIDE2connectionUpdateItem),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
		Value: &e2ap_pdu_contents.E2ConnectionUpdateItem{
			TnlInformation: &e2ap_ies.Tnlinformation{
				TnlPort: &e2ap_commondatatypes.BitString{
					Value: 0x89bcd,
					Len:   16,
				},
				TnlAddress: &e2ap_commondatatypes.BitString{
					Value: 0x89abcdef01234567,
					Len:   64,
				},
			},
			TnlUsage: e2ap_ies.Tnlusage_TNLUSAGE_BOTH,
		},
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
	}

	e2connectionUpdateList.Value = append(e2connectionUpdateList.Value, item)

	if err := e2connectionUpdateList.Validate(); err != nil {
		return nil, fmt.Errorf("error validating E2connectionUpdateList %s", err.Error())
	}
	return &e2connectionUpdateList, nil
}

func Test_xerEncodingE2connectionUpdateList(t *testing.T) {

	e2connectionUpdateList, err := createE2connectionUpdateListMsg()
	assert.NilError(t, err, "Error creating E2connectionUpdateList PDU")

	xer, err := xerEncodeE2connectionUpdateList(e2connectionUpdateList)
	assert.NilError(t, err)
	assert.Equal(t, 665, len(xer))
	t.Logf("E2connectionUpdateList XER\n%s", string(xer))

	result, err := xerDecodeE2connectionUpdateList(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2connectionUpdateList XER - decoded\n%v", result)
	assert.Equal(t, 1, len(result.GetValue()))
	assert.Equal(t, e2connectionUpdateList.GetValue()[0].GetValue().GetTnlUsage(), result.GetValue()[0].GetValue().GetTnlUsage())
	assert.Equal(t, e2connectionUpdateList.GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetLen(), result.GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetLen())
	assert.Equal(t, e2connectionUpdateList.GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetLen(), result.GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetLen())
}

func Test_perEncodingE2connectionUpdateList(t *testing.T) {

	e2connectionUpdateList, err := createE2connectionUpdateListMsg()
	assert.NilError(t, err, "Error creating E2connectionUpdateList PDU")

	per, err := perEncodeE2connectionUpdateList(e2connectionUpdateList)
	assert.NilError(t, err)
	assert.Equal(t, 18, len(per))
	t.Logf("E2connectionUpdateList PER\n%v", hex.Dump(per))

	result, err := perDecodeE2connectionUpdateList(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2connectionUpdateList PER - decoded\n%v", result)
	assert.Equal(t, 1, len(result.GetValue()))
	assert.Equal(t, e2connectionUpdateList.GetValue()[0].GetValue().GetTnlUsage(), result.GetValue()[0].GetValue().GetTnlUsage())
	assert.Equal(t, e2connectionUpdateList.GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetLen(), result.GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetLen())
	assert.Equal(t, e2connectionUpdateList.GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetLen(), result.GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetLen())

}
