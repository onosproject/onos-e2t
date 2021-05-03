// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "ENGNB-ID.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	"unsafe"
)

func xerEncodeEngnbID(engnbID *e2ap_ies.EngnbId) ([]byte, error) {
	engnbIDCP, err := newEngnbID(engnbID)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeEngnbID() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_ENGNB_ID, unsafe.Pointer(engnbIDCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeEngnbID() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeEngnbID(engnbID *e2ap_ies.EngnbId) ([]byte, error) {
	engnbIDCP, err := newEngnbID(engnbID)
	if err != nil {
		return nil, fmt.Errorf("perEncodeEngnbID() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_ENGNB_ID, unsafe.Pointer(engnbIDCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeEngnbID() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeEngnbID(bytes []byte) (*e2ap_ies.EngnbId, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_ENGNB_ID)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeEngnbID((*C.ENGNB_ID_t)(unsafePtr))
}

func perDecodeEngnbID(bytes []byte) (*e2ap_ies.EngnbId, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_ENGNB_ID)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeEngnbID((*C.ENGNB_ID_t)(unsafePtr))
}

func newEngnbID(engnbID *e2ap_ies.EngnbId) (*C.ENGNB_ID_t, error) {

	var pr C.ENGNB_ID_PR
	choiceC := [48]byte{} //ToDo - Check if number of bytes is sufficient

	switch choice := engnbID.EngnbId.(type) {
	case *e2ap_ies.EngnbId_GNbId:
		pr = C.ENGNB_ID_PR_gNB_ID

		bsC := newBitString(choice.GNbId)
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(bsC.buf))))
		binary.LittleEndian.PutUint64(choiceC[8:], uint64(bsC.size))
		binary.LittleEndian.PutUint32(choiceC[16:], uint32(bsC.bits_unused))
	default:
		return nil, fmt.Errorf("newEngnbID() %T not yet implemented", choice)
	}

	engnbIDC := C.ENGNB_ID_t{
		present: pr,
		choice:  choiceC,
	}

	return &engnbIDC, nil
}

func decodeEngnbID(engnbIDC *C.ENGNB_ID_t) (*e2ap_ies.EngnbId, error) {

	engnbID := new(e2ap_ies.EngnbId)

	switch engnbIDC.present {
	case C.ENGNB_ID_PR_gNB_ID:
		engnbIDstructC := newBitStringFromArray(engnbIDC.choice)
		engNbID, err := decodeBitString(engnbIDstructC)
		if err != nil {
			return nil, err
		}
		engnbID.EngnbId = &e2ap_ies.EngnbId_GNbId{
			GNbId: engNbID,
		}
	default:
		return nil, fmt.Errorf("decodeEngnbID() %v not yet implemented", engnbIDC.present)
	}

	return engnbID, nil
}

func decodeEngnbIDBytes(array [8]byte) (*e2ap_ies.EngnbId, error) {
	engnbIDC := (*C.ENGNB_ID_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeEngnbID(engnbIDC)
}
