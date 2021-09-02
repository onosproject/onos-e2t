// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "CauseProtocolVone.h"
import "C"
import (
	"fmt"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	"unsafe"
)

func xerEncodeCauseProtocol(causeProtocol *e2ap_ies.CauseProtocol) ([]byte, error) {
	causeProtocolCP, err := newCauseProtocol(causeProtocol)
	if err != nil {
		return nil, err
	}

	bytes, err := encodeXer(&C.asn_DEF_CauseProtocolVone, unsafe.Pointer(causeProtocolCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeCauseProtocol() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeCauseProtocol(causeProtocol *e2ap_ies.CauseProtocol) ([]byte, error) {
	causeProtocolCP, err := newCauseProtocol(causeProtocol)
	if err != nil {
		return nil, err
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_CauseProtocolVone, unsafe.Pointer(causeProtocolCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeCauseProtocol() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeCauseProtocol(bytes []byte) (*e2ap_ies.CauseProtocol, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_CauseProtocolVone)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeCauseProtocol((*C.CauseProtocolVone_t)(unsafePtr))
}

func perDecodeCauseProtocol(bytes []byte) (*e2ap_ies.CauseProtocol, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_CauseProtocolVone)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeCauseProtocol((*C.CauseProtocolVone_t)(unsafePtr))
}

func newCauseProtocol(causeProtocol *e2ap_ies.CauseProtocol) (*C.CauseProtocolVone_t, error) {
	var ret C.CauseProtocolVone_t
	switch *causeProtocol {
	case e2ap_ies.CauseProtocol_CAUSE_PROTOCOL_TRANSFER_SYNTAX_ERROR:
		ret = C.CauseProtocolVone_transfer_syntax_error
	case e2ap_ies.CauseProtocol_CAUSE_PROTOCOL_ABSTRACT_SYNTAX_ERROR_REJECT:
		ret = C.CauseProtocolVone_abstract_syntax_error_reject
	case e2ap_ies.CauseProtocol_CAUSE_PROTOCOL_ABSTRACT_SYNTAX_ERROR_IGNORE_AND_NOTIFY:
		ret = C.CauseProtocolVone_abstract_syntax_error_ignore_and_notify
	case e2ap_ies.CauseProtocol_CAUSE_PROTOCOL_MESSAGE_NOT_COMPATIBLE_WITH_RECEIVER_STATE:
		ret = C.CauseProtocolVone_message_not_compatible_with_receiver_state
	case e2ap_ies.CauseProtocol_CAUSE_PROTOCOL_SEMANTIC_ERROR:
		ret = C.CauseProtocolVone_semantic_error
	case e2ap_ies.CauseProtocol_CAUSE_PROTOCOL_ABSTRACT_SYNTAX_ERROR_FALSELY_CONSTRUCTED_MESSAGE:
		ret = C.CauseProtocolVone_abstract_syntax_error_falsely_constructed_message
	case e2ap_ies.CauseProtocol_CAUSE_PROTOCOL_UNSPECIFIED:
		ret = C.CauseProtocolVone_unspecified
	default:
		return nil, fmt.Errorf("unexpected CauseProtocol %v", causeProtocol)
	}

	return &ret, nil
}

func decodeCauseProtocol(causeProtocolC *C.CauseProtocolVone_t) (*e2ap_ies.CauseProtocol, error) {

	causeProtocol := e2ap_ies.CauseProtocol(int32(*causeProtocolC))

	return &causeProtocol, nil
}
