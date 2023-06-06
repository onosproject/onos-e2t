// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2

import (
	"github.com/onosproject/onos-api/go/onos/topo"
	"github.com/onosproject/onos-e2t/test/utils"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"time"
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
func (s *TestSuite) getE2Pods() map[string]*v1.Pod {

	e2tPodList, err := s.CoreV1().Pods(s.Namespace()).List(s.Context(), metav1.ListOptions{
		LabelSelector: "app=onos,type=e2t",
	})
	s.NoError(err)

	e2tPods := make(map[string]*v1.Pod)
	for i, pod := range e2tPodList.Items {
		e2tPods[pod.Name] = &e2tPodList.Items[i]
	}
	return e2tPods
}

// getExpiration extracts the expiration time for an object from its aspects
func (s *TestSuite) getExpiration(node topo.Object) time.Time {
	lease := topo.Lease{}
	err := node.GetAspect(&lease)
	s.NoError(err)
	return *lease.Expiration
}

// getE2TNodes waits for the specified number of e2t nodes to appear in the topo
func (s *TestSuite) getE2TNodes(count int64) []topo.Object {
	topoSdkClient, err := utils.NewTopoClient()
	s.NoError(err)

	iterations := 10

	for i := 1; i <= iterations; i++ {
		nodes, err := topoSdkClient.E2TNodes(s.Context())
		s.NoError(err)
		if len(nodes) == int(count) {
			return nodes
		}
		time.Sleep(2 * time.Second)
	}
	s.Fail("Nodes never became active")
	return nil
}

// deletePods deletes a pod
func (s *TestSuite) deletePod(pod *v1.Pod) {
	err := s.CoreV1().Pods(s.Namespace()).Delete(s.Context(), pod.Name, metav1.DeleteOptions{})
	s.NoError(err)
}

func (s *TestSuite) waitForE2TDeleted(ch chan topo.Event) {
	for event := range ch {
		if event.Type == topo.EventType_REMOVED {
			return
		}
	}
	s.Fail("REMOVED event not seen")
}

// TestE2TLeaseExpiration checks that when an E2T pod is deleted, topo is updated properly
func (s *TestSuite) TestE2TLeaseExpiration() {
	// check that the E2T pods are all registered properly with topo
	e2tPods := s.getE2Pods()
	s.Equal(int(s.E2TReplicaCount), len(e2tPods))
	nodes := s.getE2TNodes(s.E2TReplicaCount)

	topoSdkClient, err := utils.NewTopoClient()
	s.NoError(err)

	deleteEventChan := make(chan topo.Event)
	err = topoSdkClient.WatchE2TNodes(s.Context(), deleteEventChan)
	s.NoError(err)

	// delete the first e2t node in the list - k8s will make a new pod
	firstNode := nodes[0]
	firstPod := e2tPods[podID(firstNode)]
	s.deletePod(firstPod)

	// wait for the topo object to expire
	expiration := s.getExpiration(firstNode)
	untilExpiration := time.Until(expiration)
	<-time.After(untilExpiration)

	s.waitForE2TDeleted(deleteEventChan)

	// Check that the new pod was properly registered
	e2tPods = s.getE2Pods()
	s.Equal(int(s.E2TReplicaCount), len(e2tPods))
	nodesAfterDelete := s.getE2TNodes(s.E2TReplicaCount)

	// check that the expired node was removed
	for _, node := range nodesAfterDelete {
		if string(node.ID) == nodeID(firstPod) {
			s.Fail("Crashed e2 pod not removed from topo")
		}
	}
}
