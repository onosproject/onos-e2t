// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package e2

import (
	"context"
	"github.com/onosproject/helmit/pkg/kubernetes"
	v1 "github.com/onosproject/helmit/pkg/kubernetes/core/v1"
	"github.com/onosproject/onos-api/go/onos/topo"
	"sync"
	"testing"
	"time"

	"encoding/json"
	"github.com/onosproject/onos-e2t/test/utils"
	"github.com/stretchr/testify/assert"
)

func podID(node topo.Object) string {
	return string(node.ID[3:])
}

func nodeID(pod *v1.Pod) string {
	return "e2:" + pod.Name
}

func getCtx() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 1*time.Minute)
}

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

func deletePod(t *testing.T, pod *v1.Pod) {
	ctx, cancel := getCtx()
	err := pod.Delete(ctx)
	assert.NoError(t, err)
	cancel()
}

func (s *TestSuite) TestE2TClustering(t *testing.T) {
	sim := utils.CreateRanSimulatorWithNameOrDie(t, s.c, "e2-clustering")
	assert.NotNil(t, sim)

	e2tPods := getE2Pods(t, s)

	nodes := checkNodes(t, e2tPods)
	firstNode := nodes[0]
	firstPod := e2tPods[podID(firstNode)]
	expiration := getExpiration(t, firstNode)

	deletePod(t, firstPod)

	untilExpiration := time.Until(expiration)

	wg := sync.WaitGroup{}
	wg.Add(1)
	time.AfterFunc(untilExpiration, func() { wg.Done() })
	wg.Wait()

	e2tPods = getE2Pods(t, s)
	nodes = checkNodes(t, e2tPods)
	for _, node := range nodes {
		if string(node.ID) == nodeID(firstPod) {
			assert.Fail(t, "Crashed e2 pod not removed from topo")
		}
	}

	assert.NoError(t, sim.Uninstall())
}
