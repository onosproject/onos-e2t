// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package xnapiesv1

import (
	xnapcommondatatypesv1 "github.com/onosproject/onos-e2t/api/xnap/v1/xnap-commondatatypes"

	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
)

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

func (m *AreaOfInterestItem) SetListOfTaisinAoI(listOfTaisinAoI *ListOfTaisinAoI) *AreaOfInterestItem {
	m.ListOfTaisinAoI = listOfTaisinAoI
	return m
}

func (m *AreaOfInterestItem) SetListOfCellsinAoI(listOfCellsinAoI *ListOfCells) *AreaOfInterestItem {
	m.ListOfCellsinAoI = listOfCellsinAoI
	return m
}

func (m *AreaOfInterestItem) SetListOfRannodesinAoI(listOfRannodesinAoI *ListOfRannodesinAoI) *AreaOfInterestItem {
	m.ListOfRannodesinAoI = listOfRannodesinAoI
	return m
}

func (m *AreaScopeOfNeighCellsItem) SetPciListForMdt(pciListForMdt *PcilistForMdt) *AreaScopeOfNeighCellsItem {
	m.PciListForMdt = pciListForMdt
	return m
}

func (m *AssistanceDataForRanpaging) SetRanPagingAttemptInfo(ranPagingAttemptInfo *RanpagingAttemptInfo) *AssistanceDataForRanpaging {
	m.RanPagingAttemptInfo = ranPagingAttemptInfo
	return m
}

func (m *BluetoothMeasurementConfiguration) SetBluetoothMeasConfigNameList(bluetoothMeasConfigNameList *BluetoothMeasConfigNameList) *BluetoothMeasurementConfiguration {
	m.BluetoothMeasConfigNameList = bluetoothMeasConfigNameList
	return m
}

func (m *BluetoothMeasurementConfiguration) SetBtRssi(btRssi BtrssiBluetoothMeasurementConfiguration) *BluetoothMeasurementConfiguration {
	m.BtRssi = &btRssi
	return m
}

func (m *BplmnIDInfoEUtraItem) SetRanac(ranac *Ranac) *BplmnIDInfoEUtraItem {
	m.Ranac = ranac
	return m
}

func (m *BplmnIDInfoNRItem) SetRanac(ranac *Ranac) *BplmnIDInfoNRItem {
	m.Ranac = ranac
	return m
}

func (m *CapacityValueInfo) SetSsbAreaCapacityValueList(ssbAreaCapacityValueList *SsbareaCapacityValueList) *CapacityValueInfo {
	m.SsbAreaCapacityValueList = ssbAreaCapacityValueList
	return m
}

func (m *CellAndCapacityAssistanceInfoNR) SetMaximumCellListSize(maximumCellListSize *MaximumCellListSize) *CellAndCapacityAssistanceInfoNR {
	m.MaximumCellListSize = maximumCellListSize
	return m
}

func (m *CellAndCapacityAssistanceInfoNR) SetCellAssistanceInfoNr(cellAssistanceInfoNr *CellAssistanceInfoNR) *CellAndCapacityAssistanceInfoNR {
	m.CellAssistanceInfoNr = cellAssistanceInfoNr
	return m
}

func (m *CellAndCapacityAssistanceInfoEUtra) SetMaximumCellListSize(maximumCellListSize *MaximumCellListSize) *CellAndCapacityAssistanceInfoEUtra {
	m.MaximumCellListSize = maximumCellListSize
	return m
}

func (m *CellAndCapacityAssistanceInfoEUtra) SetCellAssistanceInfoEutra(cellAssistanceInfoEutra *CellAssistanceInfoEUtra) *CellAndCapacityAssistanceInfoEUtra {
	m.CellAssistanceInfoEutra = cellAssistanceInfoEutra
	return m
}

func (m *CellMeasurementResultItem) SetRadioResourceStatus(radioResourceStatus *RadioResourceStatus) *CellMeasurementResultItem {
	m.RadioResourceStatus = radioResourceStatus
	return m
}

func (m *CellMeasurementResultItem) SetTNlcapacityIndicator(tNlcapacityIndicator *TnlcapacityIndicator) *CellMeasurementResultItem {
	m.TNlcapacityIndicator = tNlcapacityIndicator
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

func (m *CellMeasurementResultItem) SetRRcconnections(rRcconnections *Rrcconnections) *CellMeasurementResultItem {
	m.RRcconnections = rRcconnections
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

func (m *CompositeAvailableCapacity) SetCellCapacityClassValue(cellCapacityClassValue *CellCapacityClassValue) *CompositeAvailableCapacity {
	m.CellCapacityClassValue = cellCapacityClassValue
	return m
}

func (m *ChoinformationReq) SetTargetNgRannodeUexnApID(targetNgRannodeUexnApID *NgRAnnodeUexnApid) *ChoinformationReq {
	m.TargetNgRannodeUexnApid = targetNgRannodeUexnApID
	return m
}

func (m *ChoinformationReq) SetCHoEstimatedArrivalProbability(cHoEstimatedArrivalProbability *ChoProbability) *ChoinformationReq {
	m.CHoEstimatedArrivalProbability = cHoEstimatedArrivalProbability
	return m
}

func (m *ChoinformationAck) SetMaxChooperations(maxChooperations *MaxChopreparations) *ChoinformationAck {
	m.MaxChooperations = maxChooperations
	return m
}

func (m *CriticalityDiagnostics) SetProcedureCode(procedureCode *xnapcommondatatypesv1.ProcedureCode) *CriticalityDiagnostics {
	m.ProcedureCode = procedureCode
	return m
}

func (m *CriticalityDiagnostics) SetTriggeringMessage(triggeringMessage xnapcommondatatypesv1.TriggeringMessage) *CriticalityDiagnostics {
	m.TriggeringMessage = &triggeringMessage
	return m
}

func (m *CriticalityDiagnostics) SetProcedureCriticality(procedureCriticality xnapcommondatatypesv1.Criticality) *CriticalityDiagnostics {
	m.ProcedureCriticality = &procedureCriticality
	return m
}

func (m *CriticalityDiagnostics) SetIEsCriticalityDiagnostics(iEsCriticalityDiagnostics *CriticalityDiagnosticsIEList) *CriticalityDiagnostics {
	m.IEsCriticalityDiagnostics = iEsCriticalityDiagnostics
	return m
}

func (m *XnUaddressInfoperPdusessionItem) SetDataForwardingInfoFromTargetNgrannode(dataForwardingInfoFromTargetNgrannode *DataForwardingInfoFromTargetNgrannode) *XnUaddressInfoperPdusessionItem {
	m.DataForwardingInfoFromTargetNgrannode = dataForwardingInfoFromTargetNgrannode
	return m
}

func (m *XnUaddressInfoperPdusessionItem) SetPduSessionResourceSetupCompleteInfoSnterm(pduSessionResourceSetupCompleteInfoSnterm *PdusessionResourceBearerSetupCompleteInfoSNterminated) *XnUaddressInfoperPdusessionItem {
	m.PduSessionResourceSetupCompleteInfoSnterm = pduSessionResourceSetupCompleteInfoSnterm
	return m
}

func (m *DataForwardingInfoFromTargetNgrannode) SetPduSessionLevelDldataForwardingInfo(pduSessionLevelDldataForwardingInfo *UptransportLayerInformation) *DataForwardingInfoFromTargetNgrannode {
	m.PduSessionLevelDldataForwardingInfo = pduSessionLevelDldataForwardingInfo
	return m
}

func (m *DataForwardingInfoFromTargetNgrannode) SetPduSessionLevelUldataForwardingInfo(pduSessionLevelUldataForwardingInfo *UptransportLayerInformation) *DataForwardingInfoFromTargetNgrannode {
	m.PduSessionLevelUldataForwardingInfo = pduSessionLevelUldataForwardingInfo
	return m
}

func (m *DataForwardingInfoFromTargetNgrannode) SetDataForwardingResponseDrbitemList(dataForwardingResponseDrbitemList *DataForwardingResponseDrbitemList) *DataForwardingInfoFromTargetNgrannode {
	m.DataForwardingResponseDrbitemList = dataForwardingResponseDrbitemList
	return m
}

func (m *DataforwardingandOffloadingInfofromSource) SetSourceDrbtoQoSflowMapping(sourceDrbtoQoSflowMapping *DrbtoQoSflowMappingList) *DataforwardingandOffloadingInfofromSource {
	m.SourceDrbtoQoSflowMapping = sourceDrbtoQoSflowMapping
	return m
}

func (m *DataForwardingResponseDrbitem) SetDlForwardingUptnl(dlForwardingUptnl *UptransportLayerInformation) *DataForwardingResponseDrbitem {
	m.DlForwardingUptnl = dlForwardingUptnl
	return m
}

func (m *DataForwardingResponseDrbitem) SetUlForwardingUptnl(ulForwardingUptnl *UptransportLayerInformation) *DataForwardingResponseDrbitem {
	m.UlForwardingUptnl = ulForwardingUptnl
	return m
}

func (m *DataTrafficResourceIndication) SetReservedSubframePattern(reservedSubframePattern *ReservedSubframePattern) *DataTrafficResourceIndication {
	m.ReservedSubframePattern = reservedSubframePattern
	return m
}

func (m *DrbListwithCauseItem) SetRLcMode(rLcMode Rlcmode) *DrbListwithCauseItem {
	m.RLcMode = &rLcMode
	return m
}

func (m *DrbbstatusTransfer12BitsSn) SetReceiveStatusofPdcpsdu(receiveStatusofPdcpsdu *asn1.BitString) *DrbbstatusTransfer12BitsSn {
	m.ReceiveStatusofPdcpsdu = receiveStatusofPdcpsdu
	return m
}

func (m *DrbbstatusTransfer18BitsSn) SetReceiveStatusofPdcpsdu(receiveStatusofPdcpsdu *asn1.BitString) *DrbbstatusTransfer18BitsSn {
	m.ReceiveStatusofPdcpsdu = receiveStatusofPdcpsdu
	return m
}

func (m *DrbtoQoSflowMappingItem) SetRLcMode(rLcMode Rlcmode) *DrbtoQoSflowMappingItem {
	m.RLcMode = &rLcMode
	return m
}

func (m *Dynamic5Qidescriptor) SetFiveQi(fiveQi *FiveQi) *Dynamic5Qidescriptor {
	m.FiveQi = fiveQi
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

func (m *Dynamic5Qidescriptor) SetMaximumDataBurstVolume(maximumDataBurstVolume *MaximumDataBurstVolume) *Dynamic5Qidescriptor {
	m.MaximumDataBurstVolume = maximumDataBurstVolume
	return m
}

func (m *EUTraprachconfiguration) SetPrachConfigIndex(prachConfigIndex int32) *EUTraprachconfiguration {
	m.PrachConfigIndex = &prachConfigIndex
	return m
}

func (m *ExpectedUeactivityBehaviour) SetExpectedActivityPeriod(expectedActivityPeriod *ExpectedActivityPeriod) *ExpectedUeactivityBehaviour {
	m.ExpectedActivityPeriod = expectedActivityPeriod
	return m
}

func (m *ExpectedUeactivityBehaviour) SetExpectedIDlePeriod(expectedIDlePeriod *ExpectedIdlePeriod) *ExpectedUeactivityBehaviour {
	m.ExpectedIdlePeriod = expectedIDlePeriod
	return m
}

func (m *ExpectedUeactivityBehaviour) SetSourceOfUeactivityBehaviourInformation(sourceOfUeactivityBehaviourInformation SourceOfUeactivityBehaviourInformation) *ExpectedUeactivityBehaviour {
	m.SourceOfUeactivityBehaviourInformation = &sourceOfUeactivityBehaviourInformation
	return m
}

func (m *ExpectedUebehaviour) SetExpectedUeactivityBehaviour(expectedUeactivityBehaviour *ExpectedUeactivityBehaviour) *ExpectedUebehaviour {
	m.ExpectedUeactivityBehaviour = expectedUeactivityBehaviour
	return m
}

func (m *ExpectedUebehaviour) SetExpectedHointerval(expectedHointerval ExpectedHointerval) *ExpectedUebehaviour {
	m.ExpectedHointerval = &expectedHointerval
	return m
}

func (m *ExpectedUebehaviour) SetExpectedUemobility(expectedUemobility ExpectedUemobility) *ExpectedUebehaviour {
	m.ExpectedUemobility = &expectedUemobility
	return m
}

func (m *ExpectedUebehaviour) SetExpectedUemovingTrajectory(expectedUemovingTrajectory *ExpectedUemovingTrajectory) *ExpectedUebehaviour {
	m.ExpectedUemovingTrajectory = expectedUemovingTrajectory
	return m
}

func (m *ExpectedUemovingTrajectoryItem) SetTimeStayedInCell(timeStayedInCell int32) *ExpectedUemovingTrajectoryItem {
	m.TimeStayedInCell = &timeStayedInCell
	return m
}

func (m *ExtTlaItem) SetIPsecTla(iPsecTla *TransportLayerAddress) *ExtTlaItem {
	m.IPsecTla = iPsecTla
	return m
}

func (m *ExtTlaItem) SetGTptransportLayerAddresses(gTptransportLayerAddresses *Gtptlas) *ExtTlaItem {
	m.GTptransportLayerAddresses = gTptransportLayerAddresses
	return m
}

func (m *GbrqoSflowInfo) SetNotificationControl(notificationControl NotificationControlGbrqoSflowInfo) *GbrqoSflowInfo {
	m.NotificationControl = &notificationControl
	return m
}

func (m *GbrqoSflowInfo) SetMaxPacketLossRateDl(maxPacketLossRateDl *PacketLossRate) *GbrqoSflowInfo {
	m.MaxPacketLossRateDl = maxPacketLossRateDl
	return m
}

func (m *GbrqoSflowInfo) SetMaxPacketLossRateUl(maxPacketLossRateUl *PacketLossRate) *GbrqoSflowInfo {
	m.MaxPacketLossRateUl = maxPacketLossRateUl
	return m
}

func (m *ImmediateMdtNR) SetM1Configuration(m1Configuration *M1Configuration) *ImmediateMdtNR {
	m.M1Configuration = m1Configuration
	return m
}

func (m *ImmediateMdtNR) SetM4Configuration(m4Configuration *M4Configuration) *ImmediateMdtNR {
	m.M4Configuration = m4Configuration
	return m
}

func (m *ImmediateMdtNR) SetM5Configuration(m5Configuration *M5Configuration) *ImmediateMdtNR {
	m.M5Configuration = m5Configuration
	return m
}

func (m *ImmediateMdtNR) SetMDtLocationInfo(mDtLocationInfo *MdtLocationInfo) *ImmediateMdtNR {
	m.MDtLocationInfo = mDtLocationInfo
	return m
}

func (m *ImmediateMdtNR) SetM6Configuration(m6Configuration *M6Configuration) *ImmediateMdtNR {
	m.M6Configuration = m6Configuration
	return m
}

func (m *ImmediateMdtNR) SetM7Configuration(m7Configuration *M7Configuration) *ImmediateMdtNR {
	m.M7Configuration = m7Configuration
	return m
}

func (m *ImmediateMdtNR) SetBluetoothMeasurementConfiguration(bluetoothMeasurementConfiguration *BluetoothMeasurementConfiguration) *ImmediateMdtNR {
	m.BluetoothMeasurementConfiguration = bluetoothMeasurementConfiguration
	return m
}

func (m *ImmediateMdtNR) SetWLanmeasurementConfiguration(wLanmeasurementConfiguration *WlanmeasurementConfiguration) *ImmediateMdtNR {
	m.WLanmeasurementConfiguration = wLanmeasurementConfiguration
	return m
}

func (m *ImmediateMdtNR) SetSensorMeasurementConfiguration(sensorMeasurementConfiguration *SensorMeasurementConfiguration) *ImmediateMdtNR {
	m.SensorMeasurementConfiguration = sensorMeasurementConfiguration
	return m
}

func (m *LocationReportingInformation) SetAreaOfInterest(areaOfInterest *AreaOfInterestInformation) *LocationReportingInformation {
	m.AreaOfInterest = areaOfInterest
	return m
}

func (m *LoggedMdtEUtra) SetBluetoothMeasurementConfiguration(bluetoothMeasurementConfiguration *BluetoothMeasurementConfiguration) *LoggedMdtEUtra {
	m.BluetoothMeasurementConfiguration = bluetoothMeasurementConfiguration
	return m
}

func (m *LoggedMdtEUtra) SetWLanmeasurementConfiguration(wLanmeasurementConfiguration *WlanmeasurementConfiguration) *LoggedMdtEUtra {
	m.WLanmeasurementConfiguration = wLanmeasurementConfiguration
	return m
}

func (m *LoggedMdtNR) SetBluetoothMeasurementConfiguration(bluetoothMeasurementConfiguration *BluetoothMeasurementConfiguration) *LoggedMdtNR {
	m.BluetoothMeasurementConfiguration = bluetoothMeasurementConfiguration
	return m
}

func (m *LoggedMdtNR) SetWLanmeasurementConfiguration(wLanmeasurementConfiguration *WlanmeasurementConfiguration) *LoggedMdtNR {
	m.WLanmeasurementConfiguration = wLanmeasurementConfiguration
	return m
}

func (m *LoggedMdtNR) SetSensorMeasurementConfiguration(sensorMeasurementConfiguration *SensorMeasurementConfiguration) *LoggedMdtNR {
	m.SensorMeasurementConfiguration = sensorMeasurementConfiguration
	return m
}

func (m *LoggedMdtNR) SetAreaScopeOfNeighCellsList(areaScopeOfNeighCellsList *AreaScopeOfNeighCellsList) *LoggedMdtNR {
	m.AreaScopeOfNeighCellsList = areaScopeOfNeighCellsList
	return m
}

func (m *Ltev2XservicesAuthorized) SetVehicleUe(vehicleUe VehicleUe) *Ltev2XservicesAuthorized {
	m.VehicleUe = &vehicleUe
	return m
}

func (m *Ltev2XservicesAuthorized) SetPedestrianUe(pedestrianUe PedestrianUe) *Ltev2XservicesAuthorized {
	m.PedestrianUe = &pedestrianUe
	return m
}

func (m *M1Configuration) SetM1ThresholdeventA2(m1ThresholdeventA2 *M1ThresholdEventA2) *M1Configuration {
	m.M1ThresholdeventA2 = m1ThresholdeventA2
	return m
}

func (m *M1Configuration) SetM1PeriodicReporting(m1PeriodicReporting *M1PeriodicReporting) *M1Configuration {
	m.M1PeriodicReporting = m1PeriodicReporting
	return m
}

func (m *MdtConfiguration) SetMDtConfigurationNr(mDtConfigurationNr *MdtConfigurationNR) *MdtConfiguration {
	m.MDtConfigurationNr = mDtConfigurationNr
	return m
}

func (m *MdtConfiguration) SetMDtConfigurationEutra(mDtConfigurationEutra *MdtConfigurationEUtra) *MdtConfiguration {
	m.MDtConfigurationEutra = mDtConfigurationEutra
	return m
}

func (m *MdtConfigurationNR) SetAreaScopeOfMdtNr(areaScopeOfMdtNr *AreaScopeOfMdtNR) *MdtConfigurationNR {
	m.AreaScopeOfMdtNr = areaScopeOfMdtNr
	return m
}

func (m *MdtConfigurationEUtra) SetAreaScopeOfMdtEutra(areaScopeOfMdtEutra *AreaScopeOfMdtEUtra) *MdtConfigurationEUtra {
	m.AreaScopeOfMdtEutra = areaScopeOfMdtEutra
	return m
}

func (m *MobilityRestrictionList) SetEquivalentPlmns(equivalentPlmns []*PlmnIdentity) *MobilityRestrictionList {
	m.EquivalentPlmns = equivalentPlmns
	return m
}

func (m *MobilityRestrictionList) SetRatRestrictions(ratRestrictions *RatRestrictionsList) *MobilityRestrictionList {
	m.RatRestrictions = ratRestrictions
	return m
}

func (m *MobilityRestrictionList) SetForbIDdenAreaInformation(forbIDdenAreaInformation *ForbiddenAreaList) *MobilityRestrictionList {
	m.ForbiddenAreaInformation = forbIDdenAreaInformation
	return m
}

func (m *MobilityRestrictionList) SetServiceAreaInformation(serviceAreaInformation *ServiceAreaList) *MobilityRestrictionList {
	m.ServiceAreaInformation = serviceAreaInformation
	return m
}

func (m *ServiceAreaItem) SetAllowedTacsServiceArea(allowedTacsServiceArea []*Tac) *ServiceAreaItem {
	m.AllowedTacsServiceArea = allowedTacsServiceArea
	return m
}

func (m *ServiceAreaItem) SetNotAllowedTacsServiceArea(notAllowedTacsServiceArea []*Tac) *ServiceAreaItem {
	m.NotAllowedTacsServiceArea = notAllowedTacsServiceArea
	return m
}

func (m *EUTraResourceCoordinationInfo) SetDlCoordinationInfo(dlCoordinationInfo *asn1.BitString) *EUTraResourceCoordinationInfo {
	m.DlCoordinationInfo = dlCoordinationInfo
	return m
}

func (m *EUTraResourceCoordinationInfo) SetNrCell(nrCell *NrCGi) *EUTraResourceCoordinationInfo {
	m.NrCell = nrCell
	return m
}

func (m *EUTraResourceCoordinationInfo) SetEUtraCoordinationAssistanceInfo(eUtraCoordinationAssistanceInfo EUTraCoordinationAssistanceInfo) *EUTraResourceCoordinationInfo {
	m.EUtraCoordinationAssistanceInfo = &eUtraCoordinationAssistanceInfo
	return m
}

func (m *NrResourceCoordinationInfo) SetDlCoordinationInfo(dlCoordinationInfo *asn1.BitString) *NrResourceCoordinationInfo {
	m.DlCoordinationInfo = dlCoordinationInfo
	return m
}

func (m *NrResourceCoordinationInfo) SetEUtraCell(eUtraCell *EUTraCGi) *NrResourceCoordinationInfo {
	m.EUtraCell = eUtraCell
	return m
}

func (m *NrResourceCoordinationInfo) SetNrCoordinationAssistanceInfo(nrCoordinationAssistanceInfo NrCoordinationAssistanceInfo) *NrResourceCoordinationInfo {
	m.NrCoordinationAssistanceInfo = &nrCoordinationAssistanceInfo
	return m
}

func (m *NeighbourInformationEUTraItem) SetRanac(ranac *Ranac) *NeighbourInformationEUTraItem {
	m.Ranac = ranac
	return m
}
func (m *NeighbourInformationNRItem) SetRanac(ranac *Ranac) *NeighbourInformationNRItem {
	m.Ranac = ranac
	return m
}

func (m *NonDynamic5Qidescriptor) SetPriorityLevelQoS(priorityLevelQoS *PriorityLevelQoS) *NonDynamic5Qidescriptor {
	m.PriorityLevelQoS = priorityLevelQoS
	return m
}

func (m *NonDynamic5Qidescriptor) SetAveragingWindow(averagingWindow *AveragingWindow) *NonDynamic5Qidescriptor {
	m.AveragingWindow = averagingWindow
	return m
}

func (m *NonDynamic5Qidescriptor) SetMaximumDataBurstVolume(maximumDataBurstVolume *MaximumDataBurstVolume) *NonDynamic5Qidescriptor {
	m.MaximumDataBurstVolume = maximumDataBurstVolume
	return m
}

func (m *NprachconfigurationFDd) SetAnchorCarrierEdtNprachconfig(anchorCarrierEdtNprachconfig []byte) *NprachconfigurationFDd {
	m.AnchorCarrierEdtNprachconfig = anchorCarrierEdtNprachconfig
	return m
}

func (m *NprachconfigurationFDd) SetAnchorCarrierFormat2Nprachconfig(anchorCarrierFormat2Nprachconfig []byte) *NprachconfigurationFDd {
	m.AnchorCarrierFormat2Nprachconfig = anchorCarrierFormat2Nprachconfig
	return m
}

func (m *NprachconfigurationFDd) SetAnchorCarrierFormat2EdtNprachconfig(anchorCarrierFormat2EdtNprachconfig []byte) *NprachconfigurationFDd {
	m.AnchorCarrierFormat2EdtNprachconfig = anchorCarrierFormat2EdtNprachconfig
	return m
}

func (m *NprachconfigurationFDd) SetNonAnchorCarrierNprachconfig(nonAnchorCarrierNprachconfig []byte) *NprachconfigurationFDd {
	m.NonAnchorCarrierNprachconfig = nonAnchorCarrierNprachconfig
	return m
}

func (m *NprachconfigurationFDd) SetNonAnchorCarrierFormat2Nprachconfig(nonAnchorCarrierFormat2Nprachconfig []byte) *NprachconfigurationFDd {
	m.NonAnchorCarrierFormat2Nprachconfig = nonAnchorCarrierFormat2Nprachconfig
	return m
}

func (m *NprachconfigurationTDd) SetNonAnchorCarrierFequencyConfiglist(nonAnchorCarrierFequencyConfiglist *NonAnchorCarrierFrequencylist) *NprachconfigurationTDd {
	m.NonAnchorCarrierFequencyConfiglist = nonAnchorCarrierFequencyConfiglist
	return m
}

func (m *NprachconfigurationTDd) SetNonAnchorCarrierNprachconfigTdd(nonAnchorCarrierNprachconfigTdd []byte) *NprachconfigurationTDd {
	m.NonAnchorCarrierNprachconfigTdd = nonAnchorCarrierNprachconfigTdd
	return m
}

func (m *NrfrequencyBandItem) SetSupportedSulBandList(supportedSulBandList *SupportedSulbandList) *NrfrequencyBandItem {
	m.SupportedSulBandList = supportedSulBandList
	return m
}

func (m *NrfrequencyInfo) SetSulInformation(sulInformation *SulInformation) *NrfrequencyInfo {
	m.SulInformation = sulInformation
	return m
}

func (m *Nrv2XservicesAuthorized) SetVehicleUe(vehicleUe VehicleUe) *Nrv2XservicesAuthorized {
	m.VehicleUe = &vehicleUe
	return m
}

func (m *Nrv2XservicesAuthorized) SetPedestrianUe(pedestrianUe PedestrianUe) *Nrv2XservicesAuthorized {
	m.PedestrianUe = &pedestrianUe
	return m
}

func (m *PagingeDrxinformation) SetPagingTimeWindow(pagingTimeWindow PagingTimeWindow) *PagingeDrxinformation {
	m.PagingTimeWindow = &pagingTimeWindow
	return m
}

func (m *Pc5QoSparameters) SetPc5LinkAggregateBitRates(pc5LinkAggregateBitRates *BitRate) *Pc5QoSparameters {
	m.Pc5LinkAggregateBitRates = pc5LinkAggregateBitRates
	return m
}

func (m *Pc5QoSflowItem) SetPc5FlowBitRates(pc5FlowBitRates *Pc5FlowBitRates) *Pc5QoSflowItem {
	m.Pc5FlowBitRates = pc5FlowBitRates
	return m
}

func (m *Pc5QoSflowItem) SetRange(rang Range) *Pc5QoSflowItem {
	m.Range = &rang
	return m
}

func (m *PdusessionListwithCauseItem) SetCause(cause *Cause) *PdusessionListwithCauseItem {
	m.Cause = cause
	return m
}

func (m *PdusessionListwithDataForwardingRequestItem) SetDataforwardingInfofromSource(dataforwardingInfofromSource *DataforwardingandOffloadingInfofromSource) *PdusessionListwithDataForwardingRequestItem {
	m.DataforwardingInfofromSource = dataforwardingInfofromSource
	return m
}

func (m *PdusessionListwithDataForwardingRequestItem) SetDRbtoBeReleasedList(dRbtoBeReleasedList *DrbtoQoSflowMappingList) *PdusessionListwithDataForwardingRequestItem {
	m.DRbtoBeReleasedList = dRbtoBeReleasedList
	return m
}

func (m *PdusessionResourceAdmittedInfo) SetDLNgUTnlInformationUnchanged(dLNgUTnlInformationUnchanged DlngutnlinformationUnchangedPdusessionResourceAdmittedInfo) *PdusessionResourceAdmittedInfo {
	m.DLNgUTnlInformationUnchanged = &dLNgUTnlInformationUnchanged
	return m
}

func (m *PdusessionResourceAdmittedInfo) SetQosFlowsNotAdmittedList(qosFlowsNotAdmittedList *QoSflowsListwithCause) *PdusessionResourceAdmittedInfo {
	m.QosFlowsNotAdmittedList = qosFlowsNotAdmittedList
	return m
}

func (m *PdusessionResourceAdmittedInfo) SetDataForwardingInfoFromTarget(dataForwardingInfoFromTarget *DataForwardingInfoFromTargetNgrannode) *PdusessionResourceAdmittedInfo {
	m.DataForwardingInfoFromTarget = dataForwardingInfoFromTarget
	return m
}

func (m *PdusessionResourcesNotAdmittedItem) SetCause(cause *Cause) *PdusessionResourcesNotAdmittedItem {
	m.Cause = cause
	return m
}

func (m *PdusessionResourcesToBeSetupItem) SetPduSessionAmbr(pduSessionAmbr *PdusessionAggregateMaximumBitRate) *PdusessionResourcesToBeSetupItem {
	m.PduSessionAmbr = pduSessionAmbr
	return m
}

func (m *PdusessionResourcesToBeSetupItem) SetSourceDlNgUTnlInformation(sourceDlNgUTnlInformation *UptransportLayerInformation) *PdusessionResourcesToBeSetupItem {
	m.SourceDlNgUTnlInformation = sourceDlNgUTnlInformation
	return m
}

func (m *PdusessionResourcesToBeSetupItem) SetSecurityIndication(securityIndication *SecurityIndication) *PdusessionResourcesToBeSetupItem {
	m.SecurityIndication = securityIndication
	return m
}

func (m *PdusessionResourcesToBeSetupItem) SetPduSessionNetworkInstance(pduSessionNetworkInstance *PdusessionNetworkInstance) *PdusessionResourcesToBeSetupItem {
	m.PduSessionNetworkInstance = pduSessionNetworkInstance
	return m
}

func (m *PdusessionResourcesToBeSetupItem) SetDataforwardinginfofromSource(dataforwardinginfofromSource *DataforwardingandOffloadingInfofromSource) *PdusessionResourcesToBeSetupItem {
	m.DataforwardinginfofromSource = dataforwardinginfofromSource
	return m
}

func (m *PdusessionResourceSetupInfoSNterminated) SetPduSessionNetworkInstance(pduSessionNetworkInstance *PdusessionNetworkInstance) *PdusessionResourceSetupInfoSNterminated {
	m.PduSessionNetworkInstance = pduSessionNetworkInstance
	return m
}

func (m *PdusessionResourceSetupInfoSNterminated) SetDataforwardinginfofromSource(dataforwardinginfofromSource *DataforwardingandOffloadingInfofromSource) *PdusessionResourceSetupInfoSNterminated {
	m.DataforwardinginfofromSource = dataforwardinginfofromSource
	return m
}

func (m *PdusessionResourceSetupInfoSNterminated) SetSecurityIndication(securityIndication *SecurityIndication) *PdusessionResourceSetupInfoSNterminated {
	m.SecurityIndication = securityIndication
	return m
}

func (m *QoSflowsToBeSetupListSetupSNterminatedItem) SetOfferedGbrqoSflowInfo(offeredGbrqoSflowInfo *GbrqoSflowInfo) *QoSflowsToBeSetupListSetupSNterminatedItem {
	m.OfferedGbrqoSflowInfo = offeredGbrqoSflowInfo
	return m
}

func (m *PdusessionResourceSetupResponseInfoSNterminated) SetDRbsToBeSetup(dRbsToBeSetup *DrbsToBeSetupListSetupResponseSNterminated) *PdusessionResourceSetupResponseInfoSNterminated {
	m.DRbsToBeSetup = dRbsToBeSetup
	return m
}

func (m *PdusessionResourceSetupResponseInfoSNterminated) SetDataforwardinginfoTarget(dataforwardinginfoTarget *DataForwardingInfoFromTargetNgrannode) *PdusessionResourceSetupResponseInfoSNterminated {
	m.DataforwardinginfoTarget = dataforwardinginfoTarget
	return m
}

func (m *PdusessionResourceSetupResponseInfoSNterminated) SetQosFlowsNotAdmittedList(qosFlowsNotAdmittedList *QoSflowsListwithCause) *PdusessionResourceSetupResponseInfoSNterminated {
	m.QosFlowsNotAdmittedList = qosFlowsNotAdmittedList
	return m
}

func (m *PdusessionResourceSetupResponseInfoSNterminated) SetSecurityResult(securityResult *SecurityResult) *PdusessionResourceSetupResponseInfoSNterminated {
	m.SecurityResult = securityResult
	return m
}

func (m *DrbsToBeSetupListSetupResponseSNterminatedItem) SetPDcpSnlength(pDcpSnlength *Pdcpsnlength) *DrbsToBeSetupListSetupResponseSNterminatedItem {
	m.PDcpSnlength = pDcpSnlength
	return m
}

func (m *DrbsToBeSetupListSetupResponseSNterminatedItem) SetULConfiguration(uLConfiguration *Ulconfiguration) *DrbsToBeSetupListSetupResponseSNterminatedItem {
	m.ULConfiguration = uLConfiguration
	return m
}

func (m *DrbsToBeSetupListSetupResponseSNterminatedItem) SetSecondarySnUlPdcpUpTnlinfo(secondarySnUlPdcpUpTnlinfo *UptransportParameters) *DrbsToBeSetupListSetupResponseSNterminatedItem {
	m.SecondarySnUlPdcpUpTnlinfo = secondarySnUlPdcpUpTnlinfo
	return m
}

func (m *DrbsToBeSetupListSetupResponseSNterminatedItem) SetDuplicationActivation(duplicationActivation DuplicationActivation) *DrbsToBeSetupListSetupResponseSNterminatedItem {
	m.DuplicationActivation = &duplicationActivation
	return m
}

func (m *QoSflowsMappedtoDrbSetupResponseSNterminatedItem) SetMCgrequestedGbrqoSflowInfo(mCgrequestedGbrqoSflowInfo *GbrqoSflowInfo) *QoSflowsMappedtoDrbSetupResponseSNterminatedItem {
	m.MCgrequestedGbrqoSflowInfo = mCgrequestedGbrqoSflowInfo
	return m
}

func (m *QoSflowsMappedtoDrbSetupResponseSNterminatedItem) SetQosFlowMappingIndication(qosFlowMappingIndication QoSflowMappingIndication) *QoSflowsMappedtoDrbSetupResponseSNterminatedItem {
	m.QosFlowMappingIndication = &qosFlowMappingIndication
	return m
}

func (m *DrbsToBeSetupListSetupMNterminatedItem) SetULConfiguration(uLConfiguration *Ulconfiguration) *DrbsToBeSetupListSetupMNterminatedItem {
	m.ULConfiguration = uLConfiguration
	return m
}

func (m *DrbsToBeSetupListSetupMNterminatedItem) SetPDcpSnlength(pDcpSnlength *Pdcpsnlength) *DrbsToBeSetupListSetupMNterminatedItem {
	m.PDcpSnlength = pDcpSnlength
	return m
}

func (m *DrbsToBeSetupListSetupMNterminatedItem) SetSecondaryMnUlPdcpUpTnlinfo(secondaryMnUlPdcpUpTnlinfo *UptransportParameters) *DrbsToBeSetupListSetupMNterminatedItem {
	m.SecondaryMnUlPdcpUpTnlinfo = secondaryMnUlPdcpUpTnlinfo
	return m
}

func (m *DrbsToBeSetupListSetupMNterminatedItem) SetDuplicationActivation(duplicationActivation DuplicationActivation) *DrbsToBeSetupListSetupMNterminatedItem {
	m.DuplicationActivation = &duplicationActivation
	return m
}

func (m *QoSflowsMappedtoDrbSetupMNterminatedItem) SetQosFlowMappingIndication(qosFlowMappingIndication QoSflowMappingIndication) *QoSflowsMappedtoDrbSetupMNterminatedItem {
	m.QosFlowMappingIndication = &qosFlowMappingIndication
	return m
}

func (m *DrbsAdmittedListSetupResponseMNterminatedItem) SetSecondarySnDlScgUpTnlinfo(secondarySnDlScgUpTnlinfo *UptransportParameters) *DrbsAdmittedListSetupResponseMNterminatedItem {
	m.SecondarySnDlScgUpTnlinfo = secondarySnDlScgUpTnlinfo
	return m
}

func (m *DrbsAdmittedListSetupResponseMNterminatedItem) SetLCID(lCID *Lcid) *DrbsAdmittedListSetupResponseMNterminatedItem {
	m.LCid = lCID
	return m
}

func (m *PdusessionResourceModificationInfoSNterminated) SetULNgUTnlatUpf(uLNgUTnlatUpf *UptransportLayerInformation) *PdusessionResourceModificationInfoSNterminated {
	m.ULNgUTnlatUpf = uLNgUTnlatUpf
	return m
}

func (m *PdusessionResourceModificationInfoSNterminated) SetPduSessionNetworkInstance(pduSessionNetworkInstance *PdusessionNetworkInstance) *PdusessionResourceModificationInfoSNterminated {
	m.PduSessionNetworkInstance = pduSessionNetworkInstance
	return m
}

func (m *PdusessionResourceModificationInfoSNterminated) SetQosFlowsToBeSetupList(qosFlowsToBeSetupList *QoSflowsToBeSetupListSetupSNterminated) *PdusessionResourceModificationInfoSNterminated {
	m.QosFlowsToBeSetupList = qosFlowsToBeSetupList
	return m
}

func (m *PdusessionResourceModificationInfoSNterminated) SetDataforwardinginfofromSource(dataforwardinginfofromSource *DataforwardingandOffloadingInfofromSource) *PdusessionResourceModificationInfoSNterminated {
	m.DataforwardinginfofromSource = dataforwardinginfofromSource
	return m
}

func (m *PdusessionResourceModificationInfoSNterminated) SetQosFlowsToBeModifiedList(qosFlowsToBeModifiedList *QoSflowsToBeSetupListModifiedSNterminated) *PdusessionResourceModificationInfoSNterminated {
	m.QosFlowsToBeModifiedList = qosFlowsToBeModifiedList
	return m
}

func (m *PdusessionResourceModificationInfoSNterminated) SetQoSflowsToBeReleasedList(qoSflowsToBeReleasedList *QoSflowsListwithCause) *PdusessionResourceModificationInfoSNterminated {
	m.QoSflowsToBeReleasedList = qoSflowsToBeReleasedList
	return m
}

func (m *PdusessionResourceModificationInfoSNterminated) SetDrbsToBeModifiedList(drbsToBeModifiedList *DrbsToBeModifiedListModifiedSNterminated) *PdusessionResourceModificationInfoSNterminated {
	m.DrbsToBeModifiedList = drbsToBeModifiedList
	return m
}

func (m *PdusessionResourceModificationInfoSNterminated) SetDRbsToBeReleased(dRbsToBeReleased *DrbListwithCause) *PdusessionResourceModificationInfoSNterminated {
	m.DRbsToBeReleased = dRbsToBeReleased
	return m
}

func (m *QoSflowsToBeSetupListModifiedSNterminatedItem) SetQosFlowLevelQoSparameters(qosFlowLevelQoSparameters *QoSflowLevelQoSparameters) *QoSflowsToBeSetupListModifiedSNterminatedItem {
	m.QosFlowLevelQoSparameters = qosFlowLevelQoSparameters
	return m
}

func (m *QoSflowsToBeSetupListModifiedSNterminatedItem) SetOfferedGbrqoSflowInfo(offeredGbrqoSflowInfo *GbrqoSflowInfo) *QoSflowsToBeSetupListModifiedSNterminatedItem {
	m.OfferedGbrqoSflowInfo = offeredGbrqoSflowInfo
	return m
}

func (m *QoSflowsToBeSetupListModifiedSNterminatedItem) SetQosFlowMappingIndication(qosFlowMappingIndication QoSflowMappingIndication) *QoSflowsToBeSetupListModifiedSNterminatedItem {
	m.QosFlowMappingIndication = &qosFlowMappingIndication
	return m
}

func (m *DrbsToBeModifiedListModifiedSNterminatedItem) SetMNDlScgUpTnlinfo(mNDlScgUpTnlinfo *UptransportParameters) *DrbsToBeModifiedListModifiedSNterminatedItem {
	m.MNDlScgUpTnlinfo = mNDlScgUpTnlinfo
	return m
}

func (m *DrbsToBeModifiedListModifiedSNterminatedItem) SetSecondaryMnDlScgUpTnlinfo(secondaryMnDlScgUpTnlinfo *UptransportParameters) *DrbsToBeModifiedListModifiedSNterminatedItem {
	m.SecondaryMnDlScgUpTnlinfo = secondaryMnDlScgUpTnlinfo
	return m
}

func (m *DrbsToBeModifiedListModifiedSNterminatedItem) SetLCID(lCID *Lcid) *DrbsToBeModifiedListModifiedSNterminatedItem {
	m.LCid = lCID
	return m
}

func (m *DrbsToBeModifiedListModifiedSNterminatedItem) SetRlcStatus(rlcStatus *RlcStatus) *DrbsToBeModifiedListModifiedSNterminatedItem {
	m.RlcStatus = rlcStatus
	return m
}

func (m *PdusessionResourceModificationResponseInfoSNterminated) SetDLNgUTnlatNgRan(dLNgUTnlatNgRan *UptransportLayerInformation) *PdusessionResourceModificationResponseInfoSNterminated {
	m.DLNgUTnlatNgRan = dLNgUTnlatNgRan
	return m
}

func (m *PdusessionResourceModificationResponseInfoSNterminated) SetDRbsToBeSetup(dRbsToBeSetup *DrbsToBeSetupListSetupResponseSNterminated) *PdusessionResourceModificationResponseInfoSNterminated {
	m.DRbsToBeSetup = dRbsToBeSetup
	return m
}

func (m *PdusessionResourceModificationResponseInfoSNterminated) SetDataforwardinginfoTarget(dataforwardinginfoTarget *DataForwardingInfoFromTargetNgrannode) *PdusessionResourceModificationResponseInfoSNterminated {
	m.DataforwardinginfoTarget = dataforwardinginfoTarget
	return m
}

func (m *PdusessionResourceModificationResponseInfoSNterminated) SetDRbsToBeModified(dRbsToBeModified *DrbsToBeModifiedListModificationResponseSNterminated) *PdusessionResourceModificationResponseInfoSNterminated {
	m.DRbsToBeModified = dRbsToBeModified
	return m
}

func (m *PdusessionResourceModificationResponseInfoSNterminated) SetDRbsToBeReleased(dRbsToBeReleased *DrbListwithCause) *PdusessionResourceModificationResponseInfoSNterminated {
	m.DRbsToBeReleased = dRbsToBeReleased
	return m
}

func (m *PdusessionResourceModificationResponseInfoSNterminated) SetDataforwardinginfofromSource(dataforwardinginfofromSource *DataforwardingandOffloadingInfofromSource) *PdusessionResourceModificationResponseInfoSNterminated {
	m.DataforwardinginfofromSource = dataforwardinginfofromSource
	return m
}

func (m *PdusessionResourceModificationResponseInfoSNterminated) SetQosFlowsNotAdmittedTbadded(qosFlowsNotAdmittedTbadded *QoSflowsListwithCause) *PdusessionResourceModificationResponseInfoSNterminated {
	m.QosFlowsNotAdmittedTbadded = qosFlowsNotAdmittedTbadded
	return m
}

func (m *PdusessionResourceModificationResponseInfoSNterminated) SetQosFlowsReleased(qosFlowsReleased *QoSflowsListwithCause) *PdusessionResourceModificationResponseInfoSNterminated {
	m.QosFlowsReleased = qosFlowsReleased
	return m
}

func (m *DrbsToBeModifiedListModificationResponseSNterminatedItem) SetSNUlPdcpUpTnlinfo(sNUlPdcpUpTnlinfo *UptransportParameters) *DrbsToBeModifiedListModificationResponseSNterminatedItem {
	m.SNUlPdcpUpTnlinfo = sNUlPdcpUpTnlinfo
	return m
}

func (m *DrbsToBeModifiedListModificationResponseSNterminatedItem) SetDRbQoS(dRbQoS *QoSflowLevelQoSparameters) *DrbsToBeModifiedListModificationResponseSNterminatedItem {
	m.DRbQoS = dRbQoS
	return m
}

func (m *DrbsToBeModifiedListModificationResponseSNterminatedItem) SetQoSflowsMappedtoDrbSetupResponseSnterminated(qoSflowsMappedtoDrbSetupResponseSnterminated *QoSflowsMappedtoDrbSetupResponseSNterminated) *DrbsToBeModifiedListModificationResponseSNterminatedItem {
	m.QoSflowsMappedtoDrbSetupResponseSnterminated = qoSflowsMappedtoDrbSetupResponseSnterminated
	return m
}

func (m *PdusessionResourceModificationInfoMNterminated) SetDRbsToBeSetup(dRbsToBeSetup *DrbsToBeSetupListSetupMNterminated) *PdusessionResourceModificationInfoMNterminated {
	m.DRbsToBeSetup = dRbsToBeSetup
	return m
}

func (m *PdusessionResourceModificationInfoMNterminated) SetDRbsToBeModified(dRbsToBeModified *DrbsToBeModifiedListModificationMNterminated) *PdusessionResourceModificationInfoMNterminated {
	m.DRbsToBeModified = dRbsToBeModified
	return m
}

func (m *PdusessionResourceModificationInfoMNterminated) SetDRbsToBeReleased(dRbsToBeReleased *DrbListwithCause) *PdusessionResourceModificationInfoMNterminated {
	m.DRbsToBeReleased = dRbsToBeReleased
	return m
}

func (m *DrbsToBeModifiedListModificationMNterminatedItem) SetMNUlPdcpUpTnlinfo(mNUlPdcpUpTnlinfo *UptransportParameters) *DrbsToBeModifiedListModificationMNterminatedItem {
	m.MNUlPdcpUpTnlinfo = mNUlPdcpUpTnlinfo
	return m
}

func (m *DrbsToBeModifiedListModificationMNterminatedItem) SetDRbQoS(dRbQoS *QoSflowLevelQoSparameters) *DrbsToBeModifiedListModificationMNterminatedItem {
	m.DRbQoS = dRbQoS
	return m
}

func (m *DrbsToBeModifiedListModificationMNterminatedItem) SetSecondaryMnUlPdcpUpTnlinfo(secondaryMnUlPdcpUpTnlinfo *UptransportParameters) *DrbsToBeModifiedListModificationMNterminatedItem {
	m.SecondaryMnUlPdcpUpTnlinfo = secondaryMnUlPdcpUpTnlinfo
	return m
}

func (m *DrbsToBeModifiedListModificationMNterminatedItem) SetULConfiguration(uLConfiguration *Ulconfiguration) *DrbsToBeModifiedListModificationMNterminatedItem {
	m.ULConfiguration = uLConfiguration
	return m
}

func (m *DrbsToBeModifiedListModificationMNterminatedItem) SetPdcpDuplicationConfiguration(pdcpDuplicationConfiguration PdcpduplicationConfiguration) *DrbsToBeModifiedListModificationMNterminatedItem {
	m.PdcpDuplicationConfiguration = &pdcpDuplicationConfiguration
	return m
}

func (m *DrbsToBeModifiedListModificationMNterminatedItem) SetDuplicationActivation(duplicationActivation DuplicationActivation) *DrbsToBeModifiedListModificationMNterminatedItem {
	m.DuplicationActivation = &duplicationActivation
	return m
}

func (m *DrbsToBeModifiedListModificationMNterminatedItem) SetQoSflowsMappedtoDrbSetupMnterminated(qoSflowsMappedtoDrbSetupMnterminated *QoSflowsMappedtoDrbSetupMNterminated) *DrbsToBeModifiedListModificationMNterminatedItem {
	m.QoSflowsMappedtoDrbSetupMnterminated = qoSflowsMappedtoDrbSetupMnterminated
	return m
}

func (m *PdusessionResourceModificationResponseInfoMNterminated) SetDRbsReleasedList(dRbsReleasedList *DrbList) *PdusessionResourceModificationResponseInfoMNterminated {
	m.DRbsReleasedList = dRbsReleasedList
	return m
}

func (m *PdusessionResourceModificationResponseInfoMNterminated) SetDRbsNotAdmittedSetupModifyList(dRbsNotAdmittedSetupModifyList *DrbListwithCause) *PdusessionResourceModificationResponseInfoMNterminated {
	m.DRbsNotAdmittedSetupModifyList = dRbsNotAdmittedSetupModifyList
	return m
}

func (m *DrbsAdmittedListModificationResponseMNterminatedItem) SetSNDlScgUpTnlinfo(sNDlScgUpTnlinfo *UptransportParameters) *DrbsAdmittedListModificationResponseMNterminatedItem {
	m.SNDlScgUpTnlinfo = sNDlScgUpTnlinfo
	return m
}

func (m *DrbsAdmittedListModificationResponseMNterminatedItem) SetSecondarySnDlScgUpTnlinfo(secondarySnDlScgUpTnlinfo *UptransportParameters) *DrbsAdmittedListModificationResponseMNterminatedItem {
	m.SecondarySnDlScgUpTnlinfo = secondarySnDlScgUpTnlinfo
	return m
}

func (m *DrbsAdmittedListModificationResponseMNterminatedItem) SetLCID(lCID *Lcid) *DrbsAdmittedListModificationResponseMNterminatedItem {
	m.LCid = lCID
	return m
}

func (m *PdusessionResourceChangeRequiredInfoSNterminated) SetDataforwardinginfofromSource(dataforwardinginfofromSource *DataforwardingandOffloadingInfofromSource) *PdusessionResourceChangeRequiredInfoSNterminated {
	m.DataforwardinginfofromSource = dataforwardinginfofromSource
	return m
}

func (m *PdusessionResourceChangeConfirmInfoSNterminated) SetDataforwardinginfoTarget(dataforwardinginfoTarget *DataForwardingInfoFromTargetNgrannode) *PdusessionResourceChangeConfirmInfoSNterminated {
	m.DataforwardinginfoTarget = dataforwardinginfoTarget
	return m
}

func (m *PdusessionResourceModRqdInfoSNterminated) SetDLNgUTnlatNgRan(dLNgUTnlatNgRan *UptransportLayerInformation) *PdusessionResourceModRqdInfoSNterminated {
	m.DLNgUTnlatNgRan = dLNgUTnlatNgRan
	return m
}

func (m *PdusessionResourceModRqdInfoSNterminated) SetQoSflowsToBeReleasedList(qoSflowsToBeReleasedList *QoSflowsListwithCause) *PdusessionResourceModRqdInfoSNterminated {
	m.QoSflowsToBeReleasedList = qoSflowsToBeReleasedList
	return m
}

func (m *PdusessionResourceModRqdInfoSNterminated) SetDataforwardinginfofromSource(dataforwardinginfofromSource *DataforwardingandOffloadingInfofromSource) *PdusessionResourceModRqdInfoSNterminated {
	m.DataforwardinginfofromSource = dataforwardinginfofromSource
	return m
}

func (m *PdusessionResourceModRqdInfoSNterminated) SetDrbsToBeSetupList(drbsToBeSetupList *DrbsToBeSetupListModRqdSNterminated) *PdusessionResourceModRqdInfoSNterminated {
	m.DrbsToBeSetupList = drbsToBeSetupList
	return m
}

func (m *PdusessionResourceModRqdInfoSNterminated) SetDrbsToBeModifiedList(drbsToBeModifiedList *DrbsToBeModifiedListModRqdSNterminated) *PdusessionResourceModRqdInfoSNterminated {
	m.DrbsToBeModifiedList = drbsToBeModifiedList
	return m
}

func (m *PdusessionResourceModRqdInfoSNterminated) SetDRbsToBeReleased(dRbsToBeReleased *DrbListwithCause) *PdusessionResourceModRqdInfoSNterminated {
	m.DRbsToBeReleased = dRbsToBeReleased
	return m
}

func (m *DrbsToBeSetupListModRqdSNterminatedItem) SetPDcpSnlength(pDcpSnlength *Pdcpsnlength) *DrbsToBeSetupListModRqdSNterminatedItem {
	m.PDcpSnlength = pDcpSnlength
	return m
}

func (m *DrbsToBeSetupListModRqdSNterminatedItem) SetSecondarySnUlPdcpUpTnlinfo(secondarySnUlPdcpUpTnlinfo *UptransportParameters) *DrbsToBeSetupListModRqdSNterminatedItem {
	m.SecondarySnUlPdcpUpTnlinfo = secondarySnUlPdcpUpTnlinfo
	return m
}

func (m *DrbsToBeSetupListModRqdSNterminatedItem) SetDuplicationActivation(duplicationActivation DuplicationActivation) *DrbsToBeSetupListModRqdSNterminatedItem {
	m.DuplicationActivation = &duplicationActivation
	return m
}

func (m *DrbsToBeSetupListModRqdSNterminatedItem) SetULConfiguration(uLConfiguration *Ulconfiguration) *DrbsToBeSetupListModRqdSNterminatedItem {
	m.ULConfiguration = uLConfiguration
	return m
}

func (m *QoSflowsSetupMappedtoDrbModRqdSNterminatedItem) SetMCgrequestedGbrqoSflowInfo(mCgrequestedGbrqoSflowInfo *GbrqoSflowInfo) *QoSflowsSetupMappedtoDrbModRqdSNterminatedItem {
	m.MCgrequestedGbrqoSflowInfo = mCgrequestedGbrqoSflowInfo
	return m
}

func (m *DrbsToBeModifiedListModRqdSNterminatedItem) SetSNUlPdcpUpTnlinfo(sNUlPdcpUpTnlinfo *UptransportParameters) *DrbsToBeModifiedListModRqdSNterminatedItem {
	m.SNUlPdcpUpTnlinfo = sNUlPdcpUpTnlinfo
	return m
}

func (m *DrbsToBeModifiedListModRqdSNterminatedItem) SetDRbQoS(dRbQoS *QoSflowLevelQoSparameters) *DrbsToBeModifiedListModRqdSNterminatedItem {
	m.DRbQoS = dRbQoS
	return m
}

func (m *DrbsToBeModifiedListModRqdSNterminatedItem) SetSecondarySnUlPdcpUpTnlinfo(secondarySnUlPdcpUpTnlinfo *UptransportParameters) *DrbsToBeModifiedListModRqdSNterminatedItem {
	m.SecondarySnUlPdcpUpTnlinfo = secondarySnUlPdcpUpTnlinfo
	return m
}

func (m *DrbsToBeModifiedListModRqdSNterminatedItem) SetULConfiguration(uLConfiguration *Ulconfiguration) *DrbsToBeModifiedListModRqdSNterminatedItem {
	m.ULConfiguration = uLConfiguration
	return m
}

func (m *DrbsToBeModifiedListModRqdSNterminatedItem) SetPdcpDuplicationConfiguration(pdcpDuplicationConfiguration PdcpduplicationConfiguration) *DrbsToBeModifiedListModRqdSNterminatedItem {
	m.PdcpDuplicationConfiguration = &pdcpDuplicationConfiguration
	return m
}

func (m *DrbsToBeModifiedListModRqdSNterminatedItem) SetDuplicationActivation(duplicationActivation DuplicationActivation) *DrbsToBeModifiedListModRqdSNterminatedItem {
	m.DuplicationActivation = &duplicationActivation
	return m
}

func (m *DrbsToBeModifiedListModRqdSNterminatedItem) SetQoSflowsMappedtoDrbModRqdSnterminated(qoSflowsMappedtoDrbModRqdSnterminated *QoSflowsModifiedMappedtoDrbModRqdSNterminated) *DrbsToBeModifiedListModRqdSNterminatedItem {
	m.QoSflowsMappedtoDrbModRqdSnterminated = qoSflowsMappedtoDrbModRqdSnterminated
	return m
}

func (m *QoSflowsModifiedMappedtoDrbModRqdSNterminatedItem) SetMCgrequestedGbrqoSflowInfo(mCgrequestedGbrqoSflowInfo *GbrqoSflowInfo) *QoSflowsModifiedMappedtoDrbModRqdSNterminatedItem {
	m.MCgrequestedGbrqoSflowInfo = mCgrequestedGbrqoSflowInfo
	return m
}

func (m *PdusessionResourceModConfirmInfoSNterminated) SetULNgUTnlatUpf(uLNgUTnlatUpf *UptransportLayerInformation) *PdusessionResourceModConfirmInfoSNterminated {
	m.ULNgUTnlatUpf = uLNgUTnlatUpf
	return m
}

func (m *PdusessionResourceModConfirmInfoSNterminated) SetDRbsNotAdmittedSetupModifyList(dRbsNotAdmittedSetupModifyList *DrbListwithCause) *PdusessionResourceModConfirmInfoSNterminated {
	m.DRbsNotAdmittedSetupModifyList = dRbsNotAdmittedSetupModifyList
	return m
}

func (m *PdusessionResourceModConfirmInfoSNterminated) SetDataforwardinginfoTarget(dataforwardinginfoTarget *DataForwardingInfoFromTargetNgrannode) *PdusessionResourceModConfirmInfoSNterminated {
	m.DataforwardinginfoTarget = dataforwardinginfoTarget
	return m
}

func (m *DrbsAdmittedListModConfirmSNterminatedItem) SetMNDlCgUpTnlinfo(mNDlCgUpTnlinfo *UptransportParameters) *DrbsAdmittedListModConfirmSNterminatedItem {
	m.MNDlCgUpTnlinfo = mNDlCgUpTnlinfo
	return m
}

func (m *DrbsAdmittedListModConfirmSNterminatedItem) SetSecondaryMnDlCgUpTnlinfo(secondaryMnDlCgUpTnlinfo *UptransportParameters) *DrbsAdmittedListModConfirmSNterminatedItem {
	m.SecondaryMnDlCgUpTnlinfo = secondaryMnDlCgUpTnlinfo
	return m
}

func (m *DrbsAdmittedListModConfirmSNterminatedItem) SetLCID(lCID *Lcid) *DrbsAdmittedListModConfirmSNterminatedItem {
	m.LCid = lCID
	return m
}

func (m *PdusessionResourceModRqdInfoMNterminated) SetDRbsToBeModified(dRbsToBeModified *DrbsToBeModifiedListModRqdMNterminated) *PdusessionResourceModRqdInfoMNterminated {
	m.DRbsToBeModified = dRbsToBeModified
	return m
}

func (m *PdusessionResourceModRqdInfoMNterminated) SetDRbsToBeReleased(dRbsToBeReleased *DrbListwithCause) *PdusessionResourceModRqdInfoMNterminated {
	m.DRbsToBeReleased = dRbsToBeReleased
	return m
}

func (m *DrbsToBeModifiedListModRqdMNterminatedItem) SetSecondarySnDlScgUpTnlinfo(secondarySnDlScgUpTnlinfo *UptransportLayerInformation) *DrbsToBeModifiedListModRqdMNterminatedItem {
	m.SecondarySnDlScgUpTnlinfo = secondarySnDlScgUpTnlinfo
	return m
}

func (m *DrbsToBeModifiedListModRqdMNterminatedItem) SetLCID(lCID *Lcid) *DrbsToBeModifiedListModRqdMNterminatedItem {
	m.LCid = lCID
	return m
}

func (m *DrbsToBeModifiedListModRqdMNterminatedItem) SetRlcStatus(rlcStatus *RlcStatus) *DrbsToBeModifiedListModRqdMNterminatedItem {
	m.RlcStatus = rlcStatus
	return m
}

func (m *ProtectedEUTraresourceIndication) SetMbsfnControlRegionLength(mbsfnControlRegionLength *MbsfncontrolRegionLength) *ProtectedEUTraresourceIndication {
	m.MbsfnControlRegionLength = mbsfnControlRegionLength
	return m
}

func (m *QoSflowLevelQoSparameters) SetGBrqoSflowInfo(gBrqoSflowInfo *GbrqoSflowInfo) *QoSflowLevelQoSparameters {
	m.GBrqoSflowInfo = gBrqoSflowInfo
	return m
}

func (m *QoSflowLevelQoSparameters) SetRelectiveQoS(relectiveQoS ReflectiveQoSattribute) *QoSflowLevelQoSparameters {
	m.RelectiveQoS = &relectiveQoS
	return m
}

func (m *QoSflowLevelQoSparameters) SetAdditionalQoSflowInfo(additionalQoSflowInfo AdditionalQoSflowInfoQoSflowLevelQoSparameters) *QoSflowLevelQoSparameters {
	m.AdditionalQoSflowInfo = &additionalQoSflowInfo
	return m
}

func (m *QoSflowItem) SetQosFlowMappingIndication(qosFlowMappingIndication QoSflowMappingIndication) *QoSflowItem {
	m.QosFlowMappingIndication = &qosFlowMappingIndication
	return m
}

func (m *QoSflowwithCauseItem) SetCause(cause *Cause) *QoSflowwithCauseItem {
	m.Cause = cause
	return m
}

func (m *QoSMappingInformation) SetDscp(dscp *asn1.BitString) *QoSMappingInformation {
	m.Dscp = dscp
	return m
}

func (m *QoSMappingInformation) SetFlowLabel(flowLabel *asn1.BitString) *QoSMappingInformation {
	m.FlowLabel = flowLabel
	return m
}

func (m *QoSflowsToBeSetupItem) SetERabID(eRabID *ERAbID) *QoSflowsToBeSetupItem {
	m.ERabId = eRabID
	return m
}

func (m *RanareaId) SetRAnac(rAnac *Ranac) *RanareaId {
	m.RAnac = rAnac
	return m
}

func (m *RanpagingAttemptInfo) SetNextPagingAreaScope(nextPagingAreaScope NextPagingAreaScopeRanpagingAttemptInfo) *RanpagingAttemptInfo {
	m.NextPagingAreaScope = &nextPagingAreaScope
	return m
}

func (m *ReservedSubframePattern) SetMbsfnControlRegionLength(mbsfnControlRegionLength *MbsfncontrolRegionLength) *ReservedSubframePattern {
	m.MbsfnControlRegionLength = mbsfnControlRegionLength
	return m
}

func (m *ResetRequestPartialReleaseItem) SetNgRanNode1UexnApID(ngRanNode1UexnApID *NgRAnnodeUexnApid) *ResetRequestPartialReleaseItem {
	m.NgRanNode1UexnApid = ngRanNode1UexnApID
	return m
}

func (m *ResetRequestPartialReleaseItem) SetNgRanNode2UexnApID(ngRanNode2UexnApID *NgRAnnodeUexnApid) *ResetRequestPartialReleaseItem {
	m.NgRanNode2UexnApid = ngRanNode2UexnApID
	return m
}

func (m *ResetResponsePartialReleaseItem) SetNgRanNode1UexnApID(ngRanNode1UexnApID *NgRAnnodeUexnApid) *ResetResponsePartialReleaseItem {
	m.NgRanNode1UexnApid = ngRanNode1UexnApID
	return m
}

func (m *ResetResponsePartialReleaseItem) SetNgRanNode2UexnApID(ngRanNode2UexnApID *NgRAnnodeUexnApid) *ResetResponsePartialReleaseItem {
	m.NgRanNode2UexnApid = ngRanNode2UexnApID
	return m
}

func (m *RlcduplicationInformation) SetRLcPrimaryIndicator(rLcPrimaryIndicator RlcprimaryIndicatorRlcduplicationInformation) *RlcduplicationInformation {
	m.RLcPrimaryIndicator = &rLcPrimaryIndicator
	return m
}

func (m *Rrcsetupinitiated) SetUErlfreportContainer(uErlfreportContainer *UerlfreportContainer) *Rrcsetupinitiated {
	m.UErlfreportContainer = uErlfreportContainer
	return m
}

func (m *SecondaryRatusageInformation) SetPDusessionUsageReport(pDusessionUsageReport *PdusessionUsageReport) *SecondaryRatusageInformation {
	m.PDusessionUsageReport = pDusessionUsageReport
	return m
}

func (m *SecondaryRatusageInformation) SetQosFlowsUsageReportList(qosFlowsUsageReportList *QoSflowsUsageReportList) *SecondaryRatusageInformation {
	m.QosFlowsUsageReportList = qosFlowsUsageReportList
	return m
}

func (m *SecurityIndication) SetMaximumIpdatarate(maximumIpdatarate *MaximumIpdatarate) *SecurityIndication {
	m.MaximumIpdatarate = maximumIpdatarate
	return m
}

func (m *SensorMeasurementConfiguration) SetSensorMeasConfigNameList(sensorMeasConfigNameList *SensorMeasConfigNameList) *SensorMeasurementConfiguration {
	m.SensorMeasConfigNameList = sensorMeasConfigNameList
	return m
}

func (m *SensorName) SetUncompensatedBarometricConfig(uncompensatedBarometricConfig UncompensatedBarometricConfigSensorName) *SensorName {
	m.UncompensatedBarometricConfig = &uncompensatedBarometricConfig
	return m
}

func (m *SensorName) SetUeSpeedConfig(ueSpeedConfig UeSpeedConfigSensorName) *SensorName {
	m.UeSpeedConfig = &ueSpeedConfig
	return m
}

func (m *SensorName) SetUeOrientationConfig(ueOrientationConfig UeOrientationConfigSensorName) *SensorName {
	m.UeOrientationConfig = &ueOrientationConfig
	return m
}

func (m *ServedCellInformationEUTra) SetRanac(ranac *Ranac) *ServedCellInformationEUTra {
	m.Ranac = ranac
	return m
}

func (m *ServedCellInformationEUTra) SetNumberofAntennaPorts(numberofAntennaPorts NumberOfAntennaPortsEUTra) *ServedCellInformationEUTra {
	m.NumberofAntennaPorts = &numberofAntennaPorts
	return m
}

func (m *ServedCellInformationEUTra) SetPrachConfiguration(prachConfiguration *EUTraprachconfiguration) *ServedCellInformationEUTra {
	m.PrachConfiguration = prachConfiguration
	return m
}

func (m *ServedCellInformationEUTra) SetMBsfnsubframeInfo(mBsfnsubframeInfo *MbsfnsubframeInfoEUTra) *ServedCellInformationEUTra {
	m.MBsfnsubframeInfo = mBsfnsubframeInfo
	return m
}

func (m *ServedCellInformationEUTra) SetMultibandInfo(multibandInfo *EUTramultibandInfoList) *ServedCellInformationEUTra {
	m.MultibandInfo = multibandInfo
	return m
}

func (m *ServedCellInformationEUTra) SetFreqBandIndicatorPriority(freqBandIndicatorPriority FreqBandIndicatorPriorityServedCellInformationEutra) *ServedCellInformationEUTra {
	m.FreqBandIndicatorPriority = &freqBandIndicatorPriority
	return m
}

func (m *ServedCellInformationEUTra) SetBandwIDthReducedSi(bandwIDthReducedSi BandwidthReducedSiservedCellInformationEutra) *ServedCellInformationEUTra {
	m.BandwidthReducedSi = &bandwIDthReducedSi
	return m
}

func (m *ServedCellInformationEUTra) SetProtectedEUtraresourceIndication(protectedEUtraresourceIndication *ProtectedEUTraresourceIndication) *ServedCellInformationEUTra {
	m.ProtectedEUtraresourceIndication = protectedEUtraresourceIndication
	return m
}

func (m *ServedCellsEUTraItem) SetNeighbourInfoNr(neighbourInfoNr *NeighbourInformationNR) *ServedCellsEUTraItem {
	m.NeighbourInfoNr = neighbourInfoNr
	return m
}

func (m *ServedCellsEUTraItem) SetNeighbourInfoEUtra(neighbourInfoEUtra *NeighbourInformationEUTra) *ServedCellsEUTraItem {
	m.NeighbourInfoEUtra = neighbourInfoEUtra
	return m
}

func (m *ServedCellsToUpdateEUTra) SetServedCellsToAddEUtra(servedCellsToAddEUtra *ServedCellsEUTra) *ServedCellsToUpdateEUTra {
	m.ServedCellsToAddEUtra = servedCellsToAddEUtra
	return m
}

func (m *ServedCellsToUpdateEUTra) SetServedCellsToModifyEUtra(servedCellsToModifyEUtra *ServedCellsToModifyEUTra) *ServedCellsToUpdateEUTra {
	m.ServedCellsToModifyEUtra = servedCellsToModifyEUtra
	return m
}

func (m *ServedCellsToUpdateEUTra) SetServedCellsToDeleteEUtra(servedCellsToDeleteEUtra []*EUTraCGi) *ServedCellsToUpdateEUTra {
	m.ServedCellsToDeleteEUtra = servedCellsToDeleteEUtra
	return m
}

func (m *ServedCellsToModifyEUTraItem) SetNeighbourInfoNr(neighbourInfoNr *NeighbourInformationNR) *ServedCellsToModifyEUTraItem {
	m.NeighbourInfoNr = neighbourInfoNr
	return m
}

func (m *ServedCellsToModifyEUTraItem) SetNeighbourInfoEUtra(neighbourInfoEUtra *NeighbourInformationEUTra) *ServedCellsToModifyEUTraItem {
	m.NeighbourInfoEUtra = neighbourInfoEUtra
	return m
}

func (m *ServedCellsToModifyEUTraItem) SetDeactivationIndication(deactivationIndication DeactivationindicationServedCellsToModifyEutraitem) *ServedCellsToModifyEUTraItem {
	m.DeactivationIndication = &deactivationIndication
	return m
}

func (m *ServedCellInformationNR) SetRanac(ranac *Ranac) *ServedCellInformationNR {
	m.Ranac = ranac
	return m
}

func (m *ServedCellsNRItem) SetNeighbourInfoNr(neighbourInfoNr *NeighbourInformationNR) *ServedCellsNRItem {
	m.NeighbourInfoNr = neighbourInfoNr
	return m
}

func (m *ServedCellsNRItem) SetNeighbourInfoEUtra(neighbourInfoEUtra *NeighbourInformationEUTra) *ServedCellsNRItem {
	m.NeighbourInfoEUtra = neighbourInfoEUtra
	return m
}

func (m *ServedCellsToModifyNRItem) SetNeighbourInfoNr(neighbourInfoNr *NeighbourInformationNR) *ServedCellsToModifyNRItem {
	m.NeighbourInfoNr = neighbourInfoNr
	return m
}

func (m *ServedCellsToModifyNRItem) SetNeighbourInfoEUtra(neighbourInfoEUtra *NeighbourInformationEUTra) *ServedCellsToModifyNRItem {
	m.NeighbourInfoEUtra = neighbourInfoEUtra
	return m
}

func (m *ServedCellsToModifyNRItem) SetDeactivationIndication(deactivationIndication DeactivationindicationServedCellsToModifyNritem) *ServedCellsToModifyNRItem {
	m.DeactivationIndication = &deactivationIndication
	return m
}

func (m *ServedCellsToUpdateNR) SetServedCellsToAddNr(servedCellsToAddNr *ServedCellsNR) *ServedCellsToUpdateNR {
	m.ServedCellsToAddNr = servedCellsToAddNr
	return m
}

func (m *ServedCellsToUpdateNR) SetServedCellsToModifyNr(servedCellsToModifyNr *ServedCellsToModifyNR) *ServedCellsToUpdateNR {
	m.ServedCellsToModifyNr = servedCellsToModifyNr
	return m
}

func (m *ServedCellsToUpdateNR) SetServedCellsToDeleteNr(servedCellsToDeleteNr []*NrCGi) *ServedCellsToUpdateNR {
	m.ServedCellsToDeleteNr = servedCellsToDeleteNr
	return m
}

func (m *SNSsai) SetSd(sd []byte) *SNSsai {
	m.Sd = sd
	return m
}

func (m *TnlconfigurationInfo) SetExtendedUptransportLayerAddressesToAdd(extendedUptransportLayerAddressesToAdd *ExtTlas) *TnlconfigurationInfo {
	m.ExtendedUptransportLayerAddressesToAdd = extendedUptransportLayerAddressesToAdd
	return m
}

func (m *TnlconfigurationInfo) SetExtendedUptransportLayerAddressesToRemove(extendedUptransportLayerAddressesToRemove *ExtTlas) *TnlconfigurationInfo {
	m.ExtendedUptransportLayerAddressesToRemove = extendedUptransportLayerAddressesToRemove
	return m
}

func (m *TnlaToUpdateItem) SetTNlassociationUsage(tNlassociationUsage TnlassociationUsage) *TnlaToUpdateItem {
	m.TNlassociationUsage = &tNlassociationUsage
	return m
}

func (m *TsctrafficCharacteristics) SetTScassistanceInformationDownlink(tScassistanceInformationDownlink *TscassistanceInformation) *TsctrafficCharacteristics {
	m.TScassistanceInformationDownlink = tScassistanceInformationDownlink
	return m
}

func (m *TsctrafficCharacteristics) SetTScassistanceInformationUplink(tScassistanceInformationUplink *TscassistanceInformation) *TsctrafficCharacteristics {
	m.TScassistanceInformationUplink = tScassistanceInformationUplink
	return m
}

func (m *TscassistanceInformation) SetBurstArrivalTime(burstArrivalTime []byte) *TscassistanceInformation {
	m.BurstArrivalTime = burstArrivalTime
	return m
}

func (m *UecontextInfoRetrUectxtResp) SetMobilityRestrictionList(mobilityRestrictionList *MobilityRestrictionList) *UecontextInfoRetrUectxtResp {
	m.MobilityRestrictionList = mobilityRestrictionList
	return m
}

func (m *UecontextInfoRetrUectxtResp) SetIndexToRatFrequencySelectionPriority(indexToRatFrequencySelectionPriority *RfspIndex) *UecontextInfoRetrUectxtResp {
	m.IndexToRatFrequencySelectionPriority = indexToRatFrequencySelectionPriority
	return m
}

func (m *UeradioCapabilityForPaging) SetUEradioCapabilityForPagingOfNr(uEradioCapabilityForPagingOfNr *UeradioCapabilityForPagingOfNr) *UeradioCapabilityForPaging {
	m.UEradioCapabilityForPagingOfNr = uEradioCapabilityForPagingOfNr
	return m
}

func (m *UeradioCapabilityForPaging) SetUEradioCapabilityForPagingOfEutra(uEradioCapabilityForPagingOfEutra *UeradioCapabilityForPagingOfEutra) *UeradioCapabilityForPaging {
	m.UEradioCapabilityForPagingOfEutra = uEradioCapabilityForPagingOfEutra
	return m
}

func (m *WlanmeasurementConfiguration) SetWlanMeasConfigNameList(wlanMeasConfigNameList *WlanmeasConfigNameList) *WlanmeasurementConfiguration {
	m.WlanMeasConfigNameList = wlanMeasConfigNameList
	return m
}

func (m *WlanmeasurementConfiguration) SetWlanRssi(wlanRssi WlanrssiWlanmeasurementConfiguration) *WlanmeasurementConfiguration {
	m.WlanRssi = &wlanRssi
	return m
}

func (m *WlanmeasurementConfiguration) SetWlanRtt(wlanRtt WlanrttWlanmeasurementConfiguration) *WlanmeasurementConfiguration {
	m.WlanRtt = &wlanRtt
	return m
}
