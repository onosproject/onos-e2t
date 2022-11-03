// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package f1appducontentsv1

import (
	f1apiesv1 "github.com/onosproject/onos-e2t/api/f1ap/v1/f1ap_ies"
)

func (m *SemipersistentSrs) SetSRsspatialRelation(sRsspatialRelation *f1apiesv1.SpatialRelationInfo) *SemipersistentSrs {
	m.SRsspatialRelation = sRsspatialRelation
	return m
}

func (m *SemipersistentSrs) SetIEExtensions(iEExtensions []*SemipersistentSrsExtIes) *SemipersistentSrs {
	m.IEExtensions = iEExtensions
	return m
}

func (m *AperiodicSrs) SetSRsresourceTrigger(sRsresourceTrigger *f1apiesv1.SrsresourceTrigger) *AperiodicSrs {
	m.SRsresourceTrigger = sRsresourceTrigger
	return m
}

func (m *AperiodicSrs) SetIEExtensions(iEExtensions []*AperiodicSrsExtIes) *AperiodicSrs {
	m.IEExtensions = iEExtensions
	return m
}
