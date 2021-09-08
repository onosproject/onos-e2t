// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdudecoder

import (
	"fmt"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	e2appdudescriptions "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-descriptions"
)

func DecodeResetRequestPdu(e2apPdu *e2appdudescriptions.E2ApPdu) (*e2ap_ies.Cause, error) {
	if err := e2apPdu.Validate(); err != nil {
		return nil, fmt.Errorf("invalid E2APpdu %s", err.Error())
	}

	rr := e2apPdu.GetInitiatingMessage().GetProcedureCode().GetReset_()
	if rr == nil {
		return nil, fmt.Errorf("error E2APpdu does not have ResetRequest")
	}

	return e2apPdu.GetInitiatingMessage().GetProcedureCode().GetReset_().
		GetInitiatingMessage().GetProtocolIes().GetResetRequestIes1().GetValue(), nil
}
