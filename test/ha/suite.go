// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package ha

import (
	"github.com/onosproject/helmit/pkg/helm"
	"github.com/onosproject/helmit/pkg/input"
	"github.com/onosproject/helmit/pkg/test"
	"github.com/onosproject/onos-e2t/test/utils"
)

// TestSuite is the primary onos-e2t test suite
type TestSuite struct {
	test.Suite
	release *helm.HelmRelease
	c       *input.Context
}

// SetupTestSuite sets up the onos-e2t test suite
func (s *TestSuite) SetupTestSuite(c *input.Context) error {
	s.c = c
	sdran, err := utils.CreateSdranRelease(c)
	if err != nil {
		return err
	}

	registry := c.GetArg("registry").String("")
	s.release = sdran.Set("global.image.registry", registry)
	return s.release.Install(true)
}
