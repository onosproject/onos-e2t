// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "RICserviceQuery.h"
import "C"

import (
	"fmt"
	e2ap_pdu_contents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
	"unsafe"
)

func xerEncodeRicServiceQuery(rsq *e2ap_pdu_contents.RicserviceQuery) ([]byte, error) {
	rsqCP, err := newRicServiceQuery(rsq)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeRicServiceQuery() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_RICserviceQuery, unsafe.Pointer(rsqCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeRicServiceQuery() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeRicServiceQuery(rsq *e2ap_pdu_contents.RicserviceQuery) ([]byte, error) {
	rsqCP, err := newRicServiceQuery(rsq)
	if err != nil {
		return nil, fmt.Errorf("perEncodeRicServiceQuery() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_RICserviceQuery, unsafe.Pointer(rsqCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeRicServiceQuery() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeRicServiceQuery(bytes []byte) (*e2ap_pdu_contents.RicserviceQuery, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_RICserviceQuery)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeRicServiceQuery((*C.RICserviceQuery_t)(unsafePtr))
}

func perDecodeRicServiceQuery(bytes []byte) (*e2ap_pdu_contents.RicserviceQuery, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_RICserviceQuery)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeRicServiceQuery((*C.RICserviceQuery_t)(unsafePtr))
}

func newRicServiceQuery(rsq *e2ap_pdu_contents.RicserviceQuery) (*C.RICserviceQuery_t, error) {

	pIeC1710P25, err := newRicServiceQueryIe(rsq.ProtocolIes)
	if err != nil {
		return nil, err
	}
	rsqC := C.RICserviceQuery_t{
		protocolIEs: *pIeC1710P25,
	}

	return &rsqC, nil
}

func decodeRicServiceQuery(rsqC *C.RICserviceQuery_t) (*e2ap_pdu_contents.RicserviceQuery, error) {

	pIEs, err := decodeRicServiceQueryIes(&rsqC.protocolIEs)
	if err != nil {
		return nil, err
	}

	rsq := e2ap_pdu_contents.RicserviceQuery{
		ProtocolIes: pIEs,
	}

	return &rsq, nil
}
