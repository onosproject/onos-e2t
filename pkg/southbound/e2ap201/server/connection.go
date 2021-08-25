// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package server

import (
	"context"

	"github.com/onosproject/onos-lib-go/pkg/errors"

	"github.com/onosproject/onos-e2t/pkg/southbound/e2conn"

	"sync"
	"time"

	e2smtypes "github.com/onosproject/onos-api/go/onos/e2t/e2sm"
	topoapi "github.com/onosproject/onos-api/go/onos/topo"
	subscriptionv2beta1 "github.com/onosproject/onos-e2t/pkg/broker/subscription/v2beta1"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap201/types"

	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-pdu-contents"
	e2 "github.com/onosproject/onos-e2t/pkg/protocols/e2ap/v201"
)

func NewE2Conn(nodeID topoapi.ID, plmnID string, nodeIdentity *types.E2NodeIdentity, conn e2.ServerConn,
	streamsv2beta1 subscriptionv2beta1.Broker,
	serviceModels map[string]*topoapi.ServiceModelInfo, ranFunctions map[e2smtypes.OID]RANFunction,
	e2Cells []*topoapi.E2Cell, now time.Time) *E2Conn {

	baseConn := e2conn.NewE2BaseConn(nodeID, plmnID, serviceModels, e2Cells, now)
	return &E2Conn{
		ServerConn:       conn,
		E2BaseConnection: baseConn,
		NodeType:         nodeIdentity.NodeType,
		RANFunctions:     ranFunctions,
		Streamsv2beta1:   streamsv2beta1,
	}
}

// RANFunction RAN function information
type RANFunction struct {
	OID e2smtypes.OID
	ID  types.RanFunctionID
	// protobuf encoded description
	Description []byte
}

type E2Conn struct {
	*e2conn.E2BaseConnection
	e2.ServerConn
	NodeType       types.E2NodeType
	RANFunctions   map[e2smtypes.OID]RANFunction
	Streamsv2beta1 subscriptionv2beta1.Broker
	mu             sync.RWMutex
}

func (c *E2Conn) GetTimeAlive() time.Time {
	return c.E2BaseConnection.GetTimeAlive()
}

func (c *E2Conn) GetPlmnID() string {
	return c.E2BaseConnection.GetPlmnID()
}

func (c *E2Conn) GetE2Cells() []*topoapi.E2Cell {
	return c.E2BaseConnection.GetE2Cells()
}

func (c *E2Conn) GetE2NodeID() topoapi.ID {
	return c.E2BaseConnection.GetE2NodeID()
}

func (c *E2Conn) GetID() e2conn.ID {
	return c.E2BaseConnection.GetID()
}

func (c *E2Conn) GetRANFunctions() []RANFunction {
	c.mu.RLock()
	defer c.mu.RUnlock()
	ranFunctions := make([]RANFunction, 0, len(c.RANFunctions))
	for _, ranFunction := range c.RANFunctions {
		ranFunctions = append(ranFunctions, ranFunction)
	}
	return ranFunctions
}

func (c *E2Conn) GetRANFunction(oid e2smtypes.OID) (RANFunction, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	ranFunction, ok := c.RANFunctions[oid]
	return ranFunction, ok
}

func (c *E2Conn) ricIndication(ctx context.Context, request *e2appducontents.Ricindication) error {
	log.Debugf("Received RICIndication %+v", request)
	streamID := subscriptionv2beta1.StreamID(request.ProtocolIes.E2ApProtocolIes29.Value.RicRequestorId)
	stream, ok := c.Streamsv2beta1.GetWriter(streamID)
	if !ok {
		return errors.NewNotFound("cannot find stream with ID %s", streamID)
	}
	return stream.Send(request)
}

var _ e2conn.E2BaseConn = &E2Conn{}
