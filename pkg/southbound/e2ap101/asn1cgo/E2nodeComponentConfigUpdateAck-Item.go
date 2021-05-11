// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2nodeComponentConfigUpdateAck-Item.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	e2ap_pdu_contents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
	"unsafe"
)

func xerEncodeE2nodeComponentConfigUpdateAckItem(e2nodeComponentConfigUpdateAckItem *e2ap_pdu_contents.E2NodeComponentConfigUpdateAckItem) ([]byte, error) {
	e2nodeComponentConfigUpdateAckItemCP, err := newE2nodeComponentConfigUpdateAckItem(e2nodeComponentConfigUpdateAckItem)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2nodeComponentConfigUpdateAckItem() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2nodeComponentConfigUpdateAck_Item, unsafe.Pointer(e2nodeComponentConfigUpdateAckItemCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2nodeComponentConfigUpdateAckItem() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeE2nodeComponentConfigUpdateAckItem(e2nodeComponentConfigUpdateAckItem *e2ap_pdu_contents.E2NodeComponentConfigUpdateAckItem) ([]byte, error) {
	e2nodeComponentConfigUpdateAckItemCP, err := newE2nodeComponentConfigUpdateAckItem(e2nodeComponentConfigUpdateAckItem)
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2nodeComponentConfigUpdateAckItem() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2nodeComponentConfigUpdateAck_Item, unsafe.Pointer(e2nodeComponentConfigUpdateAckItemCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2nodeComponentConfigUpdateAckItem() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeE2nodeComponentConfigUpdateAckItem(bytes []byte) (*e2ap_pdu_contents.E2NodeComponentConfigUpdateAckItem, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2nodeComponentConfigUpdateAck_Item)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2nodeComponentConfigUpdateAckItem((*C.E2nodeComponentConfigUpdateAck_Item_t)(unsafePtr))
}

func perDecodeE2nodeComponentConfigUpdateAckItem(bytes []byte) (*e2ap_pdu_contents.E2NodeComponentConfigUpdateAckItem, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2nodeComponentConfigUpdateAck_Item)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2nodeComponentConfigUpdateAckItem((*C.E2nodeComponentConfigUpdateAck_Item_t)(unsafePtr))
}

func newE2nodeComponentConfigUpdateAckItem(e2nodeComponentConfigUpdateAckItem *e2ap_pdu_contents.E2NodeComponentConfigUpdateAckItem) (*C.E2nodeComponentConfigUpdateAck_Item_t, error) {

	var err error
	e2nodeComponentConfigUpdateAckItemC := C.E2nodeComponentConfigUpdateAck_Item_t{}

	e2nodeComponentTypeC, err := newE2nodeComponentType(&e2nodeComponentConfigUpdateAckItem.E2NodeComponentType)
	if err != nil {
		return nil, fmt.Errorf("newE2nodeComponentType() %s", err.Error())
	}

	e2nodeComponentIDC, err := newE2nodeComponentID(e2nodeComponentConfigUpdateAckItem.E2NodeComponentId)
	if err != nil {
		return nil, fmt.Errorf("newE2nodeComponentID() %s", err.Error())
	}

	e2nodeComponentConfigUpdateAckC, err := newE2nodeComponentConfigUpdateAck(e2nodeComponentConfigUpdateAckItem.E2NodeComponentConfigUpdateAck)
	if err != nil {
		return nil, fmt.Errorf("newE2nodeComponentConfigUpdateAck() %s", err.Error())
	}

	//ToDo - check whether pointers passed correctly with regard to C-struct's definition .h file
	e2nodeComponentConfigUpdateAckItemC.e2nodeComponentType = *e2nodeComponentTypeC
	e2nodeComponentConfigUpdateAckItemC.e2nodeComponentID = e2nodeComponentIDC
	e2nodeComponentConfigUpdateAckItemC.e2nodeComponentConfigUpdateAck = *e2nodeComponentConfigUpdateAckC

	return &e2nodeComponentConfigUpdateAckItemC, nil
}

func decodeE2nodeComponentConfigUpdateAckItem(e2nodeComponentConfigUpdateAckItemC *C.E2nodeComponentConfigUpdateAck_Item_t) (*e2ap_pdu_contents.E2NodeComponentConfigUpdateAckItem, error) {

	var err error
	e2nodeComponentConfigUpdateAckItem := e2ap_pdu_contents.E2NodeComponentConfigUpdateAckItem{}

	e2NodeComponentType, err := decodeE2nodeComponentType(&e2nodeComponentConfigUpdateAckItemC.e2nodeComponentType)
	if err != nil {
		return nil, fmt.Errorf("decodeE2nodeComponentType() %s", err.Error())
	}
	e2nodeComponentConfigUpdateAckItem.E2NodeComponentType = *e2NodeComponentType

	e2nodeComponentConfigUpdateAckItem.E2NodeComponentId, err = decodeE2nodeComponentID(e2nodeComponentConfigUpdateAckItemC.e2nodeComponentID)
	if err != nil {
		return nil, fmt.Errorf("decodeE2nodeComponentID() %s", err.Error())
	}

	e2nodeComponentConfigUpdateAckItem.E2NodeComponentConfigUpdateAck, err = decodeE2nodeComponentConfigUpdateAck(&e2nodeComponentConfigUpdateAckItemC.e2nodeComponentConfigUpdateAck)
	if err != nil {
		return nil, fmt.Errorf("decodeE2nodeComponentConfigUpdateAck() %s", err.Error())
	}

	return &e2nodeComponentConfigUpdateAckItem, nil
}

func decodeE2nodeComponentConfigUpdateAckItemBytes(bytes [80]byte) (*e2ap_pdu_contents.E2NodeComponentConfigUpdateAckItem, error) {

	e2nccuaiC := C.E2nodeComponentConfigUpdateAck_Item_t{
		e2nodeComponentType: C.long(binary.LittleEndian.Uint64(bytes[0:8])),
		e2nodeComponentID:   (*C.struct_E2nodeComponentID)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(bytes[8:16])))),
		e2nodeComponentConfigUpdateAck: C.E2nodeComponentConfigUpdateAck_t{
			updateOutcome: C.long(binary.LittleEndian.Uint64(bytes[16:24])),
			failureCause:  (*C.struct_Cause)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(bytes[24:])))),
		},
	}

	return decodeE2nodeComponentConfigUpdateAckItem(&e2nccuaiC)
}
