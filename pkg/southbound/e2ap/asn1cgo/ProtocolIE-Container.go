// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "ProtocolIE-Container.h"
//#include "ProtocolIE-Field.h"
import "C"
import (
	"fmt"
	"unsafe"

	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
)

func newE2SetupRequestIes(esv *e2appducontents.E2SetupRequestIes) (*C.ProtocolIE_Container_1710P11_t, error) {
	pIeC1710P11 := new(C.ProtocolIE_Container_1710P11_t)

	if esv.GetE2ApProtocolIes3() != nil {
		ie3C, err := newE2setupRequestIe3GlobalE2NodeID(esv.GetE2ApProtocolIes3())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1710P11), unsafe.Pointer(ie3C)); err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("newE2SetupRequestIes(): GlobalE2NodeID should be mandatory present in the message")
	}

	if esv.GetE2ApProtocolIes10() != nil {
		ie10C, err := newE2setupRequestIe10RanFunctionList(esv.GetE2ApProtocolIes10())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1710P11), unsafe.Pointer(ie10C)); err != nil {
			return nil, err
		}
	}

	if esv.GetE2ApProtocolIes33() != nil {
		ie33C, err := newE2setupRequestIe33E2nodeComponentConfigUpdateList(esv.GetE2ApProtocolIes33())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1710P11), unsafe.Pointer(ie33C)); err != nil {
			return nil, err
		}
	}

	return pIeC1710P11, nil
}

func decodeE2SetupRequestIes(protocolIEsC *C.ProtocolIE_Container_1710P11_t) (*e2appducontents.E2SetupRequestIes, error) {
	pIEs := new(e2appducontents.E2SetupRequestIes)

	ieCount := int(protocolIEsC.list.count)
	//fmt.Printf("1544P11 Type %T Count %v Size %v\n", *protocolIEsC.list.array, protocolIEsC.list.count, protocolIEsC.list.size)
	for i := 0; i < ieCount; i++ {
		offset := unsafe.Sizeof(unsafe.Pointer(*protocolIEsC.list.array)) * uintptr(i) // Forget the rest - this works - 7Nov20
		e2srIeC := *(**C.E2setupRequestIEs_t)(unsafe.Pointer(uintptr(unsafe.Pointer(protocolIEsC.list.array)) + offset))

		ie, err := decodeE2setupRequestIE(e2srIeC)
		if err != nil {
			return nil, err
		}
		if ie.E2ApProtocolIes3 != nil {
			pIEs.E2ApProtocolIes3 = ie.E2ApProtocolIes3
			//} else {
			//	return nil, fmt.Errorf("decodeE2SetupRequestIes(): obtained payload doesn't contain GlobalE2NodeID")
		}
		if ie.E2ApProtocolIes10 != nil {
			pIEs.E2ApProtocolIes10 = ie.E2ApProtocolIes10
		}
		if ie.E2ApProtocolIes33 != nil {
			pIEs.E2ApProtocolIes33 = ie.E2ApProtocolIes33
		}
	}

	return pIEs, nil
}

func newE2SetupResponseIes(e2srIEs *e2appducontents.E2SetupResponseIes) (*C.ProtocolIE_Container_1710P12_t, error) {
	pIeC1710P12 := new(C.ProtocolIE_Container_1710P12_t)

	if e2srIEs.GetE2ApProtocolIes4() != nil {
		ie4C, err := newE2setupResponseIe4GlobalRicID(e2srIEs.GetE2ApProtocolIes4())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1710P12), unsafe.Pointer(ie4C)); err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("newE2SetupResponseIes(): GlobalRicID should be mandatory present in the message")
	}

	if e2srIEs.GetE2ApProtocolIes9() != nil {
		ie9C, err := newE2setupResponseIe9RanFunctionsAccepted(e2srIEs.GetE2ApProtocolIes9())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1710P12), unsafe.Pointer(ie9C)); err != nil {
			return nil, err
		}
	}

	if e2srIEs.GetE2ApProtocolIes13() != nil {
		ie13C, err := newE2setupResponseIe13RanFunctionsRejected(e2srIEs.GetE2ApProtocolIes13())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1710P12), unsafe.Pointer(ie13C)); err != nil {
			return nil, err
		}
	}

	if e2srIEs.GetE2ApProtocolIes35() != nil {
		ie35C, err := newE2setupResponseIe35E2nodeComponentConfigUpdateAckList(e2srIEs.GetE2ApProtocolIes35())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1710P12), unsafe.Pointer(ie35C)); err != nil {
			return nil, err
		}
	}

	return pIeC1710P12, nil
}

func decodeE2SetupResponseIes(protocolIEsC *C.ProtocolIE_Container_1710P12_t) (*e2appducontents.E2SetupResponseIes, error) {
	pIEs := new(e2appducontents.E2SetupResponseIes)

	ieCount := int(protocolIEsC.list.count)
	//fmt.Printf("1544P11 Type %T Count %v Size %v\n", *protocolIEsC.list.array, protocolIEsC.list.count, protocolIEsC.list.size)
	for i := 0; i < ieCount; i++ {
		offset := unsafe.Sizeof(unsafe.Pointer(*protocolIEsC.list.array)) * uintptr(i)
		e2srIeC := *(**C.E2setupResponseIEs_t)(unsafe.Pointer(uintptr(unsafe.Pointer(protocolIEsC.list.array)) + offset))

		ie, err := decodeE2setupResponseIE(e2srIeC)
		if err != nil {
			return nil, err
		}
		if ie.E2ApProtocolIes4 != nil {
			pIEs.E2ApProtocolIes4 = ie.E2ApProtocolIes4
			//} else {
			//	return nil, fmt.Errorf("decodeE2SetupResponseIes(): obtained payload doesn't contain GlobalRicID")
		}
		if ie.E2ApProtocolIes9 != nil {
			pIEs.E2ApProtocolIes9 = ie.E2ApProtocolIes9
		}
		if ie.E2ApProtocolIes13 != nil {
			pIEs.E2ApProtocolIes13 = ie.E2ApProtocolIes13
		}
		if ie.E2ApProtocolIes35 != nil {
			pIEs.E2ApProtocolIes35 = ie.E2ApProtocolIes35
		}
	}

	return pIEs, nil
}

func newRicSubscriptionResponseIe(rsrIEs *e2appducontents.RicsubscriptionResponseIes) (*C.ProtocolIE_Container_1710P1_t, error) {
	pIeC1710P1 := new(C.ProtocolIE_Container_1710P1_t)

	if rsrIEs.GetE2ApProtocolIes5() != nil {
		ie5C, err := newRicSubscriptionResponseIe5RanFunctionID(rsrIEs.GetE2ApProtocolIes5())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1710P1), unsafe.Pointer(ie5C)); err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("newRicSubscriptionResponseIe(): RanFunctionID should be mandatory present in the message")
	}
	if rsrIEs.GetE2ApProtocolIes17() != nil {
		ie17C, err := newRicSubscriptionResponseIe17RactionAdmittedList(rsrIEs.GetE2ApProtocolIes17())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1710P1), unsafe.Pointer(ie17C)); err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("newRicSubscriptionResponseIe(): RicActionAdmittedList should be mandatory present in the message")
	}
	//TODO: Comment back in when RICactionRejected is handled
	if rsrIEs.GetE2ApProtocolIes18() != nil {
		ie18C, err := newRicSubscriptionResponseIe18RicActionNotAdmittedList(rsrIEs.GetE2ApProtocolIes18())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1710P1), unsafe.Pointer(ie18C)); err != nil {
			return nil, err
		}
	}
	if rsrIEs.GetE2ApProtocolIes29() != nil {
		ie29C, err := newRicSubscriptionResponseIe29RicRequestID(rsrIEs.GetE2ApProtocolIes29())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1710P1), unsafe.Pointer(ie29C)); err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("newRicSubscriptionResponseIe(): RicRequestID should be mandatory present in the message")
	}
	return pIeC1710P1, nil
}

func decodeRicSubscriptionResponseIes(protocolIEsC *C.ProtocolIE_Container_1710P1_t) (*e2appducontents.RicsubscriptionResponseIes, error) {
	pIEs := new(e2appducontents.RicsubscriptionResponseIes)

	ieCount := int(protocolIEsC.list.count)
	//fmt.Printf("1544P1 Type %T Count %v Size %v\n", *protocolIEsC.list.array, protocolIEsC.list.count, protocolIEsC.list.size)
	for i := 0; i < ieCount; i++ {
		offset := unsafe.Sizeof(unsafe.Pointer(*protocolIEsC.list.array)) * uintptr(i)
		rsrIeC := *(**C.RICsubscriptionResponse_IEs_t)(unsafe.Pointer(uintptr(unsafe.Pointer(protocolIEsC.list.array)) + offset))

		ie, err := decodeRicSubscriptionResponseIE(rsrIeC)
		if err != nil {
			return nil, err
		}
		if ie.E2ApProtocolIes5 != nil {
			pIEs.E2ApProtocolIes5 = ie.E2ApProtocolIes5
			//} else {
			//	return nil, fmt.Errorf("decodeRicSubscriptionResponseIes(): obtained payload doesn't contain RanFunctionID")
		}
		if ie.E2ApProtocolIes17 != nil {
			pIEs.E2ApProtocolIes17 = ie.E2ApProtocolIes17
			//} else {
			//	return nil, fmt.Errorf("decodeRicSubscriptionResponseIes(): obtained payload doesn't contain RicActionAdmittedList")
		}
		if ie.E2ApProtocolIes18 != nil {
			pIEs.E2ApProtocolIes18 = ie.E2ApProtocolIes18
		}
		if ie.E2ApProtocolIes29 != nil {
			pIEs.E2ApProtocolIes29 = ie.E2ApProtocolIes29
			//} else {
			//	return nil, fmt.Errorf("decodeRicSubscriptionResponseIes(): obtained payload doesn't contain RicRequestID")
		}
	}

	return pIEs, nil
}

func newRicSubscriptionRequestIes(rsrIEs *e2appducontents.RicsubscriptionRequestIes) (*C.ProtocolIE_Container_1710P0_t, error) {
	pIeC1710P0 := new(C.ProtocolIE_Container_1710P0_t)

	if rsrIEs.GetE2ApProtocolIes5() != nil {
		ie5C, err := newRicSubscriptionRequestIe5RanFunctionID(rsrIEs.E2ApProtocolIes5)
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1710P0), unsafe.Pointer(ie5C)); err != nil {
			return nil, err
		}
	}

	if rsrIEs.GetE2ApProtocolIes29() != nil {
		ie29C, err := newRicSubscriptionRequestIe29RicRequestID(rsrIEs.E2ApProtocolIes29)
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1710P0), unsafe.Pointer(ie29C)); err != nil {
			return nil, err
		}
	}

	if rsrIEs.GetE2ApProtocolIes30() != nil {
		ie30C, err := newRicSubscriptionRequestIe30RicSubscriptionDetails(rsrIEs.E2ApProtocolIes30)
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1710P0), unsafe.Pointer(ie30C)); err != nil {
			return nil, err
		}
	}

	return pIeC1710P0, nil
}

func decodeRicSubscriptionRequestIes(protocolIEsC *C.ProtocolIE_Container_1710P0_t) (*e2appducontents.RicsubscriptionRequestIes, error) {
	pIEs := new(e2appducontents.RicsubscriptionRequestIes)

	ieCount := int(protocolIEsC.list.count)
	//	fmt.Printf("1544P0 Type %T Count %v Size %v\n", *protocolIEsC.list.array, protocolIEsC.list.count, protocolIEsC.list.size)
	for i := 0; i < ieCount; i++ {
		offset := unsafe.Sizeof(unsafe.Pointer(*protocolIEsC.list.array)) * uintptr(i)
		rsrIeC := *(**C.RICsubscriptionRequest_IEs_t)(unsafe.Pointer(uintptr(unsafe.Pointer(protocolIEsC.list.array)) + offset))

		ie, err := decodeRicSubscriptionRequestIE(rsrIeC)
		if err != nil {
			return nil, err
		}
		if ie.E2ApProtocolIes5 != nil {
			pIEs.E2ApProtocolIes5 = ie.E2ApProtocolIes5
		}
		if ie.E2ApProtocolIes29 != nil {
			pIEs.E2ApProtocolIes29 = ie.E2ApProtocolIes29
		}
		if ie.E2ApProtocolIes30 != nil {
			pIEs.E2ApProtocolIes30 = ie.E2ApProtocolIes30
		}
	}

	return pIEs, nil
}

func newRicIndicationIEs(riIes *e2appducontents.RicindicationIes) (*C.ProtocolIE_Container_1710P6_t, error) {
	pIeC1710P6 := new(C.ProtocolIE_Container_1710P6_t)

	if riIes.GetE2ApProtocolIes5() != nil {
		ie5c, err := newRicIndicationIe5RanFunctionID(riIes.GetE2ApProtocolIes5())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1710P6), unsafe.Pointer(ie5c)); err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("newRicIndicationIEs(): RanFunctionID should be mandatory present in the message")
	}

	if riIes.GetE2ApProtocolIes15() != nil {
		ie15c, err := newRicIndicationIe15RicActionID(riIes.GetE2ApProtocolIes15())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1710P6), unsafe.Pointer(ie15c)); err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("newRicIndicationIEs(): RicActionID should be mandatory present in the message")
	}

	if riIes.GetE2ApProtocolIes20() != nil {
		ie20c, err := newRicIndicationIe20RiccallProcessID(riIes.GetE2ApProtocolIes20())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1710P6), unsafe.Pointer(ie20c)); err != nil {
			return nil, err
		}
	}

	if riIes.GetE2ApProtocolIes25() != nil {
		ie25c, err := newRicIndicationIe25RicIndicationHeader(riIes.GetE2ApProtocolIes25())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1710P6), unsafe.Pointer(ie25c)); err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("newRicIndicationIEs(): RicIndicationHeader should be mandatory present in the message")
	}

	if riIes.GetE2ApProtocolIes26() != nil {
		ie26c, err := newRicIndicationIe26RicIndicationMessage(riIes.GetE2ApProtocolIes26())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1710P6), unsafe.Pointer(ie26c)); err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("newRicIndicationIEs(): RicIndicationMessage should be mandatory present in the message")
	}

	if riIes.GetE2ApProtocolIes27() != nil {
		ie27c, err := newRicIndicationIe27RicIndicationSn(riIes.GetE2ApProtocolIes27())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1710P6), unsafe.Pointer(ie27c)); err != nil {
			return nil, err
		}
	}

	if riIes.GetE2ApProtocolIes28() != nil {
		ie28c, err := newRicIndicationIe28RicIndicationType(riIes.GetE2ApProtocolIes28())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1710P6), unsafe.Pointer(ie28c)); err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("newRicIndicationIEs(): RicIndicationType should be mandatory present in the message")
	}

	if riIes.GetE2ApProtocolIes29() != nil {
		ie29c, err := newRicIndicationIe29RicRequestID(riIes.GetE2ApProtocolIes29())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1710P6), unsafe.Pointer(ie29c)); err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("newRicIndicationIEs(): RicRequestID should be mandatory present in the message")
	}

	return pIeC1710P6, nil
}

func decodeRicIndicationIes(protocolIEsC *C.ProtocolIE_Container_1710P6_t) (*e2appducontents.RicindicationIes, error) {
	pIEs := new(e2appducontents.RicindicationIes)

	ieCount := int(protocolIEsC.list.count)
	//fmt.Printf("1544P6 Type %T Count %v Size %v\n", *protocolIEsC.list.array, protocolIEsC.list.count, protocolIEsC.list.size)
	for i := 0; i < ieCount; i++ {
		offset := unsafe.Sizeof(unsafe.Pointer(*protocolIEsC.list.array)) * uintptr(i) // Forget the rest - this works - 7Nov20
		riIeC := *(**C.RICindication_IEs_t)(unsafe.Pointer(uintptr(unsafe.Pointer(protocolIEsC.list.array)) + offset))

		ie, err := decodeRicIndicationIE(riIeC)
		if err != nil {
			return nil, err
		}
		if ie.E2ApProtocolIes5 != nil {
			pIEs.E2ApProtocolIes5 = ie.E2ApProtocolIes5
			//} else {
			//	return nil, fmt.Errorf("decodeRicIndicationIes(): obtained payload doesn't contain RanFunctionID")
		}
		if ie.E2ApProtocolIes15 != nil {
			pIEs.E2ApProtocolIes15 = ie.E2ApProtocolIes15
			//} else {
			//	return nil, fmt.Errorf("decodeRicIndicationIes(): obtained payload doesn't contain RicActionID")
		}
		if ie.E2ApProtocolIes20 != nil {
			pIEs.E2ApProtocolIes20 = ie.E2ApProtocolIes20
		}
		if ie.E2ApProtocolIes25 != nil {
			pIEs.E2ApProtocolIes25 = ie.E2ApProtocolIes25
			//} else {
			//	return nil, fmt.Errorf("decodeRicIndicationIes(): obtained payload doesn't contain RicIndicationHeader")
		}
		if ie.E2ApProtocolIes26 != nil {
			pIEs.E2ApProtocolIes26 = ie.E2ApProtocolIes26
			//} else {
			//	return nil, fmt.Errorf("decodeRicIndicationIes(): obtained payload doesn't contain RicIndicationMessage")
		}
		if ie.E2ApProtocolIes27 != nil {
			pIEs.E2ApProtocolIes27 = ie.E2ApProtocolIes27
		}
		if ie.E2ApProtocolIes28 != nil {
			pIEs.E2ApProtocolIes28 = ie.E2ApProtocolIes28
			//} else {
			//	return nil, fmt.Errorf("decodeRicIndicationIes(): obtained payload doesn't contain RicIndicationType")
		}
		if ie.E2ApProtocolIes29 != nil {
			pIEs.E2ApProtocolIes29 = ie.E2ApProtocolIes29
			//} else {
			//	return nil, fmt.Errorf("decodeRicIndicationIes(): obtained payload doesn't contain RicrequestID")
		}
	}

	return pIEs, nil
}

func newRicControlRequestIEs(rcRIes *e2appducontents.RiccontrolRequestIes) (*C.ProtocolIE_Container_1710P7_t, error) {
	pIeC1710P7 := new(C.ProtocolIE_Container_1710P7_t)

	if rcRIes.GetE2ApProtocolIes5() != nil {
		ie5c, err := newRicControlRequestIe5RanFunctionID(rcRIes.GetE2ApProtocolIes5())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1710P7), unsafe.Pointer(ie5c)); err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("newRicControlRequestIEs(): RanFunctionID should be mandatory present in the message")
	}

	if rcRIes.GetE2ApProtocolIes20() != nil {
		ie20c, err := newRicControlRequestIe20RiccallProcessID(rcRIes.GetE2ApProtocolIes20())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1710P7), unsafe.Pointer(ie20c)); err != nil {
			return nil, err
		}
	}

	if rcRIes.GetE2ApProtocolIes22() != nil {
		ie22c, err := newRicControlRequestIe22RiccontrolHeader(rcRIes.GetE2ApProtocolIes22())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1710P7), unsafe.Pointer(ie22c)); err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("newRicControlRequestIEs(): RicControlHeader should be mandatory present in the message")
	}

	if rcRIes.GetE2ApProtocolIes23() != nil {
		ie23c, err := newRicControlRequestIe23RiccontrolMessage(rcRIes.GetE2ApProtocolIes23())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1710P7), unsafe.Pointer(ie23c)); err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("newRicControlRequestIEs(): RicControlMessage should be mandatory present in the message")
	}

	if rcRIes.GetE2ApProtocolIes21() != nil {
		ie21c, err := newRicControlRequestIe21RiccontrolAckRequest(rcRIes.GetE2ApProtocolIes21())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1710P7), unsafe.Pointer(ie21c)); err != nil {
			return nil, err
		}
	}

	if rcRIes.GetE2ApProtocolIes29() != nil {
		ie29c, err := newRicControlRequestIe29RicRequestID(rcRIes.GetE2ApProtocolIes29())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1710P7), unsafe.Pointer(ie29c)); err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("newRicControlRequestIEs(): RicRequestID should be mandatory present in the message")
	}

	return pIeC1710P7, nil
}

func decodeRicControlRequestIes(protocolIEsC *C.ProtocolIE_Container_1710P7_t) (*e2appducontents.RiccontrolRequestIes, error) {
	pIEs := new(e2appducontents.RiccontrolRequestIes)

	ieCount := int(protocolIEsC.list.count)
	//fmt.Printf("1544P6 Type %T Count %v Size %v\n", *protocolIEsC.list.array, protocolIEsC.list.count, protocolIEsC.list.size)
	for i := 0; i < ieCount; i++ {
		offset := unsafe.Sizeof(unsafe.Pointer(*protocolIEsC.list.array)) * uintptr(i) // Forget the rest - this works - 7Nov20
		riIeC := *(**C.RICcontrolRequest_IEs_t)(unsafe.Pointer(uintptr(unsafe.Pointer(protocolIEsC.list.array)) + offset))

		ie, err := decodeRicControlRequestIE(riIeC)
		if err != nil {
			return nil, err
		}
		if ie.E2ApProtocolIes5 != nil {
			pIEs.E2ApProtocolIes5 = ie.E2ApProtocolIes5
			//} else {
			//	return nil, fmt.Errorf("decodeRicControlRequestIes(): obtained payload doesn't contain RanFunctionID")
		}
		if ie.E2ApProtocolIes20 != nil {
			pIEs.E2ApProtocolIes20 = ie.E2ApProtocolIes20
		}
		if ie.E2ApProtocolIes22 != nil {
			pIEs.E2ApProtocolIes22 = ie.E2ApProtocolIes22
			//} else {
			//	return nil, fmt.Errorf("decodeRicControlRequestIes(): obtained payload doesn't contain RicControlHeader")
		}
		if ie.E2ApProtocolIes23 != nil {
			pIEs.E2ApProtocolIes23 = ie.E2ApProtocolIes23
			//} else {
			//	return nil, fmt.Errorf("decodeRicControlRequestIes(): obtained payload doesn't contain RicControlMessage")
		}
		if ie.E2ApProtocolIes21 != nil {
			pIEs.E2ApProtocolIes21 = ie.E2ApProtocolIes21
		}
		if ie.E2ApProtocolIes29 != nil {
			pIEs.E2ApProtocolIes29 = ie.E2ApProtocolIes29
			//} else {
			//	return nil, fmt.Errorf("decodeRicControlRequestIes(): obtained payload doesn't contain RicRequestID")
		}
	}

	return pIEs, nil
}

func newRicControlFailureIEs(rcFIes *e2appducontents.RiccontrolFailureIes) (*C.ProtocolIE_Container_1710P9_t, error) {
	pIeC1710P9 := new(C.ProtocolIE_Container_1710P9_t)

	if rcFIes.GetE2ApProtocolIes5() != nil {
		ie5c, err := newRicControlFailureIe5RanFunctionID(rcFIes.GetE2ApProtocolIes5())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1710P9), unsafe.Pointer(ie5c)); err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("newRicControlFailureIEs() RanFunctionID should be mandatory present in the message")
	}

	if rcFIes.GetE2ApProtocolIes20() != nil {
		ie20c, err := newRicControlFailureIe20RiccallProcessID(rcFIes.GetE2ApProtocolIes20())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1710P9), unsafe.Pointer(ie20c)); err != nil {
			return nil, err
		}
	}

	if rcFIes.GetE2ApProtocolIes32() != nil {
		ie22c, err := newRicControlFailureIe32RiccontrolOutcome(rcFIes.GetE2ApProtocolIes32())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1710P9), unsafe.Pointer(ie22c)); err != nil {
			return nil, err
		}
	}

	if rcFIes.GetE2ApProtocolIes1() != nil {
		ie21c, err := newRicControlFailureIe1Cause(rcFIes.GetE2ApProtocolIes1())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1710P9), unsafe.Pointer(ie21c)); err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("newRicControlFailureIEs() Cause should be mandatory present in the message")
	}

	if rcFIes.GetE2ApProtocolIes29() != nil {
		ie29c, err := newRicControlFailureIe29RicRequestID(rcFIes.GetE2ApProtocolIes29())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1710P9), unsafe.Pointer(ie29c)); err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("newRicControlFailureIEs() RicRequestID should be mandatory present in the message")
	}

	return pIeC1710P9, nil
}

func decodeRicControlFailureIes(protocolIEsC *C.ProtocolIE_Container_1710P9_t) (*e2appducontents.RiccontrolFailureIes, error) {
	pIEs := new(e2appducontents.RiccontrolFailureIes)

	ieCount := int(protocolIEsC.list.count)
	//fmt.Printf("1544P9 Type %T Count %v Size %v\n", *protocolIEsC.list.array, protocolIEsC.list.count, protocolIEsC.list.size)
	for i := 0; i < ieCount; i++ {
		offset := unsafe.Sizeof(unsafe.Pointer(*protocolIEsC.list.array)) * uintptr(i) // Forget the rest - this works - 7Nov20
		rcfIeC := *(**C.RICcontrolFailure_IEs_t)(unsafe.Pointer(uintptr(unsafe.Pointer(protocolIEsC.list.array)) + offset))

		ie, err := decodeRicControlFailureIE(rcfIeC)
		if err != nil {
			return nil, err
		}
		if ie.E2ApProtocolIes5 != nil {
			pIEs.E2ApProtocolIes5 = ie.E2ApProtocolIes5
		}
		if ie.E2ApProtocolIes20 != nil {
			pIEs.E2ApProtocolIes20 = ie.E2ApProtocolIes20
		}
		if ie.E2ApProtocolIes32 != nil {
			pIEs.E2ApProtocolIes32 = ie.E2ApProtocolIes32
		}
		if ie.E2ApProtocolIes1 != nil {
			pIEs.E2ApProtocolIes1 = ie.E2ApProtocolIes1
		}
		if ie.E2ApProtocolIes29 != nil {
			pIEs.E2ApProtocolIes29 = ie.E2ApProtocolIes29
		}
	}

	return pIEs, nil
}

func newRicControlAcknowledgeIEs(rcaIes *e2appducontents.RiccontrolAcknowledgeIes) (*C.ProtocolIE_Container_1710P8_t, error) {
	pIeC1710P8 := new(C.ProtocolIE_Container_1710P8_t)

	if rcaIes.GetE2ApProtocolIes5() != nil {
		ie5c, err := newRicControlAcknowledgeIe5RanFunctionID(rcaIes.GetE2ApProtocolIes5())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1710P8), unsafe.Pointer(ie5c)); err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("newRicControlAcknowledgeIEs(): RanFunctionID should be mandatory present in the message")
	}

	if rcaIes.GetE2ApProtocolIes20() != nil {
		ie20c, err := newRicControlAcknowledgeIe20RiccallProcessID(rcaIes.GetE2ApProtocolIes20())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1710P8), unsafe.Pointer(ie20c)); err != nil {
			return nil, err
		}
	}

	if rcaIes.GetE2ApProtocolIes24() != nil {
		ie22c, err := newRicControlAcknowledgeIe24RiccontrolStatus(rcaIes.GetE2ApProtocolIes24())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1710P8), unsafe.Pointer(ie22c)); err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("newRicControlAcknowledgeIEs(): RicControlStatus should be mandatory present in the message")
	}

	if rcaIes.GetE2ApProtocolIes29() != nil {
		ie29c, err := newRicControlAcknowledgeIe29RicRequestID(rcaIes.GetE2ApProtocolIes29())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1710P8), unsafe.Pointer(ie29c)); err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("newRicControlAcknowledgeIEs(): RicRequestID should be mandatory present in the message")
	}

	if rcaIes.GetE2ApProtocolIes32() != nil {
		ie32c, err := newRicControlAcknowledgeIe32RiccontrolOutcome(rcaIes.GetE2ApProtocolIes32())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1710P8), unsafe.Pointer(ie32c)); err != nil {
			return nil, err
		}
	}

	return pIeC1710P8, nil
}

func decodeRicControlAcknowledgeIes(protocolIEsC *C.ProtocolIE_Container_1710P8_t) (*e2appducontents.RiccontrolAcknowledgeIes, error) {
	pIEs := new(e2appducontents.RiccontrolAcknowledgeIes)

	ieCount := int(protocolIEsC.list.count)
	//fmt.Printf("1544P6 Type %T Count %v Size %v\n", *protocolIEsC.list.array, protocolIEsC.list.count, protocolIEsC.list.size)
	for i := 0; i < ieCount; i++ {
		offset := unsafe.Sizeof(unsafe.Pointer(*protocolIEsC.list.array)) * uintptr(i) // Forget the rest - this works - 7Nov20
		riIeC := *(**C.RICcontrolAcknowledge_IEs_t)(unsafe.Pointer(uintptr(unsafe.Pointer(protocolIEsC.list.array)) + offset))

		ie, err := decodeRicControlAcknowledgeIE(riIeC)
		if err != nil {
			return nil, err
		}
		if ie.E2ApProtocolIes5 != nil {
			pIEs.E2ApProtocolIes5 = ie.E2ApProtocolIes5
			//} else {
			//	return nil, fmt.Errorf("decodeRicControlAcknowledgeIes(): obtained payload doesn't contain RanFunctionID")
		}
		if ie.E2ApProtocolIes20 != nil {
			pIEs.E2ApProtocolIes20 = ie.E2ApProtocolIes20
		}
		if ie.E2ApProtocolIes24 != nil {
			pIEs.E2ApProtocolIes24 = ie.E2ApProtocolIes24
			//} else {
			//	return nil, fmt.Errorf("decodeRicControlAcknowledgeIes(): obtained payload doesn't contain RicControlStatus")
		}
		if ie.E2ApProtocolIes29 != nil {
			pIEs.E2ApProtocolIes29 = ie.E2ApProtocolIes29
			//} else {
			//	return nil, fmt.Errorf("decodeRicControlAcknowledgeIes(): obtained payload doesn't contain RicRequestID")
		}
		if ie.E2ApProtocolIes32 != nil {
			pIEs.E2ApProtocolIes32 = ie.E2ApProtocolIes32
		}
	}

	return pIEs, nil
}

func newRicSubscriptionDeleteRequestIes(rsdrIEs *e2appducontents.RicsubscriptionDeleteRequestIes) (*C.ProtocolIE_Container_1710P3_t, error) {
	pIeC1710P3 := new(C.ProtocolIE_Container_1710P3_t)

	if rsdrIEs.GetE2ApProtocolIes5() != nil {
		ie5C, err := newRicSubscriptionDeleteRequestIe5RanFunctionID(rsdrIEs.E2ApProtocolIes5)
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1710P3), unsafe.Pointer(ie5C)); err != nil {
			return nil, err
		}
	}

	if rsdrIEs.GetE2ApProtocolIes29() != nil {
		ie29C, err := newRicSubscriptionDeleteRequestIe29RicRequestID(rsdrIEs.E2ApProtocolIes29)
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1710P3), unsafe.Pointer(ie29C)); err != nil {
			return nil, err
		}
	}

	return pIeC1710P3, nil
}

func decodeRicSubscriptionDeleteRequestIes(protocolIEsC *C.ProtocolIE_Container_1710P3_t) (*e2appducontents.RicsubscriptionDeleteRequestIes, error) {
	pIEs := new(e2appducontents.RicsubscriptionDeleteRequestIes)

	ieCount := int(protocolIEsC.list.count)
	//	fmt.Printf("1544P0 Type %T Count %v Size %v\n", *protocolIEsC.list.array, protocolIEsC.list.count, protocolIEsC.list.size)
	for i := 0; i < ieCount; i++ {
		offset := unsafe.Sizeof(unsafe.Pointer(*protocolIEsC.list.array)) * uintptr(i)
		rsrIeC := *(**C.RICsubscriptionDeleteRequest_IEs_t)(unsafe.Pointer(uintptr(unsafe.Pointer(protocolIEsC.list.array)) + offset))

		ie, err := decodeRicSubscriptionDeleteRequestIE(rsrIeC)
		if err != nil {
			return nil, err
		}
		if ie.E2ApProtocolIes5 != nil {
			pIEs.E2ApProtocolIes5 = ie.E2ApProtocolIes5
		}
		if ie.E2ApProtocolIes29 != nil {
			pIEs.E2ApProtocolIes29 = ie.E2ApProtocolIes29
		}
	}

	return pIEs, nil
}

func newRicSubscriptionDeleteResponseIe(rsrIEs *e2appducontents.RicsubscriptionDeleteResponseIes) (*C.ProtocolIE_Container_1710P4_t, error) {
	pIeC1710P4 := new(C.ProtocolIE_Container_1710P4_t)

	if rsrIEs.GetE2ApProtocolIes5() != nil {
		ie5C, err := newRicSubscriptionDeleteResponseIe5RanFunctionID(rsrIEs.GetE2ApProtocolIes5())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1710P4), unsafe.Pointer(ie5C)); err != nil {
			return nil, err
		}
	}

	if rsrIEs.GetE2ApProtocolIes29() != nil {
		ie29C, err := newRicSubscriptionDeleteResponseIe29RicRequestID(rsrIEs.GetE2ApProtocolIes29())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1710P4), unsafe.Pointer(ie29C)); err != nil {
			return nil, err
		}
	}
	return pIeC1710P4, nil
}

func decodeRicSubscriptionDeleteResponseIes(protocolIEsC *C.ProtocolIE_Container_1710P4_t) (*e2appducontents.RicsubscriptionDeleteResponseIes, error) {
	pIEs := new(e2appducontents.RicsubscriptionDeleteResponseIes)

	ieCount := int(protocolIEsC.list.count)
	//fmt.Printf("1544P1 Type %T Count %v Size %v\n", *protocolIEsC.list.array, protocolIEsC.list.count, protocolIEsC.list.size)
	for i := 0; i < ieCount; i++ {
		offset := unsafe.Sizeof(unsafe.Pointer(*protocolIEsC.list.array)) * uintptr(i)
		rsrIeC := *(**C.RICsubscriptionDeleteResponse_IEs_t)(unsafe.Pointer(uintptr(unsafe.Pointer(protocolIEsC.list.array)) + offset))

		ie, err := decodeRicSubscriptionDeleteResponseIE(rsrIeC)
		if err != nil {
			return nil, err
		}
		if ie.E2ApProtocolIes5 != nil {
			pIEs.E2ApProtocolIes5 = ie.E2ApProtocolIes5
		}
		if ie.E2ApProtocolIes29 != nil {
			pIEs.E2ApProtocolIes29 = ie.E2ApProtocolIes29
		}
	}

	return pIEs, nil
}

func newRicSubscriptionFailureIe(rsdIEs *e2appducontents.RicsubscriptionFailureIes) (*C.ProtocolIE_Container_1710P2_t, error) {
	pIeC1710P2 := new(C.ProtocolIE_Container_1710P2_t)

	if rsdIEs.GetE2ApProtocolIes2() != nil {
		ie2C, err := newRicSubscriptionFailureIe2CriticalityDiagnostics(rsdIEs.GetE2ApProtocolIes2())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1710P2), unsafe.Pointer(ie2C)); err != nil {
			return nil, err
		}
	}

	if rsdIEs.GetE2ApProtocolIes5() != nil {
		ie5C, err := newRicSubscriptionFailureIe5RanFunctionID(rsdIEs.GetE2ApProtocolIes5())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1710P2), unsafe.Pointer(ie5C)); err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("newRicSubscriptionFailureIe(): RanFunctionID should be mandatory present in the message")
	}

	if rsdIEs.GetE2ApProtocolIes18() != nil {
		ie2C, err := newRicSubscriptionFailureIe18RicActionNotAdmittedList(rsdIEs.GetE2ApProtocolIes18())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1710P2), unsafe.Pointer(ie2C)); err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("newRicSubscriptionFailureIe(): RicActionNotAdmittedList should be mandatory present in the message")
	}

	if rsdIEs.GetE2ApProtocolIes29() != nil {
		ie29C, err := newRicSubscriptionFailureIe29RicRequestID(rsdIEs.GetE2ApProtocolIes29())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1710P2), unsafe.Pointer(ie29C)); err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("newRicSubscriptionFailureIe(): RicRequestID should be mandatory present in the message")
	}

	return pIeC1710P2, nil
}

func decodeRicSubscriptionFailureIes(protocolIEsC *C.ProtocolIE_Container_1710P2_t) (*e2appducontents.RicsubscriptionFailureIes, error) {
	pIEs := new(e2appducontents.RicsubscriptionFailureIes)

	ieCount := int(protocolIEsC.list.count)
	//fmt.Printf("1544P1 Type %T Count %v Size %v\n", *protocolIEsC.list.array, protocolIEsC.list.count, protocolIEsC.list.size)
	for i := 0; i < ieCount; i++ {
		offset := unsafe.Sizeof(unsafe.Pointer(*protocolIEsC.list.array)) * uintptr(i)
		rsfIeC := *(**C.RICsubscriptionFailure_IEs_t)(unsafe.Pointer(uintptr(unsafe.Pointer(protocolIEsC.list.array)) + offset))

		ie, err := decodeRicSubscriptionFailureIE(rsfIeC)
		if err != nil {
			return nil, err
		}
		if ie.E2ApProtocolIes2 != nil {
			pIEs.E2ApProtocolIes2 = ie.E2ApProtocolIes2
		}
		if ie.E2ApProtocolIes5 != nil {
			pIEs.E2ApProtocolIes5 = ie.E2ApProtocolIes5
			//} else {
			//	return nil, fmt.Errorf("decodeRicSubscriptionFailureIes(): obtained payload doesn't contain mandatory RanFunctionID")
		}
		if ie.E2ApProtocolIes18 != nil {
			pIEs.E2ApProtocolIes18 = ie.E2ApProtocolIes18
			//} else {
			//	return nil, fmt.Errorf("decodeRicSubscriptionFailureIes(): obtained payload doesn't contain mandatory RicActionNotAdmittedList")
		}
		if ie.E2ApProtocolIes29 != nil {
			pIEs.E2ApProtocolIes29 = ie.E2ApProtocolIes29
			//} else {
			//	return nil, fmt.Errorf("decodeRicSubscriptionFailureIes(): obtained payload doesn't contain mandatory RicRequestID")
		}
	}

	return pIEs, nil
}

func newRicSubscriptionDeleteFailureIe(rsdfIEs *e2appducontents.RicsubscriptionDeleteFailureIes) (*C.ProtocolIE_Container_1710P5_t, error) {
	pIeC1710P5 := new(C.ProtocolIE_Container_1710P5_t)

	if rsdfIEs.GetE2ApProtocolIes5() != nil {
		ie5C, err := newRicSubscriptionDeleteFailureIe5RanFunctionID(rsdfIEs.GetE2ApProtocolIes5())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1710P5), unsafe.Pointer(ie5C)); err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("newRicSubscriptionDeleteFailureIe(): RanFunctionID should be mandatory present in the message")
	}

	if rsdfIEs.GetE2ApProtocolIes29() != nil {
		ie29C, err := newRicSubscriptionDeleteFailureIe29RicRequestID(rsdfIEs.GetE2ApProtocolIes29())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1710P5), unsafe.Pointer(ie29C)); err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("newRicSubscriptionDeleteFailureIe(): RicRequestID should be mandatory present in the message")
	}

	if rsdfIEs.GetE2ApProtocolIes1() != nil {
		ie1C, err := newRicSubscriptionDeleteFailureIe1Cause(rsdfIEs.GetE2ApProtocolIes1())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1710P5), unsafe.Pointer(ie1C)); err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("newRicSubscriptionDeleteFailureIe(): Cause should be mandatory present in the message")
	}

	if rsdfIEs.GetE2ApProtocolIes2() != nil {
		ie2C, err := newRicSubscriptionDeleteFailureIe2CriticalityDiagnostics(rsdfIEs.GetE2ApProtocolIes2())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1710P5), unsafe.Pointer(ie2C)); err != nil {
			return nil, err
		}
	}
	return pIeC1710P5, nil
}

func decodeRicSubscriptionDeleteFailureIes(protocolIEsC *C.ProtocolIE_Container_1710P5_t) (*e2appducontents.RicsubscriptionDeleteFailureIes, error) {
	pIEs := new(e2appducontents.RicsubscriptionDeleteFailureIes)

	ieCount := int(protocolIEsC.list.count)
	//fmt.Printf("1544P1 Type %T Count %v Size %v\n", *protocolIEsC.list.array, protocolIEsC.list.count, protocolIEsC.list.size)
	for i := 0; i < ieCount; i++ {
		offset := unsafe.Sizeof(unsafe.Pointer(*protocolIEsC.list.array)) * uintptr(i)
		rsdfIeC := *(**C.RICsubscriptionDeleteFailure_IEs_t)(unsafe.Pointer(uintptr(unsafe.Pointer(protocolIEsC.list.array)) + offset))

		ie, err := decodeRicSubscriptionDeleteFailureIE(rsdfIeC)
		if err != nil {
			return nil, err
		}
		if ie.E2ApProtocolIes5 != nil {
			pIEs.E2ApProtocolIes5 = ie.E2ApProtocolIes5
			//} else {
			//	return nil, fmt.Errorf("decodeRicSubscriptionDeleteFailureIes(): obtained payload doesn't contain RanFunctionID")
		}
		if ie.E2ApProtocolIes29 != nil {
			pIEs.E2ApProtocolIes29 = ie.E2ApProtocolIes29
			//} else {
			//	return nil, fmt.Errorf("decodeRicSubscriptionDeleteFailureIes(): obtained payload doesn't contain RicRequestID")
		}
		if ie.E2ApProtocolIes1 != nil {
			pIEs.E2ApProtocolIes1 = ie.E2ApProtocolIes1
			//} else {
			//	return nil, fmt.Errorf("decodeRicSubscriptionDeleteFailureIes(): obtained payload doesn't contain Cause")
		}
		if ie.E2ApProtocolIes2 != nil {
			pIEs.E2ApProtocolIes2 = ie.E2ApProtocolIes2
		}
	}

	return pIEs, nil
}

func newErrorIndicationIe(eiIEs *e2appducontents.ErrorIndicationIes) (*C.ProtocolIE_Container_1710P10_t, error) {
	pIeC1710P10 := new(C.ProtocolIE_Container_1710P10_t)

	if eiIEs.GetE2ApProtocolIes2() != nil {
		ie2C, err := newErrorIndicationIe2CriticalityDiagnostics(eiIEs.GetE2ApProtocolIes2())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1710P10), unsafe.Pointer(ie2C)); err != nil {
			return nil, err
		}
	}

	if eiIEs.GetE2ApProtocolIes5() != nil {
		ie5C, err := newErrorIndicationIe5RanFunctionID(eiIEs.GetE2ApProtocolIes5())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1710P10), unsafe.Pointer(ie5C)); err != nil {
			return nil, err
		}
	}

	if eiIEs.GetE2ApProtocolIes1() != nil {
		ie1C, err := newErrorIndicationIe1Cause(eiIEs.GetE2ApProtocolIes1())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1710P10), unsafe.Pointer(ie1C)); err != nil {
			return nil, err
		}
	}

	if eiIEs.GetE2ApProtocolIes29() != nil {
		ie29C, err := newErrorIndicationIe29RicRequestID(eiIEs.GetE2ApProtocolIes29())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1710P10), unsafe.Pointer(ie29C)); err != nil {
			return nil, err
		}
	}

	return pIeC1710P10, nil
}

func decodeErrorIndicationIes(protocolIEsC *C.ProtocolIE_Container_1710P10_t) (*e2appducontents.ErrorIndicationIes, error) {
	pIEs := new(e2appducontents.ErrorIndicationIes)

	ieCount := int(protocolIEsC.list.count)
	//fmt.Printf("1544P1 Type %T Count %v Size %v\n", *protocolIEsC.list.array, protocolIEsC.list.count, protocolIEsC.list.size)
	for i := 0; i < ieCount; i++ {
		offset := unsafe.Sizeof(unsafe.Pointer(*protocolIEsC.list.array)) * uintptr(i)
		eiIeC := *(**C.ErrorIndication_IEs_t)(unsafe.Pointer(uintptr(unsafe.Pointer(protocolIEsC.list.array)) + offset))

		ie, err := decodeErrorIndicationIE(eiIeC)
		if err != nil {
			return nil, err
		}
		if ie.E2ApProtocolIes2 != nil {
			pIEs.E2ApProtocolIes2 = ie.E2ApProtocolIes2
		}
		if ie.E2ApProtocolIes5 != nil {
			pIEs.E2ApProtocolIes5 = ie.E2ApProtocolIes5
		}
		if ie.E2ApProtocolIes1 != nil {
			pIEs.E2ApProtocolIes1 = ie.E2ApProtocolIes1
		}
		if ie.E2ApProtocolIes29 != nil {
			pIEs.E2ApProtocolIes29 = ie.E2ApProtocolIes29
		}
	}

	return pIEs, nil
}

func newRicServiceQueryIe(rsqIEs *e2appducontents.RicserviceQueryIes) (*C.ProtocolIE_Container_1710P25_t, error) {
	pIeC1710P25 := new(C.ProtocolIE_Container_1710P25_t)

	if rsqIEs.GetRicserviceQueryIes9() != nil {
		ie9C, err := newRicServiceQueryIe9RanFunctionsAccepted(rsqIEs.GetRicserviceQueryIes9())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1710P25), unsafe.Pointer(ie9C)); err != nil {
			return nil, err
		}
	}

	return pIeC1710P25, nil
}

func decodeRicServiceQueryIes(protocolIEsC *C.ProtocolIE_Container_1710P25_t) (*e2appducontents.RicserviceQueryIes, error) {
	pIEs := new(e2appducontents.RicserviceQueryIes)

	ieCount := int(protocolIEsC.list.count)
	//fmt.Printf("1544P1 Type %T Count %v Size %v\n", *protocolIEsC.list.array, protocolIEsC.list.count, protocolIEsC.list.size)
	for i := 0; i < ieCount; i++ {
		offset := unsafe.Sizeof(unsafe.Pointer(*protocolIEsC.list.array)) * uintptr(i)
		eiIeC := *(**C.RICserviceQuery_IEs_t)(unsafe.Pointer(uintptr(unsafe.Pointer(protocolIEsC.list.array)) + offset))

		ie, err := decodeRicServiceQueryIE(eiIeC)
		if err != nil {
			return nil, err
		}
		if ie.RicserviceQueryIes9 != nil {
			pIEs.RicserviceQueryIes9 = ie.RicserviceQueryIes9
		}
	}

	return pIEs, nil
}

func newResetRequestIe(rrIEs *e2appducontents.ResetRequestIes) (*C.ProtocolIE_Container_1710P20_t, error) {
	pIeC1710P20 := new(C.ProtocolIE_Container_1710P20_t)

	if rrIEs.GetResetRequestIes1() != nil {
		ie1C, err := newResetRequestIe1Cause(rrIEs.GetResetRequestIes1())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1710P20), unsafe.Pointer(ie1C)); err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("newResetRequestIe() Cause should be mandatory present in the message")
	}

	return pIeC1710P20, nil
}

func decodeResetRequestIes(protocolIEsC *C.ProtocolIE_Container_1710P20_t) (*e2appducontents.ResetRequestIes, error) {
	pIEs := new(e2appducontents.ResetRequestIes)

	ieCount := int(protocolIEsC.list.count)
	//fmt.Printf("1544P1 Type %T Count %v Size %v\n", *protocolIEsC.list.array, protocolIEsC.list.count, protocolIEsC.list.size)
	for i := 0; i < ieCount; i++ {
		offset := unsafe.Sizeof(unsafe.Pointer(*protocolIEsC.list.array)) * uintptr(i)
		eiIeC := *(**C.ResetRequestIEs_t)(unsafe.Pointer(uintptr(unsafe.Pointer(protocolIEsC.list.array)) + offset))

		ie, err := decodeResetRequestIE(eiIeC)
		if err != nil {
			return nil, err
		}
		if ie.ResetRequestIes1 != nil {
			pIEs.ResetRequestIes1 = ie.ResetRequestIes1
		}
	}

	return pIEs, nil
}

func newResetResponseIe(rrIEs *e2appducontents.ResetResponseIes) (*C.ProtocolIE_Container_1710P21_t, error) {
	pIeC1710P21 := new(C.ProtocolIE_Container_1710P21_t)

	if rrIEs.GetResetResponseIes2() != nil {
		ie1C, err := newResetResponseIe2CriticalityDiagnostics(rrIEs.GetResetResponseIes2())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1710P21), unsafe.Pointer(ie1C)); err != nil {
			return nil, err
		}
	}

	return pIeC1710P21, nil
}

func decodeResetResponseIes(protocolIEsC *C.ProtocolIE_Container_1710P21_t) (*e2appducontents.ResetResponseIes, error) {
	pIEs := new(e2appducontents.ResetResponseIes)

	ieCount := int(protocolIEsC.list.count)
	//fmt.Printf("1544P1 Type %T Count %v Size %v\n", *protocolIEsC.list.array, protocolIEsC.list.count, protocolIEsC.list.size)
	for i := 0; i < ieCount; i++ {
		offset := unsafe.Sizeof(unsafe.Pointer(*protocolIEsC.list.array)) * uintptr(i)
		eiIeC := *(**C.ResetResponseIEs_t)(unsafe.Pointer(uintptr(unsafe.Pointer(protocolIEsC.list.array)) + offset))

		ie, err := decodeResetResponseIE(eiIeC)
		if err != nil {
			return nil, err
		}
		if ie.ResetResponseIes2 != nil {
			pIEs.ResetResponseIes2 = ie.ResetResponseIes2
		}
	}

	return pIEs, nil
}

func newRicServiceUpdateIe(rsuIEs *e2appducontents.RicserviceUpdateIes) (*C.ProtocolIE_Container_1710P22_t, error) {
	pIeC1710P22 := new(C.ProtocolIE_Container_1710P22_t)

	if rsuIEs.GetE2ApProtocolIes10() != nil {
		ie1C, err := newRicServiceUpdateIe10RanFunctionAddedList(rsuIEs.GetE2ApProtocolIes10())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1710P22), unsafe.Pointer(ie1C)); err != nil {
			return nil, err
		}
	}

	if rsuIEs.GetE2ApProtocolIes11() != nil {
		ie1C, err := newRicServiceUpdateIe11RanFunctionDeletedList(rsuIEs.GetE2ApProtocolIes11())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1710P22), unsafe.Pointer(ie1C)); err != nil {
			return nil, err
		}
	}

	if rsuIEs.GetE2ApProtocolIes12() != nil {
		ie1C, err := newRicServiceUpdateIe12RanFunctionModifiedList(rsuIEs.GetE2ApProtocolIes12())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1710P22), unsafe.Pointer(ie1C)); err != nil {
			return nil, err
		}
	}

	return pIeC1710P22, nil
}

func decodeRicServiceUpdateIes(protocolIEsC *C.ProtocolIE_Container_1710P22_t) (*e2appducontents.RicserviceUpdateIes, error) {
	pIEs := new(e2appducontents.RicserviceUpdateIes)

	ieCount := int(protocolIEsC.list.count)
	//fmt.Printf("1544P1 Type %T Count %v Size %v\n", *protocolIEsC.list.array, protocolIEsC.list.count, protocolIEsC.list.size)
	for i := 0; i < ieCount; i++ {
		offset := unsafe.Sizeof(unsafe.Pointer(*protocolIEsC.list.array)) * uintptr(i)
		rsuIeC := *(**C.RICserviceUpdate_IEs_t)(unsafe.Pointer(uintptr(unsafe.Pointer(protocolIEsC.list.array)) + offset))

		ie, err := decodeRicServiceUpdateIE(rsuIeC)
		if err != nil {
			return nil, err
		}
		if ie.GetE2ApProtocolIes10() != nil {
			pIEs.E2ApProtocolIes10 = ie.E2ApProtocolIes10
		}

		if ie.GetE2ApProtocolIes11() != nil {
			pIEs.E2ApProtocolIes11 = ie.E2ApProtocolIes11
		}

		if ie.GetE2ApProtocolIes12() != nil {
			pIEs.E2ApProtocolIes12 = ie.E2ApProtocolIes12
		}
	}

	return pIEs, nil
}

func newRicServiceUpdateAcknowledgeIe(rsuIEs *e2appducontents.RicserviceUpdateAcknowledgeIes) (*C.ProtocolIE_Container_1710P23_t, error) {
	pIeC1710P23 := new(C.ProtocolIE_Container_1710P23_t)

	if rsuIEs.GetE2ApProtocolIes9() != nil {
		ie1C, err := newRicServiceUpdateAcknowledgeIe9RanFunctionsAccepted(rsuIEs.GetE2ApProtocolIes9())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1710P23), unsafe.Pointer(ie1C)); err != nil {
			return nil, err
		}
	}

	if rsuIEs.GetE2ApProtocolIes13() != nil {
		ie1C, err := newRicServiceUpdateAcknowledgeIe13RanFunctionsRejected(rsuIEs.GetE2ApProtocolIes13())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1710P23), unsafe.Pointer(ie1C)); err != nil {
			return nil, err
		}
	}

	return pIeC1710P23, nil
}

func decodeRicServiceUpdateAcknowledgeIes(protocolIEsC *C.ProtocolIE_Container_1710P23_t) (*e2appducontents.RicserviceUpdateAcknowledgeIes, error) {
	pIEs := new(e2appducontents.RicserviceUpdateAcknowledgeIes)

	ieCount := int(protocolIEsC.list.count)
	//fmt.Printf("1544P1 Type %T Count %v Size %v\n", *protocolIEsC.list.array, protocolIEsC.list.count, protocolIEsC.list.size)
	for i := 0; i < ieCount; i++ {
		offset := unsafe.Sizeof(unsafe.Pointer(*protocolIEsC.list.array)) * uintptr(i)
		rsuIeC := *(**C.RICserviceUpdateAcknowledge_IEs_t)(unsafe.Pointer(uintptr(unsafe.Pointer(protocolIEsC.list.array)) + offset))

		ie, err := decodeRicServiceUpdateAcknowledgeIE(rsuIeC)
		if err != nil {
			return nil, err
		}
		if ie.GetE2ApProtocolIes9() != nil {
			pIEs.E2ApProtocolIes9 = ie.E2ApProtocolIes9
		}

		if ie.GetE2ApProtocolIes13() != nil {
			pIEs.E2ApProtocolIes13 = ie.E2ApProtocolIes13
		}
	}

	return pIEs, nil
}

func newRicServiceUpdateFailureIe(rsuIEs *e2appducontents.RicserviceUpdateFailureIes) (*C.ProtocolIE_Container_1710P24_t, error) {
	pIeC1710P24 := new(C.ProtocolIE_Container_1710P24_t)

	if rsuIEs.GetE2ApProtocolIes2() != nil {
		ie1C, err := newRicServiceUpdateFailureIe2CriticalityDiagnostics(rsuIEs.GetE2ApProtocolIes2())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1710P24), unsafe.Pointer(ie1C)); err != nil {
			return nil, err
		}
	}

	if rsuIEs.GetE2ApProtocolIes13() != nil {
		ie1C, err := newRicServiceUpdateFailureIe13RanFunctionsRejected(rsuIEs.GetE2ApProtocolIes13())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1710P24), unsafe.Pointer(ie1C)); err != nil {
			return nil, err
		}
	}

	if rsuIEs.GetE2ApProtocolIes31() != nil {
		ie1C, err := newRicServiceUpdateFailureIe31TimeToWait(rsuIEs.GetE2ApProtocolIes31())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1710P24), unsafe.Pointer(ie1C)); err != nil {
			return nil, err
		}
	}

	return pIeC1710P24, nil
}

func decodeRicServiceUpdateFailureIes(protocolIEsC *C.ProtocolIE_Container_1710P24_t) (*e2appducontents.RicserviceUpdateFailureIes, error) {
	pIEs := new(e2appducontents.RicserviceUpdateFailureIes)

	ieCount := int(protocolIEsC.list.count)
	//fmt.Printf("1544P1 Type %T Count %v Size %v\n", *protocolIEsC.list.array, protocolIEsC.list.count, protocolIEsC.list.size)
	for i := 0; i < ieCount; i++ {
		offset := unsafe.Sizeof(unsafe.Pointer(*protocolIEsC.list.array)) * uintptr(i)
		rsuIeC := *(**C.RICserviceUpdateFailure_IEs_t)(unsafe.Pointer(uintptr(unsafe.Pointer(protocolIEsC.list.array)) + offset))

		ie, err := decodeRicServiceUpdateFailureIE(rsuIeC)
		if err != nil {
			return nil, err
		}
		if ie.GetE2ApProtocolIes2() != nil {
			pIEs.E2ApProtocolIes2 = ie.E2ApProtocolIes2
		}

		if ie.GetE2ApProtocolIes13() != nil {
			pIEs.E2ApProtocolIes13 = ie.E2ApProtocolIes13
		}

		if ie.GetE2ApProtocolIes31() != nil {
			pIEs.E2ApProtocolIes31 = ie.E2ApProtocolIes31
		}
	}

	return pIEs, nil
}

func newE2nodeConfigurationUpdateIe(rsuIEs *e2appducontents.E2NodeConfigurationUpdateIes) (*C.ProtocolIE_Container_1710P17_t, error) {
	pIeC1710P17 := new(C.ProtocolIE_Container_1710P17_t)

	if rsuIEs.GetValue() != nil {
		ie1C, err := newE2nodeConfigurationUpdateIe33E2nodeComponentConfigUpdateList(rsuIEs)
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1710P17), unsafe.Pointer(ie1C)); err != nil {
			return nil, err
		}
	}

	return pIeC1710P17, nil
}

func decodeE2nodeConfigurationUpdateIes(protocolIEsC *C.ProtocolIE_Container_1710P17_t) (*e2appducontents.E2NodeConfigurationUpdateIes, error) {
	pIEs := new(e2appducontents.E2NodeConfigurationUpdateIes)

	ieCount := int(protocolIEsC.list.count)
	//fmt.Printf("1544P1 Type %T Count %v Size %v\n", *protocolIEsC.list.array, protocolIEsC.list.count, protocolIEsC.list.size)
	for i := 0; i < ieCount; i++ {
		offset := unsafe.Sizeof(unsafe.Pointer(*protocolIEsC.list.array)) * uintptr(i)
		e2ncuIeC := *(**C.E2nodeConfigurationUpdate_IEs_t)(unsafe.Pointer(uintptr(unsafe.Pointer(protocolIEsC.list.array)) + offset))

		ie, err := decodeE2nodeConfigurationUpdateIE(e2ncuIeC)
		if err != nil {
			return nil, err
		}
		if ie != nil {
			pIEs = ie
		}
	}

	return pIEs, nil
}

func newE2connectionUpdateIe(e2cuIEs *e2appducontents.E2ConnectionUpdateIes) (*C.ProtocolIE_Container_1710P14_t, error) {
	pIeC1710P14 := new(C.ProtocolIE_Container_1710P14_t)

	if e2cuIEs.GetE2ApProtocolIes44() != nil {
		ie44C, err := newE2connectionUpdateIe44E2connectionUpdateList(e2cuIEs.GetE2ApProtocolIes44())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1710P14), unsafe.Pointer(ie44C)); err != nil {
			return nil, err
		}
	}

	if e2cuIEs.GetE2ApProtocolIes45() != nil {
		ie45C, err := newE2connectionUpdateIe45E2connectionUpdateList(e2cuIEs.GetE2ApProtocolIes45())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1710P14), unsafe.Pointer(ie45C)); err != nil {
			return nil, err
		}
	}

	if e2cuIEs.GetE2ApProtocolIes46() != nil {
		ie46C, err := newE2connectionUpdateIe46E2connectionUpdateRemoveList(e2cuIEs.GetE2ApProtocolIes46())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1710P14), unsafe.Pointer(ie46C)); err != nil {
			return nil, err
		}
	}

	return pIeC1710P14, nil
}

func decodeE2connectionUpdateIes(protocolIEsC *C.ProtocolIE_Container_1710P14_t) (*e2appducontents.E2ConnectionUpdateIes, error) {
	pIEs := new(e2appducontents.E2ConnectionUpdateIes)

	ieCount := int(protocolIEsC.list.count)
	//fmt.Printf("1544P1 Type %T Count %v Size %v\n", *protocolIEsC.list.array, protocolIEsC.list.count, protocolIEsC.list.size)
	for i := 0; i < ieCount; i++ {
		offset := unsafe.Sizeof(unsafe.Pointer(*protocolIEsC.list.array)) * uintptr(i)
		e2cuIeC := *(**C.E2connectionUpdate_IEs_t)(unsafe.Pointer(uintptr(unsafe.Pointer(protocolIEsC.list.array)) + offset))

		ie, err := decodeE2connectionUpdateIE(e2cuIeC)
		if err != nil {
			return nil, err
		}

		if ie.E2ApProtocolIes44 != nil {
			pIEs.E2ApProtocolIes44 = ie.E2ApProtocolIes44
		}

		if ie.E2ApProtocolIes45 != nil {
			pIEs.E2ApProtocolIes45 = ie.E2ApProtocolIes45
		}

		if ie.E2ApProtocolIes46 != nil {
			pIEs.E2ApProtocolIes46 = ie.E2ApProtocolIes46
		}
	}

	return pIEs, nil
}

func newE2nodeConfigurationUpdateAcknowledgeIe(e2ncuaIEs *e2appducontents.E2NodeConfigurationUpdateAcknowledgeIes) (*C.ProtocolIE_Container_1710P18_t, error) {
	pIeC1710P18 := new(C.ProtocolIE_Container_1710P18_t)

	if e2ncuaIEs.GetValue() != nil {
		ie1C, err := newE2nodeConfigurationUpdateAcknowledgeIe35E2nodeComponentConfigUpdateAckList(e2ncuaIEs)
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1710P18), unsafe.Pointer(ie1C)); err != nil {
			return nil, err
		}
	}

	return pIeC1710P18, nil
}

func decodeE2nodeConfigurationUpdateAcknowledgeIes(protocolIEsC *C.ProtocolIE_Container_1710P18_t) (*e2appducontents.E2NodeConfigurationUpdateAcknowledgeIes, error) {
	pIEs := new(e2appducontents.E2NodeConfigurationUpdateAcknowledgeIes)

	ieCount := int(protocolIEsC.list.count)
	//fmt.Printf("1544P1 Type %T Count %v Size %v\n", *protocolIEsC.list.array, protocolIEsC.list.count, protocolIEsC.list.size)
	for i := 0; i < ieCount; i++ {
		offset := unsafe.Sizeof(unsafe.Pointer(*protocolIEsC.list.array)) * uintptr(i)
		rsuIeC := *(**C.E2nodeConfigurationUpdateAcknowledge_IEs_t)(unsafe.Pointer(uintptr(unsafe.Pointer(protocolIEsC.list.array)) + offset))

		ie, err := decodeE2nodeConfigurationUpdateAcknowledgeIE(rsuIeC)
		if err != nil {
			return nil, err
		}
		if ie != nil {
			pIEs = ie
		}
	}

	return pIEs, nil
}

func newE2connectionUpdateAcknowledgeIe(e2cuaIEs *e2appducontents.E2ConnectionUpdateAckIes) (*C.ProtocolIE_Container_1710P15_t, error) {
	pIeC1710P15 := new(C.ProtocolIE_Container_1710P15_t)

	if e2cuaIEs.GetE2ApProtocolIes39() != nil {
		ie39C, err := newE2connectionUpdateAck39E2connectionUpdateList(e2cuaIEs.GetE2ApProtocolIes39())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1710P15), unsafe.Pointer(ie39C)); err != nil {
			return nil, err
		}
	}

	if e2cuaIEs.GetE2ApProtocolIes40() != nil {
		ie40C, err := newE2connectionUpdateAck40E2connectionSetupFailedList(e2cuaIEs.GetE2ApProtocolIes40())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1710P15), unsafe.Pointer(ie40C)); err != nil {
			return nil, err
		}
	}

	return pIeC1710P15, nil
}

func decodeE2connectionUpdateAcknowledgeIes(protocolIEsC *C.ProtocolIE_Container_1710P15_t) (*e2appducontents.E2ConnectionUpdateAckIes, error) {
	pIEs := new(e2appducontents.E2ConnectionUpdateAckIes)

	ieCount := int(protocolIEsC.list.count)
	//fmt.Printf("1544P1 Type %T Count %v Size %v\n", *protocolIEsC.list.array, protocolIEsC.list.count, protocolIEsC.list.size)
	for i := 0; i < ieCount; i++ {
		offset := unsafe.Sizeof(unsafe.Pointer(*protocolIEsC.list.array)) * uintptr(i)
		e2cuaIeC := *(**C.E2connectionUpdateAck_IEs_t)(unsafe.Pointer(uintptr(unsafe.Pointer(protocolIEsC.list.array)) + offset))

		ie, err := decodeE2connectionUpdateAckIE(e2cuaIeC)
		if err != nil {
			return nil, err
		}
		if ie.E2ApProtocolIes39 != nil {
			pIEs.E2ApProtocolIes39 = ie.E2ApProtocolIes39
		}
		if ie.E2ApProtocolIes40 != nil {
			pIEs.E2ApProtocolIes40 = ie.E2ApProtocolIes40
		}
	}

	return pIEs, nil
}

func newE2nodeConfigurationUpdateFailureIe(rsuIEs *e2appducontents.E2NodeConfigurationUpdateFailureIes) (*C.ProtocolIE_Container_1710P19_t, error) {
	pIeC1710P19 := new(C.ProtocolIE_Container_1710P19_t)

	if rsuIEs.GetE2ApProtocolIes1() != nil {
		ie1C, err := newE2nodeConfigurationUpdateFailureIes1Cause(rsuIEs.GetE2ApProtocolIes1())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1710P19), unsafe.Pointer(ie1C)); err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("newE2nodeConfigurationUpdateFailureIe() Cause should be mandatory present in the message")
	}

	if rsuIEs.GetE2ApProtocolIes2() != nil {
		ie2C, err := newE2nodeConfigurationUpdateFailureIes2CriticalityDiagnostics(rsuIEs.GetE2ApProtocolIes2())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1710P19), unsafe.Pointer(ie2C)); err != nil {
			return nil, err
		}
	}

	if rsuIEs.GetE2ApProtocolIes31() != nil {
		ie31C, err := newE2nodeConfigurationUpdateFailureIes31TimeToWait(rsuIEs.GetE2ApProtocolIes31())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1710P19), unsafe.Pointer(ie31C)); err != nil {
			return nil, err
		}
	}

	return pIeC1710P19, nil
}

func decodeE2nodeConfigurationUpdateFailureIes(protocolIEsC *C.ProtocolIE_Container_1710P19_t) (*e2appducontents.E2NodeConfigurationUpdateFailureIes, error) {
	pIEs := new(e2appducontents.E2NodeConfigurationUpdateFailureIes)

	ieCount := int(protocolIEsC.list.count)
	//fmt.Printf("1544P1 Type %T Count %v Size %v\n", *protocolIEsC.list.array, protocolIEsC.list.count, protocolIEsC.list.size)
	for i := 0; i < ieCount; i++ {
		offset := unsafe.Sizeof(unsafe.Pointer(*protocolIEsC.list.array)) * uintptr(i)
		rsuIeC := *(**C.E2nodeConfigurationUpdateFailure_IEs_t)(unsafe.Pointer(uintptr(unsafe.Pointer(protocolIEsC.list.array)) + offset))

		ie, err := decodeE2nodeConfigurationUpdateFailureIE(rsuIeC)
		if err != nil {
			return nil, err
		}
		if ie.E2ApProtocolIes1 != nil {
			pIEs.E2ApProtocolIes1 = ie.E2ApProtocolIes1
		}

		if ie.E2ApProtocolIes2 != nil {
			pIEs.E2ApProtocolIes2 = ie.E2ApProtocolIes2
		}

		if ie.E2ApProtocolIes31 != nil {
			pIEs.E2ApProtocolIes31 = ie.E2ApProtocolIes31
		}
	}

	return pIEs, nil
}

func newE2connectionUpdateFailureIe(e2cufIEs *e2appducontents.E2ConnectionUpdateFailureIes) (*C.ProtocolIE_Container_1710P16_t, error) {
	pIeC1710P16 := new(C.ProtocolIE_Container_1710P16_t)

	if e2cufIEs.GetE2ApProtocolIes1() != nil {
		ie1C, err := newE2connectionUpdateFailureIes1Cause(e2cufIEs.GetE2ApProtocolIes1())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1710P16), unsafe.Pointer(ie1C)); err != nil {
			return nil, err
		}
	}

	if e2cufIEs.GetE2ApProtocolIes2() != nil {
		ie2C, err := newE2connectionUpdateFailureIes2CriticalityDiagnostics(e2cufIEs.GetE2ApProtocolIes2())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1710P16), unsafe.Pointer(ie2C)); err != nil {
			return nil, err
		}
	}

	if e2cufIEs.GetE2ApProtocolIes31() != nil {
		ie31C, err := newE2connectionUpdateFailureIes31TimeToWait(e2cufIEs.GetE2ApProtocolIes31())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1710P16), unsafe.Pointer(ie31C)); err != nil {
			return nil, err
		}
	}

	return pIeC1710P16, nil
}

func decodeE2connectionUpdateFailureIes(protocolIEsC *C.ProtocolIE_Container_1710P16_t) (*e2appducontents.E2ConnectionUpdateFailureIes, error) {
	pIEs := new(e2appducontents.E2ConnectionUpdateFailureIes)

	ieCount := int(protocolIEsC.list.count)
	//fmt.Printf("1544P1 Type %T Count %v Size %v\n", *protocolIEsC.list.array, protocolIEsC.list.count, protocolIEsC.list.size)
	for i := 0; i < ieCount; i++ {
		offset := unsafe.Sizeof(unsafe.Pointer(*protocolIEsC.list.array)) * uintptr(i)
		rsuIeC := *(**C.E2connectionUpdateFailure_IEs_t)(unsafe.Pointer(uintptr(unsafe.Pointer(protocolIEsC.list.array)) + offset))

		ie, err := decodeE2connectionUpdateFailureIE(rsuIeC)
		if err != nil {
			return nil, err
		}
		if ie.E2ApProtocolIes1 != nil {
			pIEs.E2ApProtocolIes1 = ie.E2ApProtocolIes1
		}

		if ie.E2ApProtocolIes2 != nil {
			pIEs.E2ApProtocolIes2 = ie.E2ApProtocolIes2
		}

		if ie.E2ApProtocolIes31 != nil {
			pIEs.E2ApProtocolIes31 = ie.E2ApProtocolIes31
		}
	}

	return pIEs, nil
}

func newE2setupFailureIe(e2sfIEs *e2appducontents.E2SetupFailureIes) (*C.ProtocolIE_Container_1710P13_t, error) {
	pIeC1710P13 := new(C.ProtocolIE_Container_1710P13_t)

	if e2sfIEs.GetE2ApProtocolIes1() != nil {
		ie1C, err := newE2setupFailureIe1Cause(e2sfIEs.GetE2ApProtocolIes1())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1710P13), unsafe.Pointer(ie1C)); err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("newE2setupFailureIe(): Cause should be mandatory present in the message")
	}

	if e2sfIEs.GetE2ApProtocolIes2() != nil {
		ie2C, err := newE2setupIe2CriticalityDiagnostics(e2sfIEs.GetE2ApProtocolIes2())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1710P13), unsafe.Pointer(ie2C)); err != nil {
			return nil, err
		}
	}

	if e2sfIEs.GetE2ApProtocolIes31() != nil {
		ie31C, err := newE2setupFailureIe31TimeToWait(e2sfIEs.GetE2ApProtocolIes31())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1710P13), unsafe.Pointer(ie31C)); err != nil {
			return nil, err
		}
	}

	if e2sfIEs.GetE2ApProtocolIes48() != nil {
		ie48C, err := newE2setupFailureIe48Tnlinformation(e2sfIEs.GetE2ApProtocolIes48())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1710P13), unsafe.Pointer(ie48C)); err != nil {
			return nil, err
		}
	}

	return pIeC1710P13, nil
}

func decodeE2setupFailureIes(protocolIEsC *C.ProtocolIE_Container_1710P13_t) (*e2appducontents.E2SetupFailureIes, error) {
	pIEs := new(e2appducontents.E2SetupFailureIes)

	ieCount := int(protocolIEsC.list.count)
	//fmt.Printf("1544P1 Type %T Count %v Size %v\n", *protocolIEsC.list.array, protocolIEsC.list.count, protocolIEsC.list.size)
	for i := 0; i < ieCount; i++ {
		offset := unsafe.Sizeof(unsafe.Pointer(*protocolIEsC.list.array)) * uintptr(i)
		eiIeC := *(**C.E2setupFailureIEs_t)(unsafe.Pointer(uintptr(unsafe.Pointer(protocolIEsC.list.array)) + offset))

		ie, err := decodeE2setupFailureIE(eiIeC)
		if err != nil {
			return nil, err
		}
		if ie.E2ApProtocolIes1 != nil {
			pIEs.E2ApProtocolIes1 = ie.E2ApProtocolIes1
			//} else {
			//	return nil, fmt. Errorf("decodeE2setupFailureIes(): obtained payload doesn't contain Cause")
		}
		if ie.E2ApProtocolIes2 != nil {
			pIEs.E2ApProtocolIes2 = ie.E2ApProtocolIes2
		}
		if ie.E2ApProtocolIes31 != nil {
			pIEs.E2ApProtocolIes31 = ie.E2ApProtocolIes31
		}
		if ie.E2ApProtocolIes48 != nil {
			pIEs.E2ApProtocolIes48 = ie.E2ApProtocolIes48
		}
	}

	return pIEs, nil
}
