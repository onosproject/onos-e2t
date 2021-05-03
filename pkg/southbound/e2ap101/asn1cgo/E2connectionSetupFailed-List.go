// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2connectionSetupFailed-List.h" //ToDo - if there is an anonymous C-struct option, it would require linking additional C-struct file definition (the one above or before)
//#include "E2connectionSetupFailed-Item.h" //ToDo - include correct .h file for corresponding C-struct of "Repeated" field or other anonymous structure defined in .h file
import "C"
import (
	"encoding/binary"
	"fmt"
	e2ap_pdu_contents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
	"unsafe"
)

func xerEncodeE2connectionSetupFailedList(e2connectionSetupFailedList *e2ap_pdu_contents.E2ConnectionSetupFailedList) ([]byte, error) {
	e2connectionSetupFailedListCP, err := newE2connectionSetupFailedList(e2connectionSetupFailedList)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2connectionSetupFailedList() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2connectionSetupFailed_List, unsafe.Pointer(e2connectionSetupFailedListCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2connectionSetupFailedList() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeE2connectionSetupFailedList(e2connectionSetupFailedList *e2ap_pdu_contents.E2ConnectionSetupFailedList) ([]byte, error) {
	e2connectionSetupFailedListCP, err := newE2connectionSetupFailedList(e2connectionSetupFailedList)
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2connectionSetupFailedList() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2connectionSetupFailed_List, unsafe.Pointer(e2connectionSetupFailedListCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2connectionSetupFailedList() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeE2connectionSetupFailedList(bytes []byte) (*e2ap_pdu_contents.E2ConnectionSetupFailedList, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2connectionSetupFailed_List)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2connectionSetupFailedList((*C.E2connectionSetupFailed_List_t)(unsafePtr))
}

func perDecodeE2connectionSetupFailedList(bytes []byte) (*e2ap_pdu_contents.E2ConnectionSetupFailedList, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2connectionSetupFailed_List)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2connectionSetupFailedList((*C.E2connectionSetupFailed_List_t)(unsafePtr))
}

func newE2connectionSetupFailedList(e2connectionSetupFailedList *e2ap_pdu_contents.E2ConnectionSetupFailedList) (*C.E2connectionSetupFailed_List_t, error) {

	e2connectionSetupFailedListC := C.E2connectionSetupFailed_List_t{}
	//for _, ie := range e2connectionSetupFailedList.GetValue() {
	//	ieC, err := newE2connectionSetupFailedItemIes(ie)
	//	if err != nil {
	//		return nil, fmt.Errorf("newE2connectionSetupFailedItemIes() %s", err.Error())
	//	}
	//	if _, err = C.asn_sequence_add(unsafe.Pointer(e2connectionSetupFailedListC), unsafe.Pointer(ieC)); err != nil {
	//		return nil, err
	//	}
	//}

	return &e2connectionSetupFailedListC, nil
}

func decodeE2connectionSetupFailedList(e2connectionSetupFailedListC *C.E2connectionSetupFailed_List_t) (*e2ap_pdu_contents.E2ConnectionSetupFailedList, error) {

	//var ieCount int

	e2connectionSetupFailedList := e2ap_pdu_contents.E2ConnectionSetupFailedList{}

	//ieCount = int(e2connectionSetupFailedListC.list.count)
	//for i := 0; i < ieCount; i++ {
	//	offset := unsafe.Sizeof(unsafe.Pointer(e2connectionSetupFailedListC.list.array)) * uintptr(i)
	//	ieC := *(**C.E2connectionSetupFailed_Item_t)(unsafe.Pointer(uintptr(unsafe.Pointer(e2connectionSetupFailedListC.list.array)) + offset))
	//	ie, err := decodeE2connectionSetupFailedItemIes(ieC)
	//	if err != nil {
	//		return nil, fmt.Errorf("decodeE2connectionSetupFailedItemIes() %s", err.Error())
	//	}
	//	e2connectionSetupFailedList.Value = append(e2connectionSetupFailedList.Value, ie)
	//}

	return &e2connectionSetupFailedList, nil
}

func decodeE2connectionSetupFailedListBytes(array [8]byte) (*e2ap_pdu_contents.E2ConnectionSetupFailedList, error) {
	e2connectionSetupFailedListC := (*C.E2connectionSetupFailed_List_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeE2connectionSetupFailedList(e2connectionSetupFailedListC)
}
