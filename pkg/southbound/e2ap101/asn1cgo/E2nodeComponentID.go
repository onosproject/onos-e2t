// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2nodeComponentID.h" //ToDo - if there is an anonymous C-struct option, it would require linking additional C-struct file definition (the one above or before)
import "C"

import (
	"encoding/binary"
	"fmt"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	"unsafe"
)

func xerEncodeE2nodeComponentID(e2nodeComponentID *e2ap_ies.E2NodeComponentId) ([]byte, error) {
	e2nodeComponentIDCP, err := newE2nodeComponentID(e2nodeComponentID)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2nodeComponentID() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2nodeComponentID, unsafe.Pointer(e2nodeComponentIDCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2nodeComponentID() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeE2nodeComponentID(e2nodeComponentID *e2ap_ies.E2NodeComponentId) ([]byte, error) {
	e2nodeComponentIDCP, err := newE2nodeComponentID(e2nodeComponentID)
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2nodeComponentID() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2nodeComponentID, unsafe.Pointer(e2nodeComponentIDCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2nodeComponentID() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeE2nodeComponentID(bytes []byte) (*e2ap_ies.E2NodeComponentId, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2nodeComponentID)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2nodeComponentID((*C.E2nodeComponentID_t)(unsafePtr))
}

func perDecodeE2nodeComponentID(bytes []byte) (*e2ap_ies.E2NodeComponentId, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2nodeComponentID)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2nodeComponentID((*C.E2nodeComponentID_t)(unsafePtr))
}

func newE2nodeComponentID(e2nodeComponentID *e2ap_ies.E2NodeComponentId) (*C.E2nodeComponentID_t, error) {

	var pr C.E2nodeComponentID_PR //ToDo - verify correctness of the name
	choiceC := [8]byte{}          //ToDo - Check if number of bytes is sufficient

	switch choice := e2nodeComponentID.E2NodeComponentId.(type) {
	case *e2ap_ies.E2NodeComponentId_E2NodeComponentTypeGnbCuUp:
		pr = C.E2nodeComponentID_PR_e2nodeComponentTypeGNB_CU_UP //ToDo - Check if it's correct PR's name

		im, err := newE2nodeComponentGnbCuUpID(choice.E2NodeComponentTypeGnbCuUp)
		if err != nil {
			return nil, fmt.Errorf("newE2nodeComponentGnbCuUpID() %s", err.Error())
		}
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(im))))
	case *e2ap_ies.E2NodeComponentId_E2NodeComponentTypeGnbDu:
		pr = C.E2nodeComponentID_PR_e2nodeComponentTypeGNB_DU //ToDo - Check if it's correct PR's name

		im, err := newE2nodeComponentGnbDuID(choice.E2NodeComponentTypeGnbDu)
		if err != nil {
			return nil, fmt.Errorf("newE2nodeComponentGnbDuID() %s", err.Error())
		}
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(im))))
	default:
		return nil, fmt.Errorf("newE2nodeComponentID() %T not yet implemented", choice)
	}

	e2nodeComponentIDC := C.E2nodeComponentID_t{
		present: pr,
		choice:  choiceC,
	}

	return &e2nodeComponentIDC, nil
}

func decodeE2nodeComponentID(e2nodeComponentIDC *C.E2nodeComponentID_t) (*e2ap_ies.E2NodeComponentId, error) {

	e2nodeComponentID := new(e2ap_ies.E2NodeComponentId)

	switch e2nodeComponentIDC.present {
	case C.E2nodeComponentID_PR_e2nodeComponentTypeGNB_CU_UP:
		e2nodeComponentIDstructC, err := decodeE2nodeComponentGnbCuUpIDBytes(e2nodeComponentIDC.choice) //ToDo - Verify if decodeSmthBytes function exists
		if err != nil {
			return nil, fmt.Errorf("decodeE2nodeComponentGnbCuUpIDBytes() %s", err.Error())
		}
		e2nodeComponentID.E2NodeComponentId = &e2ap_ies.E2NodeComponentId_E2NodeComponentTypeGnbCuUp{
			E2NodeComponentTypeGnbCuUp: e2nodeComponentIDstructC,
		}
	case C.E2nodeComponentID_PR_e2nodeComponentTypeGNB_DU:
		e2nodeComponentIDstructC, err := decodeE2nodeComponentGnbDuIDBytes(e2nodeComponentIDC.choice) //ToDo - Verify if decodeSmthBytes function exists
		if err != nil {
			return nil, fmt.Errorf("decodeE2nodeComponentGnbDuIDBytes() %s", err.Error())
		}
		e2nodeComponentID.E2NodeComponentId = &e2ap_ies.E2NodeComponentId_E2NodeComponentTypeGnbDu{
			E2NodeComponentTypeGnbDu: e2nodeComponentIDstructC,
		}
	default:
		return nil, fmt.Errorf("decodeE2nodeComponentID() %v not yet implemented", e2nodeComponentIDC.present)
	}

	return e2nodeComponentID, nil
}

func decodeE2nodeComponentIDBytes(array [8]byte) (*e2ap_ies.E2NodeComponentId, error) {
	e2nodeComponentIDC := (*C.E2nodeComponentID_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeE2nodeComponentID(e2nodeComponentIDC)
}
