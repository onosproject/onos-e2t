// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

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

// checkNodes checks that the IDs of the E2T pods match the topo object IDs
func checkNodes(t *testing.T, e2tPods map[string]*v1.Pod, events map[string]topo.Event) []topo.Object {
	topoSdkClient, err := utils.NewTopoClient()
	assert.NoError(t, err)

	ctx, cancel := e2utils.GetCtx()
	nodes, err := topoSdkClient.E2TNodes(ctx)
	assert.NoError(t, err)

	for _, node := range nodes {
		// check that the node found in topo is also in k8s
		idAsPod := podID(node)
		assert.Equal(t, idAsPod, e2tPods[idAsPod].Name)

		// If an event map was provided, check that the ID in the event map
		// matches the k8s pod ID
		if events != nil {
			_, ok := events[idAsPod]
			assert.True(t, ok)
		}
	}
	cancel()
	return nodes
}

// deletePods deletes a pod
func deletePod(t *testing.T, pod *v1.Pod) {
	ctx, cancel := e2utils.GetCtx()
	err := pod.Delete(ctx)
	assert.NoError(t, err)
	cancel()
}

func readTopoAddedEvents(ch chan topo.Event, expectedValue int) map[string]topo.Event {
	eventsMap := make(map[string]topo.Event)
	count := 0
	for event := range ch {
		if event.Type == topo.EventType_ADDED {
			count++
			eventsMap[podID(event.Object)] = event
			if count == expectedValue {
				break
			}
		}
	}
	return eventsMap
}

// waitForE2TNodes waits until 2 E2T nodes have registered with topo
func waitForE2TNodes(t *testing.T) map[string]topo.Event {
	topoSdkClient, err := utils.NewTopoClient()
	assert.NoError(t, err)

	topoEventChan := make(chan topo.Event)
	ctx, cancel := e2utils.GetCtx()
	err = topoSdkClient.WatchE2TNodes(ctx, topoEventChan)
	assert.NoError(t, err)
	events := readTopoAddedEvents(topoEventChan, utils.E2TReplicaCount)
	cancel()
	return events
}

func waitForE2TDeleted(t *testing.T, ch chan topo.Event) {
	time.Sleep(5 * time.Second)
	// TODO figure out why these events are not always delivered properly
	//for event := range ch {
	//	fmt.Fprintf(os.Stderr, "waiting for delete, saw %v\n", event.Type)
	//	if event.Type == topo.EventType_REMOVED {
	//		return
	//	}
	//}
	//assert.Fail(t, "REMOVED event not seen")
}

// TestE2TLeaseExpiration checks that when an E2T pod is deleted, topo is updated properly
func (s *TestSuite) TestE2TLeaseExpiration(t *testing.T) {
	eventNodes := waitForE2TNodes(t)

	// check that the E2T pods are all registered properly with topo
	e2tPods := getE2Pods(t, s)
	nodes := checkNodes(t, e2tPods, eventNodes)
	assert.Equal(t, 2, len(nodes))

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
	nodesAfterDelete := checkNodes(t, e2tPods, nil)

	// check that the expired node was removed
	for _, node := range nodesAfterDelete {
		if string(node.ID) == nodeID(firstPod) {
			assert.Fail(t, "Crashed e2 pod not removed from topo")
		}
	}

	// check that there are the correct number of registrations
	assert.Equal(t, len(nodes), len(nodesAfterDelete))
}
