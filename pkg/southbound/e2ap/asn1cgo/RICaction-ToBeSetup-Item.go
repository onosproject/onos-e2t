// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "RICaction-ToBeSetup-Item.h"
import "C"
import (
	"fmt"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2appducontents"
)

func newRicActionToBeSetupItem(ratbsItem *e2appducontents.RicactionToBeSetupItem) (*C.RICaction_ToBeSetup_Item_t, error) {
	ratC, err := newRicActionType(ratbsItem.GetRicActionType())
	if err != nil {
		return nil, fmt.Errorf("newRicActionType() %s", err.Error())
	}

	rsaC, err := newRicSubsequentAction(ratbsItem.RicSubsequentAction)
	if err != nil {
		return nil, fmt.Errorf("newRicSubsequentAction() %s", err.Error())
	}

	ratbsItemC := C.RICaction_ToBeSetup_Item_t{
		ricActionID:         *newRicActionID(ratbsItem.GetRicActionId()),
		ricActionType:       *ratC,
		ricActionDefinition: newOctetString(string(ratbsItem.GetRicActionDefinition().GetValue())),
		ricSubsequentAction: rsaC,
	}

	return &ratbsItemC, nil
}
