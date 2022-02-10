// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdudecoder

import (
	"fmt"
	v2 "github.com/onosproject/onos-e2t/api/e2ap/v2"
	e2appdudescriptions "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-descriptions"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
)

func DecodeRicSubscriptionDeleteRequiredPdu(e2apPdu *e2appdudescriptions.E2ApPdu) (
	types.RicSubscriptionWithCauseList, error) {

	if err := e2apPdu.Validate(); err != nil {
		return types.RicSubscriptionWithCauseList{}, fmt.Errorf("invalid E2APpdu %s", err.Error())
	}

	rswcl := make(types.RicSubscriptionWithCauseList)

	list := e2apPdu.GetInitiatingMessage().GetValue().GetRicSubscriptionDeleteRequired().GetProtocolIes()
	for _, item := range list {
		if item.Id == int32(v2.ProtocolIeIDRICsubscriptionToBeRemoved) {
			innerList := item.GetValue().GetRsdr().GetValue()
			for _, v := range innerList {
				if v.Id == int32(v2.ProtocolIeIDRICsubscriptionWithCauseItem) {
					rswcl[types.RanFunctionID(v.GetValue().GetE2Curi().GetRanFunctionId().GetValue())] = &types.RicSubscriptionWithCauseItem{
						RicRequestID: types.RicRequest{
							RequestorID: types.RicRequestorID(v.GetValue().GetE2Curi().GetRicRequestId().GetRicRequestorId()),
							InstanceID:  types.RicInstanceID(v.GetValue().GetE2Curi().GetRicRequestId().GetRicInstanceId()),
						},
						Cause: v.GetValue().GetE2Curi().GetCause(),
					}
				}
			}
		}
	}

	return rswcl, nil
}
