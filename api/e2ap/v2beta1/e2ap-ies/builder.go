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

package e2ap_ies

import (
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-commondatatypes"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
)

func (m *GlobalE2NodeGnbId) SetGlobalEnGNbID(plmnID []byte, bs *asn1.BitString) *GlobalE2NodeGnbId {
	m.GlobalEnGNbId = &GlobalenGnbId{
		PLmnIdentity: &e2ap_commondatatypes.PlmnIdentity{
			Value: plmnID,
		},
		GNbId: &EngnbId{
			EngnbId: &EngnbId_GNbId{
				GNbId: bs,
			},
		},
	}
	return m
}

func (m *GlobalE2NodeGnbId) SetGnbCuUpID(val int64) *GlobalE2NodeGnbId {
	m.GNbCuUpId = &GnbCuUpId{
		Value: val,
	}
	return m
}

func (m *GlobalE2NodeGnbId) SetGnbDuID(val int64) *GlobalE2NodeGnbId {
	m.GNbDuId = &GnbDuId{
		Value: val,
	}
	return m
}

func (m *GlobalE2NodeEnGnbId) SetGnbCuUpID(val int64) *GlobalE2NodeEnGnbId {
	m.EnGNbCuUpId = &GnbCuUpId{
		Value: val,
	}
	return m
}

func (m *GlobalE2NodeEnGnbId) SetGnbDuID(val int64) *GlobalE2NodeEnGnbId {
	m.EnGNbDuId = &GnbDuId{
		Value: val,
	}
	return m
}

func (m *GlobalE2NodeNgEnbId) SetGlobalEnbID(plmnID []byte, enbID *EnbId) *GlobalE2NodeNgEnbId {
	m.GlobalENbId = &GlobalEnbId{
		PLmnIdentity: &e2ap_commondatatypes.PlmnIdentity{
			Value: plmnID,
		},
		ENbId: enbID,
	}
	return m
}

func (m *GlobalE2NodeNgEnbId) SetNgEnbDuID(val int64) *GlobalE2NodeNgEnbId {
	m.NgEnbDuId = &NgenbDuId{
		Value: val,
	}
	return m
}

func (m *Tnlinformation) SetTnlPort(bs *asn1.BitString) *Tnlinformation {
	m.TnlPort = bs
	return m
}
