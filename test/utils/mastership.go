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
	relations, err := topoSdkClient.GetControlRelations()
	assert.NoError(t, err)

	for _, relationObject := range relations {
		if relationObject.ID == masterRelationID {
			return relationObject.GetRelation()
		}
	}
	return nil
}

func GetE2NodeNonMasterNodes(t *testing.T, e2NodeID topoapi.ID) []topoapi.Interface {
	nonMasters := make([]topoapi.Interface, 0)
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

		e2tIface := topoapi.Interface{}
		e2tInfo := &topoapi.E2TInfo{}
		err := e2tNode.GetAspect(e2tInfo)
		assert.NoError(t, err)
		for _, iface := range e2tInfo.Interfaces {
			if iface.Type == topoapi.Interface_INTERFACE_E2T {
				e2tIface.IP = iface.IP
				e2tIface.Port = iface.Port
				e2tIface.Type = topoapi.Interface_INTERFACE_E2T
				break
			}
		}
		if masterRelation.GetSrcEntityID() == e2tNode.GetID() {
			continue
		} else {
			nonMasters = append(nonMasters, e2tIface)
		}
	}
	t.Logf("List of non master e2t Nodes for e2 node  %s are %+v", e2NodeID, nonMasters)
	return nonMasters
}

func GetE2NodeMaster(t *testing.T, e2NodeID topoapi.ID) topoapi.Interface {
	var master topoapi.Interface
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
		e2tIface := topoapi.Interface{}
		e2tInfo := &topoapi.E2TInfo{}
		err := e2tNode.GetAspect(e2tInfo)
		assert.NoError(t, err)
		for _, iface := range e2tInfo.Interfaces {
			if iface.Type == topoapi.Interface_INTERFACE_E2T {
				e2tIface.IP = iface.IP
				e2tIface.Port = iface.Port
				break
			}
		}
		if masterRelation.GetSrcEntityID() == e2tNode.GetID() {
			master.IP = e2tIface.IP
			master.Port = e2tIface.Port
			master.Type = topoapi.Interface_INTERFACE_E2T
			break
		}
	}
	t.Logf("Master node for e2 Node %s is %+v", e2NodeID, master)
	return master
}
