// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package orane2

//#include "ProtocolIE-Container.h"
//#include "ProtocolIE-Field.h"
import "C"
import (
	"fmt"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2appducontents"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2proxy/e2ctypes"
	"unsafe"
)

// Deprecated: Do not use.
func newProtocolIeContainer1544P0(rsrIEs *e2ctypes.ProtocolIE_Container_1544P0T) (*C.ProtocolIE_Container_1544P0_t, error) {
	pIeC1544P0 := C.ProtocolIE_Container_1544P0_t{}
	for _, ie := range rsrIEs.List {
		ieC, err := newRICsubscriptionRequestIE(ie)
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(&pIeC1544P0), unsafe.Pointer(ieC)); err != nil {
			return nil, err
		}
	}

	return &pIeC1544P0, nil
}

// Deprecated: Do not use.
func newProtocolIeContainer1544P6(riv *e2ctypes.ProtocolIE_Container_1544P6T) (*C.ProtocolIE_Container_1544P6_t, error) {
	pIeC1544P6 := C.ProtocolIE_Container_1544P6_t{}

	for _, ie := range riv.GetList() {
		ieC, err := newRicIndicationIe(ie)
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(&pIeC1544P6), unsafe.Pointer(ieC)); err != nil {
			return nil, err
		}
	}

	return &pIeC1544P6, nil
}

// Deprecated: Do not use.
func newProtocolIeContainer1544P11(esv *e2ctypes.ProtocolIE_Container_1544P11T) (*C.ProtocolIE_Container_1544P11_t, error) {
	pIeC1544P11 := C.ProtocolIE_Container_1544P11_t{}

	for _, ie := range esv.GetList() {
		ieC, err := newE2setupRequestIeOld(ie)
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(&pIeC1544P11), unsafe.Pointer(ieC)); err != nil {
			return nil, err
		}
	}

	return &pIeC1544P11, nil
}

// Deprecated: Do not use.
func newProtocolIeContainer1544P12(e2srIEs *e2ctypes.ProtocolIE_Container_1544P12T) (*C.ProtocolIE_Container_1544P12_t, error) {
	pIeC1544P12 := C.ProtocolIE_Container_1544P12_t{}
	for _, ie := range e2srIEs.List {
		ieC, err := newE2setupResponseIE(ie)
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(&pIeC1544P12), unsafe.Pointer(ieC)); err != nil {
			return nil, err
		}
	}

	return &pIeC1544P12, nil
}

// Deprecated: Do not use.
func decodeProtocolIeContainer1544P0(protocolIEsC *C.ProtocolIE_Container_1544P0_t) (*e2ctypes.ProtocolIE_Container_1544P0T, error) {
	pIEs := &e2ctypes.ProtocolIE_Container_1544P0T{
		List: make([]*e2ctypes.RICsubscriptionRequest_IEsT, 0),
	}

	ieCount := int(protocolIEsC.list.count)
	fmt.Printf("1544P0 Type %T Count %v Size %v\n", *protocolIEsC.list.array, protocolIEsC.list.count, protocolIEsC.list.size)
	for i := 0; i < ieCount; i++ {
		rsrIeC := (*C.RICsubscriptionRequest_IEs_t)(unsafe.Pointer(uintptr(unsafe.Pointer(*protocolIEsC.list.array)) + uintptr(protocolIEsC.list.size*C.int(i))))

		ie, err := decodeRICsubscriptionRequestIE(rsrIeC)
		if err != nil {
			return nil, err
		}
		pIEs.List = append(pIEs.List, ie)
	}

	return pIEs, nil
}

// Deprecated: Do not use.
func decodeProtocolIeContainer1544P1(protocolIEsC *C.ProtocolIE_Container_1544P1_t) (*e2ctypes.ProtocolIE_Container_1544P1T, error) {
	pIEs := &e2ctypes.ProtocolIE_Container_1544P1T{
		List: make([]*e2ctypes.RICsubscriptionResponse_IEsT, 0),
	}

	//fmt.Printf("1544P1 Type %T Count %v Size %v\nValue[0] %v\n", protocolIEsC.list.array, protocolIEsC.list.count, protocolIEsC.list.size, **protocolIEsC.list.array)
	var i C.int
	for i = 0; i < protocolIEsC.list.count; i++ {
		rsrIeC := (**C.RICsubscriptionResponse_IEs_t)(unsafe.Pointer(uintptr(unsafe.Pointer(*protocolIEsC.list.array)) + uintptr(protocolIEsC.list.size*2*i))) //Not sure why it needs to be multiplied by 2

		ie, err := decodeRICsubscriptionResponseIEOld(*rsrIeC)
		if err != nil {
			return nil, err
		}
		pIEs.List = append(pIEs.List, ie)
	}

	return pIEs, nil
}

// Deprecated: Do not use.
func decodeProtocolIeContainer1544P10(protocolIEsC *C.ProtocolIE_Container_1544P10_t) (*e2ctypes.ProtocolIE_Container_1544P10T, error) {
	pIEs := &e2ctypes.ProtocolIE_Container_1544P10T{
		List: make([]*e2ctypes.ErrorIndication_IEsT, 0),
	}

	ieCount := int(protocolIEsC.list.count)
	//fmt.Printf("1544 Type %T Count %v Size %v\n", unsafe.Pointer(*protocolIEsC.list.array), protocolIEsC.list.count, protocolIEsC.list.size)
	for i := 0; i < ieCount; i++ {
		listC := unsafe.Pointer(*protocolIEsC.list.array)
		errIndIeC := (*C.ErrorIndication_IEs_t)(unsafe.Pointer(uintptr(listC) + uintptr(protocolIEsC.list.size*C.int(i))))

		ie, err := decodeErrorIndicationIE(errIndIeC)
		if err != nil {
			return nil, err
		}
		pIEs.List = append(pIEs.List, ie)
	}

	return pIEs, nil
}

// Deprecated: Do not use.
func decodeProtocolIeContainer1544P11(protocolIEsC *C.ProtocolIE_Container_1544P11_t) (*e2ctypes.ProtocolIE_Container_1544P11T, error) {
	pIEs := &e2ctypes.ProtocolIE_Container_1544P11T{
		List: make([]*e2ctypes.E2SetupRequestIEsT, 0),
	}

	ieCount := int(protocolIEsC.list.count)
	fmt.Printf("1544P11 Type %T Count %v Size %v\n", *protocolIEsC.list.array, protocolIEsC.list.count, protocolIEsC.list.size)
	for i := 0; i < ieCount; i++ {
		listC := unsafe.Pointer(*protocolIEsC.list.array)
		e2srIeC := (*C.E2setupRequestIEs_t)(unsafe.Pointer(uintptr(listC) + uintptr(protocolIEsC.list.size*C.int(i))))

		ie, err := decodeE2setupRequestIEOld(e2srIeC)
		if err != nil {
			return nil, err
		}
		pIEs.List = append(pIEs.List, ie)
	}

	return pIEs, nil
}

func newE2SetupRequestIes(esv *e2appducontents.E2SetupRequestIes) (*C.ProtocolIE_Container_1544P11_t, error) {
	pIeC1544P11 := new(C.ProtocolIE_Container_1544P11_t)

	if esv.GetE2ApProtocolIes3() != nil {
		ie3C, err := newE2setupRequestIe3GlobalE2NodeID(esv.GetE2ApProtocolIes3())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1544P11), unsafe.Pointer(ie3C)); err != nil {
			return nil, err
		}
	}

	if esv.GetE2ApProtocolIes10() != nil {
		ie10C, err := newE2setupRequestIe10RanFunctionList(esv.GetE2ApProtocolIes10())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1544P11), unsafe.Pointer(ie10C)); err != nil {
			return nil, err
		}
	}

	return pIeC1544P11, nil
}

func decodeE2SetupRequestIes(protocolIEsC *C.ProtocolIE_Container_1544P11_t) (*e2appducontents.E2SetupRequestIes, error) {
	pIEs := new(e2appducontents.E2SetupRequestIes)

	ieCount := int(protocolIEsC.list.count)
	fmt.Printf("1544P11 Type %T Count %v Size %v\n", *protocolIEsC.list.array, protocolIEsC.list.count, protocolIEsC.list.size)
	for i := 0; i < ieCount; i++ {
		listC := unsafe.Pointer(*protocolIEsC.list.array)
		e2srIeC := (*C.E2setupRequestIEs_t)(unsafe.Pointer(uintptr(listC) + uintptr(protocolIEsC.list.size*C.int(i))))

		ie, err := decodeE2setupRequestIE(e2srIeC)
		if err != nil {
			return nil, err
		}
		if ie.E2ApProtocolIes3 != nil {
			pIEs.E2ApProtocolIes3 = ie.E2ApProtocolIes3
		}
		if ie.E2ApProtocolIes10 != nil {
			pIEs.E2ApProtocolIes10 = ie.E2ApProtocolIes10
		}
	}

	return pIEs, nil
}

func newE2SetupResponseIes(e2srIEs *e2appducontents.E2SetupResponseIes) (*C.ProtocolIE_Container_1544P12_t, error) {
	pIeC1544P12 := new(C.ProtocolIE_Container_1544P12_t)

	if e2srIEs.GetE2ApProtocolIes4() != nil {
		ie4C, err := newE2setupResponseIe4GlobalRicId(e2srIEs.GetE2ApProtocolIes4())
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1544P12), unsafe.Pointer(ie4C)); err != nil {
			return nil, err
		}
	}

	// TODO: comment back in once these are implemented
	//if e2srIEs.GetE2ApProtocolIes9() != nil {
	//	ie9C, err := newE2setupResponseIe9RanFunctionsAccepted(e2srIEs.GetE2ApProtocolIes9())
	//	if err != nil {
	//		return nil, err
	//	}
	//	if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1544P12), unsafe.Pointer(ie9C)); err != nil {
	//		return nil, err
	//	}
	//}
	//
	//if e2srIEs.GetE2ApProtocolIes13() != nil {
	//	ie13C, err := newE2setupResponseIe13RanFunctionsRejected(e2srIEs.GetE2ApProtocolIes13())
	//	if err != nil {
	//		return nil, err
	//	}
	//	if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1544P12), unsafe.Pointer(ie13C)); err != nil {
	//		return nil, err
	//	}
	//}
	return pIeC1544P12, nil
}

func decodeRicSubscriptionResponseIes(protocolIEsC *C.ProtocolIE_Container_1544P1_t) (*e2appducontents.RicsubscriptionResponseIes, error) {
	pIEs := new(e2appducontents.RicsubscriptionResponseIes)

	ieCount := int(protocolIEsC.list.count)
	fmt.Printf("1544P1 Type %T Count %v Size %v\n", *protocolIEsC.list.array, protocolIEsC.list.count, protocolIEsC.list.size)
	for i := 0; i < ieCount; i++ {
		rsrIeC := (**C.RICsubscriptionResponse_IEs_t)(unsafe.Pointer(uintptr(unsafe.Pointer(*protocolIEsC.list.array)) + uintptr(protocolIEsC.list.size*2*C.int(i)))) //Not sure why it needs to be multiplied by 2

		ie, err := decodeRicSubscriptionResponseIE(*rsrIeC)
		if err != nil {
			return nil, err
		}
		if ie.E2ApProtocolIes5 != nil {
			pIEs.E2ApProtocolIes5 = ie.E2ApProtocolIes5
		}
		if ie.E2ApProtocolIes17 != nil {
			pIEs.E2ApProtocolIes17 = ie.E2ApProtocolIes17
		}
		if ie.E2ApProtocolIes18 != nil {
			pIEs.E2ApProtocolIes18 = ie.E2ApProtocolIes18
		}
		if ie.E2ApProtocolIes29 != nil {
			pIEs.E2ApProtocolIes29 = ie.E2ApProtocolIes29
		}
	}

	return pIEs, nil
}

func newRicSubscriptionRequestIes(rsrIEs *e2appducontents.RicsubscriptionRequestIes) (*C.ProtocolIE_Container_1544P0_t, error) {
	pIeC1544P0 := new(C.ProtocolIE_Container_1544P0_t)

	if rsrIEs.GetE2ApProtocolIes5() != nil {
		ie5C, err := newRicSubscriptionRequestIe5RanFunctionID(rsrIEs.E2ApProtocolIes5)
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1544P0), unsafe.Pointer(ie5C)); err != nil {
			return nil, err
		}
	}

	if rsrIEs.GetE2ApProtocolIes29() != nil {
		ie29C, err := newRicSubscriptionRequestIe29RicRequestID(rsrIEs.E2ApProtocolIes29)
		if err != nil {
			return nil, err
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(pIeC1544P0), unsafe.Pointer(ie29C)); err != nil {
			return nil, err
		}
	}

	//if rsrIEs.GetE2ApProtocolIes30() != nil {
	// TODO: implement me
	//}

	return pIeC1544P0, nil
}

func decodeRicSubscriptionRequestIes(protocolIEsC *C.ProtocolIE_Container_1544P0_t) (*e2appducontents.RicsubscriptionRequestIes, error) {
	pIEs := new(e2appducontents.RicsubscriptionRequestIes)

	ieCount := int(protocolIEsC.list.count)
	fmt.Printf("1544P0 Type %T Count %v Size %v\n", *protocolIEsC.list.array, protocolIEsC.list.count, protocolIEsC.list.size)
	for i := 0; i < ieCount; i++ {
		listC := unsafe.Pointer(*protocolIEsC.list.array)
		rsrIeC := (*C.RICsubscriptionRequest_IEs_t)(unsafe.Pointer(uintptr(listC) + uintptr(protocolIEsC.list.size*C.int(i))))

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
