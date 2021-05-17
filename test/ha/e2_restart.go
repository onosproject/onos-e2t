// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package ha

import (
	"context"
	"github.com/onosproject/helmit/pkg/helm"
	"github.com/onosproject/helmit/pkg/kubernetes"
	"github.com/onosproject/helmit/pkg/kubernetes/core/v1"
	"github.com/onosproject/onos-api/go/onos/e2t/e2"
	"github.com/onosproject/onos-e2t/test/utils"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"
)

const (
	onosComponentName = "sd-ran"
)

// GetPodListOrFail gets the list of pods active in the onos-config release. The test is failed if getting the list returns
// an error.
func GetPodListOrFail(t *testing.T) []*v1.Pod {
	release := helm.Chart(onosComponentName).Release(onosComponentName)
	client := kubernetes.NewForReleaseOrDie(release)
	podList, err := client.
		CoreV1().
		Pods().
		List()
	assert.NoError(t, err)
	return podList
}

// CrashPodOrFail deletes the given pod and fails the test if there is an error
func CrashPodOrFail(t *testing.T, pod *v1.Pod) {
	err := pod.Delete()
	assert.NoError(t, err)
}

// FindPodWithPrefix looks for the first pod whose name matches the given prefix string. The test is failed
// if no matching pod is found.
func FindPodWithPrefix(t *testing.T, prefix string) *v1.Pod {
	podList := GetPodListOrFail(t)
	for _, p := range podList {
		if strings.HasPrefix(p.Name, prefix) {
			return p
		}
	}
	assert.Failf(t, "No pod found matching %s", prefix)
	return nil
}

// TestE2NodeRestart :
func (s *TestSuite) TestE2NodeRestart(t *testing.T) {
	// Create a simulator
	sim := utils.CreateRanSimulatorWithNameOrDie(t, "e2node-restart")
	assert.NotNil(t, sim)

	e2Client := utils.GetE2Client(t, "e2node-restart-test")
	controlRequest := &e2.ControlRequest{
		Header: &e2.RequestHeader{
			EncodingType: 0,
			ServiceModel: nil,
		},
		E2NodeID: "e2 node",
	}
	controlResponse, err := e2Client.Control(context.Background(), controlRequest)
	assert.Nil(t, controlResponse)
	assert.Error(t, err)
	assert.Regexp(t, ".*channel 'e2 node' not found.*", err.Error())

	e2tPod := FindPodWithPrefix(t, "onos-e2t")
	CrashPodOrFail(t, e2tPod)

	time.Sleep(15 * time.Second)
	e2tPodReboot := FindPodWithPrefix(t, "onos-e2t")
	err = e2tPodReboot.Wait(45 * time.Second)
	assert.NoError(t, err)
	time.Sleep(15 * time.Second)

	e2Client2 := utils.GetE2Client(t, "e2node-restart-test")
	controlResponse2, err := e2Client2.Control(context.Background(), controlRequest)
	assert.Nil(t, controlResponse2)
	assert.Error(t, err)
	assert.Regexp(t, ".*channel 'e2 node' not found.*", err.Error())
}
