// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package utils

import (
	"context"
	"fmt"
	"time"

	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	e2smkpmv2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/pdubuilder"
	kpmv2types "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	"github.com/onosproject/onos-ric-sdk-go/pkg/e2/creds"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/protobuf/proto"
)

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
	measInfoConAvgItem, err := e2smkpmv2.CreateMeasurementInfoItem(rrcConAvgName)
	if err != nil {
		return nil, err
	}
	measInfoConMaxItem, err := e2smkpmv2.CreateMeasurementInfoItem(rrcConMaxName)
	if err != nil {
		return nil, err
	}
	measInfoConnEstabAttItem, err := e2smkpmv2.CreateMeasurementInfoItem(rrcConnEstabAtt)
	if err != nil {
		return nil, err
	}

	measInfoList := &kpmv2types.MeasurementInfoList{
		Value: make([]*kpmv2types.MeasurementInfoItem, 0),
	}
	measInfoList.Value = append(measInfoList.Value, measInfoConAvgItem)
	measInfoList.Value = append(measInfoList.Value, measInfoConMaxItem)
	measInfoList.Value = append(measInfoList.Value, measInfoConnEstabAttItem)

	actionDefinition, err := e2smkpmv2.CreateActionDefinitionFormat1(cellObjectID, measInfoList, int64(granularity), 1234)
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
	e2SmKpmEventTriggerDefinition, err := e2smkpmv2.CreateE2SmKpmEventTriggerDefinition(int64(rtPeriod))
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
