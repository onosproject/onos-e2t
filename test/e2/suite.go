// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2

import (
	"fmt"
	petname "github.com/dustinkirkland/golang-petname"
	"github.com/onosproject/helmit/pkg/helm"
	"github.com/onosproject/helmit/pkg/test"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	testutils "github.com/onosproject/onos-ric-sdk-go/pkg/utils"
	"github.com/onosproject/onos-test/pkg/onostest"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func init() {
	logging.SetLevel(logging.WarnLevel)
}

// TestSuite is the primary onos-e2t test suite
type TestSuite struct {
	test.Suite
	release         *helm.Release
	E2TReplicaCount int64
}

func getInt(value interface{}) int64 {
	if i, ok := value.(int); ok {
		return int64(i)
	} else if i, ok := value.(float64); ok {
		return int64(i)
	} else if i, ok := value.(int64); ok {
		return i
	}
	return 0
}

const E2TReplicaCount int64 = 2

// CreateSdranRelease creates a helm release for an sd-ran instance
func (s *TestSuite) CreateSdranRelease() {
	registry := s.Arg("registry").String()

	install := s.Helm().
		Install("sd-ran", "sd-ran").
		RepoURL(onostest.SdranChartRepo).
		Set("onos-topo.global.image.tag", "latest")

	if registry != "" {
		install = install.
			Set("onos-topo.global.image.registry", registry).
			Set("onos-config.global.image.registry", registry).
			Set("onos-umbrella.global.image.registry", registry).
			Set("topo-discovery.global.image.registry", registry).
			Set("device-provisioner.global.image.registry", registry).
			Set("onos-cli.global.image.registry", registry).
			Set("global.image.registry", registry)
	}

	install = install.
		Set("import.onos-config.enabled", false).
		Set("import.onos-a1t.enabled", false).
		Set("import.onos-cli.enabled", false).
		Set("import.onos-e2t.enabled", true).
		Set("global.storage.consensus.enabled", "true").
		Set("onos-topo.image.tag", "latest").
		Set("onos-e2t.image.tag", "latest").
		Set("ran-simulator.image.tag", "latest").
		Set("onos-e2t.replicaCount", E2TReplicaCount).
		Set("onos-uenib.image.tag", "latest")

	var err error
	s.release, err = install.Wait().
		Get(s.Context())
	s.NoError(err)
}

// CreateRanSimulator creates a ran simulator
func (s *TestSuite) CreateRanSimulator() *helm.Helm {
	return s.CreateRanSimulatorWithName(petname.Generate(2, "-"))
}

// CreateRanSimulatorWithNameOrDie creates a simulator and fails the test if the creation returned an error
func (s *TestSuite) CreateRanSimulatorWithNameOrDie(simName string) *helm.Helm {
	sim := s.CreateRanSimulatorWithName(simName)
	s.NotNil(sim)
	return sim
}

// UninstallRanSimulatorOrDie uninstalls a simulator and fails the test if the operation returned an error
func (s *TestSuite) UninstallRanSimulatorOrDie(sim *helm.Helm, simName string) {
	s.NoError(sim.Uninstall(simName).Do(s.Context()))
}

// FindSimulatorPodOrDie finds the pod for the target simulator
func (s *TestSuite) FindSimulatorPodOrDie(simName string) v1.Pod {
	pods, err := s.CoreV1().Pods(s.Namespace()).List(s.Context(), metav1.ListOptions{
		LabelSelector: fmt.Sprintf("name=%s-device-simulator", simName),
	})
	s.NoError(err)
	s.Len(pods.Items, 1)
	return pods.Items[0]
}

// CrashSimulatorPodOrDie crashes the target simulator
func (s *TestSuite) CrashSimulatorPodOrDie(simName string) {
	pod := s.FindSimulatorPodOrDie(simName)
	err := s.CoreV1().Pods(s.Namespace()).Delete(s.Context(), pod.Name, metav1.DeleteOptions{})
	s.NoError(err)
}

// CreateRanSimulatorWithName creates a ran simulator
func (s *TestSuite) CreateRanSimulatorWithName(name string) *helm.Helm {
	registry := s.Arg("registry").String()

	simHelm := s.Helm()
	install := simHelm.
		Install(name, "ran-simulator").
		RepoURL(onostest.SdranChartRepo).
		Set("onos-topo.global.image.tag", "latest").
		Set("image.tag", "latest").
		Set("fullnameOverride", "").
		Set("global.image.registry", registry)

	_, err := install.Wait().
		Get(s.Context())

	s.NoError(err, "could not install device simulator %v", err)

	return simHelm
}

// SetupSuite sets up the onos-e2t test suite
func (s *TestSuite) SetupSuite() {
	s.CreateSdranRelease()
	s.E2TReplicaCount = getInt(s.release.Get("onos-e2t.replicaCount"))

	testutils.StartTestProxy()
}

// TearDownSuite tears down the test ONOS proxy.
func (s *TestSuite) TearDownSuite() {
	testutils.StopTestProxy()
}

var _ test.SetupSuite = (*TestSuite)(nil)
var _ test.TearDownSuite = (*TestSuite)(nil)
