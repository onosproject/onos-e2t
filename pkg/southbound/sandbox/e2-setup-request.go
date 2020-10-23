// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
package sandbox

import (
	"github.com/onosproject/onos-e2t/api/e2ap_v01_00_00_asn/v1/e2appducontents"
	"github.com/onosproject/onos-e2t/api/e2ap_v01_00_00_asn/v1/e2appdudescriptions"
)

func CreateE2apPdu() *e2appdudescriptions.E2ApPdu {
	e2apPdu := e2appdudescriptions.E2ApPdu{
		E2ApPdu: &e2appdudescriptions.E2ApPdu_InitiatingMessage{
			InitiatingMessage: &e2appdudescriptions.InitiatingMessage{
				ProcedureCode: &e2appdudescriptions.E2ApElementaryProcedures{
					E2ApElementaryProcedures: &e2appdudescriptions.E2ApElementaryProcedures_Instance005{
						Instance005: &e2appdudescriptions.E2ApElementaryProcedures_E2ApElementaryProcedures005{
							InitiatingMessage: &e2appducontents.E2SetupRequest{
								ProtocolIes: nil,
							},
						},
					},
				},
			},
		},
	}
	return &e2apPdu
}
