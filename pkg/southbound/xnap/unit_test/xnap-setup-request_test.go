//  SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
//  SPDX-License-Identifier: Apache-2.0

package unit_test

import (
	"encoding/hex"
	v1 "github.com/onosproject/onos-e2t/api/xnap/v1"
	xnapcommondatatypesv1 "github.com/onosproject/onos-e2t/api/xnap/v1/xnap-commondatatypes"
	xnapiesv1 "github.com/onosproject/onos-e2t/api/xnap/v1/xnap-ies"
	xnappducontentsv1 "github.com/onosproject/onos-e2t/api/xnap/v1/xnap-pdu-contents"
	xnappdudescriptionsv1 "github.com/onosproject/onos-e2t/api/xnap/v1/xnap-pdu-descriptions"
	"github.com/onosproject/onos-e2t/pkg/southbound/xnap/encoder"
	"github.com/onosproject/onos-e2t/pkg/southbound/xnap/pdubuilder"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"gotest.tools/assert"
	"testing"
)

func TestXnSetupReques(t *testing.T) {

	list := make([]*xnappducontentsv1.XnSetupRequestIEs, 0)

	// creating GlobalNgRanNodeID
	plmnID, err := pdubuilder.CreatePlmnIdentity([]byte{0xAA, 0xBB, 0xAA})
	assert.NilError(t, err)

	enbID, err := pdubuilder.CreateEnbIDChoiceEnbIDLongmacro(&asn1.BitString{
		Value: []byte{0xBB, 0xBB, 0xF8},
		Len:   21,
	})
	assert.NilError(t, err)

	ngEnbID, err := pdubuilder.CreateGlobalngeNbID(plmnID, enbID)
	assert.NilError(t, err)

	ranNodeID, err := pdubuilder.CreateGlobalNgRAnnodeIDNgENb(ngEnbID)
	assert.NilError(t, err)

	val, err := pdubuilder.CreateXnSetupRequestIEsValueIDGlobalNgRanNodeID(ranNodeID)
	assert.NilError(t, err)

	item1 := &xnappducontentsv1.XnSetupRequestIEs{
		Id:          int32(v1.ProtocolIeIDGlobalNGRANnodeID),
		Criticality: xnapcommondatatypesv1.Criticality_CRITICALITY_REJECT,
		Value:       val,
	}
	list = append(list, item1)

	// creating TAIsupportList
	taiSupportList := make([]*xnapiesv1.TaisupportItem, 0)

	tac, err := pdubuilder.CreateTac([]byte{0xDD, 0xDD, 0xDD})
	assert.NilError(t, err)

	broadcastPlmns := make([]*xnapiesv1.BroadcastPlmninTaisupportItem, 0)

	sliceList := make([]*xnapiesv1.SNSsai, 0)
	snssai1 := &xnapiesv1.SNSsai{
		Sst: []byte{0xFF},
		Sd:  []byte{0xBB, 0xAA, 0xBB},
	}
	snssai2 := &xnapiesv1.SNSsai{
		Sst: []byte{0xFF},
		Sd:  []byte{0xCC, 0xBB, 0xCC},
	}
	sliceList = append(sliceList, snssai1)
	sliceList = append(sliceList, snssai2)

	taiSliceSupportList, err := pdubuilder.CreateSliceSupportList(sliceList)
	assert.NilError(t, err)

	plmnID2, err := pdubuilder.CreatePlmnIdentity([]byte{0xCC, 0xDD, 0xCC})
	assert.NilError(t, err)

	plmnItem, err := pdubuilder.CreateBroadcastPlmninTaisupportItem(plmnID2, taiSliceSupportList)
	assert.NilError(t, err)
	broadcastPlmns = append(broadcastPlmns, plmnItem)

	taiItem, err := pdubuilder.CreateTaisupportItem(tac, broadcastPlmns)
	assert.NilError(t, err)

	taiSupportList = append(taiSupportList, taiItem)

	val2, err := pdubuilder.CreateXnSetupRequestIEsValueIDTaisupportList(&xnapiesv1.TaisupportList{
		Value: taiSupportList,
	})
	assert.NilError(t, err)

	item2 := &xnappducontentsv1.XnSetupRequestIEs{
		Id:          int32(v1.ProtocolIeIDTAISupportlist),
		Criticality: xnapcommondatatypesv1.Criticality_CRITICALITY_REJECT,
		Value:       val2,
	}
	list = append(list, item2)

	// creating AMFRegionInformation
	amfInfoList := make([]*xnapiesv1.GlobalAmfRegionInformation, 0)

	plmnID3, err := pdubuilder.CreatePlmnIdentity([]byte{0xDD, 0xEE, 0xDD})
	assert.NilError(t, err)

	amfInfoItem, err := pdubuilder.CreateGlobalAmfRegionInformation(plmnID3, &asn1.BitString{
		Value: []byte{0xEE},
		Len:   8,
	})
	assert.NilError(t, err)

	amfInfoList = append(amfInfoList, amfInfoItem)
	amfInfoList = append(amfInfoList, amfInfoItem)
	amfInfoList = append(amfInfoList, amfInfoItem)

	val3, err := pdubuilder.CreateXnSetupRequestIEsValueIDAmfRegionInformation(&xnapiesv1.AmfRegionInformation{
		Value: amfInfoList,
	})
	assert.NilError(t, err)

	item3 := &xnappducontentsv1.XnSetupRequestIEs{
		Id:          int32(v1.ProtocolIeIDAMFRegionInformation),
		Criticality: xnapcommondatatypesv1.Criticality_CRITICALITY_REJECT,
		Value:       val3,
	}
	list = append(list, item3)

	// creating list of served cells NR
	servedCellsNRList := make([]*xnapiesv1.ServedCellsNRItem, 0)

	plmnID4, err := pdubuilder.CreatePlmnIdentity([]byte{0xEE, 0xFF, 0xEE})
	assert.NilError(t, err)

	nrCellidentity, err := pdubuilder.CreateNrCellIdentity(&asn1.BitString{
		Value: []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xF0},
		Len:   36,
	})
	assert.NilError(t, err)

	nrcgi, err := pdubuilder.CreateNrCGi(plmnID4, nrCellidentity)
	assert.NilError(t, err)

	broadcastPlmnsList := make([]*xnapiesv1.PlmnIdentity, 0)
	broadcastPlmnsList = append(broadcastPlmnsList, plmnID)
	broadcastPlmnsList = append(broadcastPlmnsList, plmnID2)
	broadcastPlmnsList = append(broadcastPlmnsList, plmnID3)
	broadcastPlmnsList = append(broadcastPlmnsList, plmnID4)

	freqBandList := make([]*xnapiesv1.NrfrequencyBandItem, 0)

	freqBand, err := pdubuilder.CreateNrfrequencyBand(11)
	assert.NilError(t, err)

	supportedSulBandList := make([]*xnapiesv1.SupportedSulbandItem, 0)
	sulBandItem := &xnapiesv1.SupportedSulbandItem{
		SulBandItem: &xnapiesv1.SulFrequencyBand{
			Value: 1,
		},
	}
	supportedSulBandList = append(supportedSulBandList, sulBandItem)

	freqBandItem := &xnapiesv1.NrfrequencyBandItem{
		NrFrequencyBand: freqBand,
		SupportedSulBandList: &xnapiesv1.SupportedSulbandList{
			Value: supportedSulBandList,
		},
	}
	freqBandList = append(freqBandList, freqBandItem)

	nrfreqInfo := &xnapiesv1.NrfrequencyInfo{
		NrArfcn: &xnapiesv1.Nrarfcn{
			Value: 32,
		},
		FrequencyBandList: &xnapiesv1.NrfrequencyBandList{
			Value: freqBandList,
		},
	}

	transmBW, err := pdubuilder.CreateNrtransmissionBandwIDth(pdubuilder.CreateNrscsScs120(), pdubuilder.CreateNrnrbNrb24())
	assert.NilError(t, err)

	nrModeInfo, err := pdubuilder.CreateNrmodeInfoFdd(nrfreqInfo, nrfreqInfo, transmBW, transmBW)
	assert.NilError(t, err)

	nrModeInfoch, err := pdubuilder.CreateNrmodeInfoFddChoice(nrModeInfo)
	assert.NilError(t, err)

	connSupport, err := pdubuilder.CreateConnectivitySupport(pdubuilder.CreateENdcsupportConnectivitySupportNotSupported())
	assert.NilError(t, err)

	servedCellInfoNr := &xnapiesv1.ServedCellInformationNR{
		NrPci: &xnapiesv1.Nrpci{
			Value: 1007,
		},
		CellId: nrcgi,
		Tac: &xnapiesv1.Tac{
			Value: []byte{0xAA, 0x00, 0xAA},
		},
		Ranac: &xnapiesv1.Ranac{
			Value: 255,
		},
		BroadcastPlmn: &xnapiesv1.BroadcastPlmns{
			Value: broadcastPlmnsList,
		},
		NrModeInfo:                     nrModeInfoch,
		MeasurementTimingConfiguration: []byte{0xFF, 0x00, 0x00, 0xFF},
		ConnectivitySupport:            connSupport,
	}

	neighbourInfoNrList := make([]*xnapiesv1.NeighbourInformationNRItem, 0)
	neighbourInfoNrList = append(neighbourInfoNrList, &xnapiesv1.NeighbourInformationNRItem{
		NrPci: &xnapiesv1.Nrpci{
			Value: 1007,
		},
		NrCgi: nrcgi,
		Tac: &xnapiesv1.Tac{
			Value: []byte{0xAA, 0x00, 0xAA},
		},
		Ranac: &xnapiesv1.Ranac{
			Value: 255,
		},
		NrModeInfo: &xnapiesv1.NeighbourInformationNRModeInfo{
			NeighbourInformationNrModeInfo: &xnapiesv1.NeighbourInformationNRModeInfo_FddInfo{
				FddInfo: &xnapiesv1.NeighbourInformationNRModeFddinfo{
					UlNrFreqInfo: nrfreqInfo,
					DlNrFequInfo: nrfreqInfo,
				},
			},
		},
		MeasurementTimingConfiguration: []byte{0xFF, 0x00, 0x00, 0xFF},
		ConnectivitySupport:            connSupport,
	})

	neighbourInfoEutraList := make([]*xnapiesv1.NeighbourInformationEUTraItem, 0)

	eutraCellidentity, err := pdubuilder.CreateEUTraCellIdentity(&asn1.BitString{
		Value: []byte{0xFF, 0xFF, 0xFF, 0xF0},
		Len:   28,
	})
	assert.NilError(t, err)

	eutracgi, err := pdubuilder.CreateEUTraCGi(plmnID, eutraCellidentity)
	assert.NilError(t, err)

	eutraItem := &xnapiesv1.NeighbourInformationEUTraItem{
		EUtraCgi: eutracgi,
		EUtraPci: &xnapiesv1.EUTrapci{
			Value: 503,
		},
		Earfcn: &xnapiesv1.EUTraarfcn{
			Value: 0,
		},
		Tac: tac,
		Ranac: &xnapiesv1.Ranac{
			Value: 1,
		},
	}
	neighbourInfoEutraList = append(neighbourInfoEutraList, eutraItem)

	nrItem := &xnapiesv1.ServedCellsNRItem{
		ServedCellInfoNr: servedCellInfoNr,
		NeighbourInfoNr: &xnapiesv1.NeighbourInformationNR{
			Value: neighbourInfoNrList,
		},
		NeighbourInfoEUtra: &xnapiesv1.NeighbourInformationEUTra{
			Value: neighbourInfoEutraList,
		},
	}
	servedCellsNRList = append(servedCellsNRList, nrItem)

	val4, err := pdubuilder.CreateXnSetupRequestIEsValueIDListOfServedCellsNr(&xnapiesv1.ServedCellsNR{
		Value: servedCellsNRList,
	})
	assert.NilError(t, err)

	item4 := &xnappducontentsv1.XnSetupRequestIEs{
		Id:          int32(v1.ProtocolIeIDListofservedcellsNR),
		Criticality: xnapcommondatatypesv1.Criticality_CRITICALITY_REJECT,
		Value:       val4,
	}
	list = append(list, item4)

	// creating list of served cells EUTRA
	servedCellsEutraList := make([]*xnapiesv1.ServedCellsEUTraItem, 0)

	bcPlmns := make([]*xnapiesv1.ServedCellInformationEUTraperBplmn, 0)
	bcPlmn, err := pdubuilder.CreateServedCellInformationEUTraperBplmn(plmnID2)
	assert.NilError(t, err)
	bcPlmns = append(bcPlmns, bcPlmn)

	specialSubFrameInfo, err := pdubuilder.CreateSpecialSubframeInfoEUTra(
		pdubuilder.CreateSpecialSubframePatternsEUtraSsp3(),
		pdubuilder.CreateCyclicPrefixEUtraDlExtended(),
		pdubuilder.CreateCyclicPrefixEUtraUlNormal())
	assert.NilError(t, err)

	eutraTdd := &xnapiesv1.ServedCellInformationEUTraTDdinfo{
		Earfcn: &xnapiesv1.EUTraarfcn{
			Value: 262143,
		},
		EUtraTxBw:           pdubuilder.CreateEUtratransmissionBandwidthBw15(),
		SubframeAssignmnet:  pdubuilder.CreateSubframeAssignmnetServedCellInformationEutratddinfoSa4(),
		SpecialSubframeInfo: specialSubFrameInfo,
	}

	eutraModeInfoCh, err := pdubuilder.CreateServedCellInformationEUTraModeInfoTdd(eutraTdd)
	assert.NilError(t, err)

	servedCellInfoEutra := &xnapiesv1.ServedCellInformationEUTra{
		EUtraPci: &xnapiesv1.EUTrapci{
			Value: 502,
		},
		EUtraCgi:       eutracgi,
		Tac:            tac,
		BroadcastPlmns: bcPlmns,
		EUtraModeInfo:  eutraModeInfoCh,
	}

	servedCellsEutraItem := &xnapiesv1.ServedCellsEUTraItem{
		ServedCellInfoEUtra: servedCellInfoEutra,
	}
	servedCellsEutraList = append(servedCellsEutraList, servedCellsEutraItem)

	val5, err := pdubuilder.CreateXnSetupRequestIEsValueIDListOfServedCellsEUtra(&xnapiesv1.ServedCellsEUTra{
		Value: servedCellsEutraList,
	})
	assert.NilError(t, err)

	item5 := &xnappducontentsv1.XnSetupRequestIEs{
		Id:          int32(v1.ProtocolIeIDListofservedcellsEUTRA),
		Criticality: xnapcommondatatypesv1.Criticality_CRITICALITY_REJECT,
		Value:       val5,
	}
	list = append(list, item5)

	xnSetupRequest, err := pdubuilder.CreateXnSetupRequest(list)
	assert.NilError(t, err)

	newXnApPdu := &xnappdudescriptionsv1.XnApPDu{
		XnApPdu: &xnappdudescriptionsv1.XnApPDu_InitiatingMessage{
			InitiatingMessage: &xnappdudescriptionsv1.InitiatingMessage{
				ProcedureCode: int32(v1.ProcedureCodeIDxnSetup),
				Criticality:   xnapcommondatatypesv1.Criticality_CRITICALITY_REJECT,
				Value: &xnappdudescriptionsv1.InitiatingMessageXnApElementaryProcedures{
					ImValues: &xnappdudescriptionsv1.InitiatingMessageXnApElementaryProcedures_XnSetupRequest{
						XnSetupRequest: xnSetupRequest,
					},
				},
			},
		},
	}
	t.Logf("Created message is\n%v", newXnApPdu)

	per, err := encoder.PerEncodeXnApPdu(newXnApPdu)
	assert.NilError(t, err)
	t.Logf("F1SetupRequest F1AP PDU PER with Go APER library\n%v", hex.Dump(per))

	result, err := encoder.PerDecodeXnApPdu(per)
	assert.NilError(t, err)
	t.Logf("Decoded message is\n%v", result)
	assert.DeepEqual(t, newXnApPdu.String(), result.String())
}
