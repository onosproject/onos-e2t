// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package orane2

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "ProtocolIE-SingleContainer.h"
//#include "RICaction-ToBeSetup-Item.h"
//#include "ProtocolIE-Field.h"
import "C"
import (
	"encoding/binary"
	"fmt"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2proxy/e2ctypes"
	"unsafe"
)

// Deprecated: Do not use.
func newRICactionTbsItemIEs(tbsItemIEs *e2ctypes.RICaction_ToBeSetup_ItemIEsT) (*C.RICaction_ToBeSetup_ItemIEs_t, error) {
	critC, err := criticalityToCOld(tbsItemIEs.GetCriticality())
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToCOld(tbsItemIEs.GetId())
	if err != nil {
		return nil, err
	}
	var vpr C.RICaction_ToBeSetup_ItemIEs__value_PR

	choiceC := [56]byte{}
	switch choice := tbsItemIEs.GetChoice().(type) {
	case *e2ctypes.RICaction_ToBeSetup_ItemIEsT_RICaction_ToBeSetup_Item:
		vpr = C.RICaction_ToBeSetup_ItemIEs__value_PR_RICaction_ToBeSetup_Item
		tbsItemC, err := newRICactionToBeSetupItem(choice.RICaction_ToBeSetup_Item)
		if err != nil {
			return nil, err
		}
		//fmt.Printf("RICaction_ToBeSetup_ItemIEs__value %v %v %v\n", tbsItemC,
		//	unsafe.Sizeof(tbsItemC.ricActionID),
		//	unsafe.Sizeof(tbsItemC.ricActionType))
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(tbsItemC.ricActionID))
		binary.LittleEndian.PutUint64(choiceC[8:], uint64(tbsItemC.ricActionType))

		if tbsItemC.ricActionDefinition != nil {
			binary.LittleEndian.PutUint64(choiceC[16:], uint64(uintptr(unsafe.Pointer(tbsItemC.ricActionDefinition))))
		}
		if tbsItemC.ricSubsequentAction != nil {
			binary.LittleEndian.PutUint64(choiceC[24:], uint64(uintptr(unsafe.Pointer(tbsItemC.ricSubsequentAction))))
		}
	default:
		return nil, fmt.Errorf("newRICactionTbsItemIEs() %T not yet implemented", choice)
	}

	tbsItemIEsC := C.RICaction_ToBeSetup_ItemIEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_RICaction_ToBeSetup_ItemIEs__value{
			present: vpr,
			choice:  choiceC,
		},
	}
	return &tbsItemIEsC, nil
}

//func newRANfunctionID_ItemIEs(rfIDIEs *e2ctypes.RANfunctionID_ItemIEsT) (*C.RANfunctionID_ItemIEs_t, error) {
//	critC, err := criticalityToCOld(rfIDIEs.GetCriticality())
//	if err != nil {
//		return nil, err
//	}
//	idC, err := protocolIeIDToCOld(rfIDIEs.GetId())
//	if err != nil {
//		return nil, err
//	}
//	var vpr C.RANfunctionID_ItemIEs__value_PR
//
//	choiceC := [40]byte{}
//	switch choice := rfIDIEs.GetChoice().(type) {
//	case *e2ctypes.RANfunctionID_ItemIEsT_RANfunctionIDcause_Item:
//		vpr = C.RANfunctionID_ItemIEs__value_PR_RANfunctionID_Item
//		// TODO figure out the rest
//	default:
//		return nil, fmt.Errorf("newRICactionTbsItemIEs() %T not yet implemented", choice)
//	}
//
//	rfIDIEsC := C.RANfunctionID_ItemIEs_t{
//		id:          idC,
//		criticality: critC,
//		value: C.struct_RANfunctionID_ItemIEs__value{
//			present: vpr,
//			choice:  choiceC,
//		},
//	}
//
//	return &rfIDIEsC, nil
//}
