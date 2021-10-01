// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "GlobalgNB-ID.h"
//#include "GlobalngeNB-ID.h"
//#include "GlobalNG-RANNode-ID.h"
import "C"
import (
	"encoding/binary"
	"fmt"
	"unsafe"

	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
)

func xerEncodeGlobalNgRanNodeID(ge2n *e2apies.GlobalNgRannodeId) ([]byte, error) {
	ge2nCP, err := newGlobalNgRanNodeID(ge2n)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeGlobalNgRanNodeID() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_GlobalNG_RANNode_ID, unsafe.Pointer(ge2nCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeGlobalNgRanNodeID() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeGlobalNgRanNodeID(ge2n *e2apies.GlobalNgRannodeId) ([]byte, error) {
	ge2nCP, err := newGlobalNgRanNodeID(ge2n)
	if err != nil {
		return nil, fmt.Errorf("PerEncodeGlobalNgRanNodeID() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_GlobalNG_RANNode_ID, unsafe.Pointer(ge2nCP))
	if err != nil {
		return nil, fmt.Errorf("PerEncodeGlobalNgRanNodeID() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeGlobalNgRanNodeID(bytes []byte) (*e2apies.GlobalNgRannodeId, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_GlobalNG_RANNode_ID)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeGlobalNgRanNodeID((*C.GlobalNG_RANNode_ID_t)(unsafePtr))
}

func perDecodeGlobalNgRanNodeID(bytes []byte) (*e2apies.GlobalNgRannodeId, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_GlobalNG_RANNode_ID)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeGlobalNgRanNodeID((*C.GlobalNG_RANNode_ID_t)(unsafePtr))
}

func newGlobalNgRanNodeID(gnID *e2apies.GlobalNgRannodeId) (*C.GlobalNG_RANNode_ID_t, error) {
	var prC C.GlobalNG_RANNode_ID_PR

	choiceC := [8]byte{} // The size of the GlobalE2node_ID_u
	switch choice := gnID.GetGlobalNgRannodeId().(type) {
	case *e2apies.GlobalNgRannodeId_GNb:
		prC = C.GlobalNG_RANNode_ID_PR_gNB

		globalgNBIDC, err := newGlobalgNBID(choice.GNb)
		if err != nil {
			return nil, err
		}
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(globalgNBIDC))))
	case *e2apies.GlobalNgRannodeId_NgENb:
		prC = C.GlobalNG_RANNode_ID_PR_ng_eNB

		globalEnbIDC, err := newGlobalngeNbID(choice.NgENb)
		if err != nil {
			return nil, err
		}
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(globalEnbIDC))))
	default:
		return nil, fmt.Errorf("handling of %v not yet implemented", choice)
	}

	gnIDC := C.GlobalNG_RANNode_ID_t{
		present: prC,
		choice:  choiceC,
	}

	return &gnIDC, nil
}

func decodeGlobalNgRanNodeID(globalNgRanNodeID *C.GlobalNG_RANNode_ID_t) (*e2apies.GlobalNgRannodeId, error) {

	result := new(e2apies.GlobalNgRannodeId)

	switch globalNgRanNodeID.present {
	case C.GlobalNG_RANNode_ID_PR_gNB:
		gNB, err := decodeGlobalGnbIDBytes(globalNgRanNodeID.choice)
		if err != nil {
			return nil, fmt.Errorf("decodeGlobalE2NodeID() %v", err)
		}

		result.GlobalNgRannodeId = &e2apies.GlobalNgRannodeId_GNb{
			GNb: gNB,
		}
	case C.GlobalNG_RANNode_ID_PR_ng_eNB:
		ngENb, err := decodeGlobalngeNbIDBytes(globalNgRanNodeID.choice)
		if err != nil {
			return nil, fmt.Errorf("decodeGlobalE2nodeeNBID() %v", err)
		}

		result.GlobalNgRannodeId = &e2apies.GlobalNgRannodeId_NgENb{
			NgENb: ngENb,
		}
	default:
		return nil, fmt.Errorf("decodeGlobalNgRanNodeID(). %v not yet implemneted", globalNgRanNodeID.present)
	}

	return result, nil
}

//func decodeGlobalNgRanNodeIDBytes(globalNgRanNodeIDchoice [48]byte) (*e2apies.GlobalNgRannodeId, error) {
//
//	present := C.long(binary.LittleEndian.Uint64(globalNgRanNodeIDchoice[0:8]))
//	result := new(e2apies.GlobalNgRannodeId)
//
//	switch present {
//	case C.GlobalNG_RANNode_ID_PR_gNB:
//		bufC := globalNgRanNodeIDchoice[8:16]
//		gNbC := *(**C.GlobalgNB_ID_t)(unsafe.Pointer(&bufC[0]))
//		gNB, err := decodeGlobalGnbID(gNbC)
//		if err != nil {
//			return nil, fmt.Errorf("decodeGlobalE2NodeID() %v", err)
//		}
//
//		result.GlobalNgRannodeId = &e2apies.GlobalNgRannodeId_GNb{
//			GNb: gNB,
//		}
//	case C.GlobalNG_RANNode_ID_PR_ng_eNB:
//		bufC := globalNgRanNodeIDchoice[8:16]
//		ngENbC := *(**C.GlobalngeNB_ID_t)(unsafe.Pointer(&bufC[0]))
//		ngENb, err := decodeGlobalngeNbID(ngENbC)
//		if err != nil {
//			return nil, fmt.Errorf("decodeGlobalE2nodeNgEnbID() %v", err)
//		}
//
//		result.GlobalNgRannodeId = &e2apies.GlobalNgRannodeId_NgENb{
//			NgENb: ngENb,
//		}
//	default:
//		return nil, fmt.Errorf("decodeGlobalNgRanNodeIDBytes(). %v not yet implemneted", present)
//	}
//
//	return result, nil
//}
