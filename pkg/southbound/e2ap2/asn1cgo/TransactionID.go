// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

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
	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-ies"
)

func xerDecodeTransactionID(bytes []byte) (*e2apies.TransactionId, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_TransactionID)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	ricIndicationC := (*C.TransactionID_t)(unsafePtr)
	ricIndication := decodeTransactionID(ricIndicationC)

	return ricIndication, nil
}

func newTransactionID(transactionID *e2apies.TransactionId) *C.TransactionID_t {
	res := C.long(transactionID.GetValue())
	return &res
}

func decodeTransactionIDBytes(transactionIDCchoice []byte) *e2apies.TransactionId {
	transactionID := C.long(binary.LittleEndian.Uint64(transactionIDCchoice[0:8]))

	return decodeTransactionID(&transactionID)
}

func decodeTransactionID(transactionIDC *C.TransactionID_t) *e2apies.TransactionId {

	return &e2apies.TransactionId{
		Value: int32(*transactionIDC),
	}
}