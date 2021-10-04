// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2nodeComponentConfigRemovalAck-List.h"
//#include "ProtocolIE-SingleContainer.h"
import "C"
import (
	"encoding/binary"
	"fmt"
	"unsafe"

	e2ap_pdu_contents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
)

func xerEncodeE2nodeComponentConfigRemovalAckList(e2NodeComponentConfigRemovalAckList *e2ap_pdu_contents.E2NodeComponentConfigRemovalAckList) ([]byte, error) {
	e2NodeComponentConfigRemovalAckListCP, err := newE2nodeComponentConfigRemovalAckList(e2NodeComponentConfigRemovalAckList)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2nodeComponentConfigRemovalAckList() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2nodeComponentConfigRemovalAck_List, unsafe.Pointer(e2NodeComponentConfigRemovalAckListCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2nodeComponentConfigRemovalAckList() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeE2nodeComponentConfigRemovalAckList(e2NodeComponentConfigRemovalAckList *e2ap_pdu_contents.E2NodeComponentConfigRemovalAckList) ([]byte, error) {
	e2NodeComponentConfigRemovalAckListCP, err := newE2nodeComponentConfigRemovalAckList(e2NodeComponentConfigRemovalAckList)
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2nodeComponentConfigRemovalAckList() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2nodeComponentConfigRemovalAck_List, unsafe.Pointer(e2NodeComponentConfigRemovalAckListCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2nodeComponentConfigRemovalAckList() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeE2nodeComponentConfigRemovalAckList(bytes []byte) (*e2ap_pdu_contents.E2NodeComponentConfigRemovalAckList, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2nodeComponentConfigRemovalAck_List)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2nodeComponentConfigRemovalAckList((*C.E2nodeComponentConfigRemovalAck_List_t)(unsafePtr))
}

func perDecodeE2nodeComponentConfigRemovalAckList(bytes []byte) (*e2ap_pdu_contents.E2NodeComponentConfigRemovalAckList, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2nodeComponentConfigRemovalAck_List)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2nodeComponentConfigRemovalAckList((*C.E2nodeComponentConfigRemovalAck_List_t)(unsafePtr))
}

func newE2nodeComponentConfigRemovalAckList(e2curl *e2ap_pdu_contents.E2NodeComponentConfigRemovalAckList) (*C.E2nodeComponentConfigRemovalAck_List_t, error) {

	e2curlC := new(C.E2nodeComponentConfigRemovalAck_List_t)
	for _, ie := range e2curl.GetValue() {
		ieC, err := newE2nodeComponentConfigRemovalAckItemIesSingleContainer(ie)
		if err != nil {
			return nil, fmt.Errorf("newE2nodeComponentConfigRemovalAckItemIesSingleContainer() %s", err.Error())
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(e2curlC), unsafe.Pointer(ieC)); err != nil {
			return nil, err
		}
	}

	return e2curlC, nil
}

func decodeE2nodeComponentConfigRemovalAckList(e2curlC *C.E2nodeComponentConfigRemovalAck_List_t) (*e2ap_pdu_contents.E2NodeComponentConfigRemovalAckList, error) {

	e2curl := e2ap_pdu_contents.E2NodeComponentConfigRemovalAckList{
		Value: make([]*e2ap_pdu_contents.E2NodeComponentConfigRemovalAckItemIes, 0),
	}

	ieCount := int(e2curlC.list.count)
	for i := 0; i < ieCount; i++ {
		offset := unsafe.Sizeof(unsafe.Pointer(e2curlC.list.array)) * uintptr(i)
		ieC := *(**C.ProtocolIE_SingleContainer_1911P13_t)(unsafe.Pointer(uintptr(unsafe.Pointer(e2curlC.list.array)) + offset))
		ie, err := decodeE2nodeComponentConfigRemovalAckItemIesSingleContainer(ieC)
		if err != nil {
			return nil, fmt.Errorf("decodeE2nodeComponentConfigRemovalAckItemIesSingleContainer() %s", err.Error())
		}
		e2curl.Value = append(e2curl.Value, ie)
	}

	return &e2curl, nil
}

func decodeE2nodeComponentConfigRemovalAckListBytes(e2curlC [48]byte) (*e2ap_pdu_contents.E2NodeComponentConfigRemovalAckList, error) {
	array := (**C.struct_ProtocolIE_SingleContainer)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(e2curlC[0:8]))))
	count := C.int(binary.LittleEndian.Uint32(e2curlC[8:12]))
	size := C.int(binary.LittleEndian.Uint32(e2curlC[12:16]))

	rfIDlC := C.E2nodeComponentConfigRemovalAck_List_t{
		list: C.struct___150{
			array: array,
			size:  size,
			count: count,
		},
	}

	return decodeE2nodeComponentConfigRemovalAckList(&rfIDlC)
}
