// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package orane2

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "GlobalE2node-gNB-ID.h"
import "C"
import (
	"fmt"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2apies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2proxy/e2ctypes"
)

// Deprecated: Do not use.
func newGlobalE2nodegNBIDOld(gnbID *e2ctypes.GlobalE2NodeGNB_ID) (*C.GlobalE2node_gNB_ID_t, error) {

	globalgNBID, err := newGlobalgNBIDOld(gnbID.GlobalGNB_ID)
	if err != nil {
		return nil, err
	}

	globalgNBIDC := C.GlobalE2node_gNB_ID_t{
		global_gNB_ID: *globalgNBID,
		gNB_CU_UP_ID:  nil,
		gNB_DU_ID:     nil,
	}

	return &globalgNBIDC, nil
}

// Deprecated: Do not use.
func decodeGlobalE2nodegNBIDOld(gNBC *C.GlobalE2node_gNB_ID_t) (*e2ctypes.GlobalE2NodeGNB_ID, error) {
	result := new(e2ctypes.GlobalE2NodeGNB_ID)
	var err error
	result.GlobalGNB_ID, err = decodeGlobalGnbIDOld(&gNBC.global_gNB_ID)
	if err != nil {
		return nil, fmt.Errorf("error decodeGlobalE2nodegNBIDOld() %v", err)
	}

	return result, nil
}

func newGlobalE2nodegNBID(gnbID *e2apies.GlobalE2NodeGnbId) (*C.GlobalE2node_gNB_ID_t, error) {

	globalgNBID, err := newGlobalgNBID(gnbID.GlobalGNbId)
	if err != nil {
		return nil, err
	}

	globalgNBIDC := C.GlobalE2node_gNB_ID_t{
		global_gNB_ID: *globalgNBID,
		gNB_CU_UP_ID:  nil,
		gNB_DU_ID:     nil,
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

	return result, nil
}
