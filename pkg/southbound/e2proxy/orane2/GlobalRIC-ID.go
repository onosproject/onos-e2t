// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package orane2

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "GlobalRIC-ID.h"
import "C"
import (
	"fmt"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2proxy/e2ctypes"
	"unsafe"
)

// XerGlobalRicIDT - used for test only
// Deprecated: Do not use.
func XerGlobalRicIDT(e2srIE *e2ctypes.GlobalRIC_IDT) ([]byte, error) {
	rsrIEC, err := newGlobalRicID(e2srIE)
	if err != nil {
		return nil, err
	}

	bytes, err := encodeXer(&C.asn_DEF_GlobalRIC_ID, unsafe.Pointer(rsrIEC))
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

// PerGlobalRicIDT - used for test only
// Deprecated: Do not use.
func PerGlobalRicIDT(e2srIE *e2ctypes.GlobalRIC_IDT) ([]byte, error) {
	rsrIEC, err := newGlobalRicID(e2srIE)
	if err != nil {
		return nil, err
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_GlobalRIC_ID, unsafe.Pointer(rsrIEC))
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

// Deprecated: Do not use.
func newGlobalRicID(id *e2ctypes.GlobalRIC_IDT) (*C.GlobalRIC_ID_t, error) {
	if len(id.PLMN_Identity) != 3 {
		return nil, fmt.Errorf("plmnID length is 3 - e2ap-v01.00.00.asn line 1105")
	}

	if id.Ric_ID.GetNumbits() != 20 {
		return nil, fmt.Errorf("ric-ID has to be 20 bits exactly - e2ap-v01.00.00.asn line 1076")
	}

	idC := C.GlobalRIC_ID_t{
		pLMN_Identity: *newOctetString(id.PLMN_Identity),
		ric_ID:        *newBitString(id.Ric_ID),
	}

	return &idC, nil
}
