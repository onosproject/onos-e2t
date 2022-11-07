//  SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
//  SPDX-License-Identifier: Apache-2.0

package unit_test

import (
	"encoding/hex"
	v1 "github.com/onosproject/onos-e2t/api/f1ap/v1"
	f1apcommondatatypesv1 "github.com/onosproject/onos-e2t/api/f1ap/v1/f1ap_commondatatypes"
	f1apiesv1 "github.com/onosproject/onos-e2t/api/f1ap/v1/f1ap_ies"
	f1appducontentsv1 "github.com/onosproject/onos-e2t/api/f1ap/v1/f1ap_pdu_contents"
	f1appdudescriptionsv1 "github.com/onosproject/onos-e2t/api/f1ap/v1/f1ap_pdu_descriptions"
	"github.com/onosproject/onos-e2t/pkg/southbound/f1ap/encoder"
	"github.com/onosproject/onos-e2t/pkg/southbound/f1ap/pdubuilder"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"gotest.tools/assert"
	"testing"
)

func TestF1SetupReques(t *testing.T) {

	list := make([]*f1appducontentsv1.F1SetupRequestIes, 0)

	// transaction ID
	trID, err := pdubuilder.CreateTransactionID(1)
	assert.NilError(t, err)
	ie1Value, err := pdubuilder.CreateF1SetupRequestIesValueTransactionID(trID)
	assert.NilError(t, err)
	ie1, err := pdubuilder.CreateF1SetupRequestIes(&f1apcommondatatypesv1.ProtocolIeID{
		Value: int32(v1.ProtocolIeIDTransactionID),
	}, f1apcommondatatypesv1.Criticality_CRITICALITY_REJECT, ie1Value)
	assert.NilError(t, err)
	list = append(list, ie1)

	// GnbDuID
	gnbDuID, err := pdubuilder.CreateGnbDUID(21)
	assert.NilError(t, err)
	ie2Value, err := pdubuilder.CreateF1SetupRequestIesValueGnbDuID(gnbDuID)
	assert.NilError(t, err)
	ie2, err := pdubuilder.CreateF1SetupRequestIes(&f1apcommondatatypesv1.ProtocolIeID{
		Value: int32(v1.ProtocolIeIDgNBDUID),
	}, f1apcommondatatypesv1.Criticality_CRITICALITY_REJECT, ie2Value)
	assert.NilError(t, err)
	list = append(list, ie2)

	// GnbDuRRC version
	rrcVersion, err := pdubuilder.CreateRrcVersion(&asn1.BitString{
		Value: []byte{0xE0},
		Len:   3,
	})
	assert.NilError(t, err)
	ie3Value, err := pdubuilder.CreateF1SetupRequestIesValueRrcVersion(rrcVersion)
	assert.NilError(t, err)
	ie3, err := pdubuilder.CreateF1SetupRequestIes(&f1apcommondatatypesv1.ProtocolIeID{
		Value: int32(v1.ProtocolIeIDGNBDURRCVersion),
	}, f1apcommondatatypesv1.Criticality_CRITICALITY_REJECT, ie3Value)
	assert.NilError(t, err)
	list = append(list, ie3)

	// Served Cells List
	plmnID, err := pdubuilder.CreatePlmnIdentity([]byte{0xFF, 0xDD, 0xFF})
	assert.NilError(t, err)
	nrCellID, err := pdubuilder.CreateNrcellIdentity(&asn1.BitString{
		Value: []byte{0xAA, 0xBB, 0xCC, 0xDD, 0xF0},
		Len:   36,
	})
	assert.NilError(t, err)
	nrcgi, err := pdubuilder.CreateNrcgi(plmnID, nrCellID)
	assert.NilError(t, err)

	nrpci, err := pdubuilder.CreateNrpci(17)
	assert.NilError(t, err)

	plmnList := make([]*f1apiesv1.ServedPlmnsItem, 0)
	plmnItem, err := pdubuilder.CreateServedPlmnsItem(plmnID)
	assert.NilError(t, err)
	plmnList = append(plmnList, plmnItem)
	servedPlmns, err := pdubuilder.CreateServedPlmnsList(plmnList)
	assert.NilError(t, err)

	sulList := make([]*f1apiesv1.SupportedSulfreqBandItem, 0)
	sulItem, err := pdubuilder.CreateSupportedSulfreqBandItem(32)
	assert.NilError(t, err)
	sulList = append(sulList, sulItem)

	fbnrlist := make([]*f1apiesv1.FreqBandNrItem, 0)
	fbnritem, err := pdubuilder.CreateFreqBandNrItem(1, sulList)
	assert.NilError(t, err)
	fbnrlist = append(fbnrlist, fbnritem)

	nrFreqInfo := &f1apiesv1.NrfreqInfo{
		NRarfcn:        0,
		FreqBandListNr: fbnrlist,
	}

	transmissionBW, err := pdubuilder.CreateTransmissionBandwIDth(pdubuilder.CreateNrscsScs120(),
		pdubuilder.CreateNrnrbNrb11())
	assert.NilError(t, err)

	fddInfo, err := pdubuilder.CreateFddInfo(nrFreqInfo, nrFreqInfo, transmissionBW, transmissionBW)
	assert.NilError(t, err)

	nrmode, err := pdubuilder.CreateNrModeInfoFDd(fddInfo)
	assert.NilError(t, err)

	servedCellItem := &f1apiesv1.GnbDUServedCellsItem{
		ServedCellInformation: &f1apiesv1.ServedCellInformation{
			NRcgi:                          nrcgi,
			NRpci:                          nrpci,
			ServedPlmns:                    servedPlmns,
			NRModeInfo:                     nrmode,
			MeasurementTimingConfiguration: []byte{0xF1, 0xF1, 0xF1},
		},
	}

	servedCellItemValue, err := pdubuilder.CreateGnbDUServedCellsItemIesValueGnbDUServedCellsItem(servedCellItem)
	assert.NilError(t, err)

	gnbDuServedCellsList := make([]*f1appducontentsv1.GnbDUServedCellsItemIes, 0)
	gnbDuServedCellItem, err := pdubuilder.CreateGnbDUServedCellsItemIesValue(&f1apcommondatatypesv1.ProtocolIeID{
		Value: int32(v1.ProtocolIeIDGNBDUServedCellsItem),
	}, f1apcommondatatypesv1.Criticality_CRITICALITY_REJECT, servedCellItemValue)
	assert.NilError(t, err)
	gnbDuServedCellsList = append(gnbDuServedCellsList, gnbDuServedCellItem)

	ie4Value, err := pdubuilder.CreateF1SetupRequestIesValueGnbDuServedCellsList(&f1appducontentsv1.GnbDUServedCellsList{
		Value: gnbDuServedCellsList,
	})
	assert.NilError(t, err)
	ie4, err := pdubuilder.CreateF1SetupRequestIes(&f1apcommondatatypesv1.ProtocolIeID{
		Value: int32(v1.ProtocolIeIDgNBDUServedCellsList),
	}, f1apcommondatatypesv1.Criticality_CRITICALITY_REJECT, ie4Value)
	assert.NilError(t, err)
	list = append(list, ie4)

	f1SetupRequest, err := pdubuilder.CreateF1SetupRequest(list)
	assert.NilError(t, err)

	newF1apPdu := &f1appdudescriptionsv1.F1ApPDu{
		F1ApPdu: &f1appdudescriptionsv1.F1ApPDu_InitiatingMessage{
			InitiatingMessage: &f1appdudescriptionsv1.InitiatingMessage{
				ProcedureCode: int32(v1.ProcedureCodeIDF1Setup),
				Criticality:   f1apcommondatatypesv1.Criticality_CRITICALITY_REJECT,
				Value: &f1appdudescriptionsv1.InitiatingMessageF1ApElementaryProcedures{
					ImValues: &f1appdudescriptionsv1.InitiatingMessageF1ApElementaryProcedures_F1SetupRequest{
						F1SetupRequest: f1SetupRequest,
					},
				},
			},
		},
	}
	t.Logf("Created message is\n%v", newF1apPdu)

	per, err := encoder.PerEncodeF1ApPdu(newF1apPdu)
	assert.NilError(t, err)
	t.Logf("F1SetupRequest F1AP PDU PER with Go APER library\n%v", hex.Dump(per))

	result, err := encoder.PerDecodeF1ApPdu(per)
	assert.NilError(t, err)
	t.Logf("Decoded message is\n%v", result)
	assert.DeepEqual(t, newF1apPdu.String(), result.String())
}
