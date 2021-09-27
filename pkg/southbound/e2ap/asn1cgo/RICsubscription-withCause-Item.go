// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "RICsubscription-withCause-Item.h"
import "C"
import (
	"encoding/binary"
	"fmt"
	"unsafe"

	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
)

func xerEncodeRicSubscriptionWithCauseItem(ratbsi *e2appducontents.RicsubscriptionWithCauseItem) ([]byte, error) {
	ratbsiCP, err := newRicSubscriptionWithCauseItem(ratbsi)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeRicSubscriptionWithCauseItem() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_RICsubscription_withCause_Item, unsafe.Pointer(ratbsiCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeRicSubscriptionWithCauseItem() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeRicSubscriptionWithCauseItem(ratbsi *e2appducontents.RicsubscriptionWithCauseItem) ([]byte, error) {
	ratbsiCP, err := newRicSubscriptionWithCauseItem(ratbsi)
	if err != nil {
		return nil, fmt.Errorf("perEncodeRicSubscriptionWithCauseItem() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_RICsubscription_withCause_Item, unsafe.Pointer(ratbsiCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeRicSubscriptionWithCauseItem() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeRicSubscriptionWithCauseItem(bytes []byte) (*e2appducontents.RicsubscriptionWithCauseItem, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_RICsubscription_withCause_Item)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeRicSubscriptionWithCauseItem((*C.RICsubscription_withCause_Item_t)(unsafePtr))
}

func perDecodeRicSubscriptionWithCauseItem(bytes []byte) (*e2appducontents.RicsubscriptionWithCauseItem, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_RICsubscription_withCause_Item)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeRicSubscriptionWithCauseItem((*C.RICsubscription_withCause_Item_t)(unsafePtr))
}

func newRicSubscriptionWithCauseItem(rswcItem *e2appducontents.RicsubscriptionWithCauseItem) (*C.RICsubscription_withCause_Item_t, error) {

	c, err := newCause(rswcItem.GetCause())
	if err != nil {
		return nil, fmt.Errorf("newCause() %s", err.Error())
	}

	rswcItemC := C.RICsubscription_withCause_Item_t{
		ricRequestID:  *newRicRequestID(rswcItem.GetRicRequestId()),
		ranFunctionID: newRanFunctionID(rswcItem.GetRanFunctionId()),
		cause:         *c,
	}

	return &rswcItemC, nil
}

func decodeRicSubscriptionWithCauseItemBytes(bytes [112]byte) (*e2appducontents.RicsubscriptionWithCauseItem, error) {

	rswcItemC := C.RICsubscription_withCause_Item_t{
		ricRequestID: C.RICrequestID_t{
			ricRequestorID: C.long(binary.LittleEndian.Uint64(bytes[:8])),
			ricInstanceID:  C.long(binary.LittleEndian.Uint64(bytes[8:16])),
		},
		// Gap of 24 for the asn_struct_ctx_t belonging to RICrequestID --> 40
		ranFunctionID: C.long(binary.LittleEndian.Uint64(bytes[40:48])),
		cause: C.Cause_t{
			present: C.Cause_PR(binary.LittleEndian.Uint64(bytes[48:56])),
		},
	}
	copy(rswcItemC.cause.choice[:], bytes[56:64])
	// Gap of 24 for the asn_struct_ctx_t belonging to Cause --> 88
	// Gap of 24 for the asn_struct_ctx_t belonging to RICsubscription_withCause_Item --> 112

	return decodeRicSubscriptionWithCauseItem(&rswcItemC)
}

func decodeRicSubscriptionWithCauseItem(rswcItemC *C.RICsubscription_withCause_Item_t) (*e2appducontents.RicsubscriptionWithCauseItem, error) {

	c, err := decodeCause(&rswcItemC.cause)
	if err != nil {
		return nil, err
	}

	rswcItem := e2appducontents.RicsubscriptionWithCauseItem{
		RicRequestId:  decodeRicRequestID(&rswcItemC.ricRequestID),
		RanFunctionId: decodeRanFunctionID(&rswcItemC.ranFunctionID),
		Cause:         c,
	}

	return &rswcItem, nil
}
