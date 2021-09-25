// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

// #cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
// #cgo LDFLAGS: -lm
// #include <stdio.h>
// #include <stdlib.h>
// #include <assert.h>
// #include "RICsubscription-List-withCause.h"
//#include "ProtocolIE-SingleContainer.h"
import "C"
import (
	"encoding/binary"
	"fmt"
	"unsafe"

	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
)

func xerEncodeRicSubscriptionListWithCause(ratbsl *e2appducontents.RicsubscriptionListWithCause) ([]byte, error) {
	ratbslC, err := newRicSubscriptionListWithCause(ratbsl)

	if err != nil {
		return nil, err
	}

	bytes, err := encodeXer(&C.asn_DEF_RICsubscription_List_withCause, unsafe.Pointer(ratbslC))
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func perEncodeRicSubscriptionListWithCause(ratbsl *e2appducontents.RicsubscriptionListWithCause) ([]byte, error) {
	ratbslC, err := newRicSubscriptionListWithCause(ratbsl)

	if err != nil {
		return nil, err
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_RICsubscription_List_withCause, unsafe.Pointer(ratbslC))
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func xerDecodeRicSubscriptionListWithCause(xer []byte) (*e2appducontents.RicsubscriptionListWithCause, error) {
	unsafePtr, err := decodeXer(xer, &C.asn_DEF_RICsubscription_List_withCause)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	ratsslC := (*C.RICsubscription_List_withCause_t)(unsafePtr)
	return decodeRicSubscriptionListWithCause(ratsslC)
}

func perDecodeRicSubscriptionListWithCause(bytes []byte) (*e2appducontents.RicsubscriptionListWithCause, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_RICsubscription_List_withCause)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeRicSubscriptionListWithCause((*C.RICsubscription_List_withCause_t)(unsafePtr))
}

func newRicSubscriptionListWithCause(ratbsL *e2appducontents.RicsubscriptionListWithCause) (*C.RICsubscription_List_withCause_t, error) {
	ratbsLC := new(C.RICsubscription_List_withCause_t)
	for _, ricActionToBeSetupItemIe := range ratbsL.GetValue() {
		ricActionToBeItemIesScC, err := newRicSubscriptionWithCauseItemIesSingleContainer(ricActionToBeSetupItemIe)
		if err != nil {
			return nil, fmt.Errorf("newRicActionToBeSetupItemIesSingleContainer() %s", err.Error())
		}

		if _, err = C.asn_sequence_add(unsafe.Pointer(ratbsLC), unsafe.Pointer(ricActionToBeItemIesScC)); err != nil {
			return nil, err
		}
	}

	return ratbsLC, nil
}

func decodeRicSubscriptionListWithCause(ratbsLC *C.RICsubscription_List_withCause_t) (*e2appducontents.RicsubscriptionListWithCause, error) {
	ratbsL := e2appducontents.RicsubscriptionListWithCause{
		Value: make([]*e2appducontents.RicsubscriptionWithCauseItemIes, 0),
	}

	ieCount := int(ratbsLC.list.count)
	for i := 0; i < ieCount; i++ {
		offset := unsafe.Sizeof(unsafe.Pointer(*ratbsLC.list.array)) * uintptr(i)
		ratbsIeC := *(**C.ProtocolIE_SingleContainer_1911P3_t)(unsafe.Pointer(uintptr(unsafe.Pointer(ratbsLC.list.array)) + offset))
		ratbsIe, err := decodeRicSubscriptionWithCauseItemIesSingleContainer(ratbsIeC)
		if err != nil {
			return nil, fmt.Errorf("decodeRicActionToBeSetupItemIesSingleContainer() %s", err.Error())
		}
		ratbsL.Value = append(ratbsL.Value, ratbsIe)
	}

	return &ratbsL, nil
}

func decodeRicSubscriptionListWithCauseBytes(list [48]byte) (*e2appducontents.RicsubscriptionListWithCause, error) {
	array := (**C.struct_ProtocolIE_SingleContainer)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(list[0:8]))))
	count := C.int(binary.LittleEndian.Uint32(list[8:12]))
	size := C.int(binary.LittleEndian.Uint32(list[12:16]))

	ranFunctionListChoiceC := C.RICsubscription_List_withCause_t{
		list: C.struct___111{
			array: array,
			size:  size,
			count: count,
		},
	}

	return decodeRicSubscriptionListWithCause(&ranFunctionListChoiceC)
}
