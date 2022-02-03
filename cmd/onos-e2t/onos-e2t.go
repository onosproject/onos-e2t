// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/onosproject/onos-e2t/pkg/manager"
	"github.com/onosproject/onos-lib-go/pkg/logging"
)

type arrayFlags []string

func (i *arrayFlags) String() string {
	return "my string representation"
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

var log = logging.GetLogger("main")

func main() {
	var serviceModelPlugins arrayFlags
	flag.Var(&serviceModelPlugins, "serviceModel", "names of service model plugins to load (repeated)")
	caPath := flag.String("caPath", "", "path to CA certificate")
	keyPath := flag.String("keyPath", "", "path to client private key")
	certPath := flag.String("certPath", "", "path to client certificate")
	sctpPort := flag.Uint("sctpport", 36421, "sctp server port")
	topoEndpoint := flag.String("topoEndpoint", "onos-topo:5150", "onos-topo endpoint address")

	flag.Parse()

	log.Info("Starting onos-e2t")
	cfg := manager.Config{
		CAPath:              *caPath,
		KeyPath:             *keyPath,
		CertPath:            *certPath,
		GRPCPort:            5150,
		E2Port:              int(*sctpPort),
		TopoAddress:         *topoEndpoint,
		ServiceModelPlugins: serviceModelPlugins,
	}
	mgr := manager.NewManager(cfg)
	mgr.Run()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	<-sigCh

	mgr.Close()
}
