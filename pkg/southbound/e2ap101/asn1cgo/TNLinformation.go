// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "TNLinformation.h"
import "C"

import (
	"fmt"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	"unsafe"
)

func xerEncodeTnlinformation(tnlinformation *e2ap_ies.Tnlinformation) ([]byte, error) {
	tnlinformationCP, err := newTnlinformation(tnlinformation)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeTnlinformation() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_TNLinformation, unsafe.Pointer(tnlinformationCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeTnlinformation() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeTnlinformation(tnlinformation *e2ap_ies.Tnlinformation) ([]byte, error) {
	tnlinformationCP, err := newTnlinformation(tnlinformation)
	if err != nil {
		return nil, fmt.Errorf("perEncodeTnlinformation() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_TNLinformation, unsafe.Pointer(tnlinformationCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeTnlinformation() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeTnlinformation(bytes []byte) (*e2ap_ies.Tnlinformation, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_TNLinformation)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeTnlinformation((*C.TNLinformation_t)(unsafePtr))
}

func perDecodeTnlinformation(bytes []byte) (*e2ap_ies.Tnlinformation, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_TNLinformation)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeTnlinformation((*C.TNLinformation_t)(unsafePtr))
}

func newTnlinformation(tnlinformation *e2ap_ies.Tnlinformation) (*C.TNLinformation_t, error) {

	tnlinformationC := C.TNLinformation_t{}

	tnlAddressC := newBitString(tnlinformation.TnlAddress)
	tnlinformationC.tnlAddress = *tnlAddressC

	if tnlinformation.TnlPort != nil {
		tnlPortC := newBitString(tnlinformation.TnlPort)
		tnlinformationC.tnlPort = tnlPortC
	}

	return &tnlinformationC, nil
}

func decodeTnlinformation(tnlinformationC *C.TNLinformation_t) (*e2ap_ies.Tnlinformation, error) {

	var err error
	tnlinformation := e2ap_ies.Tnlinformation{}

	tnlinformation.TnlAddress, err = decodeBitString(&tnlinformationC.tnlAddress)
	if err != nil {
		return nil, fmt.Errorf("decodeBitString() %s", err.Error())
	}

	if tnlinformationC.tnlPort != nil {
		tnlinformation.TnlPort, err = decodeBitString(tnlinformationC.tnlPort)
		if err != nil {
			return nil, fmt.Errorf("decodeBitString() %s", err.Error())
		}
	}

	return &tnlinformation, nil
}
