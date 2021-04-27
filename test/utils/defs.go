// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package utils

import (
	"strconv"
)

const (
	SubscriptionServiceHost = "onos-e2sub"
	SubscriptionServicePort = 5150
	KpmServiceModelName     = "oran-e2sm-kpm"
	Version1                = "v1"
	Version2                = "v2"
	RcServiceModelName      = "oran-e2sm-rc-pre"
	KpmServiceModelOIDV2    = "1.3.6.1.4.1.53148.1.2.2.2"

	E2TServiceHost    = "onos-e2t"
	E2TServicePort    = 5150
	RansimServicePort = 5150
)

var (
	SubscriptionServiceAddress = SubscriptionServiceHost + ":" + strconv.Itoa(SubscriptionServicePort)
)

func getRansimServiceAddress(ransimServiceHost string) string {
	return ransimServiceHost + ":" + strconv.Itoa(RansimServicePort)
}
