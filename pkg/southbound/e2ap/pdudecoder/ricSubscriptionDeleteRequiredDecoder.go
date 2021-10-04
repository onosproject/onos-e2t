// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdudecoder

import (
	e2appdudescriptions "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-descriptions"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
)

func DecodeRicSubscriptionDeleteRequiredPdu(e2apPdu *e2appdudescriptions.E2ApPdu) (
	types.RicSubscriptionWithCauseList, error) {

	//if err := e2apPdu.Validate(); err != nil {
	//	return nil, nil, fmt.Errorf("invalid E2APpdu %s", err.Error())
	//}

	rswcl := make(types.RicSubscriptionWithCauseList)

	list := e2apPdu.GetInitiatingMessage().GetProcedureCode().GetRicSubscriptionDeleteRequired().GetInitiatingMessage().GetProtocolIes().GetValue().GetValue()
	for _, item := range list {
		rswcl[types.RanFunctionID(item.GetValue().GetRanFunctionId().GetValue())] = &types.RicSubscriptionWithCauseItem{
			RicRequestID: types.RicRequest{
				RequestorID: types.RicRequestorID(item.GetValue().GetRicRequestId().GetRicRequestorId()),
				InstanceID:  types.RicInstanceID(item.GetValue().GetRicRequestId().GetRicInstanceId()),
			},
			Cause: item.GetValue().GetCause(),
		}
	}

	return rswcl, nil
}
