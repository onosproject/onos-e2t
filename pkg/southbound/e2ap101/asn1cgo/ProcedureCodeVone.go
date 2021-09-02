// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "ProcedureCodeVone.h"
import "C"
import (
	"fmt"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta2"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-commondatatypes"
)

func newProcedureCode(pc *e2ap_commondatatypes.ProcedureCode) (C.ProcedureCodeVone_t, error) {
	switch pcT := v1beta2.ProcedureCodeT(pc.GetValue()); pcT {
	case v1beta2.ProcedureCodeIDE2setup:
		return C.ProcedureCodeVone_id_E2setupVone, nil
	case v1beta2.ProcedureCodeIDErrorIndication:
		return C.ProcedureCodeVone_id_ErrorIndicationVone, nil
	case v1beta2.ProcedureCodeIDReset:
		return C.ProcedureCodeVone_id_ResetVone, nil
	case v1beta2.ProcedureCodeIDRICcontrol:
		return C.ProcedureCodeVone_id_RICcontrolVone, nil
	case v1beta2.ProcedureCodeIDRICindication:
		return C.ProcedureCodeVone_id_RICindicationVone, nil
	case v1beta2.ProcedureCodeIDRICserviceQuery:
		return C.ProcedureCodeVone_id_RICserviceQueryVone, nil
	case v1beta2.ProcedureCodeIDRICserviceUpdate:
		return C.ProcedureCodeVone_id_RICserviceUpdateVone, nil
	case v1beta2.ProcedureCodeIDRICsubscription:
		return C.ProcedureCodeVone_id_RICsubscriptionVone, nil
	case v1beta2.ProcedureCodeIDRICsubscriptionDelete:
		return C.ProcedureCodeVone_id_RICsubscriptionDeleteVone, nil
	default:
		return 0, fmt.Errorf("unexpected procedure code %v", pcT)
	}
}

func decodeProcedureCode(pc C.ProcedureCodeVone_t) *e2ap_commondatatypes.ProcedureCode {
	return &e2ap_commondatatypes.ProcedureCode{
		Value: int32(pc),
	}
}
