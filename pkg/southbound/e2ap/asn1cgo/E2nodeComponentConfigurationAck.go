// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2nodeComponentConfigurationAck.h"
import "C"

import (
	"fmt"
	"unsafe"

	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
)

func xerEncodeE2nodeComponentConfigurationAck(e2nodeComponentConfigurationAck *e2ap_ies.E2NodeComponentConfigurationAck) ([]byte, error) {
	e2nodeComponentConfigurationAckCP, err := newE2nodeComponentConfigurationAck(e2nodeComponentConfigurationAck)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2nodeComponentConfigurationAck() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2nodeComponentConfigurationAck, unsafe.Pointer(e2nodeComponentConfigurationAckCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2nodeComponentConfigurationAck() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeE2nodeComponentConfigurationAck(e2nodeComponentConfigurationAck *e2ap_ies.E2NodeComponentConfigurationAck) ([]byte, error) {
	e2nodeComponentConfigurationAckCP, err := newE2nodeComponentConfigurationAck(e2nodeComponentConfigurationAck)
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2nodeComponentConfigurationAck() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2nodeComponentConfigurationAck, unsafe.Pointer(e2nodeComponentConfigurationAckCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2nodeComponentConfigurationAck() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeE2nodeComponentConfigurationAck(bytes []byte) (*e2ap_ies.E2NodeComponentConfigurationAck, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2nodeComponentConfigurationAck)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2nodeComponentConfigurationAck((*C.E2nodeComponentConfigurationAck_t)(unsafePtr))
}

func perDecodeE2nodeComponentConfigurationAck(bytes []byte) (*e2ap_ies.E2NodeComponentConfigurationAck, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2nodeComponentConfigurationAck)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2nodeComponentConfigurationAck((*C.E2nodeComponentConfigurationAck_t)(unsafePtr))
}

func newE2nodeComponentConfigurationAck(e2nodeComponentConfigurationAck *e2ap_ies.E2NodeComponentConfigurationAck) (*C.E2nodeComponentConfigurationAck_t, error) {

	var uoC C.E2nodeComponentConfigurationAck__updateOutcome_t
	switch e2nodeComponentConfigurationAck.GetUpdateOutcome() {
	case e2ap_ies.UpdateOutcome_UPDATE_OUTCOME_SUCCESS:
		uoC = C.E2nodeComponentConfigurationAck__updateOutcome_success
	case e2ap_ies.UpdateOutcome_UPDATE_OUTCOME_FAILURE:
		uoC = C.E2nodeComponentConfigurationAck__updateOutcome_failure
	default:
		return nil, fmt.Errorf("unexpected UpdateOutcome %v", e2nodeComponentConfigurationAck.GetUpdateOutcome())
	}

	e2nodeComponentConfigurationAckC := C.E2nodeComponentConfigurationAck_t{
		updateOutcome: uoC,
		//failureCause:  fc,
	}

	if e2nodeComponentConfigurationAck.GetFailureCause() != nil {
		fc, err := newCause(e2nodeComponentConfigurationAck.GetFailureCause())
		if err != nil {
			return nil, err
		}
		e2nodeComponentConfigurationAckC.failureCause = fc
	}

	return &e2nodeComponentConfigurationAckC, nil
}

func decodeE2nodeComponentConfigurationAck(e2nodeComponentConfigurationAckC *C.E2nodeComponentConfigurationAck_t) (*e2ap_ies.E2NodeComponentConfigurationAck, error) {

	var err error
	e2nodeComponentConfigurationAck := e2ap_ies.E2NodeComponentConfigurationAck{
		UpdateOutcome: e2ap_ies.UpdateOutcome(int32(e2nodeComponentConfigurationAckC.updateOutcome)),
	}

	if e2nodeComponentConfigurationAckC.failureCause != nil {
		e2nodeComponentConfigurationAck.FailureCause, err = decodeCause(e2nodeComponentConfigurationAckC.failureCause)
		if err != nil {
			return nil, err
		}
	}

	return &e2nodeComponentConfigurationAck, nil
}
