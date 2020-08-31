// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package sctp

import (
	"fmt"
	"github.com/ishidawataru/sctp"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"net"
	"strings"
)

var log = logging.GetLogger("southbound", "e2proxy", "sctp")

// StartSctpClient - the E2Node is the client and sends an E2SetupRequest first
func StartSctpClient(endpointAddrs string, endpointPort int, recv chan []byte, send chan []byte) error {
	ips := []net.IPAddr{}

	for _, i := range strings.Split(endpointAddrs, ",") {
		if a, err := net.ResolveIPAddr("ip", i); err == nil {
			log.Infof("Resolved address '%s' to %s", i, a)
			ips = append(ips, *a)
		} else {
			return fmt.Errorf("error resolving address '%s': %v", i, err)
		}
	}

	addr := &sctp.SCTPAddr{
		IPAddrs: ips,
		Port:    endpointPort,
	}
	log.Infof("raw addr: %+v\n", addr.ToRawSockAddrBuf())

	var laddr *sctp.SCTPAddr

	conn, err := sctp.DialSCTP("sctp", laddr, addr)
	if err != nil {
		return fmt.Errorf("failed to dial: %v", err)
	}

	log.Infof("Dial LocalAddr: %s; RemoteAddr: %s", conn.LocalAddr(), conn.RemoteAddr())

	sndbuf, err := conn.GetWriteBuffer()
	if err != nil {
		log.Fatalf("failed to get write buf: %v", err)
	}
	rcvbuf, err := conn.GetReadBuffer()
	if err != nil {
		log.Fatalf("failed to get read buf: %v", err)
	}
	log.Infof("SndBufSize: %d, RcvBufSize: %d", sndbuf, rcvbuf)

	ppid := 0
	bufsize := 1024

	// Sending thread
	go func() {

		// Block here waiting for something on send channel
		for s := range send {
			ppid++

			info := &sctp.SndRcvInfo{
				Stream: uint16(ppid),
				PPID:   uint32(ppid),
			}
			n, err := conn.SCTPWrite(s, info)
			if err != nil {
				log.Errorf("failed to write %d: %v", n, err)
				break
			}
			log.Debug("Sent %d bytes", n)
		}
		log.Errorf("Closed the sending loop")
	}()

	// Receiving thread
	for {
		ppid++
		if err := conn.SubscribeEvents(sctp.SCTP_EVENT_DATA_IO); err != nil {
			log.Errorf("Unable to subscribe %v", err)
			return err
		}
		buf := make([]byte, bufsize)
		n, info, err := conn.SCTPRead(buf)
		if err != nil {
			return err
		}
		log.Debugf("read: len %d, info: %+v", n, info)
		recv <- buf[:n]
	}
}
