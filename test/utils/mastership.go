// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package utils

import (
	"context"
	topoapi "github.com/onosproject/onos-api/go/onos/topo"
	"github.com/stretchr/testify/assert"
	"testing"
)

func findMasterRelation(t *testing.T, e2Node topoapi.Object, nodeID topoapi.ID) *topoapi.Relation {
	topoSdkClient, err := NewTopoClient()
	assert.NoError(t, err)
	relations, err := topoSdkClient.GetControlRelationsForTarget()
	assert.NoError(t, err)

	// TODO - replace this with a filter when one is available
	var result *topoapi.Relation
	for _, relationObject := range relations {
		relation := relationObject.GetRelation()
		if relation.SrcEntityID == e2Node.ID &&
			relation.TgtEntityID == nodeID {
			result = relation
		}
	}
	return result
}

type IPAndPort struct {
	IP   string
	Port uint32
}

func GetE2Masters(t *testing.T, e2NodeID topoapi.ID) (IPAndPort, []IPAndPort) {
	var master IPAndPort
	nonMasters := make([]IPAndPort, 0)
	topoClient, err := NewTopoClient()
	assert.NoError(t, err)

	e2tNodes, err := topoClient.E2TNodes(context.Background())
	assert.NoError(t, err)

	for _, node := range e2tNodes {
		e2tIP := ""
		e2tPort := uint32(0)
		e2tInfo := &topoapi.E2TInfo{}
		err := node.GetAspect(e2tInfo)
		assert.NoError(t, err)
		for _, iface := range e2tInfo.Interfaces {
			if iface.Type == topoapi.Interface_INTERFACE_E2T {
				e2tIP = iface.IP
				e2tPort = iface.Port
				break
			}
		}
		rel := findMasterRelation(t, node, e2NodeID)
		if rel != nil {
			master.IP = e2tIP
			master.Port = e2tPort
		} else {
			nonMasters = append(nonMasters, IPAndPort{IP: e2tIP, Port: e2tPort})
		}
	}
	return master, nonMasters
}
