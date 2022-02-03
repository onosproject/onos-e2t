// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"github.com/onosproject/helmit/pkg/registry"
	"github.com/onosproject/helmit/pkg/simulation"
	"github.com/onosproject/onos-e2t/sim/e2"
)

func main() {
	registry.RegisterSimulationSuite("e2", &e2.SimSuite{})
	simulation.Main()
}
