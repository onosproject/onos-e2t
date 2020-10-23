// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2proxy"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2proxy/connections"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2proxy/e2ctypes"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2proxy/orane2"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2proxy/sctp"
	"github.com/onosproject/onos-lib-go/pkg/certs"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"io/ioutil"
	"os"
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

	sendChan := make(chan []byte)
	defer close(sendChan)
	recvChan := make(chan []byte)
	defer close(recvChan)
	startedChan := make(chan bool)
	defer close(startedChan)

	// nolint:staticcheck
	e2inChan := make(chan *e2ctypes.E2AP_PDUT)
	defer close(e2inChan)

	go func() {
		for r := range recvChan {
			fmt.Printf("Received %d\n", len(r))
			e2apPdu, err := orane2.PerDecodeE2apPdu(r)
			if err != nil {
				fmt.Printf("unable to parse response as PER %v\n", err)
				// Try XER instead - used for initial E2Setup Message
				e2apPdu, err = orane2.XerDecodeE2apPdu(r)
				if err != nil {
					fmt.Printf("unable to parse response as XER %v %s\n", err, string(r))
					break
					//os.Exit(-1)
				}
			}
			e2inChan <- e2apPdu
			fmt.Printf("Payload %v\n", e2apPdu)
		}
	}()

	go func() {
		for pduIn := range e2inChan {

			// A simple reactionary model
			pc, err := e2proxy.GetE2apPduType(pduIn)
			if err != nil {
				fmt.Printf("%v\n", err)
				continue
			}
			switch pc {
			case e2ctypes.ProcedureCodeT_ProcedureCode_id_E2setup:
				id := pduIn.GetInitiatingMessage().GetE2SetupRequest().ProtocolIEs.List[0].GetGlobalE2Node_ID().GetGNB().GetGlobalGNB_ID().GnbId.GetGnb_ID().BitString
				plmnId := pduIn.GetInitiatingMessage().GetE2SetupRequest().ProtocolIEs.List[0].GetGlobalE2Node_ID().GetGNB().GetGlobalGNB_ID().PlmnId
				connections.CreateConnection(pduIn.GetInitiatingMessage().GetE2SetupRequest())

				fmt.Printf("Received E2SetupRequest from 0x%x plmnid %s\n", binary.BigEndian.Uint32(id), plmnId)
				e2setupResp := e2proxy.NewE2SetupResponse()
				fmt.Printf("Sending E2SetupResponse\n")
				e2setupRespBytes, err := orane2.XerEncodeE2apPdu(e2setupResp)
				if err != nil {
					fmt.Printf("%v\n", err)
					continue
				}
				sendChan <- e2setupRespBytes

				//Now send a RIC subscription request
				fmt.Printf("Sending RICsubscriptionRequest\n")
				e2subReq := e2proxy.NewRICsubscriptionRequest(22, 6, 1,
					[]byte{0x41, 0x63, 0x74, 0x69, 0x6F, 0x6E, 0x44, 0x65, 0x66,
						0x69, 0x6F, 0x6E, 0x54, 0x72, 0x69, 0x67,
						0x67, 0x65, 0x72, 0x73}) //Spells "ActionDefionTriggers"
				fmt.Printf("Sending %+v\n", e2subReq)
				e2subReqBytes, err := orane2.PerEncodeE2apPdu(e2subReq)
				if err != nil {
					fmt.Printf("%v\n", err)
					continue
				}
				sendChan <- e2subReqBytes
			case e2ctypes.ProcedureCodeT_ProcedureCode_id_RICsubscription:
				fmt.Printf("Received RICsubscriptionResponse\n")

			default:
				fmt.Printf("Unhandled response %v", pc)
			}
			for index, conn := range connections.ListConnections() {
				fmt.Printf("Connection #%d is %v\n", index, conn)
			}
		}
	}()

	go func() {
		if err = sctp.StartSctpServer("0.0.0.0", int(*sctpPort), recvChan, sendChan, startedChan); err != nil {
			os.Exit(-1)
		}
	}()

	<-startedChan // block until server starts listening
	log.Info("Starting onos-e2t")
	if err := ioutil.WriteFile(probeFile, []byte("onos-e2t"), 0644); err != nil {
		log.Fatalf("Unable to write probe file %s", probeFile)
	}
	defer os.Remove(probeFile)

	<-startedChan // block again to stay running
}
