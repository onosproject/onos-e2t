// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package v1beta1

import (
	"context"
	"sync"

	"github.com/onosproject/onos-e2t/pkg/southbound/e2conn"

	"github.com/onosproject/onos-e2t/pkg/store/rnib"

	e2ap101pducontents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	topoapi "github.com/onosproject/onos-api/go/onos/topo"

	"github.com/onosproject/onos-e2t/pkg/oid"

	e2ap101ies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	e2ap101pdubuilder "github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/pdubuilder"

	"github.com/onosproject/onos-e2t/pkg/config"
	"github.com/onosproject/onos-e2t/pkg/modelregistry"
	e2ap101server "github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/server"
	e2ap101types "github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/types"
	"github.com/onosproject/onos-lib-go/pkg/errors"

	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/onosproject/onos-lib-go/pkg/northbound"
	"google.golang.org/grpc"
)

var log = logging.GetLogger("northbound", "e2", "v1beta1")

// NewControlService creates a new control service
func NewControlService(modelRegistry modelregistry.ModelRegistry, connections e2conn.ConnManager,
	oidRegistry oid.Registry, topo rnib.Store) northbound.Service {
	return &ControlService{
		modelRegistry: modelRegistry,
		connections:   connections,
		oidRegistry:   oidRegistry,
		topo:          topo,
	}
}

// ControlService is a Service implementation for control requests
type ControlService struct {
	northbound.Service
	modelRegistry modelregistry.ModelRegistry
	connections   e2conn.ConnManager
	oidRegistry   oid.Registry
	topo          rnib.Store
}

// Register registers the Service with the gRPC server.
func (s ControlService) Register(r *grpc.Server) {
	server := &ControlServer{
		modelRegistry: s.modelRegistry,
		connections:   s.connections,
		oidRegistry:   s.oidRegistry,
		topo:          s.topo}
	e2api.RegisterControlServiceServer(r, server)
}

// ControlServer implements the gRPC service for control
type ControlServer struct {
	modelRegistry modelregistry.ModelRegistry
	connections   e2conn.ConnManager
	oidRegistry   oid.Registry
	topo          rnib.Store
	requestID     int32
	requestMu     sync.Mutex
}

func (s *ControlServer) Control(ctx context.Context, request *e2api.ControlRequest) (*e2api.ControlResponse, error) {
	log.Infof("Received E2 Control Request %v", request)

	log.Debugf("Fetching mastership state for E2Node '%s'", request.Headers.E2NodeID)
	e2NodeEntity, err := s.topo.Get(ctx, topoapi.ID(request.Headers.E2NodeID))
	if err != nil {
		log.Warnf("Fetching mastership state for E2Node '%s' failed: %v", request.Headers.E2NodeID, err)
		return nil, errors.Status(errors.NewUnavailable(err.Error())).Err()
	}

	mastership := &topoapi.MastershipState{}
	_ = e2NodeEntity.GetAspect(mastership)
	if mastership.Term == 0 {
		err := errors.NewUnavailable("not the master for %s", request.Headers.E2NodeID)
		log.Warnf("Fetching mastership state for E2Node '%s' failed: %v", request.Headers.E2NodeID, err)
		return nil, errors.Status(err).Err()
	}

	e2NodeRelation, err := s.topo.Get(ctx, topoapi.ID(mastership.NodeId))
	if err != nil {
		log.Warnf("Fetching mastership state for E2Node '%s' failed: %v", request.Headers.E2NodeID, err)
		return nil, errors.Status(errors.NewUnavailable(err.Error())).Err()
	}

	conn, err := s.connections.Get(ctx, e2conn.ID(e2NodeRelation.ID))
	if err != nil {
		log.Warnf("Fetching mastership state for E2Node '%s' failed: %v", request.Headers.E2NodeID, err)
		return nil, errors.Status(errors.NewUnavailable(err.Error())).Err()
	}

	response := &e2api.ControlResponse{}
	serviceModelOID, err := oid.ModelIDToOid(s.oidRegistry,
		string(request.Headers.ServiceModel.Name),
		string(request.Headers.ServiceModel.Version))
	if err != nil {
		log.Warn(err)
		return nil, errors.Status(err).Err()
	}
	serviceModelPlugin, err := s.modelRegistry.GetPlugin(serviceModelOID)
	if err != nil {
		log.Warn(err)
		return nil, errors.Status(err).Err()
	}

	controlHeaderBytes := request.Message.Header
	controlMessageBytes := request.Message.Payload
	if request.Headers.Encoding == e2api.Encoding_PROTO {
		controlHeaderBytes, err = serviceModelPlugin.ControlHeaderProtoToASN1(controlHeaderBytes)
		if err != nil {
			log.Warnf("Error transforming Control Header Proto bytes to ASN: %s", err.Error())
			return nil, errors.Status(errors.NewInvalid(err.Error())).Err()
		}
		controlMessageBytes, err = serviceModelPlugin.ControlMessageProtoToASN1(controlMessageBytes)
		if err != nil {
			log.Warnf("Error transforming Control Message Proto bytes to ASN: %s", err.Error())
			return nil, errors.Status(errors.NewInvalid(err.Error())).Err()
		}
	}

	s.requestMu.Lock()
	s.requestID++
	requestID := s.requestID
	s.requestMu.Unlock()

	switch conn := conn.(type) {
	case *e2ap101server.E2Conn:
		ricRequest := e2ap101types.RicRequest{
			RequestorID: e2ap101types.RicRequestorID(requestID),
			InstanceID:  config.InstanceID,
		}
		ranFuncID, ok := conn.GetRANFunction(serviceModelOID)
		if !ok {
			log.Warn("RAN function not found for SM %s", serviceModelOID)
		}

		rcar := e2ap101ies.RiccontrolAckRequest_RICCONTROL_ACK_REQUEST_ACK
		controlRequest, err := e2ap101pdubuilder.NewControlRequest(ricRequest, ranFuncID.ID, nil, controlHeaderBytes, controlMessageBytes, &rcar)

		if err != nil {
			log.Warn(err)
			return nil, errors.Status(err).Err()
		}

		ack, failure, err := conn.RICControl(ctx, controlRequest)
		if err != nil {
			log.Warn(err)
			return nil, errors.Status(err).Err()
		}

		if ack != nil {
			outcomeProtoBytes := ack.ProtocolIes.E2ApProtocolIes32.Value.Value
			if request.Headers.Encoding == e2api.Encoding_PROTO {
				outcomeProtoBytes, err = serviceModelPlugin.ControlOutcomeASN1toProto(outcomeProtoBytes)
				if err != nil {
					log.Warnf("Error transforming Control Outcome ASN1 to Proto bytes: %s", err.Error())
					return nil, errors.Status(errors.NewInvalid(err.Error())).Err()
				}
			}
			response = &e2api.ControlResponse{
				Headers: e2api.ResponseHeaders{
					Encoding: e2api.Encoding_PROTO,
				},
				Outcome: e2api.ControlOutcome{
					Payload: outcomeProtoBytes,
				},
			}
		} else if failure != nil {
			st := status.New(codes.Aborted, "an E2AP failure occurred")
			st, err := st.WithDetails(getControlE2ap101Error(failure))
			if err != nil {
				return nil, err
			}
			return nil, st.Err()
		}
		return response, nil
	default:
		return nil, errors.Status(errors.NewNotSupported("not supported")).Err()
	}
}

func getControlE2ap101Error(failure *e2ap101pducontents.RiccontrolFailure) *e2api.Error {
	switch c := failure.ProtocolIes.E2ApProtocolIes1.Value.Cause.(type) {
	case *e2ap101ies.Cause_RicRequest:
		var errType e2api.Error_Cause_Ric_Type
		switch c.RicRequest {
		case e2ap101ies.CauseRic_CAUSE_RIC_RAN_FUNCTION_ID_INVALID:
			errType = e2api.Error_Cause_Ric_RAN_FUNCTION_ID_INVALID
		case e2ap101ies.CauseRic_CAUSE_RIC_ACTION_NOT_SUPPORTED:
			errType = e2api.Error_Cause_Ric_ACTION_NOT_SUPPORTED
		case e2ap101ies.CauseRic_CAUSE_RIC_EXCESSIVE_ACTIONS:
			errType = e2api.Error_Cause_Ric_EXCESSIVE_ACTIONS
		case e2ap101ies.CauseRic_CAUSE_RIC_DUPLICATE_ACTION:
			errType = e2api.Error_Cause_Ric_DUPLICATE_ACTION
		case e2ap101ies.CauseRic_CAUSE_RIC_DUPLICATE_EVENT:
			errType = e2api.Error_Cause_Ric_DUPLICATE_EVENT
		case e2ap101ies.CauseRic_CAUSE_RIC_FUNCTION_RESOURCE_LIMIT:
			errType = e2api.Error_Cause_Ric_FUNCTION_RESOURCE_LIMIT
		case e2ap101ies.CauseRic_CAUSE_RIC_REQUEST_ID_UNKNOWN:
			errType = e2api.Error_Cause_Ric_REQUEST_ID_UNKNOWN
		case e2ap101ies.CauseRic_CAUSE_RIC_INCONSISTENT_ACTION_SUBSEQUENT_ACTION_SEQUENCE:
			errType = e2api.Error_Cause_Ric_INCONSISTENT_ACTION_SUBSEQUENT_ACTION_SEQUENCE
		case e2ap101ies.CauseRic_CAUSE_RIC_CONTROL_MESSAGE_INVALID:
			errType = e2api.Error_Cause_Ric_CONTROL_MESSAGE_INVALID
		case e2ap101ies.CauseRic_CAUSE_RIC_CALL_PROCESS_ID_INVALID:
			errType = e2api.Error_Cause_Ric_CALL_PROCESS_ID_INVALID
		case e2ap101ies.CauseRic_CAUSE_RIC_UNSPECIFIED:
			errType = e2api.Error_Cause_Ric_UNSPECIFIED
		}
		return &e2api.Error{
			Cause: &e2api.Error_Cause{
				Cause: &e2api.Error_Cause_Ric_{
					Ric: &e2api.Error_Cause_Ric{
						Type: errType,
					},
				},
			},
		}
	case *e2ap101ies.Cause_RicService:
		var errType e2api.Error_Cause_RicService_Type
		switch c.RicService {
		case e2ap101ies.CauseRicservice_CAUSE_RICSERVICE_FUNCTION_NOT_REQUIRED:
			errType = e2api.Error_Cause_RicService_FUNCTION_NOT_REQUIRED
		case e2ap101ies.CauseRicservice_CAUSE_RICSERVICE_EXCESSIVE_FUNCTIONS:
			errType = e2api.Error_Cause_RicService_EXCESSIVE_FUNCTIONS
		case e2ap101ies.CauseRicservice_CAUSE_RICSERVICE_RIC_RESOURCE_LIMIT:
			errType = e2api.Error_Cause_RicService_RIC_RESOURCE_LIMIT
		}
		return &e2api.Error{
			Cause: &e2api.Error_Cause{
				Cause: &e2api.Error_Cause_RicService_{
					RicService: &e2api.Error_Cause_RicService{
						Type: errType,
					},
				},
			},
		}
	case *e2ap101ies.Cause_Protocol:
		var errType e2api.Error_Cause_Protocol_Type
		switch c.Protocol {
		case e2ap101ies.CauseProtocol_CAUSE_PROTOCOL_TRANSFER_SYNTAX_ERROR:
			errType = e2api.Error_Cause_Protocol_TRANSFER_SYNTAX_ERROR
		case e2ap101ies.CauseProtocol_CAUSE_PROTOCOL_ABSTRACT_SYNTAX_ERROR_REJECT:
			errType = e2api.Error_Cause_Protocol_ABSTRACT_SYNTAX_ERROR_REJECT
		case e2ap101ies.CauseProtocol_CAUSE_PROTOCOL_ABSTRACT_SYNTAX_ERROR_IGNORE_AND_NOTIFY:
			errType = e2api.Error_Cause_Protocol_ABSTRACT_SYNTAX_ERROR_IGNORE_AND_NOTIFY
		case e2ap101ies.CauseProtocol_CAUSE_PROTOCOL_MESSAGE_NOT_COMPATIBLE_WITH_RECEIVER_STATE:
			errType = e2api.Error_Cause_Protocol_MESSAGE_NOT_COMPATIBLE_WITH_RECEIVER_STATE
		case e2ap101ies.CauseProtocol_CAUSE_PROTOCOL_SEMANTIC_ERROR:
			errType = e2api.Error_Cause_Protocol_SEMANTIC_ERROR
		case e2ap101ies.CauseProtocol_CAUSE_PROTOCOL_ABSTRACT_SYNTAX_ERROR_FALSELY_CONSTRUCTED_MESSAGE:
			errType = e2api.Error_Cause_Protocol_ABSTRACT_SYNTAX_ERROR_FALSELY_CONSTRUCTED_MESSAGE
		case e2ap101ies.CauseProtocol_CAUSE_PROTOCOL_UNSPECIFIED:
			errType = e2api.Error_Cause_Protocol_UNSPECIFIED
		}
		return &e2api.Error{
			Cause: &e2api.Error_Cause{
				Cause: &e2api.Error_Cause_Protocol_{
					Protocol: &e2api.Error_Cause_Protocol{
						Type: errType,
					},
				},
			},
		}
	case *e2ap101ies.Cause_Transport:
		var errType e2api.Error_Cause_Transport_Type
		switch c.Transport {
		case e2ap101ies.CauseTransport_CAUSE_TRANSPORT_UNSPECIFIED:
			errType = e2api.Error_Cause_Transport_UNSPECIFIED
		case e2ap101ies.CauseTransport_CAUSE_TRANSPORT_TRANSPORT_RESOURCE_UNAVAILABLE:
			errType = e2api.Error_Cause_Transport_TRANSPORT_RESOURCE_UNAVAILABLE
		}
		return &e2api.Error{
			Cause: &e2api.Error_Cause{
				Cause: &e2api.Error_Cause_Transport_{
					Transport: &e2api.Error_Cause_Transport{
						Type: errType,
					},
				},
			},
		}
	case *e2ap101ies.Cause_Misc:
		var errType e2api.Error_Cause_Misc_Type
		switch c.Misc {
		case e2ap101ies.CauseMisc_CAUSE_MISC_CONTROL_PROCESSING_OVERLOAD:
			errType = e2api.Error_Cause_Misc_CONTROL_PROCESSING_OVERLOAD
		case e2ap101ies.CauseMisc_CAUSE_MISC_HARDWARE_FAILURE:
			errType = e2api.Error_Cause_Misc_HARDWARE_FAILURE
		case e2ap101ies.CauseMisc_CAUSE_MISC_OM_INTERVENTION:
			errType = e2api.Error_Cause_Misc_OM_INTERVENTION
		case e2ap101ies.CauseMisc_CAUSE_MISC_UNSPECIFIED:
			errType = e2api.Error_Cause_Misc_UNSPECIFIED
		}
		return &e2api.Error{
			Cause: &e2api.Error_Cause{
				Cause: &e2api.Error_Cause_Misc_{
					Misc: &e2api.Error_Cause_Misc{
						Type: errType,
					},
				},
			},
		}
	}
	return nil
}
