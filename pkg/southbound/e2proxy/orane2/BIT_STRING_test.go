// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package orane2

import (
	"encoding/binary"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2ap-commondatatypes"
	"gotest.tools/assert"
	"testing"
	"unsafe"
)

func Test_newBitString(t *testing.T) {
	bs1 := e2ap_commondatatypes.BitString{
		Value: 0x9ABCDEF,
		Len:   28,
	}

	cBitString := newBitString(&bs1)

	assert.Equal(t, 8, int(cBitString.size), "unexpected number of bits")
	assert.Equal(t, 36, int(cBitString.bits_unused), "unexpected number of bits_unused")
	// Can't do any further analysis as we can't have C in tests
}

func Test_decodeBitString(t *testing.T) {
	value := [4]byte{0x9A, 0xBC, 0xDE, 0xF0}
	bsBytes := make([]byte, 20) // 8 for a 64bit address, 8 for the size uint64, 4 for the unused bits
	binary.LittleEndian.PutUint64(bsBytes, uint64(uintptr(unsafe.Pointer(&value))))
	bsBytes[8] = 4 // num bytes = num bits / 8
	bsBytes[16] = 4 // unused bits

	bytes20 := [20]byte{}
	for i, b := range bsBytes {
		bytes20[i] = b
	}

	protoBitString, err := decodeBitString(bytes20)
	assert.NilError(t, err)
	assert.Assert(t, protoBitString != nil)
	assert.Equal(t, int(protoBitString.Len), 28, "unexpected bit string length")
	assert.Equal(t, int(protoBitString.Value), 0xf0debc9a, "unexpected bit string value")
}

func Test_decodeBitString2(t *testing.T) {
	value := [8]byte{0x9A, 0xBC, 0xD4, 0x00, 0x00, 0x00, 0x00, 0x00}
	bsBytes := make([]byte, 20) // 8 for a 64bit address, 8 for the size uint64, 4 for the unused bits
	binary.LittleEndian.PutUint64(bsBytes, uint64(uintptr(unsafe.Pointer(&value))))
	bsBytes[8] = 8 // num bytes = num bits / 8
	bsBytes[16] = 42 // unused bits

	bytes20 := [20]byte{}
	for i, b := range bsBytes {
		bytes20[i] = b
	}

	protoBitString, err := decodeBitString(bytes20)
	assert.NilError(t, err)
	assert.Assert(t, protoBitString != nil)
	assert.Equal(t, int(protoBitString.Len), 22, "unexpected bit string length")
	assert.Equal(t, protoBitString.Value, uint64(0xd4bc9a), "unexpected bit string value")
}
