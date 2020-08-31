// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package orane2

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "BIT_STRING.h"
import "C"
import (
	"encoding/binary"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2proxy/e2ctypes"
	"math"
	"unsafe"
)

// XerEncodeBitString - used only in tests
func XerEncodeBitString(bs *e2ctypes.BIT_STRING) ([]byte, error) {
	bsC := newBitString(bs)

	bytes, err := encodeXer(&C.asn_DEF_BIT_STRING, unsafe.Pointer(bsC))
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

// PerEncodeBitString - used only in tests
func PerEncodeBitString(bs *e2ctypes.BIT_STRING) ([]byte, error) {
	bsC := newBitString(bs)

	bytes, err := encodePerBuffer(&C.asn_DEF_BIT_STRING, unsafe.Pointer(bsC))
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func newBitString(bs *e2ctypes.BIT_STRING) *C.BIT_STRING_t {
	bitsUnused := len(bs.BitString)*8 - int(bs.Numbits)
	if len(bs.BitString)*8 < int(bs.Numbits) {
		bitsUnused = 0
	}
	bsC := C.BIT_STRING_t{
		buf:         (*C.uchar)(C.CBytes(bs.BitString)),
		size:        C.ulong(len(bs.BitString)),
		bits_unused: C.int(bitsUnused),
	}
	//fmt.Printf("Bit string %+v\n", bsC)
	return &bsC
}

func decodeBitString(bsC [48]byte) *e2ctypes.BIT_STRING {
	bufC := bsC[0:8]
	size := binary.LittleEndian.Uint64(bsC[8:16])
	bitsUnused := binary.LittleEndian.Uint32(bsC[16:20])
	//fmt.Printf("bit string %x %d %d %+x\n", buf, size, bitsUnused, bsC)
	bs := &e2ctypes.BIT_STRING{
		BitString: make([]byte, size),
		Numbits:   uint32(size*8 - uint64(bitsUnused)),
	}
	bytes := C.GoBytes(unsafe.Pointer(*(**C.char)(unsafe.Pointer(&bufC[0]))), C.int(size))
	// Need to bit shift whole array to the right by bitsUnused
	var carry byte
	mask := byte(math.Pow(2, float64(size)) - 1)
	for i := 0; i < int(size); i++ {
		prevCarry := carry << (8 - bitsUnused)
		carry = bytes[i] & mask
		bytes[i] = bytes[i] >> bitsUnused
		bs.BitString[i] = bytes[i] | prevCarry
	}

	return bs
}
