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
	"fmt"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2ap-commondatatypes"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2proxy/e2ctypes"
	"math"
	"unsafe"
)

// XerEncodeBitStringOld - used only in tests
// Deprecated: Do not use.
func XerEncodeBitStringOld(bs *e2ctypes.BIT_STRING) ([]byte, error) {
	bsC := newBitStringOld(bs)

	bytes, err := encodeXer(&C.asn_DEF_BIT_STRING, unsafe.Pointer(bsC))
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

// PerEncodeBitStringOld - used only in tests
// Deprecated: Do not use.
func PerEncodeBitStringOld(bs *e2ctypes.BIT_STRING) ([]byte, error) {
	bsC := newBitStringOld(bs)

	bytes, err := encodePerBuffer(&C.asn_DEF_BIT_STRING, unsafe.Pointer(bsC))
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

// Deprecated: Do not use.
func newBitStringOld(bs *e2ctypes.BIT_STRING) *C.BIT_STRING_t {
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

// Deprecated: Do not use.
func decodeBitStringOld(bsC [48]byte) *e2ctypes.BIT_STRING {
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

// XerEncodeGnbID - used only in tests
func XerEncodeBitString(bs *e2ap_commondatatypes.BitString) ([]byte, error) {
	bsC := newBitString(bs)

	bytes, err := encodeXer(&C.asn_DEF_BIT_STRING, unsafe.Pointer(bsC))
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

// PerEncodeGnbID - used only in tests
func PerEncodeBitString(bs *e2ap_commondatatypes.BitString) ([]byte, error) {
	bsC := newBitString(bs)

	bytes, err := encodePerBuffer(&C.asn_DEF_BIT_STRING, unsafe.Pointer(bsC))
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func newBitString(bs *e2ap_commondatatypes.BitString) *C.BIT_STRING_t {
	bitsUnused := 64 - int(bs.Len)
	valAsBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(valAsBytes, bs.Value)

	bsC := C.BIT_STRING_t{
		buf:         (*C.uchar)(C.CBytes(valAsBytes)),
		size:        C.ulong(8),
		bits_unused: C.int(bitsUnused),
	}
	fmt.Printf("Bit string %+v\n", bsC)
	return &bsC
}

// decodeBitString - byteString in C has 20 bytes
// 8 for a 64bit address of a buffer, 8 for the size in bytes of the buffer uint64, 4 for the unused bits
// The unused bits are at the end of the buffer
func decodeBitString(bsC [20]byte) (*e2ap_commondatatypes.BitString, error) {
	bufAddr := bsC[0:8]
	size := binary.LittleEndian.Uint64(bsC[8:16])
	bitsUnused := binary.LittleEndian.Uint32(bsC[16:20])
	if size > 8 {
		return nil, fmt.Errorf("max size is 8 bytes (64 bits) got %d", size)
	} else if uint64(bitsUnused) > (size * 8) {
		return nil, fmt.Errorf("bits unused (%d) is greater than bits used (64)", bitsUnused)
	}

	goBytes := make([]byte, 8)
	bytes := C.GoBytes(unsafe.Pointer(*(**C.char)(unsafe.Pointer(&bufAddr[0]))), C.int(size))
	// Need to bit shift whole array to the right by bitsUnused
	//var carry byte
	//mask := byte(math.Pow(2, float64(size)) - 1)
	for i := 0; i < int(size); i++ {
		//prevCarry := carry << (8 - bitsUnused%8)
		//carry = bytes[i] & mask
		//bytes[i] = bytes[i] >> bitsUnused%8
		//goBytes[i] = bytes[i] | prevCarry
		goBytes[i] = bytes[i]
	}
	//fmt.Printf("bit string %x %d %d %+x %+x\n", bufAddr, size, bitsUnused, bytes, goBytes)
	bs := &e2ap_commondatatypes.BitString{
		Value: binary.LittleEndian.Uint64(goBytes),
		Len: uint32(size*8 - uint64(bitsUnused)),
	}

	return bs, nil
}
