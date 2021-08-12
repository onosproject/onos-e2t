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

func createRanFunctionsIDListMsg() (*e2ap_pdu_contents.RanfunctionsIdList, error) {

	rfIDl := e2ap_pdu_contents.RanfunctionsIdList{
		Value: make([]*e2ap_pdu_contents.RanfunctionIdItemIes, 0),
	}

	rfIDi := &e2ap_pdu_contents.RanfunctionIdItemIes{
		RanFunctionIdItemIes6: &e2ap_pdu_contents.RanfunctionIdItemIes_RanfunctionIdItemIes6{
			Id:          int32(v1beta2.ProtocolIeIDRanfunctionIDItem),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
			Value: &e2ap_pdu_contents.RanfunctionIdItem{
				RanFunctionId: &e2ap_ies.RanfunctionId{
					Value: 123,
				},
				RanFunctionRevision: &e2ap_ies.RanfunctionRevision{
					Value: 8,
				},
			},
			Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		},
	}
	rfIDl.Value = append(rfIDl.Value, rfIDi)

	if err := rfIDl.Validate(); err != nil {
		return nil, fmt.Errorf("error validating RANfunctionsIDList %s", err.Error())
	}
	return &rfIDl, nil
}

func Test_xerEncodingRanFunctionIDList(t *testing.T) {

	rfIDl, err := createRanFunctionsIDListMsg()
	assert.NilError(t, err, "Error creating RANfunctionsIDList PDU")

	xer, err := xerEncodeRanFunctionsIDList(rfIDl)
	assert.NilError(t, err)
	assert.Equal(t, 388, len(xer))
	t.Logf("RANfunctionsIDList XER\n%s", string(xer))

	result, err := xerDecodeRanFunctionsIDList(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("RANfunctionsIDList XER - decoded\n%v", result)
	assert.Equal(t, rfIDl.GetValue()[0].GetRanFunctionIdItemIes6().GetValue().GetRanFunctionRevision().GetValue(), result.GetValue()[0].GetRanFunctionIdItemIes6().GetValue().GetRanFunctionRevision().GetValue())
	assert.Equal(t, rfIDl.GetValue()[0].GetRanFunctionIdItemIes6().GetValue().GetRanFunctionId().GetValue(), result.GetValue()[0].GetRanFunctionIdItemIes6().GetValue().GetRanFunctionId().GetValue())
}

func Test_perEncodingRanFunctionIDList(t *testing.T) {

	rfIDl, err := createRanFunctionsIDListMsg()
	assert.NilError(t, err, "Error creating RANfunctionsIDList PDU")

	per, err := perEncodeRanFunctionsIDList(rfIDl)
	assert.NilError(t, err)
	assert.Equal(t, 11, len(per))
	t.Logf("RANfunctionsIDList PER\n%v", hex.Dump(per))

	result, err := perDecodeRanFunctionsIDList(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("RANfunctionsIDList PER - decoded\n%v", result)
	assert.Equal(t, rfIDl.GetValue()[0].GetRanFunctionIdItemIes6().GetValue().GetRanFunctionRevision().GetValue(), result.GetValue()[0].GetRanFunctionIdItemIes6().GetValue().GetRanFunctionRevision().GetValue())
	assert.Equal(t, rfIDl.GetValue()[0].GetRanFunctionIdItemIes6().GetValue().GetRanFunctionId().GetValue(), result.GetValue()[0].GetRanFunctionIdItemIes6().GetValue().GetRanFunctionId().GetValue())
}
