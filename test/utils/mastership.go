// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package utils

import (
	"context"
	"testing"

	topoapi "github.com/onosproject/onos-api/go/onos/topo"
	"github.com/stretchr/testify/assert"
)

func getMasterRelation(t *testing.T, masterRelationID topoapi.ID) *topoapi.Relation {
	topoSdkClient, err := NewTopoClient()
	assert.NoError(t, err)
	relations, err := topoSdkClient.GetControlRelationsForTarget()
	assert.NoError(t, err)

	for _, relationObject := range relations {
		if relationObject.ID == masterRelationID {
			relation := relationObject.GetRelation()
			return relation
		}
	}
	return nil
}

type IPAndPort struct {
	IP   string
	Port uint32
}

func GetE2NodeNonMasterNodes(t *testing.T, e2NodeID topoapi.ID) []IPAndPort {
	nonMasters := make([]IPAndPort, 0)
	topoClient, err := NewTopoClient()
	assert.NoError(t, err)

	e2tNodes, err := topoClient.E2TNodes(context.Background())
	assert.NoError(t, err)

	// Gets mastership state aspect for an E2 node
	e2NodeMastershipState, err := topoClient.GetE2NodeMastershipState(context.Background(), e2NodeID)
	assert.NoError(t, err)
	assert.NotNil(t, e2NodeMastershipState)
	// find a control relation based on mastership state node ID (i.e. control relation ID)
	masterRelation := getMasterRelation(t, topoapi.ID(e2NodeMastershipState.GetNodeId()))
	assert.NotNil(t, masterRelation)

	for _, e2tNode := range e2tNodes {
		e2tIP := ""
		e2tPort := uint32(0)
		e2tInfo := &topoapi.E2TInfo{}
		err := e2tNode.GetAspect(e2tInfo)
		assert.NoError(t, err)
		for _, iface := range e2tInfo.Interfaces {
			if iface.Type == topoapi.Interface_INTERFACE_E2T {
				e2tIP = iface.IP
				e2tPort = iface.Port
				break
			}
		}
		if masterRelation.GetSrcEntityID() == e2tNode.GetID() {
			continue
		} else {
			nonMasters = append(nonMasters, IPAndPort{IP: e2tIP, Port: e2tPort})
		}
	}
	t.Logf("List of non master e2t Nodes for e2 node  %s are %+v", e2NodeID, nonMasters)
	return nonMasters
}

func GetE2NodeMaster(t *testing.T, e2NodeID topoapi.ID) IPAndPort {
	var master IPAndPort
	topoClient, err := NewTopoClient()
	assert.NoError(t, err)

	e2tNodes, err := topoClient.E2TNodes(context.Background())
	assert.NoError(t, err)

	e2NodeMastershipState, err := topoClient.GetE2NodeMastershipState(context.Background(), e2NodeID)
	assert.NoError(t, err)
	assert.NotNil(t, e2NodeMastershipState)
	masterRelation := getMasterRelation(t, topoapi.ID(e2NodeMastershipState.GetNodeId()))
	assert.NotNil(t, masterRelation)

	for _, e2tNode := range e2tNodes {
		e2tIP := ""
		e2tPort := uint32(0)
		e2tInfo := &topoapi.E2TInfo{}
		err := e2tNode.GetAspect(e2tInfo)
		assert.NoError(t, err)
		for _, iface := range e2tInfo.Interfaces {
			if iface.Type == topoapi.Interface_INTERFACE_E2T {
				e2tIP = iface.IP
				e2tPort = iface.Port
				break
			}
		}
		if masterRelation.GetSrcEntityID() == e2tNode.GetID() {
			master.IP = e2tIP
			master.Port = e2tPort
			break
		}
	}
	t.Logf("Master node for e2 Node %s is %+v", e2NodeID, master)
	return master
}
