// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "CauseRIC.h"
import "C"
import (
	"fmt"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	"unsafe"
)

func xerEncodeCauseRic(causeRic *e2ap_ies.CauseRic) ([]byte, error) {
	causeRicCP, err := newCauseRic(causeRic)
	if err != nil {
		return nil, err
	}

	bytes, err := encodeXer(&C.asn_DEF_CauseRIC, unsafe.Pointer(causeRicCP)) //ToDo - change name of C-encoder tag
	if err != nil {
		return nil, fmt.Errorf("xerEncodeCauseRic() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeCauseRic(causeRic *e2ap_ies.CauseRic) ([]byte, error) {
	causeRicCP, err := newCauseRic(causeRic)
	if err != nil {
		return nil, err
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_CauseRIC, unsafe.Pointer(causeRicCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeCauseRic() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeCauseRic(bytes []byte) (*e2ap_ies.CauseRic, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_CauseRIC)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeCauseRic((*C.CauseRIC_t)(unsafePtr)) //ToDo - change name of C-struct
}

func perDecodeCauseRic(bytes []byte) (*e2ap_ies.CauseRic, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_CauseRIC)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeCauseRic((*C.CauseRIC_t)(unsafePtr))
}

func newCauseRic(causeRic *e2ap_ies.CauseRic) (*C.CauseRIC_t, error) {
	var ret C.CauseRIC_t
	switch *causeRic {
	case e2ap_ies.CauseRic_CAUSE_RIC_RAN_FUNCTION_ID_INVALID:
		ret = C.CauseRIC_ran_function_id_Invalid
	case e2ap_ies.CauseRic_CAUSE_RIC_ACTION_NOT_SUPPORTED:
		ret = C.CauseRIC_action_not_supported
	case e2ap_ies.CauseRic_CAUSE_RIC_EXCESSIVE_ACTIONS:
		ret = C.CauseRIC_excessive_actions
	case e2ap_ies.CauseRic_CAUSE_RIC_DUPLICATE_ACTION:
		ret = C.CauseRIC_duplicate_action
	case e2ap_ies.CauseRic_CAUSE_RIC_DUPLICATE_EVENT:
		ret = C.CauseRIC_duplicate_event
	case e2ap_ies.CauseRic_CAUSE_RIC_FUNCTION_RESOURCE_LIMIT:
		ret = C.CauseRIC_function_resource_limit
	case e2ap_ies.CauseRic_CAUSE_RIC_REQUEST_ID_UNKNOWN:
		ret = C.CauseRIC_request_id_unknown
	case e2ap_ies.CauseRic_CAUSE_RIC_INCONSISTENT_ACTION_SUBSEQUENT_ACTION_SEQUENCE:
		ret = C.CauseRIC_inconsistent_action_subsequent_action_sequence
	case e2ap_ies.CauseRic_CAUSE_RIC_CONTROL_MESSAGE_INVALID:
		ret = C.CauseRIC_control_message_invalid
	case e2ap_ies.CauseRic_CAUSE_RIC_CALL_PROCESS_ID_INVALID:
		ret = C.CauseRIC_call_process_id_invalid
	case e2ap_ies.CauseRic_CAUSE_RIC_UNSPECIFIED:
		ret = C.CauseRIC_unspecified
	default:
		return nil, fmt.Errorf("unexpected CauseRic %v", causeRic)
	}

	return &ret, nil
}

func decodeCauseRic(causeRicC *C.CauseRIC_t) (*e2ap_ies.CauseRic, error) {

	//ToDo: int32 shouldn't be valid all the time -- investigate in data type conversion (casting) more
	causeRic := e2ap_ies.CauseRic(int32(*causeRicC))

	return &causeRic, nil
}
