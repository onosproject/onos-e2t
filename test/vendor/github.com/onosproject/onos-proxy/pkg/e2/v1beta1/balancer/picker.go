// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package balancer

import (
	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/balancer/base"
	"google.golang.org/grpc/metadata"
)

const e2NodeIDHeader = "e2-node-id"

func init() {
	balancer.Register(base.NewBalancerBuilder(ResolverName, &PickerBuilder{}, base.Config{}))
}

// PickerBuilder :
type PickerBuilder struct{}

// Build :
func (p *PickerBuilder) Build(info base.PickerBuildInfo) balancer.Picker {
	masters := make(map[string]balancer.SubConn)

	for sc, scInfo := range info.ReadySCs {
		nodes := scInfo.Address.Attributes.Value("nodes").(nodeList)
		for _, node := range nodes {
			log.Debugf("E2 node %s is mastered by E2T %s; conn=%+v", node, scInfo.Address.Addr, sc)
			masters[node] = sc
		}
	}
	log.Infof("Built new picker for E2T instances: %+v", masters)
	return &Picker{
		masters: masters,
	}
}

var _ base.PickerBuilder = (*PickerBuilder)(nil)

// Picker :
type Picker struct {
	masters map[string]balancer.SubConn // NodeID string to connection mapping
}

// Pick :
func (p *Picker) Pick(info balancer.PickInfo) (balancer.PickResult, error) {
	var result balancer.PickResult
	if md, ok := metadata.FromOutgoingContext(info.Ctx); ok {
		ids := md.Get(e2NodeIDHeader)
		if len(ids) > 0 {
			if subConn, ok := p.masters[ids[0]]; ok {
				log.Debugf("Picked subconn for %s: %+v", ids[0], subConn)
				result.SubConn = subConn
				return result, nil
			}
		}
	}
	log.Warn("No subconn available")
	return result, balancer.ErrNoSubConnAvailable
}

var _ balancer.Picker = (*Picker)(nil)
