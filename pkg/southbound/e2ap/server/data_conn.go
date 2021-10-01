// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package server

import (
	"context"

	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"

	"github.com/onosproject/onos-e2t/pkg/store/rnib"

	"github.com/onosproject/onos-lib-go/pkg/errors"

	"github.com/google/uuid"
	e2smtypes "github.com/onosproject/onos-api/go/onos/e2t/e2sm"
	topoapi "github.com/onosproject/onos-api/go/onos/topo"
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-pdu-contents"
	subscriptionv1beta1 "github.com/onosproject/onos-e2t/pkg/broker/subscription/v1beta1"
	e2 "github.com/onosproject/onos-e2t/pkg/protocols/e2ap"
	"github.com/onosproject/onos-lib-go/pkg/uri"
)

// NewE2APConn creates a new E2AP connection
func NewE2APConn(nodeID topoapi.ID, conn e2.ServerConn, streamsv1beta1 subscriptionv1beta1.Broker, rnib rnib.Store) *E2APConn {
	connID := ConnID(uri.NewURI(
		uri.WithScheme("uuid"),
		uri.WithOpaque(uuid.New().String())).String())

	return &E2APConn{
		ServerConn:     conn,
		ID:             connID,
		E2NodeID:       nodeID,
		streamsv1beta1: streamsv1beta1,
		rnib:           rnib,
	}

}

// E2APConn e2ap connection
type E2APConn struct {
	e2.ServerConn
	ID             ConnID
	E2NodeID       topoapi.ID
	streamsv1beta1 subscriptionv1beta1.Broker
	rnib           rnib.Store
}

func (c *E2APConn) ricIndication(ctx context.Context, request *e2appducontents.Ricindication) error {
	log.Debugf("Received RICIndication %+v", request)
	streamID := subscriptionv1beta1.StreamID(request.ProtocolIes.E2ApProtocolIes29.Value.RicRequestorId)
	stream, ok := c.streamsv1beta1.GetWriter(streamID)
	if !ok {
		return errors.NewNotFound("stream %s not found", stream.ID())
	}
	return stream.Send(request)
}

func (c *E2APConn) GetRANFunctionID(ctx context.Context, oid e2smtypes.OID) (types.RanFunctionID, bool) {
	object, err := c.rnib.Get(ctx, c.E2NodeID)
	if err != nil {
		log.Warn(err)
		return 0, false
	}
	e2Node := &topoapi.E2Node{}
	err = object.GetAspect(e2Node)
	if err != nil {
		log.Warn(err)
		return 0, false
	}
	ranFunctionIDs := e2Node.ServiceModels[string(oid)].RanFunctionIDs
	// TODO each service model can have multiple RAN functions
	// 	for now it each service model has one supported RAN function so it returns the first ran function ID
	// 	in the list
	return types.RanFunctionID(ranFunctionIDs[0]), true
}
