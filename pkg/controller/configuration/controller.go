// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package configuration

import (
	"context"
	"encoding/binary"
	"net"
	"time"

	topoapi "github.com/onosproject/onos-api/go/onos/topo"

	"github.com/onosproject/onos-e2t/pkg/controller/utils"

	"github.com/onosproject/onos-lib-go/pkg/errors"

	"github.com/onosproject/onos-e2t/api/e2ap/v2beta1"
	e2apcommondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-commondatatypes"
	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-ies"
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-pdu-contents"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"

	e2server "github.com/onosproject/onos-e2t/pkg/southbound/e2ap/server"
	"github.com/onosproject/onos-e2t/pkg/store/rnib"
	"github.com/onosproject/onos-lib-go/pkg/controller"
	"github.com/onosproject/onos-lib-go/pkg/logging"
)

var log = logging.GetLogger("controller", "configuration")

const (
	defaultTimeout = 30 * time.Second
)

// NewController returns a new E2 connection update controller
func NewController(rnib rnib.Store, mgmtConns e2server.MgmtConnManager, e2apConns e2server.E2APConnManager) *controller.Controller {
	c := controller.NewController("configuration")

	c.Watch(&MgmtConnWatcher{
		mgmtConns: mgmtConns,
	})

	c.Watch(&TopoWatcher{
		mgmtConns: mgmtConns,
		e2apConns: e2apConns,
		rnib:      rnib,
	})

	c.Reconcile(&Reconciler{
		mgmtConns: mgmtConns,
		e2apConns: e2apConns,
		rnib:      rnib,
	})

	return c
}

// Reconciler reconciles configuration of an E2 node
type Reconciler struct {
	mgmtConns e2server.MgmtConnManager
	e2apConns e2server.E2APConnManager
	rnib      rnib.Store
}

func (r *Reconciler) createConnectionUpdateReq(ip string) *e2appducontents.E2ConnectionUpdate {
	connectionAddList := &e2appducontents.E2ConnectionUpdateIes_E2ConnectionUpdateIes44{
		Id:          int32(v2beta1.ProtocolIeIDE2connectionUpdateAdd),
		Criticality: int32(e2apcommondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &e2appducontents.E2ConnectionUpdateList{
			Value: make([]*e2appducontents.E2ConnectionUpdateItemIes, 0),
		},
		Presence: int32(e2apcommondatatypes.Presence_PRESENCE_OPTIONAL),
	}

	testIP := net.ParseIP(ip)

	portBytes := make([]byte, 2)
	port := uint16(36421)
	binary.BigEndian.PutUint16(portBytes, port)
	cai := &e2appducontents.E2ConnectionUpdateItemIes{
		Id:          int32(v2beta1.ProtocolIeIDE2connectionUpdateItem),
		Criticality: int32(e2apcommondatatypes.Criticality_CRITICALITY_IGNORE),
		Value: &e2appducontents.E2ConnectionUpdateItem{
			TnlInformation: &e2apies.Tnlinformation{
				TnlPort: &asn1.BitString{
					Value: portBytes,
					Len:   16,
				},
				TnlAddress: &asn1.BitString{
					Value: testIP.To4(),
					Len:   32,
				},
			},
			TnlUsage: e2apies.Tnlusage_TNLUSAGE_BOTH,
		},
		Presence: int32(e2apcommondatatypes.Presence_PRESENCE_MANDATORY),
	}

	transactionID := &e2appducontents.E2ConnectionUpdateIes_E2ConnectionUpdateIes49{
		Id:          int32(v2beta1.ProtocolIeIDTransactionID),
		Criticality: int32(e2apcommondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &e2apies.TransactionId{
			Value: 3,
		},
		Presence: int32(e2apcommondatatypes.Presence_PRESENCE_MANDATORY),
	}
	connectionAddList.Value.Value = append(connectionAddList.Value.Value, cai)
	// Send the subscription request and await a response
	connectionUpdateRequest := &e2appducontents.E2ConnectionUpdate{
		ProtocolIes: &e2appducontents.E2ConnectionUpdateIes{
			E2ApProtocolIes44: connectionAddList,
			E2ApProtocolIes49: transactionID,
		},
	}

	return connectionUpdateRequest
}

func (r *Reconciler) connectionExist(e2NodeID topoapi.ID, e2tConn *topoapi.Interface, e2NodeConns []topoapi.Interface) bool {
	for _, e2NodeConn := range e2NodeConns {
		if e2NodeConn.IP == e2tConn.IP &&
			e2NodeConn.Port == e2tConn.Port && e2NodeConn.Type == e2tConn.Type {
			log.Debugf("Test Connection %+v already exists  for e2node: %s", e2NodeConn, e2NodeID)
			return true
		}
	}
	log.Debugf("Connection %+v does not exists", e2tConn)
	return false
}

func (r *Reconciler) Reconcile(id controller.ID) (controller.Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	connID := id.Value.(e2server.ConnID)
	log.Infof("Reconciling  configuration using mgmt connection: %s", connID)
	conn, err := r.mgmtConns.Get(ctx, connID)
	if err != nil {
		if errors.IsNotFound(err) {
			log.Warn(err)
			return controller.Result{}, nil
		}
		log.Warnf("Failed to reconcile configuration using management connection %s: %s", connID, err)
		return controller.Result{}, err
	}
	e2NodeID := conn.NodeID

	e2tNodes, err := r.rnib.List(ctx, utils.GetE2TFilter())
	if err != nil {
		return controller.Result{}, err
	}
	if len(e2tNodes) == 0 {
		return controller.Result{
			Requeue: id,
		}, nil
	}
	e2Node, err := r.rnib.Get(ctx, topoapi.ID(e2NodeID))
	if err != nil {
		return controller.Result{}, err
	}

	e2NodeConfig := &topoapi.E2NodeConfig{}
	_ = e2Node.GetAspect(e2NodeConfig)

	e2NodeConns := e2NodeConfig.Connections
	for _, e2tNode := range e2tNodes {
		e2tNodeInfo := &topoapi.E2TInfo{}
		err := e2tNode.GetAspect(e2tNodeInfo)
		if err != nil {
			log.Warnf("Failed to reconcile configuration using management connection %s: %s", connID, err)
			return controller.Result{}, err
		}
		for _, e2tConn := range e2tNodeInfo.GetInterfaces() {
			if e2tConn.Type == topoapi.Interface_INTERFACE_E2AP200 && !r.connectionExist(e2Node.ID, e2tConn, e2NodeConns) {
				connectionUpdateReq := r.createConnectionUpdateReq(e2tConn.IP)
				log.Debugf("Sending connection update %+v", connectionUpdateReq)
				connUpdateAck, connUpdateFailure, err := conn.E2ConnectionUpdate(ctx, connectionUpdateReq)
				if err != nil {
					log.Warnf("Failed to reconcile configuration using management connection %s: %s", connID, err)
					return controller.Result{}, err
				}

				if connUpdateAck != nil {
					log.Infof("Received connection update ack:%+v", connUpdateAck)
					e2NodeConfig.Connections = append(e2NodeConfig.Connections, *e2tConn)
					err := e2Node.SetAspect(e2NodeConfig)
					if err != nil {
						log.Warnf("Failed to reconcile configuration using management connection %s: %s", connID, err)
						return controller.Result{}, err
					}

					err = r.rnib.Update(ctx, e2Node)
					if err != nil {
						if !errors.IsNotFound(err) {
							log.Warnf("Failed to reconcile configuration using management connection %s: %s", connID, err)
							return controller.Result{}, err
						}
						return controller.Result{}, nil
					}
					return controller.Result{}, nil

				}
				if connUpdateFailure != nil {
					log.Infof("Received connection update failure: %+v", connUpdateFailure)
				}
			}
		}
	}

	return controller.Result{}, nil
}
