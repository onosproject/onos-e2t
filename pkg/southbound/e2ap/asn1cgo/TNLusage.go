// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "TNLusage.h"
import "C"
import (
	"fmt"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	"unsafe"
)

func xerEncodeTnlusage(tnlusage *e2ap_ies.Tnlusage) ([]byte, error) {
	tnlusageCP, err := newTnlusage(tnlusage)
	if err != nil {
		return nil, err
	}

	bytes, err := encodeXer(&C.asn_DEF_TNLusage, unsafe.Pointer(tnlusageCP)) //ToDo - change name of C-encoder tag
	if err != nil {
		return nil, fmt.Errorf("xerEncodeTnlusage() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeTnlusage(tnlusage *e2ap_ies.Tnlusage) ([]byte, error) {
	tnlusageCP, err := newTnlusage(tnlusage)
	if err != nil {
		return nil, err
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_TNLusage, unsafe.Pointer(tnlusageCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeTnlusage() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeTnlusage(bytes []byte) (*e2ap_ies.Tnlusage, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_TNLusage)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeTnlusage((*C.TNLusage_t)(unsafePtr)) //ToDo - change name of C-struct
}

func perDecodeTnlusage(bytes []byte) (*e2ap_ies.Tnlusage, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_TNLusage)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeTnlusage((*C.TNLusage_t)(unsafePtr))
}

func newTnlusage(tnlusage *e2ap_ies.Tnlusage) (*C.TNLusage_t, error) {
	var ret C.TNLusage_t
	switch *tnlusage {
	case e2ap_ies.Tnlusage_TNLUSAGE_RIC_SERVICE:
		ret = C.TNLusage_ric_service
	case e2ap_ies.Tnlusage_TNLUSAGE_SUPPORT_FUNCTION:
		ret = C.TNLusage_support_function
	case e2ap_ies.Tnlusage_TNLUSAGE_BOTH:
		ret = C.TNLusage_both
	default:
		return nil, fmt.Errorf("unexpected Tnlusage %v", tnlusage)
	}

	return &ret, nil
}

func decodeTnlusage(tnlusageC *C.TNLusage_t) (*e2ap_ies.Tnlusage, error) {

	tnlusage := e2ap_ies.Tnlusage(int32(*tnlusageC))

	return &tnlusage, nil
}
