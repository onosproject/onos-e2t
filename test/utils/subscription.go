// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package utils

import (
	"context"
	"fmt"
	"testing"
	"time"

	topoapi "github.com/onosproject/onos-api/go/onos/topo"
	"github.com/stretchr/testify/assert"

	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/pdubuilder"
	e2smkpmv2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/pdubuilder"
	kpmv2types "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	rcpdubuilder "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/pdubuilder"
	"github.com/onosproject/onos-ric-sdk-go/pkg/e2/creds"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/protobuf/proto"
)

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
	rrcConnEstabAtt, err := e2smkpmv2.CreateMeasurementTypeMeasName("RRC.ConnEstabAtt.Sum")
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

// CreateKpmV2EventTrigger creates a kpm v2 service model event trigger
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

// CreateKpmV1EventTrigger creates a kpm v1 service model event trigger
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

// Subscription subscription request for subscription SDK api
type Subscription struct {
	NodeID              string
	ServiceModelName    e2api.ServiceModelName
	ServiceModelVersion e2api.ServiceModelVersion
	Actions             []e2api.Action
	EventTrigger        []byte
}

// Create creates a subscription request using SDK
func (subRequest *Subscription) Create() (e2api.SubscriptionSpec, error) {
	spec := e2api.SubscriptionSpec{
		EventTrigger: e2api.EventTrigger{
			Payload: subRequest.EventTrigger,
		},
		Actions: subRequest.Actions,
	}
	return spec, nil
}

func (subRequest *Subscription) CreateWithActionDefinition() (e2api.SubscriptionSpec, error) {
	spec := e2api.SubscriptionSpec{
		EventTrigger: e2api.EventTrigger{
			Payload: subRequest.EventTrigger,
		},

		Actions: subRequest.Actions,
	}
	return spec, nil
}

// ConnectE2tServiceHost connects to subscription service via service name
func ConnectE2tServiceHost() (*grpc.ClientConn, error) {
	tlsConfig, err := creds.GetClientCredentials()
	if err != nil {
		return nil, err
	}
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)),
	}

	return grpc.DialContext(context.Background(), E2tServiceAddress, opts...)
}

// ConnectE2t connects to subscription service via IP/port
func ConnectE2t(IP string, port uint32) (*grpc.ClientConn, error) {
	tlsConfig, err := creds.GetClientCredentials()
	if err != nil {
		return nil, err
	}
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)),
	}

	addr := fmt.Sprintf("%s:%d", IP, port)
	return grpc.DialContext(context.Background(), addr, opts...)
}

// CreateKpmV2Sub :
func CreateKpmV2Sub(t *testing.T, nodeID topoapi.ID) e2api.SubscriptionSpec {

	topoSdkClient, err := NewTopoClient()
	assert.NoError(t, err)

	cells, err := topoSdkClient.GetCells(context.Background(), nodeID)
	assert.NoError(t, err)
	assert.Greater(t, len(cells), 0)

	return CreateKpmV2SubWithCell(t, nodeID, cells[0].CellObjectID)
}

// CreateKpmV2SubWithCell :
func CreateKpmV2SubWithCell(t *testing.T, nodeID topoapi.ID, cellObjectID string) e2api.SubscriptionSpec {
	reportPeriod := uint32(5000)
	granularity := uint32(500)

	// Kpm v2 interval is defined in ms
	eventTriggerBytes, err := CreateKpmV2EventTrigger(reportPeriod)
	assert.NoError(t, err)

	// Use one of the cell object IDs for action definition
	actionDefinitionBytes, err := CreateKpmV2ActionDefinition(cellObjectID, granularity)
	assert.NoError(t, err)

	var actions []e2api.Action
	action := e2api.Action{
		ID:   100,
		Type: e2api.ActionType_ACTION_TYPE_REPORT,
		SubsequentAction: &e2api.SubsequentAction{
			Type:       e2api.SubsequentActionType_SUBSEQUENT_ACTION_TYPE_CONTINUE,
			TimeToWait: e2api.TimeToWait_TIME_TO_WAIT_ZERO,
		},
		Payload: actionDefinitionBytes,
	}

	actions = append(actions, action)

	subRequest := Subscription{
		NodeID:              string(nodeID),
		EventTrigger:        eventTriggerBytes,
		ServiceModelName:    KpmServiceModelName,
		ServiceModelVersion: Version2,
		Actions:             actions,
	}

	subSpec, err := subRequest.CreateWithActionDefinition()
	assert.NoError(t, err)
	return subSpec
}

// ReadToEndOfChannel reads messages from a channel until an error occurs, clearing the
// channel of messages
func ReadToEndOfChannel(ch chan e2api.Indication) bool {
	for {
		select {
		case _, ok := <-ch:
			if !ok {
				return true
			}
		case <-time.After(2 * time.Minute):
			return false
		}
	}
}
