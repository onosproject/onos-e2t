// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

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
	"fmt"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta2"
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
)

func newRicActionToBeSetupItemIesSingleContainer(rfItemIes *e2appducontents.RicactionToBeSetupItemIes) (*C.ProtocolIE_SingleContainer_1713P0_t, error) {
	return newRicActionToBeSetupItemIEs(rfItemIes)
}

func newRicActionAdmittedItemIEItemIesSingleContainer(raaItemIes *e2appducontents.RicactionAdmittedItemIes) (*C.ProtocolIE_SingleContainer_1713P1_t, error) {
	return newRicActionAdmittedItemIEs(raaItemIes)
}

func newRicActionNotAdmittedItemIEItemIesSingleContainer(ranaItemIes *e2appducontents.RicactionNotAdmittedItemIes) (*C.ProtocolIE_SingleContainer_1713P2_t, error) {
	return newRicActionNotAdmittedItemIEs(ranaItemIes)
}

func newRanFunctionItemIesSingleContainer(rfItemIes *e2appducontents.RanfunctionItemIes) (*C.ProtocolIE_SingleContainer_1713P8_t, error) {
	return newRANfunctionItemIEs(rfItemIes)
}

func newRanFunctionIDItemIesSingleContainer(rfIDItemIes *e2appducontents.RanfunctionIdItemIes) (*C.ProtocolIE_SingleContainer_1713P9_t, error) {
	return newRANfunctionIDItemIEs(rfIDItemIes)
}

func newRanFunctionIDcauseItemIesSingleContainer(rfIDcauseItemIes *e2appducontents.RanfunctionIdcauseItemIes) (*C.ProtocolIE_SingleContainer_1713P10_t, error) {
	return newRANfunctionIDCauseItemIEs(rfIDcauseItemIes)
}

func newE2nodeConfigurationUpdateIesSingleContainer(e2nccuItemIes *e2appducontents.E2NodeComponentConfigUpdateItemIes) (*C.ProtocolIE_SingleContainer_1713P6_t, error) {
	return newE2nodeComponentConfigUpdateItemIEs(e2nccuItemIes)
}

func newE2nodeConfigurationUpdateAckIesSingleContainer(e2nccuaItemIes *e2appducontents.E2NodeComponentConfigUpdateAckItemIes) (*C.ProtocolIE_SingleContainer_1713P7_t, error) {
	return newE2nodeComponentConfigUpdateAckItemIEs(e2nccuaItemIes)
}

func newE2connectionUpdateIesSingleContainer(e2cuItemIes *e2appducontents.E2ConnectionUpdateItemIes) (*C.ProtocolIE_SingleContainer_1713P3_t, error) {
	return newE2connectionUpdateItemIEs(e2cuItemIes)
}

func newE2connectionUpdateRemoveIesSingleContainer(e2curItemIes *e2appducontents.E2ConnectionUpdateRemoveItemIes) (*C.ProtocolIE_SingleContainer_1713P4_t, error) {
	return newE2connectionUpdateRemoveItemIEs(e2curItemIes)
}

func newE2connectionSetupFailedIesSingleContainer(e2csfItemIes *e2appducontents.E2ConnectionSetupFailedItemIes) (*C.ProtocolIE_SingleContainer_1713P5_t, error) {
	return newE2connectionSetupFailedItemIEs(e2csfItemIes)
}

func decodeRicActionToBeSetupItemIesSingleContainer(ratbsIeScC *C.ProtocolIE_SingleContainer_1713P0_t) (*e2appducontents.RicactionToBeSetupItemIes, error) {
	//fmt.Printf("Value %T %v\n", ratbsIeScC, ratbsIeScC)
	switch id := ratbsIeScC.id; id {
	case C.long(v1beta2.ProtocolIeIDRicactionToBeSetupItem):
		return decodeRicActionToBeSetupItemIes(&ratbsIeScC.value)
	default:
		return nil, fmt.Errorf("unexpected id for RicActionToBeSetupItem %v", id)
	}

}

func decodeRicActionAdmittedItemIesSingleContainer(raaiIeScC *C.ProtocolIE_SingleContainer_1713P1_t) (*e2appducontents.RicactionAdmittedItemIes, error) {
	//fmt.Printf("Value %T %v\n", raaiIeScC, raaiIeScC)
	switch id := raaiIeScC.id; id {
	case C.long(v1beta2.ProtocolIeIDRicactionAdmittedItem):
		return decodeRicActionAdmittedIDItemIes(&raaiIeScC.value)
	default:
		return nil, fmt.Errorf("unexpected id for RicactionAdmittedItemIes %v", id)
	}

}

func decodeRicActionNotAdmittedItemIesSingleContainer(ranaiIeScC *C.ProtocolIE_SingleContainer_1713P2_t) (*e2appducontents.RicactionNotAdmittedItemIes, error) {
	//fmt.Printf("Value %T %v\n", ranaiIeScC, ranaiIeScC)
	switch id := ranaiIeScC.id; id {
	case C.long(v1beta2.ProtocolIeIDRicactionNotAdmittedItem):
		return decodeRicActionNotAdmittedIDItemIes(&ranaiIeScC.value)
	default:
		return nil, fmt.Errorf("unexpected id for RicactionNotAdmittedItemIes %v", id)
	}

}

func decodeRanFunctionItemIesSingleContainer(rfiIeScC *C.ProtocolIE_SingleContainer_1713P8_t) (*e2appducontents.RanfunctionItemIes, error) {
	//fmt.Printf("Value %T %v\n", rfiIeScC, rfiIeScC)
	switch id := rfiIeScC.id; id {
	case C.long(v1beta2.ProtocolIeIDRanfunctionItem):
		return decodeRANfunctionItemIes(&rfiIeScC.value)
	default:
		return nil, fmt.Errorf("unexpected id for RanFunctionItem %v", id)
	}

}

func decodeRanFunctionIDItemIesSingleContainer(rfIDiIeScC *C.ProtocolIE_SingleContainer_1713P9_t) (*e2appducontents.RanfunctionIdItemIes, error) {
	//fmt.Printf("Value %T %v\n", rfIDiIeScC, rfIDiIeScC)
	switch id := rfIDiIeScC.id; id {
	case C.long(v1beta2.ProtocolIeIDRanfunctionIDItem):
		return decodeRANfunctionIDItemIes(&rfIDiIeScC.value)
	default:
		return nil, fmt.Errorf("unexpected id for RanfunctionIDItem %v", id)
	}

}

func decodeE2nodeComponentConfigUpdateItemIesSingleContainer(e2nccuiIeScC *C.ProtocolIE_SingleContainer_1713P6_t) (*e2appducontents.E2NodeComponentConfigUpdateItemIes, error) {
	//fmt.Printf("Value %T %v\n", rfIDiIeScC, rfIDiIeScC)
	switch id := e2nccuiIeScC.id; id {
	case C.long(v1beta2.ProtocolIeIDE2nodeComponentConfigUpdateItem):
		return decodeE2nodeComponentConfigUpdateItemIes(&e2nccuiIeScC.value)
	default:
		return nil, fmt.Errorf("unexpected id for E2nodeComponentConfigUpdateItem %v", id)
	}

}

func decodeE2nodeComponentConfigUpdateAckItemIesSingleContainer(e2nccuaiIeScC *C.ProtocolIE_SingleContainer_1713P7_t) (*e2appducontents.E2NodeComponentConfigUpdateAckItemIes, error) {
	//fmt.Printf("Value %T %v\n", rfIDiIeScC, rfIDiIeScC)
	switch id := e2nccuaiIeScC.id; id {
	case C.long(v1beta2.ProtocolIeIDE2nodeComponentConfigUpdateAckItem):
		return decodeE2nodeComponentConfigUpdateAckItemIes(&e2nccuaiIeScC.value)
	default:
		return nil, fmt.Errorf("unexpected id for E2nodeComponentConfigUpdateAckItem %v", id)
	}

}

func decodeE2connectionUpdateItemIesSingleContainer(e2cuiIeScC *C.ProtocolIE_SingleContainer_1713P3_t) (*e2appducontents.E2ConnectionUpdateItemIes, error) {
	//fmt.Printf("Value %T %v\n", rfIDiIeScC, rfIDiIeScC)
	switch id := e2cuiIeScC.id; id {
	case C.long(v1beta2.ProtocolIeIDE2connectionUpdateItem):
		return decodeE2connectionUpdateItemIes(&e2cuiIeScC.value)
	default:
		return nil, fmt.Errorf("unexpected id for E2connectionUpdateItem %v", id)
	}

}

func decodeE2connectionUpdateRemoveItemIesSingleContainer(e2csfiIeScC *C.ProtocolIE_SingleContainer_1713P4_t) (*e2appducontents.E2ConnectionUpdateRemoveItemIes, error) {
	//fmt.Printf("Value %T %v\n", rfIDiIeScC, rfIDiIeScC)
	switch id := e2csfiIeScC.id; id {
	case C.long(v1beta2.ProtocolIeIDE2connectionUpdateRemoveItem):
		return decodeE2connectionUpdateRemoveItemIes(&e2csfiIeScC.value)
	default:
		return nil, fmt.Errorf("unexpected id for E2connectionSetupFailedItem %v", id)
	}

}

func decodeE2connectionSetupFailedItemIesSingleContainer(e2csfiIeScC *C.ProtocolIE_SingleContainer_1713P5_t) (*e2appducontents.E2ConnectionSetupFailedItemIes, error) {
	//fmt.Printf("Value %T %v\n", rfIDiIeScC, rfIDiIeScC)
	switch id := e2csfiIeScC.id; id {
	case C.long(v1beta2.ProtocolIeIDE2connectionSetupFailedItem):
		return decodeE2connectionSetupFailedItemIes(&e2csfiIeScC.value)
	default:
		return nil, fmt.Errorf("unexpected id for E2connectionSetupFailedItem %v", id)
	}

}

func decodeRanFunctionIDCauseItemIesSingleContainer(rfIDciIeScC *C.ProtocolIE_SingleContainer_1713P10_t) (*e2appducontents.RanfunctionIdcauseItemIes, error) {
	//fmt.Printf("Value %T %v\n", rfIDciIeScC, rfIDciIeScC)
	switch id := rfIDciIeScC.id; id {
	case C.long(v1beta2.ProtocolIeIDRanfunctionIeCauseItem):
		return decodeRANfunctionIDCauseItemIes(&rfIDciIeScC.value)
	default:
		return nil, fmt.Errorf("unexpected id for RanfunctionIeCauseItem %v", id)
	}

}
