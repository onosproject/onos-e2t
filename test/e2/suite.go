// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package e2

import (
	"github.com/onosproject/helmit/pkg/helm"
	"github.com/onosproject/helmit/pkg/input"
	"github.com/onosproject/helmit/pkg/test"
	"github.com/onosproject/onos-e2t/test/utils"
)

// TestSuite is the primary onos-e2t test suite
type TestSuite struct {
	test.Suite
	c               *input.Context
	release         *helm.HelmRelease
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

// SetupTestSuite sets up the onos-e2t test suite
func (s *TestSuite) SetupTestSuite(c *input.Context) error {
	s.c = c
	sdran, err := utils.CreateSdranRelease(c)
	if err != nil {
		return err
	}

	registry := c.GetArg("registry").String("")

	s.release = sdran.Set("global.image.registry", registry)
	r := s.release.Install(true)
	s.E2TReplicaCount = getInt(sdran.Get("onos-e2t.replicaCount"))
	return r
}
