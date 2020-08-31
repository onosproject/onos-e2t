// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package orane2

// #cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
// #cgo LDFLAGS: -lm
// #include <stdio.h>
// #include <stdlib.h>
// #include <assert.h>
// #include "RICrequestID.h"
import "C"
import (
	"encoding/binary"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2proxy/e2ctypes"
)

// XerDecodeRICrequestID - just used for test only
func XerDecodeRICrequestID(bytes []byte) (*e2ctypes.RICrequestIDT, error) {
	resultUnsafe, err := decodeXer(bytes, &C.asn_DEF_RICrequestID)
	if err != nil {
		return nil, err
	}

	resultC := (*C.RICrequestID_t)(resultUnsafe)
	result := e2ctypes.RICrequestIDT{
		RicRequestorID: int64(resultC.ricRequestorID),
		RicInstanceID:  int64(resultC.ricInstanceID),
	}

	return &result, nil
}

func decodeRicRequestID(ricRequestIDCchoice []byte) (*e2ctypes.RICrequestIDT, error) {
	ricRequestorID := binary.LittleEndian.Uint64(ricRequestIDCchoice[0:8])
	ricInstanceID := binary.LittleEndian.Uint64(ricRequestIDCchoice[8:16])

	result := e2ctypes.RICrequestIDT{
		RicRequestorID: int64(ricRequestorID),
		RicInstanceID:  int64(ricInstanceID),
	}

	return &result, nil
}
