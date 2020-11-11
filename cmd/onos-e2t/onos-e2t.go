// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/onosproject/onos-e2t/pkg/manager"
	"github.com/onosproject/onos-lib-go/pkg/certs"
	"github.com/onosproject/onos-lib-go/pkg/logging"
)

var log = logging.GetLogger("main")

func main() {
	caPath := flag.String("caPath", "", "path to CA certificate")
	keyPath := flag.String("keyPath", "", "path to client private key")
	certPath := flag.String("certPath", "", "path to client certificate")
	sctpPort := flag.Uint("sctpport", 36421, "sctp server port")
	flag.Parse()

	opts, err := certs.HandleCertPaths(*caPath, *keyPath, *certPath, true)
	if err != nil {
		log.Fatal(err)
	}

	log.Info("Starting onos-e2t")
	cfg := manager.Config{
		CAPath:   *caPath,
		KeyPath:  *keyPath,
		CertPath: *certPath,
		GRPCPort: 5150,
		E2Port:   int(*sctpPort),
	}
	mgr := manager.NewManager(cfg, opts...)
	mgr.Run()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	<-sigCh

	mgr.Close()
}
