// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package orane2

// #include "ProcedureCode.h"
import "C"
import (
	"fmt"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2proxy/e2ctypes"
)

func procedureCodeToC(procedureCode e2ctypes.ProcedureCodeT) (C.ProcedureCode_t, error) {
	var pcC C.ProcedureCode_t
	switch procedureCode {
	case e2ctypes.ProcedureCodeT_ProcedureCode_id_E2setup:
		pcC = C.ProcedureCode_id_E2setup
	case e2ctypes.ProcedureCodeT_ProcedureCode_id_ErrorIndication:
		pcC = C.ProcedureCode_id_ErrorIndication
	case e2ctypes.ProcedureCodeT_ProcedureCode_id_Reset:
		pcC = C.ProcedureCode_id_Reset
	case e2ctypes.ProcedureCodeT_ProcedureCode_id_RICcontrol:
		pcC = C.ProcedureCode_id_RICcontrol
	case e2ctypes.ProcedureCodeT_ProcedureCode_id_RICindication:
		pcC = C.ProcedureCode_id_RICindication
	case e2ctypes.ProcedureCodeT_ProcedureCode_id_RICserviceQuery:
		pcC = C.ProcedureCode_id_RICserviceQuery
	case e2ctypes.ProcedureCodeT_ProcedureCode_id_RICserviceUpdate:
		pcC = C.ProcedureCode_id_RICserviceUpdate
	case e2ctypes.ProcedureCodeT_ProcedureCode_id_RICsubscription:
		pcC = C.ProcedureCode_id_RICsubscription
	case e2ctypes.ProcedureCodeT_ProcedureCode_id_RICsubscriptionDelete:
		pcC = C.ProcedureCode_id_RICsubscriptionDelete
	case e2ctypes.ProcedureCodeT_ProcedureCode_id_dummy:
		fallthrough
	default:
		return C.ProcedureCode_t(-1), fmt.Errorf("unexpected initiation message procedure code: %d", procedureCode)
	}

	return pcC, nil
}
