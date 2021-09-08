// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "PLMN-Identity.h"
import "C"

import (
	"fmt"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-commondatatypes"
	"unsafe"
)

func xerEncodePlmnIdentity(plmnIdentity *e2ap_commondatatypes.PlmnIdentity) ([]byte, error) {
	plmnIdentityCP, err := newPlmnIdentity(plmnIdentity)
	if err != nil {
		return nil, fmt.Errorf("xerEncodePlmnIdentity() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_PLMN_Identity, unsafe.Pointer(plmnIdentityCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodePlmnIdentity() %s", err.Error())
	}
	return bytes, nil
}

func perEncodePlmnIdentity(plmnIdentity *e2ap_commondatatypes.PlmnIdentity) ([]byte, error) {
	plmnIdentityCP, err := newPlmnIdentity(plmnIdentity)
	if err != nil {
		return nil, fmt.Errorf("perEncodePlmnIdentity() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_PLMN_Identity, unsafe.Pointer(plmnIdentityCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodePlmnIdentity() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodePlmnIdentity(bytes []byte) (*e2ap_commondatatypes.PlmnIdentity, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_PLMN_Identity)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodePlmnIdentity((*C.PLMN_Identity_t)(unsafePtr))
}

func perDecodePlmnIdentity(bytes []byte) (*e2ap_commondatatypes.PlmnIdentity, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_PLMN_Identity)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodePlmnIdentity((*C.PLMN_Identity_t)(unsafePtr))
}

func newPlmnIdentity(plmnIdentity *e2ap_commondatatypes.PlmnIdentity) (*C.PLMN_Identity_t, error) {

	plmnIdentityC := newOctetString(string(plmnIdentity.Value))

	return plmnIdentityC, nil
}

func decodePlmnIdentity(plmnIdentityC *C.PLMN_Identity_t) (*e2ap_commondatatypes.PlmnIdentity, error) {

	plmnID := decodeOctetString(plmnIdentityC)

	plmnIdentity := e2ap_commondatatypes.PlmnIdentity{
		Value: []byte(plmnID),
	}

	return &plmnIdentity, nil
}
