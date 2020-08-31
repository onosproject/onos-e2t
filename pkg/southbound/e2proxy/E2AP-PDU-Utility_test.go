// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package e2proxy

import (
	"github.com/onosproject/onos-e2t/pkg/southbound/e2proxy/e2ctypes"
	"gotest.tools/assert"
	"testing"
)

func Test_GetE2apPduType(t *testing.T) {
	e2SetupResponse := e2ctypes.E2SetupResponseT{
		ProtocolIEs: &e2ctypes.ProtocolIE_Container_1544P12T{
			List: make([]*e2ctypes.E2SetupResponseIEsT, 0),
		},
	}

	e2apPdu := &e2ctypes.E2AP_PDUT{
		Choice: &e2ctypes.E2AP_PDUT_SuccessfulOutcome{
			SuccessfulOutcome: &e2ctypes.SuccessfulOutcomeT{
				ProcedureCode: e2ctypes.ProcedureCodeT_ProcedureCode_id_E2setup,
				Criticality:   e2ctypes.CriticalityT_Criticality_reject,
				Choice: &e2ctypes.SuccessfulOutcomeT_E2SetupResponse{
					E2SetupResponse: &e2SetupResponse,
				},
			},
		},
	}

	pc, err := GetE2apPduType(e2apPdu)
	assert.NilError(t, err, "unexpected error getting type from pdu")
	assert.Equal(t, "ProcedureCode_id_E2setup", pc.String())
}
