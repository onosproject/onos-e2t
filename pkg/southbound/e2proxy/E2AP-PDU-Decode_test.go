// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package e2proxy

import (
	"fmt"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2proxy/e2ctypes"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2proxy/orane2"
	"gotest.tools/assert"
	"io/ioutil"
	"testing"
)

const RicreqidXer = `<RICrequestID>
	<ricRequestorID>543210</ricRequestorID>
	<ricInstanceID>6789</ricInstanceID>
</RICrequestID>`

func Test_RicRequestID(t *testing.T) {
	result, err := orane2.XerDecodeRICrequestID([]byte(RicreqidXer))
	assert.NilError(t, err, "Unexpected error when decoding XER payload")
	assert.Equal(t, int64(543210), result.RicRequestorID)
	assert.Equal(t, int64(6789), result.RicInstanceID)
}

func Test_ErrorIndicationRANfunctionID(t *testing.T) {
	errorIndicationPart1Xer, err := ioutil.ReadFile("./test/ErrorIndicationPart1.xml")
	assert.NilError(t, err, "Unexpected error when loading file")
	errorIndication, err := orane2.XerDecodeErrorIndication(errorIndicationPart1Xer)
	assert.NilError(t, err, "Unexpected error when decoding XER payload")
	assert.Assert(t, errorIndication != nil, "expected a value for the errorIndication")
	assert.Equal(t, 1, len(errorIndication.ProtocolIEs.List))
	errIndIe1 := errorIndication.ProtocolIEs.List[0]
	assert.Equal(t, "ProtocolIE_ID_id_RANfunctionID", errIndIe1.Id.String())
	assert.Equal(t, "Criticality_ignore", errIndIe1.Criticality.String())
	errIndIe1Ch, ok := errIndIe1.Choice.(*e2ctypes.ErrorIndication_IEsT_RANfunctionID)
	assert.Assert(t, ok, "Expected to convert to ErrorIndication_IEs_Value_RICrequestID")
	assert.Equal(t, 12345, int(errIndIe1Ch.RANfunctionID))
}

func Test_ErrorIndicationRicRequestID(t *testing.T) {
	errorIndicationPart2Xer, err := ioutil.ReadFile("./test/ErrorIndicationPart2.xml")
	assert.NilError(t, err, "Unexpected error when loading file")
	errorIndication, err := orane2.XerDecodeErrorIndication(errorIndicationPart2Xer)
	assert.NilError(t, err, "Unexpected error when decoding XER payload")
	assert.Assert(t, errorIndication != nil, "expected a value for the errorIndication")
	assert.Equal(t, 1, len(errorIndication.ProtocolIEs.List))
	errIndIe1 := errorIndication.ProtocolIEs.List[0]
	assert.Equal(t, "ProtocolIE_ID_id_RICrequestID", errIndIe1.Id.String())
	assert.Equal(t, "Criticality_ignore", errIndIe1.Criticality.String())
	errIndIe1Ch, ok := errIndIe1.Choice.(*e2ctypes.ErrorIndication_IEsT_RICrequestID)
	assert.Assert(t, ok, "Expected to convert to ErrorIndication_IEs_Value_RICrequestID")
	assert.Equal(t, 543210, int(errIndIe1Ch.RICrequestID.RicRequestorID))
	assert.Equal(t, 6789, int(errIndIe1Ch.RICrequestID.RicInstanceID))
}

func Test_E2AP_PDU_ErrorIndication(t *testing.T) {
	errorIndicationXer, err := ioutil.ReadFile("./test/ErrorIndication.xml")
	assert.NilError(t, err, "Unexpected error when loading file")
	e2apPdu, err := orane2.XerDecodeE2apPdu(errorIndicationXer)
	assert.NilError(t, err, "Unexpected error when decoding XER payload")

	assert.Assert(t, e2apPdu != nil, "expected a value for the e2apPdu")
	initMsg, ok := e2apPdu.GetChoice().(*e2ctypes.E2AP_PDUT_InitiatingMessage)
	assert.Assert(t, ok, "expected choice to be InitiatingMessage")
	assert.Equal(t, "ProcedureCode_id_ErrorIndication", initMsg.InitiatingMessage.GetProcedureCode().String())
	assert.Equal(t, "Criticality_reject", initMsg.InitiatingMessage.GetCriticality().String())

	_, ok = initMsg.InitiatingMessage.GetChoice().(*e2ctypes.InitiatingMessageT_ErrorIndication)
	assert.Assert(t, ok, "expected choice to be ErrorIndication")
	// TODO ensure everything is being decoded correctly - hint InitiatingMessage.go line 127
	//assert.Equal(t, 1, len(errInd.ErrorIndication.ProtocolIEs.List))
}

func Test_E2AP_PDU_E2SetupRequest_DecodeXer(t *testing.T) {
	e2setupRequestXer, err := ioutil.ReadFile("./test/E2setupRequest.xml")
	assert.NilError(t, err, "Unexpected error when loading file")
	e2apPdu, err := orane2.XerDecodeE2apPdu(e2setupRequestXer)
	assert.NilError(t, err, "Unexpected error when decoding XER payload")
	assert.Assert(t, e2apPdu != nil, "expected a value for the e2apPdu")
	initMsg, ok := e2apPdu.GetChoice().(*e2ctypes.E2AP_PDUT_InitiatingMessage)

	assert.Assert(t, ok, "expected choice to be InitiatingMessage")
	assert.Equal(t, "ProcedureCode_id_E2setup", initMsg.InitiatingMessage.GetProcedureCode().String())
	assert.Equal(t, "Criticality_reject", initMsg.InitiatingMessage.GetCriticality().String())

	e2Sr, ok := initMsg.InitiatingMessage.GetChoice().(*e2ctypes.InitiatingMessageT_E2SetupRequest)
	assert.Assert(t, ok, "expected choice to be E2SetupRequest")
	assert.Assert(t, e2Sr != nil, "expected a value for the e2Sr")
	assert.Equal(t, 1, len(e2Sr.E2SetupRequest.GetProtocolIEs().GetList()))
	e2SrIE1 := e2Sr.E2SetupRequest.GetProtocolIEs().GetList()[0]
	assert.Equal(t, "ProtocolIE_ID_id_GlobalE2node_ID", e2SrIE1.Id.String())
	assert.Equal(t, "Criticality_reject", e2SrIE1.Criticality.String())

	gE2nID, ok := e2SrIE1.GetChoice().(*e2ctypes.E2SetupRequestIEsT_GlobalE2Node_ID)
	assert.Assert(t, ok, "expected choice to be GlobalE2Node_ID")
	assert.Assert(t, gE2nID.GlobalE2Node_ID != nil, "gE2nID.GlobalE2Node_ID is nil")

	gE2nbID, ok := gE2nID.GlobalE2Node_ID.GetChoice().(*e2ctypes.GlobalE2Node_IDT_GNB)
	assert.Assert(t, ok, "expected choice to be GlobalE2Node_IDT_GNB")
	assert.Equal(t, "ONF", gE2nbID.GNB.GetGlobalGNB_ID().GetPlmnId())

	gnbID, ok := gE2nbID.GNB.GetGlobalGNB_ID().GetGnbId().GetChoice().(*e2ctypes.GNB_ID_ChoiceT_Gnb_ID)
	assert.Assert(t, ok, "expected choice to be GNB_ID_ChoiceT_Gnb_ID")
	assert.Equal(t, 4, len(gnbID.Gnb_ID.GetBitString()), "comparing gnb-id length")
	assert.Equal(t, uint32(29), gnbID.Gnb_ID.GetNumbits(), "gnb-id number of bits")
	assert.Equal(t, "16b8cef1", fmt.Sprintf("%x", gnbID.Gnb_ID.GetBitString()), "gnb-id number of bits")
}

func Test_E2AP_PDU_RICsubscriptionReq_DecodeXer(t *testing.T) {
	t.Skip()
	rICsubscriptionRequestXer, err := ioutil.ReadFile("./test/RICsubscriptionRequest.xml")
	assert.NilError(t, err, "Unexpected error when loading file")
	e2apPdu, err := orane2.XerDecodeE2apPdu(rICsubscriptionRequestXer)
	assert.NilError(t, err, "Unexpected error when decoding XER payload")
	assert.Assert(t, e2apPdu != nil, "expected a value for the e2apPdu")

}

func Test_E2AP_PDU_RICsubscriptionResp_DecodeXer(t *testing.T) {
	rICsubscriptionResponseXer, err := ioutil.ReadFile("./test/RICsubscriptionResponse.xml")
	assert.NilError(t, err, "Unexpected error when loading file")
	e2apPdu, err := orane2.XerDecodeE2apPdu(rICsubscriptionResponseXer)
	assert.NilError(t, err, "Unexpected error when decoding XER payload")
	assert.Assert(t, e2apPdu != nil, "expected a value for the e2apPdu")

	successOc, ok := e2apPdu.GetChoice().(*e2ctypes.E2AP_PDUT_SuccessfulOutcome)
	assert.Assert(t, ok, "expected choice to be SuccessfulOutcome")
	assert.Equal(t, "ProcedureCode_id_RICsubscription", successOc.SuccessfulOutcome.GetProcedureCode().String())
	assert.Equal(t, "Criticality_reject", successOc.SuccessfulOutcome.GetCriticality().String())

	rsresp, ok := successOc.SuccessfulOutcome.GetChoice().(*e2ctypes.SuccessfulOutcomeT_RICsubscriptionResponse)
	assert.Assert(t, ok, "expected choice to be RICsubscriptionResponse")
	assert.Assert(t, rsresp != nil, "expected a value for the e2Sr")
	assert.Equal(t, 2, len(rsresp.RICsubscriptionResponse.GetProtocolIEs().GetList()))
	rsrespIE1 := rsresp.RICsubscriptionResponse.GetProtocolIEs().GetList()[0]
	assert.Equal(t, "ProtocolIE_ID_id_RICrequestID", rsrespIE1.Id.String())
	assert.Equal(t, "Criticality_reject", rsrespIE1.Criticality.String())
	rsrespRid, ok := rsrespIE1.GetChoice().(*e2ctypes.RICsubscriptionResponse_IEsT_RICrequestID)
	assert.Assert(t, ok, "expected choice to be RICrequestID")
	assert.Assert(t, rsrespRid != nil, "expected a value for the rsrespRid")
	assert.Equal(t, int64(6), rsrespRid.RICrequestID.GetRicInstanceID())
	assert.Equal(t, int64(22), rsrespRid.RICrequestID.GetRicRequestorID())

	rsrespIE2 := rsresp.RICsubscriptionResponse.GetProtocolIEs().GetList()[1]
	assert.Equal(t, "ProtocolIE_ID_id_RICactions_Admitted", rsrespIE2.Id.String())
	assert.Equal(t, "Criticality_reject", rsrespIE2.Criticality.String())

}
