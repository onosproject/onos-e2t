// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package orane2

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "RICaction-Admitted-List.h"
//#include "ProtocolIE-Field.h"
import "C"
import "github.com/onosproject/onos-e2t/pkg/southbound/e2proxy/e2ctypes"

// Deprecated: Do not use.
func decodeRicActionAdmittedList(raalC [48]byte) (*e2ctypes.RICaction_Admitted_ListT, error) {

	raal := e2ctypes.RICaction_Admitted_ListT{
		List: make([]*e2ctypes.RICaction_Admitted_ItemIEsT, 0),
	}
	return &raal, nil
}
