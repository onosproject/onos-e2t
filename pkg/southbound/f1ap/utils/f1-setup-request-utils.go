// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package f1utils

import (
	"fmt"
	"github.com/onosproject/onos-api/go/onos/topo"
	f1apiesv1 "github.com/onosproject/onos-e2t/api/f1ap/v1/f1ap_ies"
	f1appducontentsv1 "github.com/onosproject/onos-e2t/api/f1ap/v1/f1ap_pdu_contents"
	"github.com/onosproject/onos-e2t/pkg/utils/decode"
	"github.com/onosproject/onos-lib-go/pkg/logging"
)

var log = logging.GetLogger()

func hasCGICellID(cell *f1appducontentsv1.GnbDUServedCellsItemIes) bool {
	return cell != nil && cell.GetValue() != nil && cell.GetValue().GetGnbDUServedCellsItem() != nil &&
		cell.GetValue().GetGnbDUServedCellsItem().GetServedCellInformation() != nil &&
		cell.GetValue().GetGnbDUServedCellsItem().GetServedCellInformation().GetNRcgi() != nil &&
		cell.GetValue().GetGnbDUServedCellsItem().GetServedCellInformation().GetNRcgi().NRcellIdentity != nil &&
		cell.GetValue().GetGnbDUServedCellsItem().GetServedCellInformation().GetNRcgi().GetNRcellIdentity().GetValue() != nil
}

func GetCGICellID(cell *f1appducontentsv1.GnbDUServedCellsItemIes) (string, error) {
	if !hasCGICellID(cell) {
		errMsg := "CGI does not have Cell ID"
		log.Debug(errMsg)
		return "", fmt.Errorf(errMsg)
	}
	return fmt.Sprintf("%x", *decode.Asn1BitstringToUint64(cell.GetValue().GetGnbDUServedCellsItem().GetServedCellInformation().GetNRcgi().NRcellIdentity.GetValue())), nil
}

func hasCGIPlmnID(cell *f1appducontentsv1.GnbDUServedCellsItemIes) bool {
	return cell != nil && cell.GetValue() != nil && cell.GetValue().GetGnbDUServedCellsItem() != nil &&
		cell.GetValue().GetGnbDUServedCellsItem().GetServedCellInformation() != nil &&
		cell.GetValue().GetGnbDUServedCellsItem().GetServedCellInformation().GetNRcgi() != nil &&
		cell.GetValue().GetGnbDUServedCellsItem().GetServedCellInformation().GetNRcgi().PLmnIdentity != nil
}

func GetCGIPlmnID(cell *f1appducontentsv1.GnbDUServedCellsItemIes) (uint64, error) {
	if !hasCGIPlmnID(cell) {
		errMsg := "PlmnID does not have Cell ID"
		log.Debug(errMsg)
		return 0, fmt.Errorf(errMsg)
	}
	return *decode.Asn1BytesToUint64(cell.GetValue().GetGnbDUServedCellsItem().GetServedCellInformation().GetNRcgi().PLmnIdentity.GetValue()), nil
}

func hasServedPlmns(cell *f1appducontentsv1.GnbDUServedCellsItemIes) bool {
	return cell != nil && cell.GetValue().GetGnbDUServedCellsItem().GetServedCellInformation().GetServedPlmns() != nil &&
		cell.GetValue().GetGnbDUServedCellsItem().GetServedCellInformation().GetServedPlmns().GetValue() != nil
}

func GetServedPlmns(cell *f1appducontentsv1.GnbDUServedCellsItemIes) []uint32 {
	result := make([]uint32, 0)
	if !hasServedPlmns(cell) {
		errMsg := "served plmn is nil"
		log.Debug(errMsg)
		return result
	}

	for _, sPlmn := range cell.GetValue().GetGnbDUServedCellsItem().GetServedCellInformation().GetServedPlmns().GetValue() {
		if sPlmn.PLmnIdentity != nil {
			result = append(result, uint32(*decode.Asn1BytesToUint64(sPlmn.PLmnIdentity.GetValue())))
		}
	}

	return result
}

func hasPCI(cell *f1appducontentsv1.GnbDUServedCellsItemIes) bool {
	return cell != nil && cell.GetValue() != nil && cell.GetValue().GetGnbDUServedCellsItem() != nil &&
		cell.GetValue().GetGnbDUServedCellsItem().GetServedCellInformation() != nil && cell.GetValue().GetGnbDUServedCellsItem().GetServedCellInformation().GetNRpci() != nil
}

func GetPCI(cell *f1appducontentsv1.GnbDUServedCellsItemIes) (uint32, error) {
	if !hasPCI(cell) {
		errMsg := "pci is nil"
		log.Debug(errMsg)
		return 0, fmt.Errorf(errMsg)
	}

	return uint32(cell.GetValue().GetGnbDUServedCellsItem().GetServedCellInformation().GetNRpci().GetValue()), nil
}

func hasNRModeInfo(cell *f1appducontentsv1.GnbDUServedCellsItemIes) bool {
	return cell.GetValue().GetGnbDUServedCellsItem().GetServedCellInformation().GetNRModeInfo() != nil
}

func hasTDDInfo(cell *f1appducontentsv1.GnbDUServedCellsItemIes) bool {
	return hasNRModeInfo(cell) && cell.GetValue().GetGnbDUServedCellsItem().GetServedCellInformation().GetNRModeInfo().GetTDd() != nil
}

func hasFDDInfo(cell *f1appducontentsv1.GnbDUServedCellsItemIes) bool {
	return hasNRModeInfo(cell) && cell.GetValue().GetGnbDUServedCellsItem().GetServedCellInformation().GetNRModeInfo().GetFDd() != nil
}

func hasTDDFreqInfo(cell *f1appducontentsv1.GnbDUServedCellsItemIes) bool {
	return hasTDDInfo(cell) && cell.GetValue().GetGnbDUServedCellsItem().GetServedCellInformation().GetNRModeInfo().GetTDd().GetNRfreqInfo() != nil
}

func hasTDDTransmissionBandwidth(cell *f1appducontentsv1.GnbDUServedCellsItemIes) bool {
	return hasTDDInfo(cell) && cell.GetValue().GetGnbDUServedCellsItem().GetServedCellInformation().GetNRModeInfo().GetTDd().GetTransmissionBandwidth() != nil
}

func GetARFCN(freqInfo *f1apiesv1.NrfreqInfo) int32 {
	return freqInfo.GetNRarfcn()
}

func hasFDDUlFreqInfo(cell *f1appducontentsv1.GnbDUServedCellsItemIes) bool {
	return hasFDDInfo(cell) && cell.GetValue().GetGnbDUServedCellsItem().GetServedCellInformation().GetNRModeInfo().GetFDd().GetULNrfreqInfo() != nil
}

func hasFDDDlFreqInfo(cell *f1appducontentsv1.GnbDUServedCellsItemIes) bool {
	return hasFDDInfo(cell) && cell.GetValue().GetGnbDUServedCellsItem().GetServedCellInformation().GetNRModeInfo().GetFDd().GetDLNrfreqInfo() != nil
}

func hasFDDUlTransmissionBandwidth(cell *f1appducontentsv1.GnbDUServedCellsItemIes) bool {
	return hasFDDInfo(cell) && cell.GetValue().GetGnbDUServedCellsItem().GetServedCellInformation().GetNRModeInfo().GetFDd().GetDLTransmissionBandwidth() != nil
}

func hasFDDDlTransmissionBandwidth(cell *f1appducontentsv1.GnbDUServedCellsItemIes) bool {
	return hasFDDInfo(cell) && cell.GetValue().GetGnbDUServedCellsItem().GetServedCellInformation().GetNRModeInfo().GetFDd().GetULTransmissionBandwidth() != nil
}

func GetTDDInfo(cell *f1appducontentsv1.GnbDUServedCellsItemIes) (*topo.E2Cell_TddInfo, error) {
	result := &topo.E2Cell_TddInfo{
		TddInfo: &topo.TDDInfo{},
	}
	if !hasTDDInfo(cell) {
		errMsg := "tdd info is nil"
		log.Debug(errMsg)
		return nil, fmt.Errorf(errMsg)
	}

	if hasTDDFreqInfo(cell) {
		freqBandItems := make([]*topo.FrequencyBandItem, 0)
		for _, f := range cell.GetValue().GetGnbDUServedCellsItem().GetServedCellInformation().GetNRModeInfo().GetTDd().GetNRfreqInfo().FreqBandListNr {
			if f != nil {
				freqBandItems = append(freqBandItems, &topo.FrequencyBandItem{
					NrFrequencyBand: uint32(f.GetFreqBandIndicatorNr()),
				})
			}
		}
		result.TddInfo.NrFreqInfo = &topo.FrequencyInfo{
			NrArfcn: uint32(GetARFCN(cell.GetValue().GetGnbDUServedCellsItem().GetServedCellInformation().GetNRModeInfo().GetTDd().GetNRfreqInfo())),
			FrequencyBandList: &topo.FrequencyBandList{
				FrequencyBandItems: freqBandItems,
			},
		}
	}

	if hasTDDTransmissionBandwidth(cell) {
		nrb := cell.GetValue().GetGnbDUServedCellsItem().GetServedCellInformation().GetNRModeInfo().GetTDd().GetTransmissionBandwidth().GetNRnrb() + 1
		scs := cell.GetValue().GetGnbDUServedCellsItem().GetServedCellInformation().GetNRModeInfo().GetTDd().GetTransmissionBandwidth().GetNRscs() + 1
		result.TddInfo.TransmissionBandwidth = &topo.TransmissionBandwidth{
			Nrb:   topo.Nrb(nrb),
			NrScs: topo.NrScs(scs),
		}
	}

	return result, nil
}

func GetFDDInfo(cell *f1appducontentsv1.GnbDUServedCellsItemIes) (*topo.E2Cell_FddInfo, error) {
	result := &topo.E2Cell_FddInfo{
		FddInfo: &topo.FDDInfo{},
	}
	if !hasFDDInfo(cell) {
		errMsg := "fdd info is nil"
		log.Debug(errMsg)
		return nil, fmt.Errorf(errMsg)
	}

	if hasFDDDlFreqInfo(cell) {
		freqBandItems := make([]*topo.FrequencyBandItem, 0)
		for _, f := range cell.GetValue().GetGnbDUServedCellsItem().GetServedCellInformation().GetNRModeInfo().GetFDd().GetDLNrfreqInfo().FreqBandListNr {
			if f != nil {
				freqBandItems = append(freqBandItems, &topo.FrequencyBandItem{
					NrFrequencyBand: uint32(f.FreqBandIndicatorNr),
				})
			}
		}
		result.FddInfo.DlFreqInfo = &topo.FrequencyInfo{
			NrArfcn: uint32(GetARFCN(cell.GetValue().GetGnbDUServedCellsItem().GetServedCellInformation().GetNRModeInfo().GetFDd().GetDLNrfreqInfo())),
			FrequencyBandList: &topo.FrequencyBandList{
				FrequencyBandItems: freqBandItems,
			},
		}
	}

	if hasFDDUlFreqInfo(cell) {
		freqBandItems := make([]*topo.FrequencyBandItem, 0)
		for _, f := range cell.GetValue().GetGnbDUServedCellsItem().GetServedCellInformation().GetNRModeInfo().GetFDd().GetULNrfreqInfo().FreqBandListNr {
			if f != nil {
				freqBandItems = append(freqBandItems, &topo.FrequencyBandItem{
					NrFrequencyBand: uint32(f.FreqBandIndicatorNr),
				})
			}
		}
		result.FddInfo.UlFreqInfo = &topo.FrequencyInfo{
			NrArfcn: uint32(GetARFCN(cell.GetValue().GetGnbDUServedCellsItem().GetServedCellInformation().GetNRModeInfo().GetFDd().GetULNrfreqInfo())),
			FrequencyBandList: &topo.FrequencyBandList{
				FrequencyBandItems: freqBandItems,
			},
		}
	}

	if hasFDDDlTransmissionBandwidth(cell) {
		nrb := cell.GetValue().GetGnbDUServedCellsItem().GetServedCellInformation().GetNRModeInfo().GetFDd().GetDLTransmissionBandwidth().GetNRnrb() + 1
		scs := cell.GetValue().GetGnbDUServedCellsItem().GetServedCellInformation().GetNRModeInfo().GetFDd().GetDLTransmissionBandwidth().GetNRscs() + 1
		result.FddInfo.DlTransmissionBandwidth = &topo.TransmissionBandwidth{
			Nrb:   topo.Nrb(nrb),
			NrScs: topo.NrScs(scs),
		}
	}

	if hasFDDUlTransmissionBandwidth(cell) {
		nrb := cell.GetValue().GetGnbDUServedCellsItem().GetServedCellInformation().GetNRModeInfo().GetFDd().GetULTransmissionBandwidth().GetNRnrb() + 1
		scs := cell.GetValue().GetGnbDUServedCellsItem().GetServedCellInformation().GetNRModeInfo().GetFDd().GetULTransmissionBandwidth().GetNRscs() + 1
		result.FddInfo.UlTransmissionBandwidth = &topo.TransmissionBandwidth{
			Nrb:   topo.Nrb(nrb),
			NrScs: topo.NrScs(scs),
		}
	}
	return result, nil
}

func GetMeasurementTimingConfiguration(cell *f1appducontentsv1.GnbDUServedCellsItemIes) uint64 {
	return *decode.Asn1BytesToUint64(cell.GetValue().GetGnbDUServedCellsItem().GetServedCellInformation().GetMeasurementTimingConfiguration())
}
