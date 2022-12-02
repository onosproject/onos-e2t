// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdubuilder

import (
	xnapiesv1 "github.com/onosproject/onos-e2t/api/xnap/v1/xnap-ies"

	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"

	xnapcontainersv1 "github.com/onosproject/onos-e2t/api/xnap/v1/xnap-containers"

	xnapconstantsv1 "github.com/onosproject/onos-e2t/api/xnap/v1/xnap-constants"

	xnapcommondatatypesv1 "github.com/onosproject/onos-e2t/api/xnap/v1/xnap-commondatatypes"

	xnappdudescriptionsv1 "github.com/onosproject/onos-e2t/api/xnap/v1/xnap-pdu-descriptions"

	xnappducontentsv1 "github.com/onosproject/onos-e2t/api/xnap/v1/xnap-pdu-contents"
)

func CreateAdditionalPDcpDuplicationTNlList(value []*xnapiesv1.AdditionalPDcpDuplicationTNlItem) (*xnapiesv1.AdditionalPDcpDuplicationTNlList, error) {

	msg := &xnapiesv1.AdditionalPDcpDuplicationTNlList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateAdditionalPDcpDuplicationTNlList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateAdditionalPDcpDuplicationTNlItem(additionalPdcpDuplicationUpTnlInformation *xnapiesv1.UptransportLayerInformation) (*xnapiesv1.AdditionalPDcpDuplicationTNlItem, error) {

	msg := &xnapiesv1.AdditionalPDcpDuplicationTNlItem{}
	msg.AdditionalPdcpDuplicationUpTnlInformation = additionalPdcpDuplicationUpTnlInformation

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateAdditionalPDcpDuplicationTNlItem() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateAdditionalULNGUTNlatUpfItem(additionalUlNgUTnlatUpf *xnapiesv1.UptransportLayerInformation) (*xnapiesv1.AdditionalULNGUTNlatUpfItem, error) {

	msg := &xnapiesv1.AdditionalULNGUTNlatUpfItem{}
	msg.AdditionalUlNgUTnlatUpf = additionalUlNgUTnlatUpf

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateAdditionalULNGUTNlatUpfItem() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateAdditionalULNGUTNlatUpfList(value []*xnapiesv1.AdditionalULNGUTNlatUpfItem) (*xnapiesv1.AdditionalULNGUTNlatUpfList, error) {

	msg := &xnapiesv1.AdditionalULNGUTNlatUpfList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateAdditionalULNGUTNlatUpfList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateActivationIDforCellActivation(value int32) (*xnapiesv1.ActivationIdforCellActivation, error) {

	msg := &xnapiesv1.ActivationIdforCellActivation{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateActivationIDforCellActivation() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateAllocationandRetentionPriority(priorityLevel int32, preEmptionCapability xnapiesv1.PreemptioncapabilityAllocationandRetentionPriority, preEmptionVulnerability xnapiesv1.PreemptionvulnerabilityAllocationandRetentionPriority) (*xnapiesv1.AllocationandRetentionPriority, error) {

	msg := &xnapiesv1.AllocationandRetentionPriority{}
	msg.PriorityLevel = priorityLevel
	msg.PreEmptionCapability = preEmptionCapability
	msg.PreEmptionVulnerability = preEmptionVulnerability

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateAllocationandRetentionPriority() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateActivationSfn(value int32) (*xnapiesv1.ActivationSfn, error) {

	msg := &xnapiesv1.ActivationSfn{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateActivationSfn() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateAllowedCagIDListperPlmn(value []*xnapiesv1.CagIdentifier) (*xnapiesv1.AllowedCagIDListperPlmn, error) {

	msg := &xnapiesv1.AllowedCagIDListperPlmn{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateAllowedCagIDListperPlmn() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateAllowedPniNPnIDList(value []*xnapiesv1.AllowedPniNPnIDItem) (*xnapiesv1.AllowedPniNPnIDList, error) {

	msg := &xnapiesv1.AllowedPniNPnIDList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateAllowedPniNPnIDList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateAllowedPniNPnIDItem(plmnID *xnapiesv1.PlmnIdentity, pniNpnRestrictedInformation xnapiesv1.PniNPnRestrictedInformation, allowedCagIDListPerPlmn *xnapiesv1.AllowedCagIDListperPlmn) (*xnapiesv1.AllowedPniNPnIDItem, error) {

	msg := &xnapiesv1.AllowedPniNPnIDItem{}
	msg.PlmnId = plmnID
	msg.PniNpnRestrictedInformation = pniNpnRestrictedInformation
	msg.AllowedCagIdListPerPlmn = allowedCagIDListPerPlmn

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateAllowedPniNPnIDItem() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateAlternativeQoSparaSetList(value []*xnapiesv1.AlternativeQoSparaSetItem) (*xnapiesv1.AlternativeQoSparaSetList, error) {

	msg := &xnapiesv1.AlternativeQoSparaSetList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateAlternativeQoSparaSetList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateAmfRegionInformation(value []*xnapiesv1.GlobalAmfRegionInformation) (*xnapiesv1.AmfRegionInformation, error) {

	msg := &xnapiesv1.AmfRegionInformation{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateAmfRegionInformation() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateGlobalAmfRegionInformation(plmnID *xnapiesv1.PlmnIdentity, amfRegionID *asn1.BitString) (*xnapiesv1.GlobalAmfRegionInformation, error) {

	msg := &xnapiesv1.GlobalAmfRegionInformation{}
	msg.PlmnId = plmnID
	msg.AmfRegionId = amfRegionID

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateGlobalAmfRegionInformation() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateAmfUENGapID(value int64) (*xnapiesv1.AmfUENGapID, error) {

	msg := &xnapiesv1.AmfUENGapID{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateAmfUENGapID() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateAreaOfInterestInformation(value []*xnapiesv1.AreaOfInterestItem) (*xnapiesv1.AreaOfInterestInformation, error) {

	msg := &xnapiesv1.AreaOfInterestInformation{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateAreaOfInterestInformation() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateAreaScopeOfNeighCellsList(value []*xnapiesv1.AreaScopeOfNeighCellsItem) (*xnapiesv1.AreaScopeOfNeighCellsList, error) {

	msg := &xnapiesv1.AreaScopeOfNeighCellsList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateAreaScopeOfNeighCellsList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateAsSecurityInformation(keyNgRanStar *asn1.BitString, ncc int32) (*xnapiesv1.AsSecurityInformation, error) {

	msg := &xnapiesv1.AsSecurityInformation{}
	msg.KeyNgRanStar = keyNgRanStar
	msg.Ncc = ncc

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateAsSecurityInformation() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateAvailableCapacity(value int32) (*xnapiesv1.AvailableCapacity, error) {

	msg := &xnapiesv1.AvailableCapacity{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateAvailableCapacity() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateAvailableRrcconnectionCapacityValue(value int32) (*xnapiesv1.AvailableRrcconnectionCapacityValue, error) {

	msg := &xnapiesv1.AvailableRrcconnectionCapacityValue{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateAvailableRrcconnectionCapacityValue() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateAveragingWindow(value int32) (*xnapiesv1.AveragingWindow, error) {

	msg := &xnapiesv1.AveragingWindow{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateAveragingWindow() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateBluetoothMeasConfigNameList(value []*xnapiesv1.BluetoothName) (*xnapiesv1.BluetoothMeasConfigNameList, error) {

	msg := &xnapiesv1.BluetoothMeasConfigNameList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateBluetoothMeasConfigNameList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateBluetoothName(value []byte) (*xnapiesv1.BluetoothName, error) {

	msg := &xnapiesv1.BluetoothName{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateBluetoothName() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateBplmnIDInfoEUtra(value []*xnapiesv1.BplmnIDInfoEUtraItem) (*xnapiesv1.BplmnIDInfoEUtra, error) {

	msg := &xnapiesv1.BplmnIDInfoEUtra{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateBplmnIDInfoEUtra() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateBplmnIDInfoNR(value []*xnapiesv1.BplmnIDInfoNRItem) (*xnapiesv1.BplmnIDInfoNR, error) {

	msg := &xnapiesv1.BplmnIDInfoNR{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateBplmnIDInfoNR() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateBplmnIDInfoNRItemExtIes(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, extension *xnapiesv1.BplmnIDInfoNRItemExtIesExtension) (*xnapiesv1.BplmnIDInfoNRItemExtIes, error) {

	msg := &xnapiesv1.BplmnIDInfoNRItemExtIes{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Extension = extension

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateBplmnIDInfoNRItemExtIes() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateBitRate(value int64) (*xnapiesv1.BitRate, error) {

	msg := &xnapiesv1.BitRate{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateBitRate() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateBroadcastCagIDentifierList(value []*xnapiesv1.BroadcastCagIdentifierItem) (*xnapiesv1.BroadcastCagIdentifierList, error) {

	msg := &xnapiesv1.BroadcastCagIdentifierList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateBroadcastCagIDentifierList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateBroadcastCagIDentifierItem(cagIDentifier *xnapiesv1.CagIdentifier) (*xnapiesv1.BroadcastCagIdentifierItem, error) {

	msg := &xnapiesv1.BroadcastCagIdentifierItem{}
	msg.CagIdentifier = cagIDentifier

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateBroadcastCagIDentifierItem() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateBroadcastNIDList(value []*xnapiesv1.BroadcastNidItem) (*xnapiesv1.BroadcastNidList, error) {

	msg := &xnapiesv1.BroadcastNidList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateBroadcastNIDList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateBroadcastNIDItem(nID *xnapiesv1.Nid) (*xnapiesv1.BroadcastNidItem, error) {

	msg := &xnapiesv1.BroadcastNidItem{}
	msg.Nid = nID

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateBroadcastNIDItem() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateBroadcastPlmns(value []*xnapiesv1.PlmnIdentity) (*xnapiesv1.BroadcastPlmns, error) {

	msg := &xnapiesv1.BroadcastPlmns{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateBroadcastPlmns() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateBroadcastEutraplmns(value []*xnapiesv1.PlmnIdentity) (*xnapiesv1.BroadcastEutraplmns, error) {

	msg := &xnapiesv1.BroadcastEutraplmns{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateBroadcastEutraplmns() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateBroadcastPlmninTaisupportItemExtIes(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, extension *xnapiesv1.BroadcastPlmninTaisupportItemExtIesExtension) (*xnapiesv1.BroadcastPlmninTaisupportItemExtIes, error) {

	msg := &xnapiesv1.BroadcastPlmninTaisupportItemExtIes{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Extension = extension

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateBroadcastPlmninTaisupportItemExtIes() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateBroadcastPlmninTaisupportItem(plmnID *xnapiesv1.PlmnIdentity, tAisliceSupportList *xnapiesv1.SliceSupportList) (*xnapiesv1.BroadcastPlmninTaisupportItem, error) {

	msg := &xnapiesv1.BroadcastPlmninTaisupportItem{}
	msg.PlmnId = plmnID
	msg.TAisliceSupportList = tAisliceSupportList

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateBroadcastPlmninTaisupportItem() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateBroadcastPniNPnIDInformation(value []*xnapiesv1.BroadcastPniNPnIDInformationItem) (*xnapiesv1.BroadcastPniNPnIDInformation, error) {

	msg := &xnapiesv1.BroadcastPniNPnIDInformation{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateBroadcastPniNPnIDInformation() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateBroadcastPniNPnIDInformationItem(plmnID *xnapiesv1.PlmnIdentity, broadcastCagIDentifierList *xnapiesv1.BroadcastCagIdentifierList) (*xnapiesv1.BroadcastPniNPnIDInformationItem, error) {

	msg := &xnapiesv1.BroadcastPniNPnIDInformationItem{}
	msg.PlmnId = plmnID
	msg.BroadcastCagIdentifierList = broadcastCagIDentifierList

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateBroadcastPniNPnIDInformationItem() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateBroadcastSnpnIDList(value []*xnapiesv1.BroadcastSnpnid) (*xnapiesv1.BroadcastSnpnidList, error) {

	msg := &xnapiesv1.BroadcastSnpnidList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateBroadcastSnpnIDList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateBroadcastSnpnID(plmnID *xnapiesv1.PlmnIdentity, broadcastNIDList *xnapiesv1.BroadcastNidList) (*xnapiesv1.BroadcastSnpnid, error) {

	msg := &xnapiesv1.BroadcastSnpnid{}
	msg.PlmnId = plmnID
	msg.BroadcastNidList = broadcastNIDList

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateBroadcastSnpnID() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateCagIDentifier(value *asn1.BitString) (*xnapiesv1.CagIdentifier, error) {

	msg := &xnapiesv1.CagIdentifier{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateCagIDentifier() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateCapacityValue(value int32) (*xnapiesv1.CapacityValue, error) {

	msg := &xnapiesv1.CapacityValue{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateCapacityValue() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateLimitedNrListCellAssistanceInfoNR(limitedNrList []*xnapiesv1.NrCGi) (*xnapiesv1.LimitedNrListCellAssistanceInfoNR, error) {

	msg := &xnapiesv1.LimitedNrListCellAssistanceInfoNR{}
	msg.LimitedNrList = limitedNrList

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateLimitedNrListCellAssistanceInfoNR() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateLimitedEutraListCellAssistanceInfoEUtra(limitedEutraList []*xnapiesv1.EUTraCGi) (*xnapiesv1.LimitedEutraListCellAssistanceInfoEUtra, error) {

	msg := &xnapiesv1.LimitedEutraListCellAssistanceInfoEUtra{}
	msg.LimitedEutraList = limitedEutraList

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateLimitedEutraListCellAssistanceInfoEUtra() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateCellBasedMdtNR(cellIDListforMdtNr *xnapiesv1.CellIdListforMdtNR) (*xnapiesv1.CellBasedMdtNR, error) {

	msg := &xnapiesv1.CellBasedMdtNR{}
	msg.CellIdListforMdtNr = cellIDListforMdtNr

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateCellBasedMdtNR() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateCellIDListforMdtNR(value []*xnapiesv1.NrCGi) (*xnapiesv1.CellIdListforMdtNR, error) {

	msg := &xnapiesv1.CellIdListforMdtNR{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateCellIDListforMdtNR() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateCellBasedMdtEUtra(cellIDListforMdtEutra *xnapiesv1.CellIdListforMdtEUtra) (*xnapiesv1.CellBasedMdtEUtra, error) {

	msg := &xnapiesv1.CellBasedMdtEUtra{}
	msg.CellIdListforMdtEutra = cellIDListforMdtEutra

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateCellBasedMdtEUtra() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateCellIDListforMdtEUtra(value []*xnapiesv1.EUTraCGi) (*xnapiesv1.CellIdListforMdtEUtra, error) {

	msg := &xnapiesv1.CellIdListforMdtEUtra{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateCellIDListforMdtEUtra() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateCellCapacityClassValue(value int32) (*xnapiesv1.CellCapacityClassValue, error) {

	msg := &xnapiesv1.CellCapacityClassValue{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateCellCapacityClassValue() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateCellGroupID(value int32) (*xnapiesv1.CellGroupId, error) {

	msg := &xnapiesv1.CellGroupId{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateCellGroupID() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateCellMeasurementResult(value []*xnapiesv1.CellMeasurementResultItem) (*xnapiesv1.CellMeasurementResult, error) {

	msg := &xnapiesv1.CellMeasurementResult{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateCellMeasurementResult() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateCellToReport(value []*xnapiesv1.CellToReportItem) (*xnapiesv1.CellToReport, error) {

	msg := &xnapiesv1.CellToReport{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateCellToReport() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateCompositeAvailableCapacityGroup(compositeAvailableCapacityDownlink *xnapiesv1.CompositeAvailableCapacity, compositeAvailableCapacityUplink *xnapiesv1.CompositeAvailableCapacity) (*xnapiesv1.CompositeAvailableCapacityGroup, error) {

	msg := &xnapiesv1.CompositeAvailableCapacityGroup{}
	msg.CompositeAvailableCapacityDownlink = compositeAvailableCapacityDownlink
	msg.CompositeAvailableCapacityUplink = compositeAvailableCapacityUplink

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateCompositeAvailableCapacityGroup() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateChoProbability(value int32) (*xnapiesv1.ChoProbability, error) {

	msg := &xnapiesv1.ChoProbability{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateChoProbability() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateConnectivitySupport(eNdcSupport xnapiesv1.EndcsupportConnectivitySupport) (*xnapiesv1.ConnectivitySupport, error) {

	msg := &xnapiesv1.ConnectivitySupport{}
	msg.ENdcSupport = eNdcSupport

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateConnectivitySupport() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateCountPDcpSN12(pdcpSn12 int32, hfnPdcpSn12 int32) (*xnapiesv1.CountPDcpSN12, error) {

	msg := &xnapiesv1.CountPDcpSN12{}
	msg.PdcpSn12 = pdcpSn12
	msg.HfnPdcpSn12 = hfnPdcpSn12

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateCountPDcpSN12() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateCountPDcpSN18(pdcpSn18 int32, hfnPdcpSn18 int32) (*xnapiesv1.CountPDcpSN18, error) {

	msg := &xnapiesv1.CountPDcpSN18{}
	msg.PdcpSn18 = pdcpSn18
	msg.HfnPdcpSn18 = hfnPdcpSn18

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateCountPDcpSN18() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateNestedCriticalityDiagnosticsIEList(iEcriticality xnapcommondatatypesv1.Criticality, iEID *xnapcommondatatypesv1.ProtocolIeID, typeOfError xnapiesv1.TypeOfError) (*xnapiesv1.NestedCriticalityDiagnosticsIEList, error) {

	msg := &xnapiesv1.NestedCriticalityDiagnosticsIEList{}
	msg.IEcriticality = iEcriticality
	msg.IEId = iEID
	msg.TypeOfError = typeOfError

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNestedCriticalityDiagnosticsIEList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateCriticalityDiagnosticsIEList(value []*xnapiesv1.NestedCriticalityDiagnosticsIEList) (*xnapiesv1.CriticalityDiagnosticsIEList, error) {

	msg := &xnapiesv1.CriticalityDiagnosticsIEList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateCriticalityDiagnosticsIEList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateCRNti(value *asn1.BitString) (*xnapiesv1.CRNti, error) {

	msg := &xnapiesv1.CRNti{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateCRNti() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateXnUaddressInfoperPdusessionList(value []*xnapiesv1.XnUaddressInfoperPdusessionItem) (*xnapiesv1.XnUaddressInfoperPdusessionList, error) {

	msg := &xnapiesv1.XnUaddressInfoperPdusessionList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateXnUaddressInfoperPdusessionList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateXnUaddressInfoperPdusessionItemExtIes(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, extension *xnapiesv1.XnUaddressInfoperPdusessionItemExtIesExtension) (*xnapiesv1.XnUaddressInfoperPdusessionItemExtIes, error) {

	msg := &xnapiesv1.XnUaddressInfoperPdusessionItemExtIes{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Extension = extension

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateXnUaddressInfoperPdusessionItemExtIes() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateDataForwardingInfoFromTargetEUTrannode(dataForwardingInfoFromTargetEUtrannodeList *xnapiesv1.DataForwardingInfoFromTargetEUTrannodeList) (*xnapiesv1.DataForwardingInfoFromTargetEUTrannode, error) {

	msg := &xnapiesv1.DataForwardingInfoFromTargetEUTrannode{}
	msg.DataForwardingInfoFromTargetEUtrannodeList = dataForwardingInfoFromTargetEUtrannodeList

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateDataForwardingInfoFromTargetEUTrannode() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateDataForwardingInfoFromTargetEUTrannodeList(value []*xnapiesv1.DataForwardingInfoFromTargetEUTrannodeItem) (*xnapiesv1.DataForwardingInfoFromTargetEUTrannodeList, error) {

	msg := &xnapiesv1.DataForwardingInfoFromTargetEUTrannodeList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateDataForwardingInfoFromTargetEUTrannodeList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateDataForwardingInfoFromTargetEUTrannodeItem(dlForwardingUptnlinformation *xnapiesv1.UptransportLayerInformation, qosFlowsToBeForwardedList *xnapiesv1.QoSflowsToBeForwardedList) (*xnapiesv1.DataForwardingInfoFromTargetEUTrannodeItem, error) {

	msg := &xnapiesv1.DataForwardingInfoFromTargetEUTrannodeItem{}
	msg.DlForwardingUptnlinformation = dlForwardingUptnlinformation
	msg.QosFlowsToBeForwardedList = qosFlowsToBeForwardedList

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateDataForwardingInfoFromTargetEUTrannodeItem() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateQoSflowsToBeForwardedList(value []*xnapiesv1.QoSflowsToBeForwardedItem) (*xnapiesv1.QoSflowsToBeForwardedList, error) {

	msg := &xnapiesv1.QoSflowsToBeForwardedList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateQoSflowsToBeForwardedList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateQoSflowsToBeForwardedItem(qosFlowIDentifier *xnapiesv1.QoSflowIdentifier) (*xnapiesv1.QoSflowsToBeForwardedItem, error) {

	msg := &xnapiesv1.QoSflowsToBeForwardedItem{}
	msg.QosFlowIdentifier = qosFlowIDentifier

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateQoSflowsToBeForwardedItem() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateQoSflowsAcceptedToBeForwardedList(value []*xnapiesv1.QoSflowsAcceptedToBeForwardedItem) (*xnapiesv1.QoSflowsAcceptedToBeForwardedList, error) {

	msg := &xnapiesv1.QoSflowsAcceptedToBeForwardedList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateQoSflowsAcceptedToBeForwardedList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateQoSflowsAcceptedToBeForwardedItem(qosFlowIDentifier *xnapiesv1.QoSflowIdentifier) (*xnapiesv1.QoSflowsAcceptedToBeForwardedItem, error) {

	msg := &xnapiesv1.QoSflowsAcceptedToBeForwardedItem{}
	msg.QosFlowIdentifier = qosFlowIDentifier

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateQoSflowsAcceptedToBeForwardedItem() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateQoSflowsToBeForwardedList1(value []*xnapiesv1.QoSflowsToBeForwardedItem1) (*xnapiesv1.QoSflowsToBeForwardedList1, error) {

	msg := &xnapiesv1.QoSflowsToBeForwardedList1{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateQoSflowsToBeForwardedList1() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateQoSflowsToBeForwardedItem1(qosFlowIDentifier *xnapiesv1.QoSflowIdentifier, dlDataforwarding xnapiesv1.Dlforwarding, ulDataforwarding xnapiesv1.Ulforwarding) (*xnapiesv1.QoSflowsToBeForwardedItem1, error) {

	msg := &xnapiesv1.QoSflowsToBeForwardedItem1{}
	msg.QosFlowIdentifier = qosFlowIDentifier
	msg.DlDataforwarding = dlDataforwarding
	msg.UlDataforwarding = ulDataforwarding

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateQoSflowsToBeForwardedItem1() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateDataForwardingResponseDrbitemList(value []*xnapiesv1.DataForwardingResponseDrbitem) (*xnapiesv1.DataForwardingResponseDrbitemList, error) {

	msg := &xnapiesv1.DataForwardingResponseDrbitemList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateDataForwardingResponseDrbitemList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateDataTrafficResources(value *asn1.BitString) (*xnapiesv1.DataTrafficResources, error) {

	msg := &xnapiesv1.DataTrafficResources{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateDataTrafficResources() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateDapsrequestInfo(dapsIndicator xnapiesv1.DapsIndicatorDapsrequestInfo) (*xnapiesv1.DapsrequestInfo, error) {

	msg := &xnapiesv1.DapsrequestInfo{}
	msg.DapsIndicator = dapsIndicator

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateDapsrequestInfo() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateDapsresponseInfoList(value []*xnapiesv1.DapsresponseInfoItem) (*xnapiesv1.DapsresponseInfoList, error) {

	msg := &xnapiesv1.DapsresponseInfoList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateDapsresponseInfoList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateDapsresponseInfoItem(drbID *xnapiesv1.DrbID, dapsResponseIndicator xnapiesv1.DapsResponseIndicatorDapsresponseInfoItem) (*xnapiesv1.DapsresponseInfoItem, error) {

	msg := &xnapiesv1.DapsresponseInfoItem{}
	msg.DrbId = drbID
	msg.DapsResponseIndicator = dapsResponseIndicator

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateDapsresponseInfoItem() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateDeliveryStatus(value int32) (*xnapiesv1.DeliveryStatus, error) {

	msg := &xnapiesv1.DeliveryStatus{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateDeliveryStatus() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateDlGBrPRbusage(value int32) (*xnapiesv1.DlGBrPRbusage, error) {

	msg := &xnapiesv1.DlGBrPRbusage{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateDlGBrPRbusage() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateDlnonGBrPRbusage(value int32) (*xnapiesv1.DlnonGBrPRbusage, error) {

	msg := &xnapiesv1.DlnonGBrPRbusage{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateDlnonGBrPRbusage() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateDlTotalPRbusage(value int32) (*xnapiesv1.DlTotalPRbusage, error) {

	msg := &xnapiesv1.DlTotalPRbusage{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateDlTotalPRbusage() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateDrbID(value int32) (*xnapiesv1.DrbID, error) {

	msg := &xnapiesv1.DrbID{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateDrbID() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateDrbList(value []*xnapiesv1.DrbID) (*xnapiesv1.DrbList, error) {

	msg := &xnapiesv1.DrbList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateDrbList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateDrbListwithCause(value []*xnapiesv1.DrbListwithCauseItem) (*xnapiesv1.DrbListwithCause, error) {

	msg := &xnapiesv1.DrbListwithCause{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateDrbListwithCause() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateDrbNumber(value int32) (*xnapiesv1.DrbNumber, error) {

	msg := &xnapiesv1.DrbNumber{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateDrbNumber() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateDrbsSubjectToDldiscardingList(value []*xnapiesv1.DrbsSubjectToDldiscardingItem) (*xnapiesv1.DrbsSubjectToDldiscardingList, error) {

	msg := &xnapiesv1.DrbsSubjectToDldiscardingList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateDrbsSubjectToDldiscardingList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateDrbsSubjectToDldiscardingItem(drbID *xnapiesv1.DrbID, dlCount *xnapiesv1.DlcountChoice) (*xnapiesv1.DrbsSubjectToDldiscardingItem, error) {

	msg := &xnapiesv1.DrbsSubjectToDldiscardingItem{}
	msg.DrbId = drbID
	msg.DlCount = dlCount

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateDrbsSubjectToDldiscardingItem() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateDrbsSubjectToEarlyStatusTransferList(value []*xnapiesv1.DrbsSubjectToEarlyStatusTransferItem) (*xnapiesv1.DrbsSubjectToEarlyStatusTransferList, error) {

	msg := &xnapiesv1.DrbsSubjectToEarlyStatusTransferList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateDrbsSubjectToEarlyStatusTransferList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateDrbsSubjectToEarlyStatusTransferItem(drbID *xnapiesv1.DrbID, dlCount *xnapiesv1.DlcountChoice) (*xnapiesv1.DrbsSubjectToEarlyStatusTransferItem, error) {

	msg := &xnapiesv1.DrbsSubjectToEarlyStatusTransferItem{}
	msg.DrbId = drbID
	msg.DlCount = dlCount

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateDrbsSubjectToEarlyStatusTransferItem() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateDrbsSubjectToStatusTransferList(value []*xnapiesv1.DrbsSubjectToStatusTransferItem) (*xnapiesv1.DrbsSubjectToStatusTransferList, error) {

	msg := &xnapiesv1.DrbsSubjectToStatusTransferList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateDrbsSubjectToStatusTransferList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateDrbsSubjectToStatusTransferItem(drbID *xnapiesv1.DrbID, pdcpStatusTransferUl *xnapiesv1.DrbbstatusTransferChoice, pdcpStatusTransferDl *xnapiesv1.DrbbstatusTransferChoice) (*xnapiesv1.DrbsSubjectToStatusTransferItem, error) {

	msg := &xnapiesv1.DrbsSubjectToStatusTransferItem{}
	msg.DrbId = drbID
	msg.PdcpStatusTransferUl = pdcpStatusTransferUl
	msg.PdcpStatusTransferDl = pdcpStatusTransferDl

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateDrbsSubjectToStatusTransferItem() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateDrbtoQoSflowMappingList(value []*xnapiesv1.DrbtoQoSflowMappingItem) (*xnapiesv1.DrbtoQoSflowMappingList, error) {

	msg := &xnapiesv1.DrbtoQoSflowMappingList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateDrbtoQoSflowMappingList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateDynamic5QIDescriptorExtIes(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, extension *xnapiesv1.Dynamic5QidescriptorExtIesExtension) (*xnapiesv1.Dynamic5QidescriptorExtIes, error) {

	msg := &xnapiesv1.Dynamic5QidescriptorExtIes{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Extension = extension

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateDynamic5QIDescriptorExtIes() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateERAbID(value int32) (*xnapiesv1.ERAbID, error) {

	msg := &xnapiesv1.ERAbID{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateERAbID() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateEUTraarfcn(value int32) (*xnapiesv1.EUTraarfcn, error) {

	msg := &xnapiesv1.EUTraarfcn{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateEUTraarfcn() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateEUTraCellIdentity(value *asn1.BitString) (*xnapiesv1.EUTraCellIdentity, error) {

	msg := &xnapiesv1.EUTraCellIdentity{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateEUTraCellIdentity() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateEUTraCGi(plmnID *xnapiesv1.PlmnIdentity, eUtraCi *xnapiesv1.EUTraCellIdentity) (*xnapiesv1.EUTraCGi, error) {

	msg := &xnapiesv1.EUTraCGi{}
	msg.PlmnId = plmnID
	msg.EUtraCi = eUtraCi

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateEUTraCGi() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateEUTrafrequencyBandIndicator(value int32) (*xnapiesv1.EUTrafrequencyBandIndicator, error) {

	msg := &xnapiesv1.EUTrafrequencyBandIndicator{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateEUTrafrequencyBandIndicator() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateEUTramultibandInfoList(value []*xnapiesv1.EUTrafrequencyBandIndicator) (*xnapiesv1.EUTramultibandInfoList, error) {

	msg := &xnapiesv1.EUTramultibandInfoList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateEUTramultibandInfoList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateEUTrapci(value int32) (*xnapiesv1.EUTrapci, error) {

	msg := &xnapiesv1.EUTrapci{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateEUTrapci() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateEndpointIpaddressAndPort(endpointIpaddress *xnapiesv1.TransportLayerAddress, portNumber *xnapiesv1.PortNumber) (*xnapiesv1.EndpointIpaddressAndPort, error) {

	msg := &xnapiesv1.EndpointIpaddressAndPort{}
	msg.EndpointIpaddress = endpointIpaddress
	msg.PortNumber = portNumber

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateEndpointIpaddressAndPort() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateEventTriggered(loggedEventTriggeredConfig *xnapiesv1.LoggedEventTriggeredConfig) (*xnapiesv1.EventTriggered, error) {

	msg := &xnapiesv1.EventTriggered{}
	msg.LoggedEventTriggeredConfig = loggedEventTriggeredConfig

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateEventTriggered() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateEventL1(l1Threshold *xnapiesv1.MeasurementThresholdL1LoggedMdt, hysteresis *xnapiesv1.Hysteresis, timeToTrigger xnapiesv1.TimeToTrigger) (*xnapiesv1.EventL1, error) {

	msg := &xnapiesv1.EventL1{}
	msg.L1Threshold = l1Threshold
	msg.Hysteresis = hysteresis
	msg.TimeToTrigger = timeToTrigger

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateEventL1() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateExpectedActivityPeriod(value int32) (*xnapiesv1.ExpectedActivityPeriod, error) {

	msg := &xnapiesv1.ExpectedActivityPeriod{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateExpectedActivityPeriod() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateExpectedIDlePeriod(value int32) (*xnapiesv1.ExpectedIdlePeriod, error) {

	msg := &xnapiesv1.ExpectedIdlePeriod{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateExpectedIDlePeriod() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateExpectedUemovingTrajectory(value []*xnapiesv1.ExpectedUemovingTrajectoryItem) (*xnapiesv1.ExpectedUemovingTrajectory, error) {

	msg := &xnapiesv1.ExpectedUemovingTrajectory{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateExpectedUemovingTrajectory() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateExtendedRatrestrictionInformation(primaryRatrestriction *asn1.BitString, secondaryRatrestriction *asn1.BitString) (*xnapiesv1.ExtendedRatrestrictionInformation, error) {

	msg := &xnapiesv1.ExtendedRatrestrictionInformation{}
	msg.PrimaryRatrestriction = primaryRatrestriction
	msg.SecondaryRatrestriction = secondaryRatrestriction

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateExtendedRatrestrictionInformation() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateExtendedPacketDelayBudget(value int32) (*xnapiesv1.ExtendedPacketDelayBudget, error) {

	msg := &xnapiesv1.ExtendedPacketDelayBudget{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateExtendedPacketDelayBudget() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateExtendedSliceSupportList(value []*xnapiesv1.SNSsai) (*xnapiesv1.ExtendedSliceSupportList, error) {

	msg := &xnapiesv1.ExtendedSliceSupportList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateExtendedSliceSupportList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateExtendedUeidentityIndexValue(value *asn1.BitString) (*xnapiesv1.ExtendedUeidentityIndexValue, error) {

	msg := &xnapiesv1.ExtendedUeidentityIndexValue{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateExtendedUeidentityIndexValue() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateExtTlas(value []*xnapiesv1.ExtTlaItem) (*xnapiesv1.ExtTlas, error) {

	msg := &xnapiesv1.ExtTlas{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateExtTlas() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateGtptlas(value []*xnapiesv1.GtptlaItem) (*xnapiesv1.Gtptlas, error) {

	msg := &xnapiesv1.Gtptlas{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateGtptlas() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateGtptlaItem(gTptransportLayerAddresses *xnapiesv1.TransportLayerAddress) (*xnapiesv1.GtptlaItem, error) {

	msg := &xnapiesv1.GtptlaItem{}
	msg.GTptransportLayerAddresses = gTptransportLayerAddresses

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateGtptlaItem() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateFiveGcmobilityRestrictionListContainer(value []byte) (*xnapiesv1.FiveGcmobilityRestrictionListContainer, error) {

	msg := &xnapiesv1.FiveGcmobilityRestrictionListContainer{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateFiveGcmobilityRestrictionListContainer() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateFiveQi(value int32) (*xnapiesv1.FiveQi, error) {

	msg := &xnapiesv1.FiveQi{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateFiveQi() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateGlobalgNbID(plmnID *xnapiesv1.PlmnIdentity, gnbID *xnapiesv1.GnbIDChoice) (*xnapiesv1.GlobalgNbID, error) {

	msg := &xnapiesv1.GlobalgNbID{}
	msg.PlmnId = plmnID
	msg.GnbId = gnbID

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateGlobalgNbID() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateGnbRadioResourceStatus(ssbAreaRadioResourceStatusList *xnapiesv1.SsbareaRadioResourceStatusList) (*xnapiesv1.GnbRadioResourceStatus, error) {

	msg := &xnapiesv1.GnbRadioResourceStatus{}
	msg.SsbAreaRadioResourceStatusList = ssbAreaRadioResourceStatusList

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateGnbRadioResourceStatus() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateGlobalCellID(plmnID *xnapiesv1.PlmnIdentity, cellType *xnapiesv1.CellTypeChoice) (*xnapiesv1.GlobalCellID, error) {

	msg := &xnapiesv1.GlobalCellID{}
	msg.PlmnId = plmnID
	msg.CellType = cellType

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateGlobalCellID() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateGlobalngeNbID(plmnID *xnapiesv1.PlmnIdentity, enbID *xnapiesv1.EnbIDChoice) (*xnapiesv1.GlobalngeNbID, error) {

	msg := &xnapiesv1.GlobalngeNbID{}
	msg.PlmnId = plmnID
	msg.EnbId = enbID

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateGlobalngeNbID() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateGlobalNgRAncellID(plmnID *xnapiesv1.PlmnIdentity, ngRanCellID *xnapiesv1.NgRAnCellIdentity) (*xnapiesv1.GlobalNgRAncellID, error) {

	msg := &xnapiesv1.GlobalNgRAncellID{}
	msg.PlmnId = plmnID
	msg.NgRanCellId = ngRanCellID

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateGlobalNgRAncellID() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateGtpTEID(value []byte) (*xnapiesv1.GtpTEid, error) {

	msg := &xnapiesv1.GtpTEid{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateGtpTEID() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateGtptunnelTransportLayerInformation(tnlAddress *xnapiesv1.TransportLayerAddress, gtpTeID *xnapiesv1.GtpTEid) (*xnapiesv1.GtptunnelTransportLayerInformation, error) {

	msg := &xnapiesv1.GtptunnelTransportLayerInformation{}
	msg.TnlAddress = tnlAddress
	msg.GtpTeid = gtpTeID

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateGtptunnelTransportLayerInformation() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateGuami(plmnID *xnapiesv1.PlmnIdentity, amfRegionID *asn1.BitString, amfSetID *asn1.BitString, amfPointer *asn1.BitString) (*xnapiesv1.Guami, error) {

	msg := &xnapiesv1.Guami{}
	msg.PlmnId = plmnID
	msg.AmfRegionId = amfRegionID
	msg.AmfSetId = amfSetID
	msg.AmfPointer = amfPointer

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateGuami() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateHysteresis(value int32) (*xnapiesv1.Hysteresis, error) {

	msg := &xnapiesv1.Hysteresis{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateHysteresis() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateImmediateMdtEUtra(value []byte) (*xnapiesv1.ImmediateMdtEUtra, error) {

	msg := &xnapiesv1.ImmediateMdtEUtra{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateImmediateMdtEUtra() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIntendedTddDLULconfigurationNR(nrscs xnapiesv1.Nrscs, nrCyclicPrefix xnapiesv1.NrcyclicPrefix, nrDlUltransmissionPeriodicity xnapiesv1.NrdlULtransmissionPeriodicity, slotConfigurationList *xnapiesv1.SlotConfigurationList) (*xnapiesv1.IntendedTddDLULconfigurationNR, error) {

	msg := &xnapiesv1.IntendedTddDLULconfigurationNR{}
	msg.Nrscs = nrscs
	msg.NrCyclicPrefix = nrCyclicPrefix
	msg.NrDlUltransmissionPeriodicity = nrDlUltransmissionPeriodicity
	msg.SlotConfigurationList = slotConfigurationList

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIntendedTddDLULconfigurationNR() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateInterfaceInstanceIndication(value int32) (*xnapiesv1.InterfaceInstanceIndication, error) {

	msg := &xnapiesv1.InterfaceInstanceIndication{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateInterfaceInstanceIndication() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateInterfacesToTrace(value *asn1.BitString) (*xnapiesv1.InterfacesToTrace, error) {

	msg := &xnapiesv1.InterfacesToTrace{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateInterfacesToTrace() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateLastVisitedEutrancellInformation(value []byte) (*xnapiesv1.LastVisitedEutrancellInformation, error) {

	msg := &xnapiesv1.LastVisitedEutrancellInformation{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateLastVisitedEutrancellInformation() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateLastVisitedGerancellInformation(value []byte) (*xnapiesv1.LastVisitedGerancellInformation, error) {

	msg := &xnapiesv1.LastVisitedGerancellInformation{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateLastVisitedGerancellInformation() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateLastVisitedNgrancellInformation(value []byte) (*xnapiesv1.LastVisitedNgrancellInformation, error) {

	msg := &xnapiesv1.LastVisitedNgrancellInformation{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateLastVisitedNgrancellInformation() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateLastVisitedUtrancellInformation(value []byte) (*xnapiesv1.LastVisitedUtrancellInformation, error) {

	msg := &xnapiesv1.LastVisitedUtrancellInformation{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateLastVisitedUtrancellInformation() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateLcID(value int32) (*xnapiesv1.Lcid, error) {

	msg := &xnapiesv1.Lcid{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateLcID() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateListOfCells(value []*xnapiesv1.CellsinAoIItem) (*xnapiesv1.ListOfCells, error) {

	msg := &xnapiesv1.ListOfCells{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateListOfCells() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateCellsinAoIItem(pLmnIdentity *xnapiesv1.PlmnIdentity, ngRanCellID *xnapiesv1.NgRAnCellIdentity) (*xnapiesv1.CellsinAoIItem, error) {

	msg := &xnapiesv1.CellsinAoIItem{}
	msg.PLmnIdentity = pLmnIdentity
	msg.NgRanCellId = ngRanCellID

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateCellsinAoIItem() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateListOfRannodesinAoI(value []*xnapiesv1.GlobalNgRAnnodesinAoIItem) (*xnapiesv1.ListOfRannodesinAoI, error) {

	msg := &xnapiesv1.ListOfRannodesinAoI{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateListOfRannodesinAoI() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateGlobalNgRAnnodesinAoIItem(globalNgRanNodeID *xnapiesv1.GlobalNgRAnnodeID) (*xnapiesv1.GlobalNgRAnnodesinAoIItem, error) {

	msg := &xnapiesv1.GlobalNgRAnnodesinAoIItem{}
	msg.GlobalNgRanNodeId = globalNgRanNodeID

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateGlobalNgRAnnodesinAoIItem() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateListOfTaisinAoI(value []*xnapiesv1.TaisinAoIItem) (*xnapiesv1.ListOfTaisinAoI, error) {

	msg := &xnapiesv1.ListOfTaisinAoI{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateListOfTaisinAoI() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateTaisinAoIItem(pLmnIdentity *xnapiesv1.PlmnIdentity, tAc *xnapiesv1.Tac) (*xnapiesv1.TaisinAoIItem, error) {

	msg := &xnapiesv1.TaisinAoIItem{}
	msg.PLmnIdentity = pLmnIdentity
	msg.TAc = tAc

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateTaisinAoIItem() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateLoggedEventTriggeredConfig(eventTypeTrigger *xnapiesv1.EventTypeTrigger) (*xnapiesv1.LoggedEventTriggeredConfig, error) {

	msg := &xnapiesv1.LoggedEventTriggeredConfig{}
	msg.EventTypeTrigger = eventTypeTrigger

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateLoggedEventTriggeredConfig() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateLteuesIDelinkAggregateMaximumBitRate(uEsIDelinkAggregateMaximumBitRate *xnapiesv1.BitRate) (*xnapiesv1.LteuesidelinkAggregateMaximumBitRate, error) {

	msg := &xnapiesv1.LteuesidelinkAggregateMaximumBitRate{}
	msg.UEsidelinkAggregateMaximumBitRate = uEsIDelinkAggregateMaximumBitRate

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateLteuesIDelinkAggregateMaximumBitRate() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateM1PeriodicReporting(reportInterval xnapiesv1.ReportIntervalMdt, reportAmount xnapiesv1.ReportAmountMdt) (*xnapiesv1.M1PeriodicReporting, error) {

	msg := &xnapiesv1.M1PeriodicReporting{}
	msg.ReportInterval = reportInterval
	msg.ReportAmount = reportAmount

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateM1PeriodicReporting() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateM1ThresholdEventA2(measurementThreshold *xnapiesv1.MeasurementThresholdA2) (*xnapiesv1.M1ThresholdEventA2, error) {

	msg := &xnapiesv1.M1ThresholdEventA2{}
	msg.MeasurementThreshold = measurementThreshold

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateM1ThresholdEventA2() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateM4Configuration(m4Period xnapiesv1.M4Period, m4LinksToLog xnapiesv1.Linkstolog) (*xnapiesv1.M4Configuration, error) {

	msg := &xnapiesv1.M4Configuration{}
	msg.M4Period = m4Period
	msg.M4LinksToLog = m4LinksToLog

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateM4Configuration() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateM5Configuration(m5Period xnapiesv1.M5Period, m5LinksToLog xnapiesv1.Linkstolog) (*xnapiesv1.M5Configuration, error) {

	msg := &xnapiesv1.M5Configuration{}
	msg.M5Period = m5Period
	msg.M5LinksToLog = m5LinksToLog

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateM5Configuration() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateM6Configuration(m6ReportInterval xnapiesv1.M6ReportInterval, m6LinksToLog xnapiesv1.Linkstolog) (*xnapiesv1.M6Configuration, error) {

	msg := &xnapiesv1.M6Configuration{}
	msg.M6ReportInterval = m6ReportInterval
	msg.M6LinksToLog = m6LinksToLog

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateM6Configuration() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateM7Configuration(m7Period *xnapiesv1.M7Period, m7LinksToLog xnapiesv1.Linkstolog) (*xnapiesv1.M7Configuration, error) {

	msg := &xnapiesv1.M7Configuration{}
	msg.M7Period = m7Period
	msg.M7LinksToLog = m7LinksToLog

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateM7Configuration() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateM7period(value int32) (*xnapiesv1.M7Period, error) {

	msg := &xnapiesv1.M7Period{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateM7period() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateMacI(value *asn1.BitString) (*xnapiesv1.MacI, error) {

	msg := &xnapiesv1.MacI{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateMacI() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateMaskedImeisv(value *asn1.BitString) (*xnapiesv1.MaskedImeisv, error) {

	msg := &xnapiesv1.MaskedImeisv{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateMaskedImeisv() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateMaxChopreparations(value int32) (*xnapiesv1.MaxChopreparations, error) {

	msg := &xnapiesv1.MaxChopreparations{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateMaxChopreparations() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateMaximumDataBurstVolume(value int32) (*xnapiesv1.MaximumDataBurstVolume, error) {

	msg := &xnapiesv1.MaximumDataBurstVolume{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateMaximumDataBurstVolume() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateMaximumIpdatarate(maxIprateUl xnapiesv1.MaxIprate) (*xnapiesv1.MaximumIpdatarate, error) {

	msg := &xnapiesv1.MaximumIpdatarate{}
	msg.MaxIprateUl = maxIprateUl

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateMaximumIpdatarate() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateMbsfncontrolRegionLength(value int32) (*xnapiesv1.MbsfncontrolRegionLength, error) {

	msg := &xnapiesv1.MbsfncontrolRegionLength{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateMbsfncontrolRegionLength() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateMbsfnsubframeInfoEUTra(value []*xnapiesv1.MbsfnsubframeInfoEUTraItem) (*xnapiesv1.MbsfnsubframeInfoEUTra, error) {

	msg := &xnapiesv1.MbsfnsubframeInfoEUTra{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateMbsfnsubframeInfoEUTra() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateMbsfnsubframeInfoEUTraItem(radioframeAllocationPeriod xnapiesv1.RadioframeAllocationPeriodMbsfnsubframeInfoEutraitem, radioframeAllocationOffset int32, subframeAllocation *xnapiesv1.MbsfnsubframeAllocationEUTra) (*xnapiesv1.MbsfnsubframeInfoEUTraItem, error) {

	msg := &xnapiesv1.MbsfnsubframeInfoEUTraItem{}
	msg.RadioframeAllocationPeriod = radioframeAllocationPeriod
	msg.RadioframeAllocationOffset = radioframeAllocationOffset
	msg.SubframeAllocation = subframeAllocation

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateMbsfnsubframeInfoEUTraItem() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateMdtLocationInfo(value *asn1.BitString) (*xnapiesv1.MdtLocationInfo, error) {

	msg := &xnapiesv1.MdtLocationInfo{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateMdtLocationInfo() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateMdtplmnlist(value []*xnapiesv1.PlmnIdentity) (*xnapiesv1.Mdtplmnlist, error) {

	msg := &xnapiesv1.Mdtplmnlist{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateMdtplmnlist() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateMdtmodeNRExtension(value *xnapiesv1.MdtmodeNRExtensionIe) (*xnapiesv1.MdtmodeNRExtension, error) {

	msg := &xnapiesv1.MdtmodeNRExtension{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateMdtmodeNRExtension() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateMdtmodeEUtraExtension(value *xnapiesv1.MdtmodeEUtraExtensionIe) (*xnapiesv1.MdtmodeEUtraExtension, error) {

	msg := &xnapiesv1.MdtmodeEUtraExtension{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateMdtmodeEUtraExtension() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateMeasurementsToActivate(value *asn1.BitString) (*xnapiesv1.MeasurementsToActivate, error) {

	msg := &xnapiesv1.MeasurementsToActivate{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateMeasurementsToActivate() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateMeasurementID(value int32) (*xnapiesv1.MeasurementID, error) {

	msg := &xnapiesv1.MeasurementID{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateMeasurementID() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateMobilityInformation(value *asn1.BitString) (*xnapiesv1.MobilityInformation, error) {

	msg := &xnapiesv1.MobilityInformation{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateMobilityInformation() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateMobilityParametersModificationRange(handoverTriggerChangeLowerLimit int32, handoverTriggerChangeUpperLimit int32) (*xnapiesv1.MobilityParametersModificationRange, error) {

	msg := &xnapiesv1.MobilityParametersModificationRange{}
	msg.HandoverTriggerChangeLowerLimit = handoverTriggerChangeLowerLimit
	msg.HandoverTriggerChangeUpperLimit = handoverTriggerChangeUpperLimit

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateMobilityParametersModificationRange() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateMobilityParametersInformation(handoverTriggerChange int32) (*xnapiesv1.MobilityParametersInformation, error) {

	msg := &xnapiesv1.MobilityParametersInformation{}
	msg.HandoverTriggerChange = handoverTriggerChange

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateMobilityParametersInformation() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateMobilityRestrictionListExtIes(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, extension *xnapiesv1.MobilityRestrictionListExtIesExtension) (*xnapiesv1.MobilityRestrictionListExtIes, error) {

	msg := &xnapiesv1.MobilityRestrictionListExtIes{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Extension = extension

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateMobilityRestrictionListExtIes() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateCntypeRestrictionsForEquivalent(value []*xnapiesv1.CntypeRestrictionsForEquivalentItem) (*xnapiesv1.CntypeRestrictionsForEquivalent, error) {

	msg := &xnapiesv1.CntypeRestrictionsForEquivalent{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateCntypeRestrictionsForEquivalent() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateCntypeRestrictionsForEquivalentItem(plmnIdentity *xnapiesv1.PlmnIdentity, cnType xnapiesv1.CnTypeCntypeRestrictionsForEquivalentItem) (*xnapiesv1.CntypeRestrictionsForEquivalentItem, error) {

	msg := &xnapiesv1.CntypeRestrictionsForEquivalentItem{}
	msg.PlmnIdentity = plmnIdentity
	msg.CnType = cnType

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateCntypeRestrictionsForEquivalentItem() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateRatRestrictionsList(value []*xnapiesv1.RatRestrictionsItem) (*xnapiesv1.RatRestrictionsList, error) {

	msg := &xnapiesv1.RatRestrictionsList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRatRestrictionsList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateRatRestrictionsItem(plmnIdentity *xnapiesv1.PlmnIdentity, ratRestrictionInformation *xnapiesv1.RatRestrictionInformation) (*xnapiesv1.RatRestrictionsItem, error) {

	msg := &xnapiesv1.RatRestrictionsItem{}
	msg.PlmnIdentity = plmnIdentity
	msg.RatRestrictionInformation = ratRestrictionInformation

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRatRestrictionsItem() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateRatRestrictionInformation(value *asn1.BitString) (*xnapiesv1.RatRestrictionInformation, error) {

	msg := &xnapiesv1.RatRestrictionInformation{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRatRestrictionInformation() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateForbIDdenAreaList(value []*xnapiesv1.ForbiddenAreaItem) (*xnapiesv1.ForbiddenAreaList, error) {

	msg := &xnapiesv1.ForbiddenAreaList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateForbIDdenAreaList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateForbIDdenAreaItem(plmnIdentity *xnapiesv1.PlmnIdentity, forbIDdenTacs []*xnapiesv1.Tac) (*xnapiesv1.ForbiddenAreaItem, error) {

	msg := &xnapiesv1.ForbiddenAreaItem{}
	msg.PlmnIdentity = plmnIdentity
	msg.ForbiddenTacs = forbIDdenTacs

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateForbIDdenAreaItem() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateServiceAreaList(value []*xnapiesv1.ServiceAreaItem) (*xnapiesv1.ServiceAreaList, error) {

	msg := &xnapiesv1.ServiceAreaList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateServiceAreaList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateServiceAreaItem(plmnIdentity *xnapiesv1.PlmnIdentity) (*xnapiesv1.ServiceAreaItem, error) {

	msg := &xnapiesv1.ServiceAreaItem{}
	msg.PlmnIdentity = plmnIdentity

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateServiceAreaItem() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateMrDCResourceCoordinationInfo(ngRanNodeResourceCoordinationInfo *xnapiesv1.NgRAnNodeResourceCoordinationInfo) (*xnapiesv1.MrDCResourceCoordinationInfo, error) {

	msg := &xnapiesv1.MrDCResourceCoordinationInfo{}
	msg.NgRanNodeResourceCoordinationInfo = ngRanNodeResourceCoordinationInfo

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateMrDCResourceCoordinationInfo() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateMessageOversizeNotification(maximumCellListSize *xnapiesv1.MaximumCellListSize) (*xnapiesv1.MessageOversizeNotification, error) {

	msg := &xnapiesv1.MessageOversizeNotification{}
	msg.MaximumCellListSize = maximumCellListSize

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateMessageOversizeNotification() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateMaximumCellListSize(value int32) (*xnapiesv1.MaximumCellListSize, error) {

	msg := &xnapiesv1.MaximumCellListSize{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateMaximumCellListSize() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateNeDCTDmPattern(subframeAssignment xnapiesv1.SubframeAssignmentNedctdmpattern, harqOffset int32) (*xnapiesv1.NeDCTDmPattern, error) {

	msg := &xnapiesv1.NeDCTDmPattern{}
	msg.SubframeAssignment = subframeAssignment
	msg.HarqOffset = harqOffset

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNeDCTDmPattern() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateNeighbourInformationEUTra(value []*xnapiesv1.NeighbourInformationEUTraItem) (*xnapiesv1.NeighbourInformationEUTra, error) {

	msg := &xnapiesv1.NeighbourInformationEUTra{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNeighbourInformationEUTra() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateNeighbourInformationNR(value []*xnapiesv1.NeighbourInformationNRItem) (*xnapiesv1.NeighbourInformationNR, error) {

	msg := &xnapiesv1.NeighbourInformationNR{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNeighbourInformationNR() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateNeighbourInformationNRModeFddinfo(ulNrFreqInfo *xnapiesv1.NrfrequencyInfo, dlNrFequInfo *xnapiesv1.NrfrequencyInfo) (*xnapiesv1.NeighbourInformationNRModeFddinfo, error) {

	msg := &xnapiesv1.NeighbourInformationNRModeFddinfo{}
	msg.UlNrFreqInfo = ulNrFreqInfo
	msg.DlNrFequInfo = dlNrFequInfo

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNeighbourInformationNRModeFddinfo() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateNeighbourInformationNRModeTddinfo(nrFreqInfo *xnapiesv1.NrfrequencyInfo) (*xnapiesv1.NeighbourInformationNRModeTddinfo, error) {

	msg := &xnapiesv1.NeighbourInformationNRModeTddinfo{}
	msg.NrFreqInfo = nrFreqInfo

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNeighbourInformationNRModeTddinfo() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateNID(value *asn1.BitString) (*xnapiesv1.Nid, error) {

	msg := &xnapiesv1.Nid{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNID() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateNrcarrierList(value []*xnapiesv1.NrcarrierItem) (*xnapiesv1.NrcarrierList, error) {

	msg := &xnapiesv1.NrcarrierList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNrcarrierList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateNrcarrierItem(carrierScs xnapiesv1.Nrscs, offsetToCarrier int32, carrierBandwIDth int32) (*xnapiesv1.NrcarrierItem, error) {

	msg := &xnapiesv1.NrcarrierItem{}
	msg.CarrierScs = carrierScs
	msg.OffsetToCarrier = offsetToCarrier
	msg.CarrierBandwidth = carrierBandwIDth

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNrcarrierItem() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateNrcellPrachconfig(value []byte) (*xnapiesv1.NrcellPrachconfig, error) {

	msg := &xnapiesv1.NrcellPrachconfig{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNrcellPrachconfig() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateNgRAnnodeUexnApID(value int64) (*xnapiesv1.NgRAnnodeUexnApid, error) {

	msg := &xnapiesv1.NgRAnnodeUexnApid{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNgRAnnodeUexnApID() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateNumberofActiveUes(value int32) (*xnapiesv1.NumberofActiveUes, error) {

	msg := &xnapiesv1.NumberofActiveUes{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNumberofActiveUes() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateNoofRrcconnections(value int32) (*xnapiesv1.NoofRrcconnections, error) {

	msg := &xnapiesv1.NoofRrcconnections{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNoofRrcconnections() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateNonDynamic5QIDescriptorExtIes(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, extension *xnapiesv1.NonDynamic5QidescriptorExtIesExtension) (*xnapiesv1.NonDynamic5QidescriptorExtIes, error) {

	msg := &xnapiesv1.NonDynamic5QidescriptorExtIes{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Extension = extension

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNonDynamic5QIDescriptorExtIes() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateNrarfcn(value int32) (*xnapiesv1.Nrarfcn, error) {

	msg := &xnapiesv1.Nrarfcn{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNrarfcn() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateNgeNbRadioResourceStatusExtIes(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, extension *xnapiesv1.NgeNbRadioResourceStatusExtIesExtension) (*xnapiesv1.NgeNbRadioResourceStatusExtIes, error) {

	msg := &xnapiesv1.NgeNbRadioResourceStatusExtIes{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Extension = extension

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNgeNbRadioResourceStatusExtIes() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateNgeNbRadioResourceStatus(dLGbrPrbUsage *xnapiesv1.DlGBrPRbusage, uLGbrPrbUsage *xnapiesv1.UlGBrPRbusage, dLNonGbrPrbUsage *xnapiesv1.DlnonGBrPRbusage, uLNonGbrPrbUsage *xnapiesv1.UlnonGBrPRbusage, dLTotalPrbUsage *xnapiesv1.DlTotalPRbusage, uLTotalPrbUsage *xnapiesv1.UlTotalPRbusage) (*xnapiesv1.NgeNbRadioResourceStatus, error) {

	msg := &xnapiesv1.NgeNbRadioResourceStatus{}
	msg.DLGbrPrbUsage = dLGbrPrbUsage
	msg.ULGbrPrbUsage = uLGbrPrbUsage
	msg.DLNonGbrPrbUsage = dLNonGbrPrbUsage
	msg.ULNonGbrPrbUsage = uLNonGbrPrbUsage
	msg.DLTotalPrbUsage = dLTotalPrbUsage
	msg.ULTotalPrbUsage = uLTotalPrbUsage

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNgeNbRadioResourceStatus() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateDlschedulingPDcchCCeusage(value int32) (*xnapiesv1.DlschedulingPDcchCCeusage, error) {

	msg := &xnapiesv1.DlschedulingPDcchCCeusage{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateDlschedulingPDcchCCeusage() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateUlschedulingPDcchCCeusage(value int32) (*xnapiesv1.UlschedulingPDcchCCeusage, error) {

	msg := &xnapiesv1.UlschedulingPDcchCCeusage{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateUlschedulingPDcchCCeusage() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateTnlcapacityIndicator(dLtnlofferedCapacity *xnapiesv1.OfferedCapacity, dLtnlavailableCapacity *xnapiesv1.AvailableCapacity, uLtnlofferedCapacity *xnapiesv1.OfferedCapacity, uLtnlavailableCapacity *xnapiesv1.AvailableCapacity) (*xnapiesv1.TnlcapacityIndicator, error) {

	msg := &xnapiesv1.TnlcapacityIndicator{}
	msg.DLtnlofferedCapacity = dLtnlofferedCapacity
	msg.DLtnlavailableCapacity = dLtnlavailableCapacity
	msg.ULtnlofferedCapacity = uLtnlofferedCapacity
	msg.ULtnlavailableCapacity = uLtnlavailableCapacity

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateTnlcapacityIndicator() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateNpnBroadcastInformationSNpn(broadcastSnpnIDList *xnapiesv1.BroadcastSnpnidList) (*xnapiesv1.NpnBroadcastInformationSNpn, error) {

	msg := &xnapiesv1.NpnBroadcastInformationSNpn{}
	msg.BroadcastSnpnidList = broadcastSnpnIDList

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNpnBroadcastInformationSNpn() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateNpnBroadcastInformationPNiNPn(broadcastPniNpnIDInformation *xnapiesv1.BroadcastPniNPnIDInformation) (*xnapiesv1.NpnBroadcastInformationPNiNPn, error) {

	msg := &xnapiesv1.NpnBroadcastInformationPNiNPn{}
	msg.BroadcastPniNpnIdInformation = broadcastPniNpnIDInformation

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNpnBroadcastInformationPNiNPn() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateNpnmobilityInformationSNpn(servingNID *xnapiesv1.Nid) (*xnapiesv1.NpnmobilityInformationSNpn, error) {

	msg := &xnapiesv1.NpnmobilityInformationSNpn{}
	msg.ServingNid = servingNID

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNpnmobilityInformationSNpn() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateNpnmobilityInformationPNiNPn(allowedPniNpnIDList *xnapiesv1.AllowedPniNPnIDList) (*xnapiesv1.NpnmobilityInformationPNiNPn, error) {

	msg := &xnapiesv1.NpnmobilityInformationPNiNPn{}
	msg.AllowedPniNpnIdList = allowedPniNpnIDList

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNpnmobilityInformationPNiNPn() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateNpnpagingAssistanceInformationPNiNPn(allowedPniNpnIDList *xnapiesv1.AllowedPniNPnIDList) (*xnapiesv1.NpnpagingAssistanceInformationPNiNPn, error) {

	msg := &xnapiesv1.NpnpagingAssistanceInformationPNiNPn{}
	msg.AllowedPniNpnIdList = allowedPniNpnIDList

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNpnpagingAssistanceInformationPNiNPn() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateNpnSupportSNpn(nID *xnapiesv1.Nid) (*xnapiesv1.NpnSupportSNpn, error) {

	msg := &xnapiesv1.NpnSupportSNpn{}
	msg.Nid = nID

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNpnSupportSNpn() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateNprachconfiguration(fddOrTdd *xnapiesv1.FddortddNprachconfiguration) (*xnapiesv1.Nprachconfiguration, error) {

	msg := &xnapiesv1.Nprachconfiguration{}
	msg.FddOrTdd = fddOrTdd

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNprachconfiguration() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateNestedNonAnchorCarrierFrequencylist(nonAnchorCarrierFrquency []byte) (*xnapiesv1.NestedNonAnchorCarrierFrequencylist, error) {

	msg := &xnapiesv1.NestedNonAnchorCarrierFrequencylist{}
	msg.NonAnchorCarrierFrquency = nonAnchorCarrierFrquency

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNestedNonAnchorCarrierFrequencylist() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateNonAnchorCarrierFrequencylist(value []*xnapiesv1.NestedNonAnchorCarrierFrequencylist) (*xnapiesv1.NonAnchorCarrierFrequencylist, error) {

	msg := &xnapiesv1.NonAnchorCarrierFrequencylist{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNonAnchorCarrierFrequencylist() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateNrCellIdentity(value *asn1.BitString) (*xnapiesv1.NrCellIdentity, error) {

	msg := &xnapiesv1.NrCellIdentity{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNrCellIdentity() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateNgRAnCellIdentityListinRanpagingArea(value []*xnapiesv1.NgRAnCellIdentity) (*xnapiesv1.NgRAnCellIdentityListinRanpagingArea, error) {

	msg := &xnapiesv1.NgRAnCellIdentityListinRanpagingArea{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNgRAnCellIdentityListinRanpagingArea() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateNrCGi(plmnID *xnapiesv1.PlmnIdentity, nrCi *xnapiesv1.NrCellIdentity) (*xnapiesv1.NrCGi, error) {

	msg := &xnapiesv1.NrCGi{}
	msg.PlmnId = plmnID
	msg.NrCi = nrCi

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNrCGi() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateNrfrequencyBand(value int32) (*xnapiesv1.NrfrequencyBand, error) {

	msg := &xnapiesv1.NrfrequencyBand{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNrfrequencyBand() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateNrfrequencyBandList(value []*xnapiesv1.NrfrequencyBandItem) (*xnapiesv1.NrfrequencyBandList, error) {

	msg := &xnapiesv1.NrfrequencyBandList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNrfrequencyBandList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateNrmobilityHistoryReport(value []byte) (*xnapiesv1.NrmobilityHistoryReport, error) {

	msg := &xnapiesv1.NrmobilityHistoryReport{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNrmobilityHistoryReport() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateNrmodeInfoFddExtIes(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, extension *xnapiesv1.NrmodeInfoFddExtIesExtension) (*xnapiesv1.NrmodeInfoFddExtIes, error) {

	msg := &xnapiesv1.NrmodeInfoFddExtIes{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Extension = extension

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNrmodeInfoFddExtIes() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateNrmodeInfoFdd(ulNrfrequencyInfo *xnapiesv1.NrfrequencyInfo, dlNrfrequencyInfo *xnapiesv1.NrfrequencyInfo, ulNrtransmissonBandwIDth *xnapiesv1.NrtransmissionBandwidth, dlNrtransmissonBandwIDth *xnapiesv1.NrtransmissionBandwidth) (*xnapiesv1.NrmodeInfoFdd, error) {

	msg := &xnapiesv1.NrmodeInfoFdd{}
	msg.UlNrfrequencyInfo = ulNrfrequencyInfo
	msg.DlNrfrequencyInfo = dlNrfrequencyInfo
	msg.UlNrtransmissonBandwidth = ulNrtransmissonBandwIDth
	msg.DlNrtransmissonBandwidth = dlNrtransmissonBandwIDth

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNrmodeInfoFdd() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateNrmodeInfoTddExtIes(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, extension *xnapiesv1.NrmodeInfoTddExtIesExtension) (*xnapiesv1.NrmodeInfoTddExtIes, error) {

	msg := &xnapiesv1.NrmodeInfoTddExtIes{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Extension = extension

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNrmodeInfoTddExtIes() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateNrmodeInfoTdd(nrFrequencyInfo *xnapiesv1.NrfrequencyInfo, nrTransmissonBandwIDth *xnapiesv1.NrtransmissionBandwidth) (*xnapiesv1.NrmodeInfoTdd, error) {

	msg := &xnapiesv1.NrmodeInfoTdd{}
	msg.NrFrequencyInfo = nrFrequencyInfo
	msg.NrTransmissonBandwidth = nrTransmissonBandwIDth

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNrmodeInfoTdd() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateNrpci(value int32) (*xnapiesv1.Nrpci, error) {

	msg := &xnapiesv1.Nrpci{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNrpci() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateNrtransmissionBandwIDth(nRscs xnapiesv1.Nrscs, nRnrb xnapiesv1.Nrnrb) (*xnapiesv1.NrtransmissionBandwidth, error) {

	msg := &xnapiesv1.NrtransmissionBandwidth{}
	msg.NRscs = nRscs
	msg.NRnrb = nRnrb

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNrtransmissionBandwIDth() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateNgRAntraceID(value []byte) (*xnapiesv1.NgRAntraceId, error) {

	msg := &xnapiesv1.NgRAntraceId{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNgRAntraceID() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateNruesIDelinkAggregateMaximumBitRate(uEsIDelinkAggregateMaximumBitRate *xnapiesv1.BitRate) (*xnapiesv1.NruesidelinkAggregateMaximumBitRate, error) {

	msg := &xnapiesv1.NruesidelinkAggregateMaximumBitRate{}
	msg.UEsidelinkAggregateMaximumBitRate = uEsIDelinkAggregateMaximumBitRate

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNruesIDelinkAggregateMaximumBitRate() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateOfferedCapacity(value int32) (*xnapiesv1.OfferedCapacity, error) {

	msg := &xnapiesv1.OfferedCapacity{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateOfferedCapacity() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreatePacketDelayBudget(value int32) (*xnapiesv1.PacketDelayBudget, error) {

	msg := &xnapiesv1.PacketDelayBudget{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePacketDelayBudget() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreatePacketErrorRate(pErScalar *xnapiesv1.PerScalar, pErExponent *xnapiesv1.PerExponent) (*xnapiesv1.PacketErrorRate, error) {

	msg := &xnapiesv1.PacketErrorRate{}
	msg.PErScalar = pErScalar
	msg.PErExponent = pErExponent

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePacketErrorRate() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreatePerScalar(value int32) (*xnapiesv1.PerScalar, error) {

	msg := &xnapiesv1.PerScalar{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePerScalar() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreatePerExponent(value int32) (*xnapiesv1.PerExponent, error) {

	msg := &xnapiesv1.PerExponent{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePerExponent() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreatePacketLossRate(value int32) (*xnapiesv1.PacketLossRate, error) {

	msg := &xnapiesv1.PacketLossRate{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePacketLossRate() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreatePc5QoSflowList(value []*xnapiesv1.Pc5QoSflowItem) (*xnapiesv1.Pc5QoSflowList, error) {

	msg := &xnapiesv1.Pc5QoSflowList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePc5QoSflowList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreatePc5FlowBitRates(guaranteedFlowBitRate *xnapiesv1.BitRate, maximumFlowBitRate *xnapiesv1.BitRate) (*xnapiesv1.Pc5FlowBitRates, error) {

	msg := &xnapiesv1.Pc5FlowBitRates{}
	msg.GuaranteedFlowBitRate = guaranteedFlowBitRate
	msg.MaximumFlowBitRate = maximumFlowBitRate

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePc5FlowBitRates() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreatePdcpsnlength(ulPdcpsnlength xnapiesv1.UlPdcpsnlengthPdcpsnlength, dlPdcpsnlength xnapiesv1.DlPdcpsnlengthPdcpsnlength) (*xnapiesv1.Pdcpsnlength, error) {

	msg := &xnapiesv1.Pdcpsnlength{}
	msg.UlPdcpsnlength = ulPdcpsnlength
	msg.DlPdcpsnlength = dlPdcpsnlength

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePdcpsnlength() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreatePdusessionAggregateMaximumBitRate(downlinkSessionAmbr *xnapiesv1.BitRate, uplinkSessionAmbr *xnapiesv1.BitRate) (*xnapiesv1.PdusessionAggregateMaximumBitRate, error) {

	msg := &xnapiesv1.PdusessionAggregateMaximumBitRate{}
	msg.DownlinkSessionAmbr = downlinkSessionAmbr
	msg.UplinkSessionAmbr = uplinkSessionAmbr

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePdusessionAggregateMaximumBitRate() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreatePdusessionList(value []*xnapiesv1.PdusessionID) (*xnapiesv1.PdusessionList, error) {

	msg := &xnapiesv1.PdusessionList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePdusessionList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreatePdusessionListwithCause(value []*xnapiesv1.PdusessionListwithCauseItem) (*xnapiesv1.PdusessionListwithCause, error) {

	msg := &xnapiesv1.PdusessionListwithCause{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePdusessionListwithCause() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreatePdusessionListwithDataForwardingFromTarget(value []*xnapiesv1.PdusessionListwithDataForwardingFromTargetItem) (*xnapiesv1.PdusessionListwithDataForwardingFromTarget, error) {

	msg := &xnapiesv1.PdusessionListwithDataForwardingFromTarget{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePdusessionListwithDataForwardingFromTarget() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreatePdusessionListwithDataForwardingFromTargetItem(pduSessionID *xnapiesv1.PdusessionID, dataforwardinginfoTarget *xnapiesv1.DataForwardingInfoFromTargetNgrannode) (*xnapiesv1.PdusessionListwithDataForwardingFromTargetItem, error) {

	msg := &xnapiesv1.PdusessionListwithDataForwardingFromTargetItem{}
	msg.PduSessionId = pduSessionID
	msg.DataforwardinginfoTarget = dataforwardinginfoTarget

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePdusessionListwithDataForwardingFromTargetItem() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreatePdusessionListwithDataForwardingRequest(value []*xnapiesv1.PdusessionListwithDataForwardingRequestItem) (*xnapiesv1.PdusessionListwithDataForwardingRequest, error) {

	msg := &xnapiesv1.PdusessionListwithDataForwardingRequest{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePdusessionListwithDataForwardingRequest() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreatePdusessionResourcesAdmittedList(value []*xnapiesv1.PdusessionResourcesAdmittedItem) (*xnapiesv1.PdusessionResourcesAdmittedList, error) {

	msg := &xnapiesv1.PdusessionResourcesAdmittedList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePdusessionResourcesAdmittedList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreatePdusessionResourcesAdmittedItem(pduSessionID *xnapiesv1.PdusessionID, pduSessionResourceAdmittedInfo *xnapiesv1.PdusessionResourceAdmittedInfo) (*xnapiesv1.PdusessionResourcesAdmittedItem, error) {

	msg := &xnapiesv1.PdusessionResourcesAdmittedItem{}
	msg.PduSessionId = pduSessionID
	msg.PduSessionResourceAdmittedInfo = pduSessionResourceAdmittedInfo

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePdusessionResourcesAdmittedItem() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreatePdusessionResourcesNotAdmittedList(value []*xnapiesv1.PdusessionResourcesNotAdmittedItem) (*xnapiesv1.PdusessionResourcesNotAdmittedList, error) {

	msg := &xnapiesv1.PdusessionResourcesNotAdmittedList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePdusessionResourcesNotAdmittedList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreatePdusessionResourcesToBeSetupList(value []*xnapiesv1.PdusessionResourcesToBeSetupItem) (*xnapiesv1.PdusessionResourcesToBeSetupList, error) {

	msg := &xnapiesv1.PdusessionResourcesToBeSetupList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePdusessionResourcesToBeSetupList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreatePdusessionResourcesToBeSetupItemExtIes(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, extension *xnapiesv1.PdusessionResourcesToBeSetupItemExtIesExtension) (*xnapiesv1.PdusessionResourcesToBeSetupItemExtIes, error) {

	msg := &xnapiesv1.PdusessionResourcesToBeSetupItemExtIes{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Extension = extension

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePdusessionResourcesToBeSetupItemExtIes() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreatePdusessionResourceSetupInfoSNterminatedExtIes(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, extension *xnapiesv1.PdusessionResourceSetupInfoSNterminatedExtIesExtension) (*xnapiesv1.PdusessionResourceSetupInfoSNterminatedExtIes, error) {

	msg := &xnapiesv1.PdusessionResourceSetupInfoSNterminatedExtIes{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Extension = extension

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePdusessionResourceSetupInfoSNterminatedExtIes() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateQoSflowsToBeSetupListSetupSNterminated(value []*xnapiesv1.QoSflowsToBeSetupListSetupSNterminatedItem) (*xnapiesv1.QoSflowsToBeSetupListSetupSNterminated, error) {

	msg := &xnapiesv1.QoSflowsToBeSetupListSetupSNterminated{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateQoSflowsToBeSetupListSetupSNterminated() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateQoSflowsToBeSetupListSetupSNterminatedItemExtIes(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, extension *xnapiesv1.QoSflowsToBeSetupListSetupSNterminatedItemExtIesExtension) (*xnapiesv1.QoSflowsToBeSetupListSetupSNterminatedItemExtIes, error) {

	msg := &xnapiesv1.QoSflowsToBeSetupListSetupSNterminatedItemExtIes{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Extension = extension

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateQoSflowsToBeSetupListSetupSNterminatedItemExtIes() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreatePdusessionResourceSetupResponseInfoSNterminatedExtIes(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, extension *xnapiesv1.PdusessionResourceSetupResponseInfoSNterminatedExtIesExtension) (*xnapiesv1.PdusessionResourceSetupResponseInfoSNterminatedExtIes, error) {

	msg := &xnapiesv1.PdusessionResourceSetupResponseInfoSNterminatedExtIes{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Extension = extension

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePdusessionResourceSetupResponseInfoSNterminatedExtIes() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateDrbsToBeSetupListSetupResponseSNterminated(value []*xnapiesv1.DrbsToBeSetupListSetupResponseSNterminatedItem) (*xnapiesv1.DrbsToBeSetupListSetupResponseSNterminated, error) {

	msg := &xnapiesv1.DrbsToBeSetupListSetupResponseSNterminated{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateDrbsToBeSetupListSetupResponseSNterminated() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateDrbsToBeSetupListSetupResponseSNterminatedItemExtIes(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, extension *xnapiesv1.DrbsToBeSetupListSetupResponseSNterminatedItemExtIesExtension) (*xnapiesv1.DrbsToBeSetupListSetupResponseSNterminatedItemExtIes, error) {

	msg := &xnapiesv1.DrbsToBeSetupListSetupResponseSNterminatedItemExtIes{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Extension = extension

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateDrbsToBeSetupListSetupResponseSNterminatedItemExtIes() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateQoSflowsMappedtoDrbSetupResponseSNterminated(value []*xnapiesv1.QoSflowsMappedtoDrbSetupResponseSNterminatedItem) (*xnapiesv1.QoSflowsMappedtoDrbSetupResponseSNterminated, error) {

	msg := &xnapiesv1.QoSflowsMappedtoDrbSetupResponseSNterminated{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateQoSflowsMappedtoDrbSetupResponseSNterminated() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateQoSflowsMappedtoDrbSetupResponseSNterminatedItemExtIes(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, extension *xnapiesv1.QoSflowsMappedtoDrbSetupResponseSNterminatedItemExtIesExtension) (*xnapiesv1.QoSflowsMappedtoDrbSetupResponseSNterminatedItemExtIes, error) {

	msg := &xnapiesv1.QoSflowsMappedtoDrbSetupResponseSNterminatedItemExtIes{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Extension = extension

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateQoSflowsMappedtoDrbSetupResponseSNterminatedItemExtIes() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreatePdusessionResourceSetupInfoMNterminated(pduSessionType xnapiesv1.PdusessionType, dRbsToBeSetup *xnapiesv1.DrbsToBeSetupListSetupMNterminated) (*xnapiesv1.PdusessionResourceSetupInfoMNterminated, error) {

	msg := &xnapiesv1.PdusessionResourceSetupInfoMNterminated{}
	msg.PduSessionType = pduSessionType
	msg.DRbsToBeSetup = dRbsToBeSetup

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePdusessionResourceSetupInfoMNterminated() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateDrbsToBeSetupListSetupMNterminated(value []*xnapiesv1.DrbsToBeSetupListSetupMNterminatedItem) (*xnapiesv1.DrbsToBeSetupListSetupMNterminated, error) {

	msg := &xnapiesv1.DrbsToBeSetupListSetupMNterminated{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateDrbsToBeSetupListSetupMNterminated() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateDrbsToBeSetupListSetupMNterminatedItemExtIes(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, extension *xnapiesv1.DrbsToBeSetupListSetupMNterminatedItemExtIesExtension) (*xnapiesv1.DrbsToBeSetupListSetupMNterminatedItemExtIes, error) {

	msg := &xnapiesv1.DrbsToBeSetupListSetupMNterminatedItemExtIes{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Extension = extension

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateDrbsToBeSetupListSetupMNterminatedItemExtIes() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateQoSflowsMappedtoDrbSetupMNterminated(value []*xnapiesv1.QoSflowsMappedtoDrbSetupMNterminatedItem) (*xnapiesv1.QoSflowsMappedtoDrbSetupMNterminated, error) {

	msg := &xnapiesv1.QoSflowsMappedtoDrbSetupMNterminated{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateQoSflowsMappedtoDrbSetupMNterminated() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreatePdusessionResourceSetupResponseInfoMNterminated(dRbsAdmittedList *xnapiesv1.DrbsAdmittedListSetupResponseMNterminated) (*xnapiesv1.PdusessionResourceSetupResponseInfoMNterminated, error) {

	msg := &xnapiesv1.PdusessionResourceSetupResponseInfoMNterminated{}
	msg.DRbsAdmittedList = dRbsAdmittedList

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePdusessionResourceSetupResponseInfoMNterminated() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateDrbsAdmittedListSetupResponseMNterminated(value []*xnapiesv1.DrbsAdmittedListSetupResponseMNterminatedItem) (*xnapiesv1.DrbsAdmittedListSetupResponseMNterminated, error) {

	msg := &xnapiesv1.DrbsAdmittedListSetupResponseMNterminated{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateDrbsAdmittedListSetupResponseMNterminated() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateDrbsAdmittedListSetupResponseMNterminatedItemExtIes(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, extension *xnapiesv1.DrbsAdmittedListSetupResponseMNterminatedItemExtIesExtension) (*xnapiesv1.DrbsAdmittedListSetupResponseMNterminatedItemExtIes, error) {

	msg := &xnapiesv1.DrbsAdmittedListSetupResponseMNterminatedItemExtIes{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Extension = extension

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateDrbsAdmittedListSetupResponseMNterminatedItemExtIes() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateQoSflowsMappedtoDrbSetupResponseMNterminated(value []*xnapiesv1.QoSflowsMappedtoDrbSetupResponseMNterminatedItem) (*xnapiesv1.QoSflowsMappedtoDrbSetupResponseMNterminated, error) {

	msg := &xnapiesv1.QoSflowsMappedtoDrbSetupResponseMNterminated{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateQoSflowsMappedtoDrbSetupResponseMNterminated() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateQoSflowsMappedtoDrbSetupResponseMNterminatedItem(qoSflowIDentifier *xnapiesv1.QoSflowIdentifier, currentQoSparaSetIndex *xnapiesv1.QoSparaSetIndex) (*xnapiesv1.QoSflowsMappedtoDrbSetupResponseMNterminatedItem, error) {

	msg := &xnapiesv1.QoSflowsMappedtoDrbSetupResponseMNterminatedItem{}
	msg.QoSflowIdentifier = qoSflowIDentifier
	msg.CurrentQoSparaSetIndex = currentQoSparaSetIndex

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateQoSflowsMappedtoDrbSetupResponseMNterminatedItem() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreatePdusessionResourceModificationInfoSNterminatedExtIes(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, extension *xnapiesv1.PdusessionResourceModificationInfoSNterminatedExtIesExtension) (*xnapiesv1.PdusessionResourceModificationInfoSNterminatedExtIes, error) {

	msg := &xnapiesv1.PdusessionResourceModificationInfoSNterminatedExtIes{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Extension = extension

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePdusessionResourceModificationInfoSNterminatedExtIes() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateQoSflowsToBeSetupListModifiedSNterminated(value []*xnapiesv1.QoSflowsToBeSetupListModifiedSNterminatedItem) (*xnapiesv1.QoSflowsToBeSetupListModifiedSNterminated, error) {

	msg := &xnapiesv1.QoSflowsToBeSetupListModifiedSNterminated{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateQoSflowsToBeSetupListModifiedSNterminated() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateQoSflowsToBeSetupListModifiedSNterminatedItemExtIes(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, extension *xnapiesv1.QoSflowsToBeSetupListModifiedSNterminatedItemExtIesExtension) (*xnapiesv1.QoSflowsToBeSetupListModifiedSNterminatedItemExtIes, error) {

	msg := &xnapiesv1.QoSflowsToBeSetupListModifiedSNterminatedItemExtIes{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Extension = extension

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateQoSflowsToBeSetupListModifiedSNterminatedItemExtIes() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateDrbsToBeModifiedListModifiedSNterminated(value []*xnapiesv1.DrbsToBeModifiedListModifiedSNterminatedItem) (*xnapiesv1.DrbsToBeModifiedListModifiedSNterminated, error) {

	msg := &xnapiesv1.DrbsToBeModifiedListModifiedSNterminated{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateDrbsToBeModifiedListModifiedSNterminated() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreatePdusessionResourceModificationResponseInfoSNterminatedExtIes(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, extension *xnapiesv1.PdusessionResourceModificationResponseInfoSNterminatedExtIesExtension) (*xnapiesv1.PdusessionResourceModificationResponseInfoSNterminatedExtIes, error) {

	msg := &xnapiesv1.PdusessionResourceModificationResponseInfoSNterminatedExtIes{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Extension = extension

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePdusessionResourceModificationResponseInfoSNterminatedExtIes() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateDrbsToBeModifiedListModificationResponseSNterminated(value []*xnapiesv1.DrbsToBeModifiedListModificationResponseSNterminatedItem) (*xnapiesv1.DrbsToBeModifiedListModificationResponseSNterminated, error) {

	msg := &xnapiesv1.DrbsToBeModifiedListModificationResponseSNterminated{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateDrbsToBeModifiedListModificationResponseSNterminated() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateDrbsToBeModifiedListModificationResponseSNterminatedItemExtIes(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, extension *xnapiesv1.DrbsToBeModifiedListModificationResponseSNterminatedItemExtIesExtension) (*xnapiesv1.DrbsToBeModifiedListModificationResponseSNterminatedItemExtIes, error) {

	msg := &xnapiesv1.DrbsToBeModifiedListModificationResponseSNterminatedItemExtIes{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Extension = extension

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateDrbsToBeModifiedListModificationResponseSNterminatedItemExtIes() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateDrbsToBeModifiedListModificationMNterminated(value []*xnapiesv1.DrbsToBeModifiedListModificationMNterminatedItem) (*xnapiesv1.DrbsToBeModifiedListModificationMNterminated, error) {

	msg := &xnapiesv1.DrbsToBeModifiedListModificationMNterminated{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateDrbsToBeModifiedListModificationMNterminated() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateDrbsToBeModifiedListModificationMNterminatedItemExtIes(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, extension *xnapiesv1.DrbsToBeModifiedListModificationMNterminatedItemExtIesExtension) (*xnapiesv1.DrbsToBeModifiedListModificationMNterminatedItemExtIes, error) {

	msg := &xnapiesv1.DrbsToBeModifiedListModificationMNterminatedItemExtIes{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Extension = extension

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateDrbsToBeModifiedListModificationMNterminatedItemExtIes() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateDrbsAdmittedListModificationResponseMNterminated(value []*xnapiesv1.DrbsAdmittedListModificationResponseMNterminatedItem) (*xnapiesv1.DrbsAdmittedListModificationResponseMNterminated, error) {

	msg := &xnapiesv1.DrbsAdmittedListModificationResponseMNterminated{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateDrbsAdmittedListModificationResponseMNterminated() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateDrbsAdmittedListModificationResponseMNterminatedItemExtIes(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, extension *xnapiesv1.DrbsAdmittedListModificationResponseMNterminatedItemExtIesExtension) (*xnapiesv1.DrbsAdmittedListModificationResponseMNterminatedItemExtIes, error) {

	msg := &xnapiesv1.DrbsAdmittedListModificationResponseMNterminatedItemExtIes{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Extension = extension

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateDrbsAdmittedListModificationResponseMNterminatedItemExtIes() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreatePdusessionResourceChangeRequiredInfoMNterminated() (*xnapiesv1.PdusessionResourceChangeRequiredInfoMNterminated, error) {

	msg := &xnapiesv1.PdusessionResourceChangeRequiredInfoMNterminated{}

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePdusessionResourceChangeRequiredInfoMNterminated() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreatePdusessionResourceChangeConfirmInfoMNterminated() (*xnapiesv1.PdusessionResourceChangeConfirmInfoMNterminated, error) {

	msg := &xnapiesv1.PdusessionResourceChangeConfirmInfoMNterminated{}

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePdusessionResourceChangeConfirmInfoMNterminated() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateDrbsToBeSetupListModRqdSNterminated(value []*xnapiesv1.DrbsToBeSetupListModRqdSNterminatedItem) (*xnapiesv1.DrbsToBeSetupListModRqdSNterminated, error) {

	msg := &xnapiesv1.DrbsToBeSetupListModRqdSNterminated{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateDrbsToBeSetupListModRqdSNterminated() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateDrbsToBeSetupListModRqdSNterminatedItemExtIes(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, extension *xnapiesv1.DrbsToBeSetupListModRqdSNterminatedItemExtIesExtension) (*xnapiesv1.DrbsToBeSetupListModRqdSNterminatedItemExtIes, error) {

	msg := &xnapiesv1.DrbsToBeSetupListModRqdSNterminatedItemExtIes{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Extension = extension

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateDrbsToBeSetupListModRqdSNterminatedItemExtIes() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateQoSflowsSetupMappedtoDrbModRqdSNterminated(value []*xnapiesv1.QoSflowsSetupMappedtoDrbModRqdSNterminatedItem) (*xnapiesv1.QoSflowsSetupMappedtoDrbModRqdSNterminated, error) {

	msg := &xnapiesv1.QoSflowsSetupMappedtoDrbModRqdSNterminated{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateQoSflowsSetupMappedtoDrbModRqdSNterminated() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateDrbsToBeModifiedListModRqdSNterminated(value []*xnapiesv1.DrbsToBeModifiedListModRqdSNterminatedItem) (*xnapiesv1.DrbsToBeModifiedListModRqdSNterminated, error) {

	msg := &xnapiesv1.DrbsToBeModifiedListModRqdSNterminated{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateDrbsToBeModifiedListModRqdSNterminated() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateDrbsToBeModifiedListModRqdSNterminatedItemExtIes(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, extension *xnapiesv1.DrbsToBeModifiedListModRqdSNterminatedItemExtIesExtension) (*xnapiesv1.DrbsToBeModifiedListModRqdSNterminatedItemExtIes, error) {

	msg := &xnapiesv1.DrbsToBeModifiedListModRqdSNterminatedItemExtIes{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Extension = extension

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateDrbsToBeModifiedListModRqdSNterminatedItemExtIes() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateQoSflowsModifiedMappedtoDrbModRqdSNterminated(value []*xnapiesv1.QoSflowsModifiedMappedtoDrbModRqdSNterminatedItem) (*xnapiesv1.QoSflowsModifiedMappedtoDrbModRqdSNterminated, error) {

	msg := &xnapiesv1.QoSflowsModifiedMappedtoDrbModRqdSNterminated{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateQoSflowsModifiedMappedtoDrbModRqdSNterminated() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateDrbsAdmittedListModConfirmSNterminated(value []*xnapiesv1.DrbsAdmittedListModConfirmSNterminatedItem) (*xnapiesv1.DrbsAdmittedListModConfirmSNterminated, error) {

	msg := &xnapiesv1.DrbsAdmittedListModConfirmSNterminated{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateDrbsAdmittedListModConfirmSNterminated() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateDrbsToBeModifiedListModRqdMNterminated(value []*xnapiesv1.DrbsToBeModifiedListModRqdMNterminatedItem) (*xnapiesv1.DrbsToBeModifiedListModRqdMNterminated, error) {

	msg := &xnapiesv1.DrbsToBeModifiedListModRqdMNterminated{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateDrbsToBeModifiedListModRqdMNterminated() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreatePdusessionResourceModConfirmInfoMNterminated() (*xnapiesv1.PdusessionResourceModConfirmInfoMNterminated, error) {

	msg := &xnapiesv1.PdusessionResourceModConfirmInfoMNterminated{}

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePdusessionResourceModConfirmInfoMNterminated() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreatePdusessionResourceBearerSetupCompleteInfoSNterminated(dRbsToBeSetupList []*xnapiesv1.DrbsToBeSetupListBearerSetupCompleteSNterminatedItem) (*xnapiesv1.PdusessionResourceBearerSetupCompleteInfoSNterminated, error) {

	msg := &xnapiesv1.PdusessionResourceBearerSetupCompleteInfoSNterminated{}
	msg.DRbsToBeSetupList = dRbsToBeSetupList

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePdusessionResourceBearerSetupCompleteInfoSNterminated() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateDrbsToBeSetupListBearerSetupCompleteSNterminatedItem(dRbID *xnapiesv1.DrbID, mNXnUTnlinfoatM *xnapiesv1.UptransportLayerInformation) (*xnapiesv1.DrbsToBeSetupListBearerSetupCompleteSNterminatedItem, error) {

	msg := &xnapiesv1.DrbsToBeSetupListBearerSetupCompleteSNterminatedItem{}
	msg.DRbId = dRbID
	msg.MNXnUTnlinfoatM = mNXnUTnlinfoatM

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateDrbsToBeSetupListBearerSetupCompleteSNterminatedItem() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreatePdusessionResourceSecondaryRatusageList(value []*xnapiesv1.PdusessionResourceSecondaryRatusageItem) (*xnapiesv1.PdusessionResourceSecondaryRatusageList, error) {

	msg := &xnapiesv1.PdusessionResourceSecondaryRatusageList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePdusessionResourceSecondaryRatusageList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreatePdusessionResourceSecondaryRatusageItem(pDusessionID *xnapiesv1.PdusessionID, secondaryRatusageInformation *xnapiesv1.SecondaryRatusageInformation) (*xnapiesv1.PdusessionResourceSecondaryRatusageItem, error) {

	msg := &xnapiesv1.PdusessionResourceSecondaryRatusageItem{}
	msg.PDusessionId = pDusessionID
	msg.SecondaryRatusageInformation = secondaryRatusageInformation

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePdusessionResourceSecondaryRatusageItem() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreatePdusessionUsageReport(rAttype xnapiesv1.RattypePdusessionUsageReport, pDusessionTimedReportList *xnapiesv1.VolumeTimedReportList) (*xnapiesv1.PdusessionUsageReport, error) {

	msg := &xnapiesv1.PdusessionUsageReport{}
	msg.RAttype = rAttype
	msg.PDusessionTimedReportList = pDusessionTimedReportList

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePdusessionUsageReport() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreatePdusessionID(value int32) (*xnapiesv1.PdusessionID, error) {

	msg := &xnapiesv1.PdusessionID{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePdusessionID() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreatePdusessionNetworkInstance(value int32) (*xnapiesv1.PdusessionNetworkInstance, error) {

	msg := &xnapiesv1.PdusessionNetworkInstance{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePdusessionNetworkInstance() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreatePdusessionCommonNetworkInstance(value []byte) (*xnapiesv1.PdusessionCommonNetworkInstance, error) {

	msg := &xnapiesv1.PdusessionCommonNetworkInstance{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePdusessionCommonNetworkInstance() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreatePeriodical() (*xnapiesv1.Periodical, error) {

	msg := &xnapiesv1.Periodical{}

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePeriodical() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreatePlmnIdentity(value []byte) (*xnapiesv1.PlmnIdentity, error) {

	msg := &xnapiesv1.PlmnIdentity{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePlmnIdentity() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreatePcilistForMdt(value []*xnapiesv1.Nrpci) (*xnapiesv1.PcilistForMdt, error) {

	msg := &xnapiesv1.PcilistForMdt{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePcilistForMdt() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreatePortNumber(value *asn1.BitString) (*xnapiesv1.PortNumber, error) {

	msg := &xnapiesv1.PortNumber{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePortNumber() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreatePriorityLevelQoS(value int32) (*xnapiesv1.PriorityLevelQoS, error) {

	msg := &xnapiesv1.PriorityLevelQoS{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePriorityLevelQoS() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateProtectedEUTraresourceList(value []*xnapiesv1.ProtectedEUTraresourceItem) (*xnapiesv1.ProtectedEUTraresourceList, error) {

	msg := &xnapiesv1.ProtectedEUTraresourceList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateProtectedEUTraresourceList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateProtectedEUTraresourceItem(resourceType xnapiesv1.ResourceTypeProtectedEutraresourceItem, intraPrbprotectedResourceFootprint *asn1.BitString, protectedFootprintFrequencyPattern *asn1.BitString, protectedFootprintTimePattern *xnapiesv1.ProtectedEUTrafootprintTimePattern) (*xnapiesv1.ProtectedEUTraresourceItem, error) {

	msg := &xnapiesv1.ProtectedEUTraresourceItem{}
	msg.ResourceType = resourceType
	msg.IntraPrbprotectedResourceFootprint = intraPrbprotectedResourceFootprint
	msg.ProtectedFootprintFrequencyPattern = protectedFootprintFrequencyPattern
	msg.ProtectedFootprintTimePattern = protectedFootprintTimePattern

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateProtectedEUTraresourceItem() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateProtectedEUTrafootprintTimePattern(protectedFootprintTimeperiodicity int32, protectedFootrpintStartTime int32) (*xnapiesv1.ProtectedEUTrafootprintTimePattern, error) {

	msg := &xnapiesv1.ProtectedEUTrafootprintTimePattern{}
	msg.ProtectedFootprintTimeperiodicity = protectedFootprintTimeperiodicity
	msg.ProtectedFootrpintStartTime = protectedFootrpintStartTime

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateProtectedEUTrafootprintTimePattern() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateQoSflowIDentifier(value int32) (*xnapiesv1.QoSflowIdentifier, error) {

	msg := &xnapiesv1.QoSflowIdentifier{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateQoSflowIDentifier() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateQoSflowLevelQoSparametersExtIes(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, extension *xnapiesv1.QoSflowLevelQoSparametersExtIesExtension) (*xnapiesv1.QoSflowLevelQoSparametersExtIes, error) {

	msg := &xnapiesv1.QoSflowLevelQoSparametersExtIes{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Extension = extension

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateQoSflowLevelQoSparametersExtIes() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateQoSflowNotificationControlIndicationInfo(value []*xnapiesv1.QoSflowNotifyItem) (*xnapiesv1.QoSflowNotificationControlIndicationInfo, error) {

	msg := &xnapiesv1.QoSflowNotificationControlIndicationInfo{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateQoSflowNotificationControlIndicationInfo() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateQoSflowNotifyItem(qosFlowIDentifier *xnapiesv1.QoSflowIdentifier, notificationInformation xnapiesv1.NotificationInformationQoSflowNotifyItem) (*xnapiesv1.QoSflowNotifyItem, error) {

	msg := &xnapiesv1.QoSflowNotifyItem{}
	msg.QosFlowIdentifier = qosFlowIDentifier
	msg.NotificationInformation = notificationInformation

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateQoSflowNotifyItem() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateQoSflowsList(value []*xnapiesv1.QoSflowItem) (*xnapiesv1.QoSflowsList, error) {

	msg := &xnapiesv1.QoSflowsList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateQoSflowsList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateQoSflowsListwithCause(value []*xnapiesv1.QoSflowwithCauseItem) (*xnapiesv1.QoSflowsListwithCause, error) {

	msg := &xnapiesv1.QoSflowsListwithCause{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateQoSflowsListwithCause() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateQoSparaSetIndex(value int32) (*xnapiesv1.QoSparaSetIndex, error) {

	msg := &xnapiesv1.QoSparaSetIndex{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateQoSparaSetIndex() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateQoSparaSetNotifyIndex(value int32) (*xnapiesv1.QoSparaSetNotifyIndex, error) {

	msg := &xnapiesv1.QoSparaSetNotifyIndex{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateQoSparaSetNotifyIndex() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateQoSflowsAdmittedList(value []*xnapiesv1.QoSflowsAdmittedItem) (*xnapiesv1.QoSflowsAdmittedList, error) {

	msg := &xnapiesv1.QoSflowsAdmittedList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateQoSflowsAdmittedList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateQoSflowsAdmittedItem(qfi *xnapiesv1.QoSflowIdentifier) (*xnapiesv1.QoSflowsAdmittedItem, error) {

	msg := &xnapiesv1.QoSflowsAdmittedItem{}
	msg.Qfi = qfi

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateQoSflowsAdmittedItem() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateQoSflowsToBeSetupList(value []*xnapiesv1.QoSflowsToBeSetupItem) (*xnapiesv1.QoSflowsToBeSetupList, error) {

	msg := &xnapiesv1.QoSflowsToBeSetupList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateQoSflowsToBeSetupList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateQoSflowsToBeSetupItemExtIes(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, extension *xnapiesv1.QoSflowsToBeSetupItemExtIesExtension) (*xnapiesv1.QoSflowsToBeSetupItemExtIes, error) {

	msg := &xnapiesv1.QoSflowsToBeSetupItemExtIes{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Extension = extension

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateQoSflowsToBeSetupItemExtIes() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateQoSflowsUsageReportList(value []*xnapiesv1.QoSflowsUsageReportItem) (*xnapiesv1.QoSflowsUsageReportList, error) {

	msg := &xnapiesv1.QoSflowsUsageReportList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateQoSflowsUsageReportList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateQoSflowsUsageReportItem(qosFlowIDentifier *xnapiesv1.QoSflowIdentifier, rAttype xnapiesv1.RattypeQoSflowsUsageReportItem, qoSflowsTimedReportList *xnapiesv1.VolumeTimedReportList) (*xnapiesv1.QoSflowsUsageReportItem, error) {

	msg := &xnapiesv1.QoSflowsUsageReportItem{}
	msg.QosFlowIdentifier = qosFlowIDentifier
	msg.RAttype = rAttype
	msg.QoSflowsTimedReportList = qoSflowsTimedReportList

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateQoSflowsUsageReportItem() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateQosMonitoringReportingFrequency(value int32) (*xnapiesv1.QosMonitoringReportingFrequency, error) {

	msg := &xnapiesv1.QosMonitoringReportingFrequency{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateQosMonitoringReportingFrequency() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateRachreportInformation(value []*xnapiesv1.RachreportListItem) (*xnapiesv1.RachreportInformation, error) {

	msg := &xnapiesv1.RachreportInformation{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRachreportInformation() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateRachreportListItem(rAchreport *xnapiesv1.RachreportContainer) (*xnapiesv1.RachreportListItem, error) {

	msg := &xnapiesv1.RachreportListItem{}
	msg.RAchreport = rAchreport

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRachreportListItem() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateRachreportContainer(value []byte) (*xnapiesv1.RachreportContainer, error) {

	msg := &xnapiesv1.RachreportContainer{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRachreportContainer() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateRanac(value int32) (*xnapiesv1.Ranac, error) {

	msg := &xnapiesv1.Ranac{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRanac() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateRanareaIDList(value []*xnapiesv1.RanareaId) (*xnapiesv1.RanareaIdList, error) {

	msg := &xnapiesv1.RanareaIdList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRanareaIDList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateRanpagingArea(pLmnIdentity *xnapiesv1.PlmnIdentity, rAnpagingAreaChoice *xnapiesv1.RanpagingAreaChoice) (*xnapiesv1.RanpagingArea, error) {

	msg := &xnapiesv1.RanpagingArea{}
	msg.PLmnIdentity = pLmnIdentity
	msg.RAnpagingAreaChoice = rAnpagingAreaChoice

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRanpagingArea() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateRedundantPdusessionInformation(rSn xnapiesv1.Rsn) (*xnapiesv1.RedundantPdusessionInformation, error) {

	msg := &xnapiesv1.RedundantPdusessionInformation{}
	msg.RSn = rSn

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRedundantPdusessionInformation() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateReferenceID(value int32) (*xnapiesv1.ReferenceId, error) {

	msg := &xnapiesv1.ReferenceId{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateReferenceID() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateReportCharacteristics(value *asn1.BitString) (*xnapiesv1.ReportCharacteristics, error) {

	msg := &xnapiesv1.ReportCharacteristics{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateReportCharacteristics() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateRequestReferenceID(value int32) (*xnapiesv1.RequestReferenceId, error) {

	msg := &xnapiesv1.RequestReferenceId{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRequestReferenceID() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateResetRequestTypeInfoFull() (*xnapiesv1.ResetRequestTypeInfoFull, error) {

	msg := &xnapiesv1.ResetRequestTypeInfoFull{}

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateResetRequestTypeInfoFull() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateResetRequestTypeInfoPartial(ueContextsToBeReleasedList *xnapiesv1.ResetRequestPartialReleaseList) (*xnapiesv1.ResetRequestTypeInfoPartial, error) {

	msg := &xnapiesv1.ResetRequestTypeInfoPartial{}
	msg.UeContextsToBeReleasedList = ueContextsToBeReleasedList

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateResetRequestTypeInfoPartial() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateResetRequestPartialReleaseList(value []*xnapiesv1.ResetRequestPartialReleaseItem) (*xnapiesv1.ResetRequestPartialReleaseList, error) {

	msg := &xnapiesv1.ResetRequestPartialReleaseList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateResetRequestPartialReleaseList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateResetResponseTypeInfoFull() (*xnapiesv1.ResetResponseTypeInfoFull, error) {

	msg := &xnapiesv1.ResetResponseTypeInfoFull{}

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateResetResponseTypeInfoFull() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateResetResponseTypeInfoPartial(ueContextsAdmittedToBeReleasedList *xnapiesv1.ResetResponsePartialReleaseList) (*xnapiesv1.ResetResponseTypeInfoPartial, error) {

	msg := &xnapiesv1.ResetResponseTypeInfoPartial{}
	msg.UeContextsAdmittedToBeReleasedList = ueContextsAdmittedToBeReleasedList

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateResetResponseTypeInfoPartial() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateResetResponsePartialReleaseList(value []*xnapiesv1.ResetResponsePartialReleaseItem) (*xnapiesv1.ResetResponsePartialReleaseList, error) {

	msg := &xnapiesv1.ResetResponsePartialReleaseList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateResetResponsePartialReleaseList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateRlcStatus(reestablishmentIndication xnapiesv1.ReestablishmentIndication) (*xnapiesv1.RlcStatus, error) {

	msg := &xnapiesv1.RlcStatus{}
	msg.ReestablishmentIndication = reestablishmentIndication

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRlcStatus() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateRlcduplicationStateList(value []*xnapiesv1.RlcduplicationStateItem) (*xnapiesv1.RlcduplicationStateList, error) {

	msg := &xnapiesv1.RlcduplicationStateList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRlcduplicationStateList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateRlcduplicationStateItem(duplicationState xnapiesv1.DuplicationStateRlcduplicationStateItem) (*xnapiesv1.RlcduplicationStateItem, error) {

	msg := &xnapiesv1.RlcduplicationStateItem{}
	msg.DuplicationState = duplicationState

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRlcduplicationStateItem() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateRfspIndex(value int32) (*xnapiesv1.RfspIndex, error) {

	msg := &xnapiesv1.RfspIndex{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRfspIndex() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateRrcconnections(noofRrcconnections *xnapiesv1.NoofRrcconnections, availableRrcconnectionCapacityValue *xnapiesv1.AvailableRrcconnectionCapacityValue) (*xnapiesv1.Rrcconnections, error) {

	msg := &xnapiesv1.Rrcconnections{}
	msg.NoofRrcconnections = noofRrcconnections
	msg.AvailableRrcconnectionCapacityValue = availableRrcconnectionCapacityValue

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRrcconnections() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateRrcreestabinitiated(rRrcreestabInitiatedReporting *xnapiesv1.RrcreestabInitiatedReporting) (*xnapiesv1.Rrcreestabinitiated, error) {

	msg := &xnapiesv1.Rrcreestabinitiated{}
	msg.RRrcreestabInitiatedReporting = rRrcreestabInitiatedReporting

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRrcreestabinitiated() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateRrcreestabInitiatedReportingwoUErlfreport(failureCellPci *xnapiesv1.NgRAnCellPci, reestabCellCgi *xnapiesv1.GlobalNgRAncellID, cRnti *xnapiesv1.CRNti, shortMacI *xnapiesv1.MacI) (*xnapiesv1.RrcreestabInitiatedReportingwoUErlfreport, error) {

	msg := &xnapiesv1.RrcreestabInitiatedReportingwoUErlfreport{}
	msg.FailureCellPci = failureCellPci
	msg.ReestabCellCgi = reestabCellCgi
	msg.CRnti = cRnti
	msg.ShortMacI = shortMacI

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRrcreestabInitiatedReportingwoUErlfreport() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateRrcreestabInitiatedReportingwithUErlfreport(uErlfreportContainer *xnapiesv1.UerlfreportContainer) (*xnapiesv1.RrcreestabInitiatedReportingwithUErlfreport, error) {

	msg := &xnapiesv1.RrcreestabInitiatedReportingwithUErlfreport{}
	msg.UErlfreportContainer = uErlfreportContainer

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRrcreestabInitiatedReportingwithUErlfreport() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateRrcsetupInitiatedReportingwithUErlfreport(uErlfreportContainer *xnapiesv1.UerlfreportContainer) (*xnapiesv1.RrcsetupInitiatedReportingwithUErlfreport, error) {

	msg := &xnapiesv1.RrcsetupInitiatedReportingwithUErlfreport{}
	msg.UErlfreportContainer = uErlfreportContainer

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRrcsetupInitiatedReportingwithUErlfreport() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateSecondarydataForwardingInfoFromTargetItem(secondarydataForwardingInfoFromTarget *xnapiesv1.DataForwardingInfoFromTargetNgrannode) (*xnapiesv1.SecondarydataForwardingInfoFromTargetItem, error) {

	msg := &xnapiesv1.SecondarydataForwardingInfoFromTargetItem{}
	msg.SecondarydataForwardingInfoFromTarget = secondarydataForwardingInfoFromTarget

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSecondarydataForwardingInfoFromTargetItem() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateSecondarydataForwardingInfoFromTargetList(value []*xnapiesv1.SecondarydataForwardingInfoFromTargetItem) (*xnapiesv1.SecondarydataForwardingInfoFromTargetList, error) {

	msg := &xnapiesv1.SecondarydataForwardingInfoFromTargetList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSecondarydataForwardingInfoFromTargetList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateSecurityResult(integrityProtectionResult xnapiesv1.IntegrityProtectionResultSecurityResult, confIDentialityProtectionResult xnapiesv1.ConfidentialityProtectionResultSecurityResult) (*xnapiesv1.SecurityResult, error) {

	msg := &xnapiesv1.SecurityResult{}
	msg.IntegrityProtectionResult = integrityProtectionResult
	msg.ConfidentialityProtectionResult = confIDentialityProtectionResult

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSecurityResult() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateSensorMeasConfigNameList(value []*xnapiesv1.SensorName) (*xnapiesv1.SensorMeasConfigNameList, error) {

	msg := &xnapiesv1.SensorMeasConfigNameList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSensorMeasConfigNameList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateServedCellInformationEUTraExtIes(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, extension *xnapiesv1.ServedCellInformationEUTraExtIesExtension) (*xnapiesv1.ServedCellInformationEUTraExtIes, error) {

	msg := &xnapiesv1.ServedCellInformationEUTraExtIes{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Extension = extension

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateServedCellInformationEUTraExtIes() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateServedCellInformationEUTraperBplmn(plmnID *xnapiesv1.PlmnIdentity) (*xnapiesv1.ServedCellInformationEUTraperBplmn, error) {

	msg := &xnapiesv1.ServedCellInformationEUTraperBplmn{}
	msg.PlmnId = plmnID

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateServedCellInformationEUTraperBplmn() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateServedCellInformationEUTraFDdinfoExtIes(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, extension *xnapiesv1.ServedCellInformationEUTraFDdinfoExtIesExtension) (*xnapiesv1.ServedCellInformationEUTraFDdinfoExtIes, error) {

	msg := &xnapiesv1.ServedCellInformationEUTraFDdinfoExtIes{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Extension = extension

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateServedCellInformationEUTraFDdinfoExtIes() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateServedCellInformationEUTraFDdinfo(ulEarfcn *xnapiesv1.EUTraarfcn, dlEarfcn *xnapiesv1.EUTraarfcn, ulEUtraTxBw xnapiesv1.EUTratransmissionBandwidth, dlEUtraTxBw xnapiesv1.EUTratransmissionBandwidth) (*xnapiesv1.ServedCellInformationEUTraFDdinfo, error) {

	msg := &xnapiesv1.ServedCellInformationEUTraFDdinfo{}
	msg.UlEarfcn = ulEarfcn
	msg.DlEarfcn = dlEarfcn
	msg.UlEUtraTxBw = ulEUtraTxBw
	msg.DlEUtraTxBw = dlEUtraTxBw

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateServedCellInformationEUTraFDdinfo() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateServedCellInformationEUTraTDdinfoExtIes(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, extension *xnapiesv1.ServedCellInformationEUTraTDdinfoExtIesExtension) (*xnapiesv1.ServedCellInformationEUTraTDdinfoExtIes, error) {

	msg := &xnapiesv1.ServedCellInformationEUTraTDdinfoExtIes{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Extension = extension

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateServedCellInformationEUTraTDdinfoExtIes() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateServedCellInformationEUTraTDdinfo(earfcn *xnapiesv1.EUTraarfcn, eUtraTxBw xnapiesv1.EUTratransmissionBandwidth, subframeAssignmnet xnapiesv1.SubframeAssignmnetServedCellInformationEutratddinfo, specialSubframeInfo *xnapiesv1.SpecialSubframeInfoEUTra) (*xnapiesv1.ServedCellInformationEUTraTDdinfo, error) {

	msg := &xnapiesv1.ServedCellInformationEUTraTDdinfo{}
	msg.Earfcn = earfcn
	msg.EUtraTxBw = eUtraTxBw
	msg.SubframeAssignmnet = subframeAssignmnet
	msg.SpecialSubframeInfo = specialSubframeInfo

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateServedCellInformationEUTraTDdinfo() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateServedCellsEUTra(value []*xnapiesv1.ServedCellsEUTraItem) (*xnapiesv1.ServedCellsEUTra, error) {

	msg := &xnapiesv1.ServedCellsEUTra{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateServedCellsEUTra() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateServedCellsToModifyEUTra(value []*xnapiesv1.ServedCellsToModifyEUTraItem) (*xnapiesv1.ServedCellsToModifyEUTra, error) {

	msg := &xnapiesv1.ServedCellsToModifyEUTra{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateServedCellsToModifyEUTra() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateServedCellInformationNRExtIes(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, extension *xnapiesv1.ServedCellInformationNRExtIesExtension) (*xnapiesv1.ServedCellInformationNRExtIes, error) {

	msg := &xnapiesv1.ServedCellInformationNRExtIes{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Extension = extension

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateServedCellInformationNRExtIes() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateSfnOffset(sFnTimeOffset *asn1.BitString) (*xnapiesv1.SfnOffset, error) {

	msg := &xnapiesv1.SfnOffset{}
	msg.SFnTimeOffset = sFnTimeOffset

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSfnOffset() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateServedCellsNR(value []*xnapiesv1.ServedCellsNRItem) (*xnapiesv1.ServedCellsNR, error) {

	msg := &xnapiesv1.ServedCellsNR{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateServedCellsNR() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateServedCellsToModifyNR(value []*xnapiesv1.ServedCellsToModifyNRItem) (*xnapiesv1.ServedCellsToModifyNR, error) {

	msg := &xnapiesv1.ServedCellsToModifyNR{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateServedCellsToModifyNR() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateSharedResourceTypeULOnlySharing(ulResourceBitmap *xnapiesv1.DataTrafficResources) (*xnapiesv1.SharedResourceTypeULOnlySharing, error) {

	msg := &xnapiesv1.SharedResourceTypeULOnlySharing{}
	msg.UlResourceBitmap = ulResourceBitmap

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSharedResourceTypeULOnlySharing() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateSharedResourceTypeULdlSharingULResourcesChanged(ulResourceBitmap *xnapiesv1.DataTrafficResources) (*xnapiesv1.SharedResourceTypeULdlSharingULResourcesChanged, error) {

	msg := &xnapiesv1.SharedResourceTypeULdlSharingULResourcesChanged{}
	msg.UlResourceBitmap = ulResourceBitmap

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSharedResourceTypeULdlSharingULResourcesChanged() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateSharedResourceTypeULdlSharingDLResourcesChanged(dlResourceBitmap *xnapiesv1.DataTrafficResources) (*xnapiesv1.SharedResourceTypeULdlSharingDLResourcesChanged, error) {

	msg := &xnapiesv1.SharedResourceTypeULdlSharingDLResourcesChanged{}
	msg.DlResourceBitmap = dlResourceBitmap

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSharedResourceTypeULdlSharingDLResourcesChanged() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateSliceAvailableCapacity(value []*xnapiesv1.SliceAvailableCapacityItem) (*xnapiesv1.SliceAvailableCapacity, error) {

	msg := &xnapiesv1.SliceAvailableCapacity{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSliceAvailableCapacity() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateSliceAvailableCapacityItem(pLmnidentity *xnapiesv1.PlmnIdentity, sNssaiavailableCapacityList *xnapiesv1.SnssaiavailableCapacityList) (*xnapiesv1.SliceAvailableCapacityItem, error) {

	msg := &xnapiesv1.SliceAvailableCapacityItem{}
	msg.PLmnidentity = pLmnidentity
	msg.SNssaiavailableCapacityList = sNssaiavailableCapacityList

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSliceAvailableCapacityItem() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateSnssaiavailableCapacityList(value []*xnapiesv1.SnssaiavailableCapacityItem) (*xnapiesv1.SnssaiavailableCapacityList, error) {

	msg := &xnapiesv1.SnssaiavailableCapacityList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnssaiavailableCapacityList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateSnssaiavailableCapacityItem(sNssai *xnapiesv1.SNSsai, sliceAvailableCapacityValueDownlink int32, sliceAvailableCapacityValueUplink int32) (*xnapiesv1.SnssaiavailableCapacityItem, error) {

	msg := &xnapiesv1.SnssaiavailableCapacityItem{}
	msg.SNssai = sNssai
	msg.SliceAvailableCapacityValueDownlink = sliceAvailableCapacityValueDownlink
	msg.SliceAvailableCapacityValueUplink = sliceAvailableCapacityValueUplink

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnssaiavailableCapacityItem() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateSliceSupportList(value []*xnapiesv1.SNSsai) (*xnapiesv1.SliceSupportList, error) {

	msg := &xnapiesv1.SliceSupportList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSliceSupportList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateSliceToReportList(value []*xnapiesv1.SliceToReportListItem) (*xnapiesv1.SliceToReportList, error) {

	msg := &xnapiesv1.SliceToReportList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSliceToReportList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateSliceToReportListItem(pLmnidentity *xnapiesv1.PlmnIdentity, sNssailist *xnapiesv1.Snssailist) (*xnapiesv1.SliceToReportListItem, error) {

	msg := &xnapiesv1.SliceToReportListItem{}
	msg.PLmnidentity = pLmnidentity
	msg.SNssailist = sNssailist

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSliceToReportListItem() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateSnssailist(value []*xnapiesv1.SnssaiItem) (*xnapiesv1.Snssailist, error) {

	msg := &xnapiesv1.Snssailist{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnssailist() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateSnssaiItem(sNssai *xnapiesv1.SNSsai) (*xnapiesv1.SnssaiItem, error) {

	msg := &xnapiesv1.SnssaiItem{}
	msg.SNssai = sNssai

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnssaiItem() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateSlotConfigurationList(value []*xnapiesv1.SlotConfigurationListItem) (*xnapiesv1.SlotConfigurationList, error) {

	msg := &xnapiesv1.SlotConfigurationList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSlotConfigurationList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateSlotConfigurationListItem(slotIndex int32, symbolAllocationInSlot *xnapiesv1.SymbolAllocationinSlot) (*xnapiesv1.SlotConfigurationListItem, error) {

	msg := &xnapiesv1.SlotConfigurationListItem{}
	msg.SlotIndex = slotIndex
	msg.SymbolAllocationInSlot = symbolAllocationInSlot

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSlotConfigurationListItem() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateSNGRAnnodeSecurityKey(value *asn1.BitString) (*xnapiesv1.SNGRAnnodeSecurityKey, error) {

	msg := &xnapiesv1.SNGRAnnodeSecurityKey{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSNGRAnnodeSecurityKey() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateSpecialSubframeInfoEUTra(specialSubframePattern xnapiesv1.SpecialSubframePatternsEUTra, cyclicPrefixDl xnapiesv1.CyclicPrefixEUTraDL, cyclicPrefixUl xnapiesv1.CyclicPrefixEUTraUL) (*xnapiesv1.SpecialSubframeInfoEUTra, error) {

	msg := &xnapiesv1.SpecialSubframeInfoEUTra{}
	msg.SpecialSubframePattern = specialSubframePattern
	msg.CyclicPrefixDl = cyclicPrefixDl
	msg.CyclicPrefixUl = cyclicPrefixUl

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSpecialSubframeInfoEUTra() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateSpectrumSharingGroupID(value int32) (*xnapiesv1.SpectrumSharingGroupId, error) {

	msg := &xnapiesv1.SpectrumSharingGroupId{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSpectrumSharingGroupID() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateSsbareaCapacityValueList(value []*xnapiesv1.SsbareaCapacityValueListItem) (*xnapiesv1.SsbareaCapacityValueList, error) {

	msg := &xnapiesv1.SsbareaCapacityValueList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSsbareaCapacityValueList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateSsbareaCapacityValueListItem(sSbindex int32, ssbAreaCapacityValue int32) (*xnapiesv1.SsbareaCapacityValueListItem, error) {

	msg := &xnapiesv1.SsbareaCapacityValueListItem{}
	msg.SSbindex = sSbindex
	msg.SsbAreaCapacityValue = ssbAreaCapacityValue

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSsbareaCapacityValueListItem() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateSsbareaRadioResourceStatusList(value []*xnapiesv1.SsbareaRadioResourceStatusListItem) (*xnapiesv1.SsbareaRadioResourceStatusList, error) {

	msg := &xnapiesv1.SsbareaRadioResourceStatusList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSsbareaRadioResourceStatusList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateSsbareaRadioResourceStatusListItemExtIes(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, extension *xnapiesv1.SsbareaRadioResourceStatusListItemExtIesExtension) (*xnapiesv1.SsbareaRadioResourceStatusListItemExtIes, error) {

	msg := &xnapiesv1.SsbareaRadioResourceStatusListItemExtIes{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Extension = extension

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSsbareaRadioResourceStatusListItemExtIes() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateSsbareaRadioResourceStatusListItem(sSbindex int32, ssbAreaDlGbrPrbUsage *xnapiesv1.DlGBrPRbusage, ssbAreaUlGbrPrbUsage *xnapiesv1.UlGBrPRbusage, ssbAreaDLNonGbrPrbUsage *xnapiesv1.DlnonGBrPRbusage, ssbAreaULNonGbrPrbUsage *xnapiesv1.UlnonGBrPRbusage, ssbAreaDLTotalPrbUsage *xnapiesv1.DlTotalPRbusage, ssbAreaULTotalPrbUsage *xnapiesv1.UlTotalPRbusage) (*xnapiesv1.SsbareaRadioResourceStatusListItem, error) {

	msg := &xnapiesv1.SsbareaRadioResourceStatusListItem{}
	msg.SSbindex = sSbindex
	msg.SsbAreaDlGbrPrbUsage = ssbAreaDlGbrPrbUsage
	msg.SsbAreaUlGbrPrbUsage = ssbAreaUlGbrPrbUsage
	msg.SsbAreaDLNonGbrPrbUsage = ssbAreaDLNonGbrPrbUsage
	msg.SsbAreaULNonGbrPrbUsage = ssbAreaULNonGbrPrbUsage
	msg.SsbAreaDLTotalPrbUsage = ssbAreaDLTotalPrbUsage
	msg.SsbAreaULTotalPrbUsage = ssbAreaULTotalPrbUsage

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSsbareaRadioResourceStatusListItem() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateSsbtoReportList(value []*xnapiesv1.SsbtoReportListItem) (*xnapiesv1.SsbtoReportList, error) {

	msg := &xnapiesv1.SsbtoReportList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSsbtoReportList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateSsbtoReportListItem(sSbindex int32) (*xnapiesv1.SsbtoReportListItem, error) {

	msg := &xnapiesv1.SsbtoReportListItem{}
	msg.SSbindex = sSbindex

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSsbtoReportListItem() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateSulFrequencyBand(value int32) (*xnapiesv1.SulFrequencyBand, error) {

	msg := &xnapiesv1.SulFrequencyBand{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSulFrequencyBand() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateSulInformationExtIes(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, extension *xnapiesv1.SulInformationExtIesExtension) (*xnapiesv1.SulInformationExtIes, error) {

	msg := &xnapiesv1.SulInformationExtIes{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Extension = extension

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSulInformationExtIes() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateSulInformation(sulFrequencyInfo *xnapiesv1.Nrarfcn, sulTransmissionBandwIDth *xnapiesv1.NrtransmissionBandwidth) (*xnapiesv1.SulInformation, error) {

	msg := &xnapiesv1.SulInformation{}
	msg.SulFrequencyInfo = sulFrequencyInfo
	msg.SulTransmissionBandwidth = sulTransmissionBandwIDth

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSulInformation() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateSupportedSulbandList(value []*xnapiesv1.SupportedSulbandItem) (*xnapiesv1.SupportedSulbandList, error) {

	msg := &xnapiesv1.SupportedSulbandList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSupportedSulbandList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateSupportedSulbandItem(sulBandItem *xnapiesv1.SulFrequencyBand) (*xnapiesv1.SupportedSulbandItem, error) {

	msg := &xnapiesv1.SupportedSulbandItem{}
	msg.SulBandItem = sulBandItem

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSupportedSulbandItem() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateSymbolAllocationinSlotAllDl() (*xnapiesv1.SymbolAllocationinSlotAllDl, error) {

	msg := &xnapiesv1.SymbolAllocationinSlotAllDl{}

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSymbolAllocationinSlotAllDl() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateSymbolAllocationinSlotAllUl() (*xnapiesv1.SymbolAllocationinSlotAllUl, error) {

	msg := &xnapiesv1.SymbolAllocationinSlotAllUl{}

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSymbolAllocationinSlotAllUl() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateSymbolAllocationinSlotBothDlandUl(numberofDlsymbols int32, numberofUlsymbols int32) (*xnapiesv1.SymbolAllocationinSlotBothDlandUl, error) {

	msg := &xnapiesv1.SymbolAllocationinSlotBothDlandUl{}
	msg.NumberofDlsymbols = numberofDlsymbols
	msg.NumberofUlsymbols = numberofUlsymbols

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSymbolAllocationinSlotBothDlandUl() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateTabasedMdt(tAlistforMdt *xnapiesv1.TalistforMdt) (*xnapiesv1.TabasedMdt, error) {

	msg := &xnapiesv1.TabasedMdt{}
	msg.TAlistforMdt = tAlistforMdt

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateTabasedMdt() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateTaibasedMdt(tAilistforMdt *xnapiesv1.TailistforMdt) (*xnapiesv1.TaibasedMdt, error) {

	msg := &xnapiesv1.TaibasedMdt{}
	msg.TAilistforMdt = tAilistforMdt

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateTaibasedMdt() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateTailistforMdt(value []*xnapiesv1.TaiforMdtItem) (*xnapiesv1.TailistforMdt, error) {

	msg := &xnapiesv1.TailistforMdt{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateTailistforMdt() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateTaiforMdtItem(plmnID *xnapiesv1.PlmnIdentity, tAc *xnapiesv1.Tac) (*xnapiesv1.TaiforMdtItem, error) {

	msg := &xnapiesv1.TaiforMdtItem{}
	msg.PlmnId = plmnID
	msg.TAc = tAc

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateTaiforMdtItem() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateTac(value []byte) (*xnapiesv1.Tac, error) {

	msg := &xnapiesv1.Tac{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateTac() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateTaisupportList(value []*xnapiesv1.TaisupportItem) (*xnapiesv1.TaisupportList, error) {

	msg := &xnapiesv1.TaisupportList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateTaisupportList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateTaisupportItem(tac *xnapiesv1.Tac, broadcastPlmns []*xnapiesv1.BroadcastPlmninTaisupportItem) (*xnapiesv1.TaisupportItem, error) {

	msg := &xnapiesv1.TaisupportItem{}
	msg.Tac = tac
	msg.BroadcastPlmns = broadcastPlmns

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateTaisupportItem() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateTalistforMdt(value []*xnapiesv1.Tac) (*xnapiesv1.TalistforMdt, error) {

	msg := &xnapiesv1.TalistforMdt{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateTalistforMdt() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateTargetCellinEutran(value []byte) (*xnapiesv1.TargetCellinEutran, error) {

	msg := &xnapiesv1.TargetCellinEutran{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateTargetCellinEutran() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateTdduldlconfigurationCommonNr(value []byte) (*xnapiesv1.TdduldlconfigurationCommonNr, error) {

	msg := &xnapiesv1.TdduldlconfigurationCommonNr{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateTdduldlconfigurationCommonNr() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateTargetCellList(value []*xnapiesv1.TargetCellListItem) (*xnapiesv1.TargetCellList, error) {

	msg := &xnapiesv1.TargetCellList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateTargetCellList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateTargetCellListItem(targetCell *xnapiesv1.TargetCGi) (*xnapiesv1.TargetCellListItem, error) {

	msg := &xnapiesv1.TargetCellListItem{}
	msg.TargetCell = targetCell

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateTargetCellListItem() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateThresholdRSrq(value int32) (*xnapiesv1.ThresholdRSrq, error) {

	msg := &xnapiesv1.ThresholdRSrq{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateThresholdRSrq() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateThresholdRSrp(value int32) (*xnapiesv1.ThresholdRSrp, error) {

	msg := &xnapiesv1.ThresholdRSrp{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateThresholdRSrp() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateThresholdSInr(value int32) (*xnapiesv1.ThresholdSInr, error) {

	msg := &xnapiesv1.ThresholdSInr{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateThresholdSInr() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateTnlaToAddList(value []*xnapiesv1.TnlaToAddItem) (*xnapiesv1.TnlaToAddList, error) {

	msg := &xnapiesv1.TnlaToAddList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateTnlaToAddList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateTnlaToAddItem(tNlassociationTransportLayerAddress *xnapiesv1.CptransportLayerInformation, tNlassociationUsage xnapiesv1.TnlassociationUsage) (*xnapiesv1.TnlaToAddItem, error) {

	msg := &xnapiesv1.TnlaToAddItem{}
	msg.TNlassociationTransportLayerAddress = tNlassociationTransportLayerAddress
	msg.TNlassociationUsage = tNlassociationUsage

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateTnlaToAddItem() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateTnlaToUpdateList(value []*xnapiesv1.TnlaToUpdateItem) (*xnapiesv1.TnlaToUpdateList, error) {

	msg := &xnapiesv1.TnlaToUpdateList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateTnlaToUpdateList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateTnlaToRemoveList(value []*xnapiesv1.TnlaToRemoveItem) (*xnapiesv1.TnlaToRemoveList, error) {

	msg := &xnapiesv1.TnlaToRemoveList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateTnlaToRemoveList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateTnlaToRemoveItem(tNlassociationTransportLayerAddress *xnapiesv1.CptransportLayerInformation) (*xnapiesv1.TnlaToRemoveItem, error) {

	msg := &xnapiesv1.TnlaToRemoveItem{}
	msg.TNlassociationTransportLayerAddress = tNlassociationTransportLayerAddress

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateTnlaToRemoveItem() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateTnlaSetupList(value []*xnapiesv1.TnlaSetupItem) (*xnapiesv1.TnlaSetupList, error) {

	msg := &xnapiesv1.TnlaSetupList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateTnlaSetupList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateTnlaSetupItem(tNlassociationTransportLayerAddress *xnapiesv1.CptransportLayerInformation) (*xnapiesv1.TnlaSetupItem, error) {

	msg := &xnapiesv1.TnlaSetupItem{}
	msg.TNlassociationTransportLayerAddress = tNlassociationTransportLayerAddress

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateTnlaSetupItem() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateTnlaFailedToSetupList(value []*xnapiesv1.TnlaFailedToSetupItem) (*xnapiesv1.TnlaFailedToSetupList, error) {

	msg := &xnapiesv1.TnlaFailedToSetupList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateTnlaFailedToSetupList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateTnlaFailedToSetupItem(tNlassociationTransportLayerAddress *xnapiesv1.CptransportLayerInformation, cause *xnapiesv1.Cause) (*xnapiesv1.TnlaFailedToSetupItem, error) {

	msg := &xnapiesv1.TnlaFailedToSetupItem{}
	msg.TNlassociationTransportLayerAddress = tNlassociationTransportLayerAddress
	msg.Cause = cause

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateTnlaFailedToSetupItem() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateTransportLayerAddress(value *asn1.BitString) (*xnapiesv1.TransportLayerAddress, error) {

	msg := &xnapiesv1.TransportLayerAddress{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateTransportLayerAddress() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateTraceActivationExtIes(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, extension *xnapiesv1.TraceActivationExtIesExtension) (*xnapiesv1.TraceActivationExtIes, error) {

	msg := &xnapiesv1.TraceActivationExtIes{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Extension = extension

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateTraceActivationExtIes() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateTraceActivation(ngRanTraceID *xnapiesv1.NgRAntraceId, interfacesToTrace *asn1.BitString, traceDepth xnapiesv1.TraceDepth, traceCollAddress *xnapiesv1.TransportLayerAddress) (*xnapiesv1.TraceActivation, error) {

	msg := &xnapiesv1.TraceActivation{}
	msg.NgRanTraceId = ngRanTraceID
	msg.InterfacesToTrace = interfacesToTrace
	msg.TraceDepth = traceDepth
	msg.TraceCollAddress = traceCollAddress

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateTraceActivation() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateUeaggregateMaximumBitRate(dlUeAmbr *xnapiesv1.BitRate, ulUeAmbr *xnapiesv1.BitRate) (*xnapiesv1.UeaggregateMaximumBitRate, error) {

	msg := &xnapiesv1.UeaggregateMaximumBitRate{}
	msg.DlUeAmbr = dlUeAmbr
	msg.UlUeAmbr = ulUeAmbr

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateUeaggregateMaximumBitRate() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateUecontextIDforRrcresume(iRnti *xnapiesv1.IRNti, allocatedCRnti *xnapiesv1.CRNti, accessPci *xnapiesv1.NgRAnCellPci) (*xnapiesv1.UecontextIdforRrcresume, error) {

	msg := &xnapiesv1.UecontextIdforRrcresume{}
	msg.IRnti = iRnti
	msg.AllocatedCRnti = allocatedCRnti
	msg.AccessPci = accessPci

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateUecontextIDforRrcresume() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateUecontextIDforRrcreestablishment(cRnti *xnapiesv1.CRNti, failureCellPci *xnapiesv1.NgRAnCellPci) (*xnapiesv1.UecontextIdforRrcreestablishment, error) {

	msg := &xnapiesv1.UecontextIdforRrcreestablishment{}
	msg.CRnti = cRnti
	msg.FailureCellPci = failureCellPci

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateUecontextIDforRrcreestablishment() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateUecontextInfoRetrUectxtRespExtIes(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, extension *xnapiesv1.UecontextInfoRetrUectxtRespExtIesExtension) (*xnapiesv1.UecontextInfoRetrUectxtRespExtIes, error) {

	msg := &xnapiesv1.UecontextInfoRetrUectxtRespExtIes{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Extension = extension

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateUecontextInfoRetrUectxtRespExtIes() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateUehistoryInformation(value []*xnapiesv1.LastVisitedCellItem) (*xnapiesv1.UehistoryInformation, error) {

	msg := &xnapiesv1.UehistoryInformation{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateUehistoryInformation() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateUeradioCapabilityForPagingOfNr(value []byte) (*xnapiesv1.UeradioCapabilityForPagingOfNr, error) {

	msg := &xnapiesv1.UeradioCapabilityForPagingOfNr{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateUeradioCapabilityForPagingOfNr() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateUeradioCapabilityForPagingOfEutra(value []byte) (*xnapiesv1.UeradioCapabilityForPagingOfEutra, error) {

	msg := &xnapiesv1.UeradioCapabilityForPagingOfEutra{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateUeradioCapabilityForPagingOfEutra() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateUeradioCapabilityID(value []byte) (*xnapiesv1.UeradioCapabilityId, error) {

	msg := &xnapiesv1.UeradioCapabilityId{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateUeradioCapabilityID() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateUerlfreportContainerLte(value []byte) (*xnapiesv1.UerlfreportContainerLte, error) {

	msg := &xnapiesv1.UerlfreportContainerLte{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateUerlfreportContainerLte() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateUerlfreportContainerNr(value []byte) (*xnapiesv1.UerlfreportContainerNr, error) {

	msg := &xnapiesv1.UerlfreportContainerNr{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateUerlfreportContainerNr() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateUesecurityCapabilities(nrEncyptionAlgorithms *asn1.BitString, nrIntegrityProtectionAlgorithms *asn1.BitString, eUtraEncyptionAlgorithms *asn1.BitString, eUtraIntegrityProtectionAlgorithms *asn1.BitString) (*xnapiesv1.UesecurityCapabilities, error) {

	msg := &xnapiesv1.UesecurityCapabilities{}
	msg.NrEncyptionAlgorithms = nrEncyptionAlgorithms
	msg.NrIntegrityProtectionAlgorithms = nrIntegrityProtectionAlgorithms
	msg.EUtraEncyptionAlgorithms = eUtraEncyptionAlgorithms
	msg.EUtraIntegrityProtectionAlgorithms = eUtraIntegrityProtectionAlgorithms

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateUesecurityCapabilities() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateUlconfiguration(uLPdcp xnapiesv1.UlUEConfiguration) (*xnapiesv1.Ulconfiguration, error) {

	msg := &xnapiesv1.Ulconfiguration{}
	msg.ULPdcp = uLPdcp

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateUlconfiguration() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateUlGBrPRbusage(value int32) (*xnapiesv1.UlGBrPRbusage, error) {

	msg := &xnapiesv1.UlGBrPRbusage{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateUlGBrPRbusage() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateUlnonGBrPRbusage(value int32) (*xnapiesv1.UlnonGBrPRbusage, error) {

	msg := &xnapiesv1.UlnonGBrPRbusage{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateUlnonGBrPRbusage() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateUlTotalPRbusage(value int32) (*xnapiesv1.UlTotalPRbusage, error) {

	msg := &xnapiesv1.UlTotalPRbusage{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateUlTotalPRbusage() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateUptransportParameters(value []*xnapiesv1.UptransportParametersItem) (*xnapiesv1.UptransportParameters, error) {

	msg := &xnapiesv1.UptransportParameters{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateUptransportParameters() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateUptransportParametersItem(upTnlinfo *xnapiesv1.UptransportLayerInformation, cellGroupID *xnapiesv1.CellGroupId) (*xnapiesv1.UptransportParametersItem, error) {

	msg := &xnapiesv1.UptransportParametersItem{}
	msg.UpTnlinfo = upTnlinfo
	msg.CellGroupId = cellGroupID

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateUptransportParametersItem() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateUriaddress(value int32) (*xnapiesv1.Uriaddress, error) {

	msg := &xnapiesv1.Uriaddress{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateUriaddress() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateVolumeTimedReportList(value []*xnapiesv1.VolumeTimedReportItem) (*xnapiesv1.VolumeTimedReportList, error) {

	msg := &xnapiesv1.VolumeTimedReportList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateVolumeTimedReportList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateVolumeTimedReportItem(startTimeStamp []byte, endTimeStamp []byte, usageCountUl int32, usageCountDl int32) (*xnapiesv1.VolumeTimedReportItem, error) {

	msg := &xnapiesv1.VolumeTimedReportItem{}
	msg.StartTimeStamp = startTimeStamp
	msg.EndTimeStamp = endTimeStamp
	msg.UsageCountUl = usageCountUl
	msg.UsageCountDl = usageCountDl

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateVolumeTimedReportItem() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateWlanmeasConfigNameList(value []*xnapiesv1.Wlanname) (*xnapiesv1.WlanmeasConfigNameList, error) {

	msg := &xnapiesv1.WlanmeasConfigNameList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateWlanmeasConfigNameList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateWlanname(value []byte) (*xnapiesv1.Wlanname, error) {

	msg := &xnapiesv1.Wlanname{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateWlanname() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateXnBenefitValue(value int32) (*xnapiesv1.XnBenefitValue, error) {

	msg := &xnapiesv1.XnBenefitValue{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateXnBenefitValue() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateProtocolIeContainer001(value []int32) (*xnapcontainersv1.ProtocolIeContainer001, error) {

	msg := &xnapcontainersv1.ProtocolIeContainer001{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateProtocolIeContainer001() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateProtocolIeField001(ID int32, criticality int32, value int32) (*xnapcontainersv1.ProtocolIeField001, error) {

	msg := &xnapcontainersv1.ProtocolIeField001{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateProtocolIeField001() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateProtocolIeSingleContainer001(value *xnapcontainersv1.ProtocolIeField001) (*xnapcontainersv1.ProtocolIeSingleContainer001, error) {

	msg := &xnapcontainersv1.ProtocolIeSingleContainer001{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateProtocolIeSingleContainer001() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateProtocolIeContainerPair(value []int32) (*xnapcontainersv1.ProtocolIeContainerPair, error) {

	msg := &xnapcontainersv1.ProtocolIeContainerPair{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateProtocolIeContainerPair() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateProtocolIeFieldPair(ID int32) (*xnapcontainersv1.ProtocolIeFieldPair, error) {

	msg := &xnapcontainersv1.ProtocolIeFieldPair{}
	msg.Id = ID

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateProtocolIeFieldPair() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateProtocolIeContainerList(value []int32) (*xnapcontainersv1.ProtocolIeContainerList, error) {

	msg := &xnapcontainersv1.ProtocolIeContainerList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateProtocolIeContainerList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateProtocolIeContainerPairList(value []int32) (*xnapcontainersv1.ProtocolIeContainerPairList, error) {

	msg := &xnapcontainersv1.ProtocolIeContainerPairList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateProtocolIeContainerPairList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateProtocolExtensionContainer001(value []int32) (*xnapcontainersv1.ProtocolExtensionContainer001, error) {

	msg := &xnapcontainersv1.ProtocolExtensionContainer001{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateProtocolExtensionContainer001() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDhandoverPreparation(value int32) (*xnapconstantsv1.IdhandoverPreparation, error) {

	msg := &xnapconstantsv1.IdhandoverPreparation{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDhandoverPreparation() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDsNstatusTransfer(value int32) (*xnapconstantsv1.IdsNstatusTransfer, error) {

	msg := &xnapconstantsv1.IdsNstatusTransfer{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDsNstatusTransfer() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDhandoverCancel(value int32) (*xnapconstantsv1.IdhandoverCancel, error) {

	msg := &xnapconstantsv1.IdhandoverCancel{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDhandoverCancel() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDretrieveUecontext(value int32) (*xnapconstantsv1.IdretrieveUecontext, error) {

	msg := &xnapconstantsv1.IdretrieveUecontext{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDretrieveUecontext() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDrAnpaging(value int32) (*xnapconstantsv1.IdrAnpaging, error) {

	msg := &xnapconstantsv1.IdrAnpaging{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDrAnpaging() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDxnUaddressIndication(value int32) (*xnapconstantsv1.IdxnUaddressIndication, error) {

	msg := &xnapconstantsv1.IdxnUaddressIndication{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDxnUaddressIndication() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDuEcontextRelease(value int32) (*xnapconstantsv1.IduEcontextRelease, error) {

	msg := &xnapconstantsv1.IduEcontextRelease{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDuEcontextRelease() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDsNgrannodeAdditionPreparation(value int32) (*xnapconstantsv1.IdsNgrannodeAdditionPreparation, error) {

	msg := &xnapconstantsv1.IdsNgrannodeAdditionPreparation{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDsNgrannodeAdditionPreparation() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDsNgrannodeReconfigurationCompletion(value int32) (*xnapconstantsv1.IdsNgrannodeReconfigurationCompletion, error) {

	msg := &xnapconstantsv1.IdsNgrannodeReconfigurationCompletion{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDsNgrannodeReconfigurationCompletion() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDmNgrannodeinitiatedSngrannodeModificationPreparation(value int32) (*xnapconstantsv1.IdmNgrannodeinitiatedSngrannodeModificationPreparation, error) {

	msg := &xnapconstantsv1.IdmNgrannodeinitiatedSngrannodeModificationPreparation{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDmNgrannodeinitiatedSngrannodeModificationPreparation() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDsNgrannodeinitiatedSngrannodeModificationPreparation(value int32) (*xnapconstantsv1.IdsNgrannodeinitiatedSngrannodeModificationPreparation, error) {

	msg := &xnapconstantsv1.IdsNgrannodeinitiatedSngrannodeModificationPreparation{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDsNgrannodeinitiatedSngrannodeModificationPreparation() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDmNgrannodeinitiatedSngrannodeRelease(value int32) (*xnapconstantsv1.IdmNgrannodeinitiatedSngrannodeRelease, error) {

	msg := &xnapconstantsv1.IdmNgrannodeinitiatedSngrannodeRelease{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDmNgrannodeinitiatedSngrannodeRelease() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDsNgrannodeinitiatedSngrannodeRelease(value int32) (*xnapconstantsv1.IdsNgrannodeinitiatedSngrannodeRelease, error) {

	msg := &xnapconstantsv1.IdsNgrannodeinitiatedSngrannodeRelease{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDsNgrannodeinitiatedSngrannodeRelease() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDsNgrannodeCounterCheck(value int32) (*xnapconstantsv1.IdsNgrannodeCounterCheck, error) {

	msg := &xnapconstantsv1.IdsNgrannodeCounterCheck{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDsNgrannodeCounterCheck() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDsNgrannodeChange(value int32) (*xnapconstantsv1.IdsNgrannodeChange, error) {

	msg := &xnapconstantsv1.IdsNgrannodeChange{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDsNgrannodeChange() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDrRctransfer(value int32) (*xnapconstantsv1.IdrRctransfer, error) {

	msg := &xnapconstantsv1.IdrRctransfer{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDrRctransfer() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDxnRemoval(value int32) (*xnapconstantsv1.IdxnRemoval, error) {

	msg := &xnapconstantsv1.IdxnRemoval{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDxnRemoval() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDxnSetup(value int32) (*xnapconstantsv1.IdxnSetup, error) {

	msg := &xnapconstantsv1.IdxnSetup{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDxnSetup() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDnGrannodeConfigurationUpdate(value int32) (*xnapconstantsv1.IdnGrannodeConfigurationUpdate, error) {

	msg := &xnapconstantsv1.IdnGrannodeConfigurationUpdate{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDnGrannodeConfigurationUpdate() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDcellActivation(value int32) (*xnapconstantsv1.IdcellActivation, error) {

	msg := &xnapconstantsv1.IdcellActivation{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDcellActivation() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDreset(value int32) (*xnapconstantsv1.Idreset, error) {

	msg := &xnapconstantsv1.Idreset{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDreset() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDerrorIndication(value int32) (*xnapconstantsv1.IderrorIndication, error) {

	msg := &xnapconstantsv1.IderrorIndication{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDerrorIndication() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDprivateMessage(value int32) (*xnapconstantsv1.IdprivateMessage, error) {

	msg := &xnapconstantsv1.IdprivateMessage{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDprivateMessage() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDnotificationControl(value int32) (*xnapconstantsv1.IdnotificationControl, error) {

	msg := &xnapconstantsv1.IdnotificationControl{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDnotificationControl() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDactivityNotification(value int32) (*xnapconstantsv1.IdactivityNotification, error) {

	msg := &xnapconstantsv1.IdactivityNotification{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDactivityNotification() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDeUTraNRCellResourceCoordination(value int32) (*xnapconstantsv1.IdeUTraNRCellResourceCoordination, error) {

	msg := &xnapconstantsv1.IdeUTraNRCellResourceCoordination{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDeUTraNRCellResourceCoordination() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDsecondaryRatdataUsageReport(value int32) (*xnapconstantsv1.IdsecondaryRatdataUsageReport, error) {

	msg := &xnapconstantsv1.IdsecondaryRatdataUsageReport{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDsecondaryRatdataUsageReport() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDdeactivateTrace(value int32) (*xnapconstantsv1.IddeactivateTrace, error) {

	msg := &xnapconstantsv1.IddeactivateTrace{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDdeactivateTrace() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDtraceStart(value int32) (*xnapconstantsv1.IdtraceStart, error) {

	msg := &xnapconstantsv1.IdtraceStart{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDtraceStart() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDhandoverSuccess(value int32) (*xnapconstantsv1.IdhandoverSuccess, error) {

	msg := &xnapconstantsv1.IdhandoverSuccess{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDhandoverSuccess() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDconditionalHandoverCancel(value int32) (*xnapconstantsv1.IdconditionalHandoverCancel, error) {

	msg := &xnapconstantsv1.IdconditionalHandoverCancel{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDconditionalHandoverCancel() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDearlyStatusTransfer(value int32) (*xnapconstantsv1.IdearlyStatusTransfer, error) {

	msg := &xnapconstantsv1.IdearlyStatusTransfer{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDearlyStatusTransfer() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDfailureIndication(value int32) (*xnapconstantsv1.IdfailureIndication, error) {

	msg := &xnapconstantsv1.IdfailureIndication{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDfailureIndication() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDhandoverReport(value int32) (*xnapconstantsv1.IdhandoverReport, error) {

	msg := &xnapconstantsv1.IdhandoverReport{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDhandoverReport() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDresourceStatusReportingInitiation(value int32) (*xnapconstantsv1.IdresourceStatusReportingInitiation, error) {

	msg := &xnapconstantsv1.IdresourceStatusReportingInitiation{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDresourceStatusReportingInitiation() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDresourceStatusReporting(value int32) (*xnapconstantsv1.IdresourceStatusReporting, error) {

	msg := &xnapconstantsv1.IdresourceStatusReporting{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDresourceStatusReporting() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDmobilitySettingsChange(value int32) (*xnapconstantsv1.IdmobilitySettingsChange, error) {

	msg := &xnapconstantsv1.IdmobilitySettingsChange{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDmobilitySettingsChange() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDaccessAndMobilityIndication(value int32) (*xnapconstantsv1.IdaccessAndMobilityIndication, error) {

	msg := &xnapconstantsv1.IdaccessAndMobilityIndication{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDaccessAndMobilityIndication() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDActivatedServedCells(value int32) (*xnapconstantsv1.IdActivatedServedCells, error) {

	msg := &xnapconstantsv1.IdActivatedServedCells{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDActivatedServedCells() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDActivationIDforCellActivation(value int32) (*xnapconstantsv1.IdActivationIdforCellActivation, error) {

	msg := &xnapconstantsv1.IdActivationIdforCellActivation{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDActivationIDforCellActivation() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDadmittedSplitSrb(value int32) (*xnapconstantsv1.IdadmittedSplitSrb, error) {

	msg := &xnapconstantsv1.IdadmittedSplitSrb{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDadmittedSplitSrb() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDadmittedSplitSrbrelease(value int32) (*xnapconstantsv1.IdadmittedSplitSrbrelease, error) {

	msg := &xnapconstantsv1.IdadmittedSplitSrbrelease{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDadmittedSplitSrbrelease() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDAMfRegionInformation(value int32) (*xnapconstantsv1.IdAMfRegionInformation, error) {

	msg := &xnapconstantsv1.IdAMfRegionInformation{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDAMfRegionInformation() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDAssistanceDataForRanpaging(value int32) (*xnapconstantsv1.IdAssistanceDataForRanpaging, error) {

	msg := &xnapconstantsv1.IdAssistanceDataForRanpaging{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDAssistanceDataForRanpaging() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDBearersSubjectToCounterCheck(value int32) (*xnapconstantsv1.IdBearersSubjectToCounterCheck, error) {

	msg := &xnapconstantsv1.IdBearersSubjectToCounterCheck{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDBearersSubjectToCounterCheck() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDCause(value int32) (*xnapconstantsv1.IdCause, error) {

	msg := &xnapconstantsv1.IdCause{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDCause() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDcellAssistanceInfoNR(value int32) (*xnapconstantsv1.IdcellAssistanceInfoNR, error) {

	msg := &xnapconstantsv1.IdcellAssistanceInfoNR{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDcellAssistanceInfoNR() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDConfigurationUpdateInitiatingNodeChoice(value int32) (*xnapconstantsv1.IdConfigurationUpdateInitiatingNodeChoice, error) {

	msg := &xnapconstantsv1.IdConfigurationUpdateInitiatingNodeChoice{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDConfigurationUpdateInitiatingNodeChoice() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDCriticalityDiagnostics(value int32) (*xnapconstantsv1.IdCriticalityDiagnostics, error) {

	msg := &xnapconstantsv1.IdCriticalityDiagnostics{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDCriticalityDiagnostics() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDXnUaddressInfoperPdusessionList(value int32) (*xnapconstantsv1.IdXnUaddressInfoperPdusessionList, error) {

	msg := &xnapconstantsv1.IdXnUaddressInfoperPdusessionList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDXnUaddressInfoperPdusessionList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDDRbsSubjectToStatusTransferList(value int32) (*xnapconstantsv1.IdDRbsSubjectToStatusTransferList, error) {

	msg := &xnapconstantsv1.IdDRbsSubjectToStatusTransferList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDDRbsSubjectToStatusTransferList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDExpectedUebehaviour(value int32) (*xnapconstantsv1.IdExpectedUebehaviour, error) {

	msg := &xnapconstantsv1.IdExpectedUebehaviour{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDExpectedUebehaviour() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDGlobalNgRAnnodeID(value int32) (*xnapconstantsv1.IdGlobalNgRAnnodeID, error) {

	msg := &xnapconstantsv1.IdGlobalNgRAnnodeID{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDGlobalNgRAnnodeID() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDGUami(value int32) (*xnapconstantsv1.IdGUami, error) {

	msg := &xnapconstantsv1.IdGUami{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDGUami() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDindexToRatFrequSelectionPriority(value int32) (*xnapconstantsv1.IdindexToRatFrequSelectionPriority, error) {

	msg := &xnapconstantsv1.IdindexToRatFrequSelectionPriority{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDindexToRatFrequSelectionPriority() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDinitiatingNodeTypeResourceCoordRequest(value int32) (*xnapconstantsv1.IdinitiatingNodeTypeResourceCoordRequest, error) {

	msg := &xnapconstantsv1.IdinitiatingNodeTypeResourceCoordRequest{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDinitiatingNodeTypeResourceCoordRequest() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDListofservedcellsEUTra(value int32) (*xnapconstantsv1.IdListofservedcellsEUTra, error) {

	msg := &xnapconstantsv1.IdListofservedcellsEUTra{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDListofservedcellsEUTra() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDListofservedcellsNR(value int32) (*xnapconstantsv1.IdListofservedcellsNR, error) {

	msg := &xnapconstantsv1.IdListofservedcellsNR{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDListofservedcellsNR() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDLocationReportingInformation(value int32) (*xnapconstantsv1.IdLocationReportingInformation, error) {

	msg := &xnapconstantsv1.IdLocationReportingInformation{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDLocationReportingInformation() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDMAcI(value int32) (*xnapconstantsv1.IdMAcI, error) {

	msg := &xnapconstantsv1.IdMAcI{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDMAcI() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDMaskedImeisv(value int32) (*xnapconstantsv1.IdMaskedImeisv, error) {

	msg := &xnapconstantsv1.IdMaskedImeisv{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDMaskedImeisv() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDMNGRAnnodeUexnApID(value int32) (*xnapconstantsv1.IdMNGRAnnodeUexnApid, error) {

	msg := &xnapconstantsv1.IdMNGRAnnodeUexnApid{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDMNGRAnnodeUexnApID() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDMNtoSNContainer(value int32) (*xnapconstantsv1.IdMNtoSNContainer, error) {

	msg := &xnapconstantsv1.IdMNtoSNContainer{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDMNtoSNContainer() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDMobilityRestrictionList(value int32) (*xnapconstantsv1.IdMobilityRestrictionList, error) {

	msg := &xnapconstantsv1.IdMobilityRestrictionList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDMobilityRestrictionList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIdnewNGRAnCellIdentity(value int32) (*xnapconstantsv1.IdnewNGRAnCellIdentity, error) {

	msg := &xnapconstantsv1.IdnewNGRAnCellIdentity{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIdnewNGRAnCellIdentity() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDnewNgRAnnodeUexnApID(value int32) (*xnapconstantsv1.IdnewNgRAnnodeUexnApid, error) {

	msg := &xnapconstantsv1.IdnewNgRAnnodeUexnApid{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDnewNgRAnnodeUexnApID() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDUEreportRrctransfer(value int32) (*xnapconstantsv1.IdUEreportRrctransfer, error) {

	msg := &xnapconstantsv1.IdUEreportRrctransfer{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDUEreportRrctransfer() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDoldNgRAnnodeUexnApID(value int32) (*xnapconstantsv1.IdoldNgRAnnodeUexnApid, error) {

	msg := &xnapconstantsv1.IdoldNgRAnnodeUexnApid{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDoldNgRAnnodeUexnApID() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDOldtoNewNgRAnnodeResumeContainer(value int32) (*xnapconstantsv1.IdOldtoNewNgRAnnodeResumeContainer, error) {

	msg := &xnapconstantsv1.IdOldtoNewNgRAnnodeResumeContainer{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDOldtoNewNgRAnnodeResumeContainer() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDPagingDrx(value int32) (*xnapconstantsv1.IdPagingDrx, error) {

	msg := &xnapconstantsv1.IdPagingDrx{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDPagingDrx() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDPCellID(value int32) (*xnapconstantsv1.IdPCellId, error) {

	msg := &xnapconstantsv1.IdPCellId{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDPCellID() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDPDcpchangeIndication(value int32) (*xnapconstantsv1.IdPDcpchangeIndication, error) {

	msg := &xnapconstantsv1.IdPDcpchangeIndication{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDPDcpchangeIndication() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDPDusessionAdmittedAddedAddReqAck(value int32) (*xnapconstantsv1.IdPDusessionAdmittedAddedAddReqAck, error) {

	msg := &xnapconstantsv1.IdPDusessionAdmittedAddedAddReqAck{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDPDusessionAdmittedAddedAddReqAck() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDPDusessionAdmittedModSnmodConfirm(value int32) (*xnapconstantsv1.IdPDusessionAdmittedModSnmodConfirm, error) {

	msg := &xnapconstantsv1.IdPDusessionAdmittedModSnmodConfirm{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDPDusessionAdmittedModSnmodConfirm() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDPDusessionAdmittedSNmodResponse(value int32) (*xnapconstantsv1.IdPDusessionAdmittedSNmodResponse, error) {

	msg := &xnapconstantsv1.IdPDusessionAdmittedSNmodResponse{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDPDusessionAdmittedSNmodResponse() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDPDusessionNotAdmittedAddReqAck(value int32) (*xnapconstantsv1.IdPDusessionNotAdmittedAddReqAck, error) {

	msg := &xnapconstantsv1.IdPDusessionNotAdmittedAddReqAck{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDPDusessionNotAdmittedAddReqAck() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDPDusessionNotAdmittedSNmodResponse(value int32) (*xnapconstantsv1.IdPDusessionNotAdmittedSNmodResponse, error) {

	msg := &xnapconstantsv1.IdPDusessionNotAdmittedSNmodResponse{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDPDusessionNotAdmittedSNmodResponse() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDPDusessionReleasedListRelConf(value int32) (*xnapconstantsv1.IdPDusessionReleasedListRelConf, error) {

	msg := &xnapconstantsv1.IdPDusessionReleasedListRelConf{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDPDusessionReleasedListRelConf() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDPDusessionReleasedSnmodConfirm(value int32) (*xnapconstantsv1.IdPDusessionReleasedSnmodConfirm, error) {

	msg := &xnapconstantsv1.IdPDusessionReleasedSnmodConfirm{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDPDusessionReleasedSnmodConfirm() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDPDusessionResourcesActivityNotifyList(value int32) (*xnapconstantsv1.IdPDusessionResourcesActivityNotifyList, error) {

	msg := &xnapconstantsv1.IdPDusessionResourcesActivityNotifyList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDPDusessionResourcesActivityNotifyList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDPDusessionResourcesAdmittedList(value int32) (*xnapconstantsv1.IdPDusessionResourcesAdmittedList, error) {

	msg := &xnapconstantsv1.IdPDusessionResourcesAdmittedList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDPDusessionResourcesAdmittedList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDPDusessionResourcesNotAdmittedList(value int32) (*xnapconstantsv1.IdPDusessionResourcesNotAdmittedList, error) {

	msg := &xnapconstantsv1.IdPDusessionResourcesNotAdmittedList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDPDusessionResourcesNotAdmittedList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDPDusessionResourcesNotifyList(value int32) (*xnapconstantsv1.IdPDusessionResourcesNotifyList, error) {

	msg := &xnapconstantsv1.IdPDusessionResourcesNotifyList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDPDusessionResourcesNotifyList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDPDusessionSNchangeConfirmList(value int32) (*xnapconstantsv1.IdPDusessionSNchangeConfirmList, error) {

	msg := &xnapconstantsv1.IdPDusessionSNchangeConfirmList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDPDusessionSNchangeConfirmList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDPDusessionSNchangeRequiredList(value int32) (*xnapconstantsv1.IdPDusessionSNchangeRequiredList, error) {

	msg := &xnapconstantsv1.IdPDusessionSNchangeRequiredList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDPDusessionSNchangeRequiredList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDPDusessionToBeAddedAddReq(value int32) (*xnapconstantsv1.IdPDusessionToBeAddedAddReq, error) {

	msg := &xnapconstantsv1.IdPDusessionToBeAddedAddReq{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDPDusessionToBeAddedAddReq() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDPDusessionToBeModifiedSnmodRequired(value int32) (*xnapconstantsv1.IdPDusessionToBeModifiedSnmodRequired, error) {

	msg := &xnapconstantsv1.IdPDusessionToBeModifiedSnmodRequired{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDPDusessionToBeModifiedSnmodRequired() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDPDusessionToBeReleasedListRelRqd(value int32) (*xnapconstantsv1.IdPDusessionToBeReleasedListRelRqd, error) {

	msg := &xnapconstantsv1.IdPDusessionToBeReleasedListRelRqd{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDPDusessionToBeReleasedListRelRqd() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDPDusessionToBeReleasedRelReq(value int32) (*xnapconstantsv1.IdPDusessionToBeReleasedRelReq, error) {

	msg := &xnapconstantsv1.IdPDusessionToBeReleasedRelReq{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDPDusessionToBeReleasedRelReq() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDPDusessionToBeReleasedSnmodRequired(value int32) (*xnapconstantsv1.IdPDusessionToBeReleasedSnmodRequired, error) {

	msg := &xnapconstantsv1.IdPDusessionToBeReleasedSnmodRequired{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDPDusessionToBeReleasedSnmodRequired() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDRAnpagingArea(value int32) (*xnapconstantsv1.IdRAnpagingArea, error) {

	msg := &xnapconstantsv1.IdRAnpagingArea{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDRAnpagingArea() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDPagingPriority(value int32) (*xnapconstantsv1.IdPagingPriority, error) {

	msg := &xnapconstantsv1.IdPagingPriority{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDPagingPriority() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDrequestedSplitSrb(value int32) (*xnapconstantsv1.IdrequestedSplitSrb, error) {

	msg := &xnapconstantsv1.IdrequestedSplitSrb{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDrequestedSplitSrb() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDrequestedSplitSrbrelease(value int32) (*xnapconstantsv1.IdrequestedSplitSrbrelease, error) {

	msg := &xnapconstantsv1.IdrequestedSplitSrbrelease{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDrequestedSplitSrbrelease() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDResetRequestTypeInfo(value int32) (*xnapconstantsv1.IdResetRequestTypeInfo, error) {

	msg := &xnapconstantsv1.IdResetRequestTypeInfo{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDResetRequestTypeInfo() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDResetResponseTypeInfo(value int32) (*xnapconstantsv1.IdResetResponseTypeInfo, error) {

	msg := &xnapconstantsv1.IdResetResponseTypeInfo{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDResetResponseTypeInfo() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDRespondingNodeTypeConfigUpdateAck(value int32) (*xnapconstantsv1.IdRespondingNodeTypeConfigUpdateAck, error) {

	msg := &xnapconstantsv1.IdRespondingNodeTypeConfigUpdateAck{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDRespondingNodeTypeConfigUpdateAck() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDrespondingNodeTypeResourceCoordResponse(value int32) (*xnapconstantsv1.IdrespondingNodeTypeResourceCoordResponse, error) {

	msg := &xnapconstantsv1.IdrespondingNodeTypeResourceCoordResponse{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDrespondingNodeTypeResourceCoordResponse() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDResponseInfoReconfCompl(value int32) (*xnapconstantsv1.IdResponseInfoReconfCompl, error) {

	msg := &xnapconstantsv1.IdResponseInfoReconfCompl{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDResponseInfoReconfCompl() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDRRcconfigIndication(value int32) (*xnapconstantsv1.IdRRcconfigIndication, error) {

	msg := &xnapconstantsv1.IdRRcconfigIndication{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDRRcconfigIndication() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDRRcresumeCause(value int32) (*xnapconstantsv1.IdRRcresumeCause, error) {

	msg := &xnapconstantsv1.IdRRcresumeCause{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDRRcresumeCause() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDSCgconfigurationQuery(value int32) (*xnapconstantsv1.IdSCgconfigurationQuery, error) {

	msg := &xnapconstantsv1.IdSCgconfigurationQuery{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDSCgconfigurationQuery() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDselectedPlmn(value int32) (*xnapconstantsv1.IdselectedPlmn, error) {

	msg := &xnapconstantsv1.IdselectedPlmn{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDselectedPlmn() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDServedCellsToActivate(value int32) (*xnapconstantsv1.IdServedCellsToActivate, error) {

	msg := &xnapconstantsv1.IdServedCellsToActivate{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDServedCellsToActivate() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDservedCellsToUpdateEUTra(value int32) (*xnapconstantsv1.IdservedCellsToUpdateEUTra, error) {

	msg := &xnapconstantsv1.IdservedCellsToUpdateEUTra{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDservedCellsToUpdateEUTra() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDServedCellsToUpdateInitiatingNodeChoice(value int32) (*xnapconstantsv1.IdServedCellsToUpdateInitiatingNodeChoice, error) {

	msg := &xnapconstantsv1.IdServedCellsToUpdateInitiatingNodeChoice{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDServedCellsToUpdateInitiatingNodeChoice() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDservedCellsToUpdateNR(value int32) (*xnapconstantsv1.IdservedCellsToUpdateNR, error) {

	msg := &xnapconstantsv1.IdservedCellsToUpdateNR{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDservedCellsToUpdateNR() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDsngRAnnodeSecurityKey(value int32) (*xnapconstantsv1.IdsngRAnnodeSecurityKey, error) {

	msg := &xnapconstantsv1.IdsngRAnnodeSecurityKey{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDsngRAnnodeSecurityKey() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDSNGRAnnodeUeAMbr(value int32) (*xnapconstantsv1.IdSNGRAnnodeUeAMbr, error) {

	msg := &xnapconstantsv1.IdSNGRAnnodeUeAMbr{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDSNGRAnnodeUeAMbr() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDSNGRAnnodeUexnApID(value int32) (*xnapconstantsv1.IdSNGRAnnodeUexnApid, error) {

	msg := &xnapconstantsv1.IdSNGRAnnodeUexnApid{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDSNGRAnnodeUexnApID() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDSNtoMNContainer(value int32) (*xnapconstantsv1.IdSNtoMNContainer, error) {

	msg := &xnapconstantsv1.IdSNtoMNContainer{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDSNtoMNContainer() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDsourceNgRAnnodeUexnApID(value int32) (*xnapconstantsv1.IdsourceNgRAnnodeUexnApid, error) {

	msg := &xnapconstantsv1.IdsourceNgRAnnodeUexnApid{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDsourceNgRAnnodeUexnApID() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDSplitSrbRRctransfer(value int32) (*xnapconstantsv1.IdSplitSrbRRctransfer, error) {

	msg := &xnapconstantsv1.IdSplitSrbRRctransfer{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDSplitSrbRRctransfer() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDTAisupportlist(value int32) (*xnapconstantsv1.IdTAisupportlist, error) {

	msg := &xnapconstantsv1.IdTAisupportlist{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDTAisupportlist() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDTimeToWait(value int32) (*xnapconstantsv1.IdTimeToWait, error) {

	msg := &xnapconstantsv1.IdTimeToWait{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDTimeToWait() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDTarget2SourceNgRAnnodeTranspContainer(value int32) (*xnapconstantsv1.IdTarget2SourceNgRAnnodeTranspContainer, error) {

	msg := &xnapconstantsv1.IdTarget2SourceNgRAnnodeTranspContainer{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDTarget2SourceNgRAnnodeTranspContainer() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDtargetCellGlobalID(value int32) (*xnapconstantsv1.IdtargetCellGlobalId, error) {

	msg := &xnapconstantsv1.IdtargetCellGlobalId{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDtargetCellGlobalID() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDtargetNgRAnnodeUexnApID(value int32) (*xnapconstantsv1.IdtargetNgRAnnodeUexnApid, error) {

	msg := &xnapconstantsv1.IdtargetNgRAnnodeUexnApid{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDtargetNgRAnnodeUexnApID() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDtargetSNGRAnnodeID(value int32) (*xnapconstantsv1.IdtargetSNGRAnnodeId, error) {

	msg := &xnapconstantsv1.IdtargetSNGRAnnodeId{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDtargetSNGRAnnodeID() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDTraceActivation(value int32) (*xnapconstantsv1.IdTraceActivation, error) {

	msg := &xnapconstantsv1.IdTraceActivation{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDTraceActivation() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDUEcontextID(value int32) (*xnapconstantsv1.IdUEcontextId, error) {

	msg := &xnapconstantsv1.IdUEcontextId{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDUEcontextID() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDUEcontextInfoHorequest(value int32) (*xnapconstantsv1.IdUEcontextInfoHorequest, error) {

	msg := &xnapconstantsv1.IdUEcontextInfoHorequest{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDUEcontextInfoHorequest() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDUEcontextInfoRetrUectxtResp(value int32) (*xnapconstantsv1.IdUEcontextInfoRetrUectxtResp, error) {

	msg := &xnapconstantsv1.IdUEcontextInfoRetrUectxtResp{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDUEcontextInfoRetrUectxtResp() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDUEcontextInfoSNmodRequest(value int32) (*xnapconstantsv1.IdUEcontextInfoSNmodRequest, error) {

	msg := &xnapconstantsv1.IdUEcontextInfoSNmodRequest{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDUEcontextInfoSNmodRequest() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDUEcontextKeptIndicator(value int32) (*xnapconstantsv1.IdUEcontextKeptIndicator, error) {

	msg := &xnapconstantsv1.IdUEcontextKeptIndicator{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDUEcontextKeptIndicator() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDUEcontextRefAtSnHOrequest(value int32) (*xnapconstantsv1.IdUEcontextRefAtSnHOrequest, error) {

	msg := &xnapconstantsv1.IdUEcontextRefAtSnHOrequest{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDUEcontextRefAtSnHOrequest() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDUEhistoryInformation(value int32) (*xnapconstantsv1.IdUEhistoryInformation, error) {

	msg := &xnapconstantsv1.IdUEhistoryInformation{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDUEhistoryInformation() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDUEidentityIndexValue(value int32) (*xnapconstantsv1.IdUEidentityIndexValue, error) {

	msg := &xnapconstantsv1.IdUEidentityIndexValue{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDUEidentityIndexValue() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDUEranpagingIdentity(value int32) (*xnapconstantsv1.IdUEranpagingIdentity, error) {

	msg := &xnapconstantsv1.IdUEranpagingIdentity{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDUEranpagingIdentity() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDUEsecurityCapabilities(value int32) (*xnapconstantsv1.IdUEsecurityCapabilities, error) {

	msg := &xnapconstantsv1.IdUEsecurityCapabilities{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDUEsecurityCapabilities() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDUserPlaneTrafficActivityReport(value int32) (*xnapconstantsv1.IdUserPlaneTrafficActivityReport, error) {

	msg := &xnapconstantsv1.IdUserPlaneTrafficActivityReport{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDUserPlaneTrafficActivityReport() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDXnRemovalThreshold(value int32) (*xnapconstantsv1.IdXnRemovalThreshold, error) {

	msg := &xnapconstantsv1.IdXnRemovalThreshold{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDXnRemovalThreshold() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDDesiredActNotificationLevel(value int32) (*xnapconstantsv1.IdDesiredActNotificationLevel, error) {

	msg := &xnapconstantsv1.IdDesiredActNotificationLevel{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDDesiredActNotificationLevel() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDAvailableDrbIDs(value int32) (*xnapconstantsv1.IdAvailableDrbids, error) {

	msg := &xnapconstantsv1.IdAvailableDrbids{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDAvailableDrbIDs() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDAdditionalDrbIDs(value int32) (*xnapconstantsv1.IdAdditionalDrbids, error) {

	msg := &xnapconstantsv1.IdAdditionalDrbids{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDAdditionalDrbIDs() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDSpareDrbIDs(value int32) (*xnapconstantsv1.IdSpareDrbids, error) {

	msg := &xnapconstantsv1.IdSpareDrbids{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDSpareDrbIDs() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDRequiredNumberOfDrbIDs(value int32) (*xnapconstantsv1.IdRequiredNumberOfDrbids, error) {

	msg := &xnapconstantsv1.IdRequiredNumberOfDrbids{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDRequiredNumberOfDrbIDs() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDTNlaToAddList(value int32) (*xnapconstantsv1.IdTNlaToAddList, error) {

	msg := &xnapconstantsv1.IdTNlaToAddList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDTNlaToAddList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDTNlaToUpdateList(value int32) (*xnapconstantsv1.IdTNlaToUpdateList, error) {

	msg := &xnapconstantsv1.IdTNlaToUpdateList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDTNlaToUpdateList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDTNlaToRemoveList(value int32) (*xnapconstantsv1.IdTNlaToRemoveList, error) {

	msg := &xnapconstantsv1.IdTNlaToRemoveList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDTNlaToRemoveList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDTNlaSetupList(value int32) (*xnapconstantsv1.IdTNlaSetupList, error) {

	msg := &xnapconstantsv1.IdTNlaSetupList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDTNlaSetupList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDTNlaFailedToSetupList(value int32) (*xnapconstantsv1.IdTNlaFailedToSetupList, error) {

	msg := &xnapconstantsv1.IdTNlaFailedToSetupList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDTNlaFailedToSetupList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDPDusessionToBeReleasedRelReqAck(value int32) (*xnapconstantsv1.IdPDusessionToBeReleasedRelReqAck, error) {

	msg := &xnapconstantsv1.IdPDusessionToBeReleasedRelReqAck{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDPDusessionToBeReleasedRelReqAck() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDSNGRAnnodeMaxIpdataRateUL(value int32) (*xnapconstantsv1.IdSNGRAnnodeMaxIpdataRateUL, error) {

	msg := &xnapconstantsv1.IdSNGRAnnodeMaxIpdataRateUL{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDSNGRAnnodeMaxIpdataRateUL() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDPDusessionResourceSecondaryRatusageList(value int32) (*xnapconstantsv1.IdPDusessionResourceSecondaryRatusageList, error) {

	msg := &xnapconstantsv1.IdPDusessionResourceSecondaryRatusageList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDPDusessionResourceSecondaryRatusageList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDAdditionalULNGUTNlatUpfList(value int32) (*xnapconstantsv1.IdAdditionalULNGUTNlatUpfList, error) {

	msg := &xnapconstantsv1.IdAdditionalULNGUTNlatUpfList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDAdditionalULNGUTNlatUpfList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDSecondarydataForwardingInfoFromTargetList(value int32) (*xnapconstantsv1.IdSecondarydataForwardingInfoFromTargetList, error) {

	msg := &xnapconstantsv1.IdSecondarydataForwardingInfoFromTargetList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDSecondarydataForwardingInfoFromTargetList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDLocationInformationSnreporting(value int32) (*xnapconstantsv1.IdLocationInformationSnreporting, error) {

	msg := &xnapconstantsv1.IdLocationInformationSnreporting{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDLocationInformationSnreporting() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDLocationInformationSn(value int32) (*xnapconstantsv1.IdLocationInformationSn, error) {

	msg := &xnapconstantsv1.IdLocationInformationSn{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDLocationInformationSn() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDLastEUTranplmnidentity(value int32) (*xnapconstantsv1.IdLastEUTranplmnidentity, error) {

	msg := &xnapconstantsv1.IdLastEUTranplmnidentity{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDLastEUTranplmnidentity() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDSNGRAnnodeMaxIpdataRateDL(value int32) (*xnapconstantsv1.IdSNGRAnnodeMaxIpdataRateDL, error) {

	msg := &xnapconstantsv1.IdSNGRAnnodeMaxIpdataRateDL{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDSNGRAnnodeMaxIpdataRateDL() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDMaxIprateDL(value int32) (*xnapconstantsv1.IdMaxIprateDL, error) {

	msg := &xnapconstantsv1.IdMaxIprateDL{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDMaxIprateDL() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDSecurityResult(value int32) (*xnapconstantsv1.IdSecurityResult, error) {

	msg := &xnapconstantsv1.IdSecurityResult{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDSecurityResult() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDSNSsai(value int32) (*xnapconstantsv1.IdSNSsai, error) {

	msg := &xnapconstantsv1.IdSNSsai{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDSNSsai() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDMRDCResourceCoordinationInfo(value int32) (*xnapconstantsv1.IdMRDCResourceCoordinationInfo, error) {

	msg := &xnapconstantsv1.IdMRDCResourceCoordinationInfo{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDMRDCResourceCoordinationInfo() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDAMfRegionInformationToAdd(value int32) (*xnapconstantsv1.IdAMfRegionInformationToAdd, error) {

	msg := &xnapconstantsv1.IdAMfRegionInformationToAdd{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDAMfRegionInformationToAdd() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDAMfRegionInformationToDelete(value int32) (*xnapconstantsv1.IdAMfRegionInformationToDelete, error) {

	msg := &xnapconstantsv1.IdAMfRegionInformationToDelete{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDAMfRegionInformationToDelete() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDOldQoSflowMapULendmarkerexpected(value int32) (*xnapconstantsv1.IdOldQoSflowMapULendmarkerexpected, error) {

	msg := &xnapconstantsv1.IdOldQoSflowMapULendmarkerexpected{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDOldQoSflowMapULendmarkerexpected() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDRAnpagingFailure(value int32) (*xnapconstantsv1.IdRAnpagingFailure, error) {

	msg := &xnapconstantsv1.IdRAnpagingFailure{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDRAnpagingFailure() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDUEradioCapabilityForPaging(value int32) (*xnapconstantsv1.IdUEradioCapabilityForPaging, error) {

	msg := &xnapconstantsv1.IdUEradioCapabilityForPaging{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDUEradioCapabilityForPaging() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDPDusessionDataForwardingSNmodResponse(value int32) (*xnapconstantsv1.IdPDusessionDataForwardingSNmodResponse, error) {

	msg := &xnapconstantsv1.IdPDusessionDataForwardingSNmodResponse{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDPDusessionDataForwardingSNmodResponse() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDDRbsNotAdmittedSetupModifyList(value int32) (*xnapconstantsv1.IdDRbsNotAdmittedSetupModifyList, error) {

	msg := &xnapconstantsv1.IdDRbsNotAdmittedSetupModifyList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDDRbsNotAdmittedSetupModifyList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDSecondaryMNXnUTNlinfoatM(value int32) (*xnapconstantsv1.IdSecondaryMNXnUTNlinfoatM, error) {

	msg := &xnapconstantsv1.IdSecondaryMNXnUTNlinfoatM{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDSecondaryMNXnUTNlinfoatM() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDNEDCTDmPattern(value int32) (*xnapconstantsv1.IdNEDCTDmPattern, error) {

	msg := &xnapconstantsv1.IdNEDCTDmPattern{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDNEDCTDmPattern() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDPDusessionCommonNetworkInstance(value int32) (*xnapconstantsv1.IdPDusessionCommonNetworkInstance, error) {

	msg := &xnapconstantsv1.IdPDusessionCommonNetworkInstance{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDPDusessionCommonNetworkInstance() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDBPlmnIDInfoEUtra(value int32) (*xnapconstantsv1.IdBPlmnIDInfoEUtra, error) {

	msg := &xnapconstantsv1.IdBPlmnIDInfoEUtra{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDBPlmnIDInfoEUtra() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDBPlmnIDInfoNR(value int32) (*xnapconstantsv1.IdBPlmnIDInfoNR, error) {

	msg := &xnapconstantsv1.IdBPlmnIDInfoNR{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDBPlmnIDInfoNR() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDInterfaceInstanceIndication(value int32) (*xnapconstantsv1.IdInterfaceInstanceIndication, error) {

	msg := &xnapconstantsv1.IdInterfaceInstanceIndication{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDInterfaceInstanceIndication() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDSNGRAnnodeAdditionTriggerInd(value int32) (*xnapconstantsv1.IdSNGRAnnodeAdditionTriggerInd, error) {

	msg := &xnapconstantsv1.IdSNGRAnnodeAdditionTriggerInd{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDSNGRAnnodeAdditionTriggerInd() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDDefaultDrbAllowed(value int32) (*xnapconstantsv1.IdDefaultDrbAllowed, error) {

	msg := &xnapconstantsv1.IdDefaultDrbAllowed{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDDefaultDrbAllowed() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDDRbIDstakenintouse(value int32) (*xnapconstantsv1.IdDRbIDstakenintouse, error) {

	msg := &xnapconstantsv1.IdDRbIDstakenintouse{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDDRbIDstakenintouse() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDSplitSessionIndicator(value int32) (*xnapconstantsv1.IdSplitSessionIndicator, error) {

	msg := &xnapconstantsv1.IdSplitSessionIndicator{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDSplitSessionIndicator() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDCNtypeRestrictionsForEquivalent(value int32) (*xnapconstantsv1.IdCNtypeRestrictionsForEquivalent, error) {

	msg := &xnapconstantsv1.IdCNtypeRestrictionsForEquivalent{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDCNtypeRestrictionsForEquivalent() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDCNtypeRestrictionsForServing(value int32) (*xnapconstantsv1.IdCNtypeRestrictionsForServing, error) {

	msg := &xnapconstantsv1.IdCNtypeRestrictionsForServing{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDCNtypeRestrictionsForServing() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDDRbstransferredtoMN(value int32) (*xnapconstantsv1.IdDRbstransferredtoMN, error) {

	msg := &xnapconstantsv1.IdDRbstransferredtoMN{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDDRbstransferredtoMN() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDULforwardingProposal(value int32) (*xnapconstantsv1.IdULforwardingProposal, error) {

	msg := &xnapconstantsv1.IdULforwardingProposal{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDULforwardingProposal() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDEndpointIpaddressAndPort(value int32) (*xnapconstantsv1.IdEndpointIpaddressAndPort, error) {

	msg := &xnapconstantsv1.IdEndpointIpaddressAndPort{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDEndpointIpaddressAndPort() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDIntendedTddDLULconfigurationNR(value int32) (*xnapconstantsv1.IdIntendedTddDLULconfigurationNR, error) {

	msg := &xnapconstantsv1.IdIntendedTddDLULconfigurationNR{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDIntendedTddDLULconfigurationNR() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDTNlconfigurationInfo(value int32) (*xnapconstantsv1.IdTNlconfigurationInfo, error) {

	msg := &xnapconstantsv1.IdTNlconfigurationInfo{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDTNlconfigurationInfo() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDPartialListIndicatorNR(value int32) (*xnapconstantsv1.IdPartialListIndicatorNR, error) {

	msg := &xnapconstantsv1.IdPartialListIndicatorNR{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDPartialListIndicatorNR() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDMessageOversizeNotification(value int32) (*xnapconstantsv1.IdMessageOversizeNotification, error) {

	msg := &xnapconstantsv1.IdMessageOversizeNotification{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDMessageOversizeNotification() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDCellAndCapacityAssistanceInfoNR(value int32) (*xnapconstantsv1.IdCellAndCapacityAssistanceInfoNR, error) {

	msg := &xnapconstantsv1.IdCellAndCapacityAssistanceInfoNR{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDCellAndCapacityAssistanceInfoNR() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDNGRAntraceID(value int32) (*xnapconstantsv1.IdNGRAntraceId, error) {

	msg := &xnapconstantsv1.IdNGRAntraceId{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDNGRAntraceID() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDNonGbrresourcesOffered(value int32) (*xnapconstantsv1.IdNonGbrresourcesOffered, error) {

	msg := &xnapconstantsv1.IdNonGbrresourcesOffered{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDNonGbrresourcesOffered() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDFastMcgrecoveryRrctransferSNtoMN(value int32) (*xnapconstantsv1.IdFastMcgrecoveryRrctransferSNtoMN, error) {

	msg := &xnapconstantsv1.IdFastMcgrecoveryRrctransferSNtoMN{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDFastMcgrecoveryRrctransferSNtoMN() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDRequestedFastMcgrecoveryViaSrb3(value int32) (*xnapconstantsv1.IdRequestedFastMcgrecoveryViaSrb3, error) {

	msg := &xnapconstantsv1.IdRequestedFastMcgrecoveryViaSrb3{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDRequestedFastMcgrecoveryViaSrb3() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDAvailableFastMcgrecoveryViaSrb3(value int32) (*xnapconstantsv1.IdAvailableFastMcgrecoveryViaSrb3, error) {

	msg := &xnapconstantsv1.IdAvailableFastMcgrecoveryViaSrb3{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDAvailableFastMcgrecoveryViaSrb3() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDRequestedFastMcgrecoveryViaSrb3Release(value int32) (*xnapconstantsv1.IdRequestedFastMcgrecoveryViaSrb3Release, error) {

	msg := &xnapconstantsv1.IdRequestedFastMcgrecoveryViaSrb3Release{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDRequestedFastMcgrecoveryViaSrb3Release() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDReleaseFastMcgrecoveryViaSrb3(value int32) (*xnapconstantsv1.IdReleaseFastMcgrecoveryViaSrb3, error) {

	msg := &xnapconstantsv1.IdReleaseFastMcgrecoveryViaSrb3{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDReleaseFastMcgrecoveryViaSrb3() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDFastMcgrecoveryRrctransferMNtoSN(value int32) (*xnapconstantsv1.IdFastMcgrecoveryRrctransferMNtoSN, error) {

	msg := &xnapconstantsv1.IdFastMcgrecoveryRrctransferMNtoSN{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDFastMcgrecoveryRrctransferMNtoSN() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDExtendedRatrestrictionInformation(value int32) (*xnapconstantsv1.IdExtendedRatrestrictionInformation, error) {

	msg := &xnapconstantsv1.IdExtendedRatrestrictionInformation{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDExtendedRatrestrictionInformation() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDQoSmonitoringRequest(value int32) (*xnapconstantsv1.IdQoSmonitoringRequest, error) {

	msg := &xnapconstantsv1.IdQoSmonitoringRequest{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDQoSmonitoringRequest() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDFiveGcmobilityRestrictionListContainer(value int32) (*xnapconstantsv1.IdFiveGcmobilityRestrictionListContainer, error) {

	msg := &xnapconstantsv1.IdFiveGcmobilityRestrictionListContainer{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDFiveGcmobilityRestrictionListContainer() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDPartialListIndicatorEUtra(value int32) (*xnapconstantsv1.IdPartialListIndicatorEUtra, error) {

	msg := &xnapconstantsv1.IdPartialListIndicatorEUtra{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDPartialListIndicatorEUtra() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDCellAndCapacityAssistanceInfoEUtra(value int32) (*xnapconstantsv1.IdCellAndCapacityAssistanceInfoEUtra, error) {

	msg := &xnapconstantsv1.IdCellAndCapacityAssistanceInfoEUtra{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDCellAndCapacityAssistanceInfoEUtra() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDCHoinformationReq(value int32) (*xnapconstantsv1.IdCHoinformationReq, error) {

	msg := &xnapconstantsv1.IdCHoinformationReq{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDCHoinformationReq() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDCHoinformationAck(value int32) (*xnapconstantsv1.IdCHoinformationAck, error) {

	msg := &xnapconstantsv1.IdCHoinformationAck{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDCHoinformationAck() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDtargetCellsToCancel(value int32) (*xnapconstantsv1.IdtargetCellsToCancel, error) {

	msg := &xnapconstantsv1.IdtargetCellsToCancel{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDtargetCellsToCancel() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDrequestedTargetCellGlobalID(value int32) (*xnapconstantsv1.IdrequestedTargetCellGlobalId, error) {

	msg := &xnapconstantsv1.IdrequestedTargetCellGlobalId{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDrequestedTargetCellGlobalID() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDprocedureStage(value int32) (*xnapconstantsv1.IdprocedureStage, error) {

	msg := &xnapconstantsv1.IdprocedureStage{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDprocedureStage() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDDApsrequestInfo(value int32) (*xnapconstantsv1.IdDApsrequestInfo, error) {

	msg := &xnapconstantsv1.IdDApsrequestInfo{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDDApsrequestInfo() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDDApsresponseInfoList(value int32) (*xnapconstantsv1.IdDApsresponseInfoList, error) {

	msg := &xnapconstantsv1.IdDApsresponseInfoList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDDApsresponseInfoList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDCHoMRdcIndicator(value int32) (*xnapconstantsv1.IdCHoMRdcIndicator, error) {

	msg := &xnapconstantsv1.IdCHoMRdcIndicator{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDCHoMRdcIndicator() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDOffsetOfNbiotChannelNumberToDlEArfcn(value int32) (*xnapconstantsv1.IdOffsetOfNbiotChannelNumberToDlEArfcn, error) {

	msg := &xnapconstantsv1.IdOffsetOfNbiotChannelNumberToDlEArfcn{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDOffsetOfNbiotChannelNumberToDlEArfcn() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDOffsetOfNbiotChannelNumberToUlEArfcn(value int32) (*xnapconstantsv1.IdOffsetOfNbiotChannelNumberToUlEArfcn, error) {

	msg := &xnapconstantsv1.IdOffsetOfNbiotChannelNumberToUlEArfcn{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDOffsetOfNbiotChannelNumberToUlEArfcn() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDNBioTULDLAlignmentOffset(value int32) (*xnapconstantsv1.IdNBioTULDLAlignmentOffset, error) {

	msg := &xnapconstantsv1.IdNBioTULDLAlignmentOffset{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDNBioTULDLAlignmentOffset() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDLTev2XservicesAuthorized(value int32) (*xnapconstantsv1.IdLTev2XservicesAuthorized, error) {

	msg := &xnapconstantsv1.IdLTev2XservicesAuthorized{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDLTev2XservicesAuthorized() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDNRv2XservicesAuthorized(value int32) (*xnapconstantsv1.IdNRv2XservicesAuthorized, error) {

	msg := &xnapconstantsv1.IdNRv2XservicesAuthorized{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDNRv2XservicesAuthorized() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDLTeuesIDelinkAggregateMaximumBitRate(value int32) (*xnapconstantsv1.IdLTeuesidelinkAggregateMaximumBitRate, error) {

	msg := &xnapconstantsv1.IdLTeuesidelinkAggregateMaximumBitRate{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDLTeuesIDelinkAggregateMaximumBitRate() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDNRuesIDelinkAggregateMaximumBitRate(value int32) (*xnapconstantsv1.IdNRuesidelinkAggregateMaximumBitRate, error) {

	msg := &xnapconstantsv1.IdNRuesidelinkAggregateMaximumBitRate{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDNRuesIDelinkAggregateMaximumBitRate() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDPC5QoSparameters(value int32) (*xnapconstantsv1.IdPC5QoSparameters, error) {

	msg := &xnapconstantsv1.IdPC5QoSparameters{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDPC5QoSparameters() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDAlternativeQoSparaSetList(value int32) (*xnapconstantsv1.IdAlternativeQoSparaSetList, error) {

	msg := &xnapconstantsv1.IdAlternativeQoSparaSetList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDAlternativeQoSparaSetList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDCurrentQoSparaSetIndex(value int32) (*xnapconstantsv1.IdCurrentQoSparaSetIndex, error) {

	msg := &xnapconstantsv1.IdCurrentQoSparaSetIndex{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDCurrentQoSparaSetIndex() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDMobilityInformation(value int32) (*xnapconstantsv1.IdMobilityInformation, error) {

	msg := &xnapconstantsv1.IdMobilityInformation{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDMobilityInformation() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDInitiatingConditionFailureIndication(value int32) (*xnapconstantsv1.IdInitiatingConditionFailureIndication, error) {

	msg := &xnapconstantsv1.IdInitiatingConditionFailureIndication{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDInitiatingConditionFailureIndication() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDUEhistoryInformationFromTheUe(value int32) (*xnapconstantsv1.IdUEhistoryInformationFromTheUe, error) {

	msg := &xnapconstantsv1.IdUEhistoryInformationFromTheUe{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDUEhistoryInformationFromTheUe() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDHandoverReportType(value int32) (*xnapconstantsv1.IdHandoverReportType, error) {

	msg := &xnapconstantsv1.IdHandoverReportType{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDHandoverReportType() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDHandoverCause(value int32) (*xnapconstantsv1.IdHandoverCause, error) {

	msg := &xnapconstantsv1.IdHandoverCause{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDHandoverCause() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDSourceCellCgi(value int32) (*xnapconstantsv1.IdSourceCellCgi, error) {

	msg := &xnapconstantsv1.IdSourceCellCgi{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDSourceCellCgi() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDTargetCellCgi(value int32) (*xnapconstantsv1.IdTargetCellCgi, error) {

	msg := &xnapconstantsv1.IdTargetCellCgi{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDTargetCellCgi() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDReEstablishmentCellCgi(value int32) (*xnapconstantsv1.IdReEstablishmentCellCgi, error) {

	msg := &xnapconstantsv1.IdReEstablishmentCellCgi{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDReEstablishmentCellCgi() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDTargetCellinEutran(value int32) (*xnapconstantsv1.IdTargetCellinEutran, error) {

	msg := &xnapconstantsv1.IdTargetCellinEutran{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDTargetCellinEutran() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDSourceCellCrnti(value int32) (*xnapconstantsv1.IdSourceCellCrnti, error) {

	msg := &xnapconstantsv1.IdSourceCellCrnti{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDSourceCellCrnti() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDUErlfreportContainer(value int32) (*xnapconstantsv1.IdUErlfreportContainer, error) {

	msg := &xnapconstantsv1.IdUErlfreportContainer{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDUErlfreportContainer() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDNGranNode1MeasurementID(value int32) (*xnapconstantsv1.IdNGranNode1MeasurementID, error) {

	msg := &xnapconstantsv1.IdNGranNode1MeasurementID{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDNGranNode1MeasurementID() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDNGranNode2MeasurementID(value int32) (*xnapconstantsv1.IdNGranNode2MeasurementID, error) {

	msg := &xnapconstantsv1.IdNGranNode2MeasurementID{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDNGranNode2MeasurementID() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDRegistrationRequest(value int32) (*xnapconstantsv1.IdRegistrationRequest, error) {

	msg := &xnapconstantsv1.IdRegistrationRequest{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDRegistrationRequest() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDReportCharacteristics(value int32) (*xnapconstantsv1.IdReportCharacteristics, error) {

	msg := &xnapconstantsv1.IdReportCharacteristics{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDReportCharacteristics() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDCellToReport(value int32) (*xnapconstantsv1.IdCellToReport, error) {

	msg := &xnapconstantsv1.IdCellToReport{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDCellToReport() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDReportingPeriodicity(value int32) (*xnapconstantsv1.IdReportingPeriodicity, error) {

	msg := &xnapconstantsv1.IdReportingPeriodicity{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDReportingPeriodicity() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDCellMeasurementResult(value int32) (*xnapconstantsv1.IdCellMeasurementResult, error) {

	msg := &xnapconstantsv1.IdCellMeasurementResult{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDCellMeasurementResult() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDNGRAnnode1CellID(value int32) (*xnapconstantsv1.IdNGRAnnode1CellId, error) {

	msg := &xnapconstantsv1.IdNGRAnnode1CellId{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDNGRAnnode1CellID() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDNGRAnnode2CellID(value int32) (*xnapconstantsv1.IdNGRAnnode2CellId, error) {

	msg := &xnapconstantsv1.IdNGRAnnode2CellId{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDNGRAnnode2CellID() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDNGRAnnode1MobilityParameters(value int32) (*xnapconstantsv1.IdNGRAnnode1MobilityParameters, error) {

	msg := &xnapconstantsv1.IdNGRAnnode1MobilityParameters{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDNGRAnnode1MobilityParameters() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDNGRAnnode2ProposedMobilityParameters(value int32) (*xnapconstantsv1.IdNGRAnnode2ProposedMobilityParameters, error) {

	msg := &xnapconstantsv1.IdNGRAnnode2ProposedMobilityParameters{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDNGRAnnode2ProposedMobilityParameters() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDMobilityParametersModificationRange(value int32) (*xnapconstantsv1.IdMobilityParametersModificationRange, error) {

	msg := &xnapconstantsv1.IdMobilityParametersModificationRange{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDMobilityParametersModificationRange() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDTDduldlconfigurationCommonNr(value int32) (*xnapconstantsv1.IdTDduldlconfigurationCommonNr, error) {

	msg := &xnapconstantsv1.IdTDduldlconfigurationCommonNr{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDTDduldlconfigurationCommonNr() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDCarrierList(value int32) (*xnapconstantsv1.IdCarrierList, error) {

	msg := &xnapconstantsv1.IdCarrierList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDCarrierList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDULcarrierList(value int32) (*xnapconstantsv1.IdULcarrierList, error) {

	msg := &xnapconstantsv1.IdULcarrierList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDULcarrierList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDFrequencyShift7p5khz(value int32) (*xnapconstantsv1.IdFrequencyShift7P5Khz, error) {

	msg := &xnapconstantsv1.IdFrequencyShift7P5Khz{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDFrequencyShift7p5khz() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDSSbPositionsInBurst(value int32) (*xnapconstantsv1.IdSSbPositionsInBurst, error) {

	msg := &xnapconstantsv1.IdSSbPositionsInBurst{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDSSbPositionsInBurst() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDNRcellPrachconfig(value int32) (*xnapconstantsv1.IdNRcellPrachconfig, error) {

	msg := &xnapconstantsv1.IdNRcellPrachconfig{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDNRcellPrachconfig() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDRAchreportInformation(value int32) (*xnapconstantsv1.IdRAchreportInformation, error) {

	msg := &xnapconstantsv1.IdRAchreportInformation{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDRAchreportInformation() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDIAbnodeIndication(value int32) (*xnapconstantsv1.IdIAbnodeIndication, error) {

	msg := &xnapconstantsv1.IdIAbnodeIndication{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDIAbnodeIndication() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDRedundantULNGUTNlatUpf(value int32) (*xnapconstantsv1.IdRedundantULNGUTNlatUpf, error) {

	msg := &xnapconstantsv1.IdRedundantULNGUTNlatUpf{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDRedundantULNGUTNlatUpf() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDCNpacketDelayBudgetDownlink(value int32) (*xnapconstantsv1.IdCNpacketDelayBudgetDownlink, error) {

	msg := &xnapconstantsv1.IdCNpacketDelayBudgetDownlink{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDCNpacketDelayBudgetDownlink() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDCNpacketDelayBudgetUplink(value int32) (*xnapconstantsv1.IdCNpacketDelayBudgetUplink, error) {

	msg := &xnapconstantsv1.IdCNpacketDelayBudgetUplink{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDCNpacketDelayBudgetUplink() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDAdditionalRedundantULNGUTNlatUpfList(value int32) (*xnapconstantsv1.IdAdditionalRedundantULNGUTNlatUpfList, error) {

	msg := &xnapconstantsv1.IdAdditionalRedundantULNGUTNlatUpfList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDAdditionalRedundantULNGUTNlatUpfList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDRedundantCommonNetworkInstance(value int32) (*xnapconstantsv1.IdRedundantCommonNetworkInstance, error) {

	msg := &xnapconstantsv1.IdRedundantCommonNetworkInstance{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDRedundantCommonNetworkInstance() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDTSctrafficCharacteristics(value int32) (*xnapconstantsv1.IdTSctrafficCharacteristics, error) {

	msg := &xnapconstantsv1.IdTSctrafficCharacteristics{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDTSctrafficCharacteristics() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDRedundantQoSflowIndicator(value int32) (*xnapconstantsv1.IdRedundantQoSflowIndicator, error) {

	msg := &xnapconstantsv1.IdRedundantQoSflowIndicator{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDRedundantQoSflowIndicator() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDRedundantDLNGUTNlatNgRAn(value int32) (*xnapconstantsv1.IdRedundantDLNGUTNlatNgRAn, error) {

	msg := &xnapconstantsv1.IdRedundantDLNGUTNlatNgRAn{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDRedundantDLNGUTNlatNgRAn() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDExtendedPacketDelayBudget(value int32) (*xnapconstantsv1.IdExtendedPacketDelayBudget, error) {

	msg := &xnapconstantsv1.IdExtendedPacketDelayBudget{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDExtendedPacketDelayBudget() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDAdditionalPDcpDuplicationTNlList(value int32) (*xnapconstantsv1.IdAdditionalPDcpDuplicationTNlList, error) {

	msg := &xnapconstantsv1.IdAdditionalPDcpDuplicationTNlList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDAdditionalPDcpDuplicationTNlList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDRedundantPdusessionInformation(value int32) (*xnapconstantsv1.IdRedundantPdusessionInformation, error) {

	msg := &xnapconstantsv1.IdRedundantPdusessionInformation{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDRedundantPdusessionInformation() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDUsedRsninformation(value int32) (*xnapconstantsv1.IdUsedRsninformation, error) {

	msg := &xnapconstantsv1.IdUsedRsninformation{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDUsedRsninformation() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDRLcduplicationInformation(value int32) (*xnapconstantsv1.IdRLcduplicationInformation, error) {

	msg := &xnapconstantsv1.IdRLcduplicationInformation{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDRLcduplicationInformation() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDNPnBroadcastInformation(value int32) (*xnapconstantsv1.IdNPnBroadcastInformation, error) {

	msg := &xnapconstantsv1.IdNPnBroadcastInformation{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDNPnBroadcastInformation() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDNPnpagingAssistanceInformation(value int32) (*xnapconstantsv1.IdNPnpagingAssistanceInformation, error) {

	msg := &xnapconstantsv1.IdNPnpagingAssistanceInformation{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDNPnpagingAssistanceInformation() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDNPnmobilityInformation(value int32) (*xnapconstantsv1.IdNPnmobilityInformation, error) {

	msg := &xnapconstantsv1.IdNPnmobilityInformation{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDNPnmobilityInformation() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDNPnSupport(value int32) (*xnapconstantsv1.IdNPnSupport, error) {

	msg := &xnapconstantsv1.IdNPnSupport{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDNPnSupport() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDMDtConfiguration(value int32) (*xnapconstantsv1.IdMDtConfiguration, error) {

	msg := &xnapconstantsv1.IdMDtConfiguration{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDMDtConfiguration() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDMDtplmnlist(value int32) (*xnapconstantsv1.IdMDtplmnlist, error) {

	msg := &xnapconstantsv1.IdMDtplmnlist{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDMDtplmnlist() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDTraceCollectionEntityURI(value int32) (*xnapconstantsv1.IdTraceCollectionEntityUri, error) {

	msg := &xnapconstantsv1.IdTraceCollectionEntityUri{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDTraceCollectionEntityURI() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDUEradioCapabilityID(value int32) (*xnapconstantsv1.IdUEradioCapabilityId, error) {

	msg := &xnapconstantsv1.IdUEradioCapabilityId{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDUEradioCapabilityID() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDCSiRStransmissionIndication(value int32) (*xnapconstantsv1.IdCSiRStransmissionIndication, error) {

	msg := &xnapconstantsv1.IdCSiRStransmissionIndication{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDCSiRStransmissionIndication() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDSNtriggered(value int32) (*xnapconstantsv1.IdSNtriggered, error) {

	msg := &xnapconstantsv1.IdSNtriggered{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDSNtriggered() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDDLcarrierList(value int32) (*xnapconstantsv1.IdDLcarrierList, error) {

	msg := &xnapconstantsv1.IdDLcarrierList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDDLcarrierList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDExtendedTaisliceSupportList(value int32) (*xnapconstantsv1.IdExtendedTaisliceSupportList, error) {

	msg := &xnapconstantsv1.IdExtendedTaisliceSupportList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDExtendedTaisliceSupportList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDcellAssistanceInfoEUtra(value int32) (*xnapconstantsv1.IdcellAssistanceInfoEUtra, error) {

	msg := &xnapconstantsv1.IdcellAssistanceInfoEUtra{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDcellAssistanceInfoEUtra() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDConfiguredTacindication(value int32) (*xnapconstantsv1.IdConfiguredTacindication, error) {

	msg := &xnapconstantsv1.IdConfiguredTacindication{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDConfiguredTacindication() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDsecondarySNULPDcpUPTNlinfo(value int32) (*xnapconstantsv1.IdsecondarySNULPDcpUPTNlinfo, error) {

	msg := &xnapconstantsv1.IdsecondarySNULPDcpUPTNlinfo{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDsecondarySNULPDcpUPTNlinfo() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDpdcpDuplicationConfiguration(value int32) (*xnapconstantsv1.IdpdcpDuplicationConfiguration, error) {

	msg := &xnapconstantsv1.IdpdcpDuplicationConfiguration{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDpdcpDuplicationConfiguration() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDduplicationActivation(value int32) (*xnapconstantsv1.IdduplicationActivation, error) {

	msg := &xnapconstantsv1.IdduplicationActivation{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDduplicationActivation() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDNPrachconfiguration(value int32) (*xnapconstantsv1.IdNPrachconfiguration, error) {

	msg := &xnapconstantsv1.IdNPrachconfiguration{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDNPrachconfiguration() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDQosMonitoringReportingFrequency(value int32) (*xnapconstantsv1.IdQosMonitoringReportingFrequency, error) {

	msg := &xnapconstantsv1.IdQosMonitoringReportingFrequency{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDQosMonitoringReportingFrequency() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDQoSflowsMappedtoDrbSetupResponseMNterminated(value int32) (*xnapconstantsv1.IdQoSflowsMappedtoDrbSetupResponseMNterminated, error) {

	msg := &xnapconstantsv1.IdQoSflowsMappedtoDrbSetupResponseMNterminated{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDQoSflowsMappedtoDrbSetupResponseMNterminated() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDDLschedulingPDcchCCeusage(value int32) (*xnapconstantsv1.IdDLschedulingPDcchCCeusage, error) {

	msg := &xnapconstantsv1.IdDLschedulingPDcchCCeusage{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDDLschedulingPDcchCCeusage() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDULschedulingPDcchCCeusage(value int32) (*xnapconstantsv1.IdULschedulingPDcchCCeusage, error) {

	msg := &xnapconstantsv1.IdULschedulingPDcchCCeusage{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDULschedulingPDcchCCeusage() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDSFnOffset(value int32) (*xnapconstantsv1.IdSFnOffset, error) {

	msg := &xnapconstantsv1.IdSFnOffset{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDSFnOffset() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDQoSmonitoringDisabled(value int32) (*xnapconstantsv1.IdQoSmonitoringDisabled, error) {

	msg := &xnapconstantsv1.IdQoSmonitoringDisabled{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDQoSmonitoringDisabled() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDExtendedUeidentityIndexValue(value int32) (*xnapconstantsv1.IdExtendedUeidentityIndexValue, error) {

	msg := &xnapconstantsv1.IdExtendedUeidentityIndexValue{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDExtendedUeidentityIndexValue() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDPagingeDrxinformation(value int32) (*xnapconstantsv1.IdPagingeDrxinformation, error) {

	msg := &xnapconstantsv1.IdPagingeDrxinformation{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDPagingeDrxinformation() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDCHoMRdcEarlyDataForwarding(value int32) (*xnapconstantsv1.IdCHoMRdcEarlyDataForwarding, error) {

	msg := &xnapconstantsv1.IdCHoMRdcEarlyDataForwarding{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDCHoMRdcEarlyDataForwarding() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDSCgindicator(value int32) (*xnapconstantsv1.IdSCgindicator, error) {

	msg := &xnapconstantsv1.IdSCgindicator{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDSCgindicator() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDUEspecificDrx(value int32) (*xnapconstantsv1.IdUEspecificDrx, error) {

	msg := &xnapconstantsv1.IdUEspecificDrx{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDUEspecificDrx() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDPDusessionExpectedUeactivityBehaviour(value int32) (*xnapconstantsv1.IdPDusessionExpectedUeactivityBehaviour, error) {

	msg := &xnapconstantsv1.IdPDusessionExpectedUeactivityBehaviour{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDPDusessionExpectedUeactivityBehaviour() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDQoSMappingInformation(value int32) (*xnapconstantsv1.IdQoSMappingInformation, error) {

	msg := &xnapconstantsv1.IdQoSMappingInformation{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDQoSMappingInformation() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDAdditionLocationInformation(value int32) (*xnapconstantsv1.IdAdditionLocationInformation, error) {

	msg := &xnapconstantsv1.IdAdditionLocationInformation{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDAdditionLocationInformation() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDdataForwardingInfoFromTargetEUTrannode(value int32) (*xnapconstantsv1.IddataForwardingInfoFromTargetEUTrannode, error) {

	msg := &xnapconstantsv1.IddataForwardingInfoFromTargetEUTrannode{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDdataForwardingInfoFromTargetEUTrannode() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDDirectForwardingPathAvailability(value int32) (*xnapconstantsv1.IdDirectForwardingPathAvailability, error) {

	msg := &xnapconstantsv1.IdDirectForwardingPathAvailability{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDDirectForwardingPathAvailability() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDSourceNgRAnnodeID(value int32) (*xnapconstantsv1.IdSourceNgRAnnodeID, error) {

	msg := &xnapconstantsv1.IdSourceNgRAnnodeID{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDSourceNgRAnnodeID() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDSourceDlforwardingIpaddress(value int32) (*xnapconstantsv1.IdSourceDlforwardingIpaddress, error) {

	msg := &xnapconstantsv1.IdSourceDlforwardingIpaddress{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDSourceDlforwardingIpaddress() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDSourceNodeDlforwardingIpaddress(value int32) (*xnapconstantsv1.IdSourceNodeDlforwardingIpaddress, error) {

	msg := &xnapconstantsv1.IdSourceNodeDlforwardingIpaddress{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDSourceNodeDlforwardingIpaddress() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDExtendedReportIntervalMdt(value int32) (*xnapconstantsv1.IdExtendedReportIntervalMdt, error) {

	msg := &xnapconstantsv1.IdExtendedReportIntervalMdt{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDExtendedReportIntervalMdt() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDSecurityIndication(value int32) (*xnapconstantsv1.IdSecurityIndication, error) {

	msg := &xnapconstantsv1.IdSecurityIndication{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDSecurityIndication() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDRRcconnReestabIndicator(value int32) (*xnapconstantsv1.IdRRcconnReestabIndicator, error) {

	msg := &xnapconstantsv1.IdRRcconnReestabIndicator{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDRRcconnReestabIndicator() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateIDTargetNodeID(value int32) (*xnapconstantsv1.IdTargetNodeId, error) {

	msg := &xnapconstantsv1.IdTargetNodeId{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIDTargetNodeID() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateProcedureCode(value int32) (*xnapcommondatatypesv1.ProcedureCode, error) {

	msg := &xnapcommondatatypesv1.ProcedureCode{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateProcedureCode() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateProtocolIeID(value int32) (*xnapcommondatatypesv1.ProtocolIeID, error) {

	msg := &xnapcommondatatypesv1.ProtocolIeID{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateProtocolIeID() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateInitiatingMessage(procedureCode int32, criticality xnapcommondatatypesv1.Criticality, value *xnappdudescriptionsv1.InitiatingMessageXnApElementaryProcedures) (*xnappdudescriptionsv1.InitiatingMessage, error) {

	msg := &xnappdudescriptionsv1.InitiatingMessage{}
	msg.ProcedureCode = procedureCode
	msg.Criticality = criticality
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateInitiatingMessage() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateSuccessfulOutcome(procedureCode int32, criticality xnapcommondatatypesv1.Criticality, value *xnappdudescriptionsv1.SuccessfulOutcomeXnApElementaryProcedures) (*xnappdudescriptionsv1.SuccessfulOutcome, error) {

	msg := &xnappdudescriptionsv1.SuccessfulOutcome{}
	msg.ProcedureCode = procedureCode
	msg.Criticality = criticality
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSuccessfulOutcome() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateUnsuccessfulOutcome(procedureCode int32, criticality xnapcommondatatypesv1.Criticality, value *xnappdudescriptionsv1.UnsuccessfulOutcomeXnApElementaryProcedures) (*xnappdudescriptionsv1.UnsuccessfulOutcome, error) {

	msg := &xnappdudescriptionsv1.UnsuccessfulOutcome{}
	msg.ProcedureCode = procedureCode
	msg.Criticality = criticality
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateUnsuccessfulOutcome() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateHandoverRequestIEs(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, value *xnappducontentsv1.HandoverRequestIEsValue) (*xnappducontentsv1.HandoverRequestIEs, error) {

	msg := &xnappducontentsv1.HandoverRequestIEs{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateHandoverRequestIEs() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateHandoverRequest(protocolIes []*xnappducontentsv1.HandoverRequestIEs) (*xnappducontentsv1.HandoverRequest, error) {

	msg := &xnappducontentsv1.HandoverRequest{}
	msg.ProtocolIes = protocolIes

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateHandoverRequest() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateUecontextInfoHorequestExtIes(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, extension *xnappducontentsv1.UecontextInfoHorequestExtIesExtension) (*xnappducontentsv1.UecontextInfoHorequestExtIes, error) {

	msg := &xnappducontentsv1.UecontextInfoHorequestExtIes{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Extension = extension

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateUecontextInfoHorequestExtIes() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateUecontextRefAtSnHOrequest(globalNgRannodeID *xnapiesv1.GlobalNgRAnnodeID, sNNgRannodeUexnApID *xnapiesv1.NgRAnnodeUexnApid) (*xnappducontentsv1.UecontextRefAtSnHOrequest, error) {

	msg := &xnappducontentsv1.UecontextRefAtSnHOrequest{}
	msg.GlobalNgRannodeId = globalNgRannodeID
	msg.SNNgRannodeUexnApid = sNNgRannodeUexnApID

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateUecontextRefAtSnHOrequest() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateHandoverRequestAcknowledgeIEs(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, value *xnappducontentsv1.HandoverRequestAcknowledgeIEsValue) (*xnappducontentsv1.HandoverRequestAcknowledgeIEs, error) {

	msg := &xnappducontentsv1.HandoverRequestAcknowledgeIEs{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateHandoverRequestAcknowledgeIEs() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateHandoverRequestAcknowledge(protocolIes []*xnappducontentsv1.HandoverRequestAcknowledgeIEs) (*xnappducontentsv1.HandoverRequestAcknowledge, error) {

	msg := &xnappducontentsv1.HandoverRequestAcknowledge{}
	msg.ProtocolIes = protocolIes

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateHandoverRequestAcknowledge() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateHandoverPreparationFailureIEs(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, value *xnappducontentsv1.HandoverPreparationFailureIEsValue) (*xnappducontentsv1.HandoverPreparationFailureIEs, error) {

	msg := &xnappducontentsv1.HandoverPreparationFailureIEs{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateHandoverPreparationFailureIEs() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateHandoverPreparationFailure(protocolIes []*xnappducontentsv1.HandoverPreparationFailureIEs) (*xnappducontentsv1.HandoverPreparationFailure, error) {

	msg := &xnappducontentsv1.HandoverPreparationFailure{}
	msg.ProtocolIes = protocolIes

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateHandoverPreparationFailure() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateSnstatusTransferIEs(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, value *xnappducontentsv1.SnstatusTransferIEsValue) (*xnappducontentsv1.SnstatusTransferIEs, error) {

	msg := &xnappducontentsv1.SnstatusTransferIEs{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnstatusTransferIEs() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateSnstatusTransfer(protocolIes []*xnappducontentsv1.SnstatusTransferIEs) (*xnappducontentsv1.SnstatusTransfer, error) {

	msg := &xnappducontentsv1.SnstatusTransfer{}
	msg.ProtocolIes = protocolIes

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnstatusTransfer() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateUecontextReleaseIEs(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, value *xnappducontentsv1.UecontextReleaseIEsValue) (*xnappducontentsv1.UecontextReleaseIEs, error) {

	msg := &xnappducontentsv1.UecontextReleaseIEs{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateUecontextReleaseIEs() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateUecontextRelease(protocolIes []*xnappducontentsv1.UecontextReleaseIEs) (*xnappducontentsv1.UecontextRelease, error) {

	msg := &xnappducontentsv1.UecontextRelease{}
	msg.ProtocolIes = protocolIes

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateUecontextRelease() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateHandoverCancelIEs(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, value *xnappducontentsv1.HandoverCancelIEsValue) (*xnappducontentsv1.HandoverCancelIEs, error) {

	msg := &xnappducontentsv1.HandoverCancelIEs{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateHandoverCancelIEs() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateHandoverCancel(protocolIes []*xnappducontentsv1.HandoverCancelIEs) (*xnappducontentsv1.HandoverCancel, error) {

	msg := &xnappducontentsv1.HandoverCancel{}
	msg.ProtocolIes = protocolIes

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateHandoverCancel() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateHandoverSuccessIEs(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, value *xnappducontentsv1.HandoverSuccessIEsValue) (*xnappducontentsv1.HandoverSuccessIEs, error) {

	msg := &xnappducontentsv1.HandoverSuccessIEs{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateHandoverSuccessIEs() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateHandoverSuccess(protocolIes []*xnappducontentsv1.HandoverSuccessIEs) (*xnappducontentsv1.HandoverSuccess, error) {

	msg := &xnappducontentsv1.HandoverSuccess{}
	msg.ProtocolIes = protocolIes

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateHandoverSuccess() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateConditionalHandoverCancelIEs(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, value *xnappducontentsv1.ConditionalHandoverCancelIEsValue) (*xnappducontentsv1.ConditionalHandoverCancelIEs, error) {

	msg := &xnappducontentsv1.ConditionalHandoverCancelIEs{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateConditionalHandoverCancelIEs() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateConditionalHandoverCancel(protocolIes []*xnappducontentsv1.ConditionalHandoverCancelIEs) (*xnappducontentsv1.ConditionalHandoverCancel, error) {

	msg := &xnappducontentsv1.ConditionalHandoverCancel{}
	msg.ProtocolIes = protocolIes

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateConditionalHandoverCancel() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateEarlyStatusTransferIEs(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, value *xnappducontentsv1.EarlyStatusTransferIEsValue) (*xnappducontentsv1.EarlyStatusTransferIEs, error) {

	msg := &xnappducontentsv1.EarlyStatusTransferIEs{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateEarlyStatusTransferIEs() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateEarlyStatusTransfer(protocolIes []*xnappducontentsv1.EarlyStatusTransferIEs) (*xnappducontentsv1.EarlyStatusTransfer, error) {

	msg := &xnappducontentsv1.EarlyStatusTransfer{}
	msg.ProtocolIes = protocolIes

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateEarlyStatusTransfer() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateFirstDlcount(dRbsSubjectToEarlyStatusTransfer *xnapiesv1.DrbsSubjectToEarlyStatusTransferList) (*xnappducontentsv1.FirstDlcount, error) {

	msg := &xnappducontentsv1.FirstDlcount{}
	msg.DRbsSubjectToEarlyStatusTransfer = dRbsSubjectToEarlyStatusTransfer

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateFirstDlcount() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateDldiscarding(dRbsSubjectToDldiscarding *xnapiesv1.DrbsSubjectToDldiscardingList) (*xnappducontentsv1.Dldiscarding, error) {

	msg := &xnappducontentsv1.Dldiscarding{}
	msg.DRbsSubjectToDldiscarding = dRbsSubjectToDldiscarding

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateDldiscarding() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateRanpagingIEs(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, value *xnappducontentsv1.RanpagingIEsValue) (*xnappducontentsv1.RanpagingIEs, error) {

	msg := &xnappducontentsv1.RanpagingIEs{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRanpagingIEs() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateRanpaging(protocolIes []*xnappducontentsv1.RanpagingIEs) (*xnappducontentsv1.Ranpaging, error) {

	msg := &xnappducontentsv1.Ranpaging{}
	msg.ProtocolIes = protocolIes

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRanpaging() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateRetrieveUecontextRequestIEs(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, value *xnappducontentsv1.RetrieveUecontextRequestIEsValue) (*xnappducontentsv1.RetrieveUecontextRequestIEs, error) {

	msg := &xnappducontentsv1.RetrieveUecontextRequestIEs{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRetrieveUecontextRequestIEs() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateRetrieveUecontextRequest(protocolIes []*xnappducontentsv1.RetrieveUecontextRequestIEs) (*xnappducontentsv1.RetrieveUecontextRequest, error) {

	msg := &xnappducontentsv1.RetrieveUecontextRequest{}
	msg.ProtocolIes = protocolIes

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRetrieveUecontextRequest() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateRetrieveUecontextResponseIEs(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, value *xnappducontentsv1.RetrieveUecontextResponseIEsValue) (*xnappducontentsv1.RetrieveUecontextResponseIEs, error) {

	msg := &xnappducontentsv1.RetrieveUecontextResponseIEs{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRetrieveUecontextResponseIEs() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateRetrieveUecontextResponse(protocolIes []*xnappducontentsv1.RetrieveUecontextResponseIEs) (*xnappducontentsv1.RetrieveUecontextResponse, error) {

	msg := &xnappducontentsv1.RetrieveUecontextResponse{}
	msg.ProtocolIes = protocolIes

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRetrieveUecontextResponse() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateRetrieveUecontextFailureIEs(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, value *xnappducontentsv1.RetrieveUecontextFailureIEsValue) (*xnappducontentsv1.RetrieveUecontextFailureIEs, error) {

	msg := &xnappducontentsv1.RetrieveUecontextFailureIEs{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRetrieveUecontextFailureIEs() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateRetrieveUecontextFailure(protocolIes []*xnappducontentsv1.RetrieveUecontextFailureIEs) (*xnappducontentsv1.RetrieveUecontextFailure, error) {

	msg := &xnappducontentsv1.RetrieveUecontextFailure{}
	msg.ProtocolIes = protocolIes

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRetrieveUecontextFailure() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateXnUaddressIndicationIEs(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, value *xnappducontentsv1.XnUaddressIndicationIEsValue) (*xnappducontentsv1.XnUaddressIndicationIEs, error) {

	msg := &xnappducontentsv1.XnUaddressIndicationIEs{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateXnUaddressIndicationIEs() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateXnUaddressIndication(protocolIes []*xnappducontentsv1.XnUaddressIndicationIEs) (*xnappducontentsv1.XnUaddressIndication, error) {

	msg := &xnappducontentsv1.XnUaddressIndication{}
	msg.ProtocolIes = protocolIes

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateXnUaddressIndication() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateSnodeAdditionRequestIEs(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, value *xnappducontentsv1.SnodeAdditionRequestIEsValue) (*xnappducontentsv1.SnodeAdditionRequestIEs, error) {

	msg := &xnappducontentsv1.SnodeAdditionRequestIEs{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeAdditionRequestIEs() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateSnodeAdditionRequest(protocolIes []*xnappducontentsv1.SnodeAdditionRequestIEs) (*xnappducontentsv1.SnodeAdditionRequest, error) {

	msg := &xnappducontentsv1.SnodeAdditionRequest{}
	msg.ProtocolIes = protocolIes

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeAdditionRequest() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreatePdusessionToBeAddedAddReq(value []*xnappducontentsv1.PdusessionToBeAddedAddReqItem) (*xnappducontentsv1.PdusessionToBeAddedAddReq, error) {

	msg := &xnappducontentsv1.PdusessionToBeAddedAddReq{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePdusessionToBeAddedAddReq() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateSnodeAdditionRequestAcknowledgeIEs(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, value *xnappducontentsv1.SnodeAdditionRequestAcknowledgeIEsValue) (*xnappducontentsv1.SnodeAdditionRequestAcknowledgeIEs, error) {

	msg := &xnappducontentsv1.SnodeAdditionRequestAcknowledgeIEs{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeAdditionRequestAcknowledgeIEs() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateSnodeAdditionRequestAcknowledge(protocolIes []*xnappducontentsv1.SnodeAdditionRequestAcknowledgeIEs) (*xnappducontentsv1.SnodeAdditionRequestAcknowledge, error) {

	msg := &xnappducontentsv1.SnodeAdditionRequestAcknowledge{}
	msg.ProtocolIes = protocolIes

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeAdditionRequestAcknowledge() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreatePdusessionAdmittedAddedAddReqAck(value []*xnappducontentsv1.PdusessionAdmittedAddedAddReqAckItem) (*xnappducontentsv1.PdusessionAdmittedAddedAddReqAck, error) {

	msg := &xnappducontentsv1.PdusessionAdmittedAddedAddReqAck{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePdusessionAdmittedAddedAddReqAck() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateSnodeAdditionRequestRejectIEs(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, value *xnappducontentsv1.SnodeAdditionRequestRejectIEsValue) (*xnappducontentsv1.SnodeAdditionRequestRejectIEs, error) {

	msg := &xnappducontentsv1.SnodeAdditionRequestRejectIEs{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeAdditionRequestRejectIEs() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateSnodeAdditionRequestReject(protocolIes []*xnappducontentsv1.SnodeAdditionRequestRejectIEs) (*xnappducontentsv1.SnodeAdditionRequestReject, error) {

	msg := &xnappducontentsv1.SnodeAdditionRequestReject{}
	msg.ProtocolIes = protocolIes

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeAdditionRequestReject() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateSnodeReconfigurationCompleteIEs(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, value *xnappducontentsv1.SnodeReconfigurationCompleteIEsValue) (*xnappducontentsv1.SnodeReconfigurationCompleteIEs, error) {

	msg := &xnappducontentsv1.SnodeReconfigurationCompleteIEs{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeReconfigurationCompleteIEs() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateSnodeReconfigurationComplete(protocolIes []*xnappducontentsv1.SnodeReconfigurationCompleteIEs) (*xnappducontentsv1.SnodeReconfigurationComplete, error) {

	msg := &xnappducontentsv1.SnodeReconfigurationComplete{}
	msg.ProtocolIes = protocolIes

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeReconfigurationComplete() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateResponseInfoReconfCompl(responseTypeReconfComplete *xnappducontentsv1.ResponseTypeReconfComplete) (*xnappducontentsv1.ResponseInfoReconfCompl, error) {

	msg := &xnappducontentsv1.ResponseInfoReconfCompl{}
	msg.ResponseTypeReconfComplete = responseTypeReconfComplete

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateResponseInfoReconfCompl() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateSnodeModificationRequestIEs(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, value *xnappducontentsv1.SnodeModificationRequestIEsValue) (*xnappducontentsv1.SnodeModificationRequestIEs, error) {

	msg := &xnappducontentsv1.SnodeModificationRequestIEs{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeModificationRequestIEs() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateSnodeModificationRequest(protocolIes []*xnappducontentsv1.SnodeModificationRequestIEs) (*xnappducontentsv1.SnodeModificationRequest, error) {

	msg := &xnappducontentsv1.SnodeModificationRequest{}
	msg.ProtocolIes = protocolIes

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeModificationRequest() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreatePdusessionsToBeAddedSNmodRequestList(value []*xnappducontentsv1.PdusessionsToBeAddedSNmodRequestItem) (*xnappducontentsv1.PdusessionsToBeAddedSNmodRequestList, error) {

	msg := &xnappducontentsv1.PdusessionsToBeAddedSNmodRequestList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePdusessionsToBeAddedSNmodRequestList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreatePdusessionsToBeModifiedSNmodRequestList(value []*xnappducontentsv1.PdusessionsToBeModifiedSNmodRequestItem) (*xnappducontentsv1.PdusessionsToBeModifiedSNmodRequestList, error) {

	msg := &xnappducontentsv1.PdusessionsToBeModifiedSNmodRequestList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePdusessionsToBeModifiedSNmodRequestList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreatePdusessionsToBeModifiedSNmodRequestItemExtIes(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, extension *xnappducontentsv1.PdusessionsToBeModifiedSNmodRequestItemExtIesExtension) (*xnappducontentsv1.PdusessionsToBeModifiedSNmodRequestItemExtIes, error) {

	msg := &xnappducontentsv1.PdusessionsToBeModifiedSNmodRequestItemExtIes{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Extension = extension

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePdusessionsToBeModifiedSNmodRequestItemExtIes() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateSnodeModificationRequestAcknowledgeIEs(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, value *xnappducontentsv1.SnodeModificationRequestAcknowledgeIEsValue) (*xnappducontentsv1.SnodeModificationRequestAcknowledgeIEs, error) {

	msg := &xnappducontentsv1.SnodeModificationRequestAcknowledgeIEs{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeModificationRequestAcknowledgeIEs() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateSnodeModificationRequestAcknowledge(protocolIes []*xnappducontentsv1.SnodeModificationRequestAcknowledgeIEs) (*xnappducontentsv1.SnodeModificationRequestAcknowledge, error) {

	msg := &xnappducontentsv1.SnodeModificationRequestAcknowledge{}
	msg.ProtocolIes = protocolIes

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeModificationRequestAcknowledge() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreatePdusessionAdmittedToBeAddedSnmodResponse(value []*xnappducontentsv1.PdusessionAdmittedToBeAddedSnmodResponseItem) (*xnappducontentsv1.PdusessionAdmittedToBeAddedSnmodResponse, error) {

	msg := &xnappducontentsv1.PdusessionAdmittedToBeAddedSnmodResponse{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePdusessionAdmittedToBeAddedSnmodResponse() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreatePdusessionAdmittedToBeModifiedSnmodResponse(value []*xnappducontentsv1.PdusessionAdmittedToBeModifiedSnmodResponseItem) (*xnappducontentsv1.PdusessionAdmittedToBeModifiedSnmodResponse, error) {

	msg := &xnappducontentsv1.PdusessionAdmittedToBeModifiedSnmodResponse{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePdusessionAdmittedToBeModifiedSnmodResponse() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreatePdusessionDataForwardingSNmodResponse(snTerminated *xnapiesv1.PdusessionListwithDataForwardingRequest) (*xnappducontentsv1.PdusessionDataForwardingSNmodResponse, error) {

	msg := &xnappducontentsv1.PdusessionDataForwardingSNmodResponse{}
	msg.SnTerminated = snTerminated

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePdusessionDataForwardingSNmodResponse() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateSnodeModificationRequestRejectIEs(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, value *xnappducontentsv1.SnodeModificationRequestRejectIEsValue) (*xnappducontentsv1.SnodeModificationRequestRejectIEs, error) {

	msg := &xnappducontentsv1.SnodeModificationRequestRejectIEs{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeModificationRequestRejectIEs() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateSnodeModificationRequestReject(protocolIes []*xnappducontentsv1.SnodeModificationRequestRejectIEs) (*xnappducontentsv1.SnodeModificationRequestReject, error) {

	msg := &xnappducontentsv1.SnodeModificationRequestReject{}
	msg.ProtocolIes = protocolIes

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeModificationRequestReject() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateSnodeModificationRequiredIEs(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, value *xnappducontentsv1.SnodeModificationRequiredIEsValue) (*xnappducontentsv1.SnodeModificationRequiredIEs, error) {

	msg := &xnappducontentsv1.SnodeModificationRequiredIEs{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeModificationRequiredIEs() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateSnodeModificationRequired(protocolIes []*xnappducontentsv1.SnodeModificationRequiredIEs) (*xnappducontentsv1.SnodeModificationRequired, error) {

	msg := &xnappducontentsv1.SnodeModificationRequired{}
	msg.ProtocolIes = protocolIes

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeModificationRequired() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreatePdusessionToBeModifiedSnmodRequired(value []*xnappducontentsv1.PdusessionToBeModifiedSnmodRequiredItem) (*xnappducontentsv1.PdusessionToBeModifiedSnmodRequired, error) {

	msg := &xnappducontentsv1.PdusessionToBeModifiedSnmodRequired{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePdusessionToBeModifiedSnmodRequired() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateSnodeModificationConfirmIEs(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, value *xnappducontentsv1.SnodeModificationConfirmIEsValue) (*xnappducontentsv1.SnodeModificationConfirmIEs, error) {

	msg := &xnappducontentsv1.SnodeModificationConfirmIEs{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeModificationConfirmIEs() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateSnodeModificationConfirm(protocolIes []*xnappducontentsv1.SnodeModificationConfirmIEs) (*xnappducontentsv1.SnodeModificationConfirm, error) {

	msg := &xnappducontentsv1.SnodeModificationConfirm{}
	msg.ProtocolIes = protocolIes

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeModificationConfirm() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreatePdusessionAdmittedModSnmodConfirm(value []*xnappducontentsv1.PdusessionAdmittedModSnmodConfirmItem) (*xnappducontentsv1.PdusessionAdmittedModSnmodConfirm, error) {

	msg := &xnappducontentsv1.PdusessionAdmittedModSnmodConfirm{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePdusessionAdmittedModSnmodConfirm() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateSnodeModificationRefuseIEs(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, value *xnappducontentsv1.SnodeModificationRefuseIEsValue) (*xnappducontentsv1.SnodeModificationRefuseIEs, error) {

	msg := &xnappducontentsv1.SnodeModificationRefuseIEs{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeModificationRefuseIEs() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateSnodeModificationRefuse(protocolIes []*xnappducontentsv1.SnodeModificationRefuseIEs) (*xnappducontentsv1.SnodeModificationRefuse, error) {

	msg := &xnappducontentsv1.SnodeModificationRefuse{}
	msg.ProtocolIes = protocolIes

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeModificationRefuse() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateSnodeReleaseRequestIEs(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, value *xnappducontentsv1.SnodeReleaseRequestIEsValue) (*xnappducontentsv1.SnodeReleaseRequestIEs, error) {

	msg := &xnappducontentsv1.SnodeReleaseRequestIEs{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeReleaseRequestIEs() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateSnodeReleaseRequest(protocolIes []*xnappducontentsv1.SnodeReleaseRequestIEs) (*xnappducontentsv1.SnodeReleaseRequest, error) {

	msg := &xnappducontentsv1.SnodeReleaseRequest{}
	msg.ProtocolIes = protocolIes

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeReleaseRequest() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateSnodeReleaseRequestAcknowledgeIEs(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, value *xnappducontentsv1.SnodeReleaseRequestAcknowledgeIEsValue) (*xnappducontentsv1.SnodeReleaseRequestAcknowledgeIEs, error) {

	msg := &xnappducontentsv1.SnodeReleaseRequestAcknowledgeIEs{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeReleaseRequestAcknowledgeIEs() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateSnodeReleaseRequestAcknowledge(protocolIes []*xnappducontentsv1.SnodeReleaseRequestAcknowledgeIEs) (*xnappducontentsv1.SnodeReleaseRequestAcknowledge, error) {

	msg := &xnappducontentsv1.SnodeReleaseRequestAcknowledge{}
	msg.ProtocolIes = protocolIes

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeReleaseRequestAcknowledge() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateSnodeReleaseRejectIEs(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, value *xnappducontentsv1.SnodeReleaseRejectIEsValue) (*xnappducontentsv1.SnodeReleaseRejectIEs, error) {

	msg := &xnappducontentsv1.SnodeReleaseRejectIEs{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeReleaseRejectIEs() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateSnodeReleaseReject(protocolIes []*xnappducontentsv1.SnodeReleaseRejectIEs) (*xnappducontentsv1.SnodeReleaseReject, error) {

	msg := &xnappducontentsv1.SnodeReleaseReject{}
	msg.ProtocolIes = protocolIes

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeReleaseReject() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateSnodeReleaseRequiredIEs(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, value *xnappducontentsv1.SnodeReleaseRequiredIEsValue) (*xnappducontentsv1.SnodeReleaseRequiredIEs, error) {

	msg := &xnappducontentsv1.SnodeReleaseRequiredIEs{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeReleaseRequiredIEs() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateSnodeReleaseRequired(protocolIes []*xnappducontentsv1.SnodeReleaseRequiredIEs) (*xnappducontentsv1.SnodeReleaseRequired, error) {

	msg := &xnappducontentsv1.SnodeReleaseRequired{}
	msg.ProtocolIes = protocolIes

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeReleaseRequired() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateSnodeReleaseConfirmIEs(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, value *xnappducontentsv1.SnodeReleaseConfirmIEsValue) (*xnappducontentsv1.SnodeReleaseConfirmIEs, error) {

	msg := &xnappducontentsv1.SnodeReleaseConfirmIEs{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeReleaseConfirmIEs() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateSnodeReleaseConfirm(protocolIes []*xnappducontentsv1.SnodeReleaseConfirmIEs) (*xnappducontentsv1.SnodeReleaseConfirm, error) {

	msg := &xnappducontentsv1.SnodeReleaseConfirm{}
	msg.ProtocolIes = protocolIes

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeReleaseConfirm() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateSnodeCounterCheckRequestIEs(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, value *xnappducontentsv1.SnodeCounterCheckRequestIEsValue) (*xnappducontentsv1.SnodeCounterCheckRequestIEs, error) {

	msg := &xnappducontentsv1.SnodeCounterCheckRequestIEs{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeCounterCheckRequestIEs() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateSnodeCounterCheckRequest(protocolIes []*xnappducontentsv1.SnodeCounterCheckRequestIEs) (*xnappducontentsv1.SnodeCounterCheckRequest, error) {

	msg := &xnappducontentsv1.SnodeCounterCheckRequest{}
	msg.ProtocolIes = protocolIes

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeCounterCheckRequest() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateBearersSubjectToCounterCheckList(value []*xnappducontentsv1.BearersSubjectToCounterCheckItem) (*xnappducontentsv1.BearersSubjectToCounterCheckList, error) {

	msg := &xnappducontentsv1.BearersSubjectToCounterCheckList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateBearersSubjectToCounterCheckList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateBearersSubjectToCounterCheckItem(drbID *xnapiesv1.DrbID, ulCount int32, dlCount int32) (*xnappducontentsv1.BearersSubjectToCounterCheckItem, error) {

	msg := &xnappducontentsv1.BearersSubjectToCounterCheckItem{}
	msg.DrbId = drbID
	msg.UlCount = ulCount
	msg.DlCount = dlCount

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateBearersSubjectToCounterCheckItem() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateSnodeChangeRequiredIEs(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, value *xnappducontentsv1.SnodeChangeRequiredIEsValue) (*xnappducontentsv1.SnodeChangeRequiredIEs, error) {

	msg := &xnappducontentsv1.SnodeChangeRequiredIEs{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeChangeRequiredIEs() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateSnodeChangeRequired(protocolIes []*xnappducontentsv1.SnodeChangeRequiredIEs) (*xnappducontentsv1.SnodeChangeRequired, error) {

	msg := &xnappducontentsv1.SnodeChangeRequired{}
	msg.ProtocolIes = protocolIes

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeChangeRequired() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreatePdusessionSNchangeRequiredList(value []*xnappducontentsv1.PdusessionSNchangeRequiredItem) (*xnappducontentsv1.PdusessionSNchangeRequiredList, error) {

	msg := &xnappducontentsv1.PdusessionSNchangeRequiredList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePdusessionSNchangeRequiredList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateSnodeChangeConfirmIEs(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, value *xnappducontentsv1.SnodeChangeConfirmIEsValue) (*xnappducontentsv1.SnodeChangeConfirmIEs, error) {

	msg := &xnappducontentsv1.SnodeChangeConfirmIEs{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeChangeConfirmIEs() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateSnodeChangeConfirm(protocolIes []*xnappducontentsv1.SnodeChangeConfirmIEs) (*xnappducontentsv1.SnodeChangeConfirm, error) {

	msg := &xnappducontentsv1.SnodeChangeConfirm{}
	msg.ProtocolIes = protocolIes

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeChangeConfirm() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreatePdusessionSNchangeConfirmList(value []*xnappducontentsv1.PdusessionSNchangeConfirmItem) (*xnappducontentsv1.PdusessionSNchangeConfirmList, error) {

	msg := &xnappducontentsv1.PdusessionSNchangeConfirmList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePdusessionSNchangeConfirmList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateSnodeChangeRefuseIEs(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, value *xnappducontentsv1.SnodeChangeRefuseIEsValue) (*xnappducontentsv1.SnodeChangeRefuseIEs, error) {

	msg := &xnappducontentsv1.SnodeChangeRefuseIEs{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeChangeRefuseIEs() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateSnodeChangeRefuse(protocolIes []*xnappducontentsv1.SnodeChangeRefuseIEs) (*xnappducontentsv1.SnodeChangeRefuse, error) {

	msg := &xnappducontentsv1.SnodeChangeRefuse{}
	msg.ProtocolIes = protocolIes

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeChangeRefuse() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateRrctransferIEs(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, value *xnappducontentsv1.RrctransferIEsValue) (*xnappducontentsv1.RrctransferIEs, error) {

	msg := &xnappducontentsv1.RrctransferIEs{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRrctransferIEs() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateRrctransfer(protocolIes []*xnappducontentsv1.RrctransferIEs) (*xnappducontentsv1.Rrctransfer, error) {

	msg := &xnappducontentsv1.Rrctransfer{}
	msg.ProtocolIes = protocolIes

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRrctransfer() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateUereportRrctransfer(rrcContainer []byte) (*xnappducontentsv1.UereportRrctransfer, error) {

	msg := &xnappducontentsv1.UereportRrctransfer{}
	msg.RrcContainer = rrcContainer

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateUereportRrctransfer() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateFastMcgrecoveryRrctransfer(rrcContainer []byte) (*xnappducontentsv1.FastMcgrecoveryRrctransfer, error) {

	msg := &xnappducontentsv1.FastMcgrecoveryRrctransfer{}
	msg.RrcContainer = rrcContainer

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateFastMcgrecoveryRrctransfer() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateNotificationControlIndicationIEs(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, value *xnappducontentsv1.NotificationControlIndicationIEsValue) (*xnappducontentsv1.NotificationControlIndicationIEs, error) {

	msg := &xnappducontentsv1.NotificationControlIndicationIEs{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNotificationControlIndicationIEs() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateNotificationControlIndication(protocolIes []*xnappducontentsv1.NotificationControlIndicationIEs) (*xnappducontentsv1.NotificationControlIndication, error) {

	msg := &xnappducontentsv1.NotificationControlIndication{}
	msg.ProtocolIes = protocolIes

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNotificationControlIndication() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreatePdusessionResourcesNotifyList(value []*xnappducontentsv1.PdusessionResourcesNotifyItem) (*xnappducontentsv1.PdusessionResourcesNotifyList, error) {

	msg := &xnappducontentsv1.PdusessionResourcesNotifyList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePdusessionResourcesNotifyList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreatePdusessionResourcesNotifyItem(pduSessionID *xnapiesv1.PdusessionID, qosFlowsNotificationContrIndInfo *xnapiesv1.QoSflowNotificationControlIndicationInfo) (*xnappducontentsv1.PdusessionResourcesNotifyItem, error) {

	msg := &xnappducontentsv1.PdusessionResourcesNotifyItem{}
	msg.PduSessionId = pduSessionID
	msg.QosFlowsNotificationContrIndInfo = qosFlowsNotificationContrIndInfo

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePdusessionResourcesNotifyItem() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateActivityNotificationIEs(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, value *xnappducontentsv1.ActivityNotificationIEsValue) (*xnappducontentsv1.ActivityNotificationIEs, error) {

	msg := &xnappducontentsv1.ActivityNotificationIEs{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateActivityNotificationIEs() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateActivityNotification(protocolIes []*xnappducontentsv1.ActivityNotificationIEs) (*xnappducontentsv1.ActivityNotification, error) {

	msg := &xnappducontentsv1.ActivityNotification{}
	msg.ProtocolIes = protocolIes

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateActivityNotification() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreatePdusessionResourcesActivityNotifyList(value []*xnappducontentsv1.PdusessionResourcesActivityNotifyItem) (*xnappducontentsv1.PdusessionResourcesActivityNotifyList, error) {

	msg := &xnappducontentsv1.PdusessionResourcesActivityNotifyList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePdusessionResourcesActivityNotifyList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateQoSflowsActivityNotifyList(value []*xnappducontentsv1.QoSflowsActivityNotifyItem) (*xnappducontentsv1.QoSflowsActivityNotifyList, error) {

	msg := &xnappducontentsv1.QoSflowsActivityNotifyList{}
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateQoSflowsActivityNotifyList() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateQoSflowsActivityNotifyItem(qosFlowIDentifier *xnapiesv1.QoSflowIdentifier, pduSessionLevelUpactivityreport xnapiesv1.UserPlaneTrafficActivityReport) (*xnappducontentsv1.QoSflowsActivityNotifyItem, error) {

	msg := &xnappducontentsv1.QoSflowsActivityNotifyItem{}
	msg.QosFlowIdentifier = qosFlowIDentifier
	msg.PduSessionLevelUpactivityreport = pduSessionLevelUpactivityreport

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateQoSflowsActivityNotifyItem() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateXnSetupRequestIEs(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, value *xnappducontentsv1.XnSetupRequestIEsValue) (*xnappducontentsv1.XnSetupRequestIEs, error) {

	msg := &xnappducontentsv1.XnSetupRequestIEs{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateXnSetupRequestIEs() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateXnSetupRequest(protocolIes []*xnappducontentsv1.XnSetupRequestIEs) (*xnappducontentsv1.XnSetupRequest, error) {

	msg := &xnappducontentsv1.XnSetupRequest{}
	msg.ProtocolIes = protocolIes

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateXnSetupRequest() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateXnSetupResponseIEs(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, value *xnappducontentsv1.XnSetupResponseIEsValue) (*xnappducontentsv1.XnSetupResponseIEs, error) {

	msg := &xnappducontentsv1.XnSetupResponseIEs{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateXnSetupResponseIEs() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateXnSetupResponse(protocolIes []*xnappducontentsv1.XnSetupResponseIEs) (*xnappducontentsv1.XnSetupResponse, error) {

	msg := &xnappducontentsv1.XnSetupResponse{}
	msg.ProtocolIes = protocolIes

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateXnSetupResponse() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateXnSetupFailureIEs(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, value *xnappducontentsv1.XnSetupFailureIEsValue) (*xnappducontentsv1.XnSetupFailureIEs, error) {

	msg := &xnappducontentsv1.XnSetupFailureIEs{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateXnSetupFailureIEs() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateXnSetupFailure(protocolIes []*xnappducontentsv1.XnSetupFailureIEs) (*xnappducontentsv1.XnSetupFailure, error) {

	msg := &xnappducontentsv1.XnSetupFailure{}
	msg.ProtocolIes = protocolIes

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateXnSetupFailure() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateNgrannodeConfigurationUpdateIEs(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, value *xnappducontentsv1.NgrannodeConfigurationUpdateIEsValue) (*xnappducontentsv1.NgrannodeConfigurationUpdateIEs, error) {

	msg := &xnappducontentsv1.NgrannodeConfigurationUpdateIEs{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNgrannodeConfigurationUpdateIEs() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateNgrannodeConfigurationUpdate(protocolIes []*xnappducontentsv1.NgrannodeConfigurationUpdateIEs) (*xnappducontentsv1.NgrannodeConfigurationUpdate, error) {

	msg := &xnappducontentsv1.NgrannodeConfigurationUpdate{}
	msg.ProtocolIes = protocolIes

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNgrannodeConfigurationUpdate() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateConfigurationUpdategNb(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, value *xnappducontentsv1.ConfigurationUpdategNbValue) (*xnappducontentsv1.ConfigurationUpdategNb, error) {

	msg := &xnappducontentsv1.ConfigurationUpdategNb{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateConfigurationUpdategNb() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateConfigurationUpdatengeNb(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, value *xnappducontentsv1.ConfigurationUpdatengeNbValue) (*xnappducontentsv1.ConfigurationUpdatengeNb, error) {

	msg := &xnappducontentsv1.ConfigurationUpdatengeNb{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateConfigurationUpdatengeNb() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateNgrannodeConfigurationUpdateAcknowledgeIEs(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, value *xnappducontentsv1.NgrannodeConfigurationUpdateAcknowledgeIEsValue) (*xnappducontentsv1.NgrannodeConfigurationUpdateAcknowledgeIEs, error) {

	msg := &xnappducontentsv1.NgrannodeConfigurationUpdateAcknowledgeIEs{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNgrannodeConfigurationUpdateAcknowledgeIEs() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateNgrannodeConfigurationUpdateAcknowledge(protocolIes []*xnappducontentsv1.NgrannodeConfigurationUpdateAcknowledgeIEs) (*xnappducontentsv1.NgrannodeConfigurationUpdateAcknowledge, error) {

	msg := &xnappducontentsv1.NgrannodeConfigurationUpdateAcknowledge{}
	msg.ProtocolIes = protocolIes

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNgrannodeConfigurationUpdateAcknowledge() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateRespondingNodeTypeConfigUpdateAckngeNbExtIes(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, extension *xnappducontentsv1.RespondingNodeTypeConfigUpdateAckngeNbExtIesExtension) (*xnappducontentsv1.RespondingNodeTypeConfigUpdateAckngeNbExtIes, error) {

	msg := &xnappducontentsv1.RespondingNodeTypeConfigUpdateAckngeNbExtIes{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Extension = extension

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRespondingNodeTypeConfigUpdateAckngeNbExtIes() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateRespondingNodeTypeConfigUpdateAckngeNb() (*xnappducontentsv1.RespondingNodeTypeConfigUpdateAckngeNb, error) {

	msg := &xnappducontentsv1.RespondingNodeTypeConfigUpdateAckngeNb{}

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRespondingNodeTypeConfigUpdateAckngeNb() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateRespondingNodeTypeConfigUpdateAckgNbExtIes(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, extension *xnappducontentsv1.RespondingNodeTypeConfigUpdateAckgNbExtIesExtension) (*xnappducontentsv1.RespondingNodeTypeConfigUpdateAckgNbExtIes, error) {

	msg := &xnappducontentsv1.RespondingNodeTypeConfigUpdateAckgNbExtIes{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Extension = extension

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRespondingNodeTypeConfigUpdateAckgNbExtIes() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateNgrannodeConfigurationUpdateFailureIEs(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, value *xnappducontentsv1.NgrannodeConfigurationUpdateFailureIEsValue) (*xnappducontentsv1.NgrannodeConfigurationUpdateFailureIEs, error) {

	msg := &xnappducontentsv1.NgrannodeConfigurationUpdateFailureIEs{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNgrannodeConfigurationUpdateFailureIEs() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateNgrannodeConfigurationUpdateFailure(protocolIes []*xnappducontentsv1.NgrannodeConfigurationUpdateFailureIEs) (*xnappducontentsv1.NgrannodeConfigurationUpdateFailure, error) {

	msg := &xnappducontentsv1.NgrannodeConfigurationUpdateFailure{}
	msg.ProtocolIes = protocolIes

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNgrannodeConfigurationUpdateFailure() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateEUTraNRCellResourceCoordinationRequestIEs(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, value *xnappducontentsv1.EUTraNRCellResourceCoordinationRequestIEsValue) (*xnappducontentsv1.EUTraNRCellResourceCoordinationRequestIEs, error) {

	msg := &xnappducontentsv1.EUTraNRCellResourceCoordinationRequestIEs{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateEUTraNRCellResourceCoordinationRequestIEs() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateEUTraNRCellResourceCoordinationRequest(protocolIes []*xnappducontentsv1.EUTraNRCellResourceCoordinationRequestIEs) (*xnappducontentsv1.EUTraNRCellResourceCoordinationRequest, error) {

	msg := &xnappducontentsv1.EUTraNRCellResourceCoordinationRequest{}
	msg.ProtocolIes = protocolIes

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateEUTraNRCellResourceCoordinationRequest() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateResourceCoordRequestngeNbinitiated(dataTrafficResourceIndication *xnapiesv1.DataTrafficResourceIndication, spectrumSharingGroupID *xnapiesv1.SpectrumSharingGroupId) (*xnappducontentsv1.ResourceCoordRequestngeNbinitiated, error) {

	msg := &xnappducontentsv1.ResourceCoordRequestngeNbinitiated{}
	msg.DataTrafficResourceIndication = dataTrafficResourceIndication
	msg.SpectrumSharingGroupId = spectrumSharingGroupID

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateResourceCoordRequestngeNbinitiated() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateResourceCoordRequestgNbinitiated(dataTrafficResourceIndication *xnapiesv1.DataTrafficResourceIndication, spectrumSharingGroupID *xnapiesv1.SpectrumSharingGroupId) (*xnappducontentsv1.ResourceCoordRequestgNbinitiated, error) {

	msg := &xnappducontentsv1.ResourceCoordRequestgNbinitiated{}
	msg.DataTrafficResourceIndication = dataTrafficResourceIndication
	msg.SpectrumSharingGroupId = spectrumSharingGroupID

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateResourceCoordRequestgNbinitiated() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateEUTraNRCellResourceCoordinationResponseIEs(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, value *xnappducontentsv1.EUTraNRCellResourceCoordinationResponseIEsValue) (*xnappducontentsv1.EUTraNRCellResourceCoordinationResponseIEs, error) {

	msg := &xnappducontentsv1.EUTraNRCellResourceCoordinationResponseIEs{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateEUTraNRCellResourceCoordinationResponseIEs() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateEUTraNRCellResourceCoordinationResponse(protocolIes []*xnappducontentsv1.EUTraNRCellResourceCoordinationResponseIEs) (*xnappducontentsv1.EUTraNRCellResourceCoordinationResponse, error) {

	msg := &xnappducontentsv1.EUTraNRCellResourceCoordinationResponse{}
	msg.ProtocolIes = protocolIes

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateEUTraNRCellResourceCoordinationResponse() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateResourceCoordResponsengeNbinitiated(dataTrafficResourceIndication *xnapiesv1.DataTrafficResourceIndication, spectrumSharingGroupID *xnapiesv1.SpectrumSharingGroupId) (*xnappducontentsv1.ResourceCoordResponsengeNbinitiated, error) {

	msg := &xnappducontentsv1.ResourceCoordResponsengeNbinitiated{}
	msg.DataTrafficResourceIndication = dataTrafficResourceIndication
	msg.SpectrumSharingGroupId = spectrumSharingGroupID

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateResourceCoordResponsengeNbinitiated() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateResourceCoordResponsegNbinitiated(dataTrafficResourceIndication *xnapiesv1.DataTrafficResourceIndication, spectrumSharingGroupID *xnapiesv1.SpectrumSharingGroupId) (*xnappducontentsv1.ResourceCoordResponsegNbinitiated, error) {

	msg := &xnappducontentsv1.ResourceCoordResponsegNbinitiated{}
	msg.DataTrafficResourceIndication = dataTrafficResourceIndication
	msg.SpectrumSharingGroupId = spectrumSharingGroupID

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateResourceCoordResponsegNbinitiated() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateSecondaryRatdataUsageReportIEs(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, value *xnappducontentsv1.SecondaryRatdataUsageReportIEsValue) (*xnappducontentsv1.SecondaryRatdataUsageReportIEs, error) {

	msg := &xnappducontentsv1.SecondaryRatdataUsageReportIEs{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSecondaryRatdataUsageReportIEs() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateSecondaryRatdataUsageReport(protocolIes []*xnappducontentsv1.SecondaryRatdataUsageReportIEs) (*xnappducontentsv1.SecondaryRatdataUsageReport, error) {

	msg := &xnappducontentsv1.SecondaryRatdataUsageReport{}
	msg.ProtocolIes = protocolIes

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSecondaryRatdataUsageReport() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateXnRemovalRequestIEs(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, value *xnappducontentsv1.XnRemovalRequestIEsValue) (*xnappducontentsv1.XnRemovalRequestIEs, error) {

	msg := &xnappducontentsv1.XnRemovalRequestIEs{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateXnRemovalRequestIEs() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateXnRemovalRequest(protocolIes []*xnappducontentsv1.XnRemovalRequestIEs) (*xnappducontentsv1.XnRemovalRequest, error) {

	msg := &xnappducontentsv1.XnRemovalRequest{}
	msg.ProtocolIes = protocolIes

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateXnRemovalRequest() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateXnRemovalResponseIEs(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, value *xnappducontentsv1.XnRemovalResponseIEsValue) (*xnappducontentsv1.XnRemovalResponseIEs, error) {

	msg := &xnappducontentsv1.XnRemovalResponseIEs{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateXnRemovalResponseIEs() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateXnRemovalResponse(protocolIes []*xnappducontentsv1.XnRemovalResponseIEs) (*xnappducontentsv1.XnRemovalResponse, error) {

	msg := &xnappducontentsv1.XnRemovalResponse{}
	msg.ProtocolIes = protocolIes

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateXnRemovalResponse() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateXnRemovalFailureIEs(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, value *xnappducontentsv1.XnRemovalFailureIEsValue) (*xnappducontentsv1.XnRemovalFailureIEs, error) {

	msg := &xnappducontentsv1.XnRemovalFailureIEs{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateXnRemovalFailureIEs() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateXnRemovalFailure(protocolIes []*xnappducontentsv1.XnRemovalFailureIEs) (*xnappducontentsv1.XnRemovalFailure, error) {

	msg := &xnappducontentsv1.XnRemovalFailure{}
	msg.ProtocolIes = protocolIes

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateXnRemovalFailure() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateCellActivationRequestIEs(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, value *xnappducontentsv1.CellActivationRequestIEsValue) (*xnappducontentsv1.CellActivationRequestIEs, error) {

	msg := &xnappducontentsv1.CellActivationRequestIEs{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateCellActivationRequestIEs() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateCellActivationRequest(protocolIes []*xnappducontentsv1.CellActivationRequestIEs) (*xnappducontentsv1.CellActivationRequest, error) {

	msg := &xnappducontentsv1.CellActivationRequest{}
	msg.ProtocolIes = protocolIes

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateCellActivationRequest() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateNrcellsServedCellsToActivate(nrCells []*xnapiesv1.NrCGi) (*xnappducontentsv1.NrcellsServedCellsToActivate, error) {

	msg := &xnappducontentsv1.NrcellsServedCellsToActivate{}
	msg.NrCells = nrCells

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNrcellsServedCellsToActivate() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateEutracellsServedCellsToActivate(eUtraCells []*xnapiesv1.EUTraCGi) (*xnappducontentsv1.EutracellsServedCellsToActivate, error) {

	msg := &xnappducontentsv1.EutracellsServedCellsToActivate{}
	msg.EUtraCells = eUtraCells

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateEutracellsServedCellsToActivate() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateCellActivationResponseIEs(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, value *xnappducontentsv1.CellActivationResponseIEsValue) (*xnappducontentsv1.CellActivationResponseIEs, error) {

	msg := &xnappducontentsv1.CellActivationResponseIEs{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateCellActivationResponseIEs() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateCellActivationResponse(protocolIes []*xnappducontentsv1.CellActivationResponseIEs) (*xnappducontentsv1.CellActivationResponse, error) {

	msg := &xnappducontentsv1.CellActivationResponse{}
	msg.ProtocolIes = protocolIes

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateCellActivationResponse() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateNrcellsActivatedServedCells(nrCells []*xnapiesv1.NrCGi) (*xnappducontentsv1.NrcellsActivatedServedCells, error) {

	msg := &xnappducontentsv1.NrcellsActivatedServedCells{}
	msg.NrCells = nrCells

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNrcellsActivatedServedCells() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateEutracellsActivatedServedCells(eUtraCells []*xnapiesv1.EUTraCGi) (*xnappducontentsv1.EutracellsActivatedServedCells, error) {

	msg := &xnappducontentsv1.EutracellsActivatedServedCells{}
	msg.EUtraCells = eUtraCells

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateEutracellsActivatedServedCells() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateCellActivationFailureIEs(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, value *xnappducontentsv1.CellActivationFailureIEsValue) (*xnappducontentsv1.CellActivationFailureIEs, error) {

	msg := &xnappducontentsv1.CellActivationFailureIEs{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateCellActivationFailureIEs() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateCellActivationFailure(protocolIes []*xnappducontentsv1.CellActivationFailureIEs) (*xnappducontentsv1.CellActivationFailure, error) {

	msg := &xnappducontentsv1.CellActivationFailure{}
	msg.ProtocolIes = protocolIes

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateCellActivationFailure() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateResetRequestIEs(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, value *xnappducontentsv1.ResetRequestIEsValue) (*xnappducontentsv1.ResetRequestIEs, error) {

	msg := &xnappducontentsv1.ResetRequestIEs{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateResetRequestIEs() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateResetRequest(protocolIes []*xnappducontentsv1.ResetRequestIEs) (*xnappducontentsv1.ResetRequest, error) {

	msg := &xnappducontentsv1.ResetRequest{}
	msg.ProtocolIes = protocolIes

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateResetRequest() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateResetResponseIEs(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, value *xnappducontentsv1.ResetResponseIEsValue) (*xnappducontentsv1.ResetResponseIEs, error) {

	msg := &xnappducontentsv1.ResetResponseIEs{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateResetResponseIEs() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateResetResponse(protocolIes []*xnappducontentsv1.ResetResponseIEs) (*xnappducontentsv1.ResetResponse, error) {

	msg := &xnappducontentsv1.ResetResponse{}
	msg.ProtocolIes = protocolIes

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateResetResponse() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateErrorIndicationIEs(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, value *xnappducontentsv1.ErrorIndicationIEsValue) (*xnappducontentsv1.ErrorIndicationIEs, error) {

	msg := &xnappducontentsv1.ErrorIndicationIEs{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateErrorIndicationIEs() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateErrorIndication(protocolIes []*xnappducontentsv1.ErrorIndicationIEs) (*xnappducontentsv1.ErrorIndication, error) {

	msg := &xnappducontentsv1.ErrorIndication{}
	msg.ProtocolIes = protocolIes

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateErrorIndication() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreatePrivateMessage(privateIes []*xnappducontentsv1.PrivateMessageIEs) (*xnappducontentsv1.PrivateMessage, error) {

	msg := &xnappducontentsv1.PrivateMessage{}
	msg.PrivateIes = privateIes

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePrivateMessage() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateTraceStartIes(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, value *xnappducontentsv1.TraceStartIesValue) (*xnappducontentsv1.TraceStartIes, error) {

	msg := &xnappducontentsv1.TraceStartIes{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateTraceStartIes() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateTraceStart(protocolIes []*xnappducontentsv1.TraceStartIes) (*xnappducontentsv1.TraceStart, error) {

	msg := &xnappducontentsv1.TraceStart{}
	msg.ProtocolIes = protocolIes

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateTraceStart() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateDeactivateTraceIes(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, value *xnappducontentsv1.DeactivateTraceIesValue) (*xnappducontentsv1.DeactivateTraceIes, error) {

	msg := &xnappducontentsv1.DeactivateTraceIes{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateDeactivateTraceIes() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateDeactivateTrace(protocolIes []*xnappducontentsv1.DeactivateTraceIes) (*xnappducontentsv1.DeactivateTrace, error) {

	msg := &xnappducontentsv1.DeactivateTrace{}
	msg.ProtocolIes = protocolIes

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateDeactivateTrace() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateFailureIndication(protocolIes []*xnappducontentsv1.FailureIndicationIEs) (*xnappducontentsv1.FailureIndication, error) {

	msg := &xnappducontentsv1.FailureIndication{}
	msg.ProtocolIes = protocolIes

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateFailureIndication() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateHandoverReportIEs(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, value *xnappducontentsv1.HandoverReportIEsValue) (*xnappducontentsv1.HandoverReportIEs, error) {

	msg := &xnappducontentsv1.HandoverReportIEs{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateHandoverReportIEs() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateHandoverReport(protocolIes []*xnappducontentsv1.HandoverReportIEs) (*xnappducontentsv1.HandoverReport, error) {

	msg := &xnappducontentsv1.HandoverReport{}
	msg.ProtocolIes = protocolIes

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateHandoverReport() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateResourceStatusRequestIEs(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, value *xnappducontentsv1.ResourceStatusRequestIEsValue) (*xnappducontentsv1.ResourceStatusRequestIEs, error) {

	msg := &xnappducontentsv1.ResourceStatusRequestIEs{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateResourceStatusRequestIEs() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateResourceStatusRequest(protocolIes []*xnappducontentsv1.ResourceStatusRequestIEs) (*xnappducontentsv1.ResourceStatusRequest, error) {

	msg := &xnappducontentsv1.ResourceStatusRequest{}
	msg.ProtocolIes = protocolIes

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateResourceStatusRequest() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateResourceStatusResponseIEs(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, value *xnappducontentsv1.ResourceStatusResponseIEsValue) (*xnappducontentsv1.ResourceStatusResponseIEs, error) {

	msg := &xnappducontentsv1.ResourceStatusResponseIEs{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateResourceStatusResponseIEs() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateResourceStatusResponse(protocolIes []*xnappducontentsv1.ResourceStatusResponseIEs) (*xnappducontentsv1.ResourceStatusResponse, error) {

	msg := &xnappducontentsv1.ResourceStatusResponse{}
	msg.ProtocolIes = protocolIes

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateResourceStatusResponse() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateResourceStatusFailureIEs(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, value *xnappducontentsv1.ResourceStatusFailureIEsValue) (*xnappducontentsv1.ResourceStatusFailureIEs, error) {

	msg := &xnappducontentsv1.ResourceStatusFailureIEs{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateResourceStatusFailureIEs() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateResourceStatusFailure(protocolIes []*xnappducontentsv1.ResourceStatusFailureIEs) (*xnappducontentsv1.ResourceStatusFailure, error) {

	msg := &xnappducontentsv1.ResourceStatusFailure{}
	msg.ProtocolIes = protocolIes

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateResourceStatusFailure() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateResourceStatusUpdateIEs(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, value *xnappducontentsv1.ResourceStatusUpdateIEsValue) (*xnappducontentsv1.ResourceStatusUpdateIEs, error) {

	msg := &xnappducontentsv1.ResourceStatusUpdateIEs{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateResourceStatusUpdateIEs() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateResourceStatusUpdate(protocolIes []*xnappducontentsv1.ResourceStatusUpdateIEs) (*xnappducontentsv1.ResourceStatusUpdate, error) {

	msg := &xnappducontentsv1.ResourceStatusUpdate{}
	msg.ProtocolIes = protocolIes

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateResourceStatusUpdate() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateMobilityChangeRequestIEs(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, value *xnappducontentsv1.MobilityChangeRequestIEsValue) (*xnappducontentsv1.MobilityChangeRequestIEs, error) {

	msg := &xnappducontentsv1.MobilityChangeRequestIEs{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateMobilityChangeRequestIEs() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateMobilityChangeRequest(protocolIes []*xnappducontentsv1.MobilityChangeRequestIEs) (*xnappducontentsv1.MobilityChangeRequest, error) {

	msg := &xnappducontentsv1.MobilityChangeRequest{}
	msg.ProtocolIes = protocolIes

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateMobilityChangeRequest() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateMobilityChangeAcknowledgeIEs(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, value *xnappducontentsv1.MobilityChangeAcknowledgeIEsValue) (*xnappducontentsv1.MobilityChangeAcknowledgeIEs, error) {

	msg := &xnappducontentsv1.MobilityChangeAcknowledgeIEs{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateMobilityChangeAcknowledgeIEs() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateMobilityChangeAcknowledge(protocolIes []*xnappducontentsv1.MobilityChangeAcknowledgeIEs) (*xnappducontentsv1.MobilityChangeAcknowledge, error) {

	msg := &xnappducontentsv1.MobilityChangeAcknowledge{}
	msg.ProtocolIes = protocolIes

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateMobilityChangeAcknowledge() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateMobilityChangeFailureIEs(ID *xnapcommondatatypesv1.ProtocolIeID, criticality xnapcommondatatypesv1.Criticality, value *xnappducontentsv1.MobilityChangeFailureIEsValue) (*xnappducontentsv1.MobilityChangeFailureIEs, error) {

	msg := &xnappducontentsv1.MobilityChangeFailureIEs{}
	msg.Id = ID
	msg.Criticality = criticality
	msg.Value = value

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateMobilityChangeFailureIEs() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateMobilityChangeFailure(protocolIes []*xnappducontentsv1.MobilityChangeFailureIEs) (*xnappducontentsv1.MobilityChangeFailure, error) {

	msg := &xnappducontentsv1.MobilityChangeFailure{}
	msg.ProtocolIes = protocolIes

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateMobilityChangeFailure() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateAccessAndMobilityIndication(protocolIes []*xnappducontentsv1.AccessAndMobilityIndicationIEs) (*xnappducontentsv1.AccessAndMobilityIndication, error) {

	msg := &xnappducontentsv1.AccessAndMobilityIndication{}
	msg.ProtocolIes = protocolIes

	//if err := msg.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateAccessAndMobilityIndication() error validating PDU %s", err.Error())
	//}

	return msg, nil
}

func CreateAreaScopeOfMdtNRCellBased(cellBased *xnapiesv1.CellBasedMdtNR) (*xnapiesv1.AreaScopeOfMdtNR, error) {

	item := &xnapiesv1.AreaScopeOfMdtNR{
		AreaScopeOfMdtNr: &xnapiesv1.AreaScopeOfMdtNR_CellBased{
			CellBased: cellBased,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateAreaScopeOfMdtNRCellBased() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateAreaScopeOfMdtNRTAbased(tAbased *xnapiesv1.TabasedMdt) (*xnapiesv1.AreaScopeOfMdtNR, error) {

	item := &xnapiesv1.AreaScopeOfMdtNR{
		AreaScopeOfMdtNr: &xnapiesv1.AreaScopeOfMdtNR_TAbased{
			TAbased: tAbased,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateAreaScopeOfMdtNRTAbased() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateAreaScopeOfMdtNRTAibased(tAibased *xnapiesv1.TaibasedMdt) (*xnapiesv1.AreaScopeOfMdtNR, error) {

	item := &xnapiesv1.AreaScopeOfMdtNR{
		AreaScopeOfMdtNr: &xnapiesv1.AreaScopeOfMdtNR_TAibased{
			TAibased: tAibased,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateAreaScopeOfMdtNRTAibased() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateAreaScopeOfMdtEUtraCellBased(cellBased *xnapiesv1.CellBasedMdtEUtra) (*xnapiesv1.AreaScopeOfMdtEUtra, error) {

	item := &xnapiesv1.AreaScopeOfMdtEUtra{
		AreaScopeOfMdtEutra: &xnapiesv1.AreaScopeOfMdtEUtra_CellBased{
			CellBased: cellBased,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateAreaScopeOfMdtEUtraCellBased() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateAreaScopeOfMdtEUtraTAbased(tAbased *xnapiesv1.TabasedMdt) (*xnapiesv1.AreaScopeOfMdtEUtra, error) {

	item := &xnapiesv1.AreaScopeOfMdtEUtra{
		AreaScopeOfMdtEutra: &xnapiesv1.AreaScopeOfMdtEUtra_TAbased{
			TAbased: tAbased,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateAreaScopeOfMdtEUtraTAbased() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateAreaScopeOfMdtEUtraTAibased(tAibased *xnapiesv1.TaibasedMdt) (*xnapiesv1.AreaScopeOfMdtEUtra, error) {

	item := &xnapiesv1.AreaScopeOfMdtEUtra{
		AreaScopeOfMdtEutra: &xnapiesv1.AreaScopeOfMdtEUtra_TAibased{
			TAibased: tAibased,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateAreaScopeOfMdtEUtraTAibased() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateBplmnIDInfoNRItemExtIesExtensionIDConfiguredTacindication(IDConfiguredTacindication xnapiesv1.ConfiguredTacindication) (*xnapiesv1.BplmnIDInfoNRItemExtIesExtension, error) {

	item := &xnapiesv1.BplmnIDInfoNRItemExtIesExtension{
		BplmnIdInfoNrItemExtIes: &xnapiesv1.BplmnIDInfoNRItemExtIesExtension_IdConfiguredTacindication{
			IdConfiguredTacindication: IDConfiguredTacindication,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateBplmnIDInfoNRItemExtIesExtensionIDConfiguredTacindication() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateBplmnIDInfoNRItemExtIesExtensionIDNpnBroadcastInformation(IDNpnBroadcastInformation *xnapiesv1.NpnBroadcastInformation) (*xnapiesv1.BplmnIDInfoNRItemExtIesExtension, error) {

	item := &xnapiesv1.BplmnIDInfoNRItemExtIesExtension{
		BplmnIdInfoNrItemExtIes: &xnapiesv1.BplmnIDInfoNRItemExtIesExtension_IdNpnBroadcastInformation{
			IdNpnBroadcastInformation: IDNpnBroadcastInformation,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateBplmnIDInfoNRItemExtIesExtensionIDNpnBroadcastInformation() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateBroadcastPlmninTaisupportItemExtIesExtensionIDNpnSupport(IDNpnSupport *xnapiesv1.NpnSupport) (*xnapiesv1.BroadcastPlmninTaisupportItemExtIesExtension, error) {

	item := &xnapiesv1.BroadcastPlmninTaisupportItemExtIesExtension{
		BroadcastPlmninTaisupportItemExtIes: &xnapiesv1.BroadcastPlmninTaisupportItemExtIesExtension_IdNpnSupport{
			IdNpnSupport: IDNpnSupport,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateBroadcastPlmninTaisupportItemExtIesExtensionIDNpnSupport() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateBroadcastPlmninTaisupportItemExtIesExtensionIDExtendedTaisliceSupportList(IDExtendedTaisliceSupportList *xnapiesv1.ExtendedSliceSupportList) (*xnapiesv1.BroadcastPlmninTaisupportItemExtIesExtension, error) {

	item := &xnapiesv1.BroadcastPlmninTaisupportItemExtIesExtension{
		BroadcastPlmninTaisupportItemExtIes: &xnapiesv1.BroadcastPlmninTaisupportItemExtIesExtension_IdExtendedTaisliceSupportList{
			IdExtendedTaisliceSupportList: IDExtendedTaisliceSupportList,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateBroadcastPlmninTaisupportItemExtIesExtensionIDExtendedTaisliceSupportList() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateCauseRadioNetwork(radioNetwork xnapiesv1.CauseRadioNetworkLayer) (*xnapiesv1.Cause, error) {

	item := &xnapiesv1.Cause{
		Cause: &xnapiesv1.Cause_RadioNetwork{
			RadioNetwork: radioNetwork,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateCauseRadioNetwork() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateCauseTransport(transport xnapiesv1.CauseTransportLayer) (*xnapiesv1.Cause, error) {

	item := &xnapiesv1.Cause{
		Cause: &xnapiesv1.Cause_Transport{
			Transport: transport,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateCauseTransport() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateCauseProtocol(protocol xnapiesv1.CauseProtocol) (*xnapiesv1.Cause, error) {

	item := &xnapiesv1.Cause{
		Cause: &xnapiesv1.Cause_Protocol{
			Protocol: protocol,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateCauseProtocol() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateCauseMisc(misc xnapiesv1.CauseMisc) (*xnapiesv1.Cause, error) {

	item := &xnapiesv1.Cause{
		Cause: &xnapiesv1.Cause_Misc{
			Misc: misc,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateCauseMisc() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateCauseChoiceExtension(choiceExtension *xnapiesv1.CauseExtIes) (*xnapiesv1.Cause, error) {

	item := &xnapiesv1.Cause{
		Cause: &xnapiesv1.Cause_ChoiceExtension{
			ChoiceExtension: choiceExtension,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateCauseChoiceExtension() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateCellAssistanceInfoNRLimitedNrList(limitedNrList *xnapiesv1.LimitedNrListCellAssistanceInfoNR) (*xnapiesv1.CellAssistanceInfoNR, error) {

	item := &xnapiesv1.CellAssistanceInfoNR{
		CellAssistanceInfoNr: &xnapiesv1.CellAssistanceInfoNR_LimitedNrList{
			LimitedNrList: limitedNrList,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateCellAssistanceInfoNRLimitedNrList() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateCellAssistanceInfoNRFullList(fullList xnapiesv1.FullListCellAssistanceInfoNr) (*xnapiesv1.CellAssistanceInfoNR, error) {

	item := &xnapiesv1.CellAssistanceInfoNR{
		CellAssistanceInfoNr: &xnapiesv1.CellAssistanceInfoNR_FullList{
			FullList: fullList,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateCellAssistanceInfoNRFullList() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateCellAssistanceInfoNRChoiceExtension(choiceExtension *xnapiesv1.CellAssistanceInfoNRExtIes) (*xnapiesv1.CellAssistanceInfoNR, error) {

	item := &xnapiesv1.CellAssistanceInfoNR{
		CellAssistanceInfoNr: &xnapiesv1.CellAssistanceInfoNR_ChoiceExtension{
			ChoiceExtension: choiceExtension,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateCellAssistanceInfoNRChoiceExtension() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateCellAssistanceInfoEUtraLimitedEutraList(limitedEutraList *xnapiesv1.LimitedEutraListCellAssistanceInfoEUtra) (*xnapiesv1.CellAssistanceInfoEUtra, error) {

	item := &xnapiesv1.CellAssistanceInfoEUtra{
		CellAssistanceInfoEutra: &xnapiesv1.CellAssistanceInfoEUtra_LimitedEutraList{
			LimitedEutraList: limitedEutraList,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateCellAssistanceInfoEUtraLimitedEutraList() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateCellAssistanceInfoEUtraFullList(fullList xnapiesv1.FullListCellAssistanceInfoEutra) (*xnapiesv1.CellAssistanceInfoEUtra, error) {

	item := &xnapiesv1.CellAssistanceInfoEUtra{
		CellAssistanceInfoEutra: &xnapiesv1.CellAssistanceInfoEUtra_FullList{
			FullList: fullList,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateCellAssistanceInfoEUtraFullList() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateCellAssistanceInfoEUtraChoiceExtension(choiceExtension *xnapiesv1.CellAssistanceInfoEUtraExtIes) (*xnapiesv1.CellAssistanceInfoEUtra, error) {

	item := &xnapiesv1.CellAssistanceInfoEUtra{
		CellAssistanceInfoEutra: &xnapiesv1.CellAssistanceInfoEUtra_ChoiceExtension{
			ChoiceExtension: choiceExtension,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateCellAssistanceInfoEUtraChoiceExtension() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateCellTypeChoiceNgRanEUtra(ngRanEUtra *xnapiesv1.EUTraCellIdentity) (*xnapiesv1.CellTypeChoice, error) {

	item := &xnapiesv1.CellTypeChoice{
		CellTypeChoice: &xnapiesv1.CellTypeChoice_NgRanEUtra{
			NgRanEUtra: ngRanEUtra,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateCellTypeChoiceNgRanEUtra() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateCellTypeChoiceNgRanNr(ngRanNr *xnapiesv1.NrCellIdentity) (*xnapiesv1.CellTypeChoice, error) {

	item := &xnapiesv1.CellTypeChoice{
		CellTypeChoice: &xnapiesv1.CellTypeChoice_NgRanNr{
			NgRanNr: ngRanNr,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateCellTypeChoiceNgRanNr() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateCellTypeChoiceEUtran(eUtran *xnapiesv1.EUTraCellIdentity) (*xnapiesv1.CellTypeChoice, error) {

	item := &xnapiesv1.CellTypeChoice{
		CellTypeChoice: &xnapiesv1.CellTypeChoice_EUtran{
			EUtran: eUtran,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateCellTypeChoiceEUtran() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateCellTypeChoiceChoiceExtension(choiceExtension *xnapiesv1.CellTypeChoiceExtIes) (*xnapiesv1.CellTypeChoice, error) {

	item := &xnapiesv1.CellTypeChoice{
		CellTypeChoice: &xnapiesv1.CellTypeChoice_ChoiceExtension{
			ChoiceExtension: choiceExtension,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateCellTypeChoiceChoiceExtension() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateCptransportLayerInformationEndpointIpaddress(endpointIpaddress *xnapiesv1.TransportLayerAddress) (*xnapiesv1.CptransportLayerInformation, error) {

	item := &xnapiesv1.CptransportLayerInformation{
		CptransportLayerInformation: &xnapiesv1.CptransportLayerInformation_EndpointIpaddress{
			EndpointIpaddress: endpointIpaddress,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateCptransportLayerInformationEndpointIpaddress() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateCptransportLayerInformationChoiceExtension(choiceExtension *xnapiesv1.CptransportLayerInformationExtIes) (*xnapiesv1.CptransportLayerInformation, error) {

	item := &xnapiesv1.CptransportLayerInformation{
		CptransportLayerInformation: &xnapiesv1.CptransportLayerInformation_ChoiceExtension{
			ChoiceExtension: choiceExtension,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateCptransportLayerInformationChoiceExtension() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateXnUaddressInfoperPdusessionItemExtIesExtensionIDSecondarydataForwardingInfoFromTargetList(IDSecondarydataForwardingInfoFromTargetList *xnapiesv1.SecondarydataForwardingInfoFromTargetList) (*xnapiesv1.XnUaddressInfoperPdusessionItemExtIesExtension, error) {

	item := &xnapiesv1.XnUaddressInfoperPdusessionItemExtIesExtension{
		XnUaddressInfoperPdusessionItemExtIes: &xnapiesv1.XnUaddressInfoperPdusessionItemExtIesExtension_IdSecondarydataForwardingInfoFromTargetList{
			IdSecondarydataForwardingInfoFromTargetList: IDSecondarydataForwardingInfoFromTargetList,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateXnUaddressInfoperPdusessionItemExtIesExtensionIDSecondarydataForwardingInfoFromTargetList() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateXnUaddressInfoperPdusessionItemExtIesExtensionIDDrbIDsTakenintouse(IDDrbIDsTakenintouse *xnapiesv1.DrbList) (*xnapiesv1.XnUaddressInfoperPdusessionItemExtIesExtension, error) {

	item := &xnapiesv1.XnUaddressInfoperPdusessionItemExtIesExtension{
		XnUaddressInfoperPdusessionItemExtIes: &xnapiesv1.XnUaddressInfoperPdusessionItemExtIesExtension_IdDrbIdsTakenintouse{
			IdDrbIdsTakenintouse: IDDrbIDsTakenintouse,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateXnUaddressInfoperPdusessionItemExtIesExtensionIDDrbIDsTakenintouse() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateXnUaddressInfoperPdusessionItemExtIesExtensionIDDataForwardingInfoFromTargetEUtrannode(IDDataForwardingInfoFromTargetEUtrannode *xnapiesv1.DataForwardingInfoFromTargetEUTrannode) (*xnapiesv1.XnUaddressInfoperPdusessionItemExtIesExtension, error) {

	item := &xnapiesv1.XnUaddressInfoperPdusessionItemExtIesExtension{
		XnUaddressInfoperPdusessionItemExtIes: &xnapiesv1.XnUaddressInfoperPdusessionItemExtIesExtension_IdDataForwardingInfoFromTargetEUtrannode{
			IdDataForwardingInfoFromTargetEUtrannode: IDDataForwardingInfoFromTargetEUtrannode,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateXnUaddressInfoperPdusessionItemExtIesExtensionIDDataForwardingInfoFromTargetEUtrannode() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateQoSflowsToBeForwardedItemExtIesExtensionIDUlforwardingProposal(IDUlforwardingProposal xnapiesv1.UlforwardingProposal) (*xnapiesv1.QoSflowsToBeForwardedItemExtIesExtension, error) {

	item := &xnapiesv1.QoSflowsToBeForwardedItemExtIesExtension{
		QoSflowsToBeForwardedItemExtIes: &xnapiesv1.QoSflowsToBeForwardedItemExtIesExtension_IdUlforwardingProposal{
			IdUlforwardingProposal: IDUlforwardingProposal,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateQoSflowsToBeForwardedItemExtIesExtensionIDUlforwardingProposal() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateQoSflowsToBeForwardedItemExtIesExtensionIDSourceDlforwardingIpaddress(IDSourceDlforwardingIpaddress *xnapiesv1.TransportLayerAddress) (*xnapiesv1.QoSflowsToBeForwardedItemExtIesExtension, error) {

	item := &xnapiesv1.QoSflowsToBeForwardedItemExtIesExtension{
		QoSflowsToBeForwardedItemExtIes: &xnapiesv1.QoSflowsToBeForwardedItemExtIesExtension_IdSourceDlforwardingIpaddress{
			IdSourceDlforwardingIpaddress: IDSourceDlforwardingIpaddress,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateQoSflowsToBeForwardedItemExtIesExtensionIDSourceDlforwardingIpaddress() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateQoSflowsToBeForwardedItemExtIesExtensionIDSourceNodeDlforwardingIpaddress(IDSourceNodeDlforwardingIpaddress *xnapiesv1.TransportLayerAddress) (*xnapiesv1.QoSflowsToBeForwardedItemExtIesExtension, error) {

	item := &xnapiesv1.QoSflowsToBeForwardedItemExtIesExtension{
		QoSflowsToBeForwardedItemExtIes: &xnapiesv1.QoSflowsToBeForwardedItemExtIesExtension_IdSourceNodeDlforwardingIpaddress{
			IdSourceNodeDlforwardingIpaddress: IDSourceNodeDlforwardingIpaddress,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateQoSflowsToBeForwardedItemExtIesExtensionIDSourceNodeDlforwardingIpaddress() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateDlcountChoiceCount12Bits(count12Bits *xnapiesv1.CountPDcpSN12) (*xnapiesv1.DlcountChoice, error) {

	item := &xnapiesv1.DlcountChoice{
		DlcountChoice: &xnapiesv1.DlcountChoice_Count12Bits{
			Count12Bits: count12Bits,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateDlcountChoiceCount12Bits() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateDlcountChoiceCount18Bits(count18Bits *xnapiesv1.CountPDcpSN18) (*xnapiesv1.DlcountChoice, error) {

	item := &xnapiesv1.DlcountChoice{
		DlcountChoice: &xnapiesv1.DlcountChoice_Count18Bits{
			Count18Bits: count18Bits,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateDlcountChoiceCount18Bits() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateDlcountChoiceChoiceExtension(choiceExtension *xnapiesv1.DlcountChoiceExtIes) (*xnapiesv1.DlcountChoice, error) {

	item := &xnapiesv1.DlcountChoice{
		DlcountChoice: &xnapiesv1.DlcountChoice_ChoiceExtension{
			ChoiceExtension: choiceExtension,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateDlcountChoiceChoiceExtension() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateDrbbstatusTransferChoicePdcpSn12Bits(pdcpSn12Bits *xnapiesv1.DrbbstatusTransfer12BitsSn) (*xnapiesv1.DrbbstatusTransferChoice, error) {

	item := &xnapiesv1.DrbbstatusTransferChoice{
		DrbbstatusTransferChoice: &xnapiesv1.DrbbstatusTransferChoice_PdcpSn_12Bits{
			PdcpSn_12Bits: pdcpSn12Bits,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateDrbbstatusTransferChoicePdcpSn12Bits() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateDrbbstatusTransferChoicePdcpSn18Bits(pdcpSn18Bits *xnapiesv1.DrbbstatusTransfer18BitsSn) (*xnapiesv1.DrbbstatusTransferChoice, error) {

	item := &xnapiesv1.DrbbstatusTransferChoice{
		DrbbstatusTransferChoice: &xnapiesv1.DrbbstatusTransferChoice_PdcpSn_18Bits{
			PdcpSn_18Bits: pdcpSn18Bits,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateDrbbstatusTransferChoicePdcpSn18Bits() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateDrbbstatusTransferChoiceChoiceExtension(choiceExtension *xnapiesv1.DrbbstatusTransferChoiceExtIes) (*xnapiesv1.DrbbstatusTransferChoice, error) {

	item := &xnapiesv1.DrbbstatusTransferChoice{
		DrbbstatusTransferChoice: &xnapiesv1.DrbbstatusTransferChoice_ChoiceExtension{
			ChoiceExtension: choiceExtension,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateDrbbstatusTransferChoiceChoiceExtension() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateDynamic5QIDescriptorExtIesExtensionIDExtendedPacketDelayBudget(IDExtendedPacketDelayBudget *xnapiesv1.ExtendedPacketDelayBudget) (*xnapiesv1.Dynamic5QidescriptorExtIesExtension, error) {

	item := &xnapiesv1.Dynamic5QidescriptorExtIesExtension{
		Dynamic5QidescriptorExtIes: &xnapiesv1.Dynamic5QidescriptorExtIesExtension_IdExtendedPacketDelayBudget{
			IdExtendedPacketDelayBudget: IDExtendedPacketDelayBudget,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateDynamic5QIDescriptorExtIesExtensionIDExtendedPacketDelayBudget() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateDynamic5QIDescriptorExtIesExtensionIDCnpacketDelayBudgetDownlink(IDCnpacketDelayBudgetDownlink *xnapiesv1.ExtendedPacketDelayBudget) (*xnapiesv1.Dynamic5QidescriptorExtIesExtension, error) {

	item := &xnapiesv1.Dynamic5QidescriptorExtIesExtension{
		Dynamic5QidescriptorExtIes: &xnapiesv1.Dynamic5QidescriptorExtIesExtension_IdCnpacketDelayBudgetDownlink{
			IdCnpacketDelayBudgetDownlink: IDCnpacketDelayBudgetDownlink,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateDynamic5QIDescriptorExtIesExtensionIDCnpacketDelayBudgetDownlink() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateDynamic5QIDescriptorExtIesExtensionIDCnpacketDelayBudgetUplink(IDCnpacketDelayBudgetUplink *xnapiesv1.ExtendedPacketDelayBudget) (*xnapiesv1.Dynamic5QidescriptorExtIesExtension, error) {

	item := &xnapiesv1.Dynamic5QidescriptorExtIesExtension{
		Dynamic5QidescriptorExtIes: &xnapiesv1.Dynamic5QidescriptorExtIesExtension_IdCnpacketDelayBudgetUplink{
			IdCnpacketDelayBudgetUplink: IDCnpacketDelayBudgetUplink,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateDynamic5QIDescriptorExtIesExtensionIDCnpacketDelayBudgetUplink() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateEventTypeTriggerOutOfCoverage(outOfCoverage xnapiesv1.OutOfCoverageEventTypeTrigger) (*xnapiesv1.EventTypeTrigger, error) {

	item := &xnapiesv1.EventTypeTrigger{
		EventTypeTrigger: &xnapiesv1.EventTypeTrigger_OutOfCoverage{
			OutOfCoverage: outOfCoverage,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateEventTypeTriggerOutOfCoverage() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateEventTypeTriggerEventL1(eventL1 *xnapiesv1.EventL1) (*xnapiesv1.EventTypeTrigger, error) {

	item := &xnapiesv1.EventTypeTrigger{
		EventTypeTrigger: &xnapiesv1.EventTypeTrigger_EventL1{
			EventL1: eventL1,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateEventTypeTriggerEventL1() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateEventTypeTriggerChoiceExtensions(choiceExtensions *xnapiesv1.EventTypeTriggerExtIes) (*xnapiesv1.EventTypeTrigger, error) {

	item := &xnapiesv1.EventTypeTrigger{
		EventTypeTrigger: &xnapiesv1.EventTypeTrigger_ChoiceExtensions{
			ChoiceExtensions: choiceExtensions,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateEventTypeTriggerChoiceExtensions() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateMeasurementThresholdL1LoggedMdtThresholdRsrp(thresholdRsrp *xnapiesv1.ThresholdRSrp) (*xnapiesv1.MeasurementThresholdL1LoggedMdt, error) {

	item := &xnapiesv1.MeasurementThresholdL1LoggedMdt{
		MeasurementThresholdL1LoggedMdt: &xnapiesv1.MeasurementThresholdL1LoggedMdt_ThresholdRsrp{
			ThresholdRsrp: thresholdRsrp,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateMeasurementThresholdL1LoggedMdtThresholdRsrp() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateMeasurementThresholdL1LoggedMdtThresholdRsrq(thresholdRsrq *xnapiesv1.ThresholdRSrq) (*xnapiesv1.MeasurementThresholdL1LoggedMdt, error) {

	item := &xnapiesv1.MeasurementThresholdL1LoggedMdt{
		MeasurementThresholdL1LoggedMdt: &xnapiesv1.MeasurementThresholdL1LoggedMdt_ThresholdRsrq{
			ThresholdRsrq: thresholdRsrq,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateMeasurementThresholdL1LoggedMdtThresholdRsrq() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateGnbIDChoiceGnbID(gnbID *asn1.BitString) (*xnapiesv1.GnbIDChoice, error) {

	item := &xnapiesv1.GnbIDChoice{
		GnbIdChoice: &xnapiesv1.GnbIDChoice_GnbId{
			GnbId: gnbID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateGnbIDChoiceGnbID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateGnbIDChoiceChoiceExtension(choiceExtension *xnapiesv1.GnbIDChoiceExtIes) (*xnapiesv1.GnbIDChoice, error) {

	item := &xnapiesv1.GnbIDChoice{
		GnbIdChoice: &xnapiesv1.GnbIDChoice_ChoiceExtension{
			ChoiceExtension: choiceExtension,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateGnbIDChoiceChoiceExtension() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateEnbIDChoiceEnbIDMacro(enbIDMacro *asn1.BitString) (*xnapiesv1.EnbIDChoice, error) {

	item := &xnapiesv1.EnbIDChoice{
		EnbIdChoice: &xnapiesv1.EnbIDChoice_EnbIdMacro{
			EnbIdMacro: enbIDMacro,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateEnbIDChoiceEnbIDMacro() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateEnbIDChoiceEnbIDShortmacro(enbIDShortmacro *asn1.BitString) (*xnapiesv1.EnbIDChoice, error) {

	item := &xnapiesv1.EnbIDChoice{
		EnbIdChoice: &xnapiesv1.EnbIDChoice_EnbIdShortmacro{
			EnbIdShortmacro: enbIDShortmacro,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateEnbIDChoiceEnbIDShortmacro() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateEnbIDChoiceEnbIDLongmacro(enbIDLongmacro *asn1.BitString) (*xnapiesv1.EnbIDChoice, error) {

	item := &xnapiesv1.EnbIDChoice{
		EnbIdChoice: &xnapiesv1.EnbIDChoice_EnbIdLongmacro{
			EnbIdLongmacro: enbIDLongmacro,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateEnbIDChoiceEnbIDLongmacro() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateEnbIDChoiceChoiceExtension(choiceExtension *xnapiesv1.EnbIDChoiceExtIes) (*xnapiesv1.EnbIDChoice, error) {

	item := &xnapiesv1.EnbIDChoice{
		EnbIdChoice: &xnapiesv1.EnbIDChoice_ChoiceExtension{
			ChoiceExtension: choiceExtension,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateEnbIDChoiceChoiceExtension() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateGlobalNgRAnnodeIDGNb(gNb *xnapiesv1.GlobalgNbID) (*xnapiesv1.GlobalNgRAnnodeID, error) {

	item := &xnapiesv1.GlobalNgRAnnodeID{
		GlobalNgRannodeId: &xnapiesv1.GlobalNgRAnnodeID_GNb{
			GNb: gNb,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateGlobalNgRAnnodeIDGNb() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateGlobalNgRAnnodeIDNgENb(ngENb *xnapiesv1.GlobalngeNbID) (*xnapiesv1.GlobalNgRAnnodeID, error) {

	item := &xnapiesv1.GlobalNgRAnnodeID{
		GlobalNgRannodeId: &xnapiesv1.GlobalNgRAnnodeID_NgENb{
			NgENb: ngENb,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateGlobalNgRAnnodeIDNgENb() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateGlobalNgRAnnodeIDChoiceExtension(choiceExtension *xnapiesv1.GlobalNgRAnnodeIDExtIes) (*xnapiesv1.GlobalNgRAnnodeID, error) {

	item := &xnapiesv1.GlobalNgRAnnodeID{
		GlobalNgRannodeId: &xnapiesv1.GlobalNgRAnnodeID_ChoiceExtension{
			ChoiceExtension: choiceExtension,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateGlobalNgRAnnodeIDChoiceExtension() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateInitiatingConditionFailureIndicationRRcreestab(rRcreestab *xnapiesv1.Rrcreestabinitiated) (*xnapiesv1.InitiatingConditionFailureIndication, error) {

	item := &xnapiesv1.InitiatingConditionFailureIndication{
		InitiatingConditionFailureIndication: &xnapiesv1.InitiatingConditionFailureIndication_RRcreestab{
			RRcreestab: rRcreestab,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateInitiatingConditionFailureIndicationRRcreestab() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateInitiatingConditionFailureIndicationRRcsetup(rRcsetup *xnapiesv1.Rrcsetupinitiated) (*xnapiesv1.InitiatingConditionFailureIndication, error) {

	item := &xnapiesv1.InitiatingConditionFailureIndication{
		InitiatingConditionFailureIndication: &xnapiesv1.InitiatingConditionFailureIndication_RRcsetup{
			RRcsetup: rRcsetup,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateInitiatingConditionFailureIndicationRRcsetup() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateInitiatingConditionFailureIndicationChoiceExtension(choiceExtension *xnapiesv1.InitiatingConditionFailureIndicationExtIes) (*xnapiesv1.InitiatingConditionFailureIndication, error) {

	item := &xnapiesv1.InitiatingConditionFailureIndication{
		InitiatingConditionFailureIndication: &xnapiesv1.InitiatingConditionFailureIndication_ChoiceExtension{
			ChoiceExtension: choiceExtension,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateInitiatingConditionFailureIndicationChoiceExtension() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateIRNtiIRntiFull(iRntiFull *asn1.BitString) (*xnapiesv1.IRNti, error) {

	item := &xnapiesv1.IRNti{
		IRnti: &xnapiesv1.IRNti_IRntiFull{
			IRntiFull: iRntiFull,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIRNtiIRntiFull() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateIRNtiIRntiShort(iRntiShort *asn1.BitString) (*xnapiesv1.IRNti, error) {

	item := &xnapiesv1.IRNti{
		IRnti: &xnapiesv1.IRNti_IRntiShort{
			IRntiShort: iRntiShort,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIRNtiIRntiShort() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateIRNtiChoiceExtension(choiceExtension *xnapiesv1.IRNtiExtIes) (*xnapiesv1.IRNti, error) {

	item := &xnapiesv1.IRNti{
		IRnti: &xnapiesv1.IRNti_ChoiceExtension{
			ChoiceExtension: choiceExtension,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateIRNtiChoiceExtension() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateLastVisitedCellItemNGRanCell(nGRanCell *xnapiesv1.LastVisitedNgrancellInformation) (*xnapiesv1.LastVisitedCellItem, error) {

	item := &xnapiesv1.LastVisitedCellItem{
		LastVisitedCellItem: &xnapiesv1.LastVisitedCellItem_NGRanCell{
			NGRanCell: nGRanCell,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateLastVisitedCellItemNGRanCell() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateLastVisitedCellItemEUtranCell(eUtranCell *xnapiesv1.LastVisitedEutrancellInformation) (*xnapiesv1.LastVisitedCellItem, error) {

	item := &xnapiesv1.LastVisitedCellItem{
		LastVisitedCellItem: &xnapiesv1.LastVisitedCellItem_EUtranCell{
			EUtranCell: eUtranCell,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateLastVisitedCellItemEUtranCell() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateLastVisitedCellItemUTranCell(uTranCell *xnapiesv1.LastVisitedUtrancellInformation) (*xnapiesv1.LastVisitedCellItem, error) {

	item := &xnapiesv1.LastVisitedCellItem{
		LastVisitedCellItem: &xnapiesv1.LastVisitedCellItem_UTranCell{
			UTranCell: uTranCell,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateLastVisitedCellItemUTranCell() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateLastVisitedCellItemGEranCell(gEranCell *xnapiesv1.LastVisitedGerancellInformation) (*xnapiesv1.LastVisitedCellItem, error) {

	item := &xnapiesv1.LastVisitedCellItem{
		LastVisitedCellItem: &xnapiesv1.LastVisitedCellItem_GEranCell{
			GEranCell: gEranCell,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateLastVisitedCellItemGEranCell() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateLastVisitedCellItemChoiceExtension(choiceExtension *xnapiesv1.LastVisitedCellItemExtIes) (*xnapiesv1.LastVisitedCellItem, error) {

	item := &xnapiesv1.LastVisitedCellItem{
		LastVisitedCellItem: &xnapiesv1.LastVisitedCellItem_ChoiceExtension{
			ChoiceExtension: choiceExtension,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateLastVisitedCellItemChoiceExtension() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateMbsfnsubframeAllocationEUTraOneframe(oneframe *asn1.BitString) (*xnapiesv1.MbsfnsubframeAllocationEUTra, error) {

	item := &xnapiesv1.MbsfnsubframeAllocationEUTra{
		MbsfnsubframeAllocationEUtra: &xnapiesv1.MbsfnsubframeAllocationEUTra_Oneframe{
			Oneframe: oneframe,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateMbsfnsubframeAllocationEUTraOneframe() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateMbsfnsubframeAllocationEUTraFourframes(fourframes *asn1.BitString) (*xnapiesv1.MbsfnsubframeAllocationEUTra, error) {

	item := &xnapiesv1.MbsfnsubframeAllocationEUTra{
		MbsfnsubframeAllocationEUtra: &xnapiesv1.MbsfnsubframeAllocationEUTra_Fourframes{
			Fourframes: fourframes,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateMbsfnsubframeAllocationEUTraFourframes() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateMbsfnsubframeAllocationEUTraChoiceExtension(choiceExtension *xnapiesv1.MbsfnsubframeAllocationEUTraExtIes) (*xnapiesv1.MbsfnsubframeAllocationEUTra, error) {

	item := &xnapiesv1.MbsfnsubframeAllocationEUTra{
		MbsfnsubframeAllocationEUtra: &xnapiesv1.MbsfnsubframeAllocationEUTra_ChoiceExtension{
			ChoiceExtension: choiceExtension,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateMbsfnsubframeAllocationEUTraChoiceExtension() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateMdtmodeNRImmediateMdt(immediateMdt *xnapiesv1.ImmediateMdtNR) (*xnapiesv1.MdtmodeNR, error) {

	item := &xnapiesv1.MdtmodeNR{
		MdtmodeNr: &xnapiesv1.MdtmodeNR_ImmediateMdt{
			ImmediateMdt: immediateMdt,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateMdtmodeNRImmediateMdt() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateMdtmodeNRLoggedMdt(loggedMdt *xnapiesv1.LoggedMdtNR) (*xnapiesv1.MdtmodeNR, error) {

	item := &xnapiesv1.MdtmodeNR{
		MdtmodeNr: &xnapiesv1.MdtmodeNR_LoggedMdt{
			LoggedMdt: loggedMdt,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateMdtmodeNRLoggedMdt() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateMdtmodeNRMDtmodeNrExtension(mDtmodeNrExtension *xnapiesv1.MdtmodeNRExtension) (*xnapiesv1.MdtmodeNR, error) {

	item := &xnapiesv1.MdtmodeNR{
		MdtmodeNr: &xnapiesv1.MdtmodeNR_MDtmodeNrExtension{
			MDtmodeNrExtension: mDtmodeNrExtension,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateMdtmodeNRMDtmodeNrExtension() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateMdtmodeEUtraImmediateMdt(immediateMdt *xnapiesv1.ImmediateMdtEUtra) (*xnapiesv1.MdtmodeEUtra, error) {

	item := &xnapiesv1.MdtmodeEUtra{
		MdtmodeEutra: &xnapiesv1.MdtmodeEUtra_ImmediateMdt{
			ImmediateMdt: immediateMdt,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateMdtmodeEUtraImmediateMdt() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateMdtmodeEUtraLoggedMdt(loggedMdt *xnapiesv1.LoggedMdtEUtra) (*xnapiesv1.MdtmodeEUtra, error) {

	item := &xnapiesv1.MdtmodeEUtra{
		MdtmodeEutra: &xnapiesv1.MdtmodeEUtra_LoggedMdt{
			LoggedMdt: loggedMdt,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateMdtmodeEUtraLoggedMdt() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateMdtmodeEUtraMDtmodeEutraExtension(mDtmodeEutraExtension *xnapiesv1.MdtmodeEUtraExtension) (*xnapiesv1.MdtmodeEUtra, error) {

	item := &xnapiesv1.MdtmodeEUtra{
		MdtmodeEutra: &xnapiesv1.MdtmodeEUtra_MDtmodeEutraExtension{
			MDtmodeEutraExtension: mDtmodeEutraExtension,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateMdtmodeEUtraMDtmodeEutraExtension() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateMeasurementThresholdA2ThresholdRsrp(thresholdRsrp *xnapiesv1.ThresholdRSrp) (*xnapiesv1.MeasurementThresholdA2, error) {

	item := &xnapiesv1.MeasurementThresholdA2{
		MeasurementThresholdA2: &xnapiesv1.MeasurementThresholdA2_ThresholdRsrp{
			ThresholdRsrp: thresholdRsrp,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateMeasurementThresholdA2ThresholdRsrp() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateMeasurementThresholdA2ThresholdRsrq(thresholdRsrq *xnapiesv1.ThresholdRSrq) (*xnapiesv1.MeasurementThresholdA2, error) {

	item := &xnapiesv1.MeasurementThresholdA2{
		MeasurementThresholdA2: &xnapiesv1.MeasurementThresholdA2_ThresholdRsrq{
			ThresholdRsrq: thresholdRsrq,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateMeasurementThresholdA2ThresholdRsrq() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateMeasurementThresholdA2ThresholdSinr(thresholdSinr *xnapiesv1.ThresholdSInr) (*xnapiesv1.MeasurementThresholdA2, error) {

	item := &xnapiesv1.MeasurementThresholdA2{
		MeasurementThresholdA2: &xnapiesv1.MeasurementThresholdA2_ThresholdSinr{
			ThresholdSinr: thresholdSinr,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateMeasurementThresholdA2ThresholdSinr() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateMeasurementThresholdA2ChoiceExtension(choiceExtension *xnapiesv1.MeasurementThresholdA2ExtIes) (*xnapiesv1.MeasurementThresholdA2, error) {

	item := &xnapiesv1.MeasurementThresholdA2{
		MeasurementThresholdA2: &xnapiesv1.MeasurementThresholdA2_ChoiceExtension{
			ChoiceExtension: choiceExtension,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateMeasurementThresholdA2ChoiceExtension() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateMobilityRestrictionListExtIesExtensionIDLastEUtranplmnidentity(idLastEUtranplmnidentity *xnapiesv1.PlmnIdentity) (*xnapiesv1.MobilityRestrictionListExtIesExtension, error) {

	item := &xnapiesv1.MobilityRestrictionListExtIesExtension{
		MobilityRestrictionListExtIes: &xnapiesv1.MobilityRestrictionListExtIesExtension_IdLastEUtranplmnidentity{
			IdLastEUtranplmnidentity: idLastEUtranplmnidentity,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateMobilityRestrictionListExtIesExtensionIDLastEUtranplmnidentity() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateMobilityRestrictionListExtIesExtensionIDCntypeRestrictionsForServing(IDCntypeRestrictionsForServing xnapiesv1.CntypeRestrictionsForServing) (*xnapiesv1.MobilityRestrictionListExtIesExtension, error) {

	item := &xnapiesv1.MobilityRestrictionListExtIesExtension{
		MobilityRestrictionListExtIes: &xnapiesv1.MobilityRestrictionListExtIesExtension_IdCntypeRestrictionsForServing{
			IdCntypeRestrictionsForServing: IDCntypeRestrictionsForServing,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateMobilityRestrictionListExtIesExtensionIDCntypeRestrictionsForServing() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateMobilityRestrictionListExtIesExtensionIDCntypeRestrictionsForEquivalent(IDCntypeRestrictionsForEquivalent *xnapiesv1.CntypeRestrictionsForEquivalent) (*xnapiesv1.MobilityRestrictionListExtIesExtension, error) {

	item := &xnapiesv1.MobilityRestrictionListExtIesExtension{
		MobilityRestrictionListExtIes: &xnapiesv1.MobilityRestrictionListExtIesExtension_IdCntypeRestrictionsForEquivalent{
			IdCntypeRestrictionsForEquivalent: IDCntypeRestrictionsForEquivalent,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateMobilityRestrictionListExtIesExtensionIDCntypeRestrictionsForEquivalent() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateMobilityRestrictionListExtIesExtensionIDNpnmobilityInformation(IDNpnmobilityInformation *xnapiesv1.NpnmobilityInformation) (*xnapiesv1.MobilityRestrictionListExtIesExtension, error) {

	item := &xnapiesv1.MobilityRestrictionListExtIesExtension{
		MobilityRestrictionListExtIes: &xnapiesv1.MobilityRestrictionListExtIesExtension_IdNpnmobilityInformation{
			IdNpnmobilityInformation: IDNpnmobilityInformation,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateMobilityRestrictionListExtIesExtensionIDNpnmobilityInformation() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateNgRAnNodeResourceCoordinationInfoEutraResourceCoordinationInfo(eutraResourceCoordinationInfo *xnapiesv1.EUTraResourceCoordinationInfo) (*xnapiesv1.NgRAnNodeResourceCoordinationInfo, error) {

	item := &xnapiesv1.NgRAnNodeResourceCoordinationInfo{
		NgRanNodeResourceCoordinationInfo: &xnapiesv1.NgRAnNodeResourceCoordinationInfo_EutraResourceCoordinationInfo{
			EutraResourceCoordinationInfo: eutraResourceCoordinationInfo,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNgRAnNodeResourceCoordinationInfoEutraResourceCoordinationInfo() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateNgRAnNodeResourceCoordinationInfoNrResourceCoordinationInfo(nrResourceCoordinationInfo *xnapiesv1.NrResourceCoordinationInfo) (*xnapiesv1.NgRAnNodeResourceCoordinationInfo, error) {

	item := &xnapiesv1.NgRAnNodeResourceCoordinationInfo{
		NgRanNodeResourceCoordinationInfo: &xnapiesv1.NgRAnNodeResourceCoordinationInfo_NrResourceCoordinationInfo{
			NrResourceCoordinationInfo: nrResourceCoordinationInfo,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNgRAnNodeResourceCoordinationInfoNrResourceCoordinationInfo() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateNeighbourInformationNRModeInfoFddInfo(fddInfo *xnapiesv1.NeighbourInformationNRModeFddinfo) (*xnapiesv1.NeighbourInformationNRModeInfo, error) {

	item := &xnapiesv1.NeighbourInformationNRModeInfo{
		NeighbourInformationNrModeInfo: &xnapiesv1.NeighbourInformationNRModeInfo_FddInfo{
			FddInfo: fddInfo,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNeighbourInformationNRModeInfoFddInfo() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateNeighbourInformationNRModeInfoTddInfo(tddInfo *xnapiesv1.NeighbourInformationNRModeTddinfo) (*xnapiesv1.NeighbourInformationNRModeInfo, error) {

	item := &xnapiesv1.NeighbourInformationNRModeInfo{
		NeighbourInformationNrModeInfo: &xnapiesv1.NeighbourInformationNRModeInfo_TddInfo{
			TddInfo: tddInfo,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNeighbourInformationNRModeInfoTddInfo() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateNeighbourInformationNRModeInfoChoiceExtension(choiceExtension *xnapiesv1.NeighbourInformationNRModeInfoExtIes) (*xnapiesv1.NeighbourInformationNRModeInfo, error) {

	item := &xnapiesv1.NeighbourInformationNRModeInfo{
		NeighbourInformationNrModeInfo: &xnapiesv1.NeighbourInformationNRModeInfo_ChoiceExtension{
			ChoiceExtension: choiceExtension,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNeighbourInformationNRModeInfoChoiceExtension() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateNgRAnCellIdentityNr(nr *xnapiesv1.NrCellIdentity) (*xnapiesv1.NgRAnCellIdentity, error) {

	item := &xnapiesv1.NgRAnCellIdentity{
		NgRanCellIdentity: &xnapiesv1.NgRAnCellIdentity_Nr{
			Nr: nr,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNgRAnCellIdentityNr() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateNgRAnCellIdentityEUtra(eUtra *xnapiesv1.EUTraCellIdentity) (*xnapiesv1.NgRAnCellIdentity, error) {

	item := &xnapiesv1.NgRAnCellIdentity{
		NgRanCellIdentity: &xnapiesv1.NgRAnCellIdentity_EUtra{
			EUtra: eUtra,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNgRAnCellIdentityEUtra() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateNgRAnCellIdentityChoiceExtension(choiceExtension *xnapiesv1.NgRAnCellIdentityExtIes) (*xnapiesv1.NgRAnCellIdentity, error) {

	item := &xnapiesv1.NgRAnCellIdentity{
		NgRanCellIdentity: &xnapiesv1.NgRAnCellIdentity_ChoiceExtension{
			ChoiceExtension: choiceExtension,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNgRAnCellIdentityChoiceExtension() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateNgRAnCellPciNr(nr *xnapiesv1.Nrpci) (*xnapiesv1.NgRAnCellPci, error) {

	item := &xnapiesv1.NgRAnCellPci{
		NgRanCellPci: &xnapiesv1.NgRAnCellPci_Nr{
			Nr: nr,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNgRAnCellPciNr() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateNgRAnCellPciEUtra(eUtra *xnapiesv1.EUTrapci) (*xnapiesv1.NgRAnCellPci, error) {

	item := &xnapiesv1.NgRAnCellPci{
		NgRanCellPci: &xnapiesv1.NgRAnCellPci_EUtra{
			EUtra: eUtra,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNgRAnCellPciEUtra() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateNgRAnCellPciChoiceExtension(choiceExtension *xnapiesv1.NgRAnCellPciExtIes) (*xnapiesv1.NgRAnCellPci, error) {

	item := &xnapiesv1.NgRAnCellPci{
		NgRanCellPci: &xnapiesv1.NgRAnCellPci_ChoiceExtension{
			ChoiceExtension: choiceExtension,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNgRAnCellPciChoiceExtension() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateNonDynamic5QIDescriptorExtIesExtensionIDCnpacketDelayBudgetDownlink(IDCnpacketDelayBudgetDownlink *xnapiesv1.ExtendedPacketDelayBudget) (*xnapiesv1.NonDynamic5QidescriptorExtIesExtension, error) {

	item := &xnapiesv1.NonDynamic5QidescriptorExtIesExtension{
		NonDynamic5QidescriptorExtIes: &xnapiesv1.NonDynamic5QidescriptorExtIesExtension_IdCnpacketDelayBudgetDownlink{
			IdCnpacketDelayBudgetDownlink: IDCnpacketDelayBudgetDownlink,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNonDynamic5QIDescriptorExtIesExtensionIDCnpacketDelayBudgetDownlink() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateNonDynamic5QIDescriptorExtIesExtensionIDCnpacketDelayBudgetUplink(IDCnpacketDelayBudgetUplink *xnapiesv1.ExtendedPacketDelayBudget) (*xnapiesv1.NonDynamic5QidescriptorExtIesExtension, error) {

	item := &xnapiesv1.NonDynamic5QidescriptorExtIesExtension{
		NonDynamic5QidescriptorExtIes: &xnapiesv1.NonDynamic5QidescriptorExtIesExtension_IdCnpacketDelayBudgetUplink{
			IdCnpacketDelayBudgetUplink: IDCnpacketDelayBudgetUplink,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNonDynamic5QIDescriptorExtIesExtensionIDCnpacketDelayBudgetUplink() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateNgeNbRadioResourceStatusExtIesExtensionIDDlSchedulingPdcchCceUsage(IDDlSchedulingPdcchCceUsage *xnapiesv1.DlschedulingPDcchCCeusage) (*xnapiesv1.NgeNbRadioResourceStatusExtIesExtension, error) {

	item := &xnapiesv1.NgeNbRadioResourceStatusExtIesExtension{
		NgENbRadioResourceStatusExtIes: &xnapiesv1.NgeNbRadioResourceStatusExtIesExtension_IdDlSchedulingPdcchCceUsage{
			IdDlSchedulingPdcchCceUsage: IDDlSchedulingPdcchCceUsage,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNgeNbRadioResourceStatusExtIesExtensionIDDlSchedulingPdcchCceUsage() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateNgeNbRadioResourceStatusExtIesExtensionIDUlSchedulingPdcchCceUsage(IDUlSchedulingPdcchCceUsage *xnapiesv1.UlschedulingPDcchCCeusage) (*xnapiesv1.NgeNbRadioResourceStatusExtIesExtension, error) {

	item := &xnapiesv1.NgeNbRadioResourceStatusExtIesExtension{
		NgENbRadioResourceStatusExtIes: &xnapiesv1.NgeNbRadioResourceStatusExtIesExtension_IdUlSchedulingPdcchCceUsage{
			IdUlSchedulingPdcchCceUsage: IDUlSchedulingPdcchCceUsage,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNgeNbRadioResourceStatusExtIesExtensionIDUlSchedulingPdcchCceUsage() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateNpnBroadcastInformationSnpnInformation(snpnInformation *xnapiesv1.NpnBroadcastInformationSNpn) (*xnapiesv1.NpnBroadcastInformation, error) {

	item := &xnapiesv1.NpnBroadcastInformation{
		NpnBroadcastInformation: &xnapiesv1.NpnBroadcastInformation_SnpnInformation{
			SnpnInformation: snpnInformation,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNpnBroadcastInformationSnpnInformation() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateNpnBroadcastInformationPniNpnInformation(pniNpnInformation *xnapiesv1.NpnBroadcastInformationPNiNPn) (*xnapiesv1.NpnBroadcastInformation, error) {

	item := &xnapiesv1.NpnBroadcastInformation{
		NpnBroadcastInformation: &xnapiesv1.NpnBroadcastInformation_PniNpnInformation{
			PniNpnInformation: pniNpnInformation,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNpnBroadcastInformationPniNpnInformation() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateNpnBroadcastInformationChoiceExtension(choiceExtension *xnapiesv1.NpnBroadcastInformationExtIes) (*xnapiesv1.NpnBroadcastInformation, error) {

	item := &xnapiesv1.NpnBroadcastInformation{
		NpnBroadcastInformation: &xnapiesv1.NpnBroadcastInformation_ChoiceExtension{
			ChoiceExtension: choiceExtension,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNpnBroadcastInformationChoiceExtension() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateNpnmobilityInformationSnpnMobilityInformation(snpnMobilityInformation *xnapiesv1.NpnmobilityInformationSNpn) (*xnapiesv1.NpnmobilityInformation, error) {

	item := &xnapiesv1.NpnmobilityInformation{
		NpnmobilityInformation: &xnapiesv1.NpnmobilityInformation_SnpnMobilityInformation{
			SnpnMobilityInformation: snpnMobilityInformation,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNpnmobilityInformationSnpnMobilityInformation() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateNpnmobilityInformationPniNpnMobilityInformation(pniNpnMobilityInformation *xnapiesv1.NpnmobilityInformationPNiNPn) (*xnapiesv1.NpnmobilityInformation, error) {

	item := &xnapiesv1.NpnmobilityInformation{
		NpnmobilityInformation: &xnapiesv1.NpnmobilityInformation_PniNpnMobilityInformation{
			PniNpnMobilityInformation: pniNpnMobilityInformation,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNpnmobilityInformationPniNpnMobilityInformation() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateNpnmobilityInformationChoiceExtension(choiceExtension *xnapiesv1.NpnmobilityInformationExtIes) (*xnapiesv1.NpnmobilityInformation, error) {

	item := &xnapiesv1.NpnmobilityInformation{
		NpnmobilityInformation: &xnapiesv1.NpnmobilityInformation_ChoiceExtension{
			ChoiceExtension: choiceExtension,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNpnmobilityInformationChoiceExtension() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateNpnpagingAssistanceInformationPniNpnInformation(pniNpnInformation *xnapiesv1.NpnpagingAssistanceInformationPNiNPn) (*xnapiesv1.NpnpagingAssistanceInformation, error) {

	item := &xnapiesv1.NpnpagingAssistanceInformation{
		NpnpagingAssistanceInformation: &xnapiesv1.NpnpagingAssistanceInformation_PniNpnInformation{
			PniNpnInformation: pniNpnInformation,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNpnpagingAssistanceInformationPniNpnInformation() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateNpnpagingAssistanceInformationChoiceExtension(choiceExtension *xnapiesv1.NpnpagingAssistanceInformationExtIes) (*xnapiesv1.NpnpagingAssistanceInformation, error) {

	item := &xnapiesv1.NpnpagingAssistanceInformation{
		NpnpagingAssistanceInformation: &xnapiesv1.NpnpagingAssistanceInformation_ChoiceExtension{
			ChoiceExtension: choiceExtension,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNpnpagingAssistanceInformationChoiceExtension() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateNpnSupportSNpnChoice(sNpn *xnapiesv1.NpnSupportSNpn) (*xnapiesv1.NpnSupport, error) {

	item := &xnapiesv1.NpnSupport{
		NpnSupport: &xnapiesv1.NpnSupport_SNpn{
			SNpn: sNpn,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNpnSupportSNpn() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateNpnSupportChoiceExtensions(choiceExtensions *xnapiesv1.NpnSupportExtIes) (*xnapiesv1.NpnSupport, error) {

	item := &xnapiesv1.NpnSupport{
		NpnSupport: &xnapiesv1.NpnSupport_ChoiceExtensions{
			ChoiceExtensions: choiceExtensions,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNpnSupportChoiceExtensions() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateFddortddNprachconfigurationFdd(fdd *xnapiesv1.NprachconfigurationFDd) (*xnapiesv1.FddortddNprachconfiguration, error) {

	item := &xnapiesv1.FddortddNprachconfiguration{
		FddOrTddNprachconfiguration: &xnapiesv1.FddortddNprachconfiguration_Fdd{
			Fdd: fdd,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateFddortddNprachconfigurationFdd() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateFddortddNprachconfigurationTdd(tdd *xnapiesv1.NprachconfigurationTDd) (*xnapiesv1.FddortddNprachconfiguration, error) {

	item := &xnapiesv1.FddortddNprachconfiguration{
		FddOrTddNprachconfiguration: &xnapiesv1.FddortddNprachconfiguration_Tdd{
			Tdd: tdd,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateFddortddNprachconfigurationTdd() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateFddortddNprachconfigurationChoiceExtension(choiceExtension *xnapiesv1.FddorTDdinNPrachconfigurationChoiceExtIes) (*xnapiesv1.FddortddNprachconfiguration, error) {

	item := &xnapiesv1.FddortddNprachconfiguration{
		FddOrTddNprachconfiguration: &xnapiesv1.FddortddNprachconfiguration_ChoiceExtension{
			ChoiceExtension: choiceExtension,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateFddortddNprachconfigurationChoiceExtension() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateNrmodeInfoFddChoice(fdd *xnapiesv1.NrmodeInfoFdd) (*xnapiesv1.NrmodeInfo, error) {

	item := &xnapiesv1.NrmodeInfo{
		NrmodeInfo: &xnapiesv1.NrmodeInfo_Fdd{
			Fdd: fdd,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNrmodeInfoFdd() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateNrmodeInfoTddChoice(tdd *xnapiesv1.NrmodeInfoTdd) (*xnapiesv1.NrmodeInfo, error) {

	item := &xnapiesv1.NrmodeInfo{
		NrmodeInfo: &xnapiesv1.NrmodeInfo_Tdd{
			Tdd: tdd,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNrmodeInfoTdd() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateNrmodeInfoChoiceExtension(choiceExtension *xnapiesv1.NrmodeInfoExtIes) (*xnapiesv1.NrmodeInfo, error) {

	item := &xnapiesv1.NrmodeInfo{
		NrmodeInfo: &xnapiesv1.NrmodeInfo_ChoiceExtension{
			ChoiceExtension: choiceExtension,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNrmodeInfoChoiceExtension() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateNrmodeInfoFddExtIesExtensionIDUlcarrierList(IDUlcarrierList *xnapiesv1.NrcarrierList) (*xnapiesv1.NrmodeInfoFddExtIesExtension, error) {

	item := &xnapiesv1.NrmodeInfoFddExtIesExtension{
		NrmodeInfoFddExtIes: &xnapiesv1.NrmodeInfoFddExtIesExtension_IdUlcarrierList{
			IdUlcarrierList: IDUlcarrierList,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNrmodeInfoFddExtIesExtensionIDUlcarrierList() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateNrmodeInfoFddExtIesExtensionIDDlcarrierList(IDDlcarrierList *xnapiesv1.NrcarrierList) (*xnapiesv1.NrmodeInfoFddExtIesExtension, error) {

	item := &xnapiesv1.NrmodeInfoFddExtIesExtension{
		NrmodeInfoFddExtIes: &xnapiesv1.NrmodeInfoFddExtIesExtension_IdDlcarrierList{
			IdDlcarrierList: IDDlcarrierList,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNrmodeInfoFddExtIesExtensionIDDlcarrierList() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateNrmodeInfoTddExtIesExtensionIDIntendedTddDlUlconfigurationNr(IDIntendedTddDlUlconfigurationNr *xnapiesv1.IntendedTddDLULconfigurationNR) (*xnapiesv1.NrmodeInfoTddExtIesExtension, error) {

	item := &xnapiesv1.NrmodeInfoTddExtIesExtension{
		NrmodeInfoTddExtIes: &xnapiesv1.NrmodeInfoTddExtIesExtension_IdIntendedTddDlUlconfigurationNr{
			IdIntendedTddDlUlconfigurationNr: IDIntendedTddDlUlconfigurationNr,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNrmodeInfoTddExtIesExtensionIDIntendedTddDlUlconfigurationNr() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateNrmodeInfoTddExtIesExtensionIDTdduldlconfigurationCommonNr(IDTdduldlconfigurationCommonNr *xnapiesv1.TdduldlconfigurationCommonNr) (*xnapiesv1.NrmodeInfoTddExtIesExtension, error) {

	item := &xnapiesv1.NrmodeInfoTddExtIesExtension{
		NrmodeInfoTddExtIes: &xnapiesv1.NrmodeInfoTddExtIesExtension_IdTdduldlconfigurationCommonNr{
			IdTdduldlconfigurationCommonNr: IDTdduldlconfigurationCommonNr,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNrmodeInfoTddExtIesExtensionIDTdduldlconfigurationCommonNr() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateNrmodeInfoTddExtIesExtensionIDCarrierList(IDCarrierList *xnapiesv1.NrcarrierList) (*xnapiesv1.NrmodeInfoTddExtIesExtension, error) {

	item := &xnapiesv1.NrmodeInfoTddExtIesExtension{
		NrmodeInfoTddExtIes: &xnapiesv1.NrmodeInfoTddExtIesExtension_IdCarrierList{
			IdCarrierList: IDCarrierList,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNrmodeInfoTddExtIesExtensionIDCarrierList() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreatePdcpchangeIndicationFromSNgRanNode(fromSNgRanNode xnapiesv1.FromSngrannodePdcpchangeIndication) (*xnapiesv1.PdcpchangeIndication, error) {

	item := &xnapiesv1.PdcpchangeIndication{
		PdcpchangeIndication: &xnapiesv1.PdcpchangeIndication_FromSNgRanNode{
			FromSNgRanNode: fromSNgRanNode,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePdcpchangeIndicationFromSNgRanNode() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreatePdcpchangeIndicationFromMNgRanNode(fromMNgRanNode xnapiesv1.FromMngrannodePdcpchangeIndication) (*xnapiesv1.PdcpchangeIndication, error) {

	item := &xnapiesv1.PdcpchangeIndication{
		PdcpchangeIndication: &xnapiesv1.PdcpchangeIndication_FromMNgRanNode{
			FromMNgRanNode: fromMNgRanNode,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePdcpchangeIndicationFromMNgRanNode() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreatePdcpchangeIndicationChoiceExtension(choiceExtension *xnapiesv1.PdcpchangeIndicationExtIes) (*xnapiesv1.PdcpchangeIndication, error) {

	item := &xnapiesv1.PdcpchangeIndication{
		PdcpchangeIndication: &xnapiesv1.PdcpchangeIndication_ChoiceExtension{
			ChoiceExtension: choiceExtension,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePdcpchangeIndicationChoiceExtension() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreatePdusessionResourcesToBeSetupItemExtIesExtensionIDAdditionalUlNgUTnlatUpfList(IDAdditionalUlNgUTnlatUpfList *xnapiesv1.AdditionalULNGUTNlatUpfList) (*xnapiesv1.PdusessionResourcesToBeSetupItemExtIesExtension, error) {

	item := &xnapiesv1.PdusessionResourcesToBeSetupItemExtIesExtension{
		PdusessionResourcesToBeSetupItemExtIes: &xnapiesv1.PdusessionResourcesToBeSetupItemExtIesExtension_IdAdditionalUlNgUTnlatUpfList{
			IdAdditionalUlNgUTnlatUpfList: IDAdditionalUlNgUTnlatUpfList,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePdusessionResourcesToBeSetupItemExtIesExtensionIDAdditionalUlNgUTnlatUpfList() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreatePdusessionResourcesToBeSetupItemExtIesExtensionIDPdusessionCommonNetworkInstance(IDPdusessionCommonNetworkInstance *xnapiesv1.PdusessionCommonNetworkInstance) (*xnapiesv1.PdusessionResourcesToBeSetupItemExtIesExtension, error) {

	item := &xnapiesv1.PdusessionResourcesToBeSetupItemExtIesExtension{
		PdusessionResourcesToBeSetupItemExtIes: &xnapiesv1.PdusessionResourcesToBeSetupItemExtIesExtension_IdPdusessionCommonNetworkInstance{
			IdPdusessionCommonNetworkInstance: IDPdusessionCommonNetworkInstance,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePdusessionResourcesToBeSetupItemExtIesExtensionIDPdusessionCommonNetworkInstance() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreatePdusessionResourcesToBeSetupItemExtIesExtensionIDRedundantUlNgUTnlatUpf(IDRedundantUlNgUTnlatUpf *xnapiesv1.UptransportLayerInformation) (*xnapiesv1.PdusessionResourcesToBeSetupItemExtIesExtension, error) {

	item := &xnapiesv1.PdusessionResourcesToBeSetupItemExtIesExtension{
		PdusessionResourcesToBeSetupItemExtIes: &xnapiesv1.PdusessionResourcesToBeSetupItemExtIesExtension_IdRedundantUlNgUTnlatUpf{
			IdRedundantUlNgUTnlatUpf: IDRedundantUlNgUTnlatUpf,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePdusessionResourcesToBeSetupItemExtIesExtensionIDRedundantUlNgUTnlatUpf() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreatePdusessionResourcesToBeSetupItemExtIesExtensionIDAdditionalRedundantUlNgUTnlatUpfList(IDAdditionalRedundantUlNgUTnlatUpfList *xnapiesv1.AdditionalULNGUTNlatUpfList) (*xnapiesv1.PdusessionResourcesToBeSetupItemExtIesExtension, error) {

	item := &xnapiesv1.PdusessionResourcesToBeSetupItemExtIesExtension{
		PdusessionResourcesToBeSetupItemExtIes: &xnapiesv1.PdusessionResourcesToBeSetupItemExtIesExtension_IdAdditionalRedundantUlNgUTnlatUpfList{
			IdAdditionalRedundantUlNgUTnlatUpfList: IDAdditionalRedundantUlNgUTnlatUpfList,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePdusessionResourcesToBeSetupItemExtIesExtensionIDAdditionalRedundantUlNgUTnlatUpfList() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreatePdusessionResourcesToBeSetupItemExtIesExtensionIDRedundantCommonNetworkInstance(IDRedundantCommonNetworkInstance *xnapiesv1.PdusessionCommonNetworkInstance) (*xnapiesv1.PdusessionResourcesToBeSetupItemExtIesExtension, error) {

	item := &xnapiesv1.PdusessionResourcesToBeSetupItemExtIesExtension{
		PdusessionResourcesToBeSetupItemExtIes: &xnapiesv1.PdusessionResourcesToBeSetupItemExtIesExtension_IdRedundantCommonNetworkInstance{
			IdRedundantCommonNetworkInstance: IDRedundantCommonNetworkInstance,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePdusessionResourcesToBeSetupItemExtIesExtensionIDRedundantCommonNetworkInstance() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreatePdusessionResourcesToBeSetupItemExtIesExtensionIDRedundantPdusessionInformation(IDRedundantPdusessionInformation *xnapiesv1.RedundantPdusessionInformation) (*xnapiesv1.PdusessionResourcesToBeSetupItemExtIesExtension, error) {

	item := &xnapiesv1.PdusessionResourcesToBeSetupItemExtIesExtension{
		PdusessionResourcesToBeSetupItemExtIes: &xnapiesv1.PdusessionResourcesToBeSetupItemExtIesExtension_IdRedundantPdusessionInformation{
			IdRedundantPdusessionInformation: IDRedundantPdusessionInformation,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePdusessionResourcesToBeSetupItemExtIesExtensionIDRedundantPdusessionInformation() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreatePdusessionResourceSetupInfoSNterminatedExtIesExtensionIDSecurityResult(IDSecurityResult *xnapiesv1.SecurityResult) (*xnapiesv1.PdusessionResourceSetupInfoSNterminatedExtIesExtension, error) {

	item := &xnapiesv1.PdusessionResourceSetupInfoSNterminatedExtIesExtension{
		PdusessionResourceSetupInfoSnterminatedExtIes: &xnapiesv1.PdusessionResourceSetupInfoSNterminatedExtIesExtension_IdSecurityResult{
			IdSecurityResult: IDSecurityResult,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePdusessionResourceSetupInfoSNterminatedExtIesExtensionIDSecurityResult() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreatePdusessionResourceSetupInfoSNterminatedExtIesExtensionIDPdusessionCommonNetworkInstance(IDPdusessionCommonNetworkInstance *xnapiesv1.PdusessionCommonNetworkInstance) (*xnapiesv1.PdusessionResourceSetupInfoSNterminatedExtIesExtension, error) {

	item := &xnapiesv1.PdusessionResourceSetupInfoSNterminatedExtIesExtension{
		PdusessionResourceSetupInfoSnterminatedExtIes: &xnapiesv1.PdusessionResourceSetupInfoSNterminatedExtIesExtension_IdPdusessionCommonNetworkInstance{
			IdPdusessionCommonNetworkInstance: IDPdusessionCommonNetworkInstance,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePdusessionResourceSetupInfoSNterminatedExtIesExtensionIDPdusessionCommonNetworkInstance() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreatePdusessionResourceSetupInfoSNterminatedExtIesExtensionIDDefaultDrbAllowed(IDDefaultDrbAllowed xnapiesv1.DefaultDrbAllowed) (*xnapiesv1.PdusessionResourceSetupInfoSNterminatedExtIesExtension, error) {

	item := &xnapiesv1.PdusessionResourceSetupInfoSNterminatedExtIesExtension{
		PdusessionResourceSetupInfoSnterminatedExtIes: &xnapiesv1.PdusessionResourceSetupInfoSNterminatedExtIesExtension_IdDefaultDrbAllowed{
			IdDefaultDrbAllowed: IDDefaultDrbAllowed,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePdusessionResourceSetupInfoSNterminatedExtIesExtensionIDDefaultDrbAllowed() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreatePdusessionResourceSetupInfoSNterminatedExtIesExtensionIDSplitSessionIndicator(IDSplitSessionIndicator xnapiesv1.SplitSessionIndicator) (*xnapiesv1.PdusessionResourceSetupInfoSNterminatedExtIesExtension, error) {

	item := &xnapiesv1.PdusessionResourceSetupInfoSNterminatedExtIesExtension{
		PdusessionResourceSetupInfoSnterminatedExtIes: &xnapiesv1.PdusessionResourceSetupInfoSNterminatedExtIesExtension_IdSplitSessionIndicator{
			IdSplitSessionIndicator: IDSplitSessionIndicator,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePdusessionResourceSetupInfoSNterminatedExtIesExtensionIDSplitSessionIndicator() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreatePdusessionResourceSetupInfoSNterminatedExtIesExtensionIDNonGbrresourcesOffered(IDNonGbrresourcesOffered xnapiesv1.NonGbrresourcesOffered) (*xnapiesv1.PdusessionResourceSetupInfoSNterminatedExtIesExtension, error) {

	item := &xnapiesv1.PdusessionResourceSetupInfoSNterminatedExtIesExtension{
		PdusessionResourceSetupInfoSnterminatedExtIes: &xnapiesv1.PdusessionResourceSetupInfoSNterminatedExtIesExtension_IdNonGbrresourcesOffered{
			IdNonGbrresourcesOffered: IDNonGbrresourcesOffered,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePdusessionResourceSetupInfoSNterminatedExtIesExtensionIDNonGbrresourcesOffered() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreatePdusessionResourceSetupInfoSNterminatedExtIesExtensionIDRedundantUlNgUTnlatUpf(IDRedundantUlNgUTnlatUpf *xnapiesv1.UptransportLayerInformation) (*xnapiesv1.PdusessionResourceSetupInfoSNterminatedExtIesExtension, error) {

	item := &xnapiesv1.PdusessionResourceSetupInfoSNterminatedExtIesExtension{
		PdusessionResourceSetupInfoSnterminatedExtIes: &xnapiesv1.PdusessionResourceSetupInfoSNterminatedExtIesExtension_IdRedundantUlNgUTnlatUpf{
			IdRedundantUlNgUTnlatUpf: IDRedundantUlNgUTnlatUpf,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePdusessionResourceSetupInfoSNterminatedExtIesExtensionIDRedundantUlNgUTnlatUpf() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreatePdusessionResourceSetupInfoSNterminatedExtIesExtensionIDRedundantCommonNetworkInstance(IDRedundantCommonNetworkInstance *xnapiesv1.PdusessionCommonNetworkInstance) (*xnapiesv1.PdusessionResourceSetupInfoSNterminatedExtIesExtension, error) {

	item := &xnapiesv1.PdusessionResourceSetupInfoSNterminatedExtIesExtension{
		PdusessionResourceSetupInfoSnterminatedExtIes: &xnapiesv1.PdusessionResourceSetupInfoSNterminatedExtIesExtension_IdRedundantCommonNetworkInstance{
			IdRedundantCommonNetworkInstance: IDRedundantCommonNetworkInstance,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePdusessionResourceSetupInfoSNterminatedExtIesExtensionIDRedundantCommonNetworkInstance() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreatePdusessionResourceSetupInfoSNterminatedExtIesExtensionIDRedundantPdusessionInformation(IDRedundantPdusessionInformation *xnapiesv1.RedundantPdusessionInformation) (*xnapiesv1.PdusessionResourceSetupInfoSNterminatedExtIesExtension, error) {

	item := &xnapiesv1.PdusessionResourceSetupInfoSNterminatedExtIesExtension{
		PdusessionResourceSetupInfoSnterminatedExtIes: &xnapiesv1.PdusessionResourceSetupInfoSNterminatedExtIesExtension_IdRedundantPdusessionInformation{
			IdRedundantPdusessionInformation: IDRedundantPdusessionInformation,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePdusessionResourceSetupInfoSNterminatedExtIesExtensionIDRedundantPdusessionInformation() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateQoSflowsToBeSetupListSetupSNterminatedItemExtIesExtensionIDTsctrafficCharacteristics(IDTsctrafficCharacteristics *xnapiesv1.TsctrafficCharacteristics) (*xnapiesv1.QoSflowsToBeSetupListSetupSNterminatedItemExtIesExtension, error) {

	item := &xnapiesv1.QoSflowsToBeSetupListSetupSNterminatedItemExtIesExtension{
		QoSflowsToBeSetupListSetupSnterminatedItemExtIes: &xnapiesv1.QoSflowsToBeSetupListSetupSNterminatedItemExtIesExtension_IdTsctrafficCharacteristics{
			IdTsctrafficCharacteristics: IDTsctrafficCharacteristics,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateQoSflowsToBeSetupListSetupSNterminatedItemExtIesExtensionIDTsctrafficCharacteristics() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateQoSflowsToBeSetupListSetupSNterminatedItemExtIesExtensionIDRedundantQoSflowIndicator(IDRedundantQoSflowIndicator xnapiesv1.RedundantQoSflowIndicator) (*xnapiesv1.QoSflowsToBeSetupListSetupSNterminatedItemExtIesExtension, error) {

	item := &xnapiesv1.QoSflowsToBeSetupListSetupSNterminatedItemExtIesExtension{
		QoSflowsToBeSetupListSetupSnterminatedItemExtIes: &xnapiesv1.QoSflowsToBeSetupListSetupSNterminatedItemExtIesExtension_IdRedundantQoSflowIndicator{
			IdRedundantQoSflowIndicator: IDRedundantQoSflowIndicator,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateQoSflowsToBeSetupListSetupSNterminatedItemExtIesExtensionIDRedundantQoSflowIndicator() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreatePdusessionResourceSetupResponseInfoSNterminatedExtIesExtensionIDDrbIDsTakenintouse(IDDrbIDsTakenintouse *xnapiesv1.DrbList) (*xnapiesv1.PdusessionResourceSetupResponseInfoSNterminatedExtIesExtension, error) {

	item := &xnapiesv1.PdusessionResourceSetupResponseInfoSNterminatedExtIesExtension{
		PdusessionResourceSetupResponseInfoSnterminatedExtIes: &xnapiesv1.PdusessionResourceSetupResponseInfoSNterminatedExtIesExtension_IdDrbIdsTakenintouse{
			IdDrbIdsTakenintouse: IDDrbIDsTakenintouse,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePdusessionResourceSetupResponseInfoSNterminatedExtIesExtensionIDDrbIDsTakenintouse() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreatePdusessionResourceSetupResponseInfoSNterminatedExtIesExtensionIDRedundantDlNgUTnlatNgRan(IDRedundantDlNgUTnlatNgRan *xnapiesv1.UptransportLayerInformation) (*xnapiesv1.PdusessionResourceSetupResponseInfoSNterminatedExtIesExtension, error) {

	item := &xnapiesv1.PdusessionResourceSetupResponseInfoSNterminatedExtIesExtension{
		PdusessionResourceSetupResponseInfoSnterminatedExtIes: &xnapiesv1.PdusessionResourceSetupResponseInfoSNterminatedExtIesExtension_IdRedundantDlNgUTnlatNgRan{
			IdRedundantDlNgUTnlatNgRan: IDRedundantDlNgUTnlatNgRan,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePdusessionResourceSetupResponseInfoSNterminatedExtIesExtensionIDRedundantDlNgUTnlatNgRan() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreatePdusessionResourceSetupResponseInfoSNterminatedExtIesExtensionIDUsedRsninformation(IDUsedRsninformation *xnapiesv1.RedundantPdusessionInformation) (*xnapiesv1.PdusessionResourceSetupResponseInfoSNterminatedExtIesExtension, error) {

	item := &xnapiesv1.PdusessionResourceSetupResponseInfoSNterminatedExtIesExtension{
		PdusessionResourceSetupResponseInfoSnterminatedExtIes: &xnapiesv1.PdusessionResourceSetupResponseInfoSNterminatedExtIesExtension_IdUsedRsninformation{
			IdUsedRsninformation: IDUsedRsninformation,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePdusessionResourceSetupResponseInfoSNterminatedExtIesExtensionIDUsedRsninformation() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateDrbsToBeSetupListSetupResponseSNterminatedItemExtIesExtensionIDAdditionalPdcpDuplicationTnlList(IDAdditionalPdcpDuplicationTnlList *xnapiesv1.AdditionalPDcpDuplicationTNlList) (*xnapiesv1.DrbsToBeSetupListSetupResponseSNterminatedItemExtIesExtension, error) {

	item := &xnapiesv1.DrbsToBeSetupListSetupResponseSNterminatedItemExtIesExtension{
		DrbsToBeSetupListSetupResponseSnterminatedItemExtIes: &xnapiesv1.DrbsToBeSetupListSetupResponseSNterminatedItemExtIesExtension_IdAdditionalPdcpDuplicationTnlList{
			IdAdditionalPdcpDuplicationTnlList: IDAdditionalPdcpDuplicationTnlList,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateDrbsToBeSetupListSetupResponseSNterminatedItemExtIesExtensionIDAdditionalPdcpDuplicationTnlList() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateDrbsToBeSetupListSetupResponseSNterminatedItemExtIesExtensionIDRlcduplicationInformation(IDRlcduplicationInformation *xnapiesv1.RlcduplicationInformation) (*xnapiesv1.DrbsToBeSetupListSetupResponseSNterminatedItemExtIesExtension, error) {

	item := &xnapiesv1.DrbsToBeSetupListSetupResponseSNterminatedItemExtIesExtension{
		DrbsToBeSetupListSetupResponseSnterminatedItemExtIes: &xnapiesv1.DrbsToBeSetupListSetupResponseSNterminatedItemExtIesExtension_IdRlcduplicationInformation{
			IdRlcduplicationInformation: IDRlcduplicationInformation,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateDrbsToBeSetupListSetupResponseSNterminatedItemExtIesExtensionIDRlcduplicationInformation() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateQoSflowsMappedtoDrbSetupResponseSNterminatedItemExtIesExtensionIDCurrentQoSparaSetIndex(IDCurrentQoSparaSetIndex *xnapiesv1.QoSparaSetIndex) (*xnapiesv1.QoSflowsMappedtoDrbSetupResponseSNterminatedItemExtIesExtension, error) {

	item := &xnapiesv1.QoSflowsMappedtoDrbSetupResponseSNterminatedItemExtIesExtension{
		QoSflowsMappedtoDrbSetupResponseSnterminatedItemExtIes: &xnapiesv1.QoSflowsMappedtoDrbSetupResponseSNterminatedItemExtIesExtension_IdCurrentQoSparaSetIndex{
			IdCurrentQoSparaSetIndex: IDCurrentQoSparaSetIndex,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateQoSflowsMappedtoDrbSetupResponseSNterminatedItemExtIesExtensionIDCurrentQoSparaSetIndex() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateQoSflowsMappedtoDrbSetupResponseSNterminatedItemExtIesExtensionIDSourceDlforwardingIpaddress(IDSourceDlforwardingIpaddress *xnapiesv1.TransportLayerAddress) (*xnapiesv1.QoSflowsMappedtoDrbSetupResponseSNterminatedItemExtIesExtension, error) {

	item := &xnapiesv1.QoSflowsMappedtoDrbSetupResponseSNterminatedItemExtIesExtension{
		QoSflowsMappedtoDrbSetupResponseSnterminatedItemExtIes: &xnapiesv1.QoSflowsMappedtoDrbSetupResponseSNterminatedItemExtIesExtension_IdSourceDlforwardingIpaddress{
			IdSourceDlforwardingIpaddress: IDSourceDlforwardingIpaddress,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateQoSflowsMappedtoDrbSetupResponseSNterminatedItemExtIesExtensionIDSourceDlforwardingIpaddress() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateDrbsToBeSetupListSetupMNterminatedItemExtIesExtensionIDAdditionalPdcpDuplicationTnlList(IDAdditionalPdcpDuplicationTnlList *xnapiesv1.AdditionalPDcpDuplicationTNlList) (*xnapiesv1.DrbsToBeSetupListSetupMNterminatedItemExtIesExtension, error) {

	item := &xnapiesv1.DrbsToBeSetupListSetupMNterminatedItemExtIesExtension{
		DrbsToBeSetupListSetupMnterminatedItemExtIes: &xnapiesv1.DrbsToBeSetupListSetupMNterminatedItemExtIesExtension_IdAdditionalPdcpDuplicationTnlList{
			IdAdditionalPdcpDuplicationTnlList: IDAdditionalPdcpDuplicationTnlList,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateDrbsToBeSetupListSetupMNterminatedItemExtIesExtensionIDAdditionalPdcpDuplicationTnlList() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateDrbsToBeSetupListSetupMNterminatedItemExtIesExtensionIDRlcduplicationInformation(IDRlcduplicationInformation *xnapiesv1.RlcduplicationInformation) (*xnapiesv1.DrbsToBeSetupListSetupMNterminatedItemExtIesExtension, error) {

	item := &xnapiesv1.DrbsToBeSetupListSetupMNterminatedItemExtIesExtension{
		DrbsToBeSetupListSetupMnterminatedItemExtIes: &xnapiesv1.DrbsToBeSetupListSetupMNterminatedItemExtIesExtension_IdRlcduplicationInformation{
			IdRlcduplicationInformation: IDRlcduplicationInformation,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateDrbsToBeSetupListSetupMNterminatedItemExtIesExtensionIDRlcduplicationInformation() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateDrbsAdmittedListSetupResponseMNterminatedItemExtIesExtensionIDAdditionalPdcpDuplicationTnlList(IDAdditionalPdcpDuplicationTnlList *xnapiesv1.AdditionalPDcpDuplicationTNlList) (*xnapiesv1.DrbsAdmittedListSetupResponseMNterminatedItemExtIesExtension, error) {

	item := &xnapiesv1.DrbsAdmittedListSetupResponseMNterminatedItemExtIesExtension{
		DrbsAdmittedListSetupResponseMnterminatedItemExtIes: &xnapiesv1.DrbsAdmittedListSetupResponseMNterminatedItemExtIesExtension_IdAdditionalPdcpDuplicationTnlList{
			IdAdditionalPdcpDuplicationTnlList: IDAdditionalPdcpDuplicationTnlList,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateDrbsAdmittedListSetupResponseMNterminatedItemExtIesExtensionIDAdditionalPdcpDuplicationTnlList() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateDrbsAdmittedListSetupResponseMNterminatedItemExtIesExtensionIDQoSflowsMappedtoDrbSetupResponseMnterminated(IDQoSflowsMappedtoDrbSetupResponseMnterminated *xnapiesv1.QoSflowsMappedtoDrbSetupResponseMNterminated) (*xnapiesv1.DrbsAdmittedListSetupResponseMNterminatedItemExtIesExtension, error) {

	item := &xnapiesv1.DrbsAdmittedListSetupResponseMNterminatedItemExtIesExtension{
		DrbsAdmittedListSetupResponseMnterminatedItemExtIes: &xnapiesv1.DrbsAdmittedListSetupResponseMNterminatedItemExtIesExtension_IdQoSflowsMappedtoDrbSetupResponseMnterminated{
			IdQoSflowsMappedtoDrbSetupResponseMnterminated: IDQoSflowsMappedtoDrbSetupResponseMnterminated,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateDrbsAdmittedListSetupResponseMNterminatedItemExtIesExtensionIDQoSflowsMappedtoDrbSetupResponseMnterminated() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreatePdusessionResourceModificationInfoSNterminatedExtIesExtensionIDPdusessionCommonNetworkInstance(IDPdusessionCommonNetworkInstance *xnapiesv1.PdusessionCommonNetworkInstance) (*xnapiesv1.PdusessionResourceModificationInfoSNterminatedExtIesExtension, error) {

	item := &xnapiesv1.PdusessionResourceModificationInfoSNterminatedExtIesExtension{
		PdusessionResourceModificationInfoSnterminatedExtIes: &xnapiesv1.PdusessionResourceModificationInfoSNterminatedExtIesExtension_IdPdusessionCommonNetworkInstance{
			IdPdusessionCommonNetworkInstance: IDPdusessionCommonNetworkInstance,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePdusessionResourceModificationInfoSNterminatedExtIesExtensionIDPdusessionCommonNetworkInstance() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreatePdusessionResourceModificationInfoSNterminatedExtIesExtensionIDDefaultDrbAllowed(IDDefaultDrbAllowed xnapiesv1.DefaultDrbAllowed) (*xnapiesv1.PdusessionResourceModificationInfoSNterminatedExtIesExtension, error) {

	item := &xnapiesv1.PdusessionResourceModificationInfoSNterminatedExtIesExtension{
		PdusessionResourceModificationInfoSnterminatedExtIes: &xnapiesv1.PdusessionResourceModificationInfoSNterminatedExtIesExtension_IdDefaultDrbAllowed{
			IdDefaultDrbAllowed: IDDefaultDrbAllowed,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePdusessionResourceModificationInfoSNterminatedExtIesExtensionIDDefaultDrbAllowed() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreatePdusessionResourceModificationInfoSNterminatedExtIesExtensionIDNonGbrresourcesOffered(IDNonGbrresourcesOffered xnapiesv1.NonGbrresourcesOffered) (*xnapiesv1.PdusessionResourceModificationInfoSNterminatedExtIesExtension, error) {

	item := &xnapiesv1.PdusessionResourceModificationInfoSNterminatedExtIesExtension{
		PdusessionResourceModificationInfoSnterminatedExtIes: &xnapiesv1.PdusessionResourceModificationInfoSNterminatedExtIesExtension_IdNonGbrresourcesOffered{
			IdNonGbrresourcesOffered: IDNonGbrresourcesOffered,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePdusessionResourceModificationInfoSNterminatedExtIesExtensionIDNonGbrresourcesOffered() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreatePdusessionResourceModificationInfoSNterminatedExtIesExtensionIDRedundantUlNgUTnlatUpf(IDRedundantUlNgUTnlatUpf *xnapiesv1.UptransportLayerInformation) (*xnapiesv1.PdusessionResourceModificationInfoSNterminatedExtIesExtension, error) {

	item := &xnapiesv1.PdusessionResourceModificationInfoSNterminatedExtIesExtension{
		PdusessionResourceModificationInfoSnterminatedExtIes: &xnapiesv1.PdusessionResourceModificationInfoSNterminatedExtIesExtension_IdRedundantUlNgUTnlatUpf{
			IdRedundantUlNgUTnlatUpf: IDRedundantUlNgUTnlatUpf,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePdusessionResourceModificationInfoSNterminatedExtIesExtensionIDRedundantUlNgUTnlatUpf() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreatePdusessionResourceModificationInfoSNterminatedExtIesExtensionIDRedundantCommonNetworkInstance(IDRedundantCommonNetworkInstance *xnapiesv1.PdusessionCommonNetworkInstance) (*xnapiesv1.PdusessionResourceModificationInfoSNterminatedExtIesExtension, error) {

	item := &xnapiesv1.PdusessionResourceModificationInfoSNterminatedExtIesExtension{
		PdusessionResourceModificationInfoSnterminatedExtIes: &xnapiesv1.PdusessionResourceModificationInfoSNterminatedExtIesExtension_IdRedundantCommonNetworkInstance{
			IdRedundantCommonNetworkInstance: IDRedundantCommonNetworkInstance,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePdusessionResourceModificationInfoSNterminatedExtIesExtensionIDRedundantCommonNetworkInstance() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreatePdusessionResourceModificationInfoSNterminatedExtIesExtensionIDSecurityIndication(IDSecurityIndication *xnapiesv1.SecurityIndication) (*xnapiesv1.PdusessionResourceModificationInfoSNterminatedExtIesExtension, error) {

	item := &xnapiesv1.PdusessionResourceModificationInfoSNterminatedExtIesExtension{
		PdusessionResourceModificationInfoSnterminatedExtIes: &xnapiesv1.PdusessionResourceModificationInfoSNterminatedExtIesExtension_IdSecurityIndication{
			IdSecurityIndication: IDSecurityIndication,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePdusessionResourceModificationInfoSNterminatedExtIesExtensionIDSecurityIndication() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateQoSflowsToBeSetupListModifiedSNterminatedItemExtIesExtensionIDTsctrafficCharacteristics(IDTsctrafficCharacteristics *xnapiesv1.TsctrafficCharacteristics) (*xnapiesv1.QoSflowsToBeSetupListModifiedSNterminatedItemExtIesExtension, error) {

	item := &xnapiesv1.QoSflowsToBeSetupListModifiedSNterminatedItemExtIesExtension{
		QoSflowsToBeSetupListModifiedSnterminatedItemExtIes: &xnapiesv1.QoSflowsToBeSetupListModifiedSNterminatedItemExtIesExtension_IdTsctrafficCharacteristics{
			IdTsctrafficCharacteristics: IDTsctrafficCharacteristics,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateQoSflowsToBeSetupListModifiedSNterminatedItemExtIesExtensionIDTsctrafficCharacteristics() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateQoSflowsToBeSetupListModifiedSNterminatedItemExtIesExtensionIDRedundantQoSflowIndicator(IDRedundantQoSflowIndicator xnapiesv1.RedundantQoSflowIndicator) (*xnapiesv1.QoSflowsToBeSetupListModifiedSNterminatedItemExtIesExtension, error) {

	item := &xnapiesv1.QoSflowsToBeSetupListModifiedSNterminatedItemExtIesExtension{
		QoSflowsToBeSetupListModifiedSnterminatedItemExtIes: &xnapiesv1.QoSflowsToBeSetupListModifiedSNterminatedItemExtIesExtension_IdRedundantQoSflowIndicator{
			IdRedundantQoSflowIndicator: IDRedundantQoSflowIndicator,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateQoSflowsToBeSetupListModifiedSNterminatedItemExtIesExtensionIDRedundantQoSflowIndicator() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreatePdusessionResourceModificationResponseInfoSNterminatedExtIesExtensionIDDrbIDsTakenintouse(IDDrbIDsTakenintouse *xnapiesv1.DrbList) (*xnapiesv1.PdusessionResourceModificationResponseInfoSNterminatedExtIesExtension, error) {

	item := &xnapiesv1.PdusessionResourceModificationResponseInfoSNterminatedExtIesExtension{
		PdusessionResourceModificationResponseInfoSnterminatedExtIes: &xnapiesv1.PdusessionResourceModificationResponseInfoSNterminatedExtIesExtension_IdDrbIdsTakenintouse{
			IdDrbIdsTakenintouse: IDDrbIDsTakenintouse,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePdusessionResourceModificationResponseInfoSNterminatedExtIesExtensionIDDrbIDsTakenintouse() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreatePdusessionResourceModificationResponseInfoSNterminatedExtIesExtensionIDRedundantDlNgUTnlatNgRan(IDRedundantDlNgUTnlatNgRan *xnapiesv1.UptransportLayerInformation) (*xnapiesv1.PdusessionResourceModificationResponseInfoSNterminatedExtIesExtension, error) {

	item := &xnapiesv1.PdusessionResourceModificationResponseInfoSNterminatedExtIesExtension{
		PdusessionResourceModificationResponseInfoSnterminatedExtIes: &xnapiesv1.PdusessionResourceModificationResponseInfoSNterminatedExtIesExtension_IdRedundantDlNgUTnlatNgRan{
			IdRedundantDlNgUTnlatNgRan: IDRedundantDlNgUTnlatNgRan,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePdusessionResourceModificationResponseInfoSNterminatedExtIesExtensionIDRedundantDlNgUTnlatNgRan() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreatePdusessionResourceModificationResponseInfoSNterminatedExtIesExtensionIDSecurityResult(IDSecurityResult *xnapiesv1.SecurityResult) (*xnapiesv1.PdusessionResourceModificationResponseInfoSNterminatedExtIesExtension, error) {

	item := &xnapiesv1.PdusessionResourceModificationResponseInfoSNterminatedExtIesExtension{
		PdusessionResourceModificationResponseInfoSnterminatedExtIes: &xnapiesv1.PdusessionResourceModificationResponseInfoSNterminatedExtIesExtension_IdSecurityResult{
			IdSecurityResult: IDSecurityResult,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePdusessionResourceModificationResponseInfoSNterminatedExtIesExtensionIDSecurityResult() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateDrbsToBeModifiedListModificationResponseSNterminatedItemExtIesExtensionIDAdditionalPdcpDuplicationTnlList(IDAdditionalPdcpDuplicationTnlList *xnapiesv1.AdditionalPDcpDuplicationTNlList) (*xnapiesv1.DrbsToBeModifiedListModificationResponseSNterminatedItemExtIesExtension, error) {

	item := &xnapiesv1.DrbsToBeModifiedListModificationResponseSNterminatedItemExtIesExtension{
		DrbsToBeModifiedListModificationResponseSnterminatedItemExtIes: &xnapiesv1.DrbsToBeModifiedListModificationResponseSNterminatedItemExtIesExtension_IdAdditionalPdcpDuplicationTnlList{
			IdAdditionalPdcpDuplicationTnlList: IDAdditionalPdcpDuplicationTnlList,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateDrbsToBeModifiedListModificationResponseSNterminatedItemExtIesExtensionIDAdditionalPdcpDuplicationTnlList() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateDrbsToBeModifiedListModificationResponseSNterminatedItemExtIesExtensionIDRlcduplicationInformation(IDRlcduplicationInformation *xnapiesv1.RlcduplicationInformation) (*xnapiesv1.DrbsToBeModifiedListModificationResponseSNterminatedItemExtIesExtension, error) {

	item := &xnapiesv1.DrbsToBeModifiedListModificationResponseSNterminatedItemExtIesExtension{
		DrbsToBeModifiedListModificationResponseSnterminatedItemExtIes: &xnapiesv1.DrbsToBeModifiedListModificationResponseSNterminatedItemExtIesExtension_IdRlcduplicationInformation{
			IdRlcduplicationInformation: IDRlcduplicationInformation,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateDrbsToBeModifiedListModificationResponseSNterminatedItemExtIesExtensionIDRlcduplicationInformation() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateDrbsToBeModifiedListModificationResponseSNterminatedItemExtIesExtensionIDSecondarySnUlPdcpUpTnlinfo(IDSecondarySnUlPdcpUpTnlinfo *xnapiesv1.UptransportParameters) (*xnapiesv1.DrbsToBeModifiedListModificationResponseSNterminatedItemExtIesExtension, error) {

	item := &xnapiesv1.DrbsToBeModifiedListModificationResponseSNterminatedItemExtIesExtension{
		DrbsToBeModifiedListModificationResponseSnterminatedItemExtIes: &xnapiesv1.DrbsToBeModifiedListModificationResponseSNterminatedItemExtIesExtension_IdSecondarySnUlPdcpUpTnlinfo{
			IdSecondarySnUlPdcpUpTnlinfo: IDSecondarySnUlPdcpUpTnlinfo,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateDrbsToBeModifiedListModificationResponseSNterminatedItemExtIesExtensionIDSecondarySnUlPdcpUpTnlinfo() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateDrbsToBeModifiedListModificationResponseSNterminatedItemExtIesExtensionIDPdcpDuplicationConfiguration(IDPdcpDuplicationConfiguration xnapiesv1.PdcpduplicationConfiguration) (*xnapiesv1.DrbsToBeModifiedListModificationResponseSNterminatedItemExtIesExtension, error) {

	item := &xnapiesv1.DrbsToBeModifiedListModificationResponseSNterminatedItemExtIesExtension{
		DrbsToBeModifiedListModificationResponseSnterminatedItemExtIes: &xnapiesv1.DrbsToBeModifiedListModificationResponseSNterminatedItemExtIesExtension_IdPdcpDuplicationConfiguration{
			IdPdcpDuplicationConfiguration: IDPdcpDuplicationConfiguration,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateDrbsToBeModifiedListModificationResponseSNterminatedItemExtIesExtensionIDPdcpDuplicationConfiguration() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateDrbsToBeModifiedListModificationResponseSNterminatedItemExtIesExtensionIDDuplicationActivation(IDDuplicationActivation xnapiesv1.DuplicationActivation) (*xnapiesv1.DrbsToBeModifiedListModificationResponseSNterminatedItemExtIesExtension, error) {

	item := &xnapiesv1.DrbsToBeModifiedListModificationResponseSNterminatedItemExtIesExtension{
		DrbsToBeModifiedListModificationResponseSnterminatedItemExtIes: &xnapiesv1.DrbsToBeModifiedListModificationResponseSNterminatedItemExtIesExtension_IdDuplicationActivation{
			IdDuplicationActivation: IDDuplicationActivation,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateDrbsToBeModifiedListModificationResponseSNterminatedItemExtIesExtensionIDDuplicationActivation() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateDrbsToBeModifiedListModificationMNterminatedItemExtIesExtensionIDAdditionalPdcpDuplicationTnlList(IDAdditionalPdcpDuplicationTnlList *xnapiesv1.AdditionalPDcpDuplicationTNlList) (*xnapiesv1.DrbsToBeModifiedListModificationMNterminatedItemExtIesExtension, error) {

	item := &xnapiesv1.DrbsToBeModifiedListModificationMNterminatedItemExtIesExtension{
		DrbsToBeModifiedListModificationMnterminatedItemExtIes: &xnapiesv1.DrbsToBeModifiedListModificationMNterminatedItemExtIesExtension_IdAdditionalPdcpDuplicationTnlList{
			IdAdditionalPdcpDuplicationTnlList: IDAdditionalPdcpDuplicationTnlList,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateDrbsToBeModifiedListModificationMNterminatedItemExtIesExtensionIDAdditionalPdcpDuplicationTnlList() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateDrbsToBeModifiedListModificationMNterminatedItemExtIesExtensionIDRlcduplicationInformation(IDRlcduplicationInformation *xnapiesv1.RlcduplicationInformation) (*xnapiesv1.DrbsToBeModifiedListModificationMNterminatedItemExtIesExtension, error) {

	item := &xnapiesv1.DrbsToBeModifiedListModificationMNterminatedItemExtIesExtension{
		DrbsToBeModifiedListModificationMnterminatedItemExtIes: &xnapiesv1.DrbsToBeModifiedListModificationMNterminatedItemExtIesExtension_IdRlcduplicationInformation{
			IdRlcduplicationInformation: IDRlcduplicationInformation,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateDrbsToBeModifiedListModificationMNterminatedItemExtIesExtensionIDRlcduplicationInformation() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateDrbsAdmittedListModificationResponseMNterminatedItemExtIesExtensionIDAdditionalPdcpDuplicationTnlList(IDAdditionalPdcpDuplicationTnlList *xnapiesv1.AdditionalPDcpDuplicationTNlList) (*xnapiesv1.DrbsAdmittedListModificationResponseMNterminatedItemExtIesExtension, error) {

	item := &xnapiesv1.DrbsAdmittedListModificationResponseMNterminatedItemExtIesExtension{
		DrbsAdmittedListModificationResponseMnterminatedItemExtIes: &xnapiesv1.DrbsAdmittedListModificationResponseMNterminatedItemExtIesExtension_IdAdditionalPdcpDuplicationTnlList{
			IdAdditionalPdcpDuplicationTnlList: IDAdditionalPdcpDuplicationTnlList,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateDrbsAdmittedListModificationResponseMNterminatedItemExtIesExtensionIDAdditionalPdcpDuplicationTnlList() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateDrbsAdmittedListModificationResponseMNterminatedItemExtIesExtensionIDQoSflowsMappedtoDrbSetupResponseMnterminated(IDQoSflowsMappedtoDrbSetupResponseMnterminated *xnapiesv1.QoSflowsMappedtoDrbSetupResponseMNterminated) (*xnapiesv1.DrbsAdmittedListModificationResponseMNterminatedItemExtIesExtension, error) {

	item := &xnapiesv1.DrbsAdmittedListModificationResponseMNterminatedItemExtIesExtension{
		DrbsAdmittedListModificationResponseMnterminatedItemExtIes: &xnapiesv1.DrbsAdmittedListModificationResponseMNterminatedItemExtIesExtension_IdQoSflowsMappedtoDrbSetupResponseMnterminated{
			IdQoSflowsMappedtoDrbSetupResponseMnterminated: IDQoSflowsMappedtoDrbSetupResponseMnterminated,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateDrbsAdmittedListModificationResponseMNterminatedItemExtIesExtensionIDQoSflowsMappedtoDrbSetupResponseMnterminated() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateDrbsToBeSetupListModRqdSNterminatedItemExtIesExtensionIDAdditionalPdcpDuplicationTnlList(IDAdditionalPdcpDuplicationTnlList *xnapiesv1.AdditionalPDcpDuplicationTNlList) (*xnapiesv1.DrbsToBeSetupListModRqdSNterminatedItemExtIesExtension, error) {

	item := &xnapiesv1.DrbsToBeSetupListModRqdSNterminatedItemExtIesExtension{
		DrbsToBeSetupListModRqdSnterminatedItemExtIes: &xnapiesv1.DrbsToBeSetupListModRqdSNterminatedItemExtIesExtension_IdAdditionalPdcpDuplicationTnlList{
			IdAdditionalPdcpDuplicationTnlList: IDAdditionalPdcpDuplicationTnlList,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateDrbsToBeSetupListModRqdSNterminatedItemExtIesExtensionIDAdditionalPdcpDuplicationTnlList() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateDrbsToBeSetupListModRqdSNterminatedItemExtIesExtensionIDRlcduplicationInformation(IDRlcduplicationInformation *xnapiesv1.RlcduplicationInformation) (*xnapiesv1.DrbsToBeSetupListModRqdSNterminatedItemExtIesExtension, error) {

	item := &xnapiesv1.DrbsToBeSetupListModRqdSNterminatedItemExtIesExtension{
		DrbsToBeSetupListModRqdSnterminatedItemExtIes: &xnapiesv1.DrbsToBeSetupListModRqdSNterminatedItemExtIesExtension_IdRlcduplicationInformation{
			IdRlcduplicationInformation: IDRlcduplicationInformation,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateDrbsToBeSetupListModRqdSNterminatedItemExtIesExtensionIDRlcduplicationInformation() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateDrbsToBeModifiedListModRqdSNterminatedItemExtIesExtensionIDAdditionalPdcpDuplicationTnlList(IDAdditionalPdcpDuplicationTnlList *xnapiesv1.AdditionalPDcpDuplicationTNlList) (*xnapiesv1.DrbsToBeModifiedListModRqdSNterminatedItemExtIesExtension, error) {

	item := &xnapiesv1.DrbsToBeModifiedListModRqdSNterminatedItemExtIesExtension{
		DrbsToBeModifiedListModRqdSnterminatedItemExtIes: &xnapiesv1.DrbsToBeModifiedListModRqdSNterminatedItemExtIesExtension_IdAdditionalPdcpDuplicationTnlList{
			IdAdditionalPdcpDuplicationTnlList: IDAdditionalPdcpDuplicationTnlList,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateDrbsToBeModifiedListModRqdSNterminatedItemExtIesExtensionIDAdditionalPdcpDuplicationTnlList() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateDrbsToBeModifiedListModRqdSNterminatedItemExtIesExtensionIDRlcduplicationInformation(IDRlcduplicationInformation *xnapiesv1.RlcduplicationInformation) (*xnapiesv1.DrbsToBeModifiedListModRqdSNterminatedItemExtIesExtension, error) {

	item := &xnapiesv1.DrbsToBeModifiedListModRqdSNterminatedItemExtIesExtension{
		DrbsToBeModifiedListModRqdSnterminatedItemExtIes: &xnapiesv1.DrbsToBeModifiedListModRqdSNterminatedItemExtIesExtension_IdRlcduplicationInformation{
			IdRlcduplicationInformation: IDRlcduplicationInformation,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateDrbsToBeModifiedListModRqdSNterminatedItemExtIesExtensionIDRlcduplicationInformation() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateQoScharacteristicsNonDynamic(nonDynamic *xnapiesv1.NonDynamic5Qidescriptor) (*xnapiesv1.QoScharacteristics, error) {

	item := &xnapiesv1.QoScharacteristics{
		QoScharacteristics: &xnapiesv1.QoScharacteristics_NonDynamic{
			NonDynamic: nonDynamic,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateQoScharacteristicsNonDynamic() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateQoScharacteristicsDynamic(dynamic *xnapiesv1.Dynamic5Qidescriptor) (*xnapiesv1.QoScharacteristics, error) {

	item := &xnapiesv1.QoScharacteristics{
		QoScharacteristics: &xnapiesv1.QoScharacteristics_Dynamic{
			Dynamic: dynamic,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateQoScharacteristicsDynamic() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateQoScharacteristicsChoiceExtension(choiceExtension *xnapiesv1.QoScharacteristicsExtIes) (*xnapiesv1.QoScharacteristics, error) {

	item := &xnapiesv1.QoScharacteristics{
		QoScharacteristics: &xnapiesv1.QoScharacteristics_ChoiceExtension{
			ChoiceExtension: choiceExtension,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateQoScharacteristicsChoiceExtension() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateQoSflowLevelQoSparametersExtIesExtensionIDQoSmonitoringRequest(IDQoSmonitoringRequest xnapiesv1.QosMonitoringRequest) (*xnapiesv1.QoSflowLevelQoSparametersExtIesExtension, error) {

	item := &xnapiesv1.QoSflowLevelQoSparametersExtIesExtension{
		QoSflowLevelQoSparametersExtIes: &xnapiesv1.QoSflowLevelQoSparametersExtIesExtension_IdQoSmonitoringRequest{
			IdQoSmonitoringRequest: IDQoSmonitoringRequest,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateQoSflowLevelQoSparametersExtIesExtensionIDQoSmonitoringRequest() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateQoSflowLevelQoSparametersExtIesExtensionIDQosMonitoringReportingFrequency(IDQosMonitoringReportingFrequency *xnapiesv1.QosMonitoringReportingFrequency) (*xnapiesv1.QoSflowLevelQoSparametersExtIesExtension, error) {

	item := &xnapiesv1.QoSflowLevelQoSparametersExtIesExtension{
		QoSflowLevelQoSparametersExtIes: &xnapiesv1.QoSflowLevelQoSparametersExtIesExtension_IdQosMonitoringReportingFrequency{
			IdQosMonitoringReportingFrequency: IDQosMonitoringReportingFrequency,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateQoSflowLevelQoSparametersExtIesExtensionIDQosMonitoringReportingFrequency() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateQoSflowLevelQoSparametersExtIesExtensionIDQoSmonitoringDisabled(IDQoSmonitoringDisabled xnapiesv1.QoSmonitoringDisabled) (*xnapiesv1.QoSflowLevelQoSparametersExtIesExtension, error) {

	item := &xnapiesv1.QoSflowLevelQoSparametersExtIesExtension{
		QoSflowLevelQoSparametersExtIes: &xnapiesv1.QoSflowLevelQoSparametersExtIesExtension_IdQoSmonitoringDisabled{
			IdQoSmonitoringDisabled: IDQoSmonitoringDisabled,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateQoSflowLevelQoSparametersExtIesExtensionIDQoSmonitoringDisabled() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateQoSflowsToBeSetupItemExtIesExtensionIDTsctrafficCharacteristics(IDTsctrafficCharacteristics *xnapiesv1.TsctrafficCharacteristics) (*xnapiesv1.QoSflowsToBeSetupItemExtIesExtension, error) {

	item := &xnapiesv1.QoSflowsToBeSetupItemExtIesExtension{
		QoSflowsToBeSetupItemExtIes: &xnapiesv1.QoSflowsToBeSetupItemExtIesExtension_IdTsctrafficCharacteristics{
			IdTsctrafficCharacteristics: IDTsctrafficCharacteristics,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateQoSflowsToBeSetupItemExtIesExtensionIDTsctrafficCharacteristics() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateQoSflowsToBeSetupItemExtIesExtensionIDRedundantQoSflowIndicator(IDRedundantQoSflowIndicator xnapiesv1.RedundantQoSflowIndicator) (*xnapiesv1.QoSflowsToBeSetupItemExtIesExtension, error) {

	item := &xnapiesv1.QoSflowsToBeSetupItemExtIesExtension{
		QoSflowsToBeSetupItemExtIes: &xnapiesv1.QoSflowsToBeSetupItemExtIesExtension_IdRedundantQoSflowIndicator{
			IdRedundantQoSflowIndicator: IDRedundantQoSflowIndicator,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateQoSflowsToBeSetupItemExtIesExtensionIDRedundantQoSflowIndicator() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateRadioResourceStatusNgENbRadioResourceStatus(ngENbRadioResourceStatus *xnapiesv1.NgeNbRadioResourceStatus) (*xnapiesv1.RadioResourceStatus, error) {

	item := &xnapiesv1.RadioResourceStatus{
		RadioResourceStatus: &xnapiesv1.RadioResourceStatus_NgENbRadioResourceStatus{
			NgENbRadioResourceStatus: ngENbRadioResourceStatus,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRadioResourceStatusNgENbRadioResourceStatus() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateRadioResourceStatusGNbRadioResourceStatus(gNbRadioResourceStatus *xnapiesv1.GnbRadioResourceStatus) (*xnapiesv1.RadioResourceStatus, error) {

	item := &xnapiesv1.RadioResourceStatus{
		RadioResourceStatus: &xnapiesv1.RadioResourceStatus_GNbRadioResourceStatus{
			GNbRadioResourceStatus: gNbRadioResourceStatus,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRadioResourceStatusGNbRadioResourceStatus() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateRadioResourceStatusChoiceExtension(choiceExtension *xnapiesv1.RadioResourceStatusExtIes) (*xnapiesv1.RadioResourceStatus, error) {

	item := &xnapiesv1.RadioResourceStatus{
		RadioResourceStatus: &xnapiesv1.RadioResourceStatus_ChoiceExtension{
			ChoiceExtension: choiceExtension,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRadioResourceStatusChoiceExtension() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateRanpagingAreaChoiceCellList(cellList *xnapiesv1.NgRAnCellIdentityListinRanpagingArea) (*xnapiesv1.RanpagingAreaChoice, error) {

	item := &xnapiesv1.RanpagingAreaChoice{
		RanpagingAreaChoice: &xnapiesv1.RanpagingAreaChoice_CellList{
			CellList: cellList,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRanpagingAreaChoiceCellList() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateRanpagingAreaChoiceRAnareaIDList(rAnareaIDList *xnapiesv1.RanareaIdList) (*xnapiesv1.RanpagingAreaChoice, error) {

	item := &xnapiesv1.RanpagingAreaChoice{
		RanpagingAreaChoice: &xnapiesv1.RanpagingAreaChoice_RAnareaIdList{
			RAnareaIdList: rAnareaIDList,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRanpagingAreaChoiceRAnareaIDList() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateRanpagingAreaChoiceChoiceExtension(choiceExtension *xnapiesv1.RanpagingAreaChoiceExtIes) (*xnapiesv1.RanpagingAreaChoice, error) {

	item := &xnapiesv1.RanpagingAreaChoice{
		RanpagingAreaChoice: &xnapiesv1.RanpagingAreaChoice_ChoiceExtension{
			ChoiceExtension: choiceExtension,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRanpagingAreaChoiceChoiceExtension() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateReportTypePeriodical(periodical *xnapiesv1.Periodical) (*xnapiesv1.ReportType, error) {

	item := &xnapiesv1.ReportType{
		ReportType: &xnapiesv1.ReportType_Periodical{
			Periodical: periodical,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateReportTypePeriodical() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateReportTypeEventTriggered(eventTriggered *xnapiesv1.EventTriggered) (*xnapiesv1.ReportType, error) {

	item := &xnapiesv1.ReportType{
		ReportType: &xnapiesv1.ReportType_EventTriggered{
			EventTriggered: eventTriggered,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateReportTypeEventTriggered() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateResetRequestTypeInfoFullReset(fullReset *xnapiesv1.ResetRequestTypeInfoFull) (*xnapiesv1.ResetRequestTypeInfo, error) {

	item := &xnapiesv1.ResetRequestTypeInfo{
		ResetRequestTypeInfo: &xnapiesv1.ResetRequestTypeInfo_FullReset{
			FullReset: fullReset,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateResetRequestTypeInfoFullReset() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateResetRequestTypeInfoPartialReset(partialReset *xnapiesv1.ResetRequestTypeInfoPartial) (*xnapiesv1.ResetRequestTypeInfo, error) {

	item := &xnapiesv1.ResetRequestTypeInfo{
		ResetRequestTypeInfo: &xnapiesv1.ResetRequestTypeInfo_PartialReset{
			PartialReset: partialReset,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateResetRequestTypeInfoPartialReset() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateResetRequestTypeInfoChoiceExtension(choiceExtension *xnapiesv1.ResetRequestTypeInfoExtIes) (*xnapiesv1.ResetRequestTypeInfo, error) {

	item := &xnapiesv1.ResetRequestTypeInfo{
		ResetRequestTypeInfo: &xnapiesv1.ResetRequestTypeInfo_ChoiceExtension{
			ChoiceExtension: choiceExtension,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateResetRequestTypeInfoChoiceExtension() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateResetResponseTypeInfoFullReset(fullReset *xnapiesv1.ResetResponseTypeInfoFull) (*xnapiesv1.ResetResponseTypeInfo, error) {

	item := &xnapiesv1.ResetResponseTypeInfo{
		ResetResponseTypeInfo: &xnapiesv1.ResetResponseTypeInfo_FullReset{
			FullReset: fullReset,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateResetResponseTypeInfoFullReset() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateResetResponseTypeInfoPartialReset(partialReset *xnapiesv1.ResetResponseTypeInfoPartial) (*xnapiesv1.ResetResponseTypeInfo, error) {

	item := &xnapiesv1.ResetResponseTypeInfo{
		ResetResponseTypeInfo: &xnapiesv1.ResetResponseTypeInfo_PartialReset{
			PartialReset: partialReset,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateResetResponseTypeInfoPartialReset() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateResetResponseTypeInfoChoiceExtension(choiceExtension *xnapiesv1.ResetResponseTypeInfoExtIes) (*xnapiesv1.ResetResponseTypeInfo, error) {

	item := &xnapiesv1.ResetResponseTypeInfo{
		ResetResponseTypeInfo: &xnapiesv1.ResetResponseTypeInfo_ChoiceExtension{
			ChoiceExtension: choiceExtension,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateResetResponseTypeInfoChoiceExtension() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateRrcreestabInitiatedReportingRRcreestabReportingWoUerlfreport(rRcreestabReportingWoUerlfreport *xnapiesv1.RrcreestabInitiatedReportingwoUErlfreport) (*xnapiesv1.RrcreestabInitiatedReporting, error) {

	item := &xnapiesv1.RrcreestabInitiatedReporting{
		RrcreestabInitiatedReporting: &xnapiesv1.RrcreestabInitiatedReporting_RRcreestabReportingWoUerlfreport{
			RRcreestabReportingWoUerlfreport: rRcreestabReportingWoUerlfreport,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRrcreestabInitiatedReportingRRcreestabReportingWoUerlfreport() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateRrcreestabInitiatedReportingRRcreestabReportingWithUerlfreport(rRcreestabReportingWithUerlfreport *xnapiesv1.RrcreestabInitiatedReportingwithUErlfreport) (*xnapiesv1.RrcreestabInitiatedReporting, error) {

	item := &xnapiesv1.RrcreestabInitiatedReporting{
		RrcreestabInitiatedReporting: &xnapiesv1.RrcreestabInitiatedReporting_RRcreestabReportingWithUerlfreport{
			RRcreestabReportingWithUerlfreport: rRcreestabReportingWithUerlfreport,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRrcreestabInitiatedReportingRRcreestabReportingWithUerlfreport() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateRrcreestabInitiatedReportingChoiceExtension(choiceExtension *xnapiesv1.RrcreestabInitiatedReportingExtIes) (*xnapiesv1.RrcreestabInitiatedReporting, error) {

	item := &xnapiesv1.RrcreestabInitiatedReporting{
		RrcreestabInitiatedReporting: &xnapiesv1.RrcreestabInitiatedReporting_ChoiceExtension{
			ChoiceExtension: choiceExtension,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRrcreestabInitiatedReportingChoiceExtension() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateRrcsetupInitiatedReportingRRcsetupReportingWithUerlfreport(rRcsetupReportingWithUerlfreport *xnapiesv1.RrcsetupInitiatedReportingwithUErlfreport) (*xnapiesv1.RrcsetupInitiatedReporting, error) {

	item := &xnapiesv1.RrcsetupInitiatedReporting{
		RrcsetupInitiatedReporting: &xnapiesv1.RrcsetupInitiatedReporting_RRcsetupReportingWithUerlfreport{
			RRcsetupReportingWithUerlfreport: rRcsetupReportingWithUerlfreport,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRrcsetupInitiatedReportingRRcsetupReportingWithUerlfreport() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateRrcsetupInitiatedReportingChoiceExtension(choiceExtension *xnapiesv1.RrcsetupInitiatedReportingExtIes) (*xnapiesv1.RrcsetupInitiatedReporting, error) {

	item := &xnapiesv1.RrcsetupInitiatedReporting{
		RrcsetupInitiatedReporting: &xnapiesv1.RrcsetupInitiatedReporting_ChoiceExtension{
			ChoiceExtension: choiceExtension,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRrcsetupInitiatedReportingChoiceExtension() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateServedCellInformationEUTraExtIesExtensionIDBplmnIDInfoEutra(IDBplmnIDInfoEutra *xnapiesv1.BplmnIDInfoEUtra) (*xnapiesv1.ServedCellInformationEUTraExtIesExtension, error) {

	item := &xnapiesv1.ServedCellInformationEUTraExtIesExtension{
		ServedCellInformationEUtraExtIes: &xnapiesv1.ServedCellInformationEUTraExtIesExtension_IdBplmnIdInfoEutra{
			IdBplmnIdInfoEutra: IDBplmnIDInfoEutra,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateServedCellInformationEUTraExtIesExtensionIDBplmnIDInfoEutra() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateServedCellInformationEUTraExtIesExtensionIDNprachconfiguration(IDNprachconfiguration *xnapiesv1.Nprachconfiguration) (*xnapiesv1.ServedCellInformationEUTraExtIesExtension, error) {

	item := &xnapiesv1.ServedCellInformationEUTraExtIesExtension{
		ServedCellInformationEUtraExtIes: &xnapiesv1.ServedCellInformationEUTraExtIesExtension_IdNprachconfiguration{
			IdNprachconfiguration: IDNprachconfiguration,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateServedCellInformationEUTraExtIesExtensionIDNprachconfiguration() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateServedCellInformationEUTraModeInfoFdd(fdd *xnapiesv1.ServedCellInformationEUTraFDdinfo) (*xnapiesv1.ServedCellInformationEUTraModeInfo, error) {

	item := &xnapiesv1.ServedCellInformationEUTraModeInfo{
		ServedCellInformationEUtraModeInfo: &xnapiesv1.ServedCellInformationEUTraModeInfo_Fdd{
			Fdd: fdd,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateServedCellInformationEUTraModeInfoFdd() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateServedCellInformationEUTraModeInfoTdd(tdd *xnapiesv1.ServedCellInformationEUTraTDdinfo) (*xnapiesv1.ServedCellInformationEUTraModeInfo, error) {

	item := &xnapiesv1.ServedCellInformationEUTraModeInfo{
		ServedCellInformationEUtraModeInfo: &xnapiesv1.ServedCellInformationEUTraModeInfo_Tdd{
			Tdd: tdd,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateServedCellInformationEUTraModeInfoTdd() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateServedCellInformationEUTraModeInfoChoiceExtension(choiceExtension *xnapiesv1.ServedCellInformationEUTraModeInfoExtIes) (*xnapiesv1.ServedCellInformationEUTraModeInfo, error) {

	item := &xnapiesv1.ServedCellInformationEUTraModeInfo{
		ServedCellInformationEUtraModeInfo: &xnapiesv1.ServedCellInformationEUTraModeInfo_ChoiceExtension{
			ChoiceExtension: choiceExtension,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateServedCellInformationEUTraModeInfoChoiceExtension() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateServedCellInformationEUTraFDdinfoExtIesExtensionIDOffsetOfNbiotChannelNumberToDlEarfcn(IDOffsetOfNbiotChannelNumberToDlEarfcn xnapiesv1.OffsetOfNbiotChannelNumberToEarfcn) (*xnapiesv1.ServedCellInformationEUTraFDdinfoExtIesExtension, error) {

	item := &xnapiesv1.ServedCellInformationEUTraFDdinfoExtIesExtension{
		ServedCellInformationEUtraFddinfoExtIes: &xnapiesv1.ServedCellInformationEUTraFDdinfoExtIesExtension_IdOffsetOfNbiotChannelNumberToDlEarfcn{
			IdOffsetOfNbiotChannelNumberToDlEarfcn: IDOffsetOfNbiotChannelNumberToDlEarfcn,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateServedCellInformationEUTraFDdinfoExtIesExtensionIDOffsetOfNbiotChannelNumberToDlEarfcn() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateServedCellInformationEUTraFDdinfoExtIesExtensionIDOffsetOfNbiotChannelNumberToUlEarfcn(IDOffsetOfNbiotChannelNumberToUlEarfcn xnapiesv1.OffsetOfNbiotChannelNumberToEarfcn) (*xnapiesv1.ServedCellInformationEUTraFDdinfoExtIesExtension, error) {

	item := &xnapiesv1.ServedCellInformationEUTraFDdinfoExtIesExtension{
		ServedCellInformationEUtraFddinfoExtIes: &xnapiesv1.ServedCellInformationEUTraFDdinfoExtIesExtension_IdOffsetOfNbiotChannelNumberToUlEarfcn{
			IdOffsetOfNbiotChannelNumberToUlEarfcn: IDOffsetOfNbiotChannelNumberToUlEarfcn,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateServedCellInformationEUTraFDdinfoExtIesExtensionIDOffsetOfNbiotChannelNumberToUlEarfcn() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateServedCellInformationEUTraTDdinfoExtIesExtensionIDOffsetOfNbiotChannelNumberToDlEarfcn(IDOffsetOfNbiotChannelNumberToDlEarfcn xnapiesv1.OffsetOfNbiotChannelNumberToEarfcn) (*xnapiesv1.ServedCellInformationEUTraTDdinfoExtIesExtension, error) {

	item := &xnapiesv1.ServedCellInformationEUTraTDdinfoExtIesExtension{
		ServedCellInformationEUtraTddinfoExtIes: &xnapiesv1.ServedCellInformationEUTraTDdinfoExtIesExtension_IdOffsetOfNbiotChannelNumberToDlEarfcn{
			IdOffsetOfNbiotChannelNumberToDlEarfcn: IDOffsetOfNbiotChannelNumberToDlEarfcn,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateServedCellInformationEUTraTDdinfoExtIesExtensionIDOffsetOfNbiotChannelNumberToDlEarfcn() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateServedCellInformationEUTraTDdinfoExtIesExtensionIDNbioTUlDlAlignmentOffset(IDNbioTUlDlAlignmentOffset xnapiesv1.NbioTULDLAlignmentOffset) (*xnapiesv1.ServedCellInformationEUTraTDdinfoExtIesExtension, error) {

	item := &xnapiesv1.ServedCellInformationEUTraTDdinfoExtIesExtension{
		ServedCellInformationEUtraTddinfoExtIes: &xnapiesv1.ServedCellInformationEUTraTDdinfoExtIesExtension_IdNbioTUlDlAlignmentOffset{
			IdNbioTUlDlAlignmentOffset: IDNbioTUlDlAlignmentOffset,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateServedCellInformationEUTraTDdinfoExtIesExtensionIDNbioTUlDlAlignmentOffset() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateServedCellInformationNRExtIesExtensionIDBplmnIDInfoNr(IDBplmnIDInfoNr *xnapiesv1.BplmnIDInfoNR) (*xnapiesv1.ServedCellInformationNRExtIesExtension, error) {

	item := &xnapiesv1.ServedCellInformationNRExtIesExtension{
		ServedCellInformationNrExtIes: &xnapiesv1.ServedCellInformationNRExtIesExtension_IdBplmnIdInfoNr{
			IdBplmnIdInfoNr: IDBplmnIDInfoNr,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateServedCellInformationNRExtIesExtensionIDBplmnIDInfoNr() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateServedCellInformationNRExtIesExtensionIDConfiguredTacindication(IDConfiguredTacindication xnapiesv1.ConfiguredTacindication) (*xnapiesv1.ServedCellInformationNRExtIesExtension, error) {

	item := &xnapiesv1.ServedCellInformationNRExtIesExtension{
		ServedCellInformationNrExtIes: &xnapiesv1.ServedCellInformationNRExtIesExtension_IdConfiguredTacindication{
			IdConfiguredTacindication: IDConfiguredTacindication,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateServedCellInformationNRExtIesExtensionIDConfiguredTacindication() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateServedCellInformationNRExtIesExtensionIDSsbPositionsInBurst(IDSsbPositionsInBurst *xnapiesv1.SsbPositionsInBurst) (*xnapiesv1.ServedCellInformationNRExtIesExtension, error) {

	item := &xnapiesv1.ServedCellInformationNRExtIesExtension{
		ServedCellInformationNrExtIes: &xnapiesv1.ServedCellInformationNRExtIesExtension_IdSsbPositionsInBurst{
			IdSsbPositionsInBurst: IDSsbPositionsInBurst,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateServedCellInformationNRExtIesExtensionIDSsbPositionsInBurst() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateServedCellInformationNRExtIesExtensionIDNrcellPrachconfig(IDNrcellPrachconfig *xnapiesv1.NrcellPrachconfig) (*xnapiesv1.ServedCellInformationNRExtIesExtension, error) {

	item := &xnapiesv1.ServedCellInformationNRExtIesExtension{
		ServedCellInformationNrExtIes: &xnapiesv1.ServedCellInformationNRExtIesExtension_IdNrcellPrachconfig{
			IdNrcellPrachconfig: IDNrcellPrachconfig,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateServedCellInformationNRExtIesExtensionIDNrcellPrachconfig() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateServedCellInformationNRExtIesExtensionIDNpnBroadcastInformation(IDNpnBroadcastInformation *xnapiesv1.NpnBroadcastInformation) (*xnapiesv1.ServedCellInformationNRExtIesExtension, error) {

	item := &xnapiesv1.ServedCellInformationNRExtIesExtension{
		ServedCellInformationNrExtIes: &xnapiesv1.ServedCellInformationNRExtIesExtension_IdNpnBroadcastInformation{
			IdNpnBroadcastInformation: IDNpnBroadcastInformation,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateServedCellInformationNRExtIesExtensionIDNpnBroadcastInformation() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateServedCellInformationNRExtIesExtensionIDCsiRstransmissionIndication(IDCsiRstransmissionIndication xnapiesv1.CsiRStransmissionIndication) (*xnapiesv1.ServedCellInformationNRExtIesExtension, error) {

	item := &xnapiesv1.ServedCellInformationNRExtIesExtension{
		ServedCellInformationNrExtIes: &xnapiesv1.ServedCellInformationNRExtIesExtension_IdCsiRstransmissionIndication{
			IdCsiRstransmissionIndication: IDCsiRstransmissionIndication,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateServedCellInformationNRExtIesExtensionIDCsiRstransmissionIndication() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateServedCellInformationNRExtIesExtensionIDSfnOffset(IDSfnOffset *xnapiesv1.SfnOffset) (*xnapiesv1.ServedCellInformationNRExtIesExtension, error) {

	item := &xnapiesv1.ServedCellInformationNRExtIesExtension{
		ServedCellInformationNrExtIes: &xnapiesv1.ServedCellInformationNRExtIesExtension_IdSfnOffset{
			IdSfnOffset: IDSfnOffset,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateServedCellInformationNRExtIesExtensionIDSfnOffset() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSharedResourceTypeUlOnlySharing(ulOnlySharing *xnapiesv1.SharedResourceTypeULOnlySharing) (*xnapiesv1.SharedResourceType, error) {

	item := &xnapiesv1.SharedResourceType{
		SharedResourceType: &xnapiesv1.SharedResourceType_UlOnlySharing{
			UlOnlySharing: ulOnlySharing,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSharedResourceTypeUlOnlySharing() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSharedResourceTypeUlAndDlSharing(ulAndDlSharing *xnapiesv1.SharedResourceTypeULdlSharing) (*xnapiesv1.SharedResourceType, error) {

	item := &xnapiesv1.SharedResourceType{
		SharedResourceType: &xnapiesv1.SharedResourceType_UlAndDlSharing{
			UlAndDlSharing: ulAndDlSharing,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSharedResourceTypeUlAndDlSharing() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSharedResourceTypeChoiceExtension(choiceExtension *xnapiesv1.SharedResourceTypeExtIes) (*xnapiesv1.SharedResourceType, error) {

	item := &xnapiesv1.SharedResourceType{
		SharedResourceType: &xnapiesv1.SharedResourceType_ChoiceExtension{
			ChoiceExtension: choiceExtension,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSharedResourceTypeChoiceExtension() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSharedResourceTypeULdlSharingUlResources(ulResources *xnapiesv1.SharedResourceTypeULdlSharingULResources) (*xnapiesv1.SharedResourceTypeULdlSharing, error) {

	item := &xnapiesv1.SharedResourceTypeULdlSharing{
		SharedResourceTypeUldlSharing: &xnapiesv1.SharedResourceTypeULdlSharing_UlResources{
			UlResources: ulResources,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSharedResourceTypeULdlSharingUlResources() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSharedResourceTypeULdlSharingDlResources(dlResources *xnapiesv1.SharedResourceTypeULdlSharingDLResources) (*xnapiesv1.SharedResourceTypeULdlSharing, error) {

	item := &xnapiesv1.SharedResourceTypeULdlSharing{
		SharedResourceTypeUldlSharing: &xnapiesv1.SharedResourceTypeULdlSharing_DlResources{
			DlResources: dlResources,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSharedResourceTypeULdlSharingDlResources() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSharedResourceTypeULdlSharingChoiceExtension(choiceExtension *xnapiesv1.SharedResourceTypeULdlSharingExtIes) (*xnapiesv1.SharedResourceTypeULdlSharing, error) {

	item := &xnapiesv1.SharedResourceTypeULdlSharing{
		SharedResourceTypeUldlSharing: &xnapiesv1.SharedResourceTypeULdlSharing_ChoiceExtension{
			ChoiceExtension: choiceExtension,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSharedResourceTypeULdlSharingChoiceExtension() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSharedResourceTypeULdlSharingULResourcesUnchanged(unchanged int32) (*xnapiesv1.SharedResourceTypeULdlSharingULResources, error) {

	item := &xnapiesv1.SharedResourceTypeULdlSharingULResources{
		SharedResourceTypeUldlSharingUlResources: &xnapiesv1.SharedResourceTypeULdlSharingULResources_Unchanged{
			Unchanged: unchanged,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSharedResourceTypeULdlSharingULResourcesUnchanged() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSharedResourceTypeULdlSharingULResourcesChangedChoice(changed *xnapiesv1.SharedResourceTypeULdlSharingULResourcesChanged) (*xnapiesv1.SharedResourceTypeULdlSharingULResources, error) {

	item := &xnapiesv1.SharedResourceTypeULdlSharingULResources{
		SharedResourceTypeUldlSharingUlResources: &xnapiesv1.SharedResourceTypeULdlSharingULResources_Changed{
			Changed: changed,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSharedResourceTypeULdlSharingULResourcesChanged() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSharedResourceTypeULdlSharingULResourcesChoiceExtension(choiceExtension *xnapiesv1.SharedResourceTypeULdlSharingULResourcesExtIes) (*xnapiesv1.SharedResourceTypeULdlSharingULResources, error) {

	item := &xnapiesv1.SharedResourceTypeULdlSharingULResources{
		SharedResourceTypeUldlSharingUlResources: &xnapiesv1.SharedResourceTypeULdlSharingULResources_ChoiceExtension{
			ChoiceExtension: choiceExtension,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSharedResourceTypeULdlSharingULResourcesChoiceExtension() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSharedResourceTypeULdlSharingDLResourcesUnchanged(unchanged int32) (*xnapiesv1.SharedResourceTypeULdlSharingDLResources, error) {

	item := &xnapiesv1.SharedResourceTypeULdlSharingDLResources{
		SharedResourceTypeUldlSharingDlResources: &xnapiesv1.SharedResourceTypeULdlSharingDLResources_Unchanged{
			Unchanged: unchanged,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSharedResourceTypeULdlSharingDLResourcesUnchanged() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSharedResourceTypeULdlSharingDLResourcesChangedChoice(changed *xnapiesv1.SharedResourceTypeULdlSharingDLResourcesChanged) (*xnapiesv1.SharedResourceTypeULdlSharingDLResources, error) {

	item := &xnapiesv1.SharedResourceTypeULdlSharingDLResources{
		SharedResourceTypeUldlSharingDlResources: &xnapiesv1.SharedResourceTypeULdlSharingDLResources_Changed{
			Changed: changed,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSharedResourceTypeULdlSharingDLResourcesChanged() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSharedResourceTypeULdlSharingDLResourcesChoiceExtension(choiceExtension *xnapiesv1.SharedResourceTypeULdlSharingDLResourcesExtIes) (*xnapiesv1.SharedResourceTypeULdlSharingDLResources, error) {

	item := &xnapiesv1.SharedResourceTypeULdlSharingDLResources{
		SharedResourceTypeUldlSharingDlResources: &xnapiesv1.SharedResourceTypeULdlSharingDLResources_ChoiceExtension{
			ChoiceExtension: choiceExtension,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSharedResourceTypeULdlSharingDLResourcesChoiceExtension() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSsbareaRadioResourceStatusListItemExtIesExtensionIDDlSchedulingPdcchCceUsage(IDDlSchedulingPdcchCceUsage *xnapiesv1.DlschedulingPDcchCCeusage) (*xnapiesv1.SsbareaRadioResourceStatusListItemExtIesExtension, error) {

	item := &xnapiesv1.SsbareaRadioResourceStatusListItemExtIesExtension{
		SsbareaRadioResourceStatusListItemExtIes: &xnapiesv1.SsbareaRadioResourceStatusListItemExtIesExtension_IdDlSchedulingPdcchCceUsage{
			IdDlSchedulingPdcchCceUsage: IDDlSchedulingPdcchCceUsage,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSsbareaRadioResourceStatusListItemExtIesExtensionIDDlSchedulingPdcchCceUsage() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSsbareaRadioResourceStatusListItemExtIesExtensionIDUlSchedulingPdcchCceUsage(IDUlSchedulingPdcchCceUsage *xnapiesv1.UlschedulingPDcchCCeusage) (*xnapiesv1.SsbareaRadioResourceStatusListItemExtIesExtension, error) {

	item := &xnapiesv1.SsbareaRadioResourceStatusListItemExtIesExtension{
		SsbareaRadioResourceStatusListItemExtIes: &xnapiesv1.SsbareaRadioResourceStatusListItemExtIesExtension_IdUlSchedulingPdcchCceUsage{
			IdUlSchedulingPdcchCceUsage: IDUlSchedulingPdcchCceUsage,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSsbareaRadioResourceStatusListItemExtIesExtensionIDUlSchedulingPdcchCceUsage() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSsbPositionsInBurstShortBitmap(shortBitmap *asn1.BitString) (*xnapiesv1.SsbPositionsInBurst, error) {

	item := &xnapiesv1.SsbPositionsInBurst{
		SsbPositionsInBurst: &xnapiesv1.SsbPositionsInBurst_ShortBitmap{
			ShortBitmap: shortBitmap,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSsbPositionsInBurstShortBitmap() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSsbPositionsInBurstMediumBitmap(mediumBitmap *asn1.BitString) (*xnapiesv1.SsbPositionsInBurst, error) {

	item := &xnapiesv1.SsbPositionsInBurst{
		SsbPositionsInBurst: &xnapiesv1.SsbPositionsInBurst_MediumBitmap{
			MediumBitmap: mediumBitmap,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSsbPositionsInBurstMediumBitmap() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSsbPositionsInBurstLongBitmap(longBitmap *asn1.BitString) (*xnapiesv1.SsbPositionsInBurst, error) {

	item := &xnapiesv1.SsbPositionsInBurst{
		SsbPositionsInBurst: &xnapiesv1.SsbPositionsInBurst_LongBitmap{
			LongBitmap: longBitmap,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSsbPositionsInBurstLongBitmap() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSsbPositionsInBurstChoiceExtension(choiceExtension *xnapiesv1.SsbPositionsInBurstExtIes) (*xnapiesv1.SsbPositionsInBurst, error) {

	item := &xnapiesv1.SsbPositionsInBurst{
		SsbPositionsInBurst: &xnapiesv1.SsbPositionsInBurst_ChoiceExtension{
			ChoiceExtension: choiceExtension,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSsbPositionsInBurstChoiceExtension() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSulInformationExtIesExtensionIDCarrierList(IDCarrierList *xnapiesv1.NrcarrierList) (*xnapiesv1.SulInformationExtIesExtension, error) {

	item := &xnapiesv1.SulInformationExtIesExtension{
		SulInformationExtIes: &xnapiesv1.SulInformationExtIesExtension_IdCarrierList{
			IdCarrierList: IDCarrierList,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSulInformationExtIesExtensionIDCarrierList() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSulInformationExtIesExtensionIDFrequencyShift7P5Khz(IDFrequencyShift7P5Khz xnapiesv1.FrequencyShift7P5Khz) (*xnapiesv1.SulInformationExtIesExtension, error) {

	item := &xnapiesv1.SulInformationExtIesExtension{
		SulInformationExtIes: &xnapiesv1.SulInformationExtIesExtension_IdFrequencyShift7P5Khz{
			IdFrequencyShift7P5Khz: IDFrequencyShift7P5Khz,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSulInformationExtIesExtensionIDFrequencyShift7P5Khz() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSymbolAllocationinSlotAllDlChoice(allDl *xnapiesv1.SymbolAllocationinSlotAllDl) (*xnapiesv1.SymbolAllocationinSlot, error) {

	item := &xnapiesv1.SymbolAllocationinSlot{
		SymbolAllocationInSlot: &xnapiesv1.SymbolAllocationinSlot_AllDl{
			AllDl: allDl,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSymbolAllocationinSlotAllDlChoice() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSymbolAllocationinSlotAllUlChoice(allUl *xnapiesv1.SymbolAllocationinSlotAllUl) (*xnapiesv1.SymbolAllocationinSlot, error) {

	item := &xnapiesv1.SymbolAllocationinSlot{
		SymbolAllocationInSlot: &xnapiesv1.SymbolAllocationinSlot_AllUl{
			AllUl: allUl,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSymbolAllocationinSlotAllUl() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSymbolAllocationinSlotBothDlandUlChoice(bothDlandUl *xnapiesv1.SymbolAllocationinSlotBothDlandUl) (*xnapiesv1.SymbolAllocationinSlot, error) {

	item := &xnapiesv1.SymbolAllocationinSlot{
		SymbolAllocationInSlot: &xnapiesv1.SymbolAllocationinSlot_BothDlandUl{
			BothDlandUl: bothDlandUl,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSymbolAllocationinSlotBothDlandUl() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSymbolAllocationinSlotChoiceExtension(choiceExtension *xnapiesv1.SymbolAllocationinSlotExtIes) (*xnapiesv1.SymbolAllocationinSlot, error) {

	item := &xnapiesv1.SymbolAllocationinSlot{
		SymbolAllocationInSlot: &xnapiesv1.SymbolAllocationinSlot_ChoiceExtension{
			ChoiceExtension: choiceExtension,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSymbolAllocationinSlotChoiceExtension() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateTargetCGiNr(nr *xnapiesv1.NrCGi) (*xnapiesv1.TargetCGi, error) {

	item := &xnapiesv1.TargetCGi{
		TargetCgi: &xnapiesv1.TargetCGi_Nr{
			Nr: nr,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateTargetCGiNr() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateTargetCGiEUtra(eUtra *xnapiesv1.EUTraCGi) (*xnapiesv1.TargetCGi, error) {

	item := &xnapiesv1.TargetCGi{
		TargetCgi: &xnapiesv1.TargetCGi_EUtra{
			EUtra: eUtra,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateTargetCGiEUtra() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateTargetCGiChoiceExtension(choiceExtension *xnapiesv1.TargetCgiExtIes) (*xnapiesv1.TargetCGi, error) {

	item := &xnapiesv1.TargetCGi{
		TargetCgi: &xnapiesv1.TargetCGi_ChoiceExtension{
			ChoiceExtension: choiceExtension,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateTargetCGiChoiceExtension() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateTraceActivationExtIesExtensionIDTraceCollectionEntityURI(IDTraceCollectionEntityURI *xnapiesv1.Uriaddress) (*xnapiesv1.TraceActivationExtIesExtension, error) {

	item := &xnapiesv1.TraceActivationExtIesExtension{
		TraceActivationExtIes: &xnapiesv1.TraceActivationExtIesExtension_IdTraceCollectionEntityUri{
			IdTraceCollectionEntityUri: IDTraceCollectionEntityURI,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateTraceActivationExtIesExtensionIDTraceCollectionEntityURI() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateTraceActivationExtIesExtensionIDMdtConfiguration(IDMdtConfiguration *xnapiesv1.MdtConfiguration) (*xnapiesv1.TraceActivationExtIesExtension, error) {

	item := &xnapiesv1.TraceActivationExtIesExtension{
		TraceActivationExtIes: &xnapiesv1.TraceActivationExtIesExtension_IdMdtConfiguration{
			IdMdtConfiguration: IDMdtConfiguration,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateTraceActivationExtIesExtensionIDMdtConfiguration() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateUecontextIDRRcresume(rRcresume *xnapiesv1.UecontextIdforRrcresume) (*xnapiesv1.UecontextId, error) {

	item := &xnapiesv1.UecontextId{
		UecontextId: &xnapiesv1.UecontextId_RRcresume{
			RRcresume: rRcresume,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateUecontextIDRRcresume() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateUecontextIDRRrcreestablishment(rRrcreestablishment *xnapiesv1.UecontextIdforRrcreestablishment) (*xnapiesv1.UecontextId, error) {

	item := &xnapiesv1.UecontextId{
		UecontextId: &xnapiesv1.UecontextId_RRrcreestablishment{
			RRrcreestablishment: rRrcreestablishment,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateUecontextIDRRrcreestablishment() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateUecontextIDChoiceExtension(choiceExtension *xnapiesv1.UecontextIdExtIes) (*xnapiesv1.UecontextId, error) {

	item := &xnapiesv1.UecontextId{
		UecontextId: &xnapiesv1.UecontextId_ChoiceExtension{
			ChoiceExtension: choiceExtension,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateUecontextIDChoiceExtension() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateUecontextInfoRetrUectxtRespExtIesExtensionIDFiveGcmobilityRestrictionListContainer(IDFiveGcmobilityRestrictionListContainer *xnapiesv1.FiveGcmobilityRestrictionListContainer) (*xnapiesv1.UecontextInfoRetrUectxtRespExtIesExtension, error) {

	item := &xnapiesv1.UecontextInfoRetrUectxtRespExtIesExtension{
		UecontextInfoRetrUectxtRespExtIes: &xnapiesv1.UecontextInfoRetrUectxtRespExtIesExtension_IdFiveGcmobilityRestrictionListContainer{
			IdFiveGcmobilityRestrictionListContainer: IDFiveGcmobilityRestrictionListContainer,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateUecontextInfoRetrUectxtRespExtIesExtensionIDFiveGcmobilityRestrictionListContainer() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateUecontextInfoRetrUectxtRespExtIesExtensionIDNruesIDelinkAggregateMaximumBitRate(IDNruesIDelinkAggregateMaximumBitRate *xnapiesv1.NruesidelinkAggregateMaximumBitRate) (*xnapiesv1.UecontextInfoRetrUectxtRespExtIesExtension, error) {

	item := &xnapiesv1.UecontextInfoRetrUectxtRespExtIesExtension{
		UecontextInfoRetrUectxtRespExtIes: &xnapiesv1.UecontextInfoRetrUectxtRespExtIesExtension_IdNruesidelinkAggregateMaximumBitRate{
			IdNruesidelinkAggregateMaximumBitRate: IDNruesIDelinkAggregateMaximumBitRate,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateUecontextInfoRetrUectxtRespExtIesExtensionIDNruesIDelinkAggregateMaximumBitRate() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateUecontextInfoRetrUectxtRespExtIesExtensionIDLteuesIDelinkAggregateMaximumBitRate(IDLteuesIDelinkAggregateMaximumBitRate *xnapiesv1.LteuesidelinkAggregateMaximumBitRate) (*xnapiesv1.UecontextInfoRetrUectxtRespExtIesExtension, error) {

	item := &xnapiesv1.UecontextInfoRetrUectxtRespExtIesExtension{
		UecontextInfoRetrUectxtRespExtIes: &xnapiesv1.UecontextInfoRetrUectxtRespExtIesExtension_IdLteuesidelinkAggregateMaximumBitRate{
			IdLteuesidelinkAggregateMaximumBitRate: IDLteuesIDelinkAggregateMaximumBitRate,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateUecontextInfoRetrUectxtRespExtIesExtensionIDLteuesIDelinkAggregateMaximumBitRate() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateUecontextInfoRetrUectxtRespExtIesExtensionIDUeradioCapabilityID(IDUeradioCapabilityID *xnapiesv1.UeradioCapabilityId) (*xnapiesv1.UecontextInfoRetrUectxtRespExtIesExtension, error) {

	item := &xnapiesv1.UecontextInfoRetrUectxtRespExtIesExtension{
		UecontextInfoRetrUectxtRespExtIes: &xnapiesv1.UecontextInfoRetrUectxtRespExtIesExtension_IdUeradioCapabilityId{
			IdUeradioCapabilityId: IDUeradioCapabilityID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateUecontextInfoRetrUectxtRespExtIesExtensionIDUeradioCapabilityID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateUehistoryInformationFromTheUeNR(nR *xnapiesv1.NrmobilityHistoryReport) (*xnapiesv1.UehistoryInformationFromTheUe, error) {

	item := &xnapiesv1.UehistoryInformationFromTheUe{
		UehistoryInformationFromTheUe: &xnapiesv1.UehistoryInformationFromTheUe_NR{
			NR: nR,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateUehistoryInformationFromTheUeNR() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateUehistoryInformationFromTheUeChoiceExtension(choiceExtension *xnapiesv1.UehistoryInformationFromTheUeExtIes) (*xnapiesv1.UehistoryInformationFromTheUe, error) {

	item := &xnapiesv1.UehistoryInformationFromTheUe{
		UehistoryInformationFromTheUe: &xnapiesv1.UehistoryInformationFromTheUe_ChoiceExtension{
			ChoiceExtension: choiceExtension,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateUehistoryInformationFromTheUeChoiceExtension() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateUeidentityIndexValueIndexLength10(indexLength10 *asn1.BitString) (*xnapiesv1.UeidentityIndexValue, error) {

	item := &xnapiesv1.UeidentityIndexValue{
		UeidentityIndexValue: &xnapiesv1.UeidentityIndexValue_IndexLength10{
			IndexLength10: indexLength10,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateUeidentityIndexValueIndexLength10() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateUeidentityIndexValueChoiceExtension(choiceExtension *xnapiesv1.UeidentityIndexValueExtIes) (*xnapiesv1.UeidentityIndexValue, error) {

	item := &xnapiesv1.UeidentityIndexValue{
		UeidentityIndexValue: &xnapiesv1.UeidentityIndexValue_ChoiceExtension{
			ChoiceExtension: choiceExtension,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateUeidentityIndexValueChoiceExtension() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateUeranpagingIdentityIRntiFull(iRntiFull *asn1.BitString) (*xnapiesv1.UeranpagingIdentity, error) {

	item := &xnapiesv1.UeranpagingIdentity{
		UeranpagingIdentity: &xnapiesv1.UeranpagingIdentity_IRntiFull{
			IRntiFull: iRntiFull,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateUeranpagingIdentityIRntiFull() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateUeranpagingIdentityChoiceExtension(choiceExtension *xnapiesv1.UeranpagingIdentityExtIes) (*xnapiesv1.UeranpagingIdentity, error) {

	item := &xnapiesv1.UeranpagingIdentity{
		UeranpagingIdentity: &xnapiesv1.UeranpagingIdentity_ChoiceExtension{
			ChoiceExtension: choiceExtension,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateUeranpagingIdentityChoiceExtension() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateUerlfreportContainerNRUerlfreportContainer(nRUerlfreportContainer *xnapiesv1.UerlfreportContainerNr) (*xnapiesv1.UerlfreportContainer, error) {

	item := &xnapiesv1.UerlfreportContainer{
		UerlfreportContainer: &xnapiesv1.UerlfreportContainer_NRUerlfreportContainer{
			NRUerlfreportContainer: nRUerlfreportContainer,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateUerlfreportContainerNRUerlfreportContainer() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateUerlfreportContainerLTeUerlfreportContainer(lTeUerlfreportContainer *xnapiesv1.UerlfreportContainerLte) (*xnapiesv1.UerlfreportContainer, error) {

	item := &xnapiesv1.UerlfreportContainer{
		UerlfreportContainer: &xnapiesv1.UerlfreportContainer_LTeUerlfreportContainer{
			LTeUerlfreportContainer: lTeUerlfreportContainer,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateUerlfreportContainerLTeUerlfreportContainer() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateUerlfreportContainerChoiceExtension(choiceExtension *xnapiesv1.UerlfreportContainerExtIes) (*xnapiesv1.UerlfreportContainer, error) {

	item := &xnapiesv1.UerlfreportContainer{
		UerlfreportContainer: &xnapiesv1.UerlfreportContainer_ChoiceExtension{
			ChoiceExtension: choiceExtension,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateUerlfreportContainerChoiceExtension() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateUptransportLayerInformationGtpTunnel(gtpTunnel *xnapiesv1.GtptunnelTransportLayerInformation) (*xnapiesv1.UptransportLayerInformation, error) {

	item := &xnapiesv1.UptransportLayerInformation{
		UptransportLayerInformation: &xnapiesv1.UptransportLayerInformation_GtpTunnel{
			GtpTunnel: gtpTunnel,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateUptransportLayerInformationGtpTunnel() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateUptransportLayerInformationChoiceExtension(choiceExtension *xnapiesv1.UptransportLayerInformationExtIes) (*xnapiesv1.UptransportLayerInformation, error) {

	item := &xnapiesv1.UptransportLayerInformation{
		UptransportLayerInformation: &xnapiesv1.UptransportLayerInformation_ChoiceExtension{
			ChoiceExtension: choiceExtension,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateUptransportLayerInformationChoiceExtension() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreatePrivateIeIDLocal(local int32) (*xnapcommondatatypesv1.PrivateIeID, error) {

	item := &xnapcommondatatypesv1.PrivateIeID{
		PrivateIeId: &xnapcommondatatypesv1.PrivateIeID_Local{
			Local: local,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePrivateIeIDLocal() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreatePrivateIeIDGlobal(global string) (*xnapcommondatatypesv1.PrivateIeID, error) {

	item := &xnapcommondatatypesv1.PrivateIeID{
		PrivateIeId: &xnapcommondatatypesv1.PrivateIeID_Global{
			Global: global,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePrivateIeIDGlobal() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateXnApPDuInitiatingMessage(initiatingMessage *xnappdudescriptionsv1.InitiatingMessage) (*xnappdudescriptionsv1.XnApPDu, error) {

	item := &xnappdudescriptionsv1.XnApPDu{
		XnApPdu: &xnappdudescriptionsv1.XnApPDu_InitiatingMessage{
			InitiatingMessage: initiatingMessage,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateXnApPDuInitiatingMessage() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateXnApPDuSuccessfulOutcome(successfulOutcome *xnappdudescriptionsv1.SuccessfulOutcome) (*xnappdudescriptionsv1.XnApPDu, error) {

	item := &xnappdudescriptionsv1.XnApPDu{
		XnApPdu: &xnappdudescriptionsv1.XnApPDu_SuccessfulOutcome{
			SuccessfulOutcome: successfulOutcome,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateXnApPDuSuccessfulOutcome() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateXnApPDuUnsuccessfulOutcome(unsuccessfulOutcome *xnappdudescriptionsv1.UnsuccessfulOutcome) (*xnappdudescriptionsv1.XnApPDu, error) {

	item := &xnappdudescriptionsv1.XnApPDu{
		XnApPdu: &xnappdudescriptionsv1.XnApPDu_UnsuccessfulOutcome{
			UnsuccessfulOutcome: unsuccessfulOutcome,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateXnApPDuUnsuccessfulOutcome() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateInitiatingMessageXnApElementaryProceduresXnSetupRequest(xnSetupRequest *xnappducontentsv1.XnSetupRequest) (*xnappdudescriptionsv1.InitiatingMessageXnApElementaryProcedures, error) {

	item := &xnappdudescriptionsv1.InitiatingMessageXnApElementaryProcedures{
		ImValues: &xnappdudescriptionsv1.InitiatingMessageXnApElementaryProcedures_XnSetupRequest{
			XnSetupRequest: xnSetupRequest,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateInitiatingMessageXnApElementaryProceduresXnSetupRequest() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSuccessfulOutcomeXnApElementaryProceduresXnSetupResponse(xnSetupResponse *xnappducontentsv1.XnSetupResponse) (*xnappdudescriptionsv1.SuccessfulOutcomeXnApElementaryProcedures, error) {

	item := &xnappdudescriptionsv1.SuccessfulOutcomeXnApElementaryProcedures{
		SoValues: &xnappdudescriptionsv1.SuccessfulOutcomeXnApElementaryProcedures_XnSetupResponse{
			XnSetupResponse: xnSetupResponse,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSuccessfulOutcomeXnApElementaryProceduresXnSetupResponse() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateUnsuccessfulOutcomeXnApElementaryProceduresXnSetupFailure(xnSetupFailure *xnappducontentsv1.XnSetupFailure) (*xnappdudescriptionsv1.UnsuccessfulOutcomeXnApElementaryProcedures, error) {

	item := &xnappdudescriptionsv1.UnsuccessfulOutcomeXnApElementaryProcedures{
		UoValues: &xnappdudescriptionsv1.UnsuccessfulOutcomeXnApElementaryProcedures_XnSetupFailure{
			XnSetupFailure: xnSetupFailure,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateUnsuccessfulOutcomeXnApElementaryProceduresXnSetupFailure() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateHandoverRequestIEsValueIDSourceNgRannodeUexnApID(IDSourceNgRannodeUexnApID *xnapiesv1.NgRAnnodeUexnApid) (*xnappducontentsv1.HandoverRequestIEsValue, error) {

	item := &xnappducontentsv1.HandoverRequestIEsValue{
		HandoverRequestIes: &xnappducontentsv1.HandoverRequestIEsValue_IdSourceNgRannodeUexnApid{
			IdSourceNgRannodeUexnApid: IDSourceNgRannodeUexnApID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateHandoverRequestIEsValueIDSourceNgRannodeUexnApID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateHandoverRequestIEsValueIDCause(IDCause *xnapiesv1.Cause) (*xnappducontentsv1.HandoverRequestIEsValue, error) {

	item := &xnappducontentsv1.HandoverRequestIEsValue{
		HandoverRequestIes: &xnappducontentsv1.HandoverRequestIEsValue_IdCause{
			IdCause: IDCause,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateHandoverRequestIEsValueIDCause() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateHandoverRequestIEsValueIDTargetCellGlobalID(IDTargetCellGlobalID *xnapiesv1.TargetCGi) (*xnappducontentsv1.HandoverRequestIEsValue, error) {

	item := &xnappducontentsv1.HandoverRequestIEsValue{
		HandoverRequestIes: &xnappducontentsv1.HandoverRequestIEsValue_IdTargetCellGlobalId{
			IdTargetCellGlobalId: IDTargetCellGlobalID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateHandoverRequestIEsValueIDTargetCellGlobalID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateHandoverRequestIEsValueIDGuami(IDGuami *xnapiesv1.Guami) (*xnappducontentsv1.HandoverRequestIEsValue, error) {

	item := &xnappducontentsv1.HandoverRequestIEsValue{
		HandoverRequestIes: &xnappducontentsv1.HandoverRequestIEsValue_IdGuami{
			IdGuami: IDGuami,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateHandoverRequestIEsValueIDGuami() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateHandoverRequestIEsValueIDUecontextInfoHorequest(IDUecontextInfoHorequest *xnappducontentsv1.UecontextInfoHorequest) (*xnappducontentsv1.HandoverRequestIEsValue, error) {

	item := &xnappducontentsv1.HandoverRequestIEsValue{
		HandoverRequestIes: &xnappducontentsv1.HandoverRequestIEsValue_IdUecontextInfoHorequest{
			IdUecontextInfoHorequest: IDUecontextInfoHorequest,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateHandoverRequestIEsValueIDUecontextInfoHorequest() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateHandoverRequestIEsValueIDTraceActivation(IDTraceActivation *xnapiesv1.TraceActivation) (*xnappducontentsv1.HandoverRequestIEsValue, error) {

	item := &xnappducontentsv1.HandoverRequestIEsValue{
		HandoverRequestIes: &xnappducontentsv1.HandoverRequestIEsValue_IdTraceActivation{
			IdTraceActivation: IDTraceActivation,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateHandoverRequestIEsValueIDTraceActivation() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateHandoverRequestIEsValueIDMaskedImeisv(IDMaskedImeisv *xnapiesv1.MaskedImeisv) (*xnappducontentsv1.HandoverRequestIEsValue, error) {

	item := &xnappducontentsv1.HandoverRequestIEsValue{
		HandoverRequestIes: &xnappducontentsv1.HandoverRequestIEsValue_IdMaskedImeisv{
			IdMaskedImeisv: IDMaskedImeisv,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateHandoverRequestIEsValueIDMaskedImeisv() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateHandoverRequestIEsValueIDUehistoryInformation(IDUehistoryInformation *xnapiesv1.UehistoryInformation) (*xnappducontentsv1.HandoverRequestIEsValue, error) {

	item := &xnappducontentsv1.HandoverRequestIEsValue{
		HandoverRequestIes: &xnappducontentsv1.HandoverRequestIEsValue_IdUehistoryInformation{
			IdUehistoryInformation: IDUehistoryInformation,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateHandoverRequestIEsValueIDUehistoryInformation() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateHandoverRequestIEsValueIDUecontextRefAtSnHorequest(IDUecontextRefAtSnHorequest *xnappducontentsv1.UecontextRefAtSnHOrequest) (*xnappducontentsv1.HandoverRequestIEsValue, error) {

	item := &xnappducontentsv1.HandoverRequestIEsValue{
		HandoverRequestIes: &xnappducontentsv1.HandoverRequestIEsValue_IdUecontextRefAtSnHorequest{
			IdUecontextRefAtSnHorequest: IDUecontextRefAtSnHorequest,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateHandoverRequestIEsValueIDUecontextRefAtSnHorequest() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateHandoverRequestIEsValueIDChoinformationReq(IDChoinformationReq *xnapiesv1.ChoinformationReq) (*xnappducontentsv1.HandoverRequestIEsValue, error) {

	item := &xnappducontentsv1.HandoverRequestIEsValue{
		HandoverRequestIes: &xnappducontentsv1.HandoverRequestIEsValue_IdChoinformationReq{
			IdChoinformationReq: IDChoinformationReq,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateHandoverRequestIEsValueIDChoinformationReq() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateHandoverRequestIEsValueIDNrv2XservicesAuthorized(IDNrv2XservicesAuthorized *xnapiesv1.Nrv2XservicesAuthorized) (*xnappducontentsv1.HandoverRequestIEsValue, error) {

	item := &xnappducontentsv1.HandoverRequestIEsValue{
		HandoverRequestIes: &xnappducontentsv1.HandoverRequestIEsValue_IdNrv2XservicesAuthorized{
			IdNrv2XservicesAuthorized: IDNrv2XservicesAuthorized,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateHandoverRequestIEsValueIDNrv2XservicesAuthorized() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateHandoverRequestIEsValueIDLtev2XservicesAuthorized(IDLtev2XservicesAuthorized *xnapiesv1.Ltev2XservicesAuthorized) (*xnappducontentsv1.HandoverRequestIEsValue, error) {

	item := &xnappducontentsv1.HandoverRequestIEsValue{
		HandoverRequestIes: &xnappducontentsv1.HandoverRequestIEsValue_IdLtev2XservicesAuthorized{
			IdLtev2XservicesAuthorized: IDLtev2XservicesAuthorized,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateHandoverRequestIEsValueIDLtev2XservicesAuthorized() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateHandoverRequestIEsValueIDPc5QoSparameters(IDPc5QoSparameters *xnapiesv1.Pc5QoSparameters) (*xnappducontentsv1.HandoverRequestIEsValue, error) {

	item := &xnappducontentsv1.HandoverRequestIEsValue{
		HandoverRequestIes: &xnappducontentsv1.HandoverRequestIEsValue_IdPc5QoSparameters{
			IdPc5QoSparameters: IDPc5QoSparameters,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateHandoverRequestIEsValueIDPc5QoSparameters() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateHandoverRequestIEsValueIDMobilityInformation(IDMobilityInformation *xnapiesv1.MobilityInformation) (*xnappducontentsv1.HandoverRequestIEsValue, error) {

	item := &xnappducontentsv1.HandoverRequestIEsValue{
		HandoverRequestIes: &xnappducontentsv1.HandoverRequestIEsValue_IdMobilityInformation{
			IdMobilityInformation: IDMobilityInformation,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateHandoverRequestIEsValueIDMobilityInformation() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateHandoverRequestIEsValueIDUehistoryInformationFromTheUe(IDUehistoryInformationFromTheUe *xnapiesv1.UehistoryInformationFromTheUe) (*xnappducontentsv1.HandoverRequestIEsValue, error) {

	item := &xnappducontentsv1.HandoverRequestIEsValue{
		HandoverRequestIes: &xnappducontentsv1.HandoverRequestIEsValue_IdUehistoryInformationFromTheUe{
			IdUehistoryInformationFromTheUe: IDUehistoryInformationFromTheUe,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateHandoverRequestIEsValueIDUehistoryInformationFromTheUe() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateHandoverRequestIEsValueIDIabnodeIndication(IDIabnodeIndication xnapiesv1.IabnodeIndication) (*xnappducontentsv1.HandoverRequestIEsValue, error) {

	item := &xnappducontentsv1.HandoverRequestIEsValue{
		HandoverRequestIes: &xnappducontentsv1.HandoverRequestIEsValue_IdIabnodeIndication{
			IdIabnodeIndication: IDIabnodeIndication,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateHandoverRequestIEsValueIDIabnodeIndication() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateUecontextInfoHorequestExtIesExtensionIDFiveGcmobilityRestrictionListContainer(IDFiveGcmobilityRestrictionListContainer *xnapiesv1.FiveGcmobilityRestrictionListContainer) (*xnappducontentsv1.UecontextInfoHorequestExtIesExtension, error) {

	item := &xnappducontentsv1.UecontextInfoHorequestExtIesExtension{
		UecontextInfoHorequestExtIes: &xnappducontentsv1.UecontextInfoHorequestExtIesExtension_IdFiveGcmobilityRestrictionListContainer{
			IdFiveGcmobilityRestrictionListContainer: IDFiveGcmobilityRestrictionListContainer,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateUecontextInfoHorequestExtIesExtensionIDFiveGcmobilityRestrictionListContainer() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateUecontextInfoHorequestExtIesExtensionIDNruesIDelinkAggregateMaximumBitRate(IDNruesIDelinkAggregateMaximumBitRate *xnapiesv1.NruesidelinkAggregateMaximumBitRate) (*xnappducontentsv1.UecontextInfoHorequestExtIesExtension, error) {

	item := &xnappducontentsv1.UecontextInfoHorequestExtIesExtension{
		UecontextInfoHorequestExtIes: &xnappducontentsv1.UecontextInfoHorequestExtIesExtension_IdNruesidelinkAggregateMaximumBitRate{
			IdNruesidelinkAggregateMaximumBitRate: IDNruesIDelinkAggregateMaximumBitRate,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateUecontextInfoHorequestExtIesExtensionIDNruesIDelinkAggregateMaximumBitRate() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateUecontextInfoHorequestExtIesExtensionIDLteuesIDelinkAggregateMaximumBitRate(IDLteuesIDelinkAggregateMaximumBitRate *xnapiesv1.LteuesidelinkAggregateMaximumBitRate) (*xnappducontentsv1.UecontextInfoHorequestExtIesExtension, error) {

	item := &xnappducontentsv1.UecontextInfoHorequestExtIesExtension{
		UecontextInfoHorequestExtIes: &xnappducontentsv1.UecontextInfoHorequestExtIesExtension_IdLteuesidelinkAggregateMaximumBitRate{
			IdLteuesidelinkAggregateMaximumBitRate: IDLteuesIDelinkAggregateMaximumBitRate,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateUecontextInfoHorequestExtIesExtensionIDLteuesIDelinkAggregateMaximumBitRate() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateUecontextInfoHorequestExtIesExtensionIDMdtplmnlist(IDMdtplmnlist *xnapiesv1.Mdtplmnlist) (*xnappducontentsv1.UecontextInfoHorequestExtIesExtension, error) {

	item := &xnappducontentsv1.UecontextInfoHorequestExtIesExtension{
		UecontextInfoHorequestExtIes: &xnappducontentsv1.UecontextInfoHorequestExtIesExtension_IdMdtplmnlist{
			IdMdtplmnlist: IDMdtplmnlist,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateUecontextInfoHorequestExtIesExtensionIDMdtplmnlist() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateUecontextInfoHorequestExtIesExtensionIDUeradioCapabilityID(IDUeradioCapabilityID *xnapiesv1.UeradioCapabilityId) (*xnappducontentsv1.UecontextInfoHorequestExtIesExtension, error) {

	item := &xnappducontentsv1.UecontextInfoHorequestExtIesExtension{
		UecontextInfoHorequestExtIes: &xnappducontentsv1.UecontextInfoHorequestExtIesExtension_IdUeradioCapabilityId{
			IdUeradioCapabilityId: IDUeradioCapabilityID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateUecontextInfoHorequestExtIesExtensionIDUeradioCapabilityID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateHandoverRequestAcknowledgeIEsValueIDSourceNgRannodeUexnApID(IDSourceNgRannodeUexnApID *xnapiesv1.NgRAnnodeUexnApid) (*xnappducontentsv1.HandoverRequestAcknowledgeIEsValue, error) {

	item := &xnappducontentsv1.HandoverRequestAcknowledgeIEsValue{
		HandoverRequestAcknowledgeIes: &xnappducontentsv1.HandoverRequestAcknowledgeIEsValue_IdSourceNgRannodeUexnApid{
			IdSourceNgRannodeUexnApid: IDSourceNgRannodeUexnApID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateHandoverRequestAcknowledgeIEsValueIDSourceNgRannodeUexnApID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateHandoverRequestAcknowledgeIEsValueIDTargetNgRannodeUexnApID(IDTargetNgRannodeUexnApID *xnapiesv1.NgRAnnodeUexnApid) (*xnappducontentsv1.HandoverRequestAcknowledgeIEsValue, error) {

	item := &xnappducontentsv1.HandoverRequestAcknowledgeIEsValue{
		HandoverRequestAcknowledgeIes: &xnappducontentsv1.HandoverRequestAcknowledgeIEsValue_IdTargetNgRannodeUexnApid{
			IdTargetNgRannodeUexnApid: IDTargetNgRannodeUexnApID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateHandoverRequestAcknowledgeIEsValueIDTargetNgRannodeUexnApID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateHandoverRequestAcknowledgeIEsValueIDPdusessionResourcesAdmittedList(IDPdusessionResourcesAdmittedList *xnapiesv1.PdusessionResourcesAdmittedList) (*xnappducontentsv1.HandoverRequestAcknowledgeIEsValue, error) {

	item := &xnappducontentsv1.HandoverRequestAcknowledgeIEsValue{
		HandoverRequestAcknowledgeIes: &xnappducontentsv1.HandoverRequestAcknowledgeIEsValue_IdPdusessionResourcesAdmittedList{
			IdPdusessionResourcesAdmittedList: IDPdusessionResourcesAdmittedList,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateHandoverRequestAcknowledgeIEsValueIDPdusessionResourcesAdmittedList() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateHandoverRequestAcknowledgeIEsValueIDPdusessionResourcesNotAdmittedList(IDPdusessionResourcesNotAdmittedList *xnapiesv1.PdusessionResourcesNotAdmittedList) (*xnappducontentsv1.HandoverRequestAcknowledgeIEsValue, error) {

	item := &xnappducontentsv1.HandoverRequestAcknowledgeIEsValue{
		HandoverRequestAcknowledgeIes: &xnappducontentsv1.HandoverRequestAcknowledgeIEsValue_IdPdusessionResourcesNotAdmittedList{
			IdPdusessionResourcesNotAdmittedList: IDPdusessionResourcesNotAdmittedList,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateHandoverRequestAcknowledgeIEsValueIDPdusessionResourcesNotAdmittedList() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateHandoverRequestAcknowledgeIEsValueIDTarget2SourceNgRannodeTranspContainer(IDTarget2SourceNgRannodeTranspContainer string) (*xnappducontentsv1.HandoverRequestAcknowledgeIEsValue, error) {

	item := &xnappducontentsv1.HandoverRequestAcknowledgeIEsValue{
		HandoverRequestAcknowledgeIes: &xnappducontentsv1.HandoverRequestAcknowledgeIEsValue_IdTarget2SourceNgRannodeTranspContainer{
			IdTarget2SourceNgRannodeTranspContainer: IDTarget2SourceNgRannodeTranspContainer,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateHandoverRequestAcknowledgeIEsValueIDTarget2SourceNgRannodeTranspContainer() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateHandoverRequestAcknowledgeIEsValueIDUecontextKeptIndicator(IDUecontextKeptIndicator xnapiesv1.UecontextKeptIndicator) (*xnappducontentsv1.HandoverRequestAcknowledgeIEsValue, error) {

	item := &xnappducontentsv1.HandoverRequestAcknowledgeIEsValue{
		HandoverRequestAcknowledgeIes: &xnappducontentsv1.HandoverRequestAcknowledgeIEsValue_IdUecontextKeptIndicator{
			IdUecontextKeptIndicator: IDUecontextKeptIndicator,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateHandoverRequestAcknowledgeIEsValueIDUecontextKeptIndicator() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateHandoverRequestAcknowledgeIEsValueIDCriticalityDiagnostics(IDCriticalityDiagnostics *xnapiesv1.CriticalityDiagnostics) (*xnappducontentsv1.HandoverRequestAcknowledgeIEsValue, error) {

	item := &xnappducontentsv1.HandoverRequestAcknowledgeIEsValue{
		HandoverRequestAcknowledgeIes: &xnappducontentsv1.HandoverRequestAcknowledgeIEsValue_IdCriticalityDiagnostics{
			IdCriticalityDiagnostics: IDCriticalityDiagnostics,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateHandoverRequestAcknowledgeIEsValueIDCriticalityDiagnostics() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateHandoverRequestAcknowledgeIEsValueIDDrbsTransferredToMn(IDDrbsTransferredToMn *xnapiesv1.DrbList) (*xnappducontentsv1.HandoverRequestAcknowledgeIEsValue, error) {

	item := &xnappducontentsv1.HandoverRequestAcknowledgeIEsValue{
		HandoverRequestAcknowledgeIes: &xnappducontentsv1.HandoverRequestAcknowledgeIEsValue_IdDrbsTransferredToMn{
			IdDrbsTransferredToMn: IDDrbsTransferredToMn,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateHandoverRequestAcknowledgeIEsValueIDDrbsTransferredToMn() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateHandoverRequestAcknowledgeIEsValueIDDapsresponseInfoList(IDDapsresponseInfoList *xnapiesv1.DapsresponseInfoList) (*xnappducontentsv1.HandoverRequestAcknowledgeIEsValue, error) {

	item := &xnappducontentsv1.HandoverRequestAcknowledgeIEsValue{
		HandoverRequestAcknowledgeIes: &xnappducontentsv1.HandoverRequestAcknowledgeIEsValue_IdDapsresponseInfoList{
			IdDapsresponseInfoList: IDDapsresponseInfoList,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateHandoverRequestAcknowledgeIEsValueIDDapsresponseInfoList() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateHandoverRequestAcknowledgeIEsValueIDChoinformationAck(IDChoinformationAck *xnapiesv1.ChoinformationAck) (*xnappducontentsv1.HandoverRequestAcknowledgeIEsValue, error) {

	item := &xnappducontentsv1.HandoverRequestAcknowledgeIEsValue{
		HandoverRequestAcknowledgeIes: &xnappducontentsv1.HandoverRequestAcknowledgeIEsValue_IdChoinformationAck{
			IdChoinformationAck: IDChoinformationAck,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateHandoverRequestAcknowledgeIEsValueIDChoinformationAck() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateHandoverPreparationFailureIEsValueIDSourceNgRannodeUexnApID(IDSourceNgRannodeUexnApID *xnapiesv1.NgRAnnodeUexnApid) (*xnappducontentsv1.HandoverPreparationFailureIEsValue, error) {

	item := &xnappducontentsv1.HandoverPreparationFailureIEsValue{
		HandoverPreparationFailureIes: &xnappducontentsv1.HandoverPreparationFailureIEsValue_IdSourceNgRannodeUexnApid{
			IdSourceNgRannodeUexnApid: IDSourceNgRannodeUexnApID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateHandoverPreparationFailureIEsValueIDSourceNgRannodeUexnApID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateHandoverPreparationFailureIEsValueIDCause(IDCause *xnapiesv1.Cause) (*xnappducontentsv1.HandoverPreparationFailureIEsValue, error) {

	item := &xnappducontentsv1.HandoverPreparationFailureIEsValue{
		HandoverPreparationFailureIes: &xnappducontentsv1.HandoverPreparationFailureIEsValue_IdCause{
			IdCause: IDCause,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateHandoverPreparationFailureIEsValueIDCause() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateHandoverPreparationFailureIEsValueIDCriticalityDiagnostics(IDCriticalityDiagnostics *xnapiesv1.CriticalityDiagnostics) (*xnappducontentsv1.HandoverPreparationFailureIEsValue, error) {

	item := &xnappducontentsv1.HandoverPreparationFailureIEsValue{
		HandoverPreparationFailureIes: &xnappducontentsv1.HandoverPreparationFailureIEsValue_IdCriticalityDiagnostics{
			IdCriticalityDiagnostics: IDCriticalityDiagnostics,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateHandoverPreparationFailureIEsValueIDCriticalityDiagnostics() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateHandoverPreparationFailureIEsValueIDRequestedTargetCellGlobalID(IDRequestedTargetCellGlobalID *xnapiesv1.TargetCGi) (*xnappducontentsv1.HandoverPreparationFailureIEsValue, error) {

	item := &xnappducontentsv1.HandoverPreparationFailureIEsValue{
		HandoverPreparationFailureIes: &xnappducontentsv1.HandoverPreparationFailureIEsValue_IdRequestedTargetCellGlobalId{
			IdRequestedTargetCellGlobalId: IDRequestedTargetCellGlobalID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateHandoverPreparationFailureIEsValueIDRequestedTargetCellGlobalID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnstatusTransferIEsValueIDSourceNgRannodeUexnApID(IDSourceNgRannodeUexnApID *xnapiesv1.NgRAnnodeUexnApid) (*xnappducontentsv1.SnstatusTransferIEsValue, error) {

	item := &xnappducontentsv1.SnstatusTransferIEsValue{
		SnstatusTransferIes: &xnappducontentsv1.SnstatusTransferIEsValue_IdSourceNgRannodeUexnApid{
			IdSourceNgRannodeUexnApid: IDSourceNgRannodeUexnApID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnstatusTransferIEsValueIDSourceNgRannodeUexnApID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnstatusTransferIEsValueIDTargetNgRannodeUexnApID(IDTargetNgRannodeUexnApID *xnapiesv1.NgRAnnodeUexnApid) (*xnappducontentsv1.SnstatusTransferIEsValue, error) {

	item := &xnappducontentsv1.SnstatusTransferIEsValue{
		SnstatusTransferIes: &xnappducontentsv1.SnstatusTransferIEsValue_IdTargetNgRannodeUexnApid{
			IdTargetNgRannodeUexnApid: IDTargetNgRannodeUexnApID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnstatusTransferIEsValueIDTargetNgRannodeUexnApID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnstatusTransferIEsValueIDDrbsSubjectToStatusTransferList(IDDrbsSubjectToStatusTransferList *xnapiesv1.DrbsSubjectToStatusTransferList) (*xnappducontentsv1.SnstatusTransferIEsValue, error) {

	item := &xnappducontentsv1.SnstatusTransferIEsValue{
		SnstatusTransferIes: &xnappducontentsv1.SnstatusTransferIEsValue_IdDrbsSubjectToStatusTransferList{
			IdDrbsSubjectToStatusTransferList: IDDrbsSubjectToStatusTransferList,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnstatusTransferIEsValueIDDrbsSubjectToStatusTransferList() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateUecontextReleaseIEsValueIDSourceNgRannodeUexnApID(IDSourceNgRannodeUexnApID *xnapiesv1.NgRAnnodeUexnApid) (*xnappducontentsv1.UecontextReleaseIEsValue, error) {

	item := &xnappducontentsv1.UecontextReleaseIEsValue{
		UecontextReleaseIes: &xnappducontentsv1.UecontextReleaseIEsValue_IdSourceNgRannodeUexnApid{
			IdSourceNgRannodeUexnApid: IDSourceNgRannodeUexnApID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateUecontextReleaseIEsValueIDSourceNgRannodeUexnApID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateUecontextReleaseIEsValueIDTargetNgRannodeUexnApID(IDTargetNgRannodeUexnApID *xnapiesv1.NgRAnnodeUexnApid) (*xnappducontentsv1.UecontextReleaseIEsValue, error) {

	item := &xnappducontentsv1.UecontextReleaseIEsValue{
		UecontextReleaseIes: &xnappducontentsv1.UecontextReleaseIEsValue_IdTargetNgRannodeUexnApid{
			IdTargetNgRannodeUexnApid: IDTargetNgRannodeUexnApID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateUecontextReleaseIEsValueIDTargetNgRannodeUexnApID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateHandoverCancelIEsValueIDSourceNgRannodeUexnApID(IDSourceNgRannodeUexnApID *xnapiesv1.NgRAnnodeUexnApid) (*xnappducontentsv1.HandoverCancelIEsValue, error) {

	item := &xnappducontentsv1.HandoverCancelIEsValue{
		HandoverCancelIes: &xnappducontentsv1.HandoverCancelIEsValue_IdSourceNgRannodeUexnApid{
			IdSourceNgRannodeUexnApid: IDSourceNgRannodeUexnApID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateHandoverCancelIEsValueIDSourceNgRannodeUexnApID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateHandoverCancelIEsValueIDTargetNgRannodeUexnApID(IDTargetNgRannodeUexnApID *xnapiesv1.NgRAnnodeUexnApid) (*xnappducontentsv1.HandoverCancelIEsValue, error) {

	item := &xnappducontentsv1.HandoverCancelIEsValue{
		HandoverCancelIes: &xnappducontentsv1.HandoverCancelIEsValue_IdTargetNgRannodeUexnApid{
			IdTargetNgRannodeUexnApid: IDTargetNgRannodeUexnApID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateHandoverCancelIEsValueIDTargetNgRannodeUexnApID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateHandoverCancelIEsValueIDCause(IDCause *xnapiesv1.Cause) (*xnappducontentsv1.HandoverCancelIEsValue, error) {

	item := &xnappducontentsv1.HandoverCancelIEsValue{
		HandoverCancelIes: &xnappducontentsv1.HandoverCancelIEsValue_IdCause{
			IdCause: IDCause,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateHandoverCancelIEsValueIDCause() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateHandoverCancelIEsValueIDTargetCellsToCancel(IDTargetCellsToCancel *xnapiesv1.TargetCellList) (*xnappducontentsv1.HandoverCancelIEsValue, error) {

	item := &xnappducontentsv1.HandoverCancelIEsValue{
		HandoverCancelIes: &xnappducontentsv1.HandoverCancelIEsValue_IdTargetCellsToCancel{
			IdTargetCellsToCancel: IDTargetCellsToCancel,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateHandoverCancelIEsValueIDTargetCellsToCancel() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateHandoverSuccessIEsValueIDSourceNgRannodeUexnApID(IDSourceNgRannodeUexnApID *xnapiesv1.NgRAnnodeUexnApid) (*xnappducontentsv1.HandoverSuccessIEsValue, error) {

	item := &xnappducontentsv1.HandoverSuccessIEsValue{
		HandoverSuccessIes: &xnappducontentsv1.HandoverSuccessIEsValue_IdSourceNgRannodeUexnApid{
			IdSourceNgRannodeUexnApid: IDSourceNgRannodeUexnApID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateHandoverSuccessIEsValueIDSourceNgRannodeUexnApID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateHandoverSuccessIEsValueIDTargetNgRannodeUexnApID(IDTargetNgRannodeUexnApID *xnapiesv1.NgRAnnodeUexnApid) (*xnappducontentsv1.HandoverSuccessIEsValue, error) {

	item := &xnappducontentsv1.HandoverSuccessIEsValue{
		HandoverSuccessIes: &xnappducontentsv1.HandoverSuccessIEsValue_IdTargetNgRannodeUexnApid{
			IdTargetNgRannodeUexnApid: IDTargetNgRannodeUexnApID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateHandoverSuccessIEsValueIDTargetNgRannodeUexnApID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateHandoverSuccessIEsValueIDRequestedTargetCellGlobalID(IDRequestedTargetCellGlobalID *xnapiesv1.TargetCGi) (*xnappducontentsv1.HandoverSuccessIEsValue, error) {

	item := &xnappducontentsv1.HandoverSuccessIEsValue{
		HandoverSuccessIes: &xnappducontentsv1.HandoverSuccessIEsValue_IdRequestedTargetCellGlobalId{
			IdRequestedTargetCellGlobalId: IDRequestedTargetCellGlobalID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateHandoverSuccessIEsValueIDRequestedTargetCellGlobalID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateConditionalHandoverCancelIEsValueIDSourceNgRannodeUexnApID(IDSourceNgRannodeUexnApID *xnapiesv1.NgRAnnodeUexnApid) (*xnappducontentsv1.ConditionalHandoverCancelIEsValue, error) {

	item := &xnappducontentsv1.ConditionalHandoverCancelIEsValue{
		ConditionalHandoverCancelIes: &xnappducontentsv1.ConditionalHandoverCancelIEsValue_IdSourceNgRannodeUexnApid{
			IdSourceNgRannodeUexnApid: IDSourceNgRannodeUexnApID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateConditionalHandoverCancelIEsValueIDSourceNgRannodeUexnApID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateConditionalHandoverCancelIEsValueIDTargetNgRannodeUexnApID(IDTargetNgRannodeUexnApID *xnapiesv1.NgRAnnodeUexnApid) (*xnappducontentsv1.ConditionalHandoverCancelIEsValue, error) {

	item := &xnappducontentsv1.ConditionalHandoverCancelIEsValue{
		ConditionalHandoverCancelIes: &xnappducontentsv1.ConditionalHandoverCancelIEsValue_IdTargetNgRannodeUexnApid{
			IdTargetNgRannodeUexnApid: IDTargetNgRannodeUexnApID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateConditionalHandoverCancelIEsValueIDTargetNgRannodeUexnApID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateConditionalHandoverCancelIEsValueIDCause(IDCause *xnapiesv1.Cause) (*xnappducontentsv1.ConditionalHandoverCancelIEsValue, error) {

	item := &xnappducontentsv1.ConditionalHandoverCancelIEsValue{
		ConditionalHandoverCancelIes: &xnappducontentsv1.ConditionalHandoverCancelIEsValue_IdCause{
			IdCause: IDCause,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateConditionalHandoverCancelIEsValueIDCause() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateConditionalHandoverCancelIEsValueIDTargetCellsToCancel(IDTargetCellsToCancel *xnapiesv1.TargetCellList) (*xnappducontentsv1.ConditionalHandoverCancelIEsValue, error) {

	item := &xnappducontentsv1.ConditionalHandoverCancelIEsValue{
		ConditionalHandoverCancelIes: &xnappducontentsv1.ConditionalHandoverCancelIEsValue_IdTargetCellsToCancel{
			IdTargetCellsToCancel: IDTargetCellsToCancel,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateConditionalHandoverCancelIEsValueIDTargetCellsToCancel() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateEarlyStatusTransferIEsValueIDSourceNgRannodeUexnApID(IDSourceNgRannodeUexnApID *xnapiesv1.NgRAnnodeUexnApid) (*xnappducontentsv1.EarlyStatusTransferIEsValue, error) {

	item := &xnappducontentsv1.EarlyStatusTransferIEsValue{
		EarlyStatusTransferIes: &xnappducontentsv1.EarlyStatusTransferIEsValue_IdSourceNgRannodeUexnApid{
			IdSourceNgRannodeUexnApid: IDSourceNgRannodeUexnApID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateEarlyStatusTransferIEsValueIDSourceNgRannodeUexnApID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateEarlyStatusTransferIEsValueIDTargetNgRannodeUexnApID(IDTargetNgRannodeUexnApID *xnapiesv1.NgRAnnodeUexnApid) (*xnappducontentsv1.EarlyStatusTransferIEsValue, error) {

	item := &xnappducontentsv1.EarlyStatusTransferIEsValue{
		EarlyStatusTransferIes: &xnappducontentsv1.EarlyStatusTransferIEsValue_IdTargetNgRannodeUexnApid{
			IdTargetNgRannodeUexnApid: IDTargetNgRannodeUexnApID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateEarlyStatusTransferIEsValueIDTargetNgRannodeUexnApID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateEarlyStatusTransferIEsValueIDProcedureStage(IDProcedureStage *xnappducontentsv1.ProcedureStageChoice) (*xnappducontentsv1.EarlyStatusTransferIEsValue, error) {

	item := &xnappducontentsv1.EarlyStatusTransferIEsValue{
		EarlyStatusTransferIes: &xnappducontentsv1.EarlyStatusTransferIEsValue_IdProcedureStage{
			IdProcedureStage: IDProcedureStage,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateEarlyStatusTransferIEsValueIDProcedureStage() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateProcedureStageChoiceFirstDlCount(firstDlCount *xnappducontentsv1.FirstDlcount) (*xnappducontentsv1.ProcedureStageChoice, error) {

	item := &xnappducontentsv1.ProcedureStageChoice{
		ProcedureStageChoice: &xnappducontentsv1.ProcedureStageChoice_FirstDlCount{
			FirstDlCount: firstDlCount,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateProcedureStageChoiceFirstDlCount() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateProcedureStageChoiceDlDiscarding(dlDiscarding *xnappducontentsv1.Dldiscarding) (*xnappducontentsv1.ProcedureStageChoice, error) {

	item := &xnappducontentsv1.ProcedureStageChoice{
		ProcedureStageChoice: &xnappducontentsv1.ProcedureStageChoice_DlDiscarding{
			DlDiscarding: dlDiscarding,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateProcedureStageChoiceDlDiscarding() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateProcedureStageChoiceChoiceExtension(choiceExtension *xnappducontentsv1.ProcedureStageChoiceExtIes) (*xnappducontentsv1.ProcedureStageChoice, error) {

	item := &xnappducontentsv1.ProcedureStageChoice{
		ProcedureStageChoice: &xnappducontentsv1.ProcedureStageChoice_ChoiceExtension{
			ChoiceExtension: choiceExtension,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateProcedureStageChoiceChoiceExtension() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateRanpagingIEsValueIDUeidentityIndexValue(idUeidentityIndexValue *xnapiesv1.UeidentityIndexValue) (*xnappducontentsv1.RanpagingIEsValue, error) {

	item := &xnappducontentsv1.RanpagingIEsValue{
		RanpagingIes: &xnappducontentsv1.RanpagingIEsValue_IdUeidentityIndexValue{
			IdUeidentityIndexValue: idUeidentityIndexValue,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRanpagingIEsValueIDUeidentityIndexValue() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateRanpagingIEsValueIDUeranpagingIdentity(idUeranpagingIdentity *xnapiesv1.UeranpagingIdentity) (*xnappducontentsv1.RanpagingIEsValue, error) {

	item := &xnappducontentsv1.RanpagingIEsValue{
		RanpagingIes: &xnappducontentsv1.RanpagingIEsValue_IdUeranpagingIdentity{
			IdUeranpagingIdentity: idUeranpagingIdentity,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRanpagingIEsValueIDUeranpagingIdentity() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateRanpagingIEsValueIDPagingDrx(IDPagingDrx xnapiesv1.PagingDrx) (*xnappducontentsv1.RanpagingIEsValue, error) {

	item := &xnappducontentsv1.RanpagingIEsValue{
		RanpagingIes: &xnappducontentsv1.RanpagingIEsValue_IdPagingDrx{
			IdPagingDrx: IDPagingDrx,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRanpagingIEsValueIDPagingDrx() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateRanpagingIEsValueIDRanpagingArea(IDRanpagingArea *xnapiesv1.RanpagingArea) (*xnappducontentsv1.RanpagingIEsValue, error) {

	item := &xnappducontentsv1.RanpagingIEsValue{
		RanpagingIes: &xnappducontentsv1.RanpagingIEsValue_IdRanpagingArea{
			IdRanpagingArea: IDRanpagingArea,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRanpagingIEsValueIDRanpagingArea() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateRanpagingIEsValueIDPagingPriority(IDPagingPriority xnapiesv1.PagingPriority) (*xnappducontentsv1.RanpagingIEsValue, error) {

	item := &xnappducontentsv1.RanpagingIEsValue{
		RanpagingIes: &xnappducontentsv1.RanpagingIEsValue_IdPagingPriority{
			IdPagingPriority: IDPagingPriority,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRanpagingIEsValueIDPagingPriority() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateRanpagingIEsValueIDAssistanceDataForRanpaging(IDAssistanceDataForRanpaging *xnapiesv1.AssistanceDataForRanpaging) (*xnappducontentsv1.RanpagingIEsValue, error) {

	item := &xnappducontentsv1.RanpagingIEsValue{
		RanpagingIes: &xnappducontentsv1.RanpagingIEsValue_IdAssistanceDataForRanpaging{
			IdAssistanceDataForRanpaging: IDAssistanceDataForRanpaging,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRanpagingIEsValueIDAssistanceDataForRanpaging() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateRanpagingIEsValueIDUeradioCapabilityForPaging(IDUeradioCapabilityForPaging *xnapiesv1.UeradioCapabilityForPaging) (*xnappducontentsv1.RanpagingIEsValue, error) {

	item := &xnappducontentsv1.RanpagingIEsValue{
		RanpagingIes: &xnappducontentsv1.RanpagingIEsValue_IdUeradioCapabilityForPaging{
			IdUeradioCapabilityForPaging: IDUeradioCapabilityForPaging,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRanpagingIEsValueIDUeradioCapabilityForPaging() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateRanpagingIEsValueIDExtendedUeidentityIndexValue(idExtendedUeidentityIndexValue *xnapiesv1.ExtendedUeidentityIndexValue) (*xnappducontentsv1.RanpagingIEsValue, error) {

	item := &xnappducontentsv1.RanpagingIEsValue{
		RanpagingIes: &xnappducontentsv1.RanpagingIEsValue_IdExtendedUeidentityIndexValue{
			IdExtendedUeidentityIndexValue: idExtendedUeidentityIndexValue,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRanpagingIEsValueIDExtendedUeidentityIndexValue() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateRanpagingIEsValueIDPagingeDrxinformation(IDPagingeDrxinformation *xnapiesv1.PagingeDrxinformation) (*xnappducontentsv1.RanpagingIEsValue, error) {

	item := &xnappducontentsv1.RanpagingIEsValue{
		RanpagingIes: &xnappducontentsv1.RanpagingIEsValue_IdPagingeDrxinformation{
			IdPagingeDrxinformation: IDPagingeDrxinformation,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRanpagingIEsValueIDPagingeDrxinformation() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateRanpagingIEsValueIDUespecificDrx(IDUespecificDrx xnapiesv1.UespecificDrx) (*xnappducontentsv1.RanpagingIEsValue, error) {

	item := &xnappducontentsv1.RanpagingIEsValue{
		RanpagingIes: &xnappducontentsv1.RanpagingIEsValue_IdUespecificDrx{
			IdUespecificDrx: IDUespecificDrx,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRanpagingIEsValueIDUespecificDrx() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateRetrieveUecontextRequestIEsValueIDNewNgRannodeUexnApID(IDNewNgRannodeUexnApID *xnapiesv1.NgRAnnodeUexnApid) (*xnappducontentsv1.RetrieveUecontextRequestIEsValue, error) {

	item := &xnappducontentsv1.RetrieveUecontextRequestIEsValue{
		RetrieveUecontextRequestIes: &xnappducontentsv1.RetrieveUecontextRequestIEsValue_IdNewNgRannodeUexnApid{
			IdNewNgRannodeUexnApid: IDNewNgRannodeUexnApID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRetrieveUecontextRequestIEsValueIDNewNgRannodeUexnApID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateRetrieveUecontextRequestIEsValueIDUecontextID(IDUecontextID *xnapiesv1.UecontextId) (*xnappducontentsv1.RetrieveUecontextRequestIEsValue, error) {

	item := &xnappducontentsv1.RetrieveUecontextRequestIEsValue{
		RetrieveUecontextRequestIes: &xnappducontentsv1.RetrieveUecontextRequestIEsValue_IdUecontextId{
			IdUecontextId: IDUecontextID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRetrieveUecontextRequestIEsValueIDUecontextID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateRetrieveUecontextRequestIEsValueIDMacI(IDMacI *xnapiesv1.MacI) (*xnappducontentsv1.RetrieveUecontextRequestIEsValue, error) {

	item := &xnappducontentsv1.RetrieveUecontextRequestIEsValue{
		RetrieveUecontextRequestIes: &xnappducontentsv1.RetrieveUecontextRequestIEsValue_IdMacI{
			IdMacI: IDMacI,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRetrieveUecontextRequestIEsValueIDMacI() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateRetrieveUecontextRequestIEsValueIDNewNgRanCellIdentity(idNewNgRanCellIdentity *xnapiesv1.NgRAnCellIdentity) (*xnappducontentsv1.RetrieveUecontextRequestIEsValue, error) {

	item := &xnappducontentsv1.RetrieveUecontextRequestIEsValue{
		RetrieveUecontextRequestIes: &xnappducontentsv1.RetrieveUecontextRequestIEsValue_IdNewNgRanCellIdentity{
			IdNewNgRanCellIdentity: idNewNgRanCellIdentity,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRetrieveUecontextRequestIEsValueIDNewNgRanCellIdentity() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateRetrieveUecontextRequestIEsValueIDRrcresumeCause(IDRrcresumeCause xnapiesv1.RrcresumeCause) (*xnappducontentsv1.RetrieveUecontextRequestIEsValue, error) {

	item := &xnappducontentsv1.RetrieveUecontextRequestIEsValue{
		RetrieveUecontextRequestIes: &xnappducontentsv1.RetrieveUecontextRequestIEsValue_IdRrcresumeCause{
			IdRrcresumeCause: IDRrcresumeCause,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRetrieveUecontextRequestIEsValueIDRrcresumeCause() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateRetrieveUecontextResponseIEsValueIDNewNgRannodeUexnApID(IDNewNgRannodeUexnApID *xnapiesv1.NgRAnnodeUexnApid) (*xnappducontentsv1.RetrieveUecontextResponseIEsValue, error) {

	item := &xnappducontentsv1.RetrieveUecontextResponseIEsValue{
		RetrieveUecontextResponseIes: &xnappducontentsv1.RetrieveUecontextResponseIEsValue_IdNewNgRannodeUexnApid{
			IdNewNgRannodeUexnApid: IDNewNgRannodeUexnApID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRetrieveUecontextResponseIEsValueIDNewNgRannodeUexnApID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateRetrieveUecontextResponseIEsValueIDOldNgRannodeUexnApID(IDOldNgRannodeUexnApID *xnapiesv1.NgRAnnodeUexnApid) (*xnappducontentsv1.RetrieveUecontextResponseIEsValue, error) {

	item := &xnappducontentsv1.RetrieveUecontextResponseIEsValue{
		RetrieveUecontextResponseIes: &xnappducontentsv1.RetrieveUecontextResponseIEsValue_IdOldNgRannodeUexnApid{
			IdOldNgRannodeUexnApid: IDOldNgRannodeUexnApID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRetrieveUecontextResponseIEsValueIDOldNgRannodeUexnApID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateRetrieveUecontextResponseIEsValueIDGuami(IDGuami *xnapiesv1.Guami) (*xnappducontentsv1.RetrieveUecontextResponseIEsValue, error) {

	item := &xnappducontentsv1.RetrieveUecontextResponseIEsValue{
		RetrieveUecontextResponseIes: &xnappducontentsv1.RetrieveUecontextResponseIEsValue_IdGuami{
			IdGuami: IDGuami,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRetrieveUecontextResponseIEsValueIDGuami() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateRetrieveUecontextResponseIEsValueIDUecontextInfoRetrUectxtResp(IDUecontextInfoRetrUectxtResp *xnapiesv1.UecontextInfoRetrUectxtResp) (*xnappducontentsv1.RetrieveUecontextResponseIEsValue, error) {

	item := &xnappducontentsv1.RetrieveUecontextResponseIEsValue{
		RetrieveUecontextResponseIes: &xnappducontentsv1.RetrieveUecontextResponseIEsValue_IdUecontextInfoRetrUectxtResp{
			IdUecontextInfoRetrUectxtResp: IDUecontextInfoRetrUectxtResp,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRetrieveUecontextResponseIEsValueIDUecontextInfoRetrUectxtResp() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateRetrieveUecontextResponseIEsValueIDTraceActivation(IDTraceActivation *xnapiesv1.TraceActivation) (*xnappducontentsv1.RetrieveUecontextResponseIEsValue, error) {

	item := &xnappducontentsv1.RetrieveUecontextResponseIEsValue{
		RetrieveUecontextResponseIes: &xnappducontentsv1.RetrieveUecontextResponseIEsValue_IdTraceActivation{
			IdTraceActivation: IDTraceActivation,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRetrieveUecontextResponseIEsValueIDTraceActivation() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateRetrieveUecontextResponseIEsValueIDMaskedImeisv(IDMaskedImeisv *xnapiesv1.MaskedImeisv) (*xnappducontentsv1.RetrieveUecontextResponseIEsValue, error) {

	item := &xnappducontentsv1.RetrieveUecontextResponseIEsValue{
		RetrieveUecontextResponseIes: &xnappducontentsv1.RetrieveUecontextResponseIEsValue_IdMaskedImeisv{
			IdMaskedImeisv: IDMaskedImeisv,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRetrieveUecontextResponseIEsValueIDMaskedImeisv() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateRetrieveUecontextResponseIEsValueIDLocationReportingInformation(IDLocationReportingInformation *xnapiesv1.LocationReportingInformation) (*xnappducontentsv1.RetrieveUecontextResponseIEsValue, error) {

	item := &xnappducontentsv1.RetrieveUecontextResponseIEsValue{
		RetrieveUecontextResponseIes: &xnappducontentsv1.RetrieveUecontextResponseIEsValue_IdLocationReportingInformation{
			IdLocationReportingInformation: IDLocationReportingInformation,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRetrieveUecontextResponseIEsValueIDLocationReportingInformation() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateRetrieveUecontextResponseIEsValueIDCriticalityDiagnostics(IDCriticalityDiagnostics *xnapiesv1.CriticalityDiagnostics) (*xnappducontentsv1.RetrieveUecontextResponseIEsValue, error) {

	item := &xnappducontentsv1.RetrieveUecontextResponseIEsValue{
		RetrieveUecontextResponseIes: &xnappducontentsv1.RetrieveUecontextResponseIEsValue_IdCriticalityDiagnostics{
			IdCriticalityDiagnostics: IDCriticalityDiagnostics,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRetrieveUecontextResponseIEsValueIDCriticalityDiagnostics() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateRetrieveUecontextResponseIEsValueIDNrv2XservicesAuthorized(IDNrv2XservicesAuthorized *xnapiesv1.Nrv2XservicesAuthorized) (*xnappducontentsv1.RetrieveUecontextResponseIEsValue, error) {

	item := &xnappducontentsv1.RetrieveUecontextResponseIEsValue{
		RetrieveUecontextResponseIes: &xnappducontentsv1.RetrieveUecontextResponseIEsValue_IdNrv2XservicesAuthorized{
			IdNrv2XservicesAuthorized: IDNrv2XservicesAuthorized,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRetrieveUecontextResponseIEsValueIDNrv2XservicesAuthorized() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateRetrieveUecontextResponseIEsValueIDLtev2XservicesAuthorized(IDLtev2XservicesAuthorized *xnapiesv1.Ltev2XservicesAuthorized) (*xnappducontentsv1.RetrieveUecontextResponseIEsValue, error) {

	item := &xnappducontentsv1.RetrieveUecontextResponseIEsValue{
		RetrieveUecontextResponseIes: &xnappducontentsv1.RetrieveUecontextResponseIEsValue_IdLtev2XservicesAuthorized{
			IdLtev2XservicesAuthorized: IDLtev2XservicesAuthorized,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRetrieveUecontextResponseIEsValueIDLtev2XservicesAuthorized() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateRetrieveUecontextResponseIEsValueIDPc5QoSparameters(IDPc5QoSparameters *xnapiesv1.Pc5QoSparameters) (*xnappducontentsv1.RetrieveUecontextResponseIEsValue, error) {

	item := &xnappducontentsv1.RetrieveUecontextResponseIEsValue{
		RetrieveUecontextResponseIes: &xnappducontentsv1.RetrieveUecontextResponseIEsValue_IdPc5QoSparameters{
			IdPc5QoSparameters: IDPc5QoSparameters,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRetrieveUecontextResponseIEsValueIDPc5QoSparameters() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateRetrieveUecontextResponseIEsValueIDUehistoryInformation(IDUehistoryInformation *xnapiesv1.UehistoryInformation) (*xnappducontentsv1.RetrieveUecontextResponseIEsValue, error) {

	item := &xnappducontentsv1.RetrieveUecontextResponseIEsValue{
		RetrieveUecontextResponseIes: &xnappducontentsv1.RetrieveUecontextResponseIEsValue_IdUehistoryInformation{
			IdUehistoryInformation: IDUehistoryInformation,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRetrieveUecontextResponseIEsValueIDUehistoryInformation() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateRetrieveUecontextResponseIEsValueIDUehistoryInformationFromTheUe(IDUehistoryInformationFromTheUe *xnapiesv1.UehistoryInformationFromTheUe) (*xnappducontentsv1.RetrieveUecontextResponseIEsValue, error) {

	item := &xnappducontentsv1.RetrieveUecontextResponseIEsValue{
		RetrieveUecontextResponseIes: &xnappducontentsv1.RetrieveUecontextResponseIEsValue_IdUehistoryInformationFromTheUe{
			IdUehistoryInformationFromTheUe: IDUehistoryInformationFromTheUe,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRetrieveUecontextResponseIEsValueIDUehistoryInformationFromTheUe() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateRetrieveUecontextResponseIEsValueIDMdtplmnlist(IDMdtplmnlist *xnapiesv1.Mdtplmnlist) (*xnappducontentsv1.RetrieveUecontextResponseIEsValue, error) {

	item := &xnappducontentsv1.RetrieveUecontextResponseIEsValue{
		RetrieveUecontextResponseIes: &xnappducontentsv1.RetrieveUecontextResponseIEsValue_IdMdtplmnlist{
			IdMdtplmnlist: IDMdtplmnlist,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRetrieveUecontextResponseIEsValueIDMdtplmnlist() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateRetrieveUecontextFailureIEsValueIDNewNgRannodeUexnApID(IDNewNgRannodeUexnApID *xnapiesv1.NgRAnnodeUexnApid) (*xnappducontentsv1.RetrieveUecontextFailureIEsValue, error) {

	item := &xnappducontentsv1.RetrieveUecontextFailureIEsValue{
		RetrieveUecontextFailureIes: &xnappducontentsv1.RetrieveUecontextFailureIEsValue_IdNewNgRannodeUexnApid{
			IdNewNgRannodeUexnApid: IDNewNgRannodeUexnApID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRetrieveUecontextFailureIEsValueIDNewNgRannodeUexnApID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateRetrieveUecontextFailureIEsValueIDOldtoNewNgRannodeResumeContainer(IDOldtoNewNgRannodeResumeContainer string) (*xnappducontentsv1.RetrieveUecontextFailureIEsValue, error) {

	item := &xnappducontentsv1.RetrieveUecontextFailureIEsValue{
		RetrieveUecontextFailureIes: &xnappducontentsv1.RetrieveUecontextFailureIEsValue_IdOldtoNewNgRannodeResumeContainer{
			IdOldtoNewNgRannodeResumeContainer: IDOldtoNewNgRannodeResumeContainer,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRetrieveUecontextFailureIEsValueIDOldtoNewNgRannodeResumeContainer() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateRetrieveUecontextFailureIEsValueIDCause(IDCause *xnapiesv1.Cause) (*xnappducontentsv1.RetrieveUecontextFailureIEsValue, error) {

	item := &xnappducontentsv1.RetrieveUecontextFailureIEsValue{
		RetrieveUecontextFailureIes: &xnappducontentsv1.RetrieveUecontextFailureIEsValue_IdCause{
			IdCause: IDCause,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRetrieveUecontextFailureIEsValueIDCause() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateRetrieveUecontextFailureIEsValueIDCriticalityDiagnostics(IDCriticalityDiagnostics *xnapiesv1.CriticalityDiagnostics) (*xnappducontentsv1.RetrieveUecontextFailureIEsValue, error) {

	item := &xnappducontentsv1.RetrieveUecontextFailureIEsValue{
		RetrieveUecontextFailureIes: &xnappducontentsv1.RetrieveUecontextFailureIEsValue_IdCriticalityDiagnostics{
			IdCriticalityDiagnostics: IDCriticalityDiagnostics,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRetrieveUecontextFailureIEsValueIDCriticalityDiagnostics() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateXnUaddressIndicationIEsValueIDNewNgRannodeUexnApID(IDNewNgRannodeUexnApID *xnapiesv1.NgRAnnodeUexnApid) (*xnappducontentsv1.XnUaddressIndicationIEsValue, error) {

	item := &xnappducontentsv1.XnUaddressIndicationIEsValue{
		XnUaddressIndicationIes: &xnappducontentsv1.XnUaddressIndicationIEsValue_IdNewNgRannodeUexnApid{
			IdNewNgRannodeUexnApid: IDNewNgRannodeUexnApID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateXnUaddressIndicationIEsValueIDNewNgRannodeUexnApID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateXnUaddressIndicationIEsValueIDOldNgRannodeUexnApID(IDOldNgRannodeUexnApID *xnapiesv1.NgRAnnodeUexnApid) (*xnappducontentsv1.XnUaddressIndicationIEsValue, error) {

	item := &xnappducontentsv1.XnUaddressIndicationIEsValue{
		XnUaddressIndicationIes: &xnappducontentsv1.XnUaddressIndicationIEsValue_IdOldNgRannodeUexnApid{
			IdOldNgRannodeUexnApid: IDOldNgRannodeUexnApID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateXnUaddressIndicationIEsValueIDOldNgRannodeUexnApID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateXnUaddressIndicationIEsValueIDXnUaddressInfoperPdusessionList(IDXnUaddressInfoperPdusessionList *xnapiesv1.XnUaddressInfoperPdusessionList) (*xnappducontentsv1.XnUaddressIndicationIEsValue, error) {

	item := &xnappducontentsv1.XnUaddressIndicationIEsValue{
		XnUaddressIndicationIes: &xnappducontentsv1.XnUaddressIndicationIEsValue_IdXnUaddressInfoperPdusessionList{
			IdXnUaddressInfoperPdusessionList: IDXnUaddressInfoperPdusessionList,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateXnUaddressIndicationIEsValueIDXnUaddressInfoperPdusessionList() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateXnUaddressIndicationIEsValueIDChoMrdcIndicator(IDChoMrdcIndicator xnapiesv1.ChoMRdcIndicator) (*xnappducontentsv1.XnUaddressIndicationIEsValue, error) {

	item := &xnappducontentsv1.XnUaddressIndicationIEsValue{
		XnUaddressIndicationIes: &xnappducontentsv1.XnUaddressIndicationIEsValue_IdChoMrdcIndicator{
			IdChoMrdcIndicator: IDChoMrdcIndicator,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateXnUaddressIndicationIEsValueIDChoMrdcIndicator() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateXnUaddressIndicationIEsValueIDChoMrdcEarlyDataForwarding(IDChoMrdcEarlyDataForwarding xnapiesv1.ChoMRdcEarlyDataForwarding) (*xnappducontentsv1.XnUaddressIndicationIEsValue, error) {

	item := &xnappducontentsv1.XnUaddressIndicationIEsValue{
		XnUaddressIndicationIes: &xnappducontentsv1.XnUaddressIndicationIEsValue_IdChoMrdcEarlyDataForwarding{
			IdChoMrdcEarlyDataForwarding: IDChoMrdcEarlyDataForwarding,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateXnUaddressIndicationIEsValueIDChoMrdcEarlyDataForwarding() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeAdditionRequestIEsValueIDMNgRannodeUexnApID(IDMNgRannodeUexnApID *xnapiesv1.NgRAnnodeUexnApid) (*xnappducontentsv1.SnodeAdditionRequestIEsValue, error) {

	item := &xnappducontentsv1.SnodeAdditionRequestIEsValue{
		SnodeAdditionRequestIes: &xnappducontentsv1.SnodeAdditionRequestIEsValue_IdMNgRannodeUexnApid{
			IdMNgRannodeUexnApid: IDMNgRannodeUexnApID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeAdditionRequestIEsValueIDMNgRannodeUexnApID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeAdditionRequestIEsValueIDUesecurityCapabilities(IDUesecurityCapabilities *xnapiesv1.UesecurityCapabilities) (*xnappducontentsv1.SnodeAdditionRequestIEsValue, error) {

	item := &xnappducontentsv1.SnodeAdditionRequestIEsValue{
		SnodeAdditionRequestIes: &xnappducontentsv1.SnodeAdditionRequestIEsValue_IdUesecurityCapabilities{
			IdUesecurityCapabilities: IDUesecurityCapabilities,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeAdditionRequestIEsValueIDUesecurityCapabilities() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeAdditionRequestIEsValueIDSNgRannodeSecurityKey(IDSNgRannodeSecurityKey *xnapiesv1.SNGRAnnodeSecurityKey) (*xnappducontentsv1.SnodeAdditionRequestIEsValue, error) {

	item := &xnappducontentsv1.SnodeAdditionRequestIEsValue{
		SnodeAdditionRequestIes: &xnappducontentsv1.SnodeAdditionRequestIEsValue_IdSNgRannodeSecurityKey{
			IdSNgRannodeSecurityKey: IDSNgRannodeSecurityKey,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeAdditionRequestIEsValueIDSNgRannodeSecurityKey() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeAdditionRequestIEsValueIDSNgRannodeUeAmbr(IDSNgRannodeUeAmbr *xnapiesv1.UeaggregateMaximumBitRate) (*xnappducontentsv1.SnodeAdditionRequestIEsValue, error) {

	item := &xnappducontentsv1.SnodeAdditionRequestIEsValue{
		SnodeAdditionRequestIes: &xnappducontentsv1.SnodeAdditionRequestIEsValue_IdSNgRannodeUeAmbr{
			IdSNgRannodeUeAmbr: IDSNgRannodeUeAmbr,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeAdditionRequestIEsValueIDSNgRannodeUeAmbr() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeAdditionRequestIEsValueIDSelectedPlmn(IDSelectedPlmn *xnapiesv1.PlmnIdentity) (*xnappducontentsv1.SnodeAdditionRequestIEsValue, error) {

	item := &xnappducontentsv1.SnodeAdditionRequestIEsValue{
		SnodeAdditionRequestIes: &xnappducontentsv1.SnodeAdditionRequestIEsValue_IdSelectedPlmn{
			IdSelectedPlmn: IDSelectedPlmn,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeAdditionRequestIEsValueIDSelectedPlmn() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeAdditionRequestIEsValueIDMobilityRestrictionList(IDMobilityRestrictionList *xnapiesv1.MobilityRestrictionList) (*xnappducontentsv1.SnodeAdditionRequestIEsValue, error) {

	item := &xnappducontentsv1.SnodeAdditionRequestIEsValue{
		SnodeAdditionRequestIes: &xnappducontentsv1.SnodeAdditionRequestIEsValue_IdMobilityRestrictionList{
			IdMobilityRestrictionList: IDMobilityRestrictionList,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeAdditionRequestIEsValueIDMobilityRestrictionList() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeAdditionRequestIEsValueIDIndexToRatFrequSelectionPriority(IDIndexToRatFrequSelectionPriority *xnapiesv1.RfspIndex) (*xnappducontentsv1.SnodeAdditionRequestIEsValue, error) {

	item := &xnappducontentsv1.SnodeAdditionRequestIEsValue{
		SnodeAdditionRequestIes: &xnappducontentsv1.SnodeAdditionRequestIEsValue_IdIndexToRatFrequSelectionPriority{
			IdIndexToRatFrequSelectionPriority: IDIndexToRatFrequSelectionPriority,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeAdditionRequestIEsValueIDIndexToRatFrequSelectionPriority() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeAdditionRequestIEsValueIDPdusessionToBeAddedAddReq(IDPdusessionToBeAddedAddReq *xnappducontentsv1.PdusessionToBeAddedAddReq) (*xnappducontentsv1.SnodeAdditionRequestIEsValue, error) {

	item := &xnappducontentsv1.SnodeAdditionRequestIEsValue{
		SnodeAdditionRequestIes: &xnappducontentsv1.SnodeAdditionRequestIEsValue_IdPdusessionToBeAddedAddReq{
			IdPdusessionToBeAddedAddReq: IDPdusessionToBeAddedAddReq,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeAdditionRequestIEsValueIDPdusessionToBeAddedAddReq() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeAdditionRequestIEsValueIDMnToSnContainer(IDMnToSnContainer string) (*xnappducontentsv1.SnodeAdditionRequestIEsValue, error) {

	item := &xnappducontentsv1.SnodeAdditionRequestIEsValue{
		SnodeAdditionRequestIes: &xnappducontentsv1.SnodeAdditionRequestIEsValue_IdMnToSnContainer{
			IdMnToSnContainer: IDMnToSnContainer,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeAdditionRequestIEsValueIDMnToSnContainer() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeAdditionRequestIEsValueIDSNgRannodeUexnApID(IDSNgRannodeUexnApID *xnapiesv1.NgRAnnodeUexnApid) (*xnappducontentsv1.SnodeAdditionRequestIEsValue, error) {

	item := &xnappducontentsv1.SnodeAdditionRequestIEsValue{
		SnodeAdditionRequestIes: &xnappducontentsv1.SnodeAdditionRequestIEsValue_IdSNgRannodeUexnApid{
			IdSNgRannodeUexnApid: IDSNgRannodeUexnApID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeAdditionRequestIEsValueIDSNgRannodeUexnApID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeAdditionRequestIEsValueIDExpectedUebehaviour(IDExpectedUebehaviour *xnapiesv1.ExpectedUebehaviour) (*xnappducontentsv1.SnodeAdditionRequestIEsValue, error) {

	item := &xnappducontentsv1.SnodeAdditionRequestIEsValue{
		SnodeAdditionRequestIes: &xnappducontentsv1.SnodeAdditionRequestIEsValue_IdExpectedUebehaviour{
			IdExpectedUebehaviour: IDExpectedUebehaviour,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeAdditionRequestIEsValueIDExpectedUebehaviour() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeAdditionRequestIEsValueIDRequestedSplitSrb(IDRequestedSplitSrb xnapiesv1.SplitSrbsTypes) (*xnappducontentsv1.SnodeAdditionRequestIEsValue, error) {

	item := &xnappducontentsv1.SnodeAdditionRequestIEsValue{
		SnodeAdditionRequestIes: &xnappducontentsv1.SnodeAdditionRequestIEsValue_IdRequestedSplitSrb{
			IdRequestedSplitSrb: IDRequestedSplitSrb,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeAdditionRequestIEsValueIDRequestedSplitSrb() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeAdditionRequestIEsValueIDPcellID(IDPcellID *xnapiesv1.GlobalNgRAncellID) (*xnappducontentsv1.SnodeAdditionRequestIEsValue, error) {

	item := &xnappducontentsv1.SnodeAdditionRequestIEsValue{
		SnodeAdditionRequestIes: &xnappducontentsv1.SnodeAdditionRequestIEsValue_IdPcellId{
			IdPcellId: IDPcellID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeAdditionRequestIEsValueIDPcellID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeAdditionRequestIEsValueIDDesiredActNotificationLevel(IDDesiredActNotificationLevel xnapiesv1.DesiredActNotificationLevel) (*xnappducontentsv1.SnodeAdditionRequestIEsValue, error) {

	item := &xnappducontentsv1.SnodeAdditionRequestIEsValue{
		SnodeAdditionRequestIes: &xnappducontentsv1.SnodeAdditionRequestIEsValue_IdDesiredActNotificationLevel{
			IdDesiredActNotificationLevel: IDDesiredActNotificationLevel,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeAdditionRequestIEsValueIDDesiredActNotificationLevel() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeAdditionRequestIEsValueIDAvailableDrbIDs(IDAvailableDrbIDs *xnapiesv1.DrbList) (*xnappducontentsv1.SnodeAdditionRequestIEsValue, error) {

	item := &xnappducontentsv1.SnodeAdditionRequestIEsValue{
		SnodeAdditionRequestIes: &xnappducontentsv1.SnodeAdditionRequestIEsValue_IdAvailableDrbids{
			IdAvailableDrbids: IDAvailableDrbIDs,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeAdditionRequestIEsValueIDAvailableDrbIDs() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeAdditionRequestIEsValueIDSNgRannodeMaxIpdataRateUl(IDSNgRannodeMaxIpdataRateUl *xnapiesv1.BitRate) (*xnappducontentsv1.SnodeAdditionRequestIEsValue, error) {

	item := &xnappducontentsv1.SnodeAdditionRequestIEsValue{
		SnodeAdditionRequestIes: &xnappducontentsv1.SnodeAdditionRequestIEsValue_IdSNgRannodeMaxIpdataRateUl{
			IdSNgRannodeMaxIpdataRateUl: IDSNgRannodeMaxIpdataRateUl,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeAdditionRequestIEsValueIDSNgRannodeMaxIpdataRateUl() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeAdditionRequestIEsValueIDSNgRannodeMaxIpdataRateDl(IDSNgRannodeMaxIpdataRateDl *xnapiesv1.BitRate) (*xnappducontentsv1.SnodeAdditionRequestIEsValue, error) {

	item := &xnappducontentsv1.SnodeAdditionRequestIEsValue{
		SnodeAdditionRequestIes: &xnappducontentsv1.SnodeAdditionRequestIEsValue_IdSNgRannodeMaxIpdataRateDl{
			IdSNgRannodeMaxIpdataRateDl: IDSNgRannodeMaxIpdataRateDl,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeAdditionRequestIEsValueIDSNgRannodeMaxIpdataRateDl() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeAdditionRequestIEsValueIDLocationInformationSnreporting(IDLocationInformationSnreporting xnapiesv1.LocationInformationSnreporting) (*xnappducontentsv1.SnodeAdditionRequestIEsValue, error) {

	item := &xnappducontentsv1.SnodeAdditionRequestIEsValue{
		SnodeAdditionRequestIes: &xnappducontentsv1.SnodeAdditionRequestIEsValue_IdLocationInformationSnreporting{
			IdLocationInformationSnreporting: IDLocationInformationSnreporting,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeAdditionRequestIEsValueIDLocationInformationSnreporting() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeAdditionRequestIEsValueIDMrDcResourceCoordinationInfo(IDMrDcResourceCoordinationInfo *xnapiesv1.MrDCResourceCoordinationInfo) (*xnappducontentsv1.SnodeAdditionRequestIEsValue, error) {

	item := &xnappducontentsv1.SnodeAdditionRequestIEsValue{
		SnodeAdditionRequestIes: &xnappducontentsv1.SnodeAdditionRequestIEsValue_IdMrDcResourceCoordinationInfo{
			IdMrDcResourceCoordinationInfo: IDMrDcResourceCoordinationInfo,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeAdditionRequestIEsValueIDMrDcResourceCoordinationInfo() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeAdditionRequestIEsValueIDMaskedImeisv(IDMaskedImeisv *xnapiesv1.MaskedImeisv) (*xnappducontentsv1.SnodeAdditionRequestIEsValue, error) {

	item := &xnappducontentsv1.SnodeAdditionRequestIEsValue{
		SnodeAdditionRequestIes: &xnappducontentsv1.SnodeAdditionRequestIEsValue_IdMaskedImeisv{
			IdMaskedImeisv: IDMaskedImeisv,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeAdditionRequestIEsValueIDMaskedImeisv() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeAdditionRequestIEsValueIDNeDcTdmPattern(IDNeDcTdmPattern *xnapiesv1.NeDCTDmPattern) (*xnappducontentsv1.SnodeAdditionRequestIEsValue, error) {

	item := &xnappducontentsv1.SnodeAdditionRequestIEsValue{
		SnodeAdditionRequestIes: &xnappducontentsv1.SnodeAdditionRequestIEsValue_IdNeDcTdmPattern{
			IdNeDcTdmPattern: IDNeDcTdmPattern,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeAdditionRequestIEsValueIDNeDcTdmPattern() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeAdditionRequestIEsValueIDSNgRannodeAdditionTriggerInd(IDSNgRannodeAdditionTriggerInd xnapiesv1.SNGRAnnodeAdditionTriggerInd) (*xnappducontentsv1.SnodeAdditionRequestIEsValue, error) {

	item := &xnappducontentsv1.SnodeAdditionRequestIEsValue{
		SnodeAdditionRequestIes: &xnappducontentsv1.SnodeAdditionRequestIEsValue_IdSNgRannodeAdditionTriggerInd{
			IdSNgRannodeAdditionTriggerInd: IDSNgRannodeAdditionTriggerInd,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeAdditionRequestIEsValueIDSNgRannodeAdditionTriggerInd() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeAdditionRequestIEsValueIDTraceActivation(IDTraceActivation *xnapiesv1.TraceActivation) (*xnappducontentsv1.SnodeAdditionRequestIEsValue, error) {

	item := &xnappducontentsv1.SnodeAdditionRequestIEsValue{
		SnodeAdditionRequestIes: &xnappducontentsv1.SnodeAdditionRequestIEsValue_IdTraceActivation{
			IdTraceActivation: IDTraceActivation,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeAdditionRequestIEsValueIDTraceActivation() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeAdditionRequestIEsValueIDRequestedFastMcgrecoveryViaSrb3(IDRequestedFastMcgrecoveryViaSrb3 xnappducontentsv1.RequestedFastMcgrecoveryViaSrb3) (*xnappducontentsv1.SnodeAdditionRequestIEsValue, error) {

	item := &xnappducontentsv1.SnodeAdditionRequestIEsValue{
		SnodeAdditionRequestIes: &xnappducontentsv1.SnodeAdditionRequestIEsValue_IdRequestedFastMcgrecoveryViaSrb3{
			IdRequestedFastMcgrecoveryViaSrb3: IDRequestedFastMcgrecoveryViaSrb3,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeAdditionRequestIEsValueIDRequestedFastMcgrecoveryViaSrb3() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeAdditionRequestIEsValueIDUeradioCapabilityID(IDUeradioCapabilityID *xnapiesv1.UeradioCapabilityId) (*xnappducontentsv1.SnodeAdditionRequestIEsValue, error) {

	item := &xnappducontentsv1.SnodeAdditionRequestIEsValue{
		SnodeAdditionRequestIes: &xnappducontentsv1.SnodeAdditionRequestIEsValue_IdUeradioCapabilityId{
			IdUeradioCapabilityId: IDUeradioCapabilityID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeAdditionRequestIEsValueIDUeradioCapabilityID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeAdditionRequestIEsValueIDSourceNgRanNodeID(IDSourceNgRanNodeID *xnapiesv1.GlobalNgRAnnodeID) (*xnappducontentsv1.SnodeAdditionRequestIEsValue, error) {

	item := &xnappducontentsv1.SnodeAdditionRequestIEsValue{
		SnodeAdditionRequestIes: &xnappducontentsv1.SnodeAdditionRequestIEsValue_IdSourceNgRanNodeId{
			IdSourceNgRanNodeId: IDSourceNgRanNodeID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeAdditionRequestIEsValueIDSourceNgRanNodeID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeAdditionRequestAcknowledgeIEsValueIDMNgRannodeUexnApID(IDMNgRannodeUexnApID *xnapiesv1.NgRAnnodeUexnApid) (*xnappducontentsv1.SnodeAdditionRequestAcknowledgeIEsValue, error) {

	item := &xnappducontentsv1.SnodeAdditionRequestAcknowledgeIEsValue{
		SnodeAdditionRequestAcknowledgeIes: &xnappducontentsv1.SnodeAdditionRequestAcknowledgeIEsValue_IdMNgRannodeUexnApid{
			IdMNgRannodeUexnApid: IDMNgRannodeUexnApID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeAdditionRequestAcknowledgeIEsValueIDMNgRannodeUexnApID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeAdditionRequestAcknowledgeIEsValueIDSNgRannodeUexnApID(IDSNgRannodeUexnApID *xnapiesv1.NgRAnnodeUexnApid) (*xnappducontentsv1.SnodeAdditionRequestAcknowledgeIEsValue, error) {

	item := &xnappducontentsv1.SnodeAdditionRequestAcknowledgeIEsValue{
		SnodeAdditionRequestAcknowledgeIes: &xnappducontentsv1.SnodeAdditionRequestAcknowledgeIEsValue_IdSNgRannodeUexnApid{
			IdSNgRannodeUexnApid: IDSNgRannodeUexnApID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeAdditionRequestAcknowledgeIEsValueIDSNgRannodeUexnApID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeAdditionRequestAcknowledgeIEsValueIDPdusessionAdmittedAddedAddReqAck(IDPdusessionAdmittedAddedAddReqAck *xnappducontentsv1.PdusessionAdmittedAddedAddReqAck) (*xnappducontentsv1.SnodeAdditionRequestAcknowledgeIEsValue, error) {

	item := &xnappducontentsv1.SnodeAdditionRequestAcknowledgeIEsValue{
		SnodeAdditionRequestAcknowledgeIes: &xnappducontentsv1.SnodeAdditionRequestAcknowledgeIEsValue_IdPdusessionAdmittedAddedAddReqAck{
			IdPdusessionAdmittedAddedAddReqAck: IDPdusessionAdmittedAddedAddReqAck,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeAdditionRequestAcknowledgeIEsValueIDPdusessionAdmittedAddedAddReqAck() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeAdditionRequestAcknowledgeIEsValueIDPdusessionNotAdmittedAddReqAck(IDPdusessionNotAdmittedAddReqAck *xnappducontentsv1.PdusessionNotAdmittedAddReqAck) (*xnappducontentsv1.SnodeAdditionRequestAcknowledgeIEsValue, error) {

	item := &xnappducontentsv1.SnodeAdditionRequestAcknowledgeIEsValue{
		SnodeAdditionRequestAcknowledgeIes: &xnappducontentsv1.SnodeAdditionRequestAcknowledgeIEsValue_IdPdusessionNotAdmittedAddReqAck{
			IdPdusessionNotAdmittedAddReqAck: IDPdusessionNotAdmittedAddReqAck,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeAdditionRequestAcknowledgeIEsValueIDPdusessionNotAdmittedAddReqAck() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeAdditionRequestAcknowledgeIEsValueIDSnToMnContainer(IDSnToMnContainer string) (*xnappducontentsv1.SnodeAdditionRequestAcknowledgeIEsValue, error) {

	item := &xnappducontentsv1.SnodeAdditionRequestAcknowledgeIEsValue{
		SnodeAdditionRequestAcknowledgeIes: &xnappducontentsv1.SnodeAdditionRequestAcknowledgeIEsValue_IdSnToMnContainer{
			IdSnToMnContainer: IDSnToMnContainer,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeAdditionRequestAcknowledgeIEsValueIDSnToMnContainer() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeAdditionRequestAcknowledgeIEsValueIDAdmittedSplitSrb(IDAdmittedSplitSrb xnapiesv1.SplitSrbsTypes) (*xnappducontentsv1.SnodeAdditionRequestAcknowledgeIEsValue, error) {

	item := &xnappducontentsv1.SnodeAdditionRequestAcknowledgeIEsValue{
		SnodeAdditionRequestAcknowledgeIes: &xnappducontentsv1.SnodeAdditionRequestAcknowledgeIEsValue_IdAdmittedSplitSrb{
			IdAdmittedSplitSrb: IDAdmittedSplitSrb,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeAdditionRequestAcknowledgeIEsValueIDAdmittedSplitSrb() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeAdditionRequestAcknowledgeIEsValueIDRrcconfigIndication(IDRrcconfigIndication xnapiesv1.RrcconfigIndication) (*xnappducontentsv1.SnodeAdditionRequestAcknowledgeIEsValue, error) {

	item := &xnappducontentsv1.SnodeAdditionRequestAcknowledgeIEsValue{
		SnodeAdditionRequestAcknowledgeIes: &xnappducontentsv1.SnodeAdditionRequestAcknowledgeIEsValue_IdRrcconfigIndication{
			IdRrcconfigIndication: IDRrcconfigIndication,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeAdditionRequestAcknowledgeIEsValueIDRrcconfigIndication() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeAdditionRequestAcknowledgeIEsValueIDCriticalityDiagnostics(IDCriticalityDiagnostics *xnapiesv1.CriticalityDiagnostics) (*xnappducontentsv1.SnodeAdditionRequestAcknowledgeIEsValue, error) {

	item := &xnappducontentsv1.SnodeAdditionRequestAcknowledgeIEsValue{
		SnodeAdditionRequestAcknowledgeIes: &xnappducontentsv1.SnodeAdditionRequestAcknowledgeIEsValue_IdCriticalityDiagnostics{
			IdCriticalityDiagnostics: IDCriticalityDiagnostics,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeAdditionRequestAcknowledgeIEsValueIDCriticalityDiagnostics() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeAdditionRequestAcknowledgeIEsValueIDLocationInformationSn(IDLocationInformationSn *xnapiesv1.TargetCGi) (*xnappducontentsv1.SnodeAdditionRequestAcknowledgeIEsValue, error) {

	item := &xnappducontentsv1.SnodeAdditionRequestAcknowledgeIEsValue{
		SnodeAdditionRequestAcknowledgeIes: &xnappducontentsv1.SnodeAdditionRequestAcknowledgeIEsValue_IdLocationInformationSn{
			IdLocationInformationSn: IDLocationInformationSn,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeAdditionRequestAcknowledgeIEsValueIDLocationInformationSn() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeAdditionRequestAcknowledgeIEsValueIDMrDcResourceCoordinationInfo(IDMrDcResourceCoordinationInfo *xnapiesv1.MrDCResourceCoordinationInfo) (*xnappducontentsv1.SnodeAdditionRequestAcknowledgeIEsValue, error) {

	item := &xnappducontentsv1.SnodeAdditionRequestAcknowledgeIEsValue{
		SnodeAdditionRequestAcknowledgeIes: &xnappducontentsv1.SnodeAdditionRequestAcknowledgeIEsValue_IdMrDcResourceCoordinationInfo{
			IdMrDcResourceCoordinationInfo: IDMrDcResourceCoordinationInfo,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeAdditionRequestAcknowledgeIEsValueIDMrDcResourceCoordinationInfo() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeAdditionRequestAcknowledgeIEsValueIDAvailableFastMcgrecoveryViaSrb3(IDAvailableFastMcgrecoveryViaSrb3 xnappducontentsv1.AvailableFastMcgrecoveryViaSrb3) (*xnappducontentsv1.SnodeAdditionRequestAcknowledgeIEsValue, error) {

	item := &xnappducontentsv1.SnodeAdditionRequestAcknowledgeIEsValue{
		SnodeAdditionRequestAcknowledgeIes: &xnappducontentsv1.SnodeAdditionRequestAcknowledgeIEsValue_IdAvailableFastMcgrecoveryViaSrb3{
			IdAvailableFastMcgrecoveryViaSrb3: IDAvailableFastMcgrecoveryViaSrb3,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeAdditionRequestAcknowledgeIEsValueIDAvailableFastMcgrecoveryViaSrb3() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeAdditionRequestAcknowledgeIEsValueIDDirectForwardingPathAvailability(IDDirectForwardingPathAvailability xnapiesv1.DirectForwardingPathAvailability) (*xnappducontentsv1.SnodeAdditionRequestAcknowledgeIEsValue, error) {

	item := &xnappducontentsv1.SnodeAdditionRequestAcknowledgeIEsValue{
		SnodeAdditionRequestAcknowledgeIes: &xnappducontentsv1.SnodeAdditionRequestAcknowledgeIEsValue_IdDirectForwardingPathAvailability{
			IdDirectForwardingPathAvailability: IDDirectForwardingPathAvailability,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeAdditionRequestAcknowledgeIEsValueIDDirectForwardingPathAvailability() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeAdditionRequestRejectIEsValueIDMNgRannodeUexnApID(IDMNgRannodeUexnApID *xnapiesv1.NgRAnnodeUexnApid) (*xnappducontentsv1.SnodeAdditionRequestRejectIEsValue, error) {

	item := &xnappducontentsv1.SnodeAdditionRequestRejectIEsValue{
		SnodeAdditionRequestRejectIes: &xnappducontentsv1.SnodeAdditionRequestRejectIEsValue_IdMNgRannodeUexnApid{
			IdMNgRannodeUexnApid: IDMNgRannodeUexnApID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeAdditionRequestRejectIEsValueIDMNgRannodeUexnApID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeAdditionRequestRejectIEsValueIDSNgRannodeUexnApID(IDSNgRannodeUexnApID *xnapiesv1.NgRAnnodeUexnApid) (*xnappducontentsv1.SnodeAdditionRequestRejectIEsValue, error) {

	item := &xnappducontentsv1.SnodeAdditionRequestRejectIEsValue{
		SnodeAdditionRequestRejectIes: &xnappducontentsv1.SnodeAdditionRequestRejectIEsValue_IdSNgRannodeUexnApid{
			IdSNgRannodeUexnApid: IDSNgRannodeUexnApID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeAdditionRequestRejectIEsValueIDSNgRannodeUexnApID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeAdditionRequestRejectIEsValueIDCause(IDCause *xnapiesv1.Cause) (*xnappducontentsv1.SnodeAdditionRequestRejectIEsValue, error) {

	item := &xnappducontentsv1.SnodeAdditionRequestRejectIEsValue{
		SnodeAdditionRequestRejectIes: &xnappducontentsv1.SnodeAdditionRequestRejectIEsValue_IdCause{
			IdCause: IDCause,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeAdditionRequestRejectIEsValueIDCause() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeAdditionRequestRejectIEsValueIDCriticalityDiagnostics(IDCriticalityDiagnostics *xnapiesv1.CriticalityDiagnostics) (*xnappducontentsv1.SnodeAdditionRequestRejectIEsValue, error) {

	item := &xnappducontentsv1.SnodeAdditionRequestRejectIEsValue{
		SnodeAdditionRequestRejectIes: &xnappducontentsv1.SnodeAdditionRequestRejectIEsValue_IdCriticalityDiagnostics{
			IdCriticalityDiagnostics: IDCriticalityDiagnostics,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeAdditionRequestRejectIEsValueIDCriticalityDiagnostics() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeReconfigurationCompleteIEsValueIDMNgRannodeUexnApID(IDMNgRannodeUexnApID *xnapiesv1.NgRAnnodeUexnApid) (*xnappducontentsv1.SnodeReconfigurationCompleteIEsValue, error) {

	item := &xnappducontentsv1.SnodeReconfigurationCompleteIEsValue{
		SnodeReconfigurationCompleteIes: &xnappducontentsv1.SnodeReconfigurationCompleteIEsValue_IdMNgRannodeUexnApid{
			IdMNgRannodeUexnApid: IDMNgRannodeUexnApID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeReconfigurationCompleteIEsValueIDMNgRannodeUexnApID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeReconfigurationCompleteIEsValueIDSNgRannodeUexnApID(IDSNgRannodeUexnApID *xnapiesv1.NgRAnnodeUexnApid) (*xnappducontentsv1.SnodeReconfigurationCompleteIEsValue, error) {

	item := &xnappducontentsv1.SnodeReconfigurationCompleteIEsValue{
		SnodeReconfigurationCompleteIes: &xnappducontentsv1.SnodeReconfigurationCompleteIEsValue_IdSNgRannodeUexnApid{
			IdSNgRannodeUexnApid: IDSNgRannodeUexnApID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeReconfigurationCompleteIEsValueIDSNgRannodeUexnApID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeReconfigurationCompleteIEsValueIDResponseInfoReconfCompl(IDResponseInfoReconfCompl *xnappducontentsv1.ResponseInfoReconfCompl) (*xnappducontentsv1.SnodeReconfigurationCompleteIEsValue, error) {

	item := &xnappducontentsv1.SnodeReconfigurationCompleteIEsValue{
		SnodeReconfigurationCompleteIes: &xnappducontentsv1.SnodeReconfigurationCompleteIEsValue_IdResponseInfoReconfCompl{
			IdResponseInfoReconfCompl: IDResponseInfoReconfCompl,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeReconfigurationCompleteIEsValueIDResponseInfoReconfCompl() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateResponseTypeReconfCompleteConfigurationSuccessfullyApplied(configurationSuccessfullyApplied *xnappducontentsv1.Configurationsuccessfullyapplied) (*xnappducontentsv1.ResponseTypeReconfComplete, error) {

	item := &xnappducontentsv1.ResponseTypeReconfComplete{
		ResponseTypeReconfComplete: &xnappducontentsv1.ResponseTypeReconfComplete_ConfigurationSuccessfullyApplied{
			ConfigurationSuccessfullyApplied: configurationSuccessfullyApplied,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateResponseTypeReconfCompleteConfigurationSuccessfullyApplied() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateResponseTypeReconfCompleteConfigurationRejectedByMNgRannode(configurationRejectedByMNgRannode *xnappducontentsv1.ConfigurationrejectedbyMNGRAnnode) (*xnappducontentsv1.ResponseTypeReconfComplete, error) {

	item := &xnappducontentsv1.ResponseTypeReconfComplete{
		ResponseTypeReconfComplete: &xnappducontentsv1.ResponseTypeReconfComplete_ConfigurationRejectedByMNgRannode{
			ConfigurationRejectedByMNgRannode: configurationRejectedByMNgRannode,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateResponseTypeReconfCompleteConfigurationRejectedByMNgRannode() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateResponseTypeReconfCompleteChoiceExtension(choiceExtension *xnappducontentsv1.ResponseTypeReconfCompleteExtIes) (*xnappducontentsv1.ResponseTypeReconfComplete, error) {

	item := &xnappducontentsv1.ResponseTypeReconfComplete{
		ResponseTypeReconfComplete: &xnappducontentsv1.ResponseTypeReconfComplete_ChoiceExtension{
			ChoiceExtension: choiceExtension,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateResponseTypeReconfCompleteChoiceExtension() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeModificationRequestIEsValueIDMNgRannodeUexnApID(IDMNgRannodeUexnApID *xnapiesv1.NgRAnnodeUexnApid) (*xnappducontentsv1.SnodeModificationRequestIEsValue, error) {

	item := &xnappducontentsv1.SnodeModificationRequestIEsValue{
		SnodeModificationRequestIes: &xnappducontentsv1.SnodeModificationRequestIEsValue_IdMNgRannodeUexnApid{
			IdMNgRannodeUexnApid: IDMNgRannodeUexnApID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeModificationRequestIEsValueIDMNgRannodeUexnApID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeModificationRequestIEsValueIDSNgRannodeUexnApID(IDSNgRannodeUexnApID *xnapiesv1.NgRAnnodeUexnApid) (*xnappducontentsv1.SnodeModificationRequestIEsValue, error) {

	item := &xnappducontentsv1.SnodeModificationRequestIEsValue{
		SnodeModificationRequestIes: &xnappducontentsv1.SnodeModificationRequestIEsValue_IdSNgRannodeUexnApid{
			IdSNgRannodeUexnApid: IDSNgRannodeUexnApID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeModificationRequestIEsValueIDSNgRannodeUexnApID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeModificationRequestIEsValueIDCause(IDCause *xnapiesv1.Cause) (*xnappducontentsv1.SnodeModificationRequestIEsValue, error) {

	item := &xnappducontentsv1.SnodeModificationRequestIEsValue{
		SnodeModificationRequestIes: &xnappducontentsv1.SnodeModificationRequestIEsValue_IdCause{
			IdCause: IDCause,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeModificationRequestIEsValueIDCause() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeModificationRequestIEsValueIDPdcpchangeIndication(IDPdcpchangeIndication *xnapiesv1.PdcpchangeIndication) (*xnappducontentsv1.SnodeModificationRequestIEsValue, error) {

	item := &xnappducontentsv1.SnodeModificationRequestIEsValue{
		SnodeModificationRequestIes: &xnappducontentsv1.SnodeModificationRequestIEsValue_IdPdcpchangeIndication{
			IdPdcpchangeIndication: IDPdcpchangeIndication,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeModificationRequestIEsValueIDPdcpchangeIndication() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeModificationRequestIEsValueIDSelectedPlmn(IDSelectedPlmn *xnapiesv1.PlmnIdentity) (*xnappducontentsv1.SnodeModificationRequestIEsValue, error) {

	item := &xnappducontentsv1.SnodeModificationRequestIEsValue{
		SnodeModificationRequestIes: &xnappducontentsv1.SnodeModificationRequestIEsValue_IdSelectedPlmn{
			IdSelectedPlmn: IDSelectedPlmn,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeModificationRequestIEsValueIDSelectedPlmn() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeModificationRequestIEsValueIDMobilityRestrictionList(IDMobilityRestrictionList *xnapiesv1.MobilityRestrictionList) (*xnappducontentsv1.SnodeModificationRequestIEsValue, error) {

	item := &xnappducontentsv1.SnodeModificationRequestIEsValue{
		SnodeModificationRequestIes: &xnappducontentsv1.SnodeModificationRequestIEsValue_IdMobilityRestrictionList{
			IdMobilityRestrictionList: IDMobilityRestrictionList,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeModificationRequestIEsValueIDMobilityRestrictionList() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeModificationRequestIEsValueIDScgconfigurationQuery(IDScgconfigurationQuery xnapiesv1.ScgconfigurationQuery) (*xnappducontentsv1.SnodeModificationRequestIEsValue, error) {

	item := &xnappducontentsv1.SnodeModificationRequestIEsValue{
		SnodeModificationRequestIes: &xnappducontentsv1.SnodeModificationRequestIEsValue_IdScgconfigurationQuery{
			IdScgconfigurationQuery: IDScgconfigurationQuery,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeModificationRequestIEsValueIDScgconfigurationQuery() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeModificationRequestIEsValueIDUecontextInfoSnmodRequest(IDUecontextInfoSnmodRequest *xnappducontentsv1.UecontextInfoSNmodRequest) (*xnappducontentsv1.SnodeModificationRequestIEsValue, error) {

	item := &xnappducontentsv1.SnodeModificationRequestIEsValue{
		SnodeModificationRequestIes: &xnappducontentsv1.SnodeModificationRequestIEsValue_IdUecontextInfoSnmodRequest{
			IdUecontextInfoSnmodRequest: IDUecontextInfoSnmodRequest,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeModificationRequestIEsValueIDUecontextInfoSnmodRequest() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeModificationRequestIEsValueIDMnToSnContainer(IDMnToSnContainer string) (*xnappducontentsv1.SnodeModificationRequestIEsValue, error) {

	item := &xnappducontentsv1.SnodeModificationRequestIEsValue{
		SnodeModificationRequestIes: &xnappducontentsv1.SnodeModificationRequestIEsValue_IdMnToSnContainer{
			IdMnToSnContainer: IDMnToSnContainer,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeModificationRequestIEsValueIDMnToSnContainer() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeModificationRequestIEsValueIDRequestedSplitSrb(IDRequestedSplitSrb xnapiesv1.SplitSrbsTypes) (*xnappducontentsv1.SnodeModificationRequestIEsValue, error) {

	item := &xnappducontentsv1.SnodeModificationRequestIEsValue{
		SnodeModificationRequestIes: &xnappducontentsv1.SnodeModificationRequestIEsValue_IdRequestedSplitSrb{
			IdRequestedSplitSrb: IDRequestedSplitSrb,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeModificationRequestIEsValueIDRequestedSplitSrb() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeModificationRequestIEsValueIDRequestedSplitSrbrelease(IDRequestedSplitSrbrelease xnapiesv1.SplitSrbsTypes) (*xnappducontentsv1.SnodeModificationRequestIEsValue, error) {

	item := &xnappducontentsv1.SnodeModificationRequestIEsValue{
		SnodeModificationRequestIes: &xnappducontentsv1.SnodeModificationRequestIEsValue_IdRequestedSplitSrbrelease{
			IdRequestedSplitSrbrelease: IDRequestedSplitSrbrelease,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeModificationRequestIEsValueIDRequestedSplitSrbrelease() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeModificationRequestIEsValueIDDesiredActNotificationLevel(IDDesiredActNotificationLevel xnapiesv1.DesiredActNotificationLevel) (*xnappducontentsv1.SnodeModificationRequestIEsValue, error) {

	item := &xnappducontentsv1.SnodeModificationRequestIEsValue{
		SnodeModificationRequestIes: &xnappducontentsv1.SnodeModificationRequestIEsValue_IdDesiredActNotificationLevel{
			IdDesiredActNotificationLevel: IDDesiredActNotificationLevel,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeModificationRequestIEsValueIDDesiredActNotificationLevel() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeModificationRequestIEsValueIDAdditionalDrbIDs(IDAdditionalDrbIDs *xnapiesv1.DrbList) (*xnappducontentsv1.SnodeModificationRequestIEsValue, error) {

	item := &xnappducontentsv1.SnodeModificationRequestIEsValue{
		SnodeModificationRequestIes: &xnappducontentsv1.SnodeModificationRequestIEsValue_IdAdditionalDrbids{
			IdAdditionalDrbids: IDAdditionalDrbIDs,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeModificationRequestIEsValueIDAdditionalDrbIDs() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeModificationRequestIEsValueIDSNgRannodeMaxIpdataRateUl(IDSNgRannodeMaxIpdataRateUl *xnapiesv1.BitRate) (*xnappducontentsv1.SnodeModificationRequestIEsValue, error) {

	item := &xnappducontentsv1.SnodeModificationRequestIEsValue{
		SnodeModificationRequestIes: &xnappducontentsv1.SnodeModificationRequestIEsValue_IdSNgRannodeMaxIpdataRateUl{
			IdSNgRannodeMaxIpdataRateUl: IDSNgRannodeMaxIpdataRateUl,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeModificationRequestIEsValueIDSNgRannodeMaxIpdataRateUl() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeModificationRequestIEsValueIDSNgRannodeMaxIpdataRateDl(IDSNgRannodeMaxIpdataRateDl *xnapiesv1.BitRate) (*xnappducontentsv1.SnodeModificationRequestIEsValue, error) {

	item := &xnappducontentsv1.SnodeModificationRequestIEsValue{
		SnodeModificationRequestIes: &xnappducontentsv1.SnodeModificationRequestIEsValue_IdSNgRannodeMaxIpdataRateDl{
			IdSNgRannodeMaxIpdataRateDl: IDSNgRannodeMaxIpdataRateDl,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeModificationRequestIEsValueIDSNgRannodeMaxIpdataRateDl() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeModificationRequestIEsValueIDLocationInformationSnreporting(IDLocationInformationSnreporting xnapiesv1.LocationInformationSnreporting) (*xnappducontentsv1.SnodeModificationRequestIEsValue, error) {

	item := &xnappducontentsv1.SnodeModificationRequestIEsValue{
		SnodeModificationRequestIes: &xnappducontentsv1.SnodeModificationRequestIEsValue_IdLocationInformationSnreporting{
			IdLocationInformationSnreporting: IDLocationInformationSnreporting,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeModificationRequestIEsValueIDLocationInformationSnreporting() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeModificationRequestIEsValueIDMrDcResourceCoordinationInfo(IDMrDcResourceCoordinationInfo *xnapiesv1.MrDCResourceCoordinationInfo) (*xnappducontentsv1.SnodeModificationRequestIEsValue, error) {

	item := &xnappducontentsv1.SnodeModificationRequestIEsValue{
		SnodeModificationRequestIes: &xnappducontentsv1.SnodeModificationRequestIEsValue_IdMrDcResourceCoordinationInfo{
			IdMrDcResourceCoordinationInfo: IDMrDcResourceCoordinationInfo,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeModificationRequestIEsValueIDMrDcResourceCoordinationInfo() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeModificationRequestIEsValueIDPcellID(IDPcellID *xnapiesv1.GlobalNgRAncellID) (*xnappducontentsv1.SnodeModificationRequestIEsValue, error) {

	item := &xnappducontentsv1.SnodeModificationRequestIEsValue{
		SnodeModificationRequestIes: &xnappducontentsv1.SnodeModificationRequestIEsValue_IdPcellId{
			IdPcellId: IDPcellID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeModificationRequestIEsValueIDPcellID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeModificationRequestIEsValueIDNeDcTdmPattern(IDNeDcTdmPattern *xnapiesv1.NeDCTDmPattern) (*xnappducontentsv1.SnodeModificationRequestIEsValue, error) {

	item := &xnappducontentsv1.SnodeModificationRequestIEsValue{
		SnodeModificationRequestIes: &xnappducontentsv1.SnodeModificationRequestIEsValue_IdNeDcTdmPattern{
			IdNeDcTdmPattern: IDNeDcTdmPattern,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeModificationRequestIEsValueIDNeDcTdmPattern() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeModificationRequestIEsValueIDRequestedFastMcgrecoveryViaSrb3(IDRequestedFastMcgrecoveryViaSrb3 xnappducontentsv1.RequestedFastMcgrecoveryViaSrb3) (*xnappducontentsv1.SnodeModificationRequestIEsValue, error) {

	item := &xnappducontentsv1.SnodeModificationRequestIEsValue{
		SnodeModificationRequestIes: &xnappducontentsv1.SnodeModificationRequestIEsValue_IdRequestedFastMcgrecoveryViaSrb3{
			IdRequestedFastMcgrecoveryViaSrb3: IDRequestedFastMcgrecoveryViaSrb3,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeModificationRequestIEsValueIDRequestedFastMcgrecoveryViaSrb3() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeModificationRequestIEsValueIDRequestedFastMcgrecoveryViaSrb3Release(IDRequestedFastMcgrecoveryViaSrb3Release xnappducontentsv1.RequestedFastMcgrecoveryViaSrb3Release) (*xnappducontentsv1.SnodeModificationRequestIEsValue, error) {

	item := &xnappducontentsv1.SnodeModificationRequestIEsValue{
		SnodeModificationRequestIes: &xnappducontentsv1.SnodeModificationRequestIEsValue_IdRequestedFastMcgrecoveryViaSrb3Release{
			IdRequestedFastMcgrecoveryViaSrb3Release: IDRequestedFastMcgrecoveryViaSrb3Release,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeModificationRequestIEsValueIDRequestedFastMcgrecoveryViaSrb3Release() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeModificationRequestIEsValueIDSntriggered(IDSntriggered xnapiesv1.Sntriggered) (*xnappducontentsv1.SnodeModificationRequestIEsValue, error) {

	item := &xnappducontentsv1.SnodeModificationRequestIEsValue{
		SnodeModificationRequestIes: &xnappducontentsv1.SnodeModificationRequestIEsValue_IdSntriggered{
			IdSntriggered: IDSntriggered,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeModificationRequestIEsValueIDSntriggered() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeModificationRequestIEsValueIDTargetNodeID(IDTargetNodeID *xnapiesv1.GlobalNgRAnnodeID) (*xnappducontentsv1.SnodeModificationRequestIEsValue, error) {

	item := &xnappducontentsv1.SnodeModificationRequestIEsValue{
		SnodeModificationRequestIes: &xnappducontentsv1.SnodeModificationRequestIEsValue_IdTargetNodeId{
			IdTargetNodeId: IDTargetNodeID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeModificationRequestIEsValueIDTargetNodeID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreatePdusessionsToBeModifiedSNmodRequestItemExtIesExtensionIDSNssai(IDSNssai *xnapiesv1.SNSsai) (*xnappducontentsv1.PdusessionsToBeModifiedSNmodRequestItemExtIesExtension, error) {

	item := &xnappducontentsv1.PdusessionsToBeModifiedSNmodRequestItemExtIesExtension{
		PdusessionsToBeModifiedSnmodRequestItemExtIes: &xnappducontentsv1.PdusessionsToBeModifiedSNmodRequestItemExtIesExtension_IdSNssai{
			IdSNssai: IDSNssai,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePdusessionsToBeModifiedSNmodRequestItemExtIesExtensionIDSNssai() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreatePdusessionsToBeModifiedSNmodRequestItemExtIesExtensionIDPdusessionExpectedUeactivityBehaviour(IDPdusessionExpectedUeactivityBehaviour *xnapiesv1.ExpectedUeactivityBehaviour) (*xnappducontentsv1.PdusessionsToBeModifiedSNmodRequestItemExtIesExtension, error) {

	item := &xnappducontentsv1.PdusessionsToBeModifiedSNmodRequestItemExtIesExtension{
		PdusessionsToBeModifiedSnmodRequestItemExtIes: &xnappducontentsv1.PdusessionsToBeModifiedSNmodRequestItemExtIesExtension_IdPdusessionExpectedUeactivityBehaviour{
			IdPdusessionExpectedUeactivityBehaviour: IDPdusessionExpectedUeactivityBehaviour,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreatePdusessionsToBeModifiedSNmodRequestItemExtIesExtensionIDPdusessionExpectedUeactivityBehaviour() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeModificationRequestAcknowledgeIEsValueIDMNgRannodeUexnApID(IDMNgRannodeUexnApID *xnapiesv1.NgRAnnodeUexnApid) (*xnappducontentsv1.SnodeModificationRequestAcknowledgeIEsValue, error) {

	item := &xnappducontentsv1.SnodeModificationRequestAcknowledgeIEsValue{
		SnodeModificationRequestAcknowledgeIes: &xnappducontentsv1.SnodeModificationRequestAcknowledgeIEsValue_IdMNgRannodeUexnApid{
			IdMNgRannodeUexnApid: IDMNgRannodeUexnApID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeModificationRequestAcknowledgeIEsValueIDMNgRannodeUexnApID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeModificationRequestAcknowledgeIEsValueIDSNgRannodeUexnApID(IDSNgRannodeUexnApID *xnapiesv1.NgRAnnodeUexnApid) (*xnappducontentsv1.SnodeModificationRequestAcknowledgeIEsValue, error) {

	item := &xnappducontentsv1.SnodeModificationRequestAcknowledgeIEsValue{
		SnodeModificationRequestAcknowledgeIes: &xnappducontentsv1.SnodeModificationRequestAcknowledgeIEsValue_IdSNgRannodeUexnApid{
			IdSNgRannodeUexnApid: IDSNgRannodeUexnApID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeModificationRequestAcknowledgeIEsValueIDSNgRannodeUexnApID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeModificationRequestAcknowledgeIEsValueIDPdusessionAdmittedSnmodResponse(IDPdusessionAdmittedSnmodResponse *xnappducontentsv1.PdusessionAdmittedSNmodResponse) (*xnappducontentsv1.SnodeModificationRequestAcknowledgeIEsValue, error) {

	item := &xnappducontentsv1.SnodeModificationRequestAcknowledgeIEsValue{
		SnodeModificationRequestAcknowledgeIes: &xnappducontentsv1.SnodeModificationRequestAcknowledgeIEsValue_IdPdusessionAdmittedSnmodResponse{
			IdPdusessionAdmittedSnmodResponse: IDPdusessionAdmittedSnmodResponse,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeModificationRequestAcknowledgeIEsValueIDPdusessionAdmittedSnmodResponse() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeModificationRequestAcknowledgeIEsValueIDPdusessionNotAdmittedSnmodResponse(IDPdusessionNotAdmittedSnmodResponse *xnappducontentsv1.PdusessionNotAdmittedSNmodResponse) (*xnappducontentsv1.SnodeModificationRequestAcknowledgeIEsValue, error) {

	item := &xnappducontentsv1.SnodeModificationRequestAcknowledgeIEsValue{
		SnodeModificationRequestAcknowledgeIes: &xnappducontentsv1.SnodeModificationRequestAcknowledgeIEsValue_IdPdusessionNotAdmittedSnmodResponse{
			IdPdusessionNotAdmittedSnmodResponse: IDPdusessionNotAdmittedSnmodResponse,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeModificationRequestAcknowledgeIEsValueIDPdusessionNotAdmittedSnmodResponse() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeModificationRequestAcknowledgeIEsValueIDSnToMnContainer(IDSnToMnContainer string) (*xnappducontentsv1.SnodeModificationRequestAcknowledgeIEsValue, error) {

	item := &xnappducontentsv1.SnodeModificationRequestAcknowledgeIEsValue{
		SnodeModificationRequestAcknowledgeIes: &xnappducontentsv1.SnodeModificationRequestAcknowledgeIEsValue_IdSnToMnContainer{
			IdSnToMnContainer: IDSnToMnContainer,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeModificationRequestAcknowledgeIEsValueIDSnToMnContainer() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeModificationRequestAcknowledgeIEsValueIDAdmittedSplitSrb(IDAdmittedSplitSrb xnapiesv1.SplitSrbsTypes) (*xnappducontentsv1.SnodeModificationRequestAcknowledgeIEsValue, error) {

	item := &xnappducontentsv1.SnodeModificationRequestAcknowledgeIEsValue{
		SnodeModificationRequestAcknowledgeIes: &xnappducontentsv1.SnodeModificationRequestAcknowledgeIEsValue_IdAdmittedSplitSrb{
			IdAdmittedSplitSrb: IDAdmittedSplitSrb,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeModificationRequestAcknowledgeIEsValueIDAdmittedSplitSrb() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeModificationRequestAcknowledgeIEsValueIDAdmittedSplitSrbrelease(IDAdmittedSplitSrbrelease xnapiesv1.SplitSrbsTypes) (*xnappducontentsv1.SnodeModificationRequestAcknowledgeIEsValue, error) {

	item := &xnappducontentsv1.SnodeModificationRequestAcknowledgeIEsValue{
		SnodeModificationRequestAcknowledgeIes: &xnappducontentsv1.SnodeModificationRequestAcknowledgeIEsValue_IdAdmittedSplitSrbrelease{
			IdAdmittedSplitSrbrelease: IDAdmittedSplitSrbrelease,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeModificationRequestAcknowledgeIEsValueIDAdmittedSplitSrbrelease() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeModificationRequestAcknowledgeIEsValueIDCriticalityDiagnostics(IDCriticalityDiagnostics *xnapiesv1.CriticalityDiagnostics) (*xnappducontentsv1.SnodeModificationRequestAcknowledgeIEsValue, error) {

	item := &xnappducontentsv1.SnodeModificationRequestAcknowledgeIEsValue{
		SnodeModificationRequestAcknowledgeIes: &xnappducontentsv1.SnodeModificationRequestAcknowledgeIEsValue_IdCriticalityDiagnostics{
			IdCriticalityDiagnostics: IDCriticalityDiagnostics,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeModificationRequestAcknowledgeIEsValueIDCriticalityDiagnostics() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeModificationRequestAcknowledgeIEsValueIDLocationInformationSn(IDLocationInformationSn *xnapiesv1.TargetCGi) (*xnappducontentsv1.SnodeModificationRequestAcknowledgeIEsValue, error) {

	item := &xnappducontentsv1.SnodeModificationRequestAcknowledgeIEsValue{
		SnodeModificationRequestAcknowledgeIes: &xnappducontentsv1.SnodeModificationRequestAcknowledgeIEsValue_IdLocationInformationSn{
			IdLocationInformationSn: IDLocationInformationSn,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeModificationRequestAcknowledgeIEsValueIDLocationInformationSn() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeModificationRequestAcknowledgeIEsValueIDMrDcResourceCoordinationInfo(IDMrDcResourceCoordinationInfo *xnapiesv1.MrDCResourceCoordinationInfo) (*xnappducontentsv1.SnodeModificationRequestAcknowledgeIEsValue, error) {

	item := &xnappducontentsv1.SnodeModificationRequestAcknowledgeIEsValue{
		SnodeModificationRequestAcknowledgeIes: &xnappducontentsv1.SnodeModificationRequestAcknowledgeIEsValue_IdMrDcResourceCoordinationInfo{
			IdMrDcResourceCoordinationInfo: IDMrDcResourceCoordinationInfo,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeModificationRequestAcknowledgeIEsValueIDMrDcResourceCoordinationInfo() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeModificationRequestAcknowledgeIEsValueIDPdusessionDataForwardingSnmodResponse(IDPdusessionDataForwardingSnmodResponse *xnappducontentsv1.PdusessionDataForwardingSNmodResponse) (*xnappducontentsv1.SnodeModificationRequestAcknowledgeIEsValue, error) {

	item := &xnappducontentsv1.SnodeModificationRequestAcknowledgeIEsValue{
		SnodeModificationRequestAcknowledgeIes: &xnappducontentsv1.SnodeModificationRequestAcknowledgeIEsValue_IdPdusessionDataForwardingSnmodResponse{
			IdPdusessionDataForwardingSnmodResponse: IDPdusessionDataForwardingSnmodResponse,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeModificationRequestAcknowledgeIEsValueIDPdusessionDataForwardingSnmodResponse() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeModificationRequestAcknowledgeIEsValueIDRrcconfigIndication(IDRrcconfigIndication xnapiesv1.RrcconfigIndication) (*xnappducontentsv1.SnodeModificationRequestAcknowledgeIEsValue, error) {

	item := &xnappducontentsv1.SnodeModificationRequestAcknowledgeIEsValue{
		SnodeModificationRequestAcknowledgeIes: &xnappducontentsv1.SnodeModificationRequestAcknowledgeIEsValue_IdRrcconfigIndication{
			IdRrcconfigIndication: IDRrcconfigIndication,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeModificationRequestAcknowledgeIEsValueIDRrcconfigIndication() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeModificationRequestAcknowledgeIEsValueIDAvailableFastMcgrecoveryViaSrb3(IDAvailableFastMcgrecoveryViaSrb3 xnappducontentsv1.AvailableFastMcgrecoveryViaSrb3) (*xnappducontentsv1.SnodeModificationRequestAcknowledgeIEsValue, error) {

	item := &xnappducontentsv1.SnodeModificationRequestAcknowledgeIEsValue{
		SnodeModificationRequestAcknowledgeIes: &xnappducontentsv1.SnodeModificationRequestAcknowledgeIEsValue_IdAvailableFastMcgrecoveryViaSrb3{
			IdAvailableFastMcgrecoveryViaSrb3: IDAvailableFastMcgrecoveryViaSrb3,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeModificationRequestAcknowledgeIEsValueIDAvailableFastMcgrecoveryViaSrb3() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeModificationRequestAcknowledgeIEsValueIDReleaseFastMcgrecoveryViaSrb3(IDReleaseFastMcgrecoveryViaSrb3 xnappducontentsv1.ReleaseFastMcgrecoveryViaSrb3) (*xnappducontentsv1.SnodeModificationRequestAcknowledgeIEsValue, error) {

	item := &xnappducontentsv1.SnodeModificationRequestAcknowledgeIEsValue{
		SnodeModificationRequestAcknowledgeIes: &xnappducontentsv1.SnodeModificationRequestAcknowledgeIEsValue_IdReleaseFastMcgrecoveryViaSrb3{
			IdReleaseFastMcgrecoveryViaSrb3: IDReleaseFastMcgrecoveryViaSrb3,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeModificationRequestAcknowledgeIEsValueIDReleaseFastMcgrecoveryViaSrb3() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeModificationRequestAcknowledgeIEsValueIDDirectForwardingPathAvailability(IDDirectForwardingPathAvailability xnapiesv1.DirectForwardingPathAvailability) (*xnappducontentsv1.SnodeModificationRequestAcknowledgeIEsValue, error) {

	item := &xnappducontentsv1.SnodeModificationRequestAcknowledgeIEsValue{
		SnodeModificationRequestAcknowledgeIes: &xnappducontentsv1.SnodeModificationRequestAcknowledgeIEsValue_IdDirectForwardingPathAvailability{
			IdDirectForwardingPathAvailability: IDDirectForwardingPathAvailability,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeModificationRequestAcknowledgeIEsValueIDDirectForwardingPathAvailability() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeModificationRequestRejectIEsValueIDMNgRannodeUexnApID(IDMNgRannodeUexnApID *xnapiesv1.NgRAnnodeUexnApid) (*xnappducontentsv1.SnodeModificationRequestRejectIEsValue, error) {

	item := &xnappducontentsv1.SnodeModificationRequestRejectIEsValue{
		SnodeModificationRequestRejectIes: &xnappducontentsv1.SnodeModificationRequestRejectIEsValue_IdMNgRannodeUexnApid{
			IdMNgRannodeUexnApid: IDMNgRannodeUexnApID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeModificationRequestRejectIEsValueIDMNgRannodeUexnApID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeModificationRequestRejectIEsValueIDSNgRannodeUexnApID(IDSNgRannodeUexnApID *xnapiesv1.NgRAnnodeUexnApid) (*xnappducontentsv1.SnodeModificationRequestRejectIEsValue, error) {

	item := &xnappducontentsv1.SnodeModificationRequestRejectIEsValue{
		SnodeModificationRequestRejectIes: &xnappducontentsv1.SnodeModificationRequestRejectIEsValue_IdSNgRannodeUexnApid{
			IdSNgRannodeUexnApid: IDSNgRannodeUexnApID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeModificationRequestRejectIEsValueIDSNgRannodeUexnApID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeModificationRequestRejectIEsValueIDCause(IDCause *xnapiesv1.Cause) (*xnappducontentsv1.SnodeModificationRequestRejectIEsValue, error) {

	item := &xnappducontentsv1.SnodeModificationRequestRejectIEsValue{
		SnodeModificationRequestRejectIes: &xnappducontentsv1.SnodeModificationRequestRejectIEsValue_IdCause{
			IdCause: IDCause,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeModificationRequestRejectIEsValueIDCause() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeModificationRequestRejectIEsValueIDCriticalityDiagnostics(IDCriticalityDiagnostics *xnapiesv1.CriticalityDiagnostics) (*xnappducontentsv1.SnodeModificationRequestRejectIEsValue, error) {

	item := &xnappducontentsv1.SnodeModificationRequestRejectIEsValue{
		SnodeModificationRequestRejectIes: &xnappducontentsv1.SnodeModificationRequestRejectIEsValue_IdCriticalityDiagnostics{
			IdCriticalityDiagnostics: IDCriticalityDiagnostics,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeModificationRequestRejectIEsValueIDCriticalityDiagnostics() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeModificationRequiredIEsValueIDMNgRannodeUexnApID(IDMNgRannodeUexnApID *xnapiesv1.NgRAnnodeUexnApid) (*xnappducontentsv1.SnodeModificationRequiredIEsValue, error) {

	item := &xnappducontentsv1.SnodeModificationRequiredIEsValue{
		SnodeModificationRequiredIes: &xnappducontentsv1.SnodeModificationRequiredIEsValue_IdMNgRannodeUexnApid{
			IdMNgRannodeUexnApid: IDMNgRannodeUexnApID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeModificationRequiredIEsValueIDMNgRannodeUexnApID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeModificationRequiredIEsValueIDSNgRannodeUexnApID(IDSNgRannodeUexnApID *xnapiesv1.NgRAnnodeUexnApid) (*xnappducontentsv1.SnodeModificationRequiredIEsValue, error) {

	item := &xnappducontentsv1.SnodeModificationRequiredIEsValue{
		SnodeModificationRequiredIes: &xnappducontentsv1.SnodeModificationRequiredIEsValue_IdSNgRannodeUexnApid{
			IdSNgRannodeUexnApid: IDSNgRannodeUexnApID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeModificationRequiredIEsValueIDSNgRannodeUexnApID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeModificationRequiredIEsValueIDCause(IDCause *xnapiesv1.Cause) (*xnappducontentsv1.SnodeModificationRequiredIEsValue, error) {

	item := &xnappducontentsv1.SnodeModificationRequiredIEsValue{
		SnodeModificationRequiredIes: &xnappducontentsv1.SnodeModificationRequiredIEsValue_IdCause{
			IdCause: IDCause,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeModificationRequiredIEsValueIDCause() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeModificationRequiredIEsValueIDPdcpchangeIndication(IDPdcpchangeIndication *xnapiesv1.PdcpchangeIndication) (*xnappducontentsv1.SnodeModificationRequiredIEsValue, error) {

	item := &xnappducontentsv1.SnodeModificationRequiredIEsValue{
		SnodeModificationRequiredIes: &xnappducontentsv1.SnodeModificationRequiredIEsValue_IdPdcpchangeIndication{
			IdPdcpchangeIndication: IDPdcpchangeIndication,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeModificationRequiredIEsValueIDPdcpchangeIndication() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeModificationRequiredIEsValueIDPdusessionToBeModifiedSnmodRequired(IDPdusessionToBeModifiedSnmodRequired *xnappducontentsv1.PdusessionToBeModifiedSnmodRequired) (*xnappducontentsv1.SnodeModificationRequiredIEsValue, error) {

	item := &xnappducontentsv1.SnodeModificationRequiredIEsValue{
		SnodeModificationRequiredIes: &xnappducontentsv1.SnodeModificationRequiredIEsValue_IdPdusessionToBeModifiedSnmodRequired{
			IdPdusessionToBeModifiedSnmodRequired: IDPdusessionToBeModifiedSnmodRequired,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeModificationRequiredIEsValueIDPdusessionToBeModifiedSnmodRequired() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeModificationRequiredIEsValueIDPdusessionToBeReleasedSnmodRequired(IDPdusessionToBeReleasedSnmodRequired *xnappducontentsv1.PdusessionToBeReleasedSnmodRequired) (*xnappducontentsv1.SnodeModificationRequiredIEsValue, error) {

	item := &xnappducontentsv1.SnodeModificationRequiredIEsValue{
		SnodeModificationRequiredIes: &xnappducontentsv1.SnodeModificationRequiredIEsValue_IdPdusessionToBeReleasedSnmodRequired{
			IdPdusessionToBeReleasedSnmodRequired: IDPdusessionToBeReleasedSnmodRequired,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeModificationRequiredIEsValueIDPdusessionToBeReleasedSnmodRequired() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeModificationRequiredIEsValueIDSnToMnContainer(IDSnToMnContainer string) (*xnappducontentsv1.SnodeModificationRequiredIEsValue, error) {

	item := &xnappducontentsv1.SnodeModificationRequiredIEsValue{
		SnodeModificationRequiredIes: &xnappducontentsv1.SnodeModificationRequiredIEsValue_IdSnToMnContainer{
			IdSnToMnContainer: IDSnToMnContainer,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeModificationRequiredIEsValueIDSnToMnContainer() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeModificationRequiredIEsValueIDSpareDrbIDs(IDSpareDrbIDs *xnapiesv1.DrbList) (*xnappducontentsv1.SnodeModificationRequiredIEsValue, error) {

	item := &xnappducontentsv1.SnodeModificationRequiredIEsValue{
		SnodeModificationRequiredIes: &xnappducontentsv1.SnodeModificationRequiredIEsValue_IdSpareDrbids{
			IdSpareDrbids: IDSpareDrbIDs,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeModificationRequiredIEsValueIDSpareDrbIDs() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeModificationRequiredIEsValueIDRequiredNumberOfDrbIDs(IDRequiredNumberOfDrbIDs *xnapiesv1.DrbNumber) (*xnappducontentsv1.SnodeModificationRequiredIEsValue, error) {

	item := &xnappducontentsv1.SnodeModificationRequiredIEsValue{
		SnodeModificationRequiredIes: &xnappducontentsv1.SnodeModificationRequiredIEsValue_IdRequiredNumberOfDrbids{
			IdRequiredNumberOfDrbids: IDRequiredNumberOfDrbIDs,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeModificationRequiredIEsValueIDRequiredNumberOfDrbIDs() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeModificationRequiredIEsValueIDLocationInformationSn(IDLocationInformationSn *xnapiesv1.TargetCGi) (*xnappducontentsv1.SnodeModificationRequiredIEsValue, error) {

	item := &xnappducontentsv1.SnodeModificationRequiredIEsValue{
		SnodeModificationRequiredIes: &xnappducontentsv1.SnodeModificationRequiredIEsValue_IdLocationInformationSn{
			IdLocationInformationSn: IDLocationInformationSn,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeModificationRequiredIEsValueIDLocationInformationSn() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeModificationRequiredIEsValueIDMrDcResourceCoordinationInfo(IDMrDcResourceCoordinationInfo *xnapiesv1.MrDCResourceCoordinationInfo) (*xnappducontentsv1.SnodeModificationRequiredIEsValue, error) {

	item := &xnappducontentsv1.SnodeModificationRequiredIEsValue{
		SnodeModificationRequiredIes: &xnappducontentsv1.SnodeModificationRequiredIEsValue_IdMrDcResourceCoordinationInfo{
			IdMrDcResourceCoordinationInfo: IDMrDcResourceCoordinationInfo,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeModificationRequiredIEsValueIDMrDcResourceCoordinationInfo() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeModificationRequiredIEsValueIDRrcconfigIndication(IDRrcconfigIndication xnapiesv1.RrcconfigIndication) (*xnappducontentsv1.SnodeModificationRequiredIEsValue, error) {

	item := &xnappducontentsv1.SnodeModificationRequiredIEsValue{
		SnodeModificationRequiredIes: &xnappducontentsv1.SnodeModificationRequiredIEsValue_IdRrcconfigIndication{
			IdRrcconfigIndication: IDRrcconfigIndication,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeModificationRequiredIEsValueIDRrcconfigIndication() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeModificationRequiredIEsValueIDAvailableFastMcgrecoveryViaSrb3(IDAvailableFastMcgrecoveryViaSrb3 xnappducontentsv1.AvailableFastMcgrecoveryViaSrb3) (*xnappducontentsv1.SnodeModificationRequiredIEsValue, error) {

	item := &xnappducontentsv1.SnodeModificationRequiredIEsValue{
		SnodeModificationRequiredIes: &xnappducontentsv1.SnodeModificationRequiredIEsValue_IdAvailableFastMcgrecoveryViaSrb3{
			IdAvailableFastMcgrecoveryViaSrb3: IDAvailableFastMcgrecoveryViaSrb3,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeModificationRequiredIEsValueIDAvailableFastMcgrecoveryViaSrb3() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeModificationRequiredIEsValueIDReleaseFastMcgrecoveryViaSrb3(IDReleaseFastMcgrecoveryViaSrb3 xnappducontentsv1.ReleaseFastMcgrecoveryViaSrb3) (*xnappducontentsv1.SnodeModificationRequiredIEsValue, error) {

	item := &xnappducontentsv1.SnodeModificationRequiredIEsValue{
		SnodeModificationRequiredIes: &xnappducontentsv1.SnodeModificationRequiredIEsValue_IdReleaseFastMcgrecoveryViaSrb3{
			IdReleaseFastMcgrecoveryViaSrb3: IDReleaseFastMcgrecoveryViaSrb3,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeModificationRequiredIEsValueIDReleaseFastMcgrecoveryViaSrb3() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeModificationRequiredIEsValueIDScgindicator(IDScgindicator xnapiesv1.Scgindicator) (*xnappducontentsv1.SnodeModificationRequiredIEsValue, error) {

	item := &xnappducontentsv1.SnodeModificationRequiredIEsValue{
		SnodeModificationRequiredIes: &xnappducontentsv1.SnodeModificationRequiredIEsValue_IdScgindicator{
			IdScgindicator: IDScgindicator,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeModificationRequiredIEsValueIDScgindicator() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeModificationConfirmIEsValueIDMNgRannodeUexnApID(IDMNgRannodeUexnApID *xnapiesv1.NgRAnnodeUexnApid) (*xnappducontentsv1.SnodeModificationConfirmIEsValue, error) {

	item := &xnappducontentsv1.SnodeModificationConfirmIEsValue{
		SnodeModificationConfirmIes: &xnappducontentsv1.SnodeModificationConfirmIEsValue_IdMNgRannodeUexnApid{
			IdMNgRannodeUexnApid: IDMNgRannodeUexnApID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeModificationConfirmIEsValueIDMNgRannodeUexnApID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeModificationConfirmIEsValueIDSNgRannodeUexnApID(IDSNgRannodeUexnApID *xnapiesv1.NgRAnnodeUexnApid) (*xnappducontentsv1.SnodeModificationConfirmIEsValue, error) {

	item := &xnappducontentsv1.SnodeModificationConfirmIEsValue{
		SnodeModificationConfirmIes: &xnappducontentsv1.SnodeModificationConfirmIEsValue_IdSNgRannodeUexnApid{
			IdSNgRannodeUexnApid: IDSNgRannodeUexnApID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeModificationConfirmIEsValueIDSNgRannodeUexnApID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeModificationConfirmIEsValueIDPdusessionAdmittedModSnmodConfirm(IDPdusessionAdmittedModSnmodConfirm *xnappducontentsv1.PdusessionAdmittedModSnmodConfirm) (*xnappducontentsv1.SnodeModificationConfirmIEsValue, error) {

	item := &xnappducontentsv1.SnodeModificationConfirmIEsValue{
		SnodeModificationConfirmIes: &xnappducontentsv1.SnodeModificationConfirmIEsValue_IdPdusessionAdmittedModSnmodConfirm{
			IdPdusessionAdmittedModSnmodConfirm: IDPdusessionAdmittedModSnmodConfirm,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeModificationConfirmIEsValueIDPdusessionAdmittedModSnmodConfirm() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeModificationConfirmIEsValueIDPdusessionReleasedSnmodConfirm(IDPdusessionReleasedSnmodConfirm *xnappducontentsv1.PdusessionReleasedSnmodConfirm) (*xnappducontentsv1.SnodeModificationConfirmIEsValue, error) {

	item := &xnappducontentsv1.SnodeModificationConfirmIEsValue{
		SnodeModificationConfirmIes: &xnappducontentsv1.SnodeModificationConfirmIEsValue_IdPdusessionReleasedSnmodConfirm{
			IdPdusessionReleasedSnmodConfirm: IDPdusessionReleasedSnmodConfirm,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeModificationConfirmIEsValueIDPdusessionReleasedSnmodConfirm() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeModificationConfirmIEsValueIDMnToSnContainer(IDMnToSnContainer string) (*xnappducontentsv1.SnodeModificationConfirmIEsValue, error) {

	item := &xnappducontentsv1.SnodeModificationConfirmIEsValue{
		SnodeModificationConfirmIes: &xnappducontentsv1.SnodeModificationConfirmIEsValue_IdMnToSnContainer{
			IdMnToSnContainer: IDMnToSnContainer,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeModificationConfirmIEsValueIDMnToSnContainer() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeModificationConfirmIEsValueIDAdditionalDrbIDs(IDAdditionalDrbIDs *xnapiesv1.DrbList) (*xnappducontentsv1.SnodeModificationConfirmIEsValue, error) {

	item := &xnappducontentsv1.SnodeModificationConfirmIEsValue{
		SnodeModificationConfirmIes: &xnappducontentsv1.SnodeModificationConfirmIEsValue_IdAdditionalDrbids{
			IdAdditionalDrbids: IDAdditionalDrbIDs,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeModificationConfirmIEsValueIDAdditionalDrbIDs() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeModificationConfirmIEsValueIDCriticalityDiagnostics(IDCriticalityDiagnostics *xnapiesv1.CriticalityDiagnostics) (*xnappducontentsv1.SnodeModificationConfirmIEsValue, error) {

	item := &xnappducontentsv1.SnodeModificationConfirmIEsValue{
		SnodeModificationConfirmIes: &xnappducontentsv1.SnodeModificationConfirmIEsValue_IdCriticalityDiagnostics{
			IdCriticalityDiagnostics: IDCriticalityDiagnostics,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeModificationConfirmIEsValueIDCriticalityDiagnostics() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeModificationConfirmIEsValueIDMrDcResourceCoordinationInfo(IDMrDcResourceCoordinationInfo *xnapiesv1.MrDCResourceCoordinationInfo) (*xnappducontentsv1.SnodeModificationConfirmIEsValue, error) {

	item := &xnappducontentsv1.SnodeModificationConfirmIEsValue{
		SnodeModificationConfirmIes: &xnappducontentsv1.SnodeModificationConfirmIEsValue_IdMrDcResourceCoordinationInfo{
			IdMrDcResourceCoordinationInfo: IDMrDcResourceCoordinationInfo,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeModificationConfirmIEsValueIDMrDcResourceCoordinationInfo() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeModificationRefuseIEsValueIDMNgRannodeUexnApID(IDMNgRannodeUexnApID *xnapiesv1.NgRAnnodeUexnApid) (*xnappducontentsv1.SnodeModificationRefuseIEsValue, error) {

	item := &xnappducontentsv1.SnodeModificationRefuseIEsValue{
		SnodeModificationRefuseIes: &xnappducontentsv1.SnodeModificationRefuseIEsValue_IdMNgRannodeUexnApid{
			IdMNgRannodeUexnApid: IDMNgRannodeUexnApID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeModificationRefuseIEsValueIDMNgRannodeUexnApID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeModificationRefuseIEsValueIDSNgRannodeUexnApID(IDSNgRannodeUexnApID *xnapiesv1.NgRAnnodeUexnApid) (*xnappducontentsv1.SnodeModificationRefuseIEsValue, error) {

	item := &xnappducontentsv1.SnodeModificationRefuseIEsValue{
		SnodeModificationRefuseIes: &xnappducontentsv1.SnodeModificationRefuseIEsValue_IdSNgRannodeUexnApid{
			IdSNgRannodeUexnApid: IDSNgRannodeUexnApID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeModificationRefuseIEsValueIDSNgRannodeUexnApID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeModificationRefuseIEsValueIDCause(IDCause *xnapiesv1.Cause) (*xnappducontentsv1.SnodeModificationRefuseIEsValue, error) {

	item := &xnappducontentsv1.SnodeModificationRefuseIEsValue{
		SnodeModificationRefuseIes: &xnappducontentsv1.SnodeModificationRefuseIEsValue_IdCause{
			IdCause: IDCause,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeModificationRefuseIEsValueIDCause() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeModificationRefuseIEsValueIDMnToSnContainer(IDMnToSnContainer string) (*xnappducontentsv1.SnodeModificationRefuseIEsValue, error) {

	item := &xnappducontentsv1.SnodeModificationRefuseIEsValue{
		SnodeModificationRefuseIes: &xnappducontentsv1.SnodeModificationRefuseIEsValue_IdMnToSnContainer{
			IdMnToSnContainer: IDMnToSnContainer,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeModificationRefuseIEsValueIDMnToSnContainer() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeModificationRefuseIEsValueIDCriticalityDiagnostics(IDCriticalityDiagnostics *xnapiesv1.CriticalityDiagnostics) (*xnappducontentsv1.SnodeModificationRefuseIEsValue, error) {

	item := &xnappducontentsv1.SnodeModificationRefuseIEsValue{
		SnodeModificationRefuseIes: &xnappducontentsv1.SnodeModificationRefuseIEsValue_IdCriticalityDiagnostics{
			IdCriticalityDiagnostics: IDCriticalityDiagnostics,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeModificationRefuseIEsValueIDCriticalityDiagnostics() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeReleaseRequestIEsValueIDMNgRannodeUexnApID(IDMNgRannodeUexnApID *xnapiesv1.NgRAnnodeUexnApid) (*xnappducontentsv1.SnodeReleaseRequestIEsValue, error) {

	item := &xnappducontentsv1.SnodeReleaseRequestIEsValue{
		SnodeReleaseRequestIes: &xnappducontentsv1.SnodeReleaseRequestIEsValue_IdMNgRannodeUexnApid{
			IdMNgRannodeUexnApid: IDMNgRannodeUexnApID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeReleaseRequestIEsValueIDMNgRannodeUexnApID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeReleaseRequestIEsValueIDSNgRannodeUexnApID(IDSNgRannodeUexnApID *xnapiesv1.NgRAnnodeUexnApid) (*xnappducontentsv1.SnodeReleaseRequestIEsValue, error) {

	item := &xnappducontentsv1.SnodeReleaseRequestIEsValue{
		SnodeReleaseRequestIes: &xnappducontentsv1.SnodeReleaseRequestIEsValue_IdSNgRannodeUexnApid{
			IdSNgRannodeUexnApid: IDSNgRannodeUexnApID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeReleaseRequestIEsValueIDSNgRannodeUexnApID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeReleaseRequestIEsValueIDCause(IDCause *xnapiesv1.Cause) (*xnappducontentsv1.SnodeReleaseRequestIEsValue, error) {

	item := &xnappducontentsv1.SnodeReleaseRequestIEsValue{
		SnodeReleaseRequestIes: &xnappducontentsv1.SnodeReleaseRequestIEsValue_IdCause{
			IdCause: IDCause,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeReleaseRequestIEsValueIDCause() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeReleaseRequestIEsValueIDPdusessionToBeReleasedRelReq(IDPdusessionToBeReleasedRelReq *xnapiesv1.PdusessionListwithCause) (*xnappducontentsv1.SnodeReleaseRequestIEsValue, error) {

	item := &xnappducontentsv1.SnodeReleaseRequestIEsValue{
		SnodeReleaseRequestIes: &xnappducontentsv1.SnodeReleaseRequestIEsValue_IdPdusessionToBeReleasedRelReq{
			IdPdusessionToBeReleasedRelReq: IDPdusessionToBeReleasedRelReq,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeReleaseRequestIEsValueIDPdusessionToBeReleasedRelReq() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeReleaseRequestIEsValueIDUecontextKeptIndicator(IDUecontextKeptIndicator xnapiesv1.UecontextKeptIndicator) (*xnappducontentsv1.SnodeReleaseRequestIEsValue, error) {

	item := &xnappducontentsv1.SnodeReleaseRequestIEsValue{
		SnodeReleaseRequestIes: &xnappducontentsv1.SnodeReleaseRequestIEsValue_IdUecontextKeptIndicator{
			IdUecontextKeptIndicator: IDUecontextKeptIndicator,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeReleaseRequestIEsValueIDUecontextKeptIndicator() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeReleaseRequestIEsValueIDMnToSnContainer(IDMnToSnContainer string) (*xnappducontentsv1.SnodeReleaseRequestIEsValue, error) {

	item := &xnappducontentsv1.SnodeReleaseRequestIEsValue{
		SnodeReleaseRequestIes: &xnappducontentsv1.SnodeReleaseRequestIEsValue_IdMnToSnContainer{
			IdMnToSnContainer: IDMnToSnContainer,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeReleaseRequestIEsValueIDMnToSnContainer() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeReleaseRequestIEsValueIDDrbsTransferredToMn(IDDrbsTransferredToMn *xnapiesv1.DrbList) (*xnappducontentsv1.SnodeReleaseRequestIEsValue, error) {

	item := &xnappducontentsv1.SnodeReleaseRequestIEsValue{
		SnodeReleaseRequestIes: &xnappducontentsv1.SnodeReleaseRequestIEsValue_IdDrbsTransferredToMn{
			IdDrbsTransferredToMn: IDDrbsTransferredToMn,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeReleaseRequestIEsValueIDDrbsTransferredToMn() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeReleaseRequestAcknowledgeIEsValueIDMNgRannodeUexnApID(IDMNgRannodeUexnApID *xnapiesv1.NgRAnnodeUexnApid) (*xnappducontentsv1.SnodeReleaseRequestAcknowledgeIEsValue, error) {

	item := &xnappducontentsv1.SnodeReleaseRequestAcknowledgeIEsValue{
		SnodeReleaseRequestAcknowledgeIes: &xnappducontentsv1.SnodeReleaseRequestAcknowledgeIEsValue_IdMNgRannodeUexnApid{
			IdMNgRannodeUexnApid: IDMNgRannodeUexnApID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeReleaseRequestAcknowledgeIEsValueIDMNgRannodeUexnApID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeReleaseRequestAcknowledgeIEsValueIDSNgRannodeUexnApID(IDSNgRannodeUexnApID *xnapiesv1.NgRAnnodeUexnApid) (*xnappducontentsv1.SnodeReleaseRequestAcknowledgeIEsValue, error) {

	item := &xnappducontentsv1.SnodeReleaseRequestAcknowledgeIEsValue{
		SnodeReleaseRequestAcknowledgeIes: &xnappducontentsv1.SnodeReleaseRequestAcknowledgeIEsValue_IdSNgRannodeUexnApid{
			IdSNgRannodeUexnApid: IDSNgRannodeUexnApID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeReleaseRequestAcknowledgeIEsValueIDSNgRannodeUexnApID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeReleaseRequestAcknowledgeIEsValueIDPdusessionToBeReleasedRelReqAck(IDPdusessionToBeReleasedRelReqAck *xnappducontentsv1.PdusessionToBeReleasedListRelReqAck) (*xnappducontentsv1.SnodeReleaseRequestAcknowledgeIEsValue, error) {

	item := &xnappducontentsv1.SnodeReleaseRequestAcknowledgeIEsValue{
		SnodeReleaseRequestAcknowledgeIes: &xnappducontentsv1.SnodeReleaseRequestAcknowledgeIEsValue_IdPdusessionToBeReleasedRelReqAck{
			IdPdusessionToBeReleasedRelReqAck: IDPdusessionToBeReleasedRelReqAck,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeReleaseRequestAcknowledgeIEsValueIDPdusessionToBeReleasedRelReqAck() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeReleaseRequestAcknowledgeIEsValueIDCriticalityDiagnostics(IDCriticalityDiagnostics *xnapiesv1.CriticalityDiagnostics) (*xnappducontentsv1.SnodeReleaseRequestAcknowledgeIEsValue, error) {

	item := &xnappducontentsv1.SnodeReleaseRequestAcknowledgeIEsValue{
		SnodeReleaseRequestAcknowledgeIes: &xnappducontentsv1.SnodeReleaseRequestAcknowledgeIEsValue_IdCriticalityDiagnostics{
			IdCriticalityDiagnostics: IDCriticalityDiagnostics,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeReleaseRequestAcknowledgeIEsValueIDCriticalityDiagnostics() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeReleaseRejectIEsValueIDMNgRannodeUexnApID(IDMNgRannodeUexnApID *xnapiesv1.NgRAnnodeUexnApid) (*xnappducontentsv1.SnodeReleaseRejectIEsValue, error) {

	item := &xnappducontentsv1.SnodeReleaseRejectIEsValue{
		SnodeReleaseRejectIes: &xnappducontentsv1.SnodeReleaseRejectIEsValue_IdMNgRannodeUexnApid{
			IdMNgRannodeUexnApid: IDMNgRannodeUexnApID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeReleaseRejectIEsValueIDMNgRannodeUexnApID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeReleaseRejectIEsValueIDSNgRannodeUexnApID(IDSNgRannodeUexnApID *xnapiesv1.NgRAnnodeUexnApid) (*xnappducontentsv1.SnodeReleaseRejectIEsValue, error) {

	item := &xnappducontentsv1.SnodeReleaseRejectIEsValue{
		SnodeReleaseRejectIes: &xnappducontentsv1.SnodeReleaseRejectIEsValue_IdSNgRannodeUexnApid{
			IdSNgRannodeUexnApid: IDSNgRannodeUexnApID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeReleaseRejectIEsValueIDSNgRannodeUexnApID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeReleaseRejectIEsValueIDCause(IDCause *xnapiesv1.Cause) (*xnappducontentsv1.SnodeReleaseRejectIEsValue, error) {

	item := &xnappducontentsv1.SnodeReleaseRejectIEsValue{
		SnodeReleaseRejectIes: &xnappducontentsv1.SnodeReleaseRejectIEsValue_IdCause{
			IdCause: IDCause,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeReleaseRejectIEsValueIDCause() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeReleaseRejectIEsValueIDCriticalityDiagnostics(IDCriticalityDiagnostics *xnapiesv1.CriticalityDiagnostics) (*xnappducontentsv1.SnodeReleaseRejectIEsValue, error) {

	item := &xnappducontentsv1.SnodeReleaseRejectIEsValue{
		SnodeReleaseRejectIes: &xnappducontentsv1.SnodeReleaseRejectIEsValue_IdCriticalityDiagnostics{
			IdCriticalityDiagnostics: IDCriticalityDiagnostics,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeReleaseRejectIEsValueIDCriticalityDiagnostics() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeReleaseRequiredIEsValueIDMNgRannodeUexnApID(IDMNgRannodeUexnApID *xnapiesv1.NgRAnnodeUexnApid) (*xnappducontentsv1.SnodeReleaseRequiredIEsValue, error) {

	item := &xnappducontentsv1.SnodeReleaseRequiredIEsValue{
		SnodeReleaseRequiredIes: &xnappducontentsv1.SnodeReleaseRequiredIEsValue_IdMNgRannodeUexnApid{
			IdMNgRannodeUexnApid: IDMNgRannodeUexnApID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeReleaseRequiredIEsValueIDMNgRannodeUexnApID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeReleaseRequiredIEsValueIDSNgRannodeUexnApID(IDSNgRannodeUexnApID *xnapiesv1.NgRAnnodeUexnApid) (*xnappducontentsv1.SnodeReleaseRequiredIEsValue, error) {

	item := &xnappducontentsv1.SnodeReleaseRequiredIEsValue{
		SnodeReleaseRequiredIes: &xnappducontentsv1.SnodeReleaseRequiredIEsValue_IdSNgRannodeUexnApid{
			IdSNgRannodeUexnApid: IDSNgRannodeUexnApID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeReleaseRequiredIEsValueIDSNgRannodeUexnApID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeReleaseRequiredIEsValueIDPdusessionToBeReleasedListRelRqd(IDPdusessionToBeReleasedListRelRqd *xnappducontentsv1.PdusessionToBeReleasedListRelRqd) (*xnappducontentsv1.SnodeReleaseRequiredIEsValue, error) {

	item := &xnappducontentsv1.SnodeReleaseRequiredIEsValue{
		SnodeReleaseRequiredIes: &xnappducontentsv1.SnodeReleaseRequiredIEsValue_IdPdusessionToBeReleasedListRelRqd{
			IdPdusessionToBeReleasedListRelRqd: IDPdusessionToBeReleasedListRelRqd,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeReleaseRequiredIEsValueIDPdusessionToBeReleasedListRelRqd() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeReleaseRequiredIEsValueIDCause(IDCause *xnapiesv1.Cause) (*xnappducontentsv1.SnodeReleaseRequiredIEsValue, error) {

	item := &xnappducontentsv1.SnodeReleaseRequiredIEsValue{
		SnodeReleaseRequiredIes: &xnappducontentsv1.SnodeReleaseRequiredIEsValue_IdCause{
			IdCause: IDCause,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeReleaseRequiredIEsValueIDCause() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeReleaseRequiredIEsValueIDSnToMnContainer(IDSnToMnContainer string) (*xnappducontentsv1.SnodeReleaseRequiredIEsValue, error) {

	item := &xnappducontentsv1.SnodeReleaseRequiredIEsValue{
		SnodeReleaseRequiredIes: &xnappducontentsv1.SnodeReleaseRequiredIEsValue_IdSnToMnContainer{
			IdSnToMnContainer: IDSnToMnContainer,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeReleaseRequiredIEsValueIDSnToMnContainer() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeReleaseConfirmIEsValueIDMNgRannodeUexnApID(IDMNgRannodeUexnApID *xnapiesv1.NgRAnnodeUexnApid) (*xnappducontentsv1.SnodeReleaseConfirmIEsValue, error) {

	item := &xnappducontentsv1.SnodeReleaseConfirmIEsValue{
		SnodeReleaseConfirmIes: &xnappducontentsv1.SnodeReleaseConfirmIEsValue_IdMNgRannodeUexnApid{
			IdMNgRannodeUexnApid: IDMNgRannodeUexnApID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeReleaseConfirmIEsValueIDMNgRannodeUexnApID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeReleaseConfirmIEsValueIDSNgRannodeUexnApID(IDSNgRannodeUexnApID *xnapiesv1.NgRAnnodeUexnApid) (*xnappducontentsv1.SnodeReleaseConfirmIEsValue, error) {

	item := &xnappducontentsv1.SnodeReleaseConfirmIEsValue{
		SnodeReleaseConfirmIes: &xnappducontentsv1.SnodeReleaseConfirmIEsValue_IdSNgRannodeUexnApid{
			IdSNgRannodeUexnApid: IDSNgRannodeUexnApID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeReleaseConfirmIEsValueIDSNgRannodeUexnApID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeReleaseConfirmIEsValueIDPdusessionReleasedListRelConf(IDPdusessionReleasedListRelConf *xnappducontentsv1.PdusessionReleasedListRelConf) (*xnappducontentsv1.SnodeReleaseConfirmIEsValue, error) {

	item := &xnappducontentsv1.SnodeReleaseConfirmIEsValue{
		SnodeReleaseConfirmIes: &xnappducontentsv1.SnodeReleaseConfirmIEsValue_IdPdusessionReleasedListRelConf{
			IdPdusessionReleasedListRelConf: IDPdusessionReleasedListRelConf,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeReleaseConfirmIEsValueIDPdusessionReleasedListRelConf() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeReleaseConfirmIEsValueIDCriticalityDiagnostics(IDCriticalityDiagnostics *xnapiesv1.CriticalityDiagnostics) (*xnappducontentsv1.SnodeReleaseConfirmIEsValue, error) {

	item := &xnappducontentsv1.SnodeReleaseConfirmIEsValue{
		SnodeReleaseConfirmIes: &xnappducontentsv1.SnodeReleaseConfirmIEsValue_IdCriticalityDiagnostics{
			IdCriticalityDiagnostics: IDCriticalityDiagnostics,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeReleaseConfirmIEsValueIDCriticalityDiagnostics() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeCounterCheckRequestIEsValueIDMNgRannodeUexnApID(IDMNgRannodeUexnApID *xnapiesv1.NgRAnnodeUexnApid) (*xnappducontentsv1.SnodeCounterCheckRequestIEsValue, error) {

	item := &xnappducontentsv1.SnodeCounterCheckRequestIEsValue{
		SnodeCounterCheckRequestIes: &xnappducontentsv1.SnodeCounterCheckRequestIEsValue_IdMNgRannodeUexnApid{
			IdMNgRannodeUexnApid: IDMNgRannodeUexnApID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeCounterCheckRequestIEsValueIDMNgRannodeUexnApID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeCounterCheckRequestIEsValueIDSNgRannodeUexnApID(IDSNgRannodeUexnApID *xnapiesv1.NgRAnnodeUexnApid) (*xnappducontentsv1.SnodeCounterCheckRequestIEsValue, error) {

	item := &xnappducontentsv1.SnodeCounterCheckRequestIEsValue{
		SnodeCounterCheckRequestIes: &xnappducontentsv1.SnodeCounterCheckRequestIEsValue_IdSNgRannodeUexnApid{
			IdSNgRannodeUexnApid: IDSNgRannodeUexnApID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeCounterCheckRequestIEsValueIDSNgRannodeUexnApID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeCounterCheckRequestIEsValueIDBearersSubjectToCounterCheck(IDBearersSubjectToCounterCheck *xnappducontentsv1.BearersSubjectToCounterCheckList) (*xnappducontentsv1.SnodeCounterCheckRequestIEsValue, error) {

	item := &xnappducontentsv1.SnodeCounterCheckRequestIEsValue{
		SnodeCounterCheckRequestIes: &xnappducontentsv1.SnodeCounterCheckRequestIEsValue_IdBearersSubjectToCounterCheck{
			IdBearersSubjectToCounterCheck: IDBearersSubjectToCounterCheck,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeCounterCheckRequestIEsValueIDBearersSubjectToCounterCheck() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeChangeRequiredIEsValueIDMNgRannodeUexnApID(IDMNgRannodeUexnApID *xnapiesv1.NgRAnnodeUexnApid) (*xnappducontentsv1.SnodeChangeRequiredIEsValue, error) {

	item := &xnappducontentsv1.SnodeChangeRequiredIEsValue{
		SnodeChangeRequiredIes: &xnappducontentsv1.SnodeChangeRequiredIEsValue_IdMNgRannodeUexnApid{
			IdMNgRannodeUexnApid: IDMNgRannodeUexnApID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeChangeRequiredIEsValueIDMNgRannodeUexnApID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeChangeRequiredIEsValueIDSNgRannodeUexnApID(IDSNgRannodeUexnApID *xnapiesv1.NgRAnnodeUexnApid) (*xnappducontentsv1.SnodeChangeRequiredIEsValue, error) {

	item := &xnappducontentsv1.SnodeChangeRequiredIEsValue{
		SnodeChangeRequiredIes: &xnappducontentsv1.SnodeChangeRequiredIEsValue_IdSNgRannodeUexnApid{
			IdSNgRannodeUexnApid: IDSNgRannodeUexnApID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeChangeRequiredIEsValueIDSNgRannodeUexnApID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeChangeRequiredIEsValueIDTargetSNgRannodeID(IDTargetSNgRannodeID *xnapiesv1.GlobalNgRAnnodeID) (*xnappducontentsv1.SnodeChangeRequiredIEsValue, error) {

	item := &xnappducontentsv1.SnodeChangeRequiredIEsValue{
		SnodeChangeRequiredIes: &xnappducontentsv1.SnodeChangeRequiredIEsValue_IdTargetSNgRannodeId{
			IdTargetSNgRannodeId: IDTargetSNgRannodeID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeChangeRequiredIEsValueIDTargetSNgRannodeID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeChangeRequiredIEsValueIDCause(IDCause *xnapiesv1.Cause) (*xnappducontentsv1.SnodeChangeRequiredIEsValue, error) {

	item := &xnappducontentsv1.SnodeChangeRequiredIEsValue{
		SnodeChangeRequiredIes: &xnappducontentsv1.SnodeChangeRequiredIEsValue_IdCause{
			IdCause: IDCause,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeChangeRequiredIEsValueIDCause() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeChangeRequiredIEsValueIDPdusessionSnchangeRequiredList(IDPdusessionSnchangeRequiredList *xnappducontentsv1.PdusessionSNchangeRequiredList) (*xnappducontentsv1.SnodeChangeRequiredIEsValue, error) {

	item := &xnappducontentsv1.SnodeChangeRequiredIEsValue{
		SnodeChangeRequiredIes: &xnappducontentsv1.SnodeChangeRequiredIEsValue_IdPdusessionSnchangeRequiredList{
			IdPdusessionSnchangeRequiredList: IDPdusessionSnchangeRequiredList,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeChangeRequiredIEsValueIDPdusessionSnchangeRequiredList() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeChangeRequiredIEsValueIDSnToMnContainer(IDSnToMnContainer string) (*xnappducontentsv1.SnodeChangeRequiredIEsValue, error) {

	item := &xnappducontentsv1.SnodeChangeRequiredIEsValue{
		SnodeChangeRequiredIes: &xnappducontentsv1.SnodeChangeRequiredIEsValue_IdSnToMnContainer{
			IdSnToMnContainer: IDSnToMnContainer,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeChangeRequiredIEsValueIDSnToMnContainer() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeChangeConfirmIEsValueIDMNgRannodeUexnApID(IDMNgRannodeUexnApID *xnapiesv1.NgRAnnodeUexnApid) (*xnappducontentsv1.SnodeChangeConfirmIEsValue, error) {

	item := &xnappducontentsv1.SnodeChangeConfirmIEsValue{
		SnodeChangeConfirmIes: &xnappducontentsv1.SnodeChangeConfirmIEsValue_IdMNgRannodeUexnApid{
			IdMNgRannodeUexnApid: IDMNgRannodeUexnApID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeChangeConfirmIEsValueIDMNgRannodeUexnApID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeChangeConfirmIEsValueIDSNgRannodeUexnApID(IDSNgRannodeUexnApID *xnapiesv1.NgRAnnodeUexnApid) (*xnappducontentsv1.SnodeChangeConfirmIEsValue, error) {

	item := &xnappducontentsv1.SnodeChangeConfirmIEsValue{
		SnodeChangeConfirmIes: &xnappducontentsv1.SnodeChangeConfirmIEsValue_IdSNgRannodeUexnApid{
			IdSNgRannodeUexnApid: IDSNgRannodeUexnApID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeChangeConfirmIEsValueIDSNgRannodeUexnApID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeChangeConfirmIEsValueIDPdusessionSnchangeConfirmList(IDPdusessionSnchangeConfirmList *xnappducontentsv1.PdusessionSNchangeConfirmList) (*xnappducontentsv1.SnodeChangeConfirmIEsValue, error) {

	item := &xnappducontentsv1.SnodeChangeConfirmIEsValue{
		SnodeChangeConfirmIes: &xnappducontentsv1.SnodeChangeConfirmIEsValue_IdPdusessionSnchangeConfirmList{
			IdPdusessionSnchangeConfirmList: IDPdusessionSnchangeConfirmList,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeChangeConfirmIEsValueIDPdusessionSnchangeConfirmList() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeChangeConfirmIEsValueIDCriticalityDiagnostics(IDCriticalityDiagnostics *xnapiesv1.CriticalityDiagnostics) (*xnappducontentsv1.SnodeChangeConfirmIEsValue, error) {

	item := &xnappducontentsv1.SnodeChangeConfirmIEsValue{
		SnodeChangeConfirmIes: &xnappducontentsv1.SnodeChangeConfirmIEsValue_IdCriticalityDiagnostics{
			IdCriticalityDiagnostics: IDCriticalityDiagnostics,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeChangeConfirmIEsValueIDCriticalityDiagnostics() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeChangeRefuseIEsValueIDMNgRannodeUexnApID(IDMNgRannodeUexnApID *xnapiesv1.NgRAnnodeUexnApid) (*xnappducontentsv1.SnodeChangeRefuseIEsValue, error) {

	item := &xnappducontentsv1.SnodeChangeRefuseIEsValue{
		SnodeChangeRefuseIes: &xnappducontentsv1.SnodeChangeRefuseIEsValue_IdMNgRannodeUexnApid{
			IdMNgRannodeUexnApid: IDMNgRannodeUexnApID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeChangeRefuseIEsValueIDMNgRannodeUexnApID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeChangeRefuseIEsValueIDSNgRannodeUexnApID(IDSNgRannodeUexnApID *xnapiesv1.NgRAnnodeUexnApid) (*xnappducontentsv1.SnodeChangeRefuseIEsValue, error) {

	item := &xnappducontentsv1.SnodeChangeRefuseIEsValue{
		SnodeChangeRefuseIes: &xnappducontentsv1.SnodeChangeRefuseIEsValue_IdSNgRannodeUexnApid{
			IdSNgRannodeUexnApid: IDSNgRannodeUexnApID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeChangeRefuseIEsValueIDSNgRannodeUexnApID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeChangeRefuseIEsValueIDCause(IDCause *xnapiesv1.Cause) (*xnappducontentsv1.SnodeChangeRefuseIEsValue, error) {

	item := &xnappducontentsv1.SnodeChangeRefuseIEsValue{
		SnodeChangeRefuseIes: &xnappducontentsv1.SnodeChangeRefuseIEsValue_IdCause{
			IdCause: IDCause,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeChangeRefuseIEsValueIDCause() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSnodeChangeRefuseIEsValueIDCriticalityDiagnostics(IDCriticalityDiagnostics *xnapiesv1.CriticalityDiagnostics) (*xnappducontentsv1.SnodeChangeRefuseIEsValue, error) {

	item := &xnappducontentsv1.SnodeChangeRefuseIEsValue{
		SnodeChangeRefuseIes: &xnappducontentsv1.SnodeChangeRefuseIEsValue_IdCriticalityDiagnostics{
			IdCriticalityDiagnostics: IDCriticalityDiagnostics,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSnodeChangeRefuseIEsValueIDCriticalityDiagnostics() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateRrctransferIEsValueIDMNgRannodeUexnApID(IDMNgRannodeUexnApID *xnapiesv1.NgRAnnodeUexnApid) (*xnappducontentsv1.RrctransferIEsValue, error) {

	item := &xnappducontentsv1.RrctransferIEsValue{
		RrctransferIes: &xnappducontentsv1.RrctransferIEsValue_IdMNgRannodeUexnApid{
			IdMNgRannodeUexnApid: IDMNgRannodeUexnApID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRrctransferIEsValueIDMNgRannodeUexnApID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateRrctransferIEsValueIDSNgRannodeUexnApID(IDSNgRannodeUexnApID *xnapiesv1.NgRAnnodeUexnApid) (*xnappducontentsv1.RrctransferIEsValue, error) {

	item := &xnappducontentsv1.RrctransferIEsValue{
		RrctransferIes: &xnappducontentsv1.RrctransferIEsValue_IdSNgRannodeUexnApid{
			IdSNgRannodeUexnApid: IDSNgRannodeUexnApID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRrctransferIEsValueIDSNgRannodeUexnApID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateRrctransferIEsValueIDSplitSrbRrctransfer(IDSplitSrbRrctransfer *xnappducontentsv1.SplitSrbRRctransfer) (*xnappducontentsv1.RrctransferIEsValue, error) {

	item := &xnappducontentsv1.RrctransferIEsValue{
		RrctransferIes: &xnappducontentsv1.RrctransferIEsValue_IdSplitSrbRrctransfer{
			IdSplitSrbRrctransfer: IDSplitSrbRrctransfer,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRrctransferIEsValueIDSplitSrbRrctransfer() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateRrctransferIEsValueIDUereportRrctransfer(IDUereportRrctransfer *xnappducontentsv1.UereportRrctransfer) (*xnappducontentsv1.RrctransferIEsValue, error) {

	item := &xnappducontentsv1.RrctransferIEsValue{
		RrctransferIes: &xnappducontentsv1.RrctransferIEsValue_IdUereportRrctransfer{
			IdUereportRrctransfer: IDUereportRrctransfer,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRrctransferIEsValueIDUereportRrctransfer() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateRrctransferIEsValueIDFastMcgrecoveryRrctransferSnToMn(IDFastMcgrecoveryRrctransferSnToMn *xnappducontentsv1.FastMcgrecoveryRrctransfer) (*xnappducontentsv1.RrctransferIEsValue, error) {

	item := &xnappducontentsv1.RrctransferIEsValue{
		RrctransferIes: &xnappducontentsv1.RrctransferIEsValue_IdFastMcgrecoveryRrctransferSnToMn{
			IdFastMcgrecoveryRrctransferSnToMn: IDFastMcgrecoveryRrctransferSnToMn,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRrctransferIEsValueIDFastMcgrecoveryRrctransferSnToMn() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateRrctransferIEsValueIDFastMcgrecoveryRrctransferMnToSn(IDFastMcgrecoveryRrctransferMnToSn *xnappducontentsv1.FastMcgrecoveryRrctransfer) (*xnappducontentsv1.RrctransferIEsValue, error) {

	item := &xnappducontentsv1.RrctransferIEsValue{
		RrctransferIes: &xnappducontentsv1.RrctransferIEsValue_IdFastMcgrecoveryRrctransferMnToSn{
			IdFastMcgrecoveryRrctransferMnToSn: IDFastMcgrecoveryRrctransferMnToSn,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRrctransferIEsValueIDFastMcgrecoveryRrctransferMnToSn() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateNotificationControlIndicationIEsValueIDMNgRannodeUexnApID(IDMNgRannodeUexnApID *xnapiesv1.NgRAnnodeUexnApid) (*xnappducontentsv1.NotificationControlIndicationIEsValue, error) {

	item := &xnappducontentsv1.NotificationControlIndicationIEsValue{
		NotificationControlIndicationIes: &xnappducontentsv1.NotificationControlIndicationIEsValue_IdMNgRannodeUexnApid{
			IdMNgRannodeUexnApid: IDMNgRannodeUexnApID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNotificationControlIndicationIEsValueIDMNgRannodeUexnApID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateNotificationControlIndicationIEsValueIDSNgRannodeUexnApID(IDSNgRannodeUexnApID *xnapiesv1.NgRAnnodeUexnApid) (*xnappducontentsv1.NotificationControlIndicationIEsValue, error) {

	item := &xnappducontentsv1.NotificationControlIndicationIEsValue{
		NotificationControlIndicationIes: &xnappducontentsv1.NotificationControlIndicationIEsValue_IdSNgRannodeUexnApid{
			IdSNgRannodeUexnApid: IDSNgRannodeUexnApID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNotificationControlIndicationIEsValueIDSNgRannodeUexnApID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateNotificationControlIndicationIEsValueIDPdusessionResourcesNotifyList(IDPdusessionResourcesNotifyList *xnappducontentsv1.PdusessionResourcesNotifyList) (*xnappducontentsv1.NotificationControlIndicationIEsValue, error) {

	item := &xnappducontentsv1.NotificationControlIndicationIEsValue{
		NotificationControlIndicationIes: &xnappducontentsv1.NotificationControlIndicationIEsValue_IdPdusessionResourcesNotifyList{
			IdPdusessionResourcesNotifyList: IDPdusessionResourcesNotifyList,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNotificationControlIndicationIEsValueIDPdusessionResourcesNotifyList() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateActivityNotificationIEsValueIDMNgRannodeUexnApID(IDMNgRannodeUexnApID *xnapiesv1.NgRAnnodeUexnApid) (*xnappducontentsv1.ActivityNotificationIEsValue, error) {

	item := &xnappducontentsv1.ActivityNotificationIEsValue{
		ActivityNotificationIes: &xnappducontentsv1.ActivityNotificationIEsValue_IdMNgRannodeUexnApid{
			IdMNgRannodeUexnApid: IDMNgRannodeUexnApID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateActivityNotificationIEsValueIDMNgRannodeUexnApID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateActivityNotificationIEsValueIDSNgRannodeUexnApID(IDSNgRannodeUexnApID *xnapiesv1.NgRAnnodeUexnApid) (*xnappducontentsv1.ActivityNotificationIEsValue, error) {

	item := &xnappducontentsv1.ActivityNotificationIEsValue{
		ActivityNotificationIes: &xnappducontentsv1.ActivityNotificationIEsValue_IdSNgRannodeUexnApid{
			IdSNgRannodeUexnApid: IDSNgRannodeUexnApID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateActivityNotificationIEsValueIDSNgRannodeUexnApID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateActivityNotificationIEsValueIDUserPlaneTrafficActivityReport(IDUserPlaneTrafficActivityReport xnapiesv1.UserPlaneTrafficActivityReport) (*xnappducontentsv1.ActivityNotificationIEsValue, error) {

	item := &xnappducontentsv1.ActivityNotificationIEsValue{
		ActivityNotificationIes: &xnappducontentsv1.ActivityNotificationIEsValue_IdUserPlaneTrafficActivityReport{
			IdUserPlaneTrafficActivityReport: IDUserPlaneTrafficActivityReport,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateActivityNotificationIEsValueIDUserPlaneTrafficActivityReport() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateActivityNotificationIEsValueIDPdusessionResourcesActivityNotifyList(IDPdusessionResourcesActivityNotifyList *xnappducontentsv1.PdusessionResourcesActivityNotifyList) (*xnappducontentsv1.ActivityNotificationIEsValue, error) {

	item := &xnappducontentsv1.ActivityNotificationIEsValue{
		ActivityNotificationIes: &xnappducontentsv1.ActivityNotificationIEsValue_IdPdusessionResourcesActivityNotifyList{
			IdPdusessionResourcesActivityNotifyList: IDPdusessionResourcesActivityNotifyList,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateActivityNotificationIEsValueIDPdusessionResourcesActivityNotifyList() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateActivityNotificationIEsValueIDRanpagingFailure(IDRanpagingFailure xnapiesv1.RanpagingFailure) (*xnappducontentsv1.ActivityNotificationIEsValue, error) {

	item := &xnappducontentsv1.ActivityNotificationIEsValue{
		ActivityNotificationIes: &xnappducontentsv1.ActivityNotificationIEsValue_IdRanpagingFailure{
			IdRanpagingFailure: IDRanpagingFailure,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateActivityNotificationIEsValueIDRanpagingFailure() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateXnSetupRequestIEsValueIDGlobalNgRanNodeID(IDGlobalNgRanNodeID *xnapiesv1.GlobalNgRAnnodeID) (*xnappducontentsv1.XnSetupRequestIEsValue, error) {

	item := &xnappducontentsv1.XnSetupRequestIEsValue{
		XnSetupRequestIes: &xnappducontentsv1.XnSetupRequestIEsValue_IdGlobalNgRanNodeId{
			IdGlobalNgRanNodeId: IDGlobalNgRanNodeID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateXnSetupRequestIEsValueIDGlobalNgRanNodeID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateXnSetupRequestIEsValueIDTaisupportList(IDTaisupportList *xnapiesv1.TaisupportList) (*xnappducontentsv1.XnSetupRequestIEsValue, error) {

	item := &xnappducontentsv1.XnSetupRequestIEsValue{
		XnSetupRequestIes: &xnappducontentsv1.XnSetupRequestIEsValue_IdTaisupportList{
			IdTaisupportList: IDTaisupportList,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateXnSetupRequestIEsValueIDTaisupportList() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateXnSetupRequestIEsValueIDAmfRegionInformation(IDAmfRegionInformation *xnapiesv1.AmfRegionInformation) (*xnappducontentsv1.XnSetupRequestIEsValue, error) {

	item := &xnappducontentsv1.XnSetupRequestIEsValue{
		XnSetupRequestIes: &xnappducontentsv1.XnSetupRequestIEsValue_IdAmfRegionInformation{
			IdAmfRegionInformation: IDAmfRegionInformation,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateXnSetupRequestIEsValueIDAmfRegionInformation() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateXnSetupRequestIEsValueIDListOfServedCellsNr(IDListOfServedCellsNr *xnapiesv1.ServedCellsNR) (*xnappducontentsv1.XnSetupRequestIEsValue, error) {

	item := &xnappducontentsv1.XnSetupRequestIEsValue{
		XnSetupRequestIes: &xnappducontentsv1.XnSetupRequestIEsValue_IdListOfServedCellsNr{
			IdListOfServedCellsNr: IDListOfServedCellsNr,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateXnSetupRequestIEsValueIDListOfServedCellsNr() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateXnSetupRequestIEsValueIDListOfServedCellsEUtra(IDListOfServedCellsEUtra *xnapiesv1.ServedCellsEUTra) (*xnappducontentsv1.XnSetupRequestIEsValue, error) {

	item := &xnappducontentsv1.XnSetupRequestIEsValue{
		XnSetupRequestIes: &xnappducontentsv1.XnSetupRequestIEsValue_IdListOfServedCellsEUtra{
			IdListOfServedCellsEUtra: IDListOfServedCellsEUtra,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateXnSetupRequestIEsValueIDListOfServedCellsEUtra() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateXnSetupRequestIEsValueIDInterfaceInstanceIndication(IDInterfaceInstanceIndication *xnapiesv1.InterfaceInstanceIndication) (*xnappducontentsv1.XnSetupRequestIEsValue, error) {

	item := &xnappducontentsv1.XnSetupRequestIEsValue{
		XnSetupRequestIes: &xnappducontentsv1.XnSetupRequestIEsValue_IdInterfaceInstanceIndication{
			IdInterfaceInstanceIndication: IDInterfaceInstanceIndication,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateXnSetupRequestIEsValueIDInterfaceInstanceIndication() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateXnSetupRequestIEsValueIDTnlconfigurationInfo(IDTnlconfigurationInfo *xnapiesv1.TnlconfigurationInfo) (*xnappducontentsv1.XnSetupRequestIEsValue, error) {

	item := &xnappducontentsv1.XnSetupRequestIEsValue{
		XnSetupRequestIes: &xnappducontentsv1.XnSetupRequestIEsValue_IdTnlconfigurationInfo{
			IdTnlconfigurationInfo: IDTnlconfigurationInfo,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateXnSetupRequestIEsValueIDTnlconfigurationInfo() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateXnSetupRequestIEsValueIDPartialListIndicatorNr(IDPartialListIndicatorNr xnapiesv1.PartialListIndicator) (*xnappducontentsv1.XnSetupRequestIEsValue, error) {

	item := &xnappducontentsv1.XnSetupRequestIEsValue{
		XnSetupRequestIes: &xnappducontentsv1.XnSetupRequestIEsValue_IdPartialListIndicatorNr{
			IdPartialListIndicatorNr: IDPartialListIndicatorNr,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateXnSetupRequestIEsValueIDPartialListIndicatorNr() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateXnSetupRequestIEsValueIDCellAndCapacityAssistanceInfoNr(IDCellAndCapacityAssistanceInfoNr *xnapiesv1.CellAndCapacityAssistanceInfoNR) (*xnappducontentsv1.XnSetupRequestIEsValue, error) {

	item := &xnappducontentsv1.XnSetupRequestIEsValue{
		XnSetupRequestIes: &xnappducontentsv1.XnSetupRequestIEsValue_IdCellAndCapacityAssistanceInfoNr{
			IdCellAndCapacityAssistanceInfoNr: IDCellAndCapacityAssistanceInfoNr,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateXnSetupRequestIEsValueIDCellAndCapacityAssistanceInfoNr() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateXnSetupRequestIEsValueIDPartialListIndicatorEutra(IDPartialListIndicatorEutra xnapiesv1.PartialListIndicator) (*xnappducontentsv1.XnSetupRequestIEsValue, error) {

	item := &xnappducontentsv1.XnSetupRequestIEsValue{
		XnSetupRequestIes: &xnappducontentsv1.XnSetupRequestIEsValue_IdPartialListIndicatorEutra{
			IdPartialListIndicatorEutra: IDPartialListIndicatorEutra,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateXnSetupRequestIEsValueIDPartialListIndicatorEutra() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateXnSetupRequestIEsValueIDCellAndCapacityAssistanceInfoEutra(IDCellAndCapacityAssistanceInfoEutra *xnapiesv1.CellAndCapacityAssistanceInfoEUtra) (*xnappducontentsv1.XnSetupRequestIEsValue, error) {

	item := &xnappducontentsv1.XnSetupRequestIEsValue{
		XnSetupRequestIes: &xnappducontentsv1.XnSetupRequestIEsValue_IdCellAndCapacityAssistanceInfoEutra{
			IdCellAndCapacityAssistanceInfoEutra: IDCellAndCapacityAssistanceInfoEutra,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateXnSetupRequestIEsValueIDCellAndCapacityAssistanceInfoEutra() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateXnSetupResponseIEsValueIDGlobalNgRanNodeID(IDGlobalNgRanNodeID *xnapiesv1.GlobalNgRAnnodeID) (*xnappducontentsv1.XnSetupResponseIEsValue, error) {

	item := &xnappducontentsv1.XnSetupResponseIEsValue{
		XnSetupResponseIes: &xnappducontentsv1.XnSetupResponseIEsValue_IdGlobalNgRanNodeId{
			IdGlobalNgRanNodeId: IDGlobalNgRanNodeID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateXnSetupResponseIEsValueIDGlobalNgRanNodeID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateXnSetupResponseIEsValueIDTaisupportList(IDTaisupportList *xnapiesv1.TaisupportList) (*xnappducontentsv1.XnSetupResponseIEsValue, error) {

	item := &xnappducontentsv1.XnSetupResponseIEsValue{
		XnSetupResponseIes: &xnappducontentsv1.XnSetupResponseIEsValue_IdTaisupportList{
			IdTaisupportList: IDTaisupportList,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateXnSetupResponseIEsValueIDTaisupportList() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateXnSetupResponseIEsValueIDListOfServedCellsNr(IDListOfServedCellsNr *xnapiesv1.ServedCellsNR) (*xnappducontentsv1.XnSetupResponseIEsValue, error) {

	item := &xnappducontentsv1.XnSetupResponseIEsValue{
		XnSetupResponseIes: &xnappducontentsv1.XnSetupResponseIEsValue_IdListOfServedCellsNr{
			IdListOfServedCellsNr: IDListOfServedCellsNr,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateXnSetupResponseIEsValueIDListOfServedCellsNr() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateXnSetupResponseIEsValueIDListOfServedCellsEUtra(IDListOfServedCellsEUtra *xnapiesv1.ServedCellsEUTra) (*xnappducontentsv1.XnSetupResponseIEsValue, error) {

	item := &xnappducontentsv1.XnSetupResponseIEsValue{
		XnSetupResponseIes: &xnappducontentsv1.XnSetupResponseIEsValue_IdListOfServedCellsEUtra{
			IdListOfServedCellsEUtra: IDListOfServedCellsEUtra,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateXnSetupResponseIEsValueIDListOfServedCellsEUtra() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateXnSetupResponseIEsValueIDCriticalityDiagnostics(IDCriticalityDiagnostics *xnapiesv1.CriticalityDiagnostics) (*xnappducontentsv1.XnSetupResponseIEsValue, error) {

	item := &xnappducontentsv1.XnSetupResponseIEsValue{
		XnSetupResponseIes: &xnappducontentsv1.XnSetupResponseIEsValue_IdCriticalityDiagnostics{
			IdCriticalityDiagnostics: IDCriticalityDiagnostics,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateXnSetupResponseIEsValueIDCriticalityDiagnostics() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateXnSetupResponseIEsValueIDAmfRegionInformation(IDAmfRegionInformation *xnapiesv1.AmfRegionInformation) (*xnappducontentsv1.XnSetupResponseIEsValue, error) {

	item := &xnappducontentsv1.XnSetupResponseIEsValue{
		XnSetupResponseIes: &xnappducontentsv1.XnSetupResponseIEsValue_IdAmfRegionInformation{
			IdAmfRegionInformation: IDAmfRegionInformation,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateXnSetupResponseIEsValueIDAmfRegionInformation() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateXnSetupResponseIEsValueIDInterfaceInstanceIndication(IDInterfaceInstanceIndication *xnapiesv1.InterfaceInstanceIndication) (*xnappducontentsv1.XnSetupResponseIEsValue, error) {

	item := &xnappducontentsv1.XnSetupResponseIEsValue{
		XnSetupResponseIes: &xnappducontentsv1.XnSetupResponseIEsValue_IdInterfaceInstanceIndication{
			IdInterfaceInstanceIndication: IDInterfaceInstanceIndication,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateXnSetupResponseIEsValueIDInterfaceInstanceIndication() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateXnSetupResponseIEsValueIDTnlconfigurationInfo(IDTnlconfigurationInfo *xnapiesv1.TnlconfigurationInfo) (*xnappducontentsv1.XnSetupResponseIEsValue, error) {

	item := &xnappducontentsv1.XnSetupResponseIEsValue{
		XnSetupResponseIes: &xnappducontentsv1.XnSetupResponseIEsValue_IdTnlconfigurationInfo{
			IdTnlconfigurationInfo: IDTnlconfigurationInfo,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateXnSetupResponseIEsValueIDTnlconfigurationInfo() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateXnSetupResponseIEsValueIDPartialListIndicatorNr(IDPartialListIndicatorNr xnapiesv1.PartialListIndicator) (*xnappducontentsv1.XnSetupResponseIEsValue, error) {

	item := &xnappducontentsv1.XnSetupResponseIEsValue{
		XnSetupResponseIes: &xnappducontentsv1.XnSetupResponseIEsValue_IdPartialListIndicatorNr{
			IdPartialListIndicatorNr: IDPartialListIndicatorNr,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateXnSetupResponseIEsValueIDPartialListIndicatorNr() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateXnSetupResponseIEsValueIDCellAndCapacityAssistanceInfoNr(IDCellAndCapacityAssistanceInfoNr *xnapiesv1.CellAndCapacityAssistanceInfoNR) (*xnappducontentsv1.XnSetupResponseIEsValue, error) {

	item := &xnappducontentsv1.XnSetupResponseIEsValue{
		XnSetupResponseIes: &xnappducontentsv1.XnSetupResponseIEsValue_IdCellAndCapacityAssistanceInfoNr{
			IdCellAndCapacityAssistanceInfoNr: IDCellAndCapacityAssistanceInfoNr,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateXnSetupResponseIEsValueIDCellAndCapacityAssistanceInfoNr() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateXnSetupResponseIEsValueIDPartialListIndicatorEutra(IDPartialListIndicatorEutra xnapiesv1.PartialListIndicator) (*xnappducontentsv1.XnSetupResponseIEsValue, error) {

	item := &xnappducontentsv1.XnSetupResponseIEsValue{
		XnSetupResponseIes: &xnappducontentsv1.XnSetupResponseIEsValue_IdPartialListIndicatorEutra{
			IdPartialListIndicatorEutra: IDPartialListIndicatorEutra,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateXnSetupResponseIEsValueIDPartialListIndicatorEutra() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateXnSetupResponseIEsValueIDCellAndCapacityAssistanceInfoEutra(IDCellAndCapacityAssistanceInfoEutra *xnapiesv1.CellAndCapacityAssistanceInfoEUtra) (*xnappducontentsv1.XnSetupResponseIEsValue, error) {

	item := &xnappducontentsv1.XnSetupResponseIEsValue{
		XnSetupResponseIes: &xnappducontentsv1.XnSetupResponseIEsValue_IdCellAndCapacityAssistanceInfoEutra{
			IdCellAndCapacityAssistanceInfoEutra: IDCellAndCapacityAssistanceInfoEutra,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateXnSetupResponseIEsValueIDCellAndCapacityAssistanceInfoEutra() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateXnSetupFailureIEsValueIDCause(IDCause *xnapiesv1.Cause) (*xnappducontentsv1.XnSetupFailureIEsValue, error) {

	item := &xnappducontentsv1.XnSetupFailureIEsValue{
		XnSetupFailureIes: &xnappducontentsv1.XnSetupFailureIEsValue_IdCause{
			IdCause: IDCause,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateXnSetupFailureIEsValueIDCause() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateXnSetupFailureIEsValueIDTimeToWait(IDTimeToWait xnapiesv1.TimeToWait) (*xnappducontentsv1.XnSetupFailureIEsValue, error) {

	item := &xnappducontentsv1.XnSetupFailureIEsValue{
		XnSetupFailureIes: &xnappducontentsv1.XnSetupFailureIEsValue_IdTimeToWait{
			IdTimeToWait: IDTimeToWait,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateXnSetupFailureIEsValueIDTimeToWait() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateXnSetupFailureIEsValueIDCriticalityDiagnostics(IDCriticalityDiagnostics *xnapiesv1.CriticalityDiagnostics) (*xnappducontentsv1.XnSetupFailureIEsValue, error) {

	item := &xnappducontentsv1.XnSetupFailureIEsValue{
		XnSetupFailureIes: &xnappducontentsv1.XnSetupFailureIEsValue_IdCriticalityDiagnostics{
			IdCriticalityDiagnostics: IDCriticalityDiagnostics,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateXnSetupFailureIEsValueIDCriticalityDiagnostics() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateXnSetupFailureIEsValueIDInterfaceInstanceIndication(IDInterfaceInstanceIndication *xnapiesv1.InterfaceInstanceIndication) (*xnappducontentsv1.XnSetupFailureIEsValue, error) {

	item := &xnappducontentsv1.XnSetupFailureIEsValue{
		XnSetupFailureIes: &xnappducontentsv1.XnSetupFailureIEsValue_IdInterfaceInstanceIndication{
			IdInterfaceInstanceIndication: IDInterfaceInstanceIndication,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateXnSetupFailureIEsValueIDInterfaceInstanceIndication() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateXnSetupFailureIEsValueIDMessageOversizeNotification(IDMessageOversizeNotification *xnapiesv1.MessageOversizeNotification) (*xnappducontentsv1.XnSetupFailureIEsValue, error) {

	item := &xnappducontentsv1.XnSetupFailureIEsValue{
		XnSetupFailureIes: &xnappducontentsv1.XnSetupFailureIEsValue_IdMessageOversizeNotification{
			IdMessageOversizeNotification: IDMessageOversizeNotification,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateXnSetupFailureIEsValueIDMessageOversizeNotification() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateNgrannodeConfigurationUpdateIEsValueIDTaisupportList(IDTaisupportList *xnapiesv1.TaisupportList) (*xnappducontentsv1.NgrannodeConfigurationUpdateIEsValue, error) {

	item := &xnappducontentsv1.NgrannodeConfigurationUpdateIEsValue{
		NgrannodeConfigurationUpdateIes: &xnappducontentsv1.NgrannodeConfigurationUpdateIEsValue_IdTaisupportList{
			IdTaisupportList: IDTaisupportList,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNgrannodeConfigurationUpdateIEsValueIDTaisupportList() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateNgrannodeConfigurationUpdateIEsValueIDConfigurationUpdateInitiatingNodeChoice(IDConfigurationUpdateInitiatingNodeChoice *xnappducontentsv1.ConfigurationUpdateInitiatingNodeChoice) (*xnappducontentsv1.NgrannodeConfigurationUpdateIEsValue, error) {

	item := &xnappducontentsv1.NgrannodeConfigurationUpdateIEsValue{
		NgrannodeConfigurationUpdateIes: &xnappducontentsv1.NgrannodeConfigurationUpdateIEsValue_IdConfigurationUpdateInitiatingNodeChoice{
			IdConfigurationUpdateInitiatingNodeChoice: IDConfigurationUpdateInitiatingNodeChoice,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNgrannodeConfigurationUpdateIEsValueIDConfigurationUpdateInitiatingNodeChoice() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateNgrannodeConfigurationUpdateIEsValueIDTnlaToAddList(IDTnlaToAddList *xnapiesv1.TnlaToAddList) (*xnappducontentsv1.NgrannodeConfigurationUpdateIEsValue, error) {

	item := &xnappducontentsv1.NgrannodeConfigurationUpdateIEsValue{
		NgrannodeConfigurationUpdateIes: &xnappducontentsv1.NgrannodeConfigurationUpdateIEsValue_IdTnlaToAddList{
			IdTnlaToAddList: IDTnlaToAddList,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNgrannodeConfigurationUpdateIEsValueIDTnlaToAddList() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateNgrannodeConfigurationUpdateIEsValueIDTnlaToRemoveList(IDTnlaToRemoveList *xnapiesv1.TnlaToRemoveList) (*xnappducontentsv1.NgrannodeConfigurationUpdateIEsValue, error) {

	item := &xnappducontentsv1.NgrannodeConfigurationUpdateIEsValue{
		NgrannodeConfigurationUpdateIes: &xnappducontentsv1.NgrannodeConfigurationUpdateIEsValue_IdTnlaToRemoveList{
			IdTnlaToRemoveList: IDTnlaToRemoveList,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNgrannodeConfigurationUpdateIEsValueIDTnlaToRemoveList() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateNgrannodeConfigurationUpdateIEsValueIDTnlaToUpdateList(IDTnlaToUpdateList *xnapiesv1.TnlaToUpdateList) (*xnappducontentsv1.NgrannodeConfigurationUpdateIEsValue, error) {

	item := &xnappducontentsv1.NgrannodeConfigurationUpdateIEsValue{
		NgrannodeConfigurationUpdateIes: &xnappducontentsv1.NgrannodeConfigurationUpdateIEsValue_IdTnlaToUpdateList{
			IdTnlaToUpdateList: IDTnlaToUpdateList,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNgrannodeConfigurationUpdateIEsValueIDTnlaToUpdateList() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateNgrannodeConfigurationUpdateIEsValueIDGlobalNgRanNodeID(IDGlobalNgRanNodeID *xnapiesv1.GlobalNgRAnnodeID) (*xnappducontentsv1.NgrannodeConfigurationUpdateIEsValue, error) {

	item := &xnappducontentsv1.NgrannodeConfigurationUpdateIEsValue{
		NgrannodeConfigurationUpdateIes: &xnappducontentsv1.NgrannodeConfigurationUpdateIEsValue_IdGlobalNgRanNodeId{
			IdGlobalNgRanNodeId: IDGlobalNgRanNodeID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNgrannodeConfigurationUpdateIEsValueIDGlobalNgRanNodeID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateNgrannodeConfigurationUpdateIEsValueIDAmfRegionInformationToAdd(IDAmfRegionInformationToAdd *xnapiesv1.AmfRegionInformation) (*xnappducontentsv1.NgrannodeConfigurationUpdateIEsValue, error) {

	item := &xnappducontentsv1.NgrannodeConfigurationUpdateIEsValue{
		NgrannodeConfigurationUpdateIes: &xnappducontentsv1.NgrannodeConfigurationUpdateIEsValue_IdAmfRegionInformationToAdd{
			IdAmfRegionInformationToAdd: IDAmfRegionInformationToAdd,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNgrannodeConfigurationUpdateIEsValueIDAmfRegionInformationToAdd() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateNgrannodeConfigurationUpdateIEsValueIDAmfRegionInformationToDelete(IDAmfRegionInformationToDelete *xnapiesv1.AmfRegionInformation) (*xnappducontentsv1.NgrannodeConfigurationUpdateIEsValue, error) {

	item := &xnappducontentsv1.NgrannodeConfigurationUpdateIEsValue{
		NgrannodeConfigurationUpdateIes: &xnappducontentsv1.NgrannodeConfigurationUpdateIEsValue_IdAmfRegionInformationToDelete{
			IdAmfRegionInformationToDelete: IDAmfRegionInformationToDelete,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNgrannodeConfigurationUpdateIEsValueIDAmfRegionInformationToDelete() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateNgrannodeConfigurationUpdateIEsValueIDInterfaceInstanceIndication(IDInterfaceInstanceIndication *xnapiesv1.InterfaceInstanceIndication) (*xnappducontentsv1.NgrannodeConfigurationUpdateIEsValue, error) {

	item := &xnappducontentsv1.NgrannodeConfigurationUpdateIEsValue{
		NgrannodeConfigurationUpdateIes: &xnappducontentsv1.NgrannodeConfigurationUpdateIEsValue_IdInterfaceInstanceIndication{
			IdInterfaceInstanceIndication: IDInterfaceInstanceIndication,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNgrannodeConfigurationUpdateIEsValueIDInterfaceInstanceIndication() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateNgrannodeConfigurationUpdateIEsValueIDTnlconfigurationInfo(IDTnlconfigurationInfo *xnapiesv1.TnlconfigurationInfo) (*xnappducontentsv1.NgrannodeConfigurationUpdateIEsValue, error) {

	item := &xnappducontentsv1.NgrannodeConfigurationUpdateIEsValue{
		NgrannodeConfigurationUpdateIes: &xnappducontentsv1.NgrannodeConfigurationUpdateIEsValue_IdTnlconfigurationInfo{
			IdTnlconfigurationInfo: IDTnlconfigurationInfo,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNgrannodeConfigurationUpdateIEsValueIDTnlconfigurationInfo() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateConfigurationUpdategNbValueIDServedCellsToUpdateNr(IDServedCellsToUpdateNr *xnapiesv1.ServedCellsToUpdateNR) (*xnappducontentsv1.ConfigurationUpdategNbValue, error) {

	item := &xnappducontentsv1.ConfigurationUpdategNbValue{
		ConfigurationUpdateGNb: &xnappducontentsv1.ConfigurationUpdategNbValue_IdServedCellsToUpdateNr{
			IdServedCellsToUpdateNr: IDServedCellsToUpdateNr,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateConfigurationUpdategNbValueIDServedCellsToUpdateNr() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateConfigurationUpdategNbValueIDCellAssistanceInfoNr(IDCellAssistanceInfoNr *xnapiesv1.CellAssistanceInfoNR) (*xnappducontentsv1.ConfigurationUpdategNbValue, error) {

	item := &xnappducontentsv1.ConfigurationUpdategNbValue{
		ConfigurationUpdateGNb: &xnappducontentsv1.ConfigurationUpdategNbValue_IdCellAssistanceInfoNr{
			IdCellAssistanceInfoNr: IDCellAssistanceInfoNr,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateConfigurationUpdategNbValueIDCellAssistanceInfoNr() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateConfigurationUpdategNbValueIDCellAssistanceInfoEutra(IDCellAssistanceInfoEutra *xnapiesv1.CellAssistanceInfoEUtra) (*xnappducontentsv1.ConfigurationUpdategNbValue, error) {

	item := &xnappducontentsv1.ConfigurationUpdategNbValue{
		ConfigurationUpdateGNb: &xnappducontentsv1.ConfigurationUpdategNbValue_IdCellAssistanceInfoEutra{
			IdCellAssistanceInfoEutra: IDCellAssistanceInfoEutra,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateConfigurationUpdategNbValueIDCellAssistanceInfoEutra() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateConfigurationUpdatengeNbValueIDServedCellsToUpdateEUtra(IDServedCellsToUpdateEUtra *xnapiesv1.ServedCellsToUpdateEUTra) (*xnappducontentsv1.ConfigurationUpdatengeNbValue, error) {

	item := &xnappducontentsv1.ConfigurationUpdatengeNbValue{
		ConfigurationUpdateNgENb: &xnappducontentsv1.ConfigurationUpdatengeNbValue_IdServedCellsToUpdateEUtra{
			IdServedCellsToUpdateEUtra: IDServedCellsToUpdateEUtra,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateConfigurationUpdatengeNbValueIDServedCellsToUpdateEUtra() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateConfigurationUpdatengeNbValueIDCellAssistanceInfoNr(IDCellAssistanceInfoNr *xnapiesv1.CellAssistanceInfoNR) (*xnappducontentsv1.ConfigurationUpdatengeNbValue, error) {

	item := &xnappducontentsv1.ConfigurationUpdatengeNbValue{
		ConfigurationUpdateNgENb: &xnappducontentsv1.ConfigurationUpdatengeNbValue_IdCellAssistanceInfoNr{
			IdCellAssistanceInfoNr: IDCellAssistanceInfoNr,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateConfigurationUpdatengeNbValueIDCellAssistanceInfoNr() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateConfigurationUpdatengeNbValueIDCellAssistanceInfoEutra(IDCellAssistanceInfoEutra *xnapiesv1.CellAssistanceInfoEUtra) (*xnappducontentsv1.ConfigurationUpdatengeNbValue, error) {

	item := &xnappducontentsv1.ConfigurationUpdatengeNbValue{
		ConfigurationUpdateNgENb: &xnappducontentsv1.ConfigurationUpdatengeNbValue_IdCellAssistanceInfoEutra{
			IdCellAssistanceInfoEutra: IDCellAssistanceInfoEutra,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateConfigurationUpdatengeNbValueIDCellAssistanceInfoEutra() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateConfigurationUpdateInitiatingNodeChoiceGNb(gNb *xnappducontentsv1.ConfigurationUpdategNb) (*xnappducontentsv1.ConfigurationUpdateInitiatingNodeChoice, error) {

	item := &xnappducontentsv1.ConfigurationUpdateInitiatingNodeChoice{
		ConfigurationUpdateInitiatingNodeChoice: &xnappducontentsv1.ConfigurationUpdateInitiatingNodeChoice_GNb{
			GNb: gNb,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateConfigurationUpdateInitiatingNodeChoiceGNb() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateConfigurationUpdateInitiatingNodeChoiceNgENb(ngENb *xnappducontentsv1.ConfigurationUpdatengeNb) (*xnappducontentsv1.ConfigurationUpdateInitiatingNodeChoice, error) {

	item := &xnappducontentsv1.ConfigurationUpdateInitiatingNodeChoice{
		ConfigurationUpdateInitiatingNodeChoice: &xnappducontentsv1.ConfigurationUpdateInitiatingNodeChoice_NgENb{
			NgENb: ngENb,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateConfigurationUpdateInitiatingNodeChoiceNgENb() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateConfigurationUpdateInitiatingNodeChoiceChoiceExtension(choiceExtension *xnappducontentsv1.ServedCellsToUpdateInitiatingNodeChoiceExtIes) (*xnappducontentsv1.ConfigurationUpdateInitiatingNodeChoice, error) {

	item := &xnappducontentsv1.ConfigurationUpdateInitiatingNodeChoice{
		ConfigurationUpdateInitiatingNodeChoice: &xnappducontentsv1.ConfigurationUpdateInitiatingNodeChoice_ChoiceExtension{
			ChoiceExtension: choiceExtension,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateConfigurationUpdateInitiatingNodeChoiceChoiceExtension() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateNgrannodeConfigurationUpdateAcknowledgeIEsValueIDRespondingNodeTypeConfigUpdateAck(IDRespondingNodeTypeConfigUpdateAck *xnappducontentsv1.RespondingNodeTypeConfigUpdateAck) (*xnappducontentsv1.NgrannodeConfigurationUpdateAcknowledgeIEsValue, error) {

	item := &xnappducontentsv1.NgrannodeConfigurationUpdateAcknowledgeIEsValue{
		NgrannodeConfigurationUpdateAcknowledgeIes: &xnappducontentsv1.NgrannodeConfigurationUpdateAcknowledgeIEsValue_IdRespondingNodeTypeConfigUpdateAck{
			IdRespondingNodeTypeConfigUpdateAck: IDRespondingNodeTypeConfigUpdateAck,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNgrannodeConfigurationUpdateAcknowledgeIEsValueIDRespondingNodeTypeConfigUpdateAck() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateNgrannodeConfigurationUpdateAcknowledgeIEsValueIDTnlaSetupList(IDTnlaSetupList *xnapiesv1.TnlaSetupList) (*xnappducontentsv1.NgrannodeConfigurationUpdateAcknowledgeIEsValue, error) {

	item := &xnappducontentsv1.NgrannodeConfigurationUpdateAcknowledgeIEsValue{
		NgrannodeConfigurationUpdateAcknowledgeIes: &xnappducontentsv1.NgrannodeConfigurationUpdateAcknowledgeIEsValue_IdTnlaSetupList{
			IdTnlaSetupList: IDTnlaSetupList,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNgrannodeConfigurationUpdateAcknowledgeIEsValueIDTnlaSetupList() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateNgrannodeConfigurationUpdateAcknowledgeIEsValueIDTnlaFailedToSetupList(IDTnlaFailedToSetupList *xnapiesv1.TnlaFailedToSetupList) (*xnappducontentsv1.NgrannodeConfigurationUpdateAcknowledgeIEsValue, error) {

	item := &xnappducontentsv1.NgrannodeConfigurationUpdateAcknowledgeIEsValue{
		NgrannodeConfigurationUpdateAcknowledgeIes: &xnappducontentsv1.NgrannodeConfigurationUpdateAcknowledgeIEsValue_IdTnlaFailedToSetupList{
			IdTnlaFailedToSetupList: IDTnlaFailedToSetupList,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNgrannodeConfigurationUpdateAcknowledgeIEsValueIDTnlaFailedToSetupList() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateNgrannodeConfigurationUpdateAcknowledgeIEsValueIDCriticalityDiagnostics(IDCriticalityDiagnostics *xnapiesv1.CriticalityDiagnostics) (*xnappducontentsv1.NgrannodeConfigurationUpdateAcknowledgeIEsValue, error) {

	item := &xnappducontentsv1.NgrannodeConfigurationUpdateAcknowledgeIEsValue{
		NgrannodeConfigurationUpdateAcknowledgeIes: &xnappducontentsv1.NgrannodeConfigurationUpdateAcknowledgeIEsValue_IdCriticalityDiagnostics{
			IdCriticalityDiagnostics: IDCriticalityDiagnostics,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNgrannodeConfigurationUpdateAcknowledgeIEsValueIDCriticalityDiagnostics() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateNgrannodeConfigurationUpdateAcknowledgeIEsValueIDInterfaceInstanceIndication(IDInterfaceInstanceIndication *xnapiesv1.InterfaceInstanceIndication) (*xnappducontentsv1.NgrannodeConfigurationUpdateAcknowledgeIEsValue, error) {

	item := &xnappducontentsv1.NgrannodeConfigurationUpdateAcknowledgeIEsValue{
		NgrannodeConfigurationUpdateAcknowledgeIes: &xnappducontentsv1.NgrannodeConfigurationUpdateAcknowledgeIEsValue_IdInterfaceInstanceIndication{
			IdInterfaceInstanceIndication: IDInterfaceInstanceIndication,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNgrannodeConfigurationUpdateAcknowledgeIEsValueIDInterfaceInstanceIndication() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateNgrannodeConfigurationUpdateAcknowledgeIEsValueIDTnlconfigurationInfo(IDTnlconfigurationInfo *xnapiesv1.TnlconfigurationInfo) (*xnappducontentsv1.NgrannodeConfigurationUpdateAcknowledgeIEsValue, error) {

	item := &xnappducontentsv1.NgrannodeConfigurationUpdateAcknowledgeIEsValue{
		NgrannodeConfigurationUpdateAcknowledgeIes: &xnappducontentsv1.NgrannodeConfigurationUpdateAcknowledgeIEsValue_IdTnlconfigurationInfo{
			IdTnlconfigurationInfo: IDTnlconfigurationInfo,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNgrannodeConfigurationUpdateAcknowledgeIEsValueIDTnlconfigurationInfo() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateRespondingNodeTypeConfigUpdateAckNgENb(ngENb *xnappducontentsv1.RespondingNodeTypeConfigUpdateAckngeNb) (*xnappducontentsv1.RespondingNodeTypeConfigUpdateAck, error) {

	item := &xnappducontentsv1.RespondingNodeTypeConfigUpdateAck{
		RespondingNodeTypeConfigUpdateAck: &xnappducontentsv1.RespondingNodeTypeConfigUpdateAck_NgENb{
			NgENb: ngENb,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRespondingNodeTypeConfigUpdateAckNgENb() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateRespondingNodeTypeConfigUpdateAckGNb(gNb *xnappducontentsv1.RespondingNodeTypeConfigUpdateAckgNb) (*xnappducontentsv1.RespondingNodeTypeConfigUpdateAck, error) {

	item := &xnappducontentsv1.RespondingNodeTypeConfigUpdateAck{
		RespondingNodeTypeConfigUpdateAck: &xnappducontentsv1.RespondingNodeTypeConfigUpdateAck_GNb{
			GNb: gNb,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRespondingNodeTypeConfigUpdateAckGNb() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateRespondingNodeTypeConfigUpdateAckChoiceExtension(choiceExtension *xnappducontentsv1.RespondingNodeTypeConfigUpdateAckExtIes) (*xnappducontentsv1.RespondingNodeTypeConfigUpdateAck, error) {

	item := &xnappducontentsv1.RespondingNodeTypeConfigUpdateAck{
		RespondingNodeTypeConfigUpdateAck: &xnappducontentsv1.RespondingNodeTypeConfigUpdateAck_ChoiceExtension{
			ChoiceExtension: choiceExtension,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRespondingNodeTypeConfigUpdateAckChoiceExtension() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateRespondingNodeTypeConfigUpdateAckngeNbExtIesExtensionIDListOfServedCellsEUtra(IDListOfServedCellsEUtra *xnapiesv1.ServedCellsEUTra) (*xnappducontentsv1.RespondingNodeTypeConfigUpdateAckngeNbExtIesExtension, error) {

	item := &xnappducontentsv1.RespondingNodeTypeConfigUpdateAckngeNbExtIesExtension{
		RespondingNodeTypeConfigUpdateAckNgENbExtIes: &xnappducontentsv1.RespondingNodeTypeConfigUpdateAckngeNbExtIesExtension_IdListOfServedCellsEUtra{
			IdListOfServedCellsEUtra: IDListOfServedCellsEUtra,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRespondingNodeTypeConfigUpdateAckngeNbExtIesExtensionIDListOfServedCellsEUtra() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateRespondingNodeTypeConfigUpdateAckngeNbExtIesExtensionIDPartialListIndicatorEutra(IDPartialListIndicatorEutra xnapiesv1.PartialListIndicator) (*xnappducontentsv1.RespondingNodeTypeConfigUpdateAckngeNbExtIesExtension, error) {

	item := &xnappducontentsv1.RespondingNodeTypeConfigUpdateAckngeNbExtIesExtension{
		RespondingNodeTypeConfigUpdateAckNgENbExtIes: &xnappducontentsv1.RespondingNodeTypeConfigUpdateAckngeNbExtIesExtension_IdPartialListIndicatorEutra{
			IdPartialListIndicatorEutra: IDPartialListIndicatorEutra,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRespondingNodeTypeConfigUpdateAckngeNbExtIesExtensionIDPartialListIndicatorEutra() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateRespondingNodeTypeConfigUpdateAckngeNbExtIesExtensionIDCellAndCapacityAssistanceInfoEutra(IDCellAndCapacityAssistanceInfoEutra *xnapiesv1.CellAndCapacityAssistanceInfoEUtra) (*xnappducontentsv1.RespondingNodeTypeConfigUpdateAckngeNbExtIesExtension, error) {

	item := &xnappducontentsv1.RespondingNodeTypeConfigUpdateAckngeNbExtIesExtension{
		RespondingNodeTypeConfigUpdateAckNgENbExtIes: &xnappducontentsv1.RespondingNodeTypeConfigUpdateAckngeNbExtIesExtension_IdCellAndCapacityAssistanceInfoEutra{
			IdCellAndCapacityAssistanceInfoEutra: IDCellAndCapacityAssistanceInfoEutra,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRespondingNodeTypeConfigUpdateAckngeNbExtIesExtensionIDCellAndCapacityAssistanceInfoEutra() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateRespondingNodeTypeConfigUpdateAckgNbExtIesExtensionIDPartialListIndicatorNr(IDPartialListIndicatorNr xnapiesv1.PartialListIndicator) (*xnappducontentsv1.RespondingNodeTypeConfigUpdateAckgNbExtIesExtension, error) {

	item := &xnappducontentsv1.RespondingNodeTypeConfigUpdateAckgNbExtIesExtension{
		RespondingNodeTypeConfigUpdateAckGNbExtIes: &xnappducontentsv1.RespondingNodeTypeConfigUpdateAckgNbExtIesExtension_IdPartialListIndicatorNr{
			IdPartialListIndicatorNr: IDPartialListIndicatorNr,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRespondingNodeTypeConfigUpdateAckgNbExtIesExtensionIDPartialListIndicatorNr() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateRespondingNodeTypeConfigUpdateAckgNbExtIesExtensionIDCellAndCapacityAssistanceInfoNr(IDCellAndCapacityAssistanceInfoNr *xnapiesv1.CellAndCapacityAssistanceInfoNR) (*xnappducontentsv1.RespondingNodeTypeConfigUpdateAckgNbExtIesExtension, error) {

	item := &xnappducontentsv1.RespondingNodeTypeConfigUpdateAckgNbExtIesExtension{
		RespondingNodeTypeConfigUpdateAckGNbExtIes: &xnappducontentsv1.RespondingNodeTypeConfigUpdateAckgNbExtIesExtension_IdCellAndCapacityAssistanceInfoNr{
			IdCellAndCapacityAssistanceInfoNr: IDCellAndCapacityAssistanceInfoNr,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRespondingNodeTypeConfigUpdateAckgNbExtIesExtensionIDCellAndCapacityAssistanceInfoNr() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateNgrannodeConfigurationUpdateFailureIEsValueIDCause(IDCause *xnapiesv1.Cause) (*xnappducontentsv1.NgrannodeConfigurationUpdateFailureIEsValue, error) {

	item := &xnappducontentsv1.NgrannodeConfigurationUpdateFailureIEsValue{
		NgrannodeConfigurationUpdateFailureIes: &xnappducontentsv1.NgrannodeConfigurationUpdateFailureIEsValue_IdCause{
			IdCause: IDCause,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNgrannodeConfigurationUpdateFailureIEsValueIDCause() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateNgrannodeConfigurationUpdateFailureIEsValueIDTimeToWait(IDTimeToWait xnapiesv1.TimeToWait) (*xnappducontentsv1.NgrannodeConfigurationUpdateFailureIEsValue, error) {

	item := &xnappducontentsv1.NgrannodeConfigurationUpdateFailureIEsValue{
		NgrannodeConfigurationUpdateFailureIes: &xnappducontentsv1.NgrannodeConfigurationUpdateFailureIEsValue_IdTimeToWait{
			IdTimeToWait: IDTimeToWait,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNgrannodeConfigurationUpdateFailureIEsValueIDTimeToWait() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateNgrannodeConfigurationUpdateFailureIEsValueIDCriticalityDiagnostics(IDCriticalityDiagnostics *xnapiesv1.CriticalityDiagnostics) (*xnappducontentsv1.NgrannodeConfigurationUpdateFailureIEsValue, error) {

	item := &xnappducontentsv1.NgrannodeConfigurationUpdateFailureIEsValue{
		NgrannodeConfigurationUpdateFailureIes: &xnappducontentsv1.NgrannodeConfigurationUpdateFailureIEsValue_IdCriticalityDiagnostics{
			IdCriticalityDiagnostics: IDCriticalityDiagnostics,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNgrannodeConfigurationUpdateFailureIEsValueIDCriticalityDiagnostics() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateNgrannodeConfigurationUpdateFailureIEsValueIDInterfaceInstanceIndication(IDInterfaceInstanceIndication *xnapiesv1.InterfaceInstanceIndication) (*xnappducontentsv1.NgrannodeConfigurationUpdateFailureIEsValue, error) {

	item := &xnappducontentsv1.NgrannodeConfigurationUpdateFailureIEsValue{
		NgrannodeConfigurationUpdateFailureIes: &xnappducontentsv1.NgrannodeConfigurationUpdateFailureIEsValue_IdInterfaceInstanceIndication{
			IdInterfaceInstanceIndication: IDInterfaceInstanceIndication,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateNgrannodeConfigurationUpdateFailureIEsValueIDInterfaceInstanceIndication() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateEUTraNRCellResourceCoordinationRequestIEsValueIDInitiatingNodeTypeResourceCoordRequest(IDInitiatingNodeTypeResourceCoordRequest *xnappducontentsv1.InitiatingNodeTypeResourceCoordRequest) (*xnappducontentsv1.EUTraNRCellResourceCoordinationRequestIEsValue, error) {

	item := &xnappducontentsv1.EUTraNRCellResourceCoordinationRequestIEsValue{
		EUtraNrCellResourceCoordinationRequestIes: &xnappducontentsv1.EUTraNRCellResourceCoordinationRequestIEsValue_IdInitiatingNodeTypeResourceCoordRequest{
			IdInitiatingNodeTypeResourceCoordRequest: IDInitiatingNodeTypeResourceCoordRequest,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateEUTraNRCellResourceCoordinationRequestIEsValueIDInitiatingNodeTypeResourceCoordRequest() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateEUTraNRCellResourceCoordinationRequestIEsValueIDInterfaceInstanceIndication(IDInterfaceInstanceIndication *xnapiesv1.InterfaceInstanceIndication) (*xnappducontentsv1.EUTraNRCellResourceCoordinationRequestIEsValue, error) {

	item := &xnappducontentsv1.EUTraNRCellResourceCoordinationRequestIEsValue{
		EUtraNrCellResourceCoordinationRequestIes: &xnappducontentsv1.EUTraNRCellResourceCoordinationRequestIEsValue_IdInterfaceInstanceIndication{
			IdInterfaceInstanceIndication: IDInterfaceInstanceIndication,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateEUTraNRCellResourceCoordinationRequestIEsValueIDInterfaceInstanceIndication() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateInitiatingNodeTypeResourceCoordRequestNgENb(ngENb *xnappducontentsv1.ResourceCoordRequestngeNbinitiated) (*xnappducontentsv1.InitiatingNodeTypeResourceCoordRequest, error) {

	item := &xnappducontentsv1.InitiatingNodeTypeResourceCoordRequest{
		InitiatingNodeTypeResourceCoordRequest: &xnappducontentsv1.InitiatingNodeTypeResourceCoordRequest_NgENb{
			NgENb: ngENb,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateInitiatingNodeTypeResourceCoordRequestNgENb() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateInitiatingNodeTypeResourceCoordRequestGNb(gNb *xnappducontentsv1.ResourceCoordRequestgNbinitiated) (*xnappducontentsv1.InitiatingNodeTypeResourceCoordRequest, error) {

	item := &xnappducontentsv1.InitiatingNodeTypeResourceCoordRequest{
		InitiatingNodeTypeResourceCoordRequest: &xnappducontentsv1.InitiatingNodeTypeResourceCoordRequest_GNb{
			GNb: gNb,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateInitiatingNodeTypeResourceCoordRequestGNb() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateInitiatingNodeTypeResourceCoordRequestChoiceExtension(choiceExtension *xnappducontentsv1.InitiatingNodeTypeResourceCoordRequestExtIes) (*xnappducontentsv1.InitiatingNodeTypeResourceCoordRequest, error) {

	item := &xnappducontentsv1.InitiatingNodeTypeResourceCoordRequest{
		InitiatingNodeTypeResourceCoordRequest: &xnappducontentsv1.InitiatingNodeTypeResourceCoordRequest_ChoiceExtension{
			ChoiceExtension: choiceExtension,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateInitiatingNodeTypeResourceCoordRequestChoiceExtension() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateEUTraNRCellResourceCoordinationResponseIEsValueIDRespondingNodeTypeResourceCoordResponse(IDRespondingNodeTypeResourceCoordResponse *xnappducontentsv1.RespondingNodeTypeResourceCoordResponse) (*xnappducontentsv1.EUTraNRCellResourceCoordinationResponseIEsValue, error) {

	item := &xnappducontentsv1.EUTraNRCellResourceCoordinationResponseIEsValue{
		EUtraNrCellResourceCoordinationResponseIes: &xnappducontentsv1.EUTraNRCellResourceCoordinationResponseIEsValue_IdRespondingNodeTypeResourceCoordResponse{
			IdRespondingNodeTypeResourceCoordResponse: IDRespondingNodeTypeResourceCoordResponse,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateEUTraNRCellResourceCoordinationResponseIEsValueIDRespondingNodeTypeResourceCoordResponse() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateEUTraNRCellResourceCoordinationResponseIEsValueIDInterfaceInstanceIndication(IDInterfaceInstanceIndication *xnapiesv1.InterfaceInstanceIndication) (*xnappducontentsv1.EUTraNRCellResourceCoordinationResponseIEsValue, error) {

	item := &xnappducontentsv1.EUTraNRCellResourceCoordinationResponseIEsValue{
		EUtraNrCellResourceCoordinationResponseIes: &xnappducontentsv1.EUTraNRCellResourceCoordinationResponseIEsValue_IdInterfaceInstanceIndication{
			IdInterfaceInstanceIndication: IDInterfaceInstanceIndication,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateEUTraNRCellResourceCoordinationResponseIEsValueIDInterfaceInstanceIndication() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateRespondingNodeTypeResourceCoordResponseNgENb(ngENb *xnappducontentsv1.ResourceCoordResponsengeNbinitiated) (*xnappducontentsv1.RespondingNodeTypeResourceCoordResponse, error) {

	item := &xnappducontentsv1.RespondingNodeTypeResourceCoordResponse{
		RespondingNodeTypeResourceCoordResponse: &xnappducontentsv1.RespondingNodeTypeResourceCoordResponse_NgENb{
			NgENb: ngENb,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRespondingNodeTypeResourceCoordResponseNgENb() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateRespondingNodeTypeResourceCoordResponseGNb(gNb *xnappducontentsv1.ResourceCoordResponsegNbinitiated) (*xnappducontentsv1.RespondingNodeTypeResourceCoordResponse, error) {

	item := &xnappducontentsv1.RespondingNodeTypeResourceCoordResponse{
		RespondingNodeTypeResourceCoordResponse: &xnappducontentsv1.RespondingNodeTypeResourceCoordResponse_GNb{
			GNb: gNb,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRespondingNodeTypeResourceCoordResponseGNb() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateRespondingNodeTypeResourceCoordResponseChoiceExtension(choiceExtension *xnappducontentsv1.RespondingNodeTypeResourceCoordResponseExtIes) (*xnappducontentsv1.RespondingNodeTypeResourceCoordResponse, error) {

	item := &xnappducontentsv1.RespondingNodeTypeResourceCoordResponse{
		RespondingNodeTypeResourceCoordResponse: &xnappducontentsv1.RespondingNodeTypeResourceCoordResponse_ChoiceExtension{
			ChoiceExtension: choiceExtension,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateRespondingNodeTypeResourceCoordResponseChoiceExtension() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSecondaryRatdataUsageReportIEsValueIDMNgRannodeUexnApID(IDMNgRannodeUexnApID *xnapiesv1.NgRAnnodeUexnApid) (*xnappducontentsv1.SecondaryRatdataUsageReportIEsValue, error) {

	item := &xnappducontentsv1.SecondaryRatdataUsageReportIEsValue{
		SecondaryRatdataUsageReportIes: &xnappducontentsv1.SecondaryRatdataUsageReportIEsValue_IdMNgRannodeUexnApid{
			IdMNgRannodeUexnApid: IDMNgRannodeUexnApID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSecondaryRatdataUsageReportIEsValueIDMNgRannodeUexnApID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSecondaryRatdataUsageReportIEsValueIDSNgRannodeUexnApID(IDSNgRannodeUexnApID *xnapiesv1.NgRAnnodeUexnApid) (*xnappducontentsv1.SecondaryRatdataUsageReportIEsValue, error) {

	item := &xnappducontentsv1.SecondaryRatdataUsageReportIEsValue{
		SecondaryRatdataUsageReportIes: &xnappducontentsv1.SecondaryRatdataUsageReportIEsValue_IdSNgRannodeUexnApid{
			IdSNgRannodeUexnApid: IDSNgRannodeUexnApID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSecondaryRatdataUsageReportIEsValueIDSNgRannodeUexnApID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateSecondaryRatdataUsageReportIEsValueIDPdusessionResourceSecondaryRatusageList(IDPdusessionResourceSecondaryRatusageList *xnapiesv1.PdusessionResourceSecondaryRatusageList) (*xnappducontentsv1.SecondaryRatdataUsageReportIEsValue, error) {

	item := &xnappducontentsv1.SecondaryRatdataUsageReportIEsValue{
		SecondaryRatdataUsageReportIes: &xnappducontentsv1.SecondaryRatdataUsageReportIEsValue_IdPdusessionResourceSecondaryRatusageList{
			IdPdusessionResourceSecondaryRatusageList: IDPdusessionResourceSecondaryRatusageList,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateSecondaryRatdataUsageReportIEsValueIDPdusessionResourceSecondaryRatusageList() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateXnRemovalRequestIEsValueIDGlobalNgRanNodeID(IDGlobalNgRanNodeID *xnapiesv1.GlobalNgRAnnodeID) (*xnappducontentsv1.XnRemovalRequestIEsValue, error) {

	item := &xnappducontentsv1.XnRemovalRequestIEsValue{
		XnRemovalRequestIes: &xnappducontentsv1.XnRemovalRequestIEsValue_IdGlobalNgRanNodeId{
			IdGlobalNgRanNodeId: IDGlobalNgRanNodeID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateXnRemovalRequestIEsValueIDGlobalNgRanNodeID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateXnRemovalRequestIEsValueIDXnRemovalThreshold(IDXnRemovalThreshold *xnapiesv1.XnBenefitValue) (*xnappducontentsv1.XnRemovalRequestIEsValue, error) {

	item := &xnappducontentsv1.XnRemovalRequestIEsValue{
		XnRemovalRequestIes: &xnappducontentsv1.XnRemovalRequestIEsValue_IdXnRemovalThreshold{
			IdXnRemovalThreshold: IDXnRemovalThreshold,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateXnRemovalRequestIEsValueIDXnRemovalThreshold() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateXnRemovalRequestIEsValueIDInterfaceInstanceIndication(IDInterfaceInstanceIndication *xnapiesv1.InterfaceInstanceIndication) (*xnappducontentsv1.XnRemovalRequestIEsValue, error) {

	item := &xnappducontentsv1.XnRemovalRequestIEsValue{
		XnRemovalRequestIes: &xnappducontentsv1.XnRemovalRequestIEsValue_IdInterfaceInstanceIndication{
			IdInterfaceInstanceIndication: IDInterfaceInstanceIndication,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateXnRemovalRequestIEsValueIDInterfaceInstanceIndication() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateXnRemovalResponseIEsValueIDGlobalNgRanNodeID(IDGlobalNgRanNodeID *xnapiesv1.GlobalNgRAnnodeID) (*xnappducontentsv1.XnRemovalResponseIEsValue, error) {

	item := &xnappducontentsv1.XnRemovalResponseIEsValue{
		XnRemovalResponseIes: &xnappducontentsv1.XnRemovalResponseIEsValue_IdGlobalNgRanNodeId{
			IdGlobalNgRanNodeId: IDGlobalNgRanNodeID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateXnRemovalResponseIEsValueIDGlobalNgRanNodeID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateXnRemovalResponseIEsValueIDCriticalityDiagnostics(IDCriticalityDiagnostics *xnapiesv1.CriticalityDiagnostics) (*xnappducontentsv1.XnRemovalResponseIEsValue, error) {

	item := &xnappducontentsv1.XnRemovalResponseIEsValue{
		XnRemovalResponseIes: &xnappducontentsv1.XnRemovalResponseIEsValue_IdCriticalityDiagnostics{
			IdCriticalityDiagnostics: IDCriticalityDiagnostics,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateXnRemovalResponseIEsValueIDCriticalityDiagnostics() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateXnRemovalResponseIEsValueIDInterfaceInstanceIndication(IDInterfaceInstanceIndication *xnapiesv1.InterfaceInstanceIndication) (*xnappducontentsv1.XnRemovalResponseIEsValue, error) {

	item := &xnappducontentsv1.XnRemovalResponseIEsValue{
		XnRemovalResponseIes: &xnappducontentsv1.XnRemovalResponseIEsValue_IdInterfaceInstanceIndication{
			IdInterfaceInstanceIndication: IDInterfaceInstanceIndication,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateXnRemovalResponseIEsValueIDInterfaceInstanceIndication() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateXnRemovalFailureIEsValueIDCause(IDCause *xnapiesv1.Cause) (*xnappducontentsv1.XnRemovalFailureIEsValue, error) {

	item := &xnappducontentsv1.XnRemovalFailureIEsValue{
		XnRemovalFailureIes: &xnappducontentsv1.XnRemovalFailureIEsValue_IdCause{
			IdCause: IDCause,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateXnRemovalFailureIEsValueIDCause() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateXnRemovalFailureIEsValueIDCriticalityDiagnostics(IDCriticalityDiagnostics *xnapiesv1.CriticalityDiagnostics) (*xnappducontentsv1.XnRemovalFailureIEsValue, error) {

	item := &xnappducontentsv1.XnRemovalFailureIEsValue{
		XnRemovalFailureIes: &xnappducontentsv1.XnRemovalFailureIEsValue_IdCriticalityDiagnostics{
			IdCriticalityDiagnostics: IDCriticalityDiagnostics,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateXnRemovalFailureIEsValueIDCriticalityDiagnostics() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateXnRemovalFailureIEsValueIDInterfaceInstanceIndication(IDInterfaceInstanceIndication *xnapiesv1.InterfaceInstanceIndication) (*xnappducontentsv1.XnRemovalFailureIEsValue, error) {

	item := &xnappducontentsv1.XnRemovalFailureIEsValue{
		XnRemovalFailureIes: &xnappducontentsv1.XnRemovalFailureIEsValue_IdInterfaceInstanceIndication{
			IdInterfaceInstanceIndication: IDInterfaceInstanceIndication,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateXnRemovalFailureIEsValueIDInterfaceInstanceIndication() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateCellActivationRequestIEsValueIDServedCellsToActivate(IDServedCellsToActivate *xnappducontentsv1.ServedCellsToActivate) (*xnappducontentsv1.CellActivationRequestIEsValue, error) {

	item := &xnappducontentsv1.CellActivationRequestIEsValue{
		CellActivationRequestIes: &xnappducontentsv1.CellActivationRequestIEsValue_IdServedCellsToActivate{
			IdServedCellsToActivate: IDServedCellsToActivate,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateCellActivationRequestIEsValueIDServedCellsToActivate() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateCellActivationRequestIEsValueIDActivationIDforCellActivation(IDActivationIDforCellActivation *xnapiesv1.ActivationIdforCellActivation) (*xnappducontentsv1.CellActivationRequestIEsValue, error) {

	item := &xnappducontentsv1.CellActivationRequestIEsValue{
		CellActivationRequestIes: &xnappducontentsv1.CellActivationRequestIEsValue_IdActivationIdforCellActivation{
			IdActivationIdforCellActivation: IDActivationIDforCellActivation,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateCellActivationRequestIEsValueIDActivationIDforCellActivation() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateCellActivationRequestIEsValueIDInterfaceInstanceIndication(IDInterfaceInstanceIndication *xnapiesv1.InterfaceInstanceIndication) (*xnappducontentsv1.CellActivationRequestIEsValue, error) {

	item := &xnappducontentsv1.CellActivationRequestIEsValue{
		CellActivationRequestIes: &xnappducontentsv1.CellActivationRequestIEsValue_IdInterfaceInstanceIndication{
			IdInterfaceInstanceIndication: IDInterfaceInstanceIndication,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateCellActivationRequestIEsValueIDInterfaceInstanceIndication() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateServedCellsToActivateNrCells(nrCells *xnappducontentsv1.NrcellsServedCellsToActivate) (*xnappducontentsv1.ServedCellsToActivate, error) {

	item := &xnappducontentsv1.ServedCellsToActivate{
		ServedCellsToActivate: &xnappducontentsv1.ServedCellsToActivate_NrCells{
			NrCells: nrCells,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateServedCellsToActivateNrCells() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateServedCellsToActivateEUtraCells(eUtraCells *xnappducontentsv1.EutracellsServedCellsToActivate) (*xnappducontentsv1.ServedCellsToActivate, error) {

	item := &xnappducontentsv1.ServedCellsToActivate{
		ServedCellsToActivate: &xnappducontentsv1.ServedCellsToActivate_EUtraCells{
			EUtraCells: eUtraCells,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateServedCellsToActivateEUtraCells() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateServedCellsToActivateChoiceExtension(choiceExtension *xnappducontentsv1.ServedCellsToActivateExtIes) (*xnappducontentsv1.ServedCellsToActivate, error) {

	item := &xnappducontentsv1.ServedCellsToActivate{
		ServedCellsToActivate: &xnappducontentsv1.ServedCellsToActivate_ChoiceExtension{
			ChoiceExtension: choiceExtension,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateServedCellsToActivateChoiceExtension() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateCellActivationResponseIEsValueIDActivatedServedCells(IDActivatedServedCells *xnappducontentsv1.ActivatedServedCells) (*xnappducontentsv1.CellActivationResponseIEsValue, error) {

	item := &xnappducontentsv1.CellActivationResponseIEsValue{
		CellActivationResponseIes: &xnappducontentsv1.CellActivationResponseIEsValue_IdActivatedServedCells{
			IdActivatedServedCells: IDActivatedServedCells,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateCellActivationResponseIEsValueIDActivatedServedCells() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateCellActivationResponseIEsValueIDActivationIDforCellActivation(IDActivationIDforCellActivation *xnapiesv1.ActivationIdforCellActivation) (*xnappducontentsv1.CellActivationResponseIEsValue, error) {

	item := &xnappducontentsv1.CellActivationResponseIEsValue{
		CellActivationResponseIes: &xnappducontentsv1.CellActivationResponseIEsValue_IdActivationIdforCellActivation{
			IdActivationIdforCellActivation: IDActivationIDforCellActivation,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateCellActivationResponseIEsValueIDActivationIDforCellActivation() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateCellActivationResponseIEsValueIDCriticalityDiagnostics(IDCriticalityDiagnostics *xnapiesv1.CriticalityDiagnostics) (*xnappducontentsv1.CellActivationResponseIEsValue, error) {

	item := &xnappducontentsv1.CellActivationResponseIEsValue{
		CellActivationResponseIes: &xnappducontentsv1.CellActivationResponseIEsValue_IdCriticalityDiagnostics{
			IdCriticalityDiagnostics: IDCriticalityDiagnostics,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateCellActivationResponseIEsValueIDCriticalityDiagnostics() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateCellActivationResponseIEsValueIDInterfaceInstanceIndication(IDInterfaceInstanceIndication *xnapiesv1.InterfaceInstanceIndication) (*xnappducontentsv1.CellActivationResponseIEsValue, error) {

	item := &xnappducontentsv1.CellActivationResponseIEsValue{
		CellActivationResponseIes: &xnappducontentsv1.CellActivationResponseIEsValue_IdInterfaceInstanceIndication{
			IdInterfaceInstanceIndication: IDInterfaceInstanceIndication,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateCellActivationResponseIEsValueIDInterfaceInstanceIndication() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateActivatedServedCellsNrCells(nrCells *xnappducontentsv1.NrcellsActivatedServedCells) (*xnappducontentsv1.ActivatedServedCells, error) {

	item := &xnappducontentsv1.ActivatedServedCells{
		ActivatedServedCells: &xnappducontentsv1.ActivatedServedCells_NrCells{
			NrCells: nrCells,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateActivatedServedCellsNrCells() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateActivatedServedCellsEUtraCells(eUtraCells *xnappducontentsv1.EutracellsActivatedServedCells) (*xnappducontentsv1.ActivatedServedCells, error) {

	item := &xnappducontentsv1.ActivatedServedCells{
		ActivatedServedCells: &xnappducontentsv1.ActivatedServedCells_EUtraCells{
			EUtraCells: eUtraCells,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateActivatedServedCellsEUtraCells() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateActivatedServedCellsChoiceExtension(choiceExtension *xnappducontentsv1.ActivatedServedCellsExtIes) (*xnappducontentsv1.ActivatedServedCells, error) {

	item := &xnappducontentsv1.ActivatedServedCells{
		ActivatedServedCells: &xnappducontentsv1.ActivatedServedCells_ChoiceExtension{
			ChoiceExtension: choiceExtension,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateActivatedServedCellsChoiceExtension() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateCellActivationFailureIEsValueIDActivationIDforCellActivation(IDActivationIDforCellActivation *xnapiesv1.ActivationIdforCellActivation) (*xnappducontentsv1.CellActivationFailureIEsValue, error) {

	item := &xnappducontentsv1.CellActivationFailureIEsValue{
		CellActivationFailureIes: &xnappducontentsv1.CellActivationFailureIEsValue_IdActivationIdforCellActivation{
			IdActivationIdforCellActivation: IDActivationIDforCellActivation,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateCellActivationFailureIEsValueIDActivationIDforCellActivation() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateCellActivationFailureIEsValueIDCause(IDCause *xnapiesv1.Cause) (*xnappducontentsv1.CellActivationFailureIEsValue, error) {

	item := &xnappducontentsv1.CellActivationFailureIEsValue{
		CellActivationFailureIes: &xnappducontentsv1.CellActivationFailureIEsValue_IdCause{
			IdCause: IDCause,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateCellActivationFailureIEsValueIDCause() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateCellActivationFailureIEsValueIDCriticalityDiagnostics(IDCriticalityDiagnostics *xnapiesv1.CriticalityDiagnostics) (*xnappducontentsv1.CellActivationFailureIEsValue, error) {

	item := &xnappducontentsv1.CellActivationFailureIEsValue{
		CellActivationFailureIes: &xnappducontentsv1.CellActivationFailureIEsValue_IdCriticalityDiagnostics{
			IdCriticalityDiagnostics: IDCriticalityDiagnostics,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateCellActivationFailureIEsValueIDCriticalityDiagnostics() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateCellActivationFailureIEsValueIDInterfaceInstanceIndication(IDInterfaceInstanceIndication *xnapiesv1.InterfaceInstanceIndication) (*xnappducontentsv1.CellActivationFailureIEsValue, error) {

	item := &xnappducontentsv1.CellActivationFailureIEsValue{
		CellActivationFailureIes: &xnappducontentsv1.CellActivationFailureIEsValue_IdInterfaceInstanceIndication{
			IdInterfaceInstanceIndication: IDInterfaceInstanceIndication,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateCellActivationFailureIEsValueIDInterfaceInstanceIndication() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateResetRequestIEsValueIDResetRequestTypeInfo(IDResetRequestTypeInfo *xnapiesv1.ResetRequestTypeInfo) (*xnappducontentsv1.ResetRequestIEsValue, error) {

	item := &xnappducontentsv1.ResetRequestIEsValue{
		ResetRequestIes: &xnappducontentsv1.ResetRequestIEsValue_IdResetRequestTypeInfo{
			IdResetRequestTypeInfo: IDResetRequestTypeInfo,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateResetRequestIEsValueIDResetRequestTypeInfo() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateResetRequestIEsValueIDCause(IDCause *xnapiesv1.Cause) (*xnappducontentsv1.ResetRequestIEsValue, error) {

	item := &xnappducontentsv1.ResetRequestIEsValue{
		ResetRequestIes: &xnappducontentsv1.ResetRequestIEsValue_IdCause{
			IdCause: IDCause,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateResetRequestIEsValueIDCause() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateResetRequestIEsValueIDInterfaceInstanceIndication(IDInterfaceInstanceIndication *xnapiesv1.InterfaceInstanceIndication) (*xnappducontentsv1.ResetRequestIEsValue, error) {

	item := &xnappducontentsv1.ResetRequestIEsValue{
		ResetRequestIes: &xnappducontentsv1.ResetRequestIEsValue_IdInterfaceInstanceIndication{
			IdInterfaceInstanceIndication: IDInterfaceInstanceIndication,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateResetRequestIEsValueIDInterfaceInstanceIndication() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateResetResponseIEsValueIDResetResponseTypeInfo(IDResetResponseTypeInfo *xnapiesv1.ResetResponseTypeInfo) (*xnappducontentsv1.ResetResponseIEsValue, error) {

	item := &xnappducontentsv1.ResetResponseIEsValue{
		ResetResponseIes: &xnappducontentsv1.ResetResponseIEsValue_IdResetResponseTypeInfo{
			IdResetResponseTypeInfo: IDResetResponseTypeInfo,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateResetResponseIEsValueIDResetResponseTypeInfo() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateResetResponseIEsValueIDCriticalityDiagnostics(IDCriticalityDiagnostics *xnapiesv1.CriticalityDiagnostics) (*xnappducontentsv1.ResetResponseIEsValue, error) {

	item := &xnappducontentsv1.ResetResponseIEsValue{
		ResetResponseIes: &xnappducontentsv1.ResetResponseIEsValue_IdCriticalityDiagnostics{
			IdCriticalityDiagnostics: IDCriticalityDiagnostics,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateResetResponseIEsValueIDCriticalityDiagnostics() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateResetResponseIEsValueIDInterfaceInstanceIndication(IDInterfaceInstanceIndication *xnapiesv1.InterfaceInstanceIndication) (*xnappducontentsv1.ResetResponseIEsValue, error) {

	item := &xnappducontentsv1.ResetResponseIEsValue{
		ResetResponseIes: &xnappducontentsv1.ResetResponseIEsValue_IdInterfaceInstanceIndication{
			IdInterfaceInstanceIndication: IDInterfaceInstanceIndication,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateResetResponseIEsValueIDInterfaceInstanceIndication() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateErrorIndicationIEsValueIDOldNgRannodeUexnApID(IDOldNgRannodeUexnApID *xnapiesv1.NgRAnnodeUexnApid) (*xnappducontentsv1.ErrorIndicationIEsValue, error) {

	item := &xnappducontentsv1.ErrorIndicationIEsValue{
		ErrorIndicationIes: &xnappducontentsv1.ErrorIndicationIEsValue_IdOldNgRannodeUexnApid{
			IdOldNgRannodeUexnApid: IDOldNgRannodeUexnApID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateErrorIndicationIEsValueIDOldNgRannodeUexnApID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateErrorIndicationIEsValueIDNewNgRannodeUexnApID(IDNewNgRannodeUexnApID *xnapiesv1.NgRAnnodeUexnApid) (*xnappducontentsv1.ErrorIndicationIEsValue, error) {

	item := &xnappducontentsv1.ErrorIndicationIEsValue{
		ErrorIndicationIes: &xnappducontentsv1.ErrorIndicationIEsValue_IdNewNgRannodeUexnApid{
			IdNewNgRannodeUexnApid: IDNewNgRannodeUexnApID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateErrorIndicationIEsValueIDNewNgRannodeUexnApID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateErrorIndicationIEsValueIDCause(IDCause *xnapiesv1.Cause) (*xnappducontentsv1.ErrorIndicationIEsValue, error) {

	item := &xnappducontentsv1.ErrorIndicationIEsValue{
		ErrorIndicationIes: &xnappducontentsv1.ErrorIndicationIEsValue_IdCause{
			IdCause: IDCause,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateErrorIndicationIEsValueIDCause() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateErrorIndicationIEsValueIDCriticalityDiagnostics(IDCriticalityDiagnostics *xnapiesv1.CriticalityDiagnostics) (*xnappducontentsv1.ErrorIndicationIEsValue, error) {

	item := &xnappducontentsv1.ErrorIndicationIEsValue{
		ErrorIndicationIes: &xnappducontentsv1.ErrorIndicationIEsValue_IdCriticalityDiagnostics{
			IdCriticalityDiagnostics: IDCriticalityDiagnostics,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateErrorIndicationIEsValueIDCriticalityDiagnostics() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateErrorIndicationIEsValueIDInterfaceInstanceIndication(IDInterfaceInstanceIndication *xnapiesv1.InterfaceInstanceIndication) (*xnappducontentsv1.ErrorIndicationIEsValue, error) {

	item := &xnappducontentsv1.ErrorIndicationIEsValue{
		ErrorIndicationIes: &xnappducontentsv1.ErrorIndicationIEsValue_IdInterfaceInstanceIndication{
			IdInterfaceInstanceIndication: IDInterfaceInstanceIndication,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateErrorIndicationIEsValueIDInterfaceInstanceIndication() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateTraceStartIesValueIDMNgRannodeUexnApID(IDMNgRannodeUexnApID *xnapiesv1.NgRAnnodeUexnApid) (*xnappducontentsv1.TraceStartIesValue, error) {

	item := &xnappducontentsv1.TraceStartIesValue{
		TraceStartIes: &xnappducontentsv1.TraceStartIesValue_IdMNgRannodeUexnApid{
			IdMNgRannodeUexnApid: IDMNgRannodeUexnApID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateTraceStartIesValueIDMNgRannodeUexnApID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateTraceStartIesValueIDSNgRannodeUexnApID(IDSNgRannodeUexnApID *xnapiesv1.NgRAnnodeUexnApid) (*xnappducontentsv1.TraceStartIesValue, error) {

	item := &xnappducontentsv1.TraceStartIesValue{
		TraceStartIes: &xnappducontentsv1.TraceStartIesValue_IdSNgRannodeUexnApid{
			IdSNgRannodeUexnApid: IDSNgRannodeUexnApID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateTraceStartIesValueIDSNgRannodeUexnApID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateTraceStartIesValueIDTraceActivation(IDTraceActivation *xnapiesv1.TraceActivation) (*xnappducontentsv1.TraceStartIesValue, error) {

	item := &xnappducontentsv1.TraceStartIesValue{
		TraceStartIes: &xnappducontentsv1.TraceStartIesValue_IdTraceActivation{
			IdTraceActivation: IDTraceActivation,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateTraceStartIesValueIDTraceActivation() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateDeactivateTraceIesValueIDMNgRannodeUexnApID(IDMNgRannodeUexnApID *xnapiesv1.NgRAnnodeUexnApid) (*xnappducontentsv1.DeactivateTraceIesValue, error) {

	item := &xnappducontentsv1.DeactivateTraceIesValue{
		DeactivateTraceIes: &xnappducontentsv1.DeactivateTraceIesValue_IdMNgRannodeUexnApid{
			IdMNgRannodeUexnApid: IDMNgRannodeUexnApID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateDeactivateTraceIesValueIDMNgRannodeUexnApID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateDeactivateTraceIesValueIDSNgRannodeUexnApID(IDSNgRannodeUexnApID *xnapiesv1.NgRAnnodeUexnApid) (*xnappducontentsv1.DeactivateTraceIesValue, error) {

	item := &xnappducontentsv1.DeactivateTraceIesValue{
		DeactivateTraceIes: &xnappducontentsv1.DeactivateTraceIesValue_IdSNgRannodeUexnApid{
			IdSNgRannodeUexnApid: IDSNgRannodeUexnApID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateDeactivateTraceIesValueIDSNgRannodeUexnApID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateDeactivateTraceIesValueIDNgRantraceID(IDNgRantraceID *xnapiesv1.NgRAntraceId) (*xnappducontentsv1.DeactivateTraceIesValue, error) {

	item := &xnappducontentsv1.DeactivateTraceIesValue{
		DeactivateTraceIes: &xnappducontentsv1.DeactivateTraceIesValue_IdNgRantraceId{
			IdNgRantraceId: IDNgRantraceID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateDeactivateTraceIesValueIDNgRantraceID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateHandoverReportIEsValueIDHandoverReportType(IDHandoverReportType xnapiesv1.HandoverReportType) (*xnappducontentsv1.HandoverReportIEsValue, error) {

	item := &xnappducontentsv1.HandoverReportIEsValue{
		HandoverReportIes: &xnappducontentsv1.HandoverReportIEsValue_IdHandoverReportType{
			IdHandoverReportType: IDHandoverReportType,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateHandoverReportIEsValueIDHandoverReportType() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateHandoverReportIEsValueIDHandoverCause(IDHandoverCause *xnapiesv1.Cause) (*xnappducontentsv1.HandoverReportIEsValue, error) {

	item := &xnappducontentsv1.HandoverReportIEsValue{
		HandoverReportIes: &xnappducontentsv1.HandoverReportIEsValue_IdHandoverCause{
			IdHandoverCause: IDHandoverCause,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateHandoverReportIEsValueIDHandoverCause() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateHandoverReportIEsValueIDSourceCellCgi(IDSourceCellCgi *xnapiesv1.GlobalNgRAncellID) (*xnappducontentsv1.HandoverReportIEsValue, error) {

	item := &xnappducontentsv1.HandoverReportIEsValue{
		HandoverReportIes: &xnappducontentsv1.HandoverReportIEsValue_IdSourceCellCgi{
			IdSourceCellCgi: IDSourceCellCgi,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateHandoverReportIEsValueIDSourceCellCgi() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateHandoverReportIEsValueIDTargetCellCgi(IDTargetCellCgi *xnapiesv1.GlobalNgRAncellID) (*xnappducontentsv1.HandoverReportIEsValue, error) {

	item := &xnappducontentsv1.HandoverReportIEsValue{
		HandoverReportIes: &xnappducontentsv1.HandoverReportIEsValue_IdTargetCellCgi{
			IdTargetCellCgi: IDTargetCellCgi,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateHandoverReportIEsValueIDTargetCellCgi() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateHandoverReportIEsValueIDReEstablishmentCellCgi(IDReEstablishmentCellCgi *xnapiesv1.GlobalCellID) (*xnappducontentsv1.HandoverReportIEsValue, error) {

	item := &xnappducontentsv1.HandoverReportIEsValue{
		HandoverReportIes: &xnappducontentsv1.HandoverReportIEsValue_IdReEstablishmentCellCgi{
			IdReEstablishmentCellCgi: IDReEstablishmentCellCgi,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateHandoverReportIEsValueIDReEstablishmentCellCgi() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateHandoverReportIEsValueIDTargetCellinEutran(IDTargetCellinEutran *xnapiesv1.TargetCellinEutran) (*xnappducontentsv1.HandoverReportIEsValue, error) {

	item := &xnappducontentsv1.HandoverReportIEsValue{
		HandoverReportIes: &xnappducontentsv1.HandoverReportIEsValue_IdTargetCellinEutran{
			IdTargetCellinEutran: IDTargetCellinEutran,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateHandoverReportIEsValueIDTargetCellinEutran() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateHandoverReportIEsValueIDSourceCellCrnti(IDSourceCellCrnti *xnapiesv1.CRNti) (*xnappducontentsv1.HandoverReportIEsValue, error) {

	item := &xnappducontentsv1.HandoverReportIEsValue{
		HandoverReportIes: &xnappducontentsv1.HandoverReportIEsValue_IdSourceCellCrnti{
			IdSourceCellCrnti: IDSourceCellCrnti,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateHandoverReportIEsValueIDSourceCellCrnti() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateHandoverReportIEsValueIDMobilityInformation(IDMobilityInformation *xnapiesv1.MobilityInformation) (*xnappducontentsv1.HandoverReportIEsValue, error) {

	item := &xnappducontentsv1.HandoverReportIEsValue{
		HandoverReportIes: &xnappducontentsv1.HandoverReportIEsValue_IdMobilityInformation{
			IdMobilityInformation: IDMobilityInformation,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateHandoverReportIEsValueIDMobilityInformation() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateHandoverReportIEsValueIDUerlfreportContainer(IDUerlfreportContainer *xnapiesv1.UerlfreportContainer) (*xnappducontentsv1.HandoverReportIEsValue, error) {

	item := &xnappducontentsv1.HandoverReportIEsValue{
		HandoverReportIes: &xnappducontentsv1.HandoverReportIEsValue_IdUerlfreportContainer{
			IdUerlfreportContainer: IDUerlfreportContainer,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateHandoverReportIEsValueIDUerlfreportContainer() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateResourceStatusRequestIEsValueIDNgranNode1MeasurementID(IDNgranNode1MeasurementID *xnapiesv1.MeasurementID) (*xnappducontentsv1.ResourceStatusRequestIEsValue, error) {

	item := &xnappducontentsv1.ResourceStatusRequestIEsValue{
		ResourceStatusRequestIes: &xnappducontentsv1.ResourceStatusRequestIEsValue_IdNgranNode1MeasurementId{
			IdNgranNode1MeasurementId: IDNgranNode1MeasurementID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateResourceStatusRequestIEsValueIDNgranNode1MeasurementID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateResourceStatusRequestIEsValueIDNgranNode2MeasurementID(IDNgranNode2MeasurementID *xnapiesv1.MeasurementID) (*xnappducontentsv1.ResourceStatusRequestIEsValue, error) {

	item := &xnappducontentsv1.ResourceStatusRequestIEsValue{
		ResourceStatusRequestIes: &xnappducontentsv1.ResourceStatusRequestIEsValue_IdNgranNode2MeasurementId{
			IdNgranNode2MeasurementId: IDNgranNode2MeasurementID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateResourceStatusRequestIEsValueIDNgranNode2MeasurementID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateResourceStatusRequestIEsValueIDRegistrationRequest(IDRegistrationRequest xnapiesv1.RegistrationRequest) (*xnappducontentsv1.ResourceStatusRequestIEsValue, error) {

	item := &xnappducontentsv1.ResourceStatusRequestIEsValue{
		ResourceStatusRequestIes: &xnappducontentsv1.ResourceStatusRequestIEsValue_IdRegistrationRequest{
			IdRegistrationRequest: IDRegistrationRequest,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateResourceStatusRequestIEsValueIDRegistrationRequest() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateResourceStatusRequestIEsValueIDReportCharacteristics(IDReportCharacteristics *xnapiesv1.ReportCharacteristics) (*xnappducontentsv1.ResourceStatusRequestIEsValue, error) {

	item := &xnappducontentsv1.ResourceStatusRequestIEsValue{
		ResourceStatusRequestIes: &xnappducontentsv1.ResourceStatusRequestIEsValue_IdReportCharacteristics{
			IdReportCharacteristics: IDReportCharacteristics,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateResourceStatusRequestIEsValueIDReportCharacteristics() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateResourceStatusRequestIEsValueIDCellToReport(IDCellToReport *xnapiesv1.CellToReport) (*xnappducontentsv1.ResourceStatusRequestIEsValue, error) {

	item := &xnappducontentsv1.ResourceStatusRequestIEsValue{
		ResourceStatusRequestIes: &xnappducontentsv1.ResourceStatusRequestIEsValue_IdCellToReport{
			IdCellToReport: IDCellToReport,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateResourceStatusRequestIEsValueIDCellToReport() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateResourceStatusRequestIEsValueIDReportingPeriodicity(IDReportingPeriodicity xnapiesv1.ReportingPeriodicity) (*xnappducontentsv1.ResourceStatusRequestIEsValue, error) {

	item := &xnappducontentsv1.ResourceStatusRequestIEsValue{
		ResourceStatusRequestIes: &xnappducontentsv1.ResourceStatusRequestIEsValue_IdReportingPeriodicity{
			IdReportingPeriodicity: IDReportingPeriodicity,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateResourceStatusRequestIEsValueIDReportingPeriodicity() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateResourceStatusResponseIEsValueIDNgranNode1MeasurementID(IDNgranNode1MeasurementID *xnapiesv1.MeasurementID) (*xnappducontentsv1.ResourceStatusResponseIEsValue, error) {

	item := &xnappducontentsv1.ResourceStatusResponseIEsValue{
		ResourceStatusResponseIes: &xnappducontentsv1.ResourceStatusResponseIEsValue_IdNgranNode1MeasurementId{
			IdNgranNode1MeasurementId: IDNgranNode1MeasurementID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateResourceStatusResponseIEsValueIDNgranNode1MeasurementID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateResourceStatusResponseIEsValueIDNgranNode2MeasurementID(IDNgranNode2MeasurementID *xnapiesv1.MeasurementID) (*xnappducontentsv1.ResourceStatusResponseIEsValue, error) {

	item := &xnappducontentsv1.ResourceStatusResponseIEsValue{
		ResourceStatusResponseIes: &xnappducontentsv1.ResourceStatusResponseIEsValue_IdNgranNode2MeasurementId{
			IdNgranNode2MeasurementId: IDNgranNode2MeasurementID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateResourceStatusResponseIEsValueIDNgranNode2MeasurementID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateResourceStatusResponseIEsValueIDCriticalityDiagnostics(IDCriticalityDiagnostics *xnapiesv1.CriticalityDiagnostics) (*xnappducontentsv1.ResourceStatusResponseIEsValue, error) {

	item := &xnappducontentsv1.ResourceStatusResponseIEsValue{
		ResourceStatusResponseIes: &xnappducontentsv1.ResourceStatusResponseIEsValue_IdCriticalityDiagnostics{
			IdCriticalityDiagnostics: IDCriticalityDiagnostics,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateResourceStatusResponseIEsValueIDCriticalityDiagnostics() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateResourceStatusFailureIEsValueIDNgranNode1MeasurementID(IDNgranNode1MeasurementID *xnapiesv1.MeasurementID) (*xnappducontentsv1.ResourceStatusFailureIEsValue, error) {

	item := &xnappducontentsv1.ResourceStatusFailureIEsValue{
		ResourceStatusFailureIes: &xnappducontentsv1.ResourceStatusFailureIEsValue_IdNgranNode1MeasurementId{
			IdNgranNode1MeasurementId: IDNgranNode1MeasurementID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateResourceStatusFailureIEsValueIDNgranNode1MeasurementID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateResourceStatusFailureIEsValueIDNgranNode2MeasurementID(IDNgranNode2MeasurementID *xnapiesv1.MeasurementID) (*xnappducontentsv1.ResourceStatusFailureIEsValue, error) {

	item := &xnappducontentsv1.ResourceStatusFailureIEsValue{
		ResourceStatusFailureIes: &xnappducontentsv1.ResourceStatusFailureIEsValue_IdNgranNode2MeasurementId{
			IdNgranNode2MeasurementId: IDNgranNode2MeasurementID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateResourceStatusFailureIEsValueIDNgranNode2MeasurementID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateResourceStatusFailureIEsValueIDCause(IDCause *xnapiesv1.Cause) (*xnappducontentsv1.ResourceStatusFailureIEsValue, error) {

	item := &xnappducontentsv1.ResourceStatusFailureIEsValue{
		ResourceStatusFailureIes: &xnappducontentsv1.ResourceStatusFailureIEsValue_IdCause{
			IdCause: IDCause,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateResourceStatusFailureIEsValueIDCause() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateResourceStatusFailureIEsValueIDCriticalityDiagnostics(IDCriticalityDiagnostics *xnapiesv1.CriticalityDiagnostics) (*xnappducontentsv1.ResourceStatusFailureIEsValue, error) {

	item := &xnappducontentsv1.ResourceStatusFailureIEsValue{
		ResourceStatusFailureIes: &xnappducontentsv1.ResourceStatusFailureIEsValue_IdCriticalityDiagnostics{
			IdCriticalityDiagnostics: IDCriticalityDiagnostics,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateResourceStatusFailureIEsValueIDCriticalityDiagnostics() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateResourceStatusUpdateIEsValueIDNgranNode1MeasurementID(IDNgranNode1MeasurementID *xnapiesv1.MeasurementID) (*xnappducontentsv1.ResourceStatusUpdateIEsValue, error) {

	item := &xnappducontentsv1.ResourceStatusUpdateIEsValue{
		ResourceStatusUpdateIes: &xnappducontentsv1.ResourceStatusUpdateIEsValue_IdNgranNode1MeasurementId{
			IdNgranNode1MeasurementId: IDNgranNode1MeasurementID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateResourceStatusUpdateIEsValueIDNgranNode1MeasurementID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateResourceStatusUpdateIEsValueIDNgranNode2MeasurementID(IDNgranNode2MeasurementID *xnapiesv1.MeasurementID) (*xnappducontentsv1.ResourceStatusUpdateIEsValue, error) {

	item := &xnappducontentsv1.ResourceStatusUpdateIEsValue{
		ResourceStatusUpdateIes: &xnappducontentsv1.ResourceStatusUpdateIEsValue_IdNgranNode2MeasurementId{
			IdNgranNode2MeasurementId: IDNgranNode2MeasurementID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateResourceStatusUpdateIEsValueIDNgranNode2MeasurementID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateResourceStatusUpdateIEsValueIDCellMeasurementResult(IDCellMeasurementResult *xnapiesv1.CellMeasurementResult) (*xnappducontentsv1.ResourceStatusUpdateIEsValue, error) {

	item := &xnappducontentsv1.ResourceStatusUpdateIEsValue{
		ResourceStatusUpdateIes: &xnappducontentsv1.ResourceStatusUpdateIEsValue_IdCellMeasurementResult{
			IdCellMeasurementResult: IDCellMeasurementResult,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateResourceStatusUpdateIEsValueIDCellMeasurementResult() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateMobilityChangeRequestIEsValueIDNgRannode1CellID(IDNgRannode1CellID *xnapiesv1.GlobalNgRAncellID) (*xnappducontentsv1.MobilityChangeRequestIEsValue, error) {

	item := &xnappducontentsv1.MobilityChangeRequestIEsValue{
		MobilityChangeRequestIes: &xnappducontentsv1.MobilityChangeRequestIEsValue_IdNgRannode1CellId{
			IdNgRannode1CellId: IDNgRannode1CellID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateMobilityChangeRequestIEsValueIDNgRannode1CellID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateMobilityChangeRequestIEsValueIDNgRannode2CellID(IDNgRannode2CellID *xnapiesv1.GlobalNgRAncellID) (*xnappducontentsv1.MobilityChangeRequestIEsValue, error) {

	item := &xnappducontentsv1.MobilityChangeRequestIEsValue{
		MobilityChangeRequestIes: &xnappducontentsv1.MobilityChangeRequestIEsValue_IdNgRannode2CellId{
			IdNgRannode2CellId: IDNgRannode2CellID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateMobilityChangeRequestIEsValueIDNgRannode2CellID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateMobilityChangeRequestIEsValueIDNgRannode1MobilityParameters(IDNgRannode1MobilityParameters *xnapiesv1.MobilityParametersInformation) (*xnappducontentsv1.MobilityChangeRequestIEsValue, error) {

	item := &xnappducontentsv1.MobilityChangeRequestIEsValue{
		MobilityChangeRequestIes: &xnappducontentsv1.MobilityChangeRequestIEsValue_IdNgRannode1MobilityParameters{
			IdNgRannode1MobilityParameters: IDNgRannode1MobilityParameters,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateMobilityChangeRequestIEsValueIDNgRannode1MobilityParameters() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateMobilityChangeRequestIEsValueIDNgRannode2ProposedMobilityParameters(IDNgRannode2ProposedMobilityParameters *xnapiesv1.MobilityParametersInformation) (*xnappducontentsv1.MobilityChangeRequestIEsValue, error) {

	item := &xnappducontentsv1.MobilityChangeRequestIEsValue{
		MobilityChangeRequestIes: &xnappducontentsv1.MobilityChangeRequestIEsValue_IdNgRannode2ProposedMobilityParameters{
			IdNgRannode2ProposedMobilityParameters: IDNgRannode2ProposedMobilityParameters,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateMobilityChangeRequestIEsValueIDNgRannode2ProposedMobilityParameters() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateMobilityChangeRequestIEsValueIDCause(IDCause *xnapiesv1.Cause) (*xnappducontentsv1.MobilityChangeRequestIEsValue, error) {

	item := &xnappducontentsv1.MobilityChangeRequestIEsValue{
		MobilityChangeRequestIes: &xnappducontentsv1.MobilityChangeRequestIEsValue_IdCause{
			IdCause: IDCause,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateMobilityChangeRequestIEsValueIDCause() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateMobilityChangeAcknowledgeIEsValueIDNgRannode1CellID(IDNgRannode1CellID *xnapiesv1.GlobalNgRAncellID) (*xnappducontentsv1.MobilityChangeAcknowledgeIEsValue, error) {

	item := &xnappducontentsv1.MobilityChangeAcknowledgeIEsValue{
		MobilityChangeAcknowledgeIes: &xnappducontentsv1.MobilityChangeAcknowledgeIEsValue_IdNgRannode1CellId{
			IdNgRannode1CellId: IDNgRannode1CellID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateMobilityChangeAcknowledgeIEsValueIDNgRannode1CellID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateMobilityChangeAcknowledgeIEsValueIDNgRannode2CellID(IDNgRannode2CellID *xnapiesv1.GlobalNgRAncellID) (*xnappducontentsv1.MobilityChangeAcknowledgeIEsValue, error) {

	item := &xnappducontentsv1.MobilityChangeAcknowledgeIEsValue{
		MobilityChangeAcknowledgeIes: &xnappducontentsv1.MobilityChangeAcknowledgeIEsValue_IdNgRannode2CellId{
			IdNgRannode2CellId: IDNgRannode2CellID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateMobilityChangeAcknowledgeIEsValueIDNgRannode2CellID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateMobilityChangeAcknowledgeIEsValueIDCriticalityDiagnostics(IDCriticalityDiagnostics *xnapiesv1.CriticalityDiagnostics) (*xnappducontentsv1.MobilityChangeAcknowledgeIEsValue, error) {

	item := &xnappducontentsv1.MobilityChangeAcknowledgeIEsValue{
		MobilityChangeAcknowledgeIes: &xnappducontentsv1.MobilityChangeAcknowledgeIEsValue_IdCriticalityDiagnostics{
			IdCriticalityDiagnostics: IDCriticalityDiagnostics,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateMobilityChangeAcknowledgeIEsValueIDCriticalityDiagnostics() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateMobilityChangeFailureIEsValueIDNgRannode1CellID(IDNgRannode1CellID *xnapiesv1.GlobalNgRAncellID) (*xnappducontentsv1.MobilityChangeFailureIEsValue, error) {

	item := &xnappducontentsv1.MobilityChangeFailureIEsValue{
		MobilityChangeFailureIes: &xnappducontentsv1.MobilityChangeFailureIEsValue_IdNgRannode1CellId{
			IdNgRannode1CellId: IDNgRannode1CellID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateMobilityChangeFailureIEsValueIDNgRannode1CellID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateMobilityChangeFailureIEsValueIDNgRannode2CellID(IDNgRannode2CellID *xnapiesv1.GlobalNgRAncellID) (*xnappducontentsv1.MobilityChangeFailureIEsValue, error) {

	item := &xnappducontentsv1.MobilityChangeFailureIEsValue{
		MobilityChangeFailureIes: &xnappducontentsv1.MobilityChangeFailureIEsValue_IdNgRannode2CellId{
			IdNgRannode2CellId: IDNgRannode2CellID,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateMobilityChangeFailureIEsValueIDNgRannode2CellID() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateMobilityChangeFailureIEsValueIDCause(IDCause *xnapiesv1.Cause) (*xnappducontentsv1.MobilityChangeFailureIEsValue, error) {

	item := &xnappducontentsv1.MobilityChangeFailureIEsValue{
		MobilityChangeFailureIes: &xnappducontentsv1.MobilityChangeFailureIEsValue_IdCause{
			IdCause: IDCause,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateMobilityChangeFailureIEsValueIDCause() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateMobilityChangeFailureIEsValueIDMobilityParametersModificationRange(IDMobilityParametersModificationRange *xnapiesv1.MobilityParametersModificationRange) (*xnappducontentsv1.MobilityChangeFailureIEsValue, error) {

	item := &xnappducontentsv1.MobilityChangeFailureIEsValue{
		MobilityChangeFailureIes: &xnappducontentsv1.MobilityChangeFailureIEsValue_IdMobilityParametersModificationRange{
			IdMobilityParametersModificationRange: IDMobilityParametersModificationRange,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateMobilityChangeFailureIEsValueIDMobilityParametersModificationRange() error validating PDU %s", err.Error())
	//}

	return item, nil
}
func CreateMobilityChangeFailureIEsValueIDCriticalityDiagnostics(IDCriticalityDiagnostics *xnapiesv1.CriticalityDiagnostics) (*xnappducontentsv1.MobilityChangeFailureIEsValue, error) {

	item := &xnappducontentsv1.MobilityChangeFailureIEsValue{
		MobilityChangeFailureIes: &xnappducontentsv1.MobilityChangeFailureIEsValue_IdCriticalityDiagnostics{
			IdCriticalityDiagnostics: IDCriticalityDiagnostics,
		},
	}

	//if err := item.Validate(); err != nil {
	//	return nil, errors.NewInvalid("CreateMobilityChangeFailureIEsValueIDCriticalityDiagnostics() error validating PDU %s", err.Error())
	//}

	return item, nil
}

func CreateAdditionLocationInformationIncludePscell() xnapiesv1.AdditionLocationInformation {
	return xnapiesv1.AdditionLocationInformation_ADDITION_LOCATION_INFORMATION_INCLUDE_PSCELL
}
func CreatePreemptioncapabilityAllocationandRetentionPriorityShallNotTriggerPreemptdatDion() xnapiesv1.PreemptioncapabilityAllocationandRetentionPriority {
	return xnapiesv1.PreemptioncapabilityAllocationandRetentionPriority_PREEMPTIONCAPABILITY_ALLOCATIONAND_RETENTION_PRIORITY_SHALL_NOT_TRIGGER_PREEMPTDAT_DION
}
func CreatePreemptioncapabilityAllocationandRetentionPriorityMayTriggerPreemption() xnapiesv1.PreemptioncapabilityAllocationandRetentionPriority {
	return xnapiesv1.PreemptioncapabilityAllocationandRetentionPriority_PREEMPTIONCAPABILITY_ALLOCATIONAND_RETENTION_PRIORITY_MAY_TRIGGER_PREEMPTION
}
func CreatePreemptionvulnerabilityAllocationandRetentionPriorityNotPreemptable() xnapiesv1.PreemptionvulnerabilityAllocationandRetentionPriority {
	return xnapiesv1.PreemptionvulnerabilityAllocationandRetentionPriority_PREEMPTIONVULNERABILITY_ALLOCATIONAND_RETENTION_PRIORITY_NOT_PREEMPTABLE
}
func CreatePreemptionvulnerabilityAllocationandRetentionPriorityPreemptable() xnapiesv1.PreemptionvulnerabilityAllocationandRetentionPriority {
	return xnapiesv1.PreemptionvulnerabilityAllocationandRetentionPriority_PREEMPTIONVULNERABILITY_ALLOCATIONAND_RETENTION_PRIORITY_PREEMPTABLE
}
func CreateBtrssiBluetoothMeasurementConfigurationTrue() xnapiesv1.BtrssiBluetoothMeasurementConfiguration {
	return xnapiesv1.BtrssiBluetoothMeasurementConfiguration_BTRSSI_BLUETOOTH_MEASUREMENT_CONFIGURATION_TRUE
}
func CreateBluetoothMeasConfigSetup() xnapiesv1.BluetoothMeasConfig {
	return xnapiesv1.BluetoothMeasConfig_BLUETOOTH_MEAS_CONFIG_SETUP
}
func CreateCauseRadioNetworkLayerCellNotAvailable() xnapiesv1.CauseRadioNetworkLayer {
	return xnapiesv1.CauseRadioNetworkLayer_CAUSE_RADIO_NETWORK_LAYER_CELL_NOT_AVAILABLE
}
func CreateCauseRadioNetworkLayerHandoverDesirableForRadioReasons() xnapiesv1.CauseRadioNetworkLayer {
	return xnapiesv1.CauseRadioNetworkLayer_CAUSE_RADIO_NETWORK_LAYER_HANDOVER_DESIRABLE_FOR_RADIO_REASONS
}
func CreateCauseRadioNetworkLayerHandoverTargetNotAllowed() xnapiesv1.CauseRadioNetworkLayer {
	return xnapiesv1.CauseRadioNetworkLayer_CAUSE_RADIO_NETWORK_LAYER_HANDOVER_TARGET_NOT_ALLOWED
}
func CreateCauseRadioNetworkLayerInvalidAmfSetID() xnapiesv1.CauseRadioNetworkLayer {
	return xnapiesv1.CauseRadioNetworkLayer_CAUSE_RADIO_NETWORK_LAYER_INVALID_AMF_SET_ID
}
func CreateCauseRadioNetworkLayerNoRadioResourcesAvailableInTargetCell() xnapiesv1.CauseRadioNetworkLayer {
	return xnapiesv1.CauseRadioNetworkLayer_CAUSE_RADIO_NETWORK_LAYER_NO_RADIO_RESOURCES_AVAILABLE_IN_TARGET_CELL
}
func CreateCauseRadioNetworkLayerPartialHandover() xnapiesv1.CauseRadioNetworkLayer {
	return xnapiesv1.CauseRadioNetworkLayer_CAUSE_RADIO_NETWORK_LAYER_PARTIAL_HANDOVER
}
func CreateCauseRadioNetworkLayerReduceLoadInServingCell() xnapiesv1.CauseRadioNetworkLayer {
	return xnapiesv1.CauseRadioNetworkLayer_CAUSE_RADIO_NETWORK_LAYER_REDUCE_LOAD_IN_SERVING_CELL
}
func CreateCauseRadioNetworkLayerResourceOptimisationHandover() xnapiesv1.CauseRadioNetworkLayer {
	return xnapiesv1.CauseRadioNetworkLayer_CAUSE_RADIO_NETWORK_LAYER_RESOURCE_OPTIMISATION_HANDOVER
}
func CreateCauseRadioNetworkLayerTimeCriticalHandover() xnapiesv1.CauseRadioNetworkLayer {
	return xnapiesv1.CauseRadioNetworkLayer_CAUSE_RADIO_NETWORK_LAYER_TIME_CRITICAL_HANDOVER
}
func CreateCauseRadioNetworkLayerTXnRelocoverallExpiry() xnapiesv1.CauseRadioNetworkLayer {
	return xnapiesv1.CauseRadioNetworkLayer_CAUSE_RADIO_NETWORK_LAYER_T_XN_RELOCOVERALL_EXPIRY
}
func CreateCauseRadioNetworkLayerTXnRelocprepExpiry() xnapiesv1.CauseRadioNetworkLayer {
	return xnapiesv1.CauseRadioNetworkLayer_CAUSE_RADIO_NETWORK_LAYER_T_XN_RELOCPREP_EXPIRY
}
func CreateCauseRadioNetworkLayerUnknownGuamiID() xnapiesv1.CauseRadioNetworkLayer {
	return xnapiesv1.CauseRadioNetworkLayer_CAUSE_RADIO_NETWORK_LAYER_UNKNOWN_GUAMI_ID
}
func CreateCauseRadioNetworkLayerUnknownLocalNgRanNodeUeXnApID() xnapiesv1.CauseRadioNetworkLayer {
	return xnapiesv1.CauseRadioNetworkLayer_CAUSE_RADIO_NETWORK_LAYER_UNKNOWN_LOCAL_NG_RAN_NODE_UE_XN_AP_ID
}
func CreateCauseRadioNetworkLayerInconsistentRemoteNgRanNodeUeXnApID() xnapiesv1.CauseRadioNetworkLayer {
	return xnapiesv1.CauseRadioNetworkLayer_CAUSE_RADIO_NETWORK_LAYER_INCONSISTENT_REMOTE_NG_RAN_NODE_UE_XN_AP_ID
}
func CreateCauseRadioNetworkLayerEncryptionAndOrIntegrityProtectionAlgorithmsNotSupported() xnapiesv1.CauseRadioNetworkLayer {
	return xnapiesv1.CauseRadioNetworkLayer_CAUSE_RADIO_NETWORK_LAYER_ENCRYPTION_AND_OR_INTEGRITY_PROTECTION_ALGORITHMS_NOT_SUPPORTED
}
func CreateCauseRadioNetworkLayerProtectionAlgorithmsNotSupported() xnapiesv1.CauseRadioNetworkLayer {
	return xnapiesv1.CauseRadioNetworkLayer_CAUSE_RADIO_NETWORK_LAYER_PROTECTION_ALGORITHMS_NOT_SUPPORTED
}
func CreateCauseRadioNetworkLayerMultiplePduSessionIDInstances() xnapiesv1.CauseRadioNetworkLayer {
	return xnapiesv1.CauseRadioNetworkLayer_CAUSE_RADIO_NETWORK_LAYER_MULTIPLE_PDU_SESSION_ID_INSTANCES
}
func CreateCauseRadioNetworkLayerUnknownPduSessionID() xnapiesv1.CauseRadioNetworkLayer {
	return xnapiesv1.CauseRadioNetworkLayer_CAUSE_RADIO_NETWORK_LAYER_UNKNOWN_PDU_SESSION_ID
}
func CreateCauseRadioNetworkLayerUnknownQoSFlowID() xnapiesv1.CauseRadioNetworkLayer {
	return xnapiesv1.CauseRadioNetworkLayer_CAUSE_RADIO_NETWORK_LAYER_UNKNOWN_QO_S_FLOW_ID
}
func CreateCauseRadioNetworkLayerMultipleQoSFlowIDInstances() xnapiesv1.CauseRadioNetworkLayer {
	return xnapiesv1.CauseRadioNetworkLayer_CAUSE_RADIO_NETWORK_LAYER_MULTIPLE_QO_S_FLOW_ID_INSTANCES
}
func CreateCauseRadioNetworkLayerSwitchOffOngoing() xnapiesv1.CauseRadioNetworkLayer {
	return xnapiesv1.CauseRadioNetworkLayer_CAUSE_RADIO_NETWORK_LAYER_SWITCH_OFF_ONGOING
}
func CreateCauseRadioNetworkLayerNotSupported5QiValue() xnapiesv1.CauseRadioNetworkLayer {
	return xnapiesv1.CauseRadioNetworkLayer_CAUSE_RADIO_NETWORK_LAYER_NOT_SUPPORTED_5_QI_VALUE
}
func CreateCauseRadioNetworkLayerTXnDcoverallExpiry() xnapiesv1.CauseRadioNetworkLayer {
	return xnapiesv1.CauseRadioNetworkLayer_CAUSE_RADIO_NETWORK_LAYER_T_XN_DCOVERALL_EXPIRY
}
func CreateCauseRadioNetworkLayerTXnDcprepExpiry() xnapiesv1.CauseRadioNetworkLayer {
	return xnapiesv1.CauseRadioNetworkLayer_CAUSE_RADIO_NETWORK_LAYER_T_XN_DCPREP_EXPIRY
}
func CreateCauseRadioNetworkLayerActionDesirableForRadioReasons() xnapiesv1.CauseRadioNetworkLayer {
	return xnapiesv1.CauseRadioNetworkLayer_CAUSE_RADIO_NETWORK_LAYER_ACTION_DESIRABLE_FOR_RADIO_REASONS
}
func CreateCauseRadioNetworkLayerReduceLoad() xnapiesv1.CauseRadioNetworkLayer {
	return xnapiesv1.CauseRadioNetworkLayer_CAUSE_RADIO_NETWORK_LAYER_REDUCE_LOAD
}
func CreateCauseRadioNetworkLayerResourceOptimisation() xnapiesv1.CauseRadioNetworkLayer {
	return xnapiesv1.CauseRadioNetworkLayer_CAUSE_RADIO_NETWORK_LAYER_RESOURCE_OPTIMISATION
}
func CreateCauseRadioNetworkLayerTimeCriticalAction() xnapiesv1.CauseRadioNetworkLayer {
	return xnapiesv1.CauseRadioNetworkLayer_CAUSE_RADIO_NETWORK_LAYER_TIME_CRITICAL_ACTION
}
func CreateCauseRadioNetworkLayerTargetNotAllowed() xnapiesv1.CauseRadioNetworkLayer {
	return xnapiesv1.CauseRadioNetworkLayer_CAUSE_RADIO_NETWORK_LAYER_TARGET_NOT_ALLOWED
}
func CreateCauseRadioNetworkLayerNoRadioResourcesAvailable() xnapiesv1.CauseRadioNetworkLayer {
	return xnapiesv1.CauseRadioNetworkLayer_CAUSE_RADIO_NETWORK_LAYER_NO_RADIO_RESOURCES_AVAILABLE
}
func CreateCauseRadioNetworkLayerInvalidQoSCombination() xnapiesv1.CauseRadioNetworkLayer {
	return xnapiesv1.CauseRadioNetworkLayer_CAUSE_RADIO_NETWORK_LAYER_INVALID_QO_S_COMBINATION
}
func CreateCauseRadioNetworkLayerEncryptionAlgorithmsNotSupported() xnapiesv1.CauseRadioNetworkLayer {
	return xnapiesv1.CauseRadioNetworkLayer_CAUSE_RADIO_NETWORK_LAYER_ENCRYPTION_ALGORITHMS_NOT_SUPPORTED
}
func CreateCauseRadioNetworkLayerProcedureCancelled() xnapiesv1.CauseRadioNetworkLayer {
	return xnapiesv1.CauseRadioNetworkLayer_CAUSE_RADIO_NETWORK_LAYER_PROCEDURE_CANCELLED
}
func CreateCauseRadioNetworkLayerRRmPurpose() xnapiesv1.CauseRadioNetworkLayer {
	return xnapiesv1.CauseRadioNetworkLayer_CAUSE_RADIO_NETWORK_LAYER_R_RM_PURPOSE
}
func CreateCauseRadioNetworkLayerImproveUserBitRate() xnapiesv1.CauseRadioNetworkLayer {
	return xnapiesv1.CauseRadioNetworkLayer_CAUSE_RADIO_NETWORK_LAYER_IMPROVE_USER_BIT_RATE
}
func CreateCauseRadioNetworkLayerUserInactivity() xnapiesv1.CauseRadioNetworkLayer {
	return xnapiesv1.CauseRadioNetworkLayer_CAUSE_RADIO_NETWORK_LAYER_USER_INACTIVITY
}
func CreateCauseRadioNetworkLayerRadioConnectionWithUeLost() xnapiesv1.CauseRadioNetworkLayer {
	return xnapiesv1.CauseRadioNetworkLayer_CAUSE_RADIO_NETWORK_LAYER_RADIO_CONNECTION_WITH_UE_LOST
}
func CreateCauseRadioNetworkLayerFailureInTheRadioInterfaceProcedure() xnapiesv1.CauseRadioNetworkLayer {
	return xnapiesv1.CauseRadioNetworkLayer_CAUSE_RADIO_NETWORK_LAYER_FAILURE_IN_THE_RADIO_INTERFACE_PROCEDURE
}
func CreateCauseRadioNetworkLayerBearerOptionNotSupported() xnapiesv1.CauseRadioNetworkLayer {
	return xnapiesv1.CauseRadioNetworkLayer_CAUSE_RADIO_NETWORK_LAYER_BEARER_OPTION_NOT_SUPPORTED
}
func CreateCauseRadioNetworkLayerUpIntegrityProtectionNotPossible() xnapiesv1.CauseRadioNetworkLayer {
	return xnapiesv1.CauseRadioNetworkLayer_CAUSE_RADIO_NETWORK_LAYER_UP_INTEGRITY_PROTECTION_NOT_POSSIBLE
}
func CreateCauseRadioNetworkLayerUpConfidentialityProtectionNotPossible() xnapiesv1.CauseRadioNetworkLayer {
	return xnapiesv1.CauseRadioNetworkLayer_CAUSE_RADIO_NETWORK_LAYER_UP_CONFIDENTIALITY_PROTECTION_NOT_POSSIBLE
}
func CreateCauseRadioNetworkLayerResourcesNotAvailableForTheSliceS() xnapiesv1.CauseRadioNetworkLayer {
	return xnapiesv1.CauseRadioNetworkLayer_CAUSE_RADIO_NETWORK_LAYER_RESOURCES_NOT_AVAILABLE_FOR_THE_SLICE_S
}
func CreateCauseRadioNetworkLayerUeMaxIPDataRateReason() xnapiesv1.CauseRadioNetworkLayer {
	return xnapiesv1.CauseRadioNetworkLayer_CAUSE_RADIO_NETWORK_LAYER_UE_MAX_IP_DATA_RATE_REASON
}
func CreateCauseRadioNetworkLayerCPIntegrityProtectionFailure() xnapiesv1.CauseRadioNetworkLayer {
	return xnapiesv1.CauseRadioNetworkLayer_CAUSE_RADIO_NETWORK_LAYER_C_P_INTEGRITY_PROTECTION_FAILURE
}
func CreateCauseRadioNetworkLayerUPIntegrityProtectionFailure() xnapiesv1.CauseRadioNetworkLayer {
	return xnapiesv1.CauseRadioNetworkLayer_CAUSE_RADIO_NETWORK_LAYER_U_P_INTEGRITY_PROTECTION_FAILURE
}
func CreateCauseRadioNetworkLayerSliceNotSupportedByNgRan() xnapiesv1.CauseRadioNetworkLayer {
	return xnapiesv1.CauseRadioNetworkLayer_CAUSE_RADIO_NETWORK_LAYER_SLICE_NOT_SUPPORTED_BY_NG_RAN
}
func CreateCauseRadioNetworkLayerMNMobility() xnapiesv1.CauseRadioNetworkLayer {
	return xnapiesv1.CauseRadioNetworkLayer_CAUSE_RADIO_NETWORK_LAYER_M_N_MOBILITY
}
func CreateCauseRadioNetworkLayerSNMobility() xnapiesv1.CauseRadioNetworkLayer {
	return xnapiesv1.CauseRadioNetworkLayer_CAUSE_RADIO_NETWORK_LAYER_S_N_MOBILITY
}
func CreateCauseRadioNetworkLayerCountReachesMaxValue() xnapiesv1.CauseRadioNetworkLayer {
	return xnapiesv1.CauseRadioNetworkLayer_CAUSE_RADIO_NETWORK_LAYER_COUNT_REACHES_MAX_VALUE
}
func CreateCauseRadioNetworkLayerUnknownOldNgRanNodeUeXnApID() xnapiesv1.CauseRadioNetworkLayer {
	return xnapiesv1.CauseRadioNetworkLayer_CAUSE_RADIO_NETWORK_LAYER_UNKNOWN_OLD_NG_RAN_NODE_UE_XN_AP_ID
}
func CreateCauseRadioNetworkLayerPDcpOverload() xnapiesv1.CauseRadioNetworkLayer {
	return xnapiesv1.CauseRadioNetworkLayer_CAUSE_RADIO_NETWORK_LAYER_P_DCP_OVERLOAD
}
func CreateCauseRadioNetworkLayerDrbIDNotAvailable() xnapiesv1.CauseRadioNetworkLayer {
	return xnapiesv1.CauseRadioNetworkLayer_CAUSE_RADIO_NETWORK_LAYER_DRB_ID_NOT_AVAILABLE
}
func CreateCauseRadioNetworkLayerUnspecified() xnapiesv1.CauseRadioNetworkLayer {
	return xnapiesv1.CauseRadioNetworkLayer_CAUSE_RADIO_NETWORK_LAYER_UNSPECIFIED
}
func CreateCauseRadioNetworkLayerUeContextIDNotKnown() xnapiesv1.CauseRadioNetworkLayer {
	return xnapiesv1.CauseRadioNetworkLayer_CAUSE_RADIO_NETWORK_LAYER_UE_CONTEXT_ID_NOT_KNOWN
}
func CreateCauseRadioNetworkLayerNonRelocationOfContext() xnapiesv1.CauseRadioNetworkLayer {
	return xnapiesv1.CauseRadioNetworkLayer_CAUSE_RADIO_NETWORK_LAYER_NON_RELOCATION_OF_CONTEXT
}
func CreateCauseRadioNetworkLayerChoCpcResourcesTobechanged() xnapiesv1.CauseRadioNetworkLayer {
	return xnapiesv1.CauseRadioNetworkLayer_CAUSE_RADIO_NETWORK_LAYER_CHO_CPC_RESOURCES_TOBECHANGED
}
func CreateCauseRadioNetworkLayerRSnNotAvailableForTheUp() xnapiesv1.CauseRadioNetworkLayer {
	return xnapiesv1.CauseRadioNetworkLayer_CAUSE_RADIO_NETWORK_LAYER_R_SN_NOT_AVAILABLE_FOR_THE_UP
}
func CreateCauseRadioNetworkLayerNpnAccessDenied() xnapiesv1.CauseRadioNetworkLayer {
	return xnapiesv1.CauseRadioNetworkLayer_CAUSE_RADIO_NETWORK_LAYER_NPN_ACCESS_DENIED
}
func CreateCauseRadioNetworkLayerReportCharacteristicsEmpty() xnapiesv1.CauseRadioNetworkLayer {
	return xnapiesv1.CauseRadioNetworkLayer_CAUSE_RADIO_NETWORK_LAYER_REPORT_CHARACTERISTICS_EMPTY
}
func CreateCauseRadioNetworkLayerExistingMeasurementID() xnapiesv1.CauseRadioNetworkLayer {
	return xnapiesv1.CauseRadioNetworkLayer_CAUSE_RADIO_NETWORK_LAYER_EXISTING_MEASUREMENT_ID
}
func CreateCauseRadioNetworkLayerMeasurementTemporarilyNotAvailable() xnapiesv1.CauseRadioNetworkLayer {
	return xnapiesv1.CauseRadioNetworkLayer_CAUSE_RADIO_NETWORK_LAYER_MEASUREMENT_TEMPORARILY_NOT_AVAILABLE
}
func CreateCauseRadioNetworkLayerMeasurementNotSupportedForTheObject() xnapiesv1.CauseRadioNetworkLayer {
	return xnapiesv1.CauseRadioNetworkLayer_CAUSE_RADIO_NETWORK_LAYER_MEASUREMENT_NOT_SUPPORTED_FOR_THE_OBJECT
}
func CreateCauseRadioNetworkLayerUePowerSaving() xnapiesv1.CauseRadioNetworkLayer {
	return xnapiesv1.CauseRadioNetworkLayer_CAUSE_RADIO_NETWORK_LAYER_UE_POWER_SAVING
}
func CreateCauseRadioNetworkLayerUnknownNgRanNode2MeasurementID() xnapiesv1.CauseRadioNetworkLayer {
	return xnapiesv1.CauseRadioNetworkLayer_CAUSE_RADIO_NETWORK_LAYER_UNKNOWN_NG_RAN_NODE2_MEASUREMENT_ID
}
func CreateCauseRadioNetworkLayerInsufficientUeCapabilities() xnapiesv1.CauseRadioNetworkLayer {
	return xnapiesv1.CauseRadioNetworkLayer_CAUSE_RADIO_NETWORK_LAYER_INSUFFICIENT_UE_CAPABILITIES
}
func CreateCauseRadioNetworkLayerNormalRelease() xnapiesv1.CauseRadioNetworkLayer {
	return xnapiesv1.CauseRadioNetworkLayer_CAUSE_RADIO_NETWORK_LAYER_NORMAL_RELEASE
}
func CreateCauseRadioNetworkLayerValueOutOfAllowedRange() xnapiesv1.CauseRadioNetworkLayer {
	return xnapiesv1.CauseRadioNetworkLayer_CAUSE_RADIO_NETWORK_LAYER_VALUE_OUT_OF_ALLOWED_RANGE
}
func CreateCauseTransportLayerTransportResourceUnavailable() xnapiesv1.CauseTransportLayer {
	return xnapiesv1.CauseTransportLayer_CAUSE_TRANSPORT_LAYER_TRANSPORT_RESOURCE_UNAVAILABLE
}
func CreateCauseTransportLayerUnspecified() xnapiesv1.CauseTransportLayer {
	return xnapiesv1.CauseTransportLayer_CAUSE_TRANSPORT_LAYER_UNSPECIFIED
}
func CreateCauseProtocolTransferSyntaxError() xnapiesv1.CauseProtocol {
	return xnapiesv1.CauseProtocol_CAUSE_PROTOCOL_TRANSFER_SYNTAX_ERROR
}
func CreateCauseProtocolAbstractSyntaxErrorReject() xnapiesv1.CauseProtocol {
	return xnapiesv1.CauseProtocol_CAUSE_PROTOCOL_ABSTRACT_SYNTAX_ERROR_REJECT
}
func CreateCauseProtocolAbstractSyntaxErrorIgnoreAndNotify() xnapiesv1.CauseProtocol {
	return xnapiesv1.CauseProtocol_CAUSE_PROTOCOL_ABSTRACT_SYNTAX_ERROR_IGNORE_AND_NOTIFY
}
func CreateCauseProtocolMessageNotCompatibleWithReceiverState() xnapiesv1.CauseProtocol {
	return xnapiesv1.CauseProtocol_CAUSE_PROTOCOL_MESSAGE_NOT_COMPATIBLE_WITH_RECEIVER_STATE
}
func CreateCauseProtocolSemanticError() xnapiesv1.CauseProtocol {
	return xnapiesv1.CauseProtocol_CAUSE_PROTOCOL_SEMANTIC_ERROR
}
func CreateCauseProtocolAbstractSyntaxErrorFalselyConstructedMessage() xnapiesv1.CauseProtocol {
	return xnapiesv1.CauseProtocol_CAUSE_PROTOCOL_ABSTRACT_SYNTAX_ERROR_FALSELY_CONSTRUCTED_MESSAGE
}
func CreateCauseProtocolUnspecified() xnapiesv1.CauseProtocol {
	return xnapiesv1.CauseProtocol_CAUSE_PROTOCOL_UNSPECIFIED
}
func CreateCauseMiscControlProcessingOverload() xnapiesv1.CauseMisc {
	return xnapiesv1.CauseMisc_CAUSE_MISC_CONTROL_PROCESSING_OVERLOAD
}
func CreateCauseMiscHardwareFailure() xnapiesv1.CauseMisc {
	return xnapiesv1.CauseMisc_CAUSE_MISC_HARDWARE_FAILURE
}
func CreateCauseMiscOAndMIntervention() xnapiesv1.CauseMisc {
	return xnapiesv1.CauseMisc_CAUSE_MISC_O_AND_M_INTERVENTION
}
func CreateCauseMiscNotEnoughUserPlaneProcessingResources() xnapiesv1.CauseMisc {
	return xnapiesv1.CauseMisc_CAUSE_MISC_NOT_ENOUGH_USER_PLANE_PROCESSING_RESOURCES
}
func CreateCauseMiscUnspecified() xnapiesv1.CauseMisc {
	return xnapiesv1.CauseMisc_CAUSE_MISC_UNSPECIFIED
}
func CreateFullListCellAssistanceInfoNrAllServedCellsNr() xnapiesv1.FullListCellAssistanceInfoNr {
	return xnapiesv1.FullListCellAssistanceInfoNr_FULL_LIST_CELL_ASSISTANCE_INFO_NR_ALL_SERVED_CELLS_NR
}
func CreateFullListCellAssistanceInfoEutraAllServedCellsEUtra() xnapiesv1.FullListCellAssistanceInfoEutra {
	return xnapiesv1.FullListCellAssistanceInfoEutra_FULL_LIST_CELL_ASSISTANCE_INFO_EUTRA_ALL_SERVED_CELLS_E_UTRA
}
func CreateChoMrdcEarlyDataForwardingStop() xnapiesv1.ChoMRdcEarlyDataForwarding {
	return xnapiesv1.ChoMRdcEarlyDataForwarding_CHO_MRDC_EARLY_DATA_FORWARDING_STOP
}
func CreateChoMrdcIndicatorTrue() xnapiesv1.ChoMRdcIndicator {
	return xnapiesv1.ChoMRdcIndicator_CHO_MRDC_INDICATOR_TRUE
}
func CreateChotriggerChoInitiation() xnapiesv1.Chotrigger {
	return xnapiesv1.Chotrigger_CHOTRIGGER_CHO_INITIATION
}
func CreateChotriggerChoReplace() xnapiesv1.Chotrigger {
	return xnapiesv1.Chotrigger_CHOTRIGGER_CHO_REPLACE
}
func CreateConfiguredTacindicationTrue() xnapiesv1.ConfiguredTacindication {
	return xnapiesv1.ConfiguredTacindication_CONFIGURED_TACINDICATION_TRUE
}
func CreateENdcsupportConnectivitySupportSupported() xnapiesv1.EndcsupportConnectivitySupport {
	return xnapiesv1.EndcsupportConnectivitySupport_E_NDCSUPPORT_CONNECTIVITY_SUPPORT_SUPPORTED
}
func CreateENdcsupportConnectivitySupportNotSupported() xnapiesv1.EndcsupportConnectivitySupport {
	return xnapiesv1.EndcsupportConnectivitySupport_E_NDCSUPPORT_CONNECTIVITY_SUPPORT_NOT_SUPPORTED
}
func CreateCyclicPrefixEUtraDlNormal() xnapiesv1.CyclicPrefixEUTraDL {
	return xnapiesv1.CyclicPrefixEUTraDL_CYCLIC_PREFIX_E_UTRA_DL_NORMAL
}
func CreateCyclicPrefixEUtraDlExtended() xnapiesv1.CyclicPrefixEUTraDL {
	return xnapiesv1.CyclicPrefixEUTraDL_CYCLIC_PREFIX_E_UTRA_DL_EXTENDED
}
func CreateCyclicPrefixEUtraUlNormal() xnapiesv1.CyclicPrefixEUTraUL {
	return xnapiesv1.CyclicPrefixEUTraUL_CYCLIC_PREFIX_E_UTRA_UL_NORMAL
}
func CreateCyclicPrefixEUtraUlExtended() xnapiesv1.CyclicPrefixEUTraUL {
	return xnapiesv1.CyclicPrefixEUTraUL_CYCLIC_PREFIX_E_UTRA_UL_EXTENDED
}
func CreateCsiRstransmissionIndicationActivated() xnapiesv1.CsiRStransmissionIndication {
	return xnapiesv1.CsiRStransmissionIndication_CSI_RSTRANSMISSION_INDICATION_ACTIVATED
}
func CreateCsiRstransmissionIndicationDeactivated() xnapiesv1.CsiRStransmissionIndication {
	return xnapiesv1.CsiRStransmissionIndication_CSI_RSTRANSMISSION_INDICATION_DEACTIVATED
}
func CreateDataForwardingAcceptedDataForwardingAccepted() xnapiesv1.DataForwardingAccepted {
	return xnapiesv1.DataForwardingAccepted_DATA_FORWARDING_ACCEPTED_DATA_FORWARDING_ACCEPTED
}
func CreateDapsIndicatorDapsrequestInfoDapsHoRequired() xnapiesv1.DapsIndicatorDapsrequestInfo {
	return xnapiesv1.DapsIndicatorDapsrequestInfo_DAPS_INDICATOR_DAPSREQUEST_INFO_DAPS_HO_REQUIRED
}
func CreateDapsResponseIndicatorDapsresponseInfoItemDapsHoAccepted() xnapiesv1.DapsResponseIndicatorDapsresponseInfoItem {
	return xnapiesv1.DapsResponseIndicatorDapsresponseInfoItem_DAPS_RESPONSE_INDICATOR_DAPSRESPONSE_INFO_ITEM_DAPS_HO_ACCEPTED
}
func CreateDapsResponseIndicatorDapsresponseInfoItemDapsHoNotAccepted() xnapiesv1.DapsResponseIndicatorDapsresponseInfoItem {
	return xnapiesv1.DapsResponseIndicatorDapsresponseInfoItem_DAPS_RESPONSE_INDICATOR_DAPSRESPONSE_INFO_ITEM_DAPS_HO_NOT_ACCEPTED
}
func CreateDesiredActNotificationLevelNone() xnapiesv1.DesiredActNotificationLevel {
	return xnapiesv1.DesiredActNotificationLevel_DESIRED_ACT_NOTIFICATION_LEVEL_NONE
}
func CreateDesiredActNotificationLevelQosFlow() xnapiesv1.DesiredActNotificationLevel {
	return xnapiesv1.DesiredActNotificationLevel_DESIRED_ACT_NOTIFICATION_LEVEL_QOS_FLOW
}
func CreateDesiredActNotificationLevelPduSession() xnapiesv1.DesiredActNotificationLevel {
	return xnapiesv1.DesiredActNotificationLevel_DESIRED_ACT_NOTIFICATION_LEVEL_PDU_SESSION
}
func CreateDesiredActNotificationLevelUeLevel() xnapiesv1.DesiredActNotificationLevel {
	return xnapiesv1.DesiredActNotificationLevel_DESIRED_ACT_NOTIFICATION_LEVEL_UE_LEVEL
}
func CreateDefaultDrbAllowedTrue() xnapiesv1.DefaultDrbAllowed {
	return xnapiesv1.DefaultDrbAllowed_DEFAULT_DRB_ALLOWED_TRUE
}
func CreateDefaultDrbAllowedFalse() xnapiesv1.DefaultDrbAllowed {
	return xnapiesv1.DefaultDrbAllowed_DEFAULT_DRB_ALLOWED_FALSE
}
func CreateDirectForwardingPathAvailabilityDirectPathAvailable() xnapiesv1.DirectForwardingPathAvailability {
	return xnapiesv1.DirectForwardingPathAvailability_DIRECT_FORWARDING_PATH_AVAILABILITY_DIRECT_PATH_AVAILABLE
}
func CreateDlforwardingDlForwardingProposed() xnapiesv1.Dlforwarding {
	return xnapiesv1.Dlforwarding_DLFORWARDING_DL_FORWARDING_PROPOSED
}
func CreateDuplicationActivationActive() xnapiesv1.DuplicationActivation {
	return xnapiesv1.DuplicationActivation_DUPLICATION_ACTIVATION_ACTIVE
}
func CreateDuplicationActivationInactive() xnapiesv1.DuplicationActivation {
	return xnapiesv1.DuplicationActivation_DUPLICATION_ACTIVATION_INACTIVE
}
func CreateDelayCriticalDynamic5QidescriptorDelayCritical() xnapiesv1.DelayCriticalDynamic5Qidescriptor {
	return xnapiesv1.DelayCriticalDynamic5Qidescriptor_DELAY_CRITICAL_DYNAMIC5_QIDESCRIPTOR_DELAY_CRITICAL
}
func CreateDelayCriticalDynamic5QidescriptorNonDelayCritical() xnapiesv1.DelayCriticalDynamic5Qidescriptor {
	return xnapiesv1.DelayCriticalDynamic5Qidescriptor_DELAY_CRITICAL_DYNAMIC5_QIDESCRIPTOR_NON_DELAY_CRITICAL
}
func CreateHighSpeedFlagEutraprachconfigurationTrue() xnapiesv1.HighSpeedFlagEutraprachconfiguration {
	return xnapiesv1.HighSpeedFlagEutraprachconfiguration_HIGH_SPEED_FLAG_EUTRAPRACHCONFIGURATION_TRUE
}
func CreateHighSpeedFlagEutraprachconfigurationFalse() xnapiesv1.HighSpeedFlagEutraprachconfiguration {
	return xnapiesv1.HighSpeedFlagEutraprachconfiguration_HIGH_SPEED_FLAG_EUTRAPRACHCONFIGURATION_FALSE
}
func CreateEUtratransmissionBandwidthBw6() xnapiesv1.EUTratransmissionBandwidth {
	return xnapiesv1.EUTratransmissionBandwidth_E_UTRATRANSMISSION_BANDWIDTH_BW6
}
func CreateEUtratransmissionBandwidthBw15() xnapiesv1.EUTratransmissionBandwidth {
	return xnapiesv1.EUTratransmissionBandwidth_E_UTRATRANSMISSION_BANDWIDTH_BW15
}
func CreateEUtratransmissionBandwidthBw25() xnapiesv1.EUTratransmissionBandwidth {
	return xnapiesv1.EUTratransmissionBandwidth_E_UTRATRANSMISSION_BANDWIDTH_BW25
}
func CreateEUtratransmissionBandwidthBw50() xnapiesv1.EUTratransmissionBandwidth {
	return xnapiesv1.EUTratransmissionBandwidth_E_UTRATRANSMISSION_BANDWIDTH_BW50
}
func CreateEUtratransmissionBandwidthBw75() xnapiesv1.EUTratransmissionBandwidth {
	return xnapiesv1.EUTratransmissionBandwidth_E_UTRATRANSMISSION_BANDWIDTH_BW75
}
func CreateEUtratransmissionBandwidthBw100() xnapiesv1.EUTratransmissionBandwidth {
	return xnapiesv1.EUTratransmissionBandwidth_E_UTRATRANSMISSION_BANDWIDTH_BW100
}
func CreateEUtratransmissionBandwidthBw1() xnapiesv1.EUTratransmissionBandwidth {
	return xnapiesv1.EUTratransmissionBandwidth_E_UTRATRANSMISSION_BANDWIDTH_BW1
}
func CreateEventTypeReportUponChangeOfServingCell() xnapiesv1.EventType {
	return xnapiesv1.EventType_EVENT_TYPE_REPORT_UPON_CHANGE_OF_SERVING_CELL
}
func CreateEventTypeReportUeMovingPresenceIntoOrOutOfTheAreaOfInterest() xnapiesv1.EventType {
	return xnapiesv1.EventType_EVENT_TYPE_REPORT_UE_MOVING_PRESENCE_INTO_OR_OUT_OF_THE_AREA_OF_INTEREST
}
func CreateEventTypeReportUponChangeOfServingCellAndAreaOfInterest() xnapiesv1.EventType {
	return xnapiesv1.EventType_EVENT_TYPE_REPORT_UPON_CHANGE_OF_SERVING_CELL_AND_AREA_OF_INTEREST
}
func CreateOutOfCoverageEventTypeTriggerTrue() xnapiesv1.OutOfCoverageEventTypeTrigger {
	return xnapiesv1.OutOfCoverageEventTypeTrigger_OUT_OF_COVERAGE_EVENT_TYPE_TRIGGER_TRUE
}
func CreateExpectedHointervalSec15() xnapiesv1.ExpectedHointerval {
	return xnapiesv1.ExpectedHointerval_EXPECTED_HOINTERVAL_SEC15
}
func CreateExpectedHointervalSec30() xnapiesv1.ExpectedHointerval {
	return xnapiesv1.ExpectedHointerval_EXPECTED_HOINTERVAL_SEC30
}
func CreateExpectedHointervalSec60() xnapiesv1.ExpectedHointerval {
	return xnapiesv1.ExpectedHointerval_EXPECTED_HOINTERVAL_SEC60
}
func CreateExpectedHointervalSec90() xnapiesv1.ExpectedHointerval {
	return xnapiesv1.ExpectedHointerval_EXPECTED_HOINTERVAL_SEC90
}
func CreateExpectedHointervalSec120() xnapiesv1.ExpectedHointerval {
	return xnapiesv1.ExpectedHointerval_EXPECTED_HOINTERVAL_SEC120
}
func CreateExpectedHointervalSec180() xnapiesv1.ExpectedHointerval {
	return xnapiesv1.ExpectedHointerval_EXPECTED_HOINTERVAL_SEC180
}
func CreateExpectedHointervalLongTime() xnapiesv1.ExpectedHointerval {
	return xnapiesv1.ExpectedHointerval_EXPECTED_HOINTERVAL_LONG_TIME
}
func CreateExpectedUemobilityStationary() xnapiesv1.ExpectedUemobility {
	return xnapiesv1.ExpectedUemobility_EXPECTED_UEMOBILITY_STATIONARY
}
func CreateExpectedUemobilityMobile() xnapiesv1.ExpectedUemobility {
	return xnapiesv1.ExpectedUemobility_EXPECTED_UEMOBILITY_MOBILE
}
func CreateSourceOfUeactivityBehaviourInformationSubscriptionInformation() xnapiesv1.SourceOfUeactivityBehaviourInformation {
	return xnapiesv1.SourceOfUeactivityBehaviourInformation_SOURCE_OF_UEACTIVITY_BEHAVIOUR_INFORMATION_SUBSCRIPTION_INFORMATION
}
func CreateSourceOfUeactivityBehaviourInformationStatistics() xnapiesv1.SourceOfUeactivityBehaviourInformation {
	return xnapiesv1.SourceOfUeactivityBehaviourInformation_SOURCE_OF_UEACTIVITY_BEHAVIOUR_INFORMATION_STATISTICS
}
func CreateFrequencyShift7p5khzFalse() xnapiesv1.FrequencyShift7P5Khz {
	return xnapiesv1.FrequencyShift7P5Khz_FREQUENCY_SHIFT7P5KHZ_FALSE
}
func CreateFrequencyShift7p5khzTrue() xnapiesv1.FrequencyShift7P5Khz {
	return xnapiesv1.FrequencyShift7P5Khz_FREQUENCY_SHIFT7P5KHZ_TRUE
}
func CreateNotificationControlGbrqoSflowInfoNotificationRequested() xnapiesv1.NotificationControlGbrqoSflowInfo {
	return xnapiesv1.NotificationControlGbrqoSflowInfo_NOTIFICATION_CONTROL_GBRQO_SFLOW_INFO_NOTIFICATION_REQUESTED
}
func CreateHandoverReportTypeHoTooEarly() xnapiesv1.HandoverReportType {
	return xnapiesv1.HandoverReportType_HANDOVER_REPORT_TYPE_HO_TOO_EARLY
}
func CreateHandoverReportTypeHoToWrongCell() xnapiesv1.HandoverReportType {
	return xnapiesv1.HandoverReportType_HANDOVER_REPORT_TYPE_HO_TO_WRONG_CELL
}
func CreateHandoverReportTypeIntersystempingpong() xnapiesv1.HandoverReportType {
	return xnapiesv1.HandoverReportType_HANDOVER_REPORT_TYPE_INTERSYSTEMPINGPONG
}
func CreateIabnodeIndicationTrue() xnapiesv1.IabnodeIndication {
	return xnapiesv1.IabnodeIndication_IABNODE_INDICATION_TRUE
}
func CreateLinksToLogUplink() xnapiesv1.Linkstolog {
	return xnapiesv1.Linkstolog_LINKS_TO_LOG_UPLINK
}
func CreateLinksToLogDownlink() xnapiesv1.Linkstolog {
	return xnapiesv1.Linkstolog_LINKS_TO_LOG_DOWNLINK
}
func CreateLinksToLogBothUplinkAndDownlink() xnapiesv1.Linkstolog {
	return xnapiesv1.Linkstolog_LINKS_TO_LOG_BOTH_UPLINK_AND_DOWNLINK
}
func CreateLocationInformationSnreportingPScell() xnapiesv1.LocationInformationSnreporting {
	return xnapiesv1.LocationInformationSnreporting_LOCATION_INFORMATION_SNREPORTING_P_SCELL
}
func CreateLoggingIntervalMs320() xnapiesv1.LoggingInterval {
	return xnapiesv1.LoggingInterval_LOGGING_INTERVAL_MS320
}
func CreateLoggingIntervalMs640() xnapiesv1.LoggingInterval {
	return xnapiesv1.LoggingInterval_LOGGING_INTERVAL_MS640
}
func CreateLoggingIntervalMs1280() xnapiesv1.LoggingInterval {
	return xnapiesv1.LoggingInterval_LOGGING_INTERVAL_MS1280
}
func CreateLoggingIntervalMs2560() xnapiesv1.LoggingInterval {
	return xnapiesv1.LoggingInterval_LOGGING_INTERVAL_MS2560
}
func CreateLoggingIntervalMs5120() xnapiesv1.LoggingInterval {
	return xnapiesv1.LoggingInterval_LOGGING_INTERVAL_MS5120
}
func CreateLoggingIntervalMs10240() xnapiesv1.LoggingInterval {
	return xnapiesv1.LoggingInterval_LOGGING_INTERVAL_MS10240
}
func CreateLoggingIntervalMs20480() xnapiesv1.LoggingInterval {
	return xnapiesv1.LoggingInterval_LOGGING_INTERVAL_MS20480
}
func CreateLoggingIntervalMs30720() xnapiesv1.LoggingInterval {
	return xnapiesv1.LoggingInterval_LOGGING_INTERVAL_MS30720
}
func CreateLoggingIntervalMs40960() xnapiesv1.LoggingInterval {
	return xnapiesv1.LoggingInterval_LOGGING_INTERVAL_MS40960
}
func CreateLoggingIntervalMs61440() xnapiesv1.LoggingInterval {
	return xnapiesv1.LoggingInterval_LOGGING_INTERVAL_MS61440
}
func CreateLoggingDurationM10() xnapiesv1.LoggingDuration {
	return xnapiesv1.LoggingDuration_LOGGING_DURATION_M10
}
func CreateLoggingDurationM20() xnapiesv1.LoggingDuration {
	return xnapiesv1.LoggingDuration_LOGGING_DURATION_M20
}
func CreateLoggingDurationM40() xnapiesv1.LoggingDuration {
	return xnapiesv1.LoggingDuration_LOGGING_DURATION_M40
}
func CreateLoggingDurationM60() xnapiesv1.LoggingDuration {
	return xnapiesv1.LoggingDuration_LOGGING_DURATION_M60
}
func CreateLoggingDurationM90() xnapiesv1.LoggingDuration {
	return xnapiesv1.LoggingDuration_LOGGING_DURATION_M90
}
func CreateLoggingDurationM120() xnapiesv1.LoggingDuration {
	return xnapiesv1.LoggingDuration_LOGGING_DURATION_M120
}
func CreateLowerLayerPresenceStatusChangeReleaseLowerLayers() xnapiesv1.LowerLayerPresenceStatusChange {
	return xnapiesv1.LowerLayerPresenceStatusChange_LOWER_LAYER_PRESENCE_STATUS_CHANGE_RELEASE_LOWER_LAYERS
}
func CreateLowerLayerPresenceStatusChangeReEstablishLowerLayers() xnapiesv1.LowerLayerPresenceStatusChange {
	return xnapiesv1.LowerLayerPresenceStatusChange_LOWER_LAYER_PRESENCE_STATUS_CHANGE_RE_ESTABLISH_LOWER_LAYERS
}
func CreateLowerLayerPresenceStatusChangeSuspendLowerLayers() xnapiesv1.LowerLayerPresenceStatusChange {
	return xnapiesv1.LowerLayerPresenceStatusChange_LOWER_LAYER_PRESENCE_STATUS_CHANGE_SUSPEND_LOWER_LAYERS
}
func CreateLowerLayerPresenceStatusChangeResumeLowerLayers() xnapiesv1.LowerLayerPresenceStatusChange {
	return xnapiesv1.LowerLayerPresenceStatusChange_LOWER_LAYER_PRESENCE_STATUS_CHANGE_RESUME_LOWER_LAYERS
}
func CreateM1ReportingTriggerPeriodic() xnapiesv1.M1ReportingTrigger {
	return xnapiesv1.M1ReportingTrigger_M1_REPORTING_TRIGGER_PERIODIC
}
func CreateM1ReportingTriggerA2eventtriggered() xnapiesv1.M1ReportingTrigger {
	return xnapiesv1.M1ReportingTrigger_M1_REPORTING_TRIGGER_A2EVENTTRIGGERED
}
func CreateM1ReportingTriggerA2eventtriggeredPeriodic() xnapiesv1.M1ReportingTrigger {
	return xnapiesv1.M1ReportingTrigger_M1_REPORTING_TRIGGER_A2EVENTTRIGGERED_PERIODIC
}
func CreateM4periodMs1024() xnapiesv1.M4Period {
	return xnapiesv1.M4Period_M4PERIOD_MS1024
}
func CreateM4periodMs2048() xnapiesv1.M4Period {
	return xnapiesv1.M4Period_M4PERIOD_MS2048
}
func CreateM4periodMs5120() xnapiesv1.M4Period {
	return xnapiesv1.M4Period_M4PERIOD_MS5120
}
func CreateM4periodMs10240() xnapiesv1.M4Period {
	return xnapiesv1.M4Period_M4PERIOD_MS10240
}
func CreateM4periodMin1() xnapiesv1.M4Period {
	return xnapiesv1.M4Period_M4PERIOD_MIN1
}
func CreateM5periodMs1024() xnapiesv1.M5Period {
	return xnapiesv1.M5Period_M5PERIOD_MS1024
}
func CreateM5periodMs2048() xnapiesv1.M5Period {
	return xnapiesv1.M5Period_M5PERIOD_MS2048
}
func CreateM5periodMs5120() xnapiesv1.M5Period {
	return xnapiesv1.M5Period_M5PERIOD_MS5120
}
func CreateM5periodMs10240() xnapiesv1.M5Period {
	return xnapiesv1.M5Period_M5PERIOD_MS10240
}
func CreateM5periodMin1() xnapiesv1.M5Period {
	return xnapiesv1.M5Period_M5PERIOD_MIN1
}
func CreateM6reportIntervalMs120() xnapiesv1.M6ReportInterval {
	return xnapiesv1.M6ReportInterval_M6REPORT_INTERVAL_MS120
}
func CreateM6reportIntervalMs240() xnapiesv1.M6ReportInterval {
	return xnapiesv1.M6ReportInterval_M6REPORT_INTERVAL_MS240
}
func CreateM6reportIntervalMs480() xnapiesv1.M6ReportInterval {
	return xnapiesv1.M6ReportInterval_M6REPORT_INTERVAL_MS480
}
func CreateM6reportIntervalMs640() xnapiesv1.M6ReportInterval {
	return xnapiesv1.M6ReportInterval_M6REPORT_INTERVAL_MS640
}
func CreateM6reportIntervalMs1024() xnapiesv1.M6ReportInterval {
	return xnapiesv1.M6ReportInterval_M6REPORT_INTERVAL_MS1024
}
func CreateM6reportIntervalMs2048() xnapiesv1.M6ReportInterval {
	return xnapiesv1.M6ReportInterval_M6REPORT_INTERVAL_MS2048
}
func CreateM6reportIntervalMs5120() xnapiesv1.M6ReportInterval {
	return xnapiesv1.M6ReportInterval_M6REPORT_INTERVAL_MS5120
}
func CreateM6reportIntervalMs10240() xnapiesv1.M6ReportInterval {
	return xnapiesv1.M6ReportInterval_M6REPORT_INTERVAL_MS10240
}
func CreateM6reportIntervalMs20480() xnapiesv1.M6ReportInterval {
	return xnapiesv1.M6ReportInterval_M6REPORT_INTERVAL_MS20480
}
func CreateM6reportIntervalMs40960() xnapiesv1.M6ReportInterval {
	return xnapiesv1.M6ReportInterval_M6REPORT_INTERVAL_MS40960
}
func CreateM6reportIntervalMin1() xnapiesv1.M6ReportInterval {
	return xnapiesv1.M6ReportInterval_M6REPORT_INTERVAL_MIN1
}
func CreateM6reportIntervalMin6() xnapiesv1.M6ReportInterval {
	return xnapiesv1.M6ReportInterval_M6REPORT_INTERVAL_MIN6
}
func CreateM6reportIntervalMin12() xnapiesv1.M6ReportInterval {
	return xnapiesv1.M6ReportInterval_M6REPORT_INTERVAL_MIN12
}
func CreateM6reportIntervalMin30() xnapiesv1.M6ReportInterval {
	return xnapiesv1.M6ReportInterval_M6REPORT_INTERVAL_MIN30
}
func CreateMaxIprateBitrate64kbs() xnapiesv1.MaxIprate {
	return xnapiesv1.MaxIprate_MAX_IPRATE_BITRATE64KBS
}
func CreateMaxIprateMaxUerate() xnapiesv1.MaxIprate {
	return xnapiesv1.MaxIprate_MAX_IPRATE_MAX_UERATE
}
func CreateRadioframeAllocationPeriodMbsfnsubframeInfoEutraitemN1() xnapiesv1.RadioframeAllocationPeriodMbsfnsubframeInfoEutraitem {
	return xnapiesv1.RadioframeAllocationPeriodMbsfnsubframeInfoEutraitem_RADIOFRAME_ALLOCATION_PERIOD_MBSFNSUBFRAME_INFO_EUTRAITEM_N1
}
func CreateRadioframeAllocationPeriodMbsfnsubframeInfoEutraitemN2() xnapiesv1.RadioframeAllocationPeriodMbsfnsubframeInfoEutraitem {
	return xnapiesv1.RadioframeAllocationPeriodMbsfnsubframeInfoEutraitem_RADIOFRAME_ALLOCATION_PERIOD_MBSFNSUBFRAME_INFO_EUTRAITEM_N2
}
func CreateRadioframeAllocationPeriodMbsfnsubframeInfoEutraitemN4() xnapiesv1.RadioframeAllocationPeriodMbsfnsubframeInfoEutraitem {
	return xnapiesv1.RadioframeAllocationPeriodMbsfnsubframeInfoEutraitem_RADIOFRAME_ALLOCATION_PERIOD_MBSFNSUBFRAME_INFO_EUTRAITEM_N4
}
func CreateRadioframeAllocationPeriodMbsfnsubframeInfoEutraitemN8() xnapiesv1.RadioframeAllocationPeriodMbsfnsubframeInfoEutraitem {
	return xnapiesv1.RadioframeAllocationPeriodMbsfnsubframeInfoEutraitem_RADIOFRAME_ALLOCATION_PERIOD_MBSFNSUBFRAME_INFO_EUTRAITEM_N8
}
func CreateRadioframeAllocationPeriodMbsfnsubframeInfoEutraitemN16() xnapiesv1.RadioframeAllocationPeriodMbsfnsubframeInfoEutraitem {
	return xnapiesv1.RadioframeAllocationPeriodMbsfnsubframeInfoEutraitem_RADIOFRAME_ALLOCATION_PERIOD_MBSFNSUBFRAME_INFO_EUTRAITEM_N16
}
func CreateRadioframeAllocationPeriodMbsfnsubframeInfoEutraitemN32() xnapiesv1.RadioframeAllocationPeriodMbsfnsubframeInfoEutraitem {
	return xnapiesv1.RadioframeAllocationPeriodMbsfnsubframeInfoEutraitem_RADIOFRAME_ALLOCATION_PERIOD_MBSFNSUBFRAME_INFO_EUTRAITEM_N32
}
func CreateMdtActivationImmediateMdtOnly() xnapiesv1.MdtActivation {
	return xnapiesv1.MdtActivation_MDT_ACTIVATION_IMMEDIATE_MDT_ONLY
}
func CreateMdtActivationImmediateMdtAndTrace() xnapiesv1.MdtActivation {
	return xnapiesv1.MdtActivation_MDT_ACTIVATION_IMMEDIATE_MDT_AND_TRACE
}
func CreateMdtActivationLoggedMdtOnly() xnapiesv1.MdtActivation {
	return xnapiesv1.MdtActivation_MDT_ACTIVATION_LOGGED_MDT_ONLY
}
func CreateCnTypeCntypeRestrictionsForEquivalentItemEpcForbidden() xnapiesv1.CnTypeCntypeRestrictionsForEquivalentItem {
	return xnapiesv1.CnTypeCntypeRestrictionsForEquivalentItem_CN_TYPE_CNTYPE_RESTRICTIONS_FOR_EQUIVALENT_ITEM_EPC_FORBIDDEN
}
func CreateCnTypeCntypeRestrictionsForEquivalentItemFiveGcForbidden() xnapiesv1.CnTypeCntypeRestrictionsForEquivalentItem {
	return xnapiesv1.CnTypeCntypeRestrictionsForEquivalentItem_CN_TYPE_CNTYPE_RESTRICTIONS_FOR_EQUIVALENT_ITEM_FIVE_GC_FORBIDDEN
}
func CreateCntypeRestrictionsForServingEpcForbidden() xnapiesv1.CntypeRestrictionsForServing {
	return xnapiesv1.CntypeRestrictionsForServing_CNTYPE_RESTRICTIONS_FOR_SERVING_EPC_FORBIDDEN
}
func CreateEUtraCoordinationAssistanceInfoCoordinationNotRequired() xnapiesv1.EUTraCoordinationAssistanceInfo {
	return xnapiesv1.EUTraCoordinationAssistanceInfo_E_UTRA_COORDINATION_ASSISTANCE_INFO_COORDINATION_NOT_REQUIRED
}
func CreateNrCoordinationAssistanceInfoCoordinationNotRequired() xnapiesv1.NrCoordinationAssistanceInfo {
	return xnapiesv1.NrCoordinationAssistanceInfo_NR_COORDINATION_ASSISTANCE_INFO_COORDINATION_NOT_REQUIRED
}
func CreateNbioTUlDlAlignmentOffsetKhz7dot5() xnapiesv1.NbioTULDLAlignmentOffset {
	return xnapiesv1.NbioTULDLAlignmentOffset_NBIO_T_UL_DL_ALIGNMENT_OFFSET_KHZ_7DOT5
}
func CreateNbioTUlDlAlignmentOffsetKhz0() xnapiesv1.NbioTULDLAlignmentOffset {
	return xnapiesv1.NbioTULDLAlignmentOffset_NBIO_T_UL_DL_ALIGNMENT_OFFSET_KHZ0
}
func CreateNbioTUlDlAlignmentOffsetKhz7dot5Second() xnapiesv1.NbioTULDLAlignmentOffset {
	return xnapiesv1.NbioTULDLAlignmentOffset_NBIO_T_UL_DL_ALIGNMENT_OFFSET_KHZ7DOT5_SECOND
}
func CreateSubframeAssignmentNedctdmpatternSa0() xnapiesv1.SubframeAssignmentNedctdmpattern {
	return xnapiesv1.SubframeAssignmentNedctdmpattern_SUBFRAME_ASSIGNMENT_NEDCTDMPATTERN_SA0
}
func CreateSubframeAssignmentNedctdmpatternSa1() xnapiesv1.SubframeAssignmentNedctdmpattern {
	return xnapiesv1.SubframeAssignmentNedctdmpattern_SUBFRAME_ASSIGNMENT_NEDCTDMPATTERN_SA1
}
func CreateSubframeAssignmentNedctdmpatternSa2() xnapiesv1.SubframeAssignmentNedctdmpattern {
	return xnapiesv1.SubframeAssignmentNedctdmpattern_SUBFRAME_ASSIGNMENT_NEDCTDMPATTERN_SA2
}
func CreateSubframeAssignmentNedctdmpatternSa3() xnapiesv1.SubframeAssignmentNedctdmpattern {
	return xnapiesv1.SubframeAssignmentNedctdmpattern_SUBFRAME_ASSIGNMENT_NEDCTDMPATTERN_SA3
}
func CreateSubframeAssignmentNedctdmpatternSa4() xnapiesv1.SubframeAssignmentNedctdmpattern {
	return xnapiesv1.SubframeAssignmentNedctdmpattern_SUBFRAME_ASSIGNMENT_NEDCTDMPATTERN_SA4
}
func CreateSubframeAssignmentNedctdmpatternSa5() xnapiesv1.SubframeAssignmentNedctdmpattern {
	return xnapiesv1.SubframeAssignmentNedctdmpattern_SUBFRAME_ASSIGNMENT_NEDCTDMPATTERN_SA5
}
func CreateSubframeAssignmentNedctdmpatternSa6() xnapiesv1.SubframeAssignmentNedctdmpattern {
	return xnapiesv1.SubframeAssignmentNedctdmpattern_SUBFRAME_ASSIGNMENT_NEDCTDMPATTERN_SA6
}
func CreateNprachCpLengthUs66dot7() xnapiesv1.NprachCPLength {
	return xnapiesv1.NprachCPLength_NPRACH_CP_LENGTH_US66DOT7
}
func CreateNprachCpLengthUs266dot7() xnapiesv1.NprachCPLength {
	return xnapiesv1.NprachCPLength_NPRACH_CP_LENGTH_US266DOT7
}
func CreateNprachPreambleFormatFmt0() xnapiesv1.NprachpreambleFormat {
	return xnapiesv1.NprachpreambleFormat_NPRACH_PREAMBLE_FORMAT_FMT0
}
func CreateNprachPreambleFormatFmt1() xnapiesv1.NprachpreambleFormat {
	return xnapiesv1.NprachpreambleFormat_NPRACH_PREAMBLE_FORMAT_FMT1
}
func CreateNprachPreambleFormatFmt2() xnapiesv1.NprachpreambleFormat {
	return xnapiesv1.NprachpreambleFormat_NPRACH_PREAMBLE_FORMAT_FMT2
}
func CreateNprachPreambleFormatFmt0a() xnapiesv1.NprachpreambleFormat {
	return xnapiesv1.NprachpreambleFormat_NPRACH_PREAMBLE_FORMAT_FMT0A
}
func CreateNprachPreambleFormatFmt1a() xnapiesv1.NprachpreambleFormat {
	return xnapiesv1.NprachpreambleFormat_NPRACH_PREAMBLE_FORMAT_FMT1A
}
func CreateNrcyclicPrefixNormal() xnapiesv1.NrcyclicPrefix {
	return xnapiesv1.NrcyclicPrefix_NRCYCLIC_PREFIX_NORMAL
}
func CreateNrcyclicPrefixExtended() xnapiesv1.NrcyclicPrefix {
	return xnapiesv1.NrcyclicPrefix_NRCYCLIC_PREFIX_EXTENDED
}
func CreateNrdlUltransmissionPeriodicityMs0p5() xnapiesv1.NrdlULtransmissionPeriodicity {
	return xnapiesv1.NrdlULtransmissionPeriodicity_NRDL_ULTRANSMISSION_PERIODICITY_MS0P5
}
func CreateNrdlUltransmissionPeriodicityMs0p625() xnapiesv1.NrdlULtransmissionPeriodicity {
	return xnapiesv1.NrdlULtransmissionPeriodicity_NRDL_ULTRANSMISSION_PERIODICITY_MS0P625
}
func CreateNrdlUltransmissionPeriodicityMs1() xnapiesv1.NrdlULtransmissionPeriodicity {
	return xnapiesv1.NrdlULtransmissionPeriodicity_NRDL_ULTRANSMISSION_PERIODICITY_MS1
}
func CreateNrdlUltransmissionPeriodicityMs1p25() xnapiesv1.NrdlULtransmissionPeriodicity {
	return xnapiesv1.NrdlULtransmissionPeriodicity_NRDL_ULTRANSMISSION_PERIODICITY_MS1P25
}
func CreateNrdlUltransmissionPeriodicityMs2() xnapiesv1.NrdlULtransmissionPeriodicity {
	return xnapiesv1.NrdlULtransmissionPeriodicity_NRDL_ULTRANSMISSION_PERIODICITY_MS2
}
func CreateNrdlUltransmissionPeriodicityMs2p5() xnapiesv1.NrdlULtransmissionPeriodicity {
	return xnapiesv1.NrdlULtransmissionPeriodicity_NRDL_ULTRANSMISSION_PERIODICITY_MS2P5
}
func CreateNrdlUltransmissionPeriodicityMs3() xnapiesv1.NrdlULtransmissionPeriodicity {
	return xnapiesv1.NrdlULtransmissionPeriodicity_NRDL_ULTRANSMISSION_PERIODICITY_MS3
}
func CreateNrdlUltransmissionPeriodicityMs4() xnapiesv1.NrdlULtransmissionPeriodicity {
	return xnapiesv1.NrdlULtransmissionPeriodicity_NRDL_ULTRANSMISSION_PERIODICITY_MS4
}
func CreateNrdlUltransmissionPeriodicityMs5() xnapiesv1.NrdlULtransmissionPeriodicity {
	return xnapiesv1.NrdlULtransmissionPeriodicity_NRDL_ULTRANSMISSION_PERIODICITY_MS5
}
func CreateNrdlUltransmissionPeriodicityMs10() xnapiesv1.NrdlULtransmissionPeriodicity {
	return xnapiesv1.NrdlULtransmissionPeriodicity_NRDL_ULTRANSMISSION_PERIODICITY_MS10
}
func CreateNrdlUltransmissionPeriodicityMs20() xnapiesv1.NrdlULtransmissionPeriodicity {
	return xnapiesv1.NrdlULtransmissionPeriodicity_NRDL_ULTRANSMISSION_PERIODICITY_MS20
}
func CreateNrdlUltransmissionPeriodicityMs40() xnapiesv1.NrdlULtransmissionPeriodicity {
	return xnapiesv1.NrdlULtransmissionPeriodicity_NRDL_ULTRANSMISSION_PERIODICITY_MS40
}
func CreateNrdlUltransmissionPeriodicityMs60() xnapiesv1.NrdlULtransmissionPeriodicity {
	return xnapiesv1.NrdlULtransmissionPeriodicity_NRDL_ULTRANSMISSION_PERIODICITY_MS60
}
func CreateNrdlUltransmissionPeriodicityMs80() xnapiesv1.NrdlULtransmissionPeriodicity {
	return xnapiesv1.NrdlULtransmissionPeriodicity_NRDL_ULTRANSMISSION_PERIODICITY_MS80
}
func CreateNrdlUltransmissionPeriodicityMs100() xnapiesv1.NrdlULtransmissionPeriodicity {
	return xnapiesv1.NrdlULtransmissionPeriodicity_NRDL_ULTRANSMISSION_PERIODICITY_MS100
}
func CreateNrdlUltransmissionPeriodicityMs120() xnapiesv1.NrdlULtransmissionPeriodicity {
	return xnapiesv1.NrdlULtransmissionPeriodicity_NRDL_ULTRANSMISSION_PERIODICITY_MS120
}
func CreateNrdlUltransmissionPeriodicityMs140() xnapiesv1.NrdlULtransmissionPeriodicity {
	return xnapiesv1.NrdlULtransmissionPeriodicity_NRDL_ULTRANSMISSION_PERIODICITY_MS140
}
func CreateNrdlUltransmissionPeriodicityMs160() xnapiesv1.NrdlULtransmissionPeriodicity {
	return xnapiesv1.NrdlULtransmissionPeriodicity_NRDL_ULTRANSMISSION_PERIODICITY_MS160
}
func CreateNrnrbNrb11() xnapiesv1.Nrnrb {
	return xnapiesv1.Nrnrb_NRNRB_NRB11
}
func CreateNrnrbNrb18() xnapiesv1.Nrnrb {
	return xnapiesv1.Nrnrb_NRNRB_NRB18
}
func CreateNrnrbNrb24() xnapiesv1.Nrnrb {
	return xnapiesv1.Nrnrb_NRNRB_NRB24
}
func CreateNrnrbNrb25() xnapiesv1.Nrnrb {
	return xnapiesv1.Nrnrb_NRNRB_NRB25
}
func CreateNrnrbNrb31() xnapiesv1.Nrnrb {
	return xnapiesv1.Nrnrb_NRNRB_NRB31
}
func CreateNrnrbNrb32() xnapiesv1.Nrnrb {
	return xnapiesv1.Nrnrb_NRNRB_NRB32
}
func CreateNrnrbNrb38() xnapiesv1.Nrnrb {
	return xnapiesv1.Nrnrb_NRNRB_NRB38
}
func CreateNrnrbNrb51() xnapiesv1.Nrnrb {
	return xnapiesv1.Nrnrb_NRNRB_NRB51
}
func CreateNrnrbNrb52() xnapiesv1.Nrnrb {
	return xnapiesv1.Nrnrb_NRNRB_NRB52
}
func CreateNrnrbNrb65() xnapiesv1.Nrnrb {
	return xnapiesv1.Nrnrb_NRNRB_NRB65
}
func CreateNrnrbNrb66() xnapiesv1.Nrnrb {
	return xnapiesv1.Nrnrb_NRNRB_NRB66
}
func CreateNrnrbNrb78() xnapiesv1.Nrnrb {
	return xnapiesv1.Nrnrb_NRNRB_NRB78
}
func CreateNrnrbNrb79() xnapiesv1.Nrnrb {
	return xnapiesv1.Nrnrb_NRNRB_NRB79
}
func CreateNrnrbNrb93() xnapiesv1.Nrnrb {
	return xnapiesv1.Nrnrb_NRNRB_NRB93
}
func CreateNrnrbNrb106() xnapiesv1.Nrnrb {
	return xnapiesv1.Nrnrb_NRNRB_NRB106
}
func CreateNrnrbNrb107() xnapiesv1.Nrnrb {
	return xnapiesv1.Nrnrb_NRNRB_NRB107
}
func CreateNrnrbNrb121() xnapiesv1.Nrnrb {
	return xnapiesv1.Nrnrb_NRNRB_NRB121
}
func CreateNrnrbNrb132() xnapiesv1.Nrnrb {
	return xnapiesv1.Nrnrb_NRNRB_NRB132
}
func CreateNrnrbNrb133() xnapiesv1.Nrnrb {
	return xnapiesv1.Nrnrb_NRNRB_NRB133
}
func CreateNrnrbNrb135() xnapiesv1.Nrnrb {
	return xnapiesv1.Nrnrb_NRNRB_NRB135
}
func CreateNrnrbNrb160() xnapiesv1.Nrnrb {
	return xnapiesv1.Nrnrb_NRNRB_NRB160
}
func CreateNrnrbNrb162() xnapiesv1.Nrnrb {
	return xnapiesv1.Nrnrb_NRNRB_NRB162
}
func CreateNrnrbNrb189() xnapiesv1.Nrnrb {
	return xnapiesv1.Nrnrb_NRNRB_NRB189
}
func CreateNrnrbNrb216() xnapiesv1.Nrnrb {
	return xnapiesv1.Nrnrb_NRNRB_NRB216
}
func CreateNrnrbNrb217() xnapiesv1.Nrnrb {
	return xnapiesv1.Nrnrb_NRNRB_NRB217
}
func CreateNrnrbNrb245() xnapiesv1.Nrnrb {
	return xnapiesv1.Nrnrb_NRNRB_NRB245
}
func CreateNrnrbNrb264() xnapiesv1.Nrnrb {
	return xnapiesv1.Nrnrb_NRNRB_NRB264
}
func CreateNrnrbNrb270() xnapiesv1.Nrnrb {
	return xnapiesv1.Nrnrb_NRNRB_NRB270
}
func CreateNrnrbNrb273() xnapiesv1.Nrnrb {
	return xnapiesv1.Nrnrb_NRNRB_NRB273
}
func CreateNrscsScs15() xnapiesv1.Nrscs {
	return xnapiesv1.Nrscs_NRSCS_SCS15
}
func CreateNrscsScs30() xnapiesv1.Nrscs {
	return xnapiesv1.Nrscs_NRSCS_SCS30
}
func CreateNrscsScs60() xnapiesv1.Nrscs {
	return xnapiesv1.Nrscs_NRSCS_SCS60
}
func CreateNrscsScs120() xnapiesv1.Nrscs {
	return xnapiesv1.Nrscs_NRSCS_SCS120
}
func CreateNumberOfAntennaPortsEUtraAn1() xnapiesv1.NumberOfAntennaPortsEUTra {
	return xnapiesv1.NumberOfAntennaPortsEUTra_NUMBER_OF_ANTENNA_PORTS_E_UTRA_AN1
}
func CreateNumberOfAntennaPortsEUtraAn2() xnapiesv1.NumberOfAntennaPortsEUTra {
	return xnapiesv1.NumberOfAntennaPortsEUTra_NUMBER_OF_ANTENNA_PORTS_E_UTRA_AN2
}
func CreateNumberOfAntennaPortsEUtraAn4() xnapiesv1.NumberOfAntennaPortsEUTra {
	return xnapiesv1.NumberOfAntennaPortsEUTra_NUMBER_OF_ANTENNA_PORTS_E_UTRA_AN4
}
func CreateNonGbrresourcesOfferedTrue() xnapiesv1.NonGbrresourcesOffered {
	return xnapiesv1.NonGbrresourcesOffered_NON_GBRRESOURCES_OFFERED_TRUE
}
func CreateOffsetOfNbiotChannelNumberToEarfcnMinusTen() xnapiesv1.OffsetOfNbiotChannelNumberToEarfcn {
	return xnapiesv1.OffsetOfNbiotChannelNumberToEarfcn_OFFSET_OF_NBIOT_CHANNEL_NUMBER_TO_EARFCN_MINUS_TEN
}
func CreateOffsetOfNbiotChannelNumberToEarfcnMinusNine() xnapiesv1.OffsetOfNbiotChannelNumberToEarfcn {
	return xnapiesv1.OffsetOfNbiotChannelNumberToEarfcn_OFFSET_OF_NBIOT_CHANNEL_NUMBER_TO_EARFCN_MINUS_NINE
}
func CreateOffsetOfNbiotChannelNumberToEarfcnMinusEightDotFive() xnapiesv1.OffsetOfNbiotChannelNumberToEarfcn {
	return xnapiesv1.OffsetOfNbiotChannelNumberToEarfcn_OFFSET_OF_NBIOT_CHANNEL_NUMBER_TO_EARFCN_MINUS_EIGHT_DOT_FIVE
}
func CreateOffsetOfNbiotChannelNumberToEarfcnMinusEight() xnapiesv1.OffsetOfNbiotChannelNumberToEarfcn {
	return xnapiesv1.OffsetOfNbiotChannelNumberToEarfcn_OFFSET_OF_NBIOT_CHANNEL_NUMBER_TO_EARFCN_MINUS_EIGHT
}
func CreateOffsetOfNbiotChannelNumberToEarfcnMinusSeven() xnapiesv1.OffsetOfNbiotChannelNumberToEarfcn {
	return xnapiesv1.OffsetOfNbiotChannelNumberToEarfcn_OFFSET_OF_NBIOT_CHANNEL_NUMBER_TO_EARFCN_MINUS_SEVEN
}
func CreateOffsetOfNbiotChannelNumberToEarfcnMinusSix() xnapiesv1.OffsetOfNbiotChannelNumberToEarfcn {
	return xnapiesv1.OffsetOfNbiotChannelNumberToEarfcn_OFFSET_OF_NBIOT_CHANNEL_NUMBER_TO_EARFCN_MINUS_SIX
}
func CreateOffsetOfNbiotChannelNumberToEarfcnMinusFive() xnapiesv1.OffsetOfNbiotChannelNumberToEarfcn {
	return xnapiesv1.OffsetOfNbiotChannelNumberToEarfcn_OFFSET_OF_NBIOT_CHANNEL_NUMBER_TO_EARFCN_MINUS_FIVE
}
func CreateOffsetOfNbiotChannelNumberToEarfcnMinusFourDotFive() xnapiesv1.OffsetOfNbiotChannelNumberToEarfcn {
	return xnapiesv1.OffsetOfNbiotChannelNumberToEarfcn_OFFSET_OF_NBIOT_CHANNEL_NUMBER_TO_EARFCN_MINUS_FOUR_DOT_FIVE
}
func CreateOffsetOfNbiotChannelNumberToEarfcnMinusFour() xnapiesv1.OffsetOfNbiotChannelNumberToEarfcn {
	return xnapiesv1.OffsetOfNbiotChannelNumberToEarfcn_OFFSET_OF_NBIOT_CHANNEL_NUMBER_TO_EARFCN_MINUS_FOUR
}
func CreateOffsetOfNbiotChannelNumberToEarfcnMinusThree() xnapiesv1.OffsetOfNbiotChannelNumberToEarfcn {
	return xnapiesv1.OffsetOfNbiotChannelNumberToEarfcn_OFFSET_OF_NBIOT_CHANNEL_NUMBER_TO_EARFCN_MINUS_THREE
}
func CreateOffsetOfNbiotChannelNumberToEarfcnMinusTwo() xnapiesv1.OffsetOfNbiotChannelNumberToEarfcn {
	return xnapiesv1.OffsetOfNbiotChannelNumberToEarfcn_OFFSET_OF_NBIOT_CHANNEL_NUMBER_TO_EARFCN_MINUS_TWO
}
func CreateOffsetOfNbiotChannelNumberToEarfcnMinusOne() xnapiesv1.OffsetOfNbiotChannelNumberToEarfcn {
	return xnapiesv1.OffsetOfNbiotChannelNumberToEarfcn_OFFSET_OF_NBIOT_CHANNEL_NUMBER_TO_EARFCN_MINUS_ONE
}
func CreateOffsetOfNbiotChannelNumberToEarfcnMinusZeroDotFive() xnapiesv1.OffsetOfNbiotChannelNumberToEarfcn {
	return xnapiesv1.OffsetOfNbiotChannelNumberToEarfcn_OFFSET_OF_NBIOT_CHANNEL_NUMBER_TO_EARFCN_MINUS_ZERO_DOT_FIVE
}
func CreateOffsetOfNbiotChannelNumberToEarfcnZero() xnapiesv1.OffsetOfNbiotChannelNumberToEarfcn {
	return xnapiesv1.OffsetOfNbiotChannelNumberToEarfcn_OFFSET_OF_NBIOT_CHANNEL_NUMBER_TO_EARFCN_ZERO
}
func CreateOffsetOfNbiotChannelNumberToEarfcnOne() xnapiesv1.OffsetOfNbiotChannelNumberToEarfcn {
	return xnapiesv1.OffsetOfNbiotChannelNumberToEarfcn_OFFSET_OF_NBIOT_CHANNEL_NUMBER_TO_EARFCN_ONE
}
func CreateOffsetOfNbiotChannelNumberToEarfcnTwo() xnapiesv1.OffsetOfNbiotChannelNumberToEarfcn {
	return xnapiesv1.OffsetOfNbiotChannelNumberToEarfcn_OFFSET_OF_NBIOT_CHANNEL_NUMBER_TO_EARFCN_TWO
}
func CreateOffsetOfNbiotChannelNumberToEarfcnThree() xnapiesv1.OffsetOfNbiotChannelNumberToEarfcn {
	return xnapiesv1.OffsetOfNbiotChannelNumberToEarfcn_OFFSET_OF_NBIOT_CHANNEL_NUMBER_TO_EARFCN_THREE
}
func CreateOffsetOfNbiotChannelNumberToEarfcnThreeDotFive() xnapiesv1.OffsetOfNbiotChannelNumberToEarfcn {
	return xnapiesv1.OffsetOfNbiotChannelNumberToEarfcn_OFFSET_OF_NBIOT_CHANNEL_NUMBER_TO_EARFCN_THREE_DOT_FIVE
}
func CreateOffsetOfNbiotChannelNumberToEarfcnFour() xnapiesv1.OffsetOfNbiotChannelNumberToEarfcn {
	return xnapiesv1.OffsetOfNbiotChannelNumberToEarfcn_OFFSET_OF_NBIOT_CHANNEL_NUMBER_TO_EARFCN_FOUR
}
func CreateOffsetOfNbiotChannelNumberToEarfcnFive() xnapiesv1.OffsetOfNbiotChannelNumberToEarfcn {
	return xnapiesv1.OffsetOfNbiotChannelNumberToEarfcn_OFFSET_OF_NBIOT_CHANNEL_NUMBER_TO_EARFCN_FIVE
}
func CreateOffsetOfNbiotChannelNumberToEarfcnSix() xnapiesv1.OffsetOfNbiotChannelNumberToEarfcn {
	return xnapiesv1.OffsetOfNbiotChannelNumberToEarfcn_OFFSET_OF_NBIOT_CHANNEL_NUMBER_TO_EARFCN_SIX
}
func CreateOffsetOfNbiotChannelNumberToEarfcnSeven() xnapiesv1.OffsetOfNbiotChannelNumberToEarfcn {
	return xnapiesv1.OffsetOfNbiotChannelNumberToEarfcn_OFFSET_OF_NBIOT_CHANNEL_NUMBER_TO_EARFCN_SEVEN
}
func CreateOffsetOfNbiotChannelNumberToEarfcnSevenDotFive() xnapiesv1.OffsetOfNbiotChannelNumberToEarfcn {
	return xnapiesv1.OffsetOfNbiotChannelNumberToEarfcn_OFFSET_OF_NBIOT_CHANNEL_NUMBER_TO_EARFCN_SEVEN_DOT_FIVE
}
func CreateOffsetOfNbiotChannelNumberToEarfcnEight() xnapiesv1.OffsetOfNbiotChannelNumberToEarfcn {
	return xnapiesv1.OffsetOfNbiotChannelNumberToEarfcn_OFFSET_OF_NBIOT_CHANNEL_NUMBER_TO_EARFCN_EIGHT
}
func CreateOffsetOfNbiotChannelNumberToEarfcnNine() xnapiesv1.OffsetOfNbiotChannelNumberToEarfcn {
	return xnapiesv1.OffsetOfNbiotChannelNumberToEarfcn_OFFSET_OF_NBIOT_CHANNEL_NUMBER_TO_EARFCN_NINE
}
func CreatePedestrianUeAuthorized() xnapiesv1.PedestrianUe {
	return xnapiesv1.PedestrianUe_PEDESTRIAN_UE_AUTHORIZED
}
func CreatePedestrianUeNotAuthorized() xnapiesv1.PedestrianUe {
	return xnapiesv1.PedestrianUe_PEDESTRIAN_UE_NOT_AUTHORIZED
}
func CreatePagingDrxV32() xnapiesv1.PagingDrx {
	return xnapiesv1.PagingDrx_PAGING_DRX_V32
}
func CreatePagingDrxV64() xnapiesv1.PagingDrx {
	return xnapiesv1.PagingDrx_PAGING_DRX_V64
}
func CreatePagingDrxV128() xnapiesv1.PagingDrx {
	return xnapiesv1.PagingDrx_PAGING_DRX_V128
}
func CreatePagingDrxV256() xnapiesv1.PagingDrx {
	return xnapiesv1.PagingDrx_PAGING_DRX_V256
}
func CreatePagingDrxV512() xnapiesv1.PagingDrx {
	return xnapiesv1.PagingDrx_PAGING_DRX_V512
}
func CreatePagingDrxV1024() xnapiesv1.PagingDrx {
	return xnapiesv1.PagingDrx_PAGING_DRX_V1024
}
func CreatePagingEDrxCycleHfhalf() xnapiesv1.PagingeDrxCycle {
	return xnapiesv1.PagingeDrxCycle_PAGING_E_DRX_CYCLE_HFHALF
}
func CreatePagingEDrxCycleHf1() xnapiesv1.PagingeDrxCycle {
	return xnapiesv1.PagingeDrxCycle_PAGING_E_DRX_CYCLE_HF1
}
func CreatePagingEDrxCycleHf2() xnapiesv1.PagingeDrxCycle {
	return xnapiesv1.PagingeDrxCycle_PAGING_E_DRX_CYCLE_HF2
}
func CreatePagingEDrxCycleHf4() xnapiesv1.PagingeDrxCycle {
	return xnapiesv1.PagingeDrxCycle_PAGING_E_DRX_CYCLE_HF4
}
func CreatePagingEDrxCycleHf6() xnapiesv1.PagingeDrxCycle {
	return xnapiesv1.PagingeDrxCycle_PAGING_E_DRX_CYCLE_HF6
}
func CreatePagingEDrxCycleHf8() xnapiesv1.PagingeDrxCycle {
	return xnapiesv1.PagingeDrxCycle_PAGING_E_DRX_CYCLE_HF8
}
func CreatePagingEDrxCycleHf10() xnapiesv1.PagingeDrxCycle {
	return xnapiesv1.PagingeDrxCycle_PAGING_E_DRX_CYCLE_HF10
}
func CreatePagingEDrxCycleHf12() xnapiesv1.PagingeDrxCycle {
	return xnapiesv1.PagingeDrxCycle_PAGING_E_DRX_CYCLE_HF12
}
func CreatePagingEDrxCycleHf14() xnapiesv1.PagingeDrxCycle {
	return xnapiesv1.PagingeDrxCycle_PAGING_E_DRX_CYCLE_HF14
}
func CreatePagingEDrxCycleHf16() xnapiesv1.PagingeDrxCycle {
	return xnapiesv1.PagingeDrxCycle_PAGING_E_DRX_CYCLE_HF16
}
func CreatePagingEDrxCycleHf32() xnapiesv1.PagingeDrxCycle {
	return xnapiesv1.PagingeDrxCycle_PAGING_E_DRX_CYCLE_HF32
}
func CreatePagingEDrxCycleHf64() xnapiesv1.PagingeDrxCycle {
	return xnapiesv1.PagingeDrxCycle_PAGING_E_DRX_CYCLE_HF64
}
func CreatePagingEDrxCycleHf128() xnapiesv1.PagingeDrxCycle {
	return xnapiesv1.PagingeDrxCycle_PAGING_E_DRX_CYCLE_HF128
}
func CreatePagingEDrxCycleHf256() xnapiesv1.PagingeDrxCycle {
	return xnapiesv1.PagingeDrxCycle_PAGING_E_DRX_CYCLE_HF256
}
func CreatePagingTimeWindowS1() xnapiesv1.PagingTimeWindow {
	return xnapiesv1.PagingTimeWindow_PAGING_TIME_WINDOW_S1
}
func CreatePagingTimeWindowS2() xnapiesv1.PagingTimeWindow {
	return xnapiesv1.PagingTimeWindow_PAGING_TIME_WINDOW_S2
}
func CreatePagingTimeWindowS3() xnapiesv1.PagingTimeWindow {
	return xnapiesv1.PagingTimeWindow_PAGING_TIME_WINDOW_S3
}
func CreatePagingTimeWindowS4() xnapiesv1.PagingTimeWindow {
	return xnapiesv1.PagingTimeWindow_PAGING_TIME_WINDOW_S4
}
func CreatePagingTimeWindowS5() xnapiesv1.PagingTimeWindow {
	return xnapiesv1.PagingTimeWindow_PAGING_TIME_WINDOW_S5
}
func CreatePagingTimeWindowS6() xnapiesv1.PagingTimeWindow {
	return xnapiesv1.PagingTimeWindow_PAGING_TIME_WINDOW_S6
}
func CreatePagingTimeWindowS7() xnapiesv1.PagingTimeWindow {
	return xnapiesv1.PagingTimeWindow_PAGING_TIME_WINDOW_S7
}
func CreatePagingTimeWindowS8() xnapiesv1.PagingTimeWindow {
	return xnapiesv1.PagingTimeWindow_PAGING_TIME_WINDOW_S8
}
func CreatePagingTimeWindowS9() xnapiesv1.PagingTimeWindow {
	return xnapiesv1.PagingTimeWindow_PAGING_TIME_WINDOW_S9
}
func CreatePagingTimeWindowS10() xnapiesv1.PagingTimeWindow {
	return xnapiesv1.PagingTimeWindow_PAGING_TIME_WINDOW_S10
}
func CreatePagingTimeWindowS11() xnapiesv1.PagingTimeWindow {
	return xnapiesv1.PagingTimeWindow_PAGING_TIME_WINDOW_S11
}
func CreatePagingTimeWindowS12() xnapiesv1.PagingTimeWindow {
	return xnapiesv1.PagingTimeWindow_PAGING_TIME_WINDOW_S12
}
func CreatePagingTimeWindowS13() xnapiesv1.PagingTimeWindow {
	return xnapiesv1.PagingTimeWindow_PAGING_TIME_WINDOW_S13
}
func CreatePagingTimeWindowS14() xnapiesv1.PagingTimeWindow {
	return xnapiesv1.PagingTimeWindow_PAGING_TIME_WINDOW_S14
}
func CreatePagingTimeWindowS15() xnapiesv1.PagingTimeWindow {
	return xnapiesv1.PagingTimeWindow_PAGING_TIME_WINDOW_S15
}
func CreatePagingTimeWindowS16() xnapiesv1.PagingTimeWindow {
	return xnapiesv1.PagingTimeWindow_PAGING_TIME_WINDOW_S16
}
func CreatePagingPriorityPriolevel1() xnapiesv1.PagingPriority {
	return xnapiesv1.PagingPriority_PAGING_PRIORITY_PRIOLEVEL1
}
func CreatePagingPriorityPriolevel2() xnapiesv1.PagingPriority {
	return xnapiesv1.PagingPriority_PAGING_PRIORITY_PRIOLEVEL2
}
func CreatePagingPriorityPriolevel3() xnapiesv1.PagingPriority {
	return xnapiesv1.PagingPriority_PAGING_PRIORITY_PRIOLEVEL3
}
func CreatePagingPriorityPriolevel4() xnapiesv1.PagingPriority {
	return xnapiesv1.PagingPriority_PAGING_PRIORITY_PRIOLEVEL4
}
func CreatePagingPriorityPriolevel5() xnapiesv1.PagingPriority {
	return xnapiesv1.PagingPriority_PAGING_PRIORITY_PRIOLEVEL5
}
func CreatePagingPriorityPriolevel6() xnapiesv1.PagingPriority {
	return xnapiesv1.PagingPriority_PAGING_PRIORITY_PRIOLEVEL6
}
func CreatePagingPriorityPriolevel7() xnapiesv1.PagingPriority {
	return xnapiesv1.PagingPriority_PAGING_PRIORITY_PRIOLEVEL7
}
func CreatePagingPriorityPriolevel8() xnapiesv1.PagingPriority {
	return xnapiesv1.PagingPriority_PAGING_PRIORITY_PRIOLEVEL8
}
func CreatePartialListIndicatorPartial() xnapiesv1.PartialListIndicator {
	return xnapiesv1.PartialListIndicator_PARTIAL_LIST_INDICATOR_PARTIAL
}
func CreateFromSngrannodePdcpchangeIndicationSNgRanNodeKeyUpdateRequired() xnapiesv1.FromSngrannodePdcpchangeIndication {
	return xnapiesv1.FromSngrannodePdcpchangeIndication_FROM_SNGRANNODE_PDCPCHANGE_INDICATION_S_NG_RAN_NODE_KEY_UPDATE_REQUIRED
}
func CreateFromSngrannodePdcpchangeIndicationPdcpDataRecoveryRequired() xnapiesv1.FromSngrannodePdcpchangeIndication {
	return xnapiesv1.FromSngrannodePdcpchangeIndication_FROM_SNGRANNODE_PDCPCHANGE_INDICATION_PDCP_DATA_RECOVERY_REQUIRED
}
func CreateFromMngrannodePdcpchangeIndicationPdcpDataRecoveryRequired() xnapiesv1.FromMngrannodePdcpchangeIndication {
	return xnapiesv1.FromMngrannodePdcpchangeIndication_FROM_MNGRANNODE_PDCPCHANGE_INDICATION_PDCP_DATA_RECOVERY_REQUIRED
}
func CreatePdcpduplicationConfigurationConfigured() xnapiesv1.PdcpduplicationConfiguration {
	return xnapiesv1.PdcpduplicationConfiguration_PDCPDUPLICATION_CONFIGURATION_CONFIGURED
}
func CreatePdcpduplicationConfigurationDeConfigured() xnapiesv1.PdcpduplicationConfiguration {
	return xnapiesv1.PdcpduplicationConfiguration_PDCPDUPLICATION_CONFIGURATION_DE_CONFIGURED
}
func CreateUlPdcpsnlengthPdcpsnlengthV12bits() xnapiesv1.UlPdcpsnlengthPdcpsnlength {
	return xnapiesv1.UlPdcpsnlengthPdcpsnlength_UL_PDCPSNLENGTH_PDCPSNLENGTH_V12BITS
}
func CreateUlPdcpsnlengthPdcpsnlengthV18bits() xnapiesv1.UlPdcpsnlengthPdcpsnlength {
	return xnapiesv1.UlPdcpsnlengthPdcpsnlength_UL_PDCPSNLENGTH_PDCPSNLENGTH_V18BITS
}
func CreateDlPdcpsnlengthPdcpsnlengthV12bits() xnapiesv1.DlPdcpsnlengthPdcpsnlength {
	return xnapiesv1.DlPdcpsnlengthPdcpsnlength_DL_PDCPSNLENGTH_PDCPSNLENGTH_V12BITS
}
func CreateDlPdcpsnlengthPdcpsnlengthV18bits() xnapiesv1.DlPdcpsnlengthPdcpsnlength {
	return xnapiesv1.DlPdcpsnlengthPdcpsnlength_DL_PDCPSNLENGTH_PDCPSNLENGTH_V18BITS
}
func CreateDLngutnlinformationUnchangedPdusessionResourceAdmittedInfoTrue() xnapiesv1.DlngutnlinformationUnchangedPdusessionResourceAdmittedInfo {
	return xnapiesv1.DlngutnlinformationUnchangedPdusessionResourceAdmittedInfo_D_LNGUTNLINFORMATION_UNCHANGED_PDUSESSION_RESOURCE_ADMITTED_INFO_TRUE
}
func CreateRAttypePdusessionUsageReportNr() xnapiesv1.RattypePdusessionUsageReport {
	return xnapiesv1.RattypePdusessionUsageReport_R_ATTYPE_PDUSESSION_USAGE_REPORT_NR
}
func CreateRAttypePdusessionUsageReportEutra() xnapiesv1.RattypePdusessionUsageReport {
	return xnapiesv1.RattypePdusessionUsageReport_R_ATTYPE_PDUSESSION_USAGE_REPORT_EUTRA
}
func CreateRAttypePdusessionUsageReportNrUnlicensed() xnapiesv1.RattypePdusessionUsageReport {
	return xnapiesv1.RattypePdusessionUsageReport_R_ATTYPE_PDUSESSION_USAGE_REPORT_NR_UNLICENSED
}
func CreateRAttypePdusessionUsageReportEUtraUnlicensed() xnapiesv1.RattypePdusessionUsageReport {
	return xnapiesv1.RattypePdusessionUsageReport_R_ATTYPE_PDUSESSION_USAGE_REPORT_E_UTRA_UNLICENSED
}
func CreatePdusessionTypeIpv4() xnapiesv1.PdusessionType {
	return xnapiesv1.PdusessionType_PDUSESSION_TYPE_IPV4
}
func CreatePdusessionTypeIpv6() xnapiesv1.PdusessionType {
	return xnapiesv1.PdusessionType_PDUSESSION_TYPE_IPV6
}
func CreatePdusessionTypeIpv4v6() xnapiesv1.PdusessionType {
	return xnapiesv1.PdusessionType_PDUSESSION_TYPE_IPV4V6
}
func CreatePdusessionTypeEthernet() xnapiesv1.PdusessionType {
	return xnapiesv1.PdusessionType_PDUSESSION_TYPE_ETHERNET
}
func CreatePdusessionTypeUnstructured() xnapiesv1.PdusessionType {
	return xnapiesv1.PdusessionType_PDUSESSION_TYPE_UNSTRUCTURED
}
func CreatePniNpnRestrictedInformationRestriced() xnapiesv1.PniNPnRestrictedInformation {
	return xnapiesv1.PniNPnRestrictedInformation_PNI_NPN_RESTRICTED_INFORMATION_RESTRICED
}
func CreatePniNpnRestrictedInformationNotRestricted() xnapiesv1.PniNPnRestrictedInformation {
	return xnapiesv1.PniNPnRestrictedInformation_PNI_NPN_RESTRICTED_INFORMATION_NOT_RESTRICTED
}
func CreateResourceTypeProtectedEutraresourceItemDownlinknonCrs() xnapiesv1.ResourceTypeProtectedEutraresourceItem {
	return xnapiesv1.ResourceTypeProtectedEutraresourceItem_RESOURCE_TYPE_PROTECTED_EUTRARESOURCE_ITEM_DOWNLINKNON_CRS
}
func CreateResourceTypeProtectedEutraresourceItemCRs() xnapiesv1.ResourceTypeProtectedEutraresourceItem {
	return xnapiesv1.ResourceTypeProtectedEutraresourceItem_RESOURCE_TYPE_PROTECTED_EUTRARESOURCE_ITEM_C_RS
}
func CreateResourceTypeProtectedEutraresourceItemUplink() xnapiesv1.ResourceTypeProtectedEutraresourceItem {
	return xnapiesv1.ResourceTypeProtectedEutraresourceItem_RESOURCE_TYPE_PROTECTED_EUTRARESOURCE_ITEM_UPLINK
}
func CreateAdditionalQoSflowInfoQoSflowLevelQoSparametersMoreLikely() xnapiesv1.AdditionalQoSflowInfoQoSflowLevelQoSparameters {
	return xnapiesv1.AdditionalQoSflowInfoQoSflowLevelQoSparameters_ADDITIONAL_QO_SFLOW_INFO_QO_SFLOW_LEVEL_QO_SPARAMETERS_MORE_LIKELY
}
func CreateQoSflowMappingIndicationUl() xnapiesv1.QoSflowMappingIndication {
	return xnapiesv1.QoSflowMappingIndication_QO_SFLOW_MAPPING_INDICATION_UL
}
func CreateQoSflowMappingIndicationDl() xnapiesv1.QoSflowMappingIndication {
	return xnapiesv1.QoSflowMappingIndication_QO_SFLOW_MAPPING_INDICATION_DL
}
func CreateNotificationInformationQoSflowNotifyItemFulfilled() xnapiesv1.NotificationInformationQoSflowNotifyItem {
	return xnapiesv1.NotificationInformationQoSflowNotifyItem_NOTIFICATION_INFORMATION_QO_SFLOW_NOTIFY_ITEM_FULFILLED
}
func CreateNotificationInformationQoSflowNotifyItemNotFulfilled() xnapiesv1.NotificationInformationQoSflowNotifyItem {
	return xnapiesv1.NotificationInformationQoSflowNotifyItem_NOTIFICATION_INFORMATION_QO_SFLOW_NOTIFY_ITEM_NOT_FULFILLED
}
func CreateRAttypeQoSflowsUsageReportItemNr() xnapiesv1.RattypeQoSflowsUsageReportItem {
	return xnapiesv1.RattypeQoSflowsUsageReportItem_R_ATTYPE_QO_SFLOWS_USAGE_REPORT_ITEM_NR
}
func CreateRAttypeQoSflowsUsageReportItemEutra() xnapiesv1.RattypeQoSflowsUsageReportItem {
	return xnapiesv1.RattypeQoSflowsUsageReportItem_R_ATTYPE_QO_SFLOWS_USAGE_REPORT_ITEM_EUTRA
}
func CreateRAttypeQoSflowsUsageReportItemNrUnlicensed() xnapiesv1.RattypeQoSflowsUsageReportItem {
	return xnapiesv1.RattypeQoSflowsUsageReportItem_R_ATTYPE_QO_SFLOWS_USAGE_REPORT_ITEM_NR_UNLICENSED
}
func CreateRAttypeQoSflowsUsageReportItemEUtraUnlicensed() xnapiesv1.RattypeQoSflowsUsageReportItem {
	return xnapiesv1.RattypeQoSflowsUsageReportItem_R_ATTYPE_QO_SFLOWS_USAGE_REPORT_ITEM_E_UTRA_UNLICENSED
}
func CreateQosMonitoringRequestUl() xnapiesv1.QosMonitoringRequest {
	return xnapiesv1.QosMonitoringRequest_QOS_MONITORING_REQUEST_UL
}
func CreateQosMonitoringRequestDl() xnapiesv1.QosMonitoringRequest {
	return xnapiesv1.QosMonitoringRequest_QOS_MONITORING_REQUEST_DL
}
func CreateQosMonitoringRequestBoth() xnapiesv1.QosMonitoringRequest {
	return xnapiesv1.QosMonitoringRequest_QOS_MONITORING_REQUEST_BOTH
}
func CreateQoSmonitoringDisabledTrue() xnapiesv1.QoSmonitoringDisabled {
	return xnapiesv1.QoSmonitoringDisabled_QO_SMONITORING_DISABLED_TRUE
}
func CreateRangeM50() xnapiesv1.Range {
	return xnapiesv1.Range_RANGE_M50
}
func CreateRangeM80() xnapiesv1.Range {
	return xnapiesv1.Range_RANGE_M80
}
func CreateRangeM180() xnapiesv1.Range {
	return xnapiesv1.Range_RANGE_M180
}
func CreateRangeM200() xnapiesv1.Range {
	return xnapiesv1.Range_RANGE_M200
}
func CreateRangeM350() xnapiesv1.Range {
	return xnapiesv1.Range_RANGE_M350
}
func CreateRangeM400() xnapiesv1.Range {
	return xnapiesv1.Range_RANGE_M400
}
func CreateRangeM500() xnapiesv1.Range {
	return xnapiesv1.Range_RANGE_M500
}
func CreateRangeM700() xnapiesv1.Range {
	return xnapiesv1.Range_RANGE_M700
}
func CreateRangeM1000() xnapiesv1.Range {
	return xnapiesv1.Range_RANGE_M1000
}
func CreateNextPagingAreaScopeRanpagingAttemptInfoSame() xnapiesv1.NextPagingAreaScopeRanpagingAttemptInfo {
	return xnapiesv1.NextPagingAreaScopeRanpagingAttemptInfo_NEXT_PAGING_AREA_SCOPE_RANPAGING_ATTEMPT_INFO_SAME
}
func CreateNextPagingAreaScopeRanpagingAttemptInfoChanged() xnapiesv1.NextPagingAreaScopeRanpagingAttemptInfo {
	return xnapiesv1.NextPagingAreaScopeRanpagingAttemptInfo_NEXT_PAGING_AREA_SCOPE_RANPAGING_ATTEMPT_INFO_CHANGED
}
func CreateRanpagingFailureTrue() xnapiesv1.RanpagingFailure {
	return xnapiesv1.RanpagingFailure_RANPAGING_FAILURE_TRUE
}
func CreateRedundantQoSflowIndicatorTrue() xnapiesv1.RedundantQoSflowIndicator {
	return xnapiesv1.RedundantQoSflowIndicator_REDUNDANT_QO_SFLOW_INDICATOR_TRUE
}
func CreateRedundantQoSflowIndicatorFalse() xnapiesv1.RedundantQoSflowIndicator {
	return xnapiesv1.RedundantQoSflowIndicator_REDUNDANT_QO_SFLOW_INDICATOR_FALSE
}
func CreateRsnV1() xnapiesv1.Rsn {
	return xnapiesv1.Rsn_RSN_V1
}
func CreateRsnV2() xnapiesv1.Rsn {
	return xnapiesv1.Rsn_RSN_V2
}
func CreateReflectiveQoSattributeSubjectToReflectiveQoS() xnapiesv1.ReflectiveQoSattribute {
	return xnapiesv1.ReflectiveQoSattribute_REFLECTIVE_QO_SATTRIBUTE_SUBJECT_TO_REFLECTIVE_QO_S
}
func CreateReportAmountMdtR1() xnapiesv1.ReportAmountMdt {
	return xnapiesv1.ReportAmountMdt_REPORT_AMOUNT_MDT_R1
}
func CreateReportAmountMdtR2() xnapiesv1.ReportAmountMdt {
	return xnapiesv1.ReportAmountMdt_REPORT_AMOUNT_MDT_R2
}
func CreateReportAmountMdtR4() xnapiesv1.ReportAmountMdt {
	return xnapiesv1.ReportAmountMdt_REPORT_AMOUNT_MDT_R4
}
func CreateReportAmountMdtR8() xnapiesv1.ReportAmountMdt {
	return xnapiesv1.ReportAmountMdt_REPORT_AMOUNT_MDT_R8
}
func CreateReportAmountMdtR16() xnapiesv1.ReportAmountMdt {
	return xnapiesv1.ReportAmountMdt_REPORT_AMOUNT_MDT_R16
}
func CreateReportAmountMdtR32() xnapiesv1.ReportAmountMdt {
	return xnapiesv1.ReportAmountMdt_REPORT_AMOUNT_MDT_R32
}
func CreateReportAmountMdtR64() xnapiesv1.ReportAmountMdt {
	return xnapiesv1.ReportAmountMdt_REPORT_AMOUNT_MDT_R64
}
func CreateReportAmountMdtInfinity() xnapiesv1.ReportAmountMdt {
	return xnapiesv1.ReportAmountMdt_REPORT_AMOUNT_MDT_INFINITY
}
func CreateReportAreaCell() xnapiesv1.ReportArea {
	return xnapiesv1.ReportArea_REPORT_AREA_CELL
}
func CreateReportIntervalMdtMs120() xnapiesv1.ReportIntervalMdt {
	return xnapiesv1.ReportIntervalMdt_REPORT_INTERVAL_MDT_MS120
}
func CreateReportIntervalMdtMs240() xnapiesv1.ReportIntervalMdt {
	return xnapiesv1.ReportIntervalMdt_REPORT_INTERVAL_MDT_MS240
}
func CreateReportIntervalMdtMs480() xnapiesv1.ReportIntervalMdt {
	return xnapiesv1.ReportIntervalMdt_REPORT_INTERVAL_MDT_MS480
}
func CreateReportIntervalMdtMs640() xnapiesv1.ReportIntervalMdt {
	return xnapiesv1.ReportIntervalMdt_REPORT_INTERVAL_MDT_MS640
}
func CreateReportIntervalMdtMs1024() xnapiesv1.ReportIntervalMdt {
	return xnapiesv1.ReportIntervalMdt_REPORT_INTERVAL_MDT_MS1024
}
func CreateReportIntervalMdtMs2048() xnapiesv1.ReportIntervalMdt {
	return xnapiesv1.ReportIntervalMdt_REPORT_INTERVAL_MDT_MS2048
}
func CreateReportIntervalMdtMs5120() xnapiesv1.ReportIntervalMdt {
	return xnapiesv1.ReportIntervalMdt_REPORT_INTERVAL_MDT_MS5120
}
func CreateReportIntervalMdtMs10240() xnapiesv1.ReportIntervalMdt {
	return xnapiesv1.ReportIntervalMdt_REPORT_INTERVAL_MDT_MS10240
}
func CreateReportIntervalMdtMin1() xnapiesv1.ReportIntervalMdt {
	return xnapiesv1.ReportIntervalMdt_REPORT_INTERVAL_MDT_MIN1
}
func CreateReportIntervalMdtMin6() xnapiesv1.ReportIntervalMdt {
	return xnapiesv1.ReportIntervalMdt_REPORT_INTERVAL_MDT_MIN6
}
func CreateReportIntervalMdtMin12() xnapiesv1.ReportIntervalMdt {
	return xnapiesv1.ReportIntervalMdt_REPORT_INTERVAL_MDT_MIN12
}
func CreateReportIntervalMdtMin30() xnapiesv1.ReportIntervalMdt {
	return xnapiesv1.ReportIntervalMdt_REPORT_INTERVAL_MDT_MIN30
}
func CreateReportIntervalMdtMin60() xnapiesv1.ReportIntervalMdt {
	return xnapiesv1.ReportIntervalMdt_REPORT_INTERVAL_MDT_MIN60
}
func CreateExtendedReportIntervalMdtMs20480() xnapiesv1.ExtendedReportIntervalMdt {
	return xnapiesv1.ExtendedReportIntervalMdt_EXTENDED_REPORT_INTERVAL_MDT_MS20480
}
func CreateExtendedReportIntervalMdtMs40960() xnapiesv1.ExtendedReportIntervalMdt {
	return xnapiesv1.ExtendedReportIntervalMdt_EXTENDED_REPORT_INTERVAL_MDT_MS40960
}
func CreateReportingPeriodicityHalfThousandMs() xnapiesv1.ReportingPeriodicity {
	return xnapiesv1.ReportingPeriodicity_REPORTING_PERIODICITY_HALF_THOUSAND_MS
}
func CreateReportingPeriodicityOneThousandMs() xnapiesv1.ReportingPeriodicity {
	return xnapiesv1.ReportingPeriodicity_REPORTING_PERIODICITY_ONE_THOUSAND_MS
}
func CreateReportingPeriodicityTwoThousandMs() xnapiesv1.ReportingPeriodicity {
	return xnapiesv1.ReportingPeriodicity_REPORTING_PERIODICITY_TWO_THOUSAND_MS
}
func CreateReportingPeriodicityFiveThousandMs() xnapiesv1.ReportingPeriodicity {
	return xnapiesv1.ReportingPeriodicity_REPORTING_PERIODICITY_FIVE_THOUSAND_MS
}
func CreateReportingPeriodicityTenThousandMs() xnapiesv1.ReportingPeriodicity {
	return xnapiesv1.ReportingPeriodicity_REPORTING_PERIODICITY_TEN_THOUSAND_MS
}
func CreateRegistrationRequestStart() xnapiesv1.RegistrationRequest {
	return xnapiesv1.RegistrationRequest_REGISTRATION_REQUEST_START
}
func CreateRegistrationRequestStop() xnapiesv1.RegistrationRequest {
	return xnapiesv1.RegistrationRequest_REGISTRATION_REQUEST_STOP
}
func CreateRegistrationRequestAdd() xnapiesv1.RegistrationRequest {
	return xnapiesv1.RegistrationRequest_REGISTRATION_REQUEST_ADD
}
func CreateSubframeTypeReservedSubframePatternMbsfn() xnapiesv1.SubframeTypeReservedSubframePattern {
	return xnapiesv1.SubframeTypeReservedSubframePattern_SUBFRAME_TYPE_RESERVED_SUBFRAME_PATTERN_MBSFN
}
func CreateSubframeTypeReservedSubframePatternNonMbsfn() xnapiesv1.SubframeTypeReservedSubframePattern {
	return xnapiesv1.SubframeTypeReservedSubframePattern_SUBFRAME_TYPE_RESERVED_SUBFRAME_PATTERN_NON_MBSFN
}
func CreateRlcmodeRlcAm() xnapiesv1.Rlcmode {
	return xnapiesv1.Rlcmode_RLCMODE_RLC_AM
}
func CreateRlcmodeRlcUmBidirectional() xnapiesv1.Rlcmode {
	return xnapiesv1.Rlcmode_RLCMODE_RLC_UM_BIDIRECTIONAL
}
func CreateRlcmodeRlcUmUnidirectionalUl() xnapiesv1.Rlcmode {
	return xnapiesv1.Rlcmode_RLCMODE_RLC_UM_UNIDIRECTIONAL_UL
}
func CreateRlcmodeRlcUmUnidirectionalDl() xnapiesv1.Rlcmode {
	return xnapiesv1.Rlcmode_RLCMODE_RLC_UM_UNIDIRECTIONAL_DL
}
func CreateRLcprimaryIndicatorRlcduplicationInformationTrue() xnapiesv1.RlcprimaryIndicatorRlcduplicationInformation {
	return xnapiesv1.RlcprimaryIndicatorRlcduplicationInformation_R_LCPRIMARY_INDICATOR_RLCDUPLICATION_INFORMATION_TRUE
}
func CreateRLcprimaryIndicatorRlcduplicationInformationFalse() xnapiesv1.RlcprimaryIndicatorRlcduplicationInformation {
	return xnapiesv1.RlcprimaryIndicatorRlcduplicationInformation_R_LCPRIMARY_INDICATOR_RLCDUPLICATION_INFORMATION_FALSE
}
func CreateDuplicationStateRlcduplicationStateItemActive() xnapiesv1.DuplicationStateRlcduplicationStateItem {
	return xnapiesv1.DuplicationStateRlcduplicationStateItem_DUPLICATION_STATE_RLCDUPLICATION_STATE_ITEM_ACTIVE
}
func CreateDuplicationStateRlcduplicationStateItemInactive() xnapiesv1.DuplicationStateRlcduplicationStateItem {
	return xnapiesv1.DuplicationStateRlcduplicationStateItem_DUPLICATION_STATE_RLCDUPLICATION_STATE_ITEM_INACTIVE
}
func CreateReestablishmentIndicationReestablished() xnapiesv1.ReestablishmentIndication {
	return xnapiesv1.ReestablishmentIndication_REESTABLISHMENT_INDICATION_REESTABLISHED
}
func CreateRrcconfigIndicationFullConfig() xnapiesv1.RrcconfigIndication {
	return xnapiesv1.RrcconfigIndication_RRCCONFIG_INDICATION_FULL_CONFIG
}
func CreateRrcconfigIndicationDeltaConfig() xnapiesv1.RrcconfigIndication {
	return xnapiesv1.RrcconfigIndication_RRCCONFIG_INDICATION_DELTA_CONFIG
}
func CreateRrcconnReestabIndicatorReconfigurationFailure() xnapiesv1.RrcconnReestabIndicator {
	return xnapiesv1.RrcconnReestabIndicator_RRCCONN_REESTAB_INDICATOR_RECONFIGURATION_FAILURE
}
func CreateRrcconnReestabIndicatorHandoverFailure() xnapiesv1.RrcconnReestabIndicator {
	return xnapiesv1.RrcconnReestabIndicator_RRCCONN_REESTAB_INDICATOR_HANDOVER_FAILURE
}
func CreateRrcconnReestabIndicatorOtherFailure() xnapiesv1.RrcconnReestabIndicator {
	return xnapiesv1.RrcconnReestabIndicator_RRCCONN_REESTAB_INDICATOR_OTHER_FAILURE
}
func CreateRrcresumeCauseRnaUpdate() xnapiesv1.RrcresumeCause {
	return xnapiesv1.RrcresumeCause_RRCRESUME_CAUSE_RNA_UPDATE
}
func CreateScgconfigurationQueryTrue() xnapiesv1.ScgconfigurationQuery {
	return xnapiesv1.ScgconfigurationQuery_SCGCONFIGURATION_QUERY_TRUE
}
func CreateScgindicatorReleased() xnapiesv1.Scgindicator {
	return xnapiesv1.Scgindicator_SCGINDICATOR_RELEASED
}
func CreateIntegrityProtectionIndicationSecurityIndicationRequired() xnapiesv1.IntegrityProtectionIndicationSecurityIndication {
	return xnapiesv1.IntegrityProtectionIndicationSecurityIndication_INTEGRITY_PROTECTION_INDICATION_SECURITY_INDICATION_REQUIRED
}
func CreateIntegrityProtectionIndicationSecurityIndicationPreferred() xnapiesv1.IntegrityProtectionIndicationSecurityIndication {
	return xnapiesv1.IntegrityProtectionIndicationSecurityIndication_INTEGRITY_PROTECTION_INDICATION_SECURITY_INDICATION_PREFERRED
}
func CreateIntegrityProtectionIndicationSecurityIndicationNotNeeded() xnapiesv1.IntegrityProtectionIndicationSecurityIndication {
	return xnapiesv1.IntegrityProtectionIndicationSecurityIndication_INTEGRITY_PROTECTION_INDICATION_SECURITY_INDICATION_NOT_NEEDED
}
func CreateConfidentialityProtectionIndicationSecurityIndicationRequired() xnapiesv1.ConfidentialityProtectionIndicationSecurityIndication {
	return xnapiesv1.ConfidentialityProtectionIndicationSecurityIndication_CONFIDENTIALITY_PROTECTION_INDICATION_SECURITY_INDICATION_REQUIRED
}
func CreateConfidentialityProtectionIndicationSecurityIndicationPreferred() xnapiesv1.ConfidentialityProtectionIndicationSecurityIndication {
	return xnapiesv1.ConfidentialityProtectionIndicationSecurityIndication_CONFIDENTIALITY_PROTECTION_INDICATION_SECURITY_INDICATION_PREFERRED
}
func CreateConfidentialityProtectionIndicationSecurityIndicationNotNeeded() xnapiesv1.ConfidentialityProtectionIndicationSecurityIndication {
	return xnapiesv1.ConfidentialityProtectionIndicationSecurityIndication_CONFIDENTIALITY_PROTECTION_INDICATION_SECURITY_INDICATION_NOT_NEEDED
}
func CreateIntegrityProtectionResultSecurityResultPerformed() xnapiesv1.IntegrityProtectionResultSecurityResult {
	return xnapiesv1.IntegrityProtectionResultSecurityResult_INTEGRITY_PROTECTION_RESULT_SECURITY_RESULT_PERFORMED
}
func CreateIntegrityProtectionResultSecurityResultNotPerformed() xnapiesv1.IntegrityProtectionResultSecurityResult {
	return xnapiesv1.IntegrityProtectionResultSecurityResult_INTEGRITY_PROTECTION_RESULT_SECURITY_RESULT_NOT_PERFORMED
}
func CreateConfidentialityProtectionResultSecurityResultPerformed() xnapiesv1.ConfidentialityProtectionResultSecurityResult {
	return xnapiesv1.ConfidentialityProtectionResultSecurityResult_CONFIDENTIALITY_PROTECTION_RESULT_SECURITY_RESULT_PERFORMED
}
func CreateConfidentialityProtectionResultSecurityResultNotPerformed() xnapiesv1.ConfidentialityProtectionResultSecurityResult {
	return xnapiesv1.ConfidentialityProtectionResultSecurityResult_CONFIDENTIALITY_PROTECTION_RESULT_SECURITY_RESULT_NOT_PERFORMED
}
func CreateSensorMeasConfigSetup() xnapiesv1.SensorMeasConfig {
	return xnapiesv1.SensorMeasConfig_SENSOR_MEAS_CONFIG_SETUP
}
func CreateUncompensatedBarometricConfigSensorNameTrue() xnapiesv1.UncompensatedBarometricConfigSensorName {
	return xnapiesv1.UncompensatedBarometricConfigSensorName_UNCOMPENSATED_BAROMETRIC_CONFIG_SENSOR_NAME_TRUE
}
func CreateUeSpeedConfigSensorNameTrue() xnapiesv1.UeSpeedConfigSensorName {
	return xnapiesv1.UeSpeedConfigSensorName_UE_SPEED_CONFIG_SENSOR_NAME_TRUE
}
func CreateUeOrientationConfigSensorNameTrue() xnapiesv1.UeOrientationConfigSensorName {
	return xnapiesv1.UeOrientationConfigSensorName_UE_ORIENTATION_CONFIG_SENSOR_NAME_TRUE
}
func CreateFreqBandIndicatorPriorityServedCellInformationEutraNotBroadcast() xnapiesv1.FreqBandIndicatorPriorityServedCellInformationEutra {
	return xnapiesv1.FreqBandIndicatorPriorityServedCellInformationEutra_FREQ_BAND_INDICATOR_PRIORITY_SERVED_CELL_INFORMATION_EUTRA_NOT_BROADCAST
}
func CreateFreqBandIndicatorPriorityServedCellInformationEutraBroadcast() xnapiesv1.FreqBandIndicatorPriorityServedCellInformationEutra {
	return xnapiesv1.FreqBandIndicatorPriorityServedCellInformationEutra_FREQ_BAND_INDICATOR_PRIORITY_SERVED_CELL_INFORMATION_EUTRA_BROADCAST
}
func CreateBandwidthReducedSiservedCellInformationEutraScheduled() xnapiesv1.BandwidthReducedSiservedCellInformationEutra {
	return xnapiesv1.BandwidthReducedSiservedCellInformationEutra_BANDWIDTH_REDUCED_SISERVED_CELL_INFORMATION_EUTRA_SCHEDULED
}
func CreateSubframeAssignmnetServedCellInformationEutratddinfoSa0() xnapiesv1.SubframeAssignmnetServedCellInformationEutratddinfo {
	return xnapiesv1.SubframeAssignmnetServedCellInformationEutratddinfo_SUBFRAME_ASSIGNMNET_SERVED_CELL_INFORMATION_EUTRATDDINFO_SA0
}
func CreateSubframeAssignmnetServedCellInformationEutratddinfoSa1() xnapiesv1.SubframeAssignmnetServedCellInformationEutratddinfo {
	return xnapiesv1.SubframeAssignmnetServedCellInformationEutratddinfo_SUBFRAME_ASSIGNMNET_SERVED_CELL_INFORMATION_EUTRATDDINFO_SA1
}
func CreateSubframeAssignmnetServedCellInformationEutratddinfoSa2() xnapiesv1.SubframeAssignmnetServedCellInformationEutratddinfo {
	return xnapiesv1.SubframeAssignmnetServedCellInformationEutratddinfo_SUBFRAME_ASSIGNMNET_SERVED_CELL_INFORMATION_EUTRATDDINFO_SA2
}
func CreateSubframeAssignmnetServedCellInformationEutratddinfoSa3() xnapiesv1.SubframeAssignmnetServedCellInformationEutratddinfo {
	return xnapiesv1.SubframeAssignmnetServedCellInformationEutratddinfo_SUBFRAME_ASSIGNMNET_SERVED_CELL_INFORMATION_EUTRATDDINFO_SA3
}
func CreateSubframeAssignmnetServedCellInformationEutratddinfoSa4() xnapiesv1.SubframeAssignmnetServedCellInformationEutratddinfo {
	return xnapiesv1.SubframeAssignmnetServedCellInformationEutratddinfo_SUBFRAME_ASSIGNMNET_SERVED_CELL_INFORMATION_EUTRATDDINFO_SA4
}
func CreateSubframeAssignmnetServedCellInformationEutratddinfoSa5() xnapiesv1.SubframeAssignmnetServedCellInformationEutratddinfo {
	return xnapiesv1.SubframeAssignmnetServedCellInformationEutratddinfo_SUBFRAME_ASSIGNMNET_SERVED_CELL_INFORMATION_EUTRATDDINFO_SA5
}
func CreateSubframeAssignmnetServedCellInformationEutratddinfoSa6() xnapiesv1.SubframeAssignmnetServedCellInformationEutratddinfo {
	return xnapiesv1.SubframeAssignmnetServedCellInformationEutratddinfo_SUBFRAME_ASSIGNMNET_SERVED_CELL_INFORMATION_EUTRATDDINFO_SA6
}
func CreateDeactivationindicationServedCellsToModifyEutraitemDeactivated() xnapiesv1.DeactivationindicationServedCellsToModifyEutraitem {
	return xnapiesv1.DeactivationindicationServedCellsToModifyEutraitem_DEACTIVATIONINDICATION_SERVED_CELLS_TO_MODIFY_EUTRAITEM_DEACTIVATED
}
func CreateDeactivationindicationServedCellsToModifyNritemDeactivated() xnapiesv1.DeactivationindicationServedCellsToModifyNritem {
	return xnapiesv1.DeactivationindicationServedCellsToModifyNritem_DEACTIVATIONINDICATION_SERVED_CELLS_TO_MODIFY_NRITEM_DEACTIVATED
}
func CreateSNgRannodeAdditionTriggerIndSnChange() xnapiesv1.SNGRAnnodeAdditionTriggerInd {
	return xnapiesv1.SNGRAnnodeAdditionTriggerInd_S_NG_RANNODE_ADDITION_TRIGGER_IND_SN_CHANGE
}
func CreateSNgRannodeAdditionTriggerIndInterMnHo() xnapiesv1.SNGRAnnodeAdditionTriggerInd {
	return xnapiesv1.SNGRAnnodeAdditionTriggerInd_S_NG_RANNODE_ADDITION_TRIGGER_IND_INTER_MN_HO
}
func CreateSNgRannodeAdditionTriggerIndIntraMnHo() xnapiesv1.SNGRAnnodeAdditionTriggerInd {
	return xnapiesv1.SNGRAnnodeAdditionTriggerInd_S_NG_RANNODE_ADDITION_TRIGGER_IND_INTRA_MN_HO
}
func CreateSntriggeredTrue() xnapiesv1.Sntriggered {
	return xnapiesv1.Sntriggered_SNTRIGGERED_TRUE
}
func CreateSpecialSubframePatternsEUtraSsp0() xnapiesv1.SpecialSubframePatternsEUTra {
	return xnapiesv1.SpecialSubframePatternsEUTra_SPECIAL_SUBFRAME_PATTERNS_E_UTRA_SSP0
}
func CreateSpecialSubframePatternsEUtraSsp1() xnapiesv1.SpecialSubframePatternsEUTra {
	return xnapiesv1.SpecialSubframePatternsEUTra_SPECIAL_SUBFRAME_PATTERNS_E_UTRA_SSP1
}
func CreateSpecialSubframePatternsEUtraSsp2() xnapiesv1.SpecialSubframePatternsEUTra {
	return xnapiesv1.SpecialSubframePatternsEUTra_SPECIAL_SUBFRAME_PATTERNS_E_UTRA_SSP2
}
func CreateSpecialSubframePatternsEUtraSsp3() xnapiesv1.SpecialSubframePatternsEUTra {
	return xnapiesv1.SpecialSubframePatternsEUTra_SPECIAL_SUBFRAME_PATTERNS_E_UTRA_SSP3
}
func CreateSpecialSubframePatternsEUtraSsp4() xnapiesv1.SpecialSubframePatternsEUTra {
	return xnapiesv1.SpecialSubframePatternsEUTra_SPECIAL_SUBFRAME_PATTERNS_E_UTRA_SSP4
}
func CreateSpecialSubframePatternsEUtraSsp5() xnapiesv1.SpecialSubframePatternsEUTra {
	return xnapiesv1.SpecialSubframePatternsEUTra_SPECIAL_SUBFRAME_PATTERNS_E_UTRA_SSP5
}
func CreateSpecialSubframePatternsEUtraSsp6() xnapiesv1.SpecialSubframePatternsEUTra {
	return xnapiesv1.SpecialSubframePatternsEUTra_SPECIAL_SUBFRAME_PATTERNS_E_UTRA_SSP6
}
func CreateSpecialSubframePatternsEUtraSsp7() xnapiesv1.SpecialSubframePatternsEUTra {
	return xnapiesv1.SpecialSubframePatternsEUTra_SPECIAL_SUBFRAME_PATTERNS_E_UTRA_SSP7
}
func CreateSpecialSubframePatternsEUtraSsp8() xnapiesv1.SpecialSubframePatternsEUTra {
	return xnapiesv1.SpecialSubframePatternsEUTra_SPECIAL_SUBFRAME_PATTERNS_E_UTRA_SSP8
}
func CreateSpecialSubframePatternsEUtraSsp9() xnapiesv1.SpecialSubframePatternsEUTra {
	return xnapiesv1.SpecialSubframePatternsEUTra_SPECIAL_SUBFRAME_PATTERNS_E_UTRA_SSP9
}
func CreateSpecialSubframePatternsEUtraSsp10() xnapiesv1.SpecialSubframePatternsEUTra {
	return xnapiesv1.SpecialSubframePatternsEUTra_SPECIAL_SUBFRAME_PATTERNS_E_UTRA_SSP10
}
func CreateSplitSessionIndicatorSplit() xnapiesv1.SplitSessionIndicator {
	return xnapiesv1.SplitSessionIndicator_SPLIT_SESSION_INDICATOR_SPLIT
}
func CreateSplitSrbsTypesSrb1() xnapiesv1.SplitSrbsTypes {
	return xnapiesv1.SplitSrbsTypes_SPLIT_SRBS_TYPES_SRB1
}
func CreateSplitSrbsTypesSrb2() xnapiesv1.SplitSrbsTypes {
	return xnapiesv1.SplitSrbsTypes_SPLIT_SRBS_TYPES_SRB2
}
func CreateSplitSrbsTypesSrb1and2() xnapiesv1.SplitSrbsTypes {
	return xnapiesv1.SplitSrbsTypes_SPLIT_SRBS_TYPES_SRB1AND2
}
func CreateTimeToTriggerMs0() xnapiesv1.TimeToTrigger {
	return xnapiesv1.TimeToTrigger_TIME_TO_TRIGGER_MS0
}
func CreateTimeToTriggerMs40() xnapiesv1.TimeToTrigger {
	return xnapiesv1.TimeToTrigger_TIME_TO_TRIGGER_MS40
}
func CreateTimeToTriggerMs64() xnapiesv1.TimeToTrigger {
	return xnapiesv1.TimeToTrigger_TIME_TO_TRIGGER_MS64
}
func CreateTimeToTriggerMs80() xnapiesv1.TimeToTrigger {
	return xnapiesv1.TimeToTrigger_TIME_TO_TRIGGER_MS80
}
func CreateTimeToTriggerMs100() xnapiesv1.TimeToTrigger {
	return xnapiesv1.TimeToTrigger_TIME_TO_TRIGGER_MS100
}
func CreateTimeToTriggerMs128() xnapiesv1.TimeToTrigger {
	return xnapiesv1.TimeToTrigger_TIME_TO_TRIGGER_MS128
}
func CreateTimeToTriggerMs160() xnapiesv1.TimeToTrigger {
	return xnapiesv1.TimeToTrigger_TIME_TO_TRIGGER_MS160
}
func CreateTimeToTriggerMs256() xnapiesv1.TimeToTrigger {
	return xnapiesv1.TimeToTrigger_TIME_TO_TRIGGER_MS256
}
func CreateTimeToTriggerMs320() xnapiesv1.TimeToTrigger {
	return xnapiesv1.TimeToTrigger_TIME_TO_TRIGGER_MS320
}
func CreateTimeToTriggerMs480() xnapiesv1.TimeToTrigger {
	return xnapiesv1.TimeToTrigger_TIME_TO_TRIGGER_MS480
}
func CreateTimeToTriggerMs512() xnapiesv1.TimeToTrigger {
	return xnapiesv1.TimeToTrigger_TIME_TO_TRIGGER_MS512
}
func CreateTimeToTriggerMs640() xnapiesv1.TimeToTrigger {
	return xnapiesv1.TimeToTrigger_TIME_TO_TRIGGER_MS640
}
func CreateTimeToTriggerMs1024() xnapiesv1.TimeToTrigger {
	return xnapiesv1.TimeToTrigger_TIME_TO_TRIGGER_MS1024
}
func CreateTimeToTriggerMs1280() xnapiesv1.TimeToTrigger {
	return xnapiesv1.TimeToTrigger_TIME_TO_TRIGGER_MS1280
}
func CreateTimeToTriggerMs2560() xnapiesv1.TimeToTrigger {
	return xnapiesv1.TimeToTrigger_TIME_TO_TRIGGER_MS2560
}
func CreateTimeToTriggerMs5120() xnapiesv1.TimeToTrigger {
	return xnapiesv1.TimeToTrigger_TIME_TO_TRIGGER_MS5120
}
func CreateTimeToWaitV1s() xnapiesv1.TimeToWait {
	return xnapiesv1.TimeToWait_TIME_TO_WAIT_V1S
}
func CreateTimeToWaitV2s() xnapiesv1.TimeToWait {
	return xnapiesv1.TimeToWait_TIME_TO_WAIT_V2S
}
func CreateTimeToWaitV5s() xnapiesv1.TimeToWait {
	return xnapiesv1.TimeToWait_TIME_TO_WAIT_V5S
}
func CreateTimeToWaitV10s() xnapiesv1.TimeToWait {
	return xnapiesv1.TimeToWait_TIME_TO_WAIT_V10S
}
func CreateTimeToWaitV20s() xnapiesv1.TimeToWait {
	return xnapiesv1.TimeToWait_TIME_TO_WAIT_V20S
}
func CreateTimeToWaitV60s() xnapiesv1.TimeToWait {
	return xnapiesv1.TimeToWait_TIME_TO_WAIT_V60S
}
func CreateTnlassociationUsageUe() xnapiesv1.TnlassociationUsage {
	return xnapiesv1.TnlassociationUsage_TNLASSOCIATION_USAGE_UE
}
func CreateTnlassociationUsageNonUe() xnapiesv1.TnlassociationUsage {
	return xnapiesv1.TnlassociationUsage_TNLASSOCIATION_USAGE_NON_UE
}
func CreateTnlassociationUsageBoth() xnapiesv1.TnlassociationUsage {
	return xnapiesv1.TnlassociationUsage_TNLASSOCIATION_USAGE_BOTH
}
func CreateTraceDepthMinimum() xnapiesv1.TraceDepth {
	return xnapiesv1.TraceDepth_TRACE_DEPTH_MINIMUM
}
func CreateTraceDepthMedium() xnapiesv1.TraceDepth {
	return xnapiesv1.TraceDepth_TRACE_DEPTH_MEDIUM
}
func CreateTraceDepthMaximum() xnapiesv1.TraceDepth {
	return xnapiesv1.TraceDepth_TRACE_DEPTH_MAXIMUM
}
func CreateTraceDepthMinimumWithoutVendorSpecificExtension() xnapiesv1.TraceDepth {
	return xnapiesv1.TraceDepth_TRACE_DEPTH_MINIMUM_WITHOUT_VENDOR_SPECIFIC_EXTENSION
}
func CreateTraceDepthMediumWithoutVendorSpecificExtension() xnapiesv1.TraceDepth {
	return xnapiesv1.TraceDepth_TRACE_DEPTH_MEDIUM_WITHOUT_VENDOR_SPECIFIC_EXTENSION
}
func CreateTraceDepthMaximumWithoutVendorSpecificExtension() xnapiesv1.TraceDepth {
	return xnapiesv1.TraceDepth_TRACE_DEPTH_MAXIMUM_WITHOUT_VENDOR_SPECIFIC_EXTENSION
}
func CreateTypeOfErrorNotUnderstood() xnapiesv1.TypeOfError {
	return xnapiesv1.TypeOfError_TYPE_OF_ERROR_NOT_UNDERSTOOD
}
func CreateTypeOfErrorMissing() xnapiesv1.TypeOfError {
	return xnapiesv1.TypeOfError_TYPE_OF_ERROR_MISSING
}
func CreateUecontextKeptIndicatorTrue() xnapiesv1.UecontextKeptIndicator {
	return xnapiesv1.UecontextKeptIndicator_UECONTEXT_KEPT_INDICATOR_TRUE
}
func CreateUespecificDrxV32() xnapiesv1.UespecificDrx {
	return xnapiesv1.UespecificDrx_UESPECIFIC_DRX_V32
}
func CreateUespecificDrxV64() xnapiesv1.UespecificDrx {
	return xnapiesv1.UespecificDrx_UESPECIFIC_DRX_V64
}
func CreateUespecificDrxV128() xnapiesv1.UespecificDrx {
	return xnapiesv1.UespecificDrx_UESPECIFIC_DRX_V128
}
func CreateUespecificDrxV256() xnapiesv1.UespecificDrx {
	return xnapiesv1.UespecificDrx_UESPECIFIC_DRX_V256
}
func CreateUlUeConfigurationNoData() xnapiesv1.UlUEConfiguration {
	return xnapiesv1.UlUEConfiguration_UL_UE_CONFIGURATION_NO_DATA
}
func CreateUlUeConfigurationShared() xnapiesv1.UlUEConfiguration {
	return xnapiesv1.UlUEConfiguration_UL_UE_CONFIGURATION_SHARED
}
func CreateUlUeConfigurationOnly() xnapiesv1.UlUEConfiguration {
	return xnapiesv1.UlUEConfiguration_UL_UE_CONFIGURATION_ONLY
}
func CreateUlforwardingUlForwardingProposed() xnapiesv1.Ulforwarding {
	return xnapiesv1.Ulforwarding_ULFORWARDING_UL_FORWARDING_PROPOSED
}
func CreateUlforwardingProposalUlForwardingProposed() xnapiesv1.UlforwardingProposal {
	return xnapiesv1.UlforwardingProposal_ULFORWARDING_PROPOSAL_UL_FORWARDING_PROPOSED
}
func CreateUserPlaneTrafficActivityReportInactive() xnapiesv1.UserPlaneTrafficActivityReport {
	return xnapiesv1.UserPlaneTrafficActivityReport_USER_PLANE_TRAFFIC_ACTIVITY_REPORT_INACTIVE
}
func CreateUserPlaneTrafficActivityReportReActivated() xnapiesv1.UserPlaneTrafficActivityReport {
	return xnapiesv1.UserPlaneTrafficActivityReport_USER_PLANE_TRAFFIC_ACTIVITY_REPORT_RE_ACTIVATED
}
func CreateVehicleUeAuthorized() xnapiesv1.VehicleUe {
	return xnapiesv1.VehicleUe_VEHICLE_UE_AUTHORIZED
}
func CreateVehicleUeNotAuthorized() xnapiesv1.VehicleUe {
	return xnapiesv1.VehicleUe_VEHICLE_UE_NOT_AUTHORIZED
}
func CreateWlanrssiWlanmeasurementConfigurationTrue() xnapiesv1.WlanrssiWlanmeasurementConfiguration {
	return xnapiesv1.WlanrssiWlanmeasurementConfiguration_WLANRSSI_WLANMEASUREMENT_CONFIGURATION_TRUE
}
func CreateWlanrttWlanmeasurementConfigurationTrue() xnapiesv1.WlanrttWlanmeasurementConfiguration {
	return xnapiesv1.WlanrttWlanmeasurementConfiguration_WLANRTT_WLANMEASUREMENT_CONFIGURATION_TRUE
}
func CreateWlanmeasConfigSetup() xnapiesv1.WlanmeasConfig {
	return xnapiesv1.WlanmeasConfig_WLANMEAS_CONFIG_SETUP
}
func CreateCriticalityReject() xnapcommondatatypesv1.Criticality {
	return xnapcommondatatypesv1.Criticality_CRITICALITY_REJECT
}
func CreateCriticalityIgnore() xnapcommondatatypesv1.Criticality {
	return xnapcommondatatypesv1.Criticality_CRITICALITY_IGNORE
}
func CreateCriticalityNotify() xnapcommondatatypesv1.Criticality {
	return xnapcommondatatypesv1.Criticality_CRITICALITY_NOTIFY
}
func CreatePresenceOptional() xnapcommondatatypesv1.Presence {
	return xnapcommondatatypesv1.Presence_PRESENCE_OPTIONAL
}
func CreatePresenceConditional() xnapcommondatatypesv1.Presence {
	return xnapcommondatatypesv1.Presence_PRESENCE_CONDITIONAL
}
func CreatePresenceMandatory() xnapcommondatatypesv1.Presence {
	return xnapcommondatatypesv1.Presence_PRESENCE_MANDATORY
}
func CreateTriggeringMessageInitiatingMessage() xnapcommondatatypesv1.TriggeringMessage {
	return xnapcommondatatypesv1.TriggeringMessage_TRIGGERING_MESSAGE_INITIATING_MESSAGE
}
func CreateTriggeringMessageSuccessfulOutcome() xnapcommondatatypesv1.TriggeringMessage {
	return xnapcommondatatypesv1.TriggeringMessage_TRIGGERING_MESSAGE_SUCCESSFUL_OUTCOME
}
func CreateTriggeringMessageUnsuccessfulOutcome() xnapcommondatatypesv1.TriggeringMessage {
	return xnapcommondatatypesv1.TriggeringMessage_TRIGGERING_MESSAGE_UNSUCCESSFUL_OUTCOME
}
func CreateRequestedFastMcgrecoveryViaSrb3True() xnappducontentsv1.RequestedFastMcgrecoveryViaSrb3 {
	return xnappducontentsv1.RequestedFastMcgrecoveryViaSrb3_REQUESTED_FAST_MCGRECOVERY_VIA_SRB3_TRUE
}
func CreateAvailableFastMcgrecoveryViaSrb3True() xnappducontentsv1.AvailableFastMcgrecoveryViaSrb3 {
	return xnappducontentsv1.AvailableFastMcgrecoveryViaSrb3_AVAILABLE_FAST_MCGRECOVERY_VIA_SRB3_TRUE
}
func CreateRequestedFastMcgrecoveryViaSrb3ReleaseTrue() xnappducontentsv1.RequestedFastMcgrecoveryViaSrb3Release {
	return xnappducontentsv1.RequestedFastMcgrecoveryViaSrb3Release_REQUESTED_FAST_MCGRECOVERY_VIA_SRB3_RELEASE_TRUE
}
func CreateReleaseFastMcgrecoveryViaSrb3True() xnappducontentsv1.ReleaseFastMcgrecoveryViaSrb3 {
	return xnappducontentsv1.ReleaseFastMcgrecoveryViaSrb3_RELEASE_FAST_MCGRECOVERY_VIA_SRB3_TRUE
}
func CreateSrbTypeSplitSrbrrctransferSrb1() xnappducontentsv1.SrbTypeSplitSrbrrctransfer {
	return xnappducontentsv1.SrbTypeSplitSrbrrctransfer_SRB_TYPE_SPLIT_SRBRRCTRANSFER_SRB1
}
func CreateSrbTypeSplitSrbrrctransferSrb2() xnappducontentsv1.SrbTypeSplitSrbrrctransfer {
	return xnappducontentsv1.SrbTypeSplitSrbrrctransfer_SRB_TYPE_SPLIT_SRBRRCTRANSFER_SRB2
}
