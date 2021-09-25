// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2nodeComponentConfigRemoval-Item.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	"unsafe"

	e2ap_pdu_contents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
)

func xerEncodeE2nodeComponentConfigRemovalItem(e2nodeComponentConfigRemovalItem *e2ap_pdu_contents.E2NodeComponentConfigRemovalItem) ([]byte, error) {
	e2nodeComponentConfigRemovalItemCP, err := newE2nodeComponentConfigRemovalItem(e2nodeComponentConfigRemovalItem)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2nodeComponentConfigRemovalItem() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2nodeComponentConfigRemoval_Item, unsafe.Pointer(e2nodeComponentConfigRemovalItemCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2nodeComponentConfigRemovalItem() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeE2nodeComponentConfigRemovalItem(e2nodeComponentConfigRemovalItem *e2ap_pdu_contents.E2NodeComponentConfigRemovalItem) ([]byte, error) {
	e2nodeComponentConfigRemovalItemCP, err := newE2nodeComponentConfigRemovalItem(e2nodeComponentConfigRemovalItem)
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2nodeComponentConfigRemovalItem() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2nodeComponentConfigRemoval_Item, unsafe.Pointer(e2nodeComponentConfigRemovalItemCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2nodeComponentConfigRemovalItem() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeE2nodeComponentConfigRemovalItem(bytes []byte) (*e2ap_pdu_contents.E2NodeComponentConfigRemovalItem, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2nodeComponentConfigRemoval_Item)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2nodeComponentConfigRemovalItem((*C.E2nodeComponentConfigRemoval_Item_t)(unsafePtr))
}

func perDecodeE2nodeComponentConfigRemovalItem(bytes []byte) (*e2ap_pdu_contents.E2NodeComponentConfigRemovalItem, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2nodeComponentConfigRemoval_Item)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2nodeComponentConfigRemovalItem((*C.E2nodeComponentConfigRemoval_Item_t)(unsafePtr))
}

func newE2nodeComponentConfigRemovalItem(e2nodeComponentConfigRemovalItem *e2ap_pdu_contents.E2NodeComponentConfigRemovalItem) (*C.E2nodeComponentConfigRemoval_Item_t, error) {

	var err error
	e2nodeComponentConfigRemovalItemC := C.E2nodeComponentConfigRemoval_Item_t{}

	e2nodeComponentInterfaceTypeC, err := newE2nodeComponentInterfaceType(&e2nodeComponentConfigRemovalItem.E2NodeComponentInterfaceType)
	if err != nil {
		return nil, fmt.Errorf("newE2nodeComponentInterfaceType() %s", err.Error())
	}

	e2nodeComponentConfigRemovalItemC.e2nodeComponentID, err = newE2nodeComponentID(e2nodeComponentConfigRemovalItem.E2NodeComponentId)
	if err != nil {
		return nil, fmt.Errorf("newE2nodeComponentID() %s", err.Error())
	}

	e2nodeComponentConfigRemovalItemC.e2nodeComponentInterfaceType = *e2nodeComponentInterfaceTypeC

	return &e2nodeComponentConfigRemovalItemC, nil
}

func decodeE2nodeComponentConfigRemovalItem(e2nodeComponentConfigRemovalItemC *C.E2nodeComponentConfigRemoval_Item_t) (*e2ap_pdu_contents.E2NodeComponentConfigRemovalItem, error) {

	var err error
	e2nodeComponentConfigRemovalItem := e2ap_pdu_contents.E2NodeComponentConfigRemovalItem{}

	componentType, err := decodeE2nodeComponentType(&e2nodeComponentConfigRemovalItemC.e2nodeComponentInterfaceType)
	if err != nil {
		return nil, fmt.Errorf("decodeE2nodeComponentType() %s", err.Error())
	}
	e2nodeComponentConfigRemovalItem.E2NodeComponentInterfaceType = *componentType

	if e2nodeComponentConfigRemovalItemC.e2nodeComponentID != nil {
		e2nodeComponentConfigRemovalItem.E2NodeComponentId, err = decodeE2nodeComponentID(e2nodeComponentConfigRemovalItemC.e2nodeComponentID)
		if err != nil {
			return nil, fmt.Errorf("decodeE2nodeComponentID() %s", err.Error())
		}
	}

	return &e2nodeComponentConfigRemovalItem, nil
}

func decodeE2nodeComponentConfigRemovalItemBytes(bytes [80]byte) (*e2ap_pdu_contents.E2NodeComponentConfigRemovalItem, error) {

	//e2nccuC := C.E2nodeComponentConfigAdditionAck_t{
	//	present: C.E2nodeComponentConfigAdditionAck_PR(binary.LittleEndian.Uint64(bytes[40:48])),
	//}
	//copy(e2nccuC.choice[:8], bytes[48:56])

	e2nodeComponentConfigRemovalItemC := C.E2nodeComponentConfigRemoval_Item_t{
		e2nodeComponentInterfaceType:    C.long(binary.LittleEndian.Uint64(bytes[0:8])),
		e2nodeComponentID:               (*C.struct_E2nodeComponentID)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(bytes[8:16])))),
	}

	return decodeE2nodeComponentConfigRemovalItem(&e2nodeComponentConfigRemovalItemC)
}
