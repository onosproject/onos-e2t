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

func newRICtimeToWait(rttw e2ctypes.RICtimeToWaitT) (C.e_RICtimeToWait, error) {
	switch rttw {
	case e2ctypes.RICtimeToWaitT_RICtimeToWait_zero:
		return C.RICtimeToWait_zero, nil
	case e2ctypes.RICtimeToWaitT_RICtimeToWait_w1ms:
		return C.RICtimeToWait_w1ms, nil
	case e2ctypes.RICtimeToWaitT_RICtimeToWait_w2ms:
		return C.RICtimeToWait_w2ms, nil
	case e2ctypes.RICtimeToWaitT_RICtimeToWait_w5ms:
		return C.RICtimeToWait_w5ms, nil
	case e2ctypes.RICtimeToWaitT_RICtimeToWait_w10ms:
		return C.RICtimeToWait_w10ms, nil
	case e2ctypes.RICtimeToWaitT_RICtimeToWait_w20ms:
		return C.RICtimeToWait_w20ms, nil
	case e2ctypes.RICtimeToWaitT_RICtimeToWait_w30ms:
		return C.RICtimeToWait_w30ms, nil
	case e2ctypes.RICtimeToWaitT_RICtimeToWait_w40ms:
		return C.RICtimeToWait_w40ms, nil
	case e2ctypes.RICtimeToWaitT_RICtimeToWait_w50ms:
		return C.RICtimeToWait_w50ms, nil
	case e2ctypes.RICtimeToWaitT_RICtimeToWait_w100ms:
		return C.RICtimeToWait_w100ms, nil
	case e2ctypes.RICtimeToWaitT_RICtimeToWait_w200ms:
		return C.RICtimeToWait_w200ms, nil
	case e2ctypes.RICtimeToWaitT_RICtimeToWait_w500ms:
		return C.RICtimeToWait_w500ms, nil
	case e2ctypes.RICtimeToWaitT_RICtimeToWait_w1s:
		return C.RICtimeToWait_w1s, nil
	case e2ctypes.RICtimeToWaitT_RICtimeToWait_w2s:
		return C.RICtimeToWait_w2s, nil
	case e2ctypes.RICtimeToWaitT_RICtimeToWait_w5s:
		return C.RICtimeToWait_w5s, nil
	case e2ctypes.RICtimeToWaitT_RICtimeToWait_w10s:
		return C.RICtimeToWait_w10s, nil
	case e2ctypes.RICtimeToWaitT_RICtimeToWait_w20s:
		return C.RICtimeToWait_w20s, nil
	case e2ctypes.RICtimeToWaitT_RICtimeToWait_w60s:
		return C.RICtimeToWait_w60s, nil
	default:
		return C.RICtimeToWait_zero, fmt.Errorf("newRICtimeToWait() unexpected %v", rttw)
	}
}
