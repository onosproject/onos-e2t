// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2

import (
	"github.com/onosproject/helmit/pkg/kubernetes"
	"github.com/onosproject/helmit/pkg/kubernetes/core/v1"
	"github.com/onosproject/onos-api/go/onos/topo"
	"github.com/onosproject/onos-e2t/test/e2utils"
	"testing"
	"time"

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

// getE2Pods returns a map of E2T pod names to pods
func getE2Pods(t *testing.T, s *TestSuite) map[string]*v1.Pod {
	sdranClient, err := kubernetes.NewForRelease(s.release)
	assert.NoError(t, err)

	ctx, cancel := e2utils.GetCtx()
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
	lease := topo.Lease{}
	err := node.GetAspect(&lease)
	assert.NoError(t, err)
	return *lease.Expiration
}

// getE2TNodes waits for the specified number of e2t nodes to appear in the topo
func getE2TNodes(t *testing.T, count int64) []topo.Object {
	topoSdkClient, err := utils.NewTopoClient()
	assert.NoError(t, err)

	iterations := 10

	for i := 1; i <= iterations; i++ {
		ctx, cancel := e2utils.GetCtx()
		nodes, err := topoSdkClient.E2TNodes(ctx)
		assert.NoError(t, err)
		cancel()
		if len(nodes) == int(count) {
			return nodes
		}
		time.Sleep(2 * time.Second)
	}
	assert.Fail(t, "Nodes never became active")
	return nil
}

// deletePods deletes a pod
func deletePod(t *testing.T, pod *v1.Pod) {
	ctx, cancel := e2utils.GetCtx()
	err := pod.Delete(ctx)
	assert.NoError(t, err)
	cancel()
}

func waitForE2TDeleted(t *testing.T, ch chan topo.Event) {
	for event := range ch {
		if event.Type == topo.EventType_REMOVED {
			return
		}
	}
	assert.Fail(t, "REMOVED event not seen")
}

// TestE2TLeaseExpiration checks that when an E2T pod is deleted, topo is updated properly
func (s *TestSuite) TestE2TLeaseExpiration(t *testing.T) {
	// check that the E2T pods are all registered properly with topo
	e2tPods := getE2Pods(t, s)
	assert.Equal(t, int(s.E2TReplicaCount), len(e2tPods))
	nodes := getE2TNodes(t, s.E2TReplicaCount)

	topoSdkClient, err := utils.NewTopoClient()
	assert.NoError(t, err)

	deleteEventChan := make(chan topo.Event)
	ctx, cancel := e2utils.GetCtx()
	defer cancel()
	err = topoSdkClient.WatchE2TNodes(ctx, deleteEventChan)
	assert.NoError(t, err)

	// delete the first e2t node in the list - k8s will make a new pod
	firstNode := nodes[0]
	firstPod := e2tPods[podID(firstNode)]
	deletePod(t, firstPod)

	// wait for the topo object to expire
	expiration := getExpiration(t, firstNode)
	untilExpiration := time.Until(expiration)
	<-time.After(untilExpiration)

	waitForE2TDeleted(t, deleteEventChan)

	// Check that the new pod was properly registered
	e2tPods = getE2Pods(t, s)
	assert.Equal(t, int(s.E2TReplicaCount), len(e2tPods))
	nodesAfterDelete := getE2TNodes(t, s.E2TReplicaCount)

	// check that the expired node was removed
	for _, node := range nodesAfterDelete {
		if string(node.ID) == nodeID(firstPod) {
			assert.Fail(t, "Crashed e2 pod not removed from topo")
		}
	}
}
