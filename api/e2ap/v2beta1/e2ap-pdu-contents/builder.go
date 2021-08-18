// Copyright 2021-present Open Networking Foundation.
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

package e2ap_pdu_contents

import (
	"github.com/onosproject/onos-e2t/api/e2ap/v2beta1"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-commondatatypes"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-ies"
)

func (m *E2SetupFailure) SetTnlInformation(tnl *e2ap_ies.Tnlinformation) *E2SetupFailure {
	m.GetProtocolIes().E2ApProtocolIes48 = &E2SetupFailureIes_E2SetupFailureIes48{
		Id:          int32(v2beta1.ProtocolIeIDTNLinformation),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
		Value:       tnl,
		Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
	}
	return m
}
