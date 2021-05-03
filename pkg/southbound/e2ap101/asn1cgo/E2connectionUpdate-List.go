// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2connectionUpdate-List.h"
//#include "E2connectionUpdate-Item.h"
import "C"
import (
	"encoding/binary"
	"fmt"
	e2ap_pdu_contents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
	"unsafe"
)

func xerEncodeE2connectionUpdateList(e2connectionUpdateList *e2ap_pdu_contents.E2ConnectionUpdateList) ([]byte, error) {
	e2connectionUpdateListCP, err := newE2connectionUpdateList(e2connectionUpdateList)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2connectionUpdateList() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2connectionUpdate_List, unsafe.Pointer(e2connectionUpdateListCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2connectionUpdateList() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeE2connectionUpdateList(e2connectionUpdateList *e2ap_pdu_contents.E2ConnectionUpdateList) ([]byte, error) {
	e2connectionUpdateListCP, err := newE2connectionUpdateList(e2connectionUpdateList)
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2connectionUpdateList() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2connectionUpdate_List, unsafe.Pointer(e2connectionUpdateListCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2connectionUpdateList() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeE2connectionUpdateList(bytes []byte) (*e2ap_pdu_contents.E2ConnectionUpdateList, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2connectionUpdate_List)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2connectionUpdateList((*C.E2connectionUpdate_List_t)(unsafePtr))
}

func perDecodeE2connectionUpdateList(bytes []byte) (*e2ap_pdu_contents.E2ConnectionUpdateList, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2connectionUpdate_List)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2connectionUpdateList((*C.E2connectionUpdate_List_t)(unsafePtr))
}

func newE2connectionUpdateList(e2connectionUpdateList *e2ap_pdu_contents.E2ConnectionUpdateList) (*C.E2connectionUpdate_List_t, error) {

	e2connectionUpdateListC := C.E2connectionUpdate_List_t{}
	//for _, ie := range e2connectionUpdateList.GetValue() {     //ToDo - Verify if GetSmth() function is called correctly
	//	ieC, err := newE2connectionUpdateItemIes(ie)
	//	if err != nil {
	//		return nil, fmt.Errorf("newE2connectionUpdateItemIes() %s", err.Error())
	//	}
	//	if _, err = C.asn_sequence_add(unsafe.Pointer(e2connectionUpdateListC), unsafe.Pointer(ieC)); err != nil {
	//		return nil, err
	//	}
	//}

	return &e2connectionUpdateListC, nil
}

func decodeE2connectionUpdateList(e2connectionUpdateListC *C.E2connectionUpdate_List_t) (*e2ap_pdu_contents.E2ConnectionUpdateList, error) {

	//var ieCount int

	e2connectionUpdateList := e2ap_pdu_contents.E2ConnectionUpdateList{}

	//ieCount = int(e2connectionUpdateListC.list.count)
	//for i := 0; i < ieCount; i++ {
	//	offset := unsafe.Sizeof(unsafe.Pointer(e2connectionUpdateListC.list.array)) * uintptr(i)
	//	ieC := *(**C.E2connectionUpdate_Item_t)(unsafe.Pointer(uintptr(unsafe.Pointer(e2connectionUpdateListC.list.array)) + offset))
	//	ie, err := decodeE2connectionUpdateItemIes(ieC)
	//	if err != nil {
	//		return nil, fmt.Errorf("decodeE2connectionUpdateItemIes() %s", err.Error())
	//	}
	//	e2connectionUpdateList.Value = append(e2connectionUpdateList.Value, ie)
	//}

	return &e2connectionUpdateList, nil
}

func decodeE2connectionUpdateListBytes(array [8]byte) (*e2ap_pdu_contents.E2ConnectionUpdateList, error) {
	e2connectionUpdateListC := (*C.E2connectionUpdate_List_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeE2connectionUpdateList(e2connectionUpdateListC)
}
