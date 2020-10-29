// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package orane2

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "GlobalE2node-ID.h"
//#include "GlobalE2node-gNB-ID.h"
import "C"
import (
	"encoding/binary"
	"fmt"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2apies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2proxy/e2ctypes"
	"unsafe"
)

// XerEncodeGlobalE2nodeIDOld - used only in tests
// Deprecated: Do not use.
func XerEncodeGlobalE2nodeIDOld(globalE2nodeID *e2ctypes.GlobalE2Node_IDT) ([]byte, error) {
	globalE2nodeIDC, err := newGlobalE2nodeIDOld(globalE2nodeID)
	if err != nil {
		return nil, err
	}

	bytes, err := encodeXer(&C.asn_DEF_GlobalE2node_ID, unsafe.Pointer(globalE2nodeIDC))
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

// PerEncodeGlobalE2nodeIDOld - used only in tests
// Deprecated: Do not use.
func PerEncodeGlobalE2nodeIDOld(globalE2nodeID *e2ctypes.GlobalE2Node_IDT) ([]byte, error) {
	globalE2nodeIDC, err := newGlobalE2nodeIDOld(globalE2nodeID)
	if err != nil {
		return nil, err
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_GlobalE2node_ID, unsafe.Pointer(globalE2nodeIDC))
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

// Deprecated: Do not use.
func newGlobalE2nodeIDOld(gnID *e2ctypes.GlobalE2Node_IDT) (*C.GlobalE2node_ID_t, error) {
	var prC C.GlobalE2node_ID_PR

	choiceC := [8]byte{} // The size of the GlobalE2node_ID_u
	switch choice := gnID.GetChoice().(type) {
	case *e2ctypes.GlobalE2Node_IDT_GNB:
		prC = C.GlobalE2node_ID_PR_gNB

		globalgNBIDC, err := newGlobalE2nodegNBIDOld(choice.GNB)
		if err != nil {
			return nil, err
		}
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(globalgNBIDC))))
	default:
		return nil, fmt.Errorf("handling of %v not yet implemented", choice)
	}

	gnIDC := C.GlobalE2node_ID_t{
		present: prC,
		choice:  choiceC,
	}

	return &gnIDC, nil
}

// Deprecated: Do not use.
func decodeGlobalE2NodeIDOld(globalE2nodeIDchoice [48]byte) (*e2ctypes.GlobalE2Node_IDT, error) {

	present := C.long(binary.LittleEndian.Uint64(globalE2nodeIDchoice[0:8]))
	result := e2ctypes.GlobalE2Node_IDT{}

	switch present {
	case C.GlobalE2node_ID_PR_gNB:
		bufC := globalE2nodeIDchoice[8:16]
		gNbC := *(**C.GlobalE2node_gNB_ID_t)(unsafe.Pointer(&bufC[0]))
		gNB, err := decodeGlobalE2nodegNBIDOld(gNbC)
		if err != nil {
			return nil, fmt.Errorf("decodeGlobalE2NodeIDOld() %v", err)
		}

		result.Choice = &e2ctypes.GlobalE2Node_IDT_GNB{
			GNB: gNB,
		}
	default:
		return nil, fmt.Errorf("decodeGlobalE2NodeIDOld(). %v not yet implemneted", present)
	}

	return &result, nil
}

// XerEncodeGlobalE2nodeID - used only in tests
func XerEncodeGlobalE2nodeID(globalE2nodeID *e2apies.GlobalE2NodeId) ([]byte, error) {
	globalE2nodeIDC, err := newGlobalE2nodeID(globalE2nodeID)
	if err != nil {
		return nil, err
	}

	bytes, err := encodeXer(&C.asn_DEF_GlobalE2node_ID, unsafe.Pointer(globalE2nodeIDC))
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

// PerEncodeGlobalE2nodeID - used only in tests
func PerEncodeGlobalE2nodeID(globalE2nodeID *e2apies.GlobalE2NodeId) ([]byte, error) {
	globalE2nodeIDC, err := newGlobalE2nodeID(globalE2nodeID)
	if err != nil {
		return nil, err
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_GlobalE2node_ID, unsafe.Pointer(globalE2nodeIDC))
	if err != nil {
		return nil, err
	}

	return bytes, nil
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
	default:
		return nil, fmt.Errorf("handling of %v not yet implemented", choice)
	}

	gnIDC := C.GlobalE2node_ID_t{
		present: prC,
		choice:  choiceC,
	}

	return &gnIDC, nil
}

func decodeGlobalE2NodeID(globalE2nodeIDchoice [48]byte) (*e2apies.GlobalE2NodeId, error) {

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
	default:
		return nil, fmt.Errorf("decodeGlobalE2NodeID(). %v not yet implemneted", present)
	}

	return result, nil
}
