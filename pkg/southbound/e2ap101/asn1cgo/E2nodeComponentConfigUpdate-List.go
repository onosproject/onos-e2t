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
//#include "ProtocolIE-SingleContainer.h"
import "C"
import (
	"encoding/binary"
	"fmt"
	e2ap_pdu_contents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
	"unsafe"
)

func newE2nodeComponentConfigUpdateList(e2nodeComponentConfigUpdateList *e2ap_pdu_contents.E2NodeComponentConfigUpdateList) (*C.E2nodeComponentConfigUpdate_List_t, error) {

	e2nodeComponentConfigUpdateListC := C.E2nodeComponentConfigUpdate_List_t{}
	for _, ie := range e2nodeComponentConfigUpdateList.GetValue() {
		ieC, err := newE2nodeConfigurationUpdateIesSingleContainer(ie)
		if err != nil {
			return nil, fmt.Errorf("newE2nodeConfigurationUpdateIesSingleContainer() %s", err.Error())
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(e2nodeComponentConfigUpdateListC), unsafe.Pointer(ieC)); err != nil {
			return nil, err
		}
	}

	return &e2nodeComponentConfigUpdateListC, nil
}

func decodeE2nodeComponentConfigUpdateList(e2nodeComponentConfigUpdateListC *C.E2nodeComponentConfigUpdate_List_t) (*e2ap_pdu_contents.E2NodeComponentConfigUpdateList, error) {

	e2nodeComponentConfigUpdateList := e2ap_pdu_contents.E2NodeComponentConfigUpdateList{
		Value: make([]*e2ap_pdu_contents.E2NodeComponentConfigUpdateItemIes, 0),
	}

	ieCount := int(e2nodeComponentConfigUpdateListC.list.count)
	for i := 0; i < ieCount; i++ {
		offset := unsafe.Sizeof(unsafe.Pointer(e2nodeComponentConfigUpdateListC.list.array)) * uintptr(i)
		ieC := *(**C.ProtocolIE_SingleContainer_1713P6_t)(unsafe.Pointer(uintptr(unsafe.Pointer(e2nodeComponentConfigUpdateListC.list.array)) + offset))
		ie, err := decodeE2nodeComponentConfigUpdateItemIesSingleContainer(ieC)
		if err != nil {
			return nil, fmt.Errorf("decodeE2nodeComponentConfigUpdateItemIesSingleContainer() %s", err.Error())
		}
		e2nodeComponentConfigUpdateList.Value = append(e2nodeComponentConfigUpdateList.Value, ie)
	}

	return &e2nodeComponentConfigUpdateList, nil
}

func decodeE2nodeComponentConfigUpdateListBytes(e2ncculC [16]byte) (*e2ap_pdu_contents.E2NodeComponentConfigUpdateList, error) {
	array := (**C.struct_ProtocolIE_SingleContainer)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(e2ncculC[0:8]))))
	count := C.int(binary.LittleEndian.Uint32(e2ncculC[8:12]))
	size := C.int(binary.LittleEndian.Uint32(e2ncculC[12:16]))

	rfIDlC := C.E2nodeComponentConfigUpdate_List_t{
		list: C.struct___127{
			array: array,
			size:  size,
			count: count,
		},
	}

	return decodeE2nodeComponentConfigUpdateList(&rfIDlC)
}
