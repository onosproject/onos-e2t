// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	e2sim "github.com/onosproject/onos-e2t/sim/e2"
	e2 "github.com/onosproject/onos-ric-sdk-go/pkg/e2/v1beta1"
	"os"
	"os/signal"
)

func main() {
	appID := os.Args[1]
	instanceID := os.Args[2]

	app := e2sim.NewApp(e2.AppID(appID), e2.InstanceID(instanceID))
	if err := app.Start(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch

	if err := app.Stop(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
