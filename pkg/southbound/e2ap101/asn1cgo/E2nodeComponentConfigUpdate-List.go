// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2nodeComponentConfigUpdate-List.h"
//#include "E2nodeComponentConfigUpdate-Item.h"
import "C"
import (
	"encoding/binary"
	"fmt"
	e2ap_pdu_contents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
	"unsafe"
)

func xerEncodeE2nodeComponentConfigUpdateList(e2nodeComponentConfigUpdateList *e2ap_pdu_contents.E2NodeComponentConfigUpdateList) ([]byte, error) {
	e2nodeComponentConfigUpdateListCP, err := newE2nodeComponentConfigUpdateList(e2nodeComponentConfigUpdateList)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2nodeComponentConfigUpdateList() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2nodeComponentConfigUpdate_List, unsafe.Pointer(e2nodeComponentConfigUpdateListCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2nodeComponentConfigUpdateList() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeE2nodeComponentConfigUpdateList(e2nodeComponentConfigUpdateList *e2ap_pdu_contents.E2NodeComponentConfigUpdateList) ([]byte, error) {
	e2nodeComponentConfigUpdateListCP, err := newE2nodeComponentConfigUpdateList(e2nodeComponentConfigUpdateList)
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2nodeComponentConfigUpdateList() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2nodeComponentConfigUpdate_List, unsafe.Pointer(e2nodeComponentConfigUpdateListCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2nodeComponentConfigUpdateList() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeE2nodeComponentConfigUpdateList(bytes []byte) (*e2ap_pdu_contents.E2NodeComponentConfigUpdateList, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2nodeComponentConfigUpdate_List)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2nodeComponentConfigUpdateList((*C.E2nodeComponentConfigUpdate_List_t)(unsafePtr))
}

func perDecodeE2nodeComponentConfigUpdateList(bytes []byte) (*e2ap_pdu_contents.E2NodeComponentConfigUpdateList, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2nodeComponentConfigUpdate_List)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2nodeComponentConfigUpdateList((*C.E2nodeComponentConfigUpdate_List_t)(unsafePtr))
}

func newE2nodeComponentConfigUpdateList(e2nodeComponentConfigUpdateList *e2ap_pdu_contents.E2NodeComponentConfigUpdateList) (*C.E2nodeComponentConfigUpdate_List_t, error) {

	e2nodeComponentConfigUpdateListC := C.E2nodeComponentConfigUpdate_List_t{}
	//for _, ie := range e2nodeComponentConfigUpdateList.GetValue() {
	//	ieC, err := newE2nodeComponentConfigUpdateItemIes(ie)
	//	if err != nil {
	//		return nil, fmt.Errorf("newE2nodeComponentConfigUpdateItemIes() %s", err.Error())
	//	}
	//	if _, err = C.asn_sequence_add(unsafe.Pointer(e2nodeComponentConfigUpdateListC), unsafe.Pointer(ieC)); err != nil {
	//		return nil, err
	//	}
	//}

	return &e2nodeComponentConfigUpdateListC, nil
}

func decodeE2nodeComponentConfigUpdateList(e2nodeComponentConfigUpdateListC *C.E2nodeComponentConfigUpdate_List_t) (*e2ap_pdu_contents.E2NodeComponentConfigUpdateList, error) {

	//var ieCount int

	e2nodeComponentConfigUpdateList := e2ap_pdu_contents.E2NodeComponentConfigUpdateList{}

	//ieCount = int(e2nodeComponentConfigUpdateListC.list.count)
	//for i := 0; i < ieCount; i++ {
	//	offset := unsafe.Sizeof(unsafe.Pointer(e2nodeComponentConfigUpdateListC.list.array)) * uintptr(i)
	//	ieC := *(**C.E2nodeComponentConfigUpdate_Item_t)(unsafe.Pointer(uintptr(unsafe.Pointer(e2nodeComponentConfigUpdateListC.list.array)) + offset))
	//	ie, err := decodeE2nodeComponentConfigUpdateItemIes(ieC)
	//	if err != nil {
	//		return nil, fmt.Errorf("decodeE2nodeComponentConfigUpdateItemIes() %s", err.Error())
	//	}
	//	e2nodeComponentConfigUpdateList.Value = append(e2nodeComponentConfigUpdateList.Value, ie)
	//}

	return &e2nodeComponentConfigUpdateList, nil
}

func decodeE2nodeComponentConfigUpdateListBytes(array [8]byte) (*e2ap_pdu_contents.E2NodeComponentConfigUpdateList, error) {
	e2nodeComponentConfigUpdateListC := (*C.E2nodeComponentConfigUpdate_List_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeE2nodeComponentConfigUpdateList(e2nodeComponentConfigUpdateListC)
}
