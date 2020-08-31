// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package sctp

import (
	"fmt"
	"github.com/ishidawataru/sctp"
	"net"
	"strings"
)

// StartSctpServer - the RIC is the SCTP server
func StartSctpServer(endpointAddrs string, endpointPort int, recv chan []byte, send chan []byte) error {

	var ips []net.IPAddr

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

	ln, err := sctp.ListenSCTP("sctp", addr)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}
	log.Infof("Listen on %s", ln.Addr())
	readBufsize := 2048
	for {
		conn, err := ln.Accept()
		if err != nil {
			return fmt.Errorf("failed to accept: %v", err)
		}
		log.Infof("Accepted Connection from RemoteAddr: %s", conn.RemoteAddr())
		wconn := sctp.NewSCTPSndRcvInfoWrappedConn(conn.(*sctp.SCTPConn))

		sndbuf, err := wconn.GetWriteBuffer()
		if err != nil {
			return fmt.Errorf("failed to get write buf: %v", err)
		}
		rcvbuf, err := wconn.GetReadBuffer()
		if err != nil {
			return fmt.Errorf("failed to get read buf: %v", err)
		}
		log.Infof("SndBufSize: %d, RcvBufSize: %d", sndbuf, rcvbuf)

		go func() {
			for {
				buf := make([]byte, readBufsize+128) // add overhead of SCTPSndRcvInfoWrappedConn
				// Blocking call
				n, err := conn.Read(buf)
				if err != nil {
					log.Infof("read failed: %v", err)
					break
				}
				recv <- buf[:n]
				log.Infof("read: %d", n)
			}
		}()

		go func() {
			// Blocking call
			for s := range send {
				n, err := conn.Write(s)
				if err != nil {
					log.Infof("write failed: %v", err)
					break
				}
				log.Infof("write: %d", n)
			}
		}()
	}
}
