// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "TriggeringMessageVone.h"
import "C"
import (
	"fmt"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-commondatatypes"
)

func newTriggeringMessage(tm e2ap_commondatatypes.TriggeringMessage) (C.TriggeringMessageVone_t, error) {
	var ret C.TriggeringMessageVone_t
	switch tm {
	case e2ap_commondatatypes.TriggeringMessage_TRIGGERING_MESSAGE_INITIATING_MESSAGE:
		ret = C.TriggeringMessageVone_initiating_message
	case e2ap_commondatatypes.TriggeringMessage_TRIGGERING_MESSAGE_SUCCESSFUL_OUTCOME:
		ret = C.TriggeringMessageVone_successful_outcome
	case e2ap_commondatatypes.TriggeringMessage_TRIGGERING_MESSAGE_UNSUCCESSFULL_OUTCOME:
		ret = C.TriggeringMessageVone_unsuccessfull_outcome
	default:
		return 0, fmt.Errorf("unexpected TriggeringMessage %v", tm)
	}
	return ret, nil
}

func decodeTriggeringMessage(tmC C.TriggeringMessageVone_t) e2ap_commondatatypes.TriggeringMessage {
	return e2ap_commondatatypes.TriggeringMessage(tmC)
}
