// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "CauseRICrequest.h"
import "C"
import (
	"fmt"
	"unsafe"

	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
)

func xerEncodeCauseRic(causeRic *e2ap_ies.CauseRicrequest) ([]byte, error) {
	causeRicCP, err := newCauseRic(causeRic)
	if err != nil {
		return nil, err
	}

	bytes, err := encodeXer(&C.asn_DEF_CauseRICrequest, unsafe.Pointer(causeRicCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeCauseRic() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeCauseRic(causeRic *e2ap_ies.CauseRicrequest) ([]byte, error) {
	causeRicCP, err := newCauseRic(causeRic)
	if err != nil {
		return nil, err
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_CauseRICrequest, unsafe.Pointer(causeRicCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeCauseRic() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeCauseRic(bytes []byte) (*e2ap_ies.CauseRicrequest, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_CauseRICrequest)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeCauseRic((*C.CauseRICrequest_t)(unsafePtr))
}

func perDecodeCauseRic(bytes []byte) (*e2ap_ies.CauseRicrequest, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_CauseRICrequest)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeCauseRic((*C.CauseRICrequest_t)(unsafePtr))
}

func newCauseRic(causeRic *e2ap_ies.CauseRicrequest) (*C.CauseRICrequest_t, error) {
	var ret C.CauseRICrequest_t
	switch *causeRic {
	case e2ap_ies.CauseRicrequest_CAUSE_RICREQUEST_RAN_FUNCTION_ID_INVALID:
		ret = C.CauseRICrequest_ran_function_id_invalid
	case e2ap_ies.CauseRicrequest_CAUSE_RICREQUEST_ACTION_NOT_SUPPORTED:
		ret = C.CauseRICrequest_action_not_supported
	case e2ap_ies.CauseRicrequest_CAUSE_RICREQUEST_EXCESSIVE_ACTIONS:
		ret = C.CauseRICrequest_excessive_actions
	case e2ap_ies.CauseRicrequest_CAUSE_RICREQUEST_DUPLICATE_ACTION:
		ret = C.CauseRICrequest_duplicate_action
	case e2ap_ies.CauseRicrequest_CAUSE_RICREQUEST_DUPLICATE_EVENT_TRIGGER:
		ret = C.CauseRICrequest_duplicate_event_trigger
	case e2ap_ies.CauseRicrequest_CAUSE_RICREQUEST_FUNCTION_RESOURCE_LIMIT:
		ret = C.CauseRICrequest_function_resource_limit
	case e2ap_ies.CauseRicrequest_CAUSE_RICREQUEST_REQUEST_ID_UNKNOWN:
		ret = C.CauseRICrequest_request_id_unknown
	case e2ap_ies.CauseRicrequest_CAUSE_RICREQUEST_INCONSISTENT_ACTION_SUBSEQUENT_ACTION_SEQUENCE:
		ret = C.CauseRICrequest_inconsistent_action_subsequent_action_sequence
	case e2ap_ies.CauseRicrequest_CAUSE_RICREQUEST_CONTROL_MESSAGE_INVALID:
		ret = C.CauseRICrequest_control_message_invalid
	case e2ap_ies.CauseRicrequest_CAUSE_RICREQUEST_RIC_CALL_PROCESS_ID_INVALID:
		ret = C.CauseRICrequest_ric_call_process_id_invalid
	case e2ap_ies.CauseRicrequest_CAUSE_RICREQUEST_CONTROL_TIMER_EXPIRED:
		ret = C.CauseRICrequest_control_timer_expired
	case e2ap_ies.CauseRicrequest_CAUSE_RICREQUEST_CONTROL_FAILED_TO_EXECUTE:
		ret = C.CauseRICrequest_control_failed_to_execute
	case e2ap_ies.CauseRicrequest_CAUSE_RICREQUEST_SYSTEM_NOT_READY:
		ret = C.CauseRICrequest_system_not_ready
	case e2ap_ies.CauseRicrequest_CAUSE_RICREQUEST_UNSPECIFIED:
		ret = C.CauseRICrequest_unspecified
	default:
		return nil, fmt.Errorf("unexpected CauseRic %v", causeRic)
	}

	return &ret, nil
}

func decodeCauseRic(causeRicC *C.CauseRICrequest_t) (*e2ap_ies.CauseRicrequest, error) {

	causeRic := e2ap_ies.CauseRicrequest(int32(*causeRicC))

	return &causeRic, nil
}
