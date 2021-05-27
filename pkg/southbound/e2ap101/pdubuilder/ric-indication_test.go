// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
package pdubuilder

import (
	"encoding/hex"
	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/asn1cgo"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/types"
	"gotest.tools/assert"
	"testing"
)

func TestRicIndication(t *testing.T) {
	ricRequestID := types.RicRequest{
		RequestorID: 21,
		InstanceID:  22,
	}
	var ranFuncID types.RanFunctionID = 9
	var ricAction = e2apies.RicactionType_RICACTION_TYPE_POLICY
	var ricIndicationType = e2apies.RicindicationType_RICINDICATION_TYPE_INSERT
	var ricSn types.RicIndicationSn = 1
	var ricIndHd types.RicIndicationHeader = []byte("123")
	var ricIndMsg types.RicIndicationMessage = []byte("456")
	var ricCallPrID types.RicCallProcessID = []byte("789")
	newE2apPdu, err := RicIndicationE2apPdu(ricRequestID,
		ranFuncID, ricAction, ricSn, ricIndicationType, ricIndHd, ricIndMsg, ricCallPrID)
	assert.NilError(t, err)
	assert.Assert(t, newE2apPdu != nil)

	xer, err := asn1cgo.XerEncodeE2apPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("RIC Indication XER\n%s", string(xer))

	per, err := asn1cgo.PerEncodeE2apPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("RIC Indication PER\n%s", hex.Dump(per))
}
