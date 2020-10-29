// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package main

import (
	"flag"
	"github.com/onosproject/onos-e2t/pkg/manager"
	"github.com/onosproject/onos-lib-go/pkg/certs"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"
)

var log = logging.GetLogger("main")

const probeFile = "/tmp/healthy"

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
	log.Infof("not using gRPC server just yet %p", opts)

	log.Info("Starting onos-e2t")
	cfg := manager.Config{
		CAPath:   *caPath,
		KeyPath:  *keyPath,
		CertPath: *certPath,
		GRPCPort: 5150,
		E2Port:   int(*sctpPort),
	}
	mgr := manager.NewManager(cfg)
	mgr.Run()

	if err := ioutil.WriteFile(probeFile, []byte("onos-e2t"), 0644); err != nil {
		log.Fatalf("Unable to write probe file %s", probeFile)
	}
	defer os.Remove(probeFile)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	<-sigCh
}
