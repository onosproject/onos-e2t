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
	return mastership,err
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
	mastership := topoapi.MastershipState{}

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

	// wait for e2t to connect to the simulator
	topoConnectionEventChan := make(chan topoapi.Event)
	err = topoSdkClient.WatchE2Connections(ctx, topoConnectionEventChan)
	assert.NoError(t, err)

	// wait for the first mastership update event
	mastership, err = waitForMastershipTerm(t, topoNodeEventChan, 1)
	assert.NoError(t, err)
	assert.Equal(t, uint64(1), mastership.GetTerm())

	// Get the API objects for the e2T node and the E2 node
	e2Node := getE2Node(t)
	assert.NotNil(t, e2Node)
	checkE2tNodeRelation(t, *e2Node)

	// Delete the e2Node
	_, err = nodeClient.DeleteNode(ctx, &modelapi.DeleteNodeRequest{
		GnbID: gnbID,
	})

	// Make a new e2Node with the same gnbID
	createE2Node(t, sim)

	// wait for the second mastership update event
	mastership, err = waitForMastershipTerm(t, topoNodeEventChan, 2)
	assert.NoError(t, err)
	assert.Equal(t, uint64(2), mastership.GetTerm())

	// tear down the simulator
	assert.NoError(t, sim.Uninstall())
}
