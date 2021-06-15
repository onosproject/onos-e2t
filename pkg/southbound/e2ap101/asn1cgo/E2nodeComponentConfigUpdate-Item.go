// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2nodeComponentConfigUpdate-Item.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	e2ap_pdu_contents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
	"unsafe"
)

func xerEncodeE2nodeComponentConfigUpdateItem(e2nodeComponentConfigUpdateItem *e2ap_pdu_contents.E2NodeComponentConfigUpdateItem) ([]byte, error) {
	e2nodeComponentConfigUpdateItemCP, err := newE2nodeComponentConfigUpdateItem(e2nodeComponentConfigUpdateItem)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2nodeComponentConfigUpdateItem() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2nodeComponentConfigUpdate_Item, unsafe.Pointer(e2nodeComponentConfigUpdateItemCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2nodeComponentConfigUpdateItem() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeE2nodeComponentConfigUpdateItem(e2nodeComponentConfigUpdateItem *e2ap_pdu_contents.E2NodeComponentConfigUpdateItem) ([]byte, error) {
	e2nodeComponentConfigUpdateItemCP, err := newE2nodeComponentConfigUpdateItem(e2nodeComponentConfigUpdateItem)
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2nodeComponentConfigUpdateItem() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2nodeComponentConfigUpdate_Item, unsafe.Pointer(e2nodeComponentConfigUpdateItemCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2nodeComponentConfigUpdateItem() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeE2nodeComponentConfigUpdateItem(bytes []byte) (*e2ap_pdu_contents.E2NodeComponentConfigUpdateItem, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2nodeComponentConfigUpdate_Item)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2nodeComponentConfigUpdateItem((*C.E2nodeComponentConfigUpdate_Item_t)(unsafePtr))
}

func perDecodeE2nodeComponentConfigUpdateItem(bytes []byte) (*e2ap_pdu_contents.E2NodeComponentConfigUpdateItem, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2nodeComponentConfigUpdate_Item)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2nodeComponentConfigUpdateItem((*C.E2nodeComponentConfigUpdate_Item_t)(unsafePtr))
}

func newE2nodeComponentConfigUpdateItem(e2nodeComponentConfigUpdateItem *e2ap_pdu_contents.E2NodeComponentConfigUpdateItem) (*C.E2nodeComponentConfigUpdate_Item_t, error) {

	var err error
	e2nodeComponentConfigUpdateItemC := C.E2nodeComponentConfigUpdate_Item_t{}

	e2nodeComponentTypeC, err := newE2nodeComponentType(&e2nodeComponentConfigUpdateItem.E2NodeComponentType)
	if err != nil {
		return nil, fmt.Errorf("newE2nodeComponentType() %s", err.Error())
	}

	//optional structure
	if e2nodeComponentConfigUpdateItem.GetE2NodeComponentId() != nil {
		e2nodeComponentConfigUpdateItemC.e2nodeComponentID, err = newE2nodeComponentID(e2nodeComponentConfigUpdateItem.E2NodeComponentId)
		if err != nil {
			return nil, fmt.Errorf("newE2nodeComponentID() %s", err.Error())
		}
	}
	e2nodeComponentConfigUpdateC, err := newE2nodeComponentConfigUpdate(e2nodeComponentConfigUpdateItem.E2NodeComponentConfigUpdate)
	if err != nil {
		return nil, fmt.Errorf("newE2nodeComponentConfigUpdate() %s", err.Error())
	}

	e2nodeComponentConfigUpdateItemC.e2nodeComponentType = *e2nodeComponentTypeC
	e2nodeComponentConfigUpdateItemC.e2nodeComponentConfigUpdate = *e2nodeComponentConfigUpdateC

	return &e2nodeComponentConfigUpdateItemC, nil
}

func decodeE2nodeComponentConfigUpdateItem(e2nodeComponentConfigUpdateItemC *C.E2nodeComponentConfigUpdate_Item_t) (*e2ap_pdu_contents.E2NodeComponentConfigUpdateItem, error) {

	var err error
	e2nodeComponentConfigUpdateItem := e2ap_pdu_contents.E2NodeComponentConfigUpdateItem{}

	componentType, err := decodeE2nodeComponentType(&e2nodeComponentConfigUpdateItemC.e2nodeComponentType)
	if err != nil {
		return nil, fmt.Errorf("decodeE2nodeComponentType() %s", err.Error())
	}
	e2nodeComponentConfigUpdateItem.E2NodeComponentType = *componentType

	if e2nodeComponentConfigUpdateItemC.e2nodeComponentID != nil{
		e2nodeComponentConfigUpdateItem.E2NodeComponentId, err = decodeE2nodeComponentID(e2nodeComponentConfigUpdateItemC.e2nodeComponentID)
		if err != nil{
		return nil, fmt.Errorf("decodeE2nodeComponentID() %s", err.Error())
	}
	}
	e2nodeComponentConfigUpdateItem.E2NodeComponentConfigUpdate, err = decodeE2nodeComponentConfigUpdate(&e2nodeComponentConfigUpdateItemC.e2nodeComponentConfigUpdate)
	if err != nil {
		return nil, fmt.Errorf("decodeE2nodeComponentConfigUpdate() %s", err.Error())
	}

	return &e2nodeComponentConfigUpdateItem, nil
}

func decodeE2nodeComponentConfigUpdateItemBytes(bytes [80]byte) (*e2ap_pdu_contents.E2NodeComponentConfigUpdateItem, error) {

	e2nccuC := C.E2nodeComponentConfigUpdate_t{
		present: C.E2nodeComponentConfigUpdate_PR(binary.LittleEndian.Uint64(bytes[16:24])),
	}
	copy(e2nccuC.choice[:8], bytes[24:32])

	e2nodeComponentConfigUpdateItemC := C.E2nodeComponentConfigUpdate_Item_t{
		e2nodeComponentType:         C.long(binary.LittleEndian.Uint64(bytes[0:8])),
		e2nodeComponentID:           (*C.struct_E2nodeComponentID)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(bytes[8:16])))),
		e2nodeComponentConfigUpdate: e2nccuC,
	}

	return decodeE2nodeComponentConfigUpdateItem(&e2nodeComponentConfigUpdateItemC)
}
