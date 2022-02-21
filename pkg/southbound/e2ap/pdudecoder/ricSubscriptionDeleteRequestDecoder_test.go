// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdudecoder

import (
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/pdubuilder"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
	"gotest.tools/assert"
	"testing"
)

func Test_DecodeRicSubscriptionDeleteRequestPdu(t *testing.T) {
	e2apPdu, err := pdubuilder.CreateRicSubscriptionDeleteRequestE2apPdu(
		types.RicRequest{RequestorID: 1, InstanceID: 2},
		3)

	assert.NilError(t, err)
	assert.Assert(t, e2apPdu != nil)

	ricReq, ranFuncID, err := DecodeRicSubscriptionDeleteRequestPdu(e2apPdu)
	assert.NilError(t, err)
	assert.Equal(t, 1, int(ricReq.RequestorID))
	assert.Equal(t, 2, int(ricReq.InstanceID))
	assert.Equal(t, 3, int(ranFuncID))
}
