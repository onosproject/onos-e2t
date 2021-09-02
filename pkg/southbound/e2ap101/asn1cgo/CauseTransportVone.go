// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "CauseTransportVone.h"
import "C"
import (
	"fmt"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	"unsafe"
)

func xerEncodeCauseTransport(causeTransport *e2ap_ies.CauseTransport) ([]byte, error) {
	causeTransportCP, err := newCauseTransport(causeTransport)
	if err != nil {
		return nil, err
	}

	bytes, err := encodeXer(&C.asn_DEF_CauseTransportVone, unsafe.Pointer(causeTransportCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeCauseTransport() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeCauseTransport(causeTransport *e2ap_ies.CauseTransport) ([]byte, error) {
	causeTransportCP, err := newCauseTransport(causeTransport)
	if err != nil {
		return nil, err
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_CauseTransportVone, unsafe.Pointer(causeTransportCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeCauseTransport() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeCauseTransport(bytes []byte) (*e2ap_ies.CauseTransport, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_CauseTransportVone)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeCauseTransport((*C.CauseTransportVone_t)(unsafePtr))
}

func perDecodeCauseTransport(bytes []byte) (*e2ap_ies.CauseTransport, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_CauseTransportVone)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeCauseTransport((*C.CauseTransportVone_t)(unsafePtr))
}

func newCauseTransport(causeTransport *e2ap_ies.CauseTransport) (*C.CauseTransportVone_t, error) {
	var ret C.CauseTransportVone_t
	switch *causeTransport {
	case e2ap_ies.CauseTransport_CAUSE_TRANSPORT_UNSPECIFIED:
		ret = C.CauseTransportVone_unspecified
	case e2ap_ies.CauseTransport_CAUSE_TRANSPORT_TRANSPORT_RESOURCE_UNAVAILABLE:
		ret = C.CauseTransportVone_transport_resource_unavailable
	default:
		return nil, fmt.Errorf("unexpected CauseTransport %v", causeTransport)
	}

	return &ret, nil
}

func decodeCauseTransport(causeTransportC *C.CauseTransportVone_t) (*e2ap_ies.CauseTransport, error) {

	causeTransport := e2ap_ies.CauseTransport(int32(*causeTransportC))

	return &causeTransport, nil
}
