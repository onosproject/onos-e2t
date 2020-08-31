// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package orane2

// #include "Criticality.h"
import "C"
import (
	"fmt"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2proxy/e2ctypes"
)

func criticalityToC(criticality e2ctypes.CriticalityT) (C.Criticality_t, error) {
	var critC C.Criticality_t
	switch criticality {
	case e2ctypes.CriticalityT_Criticality_reject:
		critC = C.Criticality_reject
	case e2ctypes.CriticalityT_Criticality_ignore:
		critC = C.Criticality_ignore
	case e2ctypes.CriticalityT_Criticality_notify:
		critC = C.Criticality_notify
	default:
		return C.Criticality_t(-1), fmt.Errorf("unexpected value for criticality %d", criticality)
	}

	return critC, nil
}
