// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package e2

import (
	"context"
	"fmt"
	"github.com/onosproject/helmit/pkg/helm"
	modelapi "github.com/onosproject/onos-api/go/onos/ransim/model"
	ransimtypes "github.com/onosproject/onos-api/go/onos/ransim/types"
	topoapi "github.com/onosproject/onos-api/go/onos/topo"
	"github.com/onosproject/onos-e2t/test/utils"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var gnbID = ransimtypes.GnbID(166000)

// createE2Node Creates an E2 node in the simulator
func createE2Node(t *testing.T, sim *helm.HelmRelease) {
	nodeClient := utils.GetRansimNodeClient(t, sim)
	assert.NotNil(t, nodeClient)
	createNodeRequest := &modelapi.CreateNodeRequest{
		Node: &ransimtypes.Node{
			GnbID:         gnbID,
			ServiceModels: serviceModels,
			Controllers:   controllers,
			CellNCGIs:     []ransimtypes.NCGI{},
		},
	}
	ctx, cancel := getCtx()
	createNodeResponse, err := nodeClient.CreateNode(ctx, createNodeRequest)
	assert.NoError(t, err)
	assert.NotNil(t, createNodeResponse)
	cancel()
}

func getE2Node(t *testing.T) *topoapi.Object {
	topoSdkClient, err := utils.NewTopoClient()
	assert.NoError(t, err)
	ctx, cancel := getCtx()
	nodes, err := topoSdkClient.E2Nodes(ctx)
	assert.NoError(t, err)
	cancel()

	for i, node := range nodes {
		e2Node := topoapi.E2Node{}
		err := node.GetAspect(&e2Node)
		if err != nil {
			continue
		}
		if e2Node.ServiceModels["1.3.6.1.4.1.53148.1.2.2.100"] != nil {
			m, err := getMastershipAspect(node)
			if err == nil {
				fmt.Fprintf(os.Stderr, "Node %d matches : %v\n", i, m)
				return &node
			}
		}
	}

	return nil
}

func checkE2tNodeRelation(t *testing.T, e2Node topoapi.Object) *topoapi.Relation {
	topoSdkClient, err := utils.NewTopoClient()
	assert.NoError(t, err)
	relations, err := topoSdkClient.GetControlRelationsForTarget()
	assert.NoError(t, err)

	// TODO - replace this with a filter when one is available
	var result *topoapi.Relation
	for _, relationObject := range relations {
		relation := relationObject.GetRelation()
		if relation.TgtEntityID == e2Node.ID {
			result = relation
		}
	}
	assert.NotNil(t, result, "No relation found for E2 Node")
	return result
}

func getMastershipAspect(e2Node topoapi.Object) (topoapi.MastershipState, error) {
	mastership := topoapi.MastershipState{}
	err := e2Node.GetAspect(&mastership)
	return mastership, err
}

func waitForMastershipTerm(t *testing.T, topoNodeEventChan chan topoapi.Event, term uint64) (topoapi.MastershipState, error) {
	mastership := topoapi.MastershipState{}
	var err error
	for {
		fmt.Fprintf(os.Stderr, "waiting for term %d\n", term)
		addedEvent, err := utils.GetUpdatedEvent(topoNodeEventChan)
		assert.NoError(t, err)
		if err != nil {
			break
		}
		err = addedEvent.Object.GetAspect(&mastership)
		if err == nil {
			fmt.Fprintf(os.Stderr, "no error. want %d found %d\n", term, mastership.Term)
		} else {
			fmt.Fprintf(os.Stderr, "error %v\n", err)
		}
		if err == nil && mastership.GetTerm() == term {
			break
		}
	}
	return mastership, err
}

// TestE2TMastershipClustering checks mastership in a clustered environment
func (s *TestSuite) TestE2TMastershipClustering(t *testing.T) {
	const (
		firstTerm  = uint64(1)
		secondTerm = uint64(2)
	)
	ctx := context.Background()
	topoSdkClient, err := utils.NewTopoClient()
	assert.NoError(t, err)
	topoNodeEventChan := make(chan topoapi.Event)
	err = topoSdkClient.WatchE2Nodes(ctx, topoNodeEventChan)
	assert.NoError(t, err)

	// create a simulator
	sim := utils.CreateRanSimulatorWithNameOrDie(t, s.c, "mastership-clustering")
	assert.NotNil(t, sim)

	// Create an E2 node in the simulator
	nodeClient := utils.GetRansimNodeClient(t, sim)
	assert.NotNil(t, nodeClient)
	createE2Node(t, sim)

	// wait for the first mastership update event
	mastership, err := waitForMastershipTerm(t, topoNodeEventChan, firstTerm)
	assert.NoError(t, err)
	assert.Equal(t, firstTerm, mastership.GetTerm())

	// Check that the e2t control relation is correct
	e2Node := getE2Node(t)
	assert.NotNil(t, e2Node)
	checkE2tNodeRelation(t, *e2Node)

	// Delete the e2Node
	_, err = nodeClient.DeleteNode(ctx, &modelapi.DeleteNodeRequest{
		GnbID: gnbID,
	})
	assert.NoError(t, err)

	// Make a new e2Node with the same gnbID
	createE2Node(t, sim)

	// wait for the second mastership update event - term should be 2
	mastership, err = waitForMastershipTerm(t, topoNodeEventChan, secondTerm)
	assert.NoError(t, err)
	assert.Equal(t, secondTerm, mastership.GetTerm())

	// tear down the simulator
	assert.NoError(t, sim.Uninstall())
}
