// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

// #cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
// #cgo LDFLAGS: -lm
// #include <stdio.h>
// #include <stdlib.h>
// #include <assert.h>
// #include "RICactions-ToBeSetup-List.h"
import "C"
import (
	"fmt"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2appducontents"
	"unsafe"
)

func xerEncodeRicActionsToBeSetupList(ratbsl *e2appducontents.RicactionsToBeSetupList) ([]byte, error) {
	ratbslC, err := newRicActionToBeSetupList(ratbsl)

	if err != nil {
		return nil, err
	}

	bytes, err := encodeXer(&C.asn_DEF_RICactions_ToBeSetup_List, unsafe.Pointer(ratbslC))
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func perEncodeRicActionsToBeSetupList(ratbsl *e2appducontents.RicactionsToBeSetupList) ([]byte, error) {
	ratbslC, err := newRicActionToBeSetupList(ratbsl)

	if err != nil {
		return nil, err
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_RICactions_ToBeSetup_List, unsafe.Pointer(ratbslC))
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func newRicActionToBeSetupList(ratbsL *e2appducontents.RicactionsToBeSetupList) (*C.RICactions_ToBeSetup_List_t, error) {
	ratbsLC := new(C.RICactions_ToBeSetup_List_t)
	for _, ricActionToBeSetupItemIe := range ratbsL.GetValue() {
		ricActionToBeItemIesScC, err := newRicActionToBeSetupItemIesSingleContainer(ricActionToBeSetupItemIe)
		if err != nil {
			return nil, fmt.Errorf("newRanFunctionsList() %s", err.Error())
		}

		if _, err = C.asn_sequence_add(unsafe.Pointer(ratbsLC), unsafe.Pointer(ricActionToBeItemIesScC)); err != nil {
			return nil, err
		}
	}

	return ratbsLC, nil
}
