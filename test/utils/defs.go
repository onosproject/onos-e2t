// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package utils

import (
	"strconv"
)

const (
	KpmServiceModelName  = "oran-e2sm-kpm"
	Version1             = "v1"
	Version2             = "v2"
	RcServiceModelName   = "oran-e2sm-rc-pre"
	KpmServiceModelOIDV2 = "1.3.6.1.4.1.53148.1.2.2.2"

	E2TServiceHost    = "onos-e2t"
	E2TServicePort    = 5150
	RansimServicePort = 5150
)

var (
	E2tServiceAddress = E2TServiceHost + ":" + strconv.Itoa(E2TServicePort)
)

func getRansimServiceAddress(ransimServiceHost string) string {
	return ransimServiceHost + ":" + strconv.Itoa(RansimServicePort)
}
