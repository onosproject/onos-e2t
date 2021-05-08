// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2connectionUpdateRemove-List.h"
//#include "ProtocolIE-SingleContainer.h"
import "C"
import (
	"encoding/binary"
	"fmt"
	e2ap_pdu_contents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
	"unsafe"
)

func newE2connectionUpdateRemoveList(e2curl *e2ap_pdu_contents.E2ConnectionUpdateRemoveList) (*C.E2connectionUpdateRemove_List_t, error) {

	e2curlC := C.E2nodeComponentConfigUpdate_List_t{}
	for _, ie := range e2curl.GetValue() {
		ieC, err := newE2connectionUpdateRemoveIesSingleContainer(ie)
		if err != nil {
			return nil, fmt.Errorf("newE2connectionUpdateRemoveItemIesSingleContainer() %s", err.Error())
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(e2curlC), unsafe.Pointer(ieC)); err != nil {
			return nil, err
		}
	}

	return &e2curlC, nil
}

func decodeE2connectionUpdateRemoveList(e2curlC *C.E2connectionUpdateRemove_List_t) (*e2ap_pdu_contents.E2ConnectionUpdateRemoveList, error) {

	e2curl := e2ap_pdu_contents.E2ConnectionUpdateRemoveList{
		Value: make([]*e2ap_pdu_contents.E2ConnectionUpdateRemoveItemIes, 0),
	}

	ieCount := int(e2curlC.list.count)
	for i := 0; i < ieCount; i++ {
		offset := unsafe.Sizeof(unsafe.Pointer(e2curlC.list.array)) * uintptr(i)
		ieC := *(**C.ProtocolIE_SingleContainer_1713P6_t)(unsafe.Pointer(uintptr(unsafe.Pointer(e2curlC.list.array)) + offset))
		ie, err := decodeE2connectionUpdateRemoveItemIesSingleContainer(ieC)
		if err != nil {
			return nil, fmt.Errorf("decodeE2connectionUpdateRemoveItemIesSingleContainer() %s", err.Error())
		}
		e2curl.Value = append(e2curl.Value, ie)
	}

	return &e2curl, nil
}

func decodeE2connectionUpdateRemoveListBytes(e2curlC [16]byte) (*e2ap_pdu_contents.E2ConnectionUpdateRemoveList, error) {
	array := (**C.struct_ProtocolIE_SingleContainer)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(e2curlC[0:8]))))
	count := C.int(binary.LittleEndian.Uint32(e2curlC[8:12]))
	size := C.int(binary.LittleEndian.Uint32(e2curlC[12:16]))

	rfIDlC := C.E2connectionUpdateRemove_List_t{
		list: C.struct___97{
			array: array,
			size:  size,
			count: count,
		},
	}

	return decodeE2connectionUpdateRemoveList(&rfIDlC)
}