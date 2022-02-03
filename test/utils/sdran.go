// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package utils

import (
	"context"
	"testing"
	"time"

	"github.com/onosproject/helmit/pkg/helm"
	"github.com/onosproject/helmit/pkg/input"
	"github.com/onosproject/helmit/pkg/kubernetes"
	"github.com/onosproject/helmit/pkg/util/random"
	"github.com/onosproject/onos-test/pkg/onostest"
	"github.com/stretchr/testify/assert"
)

const E2TReplicaCount int64 = 2

func getCredentials() (string, string, error) {
	kubClient, err := kubernetes.New()
	if err != nil {
		return "", "", err
	}
	secrets, err := kubClient.CoreV1().Secrets().Get(context.Background(), onostest.SecretsName)
	if err != nil {
		return "", "", err
	}
	username := string(secrets.Object.Data["sd-ran-username"])
	password := string(secrets.Object.Data["sd-ran-password"])

	return username, password, nil
}

// CreateSdranRelease creates a helm release for an sd-ran instance
func CreateSdranRelease(c *input.Context) (*helm.HelmRelease, error) {
	username, password, err := getCredentials()
	registry := c.GetArg("registry").String("")
	if err != nil {
		return nil, err
	}

	sdran := helm.Chart("sd-ran", onostest.SdranChartRepo).
		Release("sd-ran").
		SetUsername(username).
		SetPassword(password).
		WithTimeout(6*time.Minute).
		Set("import.onos-config.enabled", false).
		Set("global.storage.consensus.enabled", "true").
		Set("onos-topo.image.tag", "latest").
		Set("onos-e2t.image.tag", "latest").
		Set("ran-simulator.image.tag", "latest").
		Set("onos-e2t.replicaCount", E2TReplicaCount).
		Set("onos-uenib.image.tag", "latest").
		Set("global.image.registry", registry)

	return sdran, nil
}

// CreateRanSimulator creates a ran simulator
func CreateRanSimulator(t *testing.T, c *input.Context) *helm.HelmRelease {
	return CreateRanSimulatorWithName(t, c, random.NewPetName(2))
}

// CreateRanSimulatorWithNameOrDie creates a simulator and fails the test if the creation returned an error
func CreateRanSimulatorWithNameOrDie(t *testing.T, c *input.Context, simName string) *helm.HelmRelease {
	sim := CreateRanSimulatorWithName(t, c, simName)
	assert.NotNil(t, sim)
	return sim
}

// UninstallRanSimulatorOrDie uninstalls a simulator and fails the test if the operation returned an error
func UninstallRanSimulatorOrDie(t *testing.T, sim *helm.HelmRelease) {
	assert.NoError(t, sim.Uninstall())
}

// CreateRanSimulatorWithName creates a ran simulator
func CreateRanSimulatorWithName(t *testing.T, c *input.Context, name string) *helm.HelmRelease {
	username, password, err := getCredentials()
	assert.NoError(t, err)

	registry := c.GetArg("registry").String("")

	simulator := helm.
		Chart("ran-simulator", onostest.SdranChartRepo).
		Release(name).
		SetUsername(username).
		SetPassword(password).
		Set("image.tag", "latest").
		Set("fullnameOverride", "").
		Set("global.image.registry", registry)
	err = simulator.Install(true)
	assert.NoError(t, err, "could not install device simulator %v", err)

	return simulator
}
