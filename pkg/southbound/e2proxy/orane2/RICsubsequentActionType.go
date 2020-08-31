// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package orane2

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "RICsubsequentAction.h"
import "C"
import (
	"fmt"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2proxy/e2ctypes"
)

func newRicSubsequentActionType(rsat e2ctypes.RICsubsequentActionTypeT) (C.e_RICsubsequentActionType, error) {
	switch rsat {
	case e2ctypes.RICsubsequentActionTypeT_RICsubsequentActionType_wait:
		return C.RICsubsequentActionType_wait, nil
	case e2ctypes.RICsubsequentActionTypeT_RICsubsequentActionType_continue:
		return C.RICsubsequentActionType_continue, nil
	default:
		return C.RICsubsequentActionType_wait, fmt.Errorf("unexpected RicSubsequentActionType %v", rsat)
	}
}
