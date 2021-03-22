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
	KpmServiceModelID       = "oran-e2sm-kpm.v1"
	RcServiceModelID        = "oran-e2sm-rc-pre.v1"
	E2TServiceHost          = "onos-e2t"
	E2TServicePort          = 5150
	RansimServicePort       = 5150
)

var (
	SubscriptionServiceAddress = SubscriptionServiceHost + ":" + strconv.Itoa(SubscriptionServicePort)
)

func getRansimServiceAddress(ransimServiceHost string) string {
	return ransimServiceHost + ":" + strconv.Itoa(RansimServicePort)
}
