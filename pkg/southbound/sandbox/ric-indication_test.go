// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
package sandbox

import (
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2apies"
	"gotest.tools/assert"
	"testing"
)

func TestRicIndication(t *testing.T) {
	var ricReqID int32 = 21
	var ricInstanceID int32 = 22
	var ranFuncID int32 = 9
	var ricAction = e2apies.RicactionType_RICACTION_TYPE_POLICY
	var ricIndicationType = e2apies.RicindicationType_RICINDICATION_TYPE_INSERT
	var ricSn int32 = 1
	var ricIndHd = "123"
	var ricIndMsg = "456"
	var ricCallPrID = "789"
	newE2apPdu, err := RicIndicationE2apPdu(ricReqID, ricInstanceID,
		ranFuncID, ricAction, ricSn, ricIndicationType, ricIndHd, ricIndMsg, ricCallPrID)
	assert.NilError(t, err)
	assert.Assert(t, newE2apPdu != nil)
}
