// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2nodeComponentConfiguration.h"
import "C"

import (
	"fmt"
	"unsafe"

	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
)

func xerEncodeE2nodeComponentConfiguration(e2nodeComponentConfiguration *e2ap_ies.E2NodeComponentConfiguration) ([]byte, error) {
	e2nodeComponentConfigurationCP, err := newE2nodeComponentConfiguration(e2nodeComponentConfiguration)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2nodeComponentConfiguration() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2nodeComponentConfiguration, unsafe.Pointer(e2nodeComponentConfigurationCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2nodeComponentConfiguration() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeE2nodeComponentConfiguration(e2nodeComponentConfiguration *e2ap_ies.E2NodeComponentConfiguration) ([]byte, error) {
	e2nodeComponentConfigurationCP, err := newE2nodeComponentConfiguration(e2nodeComponentConfiguration)
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2nodeComponentConfiguration() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2nodeComponentConfiguration, unsafe.Pointer(e2nodeComponentConfigurationCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2nodeComponentConfiguration() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeE2nodeComponentConfiguration(bytes []byte) (*e2ap_ies.E2NodeComponentConfiguration, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2nodeComponentConfiguration)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2nodeComponentConfiguration((*C.E2nodeComponentConfiguration_t)(unsafePtr))
}

func perDecodeE2nodeComponentConfiguration(bytes []byte) (*e2ap_ies.E2NodeComponentConfiguration, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2nodeComponentConfiguration)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2nodeComponentConfiguration((*C.E2nodeComponentConfiguration_t)(unsafePtr))
}

func newE2nodeComponentConfiguration(e2nodeComponentConfiguration *e2ap_ies.E2NodeComponentConfiguration) (*C.E2nodeComponentConfiguration_t, error) {

	e2nodeComponentConfigurationC := C.E2nodeComponentConfiguration_t{
		e2nodeComponentRequestPart:  *newOctetString(e2nodeComponentConfiguration.GetE2NodeComponentRequestPart()),
		e2nodeComponentResponsePart: *newOctetString(e2nodeComponentConfiguration.GetE2NodeComponentResponsePart()),
	}

	return &e2nodeComponentConfigurationC, nil
}

func decodeE2nodeComponentConfiguration(e2nodeComponentConfigurationC *C.E2nodeComponentConfiguration_t) (*e2ap_ies.E2NodeComponentConfiguration, error) {

	e2nodeComponentConfiguration := e2ap_ies.E2NodeComponentConfiguration{
		E2NodeComponentRequestPart:  decodeOctetString(&e2nodeComponentConfigurationC.e2nodeComponentRequestPart),
		E2NodeComponentResponsePart: decodeOctetString(&e2nodeComponentConfigurationC.e2nodeComponentResponsePart),
	}

	return &e2nodeComponentConfiguration, nil
}
