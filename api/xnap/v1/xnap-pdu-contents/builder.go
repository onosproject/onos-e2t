// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package xnappducontentsv1

import (
	xnapiesv1 "github.com/onosproject/onos-e2t/api/xnap/v1/xnap-ies"
)

func (m *UecontextInfoHorequest) SetIndexToRatFrequencySelectionPriority(indexToRatFrequencySelectionPriority *xnapiesv1.RfspIndex) *UecontextInfoHorequest {
	m.IndexToRatFrequencySelectionPriority = indexToRatFrequencySelectionPriority
	return m
}

func (m *UecontextInfoHorequest) SetLocationReportingInformation(locationReportingInformation *xnapiesv1.LocationReportingInformation) *UecontextInfoHorequest {
	m.LocationReportingInformation = locationReportingInformation
	return m
}

func (m *UecontextInfoHorequest) SetMrl(mrl *xnapiesv1.MobilityRestrictionList) *UecontextInfoHorequest {
	m.Mrl = mrl
	return m
}

func (m *PdusessionToBeAddedAddReqItem) SetSNPdusessionAmbr(sNPdusessionAmbr *xnapiesv1.PdusessionAggregateMaximumBitRate) *PdusessionToBeAddedAddReqItem {
	m.SNPdusessionAmbr = sNPdusessionAmbr
	return m
}

func (m *PdusessionToBeAddedAddReqItem) SetSnTerminated(snTerminated *xnapiesv1.PdusessionResourceSetupInfoSNterminated) *PdusessionToBeAddedAddReqItem {
	m.SnTerminated = snTerminated
	return m
}

func (m *PdusessionToBeAddedAddReqItem) SetMnTerminated(mnTerminated *xnapiesv1.PdusessionResourceSetupInfoMNterminated) *PdusessionToBeAddedAddReqItem {
	m.MnTerminated = mnTerminated
	return m
}

func (m *PdusessionAdmittedAddedAddReqAckItem) SetSnTerminated(snTerminated *xnapiesv1.PdusessionResourceSetupResponseInfoSNterminated) *PdusessionAdmittedAddedAddReqAckItem {
	m.SnTerminated = snTerminated
	return m
}

func (m *PdusessionAdmittedAddedAddReqAckItem) SetMnTerminated(mnTerminated *xnapiesv1.PdusessionResourceSetupResponseInfoMNterminated) *PdusessionAdmittedAddedAddReqAckItem {
	m.MnTerminated = mnTerminated
	return m
}

func (m *PdusessionNotAdmittedAddReqAck) SetPduSessionResourcesNotAdmittedSnterminated(pduSessionResourcesNotAdmittedSnterminated *xnapiesv1.PdusessionResourcesNotAdmittedList) *PdusessionNotAdmittedAddReqAck {
	m.PduSessionResourcesNotAdmittedSnterminated = pduSessionResourcesNotAdmittedSnterminated
	return m
}

func (m *PdusessionNotAdmittedAddReqAck) SetPduSessionResourcesNotAdmittedMnterminated(pduSessionResourcesNotAdmittedMnterminated *xnapiesv1.PdusessionResourcesNotAdmittedList) *PdusessionNotAdmittedAddReqAck {
	m.PduSessionResourcesNotAdmittedMnterminated = pduSessionResourcesNotAdmittedMnterminated
	return m
}

func (m *Configurationsuccessfullyapplied) SetMNgRannodeToSNgRannodeContainer(mNgRannodeToSNgRannodeContainer []byte) *Configurationsuccessfullyapplied {
	m.MNgRannodeToSNgRannodeContainer = mNgRannodeToSNgRannodeContainer
	return m
}

func (m *ConfigurationrejectedbyMNGRAnnode) SetMNgRannodeToSNgRannodeContainer(mNgRannodeToSNgRannodeContainer []byte) *ConfigurationrejectedbyMNGRAnnode {
	m.MNgRannodeToSNgRannodeContainer = mNgRannodeToSNgRannodeContainer
	return m
}

func (m *UecontextInfoSNmodRequest) SetUeSecurityCapabilities(ueSecurityCapabilities *xnapiesv1.UesecurityCapabilities) *UecontextInfoSNmodRequest {
	m.UeSecurityCapabilities = ueSecurityCapabilities
	return m
}

func (m *UecontextInfoSNmodRequest) SetSNgRannodeSecurityKey(sNgRannodeSecurityKey *xnapiesv1.SNGRAnnodeSecurityKey) *UecontextInfoSNmodRequest {
	m.SNgRannodeSecurityKey = sNgRannodeSecurityKey
	return m
}

func (m *UecontextInfoSNmodRequest) SetSNgRannodeUeAmbr(sNgRannodeUeAmbr *xnapiesv1.UeaggregateMaximumBitRate) *UecontextInfoSNmodRequest {
	m.SNgRannodeUeAmbr = sNgRannodeUeAmbr
	return m
}

func (m *UecontextInfoSNmodRequest) SetIndexToRatFrequencySelectionPriority(indexToRatFrequencySelectionPriority *xnapiesv1.RfspIndex) *UecontextInfoSNmodRequest {
	m.IndexToRatFrequencySelectionPriority = indexToRatFrequencySelectionPriority
	return m
}

func (m *UecontextInfoSNmodRequest) SetLowerLayerPresenceStatusChange(lowerLayerPresenceStatusChange xnapiesv1.LowerLayerPresenceStatusChange) *UecontextInfoSNmodRequest {
	m.LowerLayerPresenceStatusChange = &lowerLayerPresenceStatusChange
	return m
}

func (m *UecontextInfoSNmodRequest) SetPduSessionResourceToBeAdded(pduSessionResourceToBeAdded *PdusessionsToBeAddedSNmodRequestList) *UecontextInfoSNmodRequest {
	m.PduSessionResourceToBeAdded = pduSessionResourceToBeAdded
	return m
}

func (m *UecontextInfoSNmodRequest) SetPduSessionResourceToBeModified(pduSessionResourceToBeModified *PdusessionsToBeModifiedSNmodRequestList) *UecontextInfoSNmodRequest {
	m.PduSessionResourceToBeModified = pduSessionResourceToBeModified
	return m
}

func (m *UecontextInfoSNmodRequest) SetPduSessionResourceToBeReleased(pduSessionResourceToBeReleased *PdusessionsToBeReleasedSNmodRequestList) *UecontextInfoSNmodRequest {
	m.PduSessionResourceToBeReleased = pduSessionResourceToBeReleased
	return m
}

func (m *PdusessionsToBeAddedSNmodRequestItem) SetSNPdusessionAmbr(sNPdusessionAmbr *xnapiesv1.PdusessionAggregateMaximumBitRate) *PdusessionsToBeAddedSNmodRequestItem {
	m.SNPdusessionAmbr = sNPdusessionAmbr
	return m
}

func (m *PdusessionsToBeAddedSNmodRequestItem) SetSnTerminated(snTerminated *xnapiesv1.PdusessionResourceSetupInfoSNterminated) *PdusessionsToBeAddedSNmodRequestItem {
	m.SnTerminated = snTerminated
	return m
}

func (m *PdusessionsToBeAddedSNmodRequestItem) SetMnTerminated(mnTerminated *xnapiesv1.PdusessionResourceSetupInfoMNterminated) *PdusessionsToBeAddedSNmodRequestItem {
	m.MnTerminated = mnTerminated
	return m
}

func (m *PdusessionsToBeModifiedSNmodRequestItem) SetSNPdusessionAmbr(sNPdusessionAmbr *xnapiesv1.PdusessionAggregateMaximumBitRate) *PdusessionsToBeModifiedSNmodRequestItem {
	m.SNPdusessionAmbr = sNPdusessionAmbr
	return m
}

func (m *PdusessionsToBeModifiedSNmodRequestItem) SetSnTerminated(snTerminated *xnapiesv1.PdusessionResourceModificationInfoSNterminated) *PdusessionsToBeModifiedSNmodRequestItem {
	m.SnTerminated = snTerminated
	return m
}

func (m *PdusessionsToBeModifiedSNmodRequestItem) SetMnTerminated(mnTerminated *xnapiesv1.PdusessionResourceModificationInfoMNterminated) *PdusessionsToBeModifiedSNmodRequestItem {
	m.MnTerminated = mnTerminated
	return m
}

func (m *PdusessionsToBeReleasedSNmodRequestList) SetPduSessionList(pduSessionList *xnapiesv1.PdusessionListwithCause) *PdusessionsToBeReleasedSNmodRequestList {
	m.PduSessionList = pduSessionList
	return m
}

func (m *PdusessionAdmittedSNmodResponse) SetPduSessionResourcesAdmittedToBeAdded(pduSessionResourcesAdmittedToBeAdded *PdusessionAdmittedToBeAddedSnmodResponse) *PdusessionAdmittedSNmodResponse {
	m.PduSessionResourcesAdmittedToBeAdded = pduSessionResourcesAdmittedToBeAdded
	return m
}

func (m *PdusessionAdmittedSNmodResponse) SetPduSessionResourcesAdmittedToBeModified(pduSessionResourcesAdmittedToBeModified *PdusessionAdmittedToBeModifiedSnmodResponse) *PdusessionAdmittedSNmodResponse {
	m.PduSessionResourcesAdmittedToBeModified = pduSessionResourcesAdmittedToBeModified
	return m
}

func (m *PdusessionAdmittedSNmodResponse) SetPduSessionResourcesAdmittedToBeReleased(pduSessionResourcesAdmittedToBeReleased *PdusessionAdmittedToBeReleasedSnmodResponse) *PdusessionAdmittedSNmodResponse {
	m.PduSessionResourcesAdmittedToBeReleased = pduSessionResourcesAdmittedToBeReleased
	return m
}

func (m *PdusessionAdmittedToBeAddedSnmodResponseItem) SetSnTerminated(snTerminated *xnapiesv1.PdusessionResourceSetupResponseInfoSNterminated) *PdusessionAdmittedToBeAddedSnmodResponseItem {
	m.SnTerminated = snTerminated
	return m
}

func (m *PdusessionAdmittedToBeAddedSnmodResponseItem) SetMnTerminated(mnTerminated *xnapiesv1.PdusessionResourceSetupResponseInfoMNterminated) *PdusessionAdmittedToBeAddedSnmodResponseItem {
	m.MnTerminated = mnTerminated
	return m
}

func (m *PdusessionAdmittedToBeModifiedSnmodResponseItem) SetSnTerminated(snTerminated *xnapiesv1.PdusessionResourceModificationResponseInfoSNterminated) *PdusessionAdmittedToBeModifiedSnmodResponseItem {
	m.SnTerminated = snTerminated
	return m
}

func (m *PdusessionAdmittedToBeModifiedSnmodResponseItem) SetMnTerminated(mnTerminated *xnapiesv1.PdusessionResourceModificationResponseInfoMNterminated) *PdusessionAdmittedToBeModifiedSnmodResponseItem {
	m.MnTerminated = mnTerminated
	return m
}

func (m *PdusessionAdmittedToBeReleasedSnmodResponse) SetSnTerminated(snTerminated *xnapiesv1.PdusessionListwithDataForwardingRequest) *PdusessionAdmittedToBeReleasedSnmodResponse {
	m.SnTerminated = snTerminated
	return m
}

func (m *PdusessionAdmittedToBeReleasedSnmodResponse) SetMnTerminated(mnTerminated *xnapiesv1.PdusessionListwithCause) *PdusessionAdmittedToBeReleasedSnmodResponse {
	m.MnTerminated = mnTerminated
	return m
}

func (m *PdusessionNotAdmittedSNmodResponse) SetPduSessionList(pduSessionList *xnapiesv1.PdusessionList) *PdusessionNotAdmittedSNmodResponse {
	m.PduSessionList = pduSessionList
	return m
}

func (m *PdusessionToBeModifiedSnmodRequiredItem) SetSnTerminated(snTerminated *xnapiesv1.PdusessionResourceModRqdInfoSNterminated) *PdusessionToBeModifiedSnmodRequiredItem {
	m.SnTerminated = snTerminated
	return m
}

func (m *PdusessionToBeModifiedSnmodRequiredItem) SetMnTerminated(mnTerminated *xnapiesv1.PdusessionResourceModRqdInfoMNterminated) *PdusessionToBeModifiedSnmodRequiredItem {
	m.MnTerminated = mnTerminated
	return m
}

func (m *PdusessionToBeReleasedSnmodRequired) SetSnTerminated(snTerminated *xnapiesv1.PdusessionListwithDataForwardingRequest) *PdusessionToBeReleasedSnmodRequired {
	m.SnTerminated = snTerminated
	return m
}

func (m *PdusessionToBeReleasedSnmodRequired) SetMnTerminated(mnTerminated *xnapiesv1.PdusessionListwithCause) *PdusessionToBeReleasedSnmodRequired {
	m.MnTerminated = mnTerminated
	return m
}

func (m *PdusessionAdmittedModSnmodConfirmItem) SetSnTerminated(snTerminated *xnapiesv1.PdusessionResourceModConfirmInfoSNterminated) *PdusessionAdmittedModSnmodConfirmItem {
	m.SnTerminated = snTerminated
	return m
}

func (m *PdusessionAdmittedModSnmodConfirmItem) SetMnTerminated(mnTerminated *xnapiesv1.PdusessionResourceModConfirmInfoMNterminated) *PdusessionAdmittedModSnmodConfirmItem {
	m.MnTerminated = mnTerminated
	return m
}

func (m *PdusessionReleasedSnmodConfirm) SetSnTerminated(snTerminated *xnapiesv1.PdusessionListwithDataForwardingFromTarget) *PdusessionReleasedSnmodConfirm {
	m.SnTerminated = snTerminated
	return m
}

func (m *PdusessionReleasedSnmodConfirm) SetMnTerminated(mnTerminated *xnapiesv1.PdusessionList) *PdusessionReleasedSnmodConfirm {
	m.MnTerminated = mnTerminated
	return m
}

func (m *PdusessionToBeReleasedListRelReqAck) SetPduSessionsToBeReleasedListSnterminated(pduSessionsToBeReleasedListSnterminated *xnapiesv1.PdusessionListwithDataForwardingRequest) *PdusessionToBeReleasedListRelReqAck {
	m.PduSessionsToBeReleasedListSnterminated = pduSessionsToBeReleasedListSnterminated
	return m
}

func (m *PdusessionToBeReleasedListRelRqd) SetPduSessionsToBeReleasedListSnterminated(pduSessionsToBeReleasedListSnterminated *xnapiesv1.PdusessionListwithDataForwardingRequest) *PdusessionToBeReleasedListRelRqd {
	m.PduSessionsToBeReleasedListSnterminated = pduSessionsToBeReleasedListSnterminated
	return m
}

func (m *PdusessionReleasedListRelConf) SetPduSessionsReleasedListSnterminated(pduSessionsReleasedListSnterminated *xnapiesv1.PdusessionListwithDataForwardingFromTarget) *PdusessionReleasedListRelConf {
	m.PduSessionsReleasedListSnterminated = pduSessionsReleasedListSnterminated
	return m
}

func (m *PdusessionSNchangeRequiredItem) SetSnTerminated(snTerminated *xnapiesv1.PdusessionResourceChangeRequiredInfoSNterminated) *PdusessionSNchangeRequiredItem {
	m.SnTerminated = snTerminated
	return m
}

func (m *PdusessionSNchangeRequiredItem) SetMnTerminated(mnTerminated *xnapiesv1.PdusessionResourceChangeRequiredInfoMNterminated) *PdusessionSNchangeRequiredItem {
	m.MnTerminated = mnTerminated
	return m
}

func (m *PdusessionSNchangeConfirmItem) SetSnTerminated(snTerminated *xnapiesv1.PdusessionResourceChangeConfirmInfoSNterminated) *PdusessionSNchangeConfirmItem {
	m.SnTerminated = snTerminated
	return m
}

func (m *PdusessionSNchangeConfirmItem) SetMnTerminated(mnTerminated *xnapiesv1.PdusessionResourceChangeConfirmInfoMNterminated) *PdusessionSNchangeConfirmItem {
	m.MnTerminated = mnTerminated
	return m
}

func (m *SplitSrbRRctransfer) SetRrcContainer(rrcContainer []byte) *SplitSrbRRctransfer {
	m.RrcContainer = rrcContainer
	return m
}

func (m *SplitSrbRRctransfer) SetDeliveryStatus(deliveryStatus *xnapiesv1.DeliveryStatus) *SplitSrbRRctransfer {
	m.DeliveryStatus = deliveryStatus
	return m
}

func (m *PdusessionResourcesActivityNotifyItem) SetPduSessionLevelUpactivityreport(pduSessionLevelUpactivityreport xnapiesv1.UserPlaneTrafficActivityReport) *PdusessionResourcesActivityNotifyItem {
	m.PduSessionLevelUpactivityreport = &pduSessionLevelUpactivityreport
	return m
}

func (m *PdusessionResourcesActivityNotifyItem) SetQosFlowsActivityNotifyList(qosFlowsActivityNotifyList *QoSflowsActivityNotifyList) *PdusessionResourcesActivityNotifyItem {
	m.QosFlowsActivityNotifyList = qosFlowsActivityNotifyList
	return m
}

func (m *RespondingNodeTypeConfigUpdateAckgNb) SetServedNrCells(servedNrCells *xnapiesv1.ServedCellsNR) *RespondingNodeTypeConfigUpdateAckgNb {
	m.ServedNrCells = servedNrCells
	return m
}

func (m *ResourceCoordRequestngeNbinitiated) SetListofEUtracells(listofEUtracells []*xnapiesv1.EUTraCGi) *ResourceCoordRequestngeNbinitiated {
	m.ListofEUtracells = listofEUtracells
	return m
}

func (m *ResourceCoordRequestgNbinitiated) SetListofEUtracells(listofEUtracells []*xnapiesv1.EUTraCGi) *ResourceCoordRequestgNbinitiated {
	m.ListofEUtracells = listofEUtracells
	return m
}

func (m *ResourceCoordRequestgNbinitiated) SetListofNrcells(listofNrcells []*xnapiesv1.NrCGi) *ResourceCoordRequestgNbinitiated {
	m.ListofNrcells = listofNrcells
	return m
}

func (m *ResourceCoordResponsengeNbinitiated) SetListofEUtracells(listofEUtracells []*xnapiesv1.EUTraCGi) *ResourceCoordResponsengeNbinitiated {
	m.ListofEUtracells = listofEUtracells
	return m
}

func (m *ResourceCoordResponsegNbinitiated) SetListofNrcells(listofNrcells []*xnapiesv1.NrCGi) *ResourceCoordResponsegNbinitiated {
	m.ListofNrcells = listofNrcells
	return m
}
