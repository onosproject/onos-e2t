// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "RICserviceUpdateAcknowledge.h"
import "C"

import (
	"fmt"
	e2ap_pdu_contents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
	"unsafe"
)

func xerEncodeRicServiceUpdateAcknowledge(rsua *e2ap_pdu_contents.RicserviceUpdateAcknowledge) ([]byte, error) {
	rsuaCP, err := newRicServiceUpdateAcknowledge(rsua)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeRicServiceUpdateAcknowledge() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_RICserviceUpdateAcknowledge, unsafe.Pointer(rsuaCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeRicServiceUpdateAcknowledge() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeRicServiceUpdateAcknowledge(rsua *e2ap_pdu_contents.RicserviceUpdateAcknowledge) ([]byte, error) {
	rsuaCP, err := newRicServiceUpdateAcknowledge(rsua)
	if err != nil {
		return nil, fmt.Errorf("perEncodeRicServiceUpdateAcknowledge() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_RICserviceUpdateAcknowledge, unsafe.Pointer(rsuaCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeRicServiceUpdateAcknowledge() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeRicServiceUpdateAcknowledge(bytes []byte) (*e2ap_pdu_contents.RicserviceUpdateAcknowledge, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_RICserviceUpdateAcknowledge)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeRicServiceUpdateAcknowledge((*C.RICserviceUpdateAcknowledge_t)(unsafePtr))
}

func perDecodeRicServiceUpdateAcknowledge(bytes []byte) (*e2ap_pdu_contents.RicserviceUpdateAcknowledge, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_RICserviceUpdateAcknowledge)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeRicServiceUpdateAcknowledge((*C.RICserviceUpdateAcknowledge_t)(unsafePtr))
}

func newRicServiceUpdateAcknowledge(rsua *e2ap_pdu_contents.RicserviceUpdateAcknowledge) (*C.RICserviceUpdateAcknowledge_t, error) {

	pIeC1710P23, err := newRicServiceUpdateAcknowledgeIe(rsua.ProtocolIes)
	if err != nil {
		return nil, err
	}
	rsuaC := C.RICserviceUpdateAcknowledge_t{
		protocolIEs: *pIeC1710P23,
	}

	return &rsuaC, nil
}

func decodeRicServiceUpdateAcknowledge(rsuaC *C.RICserviceUpdateAcknowledge_t) (*e2ap_pdu_contents.RicserviceUpdateAcknowledge, error) {

	pIEs, err := decodeRicServiceUpdateAcknowledgeIes(&rsuaC.protocolIEs)
	if err != nil {
		return nil, err
	}

	rsua := e2ap_pdu_contents.RicserviceUpdateAcknowledge{
		ProtocolIes: pIEs,
	}

	return &rsua, nil
}
