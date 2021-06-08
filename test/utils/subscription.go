// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package utils

import (
	"context"

	subapi "github.com/onosproject/onos-api/go/onos/e2sub/subscription"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/pdubuilder"
	e2smkpmv2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/pdubuilder"
	kpmv2types "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	rcpdubuilder "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/pdubuilder"
	"github.com/onosproject/onos-ric-sdk-go/pkg/e2/creds"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/protobuf/proto"
)

// Subscription subscription request for subscription SDK api
type Subscription struct {
	NodeID              string
	ServiceModelName    subapi.ServiceModelName
	ServiceModelVersion subapi.ServiceModelVersion
	Actions             []subapi.Action
	EncodingType        subapi.Encoding
	EventTrigger        []byte
}

// CreateRcEventTrigger creates a rc service model event trigger
func CreateRcEventTrigger() ([]byte, error) {
	e2SmKpmEventTriggerDefinition, err := rcpdubuilder.CreateE2SmRcPreEventTriggerDefinitionUponChange()
	if err != nil {
		return []byte{}, err
	}
	err = e2SmKpmEventTriggerDefinition.Validate()
	if err != nil {
		return []byte{}, err
	}
	protoBytes, err := proto.Marshal(e2SmKpmEventTriggerDefinition)
	if err != nil {
		return []byte{}, err
	}
	return protoBytes, nil
}

func CreateKpmV2ActionDefinition(cellObjectID string, granularity uint32) ([]byte, error) {
	rrcConAvgName, err := e2smkpmv2.CreateMeasurementTypeMeasName("RRC.Conn.Avg")
	if err != nil {
		return nil, err
	}
	rrcConMaxName, err := e2smkpmv2.CreateMeasurementTypeMeasName("RRC.Conn.Max")
	if err != nil {
		return nil, err
	}
	rrcConnEstabAtt, err := e2smkpmv2.CreateMeasurementTypeMeasName("RRC.ConnEstabAtt.Tot")
	if err != nil {
		return nil, err
	}
	measInfoConAvgItem, err := e2smkpmv2.CreateMeasurementInfoItem(rrcConAvgName, nil)
	if err != nil {
		return nil, err
	}
	measInfoConMaxItem, err := e2smkpmv2.CreateMeasurementInfoItem(rrcConMaxName, nil)
	if err != nil {
		return nil, err
	}
	measInfoConnEstabAttItem, err := e2smkpmv2.CreateMeasurementInfoItem(rrcConnEstabAtt, nil)
	if err != nil {
		return nil, err
	}

	measInfoList := &kpmv2types.MeasurementInfoList{
		Value: make([]*kpmv2types.MeasurementInfoItem, 0),
	}
	measInfoList.Value = append(measInfoList.Value, measInfoConAvgItem)
	measInfoList.Value = append(measInfoList.Value, measInfoConMaxItem)
	measInfoList.Value = append(measInfoList.Value, measInfoConnEstabAttItem)

	actionDefinition, err := e2smkpmv2.CreateActionDefinitionFormat1(cellObjectID, measInfoList, granularity, 1234)
	if err != nil {
		return nil, err
	}

	e2SmKpmActionDefinition, err := e2smkpmv2.CreateE2SmKpmActionDefinitionFormat1(1, actionDefinition)
	if err != nil {
		return nil, err
	}

	protoBytes, err := proto.Marshal(e2SmKpmActionDefinition)
	if err != nil {
		return nil, err
	}
	return protoBytes, nil
}

// CreateKpmEventTrigger creates a kpm service model event trigger
func CreateKpmV2EventTrigger(rtPeriod uint32) ([]byte, error) {
	e2SmKpmEventTriggerDefinition, err := e2smkpmv2.CreateE2SmKpmEventTriggerDefinition(rtPeriod)
	if err != nil {
		return []byte{}, err
	}
	err = e2SmKpmEventTriggerDefinition.Validate()
	if err != nil {
		return []byte{}, err
	}
	protoBytes, err := proto.Marshal(e2SmKpmEventTriggerDefinition)
	if err != nil {
		return []byte{}, err
	}
	return protoBytes, nil
}

// CreateKpmEventTrigger creates a kpm service model event trigger
func CreateKpmV1EventTrigger(rtPeriod int32) ([]byte, error) {
	e2SmKpmEventTriggerDefinition, err := pdubuilder.CreateE2SmKpmEventTriggerDefinition(rtPeriod)
	if err != nil {
		return []byte{}, err
	}
	err = e2SmKpmEventTriggerDefinition.Validate()
	if err != nil {
		return []byte{}, err
	}
	protoBytes, err := proto.Marshal(e2SmKpmEventTriggerDefinition)
	if err != nil {
		return []byte{}, err
	}
	return protoBytes, nil
}

func (subRequest *Subscription) CreateWithActionDefinition() (subapi.SubscriptionDetails, error) {
	subReq := subapi.SubscriptionDetails{
		E2NodeID: subapi.E2NodeID(subRequest.NodeID),
		ServiceModel: subapi.ServiceModel{
			Name:    subRequest.ServiceModelName,
			Version: subRequest.ServiceModelVersion,
		},
		EventTrigger: subapi.EventTrigger{
			Payload: subapi.Payload{
				Encoding: subRequest.EncodingType,
				Data:     subRequest.EventTrigger,
			},
		},
		Actions: subRequest.Actions,
	}

	return subReq, nil

}

// Create creates a subscription request using SDK
func (subRequest *Subscription) Create() (subapi.SubscriptionDetails, error) {
	subReq := subapi.SubscriptionDetails{
		E2NodeID: subapi.E2NodeID(subRequest.NodeID),
		ServiceModel: subapi.ServiceModel{
			Name:    subRequest.ServiceModelName,
			Version: subRequest.ServiceModelVersion,
		},
		EventTrigger: subapi.EventTrigger{
			Payload: subapi.Payload{
				Encoding: subRequest.EncodingType,
				Data:     subRequest.EventTrigger,
			},
		},
		Actions: subRequest.Actions,
	}

	return subReq, nil
}

// ConnectSubscriptionServiceHost connects to subscription service
func ConnectSubscriptionServiceHost() (*grpc.ClientConn, error) {
	tlsConfig, err := creds.GetClientCredentials()
	if err != nil {
		return nil, err
	}
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)),
	}

	return grpc.DialContext(context.Background(), SubscriptionServiceAddress, opts...)
}
