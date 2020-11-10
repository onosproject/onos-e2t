// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package utils

import (
	"github.com/onosproject/helmit/pkg/helm"
	"github.com/onosproject/onos-test/pkg/onostest"
)

// CreateSdranRelease creates a helm release for an sd-ran instance
func CreateSdranRelease() (*helm.HelmRelease, error) {

	sdran := helm.Chart("sd-ran", onostest.SdranChartRepo).
		Release("sd-ran").
		Set("import.onos-config.enabled", false).
		Set("import.onos-e2sub.enabled", false).
		Set("import.onos-topo.enabled", false).
		Set("onos-e2t.image.tag", "latest")

	return sdran, nil
}
