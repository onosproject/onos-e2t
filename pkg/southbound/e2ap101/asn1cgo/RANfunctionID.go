// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "RANfunctionID.h"
import "C"
import (
	"encoding/binary"
	"fmt"
	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	"unsafe"
)

func xerEncodeRanFunctionID(rfID *e2apies.RanfunctionId) ([]byte, error) {
	rfIDC := newRanFunctionID(rfID)

	bytes, err := encodeXer(&C.asn_DEF_RANfunctionID, unsafe.Pointer(&rfIDC))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeRanFunctionID() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeRanFunctionID(rfID *e2apies.RanfunctionId) ([]byte, error) {
	rfIDC := newRanFunctionID(rfID)

	bytes, err := encodePerBuffer(&C.asn_DEF_RANfunctionID, unsafe.Pointer(&rfIDC))
	if err != nil {
		return nil, fmt.Errorf("perEncodeRanFunctionID() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeRanFunctionID(bytes []byte) (*e2apies.RanfunctionId, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_RANfunctionID)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeRanFunctionID((*C.RANfunctionID_t)(unsafePtr)), nil
}

func perDecodeRanFunctionID(bytes []byte) (*e2apies.RanfunctionId, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_RANfunctionID)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeRanFunctionID((*C.RANfunctionID_t)(unsafePtr)), nil
}

func newRanFunctionID(rfID *e2apies.RanfunctionId) C.long {
	return C.long(rfID.Value)
}

func decodeRanFunctionIDBytes(ranFunctionIDCbytes []byte) *e2apies.RanfunctionId {
	rfC := (C.RANfunctionID_t)(binary.LittleEndian.Uint64(ranFunctionIDCbytes[0:8]))

	return decodeRanFunctionID(&rfC)
}

func decodeRanFunctionID(ranFunctionIDC *C.RANfunctionID_t) *e2apies.RanfunctionId {
	result := e2apies.RanfunctionId{
		Value: int32(*ranFunctionIDC),
	}

	return &result
}
