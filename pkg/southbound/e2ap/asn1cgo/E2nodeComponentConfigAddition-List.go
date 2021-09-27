// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2nodeComponentConfigAddition-List.h"
//#include "ProtocolIE-SingleContainer.h"
import "C"
import (
	"encoding/binary"
	"fmt"
	"unsafe"

	e2ap_pdu_contents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
)

func xerEncodeE2nodeComponentConfigAdditionList(e2NodeComponentConfigAdditionList *e2ap_pdu_contents.E2NodeComponentConfigAdditionList) ([]byte, error) {
	e2NodeComponentConfigAdditionListCP, err := newE2nodeComponentConfigAdditionList(e2NodeComponentConfigAdditionList)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2nodeComponentConfigAdditionList() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2nodeComponentConfigAddition_List, unsafe.Pointer(e2NodeComponentConfigAdditionListCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2componentConfigAdditionList() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeE2nodeComponentConfigAdditionList(e2NodeComponentConfigAdditionList *e2ap_pdu_contents.E2NodeComponentConfigAdditionList) ([]byte, error) {
	e2NodeComponentConfigAdditionListCP, err := newE2nodeComponentConfigAdditionList(e2NodeComponentConfigAdditionList)
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2componentConfigAdditionList() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2nodeComponentConfigAddition_List, unsafe.Pointer(e2NodeComponentConfigAdditionListCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2nodeComponentConfigAdditionList() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeE2nodeComponentConfigAdditionList(bytes []byte) (*e2ap_pdu_contents.E2NodeComponentConfigAdditionList, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2nodeComponentConfigAddition_List)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2nodeComponentConfigAdditionList((*C.E2nodeComponentConfigAddition_List_t)(unsafePtr))
}

func perDecodeE2nodeComponentConfigAdditionList(bytes []byte) (*e2ap_pdu_contents.E2NodeComponentConfigAdditionList, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2nodeComponentConfigAddition_List)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2nodeComponentConfigAdditionList((*C.E2nodeComponentConfigAddition_List_t)(unsafePtr))
}

func newE2nodeComponentConfigAdditionList(e2curl *e2ap_pdu_contents.E2NodeComponentConfigAdditionList) (*C.E2nodeComponentConfigAddition_List_t, error) {

	e2curlC := new(C.E2nodeComponentConfigAddition_List_t)
	for _, ie := range e2curl.GetValue() {
		ieC, err := newE2nodeComponentConfigAdditionItemIesSingleContainer(ie)
		if err != nil {
			return nil, fmt.Errorf("newE2nodeComponentConfigAdditionItemIesSingleContainer() %s", err.Error())
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(e2curlC), unsafe.Pointer(ieC)); err != nil {
			return nil, err
		}
	}

	return e2curlC, nil
}

func decodeE2nodeComponentConfigAdditionList(e2curlC *C.E2nodeComponentConfigAddition_List_t) (*e2ap_pdu_contents.E2NodeComponentConfigAdditionList, error) {

	e2curl := e2ap_pdu_contents.E2NodeComponentConfigAdditionList{
		Value: make([]*e2ap_pdu_contents.E2NodeComponentConfigAdditionItemIes, 0),
	}

	ieCount := int(e2curlC.list.count)
	for i := 0; i < ieCount; i++ {
		offset := unsafe.Sizeof(unsafe.Pointer(e2curlC.list.array)) * uintptr(i)
		ieC := *(**C.ProtocolIE_SingleContainer_1911P7_t)(unsafe.Pointer(uintptr(unsafe.Pointer(e2curlC.list.array)) + offset))
		ie, err := decodeE2nodeComponentConfigAdditionItemIesSingleContainer(ieC)
		if err != nil {
			return nil, fmt.Errorf("decodeE2nodeComponentConfigAdditionItemIesSingleContainer() %s", err.Error())
		}
		e2curl.Value = append(e2curl.Value, ie)
	}

	return &e2curl, nil
}

func decodeE2nodeComponentConfigAdditionListBytes(e2curlC [48]byte) (*e2ap_pdu_contents.E2NodeComponentConfigAdditionList, error) {
	array := (**C.struct_ProtocolIE_SingleContainer)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(e2curlC[0:8]))))
	count := C.int(binary.LittleEndian.Uint32(e2curlC[8:12]))
	size := C.int(binary.LittleEndian.Uint32(e2curlC[12:16]))

	rfIDlC := C.E2nodeComponentConfigAddition_List_t{
		list: C.struct___105{
			array: array,
			size:  size,
			count: count,
		},
	}

	return decodeE2nodeComponentConfigAdditionList(&rfIDlC)
}
