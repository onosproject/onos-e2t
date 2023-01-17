// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package v1

// Driven from e2ap_constants.proto
// TODO: Automate the generation of this file

type ProcedureCodeT int32

const (
	ProcedureCodeIDReset                                    ProcedureCodeT = 0
	ProcedureCodeIDF1Setup                                  ProcedureCodeT = 1
	ProcedureCodeIDErrorIndication                          ProcedureCodeT = 2
	ProcedureCodeIDgNBDUConfigurationUpdate                 ProcedureCodeT = 3
	ProcedureCodeIDgNBCUConfigurationUpdate                 ProcedureCodeT = 4
	ProcedureCodeIDUEContextSetup                           ProcedureCodeT = 5
	ProcedureCodeIDUEContextRelease                         ProcedureCodeT = 6
	ProcedureCodeIDUEContextModification                    ProcedureCodeT = 7
	ProcedureCodeIDUEContextModificationRequired            ProcedureCodeT = 8
	ProcedureCodeIDUEMobilityCommand                        ProcedureCodeT = 9
	ProcedureCodeIDUEContextReleaseRequest                  ProcedureCodeT = 10
	ProcedureCodeIDInitialULRRCMessageTransfer              ProcedureCodeT = 11
	ProcedureCodeIDDLRRCMessageTransfer                     ProcedureCodeT = 12
	ProcedureCodeIDULRRCMessageTransfer                     ProcedureCodeT = 13
	ProcedureCodeIDprivateMessage                           ProcedureCodeT = 14
	ProcedureCodeIDUEInactivityNotification                 ProcedureCodeT = 15
	ProcedureCodeIDGNBDUResourceCoordination                ProcedureCodeT = 16
	ProcedureCodeIDSystemInformationDeliveryCommand         ProcedureCodeT = 17
	ProcedureCodeIDPaging                                   ProcedureCodeT = 18
	ProcedureCodeIDNotify                                   ProcedureCodeT = 19
	ProcedureCodeIDWriteReplaceWarning                      ProcedureCodeT = 20
	ProcedureCodeIDPWSCancel                                ProcedureCodeT = 21
	ProcedureCodeIDPWSRestartIndication                     ProcedureCodeT = 22
	ProcedureCodeIDPWSFailureIndication                     ProcedureCodeT = 23
	ProcedureCodeIDGNBDUStatusIndication                    ProcedureCodeT = 24
	ProcedureCodeIDRRCDeliveryReport                        ProcedureCodeT = 25
	ProcedureCodeIDF1Removal                                ProcedureCodeT = 26
	ProcedureCodeIDNetworkAccessRateReduction               ProcedureCodeT = 27
	ProcedureCodeIDTraceStart                               ProcedureCodeT = 28
	ProcedureCodeIDDeactivateTrace                          ProcedureCodeT = 29
	ProcedureCodeIDDUCURadioInformationTransfer             ProcedureCodeT = 30
	ProcedureCodeIDCUDURadioInformationTransfer             ProcedureCodeT = 31
	ProcedureCodeIDBAPMappingConfiguration                  ProcedureCodeT = 32
	ProcedureCodeIDGNBDUResourceConfiguration               ProcedureCodeT = 33
	ProcedureCodeIDIABTNLAddressAllocation                  ProcedureCodeT = 34
	ProcedureCodeIDIABUPConfigurationUpdate                 ProcedureCodeT = 35
	ProcedureCodeIDresourceStatusReportingInitiation        ProcedureCodeT = 36
	ProcedureCodeIDresourceStatusReporting                  ProcedureCodeT = 37
	ProcedureCodeIDaccessAndMobilityIndication              ProcedureCodeT = 38
	ProcedureCodeIDaccessSuccess                            ProcedureCodeT = 39
	ProcedureCodeIDcellTrafficTrace                         ProcedureCodeT = 40
	ProcedureCodeIDPositioningMeasurementExchange           ProcedureCodeT = 41
	ProcedureCodeIDPositioningAssistanceInformationControl  ProcedureCodeT = 42
	ProcedureCodeIDPositioningAssistanceInformationFeedback ProcedureCodeT = 43
	ProcedureCodeIDPositioningMeasurementReport             ProcedureCodeT = 44
	ProcedureCodeIDPositioningMeasurementAbort              ProcedureCodeT = 45
	ProcedureCodeIDPositioningMeasurementFailureIndication  ProcedureCodeT = 46
	ProcedureCodeIDPositioningMeasurementUpdate             ProcedureCodeT = 47
	ProcedureCodeIDTRPInformationExchange                   ProcedureCodeT = 48
	ProcedureCodeIDPositioningInformationExchange           ProcedureCodeT = 49
	ProcedureCodeIDPositioningActivation                    ProcedureCodeT = 50
	ProcedureCodeIDPositioningDeactivation                  ProcedureCodeT = 51
	ProcedureCodeIDECIDMeasurementInitiation                ProcedureCodeT = 52
	ProcedureCodeIDECIDMeasurementFailureIndication         ProcedureCodeT = 53
	ProcedureCodeIDECIDMeasurementReport                    ProcedureCodeT = 54
	ProcedureCodeIDECIDMeasurementTermination               ProcedureCodeT = 55
	ProcedureCodeIDPositioningInformationUpdate             ProcedureCodeT = 56
	ProcedureCodeIDReferenceTimeInformationReport           ProcedureCodeT = 57
	ProcedureCodeIDReferenceTimeInformationReportingControl ProcedureCodeT = 58
)

type ProtocolIeID int32

const (
	ProtocolIeIDCause                                          ProtocolIeID = 0
	ProtocolIeIDCellsFailedtobeActivatedList                   ProtocolIeID = 1
	ProtocolIeIDCellsFailedtobeActivatedListItem               ProtocolIeID = 2
	ProtocolIeIDCellstobeActivatedList                         ProtocolIeID = 3
	ProtocolIeIDCellstobeActivatedListItem                     ProtocolIeID = 4
	ProtocolIeIDCellstobeDeactivatedList                       ProtocolIeID = 5
	ProtocolIeIDCellstobeDeactivatedListItem                   ProtocolIeID = 6
	ProtocolIeIDCriticalityDiagnostics                         ProtocolIeID = 7
	ProtocolIeIDCUtoDURRCInformation                           ProtocolIeID = 9
	ProtocolIeIDDRBsFailedToBeModifiedItem                     ProtocolIeID = 12
	ProtocolIeIDDRBsFailedToBeModifiedList                     ProtocolIeID = 13
	ProtocolIeIDDRBsFailedToBeSetupItem                        ProtocolIeID = 14
	ProtocolIeIDDRBsFailedToBeSetupList                        ProtocolIeID = 15
	ProtocolIeIDDRBsFailedToBeSetupModItem                     ProtocolIeID = 16
	ProtocolIeIDDRBsFailedToBeSetupModList                     ProtocolIeID = 17
	ProtocolIeIDDRBsModifiedConfItem                           ProtocolIeID = 18
	ProtocolIeIDDRBsModifiedConfList                           ProtocolIeID = 19
	ProtocolIeIDDRBsModifiedItem                               ProtocolIeID = 20
	ProtocolIeIDDRBsModifiedList                               ProtocolIeID = 21
	ProtocolIeIDDRBsRequiredToBeModifiedItem                   ProtocolIeID = 22
	ProtocolIeIDDRBsRequiredToBeModifiedList                   ProtocolIeID = 23
	ProtocolIeIDDRBsRequiredToBeReleasedItem                   ProtocolIeID = 24
	ProtocolIeIDDRBsRequiredToBeReleasedList                   ProtocolIeID = 25
	ProtocolIeIDDRBsSetupItem                                  ProtocolIeID = 26
	ProtocolIeIDDRBsSetupList                                  ProtocolIeID = 27
	ProtocolIeIDDRBsSetupModItem                               ProtocolIeID = 28
	ProtocolIeIDDRBsSetupModList                               ProtocolIeID = 29
	ProtocolIeIDDRBsToBeModifiedItem                           ProtocolIeID = 30
	ProtocolIeIDDRBsToBeModifiedList                           ProtocolIeID = 31
	ProtocolIeIDDRBsToBeReleasedItem                           ProtocolIeID = 32
	ProtocolIeIDDRBsToBeReleasedList                           ProtocolIeID = 33
	ProtocolIeIDDRBsToBeSetupItem                              ProtocolIeID = 34
	ProtocolIeIDDRBsToBeSetupList                              ProtocolIeID = 35
	ProtocolIeIDDRBsToBeSetupModItem                           ProtocolIeID = 36
	ProtocolIeIDDRBsToBeSetupModList                           ProtocolIeID = 37
	ProtocolIeIDDRXCycle                                       ProtocolIeID = 38
	ProtocolIeIDDUtoCURRCInformation                           ProtocolIeID = 39
	ProtocolIeIDgNBCUUEF1APID                                  ProtocolIeID = 40
	ProtocolIeIDgNBDUUEF1APID                                  ProtocolIeID = 41
	ProtocolIeIDgNBDUID                                        ProtocolIeID = 42
	ProtocolIeIDGNBDUServedCellsItem                           ProtocolIeID = 43
	ProtocolIeIDgNBDUServedCellsList                           ProtocolIeID = 44
	ProtocolIeIDgNBDUName                                      ProtocolIeID = 45
	ProtocolIeIDNRCellID                                       ProtocolIeID = 46
	ProtocolIeIDoldgNBDUUEF1APID                               ProtocolIeID = 47
	ProtocolIeIDResetType                                      ProtocolIeID = 48
	ProtocolIeIDResourceCoordinationTransferContainer          ProtocolIeID = 49
	ProtocolIeIDRRCContainer                                   ProtocolIeID = 50
	ProtocolIeIDSCellToBeRemovedItem                           ProtocolIeID = 51
	ProtocolIeIDSCellToBeRemovedList                           ProtocolIeID = 52
	ProtocolIeIDSCellToBeSetupItem                             ProtocolIeID = 53
	ProtocolIeIDSCellToBeSetupList                             ProtocolIeID = 54
	ProtocolIeIDSCellToBeSetupModItem                          ProtocolIeID = 55
	ProtocolIeIDSCellToBeSetupModList                          ProtocolIeID = 56
	ProtocolIeIDServedCellsToAddItem                           ProtocolIeID = 57
	ProtocolIeIDServedCellsToAddList                           ProtocolIeID = 58
	ProtocolIeIDServedCellsToDeleteItem                        ProtocolIeID = 59
	ProtocolIeIDServedCellsToDeleteList                        ProtocolIeID = 60
	ProtocolIeIDServedCellsToModifyItem                        ProtocolIeID = 61
	ProtocolIeIDServedCellsToModifyList                        ProtocolIeID = 62
	ProtocolIeIDSpCellID                                       ProtocolIeID = 63
	ProtocolIeIDSRBID                                          ProtocolIeID = 64
	ProtocolIeIDSRBsFailedToBeSetupItem                        ProtocolIeID = 65
	ProtocolIeIDSRBsFailedToBeSetupList                        ProtocolIeID = 66
	ProtocolIeIDSRBsFailedToBeSetupModItem                     ProtocolIeID = 67
	ProtocolIeIDSRBsFailedToBeSetupModList                     ProtocolIeID = 68
	ProtocolIeIDSRBsRequiredToBeReleasedItem                   ProtocolIeID = 69
	ProtocolIeIDSRBsRequiredToBeReleasedList                   ProtocolIeID = 70
	ProtocolIeIDSRBsToBeReleasedItem                           ProtocolIeID = 71
	ProtocolIeIDSRBsToBeReleasedList                           ProtocolIeID = 72
	ProtocolIeIDSRBsToBeSetupItem                              ProtocolIeID = 73
	ProtocolIeIDSRBsToBeSetupList                              ProtocolIeID = 74
	ProtocolIeIDSRBsToBeSetupModItem                           ProtocolIeID = 75
	ProtocolIeIDSRBsToBeSetupModList                           ProtocolIeID = 76
	ProtocolIeIDTimeToWait                                     ProtocolIeID = 77
	ProtocolIeIDTransactionID                                  ProtocolIeID = 78
	ProtocolIeIDTransmissionActionIndicator                    ProtocolIeID = 79
	ProtocolIeIDUEassociatedLogicalF1ConnectionItem            ProtocolIeID = 80
	ProtocolIeIDUEassociatedLogicalF1ConnectionListResAck      ProtocolIeID = 81
	ProtocolIeIDgNBCUName                                      ProtocolIeID = 82
	ProtocolIeIDSCellFailedtoSetupList                         ProtocolIeID = 83
	ProtocolIeIDSCellFailedtoSetupItem                         ProtocolIeID = 84
	ProtocolIeIDSCellFailedtoSetupModList                      ProtocolIeID = 85
	ProtocolIeIDSCellFailedtoSetupModItem                      ProtocolIeID = 86
	ProtocolIeIDRRCReconfigurationCompleteIndicator            ProtocolIeID = 87
	ProtocolIeIDCellsStatusItem                                ProtocolIeID = 88
	ProtocolIeIDCellsStatusList                                ProtocolIeID = 89
	ProtocolIeIDCandidateSpCellList                            ProtocolIeID = 90
	ProtocolIeIDCandidateSpCellItem                            ProtocolIeID = 91
	ProtocolIeIDPotentialSpCellList                            ProtocolIeID = 92
	ProtocolIeIDPotentialSpCellItem                            ProtocolIeID = 93
	ProtocolIeIDFullConfiguration                              ProtocolIeID = 94
	ProtocolIeIDCRNTI                                          ProtocolIeID = 95
	ProtocolIeIDSpCellULConfigured                             ProtocolIeID = 96
	ProtocolIeIDInactivityMonitoringRequest                    ProtocolIeID = 97
	ProtocolIeIDInactivityMonitoringResponse                   ProtocolIeID = 98
	ProtocolIeIDDRBActivityItem                                ProtocolIeID = 99
	ProtocolIeIDDRBActivityList                                ProtocolIeID = 100
	ProtocolIeIDEUTRANRCellResourceCoordinationReqContainer    ProtocolIeID = 101
	ProtocolIeIDEUTRANRCellResourceCoordinationReqAckContainer ProtocolIeID = 102
	ProtocolIeIDProtectedEUTRAResourcesList                    ProtocolIeID = 105
	ProtocolIeIDRequestType                                    ProtocolIeID = 106
	ProtocolIeIDServCellIndex                                  ProtocolIeID = 107
	ProtocolIeIDRATFrequencyPriorityInformation                ProtocolIeID = 108
	ProtocolIeIDExecuteDuplication                             ProtocolIeID = 109
	ProtocolIeIDNRCGI                                          ProtocolIeID = 111
	ProtocolIeIDPagingCellItem                                 ProtocolIeID = 112
	ProtocolIeIDPagingCellList                                 ProtocolIeID = 113
	ProtocolIeIDPagingDRX                                      ProtocolIeID = 114
	ProtocolIeIDPagingPriority                                 ProtocolIeID = 115
	ProtocolIeIDSItypeList                                     ProtocolIeID = 116
	ProtocolIeIDUEIdentityIndexValue                           ProtocolIeID = 117
	ProtocolIeIDgNBCUSystemInformation                         ProtocolIeID = 118
	ProtocolIeIDHandoverPreparationInformation                 ProtocolIeID = 119
	ProtocolIeIDGNBCUTNLAssociationToAddItem                   ProtocolIeID = 120
	ProtocolIeIDGNBCUTNLAssociationToAddList                   ProtocolIeID = 121
	ProtocolIeIDGNBCUTNLAssociationToRemoveItem                ProtocolIeID = 122
	ProtocolIeIDGNBCUTNLAssociationToRemoveList                ProtocolIeID = 123
	ProtocolIeIDGNBCUTNLAssociationToUpdateItem                ProtocolIeID = 124
	ProtocolIeIDGNBCUTNLAssociationToUpdateList                ProtocolIeID = 125
	ProtocolIeIDMaskedIMEISV                                   ProtocolIeID = 126
	ProtocolIeIDPagingIdentity                                 ProtocolIeID = 127
	ProtocolIeIDDUtoCURRCContainer                             ProtocolIeID = 128
	ProtocolIeIDCellstobeBarredList                            ProtocolIeID = 129
	ProtocolIeIDCellstobeBarredItem                            ProtocolIeID = 130
	ProtocolIeIDTAISliceSupportList                            ProtocolIeID = 131
	ProtocolIeIDGNBCUTNLAssociationSetupList                   ProtocolIeID = 132
	ProtocolIeIDGNBCUTNLAssociationSetupItem                   ProtocolIeID = 133
	ProtocolIeIDGNBCUTNLAssociationFailedToSetupList           ProtocolIeID = 134
	ProtocolIeIDGNBCUTNLAssociationFailedToSetupItem           ProtocolIeID = 135
	ProtocolIeIDDRBNotifyItem                                  ProtocolIeID = 136
	ProtocolIeIDDRBNotifyList                                  ProtocolIeID = 137
	ProtocolIeIDNotficationControl                             ProtocolIeID = 138
	ProtocolIeIDRANAC                                          ProtocolIeID = 139
	ProtocolIeIDPWSSystemInformation                           ProtocolIeID = 140
	ProtocolIeIDRepetitionPeriod                               ProtocolIeID = 141
	ProtocolIeIDNumberofBroadcastRequest                       ProtocolIeID = 142
	ProtocolIeIDCellsToBeBroadcastList                         ProtocolIeID = 144
	ProtocolIeIDCellsToBeBroadcastItem                         ProtocolIeID = 145
	ProtocolIeIDCellsBroadcastCompletedList                    ProtocolIeID = 146
	ProtocolIeIDCellsBroadcastCompletedItem                    ProtocolIeID = 147
	ProtocolIeIDBroadcastToBeCancelledList                     ProtocolIeID = 148
	ProtocolIeIDBroadcastToBeCancelledItem                     ProtocolIeID = 149
	ProtocolIeIDCellsBroadcastCancelledList                    ProtocolIeID = 150
	ProtocolIeIDCellsBroadcastCancelledItem                    ProtocolIeID = 151
	ProtocolIeIDNRCGIListForRestartList                        ProtocolIeID = 152
	ProtocolIeIDNRCGIListForRestartItem                        ProtocolIeID = 153
	ProtocolIeIDPWSFailedNRCGIList                             ProtocolIeID = 154
	ProtocolIeIDPWSFailedNRCGIItem                             ProtocolIeID = 155
	ProtocolIeIDConfirmedUEID                                  ProtocolIeID = 156
	ProtocolIeIDCancelallWarningMessagesIndicator              ProtocolIeID = 157
	ProtocolIeIDGNBDUUEAMBRUL                                  ProtocolIeID = 158
	ProtocolIeIDDRXConfigurationIndicator                      ProtocolIeID = 159
	ProtocolIeIDRLCStatus                                      ProtocolIeID = 160
	ProtocolIeIDDLPDCPSNLength                                 ProtocolIeID = 161
	ProtocolIeIDGNBDUConfigurationQuery                        ProtocolIeID = 162
	ProtocolIeIDMeasurementTimingConfiguration                 ProtocolIeID = 163
	ProtocolIeIDDRBInformation                                 ProtocolIeID = 164
	ProtocolIeIDServingPLMN                                    ProtocolIeID = 165
	ProtocolIeIDProtectedEUTRAResourcesItem                    ProtocolIeID = 168
	ProtocolIeIDGNBCURRCVersion                                ProtocolIeID = 170
	ProtocolIeIDGNBDURRCVersion                                ProtocolIeID = 171
	ProtocolIeIDGNBDUOverloadInformation                       ProtocolIeID = 172
	ProtocolIeIDCellGroupConfig                                ProtocolIeID = 173
	ProtocolIeIDRLCFailureIndication                           ProtocolIeID = 174
	ProtocolIeIDUplinkTxDirectCurrentListInformation           ProtocolIeID = 175
	ProtocolIeIDDCBasedDuplicationConfigured                   ProtocolIeID = 176
	ProtocolIeIDDCBasedDuplicationActivation                   ProtocolIeID = 177
	ProtocolIeIDSULAccessIndication                            ProtocolIeID = 178
	ProtocolIeIDAvailablePLMNList                              ProtocolIeID = 179
	ProtocolIeIDPDUSessionID                                   ProtocolIeID = 180
	ProtocolIeIDULPDUSessionAggregateMaximumBitRate            ProtocolIeID = 181
	ProtocolIeIDServingCellMO                                  ProtocolIeID = 182
	ProtocolIeIDQoSFlowMappingIndication                       ProtocolIeID = 183
	ProtocolIeIDRRCDeliveryStatusRequest                       ProtocolIeID = 184
	ProtocolIeIDRRCDeliveryStatus                              ProtocolIeID = 185
	ProtocolIeIDBearerTypeChange                               ProtocolIeID = 186
	ProtocolIeIDRLCMode                                        ProtocolIeID = 187
	ProtocolIeIDDuplicationActivation                          ProtocolIeID = 188
	ProtocolIeIDDedicatedSIDeliveryNeededUEList                ProtocolIeID = 189
	ProtocolIeIDDedicatedSIDeliveryNeededUEItem                ProtocolIeID = 190
	ProtocolIeIDDRXLongCycleStartOffset                        ProtocolIeID = 191
	ProtocolIeIDULPDCPSNLength                                 ProtocolIeID = 192
	ProtocolIeIDSelectedBandCombinationIndex                   ProtocolIeID = 193
	ProtocolIeIDSelectedFeatureSetEntryIndex                   ProtocolIeID = 194
	ProtocolIeIDResourceCoordinationTransferInformation        ProtocolIeID = 195
	ProtocolIeIDExtendedServedPLMNsList                        ProtocolIeID = 196
	ProtocolIeIDExtendedAvailablePLMNList                      ProtocolIeID = 197
	ProtocolIeIDAssociatedSCellList                            ProtocolIeID = 198
	ProtocolIeIDlatestRRCVersionEnhanced                       ProtocolIeID = 199
	ProtocolIeIDAssociatedSCellItem                            ProtocolIeID = 200
	ProtocolIeIDCellDirection                                  ProtocolIeID = 201
	ProtocolIeIDSRBsSetupList                                  ProtocolIeID = 202
	ProtocolIeIDSRBsSetupItem                                  ProtocolIeID = 203
	ProtocolIeIDSRBsSetupModList                               ProtocolIeID = 204
	ProtocolIeIDSRBsSetupModItem                               ProtocolIeID = 205
	ProtocolIeIDSRBsModifiedList                               ProtocolIeID = 206
	ProtocolIeIDSRBsModifiedItem                               ProtocolIeID = 207
	ProtocolIeIDPhInfoSCG                                      ProtocolIeID = 208
	ProtocolIeIDRequestedBandCombinationIndex                  ProtocolIeID = 209
	ProtocolIeIDRequestedFeatureSetEntryIndex                  ProtocolIeID = 210
	ProtocolIeIDRequestedPMaxFR2                               ProtocolIeID = 211
	ProtocolIeIDDRXConfig                                      ProtocolIeID = 212
	ProtocolIeIDIgnoreResourceCoordinationContainer            ProtocolIeID = 213
	ProtocolIeIDUEAssistanceInformation                        ProtocolIeID = 214
	ProtocolIeIDNeedforGap                                     ProtocolIeID = 215
	ProtocolIeIDPagingOrigin                                   ProtocolIeID = 216
	ProtocolIeIDnewgNBCUUEF1APID                               ProtocolIeID = 217
	ProtocolIeIDRedirectedRRCmessage                           ProtocolIeID = 218
	ProtocolIeIDnewgNBDUUEF1APID                               ProtocolIeID = 219
	ProtocolIeIDNotificationInformation                        ProtocolIeID = 220
	ProtocolIeIDPLMNAssistanceInfoForNetShar                   ProtocolIeID = 221
	ProtocolIeIDUEContextNotRetrievable                        ProtocolIeID = 222
	ProtocolIeIDBPLMNIDInfoList                                ProtocolIeID = 223
	ProtocolIeIDSelectedPLMNID                                 ProtocolIeID = 224
	ProtocolIeIDUACAssistanceInfo                              ProtocolIeID = 225
	ProtocolIeIDRANUEID                                        ProtocolIeID = 226
	ProtocolIeIDGNBDUTNLAssociationToRemoveItem                ProtocolIeID = 227
	ProtocolIeIDGNBDUTNLAssociationToRemoveList                ProtocolIeID = 228
	ProtocolIeIDTNLAssociationTransportLayerAddressgNBDU       ProtocolIeID = 229
	ProtocolIeIDportNumber                                     ProtocolIeID = 230
	ProtocolIeIDAdditionalSIBMessageList                       ProtocolIeID = 231
	ProtocolIeIDCellType                                       ProtocolIeID = 232
	ProtocolIeIDIgnorePRACHConfiguration                       ProtocolIeID = 233
	ProtocolIeIDCGConfig                                       ProtocolIeID = 234
	ProtocolIeIDPDCCHBlindDetectionSCG                         ProtocolIeID = 235
	ProtocolIeIDRequestedPDCCHBlindDetectionSCG                ProtocolIeID = 236
	ProtocolIeIDPhInfoMCG                                      ProtocolIeID = 237
	ProtocolIeIDMeasGapSharingConfig                           ProtocolIeID = 238
	ProtocolIeIDsystemInformationAreaID                        ProtocolIeID = 239
	ProtocolIeIDareaScope                                      ProtocolIeID = 240
	ProtocolIeIDRRCContainerRRCSetupComplete                   ProtocolIeID = 241
	ProtocolIeIDTraceActivation                                ProtocolIeID = 242
	ProtocolIeIDTraceID                                        ProtocolIeID = 243
	ProtocolIeIDNeighbourCellInformationList                   ProtocolIeID = 244
	ProtocolIeIDSymbolAllocInSlot                              ProtocolIeID = 246
	ProtocolIeIDNumDLULSymbols                                 ProtocolIeID = 247
	ProtocolIeIDAdditionalRRMPriorityIndex                     ProtocolIeID = 248
	ProtocolIeIDDUCURadioInformationType                       ProtocolIeID = 249
	ProtocolIeIDCUDURadioInformationType                       ProtocolIeID = 250
	ProtocolIeIDAggressorgNBSetID                              ProtocolIeID = 251
	ProtocolIeIDVictimgNBSetID                                 ProtocolIeID = 252
	ProtocolIeIDLowerLayerPresenceStatusChange                 ProtocolIeID = 253
	ProtocolIeIDTransportLayerAddressInfo                      ProtocolIeID = 254
	ProtocolIeIDNeighbourCellInformationItem                   ProtocolIeID = 255
	ProtocolIeIDIntendedTDDDLULConfig                          ProtocolIeID = 256
	ProtocolIeIDQosMonitoringRequest                           ProtocolIeID = 257
	ProtocolIeIDBHChannelsToBeSetupList                        ProtocolIeID = 258
	ProtocolIeIDBHChannelsToBeSetupItem                        ProtocolIeID = 259
	ProtocolIeIDBHChannelsSetupList                            ProtocolIeID = 260
	ProtocolIeIDBHChannelsSetupItem                            ProtocolIeID = 261
	ProtocolIeIDBHChannelsToBeModifiedItem                     ProtocolIeID = 262
	ProtocolIeIDBHChannelsToBeModifiedList                     ProtocolIeID = 263
	ProtocolIeIDBHChannelsToBeReleasedItem                     ProtocolIeID = 264
	ProtocolIeIDBHChannelsToBeReleasedList                     ProtocolIeID = 265
	ProtocolIeIDBHChannelsToBeSetupModItem                     ProtocolIeID = 266
	ProtocolIeIDBHChannelsToBeSetupModList                     ProtocolIeID = 267
	ProtocolIeIDBHChannelsFailedToBeModifiedItem               ProtocolIeID = 268
	ProtocolIeIDBHChannelsFailedToBeModifiedList               ProtocolIeID = 269
	ProtocolIeIDBHChannelsFailedToBeSetupModItem               ProtocolIeID = 270
	ProtocolIeIDBHChannelsFailedToBeSetupModList               ProtocolIeID = 271
	ProtocolIeIDBHChannelsModifiedItem                         ProtocolIeID = 272
	ProtocolIeIDBHChannelsModifiedList                         ProtocolIeID = 273
	ProtocolIeIDBHChannelsSetupModItem                         ProtocolIeID = 274
	ProtocolIeIDBHChannelsSetupModList                         ProtocolIeID = 275
	ProtocolIeIDBHChannelsRequiredToBeReleasedItem             ProtocolIeID = 276
	ProtocolIeIDBHChannelsRequiredToBeReleasedList             ProtocolIeID = 277
	ProtocolIeIDBHChannelsFailedToBeSetupItem                  ProtocolIeID = 278
	ProtocolIeIDBHChannelsFailedToBeSetupList                  ProtocolIeID = 279
	ProtocolIeIDBHInfo                                         ProtocolIeID = 280
	ProtocolIeIDBAPAddress                                     ProtocolIeID = 281
	ProtocolIeIDConfiguredBAPAddress                           ProtocolIeID = 282
	ProtocolIeIDBHRoutingInformationAddedList                  ProtocolIeID = 283
	ProtocolIeIDBHRoutingInformationAddedListItem              ProtocolIeID = 284
	ProtocolIeIDBHRoutingInformationRemovedList                ProtocolIeID = 285
	ProtocolIeIDBHRoutingInformationRemovedListItem            ProtocolIeID = 286
	ProtocolIeIDULBHNonUPTrafficMapping                        ProtocolIeID = 287
	ProtocolIeIDActivatedCellstobeUpdatedList                  ProtocolIeID = 288
	ProtocolIeIDChildNodesList                                 ProtocolIeID = 289
	ProtocolIeIDIABInfoIABDU                                   ProtocolIeID = 290
	ProtocolIeIDIABInfoIABdonorCU                              ProtocolIeID = 291
	ProtocolIeIDIABTNLAddressesToRemoveList                    ProtocolIeID = 292
	ProtocolIeIDIABTNLAddressesToRemoveItem                    ProtocolIeID = 293
	ProtocolIeIDIABAllocatedTNLAddressList                     ProtocolIeID = 294
	ProtocolIeIDIABAllocatedTNLAddressItem                     ProtocolIeID = 295
	ProtocolIeIDIABIPv6RequestType                             ProtocolIeID = 296
	ProtocolIeIDIABv4AddressesRequested                        ProtocolIeID = 297
	ProtocolIeIDIABBarred                                      ProtocolIeID = 298
	ProtocolIeIDTrafficMappingInformation                      ProtocolIeID = 299
	ProtocolIeIDULUPTNLInformationtoUpdateList                 ProtocolIeID = 300
	ProtocolIeIDULUPTNLInformationtoUpdateListItem             ProtocolIeID = 301
	ProtocolIeIDULUPTNLAddresstoUpdateList                     ProtocolIeID = 302
	ProtocolIeIDULUPTNLAddresstoUpdateListItem                 ProtocolIeID = 303
	ProtocolIeIDDLUPTNLAddresstoUpdateList                     ProtocolIeID = 304
	ProtocolIeIDDLUPTNLAddresstoUpdateListItem                 ProtocolIeID = 305
	ProtocolIeIDNRV2XServicesAuthorized                        ProtocolIeID = 306
	ProtocolIeIDLTEV2XServicesAuthorized                       ProtocolIeID = 307
	ProtocolIeIDNRUESidelinkAggregateMaximumBitrate            ProtocolIeID = 308
	ProtocolIeIDLTEUESidelinkAggregateMaximumBitrate           ProtocolIeID = 309
	ProtocolIeIDSIB12message                                   ProtocolIeID = 310
	ProtocolIeIDSIB13message                                   ProtocolIeID = 311
	ProtocolIeIDSIB14message                                   ProtocolIeID = 312
	ProtocolIeIDSLDRBsFailedToBeModifiedItem                   ProtocolIeID = 313
	ProtocolIeIDSLDRBsFailedToBeModifiedList                   ProtocolIeID = 314
	ProtocolIeIDSLDRBsFailedToBeSetupItem                      ProtocolIeID = 315
	ProtocolIeIDSLDRBsFailedToBeSetupList                      ProtocolIeID = 316
	ProtocolIeIDSLDRBsModifiedItem                             ProtocolIeID = 317
	ProtocolIeIDSLDRBsModifiedList                             ProtocolIeID = 318
	ProtocolIeIDSLDRBsRequiredToBeModifiedItem                 ProtocolIeID = 319
	ProtocolIeIDSLDRBsRequiredToBeModifiedList                 ProtocolIeID = 320
	ProtocolIeIDSLDRBsRequiredToBeReleasedItem                 ProtocolIeID = 321
	ProtocolIeIDSLDRBsRequiredToBeReleasedList                 ProtocolIeID = 322
	ProtocolIeIDSLDRBsSetupItem                                ProtocolIeID = 323
	ProtocolIeIDSLDRBsSetupList                                ProtocolIeID = 324
	ProtocolIeIDSLDRBsToBeModifiedItem                         ProtocolIeID = 325
	ProtocolIeIDSLDRBsToBeModifiedList                         ProtocolIeID = 326
	ProtocolIeIDSLDRBsToBeReleasedItem                         ProtocolIeID = 327
	ProtocolIeIDSLDRBsToBeReleasedList                         ProtocolIeID = 328
	ProtocolIeIDSLDRBsToBeSetupItem                            ProtocolIeID = 329
	ProtocolIeIDSLDRBsToBeSetupList                            ProtocolIeID = 330
	ProtocolIeIDSLDRBsToBeSetupModItem                         ProtocolIeID = 331
	ProtocolIeIDSLDRBsToBeSetupModList                         ProtocolIeID = 332
	ProtocolIeIDSLDRBsSetupModList                             ProtocolIeID = 333
	ProtocolIeIDSLDRBsFailedToBeSetupModList                   ProtocolIeID = 334
	ProtocolIeIDSLDRBsSetupModItem                             ProtocolIeID = 335
	ProtocolIeIDSLDRBsFailedToBeSetupModItem                   ProtocolIeID = 336
	ProtocolIeIDSLDRBsModifiedConfList                         ProtocolIeID = 337
	ProtocolIeIDSLDRBsModifiedConfItem                         ProtocolIeID = 338
	ProtocolIeIDUEAssistanceInformationEUTRA                   ProtocolIeID = 339
	ProtocolIeIDPC5LinkAMBR                                    ProtocolIeID = 340
	ProtocolIeIDSLPHYMACRLCConfig                              ProtocolIeID = 341
	ProtocolIeIDSLConfigDedicatedEUTRAInfo                     ProtocolIeID = 342
	ProtocolIeIDAlternativeQoSParaSetList                      ProtocolIeID = 343
	ProtocolIeIDCurrentQoSParaSetIndex                         ProtocolIeID = 344
	ProtocolIeIDgNBCUMeasurementID                             ProtocolIeID = 345
	ProtocolIeIDgNBDUMeasurementID                             ProtocolIeID = 346
	ProtocolIeIDRegistrationRequest                            ProtocolIeID = 347
	ProtocolIeIDReportCharacteristics                          ProtocolIeID = 348
	ProtocolIeIDCellToReportList                               ProtocolIeID = 349
	ProtocolIeIDCellMeasurementResultList                      ProtocolIeID = 350
	ProtocolIeIDHardwareLoadIndicator                          ProtocolIeID = 351
	ProtocolIeIDReportingPeriodicity                           ProtocolIeID = 352
	ProtocolIeIDTNLCapacityIndicator                           ProtocolIeID = 353
	ProtocolIeIDCarrierList                                    ProtocolIeID = 354
	ProtocolIeIDULCarrierList                                  ProtocolIeID = 355
	ProtocolIeIDFrequencyShift7p5khz                           ProtocolIeID = 356
	ProtocolIeIDSSBPositionsInBurst                            ProtocolIeID = 357
	ProtocolIeIDNRPRACHConfig                                  ProtocolIeID = 358
	ProtocolIeIDRACHReportInformationList                      ProtocolIeID = 359
	ProtocolIeIDRLFReportInformationList                       ProtocolIeID = 360
	ProtocolIeIDTDDULDLConfigCommonNR                          ProtocolIeID = 361
	ProtocolIeIDCNPacketDelayBudgetDownlink                    ProtocolIeID = 362
	ProtocolIeIDExtendedPacketDelayBudget                      ProtocolIeID = 363
	ProtocolIeIDTSCTrafficCharacteristics                      ProtocolIeID = 364
	ProtocolIeIDReportingRequestType                           ProtocolIeID = 365
	ProtocolIeIDTimeReferenceInformation                       ProtocolIeID = 366
	ProtocolIeIDCNPacketDelayBudgetUplink                      ProtocolIeID = 369
	ProtocolIeIDAdditionalPDCPDuplicationTNLList               ProtocolIeID = 370
	ProtocolIeIDRLCDuplicationInformation                      ProtocolIeID = 371
	ProtocolIeIDAdditionalDuplicationIndication                ProtocolIeID = 372
	ProtocolIeIDConditionalInterDUMobilityInformation          ProtocolIeID = 373
	ProtocolIeIDConditionalIntraDUMobilityInformation          ProtocolIeID = 374
	ProtocolIeIDtargetCellsToCancel                            ProtocolIeID = 375
	ProtocolIeIDrequestedTargetCellGlobalID                    ProtocolIeID = 376
	ProtocolIeIDManagementBasedMDTPLMNList                     ProtocolIeID = 377
	ProtocolIeIDTraceCollectionEntityIPAddress                 ProtocolIeID = 378
	ProtocolIeIDPrivacyIndicator                               ProtocolIeID = 379
	ProtocolIeIDTraceCollectionEntityURI                       ProtocolIeID = 380
	ProtocolIeIDmdtConfiguration                               ProtocolIeID = 381
	ProtocolIeIDServingNID                                     ProtocolIeID = 382
	ProtocolIeIDNPNBroadcastInformation                        ProtocolIeID = 383
	ProtocolIeIDNPNSupportInfo                                 ProtocolIeID = 384
	ProtocolIeIDNID                                            ProtocolIeID = 385
	ProtocolIeIDAvailableSNPNIDList                            ProtocolIeID = 386
	ProtocolIeIDSIB10message                                   ProtocolIeID = 387
	ProtocolIeIDDLCarrierList                                  ProtocolIeID = 389
	ProtocolIeIDExtendedTAISliceSupportList                    ProtocolIeID = 390
	ProtocolIeIDRequestedSRSTransmissionCharacteristics        ProtocolIeID = 391
	ProtocolIeIDPosAssistanceInformation                       ProtocolIeID = 392
	ProtocolIeIDPosBroadcast                                   ProtocolIeID = 393
	ProtocolIeIDRoutingID                                      ProtocolIeID = 394
	ProtocolIeIDPosAssistanceInformationFailureList            ProtocolIeID = 395
	ProtocolIeIDPosMeasurementQuantities                       ProtocolIeID = 396
	ProtocolIeIDPosMeasurementResultList                       ProtocolIeID = 397
	ProtocolIeIDTRPInformationTypeListTRPReq                   ProtocolIeID = 398
	ProtocolIeIDTRPInformationTypeItem                         ProtocolIeID = 399
	ProtocolIeIDTRPInformationListTRPResp                      ProtocolIeID = 400
	ProtocolIeIDTRPInformationItem                             ProtocolIeID = 401
	ProtocolIeIDLMFMeasurementID                               ProtocolIeID = 402
	ProtocolIeIDSRSType                                        ProtocolIeID = 403
	ProtocolIeIDActivationTime                                 ProtocolIeID = 404
	ProtocolIeIDAbortTransmission                              ProtocolIeID = 405
	ProtocolIeIDPositioningBroadcastCells                      ProtocolIeID = 406
	ProtocolIeIDSRSConfiguration                               ProtocolIeID = 407
	ProtocolIeIDPosReportCharacteristics                       ProtocolIeID = 408
	ProtocolIeIDPosMeasurementPeriodicity                      ProtocolIeID = 409
	ProtocolIeIDTRPList                                        ProtocolIeID = 410
	ProtocolIeIDRANMeasurementID                               ProtocolIeID = 411
	ProtocolIeIDLMFUEMeasurementID                             ProtocolIeID = 412
	ProtocolIeIDRANUEMeasurementID                             ProtocolIeID = 413
	ProtocolIeIDECIDMeasurementQuantities                      ProtocolIeID = 414
	ProtocolIeIDECIDMeasurementQuantitiesItem                  ProtocolIeID = 415
	ProtocolIeIDECIDMeasurementPeriodicity                     ProtocolIeID = 416
	ProtocolIeIDECIDMeasurementResult                          ProtocolIeID = 417
	ProtocolIeIDCellPortionID                                  ProtocolIeID = 418
	ProtocolIeIDSFNInitialisationTime                          ProtocolIeID = 419
	ProtocolIeIDSystemFrameNumber                              ProtocolIeID = 420
	ProtocolIeIDSlotNumber                                     ProtocolIeID = 421
	ProtocolIeIDTRPMeasurementRequestList                      ProtocolIeID = 422
	ProtocolIeIDMeasurementBeamInfoRequest                     ProtocolIeID = 423
	ProtocolIeIDECIDReportCharacteristics                      ProtocolIeID = 424
	ProtocolIeIDConfiguredTACIndication                        ProtocolIeID = 425
	ProtocolIeIDExtendedGNBCUName                              ProtocolIeID = 426
	ProtocolIeIDExtendedGNBDUName                              ProtocolIeID = 427
	ProtocolIeIDF1CTransferPath                                ProtocolIeID = 428
	ProtocolIeIDSFNOffset                                      ProtocolIeID = 429
	ProtocolIeIDTransmissionStopIndicator                      ProtocolIeID = 430
	ProtocolIeIDSrsFrequency                                   ProtocolIeID = 431
	ProtocolIeIDSCGIndicator                                   ProtocolIeID = 432
	ProtocolIeIDEstimatedArrivalProbability                    ProtocolIeID = 433
	ProtocolIeIDTRPType                                        ProtocolIeID = 434
	ProtocolIeIDSRSSpatialRelationPerSRSResource               ProtocolIeID = 435
	ProtocolIeIDPDCPTerminatingNodeDLTNLAddrInfo               ProtocolIeID = 436
	ProtocolIeIDENBDLTNLAddress                                ProtocolIeID = 437
	ProtocolIeIDPosMeasurementPeriodicityExtended              ProtocolIeID = 438
	ProtocolIeIDPRSResourceID                                  ProtocolIeID = 439
	ProtocolIeIDLocationMeasurementInformation                 ProtocolIeID = 440
	ProtocolIeIDInterFrequencyConfigNoGap                      ProtocolIeID = 651
	ProtocolIeIDNeedForGapsInfoNR                              ProtocolIeID = 665
	ProtocolIeIDPosMeasurementPeriodicityNRAoA                 ProtocolIeID = 672
)
