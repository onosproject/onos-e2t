// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package server

import (
	"strconv"

	topoapi "github.com/onosproject/onos-api/go/onos/topo"
	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/asn1cgo"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/types"
)

func GetE2NodeID(nodeID []byte, e2NodeType types.E2NodeType) (topoapi.ID, error) {
	globalE2NodeID, err := asn1cgo.PerDecodeGlobalE2nodeID(nodeID)
	if err != nil {
		return "", err
	}

	e2NodeID := topoapi.ID("")
	switch e2NodeType {
	case types.E2NodeTypeGNB:
		value := globalE2NodeID.GetGNb().GetGlobalGNbId().GnbId.GetGnbId().GetValue()
		e2NodeID = topoapi.ID(strconv.FormatUint(value, 10))
	case types.E2NodeTypeENB:
		switch enbt := globalE2NodeID.GetENb().GlobalENbId.GetENbId().GetEnbId().(type) {
		case *e2apies.EnbId_MacroENbId:
			value := enbt.MacroENbId.Value
			e2NodeID = topoapi.ID(strconv.FormatUint(value, 10))
		case *e2apies.EnbId_HomeENbId:
			value := enbt.HomeENbId.Value
			e2NodeID = topoapi.ID(strconv.FormatUint(value, 10))
		case *e2apies.EnbId_ShortMacroENbId:
			value := enbt.ShortMacroENbId.Value
			e2NodeID = topoapi.ID(strconv.FormatUint(value, 10))
		case *e2apies.EnbId_LongMacroENbId:
			value := enbt.LongMacroENbId.Value
			e2NodeID = topoapi.ID(strconv.FormatUint(value, 10))

		}
	}
	return e2NodeID, nil
}
