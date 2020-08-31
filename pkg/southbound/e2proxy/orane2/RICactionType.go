// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package orane2

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "ProtocolIE-Field.h"
//#include "RICaction-ToBeSetup-Item.h"
import "C"
import (
	"fmt"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2proxy/e2ctypes"
)

func newRicaActionType(rat e2ctypes.RICactionTypeT) (C.RICactionType_t, error) {
	switch rat {
	case e2ctypes.RICactionTypeT_RICactionType_report:
		return C.RICactionType_report, nil
	case e2ctypes.RICactionTypeT_RICactionType_insert:
		return C.RICactionType_insert, nil
	case e2ctypes.RICactionTypeT_RICactionType_policy:
		return C.RICactionType_policy, nil
	default:
		return C.RICactionType_report, fmt.Errorf("unexpected RICationType %v", rat)
	}
}
