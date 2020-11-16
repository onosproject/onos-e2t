// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

// #cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
// #cgo LDFLAGS: -lm
// #include <stdio.h>
// #include <stdlib.h>
// #include <assert.h>
// #include "RICsubscriptionDetails.h"
import "C"
import "github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2appducontents"

func newRicSubscriptionDetails(rsDet *e2appducontents.RicsubscriptionDetails) (*C.RICsubscriptionDetails_t, error) {

	raTbsL, err := newRicActionToBeSetupList(rsDet.RicActionToBeSetupList)
	if err != nil {
		return nil, err
	}

	rsDetC := C.RICsubscriptionDetails_t{
		ricEventTriggerDefinition: *newRicEventTriggerDefinition(rsDet.GetRicEventTriggerDefinition()),
		ricAction_ToBeSetup_List:  *raTbsL,
	}

	return &rsDetC, nil
}
