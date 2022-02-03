// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package server

import (
	"github.com/onosproject/onos-lib-go/pkg/uri"

	"time"

	topoapi "github.com/onosproject/onos-api/go/onos/topo"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"

	e2 "github.com/onosproject/onos-e2t/pkg/protocols/e2ap"
)

func NewMgmtConn(nodeID topoapi.ID, plmnID string, nodeIdentity *types.E2NodeIdentity, conn e2.ServerConn,
	serviceModels map[string]*topoapi.ServiceModelInfo,
	e2Cells []*topoapi.E2Cell, now time.Time) *ManagementConn {

	connID := ConnID(uri.NewURI(
		uri.WithScheme("mgmt"),
		uri.WithOpaque(string(nodeID))).String())

	return &ManagementConn{
		ServerConn:    conn,
		ID:            connID,
		E2NodeID:      nodeID,
		PlmnID:        plmnID,
		NodeID:        string(nodeID),
		NodeType:      nodeIdentity.NodeType,
		TimeAlive:     now,
		ServiceModels: serviceModels,
		E2Cells:       e2Cells,
		E2NodeConfig:  &topoapi.E2NodeConfig{},
	}
}

type ConnID string

type ManagementConn struct {
	e2.ServerConn
	ID            ConnID
	E2NodeID      topoapi.ID
	NodeID        string
	NodeType      types.E2NodeType
	PlmnID        string
	TimeAlive     time.Time
	ServiceModels map[string]*topoapi.ServiceModelInfo
	E2Cells       []*topoapi.E2Cell
	E2NodeConfig  *topoapi.E2NodeConfig
}
