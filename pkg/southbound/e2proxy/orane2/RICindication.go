// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package orane2

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "RICindication.h"
//#include "Criticality.h"
//#include "ProtocolIE-Field.h"
import "C"
import (
	"fmt"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2proxy/e2ctypes"
	"unsafe"
)

// XerEncodeRICindication - This just dumps the XML to stdout
//func XerEncodeRICindication(ricInd *e2ctypes.RICindicationT) ([]byte, error) {
//	ricIndC, err := newRicIndication(ricInd)
//	if err != nil {
//		return nil, err
//	}
//
//	bytes, err := encodeXer(&C.asn_DEF_RICindication, unsafe.Pointer(ricIndC))
//	if err != nil {
//		return nil, err
//	}
//	return bytes, nil
//}

// PerEncodeRicIndication encodes a RICIndication to a []byte
func PerEncodeRicIndication(ri *C.RICindication_t) ([]byte, error) {
	perBuf := C.malloc(C.sizeof_uchar * 1024) // C allocated pointer
	defer C.free(perBuf)
	encRetVal, err := C.aper_encode_to_buffer(
		&C.asn_DEF_RICindication, nil, unsafe.Pointer(ri),
		perBuf, C.ulong(1024))
	if err != nil {
		return nil, err
	}
	if encRetVal.encoded == -1 {
		fmt.Printf("error on %v\n", *encRetVal.failed_type)

		return nil, fmt.Errorf("error encoding. Name: %v Tag: %v",
			C.GoString(encRetVal.failed_type.name),
			C.GoString(encRetVal.failed_type.xml_tag))
	}
	bytes := make([]byte, encRetVal.encoded)
	for i := 0; i < int(encRetVal.encoded); i++ {
		b := *(*C.uchar)(unsafe.Pointer(uintptr(perBuf) + uintptr(i)))
		bytes[i] = byte(b)
	}
	return bytes, nil
}

func newRicIndication(ri *e2ctypes.RICindicationT) (*C.RICindication_t, error) {
	pIeC1544P6, err := newProtocolIeContainer1544P6(ri.ProtocolIEs)
	if err != nil {
		return nil, err
	}
	riC := C.RICindication_t{
		protocolIEs: *pIeC1544P6,
	}

	return &riC, nil
}
