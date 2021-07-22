// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package v1beta1

import (
	"context"
	"sync"

	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	topoapi "github.com/onosproject/onos-api/go/onos/topo"

	"github.com/onosproject/onos-e2t/pkg/topo"

	"github.com/onosproject/onos-e2t/pkg/oid"

	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"

	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/pdubuilder"

	"github.com/onosproject/onos-e2t/pkg/config"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/types"

	"github.com/onosproject/onos-e2t/pkg/modelregistry"
	e2server "github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/server"

	"github.com/onosproject/onos-lib-go/pkg/errors"

	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/onosproject/onos-lib-go/pkg/northbound"
	"google.golang.org/grpc"
)

var log = logging.GetLogger("northbound", "e2", "v1beta1")

// NewControlService creates a new control service
func NewControlService(modelRegistry modelregistry.ModelRegistry, channels e2server.ChannelManager,
	oidRegistry oid.Registry, topoManager topo.Manager) northbound.Service {
	return &ControlService{
		modelRegistry: modelRegistry,
		channels:      channels,
		oidRegistry:   oidRegistry,
		topoManager:   topoManager,
	}
}

// ControlService is a Service implementation for control requests
type ControlService struct {
	northbound.Service
	modelRegistry modelregistry.ModelRegistry
	channels      e2server.ChannelManager
	oidRegistry   oid.Registry
	topoManager   topo.Manager
}

// Register registers the Service with the gRPC server.
func (s ControlService) Register(r *grpc.Server) {
	server := &ControlServer{
		modelRegistry: s.modelRegistry,
		channels:      s.channels,
		oidRegistry:   s.oidRegistry,
		topoManager:   s.topoManager}
	e2api.RegisterControlServiceServer(r, server)
}

// ControlServer implements the gRPC service for control
type ControlServer struct {
	modelRegistry modelregistry.ModelRegistry
	channels      e2server.ChannelManager
	oidRegistry   oid.Registry
	topoManager   topo.Manager
	requestID     int32
	requestMu     sync.Mutex
}

func (s *ControlServer) Control(ctx context.Context, request *e2api.ControlRequest) (*e2api.ControlResponse, error) {
	log.Infof("Received E2 Control Request %v", request)

	channel, err := s.channels.Get(ctx, topoapi.ID(request.Headers.E2NodeID))
	if err != nil {
		return nil, errors.Status(err).Err()
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

	s.requestMu.Lock()
	s.requestID++
	requestID := s.requestID
	s.requestMu.Unlock()

	ricRequest := types.RicRequest{
		RequestorID: types.RicRequestorID(requestID),
		InstanceID:  config.InstanceID,
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

	ranFuncID, ok := channel.GetRANFunction(serviceModelOID)
	if !ok {
		log.Warn("RAN function not found for SM %s", serviceModelOID)
	}

	rcar := e2apies.RiccontrolAckRequest_RICCONTROL_ACK_REQUEST_ACK
	controlRequest, err := pdubuilder.NewControlRequest(ricRequest, ranFuncID.ID, nil, controlHeaderBytes, controlMessageBytes, &rcar)

	if err != nil {
		log.Warn(err)
		return nil, errors.Status(err).Err()
	}

	ack, failure, err := channel.RICControl(ctx, controlRequest)
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
		st, err := st.WithDetails(getControlError(failure))
		if err != nil {
			return nil, err
		}
		return nil, st.Err()
	}
	return response, nil
}

func getControlError(failure *e2appducontents.RiccontrolFailure) *e2api.Error {
	switch c := failure.ProtocolIes.E2ApProtocolIes1.Value.Cause.(type) {
	case *e2apies.Cause_RicRequest:
		var errType e2api.Error_Cause_Ric_Type
		switch c.RicRequest {
		case e2apies.CauseRic_CAUSE_RIC_RAN_FUNCTION_ID_INVALID:
			errType = e2api.Error_Cause_Ric_RAN_FUNCTION_ID_INVALID
		case e2apies.CauseRic_CAUSE_RIC_ACTION_NOT_SUPPORTED:
			errType = e2api.Error_Cause_Ric_ACTION_NOT_SUPPORTED
		case e2apies.CauseRic_CAUSE_RIC_EXCESSIVE_ACTIONS:
			errType = e2api.Error_Cause_Ric_EXCESSIVE_ACTIONS
		case e2apies.CauseRic_CAUSE_RIC_DUPLICATE_ACTION:
			errType = e2api.Error_Cause_Ric_DUPLICATE_ACTION
		case e2apies.CauseRic_CAUSE_RIC_DUPLICATE_EVENT:
			errType = e2api.Error_Cause_Ric_DUPLICATE_EVENT
		case e2apies.CauseRic_CAUSE_RIC_FUNCTION_RESOURCE_LIMIT:
			errType = e2api.Error_Cause_Ric_FUNCTION_RESOURCE_LIMIT
		case e2apies.CauseRic_CAUSE_RIC_REQUEST_ID_UNKNOWN:
			errType = e2api.Error_Cause_Ric_REQUEST_ID_UNKNOWN
		case e2apies.CauseRic_CAUSE_RIC_INCONSISTENT_ACTION_SUBSEQUENT_ACTION_SEQUENCE:
			errType = e2api.Error_Cause_Ric_INCONSISTENT_ACTION_SUBSEQUENT_ACTION_SEQUENCE
		case e2apies.CauseRic_CAUSE_RIC_CONTROL_MESSAGE_INVALID:
			errType = e2api.Error_Cause_Ric_CONTROL_MESSAGE_INVALID
		case e2apies.CauseRic_CAUSE_RIC_CALL_PROCESS_ID_INVALID:
			errType = e2api.Error_Cause_Ric_CALL_PROCESS_ID_INVALID
		case e2apies.CauseRic_CAUSE_RIC_UNSPECIFIED:
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
	case *e2apies.Cause_RicService:
		var errType e2api.Error_Cause_RicService_Type
		switch c.RicService {
		case e2apies.CauseRicservice_CAUSE_RICSERVICE_FUNCTION_NOT_REQUIRED:
			errType = e2api.Error_Cause_RicService_FUNCTION_NOT_REQUIRED
		case e2apies.CauseRicservice_CAUSE_RICSERVICE_EXCESSIVE_FUNCTIONS:
			errType = e2api.Error_Cause_RicService_EXCESSIVE_FUNCTIONS
		case e2apies.CauseRicservice_CAUSE_RICSERVICE_RIC_RESOURCE_LIMIT:
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
	case *e2apies.Cause_Protocol:
		var errType e2api.Error_Cause_Protocol_Type
		switch c.Protocol {
		case e2apies.CauseProtocol_CAUSE_PROTOCOL_TRANSFER_SYNTAX_ERROR:
			errType = e2api.Error_Cause_Protocol_TRANSFER_SYNTAX_ERROR
		case e2apies.CauseProtocol_CAUSE_PROTOCOL_ABSTRACT_SYNTAX_ERROR_REJECT:
			errType = e2api.Error_Cause_Protocol_ABSTRACT_SYNTAX_ERROR_REJECT
		case e2apies.CauseProtocol_CAUSE_PROTOCOL_ABSTRACT_SYNTAX_ERROR_IGNORE_AND_NOTIFY:
			errType = e2api.Error_Cause_Protocol_ABSTRACT_SYNTAX_ERROR_IGNORE_AND_NOTIFY
		case e2apies.CauseProtocol_CAUSE_PROTOCOL_MESSAGE_NOT_COMPATIBLE_WITH_RECEIVER_STATE:
			errType = e2api.Error_Cause_Protocol_MESSAGE_NOT_COMPATIBLE_WITH_RECEIVER_STATE
		case e2apies.CauseProtocol_CAUSE_PROTOCOL_SEMANTIC_ERROR:
			errType = e2api.Error_Cause_Protocol_SEMANTIC_ERROR
		case e2apies.CauseProtocol_CAUSE_PROTOCOL_ABSTRACT_SYNTAX_ERROR_FALSELY_CONSTRUCTED_MESSAGE:
			errType = e2api.Error_Cause_Protocol_ABSTRACT_SYNTAX_ERROR_FALSELY_CONSTRUCTED_MESSAGE
		case e2apies.CauseProtocol_CAUSE_PROTOCOL_UNSPECIFIED:
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
	case *e2apies.Cause_Transport:
		var errType e2api.Error_Cause_Transport_Type
		switch c.Transport {
		case e2apies.CauseTransport_CAUSE_TRANSPORT_UNSPECIFIED:
			errType = e2api.Error_Cause_Transport_UNSPECIFIED
		case e2apies.CauseTransport_CAUSE_TRANSPORT_TRANSPORT_RESOURCE_UNAVAILABLE:
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
	case *e2apies.Cause_Misc:
		var errType e2api.Error_Cause_Misc_Type
		switch c.Misc {
		case e2apies.CauseMisc_CAUSE_MISC_CONTROL_PROCESSING_OVERLOAD:
			errType = e2api.Error_Cause_Misc_CONTROL_PROCESSING_OVERLOAD
		case e2apies.CauseMisc_CAUSE_MISC_HARDWARE_FAILURE:
			errType = e2api.Error_Cause_Misc_HARDWARE_FAILURE
		case e2apies.CauseMisc_CAUSE_MISC_OM_INTERVENTION:
			errType = e2api.Error_Cause_Misc_OM_INTERVENTION
		case e2apies.CauseMisc_CAUSE_MISC_UNSPECIFIED:
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
