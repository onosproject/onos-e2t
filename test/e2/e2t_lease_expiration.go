// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package e2

import (
	"context"
	"github.com/onosproject/helmit/pkg/kubernetes"
	"github.com/onosproject/helmit/pkg/kubernetes/core/v1"
	"github.com/onosproject/onos-api/go/onos/topo"
	"testing"
	"time"

	"encoding/json"
	"github.com/onosproject/onos-e2t/test/utils"
	"github.com/stretchr/testify/assert"
)

// podID converts an E2T node name ID into a pod name
func podID(node topo.Object) string {
	return string(node.ID[3:])
}

// nodeID converts a pod into an E2t node ID
func nodeID(pod *v1.Pod) string {
	return "e2:" + pod.Name
}

// getCtx returns a context to use in gRPC calls
func getCtx() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 1*time.Minute)
}

// getE2Pods returns a map of E2T pod names to pods
func getE2Pods(t *testing.T, s *TestSuite) map[string]*v1.Pod {
	sdranClient, err := kubernetes.NewForRelease(s.release)
	assert.NoError(t, err)

	ctx, cancel := getCtx()
	e2tDeployments, err := sdranClient.AppsV1().
		Deployments().
		Get(ctx, "onos-e2t")
	assert.NoError(t, err)
	e2tPodList, err := e2tDeployments.Pods().List(ctx)
	assert.NoError(t, err)
	e2tPods := make(map[string]*v1.Pod)
	for _, pod := range e2tPodList {
		e2tPods[pod.Name] = pod
	}
	cancel()
	return e2tPods
}

// getExpiration extracts the expiration time for an object from its aspects
func getExpiration(t *testing.T, node topo.Object) time.Time {
	var expiration time.Time
	for _, aspect := range node.Aspects {
		if aspect.TypeUrl == "onos.topo.Lease" {
			var jsonData map[string]interface{}
			err := json.Unmarshal(aspect.Value, &jsonData)
			assert.NoError(t, err)
			expirationString := (jsonData["expiration"]).(string)
			expiration, err = time.Parse(time.RFC3339, expirationString)
			assert.NoError(t, err)
		}
	}
	return expiration
}

// checkNodes checks that the IDs of the E2T pods match the topo object IDs
func checkNodes(t *testing.T, e2tPods map[string]*v1.Pod) []topo.Object {
	topoSdkClient, err := utils.NewTopoClient()
	assert.NoError(t, err)

	ctx, cancel := getCtx()
	nodes, err := topoSdkClient.E2TNodes(ctx)
	assert.NoError(t, err)

	for _, node := range nodes {
		idAsPod := podID(node)
		assert.Equal(t, idAsPod, e2tPods[idAsPod].Name)
	}
	cancel()
	return nodes
}

// deletePods deletes a pod
func deletePod(t *testing.T, pod *v1.Pod) {
	ctx, cancel := getCtx()
	err := pod.Delete(ctx)
	assert.NoError(t, err)
	cancel()
}

// TestE2TLeaseExpiration checks that when an E2T pod is deleted, topo is updated properly
func (s *TestSuite) TestE2TLeaseExpiration(t *testing.T) {
	// create a simulator
	sim := utils.CreateRanSimulatorWithNameOrDie(t, s.c, "e2-clustering")
	assert.NotNil(t, sim)

	// first check that the E2T pods are all registered properly with topo
	e2tPods := getE2Pods(t, s)
	nodes := checkNodes(t, e2tPods)

	// delete the first e2t node in the list - k8s will make a new pod
	firstNode := nodes[0]
	firstPod := e2tPods[podID(firstNode)]
	deletePod(t, firstPod)

	// wait for the topo object to expire
	expiration := getExpiration(t, firstNode)
	untilExpiration := time.Until(expiration)

	time.AfterFunc(untilExpiration, func() {

		// Check that the new pod was properly registered
		e2tPods = getE2Pods(t, s)
		nodesAfterDelete := checkNodes(t, e2tPods)

		// check that the expired node was removed
		for _, node := range nodesAfterDelete {
			if string(node.ID) == nodeID(firstPod) {
				assert.Fail(t, "Crashed e2 pod not removed from topo")
			}
		}

		// check that there are the correct number of registrations
		assert.Equal(t, len(nodes), len(nodesAfterDelete))

		// clean up the simulator
		assert.NoError(t, sim.Uninstall())
	})
}
