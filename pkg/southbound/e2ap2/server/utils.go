// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package server

import (
	"fmt"

	topoapi "github.com/onosproject/onos-api/go/onos/topo"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/pdudecoder"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/types"
	"github.com/onosproject/onos-lib-go/pkg/uri"
)

func GetNodeID(nodeID []byte, length int) topoapi.ID {

	e2NodeID := pdudecoder.GetE2NodeID(nodeID, length)

	e2NodeTopoID := topoapi.ID(e2NodeID)
	return e2NodeTopoID
}

func createE2NodeURI(nodeIdentity *types.E2NodeIdentity) topoapi.ID {
	e2NodeID := GetNodeID(nodeIdentity.NodeIdentifier, nodeIdentity.NodeIDLength)
	nodeType := nodeIdentity.NodeType
	var topoNodeType topoapi.NodeType
	switch nodeType {
	case types.E2NodeTypeGNB:
		topoNodeType = topoapi.NodeType_NT_GNB
	case types.E2NodeTypeEnGNB:
		topoNodeType = topoapi.NodeType_NT_EN_GNB
	case types.E2NodeTypeENB:
		topoNodeType = topoapi.NodeType_NT_ENB
	case types.E2NodeTypeNgENB:
		topoNodeType = topoapi.NodeType_NT_NG_ENB
	}
	uriOpaque := fmt.Sprintf("%d/%s", topoNodeType, e2NodeID)
	if nodeIdentity.CuID != nil {
		uriOpaque = uriOpaque + fmt.Sprintf("/%d/%x", topoapi.ComponentType_CT_CU_UP, *nodeIdentity.CuID)
	}
	if nodeIdentity.DuID != nil {
		uriOpaque = uriOpaque + fmt.Sprintf("/%d/%x", topoapi.ComponentType_CT_DU, *nodeIdentity.DuID)
	}

	uriString := uri.NewURI(
		uri.WithScheme("e2"),
		uri.WithOpaque(uriOpaque)).
		String()

	return topoapi.ID(uriString)
}
