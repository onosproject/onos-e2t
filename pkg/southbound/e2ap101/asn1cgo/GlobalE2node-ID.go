// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "GlobalE2node-ID.h"
//#include "GlobalE2node-gNB-ID.h"
//#include "GlobalE2node-en-gNB-ID.h"
//#include "GlobalE2node-eNB-ID.h"
//#include "GlobalE2node-ng-eNB-ID.h"
import "C"
import (
	"encoding/binary"
	"fmt"
	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/types"
	"unsafe"
)

func xerEncodeGlobalE2nodeID(ge2n *e2apies.GlobalE2NodeId) ([]byte, error) {
	ge2nCP, err := newGlobalE2nodeID(ge2n)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeGlobalE2nodeID() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_GlobalE2node_ID, unsafe.Pointer(ge2nCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeGlobalE2nodeID() %s", err.Error())
	}
	return bytes, nil
}

func PerEncodeGlobalE2nodeID(ge2n *e2apies.GlobalE2NodeId) ([]byte, error) {
	ge2nCP, err := newGlobalE2nodeID(ge2n)
	if err != nil {
		return nil, fmt.Errorf("perEncodeGlobalE2nodeID() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_GlobalE2node_ID, unsafe.Pointer(ge2nCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeGlobalE2nodeID() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeGlobalE2nodeID(bytes []byte) (*e2apies.GlobalE2NodeId, *types.E2NodeType, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_GlobalE2node_ID)
	if err != nil {
		return nil, nil, err
	}
	if unsafePtr == nil {
		return nil, nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeGlobalE2NodeID((*C.GlobalE2node_ID_t)(unsafePtr))
}

func PerDecodeGlobalE2nodeID(bytes []byte) (*e2apies.GlobalE2NodeId, *types.E2NodeType, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_GlobalE2node_ID)
	if err != nil {
		return nil, nil, err
	}
	if unsafePtr == nil {
		return nil, nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeGlobalE2NodeID((*C.GlobalE2node_ID_t)(unsafePtr))
}

func newGlobalE2nodeID(gnID *e2apies.GlobalE2NodeId) (*C.GlobalE2node_ID_t, error) {
	var prC C.GlobalE2node_ID_PR

	choiceC := [8]byte{} // The size of the GlobalE2node_ID_u
	switch choice := gnID.GetGlobalE2NodeId().(type) {
	case *e2apies.GlobalE2NodeId_GNb:
		prC = C.GlobalE2node_ID_PR_gNB

		globalgNBIDC, err := newGlobalE2nodegNBID(choice.GNb)
		if err != nil {
			return nil, err
		}
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(globalgNBIDC))))
	case *e2apies.GlobalE2NodeId_EnGNb:
		prC = C.GlobalE2node_ID_PR_en_gNB

		globalEngNBIDC, err := newGlobalE2nodeEnGnbID(choice.EnGNb)
		if err != nil {
			return nil, err
		}
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(globalEngNBIDC))))
	case *e2apies.GlobalE2NodeId_ENb:
		prC = C.GlobalE2node_ID_PR_eNB

		globalEnbIDC, err := newGlobalE2nodeeNBID(choice.ENb)
		if err != nil {
			return nil, err
		}
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(globalEnbIDC))))
	case *e2apies.GlobalE2NodeId_NgENb:
		prC = C.GlobalE2node_ID_PR_ng_eNB

		globalEnbIDC, err := newGlobalE2nodeNgEnbID(choice.NgENb)
		if err != nil {
			return nil, err
		}
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(globalEnbIDC))))
	default:
		return nil, fmt.Errorf("handling of %v not yet implemented", choice)
	}

	gnIDC := C.GlobalE2node_ID_t{
		present: prC,
		choice:  choiceC,
	}

	return &gnIDC, nil
}

func decodeGlobalE2NodeID(globalE2nodeID *C.GlobalE2node_ID_t) (*e2apies.GlobalE2NodeId, *types.E2NodeType, error) {

	result := new(e2apies.GlobalE2NodeId)
	var e2NodeType types.E2NodeType

	switch globalE2nodeID.present {
	case C.GlobalE2node_ID_PR_gNB:
		gNB, err := decodeGlobalE2nodegNBIDBytes(globalE2nodeID.choice)
		if err != nil {
			return nil, nil, fmt.Errorf("decodeGlobalE2NodeID() %v", err)
		}

		result.GlobalE2NodeId = &e2apies.GlobalE2NodeId_GNb{
			GNb: gNB,
		}
		e2NodeType = types.E2NodeTypeGNB
	case C.GlobalE2node_ID_PR_en_gNB:
		enGNb, err := decodeGlobalE2nodeEnGnbIDBytes(globalE2nodeID.choice)
		if err != nil {
			return nil, nil, fmt.Errorf("decodeGlobalE2NodeID() %v", err)
		}

		result.GlobalE2NodeId = &e2apies.GlobalE2NodeId_EnGNb{
			EnGNb: enGNb,
		}
		e2NodeType = types.E2NodeTypeEnGNB
	case C.GlobalE2node_ID_PR_eNB:
		eNB, err := decodeGlobalE2nodeeNBIDBytes(globalE2nodeID.choice)
		if err != nil {
			return nil, nil, fmt.Errorf("decodeGlobalE2nodeeNBID() %v", err)
		}

		result.GlobalE2NodeId = &e2apies.GlobalE2NodeId_ENb{
			ENb: eNB,
		}
		e2NodeType = types.E2NodeTypeENB
	case C.GlobalE2node_ID_PR_ng_eNB:
		ngENb, err := decodeGlobalE2nodeNgEnbIDBytes(globalE2nodeID.choice)
		if err != nil {
			return nil, nil, fmt.Errorf("decodeGlobalE2nodeeNBID() %v", err)
		}

		result.GlobalE2NodeId = &e2apies.GlobalE2NodeId_NgENb{
			NgENb: ngENb,
		}
		e2NodeType = types.E2NodeTypeNgENB
	default:
		return nil, nil, fmt.Errorf("decodeGlobalE2NodeID(). %v not yet implemneted", globalE2nodeID.present)
	}

	return result, &e2NodeType, nil
}

func decodeGlobalE2NodeIDBytes(globalE2nodeIDchoice [48]byte) (*e2apies.GlobalE2NodeId, error) {

	present := C.long(binary.LittleEndian.Uint64(globalE2nodeIDchoice[0:8]))
	result := new(e2apies.GlobalE2NodeId)

	switch present {
	case C.GlobalE2node_ID_PR_gNB:
		bufC := globalE2nodeIDchoice[8:16]
		gNbC := *(**C.GlobalE2node_gNB_ID_t)(unsafe.Pointer(&bufC[0]))
		gNB, err := decodeGlobalE2nodegNBID(gNbC)
		if err != nil {
			return nil, fmt.Errorf("decodeGlobalE2NodeID() %v", err)
		}

		result.GlobalE2NodeId = &e2apies.GlobalE2NodeId_GNb{
			GNb: gNB,
		}
	case C.GlobalE2node_ID_PR_en_gNB:
		bufC := globalE2nodeIDchoice[8:16]
		enGNbC := *(**C.GlobalE2node_en_gNB_ID_t)(unsafe.Pointer(&bufC[0]))
		enGNb, err := decodeGlobalE2nodeEnGnbID(enGNbC)
		if err != nil {
			return nil, fmt.Errorf("decodeGlobalE2nodeEnGnbID() %v", err)
		}

		result.GlobalE2NodeId = &e2apies.GlobalE2NodeId_EnGNb{
			EnGNb: enGNb,
		}
	case C.GlobalE2node_ID_PR_eNB:
		bufC := globalE2nodeIDchoice[8:16]
		eNbC := *(**C.GlobalE2node_eNB_ID_t)(unsafe.Pointer(&bufC[0]))
		eNB, err := decodeGlobalE2nodeeNBID(eNbC)
		if err != nil {
			return nil, fmt.Errorf("decodeGlobalE2nodeeNBID() %v", err)
		}

		result.GlobalE2NodeId = &e2apies.GlobalE2NodeId_ENb{
			ENb: eNB,
		}
	case C.GlobalE2node_ID_PR_ng_eNB:
		bufC := globalE2nodeIDchoice[8:16]
		ngENbC := *(**C.GlobalE2node_ng_eNB_ID_t)(unsafe.Pointer(&bufC[0]))
		ngENb, err := decodeGlobalE2nodeNgEnbID(ngENbC)
		if err != nil {
			return nil, fmt.Errorf("decodeGlobalE2nodeNgEnbID() %v", err)
		}

		result.GlobalE2NodeId = &e2apies.GlobalE2NodeId_NgENb{
			NgENb: ngENb,
		}
	default:
		return nil, fmt.Errorf("decodeGlobalE2NodeID(). %v not yet implemneted", present)
	}

	return result, nil
}
