// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package filter

import (
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2apies"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2appdudescriptions"
)

// Filter is an E2AP message filter
type Filter func(*e2appdudescriptions.E2ApPdu) bool

// RicSubscription filters responses for a RicSubscription response
func RicSubscription(requestID *e2apies.RicrequestId) Filter {
	return func(pdu *e2appdudescriptions.E2ApPdu) bool {
		success, ok := pdu.GetE2ApPdu().(*e2appdudescriptions.E2ApPdu_SuccessfulOutcome)
		if !ok {
			unsuccess, ok := pdu.GetE2ApPdu().(*e2appdudescriptions.E2ApPdu_UnsuccessfulOutcome)
			if !ok {
				return false
			}
			return unsuccess.UnsuccessfulOutcome.ProcedureCode.RicSubscription != nil &&
				(requestID.RicInstanceId == 0 || unsuccess.UnsuccessfulOutcome.ProcedureCode.RicSubscription.InitiatingMessage.ProtocolIes.E2ApProtocolIes29.Value.RicInstanceId == requestID.RicInstanceId) &&
				(requestID.RicRequestorId == 0 || unsuccess.UnsuccessfulOutcome.ProcedureCode.RicSubscription.InitiatingMessage.ProtocolIes.E2ApProtocolIes29.Value.RicRequestorId == requestID.RicRequestorId)
		}
		return success.SuccessfulOutcome.ProcedureCode.RicSubscription != nil &&
			(requestID.RicInstanceId == 0 || success.SuccessfulOutcome.ProcedureCode.RicSubscription.InitiatingMessage.ProtocolIes.E2ApProtocolIes29.Value.RicInstanceId == requestID.RicInstanceId) &&
			(requestID.RicRequestorId == 0 || success.SuccessfulOutcome.ProcedureCode.RicSubscription.InitiatingMessage.ProtocolIes.E2ApProtocolIes29.Value.RicRequestorId == requestID.RicRequestorId)
	}
}

// RicSubscriptionDelete filters responses for a RicSubscription response
func RicSubscriptionDelete(requestID *e2apies.RicrequestId) Filter {
	return func(pdu *e2appdudescriptions.E2ApPdu) bool {
		success, ok := pdu.GetE2ApPdu().(*e2appdudescriptions.E2ApPdu_SuccessfulOutcome)
		if !ok {
			unsuccess, ok := pdu.GetE2ApPdu().(*e2appdudescriptions.E2ApPdu_UnsuccessfulOutcome)
			if !ok {
				return false
			}
			return unsuccess.UnsuccessfulOutcome.ProcedureCode.RicSubscriptionDelete != nil &&
				(requestID.RicInstanceId == 0 || unsuccess.UnsuccessfulOutcome.ProcedureCode.RicSubscriptionDelete.InitiatingMessage.ProtocolIes.E2ApProtocolIes29.Value.RicInstanceId == requestID.RicInstanceId) &&
				(requestID.RicRequestorId == 0 || unsuccess.UnsuccessfulOutcome.ProcedureCode.RicSubscriptionDelete.InitiatingMessage.ProtocolIes.E2ApProtocolIes29.Value.RicRequestorId == requestID.RicRequestorId)
		}
		return success.SuccessfulOutcome.ProcedureCode.RicSubscriptionDelete != nil &&
			(requestID.RicInstanceId == 0 || success.SuccessfulOutcome.ProcedureCode.RicSubscriptionDelete.InitiatingMessage.ProtocolIes.E2ApProtocolIes29.Value.RicInstanceId == requestID.RicInstanceId) &&
			(requestID.RicRequestorId == 0 || success.SuccessfulOutcome.ProcedureCode.RicSubscriptionDelete.InitiatingMessage.ProtocolIes.E2ApProtocolIes29.Value.RicRequestorId == requestID.RicRequestorId)
	}
}

// RicIndication filters responses for a RicIndication response
func RicIndication(requestID *e2apies.RicrequestId) Filter {
	return func(pdu *e2appdudescriptions.E2ApPdu) bool {
		message, ok := pdu.GetE2ApPdu().(*e2appdudescriptions.E2ApPdu_InitiatingMessage)
		if !ok {
			return false
		}
		return message.InitiatingMessage.ProcedureCode.RicIndication != nil &&
			(requestID.RicInstanceId == 0 || message.InitiatingMessage.ProcedureCode.RicIndication.InitiatingMessage.ProtocolIes.E2ApProtocolIes29.Value.RicInstanceId == requestID.RicInstanceId) &&
			(requestID.RicRequestorId == 0 || message.InitiatingMessage.ProcedureCode.RicIndication.InitiatingMessage.ProtocolIes.E2ApProtocolIes29.Value.RicRequestorId == requestID.RicRequestorId)
	}
}
