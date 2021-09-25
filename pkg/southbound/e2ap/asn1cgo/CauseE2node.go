// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "CauseE2node.h"
import "C"
import (
	"fmt"
	"unsafe"

	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
)

func xerEncodeCauseE2node(causeE2node *e2ap_ies.CauseE2Node) ([]byte, error) {
	causeE2nodeCP, err := newCauseE2node(causeE2node)
	if err != nil {
		return nil, err
	}

	bytes, err := encodeXer(&C.asn_DEF_CauseE2node, unsafe.Pointer(causeE2nodeCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeCauseMisc() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeCauseE2node(causeE2node *e2ap_ies.CauseE2Node) ([]byte, error) {
	causeE2nodeCP, err := newCauseE2node(causeE2node)
	if err != nil {
		return nil, err
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_CauseE2node, unsafe.Pointer(causeE2nodeCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeCauseMisc() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeCauseE2node(bytes []byte) (*e2ap_ies.CauseE2Node, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_CauseE2node)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeCauseE2node((*C.CauseE2node_t)(unsafePtr))
}

func perDecodeCauseE2node(bytes []byte) (*e2ap_ies.CauseE2Node, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_CauseE2node)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeCauseE2node((*C.CauseE2node_t)(unsafePtr))
}

func newCauseE2node(causeE2node *e2ap_ies.CauseE2Node) (*C.CauseE2node_t, error) {
	var ret C.CauseE2node_t
	switch *causeE2node {
	case e2ap_ies.CauseE2Node_CAUSE_E2NODE_E2NODE_COMPONENT_UNKNOWN:
		ret = C.CauseE2node_e2node_component_unknown
	default:
		return nil, fmt.Errorf("unexpected CauseE2node %v", causeE2node)
	}

	return &ret, nil
}

func decodeCauseE2node(causeE2nodeC *C.CauseE2node_t) (*e2ap_ies.CauseE2Node, error) {

	causeE2node := e2ap_ies.CauseE2Node(int32(*causeE2nodeC))

	return &causeE2node, nil
}
