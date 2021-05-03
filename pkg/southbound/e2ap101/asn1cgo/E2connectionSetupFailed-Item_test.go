// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

import (
	"encoding/hex"
	"fmt"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-commondatatypes"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	e2ap_pdu_contents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
	"gotest.tools/assert"
	"testing"
)

func createE2connectionSetupFailedItemMsg() (*e2ap_pdu_contents.E2ConnectionSetupFailedItem, error) {

	bs1 := &e2ap_commondatatypes.BitString{
		Value: 0x89bcd,
		Len:   64,
	}

	bs2 := &e2ap_commondatatypes.BitString{
		Value: 0x89bcd,
		Len:   16,
	}

	tnlinformation := e2ap_ies.Tnlinformation{
		TnlAddress: bs1,
		TnlPort:    bs2,
	}

	e2connectionSetupFailedItem := e2ap_pdu_contents.E2ConnectionSetupFailedItem{
		TnlInformation: &tnlinformation,
		Cause: &e2ap_ies.Cause{
			Cause: &e2ap_ies.Cause_RicService{
				RicService: e2ap_ies.CauseRicservice_CAUSE_RICSERVICE_RIC_RESOURCE_LIMIT,
			},
		},
	}

	if err := e2connectionSetupFailedItem.Validate(); err != nil {
		return nil, fmt.Errorf("error validating E2connectionSetupFailedItem %s", err.Error())
	}
	return &e2connectionSetupFailedItem, nil
}

func Test_xerEncodingE2connectionSetupFailedItem(t *testing.T) {

	e2connectionSetupFailedItem, err := createE2connectionSetupFailedItemMsg()
	assert.NilError(t, err, "Error creating E2connectionSetupFailedItem PDU")

	xer, err := xerEncodeE2connectionSetupFailedItem(e2connectionSetupFailedItem)
	assert.NilError(t, err)
	assert.Equal(t, 372, len(xer))
	t.Logf("E2connectionSetupFailedItem XER\n%s", string(xer))

	result, err := xerDecodeE2connectionSetupFailedItem(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2connectionSetupFailedItem XER - decoded\n%v", result)
	assert.Equal(t, e2connectionSetupFailedItem.GetTnlInformation().GetTnlPort().GetLen(), result.GetTnlInformation().GetTnlPort().GetLen())
	assert.Equal(t, e2connectionSetupFailedItem.GetTnlInformation().GetTnlAddress().GetLen(), result.GetTnlInformation().GetTnlAddress().GetLen())
	assert.Equal(t, e2connectionSetupFailedItem.GetCause().GetRicService(), result.GetCause().GetRicService())
}

func Test_perEncodingE2connectionSetupFailedItem(t *testing.T) {

	e2connectionSetupFailedItem, err := createE2connectionSetupFailedItemMsg()
	assert.NilError(t, err, "Error creating E2connectionSetupFailedItem PDU")

	per, err := perEncodeE2connectionSetupFailedItem(e2connectionSetupFailedItem)
	assert.NilError(t, err)
	assert.Equal(t, 13, len(per))
	t.Logf("E2connectionSetupFailedItem PER\n%v", hex.Dump(per))

	result, err := perDecodeE2connectionSetupFailedItem(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2connectionSetupFailedItem PER - decoded\n%v", result)
	assert.Equal(t, e2connectionSetupFailedItem.GetTnlInformation().GetTnlPort().GetLen(), result.GetTnlInformation().GetTnlPort().GetLen())
	assert.Equal(t, e2connectionSetupFailedItem.GetTnlInformation().GetTnlAddress().GetLen(), result.GetTnlInformation().GetTnlAddress().GetLen())
	assert.Equal(t, e2connectionSetupFailedItem.GetCause().GetRicService(), result.GetCause().GetRicService())
}
