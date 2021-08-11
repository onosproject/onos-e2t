// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdudecoder

import (
	"fmt"
	e2ap_pdu_descriptions "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-descriptions"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/types"
)

func DecodeE2connectionUpdateAcknowledgePdu(e2apPdu *e2ap_pdu_descriptions.E2ApPdu) ([]*types.E2ConnectionUpdateItem,
	[]*types.E2ConnectionSetupFailedItem, error) {
	if err := e2apPdu.Validate(); err != nil {
		return nil, nil, fmt.Errorf("invalid E2APpdu %s", err.Error())
	}

	e2cua := e2apPdu.GetSuccessfulOutcome().GetProcedureCode().GetE2ConnectionUpdate()
	if e2cua == nil {
		return nil, nil, fmt.Errorf("error E2APpdu does not have E2connectionUpdateAcknowledge")
	}

	connSetup := make([]*types.E2ConnectionUpdateItem, 0)
	list := e2cua.GetSuccessfulOutcome().GetProtocolIes().GetE2ApProtocolIes39().GetConnectionSetup().GetValue()
	for _, ie := range list {
		item := types.E2ConnectionUpdateItem{}
		item.TnlInformation.TnlAddress = *ie.GetValue().GetTnlInformation().GetTnlAddress()
		item.TnlInformation.TnlPort = *ie.GetValue().GetTnlInformation().GetTnlPort()
		item.TnlUsage = ie.GetValue().GetTnlUsage()
		connSetup = append(connSetup, &item)
	}

	connSetupFail := make([]*types.E2ConnectionSetupFailedItem, 0)
	failedList := e2cua.GetSuccessfulOutcome().GetProtocolIes().GetE2ApProtocolIes40().GetConnectionSetupFailed().GetValue()
	for _, ie := range failedList {
		item := types.E2ConnectionSetupFailedItem{}
		item.TnlInformation.TnlAddress = *ie.GetValue().GetTnlInformation().GetTnlAddress()
		item.TnlInformation.TnlPort = *ie.GetValue().GetTnlInformation().GetTnlPort()
		item.Cause = *ie.GetValue().GetCause()
		connSetupFail = append(connSetupFail, &item)
	}

	return connSetup, connSetupFail, nil
}
