// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package asn1cgo

// #cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
// #cgo LDFLAGS: -lm
// #include <stdio.h>
// #include <stdlib.h>
// #include <assert.h>
// #include "TransactionID.h"
import "C"
import (
	"encoding/binary"
	"fmt"
	"unsafe"

	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
)

func xerEncodeTransactionID(transactionID *e2apies.TransactionId) ([]byte, error) {
	transactionIDCP := newTransactionID(transactionID)

	bytes, err := encodeXer(&C.asn_DEF_TransactionID, unsafe.Pointer(transactionIDCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeTnlusage() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeTransactionID(transactionID *e2apies.TransactionId) ([]byte, error) {
	transactionIDCP := newTransactionID(transactionID)

	bytes, err := encodePerBuffer(&C.asn_DEF_TransactionID, unsafe.Pointer(transactionIDCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeTnlusage() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeTransactionID(bytes []byte) (*e2apies.TransactionId, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_TransactionID)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeTransactionID((*C.TransactionID_t)(unsafePtr)), nil
}

func perDecodeTransactionID(bytes []byte) (*e2apies.TransactionId, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_TransactionID)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeTransactionID((*C.TransactionID_t)(unsafePtr)), nil
}

func newTransactionID(transactionID *e2apies.TransactionId) *C.TransactionID_t {
	res := C.long(transactionID.GetValue())
	return &res
}

func decodeTransactionIDBytes(bytes []byte) *e2apies.TransactionId {
	transactionID := C.long(binary.LittleEndian.Uint64(bytes[:8]))

	return decodeTransactionID(&transactionID)
}

func decodeTransactionID(transactionIDC *C.TransactionID_t) *e2apies.TransactionId {

	return &e2apies.TransactionId{
		Value: int32(*transactionIDC),
	}
}
