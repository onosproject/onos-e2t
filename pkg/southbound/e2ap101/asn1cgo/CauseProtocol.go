// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "CauseProtocol.h"
import "C"
import (
	"encoding/binary"
	"fmt"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	"unsafe"
)

func xerEncodeCauseProtocol(causeProtocol *e2ap_ies.CauseProtocol) ([]byte, error) {
	causeProtocolCP, err := newCauseProtocol(causeProtocol)
	if err != nil {
		return nil, err
	}

	bytes, err := encodeXer(&C.asn_DEF_CauseProtocol, unsafe.Pointer(causeProtocolCP)) //ToDo - change name of C-encoder tag
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

	bytes, err := encodePerBuffer(&C.asn_DEF_CauseProtocol, unsafe.Pointer(causeProtocolCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeCauseProtocol() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeCauseProtocol(bytes []byte) (*e2ap_ies.CauseProtocol, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_CauseProtocol)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeCauseProtocol((*C.CauseProtocol_t)(unsafePtr)) //ToDo - change name of C-struct
}

func perDecodeCauseProtocol(bytes []byte) (*e2ap_ies.CauseProtocol, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_CauseProtocol)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeCauseProtocol((*C.CauseProtocol_t)(unsafePtr))
}

func newCauseProtocol(causeProtocol *e2ap_ies.CauseProtocol) (*C.CauseProtocol_t, error) {
	var ret C.CauseProtocol_t
	switch *causeProtocol {
	case e2ap_ies.CauseProtocol_CAUSE_PROTOCOL_TRANSFER_SYNTAX_ERROR:
		ret = C.CauseProtocol_transfer_syntax_error
	case e2ap_ies.CauseProtocol_CAUSE_PROTOCOL_ABSTRACT_SYNTAX_ERROR_REJECT:
		ret = C.CauseProtocol_abstract_syntax_error_reject
	case e2ap_ies.CauseProtocol_CAUSE_PROTOCOL_ABSTRACT_SYNTAX_ERROR_IGNORE_AND_NOTIFY:
		ret = C.CauseProtocol_abstract_syntax_error_ignore_and_notify
	case e2ap_ies.CauseProtocol_CAUSE_PROTOCOL_MESSAGE_NOT_COMPATIBLE_WITH_RECEIVER_STATE:
		ret = C.CauseProtocol_message_not_compatible_with_receiver_state
	case e2ap_ies.CauseProtocol_CAUSE_PROTOCOL_SEMANTIC_ERROR:
		ret = C.CauseProtocol_semantic_error
	case e2ap_ies.CauseProtocol_CAUSE_PROTOCOL_ABSTRACT_SYNTAX_ERROR_FALSELY_CONSTRUCTED_MESSAGE:
		ret = C.CauseProtocol_abstract_syntax_error_falsely_constructed_message
	case e2ap_ies.CauseProtocol_CAUSE_PROTOCOL_UNSPECIFIED:
		ret = C.CauseProtocol_unspecified
	default:
		return nil, fmt.Errorf("unexpected CauseProtocol %v", causeProtocol)
	}

	return &ret, nil
}

func decodeCauseProtocol(causeProtocolC *C.CauseProtocol_t) (*e2ap_ies.CauseProtocol, error) {

	//ToDo: int32 shouldn't be valid all the time -- investigate in data type conversion (casting) more
	causeProtocol := e2ap_ies.CauseProtocol(int32(*causeProtocolC))

	return &causeProtocol, nil
}

func decodeCauseProtocolBytes(array [8]byte) (*e2ap_ies.CauseProtocol, error) { //ToDo - Check addressing correct structure in Protobuf
	causeProtocolC := (*C.CauseProtocol_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:]))))

	return decodeCauseProtocol(causeProtocolC)
}
