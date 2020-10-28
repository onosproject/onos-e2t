// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package orane2

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "GlobalgNB-ID.h"
import "C"
import (
	"fmt"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2ap-commondatatypes"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2apies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2proxy/e2ctypes"
)

// Deprecated: Do not use.
func newGlobalgNBIDOld(id *e2ctypes.GlobalgNB_IDT) (*C.GlobalgNB_ID_t, error) {
	if len(id.PlmnId) > 3 {
		return nil, fmt.Errorf("plmnID max length is 3 - e2ap-v01.00.00.asn line 1105")
	}

	gnbChoiceC, err := newGnbIDChoiceOld(id.GnbId)
	if err != nil {
		return nil, err
	}

	idC := C.GlobalgNB_ID_t{
		plmn_id: *newOctetString(id.PlmnId),
		gnb_id:  *gnbChoiceC,
	}

	return &idC, nil
}

// Deprecated: Do not use.
func decodeGlobalGnbIDOld(globalGnbID *C.GlobalgNB_ID_t) (*e2ctypes.GlobalgNB_IDT, error) {
	result := new(e2ctypes.GlobalgNB_IDT)
	var err error
	result.PlmnId = decodeOctetString(&globalGnbID.plmn_id)
	result.GnbId, err = decodeGnbIDChoiceOld(&globalGnbID.gnb_id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func newGlobalgNBID(id *e2apies.GlobalgNbId) (*C.GlobalgNB_ID_t, error) {
	if len(id.PlmnId.Value) > 3 {
		return nil, fmt.Errorf("plmnID max length is 3 - e2ap-v01.00.00.asn line 1105")
	}

	gnbChoiceC, err := newGnbIDChoice(id.GnbId)
	if err != nil {
		return nil, err
	}

	idC := C.GlobalgNB_ID_t{
		plmn_id: *newOctetString(string(id.PlmnId.Value)),
		gnb_id:  *gnbChoiceC,
	}

	return &idC, nil
}

func decodeGlobalGnbID(globalGnbID *C.GlobalgNB_ID_t) (*e2apies.GlobalgNbId, error) {
	result := new(e2apies.GlobalgNbId)
	result.PlmnId = new(e2ap_commondatatypes.PlmnIdentity)
	var err error
	result.PlmnId.Value = []byte(decodeOctetString(&globalGnbID.plmn_id))
	result.GnbId, err = decodeGnbIDChoice(&globalGnbID.gnb_id)
	if err != nil {
		return nil, err
	}
	return result, nil
}
