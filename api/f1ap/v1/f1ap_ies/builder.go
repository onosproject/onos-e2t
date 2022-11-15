// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package f1apiesv1

import (
	f1apcommondatatypesv1 "github.com/onosproject/onos-e2t/api/f1ap/v1/f1ap_commondatatypes"

	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
)

func (m *AccessPointPosition) SetIEExtensions(iEExtensions []*AccessPointPositionExtIes) *AccessPointPosition {
	m.IEExtensions = iEExtensions
	return m
}

func (m *ActivatedCellstobeUpdatedListItem) SetIEExtensions(iEExtensions []*ActivatedCellstobeUpdatedListItemExtIes) *ActivatedCellstobeUpdatedListItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *ActiveUlbwp) SetShift7Dot5KHz(shift7Dot5KHz Shift7Dot5KHzActiveUlbwp) *ActiveUlbwp {
	m.Shift7Dot5KHz = &shift7Dot5KHz
	return m
}

func (m *ActiveUlbwp) SetIEExtensions(iEExtensions []*ActiveUlbwpExtIes) *ActiveUlbwp {
	m.IEExtensions = iEExtensions
	return m
}

func (m *AdditionalPathItem) SetPathQuality(pathQuality *TrpmeasurementQuality) *AdditionalPathItem {
	m.PathQuality = pathQuality
	return m
}

func (m *AdditionalPathItem) SetIEExtensions(iEExtensions []*AdditionalPathItemExtIes) *AdditionalPathItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *AdditionalPdcpduplicationTnlItem) SetIEExtensions(iEExtensions []*AdditionalPdcpduplicationTnlItemExtIes) *AdditionalPdcpduplicationTnlItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *AdditionalSibmessageListItem) SetIEExtensions(iEExtensions []*AdditionalSibmessageListItemExtIes) *AdditionalSibmessageListItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *AggressorCellListItem) SetIEExtensions(iEExtensions []*AggressorCellListItemExtIes) *AggressorCellListItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *AggressorgNbsetId) SetIEExtensions(iEExtensions []*AggressorgNbsetIdExtIes) *AggressorgNbsetId {
	m.IEExtensions = iEExtensions
	return m
}

func (m *AllocationAndRetentionPriority) SetIEExtensions(iEExtensions []*AllocationAndRetentionPriorityExtIes) *AllocationAndRetentionPriority {
	m.IEExtensions = iEExtensions
	return m
}

func (m *AlternativeQoSparaSetItem) SetGuaranteedFlowBitRateDl(guaranteedFlowBitRateDl *BitRate) *AlternativeQoSparaSetItem {
	m.GuaranteedFlowBitRateDl = guaranteedFlowBitRateDl
	return m
}

func (m *AlternativeQoSparaSetItem) SetGuaranteedFlowBitRateUl(guaranteedFlowBitRateUl *BitRate) *AlternativeQoSparaSetItem {
	m.GuaranteedFlowBitRateUl = guaranteedFlowBitRateUl
	return m
}

func (m *AlternativeQoSparaSetItem) SetPacketDelayBudget(packetDelayBudget *PacketDelayBudget) *AlternativeQoSparaSetItem {
	m.PacketDelayBudget = packetDelayBudget
	return m
}

func (m *AlternativeQoSparaSetItem) SetPacketErrorRate(packetErrorRate *PacketErrorRate) *AlternativeQoSparaSetItem {
	m.PacketErrorRate = packetErrorRate
	return m
}

func (m *AlternativeQoSparaSetItem) SetIEExtensions(iEExtensions []*AlternativeQoSparaSetItemExtIes) *AlternativeQoSparaSetItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *AngleMeasurementQuality) SetZenithQuality(zenithQuality int32) *AngleMeasurementQuality {
	m.ZenithQuality = &zenithQuality
	return m
}

func (m *AngleMeasurementQuality) SetIEExtensions(iEExtensions []*AngleMeasurementQualityExtIes) *AngleMeasurementQuality {
	m.IEExtensions = iEExtensions
	return m
}

func (m *AssociatedSCellItem) SetIEExtensions(iEExtensions []*AssociatedSCellItemExtIes) *AssociatedSCellItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *AvailablePlmnlistItem) SetIEExtensions(iEExtensions []*AvailablePlmnlistItemExtIes) *AvailablePlmnlistItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *AvailableSnpnIDListItem) SetIEExtensions(iEExtensions []*AvailableSnpnIDListItemExtIes) *AvailableSnpnIDListItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *BaplayerBhrlcchannelMappingInfo) SetBAplayerBhrlcchannelMappingInfoToAdd(bAplayerBhrlcchannelMappingInfoToAdd *BaplayerBhrlcchannelMappingInfoList) *BaplayerBhrlcchannelMappingInfo {
	m.BAplayerBhrlcchannelMappingInfoToAdd = bAplayerBhrlcchannelMappingInfoToAdd
	return m
}

func (m *BaplayerBhrlcchannelMappingInfo) SetBAplayerBhrlcchannelMappingInfoToRemove(bAplayerBhrlcchannelMappingInfoToRemove *MappingInformationtoRemove) *BaplayerBhrlcchannelMappingInfo {
	m.BAplayerBhrlcchannelMappingInfoToRemove = bAplayerBhrlcchannelMappingInfoToRemove
	return m
}

func (m *BaplayerBhrlcchannelMappingInfo) SetIEExtensions(iEExtensions []*BaplayerBhrlcchannelMappingInfoExtIes) *BaplayerBhrlcchannelMappingInfo {
	m.IEExtensions = iEExtensions
	return m
}

func (m *BaplayerBhrlcchannelMappingInfoItem) SetPriorHopBapaddress(priorHopBapaddress *Bapaddress) *BaplayerBhrlcchannelMappingInfoItem {
	m.PriorHopBapaddress = priorHopBapaddress
	return m
}

func (m *BaplayerBhrlcchannelMappingInfoItem) SetIngressbHrlcchannelID(ingressbHrlcchannelID *BhrlcchannelId) *BaplayerBhrlcchannelMappingInfoItem {
	m.IngressbHrlcchannelId = ingressbHrlcchannelID
	return m
}

func (m *BaplayerBhrlcchannelMappingInfoItem) SetNextHopBapaddress(nextHopBapaddress *Bapaddress) *BaplayerBhrlcchannelMappingInfoItem {
	m.NextHopBapaddress = nextHopBapaddress
	return m
}

func (m *BaplayerBhrlcchannelMappingInfoItem) SetEgressbHrlcchannelID(egressbHrlcchannelID *BhrlcchannelId) *BaplayerBhrlcchannelMappingInfoItem {
	m.EgressbHrlcchannelId = egressbHrlcchannelID
	return m
}

func (m *BaplayerBhrlcchannelMappingInfoItem) SetIEExtensions(iEExtensions []*BaplayerBhrlcchannelMappingInfoItemExtIes) *BaplayerBhrlcchannelMappingInfoItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *BaproutingId) SetIEExtensions(iEExtensions []*BaproutingIdextIes) *BaproutingId {
	m.IEExtensions = iEExtensions
	return m
}

func (m *BhchannelsFailedToBeModifiedItem) SetCause(cause *Cause) *BhchannelsFailedToBeModifiedItem {
	m.Cause = cause
	return m
}

func (m *BhchannelsFailedToBeModifiedItem) SetIEExtensions(iEExtensions []*BhchannelsFailedToBeModifiedItemExtIes) *BhchannelsFailedToBeModifiedItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *BhchannelsFailedToBeSetupItem) SetCause(cause *Cause) *BhchannelsFailedToBeSetupItem {
	m.Cause = cause
	return m
}

func (m *BhchannelsFailedToBeSetupItem) SetIEExtensions(iEExtensions []*BhchannelsFailedToBeSetupItemExtIes) *BhchannelsFailedToBeSetupItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *BhchannelsFailedToBeSetupModItem) SetCause(cause *Cause) *BhchannelsFailedToBeSetupModItem {
	m.Cause = cause
	return m
}

func (m *BhchannelsFailedToBeSetupModItem) SetIEExtensions(iEExtensions []*BhchannelsFailedToBeSetupModItemExtIes) *BhchannelsFailedToBeSetupModItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *BhchannelsModifiedItem) SetIEExtensions(iEExtensions []*BhchannelsModifiedItemExtIes) *BhchannelsModifiedItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *BhchannelsRequiredToBeReleasedItem) SetIEExtensions(iEExtensions []*BhchannelsRequiredToBeReleasedItemExtIes) *BhchannelsRequiredToBeReleasedItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *BhchannelsSetupItem) SetIEExtensions(iEExtensions []*BhchannelsSetupItemExtIes) *BhchannelsSetupItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *BhchannelsSetupModItem) SetIEExtensions(iEExtensions []*BhchannelsSetupModItemExtIes) *BhchannelsSetupModItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *BhchannelsToBeModifiedItem) SetRLcmode(rLcmode *Rlcmode) *BhchannelsToBeModifiedItem {
	m.RLcmode = rLcmode
	return m
}

func (m *BhchannelsToBeModifiedItem) SetBApctrlPduchannel(bApctrlPduchannel *BapctrlPduchannel) *BhchannelsToBeModifiedItem {
	m.BApctrlPduchannel = bApctrlPduchannel
	return m
}

func (m *BhchannelsToBeModifiedItem) SetTrafficMappingInfo(trafficMappingInfo *TrafficMappingInfo) *BhchannelsToBeModifiedItem {
	m.TrafficMappingInfo = trafficMappingInfo
	return m
}

func (m *BhchannelsToBeModifiedItem) SetIEExtensions(iEExtensions []*BhchannelsToBeModifiedItemExtIes) *BhchannelsToBeModifiedItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *BhchannelsToBeReleasedItem) SetIEExtensions(iEExtensions []*BhchannelsToBeReleasedItemExtIes) *BhchannelsToBeReleasedItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *BhchannelsToBeSetupItem) SetBApctrlPduchannel(bApctrlPduchannel *BapctrlPduchannel) *BhchannelsToBeSetupItem {
	m.BApctrlPduchannel = bApctrlPduchannel
	return m
}

func (m *BhchannelsToBeSetupItem) SetTrafficMappingInfo(trafficMappingInfo *TrafficMappingInfo) *BhchannelsToBeSetupItem {
	m.TrafficMappingInfo = trafficMappingInfo
	return m
}

func (m *BhchannelsToBeSetupItem) SetIEExtensions(iEExtensions []*BhchannelsToBeSetupItemExtIes) *BhchannelsToBeSetupItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *BhchannelsToBeSetupModItem) SetBApctrlPduchannel(bApctrlPduchannel *BapctrlPduchannel) *BhchannelsToBeSetupModItem {
	m.BApctrlPduchannel = bApctrlPduchannel
	return m
}

func (m *BhchannelsToBeSetupModItem) SetTrafficMappingInfo(trafficMappingInfo *TrafficMappingInfo) *BhchannelsToBeSetupModItem {
	m.TrafficMappingInfo = trafficMappingInfo
	return m
}

func (m *BhchannelsToBeSetupModItem) SetIEExtensions(iEExtensions []*BhchannelsToBeSetupModItemExtIes) *BhchannelsToBeSetupModItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *Bhinfo) SetBAproutingID(bAproutingID *BaproutingId) *Bhinfo {
	m.BAproutingId = bAproutingID
	return m
}

func (m *Bhinfo) SetEgressBhrlcchlist(egressBhrlcchlist *EgressBhrlcchlist) *Bhinfo {
	m.EgressBhrlcchlist = egressBhrlcchlist
	return m
}

func (m *Bhinfo) SetIEExtensions(iEExtensions []*BhinfoExtIes) *Bhinfo {
	m.IEExtensions = iEExtensions
	return m
}

func (m *BhRoutingInformationAddedListItem) SetIEExtensions(iEExtensions []*BhRoutingInformationAddedListItemExtIes) *BhRoutingInformationAddedListItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *BhRoutingInformationRemovedListItem) SetIEExtensions(iEExtensions []*BhRoutingInformationRemovedListItemExtIes) *BhRoutingInformationRemovedListItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *BplmnIDInfoItem) SetExtendedPlmnIdentityList(extendedPlmnIdentityList *ExtendedAvailablePlmnList) *BplmnIDInfoItem {
	m.ExtendedPlmnIdentityList = extendedPlmnIdentityList
	return m
}

func (m *BplmnIDInfoItem) SetFiveGsTac(fiveGsTac *FiveGsTAc) *BplmnIDInfoItem {
	m.FiveGsTac = fiveGsTac
	return m
}

func (m *BplmnIDInfoItem) SetRanac(ranac *Ranac) *BplmnIDInfoItem {
	m.Ranac = ranac
	return m
}

func (m *BplmnIDInfoItem) SetIEExtensions(iEExtensions []*BplmnIDInfoItemExtIes) *BplmnIDInfoItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *ServedPlmnsItem) SetIEExtensions(iEExtensions []*ServedPlmnsItemExtIes) *ServedPlmnsItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *BroadcastSnpnIDListItem) SetIEExtensions(iEExtensions []*BroadcastSnpnIDListItemExtIes) *BroadcastSnpnIDListItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *BroadcastPniNPnIDListItem) SetIEExtensions(iEExtensions []*BroadcastPniNPnIDListItemExtIes) *BroadcastPniNPnIDListItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *CandidateSpCellItem) SetIEExtensions(iEExtensions []*CandidateSpCellItemExtIes) *CandidateSpCellItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *CapacityValue) SetSSbareaCapacityValueList(sSbareaCapacityValueList *SsbareaCapacityValueList) *CapacityValue {
	m.SSbareaCapacityValueList = sSbareaCapacityValueList
	return m
}

func (m *CapacityValue) SetIEExtensions(iEExtensions []*CapacityValueExtIes) *CapacityValue {
	m.IEExtensions = iEExtensions
	return m
}

func (m *CellMeasurementResultItem) SetRadioResourceStatus(radioResourceStatus *RadioResourceStatus) *CellMeasurementResultItem {
	m.RadioResourceStatus = radioResourceStatus
	return m
}

func (m *CellMeasurementResultItem) SetCompositeAvailableCapacityGroup(compositeAvailableCapacityGroup *CompositeAvailableCapacityGroup) *CellMeasurementResultItem {
	m.CompositeAvailableCapacityGroup = compositeAvailableCapacityGroup
	return m
}

func (m *CellMeasurementResultItem) SetSliceAvailableCapacity(sliceAvailableCapacity *SliceAvailableCapacity) *CellMeasurementResultItem {
	m.SliceAvailableCapacity = sliceAvailableCapacity
	return m
}

func (m *CellMeasurementResultItem) SetNumberofActiveUes(numberofActiveUes *NumberofActiveUes) *CellMeasurementResultItem {
	m.NumberofActiveUes = numberofActiveUes
	return m
}

func (m *CellMeasurementResultItem) SetIEExtensions(iEExtensions []*CellMeasurementResultItemExtIes) *CellMeasurementResultItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *CellsFailedtobeActivatedListItem) SetIEExtensions(iEExtensions []*CellsFailedtobeActivatedListItemExtIes) *CellsFailedtobeActivatedListItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *CellsStatusItem) SetIEExtensions(iEExtensions []*CellsStatusItemExtIes) *CellsStatusItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *CellsToBeBroadcastItem) SetIEExtensions(iEExtensions []*CellsToBeBroadcastItemExtIes) *CellsToBeBroadcastItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *CellsBroadcastCompletedItem) SetIEExtensions(iEExtensions []*CellsBroadcastCompletedItemExtIes) *CellsBroadcastCompletedItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *BroadcastToBeCancelledItem) SetIEExtensions(iEExtensions []*BroadcastToBeCancelledItemExtIes) *BroadcastToBeCancelledItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *CellsBroadcastCancelledItem) SetIEExtensions(iEExtensions []*CellsBroadcastCancelledItemExtIes) *CellsBroadcastCancelledItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *CellstobeActivatedListItem) SetNRpci(nRpci *Nrpci) *CellstobeActivatedListItem {
	m.NRpci = nRpci
	return m
}

func (m *CellstobeActivatedListItem) SetIEExtensions(iEExtensions []*CellstobeActivatedListItemExtIes) *CellstobeActivatedListItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *CellstobeDeactivatedListItem) SetIEExtensions(iEExtensions []*CellstobeDeactivatedListItemExtIes) *CellstobeDeactivatedListItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *CellstobeBarredItem) SetIEExtensions(iEExtensions []*CellstobeBarredItemExtIes) *CellstobeBarredItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *CellToReportItem) SetSSbtoReportList(sSbtoReportList *SsbtoReportList) *CellToReportItem {
	m.SSbtoReportList = sSbtoReportList
	return m
}

func (m *CellToReportItem) SetSliceToReportList(sliceToReportList *SliceToReportList) *CellToReportItem {
	m.SliceToReportList = sliceToReportList
	return m
}

func (m *CellToReportItem) SetIEExtensions(iEExtensions []*CellToReportItemExtIes) *CellToReportItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *CellType) SetIEExtensions(iEExtensions []*CellTypeExtIes) *CellType {
	m.IEExtensions = iEExtensions
	return m
}

func (m *ChildNodeCellsListItem) SetIAbDuCellResourceConfigurationModeInfo(iAbDuCellResourceConfigurationModeInfo *IabDUCellResourceConfigurationModeInfo) *ChildNodeCellsListItem {
	m.IAbDuCellResourceConfigurationModeInfo = iAbDuCellResourceConfigurationModeInfo
	return m
}

func (m *ChildNodeCellsListItem) SetIAbStcInfo(iAbStcInfo *IabSTcInfo) *ChildNodeCellsListItem {
	m.IAbStcInfo = iAbStcInfo
	return m
}

func (m *ChildNodeCellsListItem) SetRAchConfigCommon(rAchConfigCommon *RachConfigCommon) *ChildNodeCellsListItem {
	m.RAchConfigCommon = rAchConfigCommon
	return m
}

func (m *ChildNodeCellsListItem) SetRAchConfigCommonIab(rAchConfigCommonIab *RachConfigCommonIAb) *ChildNodeCellsListItem {
	m.RAchConfigCommonIab = rAchConfigCommonIab
	return m
}

func (m *ChildNodeCellsListItem) SetCSiRsConfiguration(cSiRsConfiguration []byte) *ChildNodeCellsListItem {
	m.CSiRsConfiguration = cSiRsConfiguration
	return m
}

func (m *ChildNodeCellsListItem) SetSRConfiguration(sRConfiguration []byte) *ChildNodeCellsListItem {
	m.SRConfiguration = sRConfiguration
	return m
}

func (m *ChildNodeCellsListItem) SetPDcchConfigSib1(pDcchConfigSib1 []byte) *ChildNodeCellsListItem {
	m.PDcchConfigSib1 = pDcchConfigSib1
	return m
}

func (m *ChildNodeCellsListItem) SetSCsCommon(sCsCommon []byte) *ChildNodeCellsListItem {
	m.SCsCommon = sCsCommon
	return m
}

func (m *ChildNodeCellsListItem) SetMultiplexingInfo(multiplexingInfo *MultiplexingInfo) *ChildNodeCellsListItem {
	m.MultiplexingInfo = multiplexingInfo
	return m
}

func (m *ChildNodeCellsListItem) SetIEExtensions(iEExtensions []*ChildNodeCellsListItemExtIes) *ChildNodeCellsListItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *ChildNodesListItem) SetChildNodeCellsList(childNodeCellsList *ChildNodeCellsList) *ChildNodesListItem {
	m.ChildNodeCellsList = childNodeCellsList
	return m
}

func (m *ChildNodesListItem) SetIEExtensions(iEExtensions []*ChildNodesListItemExtIes) *ChildNodesListItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *CompositeAvailableCapacityGroup) SetIEExtensions(iEExtensions []*CompositeAvailableCapacityGroupExtIes) *CompositeAvailableCapacityGroup {
	m.IEExtensions = iEExtensions
	return m
}

func (m *CompositeAvailableCapacity) SetCellCapacityClassValue(cellCapacityClassValue *CellCapacityClassValue) *CompositeAvailableCapacity {
	m.CellCapacityClassValue = cellCapacityClassValue
	return m
}

func (m *CompositeAvailableCapacity) SetIEExtensions(iEExtensions []*CompositeAvailableCapacityExtIes) *CompositeAvailableCapacity {
	m.IEExtensions = iEExtensions
	return m
}

func (m *ConditionalInterDumobilityInformation) SetTargetgNbDuuef1ApID(targetgNbDuuef1ApID *GnbDUUEF1ApID) *ConditionalInterDumobilityInformation {
	m.TargetgNbDuuef1Apid = targetgNbDuuef1ApID
	return m
}

func (m *ConditionalInterDumobilityInformation) SetIEExtensions(iEExtensions []*ConditionalInterDumobilityInformationExtIes) *ConditionalInterDumobilityInformation {
	m.IEExtensions = iEExtensions
	return m
}

func (m *ConditionalIntraDumobilityInformation) SetTargetCellsTocancel(targetCellsTocancel *TargetCellList) *ConditionalIntraDumobilityInformation {
	m.TargetCellsTocancel = targetCellsTocancel
	return m
}

func (m *ConditionalIntraDumobilityInformation) SetIEExtensions(iEExtensions []*ConditionalIntraDumobilityInformationExtIes) *ConditionalIntraDumobilityInformation {
	m.IEExtensions = iEExtensions
	return m
}

func (m *CriticalityDiagnostics) SetProcedureCode(procedureCode *f1apcommondatatypesv1.ProcedureCode) *CriticalityDiagnostics {
	m.ProcedureCode = procedureCode
	return m
}

func (m *CriticalityDiagnostics) SetTriggeringMessage(triggeringMessage *f1apcommondatatypesv1.TriggeringMessage) *CriticalityDiagnostics {
	m.TriggeringMessage = triggeringMessage
	return m
}

func (m *CriticalityDiagnostics) SetProcedureCriticality(procedureCriticality *f1apcommondatatypesv1.Criticality) *CriticalityDiagnostics {
	m.ProcedureCriticality = procedureCriticality
	return m
}

func (m *CriticalityDiagnostics) SetTransactionID(transactionID *TransactionId) *CriticalityDiagnostics {
	m.TransactionId = transactionID
	return m
}

func (m *CriticalityDiagnostics) SetIEsCriticalityDiagnostics(iEsCriticalityDiagnostics *CriticalityDiagnosticsIEList) *CriticalityDiagnostics {
	m.IEsCriticalityDiagnostics = iEsCriticalityDiagnostics
	return m
}

func (m *CriticalityDiagnostics) SetIEExtensions(iEExtensions []*CriticalityDiagnosticsExtIes) *CriticalityDiagnostics {
	m.IEExtensions = iEExtensions
	return m
}

func (m *CriticalityDiagnosticsIEItem) SetIEExtensions(iEExtensions []*CriticalityDiagnosticsIEItemExtIes) *CriticalityDiagnosticsIEItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *Cuduriminformation) SetIEExtensions(iEExtensions []*CuduriminformationExtIes) *Cuduriminformation {
	m.IEExtensions = iEExtensions
	return m
}

func (m *CutoDurrcinformation) SetCGConfigInfo(cGConfigInfo *CgConfigInfo) *CutoDurrcinformation {
	m.CGConfigInfo = cGConfigInfo
	return m
}

func (m *CutoDurrcinformation) SetUECapabilityRatContainerList(uECapabilityRatContainerList *UeCapabilityRatContainerList) *CutoDurrcinformation {
	m.UECapabilityRatContainerList = uECapabilityRatContainerList
	return m
}

func (m *CutoDurrcinformation) SetMeasConfig(measConfig *MeasConfig) *CutoDurrcinformation {
	m.MeasConfig = measConfig
	return m
}

func (m *CutoDurrcinformation) SetIEExtensions(iEExtensions []*CutoDurrcinformationExtIes) *CutoDurrcinformation {
	m.IEExtensions = iEExtensions
	return m
}

func (m *DedicatedSIdeliveryNeededUeItem) SetIEExtensions(iEExtensions []*DedicatedSideliveryNeededUeItemExtIes) *DedicatedSIdeliveryNeededUeItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *DlPRs) SetDlPrsresourceID(dlPrsresourceID *PrsResourceID) *DlPRs {
	m.DlPrsresourceId = dlPrsresourceID
	return m
}

func (m *DlPRs) SetIEExtensions(iEExtensions []*DlPRsExtIes) *DlPRs {
	m.IEExtensions = iEExtensions
	return m
}

func (m *DlprsresourceCoordinates) SetIEExtensions(iEExtensions []*DlprsresourceCoordinatesExtIes) *DlprsresourceCoordinates {
	m.IEExtensions = iEExtensions
	return m
}

func (m *DlprsresourceSetArp) SetIEExtensions(iEExtensions []*DlprsresourceSetArpExtIes) *DlprsresourceSetArp {
	m.IEExtensions = iEExtensions
	return m
}

func (m *DlprsresourceArp) SetIEExtensions(iEExtensions []*DlprsresourceArpExtIes) *DlprsresourceArp {
	m.IEExtensions = iEExtensions
	return m
}

func (m *DlUPTNlAddresstoUpdateListItem) SetIEExtensions(iEExtensions []*DlUPTNlAddresstoUpdateListItemExtIes) *DlUPTNlAddresstoUpdateListItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *DluptnlinformationToBeSetupItem) SetIEExtensions(iEExtensions []*DluptnlinformationToBeSetupItemExtIes) *DluptnlinformationToBeSetupItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *DrbActivityItem) SetDRbActivity(dRbActivity *DrbActivity) *DrbActivityItem {
	m.DRbActivity = dRbActivity
	return m
}

func (m *DrbActivityItem) SetIEExtensions(iEExtensions []*DrbActivityItemExtIes) *DrbActivityItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *DrbsFailedToBeModifiedItem) SetCause(cause *Cause) *DrbsFailedToBeModifiedItem {
	m.Cause = cause
	return m
}

func (m *DrbsFailedToBeModifiedItem) SetIEExtensions(iEExtensions []*DrbsFailedToBeModifiedItemExtIes) *DrbsFailedToBeModifiedItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *DrbsFailedToBeSetupItem) SetCause(cause *Cause) *DrbsFailedToBeSetupItem {
	m.Cause = cause
	return m
}

func (m *DrbsFailedToBeSetupItem) SetIEExtensions(iEExtensions []*DrbsFailedToBeSetupItemExtIes) *DrbsFailedToBeSetupItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *DrbsFailedToBeSetupModItem) SetCause(cause *Cause) *DrbsFailedToBeSetupModItem {
	m.Cause = cause
	return m
}

func (m *DrbsFailedToBeSetupModItem) SetIEExtensions(iEExtensions []*DrbsFailedToBeSetupModItemExtIes) *DrbsFailedToBeSetupModItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *DrbInformation) SetNotificationControl(notificationControl *NotificationControl) *DrbInformation {
	m.NotificationControl = notificationControl
	return m
}

func (m *DrbInformation) SetIEExtensions(iEExtensions []*DrbInformationItemExtIes) *DrbInformation {
	m.IEExtensions = iEExtensions
	return m
}

func (m *DrbsModifiedItem) SetLCID(lCID *Lcid) *DrbsModifiedItem {
	m.LCid = lCID
	return m
}

func (m *DrbsModifiedItem) SetIEExtensions(iEExtensions []*DrbsModifiedItemExtIes) *DrbsModifiedItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *DrbsModifiedConfItem) SetIEExtensions(iEExtensions []*DrbsModifiedConfItemExtIes) *DrbsModifiedConfItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *DrbNotifyItem) SetIEExtensions(iEExtensions []*DrbNotifyItemExtIes) *DrbNotifyItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *DrbsRequiredToBeModifiedItem) SetIEExtensions(iEExtensions []*DrbsRequiredToBeModifiedItemExtIes) *DrbsRequiredToBeModifiedItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *DrbsRequiredToBeReleasedItem) SetIEExtensions(iEExtensions []*DrbsRequiredToBeReleasedItemExtIes) *DrbsRequiredToBeReleasedItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *DrbsSetupItem) SetLCID(lCID *Lcid) *DrbsSetupItem {
	m.LCid = lCID
	return m
}

func (m *DrbsSetupItem) SetIEExtensions(iEExtensions []*DrbsSetupItemExtIes) *DrbsSetupItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *DrbsSetupModItem) SetLCID(lCID *Lcid) *DrbsSetupModItem {
	m.LCid = lCID
	return m
}

func (m *DrbsSetupModItem) SetIEExtensions(iEExtensions []*DrbsSetupModItemExtIes) *DrbsSetupModItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *DrbsToBeModifiedItem) SetQoSinformation(qoSinformation *QoSinformation) *DrbsToBeModifiedItem {
	m.QoSinformation = qoSinformation
	return m
}

func (m *DrbsToBeModifiedItem) SetULconfiguration(uLconfiguration *Ulconfiguration) *DrbsToBeModifiedItem {
	m.ULconfiguration = uLconfiguration
	return m
}

func (m *DrbsToBeModifiedItem) SetIEExtensions(iEExtensions []*DrbsToBeModifiedItemExtIes) *DrbsToBeModifiedItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *DrbsToBeReleasedItem) SetIEExtensions(iEExtensions []*DrbsToBeReleasedItemExtIes) *DrbsToBeReleasedItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *DrbsToBeSetupItem) SetULconfiguration(uLconfiguration *Ulconfiguration) *DrbsToBeSetupItem {
	m.ULconfiguration = uLconfiguration
	return m
}

func (m *DrbsToBeSetupItem) SetDuplicationActivation(duplicationActivation *DuplicationActivation) *DrbsToBeSetupItem {
	m.DuplicationActivation = duplicationActivation
	return m
}

func (m *DrbsToBeSetupItem) SetIEExtensions(iEExtensions []*DrbsToBeSetupItemExtIes) *DrbsToBeSetupItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *DrbsToBeSetupModItem) SetULconfiguration(uLconfiguration *Ulconfiguration) *DrbsToBeSetupModItem {
	m.ULconfiguration = uLconfiguration
	return m
}

func (m *DrbsToBeSetupModItem) SetDuplicationActivation(duplicationActivation *DuplicationActivation) *DrbsToBeSetupModItem {
	m.DuplicationActivation = duplicationActivation
	return m
}

func (m *DrbsToBeSetupModItem) SetIEExtensions(iEExtensions []*DrbsToBeSetupModItemExtIes) *DrbsToBeSetupModItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *Drxcycle) SetShortDrxcycleLength(shortDrxcycleLength *ShortDrxcycleLength) *Drxcycle {
	m.ShortDrxcycleLength = shortDrxcycleLength
	return m
}

func (m *Drxcycle) SetShortDrxcycleTimer(shortDrxcycleTimer *ShortDrxcycleTimer) *Drxcycle {
	m.ShortDrxcycleTimer = shortDrxcycleTimer
	return m
}

func (m *Drxcycle) SetIEExtensions(iEExtensions []*DrxcycleExtIes) *Drxcycle {
	m.IEExtensions = iEExtensions
	return m
}

func (m *Ducuriminformation) SetIEExtensions(iEExtensions []*DucuriminformationExtIes) *Ducuriminformation {
	m.IEExtensions = iEExtensions
	return m
}

func (m *DutoCurrcinformation) SetMeasGapConfig(measGapConfig *MeasGapConfig) *DutoCurrcinformation {
	m.MeasGapConfig = measGapConfig
	return m
}

func (m *DutoCurrcinformation) SetRequestedPMaxFr1(requestedPMaxFr1 []byte) *DutoCurrcinformation {
	m.RequestedPMaxFr1 = requestedPMaxFr1
	return m
}

func (m *DutoCurrcinformation) SetIEExtensions(iEExtensions []*DutoCurrcinformationExtIes) *DutoCurrcinformation {
	m.IEExtensions = iEExtensions
	return m
}

func (m *Dynamic5Qidescriptor) SetFiveQi(fiveQi int32) *Dynamic5Qidescriptor {
	m.FiveQi = &fiveQi
	return m
}

func (m *Dynamic5Qidescriptor) SetDelayCritical(delayCritical DelayCriticalDynamic5Qidescriptor) *Dynamic5Qidescriptor {
	m.DelayCritical = &delayCritical
	return m
}

func (m *Dynamic5Qidescriptor) SetAveragingWindow(averagingWindow *AveragingWindow) *Dynamic5Qidescriptor {
	m.AveragingWindow = averagingWindow
	return m
}

func (m *Dynamic5Qidescriptor) SetMaxDataBurstVolume(maxDataBurstVolume *MaxDataBurstVolume) *Dynamic5Qidescriptor {
	m.MaxDataBurstVolume = maxDataBurstVolume
	return m
}

func (m *Dynamic5Qidescriptor) SetIEExtensions(iEExtensions []*Dynamic5QidescriptorExtIes) *Dynamic5Qidescriptor {
	m.IEExtensions = iEExtensions
	return m
}

func (m *DynamicPqidescriptor) SetResourceType(resourceType ResourceTypeDynamicPqidescriptor) *DynamicPqidescriptor {
	m.ResourceType = &resourceType
	return m
}

func (m *DynamicPqidescriptor) SetAveragingWindow(averagingWindow *AveragingWindow) *DynamicPqidescriptor {
	m.AveragingWindow = averagingWindow
	return m
}

func (m *DynamicPqidescriptor) SetMaxDataBurstVolume(maxDataBurstVolume *MaxDataBurstVolume) *DynamicPqidescriptor {
	m.MaxDataBurstVolume = maxDataBurstVolume
	return m
}

func (m *DynamicPqidescriptor) SetIEExtensions(iEExtensions []*DynamicPqidescriptorExtIes) *DynamicPqidescriptor {
	m.IEExtensions = iEExtensions
	return m
}

func (m *ECIdMeasurementQuantitiesItem) SetIEExtensions(iEExtensions []*ECIdMeasurementQuantitiesValueExtIes) *ECIdMeasurementQuantitiesItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *ECIdMeasurementResult) SetGeographicalCoordinates(geographicalCoordinates *GeographicalCoordinates) *ECIdMeasurementResult {
	m.GeographicalCoordinates = geographicalCoordinates
	return m
}

func (m *ECIdMeasurementResult) SetMeasuredResultsList(measuredResultsList *ECIdMeasuredResultsList) *ECIdMeasurementResult {
	m.MeasuredResultsList = measuredResultsList
	return m
}

func (m *ECIdMeasurementResult) SetIEExtensions(iEExtensions []*ECIdMeasurementResultExtIes) *ECIdMeasurementResult {
	m.IEExtensions = iEExtensions
	return m
}

func (m *ECIdMeasuredResultsItem) SetIEExtensions(iEExtensions []*ECIdMeasuredResultsItemExtIes) *ECIdMeasuredResultsItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *EgressBhrlcchitem) SetIEExtensions(iEExtensions []*EgressBhrlcchitemExtIes) *EgressBhrlcchitem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *EndpointIPaddressandport) SetIEExtensions(iEExtensions []*EndpointIPaddressandportExtIes) *EndpointIPaddressandport {
	m.IEExtensions = iEExtensions
	return m
}

func (m *ExtendedAvailablePlmnItem) SetIEExtensions(iEExtensions []*ExtendedAvailablePlmnItemExtIes) *ExtendedAvailablePlmnItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *ExplicitFormat) SetNoofDownlinkSymbols(noofDownlinkSymbols *NoofDownlinkSymbols) *ExplicitFormat {
	m.NoofDownlinkSymbols = noofDownlinkSymbols
	return m
}

func (m *ExplicitFormat) SetNoofUplinkSymbols(noofUplinkSymbols *NoofUplinkSymbols) *ExplicitFormat {
	m.NoofUplinkSymbols = noofUplinkSymbols
	return m
}

func (m *ExplicitFormat) SetIEExtensions(iEExtensions []*ExplicitFormatExtIes) *ExplicitFormat {
	m.IEExtensions = iEExtensions
	return m
}

func (m *ExtendedServedPlmnsItem) SetTAisliceSupportList(tAisliceSupportList *SliceSupportList) *ExtendedServedPlmnsItem {
	m.TAisliceSupportList = tAisliceSupportList
	return m
}

func (m *ExtendedServedPlmnsItem) SetIEExtensions(iEExtensions []*ExtendedServedPlmnsItemExtIes) *ExtendedServedPlmnsItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *EutracellsListitem) SetIEExtensions(iEExtensions []*EutracellsListitemExtIes) *EutracellsListitem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *EutraCoexFDdInfo) SetULEarfcn(uLEarfcn *ExtendedEarfcn) *EutraCoexFDdInfo {
	m.ULEarfcn = uLEarfcn
	return m
}

func (m *EutraCoexFDdInfo) SetULTransmissionBandwIDth(uLTransmissionBandwIDth *EutraTransmissionBandwidth) *EutraCoexFDdInfo {
	m.ULTransmissionBandwidth = uLTransmissionBandwIDth
	return m
}

func (m *EutraCoexFDdInfo) SetIEExtensions(iEExtensions []*EutraCoexFDdInfoExtIes) *EutraCoexFDdInfo {
	m.IEExtensions = iEExtensions
	return m
}

func (m *EutraCoexTDdInfo) SetIEExtensions(iEExtensions []*EutraCoexTDdInfoExtIes) *EutraCoexTDdInfo {
	m.IEExtensions = iEExtensions
	return m
}

func (m *EutraPRachConfiguration) SetPrachConfigIndex(prachConfigIndex int32) *EutraPRachConfiguration {
	m.PrachConfigIndex = &prachConfigIndex
	return m
}

func (m *EutraPRachConfiguration) SetIEExtensions(iEExtensions []*EutraPRachConfigurationExtIes) *EutraPRachConfiguration {
	m.IEExtensions = iEExtensions
	return m
}

func (m *EutraSpecialSubframeInfo) SetIEExtensions(iEExtensions []*EutraSpecialSubframeInfoExtIes) *EutraSpecialSubframeInfo {
	m.IEExtensions = iEExtensions
	return m
}

func (m *EutranqoS) SetGbrQosInformation(gbrQosInformation *GbrQosInformation) *EutranqoS {
	m.GbrQosInformation = gbrQosInformation
	return m
}

func (m *EutranqoS) SetIEExtensions(iEExtensions []*EutranqoSExtIes) *EutranqoS {
	m.IEExtensions = iEExtensions
	return m
}

func (m *EutraFDdInfo) SetIEExtensions(iEExtensions []*EutraFDdInfoExtIes) *EutraFDdInfo {
	m.IEExtensions = iEExtensions
	return m
}

func (m *EutraTDdInfo) SetIEExtensions(iEExtensions []*EutraTDdInfoExtIes) *EutraTDdInfo {
	m.IEExtensions = iEExtensions
	return m
}

func (m *F1CtransferPath) SetIEExtensions(iEExtensions []*F1CtransferPathExtIes) *F1CtransferPath {
	m.IEExtensions = iEExtensions
	return m
}

func (m *FddInfo) SetIEExtensions(iEExtensions []*FddInfoExtIes) *FddInfo {
	m.IEExtensions = iEExtensions
	return m
}

func (m *FlowsMappedToDRbItem) SetIEExtensions(iEExtensions []*FlowsMappedToDRbItemExtIes) *FlowsMappedToDRbItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *FreqBandNrItem) SetIEExtensions(iEExtensions []*FreqBandNrItemExtIes) *FreqBandNrItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *FlowsMappedToSldrbItem) SetIEExtensions(iEExtensions []*FlowsMappedToSldrbItemExtIes) *FlowsMappedToSldrbItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *GbrQosInformation) SetIEExtensions(iEExtensions []*GbrQosInformationExtIes) *GbrQosInformation {
	m.IEExtensions = iEExtensions
	return m
}

func (m *GbrQoSflowInformation) SetMaxPacketLossRateDownlink(maxPacketLossRateDownlink *MaxPacketLossRate) *GbrQoSflowInformation {
	m.MaxPacketLossRateDownlink = maxPacketLossRateDownlink
	return m
}

func (m *GbrQoSflowInformation) SetMaxPacketLossRateUplink(maxPacketLossRateUplink *MaxPacketLossRate) *GbrQoSflowInformation {
	m.MaxPacketLossRateUplink = maxPacketLossRateUplink
	return m
}

func (m *GbrQoSflowInformation) SetIEExtensions(iEExtensions []*GbrQosFlowInformationExtIes) *GbrQoSflowInformation {
	m.IEExtensions = iEExtensions
	return m
}

func (m *GeographicalCoordinates) SetDLprsresourceCoordinates(dLprsresourceCoordinates *DlprsresourceCoordinates) *GeographicalCoordinates {
	m.DLprsresourceCoordinates = dLprsresourceCoordinates
	return m
}

func (m *GeographicalCoordinates) SetIEExtensions(iEExtensions []*GeographicalCoordinatesExtIes) *GeographicalCoordinates {
	m.IEExtensions = iEExtensions
	return m
}

func (m *GnbCUsystemInformation) SetIEExtensions(iEExtensions []*GnbCUsystemInformationExtIes) *GnbCUsystemInformation {
	m.IEExtensions = iEExtensions
	return m
}

func (m *GnbCUTNlAssociationSetupItem) SetIEExtensions(iEExtensions []*GnbCUTNlAssociationSetupItemExtIes) *GnbCUTNlAssociationSetupItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *GnbCUTNlAssociationFailedToSetupItem) SetIEExtensions(iEExtensions []*GnbCUTNlAssociationFailedToSetupItemExtIes) *GnbCUTNlAssociationFailedToSetupItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *GnbCUTNlAssociationToAddItem) SetIEExtensions(iEExtensions []*GnbCUTNlAssociationToAddItemExtIes) *GnbCUTNlAssociationToAddItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *GnbCUTNlAssociationToRemoveItem) SetIEExtensions(iEExtensions []*GnbCUTNlAssociationToRemoveItemExtIes) *GnbCUTNlAssociationToRemoveItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *GnbCUTNlAssociationToUpdateItem) SetTNlassociationUsage(tNlassociationUsage *TnlassociationUsage) *GnbCUTNlAssociationToUpdateItem {
	m.TNlassociationUsage = tNlassociationUsage
	return m
}

func (m *GnbCUTNlAssociationToUpdateItem) SetIEExtensions(iEExtensions []*GnbCUTNlAssociationToUpdateItemExtIes) *GnbCUTNlAssociationToUpdateItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *GnbDUCellResourceConfiguration) SetDUftransmissionPeriodicity(dUftransmissionPeriodicity *DuftransmissionPeriodicity) *GnbDUCellResourceConfiguration {
	m.DUftransmissionPeriodicity = dUftransmissionPeriodicity
	return m
}

func (m *GnbDUCellResourceConfiguration) SetDUfSlotConfigList(dUfSlotConfigList *DufSlotConfigList) *GnbDUCellResourceConfiguration {
	m.DUfSlotConfigList = dUfSlotConfigList
	return m
}

func (m *GnbDUCellResourceConfiguration) SetHNsaslotConfigList(hNsaslotConfigList *HsnaslotConfigList) *GnbDUCellResourceConfiguration {
	m.HNsaslotConfigList = hNsaslotConfigList
	return m
}

func (m *GnbDUCellResourceConfiguration) SetIEExtensions(iEExtensions []*GnbDUCellResourceConfigurationExtIes) *GnbDUCellResourceConfiguration {
	m.IEExtensions = iEExtensions
	return m
}

func (m *ExtendedGNbCUName) SetGNbCuNameVisibleString(gNbCuNameVisibleString *GnbCUNameVisibleString) *ExtendedGNbCUName {
	m.GNbCuNameVisibleString = gNbCuNameVisibleString
	return m
}

func (m *ExtendedGNbCUName) SetGNbCuNameUtf8String(gNbCuNameUtf8String *GnbCUNameUtf8String) *ExtendedGNbCUName {
	m.GNbCuNameUtf8String = gNbCuNameUtf8String
	return m
}

func (m *ExtendedGNbCUName) SetIEExtensions(iEExtensions []*ExtendedGNbCUNameExtIes) *ExtendedGNbCUName {
	m.IEExtensions = iEExtensions
	return m
}

func (m *ExtendedGNbDUName) SetGNbDuNameVisibleString(gNbDuNameVisibleString *GnbDUNameVisibleString) *ExtendedGNbDUName {
	m.GNbDuNameVisibleString = gNbDuNameVisibleString
	return m
}

func (m *ExtendedGNbDUName) SetGNbDuNameUtf8String(gNbDuNameUtf8String *GnbDUNameUtf8String) *ExtendedGNbDUName {
	m.GNbDuNameUtf8String = gNbDuNameUtf8String
	return m
}

func (m *ExtendedGNbDUName) SetIEExtensions(iEExtensions []*ExtendedGNbDUNameExtIes) *ExtendedGNbDUName {
	m.IEExtensions = iEExtensions
	return m
}

func (m *GnbDUServedCellsItem) SetGNbDuSystemInformation(gNbDuSystemInformation *GnbDUSystemInformation) *GnbDUServedCellsItem {
	m.GNbDuSystemInformation = gNbDuSystemInformation
	return m
}

func (m *GnbDUServedCellsItem) SetIEExtensions(iEExtensions []*GnbDUServedCellsItemExtIes) *GnbDUServedCellsItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *GnbDUSystemInformation) SetIEExtensions(iEExtensions []*GnbDUSystemInformationExtIes) *GnbDUSystemInformation {
	m.IEExtensions = iEExtensions
	return m
}

func (m *GnbDUTNlAssociationToRemoveItem) SetTNlassociationTransportLayerAddressgNbcu(tNlassociationTransportLayerAddressgNbcu *CpTransportLayerAddress) *GnbDUTNlAssociationToRemoveItem {
	m.TNlassociationTransportLayerAddressgNbcu = tNlassociationTransportLayerAddressgNbcu
	return m
}

func (m *GnbDUTNlAssociationToRemoveItem) SetIEExtensions(iEExtensions []*GnbDUTNlAssociationToRemoveItemExtIes) *GnbDUTNlAssociationToRemoveItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *GnbRxTxTimeDiff) SetAdditionalPathList(additionalPathList *AdditionalPathList) *GnbRxTxTimeDiff {
	m.AdditionalPathList = additionalPathList
	return m
}

func (m *GnbRxTxTimeDiff) SetIEExtensions(iEExtensions []*GnbRxTxTimeDiffExtIes) *GnbRxTxTimeDiff {
	m.IEExtensions = iEExtensions
	return m
}

func (m *GtptlaItem) SetIEExtensions(iEExtensions []*GtptlaItemExtIes) *GtptlaItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *Gtptunnel) SetIEExtensions(iEExtensions []*GtptunnelExtIes) *Gtptunnel {
	m.IEExtensions = iEExtensions
	return m
}

func (m *HardwareLoadIndicator) SetIEExtensions(iEExtensions []*HardwareLoadIndicatorExtIes) *HardwareLoadIndicator {
	m.IEExtensions = iEExtensions
	return m
}

func (m *HsnaslotConfigItem) SetHSnadownlink(hSnadownlink *Hsnadownlink) *HsnaslotConfigItem {
	m.HSnadownlink = hSnadownlink
	return m
}

func (m *HsnaslotConfigItem) SetHSnauplink(hSnauplink *Hsnauplink) *HsnaslotConfigItem {
	m.HSnauplink = hSnauplink
	return m
}

func (m *HsnaslotConfigItem) SetHSnaflexible(hSnaflexible *Hsnaflexible) *HsnaslotConfigItem {
	m.HSnaflexible = hSnaflexible
	return m
}

func (m *HsnaslotConfigItem) SetIEExtensions(iEExtensions []*HsnaslotConfigItemExtIes) *HsnaslotConfigItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *IabInfoIAbdonorCU) SetIAbStcInfo(iAbStcInfo *IabSTcInfo) *IabInfoIAbdonorCU {
	m.IAbStcInfo = iAbStcInfo
	return m
}

func (m *IabInfoIAbdonorCU) SetIEExtensions(iEExtensions []*IabInfoIAbdonorCUExtIes) *IabInfoIAbdonorCU {
	m.IEExtensions = iEExtensions
	return m
}

func (m *IabInfoIAbDU) SetMultiplexingInfo(multiplexingInfo *MultiplexingInfo) *IabInfoIAbDU {
	m.MultiplexingInfo = multiplexingInfo
	return m
}

func (m *IabInfoIAbDU) SetIAbStcInfo(iAbStcInfo *IabSTcInfo) *IabInfoIAbDU {
	m.IAbStcInfo = iAbStcInfo
	return m
}

func (m *IabInfoIAbDU) SetIEExtensions(iEExtensions []*IabInfoIAbDUExtIes) *IabInfoIAbDU {
	m.IEExtensions = iEExtensions
	return m
}

func (m *IabMTCellListItem) SetIEExtensions(iEExtensions []*IabMTCellListItemExtIes) *IabMTCellListItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *IabSTcInfo) SetIEExtensions(iEExtensions []*IabSTcInfoExtIes) *IabSTcInfo {
	m.IEExtensions = iEExtensions
	return m
}

func (m *IabSTcInfoItem) SetIEExtensions(iEExtensions []*IabSTcInfoItemExtIes) *IabSTcInfoItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *IabAllocatedTNlAddressItem) SetIAbtnladdressUsage(iAbtnladdressUsage *IabtnladdressUsage) *IabAllocatedTNlAddressItem {
	m.IAbtnladdressUsage = iAbtnladdressUsage
	return m
}

func (m *IabAllocatedTNlAddressItem) SetIEExtensions(iEExtensions []*IabAllocatedTNlAddressItemExtIes) *IabAllocatedTNlAddressItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *IabDUCellResourceConfigurationFDdInfo) SetIEExtensions(iEExtensions []*IabDUCellResourceConfigurationFDdInfoExtIes) *IabDUCellResourceConfigurationFDdInfo {
	m.IEExtensions = iEExtensions
	return m
}

func (m *IabDUCellResourceConfigurationTDdInfo) SetIEExtensions(iEExtensions []*IabDUCellResourceConfigurationTDdInfoExtIes) *IabDUCellResourceConfigurationTDdInfo {
	m.IEExtensions = iEExtensions
	return m
}

func (m *IabtnladdressesRequested) SetTNladdressesOrPrefixesRequestedAllTraffic(tNladdressesOrPrefixesRequestedAllTraffic int32) *IabtnladdressesRequested {
	m.TNladdressesOrPrefixesRequestedAllTraffic = &tNladdressesOrPrefixesRequestedAllTraffic
	return m
}

func (m *IabtnladdressesRequested) SetTNladdressesOrPrefixesRequestedF1C(tNladdressesOrPrefixesRequestedF1C int32) *IabtnladdressesRequested {
	m.TNladdressesOrPrefixesRequestedF1C = &tNladdressesOrPrefixesRequestedF1C
	return m
}

func (m *IabtnladdressesRequested) SetTNladdressesOrPrefixesRequestedF1U(tNladdressesOrPrefixesRequestedF1U int32) *IabtnladdressesRequested {
	m.TNladdressesOrPrefixesRequestedF1U = &tNladdressesOrPrefixesRequestedF1U
	return m
}

func (m *IabtnladdressesRequested) SetTNladdressesOrPrefixesRequestedNoNf1(tNladdressesOrPrefixesRequestedNoNf1 int32) *IabtnladdressesRequested {
	m.TNladdressesOrPrefixesRequestedNoNf1 = &tNladdressesOrPrefixesRequestedNoNf1
	return m
}

func (m *IabtnladdressesRequested) SetIEExtensions(iEExtensions []*IabtnladdressesRequestedExtIes) *IabtnladdressesRequested {
	m.IEExtensions = iEExtensions
	return m
}

func (m *IabTNlAddressesToRemoveItem) SetIEExtensions(iEExtensions []*IabTNlAddressesToRemoveItemExtIes) *IabTNlAddressesToRemoveItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *Iabv4AddressesRequested) SetIEExtensions(iEExtensions []*Iabv4AddressesRequestedExtIes) *Iabv4AddressesRequested {
	m.IEExtensions = iEExtensions
	return m
}

func (m *ImplicitFormat) SetIEExtensions(iEExtensions []*ImplicitFormatExtIes) *ImplicitFormat {
	m.IEExtensions = iEExtensions
	return m
}

func (m *IntendedTddDLULconfig) SetIEExtensions(iEExtensions []*IntendedTddDLULconfigExtIes) *IntendedTddDLULconfig {
	m.IEExtensions = iEExtensions
	return m
}

func (m *IpheaderInformation) SetDsInformationList(dsInformationList *DsinformationList) *IpheaderInformation {
	m.DsInformationList = dsInformationList
	return m
}

func (m *IpheaderInformation) SetIPv6FlowLabel(iPv6FlowLabel *asn1.BitString) *IpheaderInformation {
	m.IPv6FlowLabel = iPv6FlowLabel
	return m
}

func (m *IpheaderInformation) SetIEExtensions(iEExtensions []*IpheaderInformationItemExtIes) *IpheaderInformation {
	m.IEExtensions = iEExtensions
	return m
}

func (m *Iptolayer2TrafficMappingInfo) SetIPtolayer2TrafficMappingInfoToAdd(iPtolayer2TrafficMappingInfoToAdd *Iptolayer2TrafficMappingInfoList) *Iptolayer2TrafficMappingInfo {
	m.IPtolayer2TrafficMappingInfoToAdd = iPtolayer2TrafficMappingInfoToAdd
	return m
}

func (m *Iptolayer2TrafficMappingInfo) SetIPtolayer2TrafficMappingInfoToRemove(iPtolayer2TrafficMappingInfoToRemove *MappingInformationtoRemove) *Iptolayer2TrafficMappingInfo {
	m.IPtolayer2TrafficMappingInfoToRemove = iPtolayer2TrafficMappingInfoToRemove
	return m
}

func (m *Iptolayer2TrafficMappingInfo) SetIEExtensions(iEExtensions []*Iptolayer2TrafficMappingInfoItemExtIes) *Iptolayer2TrafficMappingInfo {
	m.IEExtensions = iEExtensions
	return m
}

func (m *Iptolayer2TrafficMappingInfoItem) SetIEExtensions(iEExtensions []*Iptolayer2TrafficMappingInfoItemExtIes) *Iptolayer2TrafficMappingInfoItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *L139Info) SetRootSequenceIndex(rootSequenceIndex int32) *L139Info {
	m.RootSequenceIndex = &rootSequenceIndex
	return m
}

func (m *L139Info) SetIEExtension(iEExtension []*L139InfoExtIes) *L139Info {
	m.IEExtension = iEExtension
	return m
}

func (m *L839Info) SetIEExtension(iEExtension []*L839InfoExtIes) *L839Info {
	m.IEExtension = iEExtension
	return m
}

func (m *LcstoGCsTranslationAoA) SetIEExtensions(iEExtensions []*LcstoGCsTranslationAoAExtIes) *LcstoGCsTranslationAoA {
	m.IEExtensions = iEExtensions
	return m
}

func (m *LcstoGcstranslation) SetAlphaFine(alphaFine int32) *LcstoGcstranslation {
	m.AlphaFine = &alphaFine
	return m
}

func (m *LcstoGcstranslation) SetBetaFine(betaFine int32) *LcstoGcstranslation {
	m.BetaFine = &betaFine
	return m
}

func (m *LcstoGcstranslation) SetGammaFine(gammaFine int32) *LcstoGcstranslation {
	m.GammaFine = &gammaFine
	return m
}

func (m *LcstoGcstranslation) SetIEExtensions(iEExtensions []*LcstoGcstranslationExtIes) *LcstoGcstranslation {
	m.IEExtensions = iEExtensions
	return m
}

func (m *LocationUncertainty) SetIEExtensions(iEExtensions []*LocationUncertaintyExtIes) *LocationUncertainty {
	m.IEExtensions = iEExtensions
	return m
}

func (m *LteuesidelinkAggregateMaximumBitrate) SetIEExtensions(iEExtensions []*LteuesidelinkAggregateMaximumBitrateExtIes) *LteuesidelinkAggregateMaximumBitrate {
	m.IEExtensions = iEExtensions
	return m
}

func (m *Ltev2XservicesAuthorized) SetVehicleUe(vehicleUe *VehicleUe) *Ltev2XservicesAuthorized {
	m.VehicleUe = vehicleUe
	return m
}

func (m *Ltev2XservicesAuthorized) SetPedestrianUe(pedestrianUe *PedestrianUe) *Ltev2XservicesAuthorized {
	m.PedestrianUe = pedestrianUe
	return m
}

func (m *Ltev2XservicesAuthorized) SetIEExtensions(iEExtensions []*Ltev2XservicesAuthorizedExtIes) *Ltev2XservicesAuthorized {
	m.IEExtensions = iEExtensions
	return m
}

func (m *MeasurementBeamInfo) SetPRsResourceID(pRsResourceID *PrsResourceID) *MeasurementBeamInfo {
	m.PRsResourceId = pRsResourceID
	return m
}

func (m *MeasurementBeamInfo) SetPRsResourceSetID(pRsResourceSetID *PrsResourceSetID) *MeasurementBeamInfo {
	m.PRsResourceSetId = pRsResourceSetID
	return m
}

func (m *MeasurementBeamInfo) SetSSbIndex(sSbIndex *SsbIndex) *MeasurementBeamInfo {
	m.SSbIndex = sSbIndex
	return m
}

func (m *MeasurementBeamInfo) SetIEExtensions(iEExtensions []*MeasurementBeamInfoExtIes) *MeasurementBeamInfo {
	m.IEExtensions = iEExtensions
	return m
}

func (m *MultiplexingInfo) SetIEExtensions(iEExtensions []*MultiplexingInfoExtIes) *MultiplexingInfo {
	m.IEExtensions = iEExtensions
	return m
}

func (m *M5Configuration) SetIEExtensions(iEExtensions []*M5ConfigurationExtIes) *M5Configuration {
	m.IEExtensions = iEExtensions
	return m
}

func (m *M6Configuration) SetIEExtensions(iEExtensions []*M6ConfigurationExtIes) *M6Configuration {
	m.IEExtensions = iEExtensions
	return m
}

func (m *M7Configuration) SetIEExtensions(iEExtensions []*M7ConfigurationExtIes) *M7Configuration {
	m.IEExtensions = iEExtensions
	return m
}

func (m *Mdtconfiguration) SetM2Configuration(m2Configuration *M2Configuration) *Mdtconfiguration {
	m.M2Configuration = m2Configuration
	return m
}

func (m *Mdtconfiguration) SetM5Configuration(m5Configuration *M5Configuration) *Mdtconfiguration {
	m.M5Configuration = m5Configuration
	return m
}

func (m *Mdtconfiguration) SetM6Configuration(m6Configuration *M6Configuration) *Mdtconfiguration {
	m.M6Configuration = m6Configuration
	return m
}

func (m *Mdtconfiguration) SetM7Configuration(m7Configuration *M7Configuration) *Mdtconfiguration {
	m.M7Configuration = m7Configuration
	return m
}

func (m *Mdtconfiguration) SetIEExtensions(iEExtensions []*MdtconfigurationExtIes) *Mdtconfiguration {
	m.IEExtensions = iEExtensions
	return m
}

func (m *NeighbourCellInformationItem) SetIntendedTddDlUlconfig(intendedTddDlUlconfig *IntendedTddDLULconfig) *NeighbourCellInformationItem {
	m.IntendedTddDlUlconfig = intendedTddDlUlconfig
	return m
}

func (m *NeighbourCellInformationItem) SetIEExtensions(iEExtensions []*NeighbourCellInformationItemExtIes) *NeighbourCellInformationItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *NgranallocationAndRetentionPriority) SetIEExtensions(iEExtensions []*NgranallocationAndRetentionPriorityExtIes) *NgranallocationAndRetentionPriority {
	m.IEExtensions = iEExtensions
	return m
}

func (m *NgranhighAccuracyAccessPointPosition) SetIEExtensions(iEExtensions []*NgranhighAccuracyAccessPointPositionExtIes) *NgranhighAccuracyAccessPointPosition {
	m.IEExtensions = iEExtensions
	return m
}

func (m *NrCGiListForRestartItem) SetIEExtensions(iEExtensions []*NrCGiListForRestartItemExtIes) *NrCGiListForRestartItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *NrPRsbeamInformation) SetLCstoGcstranslationList(lCstoGcstranslationList *LcstoGcstranslationList) *NrPRsbeamInformation {
	m.LCstoGcstranslationList = lCstoGcstranslationList
	return m
}

func (m *NrPRsbeamInformation) SetIEExtensions(iEExtensions []*NrPRsbeamInformationExtIes) *NrPRsbeamInformation {
	m.IEExtensions = iEExtensions
	return m
}

func (m *NrPRsbeamInformationItem) SetIEExtensions(iEExtensions []*NrPRsbeamInformationItemExtIes) *NrPRsbeamInformationItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *NonDynamic5Qidescriptor) SetQoSpriorityLevel(qoSpriorityLevel int32) *NonDynamic5Qidescriptor {
	m.QoSpriorityLevel = &qoSpriorityLevel
	return m
}

func (m *NonDynamic5Qidescriptor) SetAveragingWindow(averagingWindow *AveragingWindow) *NonDynamic5Qidescriptor {
	m.AveragingWindow = averagingWindow
	return m
}

func (m *NonDynamic5Qidescriptor) SetMaxDataBurstVolume(maxDataBurstVolume *MaxDataBurstVolume) *NonDynamic5Qidescriptor {
	m.MaxDataBurstVolume = maxDataBurstVolume
	return m
}

func (m *NonDynamic5Qidescriptor) SetIEExtensions(iEExtensions []*NonDynamic5QidescriptorExtIes) *NonDynamic5Qidescriptor {
	m.IEExtensions = iEExtensions
	return m
}

func (m *NonDynamicPqidescriptor) SetQoSpriorityLevel(qoSpriorityLevel int32) *NonDynamicPqidescriptor {
	m.QoSpriorityLevel = &qoSpriorityLevel
	return m
}

func (m *NonDynamicPqidescriptor) SetAveragingWindow(averagingWindow *AveragingWindow) *NonDynamicPqidescriptor {
	m.AveragingWindow = averagingWindow
	return m
}

func (m *NonDynamicPqidescriptor) SetMaxDataBurstVolume(maxDataBurstVolume *MaxDataBurstVolume) *NonDynamicPqidescriptor {
	m.MaxDataBurstVolume = maxDataBurstVolume
	return m
}

func (m *NonDynamicPqidescriptor) SetIEExtensions(iEExtensions []*NonDynamicPqidescriptorExtIes) *NonDynamicPqidescriptor {
	m.IEExtensions = iEExtensions
	return m
}

func (m *NotificationInformation) SetIEExtensions(iEExtensions []*NotificationInformationExtIes) *NotificationInformation {
	m.IEExtensions = iEExtensions
	return m
}

func (m *NpnBroadcastInformationSNpn) SetIEExtension(iEExtension []*NpnBroadcastInformationSNpnExtIes) *NpnBroadcastInformationSNpn {
	m.IEExtension = iEExtension
	return m
}

func (m *NpnBroadcastInformationPNiNPn) SetIEExtension(iEExtension []*NpnBroadcastInformationPNiNPnExtIes) *NpnBroadcastInformationPNiNPn {
	m.IEExtension = iEExtension
	return m
}

func (m *NrcarrierItem) SetIEExtension(iEExtension []*NrcarrierItemExtIes) *NrcarrierItem {
	m.IEExtension = iEExtension
	return m
}

func (m *NrfreqInfo) SetSulInformation(sulInformation *SulInformation) *NrfreqInfo {
	m.SulInformation = sulInformation
	return m
}

func (m *NrfreqInfo) SetIEExtensions(iEExtensions []*NrfreqInfoExtIes) *NrfreqInfo {
	m.IEExtensions = iEExtensions
	return m
}

func (m *Nrcgi) SetIEExtensions(iEExtensions []*NrcgiExtIes) *Nrcgi {
	m.IEExtensions = iEExtensions
	return m
}

func (m *Nrprachconfig) SetUlPrachconfigList(ulPrachconfigList *NrprachconfigList) *Nrprachconfig {
	m.UlPrachconfigList = ulPrachconfigList
	return m
}

func (m *Nrprachconfig) SetSulPrachconfigList(sulPrachconfigList *NrprachconfigList) *Nrprachconfig {
	m.SulPrachconfigList = sulPrachconfigList
	return m
}

func (m *Nrprachconfig) SetIEExtension(iEExtension []*NrprachconfigExtIes) *Nrprachconfig {
	m.IEExtension = iEExtension
	return m
}

func (m *NrprachconfigItem) SetIEExtension(iEExtension []*NrprachconfigItemExtIes) *NrprachconfigItem {
	m.IEExtension = iEExtension
	return m
}

func (m *NumDlulsymbols) SetIEExtensions(iEExtensions []*NumDlulsymbolsExtIes) *NumDlulsymbols {
	m.IEExtensions = iEExtensions
	return m
}

func (m *Nrv2XservicesAuthorized) SetVehicleUe(vehicleUe *VehicleUe) *Nrv2XservicesAuthorized {
	m.VehicleUe = vehicleUe
	return m
}

func (m *Nrv2XservicesAuthorized) SetPedestrianUe(pedestrianUe *PedestrianUe) *Nrv2XservicesAuthorized {
	m.PedestrianUe = pedestrianUe
	return m
}

func (m *Nrv2XservicesAuthorized) SetIEExtensions(iEExtensions []*Nrv2XservicesAuthorizedExtIes) *Nrv2XservicesAuthorized {
	m.IEExtensions = iEExtensions
	return m
}

func (m *NruesidelinkAggregateMaximumBitrate) SetIEExtensions(iEExtensions []*NruesidelinkAggregateMaximumBitrateExtIes) *NruesidelinkAggregateMaximumBitrate {
	m.IEExtensions = iEExtensions
	return m
}

func (m *PacketErrorRate) SetIEExtensions(iEExtensions []*PacketErrorRateExtIes) *PacketErrorRate {
	m.IEExtensions = iEExtensions
	return m
}

func (m *PagingCellItem) SetIEExtensions(iEExtensions []*PagingCellItemExtIes) *PagingCellItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *PathlossReferenceInfo) SetIEExtensions(iEExtensions []*PathlossReferenceInfoExtIes) *PathlossReferenceInfo {
	m.IEExtensions = iEExtensions
	return m
}

func (m *Pc5QoSparameters) SetPC5QoSFlowBitRates(pC5QoSFlowBitRates *Pc5FlowBitRates) *Pc5QoSparameters {
	m.PC5QoSFlowBitRates = pC5QoSFlowBitRates
	return m
}

func (m *Pc5QoSparameters) SetIEExtensions(iEExtensions []*Pc5QoSparametersExtIes) *Pc5QoSparameters {
	m.IEExtensions = iEExtensions
	return m
}

func (m *Pc5FlowBitRates) SetIEExtensions(iEExtensions []*Pc5FlowBitRatesExtIes) *Pc5FlowBitRates {
	m.IEExtensions = iEExtensions
	return m
}

func (m *PeriodicityListItem) SetIEExtensions(iEExtensions []*PeriodicityListItemExtIes) *PeriodicityListItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *PosMeasurementQuantitiesItem) SetTimingReportingGranularityFactor(timingReportingGranularityFactor int32) *PosMeasurementQuantitiesItem {
	m.TimingReportingGranularityFactor = &timingReportingGranularityFactor
	return m
}

func (m *PosMeasurementQuantitiesItem) SetIEExtensions(iEExtensions []*PosMeasurementQuantitiesItemExtIes) *PosMeasurementQuantitiesItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *PosMeasurementResultItem) SetMeasurementQuality(measurementQuality *TrpmeasurementQuality) *PosMeasurementResultItem {
	m.MeasurementQuality = measurementQuality
	return m
}

func (m *PosMeasurementResultItem) SetMeasurementBeamInfo(measurementBeamInfo *MeasurementBeamInfo) *PosMeasurementResultItem {
	m.MeasurementBeamInfo = measurementBeamInfo
	return m
}

func (m *PosMeasurementResultItem) SetIEExtensions(iEExtensions []*PosMeasurementResultItemExtIes) *PosMeasurementResultItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *PosMeasurementResultListItem) SetIEExtensions(iEExtensions []*PosMeasurementResultListItemExtIes) *PosMeasurementResultListItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *PosResourceSetTypePr) SetIEExtensions(iEExtensions []*PosResourceSetTypePrExtIes) *PosResourceSetTypePr {
	m.IEExtensions = iEExtensions
	return m
}

func (m *PosResourceSetTypeSp) SetIEExtensions(iEExtensions []*PosResourceSetTypeSpExtIes) *PosResourceSetTypeSp {
	m.IEExtensions = iEExtensions
	return m
}

func (m *PosResourceSetTypeAp) SetIEExtensions(iEExtensions []*PosResourceSetTypeApExtIes) *PosResourceSetTypeAp {
	m.IEExtensions = iEExtensions
	return m
}

func (m *PosSrsresourceItem) SetSpatialRelationPos(spatialRelationPos *SpatialRelationPos) *PosSrsresourceItem {
	m.SpatialRelationPos = spatialRelationPos
	return m
}

func (m *PosSrsresourceItem) SetIEExtensions(iEExtensions []*PosSrsresourceItemExtIes) *PosSrsresourceItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *PosSrsresourceSetItem) SetIEExtensions(iEExtensions []*PosSrsresourceSetItemExtIes) *PosSrsresourceSetItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *ProtectedEUtraResourcesItem) SetIEExtensions(iEExtensions []*ProtectedEUtraResourcesItemExtIes) *ProtectedEUtraResourcesItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *Prsconfiguration) SetIEExtensions(iEExtensions []*PrsconfigurationExtIes) *Prsconfiguration {
	m.IEExtensions = iEExtensions
	return m
}

func (m *PrsinformationPos) SetPRsResourceIDpos(pRsResourceIDpos int32) *PrsinformationPos {
	m.PRsResourceIdpos = &pRsResourceIDpos
	return m
}

func (m *PrsinformationPos) SetIEExtensions(iEExtensions []*PrsinformationPosExtIes) *PrsinformationPos {
	m.IEExtensions = iEExtensions
	return m
}

func (m *PotentialSpCellItem) SetIEExtensions(iEExtensions []*PotentialSpCellItemExtIes) *PotentialSpCellItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *PrsangleItem) SetIEExtensions(iEExtensions []*PrsangleItemItemExtIes) *PrsangleItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *Prsmuting) SetPRsmutingOption1(pRsmutingOption1 *PrsmutingOption1) *Prsmuting {
	m.PRsmutingOption1 = pRsmutingOption1
	return m
}

func (m *Prsmuting) SetPRsmutingOption2(pRsmutingOption2 *PrsmutingOption2) *Prsmuting {
	m.PRsmutingOption2 = pRsmutingOption2
	return m
}

func (m *Prsmuting) SetIEExtensions(iEExtensions []*PrsmutingExtIes) *Prsmuting {
	m.IEExtensions = iEExtensions
	return m
}

func (m *PrsmutingOption1) SetIEExtensions(iEExtensions []*PrsmutingOption1ExtIes) *PrsmutingOption1 {
	m.IEExtensions = iEExtensions
	return m
}

func (m *PrsmutingOption2) SetIEExtensions(iEExtensions []*PrsmutingOption2ExtIes) *PrsmutingOption2 {
	m.IEExtensions = iEExtensions
	return m
}

func (m *PrsresourceItem) SetQClinfo(qClinfo *PrsresourceQClinfo) *PrsresourceItem {
	m.QClinfo = qClinfo
	return m
}

func (m *PrsresourceItem) SetIEExtensions(iEExtensions []*PrsresourceItemExtIes) *PrsresourceItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *PrsresourceQClsourceSsb) SetSSbIndex(sSbIndex *SsbIndex) *PrsresourceQClsourceSsb {
	m.SSbIndex = sSbIndex
	return m
}

func (m *PrsresourceQClsourceSsb) SetIEExtensions(iEExtensions []*PrsresourceQClsourceSsbExtIes) *PrsresourceQClsourceSsb {
	m.IEExtensions = iEExtensions
	return m
}

func (m *PrsresourceQClsourcePrs) SetQClsourcePrsresourceID(qClsourcePrsresourceID *PrsResourceID) *PrsresourceQClsourcePrs {
	m.QClsourcePrsresourceId = qClsourcePrsresourceID
	return m
}

func (m *PrsresourceQClsourcePrs) SetIEExtensions(iEExtensions []*PrsresourceQClsourcePrsExtIes) *PrsresourceQClsourcePrs {
	m.IEExtensions = iEExtensions
	return m
}

func (m *PrsresourceSetItem) SetPRsmuting(pRsmuting *Prsmuting) *PrsresourceSetItem {
	m.PRsmuting = pRsmuting
	return m
}

func (m *PrsresourceSetItem) SetIEExtensions(iEExtensions []*PrsresourceSetItemExtIes) *PrsresourceSetItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *PwsFailedNRCGiItem) SetIEExtensions(iEExtensions []*PwsFailedNRCGiItemExtIes) *PwsFailedNRCGiItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *PwssystemInformation) SetIEExtensions(iEExtensions []*PwssystemInformationExtIes) *PwssystemInformation {
	m.IEExtensions = iEExtensions
	return m
}

func (m *QoSflowLevelQoSparameters) SetGBrQoSFlowInformation(gBrQoSFlowInformation *GbrQoSflowInformation) *QoSflowLevelQoSparameters {
	m.GBrQoSFlowInformation = gBrQoSFlowInformation
	return m
}

func (m *QoSflowLevelQoSparameters) SetReflectiveQoSAttribute(reflectiveQoSAttribute ReflectiveQoSattributeQoSflowLevelQoSparameters) *QoSflowLevelQoSparameters {
	m.ReflectiveQoSAttribute = &reflectiveQoSAttribute
	return m
}

func (m *QoSflowLevelQoSparameters) SetIEExtensions(iEExtensions []*QoSflowLevelQoSparametersExtIes) *QoSflowLevelQoSparameters {
	m.IEExtensions = iEExtensions
	return m
}

func (m *RachreportInformationItem) SetUEassitantIDentifier(uEassitantIDentifier *GnbDUUEF1ApID) *RachreportInformationItem {
	m.UEassitantIdentifier = uEassitantIDentifier
	return m
}

func (m *RachreportInformationItem) SetIEExtensions(iEExtensions []*RachreportInformationItemExtIes) *RachreportInformationItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *RadioResourceStatus) SetIEExtensions(iEExtensions []*RadioResourceStatusExtIes) *RadioResourceStatus {
	m.IEExtensions = iEExtensions
	return m
}

func (m *RanuepagingIdentity) SetIEExtensions(iEExtensions []*RanuepagingIdentityExtIes) *RanuepagingIdentity {
	m.IEExtensions = iEExtensions
	return m
}

func (m *RelativeCartesianLocation) SetIEExtensions(iEExtensions []*RelativeCartesianLocationExtIes) *RelativeCartesianLocation {
	m.IEExtensions = iEExtensions
	return m
}

func (m *RelativeGeodeticLocation) SetIEExtensions(iEExtensions []*RelativeGeodeticLocationExtIes) *RelativeGeodeticLocation {
	m.IEExtensions = iEExtensions
	return m
}

func (m *RequestedSrstransmissionCharacteristics) SetNumberOfTransmissions(numberOfTransmissions int32) *RequestedSrstransmissionCharacteristics {
	m.NumberOfTransmissions = &numberOfTransmissions
	return m
}

func (m *RequestedSrstransmissionCharacteristics) SetSRsresourceSetList(sRsresourceSetList *SrsresourceSetList) *RequestedSrstransmissionCharacteristics {
	m.SRsresourceSetList = sRsresourceSetList
	return m
}

func (m *RequestedSrstransmissionCharacteristics) SetSSbinformation(sSbinformation *Ssbinformation) *RequestedSrstransmissionCharacteristics {
	m.SSbinformation = sSbinformation
	return m
}

func (m *RequestedSrstransmissionCharacteristics) SetIEExtensions(iEExtensions []*RequestedSrstransmissionCharacteristicsExtIes) *RequestedSrstransmissionCharacteristics {
	m.IEExtensions = iEExtensions
	return m
}

func (m *ResourceCoordinationEutracellInfo) SetIEExtensions(iEExtensions []*ResourceCoordinationEutracellInfoExtIes) *ResourceCoordinationEutracellInfo {
	m.IEExtensions = iEExtensions
	return m
}

func (m *ResourceCoordinationTransferInformation) SetResourceCoordinationEutracellInfo(resourceCoordinationEutracellInfo *ResourceCoordinationEutracellInfo) *ResourceCoordinationTransferInformation {
	m.ResourceCoordinationEutracellInfo = resourceCoordinationEutracellInfo
	return m
}

func (m *ResourceCoordinationTransferInformation) SetIEExtensions(iEExtensions []*ResourceCoordinationTransferInformationExtIes) *ResourceCoordinationTransferInformation {
	m.IEExtensions = iEExtensions
	return m
}

func (m *ResourceSetTypePeriodic) SetIEExtensions(iEExtensions []*ResourceSetTypePeriodicExtIes) *ResourceSetTypePeriodic {
	m.IEExtensions = iEExtensions
	return m
}

func (m *ResourceSetTypeSemipersistent) SetIEExtensions(iEExtensions []*ResourceSetTypeSemipersistentExtIes) *ResourceSetTypeSemipersistent {
	m.IEExtensions = iEExtensions
	return m
}

func (m *ResourceSetTypeAperiodic) SetIEExtensions(iEExtensions []*ResourceSetTypeAperiodicExtIes) *ResourceSetTypeAperiodic {
	m.IEExtensions = iEExtensions
	return m
}

func (m *ReportingRequestType) SetReportingPeriodicityValue(reportingPeriodicityValue *ReportingPeriodicityValue) *ReportingRequestType {
	m.ReportingPeriodicityValue = reportingPeriodicityValue
	return m
}

func (m *ReportingRequestType) SetIEExtensions(iEExtensions []*ReportingRequestTypeExtIes) *ReportingRequestType {
	m.IEExtensions = iEExtensions
	return m
}

func (m *ResourceTypePeriodic) SetIEExtensions(iEExtensions []*ResourceTypePeriodicExtIes) *ResourceTypePeriodic {
	m.IEExtensions = iEExtensions
	return m
}

func (m *ResourceTypeSemipersistent) SetIEExtensions(iEExtensions []*ResourceTypeSemipersistentExtIes) *ResourceTypeSemipersistent {
	m.IEExtensions = iEExtensions
	return m
}

func (m *ResourceTypeAperiodic) SetIEExtensions(iEExtensions []*ResourceTypeAperiodicExtIes) *ResourceTypeAperiodic {
	m.IEExtensions = iEExtensions
	return m
}

func (m *ResourceTypePeriodicPos) SetIEExtensions(iEExtensions []*ResourceTypePeriodicPosExtIes) *ResourceTypePeriodicPos {
	m.IEExtensions = iEExtensions
	return m
}

func (m *ResourceTypeSemipersistentPos) SetIEExtensions(iEExtensions []*ResourceTypeSemipersistentPosExtIes) *ResourceTypeSemipersistentPos {
	m.IEExtensions = iEExtensions
	return m
}

func (m *ResourceTypeAperiodicPos) SetIEExtensions(iEExtensions []*ResourceTypeAperiodicPosExtIes) *ResourceTypeAperiodicPos {
	m.IEExtensions = iEExtensions
	return m
}

func (m *RlcduplicationInformation) SetPrimaryPathIndication(primaryPathIndication *PrimaryPathIndication) *RlcduplicationInformation {
	m.PrimaryPathIndication = primaryPathIndication
	return m
}

func (m *RlcduplicationInformation) SetIEExtensions(iEExtensions []*RlcduplicationInformationExtIes) *RlcduplicationInformation {
	m.IEExtensions = iEExtensions
	return m
}

func (m *RlcduplicationStateItem) SetIEExtensions(iEExtensions []*RlcduplicationStateItemExtIes) *RlcduplicationStateItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *RlcfailureIndication) SetIEExtensions(iEExtensions []*RlcfailureIndicationExtIes) *RlcfailureIndication {
	m.IEExtensions = iEExtensions
	return m
}

func (m *RlcStatus) SetIEExtensions(iEExtensions []*RlcStatusExtIes) *RlcStatus {
	m.IEExtensions = iEExtensions
	return m
}

func (m *RlfreportInformationItem) SetUEassitantIDentifier(uEassitantIDentifier *GnbDUUEF1ApID) *RlfreportInformationItem {
	m.UEassitantIdentifier = uEassitantIDentifier
	return m
}

func (m *RlfreportInformationItem) SetIEExtensions(iEExtensions []*RlfreportInformationItemExtIes) *RlfreportInformationItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *RrcdeliveryStatus) SetIEExtensions(iEExtensions []*RrcdeliveryStatusExtIes) *RrcdeliveryStatus {
	m.IEExtensions = iEExtensions
	return m
}

func (m *RrcVersion) SetIEExtensions(iEExtensions []*RrcVersionExtIes) *RrcVersion {
	m.IEExtensions = iEExtensions
	return m
}

func (m *ScellFailedtoSetupItem) SetCause(cause *Cause) *ScellFailedtoSetupItem {
	m.Cause = cause
	return m
}

func (m *ScellFailedtoSetupItem) SetIEExtensions(iEExtensions []*ScellFailedtoSetupItemExtIes) *ScellFailedtoSetupItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *ScellFailedtoSetupModItem) SetCause(cause *Cause) *ScellFailedtoSetupModItem {
	m.Cause = cause
	return m
}

func (m *ScellFailedtoSetupModItem) SetIEExtensions(iEExtensions []*ScellFailedtoSetupModItemExtIes) *ScellFailedtoSetupModItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *ScellToBeRemovedItem) SetIEExtensions(iEExtensions []*ScellToBeRemovedItemExtIes) *ScellToBeRemovedItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *ScellToBeSetupItem) SetSCellUlconfigured(sCellUlconfigured *CellUlconfigured) *ScellToBeSetupItem {
	m.SCellUlconfigured = sCellUlconfigured
	return m
}

func (m *ScellToBeSetupItem) SetIEExtensions(iEExtensions []*ScellToBeSetupItemExtIes) *ScellToBeSetupItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *ScellToBeSetupModItem) SetSCellUlconfigured(sCellUlconfigured *CellUlconfigured) *ScellToBeSetupModItem {
	m.SCellUlconfigured = sCellUlconfigured
	return m
}

func (m *ScellToBeSetupModItem) SetIEExtensions(iEExtensions []*ScellToBeSetupModItemExtIes) *ScellToBeSetupModItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *ScsSpecificCarrier) SetIEExtensions(iEExtensions []*ScsSpecificCarrierExtIes) *ScsSpecificCarrier {
	m.IEExtensions = iEExtensions
	return m
}

func (m *Searchwindowinformation) SetIEExtensions(iEExtensions []*SearchwindowinformationExtIes) *Searchwindowinformation {
	m.IEExtensions = iEExtensions
	return m
}

func (m *ServedCellInformation) SetFiveGsTac(fiveGsTac *FiveGsTAc) *ServedCellInformation {
	m.FiveGsTac = fiveGsTac
	return m
}

func (m *ServedCellInformation) SetConfiguredEpsTac(configuredEpsTac *ConfiguredEPsTAc) *ServedCellInformation {
	m.ConfiguredEpsTac = configuredEpsTac
	return m
}

func (m *ServedCellInformation) SetIEExtensions(iEExtensions []*ServedCellInformationExtIes) *ServedCellInformation {
	m.IEExtensions = iEExtensions
	return m
}

func (m *SfnOffset) SetIEExtensions(iEExtensions []*SfnOffsetExtIes) *SfnOffset {
	m.IEExtensions = iEExtensions
	return m
}

func (m *ServedCellsToAddItem) SetGNbDuSystemInformation(gNbDuSystemInformation *GnbDUSystemInformation) *ServedCellsToAddItem {
	m.GNbDuSystemInformation = gNbDuSystemInformation
	return m
}

func (m *ServedCellsToAddItem) SetIEExtensions(iEExtensions []*ServedCellsToAddItemExtIes) *ServedCellsToAddItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *ServedCellsToDeleteItem) SetIEExtensions(iEExtensions []*ServedCellsToDeleteItemExtIes) *ServedCellsToDeleteItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *ServedCellsToModifyItem) SetGNbDuSystemInformation(gNbDuSystemInformation *GnbDUSystemInformation) *ServedCellsToModifyItem {
	m.GNbDuSystemInformation = gNbDuSystemInformation
	return m
}

func (m *ServedCellsToModifyItem) SetIEExtensions(iEExtensions []*ServedCellsToModifyItemExtIes) *ServedCellsToModifyItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *ServedEUtraCellsInformation) SetIEExtensions(iEExtensions []*ServedEUtraCellInformationExtIes) *ServedEUtraCellsInformation {
	m.IEExtensions = iEExtensions
	return m
}

func (m *ServiceStatus) SetSwitchingOffOngoing(switchingOffOngoing SwitchingOffOngoingServiceStatus) *ServiceStatus {
	m.SwitchingOffOngoing = &switchingOffOngoing
	return m
}

func (m *ServiceStatus) SetIEExtensions(iEExtensions []*ServiceStatusExtIes) *ServiceStatus {
	m.IEExtensions = iEExtensions
	return m
}

func (m *SitypeItem) SetIEExtensions(iEExtensions []*SitypeItemExtIes) *SitypeItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *SibtypetobeupdatedListItem) SetIEExtensions(iEExtensions []*SibtypetobeupdatedListItemExtIes) *SibtypetobeupdatedListItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *SldrbsFailedToBeModifiedItem) SetCause(cause *Cause) *SldrbsFailedToBeModifiedItem {
	m.Cause = cause
	return m
}

func (m *SldrbsFailedToBeModifiedItem) SetIEExtensions(iEExtensions []*SldrbsFailedToBeModifiedItemExtIes) *SldrbsFailedToBeModifiedItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *SldrbsFailedToBeSetupItem) SetCause(cause *Cause) *SldrbsFailedToBeSetupItem {
	m.Cause = cause
	return m
}

func (m *SldrbsFailedToBeSetupItem) SetIEExtensions(iEExtensions []*SldrbsFailedToBeSetupItemExtIes) *SldrbsFailedToBeSetupItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *SldrbsFailedToBeSetupModItem) SetCause(cause *Cause) *SldrbsFailedToBeSetupModItem {
	m.Cause = cause
	return m
}

func (m *SldrbsFailedToBeSetupModItem) SetIEExtensions(iEExtensions []*SldrbsFailedToBeSetupModItemExtIes) *SldrbsFailedToBeSetupModItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *SldrbsModifiedItem) SetIEExtensions(iEExtensions []*SldrbsModifiedItemExtIes) *SldrbsModifiedItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *SldrbsModifiedConfItem) SetIEExtensions(iEExtensions []*SldrbsModifiedConfItemExtIes) *SldrbsModifiedConfItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *SldrbsRequiredToBeModifiedItem) SetIEExtensions(iEExtensions []*SldrbsRequiredToBeModifiedItemExtIes) *SldrbsRequiredToBeModifiedItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *SldrbsRequiredToBeReleasedItem) SetIEExtensions(iEExtensions []*SldrbsRequiredToBeReleasedItemExtIes) *SldrbsRequiredToBeReleasedItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *SldrbsSetupItem) SetIEExtensions(iEExtensions []*SldrbsSetupItemExtIes) *SldrbsSetupItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *SldrbsSetupModItem) SetIEExtensions(iEExtensions []*SldrbsSetupModItemExtIes) *SldrbsSetupModItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *SldrbsToBeModifiedItem) SetSLdrbinformation(sLdrbinformation *Sldrbinformation) *SldrbsToBeModifiedItem {
	m.SLdrbinformation = sLdrbinformation
	return m
}

func (m *SldrbsToBeModifiedItem) SetRLcmode(rLcmode *Rlcmode) *SldrbsToBeModifiedItem {
	m.RLcmode = rLcmode
	return m
}

func (m *SldrbsToBeModifiedItem) SetIEExtensions(iEExtensions []*SldrbsToBeModifiedItemExtIes) *SldrbsToBeModifiedItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *SldrbsToBeReleasedItem) SetIEExtensions(iEExtensions []*SldrbsToBeReleasedItemExtIes) *SldrbsToBeReleasedItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *SldrbsToBeSetupItem) SetIEExtensions(iEExtensions []*SldrbsToBeSetupItemExtIes) *SldrbsToBeSetupItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *SldrbsToBeSetupModItem) SetRLcmode(rLcmode *Rlcmode) *SldrbsToBeSetupModItem {
	m.RLcmode = rLcmode
	return m
}

func (m *SldrbsToBeSetupModItem) SetIEExtensions(iEExtensions []*SldrbsToBeSetupModItemExtIes) *SldrbsToBeSetupModItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *SliceAvailableCapacity) SetIEExtensions(iEExtensions []*SliceAvailableCapacityExtIes) *SliceAvailableCapacity {
	m.IEExtensions = iEExtensions
	return m
}

func (m *SliceAvailableCapacityItem) SetIEExtensions(iEExtensions []*SliceAvailableCapacityItemExtIes) *SliceAvailableCapacityItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *SnssaiavailableCapacityItem) SetSliceAvailableCapacityValueDownlink(sliceAvailableCapacityValueDownlink int32) *SnssaiavailableCapacityItem {
	m.SliceAvailableCapacityValueDownlink = &sliceAvailableCapacityValueDownlink
	return m
}

func (m *SnssaiavailableCapacityItem) SetSliceAvailableCapacityValueUplink(sliceAvailableCapacityValueUplink int32) *SnssaiavailableCapacityItem {
	m.SliceAvailableCapacityValueUplink = &sliceAvailableCapacityValueUplink
	return m
}

func (m *SnssaiavailableCapacityItem) SetIEExtensions(iEExtensions []*SnssaiavailableCapacityItemExtIes) *SnssaiavailableCapacityItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *SliceSupportItem) SetIEExtensions(iEExtensions []*SliceSupportItemExtIes) *SliceSupportItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *SliceToReportItem) SetIEExtensions(iEExtensions []*SliceToReportItemExtIes) *SliceToReportItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *SnssaiItem) SetIEExtensions(iEExtensions []*SnssaiItemExtIes) *SnssaiItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *SlotConfigurationItem) SetIEExtensions(iEExtensions []*SlotConfigurationItemExtIes) *SlotConfigurationItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *Snssai) SetSD(sD []byte) *Snssai {
	m.SD = sD
	return m
}

func (m *Snssai) SetIEExtensions(iEExtensions []*SnssaiExtIes) *Snssai {
	m.IEExtensions = iEExtensions
	return m
}

func (m *SpatialDirectionInformation) SetIEExtensions(iEExtensions []*SpatialDirectionInformationExtIes) *SpatialDirectionInformation {
	m.IEExtensions = iEExtensions
	return m
}

func (m *SpatialRelationInfo) SetIEExtensions(iEExtensions []*SpatialRelationInfoExtIes) *SpatialRelationInfo {
	m.IEExtensions = iEExtensions
	return m
}

func (m *SpatialRelationforResourceIditem) SetIEExtensions(iEExtensions []*SpatialRelationforResourceIditemExtIes) *SpatialRelationforResourceIditem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *SpatialRelationPerSrsresource) SetIEExtensions(iEExtensions []*SpatialRelationPerSrsresourceExtIes) *SpatialRelationPerSrsresource {
	m.IEExtensions = iEExtensions
	return m
}

func (m *SpatialRelationPerSrsresourceItem) SetIEExtensions(iEExtensions []*SpatialRelationPerSrsresourceItemExtIes) *SpatialRelationPerSrsresourceItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *SrbsFailedToBeSetupItem) SetCause(cause *Cause) *SrbsFailedToBeSetupItem {
	m.Cause = cause
	return m
}

func (m *SrbsFailedToBeSetupItem) SetIEExtensions(iEExtensions []*SrbsFailedToBeSetupItemExtIes) *SrbsFailedToBeSetupItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *SrbsFailedToBeSetupModItem) SetCause(cause *Cause) *SrbsFailedToBeSetupModItem {
	m.Cause = cause
	return m
}

func (m *SrbsFailedToBeSetupModItem) SetIEExtensions(iEExtensions []*SrbsFailedToBeSetupModItemExtIes) *SrbsFailedToBeSetupModItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *SrbsModifiedItem) SetIEExtensions(iEExtensions []*SrbsModifiedItemExtIes) *SrbsModifiedItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *SrbsRequiredToBeReleasedItem) SetIEExtensions(iEExtensions []*SrbsRequiredToBeReleasedItemExtIes) *SrbsRequiredToBeReleasedItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *SrbsSetupItem) SetIEExtensions(iEExtensions []*SrbsSetupItemExtIes) *SrbsSetupItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *SrbsSetupModItem) SetIEExtensions(iEExtensions []*SrbsSetupModItemExtIes) *SrbsSetupModItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *SrbsToBeReleasedItem) SetIEExtensions(iEExtensions []*SrbsToBeReleasedItemExtIes) *SrbsToBeReleasedItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *SrbsToBeSetupItem) SetDuplicationIndication(duplicationIndication *DuplicationIndication) *SrbsToBeSetupItem {
	m.DuplicationIndication = duplicationIndication
	return m
}

func (m *SrbsToBeSetupItem) SetIEExtensions(iEExtensions []*SrbsToBeSetupItemExtIes) *SrbsToBeSetupItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *SrbsToBeSetupModItem) SetDuplicationIndication(duplicationIndication *DuplicationIndication) *SrbsToBeSetupModItem {
	m.DuplicationIndication = duplicationIndication
	return m
}

func (m *SrbsToBeSetupModItem) SetIEExtensions(iEExtensions []*SrbsToBeSetupModItemExtIes) *SrbsToBeSetupModItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *SrscarrierListItem) SetPci(pci *Nrpci) *SrscarrierListItem {
	m.Pci = pci
	return m
}

func (m *SrscarrierListItem) SetIEExtensions(iEExtensions []*SrscarrierListItemExtIes) *SrscarrierListItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *Srsconfig) SetSRsresourceList(sRsresourceList *SrsresourceList) *Srsconfig {
	m.SRsresourceList = sRsresourceList
	return m
}

func (m *Srsconfig) SetPosSrsresourceList(posSrsresourceList *PosSrsresourceList) *Srsconfig {
	m.PosSrsresourceList = posSrsresourceList
	return m
}

func (m *Srsconfig) SetSRsresourceSetList(sRsresourceSetList *SrsresourceSetList) *Srsconfig {
	m.SRsresourceSetList = sRsresourceSetList
	return m
}

func (m *Srsconfig) SetPosSrsresourceSetList(posSrsresourceSetList *PosSrsresourceSetList) *Srsconfig {
	m.PosSrsresourceSetList = posSrsresourceSetList
	return m
}

func (m *Srsconfig) SetIEExtensions(iEExtensions []*SrsconfigExtIes) *Srsconfig {
	m.IEExtensions = iEExtensions
	return m
}

func (m *Srsconfiguration) SetIEExtensions(iEExtensions []*SrsconfigurationExtIes) *Srsconfiguration {
	m.IEExtensions = iEExtensions
	return m
}

func (m *Srsresource) SetIEExtensions(iEExtensions []*SrsresourceExtIes) *Srsresource {
	m.IEExtensions = iEExtensions
	return m
}

func (m *SrsresourceSet) SetIEExtensions(iEExtensions []*SrsresourceSetExtIes) *SrsresourceSet {
	m.IEExtensions = iEExtensions
	return m
}

func (m *SrsresourceSetItem) SetNumSrsresourcesperset(numSrsresourcesperset int32) *SrsresourceSetItem {
	m.NumSrsresourcesperset = &numSrsresourcesperset
	return m
}

func (m *SrsresourceSetItem) SetPeriodicityList(periodicityList *PeriodicityList) *SrsresourceSetItem {
	m.PeriodicityList = periodicityList
	return m
}

func (m *SrsresourceSetItem) SetSpatialRelationInfo(spatialRelationInfo *SpatialRelationInfo) *SrsresourceSetItem {
	m.SpatialRelationInfo = spatialRelationInfo
	return m
}

func (m *SrsresourceSetItem) SetPathlossReferenceInfo(pathlossReferenceInfo *PathlossReferenceInfo) *SrsresourceSetItem {
	m.PathlossReferenceInfo = pathlossReferenceInfo
	return m
}

func (m *SrsresourceSetItem) SetIEExtensions(iEExtensions []*SrsresourceSetItemExtIes) *SrsresourceSetItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *SrsresourceTrigger) SetIEExtensions(iEExtensions []*SrsresourceTriggerExtIes) *SrsresourceTrigger {
	m.IEExtensions = iEExtensions
	return m
}

func (m *Ssb) SetSsbIndex(ssbIndex *SsbIndex) *Ssb {
	m.SsbIndex = ssbIndex
	return m
}

func (m *Ssb) SetIEExtensions(iEExtensions []*SsbExtIes) *Ssb {
	m.IEExtensions = iEExtensions
	return m
}

func (m *SsbareaCapacityValueItem) SetIEExtensions(iEExtensions []*SsbareaCapacityValueItemExtIes) *SsbareaCapacityValueItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *SsbareaRadioResourceStatusItem) SetDLschedulingPdcchcceusage(dLschedulingPdcchcceusage int32) *SsbareaRadioResourceStatusItem {
	m.DLschedulingPdcchcceusage = &dLschedulingPdcchcceusage
	return m
}

func (m *SsbareaRadioResourceStatusItem) SetULschedulingPdcchcceusage(uLschedulingPdcchcceusage int32) *SsbareaRadioResourceStatusItem {
	m.ULschedulingPdcchcceusage = &uLschedulingPdcchcceusage
	return m
}

func (m *SsbareaRadioResourceStatusItem) SetIEExtensions(iEExtensions []*SsbareaRadioResourceStatusItemExtIes) *SsbareaRadioResourceStatusItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *Ssbinformation) SetIEExtensions(iEExtensions []*SsbinformationExtIes) *Ssbinformation {
	m.IEExtensions = iEExtensions
	return m
}

func (m *SsbinformationItem) SetIEExtensions(iEExtensions []*SsbinformationItemExtIes) *SsbinformationItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *SsbTFConfiguration) SetSSbPositionInBurst(sSbPositionInBurst *SsbPositionsInBurst) *SsbTFConfiguration {
	m.SSbPositionInBurst = sSbPositionInBurst
	return m
}

func (m *SsbTFConfiguration) SetSFninitialisationTime(sFninitialisationTime *RelativeTime1900) *SsbTFConfiguration {
	m.SFninitialisationTime = sFninitialisationTime
	return m
}

func (m *SsbTFConfiguration) SetIEExtensions(iEExtensions []*SsbTFConfigurationExtIes) *SsbTFConfiguration {
	m.IEExtensions = iEExtensions
	return m
}

func (m *SsbtoReportItem) SetIEExtensions(iEExtensions []*SsbtoReportItemExtIes) *SsbtoReportItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *SulInformation) SetIEExtensions(iEExtensions []*SulInformationExtIes) *SulInformation {
	m.IEExtensions = iEExtensions
	return m
}

func (m *SupportedSulfreqBandItem) SetIEExtensions(iEExtensions []*SupportedSulfreqBandItemExtIes) *SupportedSulfreqBandItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *TargetCellListItem) SetIEExtensions(iEExtensions []*TargetCellListItemExtIes) *TargetCellListItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *TddInfo) SetIEExtensions(iEExtensions []*TddInfoExtIes) *TddInfo {
	m.IEExtensions = iEExtensions
	return m
}

func (m *TimeReferenceInformation) SetIEExtensions(iEExtensions []*TimeReferenceInformationExtIes) *TimeReferenceInformation {
	m.IEExtensions = iEExtensions
	return m
}

func (m *TimeStamp) SetMeasurementTime(measurementTime *RelativeTime1900) *TimeStamp {
	m.MeasurementTime = measurementTime
	return m
}

func (m *TimeStamp) SetIEExtension(iEExtension []*TimeStampExtIes) *TimeStamp {
	m.IEExtension = iEExtension
	return m
}

func (m *TimingMeasurementQuality) SetIEExtensions(iEExtensions []*TimingMeasurementQualityExtIes) *TimingMeasurementQuality {
	m.IEExtensions = iEExtensions
	return m
}

func (m *TnlcapacityIndicator) SetIEExtensions(iEExtensions []*TnlcapacityIndicatorExtIes) *TnlcapacityIndicator {
	m.IEExtensions = iEExtensions
	return m
}

func (m *TraceActivation) SetIEExtensions(iEExtensions []*TraceActivationExtIes) *TraceActivation {
	m.IEExtensions = iEExtensions
	return m
}

func (m *TransmissionBandwidth) SetIEExtensions(iEExtensions []*TransmissionBandwidthExtIes) *TransmissionBandwidth {
	m.IEExtensions = iEExtensions
	return m
}

func (m *TransportUPLayerAddressInfoToAddItem) SetGTptransportLayerAddressToAdd(gTptransportLayerAddressToAdd *Gtptlas) *TransportUPLayerAddressInfoToAddItem {
	m.GTptransportLayerAddressToAdd = gTptransportLayerAddressToAdd
	return m
}

func (m *TransportUPLayerAddressInfoToAddItem) SetIEExtensions(iEExtensions []*TransportUPLayerAddressInfoToAddItemExtIes) *TransportUPLayerAddressInfoToAddItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *TransportUPLayerAddressInfoToRemoveItem) SetGTptransportLayerAddressToRemove(gTptransportLayerAddressToRemove *Gtptlas) *TransportUPLayerAddressInfoToRemoveItem {
	m.GTptransportLayerAddressToRemove = gTptransportLayerAddressToRemove
	return m
}

func (m *TransportUPLayerAddressInfoToRemoveItem) SetIEExtensions(iEExtensions []*TransportUPLayerAddressInfoToRemoveItemExtIes) *TransportUPLayerAddressInfoToRemoveItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *Trpinformation) SetIEExtensions(iEExtensions []*TrpinformationExtIes) *Trpinformation {
	m.IEExtensions = iEExtensions
	return m
}

func (m *TrpinformationItem) SetIEExtensions(iEExtensions []*TrpinformationItemExtIes) *TrpinformationItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *TrplistItem) SetIEExtensions(iEExtensions []*TrplistItemExtIes) *TrplistItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *TrpmeasurementQuality) SetIEExtensions(iEExtensions []*TrpmeasurementQualityExtIes) *TrpmeasurementQuality {
	m.IEExtensions = iEExtensions
	return m
}

func (m *TrpMeasurementRequestItem) SetSearchWindowInformation(searchWindowInformation *Searchwindowinformation) *TrpMeasurementRequestItem {
	m.SearchWindowInformation = searchWindowInformation
	return m
}

func (m *TrpMeasurementRequestItem) SetIEExtensions(iEExtensions []*TrpMeasurementRequestItemExtIes) *TrpMeasurementRequestItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *TrppositionDirect) SetIEExtensions(iEExtensions []*TrppositionDirectExtIes) *TrppositionDirect {
	m.IEExtensions = iEExtensions
	return m
}

func (m *TrppositionReferenced) SetIEExtensions(iEExtensions []*TrppositionReferencedExtIes) *TrppositionReferenced {
	m.IEExtensions = iEExtensions
	return m
}

func (m *TransportLayerAddressInfo) SetTransportUpLayerAddressInfoToAddList(transportUpLayerAddressInfoToAddList *TransportUPLayerAddressInfoToAddList) *TransportLayerAddressInfo {
	m.TransportUpLayerAddressInfoToAddList = transportUpLayerAddressInfoToAddList
	return m
}

func (m *TransportLayerAddressInfo) SetTransportUpLayerAddressInfoToRemoveList(transportUpLayerAddressInfoToRemoveList *TransportUPLayerAddressInfoToRemoveList) *TransportLayerAddressInfo {
	m.TransportUpLayerAddressInfoToRemoveList = transportUpLayerAddressInfoToRemoveList
	return m
}

func (m *TransportLayerAddressInfo) SetIEExtensions(iEExtensions []*TransportLayerAddressInfoExtIes) *TransportLayerAddressInfo {
	m.IEExtensions = iEExtensions
	return m
}

func (m *TscassistanceInformation) SetBurstArrivalTime(burstArrivalTime *BurstArrivalTime) *TscassistanceInformation {
	m.BurstArrivalTime = burstArrivalTime
	return m
}

func (m *TscassistanceInformation) SetIEExtensions(iEExtensions []*TscassistanceInformationExtIes) *TscassistanceInformation {
	m.IEExtensions = iEExtensions
	return m
}

func (m *TsctrafficCharacteristics) SetTScassistanceInformationDl(tScassistanceInformationDl *TscassistanceInformation) *TsctrafficCharacteristics {
	m.TScassistanceInformationDl = tScassistanceInformationDl
	return m
}

func (m *TsctrafficCharacteristics) SetTScassistanceInformationUl(tScassistanceInformationUl *TscassistanceInformation) *TsctrafficCharacteristics {
	m.TScassistanceInformationUl = tScassistanceInformationUl
	return m
}

func (m *TsctrafficCharacteristics) SetIEExtensions(iEExtensions []*TsctrafficCharacteristicsExtIes) *TsctrafficCharacteristics {
	m.IEExtensions = iEExtensions
	return m
}

func (m *UacAssistanceInfo) SetIEExtensions(iEExtensions []*UacAssistanceInfoExtIes) *UacAssistanceInfo {
	m.IEExtensions = iEExtensions
	return m
}

func (m *UacplmnItem) SetIEExtensions(iEExtensions []*UacplmnItemExtIes) *UacplmnItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *UactypeItem) SetIEExtensions(iEExtensions []*UactypeItemExtIes) *UactypeItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *UacoperatorDefined) SetIEExtensions(iEExtensions []*UacoperatorDefinedExtIes) *UacoperatorDefined {
	m.IEExtensions = iEExtensions
	return m
}

func (m *UeassociatedLogicalF1ConnectionItem) SetGNbCuUeF1ApID(gNbCuUeF1ApID *GnbCUUEF1ApID) *UeassociatedLogicalF1ConnectionItem {
	m.GNbCuUeF1ApId = gNbCuUeF1ApID
	return m
}

func (m *UeassociatedLogicalF1ConnectionItem) SetGNbDuUeF1ApID(gNbDuUeF1ApID *GnbDUUEF1ApID) *UeassociatedLogicalF1ConnectionItem {
	m.GNbDuUeF1ApId = gNbDuUeF1ApID
	return m
}

func (m *UeassociatedLogicalF1ConnectionItem) SetIEExtensions(iEExtensions []*UeassociatedLogicalF1ConnectionItemExtIes) *UeassociatedLogicalF1ConnectionItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *UlAoA) SetZenithAoA(zenithAoA int32) *UlAoA {
	m.ZenithAoA = &zenithAoA
	return m
}

func (m *UlAoA) SetLCsToGcsTranslationAoA(lCsToGcsTranslationAoA *LcstoGCsTranslationAoA) *UlAoA {
	m.LCsToGcsTranslationAoA = lCsToGcsTranslationAoA
	return m
}

func (m *UlAoA) SetIEExtensions(iEExtensions []*UlAoAExtIes) *UlAoA {
	m.IEExtensions = iEExtensions
	return m
}

func (m *UlBHNonUPTrafficMapping) SetIEExtensions(iEExtensions []*UlBHNonUPTrafficMappingExtIes) *UlBHNonUPTrafficMapping {
	m.IEExtensions = iEExtensions
	return m
}

func (m *UlBHNonUPTrafficMappingItem) SetIEExtensions(iEExtensions []*UlBHNonUPTrafficMappingItemExtIes) *UlBHNonUPTrafficMappingItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *Ulconfiguration) SetIEExtensions(iEExtensions []*UlconfigurationExtIes) *Ulconfiguration {
	m.IEExtensions = iEExtensions
	return m
}

func (m *UlRToaMeasurement) SetAdditionalPathList(additionalPathList *AdditionalPathList) *UlRToaMeasurement {
	m.AdditionalPathList = additionalPathList
	return m
}

func (m *UlRToaMeasurement) SetIEExtensions(iEExtensions []*UlRToaMeasurementExtIes) *UlRToaMeasurement {
	m.IEExtensions = iEExtensions
	return m
}

func (m *UlUPTNlInformationtoUpdateListItem) SetNewUluptnlinformation(newUluptnlinformation *UptransportLayerInformation) *UlUPTNlInformationtoUpdateListItem {
	m.NewUluptnlinformation = newUluptnlinformation
	return m
}

func (m *UlUPTNlInformationtoUpdateListItem) SetIEExtensions(iEExtensions []*UlUPTNlInformationtoUpdateListItemExtIes) *UlUPTNlInformationtoUpdateListItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *UlUPTNlAddresstoUpdateListItem) SetIEExtensions(iEExtensions []*UlUPTNlAddresstoUpdateListItemExtIes) *UlUPTNlAddresstoUpdateListItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *UluptnlinformationToBeSetupItem) SetIEExtensions(iEExtensions []*UluptnlinformationToBeSetupItemExtIes) *UluptnlinformationToBeSetupItem {
	m.IEExtensions = iEExtensions
	return m
}

func (m *VictimgNbsetId) SetIEExtensions(iEExtensions []*VictimgNbsetIdExtIes) *VictimgNbsetId {
	m.IEExtensions = iEExtensions
	return m
}
