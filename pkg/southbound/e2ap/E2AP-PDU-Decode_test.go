// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package e2ap

import (
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2ap-commondatatypes"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2apies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/asn1cgo"
	"gotest.tools/assert"
	"io/ioutil"
	"testing"
)

func Test_E2AP_PDU_E2SetupRequest_DecodeXer(t *testing.T) {
	e2setupRequestXer, err := ioutil.ReadFile("./test/E2setupRequest.xml")
	assert.NilError(t, err, "Unexpected error when loading file")
	e2apPdu, err := asn1cgo.XerDecodeE2apPdu(e2setupRequestXer)
	assert.NilError(t, err, "Unexpected error when decoding XER payload")
	assert.Assert(t, e2apPdu != nil, "expected a value for the e2apPdu")
	initMsg := e2apPdu.GetInitiatingMessage()

	assert.Assert(t, initMsg != nil, "expected choice to be InitiatingMessage")
	assert.Assert(t, initMsg.GetProcedureCode().GetE2Setup() != nil)
	assert.Equal(t, e2ap_commondatatypes.Criticality_CRITICALITY_REJECT, initMsg.GetProcedureCode().GetE2Setup().GetCriticality().GetCriticality())

	e2Sr := initMsg.GetProcedureCode().GetE2Setup().GetInitiatingMessage()
	assert.Assert(t, e2Sr != nil, "expected a value for the e2Sr")

	globalE2NodeID := e2Sr.GetProtocolIes().GetE2ApProtocolIes3()
	assert.Assert(t, globalE2NodeID != nil)
	assert.Equal(t, int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT), globalE2NodeID.GetCriticality())
	assert.Equal(t, int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY), globalE2NodeID.GetPresence())
	assert.Equal(t, int32(v1beta1.ProtocolIeIDGlobalE2nodeID), globalE2NodeID.GetId())
	globalE2nodeIdgNb, ok := globalE2NodeID.GetValue().GetGlobalE2NodeId().(*e2apies.GlobalE2NodeId_GNb)
	assert.Assert(t, ok, "expected choice to be GlobalE2Node_IDT_GNB")
	assert.Equal(t, "ONF", string(globalE2nodeIdgNb.GNb.GetGlobalGNbId().GetPlmnId().GetValue()))

	gnbID, ok := globalE2nodeIdgNb.GNb.GetGlobalGNbId().GetGnbId().GetGnbIdChoice().(*e2apies.GnbIdChoice_GnbId)
	assert.Assert(t, ok, "expected choice to be GNB_ID_ChoiceT_Gnb_ID")
	assert.Equal(t, uint32(29), gnbID.GnbId.GetLen(), "comparing gnb-id length")
	assert.Equal(t, uint64(0x8877c6b5), gnbID.GnbId.GetValue(), "gnb-id number of bits")
}

func Test_E2AP_PDU_RICsubscriptionResp_DecodeXer(t *testing.T) {
	rICsubscriptionResponseXer, err := ioutil.ReadFile("./test/RICsubscriptionResponse.xml")
	assert.NilError(t, err, "Unexpected error when loading file")
	e2apPdu, err := asn1cgo.XerDecodeE2apPdu(rICsubscriptionResponseXer)
	assert.NilError(t, err, "Unexpected error when decoding XER payload")
	assert.Assert(t, e2apPdu != nil, "expected a value for the e2apPdu")
	successOc := e2apPdu.GetSuccessfulOutcome()

	assert.Assert(t, successOc != nil, "expected choice to be SuccessfulOutcome")
	assert.Assert(t, successOc.GetProcedureCode().GetRicSubscription() != nil)
	assert.Equal(t, e2ap_commondatatypes.Criticality_CRITICALITY_REJECT, successOc.GetProcedureCode().GetRicSubscription().GetCriticality().GetCriticality())

	rsresp := successOc.GetProcedureCode().GetRicSubscription().GetSuccessfulOutcome()
	assert.Assert(t, rsresp != nil, "expected a value for the rsresp")

	ranFunctionID := rsresp.GetProtocolIes().GetE2ApProtocolIes5()
	assert.Assert(t, ranFunctionID != nil)
	assert.Equal(t, int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT), ranFunctionID.GetCriticality())
	assert.Equal(t, int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY), ranFunctionID.GetPresence())
	assert.Equal(t, int32(v1beta1.ProtocolIeIDRanfunctionID), ranFunctionID.GetId())
	assert.Equal(t, int32(20), ranFunctionID.GetValue().GetValue())

	ricActionsAdmitted := rsresp.GetProtocolIes().GetE2ApProtocolIes17()
	assert.Assert(t, ricActionsAdmitted != nil)
	assert.Equal(t, int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT), ricActionsAdmitted.GetCriticality())
	assert.Equal(t, int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY), ricActionsAdmitted.GetPresence())
	assert.Equal(t, int32(v1beta1.ProtocolIeIDRicactionsAdmitted), ricActionsAdmitted.GetId())
	// TODO check the values

	//ricActionsNotAdmitted := rsresp.GetProtocolIes().GetE2ApProtocolIes18()
	//assert.Assert(t, ricActionsNotAdmitted != nil)
	//assert.Equal(t, int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT), ricActionsNotAdmitted.GetCriticality())
	//assert.Equal(t, int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY), ricActionsNotAdmitted.GetPresence())
	//assert.Equal(t, int32(v1beta1.ProtocolIeIDRicactionsNotAdmitted), ricActionsNotAdmitted.GetId())
	// TODO check the values

	ricRequestID := rsresp.GetProtocolIes().GetE2ApProtocolIes29()
	assert.Assert(t, ricRequestID != nil)
	assert.Equal(t, int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT), ricRequestID.GetCriticality())
	assert.Equal(t, int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY), ricRequestID.GetPresence())
	assert.Equal(t, int32(v1beta1.ProtocolIeIDRicrequestID), ricRequestID.GetId())
	assert.Equal(t, int32(22), ricRequestID.GetValue().GetRicRequestorId())
	assert.Equal(t, int32(6), ricRequestID.GetValue().GetRicInstanceId())

}
