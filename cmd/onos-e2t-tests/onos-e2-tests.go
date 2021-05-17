// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package main

import (
	"github.com/onosproject/helmit/pkg/registry"
	"github.com/onosproject/helmit/pkg/test"
	"github.com/onosproject/onos-e2t/test/e2"
	"github.com/onosproject/onos-e2t/test/ha"
)

func main() {
	registry.RegisterTestSuite("e2", &e2.TestSuite{})
	registry.RegisterTestSuite("ha", &ha.TestSuite{})
	test.Main()
}
