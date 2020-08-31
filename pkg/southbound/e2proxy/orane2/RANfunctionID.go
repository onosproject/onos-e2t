// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package orane2

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
import "C"
import (
	"encoding/binary"
)

func decodeRanFunctionID(ranFunctionIDCchoice [64]byte) (*uint64, error) {
	result := binary.LittleEndian.Uint64(ranFunctionIDCchoice[0:8])

	return &result, nil
}
