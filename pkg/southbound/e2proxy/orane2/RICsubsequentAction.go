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
	"github.com/onosproject/onos-e2t/pkg/southbound/e2proxy/e2ctypes"
)

func newRicSubsequentAction(sa *e2ctypes.RICsubsequentActionT) (*C.RICsubsequentAction_t, error) {
	rsatC, err := newRicSubsequentActionType(sa.RicSubsequentActionType)
	if err != nil {
		return nil, err
	}
	rttw, err := newRICtimeToWait(sa.RicTimeToWait)
	if err != nil {
		return nil, err
	}
	saC := C.RICsubsequentAction_t{
		ricSubsequentActionType: C.long(rsatC),
		ricTimeToWait:           C.long(rttw),
	}
	return &saC, nil
}
