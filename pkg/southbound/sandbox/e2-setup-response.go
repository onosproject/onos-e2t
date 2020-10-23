// Copyright 2020-present Open Networking Foundation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package sandbox

import (
	"github.com/onosproject/onos-e2t/api/e2ap_v01_00_00_asn/v1/e2appducontents"
	"github.com/onosproject/onos-e2t/api/e2ap_v01_00_00_asn/v1/e2appdudescriptions"
)

func CreateResponseE2apPdu() *e2appdudescriptions.E2ApPdu {
	e2apPduResponse := e2appdudescriptions.E2ApPdu{
		E2ApPdu: &e2appdudescriptions.E2ApPdu_SuccessfulOutcome{
			SuccessfulOutcome: &e2appdudescriptions.SuccessfulOutcome{
				ProcedureCode: &e2appdudescriptions.E2ApElementaryProcedures{
					E2ApElementaryProcedures: &e2appdudescriptions.E2ApElementaryProcedures_Instance005{
						Instance005: &e2appdudescriptions.E2ApElementaryProcedures_E2ApElementaryProcedures005{
							SuccessfulOutcome:    &e2appducontents.E2SetupResponse{
								ProtocolIes: nil,
								},
						},
					},
				},
			},
		},
	}
	return &e2apPduResponse
}
