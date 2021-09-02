// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "RICcontrolStatusVone.h"
import "C"
import (
	"encoding/binary"
	"fmt"
	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
)

func newRicControlStatus(rcs e2apies.RiccontrolStatus) (*C.RICcontrolStatusVone_t, error) {
	var ret C.RICcontrolStatusVone_t
	switch rcs {
	case e2apies.RiccontrolStatus_RICCONTROL_STATUS_SUCCESS:
		ret = C.RICcontrolStatusVone_success
	case e2apies.RiccontrolStatus_RICCONTROL_STATUS_REJECTED:
		ret = C.RICcontrolStatusVone_rejected
	case e2apies.RiccontrolStatus_RICCONTROL_STATUS_FAILED:
		ret = C.RICcontrolStatusVone_failed
	default:
		return nil, fmt.Errorf("unexpected RICcontrolStatus %v", rcs)
	}
	return &ret, nil
}

func decodeRicControlStatusBytes(bytes []byte) e2apies.RiccontrolStatus {
	rcsC := C.long(binary.LittleEndian.Uint64(bytes[:8]))
	return decodeRicControlStatus(&rcsC)
}

func decodeRicControlStatus(rcsC *C.RICcontrolStatusVone_t) e2apies.RiccontrolStatus {
	return e2apies.RiccontrolStatus(*rcsC)
}
