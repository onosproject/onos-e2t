// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

import (
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta2"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-commondatatypes"
	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	"gotest.tools/assert"
	"testing"
)

func Test_CriticalityDiagnosticsIEItem(t *testing.T) {
	critDiagTest := e2apies.CriticalityDiagnosticsIeItem{
		IEcriticality: e2ap_commondatatypes.Criticality_CRITICALITY_NOTIFY,
		IEId: &e2ap_commondatatypes.ProtocolIeId{
			Value: int32(v1beta2.ProcedureCodeIDRICsubscription),
		},
		TypeOfError: e2apies.TypeOfError_TYPE_OF_ERROR_MISSING,
	}

	critDiatTestC, err := newCriticalityDiagnosticsIEItem(&critDiagTest)
	assert.NilError(t, err)
	assert.Assert(t, critDiatTestC != nil)

	critDiagReversed, err := decodeCriticalityDiagnosticsIEItem(critDiatTestC)
	assert.NilError(t, err)
	assert.Assert(t, critDiagReversed != nil)

	assert.Equal(t, e2ap_commondatatypes.Criticality_CRITICALITY_NOTIFY, critDiagReversed.GetIEcriticality())
	assert.Equal(t, v1beta2.ProcedureCodeIDRICsubscription, v1beta2.ProcedureCodeT(critDiagReversed.GetIEId().GetValue()))
	assert.Equal(t, e2apies.TypeOfError_TYPE_OF_ERROR_MISSING, critDiagReversed.GetTypeOfError())
}
