// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package xnutils

import (
	"fmt"
	"github.com/onosproject/onos-api/go/onos/topo"
	xnapiesv1 "github.com/onosproject/onos-e2t/api/xnap/v1/xnap-ies"
	"github.com/onosproject/onos-e2t/pkg/utils/decode"
	"github.com/onosproject/onos-lib-go/pkg/logging"
)

var log = logging.GetLogger()

func hasCGICellID(cell *xnapiesv1.ServedCellsNRItem) bool {
	return cell.ServedCellInfoNr != nil && cell.ServedCellInfoNr.GetCellId() != nil &&
		cell.ServedCellInfoNr.GetCellId().GetNrCi() != nil && cell.ServedCellInfoNr.GetCellId().GetNrCi().GetValue() != nil
}

func GetCGICellID(cell *xnapiesv1.ServedCellsNRItem) (string, error) {
	if !hasCGICellID(cell) {
		errMsg := "CGI does not have Cell ID"
		log.Debug(errMsg)
		return "", fmt.Errorf(errMsg)
	}
	return fmt.Sprintf("%x", *decode.Asn1BitstringToUint64(cell.ServedCellInfoNr.GetCellId().GetNrCi().GetValue())), nil
}

func hasCGIPlmnID(cell *xnapiesv1.ServedCellsNRItem) bool {
	return cell.ServedCellInfoNr != nil && cell.ServedCellInfoNr.GetCellId() != nil &&
		cell.ServedCellInfoNr.GetCellId().GetPlmnId() != nil
}

func GetCGIPlmnID(cell *xnapiesv1.ServedCellsNRItem) (uint64, error) {
	if !hasCGIPlmnID(cell) {
		errMsg := "PlmnID does not have Cell ID"
		log.Debug(errMsg)
		return 0, fmt.Errorf(errMsg)
	}
	return *decode.Asn1BytesToUint64(cell.ServedCellInfoNr.GetCellId().GetPlmnId().GetValue()), nil
}

func hasGlobalNgRAnnodeIDGnBID(id *xnapiesv1.GlobalNgRAnnodeID) bool {
	return id.GetGNb() != nil && id.GetGNb().GetGnbId() != nil && id.GetGNb().GetGnbId().GetGnbId() != nil && id.GetGNb().GetPlmnId() != nil
}

func hasGlobalNgRAnnodeIDNgEnBID(id *xnapiesv1.GlobalNgRAnnodeID) bool {
	return id.GetNgENb() != nil && id.GetNgENb().GetPlmnId() != nil
}

func hasGlobalNgRAnnodeIDEnBIDMacro(id *xnapiesv1.GlobalNgRAnnodeID) bool {
	return hasGlobalNgRAnnodeIDNgEnBID(id) && id.GetNgENb().GetEnbId() != nil && id.GetNgENb().GetEnbId().GetEnbIdMacro() != nil
}

func hasGlobalNgRAnnodeIDEnBIDLongMacro(id *xnapiesv1.GlobalNgRAnnodeID) bool {
	return hasGlobalNgRAnnodeIDNgEnBID(id) && id.GetNgENb().GetEnbId() != nil && id.GetNgENb().GetEnbId().GetEnbIdLongmacro() != nil
}

func hasGlobalNgRAnnodeIDEnBIDShortMacro(id *xnapiesv1.GlobalNgRAnnodeID) bool {
	return hasGlobalNgRAnnodeIDNgEnBID(id) && id.GetNgENb().GetEnbId() != nil && id.GetNgENb().GetEnbId().GetEnbIdShortmacro() != nil
}

func GetNGRanNodeID(id *xnapiesv1.GlobalNgRAnnodeID) (*topo.GlobalNgRanNodeID, error) {
	result := &topo.GlobalNgRanNodeID{}
	if id == nil {
		errMsg := "Global NG RAN Node ID is nil"
		log.Warn(errMsg)
		return nil, fmt.Errorf(errMsg)
	}

	if hasGlobalNgRAnnodeIDGnBID(id) {
		result.GlobalNgRanNodeId = &topo.GlobalNgRanNodeID_GlobalGnbId{
			GlobalGnbId: &topo.GlobalGnbID{
				PlmnId:   uint32(*decode.Asn1BytesToUint64(id.GetGNb().GetPlmnId().Value)),
				GnbId:    uint32(*decode.Asn1BytesToUint64(id.GetGNb().GetGnbId().GetGnbId().GetValue())),
				GnbIdLen: id.GetGNb().GetGnbId().GetGnbId().GetLen(),
			},
		}
		return result, nil
	}

	if hasGlobalNgRAnnodeIDNgEnBID(id) {
		if hasGlobalNgRAnnodeIDEnBIDMacro(id) {
			result.GlobalNgRanNodeId = &topo.GlobalNgRanNodeID_GlobalNgEnbId{
				GlobalNgEnbId: &topo.GlobalNgEnbID{
					PlmnId: uint32(*decode.Asn1BytesToUint64(id.GetNgENb().GetPlmnId().Value)),
					NgEnbId: &topo.GlobalNgEnbID_MacroNgEnbId{
						MacroNgEnbId: uint32(*decode.Asn1BytesToUint64(id.GetNgENb().EnbId.GetEnbIdMacro().GetValue())),
					},
				},
			}
		} else if hasGlobalNgRAnnodeIDEnBIDLongMacro(id) {
			result.GlobalNgRanNodeId = &topo.GlobalNgRanNodeID_GlobalNgEnbId{
				GlobalNgEnbId: &topo.GlobalNgEnbID{
					PlmnId: uint32(*decode.Asn1BytesToUint64(id.GetNgENb().GetPlmnId().Value)),
					NgEnbId: &topo.GlobalNgEnbID_MacroNgEnbId{
						MacroNgEnbId: uint32(*decode.Asn1BytesToUint64(id.GetNgENb().EnbId.GetEnbIdLongmacro().GetValue())),
					},
				},
			}
		} else if hasGlobalNgRAnnodeIDEnBIDShortMacro(id) {
			result.GlobalNgRanNodeId = &topo.GlobalNgRanNodeID_GlobalNgEnbId{
				GlobalNgEnbId: &topo.GlobalNgEnbID{
					PlmnId: uint32(*decode.Asn1BytesToUint64(id.GetNgENb().GetPlmnId().Value)),
					NgEnbId: &topo.GlobalNgEnbID_MacroNgEnbId{
						MacroNgEnbId: uint32(*decode.Asn1BytesToUint64(id.GetNgENb().EnbId.GetEnbIdShortmacro().GetValue())),
					},
				},
			}
		} else {
			errMsg := "Ng eNB ID does not have neither eNB ID Macro, long Macro, nor short Macro"
			log.Warn(errMsg)
			return nil, fmt.Errorf(errMsg)
		}
		return result, nil
	}

	errMsg := "Global NG RAN Node ID does not have neither gNB ID or NG EnB ID"
	log.Warn(errMsg)
	return nil, fmt.Errorf(errMsg)
}

func hasTaiTac(item *xnapiesv1.TaisupportItem) bool {
	return item != nil && item.Tac != nil
}

func hasBroadcastPlmnPlmn(bp *xnapiesv1.BroadcastPlmninTaisupportItem) bool {
	return bp != nil && bp.PlmnId != nil
}

func hasBroadcastPlmnSliceList(bp *xnapiesv1.BroadcastPlmninTaisupportItem) bool {
	return bp != nil && bp.TAisliceSupportList != nil
}

func GetTAISupportList(list []*xnapiesv1.TaisupportItem) (*topo.TaiSupportList, error) {
	result := &topo.TaiSupportList{
		TaiSupportItems: make([]*topo.TaiSupportItem, 0),
	}
	for _, item := range list {
		tai := &topo.TaiSupportItem{}
		if hasTaiTac(item) {
			tai.Tac = uint32(*decode.Asn1BytesToUint64(item.Tac.Value))
		}

		if len(item.BroadcastPlmns) > 0 {
			tai.BroadcastPlmns = make([]*topo.XnBroadcastPlmn, 0)
		}
		for _, bp := range item.GetBroadcastPlmns() {
			broadcastPlmn := &topo.XnBroadcastPlmn{}
			if hasBroadcastPlmnPlmn(bp) {
				broadcastPlmn.PlmnId = uint32(*decode.Asn1BytesToUint64(bp.PlmnId.Value))
			}

			if hasBroadcastPlmnSliceList(bp) {
				if len(bp.TAisliceSupportList.Value) > 0 {
					broadcastPlmn.TaiSliceSupportList = &topo.TaiSliceSupportList{
						SliceSupportItems: make([]*topo.SliceSupportItem, 0),
					}
				}
				for _, slice := range bp.TAisliceSupportList.Value {
					sliceSupportItem := &topo.SliceSupportItem{
						SNssai: &topo.SNssai{
							Sd:  uint32(*decode.Asn1BytesToUint64(slice.Sd)),
							Sst: uint32(*decode.Asn1BytesToUint64(slice.Sst)),
						},
					}
					broadcastPlmn.TaiSliceSupportList.SliceSupportItems = append(broadcastPlmn.TaiSliceSupportList.SliceSupportItems, sliceSupportItem)
				}
			}

			tai.BroadcastPlmns = append(tai.BroadcastPlmns, broadcastPlmn)
		}

		result.TaiSupportItems = append(result.TaiSupportItems, tai)
	}

	return result, nil
}

func hasAmfRegionInfoPlmn(amf *xnapiesv1.GlobalAmfRegionInformation) bool {
	return amf != nil && amf.PlmnId != nil
}

func hasAmfRegionInfoRegionID(amf *xnapiesv1.GlobalAmfRegionInformation) bool {
	return amf != nil && amf.AmfRegionId != nil
}

func GetAMFRegionList(list []*xnapiesv1.GlobalAmfRegionInformation) (*topo.AmfRegionInformation, error) {
	result := &topo.AmfRegionInformation{
		GlobalAmfRegionInformationItems: make([]*topo.GlobalAmfRegionInformationItem, 0),
	}
	for _, amf := range list {
		amfItem := &topo.GlobalAmfRegionInformationItem{}
		if hasAmfRegionInfoPlmn(amf) {
			amfItem.PlmnId = uint32(*decode.Asn1BytesToUint64(amf.PlmnId.Value))
		}
		if hasAmfRegionInfoRegionID(amf) {
			amfItem.AmfRegionId = uint32(*decode.Asn1BitstringToUint64(amf.AmfRegionId))
		}
		result.GlobalAmfRegionInformationItems = append(result.GetGlobalAmfRegionInformationItems(), amfItem)
	}
	return result, nil
}

func hasPCI(cell *xnapiesv1.ServedCellsNRItem) bool {
	return cell != nil && cell.ServedCellInfoNr != nil && cell.ServedCellInfoNr.NrPci != nil
}

func GetPCI(cell *xnapiesv1.ServedCellsNRItem) (uint32, error) {
	if !hasPCI(cell) {
		errMsg := "pci is nil"
		log.Debug(errMsg)
		return 0, fmt.Errorf(errMsg)
	}

	return uint32(cell.GetServedCellInfoNr().GetNrPci().GetValue()), nil
}

func hasNRModeInfo(cell *xnapiesv1.ServedCellsNRItem) bool {
	return cell != nil && cell.GetServedCellInfoNr() != nil && cell.GetServedCellInfoNr().GetNrModeInfo() != nil
}

func hasTDDInfo(cell *xnapiesv1.ServedCellsNRItem) bool {
	return hasNRModeInfo(cell) && cell.GetServedCellInfoNr().GetNrModeInfo().GetTdd() != nil
}

func hasFDDInfo(cell *xnapiesv1.ServedCellsNRItem) bool {
	return hasNRModeInfo(cell) && cell.GetServedCellInfoNr().GetNrModeInfo().GetFdd() != nil
}

func hasTDDFreqInfo(cell *xnapiesv1.ServedCellsNRItem) bool {
	return hasTDDInfo(cell) && cell.GetServedCellInfoNr().GetNrModeInfo().GetTdd().GetNrFrequencyInfo() != nil &&
		cell.GetServedCellInfoNr().GetNrModeInfo().GetTdd().GetNrFrequencyInfo() != nil
}

func hasTDDTransmissionBandwidth(cell *xnapiesv1.ServedCellsNRItem) bool {
	return hasTDDInfo(cell) && cell.GetServedCellInfoNr().GetNrModeInfo().GetTdd().GetNrTransmissonBandwidth() != nil
}

func hasTDDARFCN(cell *xnapiesv1.ServedCellsNRItem) bool {
	return hasTDDFreqInfo(cell) && cell.GetServedCellInfoNr().GetNrModeInfo().GetTdd().GetNrFrequencyInfo().GetNrArfcn() != nil
}

func hasFDDDlARFCN(cell *xnapiesv1.ServedCellsNRItem) bool {
	return hasFDDDlFreqInfo(cell) && cell.GetServedCellInfoNr().GetNrModeInfo().GetFdd().GetDlNrfrequencyInfo().GetNrArfcn() != nil
}

func hasFDDUlARFCN(cell *xnapiesv1.ServedCellsNRItem) bool {
	return hasFDDUlFreqInfo(cell) && cell.GetServedCellInfoNr().GetNrModeInfo().GetFdd().GetUlNrfrequencyInfo().GetNrArfcn() != nil
}

func hasTDDFrequencyBandList(cell *xnapiesv1.ServedCellsNRItem) bool {
	return hasTDDFreqInfo(cell) && cell.GetServedCellInfoNr().GetNrModeInfo().GetTdd().GetNrFrequencyInfo().GetFrequencyBandList() != nil
}

func hasFDDUlFreqInfo(cell *xnapiesv1.ServedCellsNRItem) bool {
	return hasFDDInfo(cell) && cell.GetServedCellInfoNr().GetNrModeInfo().GetFdd().GetUlNrfrequencyInfo() != nil
}

func hasFDDUlFreqBandList(cell *xnapiesv1.ServedCellsNRItem) bool {
	return hasFDDUlFreqInfo(cell) && cell.GetServedCellInfoNr().GetNrModeInfo().GetFdd().GetUlNrfrequencyInfo().GetFrequencyBandList() != nil
}

func hasFDDDlFreqBandList(cell *xnapiesv1.ServedCellsNRItem) bool {
	return hasFDDDlFreqInfo(cell) && cell.GetServedCellInfoNr().GetNrModeInfo().GetFdd().GetDlNrfrequencyInfo().GetFrequencyBandList() != nil
}

func hasFDDDlFreqInfo(cell *xnapiesv1.ServedCellsNRItem) bool {
	return hasFDDInfo(cell) && cell.GetServedCellInfoNr().GetNrModeInfo().GetFdd().GetDlNrfrequencyInfo() != nil
}

func hasFDDUlTransmissionBandwidth(cell *xnapiesv1.ServedCellsNRItem) bool {
	return hasFDDInfo(cell) && cell.GetServedCellInfoNr().GetNrModeInfo().GetFdd().GetUlNrtransmissonBandwidth() != nil
}

func hasFDDDlTransmissionBandwidth(cell *xnapiesv1.ServedCellsNRItem) bool {
	return hasFDDInfo(cell) && cell.GetServedCellInfoNr().GetNrModeInfo().GetFdd().GetDlNrtransmissonBandwidth() != nil
}

func GetTDDARFCN(cell *xnapiesv1.ServedCellsNRItem) (int32, error) {
	if !hasTDDARFCN(cell) {
		errMsg := "arfcn is nil"
		log.Debug(errMsg)
		return 0, fmt.Errorf(errMsg)
	}
	return cell.GetServedCellInfoNr().GetNrModeInfo().GetTdd().GetNrFrequencyInfo().GetNrArfcn().GetValue(), nil
}

func GetFDDDlARFCN(cell *xnapiesv1.ServedCellsNRItem) (int32, error) {
	if !hasFDDDlARFCN(cell) {
		errMsg := "arfcn is nil"
		log.Debug(errMsg)
		return 0, fmt.Errorf(errMsg)
	}
	return cell.GetServedCellInfoNr().GetNrModeInfo().GetFdd().GetDlNrfrequencyInfo().GetNrArfcn().GetValue(), nil
}

func GetFDDUlARFCN(cell *xnapiesv1.ServedCellsNRItem) (int32, error) {
	if !hasFDDUlARFCN(cell) {
		errMsg := "arfcn is nil"
		log.Debug(errMsg)
		return 0, fmt.Errorf(errMsg)
	}
	return cell.GetServedCellInfoNr().GetNrModeInfo().GetFdd().GetUlNrfrequencyInfo().GetNrArfcn().GetValue(), nil
}

func GetTDDInfo(cell *xnapiesv1.ServedCellsNRItem) (*topo.E2Cell_TddInfo, error) {
	result := &topo.E2Cell_TddInfo{
		TddInfo: &topo.TDDInfo{},
	}
	if !hasTDDInfo(cell) {
		errMsg := "tdd info is nil"
		log.Debug(errMsg)
		return nil, fmt.Errorf(errMsg)
	}

	if !hasTDDFrequencyBandList(cell) {
		errMsg := "frequency band list is nil in tdd info"
		log.Debug(errMsg)
		return nil, fmt.Errorf(errMsg)
	}

	if hasTDDFreqInfo(cell) {
		freqBandItems := make([]*topo.FrequencyBandItem, 0)
		for _, f := range cell.GetServedCellInfoNr().GetNrModeInfo().GetTdd().GetNrFrequencyInfo().GetFrequencyBandList().Value {
			if f.GetNrFrequencyBand() == nil {
				errMsg := "nr frequency band in TDD is nil"
				log.Debug(errMsg)
				continue
			}
			freqBandItems = append(freqBandItems, &topo.FrequencyBandItem{
				NrFrequencyBand: uint32(f.GetNrFrequencyBand().Value),
			})
		}
		result.TddInfo.NrFreqInfo = &topo.FrequencyInfo{}
		arfcn, err := GetTDDARFCN(cell)
		if err == nil {
			result.TddInfo.NrFreqInfo.NrArfcn = uint32(arfcn)
		} else {
			errMsg := "arfcn in tdd is nil"
			log.Debug(errMsg)
		}

		result.TddInfo.NrFreqInfo.FrequencyBandList = &topo.FrequencyBandList{
			FrequencyBandItems: freqBandItems,
		}
	}

	if hasTDDTransmissionBandwidth(cell) {
		nrb := cell.GetServedCellInfoNr().GetNrModeInfo().GetTdd().GetNrTransmissonBandwidth().GetNRnrb() + 1
		scs := cell.GetServedCellInfoNr().GetNrModeInfo().GetTdd().GetNrTransmissonBandwidth().GetNRscs() + 1
		result.TddInfo.TransmissionBandwidth = &topo.TransmissionBandwidth{
			Nrb:   topo.Nrb(nrb),
			NrScs: topo.NrScs(scs),
		}
	}

	return result, nil
}

func GetFDDInfo(cell *xnapiesv1.ServedCellsNRItem) (*topo.E2Cell_FddInfo, error) {
	result := &topo.E2Cell_FddInfo{
		FddInfo: &topo.FDDInfo{},
	}

	if !hasFDDInfo(cell) {
		errMsg := "fdd info is nil"
		log.Debug(errMsg)
		return nil, fmt.Errorf(errMsg)
	}

	if !hasFDDDlFreqBandList(cell) {
		errMsg := "fdd dl freq band list is nil in fdd info"
		log.Debug(errMsg)
		return nil, fmt.Errorf(errMsg)
	}

	if !hasFDDUlFreqBandList(cell) {
		errMsg := "fdd ul freq band list is nil in fdd info"
		log.Debug(errMsg)
		return nil, fmt.Errorf(errMsg)
	}

	if hasFDDDlFreqInfo(cell) {
		freqBandItems := make([]*topo.FrequencyBandItem, 0)
		for _, f := range cell.GetServedCellInfoNr().GetNrModeInfo().GetFdd().GetDlNrfrequencyInfo().GetFrequencyBandList().Value {
			if f.GetNrFrequencyBand() == nil {
				errMsg := "nr frequency band in fdd dl is nil"
				log.Debug(errMsg)
				continue
			}
			freqBandItems = append(freqBandItems, &topo.FrequencyBandItem{
				NrFrequencyBand: uint32(f.GetNrFrequencyBand().Value),
			})
		}
		result.FddInfo.DlFreqInfo = &topo.FrequencyInfo{}
		arfcn, err := GetFDDDlARFCN(cell)
		if err == nil {
			result.FddInfo.DlFreqInfo.NrArfcn = uint32(arfcn)
		} else {
			errMsg := "arfcn in fdd dl info is nil"
			log.Debug(errMsg)
		}
		result.FddInfo.DlFreqInfo.FrequencyBandList = &topo.FrequencyBandList{
			FrequencyBandItems: freqBandItems,
		}
	}

	if hasFDDUlFreqInfo(cell) {
		freqBandItems := make([]*topo.FrequencyBandItem, 0)
		for _, f := range cell.GetServedCellInfoNr().GetNrModeInfo().GetFdd().GetUlNrfrequencyInfo().GetFrequencyBandList().Value {
			if f.GetNrFrequencyBand() == nil {
				errMsg := "nr frequency band in fdd ul is nil"
				log.Debug(errMsg)
				continue
			}
			freqBandItems = append(freqBandItems, &topo.FrequencyBandItem{
				NrFrequencyBand: uint32(f.GetNrFrequencyBand().Value),
			})
		}
		result.FddInfo.UlFreqInfo = &topo.FrequencyInfo{}
		arfcn, err := GetFDDUlARFCN(cell)
		if err == nil {
			result.FddInfo.UlFreqInfo.NrArfcn = uint32(arfcn)
		} else {
			errMsg := "arfcn in fdd ul info is nil"
			log.Debug(errMsg)
		}
		result.FddInfo.UlFreqInfo.FrequencyBandList = &topo.FrequencyBandList{
			FrequencyBandItems: freqBandItems,
		}
	}

	if hasFDDDlTransmissionBandwidth(cell) {
		nrb := cell.GetServedCellInfoNr().GetNrModeInfo().GetFdd().GetDlNrtransmissonBandwidth().GetNRnrb() + 1
		scs := cell.GetServedCellInfoNr().GetNrModeInfo().GetFdd().GetDlNrtransmissonBandwidth().GetNRscs() + 1
		result.FddInfo.DlTransmissionBandwidth = &topo.TransmissionBandwidth{
			Nrb:   topo.Nrb(nrb),
			NrScs: topo.NrScs(scs),
		}
	}

	if hasFDDUlTransmissionBandwidth(cell) {
		nrb := cell.GetServedCellInfoNr().GetNrModeInfo().GetFdd().GetUlNrtransmissonBandwidth().GetNRnrb() + 1
		scs := cell.GetServedCellInfoNr().GetNrModeInfo().GetFdd().GetUlNrtransmissonBandwidth().GetNRscs() + 1
		result.FddInfo.UlTransmissionBandwidth = &topo.TransmissionBandwidth{
			Nrb:   topo.Nrb(nrb),
			NrScs: topo.NrScs(scs),
		}
	}
	return result, nil
}

func GetMeasurementTimingConfiguration(cell *xnapiesv1.ServedCellsNRItem) uint64 {
	return *decode.Asn1BytesToUint64(cell.GetServedCellInfoNr().GetMeasurementTimingConfiguration())
}

func hasNeighborInfoNR(cell *xnapiesv1.ServedCellsNRItem) bool {
	return cell != nil && cell.GetNeighbourInfoNr() != nil
}

func GetNeighborInfoNR(cell *xnapiesv1.ServedCellsNRItem) ([]*xnapiesv1.NeighbourInformationNRItem, error) {
	if !hasNeighborInfoNR(cell) {
		errMsg := "neighbor info is nil"
		log.Debug(errMsg)
		return nil, fmt.Errorf(errMsg)
	}
	return cell.GetNeighbourInfoNr().GetValue(), nil
}

func GetTopoNeighborInformationNR(cell *xnapiesv1.NeighbourInformationNRItem) (*topo.NeighborInformationNr, error) {
	result := &topo.NeighborInformationNr{}

	if cell != nil && cell.NrPci != nil {
		result.Pci = uint32(cell.NrPci.Value)
	} else {
		errMsg := "neighbor does not have nr pci"
		log.Debug(errMsg)
	}

	cgi := &topo.NeighborCellID{}
	if cell.NrCgi != nil && cell.NrCgi.PlmnId != nil {
		cgi.PlmnID = fmt.Sprintf("%x", *decode.Asn1BytesToUint64(cell.NrCgi.PlmnId.Value))
	} else {
		errMsg := "neighbor does not have plmn id"
		log.Debug(errMsg)
	}

	if cell.NrCgi != nil && cell.NrCgi.NrCi != nil {
		cgi.CellGlobalID = &topo.CellGlobalID{
			Value: fmt.Sprintf("%x", *decode.Asn1BitstringToUint64(cell.NrCgi.NrCi.Value)),
			Type:  topo.CellGlobalIDType_NRCGI,
		}
	} else {
		errMsg := "neighbor does not have cell id"
		log.Debug(errMsg)
	}
	result.NrCgi = cgi

	if cell.Tac != nil {
		result.Tac = uint32(*decode.Asn1BytesToUint64(cell.Tac.Value))
	} else {
		errMsg := "neighbor does not have tac"
		log.Debug(errMsg)
	}

	if cell.ConnectivitySupport != nil {
		result.ConnectivitySupport = &topo.ConnectivitySupport{
			EnDcSupport: topo.EnDcSupport(cell.ConnectivitySupport.ENdcSupport),
		}
	} else {
		errMsg := "neighbor does not have connectivity support"
		log.Debug(errMsg)
	}

	result.MeasurementTimingConfiguration = uint32(*decode.Asn1BytesToUint64(cell.MeasurementTimingConfiguration))

	if cell.NrModeInfo != nil && cell.NrModeInfo.GetTddInfo() != nil {
		tddInfo := &topo.NeighborInformationNr_TddInfo{
			TddInfo: &topo.TDDInfo{
				NrFreqInfo: &topo.FrequencyInfo{},
			},
		}
		if cell.NrModeInfo.GetTddInfo().GetNrFreqInfo() != nil {
			nrFreqInfo := &topo.FrequencyInfo{}
			if cell.NrModeInfo.GetTddInfo().GetNrFreqInfo().GetNrArfcn() != nil {
				nrFreqInfo.NrArfcn = uint32(cell.NrModeInfo.GetTddInfo().GetNrFreqInfo().GetNrArfcn().Value)
			} else {
				errMsg := "neighbor does not have tdd arfcn"
				log.Debug(errMsg)
			}

			if cell.NrModeInfo.GetTddInfo().GetNrFreqInfo().GetFrequencyBandList() != nil {
				nrFreqInfo.FrequencyBandList = &topo.FrequencyBandList{}
				for _, f := range cell.NrModeInfo.GetTddInfo().GetNrFreqInfo().GetFrequencyBandList().GetValue() {
					if f.GetNrFrequencyBand() == nil {
						errMsg := "neighbor does not have nr frequency band"
						log.Debug(errMsg)
						continue
					}
					nrFreqInfo.FrequencyBandList.FrequencyBandItems = append(nrFreqInfo.FrequencyBandList.FrequencyBandItems, &topo.FrequencyBandItem{
						NrFrequencyBand: uint32(f.GetNrFrequencyBand().GetValue()),
					})
				}
			} else {
				errMsg := "neighbor does not have tdd frequency band list"
				log.Debug(errMsg)
			}
			tddInfo.TddInfo.NrFreqInfo = nrFreqInfo
		}
		result.NrModeInfo = tddInfo
	} else if cell.NrModeInfo != nil && cell.NrModeInfo.GetFddInfo() != nil {
		fddInfo := &topo.NeighborInformationNr_FddInfo{
			FddInfo: &topo.FDDInfo{},
		}
		if cell.NrModeInfo.GetFddInfo().GetUlNrFreqInfo() != nil {
			nrFreqInfo := &topo.FrequencyInfo{}
			if cell.NrModeInfo.GetFddInfo().GetUlNrFreqInfo().GetNrArfcn() != nil {
				nrFreqInfo.NrArfcn = uint32(cell.NrModeInfo.GetFddInfo().GetUlNrFreqInfo().GetNrArfcn().Value)
			} else {
				errMsg := "neighbor does not have fdd ul arfcn"
				log.Debug(errMsg)
			}

			if cell.NrModeInfo.GetFddInfo().GetUlNrFreqInfo().GetFrequencyBandList() != nil {
				nrFreqInfo.FrequencyBandList = &topo.FrequencyBandList{}
				for _, f := range cell.NrModeInfo.GetFddInfo().GetUlNrFreqInfo().GetFrequencyBandList().GetValue() {
					if f.GetNrFrequencyBand() == nil {
						errMsg := "neighbor does not have nr frequency band"
						log.Debug(errMsg)
						continue
					}
					nrFreqInfo.FrequencyBandList.FrequencyBandItems = append(nrFreqInfo.FrequencyBandList.FrequencyBandItems, &topo.FrequencyBandItem{
						NrFrequencyBand: uint32(f.GetNrFrequencyBand().GetValue()),
					})
				}
			} else {
				errMsg := "neighbor does not have fdd ul frequency band list"
				log.Debug(errMsg)
			}
			fddInfo.FddInfo.UlFreqInfo = nrFreqInfo
		}
		if cell.NrModeInfo.GetFddInfo().GetDlNrFequInfo() != nil {
			nrFreqInfo := &topo.FrequencyInfo{}
			if cell.NrModeInfo.GetFddInfo().GetDlNrFequInfo().GetNrArfcn() != nil {
				nrFreqInfo.NrArfcn = uint32(cell.NrModeInfo.GetFddInfo().GetDlNrFequInfo().GetNrArfcn().Value)
			} else {
				errMsg := "neighbor does not have fdd dl arfcn"
				log.Debug(errMsg)
			}

			if cell.NrModeInfo.GetFddInfo().GetDlNrFequInfo().GetFrequencyBandList() != nil {
				nrFreqInfo.FrequencyBandList = &topo.FrequencyBandList{}
				for _, f := range cell.NrModeInfo.GetFddInfo().GetDlNrFequInfo().GetFrequencyBandList().GetValue() {
					if f.GetNrFrequencyBand() == nil {
						errMsg := "neighbor does not have nr frequency band"
						log.Debug(errMsg)
						continue
					}
					nrFreqInfo.FrequencyBandList.FrequencyBandItems = append(nrFreqInfo.FrequencyBandList.FrequencyBandItems, &topo.FrequencyBandItem{
						NrFrequencyBand: uint32(f.GetNrFrequencyBand().GetValue()),
					})
				}
			} else {
				errMsg := "neighbor does not have fdd dl frequency band list"
				log.Debug(errMsg)
			}
			fddInfo.FddInfo.DlFreqInfo = nrFreqInfo
		}
		result.NrModeInfo = fddInfo
	} else {
		errMsg := "neighbor does not have nr mode info"
		log.Debug(errMsg)
	}

	return result, nil
}
