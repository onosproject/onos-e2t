// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package orane2

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "ProtocolIE-Field.h"
//#include "RICaction-ToBeSetup-Item.h"
import "C"
import (
	"github.com/onosproject/onos-e2t/pkg/southbound/e2proxy/e2ctypes"
	"unsafe"
)

// XerEncodeRICactionToBeSetupItem - for testing only
func XerEncodeRICactionToBeSetupItem(tbs *e2ctypes.RICaction_ToBeSetup_ItemIEsT) ([]byte, error) {
	ratC, err := newRICactionTbsItemIEs(tbs)
	if err != nil {
		return nil, err
	}

	bytes, err := encodeXer(&C.asn_DEF_RICaction_ToBeSetup_ItemIEs, unsafe.Pointer(ratC))
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

// PerEncodeRICactionToBeSetupItem - for testing only
func PerEncodeRICactionToBeSetupItem(tbs *e2ctypes.RICaction_ToBeSetup_ItemIEsT) ([]byte, error) {
	ratC, err := newRICactionTbsItemIEs(tbs)
	if err != nil {
		return nil, err
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_RICaction_ToBeSetup_ItemIEs, unsafe.Pointer(ratC))
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func newRICactionToBeSetupItem(tbs *e2ctypes.RICaction_ToBeSetup_ItemT) (*C.RICaction_ToBeSetup_Item_t, error) {

	ricActionTypeC, err := newRicaActionType(tbs.GetRicActionType())
	if err != nil {
		return nil, err
	}

	tbsC := C.RICaction_ToBeSetup_Item_t{
		ricActionID:   C.long(tbs.GetRicActionID()),
		ricActionType: ricActionTypeC,
	}

	if len(tbs.RicActionDefinition) > 0 {
		def := newOctetString(tbs.RicActionDefinition)
		tbsC.ricActionDefinition = def
	}

	if tbs.RicSubsequentAction != nil {
		rsaC, err := newRicSubsequentAction(tbs.RicSubsequentAction)
		if err != nil {
			return nil, err
		}
		tbsC.ricSubsequentAction = rsaC
	}
	return &tbsC, nil
}
