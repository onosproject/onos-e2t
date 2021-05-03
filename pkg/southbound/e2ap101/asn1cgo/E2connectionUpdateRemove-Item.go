// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2connectionUpdateRemove-Item.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	e2ap_pdu_contents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
	"unsafe"
)

func xerEncodeE2connectionUpdateRemoveItem(e2connectionUpdateRemoveItem *e2ap_pdu_contents.E2ConnectionUpdateRemoveItem) ([]byte, error) {
	e2connectionUpdateRemoveItemCP, err := newE2connectionUpdateRemoveItem(e2connectionUpdateRemoveItem)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2connectionUpdateRemoveItem() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2connectionUpdateRemove_Item, unsafe.Pointer(e2connectionUpdateRemoveItemCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2connectionUpdateRemoveItem() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeE2connectionUpdateRemoveItem(e2connectionUpdateRemoveItem *e2ap_pdu_contents.E2ConnectionUpdateRemoveItem) ([]byte, error) {
	e2connectionUpdateRemoveItemCP, err := newE2connectionUpdateRemoveItem(e2connectionUpdateRemoveItem)
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2connectionUpdateRemoveItem() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2connectionUpdateRemove_Item, unsafe.Pointer(e2connectionUpdateRemoveItemCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2connectionUpdateRemoveItem() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeE2connectionUpdateRemoveItem(bytes []byte) (*e2ap_pdu_contents.E2ConnectionUpdateRemoveItem, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2connectionUpdateRemove_Item)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2connectionUpdateRemoveItem((*C.E2connectionUpdateRemove_Item_t)(unsafePtr))
}

func perDecodeE2connectionUpdateRemoveItem(bytes []byte) (*e2ap_pdu_contents.E2ConnectionUpdateRemoveItem, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2connectionUpdateRemove_Item)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2connectionUpdateRemoveItem((*C.E2connectionUpdateRemove_Item_t)(unsafePtr))
}

func newE2connectionUpdateRemoveItem(e2connectionUpdateRemoveItem *e2ap_pdu_contents.E2ConnectionUpdateRemoveItem) (*C.E2connectionUpdateRemove_Item_t, error) {

	var err error
	e2connectionUpdateRemoveItemC := C.E2connectionUpdateRemove_Item_t{}

	tnlInformationC, err := newTnlinformation(e2connectionUpdateRemoveItem.TnlInformation)
	if err != nil {
		return nil, fmt.Errorf("newTnlinformation() %s", err.Error())
	}

	//ToDo - check whether pointers passed correctly with regard to C-struct's definition .h file
	e2connectionUpdateRemoveItemC.tnlInformation = *tnlInformationC

	return &e2connectionUpdateRemoveItemC, nil
}

func decodeE2connectionUpdateRemoveItem(e2connectionUpdateRemoveItemC *C.E2connectionUpdateRemove_Item_t) (*e2ap_pdu_contents.E2ConnectionUpdateRemoveItem, error) {

	var err error
	e2connectionUpdateRemoveItem := e2ap_pdu_contents.E2ConnectionUpdateRemoveItem{
		//ToDo - check whether pointers passed correctly with regard to Protobuf's definition
		//TnlInformation: tnlInformation,
	}

	e2connectionUpdateRemoveItem.TnlInformation, err = decodeTnlinformation(&e2connectionUpdateRemoveItemC.tnlInformation)
	if err != nil {
		return nil, fmt.Errorf("decodeTnlinformation() %s", err.Error())
	}

	return &e2connectionUpdateRemoveItem, nil
}

func decodeE2connectionUpdateRemoveItemBytes(array [8]byte) (*e2ap_pdu_contents.E2ConnectionUpdateRemoveItem, error) {
	e2connectionUpdateRemoveItemC := (*C.E2connectionUpdateRemove_Item_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeE2connectionUpdateRemoveItem(e2connectionUpdateRemoveItemC)
}
