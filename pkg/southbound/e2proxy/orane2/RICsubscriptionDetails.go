// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package orane2

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "RICsubscriptionDetails.h"
//#include "RICaction-ToBeSetup-Item.h"
import "C"
import (
	"github.com/onosproject/onos-e2t/pkg/southbound/e2proxy/e2ctypes"
	"unsafe"
)

// XerEncodeRICsubscriptionDetails - for testing only
// Deprecated: Do not use.
func XerEncodeRICsubscriptionDetails(rsd *e2ctypes.RICsubscriptionDetailsT) ([]byte, error) {
	ratC, err := newRicSubscriptionDetails(rsd)
	if err != nil {
		return nil, err
	}

	bytes, err := encodeXer(&C.asn_DEF_RICsubscriptionDetails, unsafe.Pointer(ratC))
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

// PerEncodeRICsubscriptionDetails - for testing only
// Deprecated: Do not use.
func PerEncodeRICsubscriptionDetails(rsd *e2ctypes.RICsubscriptionDetailsT) ([]byte, error) {
	ratC, err := newRicSubscriptionDetails(rsd)
	if err != nil {
		return nil, err
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_RICsubscriptionDetails, unsafe.Pointer(ratC))
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

// Deprecated: Do not use.
func newRicSubscriptionDetails(rsd *e2ctypes.RICsubscriptionDetailsT) (*C.RICsubscriptionDetails_t, error) {
	ricactionTobesetupList := C.RICactions_ToBeSetup_List_t{}
	for _, tbsItem := range rsd.RicAction_ToBeSetup_List.List {
		tbsItemIEsC, err := newRICactionTbsItemIEs(tbsItem)
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(&ricactionTobesetupList), unsafe.Pointer(tbsItemIEsC)); err != nil {
			return nil, err
		}
	}

	rsdC := C.RICsubscriptionDetails_t{
		ricEventTriggerDefinition: *newOctetString(string(rsd.RicEventTriggerDefinition)),
		ricAction_ToBeSetup_List:  ricactionTobesetupList,
	}

	return &rsdC, nil
}
