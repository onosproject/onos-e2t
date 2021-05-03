// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2nodeComponentConfigUpdateAck-List.h"
//#include "E2nodeComponentConfigUpdateAck-Item.h"
import "C"
import (
	"encoding/binary"
	"fmt"
	e2ap_pdu_contents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
	"unsafe"
)

func xerEncodeE2nodeComponentConfigUpdateAckList(e2nodeComponentConfigUpdateAckList *e2ap_pdu_contents.E2NodeComponentConfigUpdateAckList) ([]byte, error) {
	e2nodeComponentConfigUpdateAckListCP, err := newE2nodeComponentConfigUpdateAckList(e2nodeComponentConfigUpdateAckList)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2nodeComponentConfigUpdateAckList() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2nodeComponentConfigUpdateAck_List, unsafe.Pointer(e2nodeComponentConfigUpdateAckListCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2nodeComponentConfigUpdateAckList() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeE2nodeComponentConfigUpdateAckList(e2nodeComponentConfigUpdateAckList *e2ap_pdu_contents.E2NodeComponentConfigUpdateAckList) ([]byte, error) {
	e2nodeComponentConfigUpdateAckListCP, err := newE2nodeComponentConfigUpdateAckList(e2nodeComponentConfigUpdateAckList)
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2nodeComponentConfigUpdateAckList() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2nodeComponentConfigUpdateAck_List, unsafe.Pointer(e2nodeComponentConfigUpdateAckListCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2nodeComponentConfigUpdateAckList() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeE2nodeComponentConfigUpdateAckList(bytes []byte) (*e2ap_pdu_contents.E2NodeComponentConfigUpdateAckList, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2nodeComponentConfigUpdateAck_List)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2nodeComponentConfigUpdateAckList((*C.E2nodeComponentConfigUpdateAck_List_t)(unsafePtr))
}

func perDecodeE2nodeComponentConfigUpdateAckList(bytes []byte) (*e2ap_pdu_contents.E2NodeComponentConfigUpdateAckList, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2nodeComponentConfigUpdateAck_List)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2nodeComponentConfigUpdateAckList((*C.E2nodeComponentConfigUpdateAck_List_t)(unsafePtr))
}

func newE2nodeComponentConfigUpdateAckList(e2nodeComponentConfigUpdateAckList *e2ap_pdu_contents.E2NodeComponentConfigUpdateAckList) (*C.E2nodeComponentConfigUpdateAck_List_t, error) {

	e2nodeComponentConfigUpdateAckListC := C.E2nodeComponentConfigUpdateAck_List_t{}
	//for _, ie := range e2nodeComponentConfigUpdateAckList.GetValue() {
	//	ieC, err := newE2nodeComponentConfigUpdateAckItemIE(ie)
	//	if err != nil {
	//		return nil, fmt.Errorf("newE2nodeComponentConfigUpdateAckItemIes() %s", err.Error())
	//	}
	//	if _, err = C.asn_sequence_add(unsafe.Pointer(e2nodeComponentConfigUpdateAckListC), unsafe.Pointer(ieC)); err != nil {
	//		return nil, err
	//	}
	//}

	return &e2nodeComponentConfigUpdateAckListC, nil
}

func decodeE2nodeComponentConfigUpdateAckList(e2nodeComponentConfigUpdateAckListC *C.E2nodeComponentConfigUpdateAck_List_t) (*e2ap_pdu_contents.E2NodeComponentConfigUpdateAckList, error) {

	//var ieCount int
	e2nodeComponentConfigUpdateAckList := e2ap_pdu_contents.E2NodeComponentConfigUpdateAckList{}

	//ieCount = int(e2nodeComponentConfigUpdateAckListC.list.count)
	//for i := 0; i < ieCount; i++ {
	//	offset := unsafe.Sizeof(unsafe.Pointer(e2nodeComponentConfigUpdateAckListC.list.array)) * uintptr(i)
	//	ieC := *(**C.E2nodeComponentConfigUpdateAck_Item_t)(unsafe.Pointer(uintptr(unsafe.Pointer(e2nodeComponentConfigUpdateAckListC.list.array)) + offset))
	//	ie, err := decodeE2nodeComponentConfigUpdateAckItemIE(ieC)
	//	if err != nil {
	//		return nil, fmt.Errorf("decodeE2nodeComponentConfigUpdateAckItemIes() %s", err.Error())
	//	}
	//	e2nodeComponentConfigUpdateAckList.Value = append(e2nodeComponentConfigUpdateAckList.Value, ie)
	//}

	return &e2nodeComponentConfigUpdateAckList, nil
}

func decodeE2nodeComponentConfigUpdateAckListBytes(array [8]byte) (*e2ap_pdu_contents.E2NodeComponentConfigUpdateAckList, error) {
	e2nodeComponentConfigUpdateAckListC := (*C.E2nodeComponentConfigUpdateAck_List_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeE2nodeComponentConfigUpdateAckList(e2nodeComponentConfigUpdateAckListC)
}
