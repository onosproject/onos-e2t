// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2ap_ies

import (
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-commondatatypes"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
)

func (m *E2NodeComponentConfigurationAck) SetFailureCause(c *Cause) *E2NodeComponentConfigurationAck {
	m.FailureCause = c
	return m
}

func (m *E2NodeComponentInterfaceX2) SetGlobalEnbID(plmnID []byte, enbID *EnbId) *E2NodeComponentInterfaceX2 {
	m.GlobalENbId = &GlobalEnbId{
		PLmnIdentity: &e2ap_commondatatypes.PlmnIdentity{
			Value: plmnID,
		},
		ENbId: enbID,
	}
	return m
}

func (m *E2NodeComponentInterfaceX2) SetGlobalEnGnbID(plmnID []byte, bs *asn1.BitString) *E2NodeComponentInterfaceX2 {
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
