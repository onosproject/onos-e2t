// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package server

import (
	"context"
	"encoding/hex"
	"fmt"
	f1apiv1 "github.com/onosproject/onos-e2t/api/f1ap/v1"
	f1appducontentsv1 "github.com/onosproject/onos-e2t/api/f1ap/v1/f1ap_pdu_contents"
	f1appdudescriptionsv1 "github.com/onosproject/onos-e2t/api/f1ap/v1/f1ap_pdu_descriptions"
	xnapiv1 "github.com/onosproject/onos-e2t/api/xnap/v1"
	xnapiesv1 "github.com/onosproject/onos-e2t/api/xnap/v1/xnap-ies"
	xnappdudescriptionsv1 "github.com/onosproject/onos-e2t/api/xnap/v1/xnap-pdu-descriptions"
	"github.com/onosproject/onos-e2t/pkg/southbound/f1ap/encoder"
	f1utils "github.com/onosproject/onos-e2t/pkg/southbound/f1ap/utils"
	encoder2 "github.com/onosproject/onos-e2t/pkg/southbound/xnap/encoder"
	xnutils "github.com/onosproject/onos-e2t/pkg/southbound/xnap/utils"
	"github.com/onosproject/onos-e2t/pkg/utils/decode"
	"time"

	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/stream"
	"github.com/onosproject/onos-lib-go/pkg/logging"

	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"

	prototypes "github.com/gogo/protobuf/types"

	"github.com/onosproject/onos-e2t/pkg/store/rnib"

	"github.com/onosproject/onos-e2t/api/e2ap/v2"
	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"

	e2smtypes "github.com/onosproject/onos-api/go/onos/e2t/e2sm"

	topoapi "github.com/onosproject/onos-api/go/onos/topo"
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
	"github.com/onosproject/onos-e2t/pkg/modelregistry"
	e2 "github.com/onosproject/onos-e2t/pkg/protocols/e2ap"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/pdubuilder"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/pdudecoder"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
)

var log = logging.GetLogger()

// TODO: Change the RIC ID to something appropriate
var ricID = types.RicIdentifier{
	RicIdentifierValue: []byte{0xDE, 0xBC, 0xA0},
	RicIdentifierLen:   20,
}

func NewE2Server(e2apConns E2APConnManager,
	mgmtConns MgmtConnManager,
	streams stream.Manager,
	modelRegistry modelregistry.ModelRegistry, rnib rnib.Store) *E2Server {
	return &E2Server{
		server:    e2.NewServer(),
		e2apConns: e2apConns,
		mgmtConns: mgmtConns,

		streams:       streams,
		modelRegistry: modelRegistry,
		rnib:          rnib,
	}
}

type E2Server struct {
	server        *e2.Server
	e2apConns     E2APConnManager
	mgmtConns     MgmtConnManager
	streams       stream.Manager
	modelRegistry modelregistry.ModelRegistry
	rnib          rnib.Store
}

func (s *E2Server) Serve() error {
	return s.server.Serve(func(conn e2.ServerConn) e2.ServerInterface {
		return &E2APServer{
			serverConn:    conn,
			e2apConns:     s.e2apConns,
			mgmtConns:     s.mgmtConns,
			streams:       s.streams,
			modelRegistry: s.modelRegistry,
			rnib:          s.rnib,
		}
	})
}

func (s *E2Server) Stop() error {
	return s.server.Stop()
}

type E2APServer struct {
	e2apConns     E2APConnManager
	mgmtConns     MgmtConnManager
	streams       stream.Manager
	serverConn    e2.ServerConn
	e2apConn      *E2APConn
	modelRegistry modelregistry.ModelRegistry
	rnib          rnib.Store
}

// uint24ToUint32 converts uint24 uint32
func uint24ToUint32(val []byte) uint32 {
	r := uint32(0)
	for i := uint32(0); i < 3; i++ {
		r |= uint32(val[i]) << (8 * i)
	}
	return r
}

func (e *E2APServer) E2Setup(ctx context.Context, request *e2appducontents.E2SetupRequest) (*e2appducontents.E2SetupResponse, *e2appducontents.E2SetupFailure, error) {
	log.Infof("Received E2 setup request: %+v", request)
	transID, nodeIdentity, ranFuncs, comps, err := pdudecoder.DecodeE2SetupRequest(request)
	if err != nil {
		cause := &e2apies.Cause{
			Cause: &e2apies.Cause_RicRequest{
				RicRequest: e2apies.CauseRicrequest_CAUSE_RICREQUEST_UNSPECIFIED,
			},
		}

		var trID int32
		for _, v := range request.GetProtocolIes() {
			if v.Id == int32(v2.ProtocolIeIDTransactionID) {
				trID = v.GetValue().GetTransactionId().GetValue()
			}
		}

		failure := &e2appducontents.E2SetupFailure{
			ProtocolIes: make([]*e2appducontents.E2SetupFailureIes, 0),
		}
		failure.SetErrorCause(cause).SetTransactionID(trID)

		return nil, failure, err
	}

	rawPlmnid := []byte{nodeIdentity.Plmn[0], nodeIdentity.Plmn[1], nodeIdentity.Plmn[2]}
	plmnID := fmt.Sprintf("%x", uint24ToUint32(rawPlmnid))

	var e2Cells []*topoapi.E2Cell
	serviceModels := make(map[string]*topoapi.ServiceModelInfo)
	rfAccepted := make(types.RanFunctionRevisions)
	rfRejected := make(types.RanFunctionCauses)
	plugins := e.modelRegistry.GetPlugins()

	for smOid, sm := range plugins {
		var ranFunctions []*prototypes.Any
		var ranFunctionIDs []uint32
		for ranFunctionID, ranFunc := range *ranFuncs {
			oid := e2smtypes.OID(ranFunc.OID)
			if smOid == oid {
				serviceModels[string(smOid)] = &topoapi.ServiceModelInfo{
					OID:          string(smOid),
					RanFunctions: ranFunctions,
				}
				ranFunctionIDs = append(ranFunctionIDs, uint32(ranFunctionID))
				if setup, ok := sm.(modelregistry.E2Setup); ok {
					onSetupRequest := &e2smtypes.OnSetupRequest{
						ServiceModels:          serviceModels,
						E2Cells:                &e2Cells,
						RANFunctionDescription: ranFunc.Description,
					}
					err := setup.OnSetup(onSetupRequest)
					if err != nil {
						log.Warn(err)
						log.Debugf("Length of RAN function Description Bytes is: %d", len(onSetupRequest.RANFunctionDescription))
						log.Debugf("RAN Function Description Bytes in hex format: %v", hex.Dump(onSetupRequest.RANFunctionDescription))
					}
				}
				rfAccepted[ranFunctionID] = ranFunc.Revision
				serviceModels[string(smOid)].RanFunctionIDs = ranFunctionIDs
			}
		}
	}

	// start e2 cell updates
	// convert e2cell slice to map so that this logic can do an e2 cell random access
	e2CellsMap := make(map[string]*topoapi.E2Cell)
	for i := 0; i < len(e2Cells); i++ {
		e2CellsMap[e2Cells[i].CellGlobalID.Value] = e2Cells[i]
	}

	f1SetupRequestMessages := make([]*f1appdudescriptionsv1.F1ApPDu, 0)
	xnSetupRequestMessages := make([]*xnappdudescriptionsv1.XnApPDu, 0)
	for _, c := range comps {
		switch c.E2NodeComponentType {
		case e2ap_ies.E2NodeComponentInterfaceType_E2NODE_COMPONENT_INTERFACE_TYPE_F1:
			f1SetupReqPer := c.E2NodeComponentConfiguration.E2NodeComponentRequestPart
			f1SetupReq, err := encoder.PerDecodeF1ApPdu(f1SetupReqPer)
			if err != nil {
				log.Warnf("f1 setup request decode fail: %+v", err)
				continue
			}
			f1SetupRequestMessages = append(f1SetupRequestMessages, f1SetupReq)
			// todo add f1setup response message
		case e2ap_ies.E2NodeComponentInterfaceType_E2NODE_COMPONENT_INTERFACE_TYPE_XN:
			xnSetupReqPer := c.E2NodeComponentConfiguration.E2NodeComponentRequestPart
			xnSetupReq, err := encoder2.PerDecodeXnApPdu(xnSetupReqPer)
			if err != nil {
				log.Warnf("xn setup request decode fail: %+v", err)
				continue
			}
			xnSetupRequestMessages = append(xnSetupRequestMessages, xnSetupReq)
			// todo add xnsetup response message
		default:
			log.Warnf("E2 Node Component Type %+v does not support", c.E2NodeComponentType)
		}
	}

	// f1 xn request messages handling logic
	// todo should be removed
	for _, f1msg := range f1SetupRequestMessages {
		log.Debugf("F1: %+v", f1msg)
		if f1msg.GetInitiatingMessage() == nil || f1msg.GetInitiatingMessage().GetValue() == nil || f1msg.GetInitiatingMessage().GetValue().GetF1SetupRequest() == nil {
			log.Warn("f1 setup request is nil")
			continue
		}
		f1IEs := f1msg.GetInitiatingMessage().GetValue().GetF1SetupRequest().ProtocolIes
		var f1TransactionID *int32
		var f1GnbDuID *int64
		var f1GnbDuRRCVersion *uint64
		var f1ServCellList []*f1appducontentsv1.GnbDUServedCellsItemIes

		for _, ie := range f1IEs {
			switch ie.Id {
			case int32(f1apiv1.ProtocolIeIDTransactionID):
				// transaction ID
				if ie.GetValue() == nil || ie.GetValue().GetTransactionId() == nil {
					log.Warn("[F1 Setup Request] transaction ID is nil")
					continue
				}
				f1TransactionID = &ie.GetValue().GetTransactionId().Value
				log.Debugf("[F1 Setup Request] F1 TransactionID: %+v", *f1TransactionID)
			case int32(f1apiv1.ProtocolIeIDgNBDUID):
				// gnb du id
				if ie.GetValue() == nil || ie.GetValue().GetGnbDuId() == nil {
					log.Warn("[F1 Setup Request] gNB DU ID is nil")
					continue
				}
				f1GnbDuID = &ie.GetValue().GetGnbDuId().Value
				log.Debugf("[F1 Setup Request] F1 gNB DU ID: %+v", *f1GnbDuID)
			case int32(f1apiv1.ProtocolIeIDGNBDURRCVersion):
				// gnb du rrc version
				if ie.GetValue() == nil || ie.GetValue().GetRrcVersion() == nil || ie.GetValue().GetRrcVersion().LatestRrcVersion == nil {
					log.Warn("[F1 Setup Request] RRC Version is nil")
					continue
				}
				f1GnbDuRRCVersion = decode.Asn1BitstringToUint64(ie.GetValue().GetRrcVersion().LatestRrcVersion)
				log.Debugf("[F1 Setup Request] F1 gNB DU RRC version: %+v", *f1GnbDuRRCVersion)
			case int32(f1apiv1.ProtocolIeIDgNBDUServedCellsList):
				// scell list
				if ie.GetValue() == nil || ie.GetValue().GetGnbDuServedCellsList() == nil {
					log.Warn("[F1 Setup Request] gNB DU served cell list is nil")
					continue
				}
				f1ServCellList = ie.GetValue().GetGnbDuServedCellsList().Value
				log.Debugf("[F1 Setup Request] F1 gNB DU Served Cell List: %+v", f1ServCellList)
			default:
				log.Warn("[F1 Setup Request] received unsupported F1 IE: %+v", ie.Id)
			}
		}

		for _, cell := range f1ServCellList {
			cgi, err := f1utils.GetCGICellID(cell)
			if err != nil {
				log.Warnf("[F1 Setup Request] Served cell does not have Cell ID: %+v", err)
				continue
			}

			var e2Cell *topoapi.E2Cell
			if _, ok := e2CellsMap[cgi]; !ok {
				e2Cell = &topoapi.E2Cell{
					CellGlobalID: &topoapi.CellGlobalID{
						Type:  topoapi.CellGlobalIDType_NRCGI, // todo 5g only now; should be updated to support 4g as well
						Value: cgi,
					},
				}
				e2Cells = append(e2Cells, e2Cell)
				e2CellsMap[cgi] = e2Cell
			} else {
				e2Cell = e2CellsMap[cgi]
			}

			plmnid, err := f1utils.GetCGIPlmnID(cell)
			if err == nil {
				e2Cell.PlmnId = uint32(plmnid)
			}

			pci, err := f1utils.GetPCI(cell)
			if err == nil {
				e2Cell.PCI = pci
			}

			if f1GnbDuID != nil {
				e2Cell.GnbDuId = uint32(*f1GnbDuID)
			}
			if f1GnbDuRRCVersion != nil {
				e2Cell.LatestRrcVersion = uint32(*f1GnbDuRRCVersion)
			}

			e2Cell.ServedPlmns = f1utils.GetServedPlmns(cell)
			tddInfo, err := f1utils.GetTDDInfo(cell)
			if err == nil {
				e2Cell.NrModeInfo = tddInfo
			}
			fddInfo, err := f1utils.GetFDDInfo(cell)
			if err == nil {
				e2Cell.NrModeInfo = fddInfo
			}

			e2Cell.MeasurementTimingConfiguration = uint32(f1utils.GetMeasurementTimingConfiguration(cell))
		}
	}
	for _, xnmsg := range xnSetupRequestMessages {
		log.Debugf("Xn: %+v", xnmsg)
		if xnmsg.GetInitiatingMessage() == nil || xnmsg.GetInitiatingMessage().GetValue() == nil || xnmsg.GetInitiatingMessage().GetValue().GetXnSetupRequest() == nil {
			log.Warn("xn setup request message is nil")
			continue
		}
		xnIEs := xnmsg.GetInitiatingMessage().GetValue().GetXnSetupRequest().ProtocolIes
		var xnGlobalNGRANnodeID *xnapiesv1.GlobalNgRAnnodeID
		var xnTAISupportList []*xnapiesv1.TaisupportItem
		var xnAMFRegionInformation []*xnapiesv1.GlobalAmfRegionInformation
		var xnListofServedCellsNR []*xnapiesv1.ServedCellsNRItem

		for _, ie := range xnIEs {
			switch ie.Id {
			case int32(xnapiv1.ProtocolIeIDGlobalNGRANnodeID):
				// GlobalNGRANnodeID
				if ie.GetValue() == nil || ie.GetValue().GetIdGlobalNgRanNodeId() == nil {
					log.Warn("[Xn Setup Request] global ng ran node id is nil")
					continue
				}
				xnGlobalNGRANnodeID = ie.GetValue().GetIdGlobalNgRanNodeId()
				log.Debugf("[Xn Setup Request] Xn Global NG RAN node ID: %+v", xnGlobalNGRANnodeID)
			case int32(xnapiv1.ProtocolIeIDTAISupportlist):
				// TAISupportlist
				if ie.GetValue() == nil || ie.GetValue().GetIdTaisupportList() == nil {
					log.Warn("[Xn Setup Request] tai support list is nil")
					continue
				}
				xnTAISupportList = ie.GetValue().GetIdTaisupportList().GetValue()
				log.Debugf("[Xn Setup Request] Xn TAI support list: %+v", xnTAISupportList)
			case int32(xnapiv1.ProtocolIeIDAMFRegionInformation):
				// AMFRegionInformation
				if ie.GetValue() == nil || ie.GetValue().GetIdAmfRegionInformation() == nil {
					log.Warn("[Xn Setup Request] amf region information is nil")
					continue
				}
				xnAMFRegionInformation = ie.GetValue().GetIdAmfRegionInformation().GetValue()
				log.Debugf("[Xn Setup Request] Xn AMF Region information: %+v", xnAMFRegionInformation)
			case int32(xnapiv1.ProtocolIeIDListofservedcellsNR):
				// ListofservedcellsNR
				if ie.GetValue() == nil || ie.GetValue().GetIdListOfServedCellsNr() == nil {
					log.Warn("[Xn Setup Request] id list of served cell nr is nil")
					continue
				}
				xnListofServedCellsNR = ie.GetValue().GetIdListOfServedCellsNr().GetValue()
				log.Debugf("[Xn Setup Request] Xn List of Served Cells NR: %+v", xnListofServedCellsNR)
			default:
				log.Warnf("[Xn Setup Request] received unsupported Xn IE: %+v", ie.Id)
			}
		}

		for _, cell := range xnListofServedCellsNR {
			cgi, err := xnutils.GetCGICellID(cell)
			if err != nil {
				log.Warnf("[Xn Setup Request] Served cell does not have Cell ID: %+v", err)
				continue
			}

			var e2Cell *topoapi.E2Cell
			if _, ok := e2CellsMap[cgi]; !ok {
				e2Cell = &topoapi.E2Cell{
					CellGlobalID: &topoapi.CellGlobalID{
						Type:  topoapi.CellGlobalIDType_NRCGI, // todo 5g only now; should be updated to support 4g as well
						Value: cgi,
					},
				}
				e2Cells = append(e2Cells, e2Cell)
				e2CellsMap[cgi] = e2Cell
			} else {
				e2Cell = e2CellsMap[cgi]
			}

			plmnid, err := xnutils.GetCGIPlmnID(cell)
			if err == nil {
				e2Cell.PlmnId = uint32(plmnid)
			}

			pci, err := xnutils.GetPCI(cell)
			if err == nil {
				e2Cell.PCI = pci
			}

			// global ng ran node ID
			globalNGRanNodeID, err := xnutils.GetNGRanNodeID(xnGlobalNGRANnodeID)
			if err == nil {
				e2Cell.GlobalNgRanNodeId = globalNGRanNodeID
			}

			// tai support list
			taiSupportList, err := xnutils.GetTAISupportList(xnTAISupportList)
			if err == nil {
				e2Cell.TaiSupportList = taiSupportList
			}

			// amf region information
			amfRegionList, err := xnutils.GetAMFRegionList(xnAMFRegionInformation)
			if err == nil {
				e2Cell.AmfRegionInformation = amfRegionList
			}

			tddInfo, err := xnutils.GetTDDInfo(cell)
			if err == nil {
				e2Cell.NrModeInfo = tddInfo
			}
			fddInfo, err := xnutils.GetFDDInfo(cell)
			if err == nil {
				e2Cell.NrModeInfo = fddInfo
			}

			e2Cell.MeasurementTimingConfiguration = uint32(xnutils.GetMeasurementTimingConfiguration(cell))

			// neighbor cell info
			neighbors, err := xnutils.GetNeighborInfoNR(cell)
			if err == nil {
				for _, n := range neighbors {
					topoNeighbor, err := xnutils.GetTopoNeighborInformationNR(n)
					if err == nil {
						e2Cell.NeighborInformationNrs = append(e2Cell.NeighborInformationNrs, topoNeighbor)
					}
				}
			}
			// todo implmenet neighbor cell EUTRA
		}
	}

	mgmtConn := NewMgmtConn(createE2NodeURI(nodeIdentity), plmnID, nodeIdentity, e.serverConn, serviceModels, e2Cells, time.Now())

	// Create an E2 setup response
	e2ncID3 := pdubuilder.CreateE2NodeComponentIDS1("S1-component")
	e2nccaal := make([]*types.E2NodeComponentConfigAdditionAckItem, 0)
	ie1 := types.E2NodeComponentConfigAdditionAckItem{
		E2NodeComponentConfigurationAck: e2ap_ies.E2NodeComponentConfigurationAck{
			UpdateOutcome: e2ap_ies.UpdateOutcome_UPDATE_OUTCOME_SUCCESS,
		},
		E2NodeComponentID:   e2ncID3,
		E2NodeComponentType: e2ap_ies.E2NodeComponentInterfaceType_E2NODE_COMPONENT_INTERFACE_TYPE_S1,
	}
	e2nccaal = append(e2nccaal, &ie1)
	response, err := pdubuilder.NewE2SetupResponse(*transID, nodeIdentity.Plmn, ricID, e2nccaal)
	if err != nil {
		cause := &e2apies.Cause{
			Cause: &e2apies.Cause_RicRequest{
				RicRequest: e2apies.CauseRicrequest_CAUSE_RICREQUEST_UNSPECIFIED,
			},
		}

		var trID int32
		for _, v := range request.GetProtocolIes() {
			if v.Id == int32(v2.ProtocolIeIDTransactionID) {
				trID = v.GetValue().GetTransactionId().GetValue()
			}
		}

		failure := &e2appducontents.E2SetupFailure{
			ProtocolIes: make([]*e2appducontents.E2SetupFailureIes, 0),
		}
		failure.SetErrorCause(cause).SetTransactionID(trID)

		return nil, failure, err
	}

	if len(rfAccepted) > 0 {
		response.SetRanFunctionAccepted(rfAccepted)
	}
	if len(rfRejected) > 0 {
		response.SetRanFunctionRejected(rfRejected)
	}
	log.Infof("Sending E2 setup response %+v", response)
	e.mgmtConns.open(mgmtConn)
	return response, nil, nil
}

func (e *E2APServer) RICIndication(ctx context.Context, request *e2appducontents.Ricindication) error {
	return e.e2apConn.ricIndication(ctx, request)
}

func (e *E2APServer) E2ConfigurationUpdate(ctx context.Context, request *e2appducontents.E2NodeConfigurationUpdate) (response *e2appducontents.E2NodeConfigurationUpdateAcknowledge, failure *e2appducontents.E2NodeConfigurationUpdateFailure, err error) {
	log.Infof("Received E2 node configuration update request: %+v", request)

	var nodeIdentity *e2apies.GlobalE2NodeId
	e2nccual := make([]*types.E2NodeComponentConfigUpdateItem, 0)
	for _, v := range request.GetProtocolIes() {
		if v.Id == int32(v2.ProtocolIeIDGlobalE2nodeID) {
			nodeIdentity = v.GetValue().GetGlobalE2NodeId()
		}
		if v.Id == int32(v2.ProtocolIeIDE2nodeComponentConfigUpdate) {
			list := v.GetValue().GetE2NodeComponentConfigUpdate().GetValue()
			for _, ie := range list {
				e2nccuai := types.E2NodeComponentConfigUpdateItem{}
				e2nccuai.E2NodeComponentType = ie.GetValue().GetE2NodeComponentConfigUpdateItem().GetE2NodeComponentInterfaceType()
				e2nccuai.E2NodeComponentID = ie.GetValue().GetE2NodeComponentConfigUpdateItem().GetE2NodeComponentId()
				e2nccuai.E2NodeComponentConfiguration = *ie.GetValue().GetE2NodeComponentConfigUpdateItem().GetE2NodeComponentConfiguration()

				e2nccual = append(e2nccual, &e2nccuai)
			}
		}
	}

	if nodeIdentity != nil {
		nodeID, err := pdudecoder.ExtractE2NodeIdentity(nodeIdentity, e2nccual)
		if err != nil {
			cause := &e2apies.Cause{
				Cause: &e2apies.Cause_RicRequest{
					RicRequest: e2apies.CauseRicrequest_CAUSE_RICREQUEST_UNSPECIFIED,
				},
			}

			var trID int32
			for _, v := range request.GetProtocolIes() {
				if v.Id == int32(v2.ProtocolIeIDTransactionID) {
					trID = v.GetValue().GetTransactionId().GetValue()
					break
				}
			}

			failure := &e2appducontents.E2NodeConfigurationUpdateFailure{
				ProtocolIes: make([]*e2appducontents.E2NodeConfigurationUpdateFailureIes, 0),
			}
			failure.SetCause(cause).SetTransactionID(trID)

			return nil, failure, nil
		}

		// Creates a new E2AP data connection
		e.e2apConn = NewE2APConn(createE2NodeURI(nodeID), e.serverConn, e.streams, e.rnib)
	}

	var trID int32
	for _, v := range request.GetProtocolIes() {
		if v.Id == int32(v2.ProtocolIeIDTransactionID) {
			trID = v.GetValue().GetTransactionId().GetValue()
			break
		}
	}

	e2ncua := &e2appducontents.E2NodeConfigurationUpdateAcknowledge{
		ProtocolIes: make([]*e2appducontents.E2NodeConfigurationUpdateAcknowledgeIes, 0),
	}
	e2ncua.SetTransactionID(trID)
	log.Debugf("Composed E2nodeConfigurationUpdateMessage is\n%v", e2ncua)
	log.Infof("Sending config update ack to e2 node: %s", e.e2apConn.E2NodeID)
	e.e2apConns.open(e.e2apConn)
	return e2ncua, nil, nil
}
