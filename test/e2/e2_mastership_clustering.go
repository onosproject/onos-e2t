// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package e2

import (
	"fmt"
	"github.com/onosproject/helmit/pkg/helm"
	modelapi "github.com/onosproject/onos-api/go/onos/ransim/model"
	ransimtypes "github.com/onosproject/onos-api/go/onos/ransim/types"
	topoapi "github.com/onosproject/onos-api/go/onos/topo"
	"github.com/onosproject/onos-e2t/test/utils"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var gnbID = ransimtypes.GnbID(166000)
var gnbIDSuffix = fmt.Sprintf("/%x", gnbID)

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

	for _, node := range nodes {
		if nodeMatchesGnbID(node) {
			return &node
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

func nodeMatchesGnbID(node topoapi.Object) bool {
	return strings.HasSuffix(string(node.ID), gnbIDSuffix)
}

func waitForMastershipTerm(topoNodeEventChan chan topoapi.Event, term uint64) (topoapi.MastershipState, error) {
	mastership := topoapi.MastershipState{}

	var err error
	for {
		addedEvent, err := utils.GetUpdatedEvent(topoNodeEventChan)
		if err != nil {
			break
		}
		if !nodeMatchesGnbID(addedEvent.Object) {
			continue
		}
		err = addedEvent.Object.GetAspect(&mastership)
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
	ctx, cancel := getCtx()
	defer cancel()
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
	mastership, err := waitForMastershipTerm(topoNodeEventChan, firstTerm)
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
	mastership, err = waitForMastershipTerm(topoNodeEventChan, secondTerm)
	assert.NoError(t, err)
	assert.Equal(t, secondTerm, mastership.GetTerm())

	// tear down the simulator
	assert.NoError(t, sim.Uninstall())
}
