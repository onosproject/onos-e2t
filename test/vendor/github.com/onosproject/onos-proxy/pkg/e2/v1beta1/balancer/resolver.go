// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package balancer

import (
	"context"
	"fmt"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/onosproject/onos-api/go/onos/topo"
	"github.com/onosproject/onos-lib-go/pkg/grpc/retry"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/attributes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/serviceconfig"
)

var log = logging.GetLogger()

const ResolverName = "e2"
const topoAddress = "onos-topo:5150"

func init() {
	resolver.Register(&ResolverBuilder{})
}

// ResolverBuilder :
type ResolverBuilder struct{}

// Scheme :
func (b *ResolverBuilder) Scheme() string {
	return ResolverName
}

// Build :
func (b *ResolverBuilder) Build(_ resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	var dialOpts []grpc.DialOption
	if opts.DialCreds != nil {
		dialOpts = append(
			dialOpts,
			grpc.WithTransportCredentials(opts.DialCreds),
		)
	} else {
		dialOpts = append(dialOpts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}
	dialOpts = append(dialOpts, grpc.WithUnaryInterceptor(retry.RetryingUnaryClientInterceptor(retry.WithRetryOn(codes.Unavailable))))
	dialOpts = append(dialOpts, grpc.WithStreamInterceptor(retry.RetryingStreamClientInterceptor(retry.WithRetryOn(codes.Unavailable))))
	dialOpts = append(dialOpts, grpc.WithContextDialer(opts.Dialer))

	topoConn, err := grpc.Dial(topoAddress, dialOpts...)
	if err != nil {
		return nil, err
	}

	serviceConfig := cc.ParseServiceConfig(
		fmt.Sprintf(`{"loadBalancingConfig":[{"%s":{}}]}`, ResolverName),
	)

	log.Infof("Built new resolver")

	resolver := &Resolver{
		clientConn:    cc,
		topoConn:      topoConn,
		serviceConfig: serviceConfig,
		masterships:   make(map[topo.ID]topo.MastershipState),
		controls:      make(map[topo.ID]topo.ID),
		addresses:     make(map[topo.ID]string),
	}
	err = resolver.start()
	if err != nil {
		return nil, err
	}
	return resolver, nil
}

var _ resolver.Builder = (*ResolverBuilder)(nil)

// Resolver :
type Resolver struct {
	clientConn    resolver.ClientConn
	topoConn      *grpc.ClientConn
	serviceConfig *serviceconfig.ParseResult
	masterships   map[topo.ID]topo.MastershipState // E2 node to mastership (controls relation ID)
	controls      map[topo.ID]topo.ID              // controls relation to E2T ID
	addresses     map[topo.ID]string               // E2T ID to address
}

func (r *Resolver) start() error {
	log.Infof("Starting resolver")

	client := topo.NewTopoClient(r.topoConn)
	request := &topo.WatchRequest{}
	stream, err := client.Watch(context.Background(), request)
	if err != nil {
		return err
	}
	go func() {
		for {
			response, err := stream.Recv()
			if err != nil {
				return
			}
			r.handleEvent(response.Event)
		}
	}()
	return nil
}

func (r *Resolver) handleEvent(event topo.Event) {
	object := event.Object
	if entity, ok := object.Obj.(*topo.Object_Entity); ok && entity.Entity.KindID == topo.E2NODE {
		// Track changes in E2 nodes
		switch event.Type {
		case topo.EventType_REMOVED:
			delete(r.masterships, object.ID)
		default:
			var mastership topo.MastershipState
			_ = object.GetAspect(&mastership)
			if mastership.Term > r.masterships[object.ID].Term {
				r.masterships[object.ID] = mastership
			}
		}
		r.updateState()
	} else if entity, ok := object.Obj.(*topo.Object_Entity); ok && entity.Entity.KindID == topo.E2T {
		// Track changes in E2T instances
		switch event.Type {
		case topo.EventType_REMOVED:
			delete(r.addresses, object.ID)
			r.updateState()
		default:
			var info topo.E2TInfo
			_ = object.GetAspect(&info)
			for _, iface := range info.Interfaces {
				if iface.Type == topo.Interface_INTERFACE_E2T {
					address := fmt.Sprintf("%s:%d", iface.IP, iface.Port)
					if r.addresses[object.ID] != address {
						r.addresses[object.ID] = address
						r.updateState()
						break
					}
				}
			}
		}
	} else if relation, ok := object.Obj.(*topo.Object_Relation); ok && relation.Relation.KindID == topo.CONTROLS {
		// Track changes in E2T/E2Node controls relations
		switch event.Type {
		case topo.EventType_REMOVED:
			delete(r.controls, object.ID)
		default:
			r.controls[object.ID] = relation.Relation.SrcEntityID
		}
		r.updateState()
	}
}

func (r *Resolver) updateState() {
	// Produce list of addresses for available E2T instances
	// Annotate each address with a list of nodes for which this instances is presently the master
	e2tE2Nodes := make(map[topo.ID]nodeList)

	// Scan over all nodes and insert their ID into the list of nodes of its master E2T instance
	for nodeID, mastership := range r.masterships {
		if e2tID, ok := r.controls[topo.ID(mastership.NodeId)]; ok {
			e2tE2Nodes[e2tID] = append(e2tE2Nodes[e2tID], string(nodeID))
		}
	}

	// Transpose the map of E2T node IDs into a list of addresses with nodes attribute
	addresses := make([]resolver.Address, 0, len(r.addresses))
	for e2tID, addr := range r.addresses {
		if nodes, ok := e2tE2Nodes[e2tID]; ok {
			addresses = append(addresses, resolver.Address{
				Addr: addr,
				Attributes: attributes.New(
					"nodes",
					nodes,
				),
			})
			log.Debugf("New resolver address: %s => %+v", addr, nodes)
		}
	}

	log.Infof("New resolver addresses: %+v", addresses)

	// Update the resolver state with list of E2T addresses annotated by nodes for which they are masters
	// TODO - this call sometimes returns an error in the 1.41 version of grpc. Need to figure out why
	_ = r.clientConn.UpdateState(resolver.State{
		Addresses:     addresses,
		ServiceConfig: r.serviceConfig,
	})
}

// ResolveNow :
func (r *Resolver) ResolveNow(resolver.ResolveNowOptions) {}

// Close :
func (r *Resolver) Close() {
	if err := r.topoConn.Close(); err != nil {
		log.Error("failed to close conn", err)
	}
}

var _ resolver.Resolver = (*Resolver)(nil)

type nodeList []string

func (l nodeList) Equal(o interface{}) bool {
	if nl, ok := o.(nodeList); ok {
		if len(l) != len(nl) {
			return false
		}
		for i := 0; i < len(l); i++ {
			if l[i] != nl[i] {
				return false
			}
		}
		return true
	}
	return false
}
