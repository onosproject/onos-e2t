// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package types

import (
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
)

type E2NodeComponentConfigUpdateItem struct {
	E2NodeComponentType          e2ap_ies.E2NodeComponentInterfaceType
	E2NodeComponentID            *e2ap_ies.E2NodeComponentId
	E2NodeComponentConfiguration e2ap_ies.E2NodeComponentConfiguration
}

type E2NodeComponentConfigUpdateAckItem struct {
	E2NodeComponentType             e2ap_ies.E2NodeComponentInterfaceType
	E2NodeComponentID               *e2ap_ies.E2NodeComponentId
	E2NodeComponentConfigurationAck E2NodeComponentConfigurationAck
}

type E2NodeComponentConfigurationAck struct {
	UpdateOutcome e2ap_ies.UpdateOutcome
	FailureCause  *e2ap_ies.Cause
}

type E2NodeComponentConfigAdditionItem struct {
	E2NodeComponentType          e2ap_ies.E2NodeComponentInterfaceType
	E2NodeComponentID            *e2ap_ies.E2NodeComponentId
	E2NodeComponentConfiguration e2ap_ies.E2NodeComponentConfiguration
}

type E2NodeComponentConfigAdditionAckItem struct {
	E2NodeComponentType             e2ap_ies.E2NodeComponentInterfaceType
	E2NodeComponentID               *e2ap_ies.E2NodeComponentId
	E2NodeComponentConfigurationAck e2ap_ies.E2NodeComponentConfigurationAck
}

type E2NodeComponentConfigRemovalItem struct {
	E2NodeComponentType e2ap_ies.E2NodeComponentInterfaceType
	E2NodeComponentID   *e2ap_ies.E2NodeComponentId
}

type E2NodeComponentConfigRemovalAckItem struct {
	E2NodeComponentType             e2ap_ies.E2NodeComponentInterfaceType
	E2NodeComponentID               *e2ap_ies.E2NodeComponentId
	E2NodeComponentConfigurationAck e2ap_ies.E2NodeComponentConfigurationAck
}
