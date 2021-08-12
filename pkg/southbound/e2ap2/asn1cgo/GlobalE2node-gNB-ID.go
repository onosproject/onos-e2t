// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "GlobalE2node-gNB-ID.h"
import "C"
import (
	"encoding/binary"
	"fmt"
	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	"unsafe"
)

func xerEncodeGlobalE2nodegNBID(ge2n *e2apies.GlobalE2NodeGnbId) ([]byte, error) {
	ge2nCP, err := newGlobalE2nodegNBID(ge2n)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeGlobalE2nodegNBID() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_GlobalE2node_gNB_ID, unsafe.Pointer(ge2nCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeGlobalE2nodegNBID() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeGlobalE2nodegNBID(ge2n *e2apies.GlobalE2NodeGnbId) ([]byte, error) {
	ge2nCP, err := newGlobalE2nodegNBID(ge2n)
	if err != nil {
		return nil, fmt.Errorf("perEncodeGlobalE2nodegNBID() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_GlobalE2node_gNB_ID, unsafe.Pointer(ge2nCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeGlobalE2nodegNBID() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeGlobalE2nodegNBID(bytes []byte) (*e2apies.GlobalE2NodeGnbId, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_GlobalE2node_gNB_ID)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeGlobalE2nodegNBID((*C.GlobalE2node_gNB_ID_t)(unsafePtr))
}

func perDecodeGlobalE2nodegNBID(bytes []byte) (*e2apies.GlobalE2NodeGnbId, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_GlobalE2node_gNB_ID)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeGlobalE2nodegNBID((*C.GlobalE2node_gNB_ID_t)(unsafePtr))
}

func newGlobalE2nodegNBID(gnbID *e2apies.GlobalE2NodeGnbId) (*C.GlobalE2node_gNB_ID_t, error) {

	globalgNBID, err := newGlobalgNBID(gnbID.GlobalGNbId)
	if err != nil {
		return nil, err
	}

	globalgNBIDC := C.GlobalE2node_gNB_ID_t{
		global_gNB_ID: *globalgNBID,
		//gNB_CU_UP_ID:  nil,
		//gNB_DU_ID:     nil,
	}

	if gnbID.GNbCuUpId != nil {
		globalgNBIDC.gNB_CU_UP_ID, err = newGnbCuUpID(gnbID.GNbCuUpId)
		if err != nil {
			return nil, err
		}
	}

	if gnbID.GNbDuId != nil {
		globalgNBIDC.gNB_DU_ID, err = newGnbDuID(gnbID.GNbDuId)
		if err != nil {
			return nil, err
		}
	}

	return &globalgNBIDC, nil
}

func decodeGlobalE2nodegNBID(gNBC *C.GlobalE2node_gNB_ID_t) (*e2apies.GlobalE2NodeGnbId, error) {
	result := new(e2apies.GlobalE2NodeGnbId)
	var err error
	result.GlobalGNbId, err = decodeGlobalGnbID(&gNBC.global_gNB_ID)
	if err != nil {
		return nil, fmt.Errorf("error decodeGlobalE2nodegNBID() %v", err)
	}

	if gNBC.gNB_CU_UP_ID != nil {
		result.GNbCuUpId, err = decodeGnbCuUpID(gNBC.gNB_CU_UP_ID)
		if err != nil {
			return nil, fmt.Errorf("error decodeGlobalE2nodegNBID() %v", err)
		}
	}

	if gNBC.gNB_DU_ID != nil {
		result.GNbDuId, err = decodeGnbDuID(gNBC.gNB_DU_ID)
		if err != nil {
			return nil, fmt.Errorf("error decodeGlobalE2nodegNBID() %v", err)
		}
	}

	return result, nil
}

func decodeGlobalE2nodegNBIDBytes(array [8]byte) (*e2apies.GlobalE2NodeGnbId, error) {
	gNBC := (*C.GlobalE2node_gNB_ID_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeGlobalE2nodegNBID(gNBC)
}
