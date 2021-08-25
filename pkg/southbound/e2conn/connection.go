// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package e2conn

import (
	"time"

	"github.com/onosproject/onos-lib-go/pkg/logging"

	"github.com/onosproject/onos-e2t/pkg/protocols/e2ap/connection"

	"github.com/google/uuid"
	"github.com/onosproject/onos-lib-go/pkg/uri"

	topoapi "github.com/onosproject/onos-api/go/onos/topo"
)

var log = logging.GetLogger("southbound", "e2", "connection")

type E2BaseConn interface {
	connection.Conn
	GetID() ID
	GetE2NodeID() topoapi.ID
	GetServiceModels() map[string]*topoapi.ServiceModelInfo
	GetE2Cells() []*topoapi.E2Cell
	GetPlmnID() string
	GetTimeAlive() time.Time
}

type ID string

func NewE2BaseConn(nodeID topoapi.ID, plmnID string,
	serviceModels map[string]*topoapi.ServiceModelInfo,
	e2Cells []*topoapi.E2Cell, now time.Time) *E2BaseConnection {

	connID := ID(uri.NewURI(
		uri.WithScheme("uuid"),
		uri.WithOpaque(uuid.New().String())).String())

	return &E2BaseConnection{
		ID:            connID,
		E2NodeID:      nodeID,
		NodeID:        string(nodeID),
		TimeAlive:     now,
		ServiceModels: serviceModels,
		E2Cells:       e2Cells,
	}

}

type E2BaseConnection struct {
	ID            ID
	E2NodeID      topoapi.ID
	NodeID        string
	PlmnID        string
	TimeAlive     time.Time
	ServiceModels map[string]*topoapi.ServiceModelInfo
	E2Cells       []*topoapi.E2Cell
}

func (c *E2BaseConnection) GetTimeAlive() time.Time {
	return c.TimeAlive
}

func (c *E2BaseConnection) GetPlmnID() string {
	return c.PlmnID
}

func (c *E2BaseConnection) GetE2Cells() []*topoapi.E2Cell {
	return c.E2Cells
}

func (c *E2BaseConnection) GetServiceModels() map[string]*topoapi.ServiceModelInfo {
	return c.ServiceModels
}

func (c *E2BaseConnection) GetID() ID {
	return c.ID
}

func (c *E2BaseConnection) GetE2NodeID() topoapi.ID {
	return c.E2NodeID
}
