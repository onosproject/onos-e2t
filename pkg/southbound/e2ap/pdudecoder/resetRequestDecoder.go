// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdudecoder

import (
	"fmt"
	v2 "github.com/onosproject/onos-e2t/api/e2ap/v2"

	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	e2appdudescriptions "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-descriptions"
)

func DecodeResetRequestPdu(e2apPdu *e2appdudescriptions.E2ApPdu) (*e2ap_ies.Cause, *int32, error) {
	if err := e2apPdu.Validate(); err != nil {
		return nil, nil, fmt.Errorf("invalid E2APpdu %s", err.Error())
	}

	rr := e2apPdu.GetInitiatingMessage().GetValue().GetReset_()
	if rr == nil {
		return nil, nil, fmt.Errorf("error E2APpdu does not have ResetRequest")
	}

	var trID int32
	var cause *e2ap_ies.Cause
	for _, v := range rr.GetProtocolIes() {
		if v.Id == int32(v2.ProtocolIeIDTransactionID) {
			trID = v.GetValue().GetTrId().GetValue()
		}
		if v.Id == int32(v2.ProtocolIeIDCause) {
			cause = v.GetValue().GetC()
		}
	}

	return cause, &trID, nil
}
