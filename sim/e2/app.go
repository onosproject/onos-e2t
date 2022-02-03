// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2

import (
	"context"
	"fmt"
	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	"github.com/onosproject/onos-e2t/test/utils"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	e2 "github.com/onosproject/onos-ric-sdk-go/pkg/e2/v1beta1"
	"google.golang.org/grpc"
	"net"
)

func NewApp(appID e2.AppID, instanceID e2.InstanceID) *App {
	client := e2.NewClient(
		e2.WithAppID(appID),
		e2.WithInstanceID(instanceID),
		e2.WithServiceModel(utils.KpmServiceModelName, utils.Version2),
		e2.WithProtoEncoding())
	return &App{
		client: client,
		log:    log.GetLogger(string(appID), string(instanceID)),
	}
}

type App struct {
	client e2.Client
	server *grpc.Server
	log    logging.Logger
}

func (a *App) startSubscription(ctx context.Context, id string, nodeID string, cellObjectID string, reportPeriod uint32, granularity uint32) error {
	a.log.Infof("Starting %s subscription %s", nodeID, id)
	eventTriggerBytes, err := utils.CreateKpmV2EventTrigger(reportPeriod)
	if err != nil {
		a.log.Error(err)
		return err
	}

	// Use one of the cell object IDs for action definition
	actionDefinitionBytes, err := utils.CreateKpmV2ActionDefinition(cellObjectID, granularity)
	if err != nil {
		a.log.Error(err)
		return err
	}

	ch := make(chan e2api.Indication)
	spec := e2api.SubscriptionSpec{
		EventTrigger: e2api.EventTrigger{
			Payload: eventTriggerBytes,
		},
		Actions: []e2api.Action{
			{
				ID:   100,
				Type: e2api.ActionType_ACTION_TYPE_REPORT,
				SubsequentAction: &e2api.SubsequentAction{
					Type:       e2api.SubsequentActionType_SUBSEQUENT_ACTION_TYPE_CONTINUE,
					TimeToWait: e2api.TimeToWait_TIME_TO_WAIT_ZERO,
				},
				Payload: actionDefinitionBytes,
			},
		},
	}

	_, err = a.client.Node(e2.NodeID(nodeID)).Subscribe(context.Background(), id, spec, ch)
	if err != nil {
		a.log.Error(err)
		return err
	}

	go func() {
		for range ch {
			a.log.Infof("Received indication for subscription %s", id)
		}
		a.log.Infof("Indication channel closed for subscription %s", id)
	}()
	return nil
}

func (a *App) stopSubscription(ctx context.Context, id string, nodeID string) error {
	a.log.Infof("Stopping %s subscription %s", nodeID, id)
	err := a.client.Node(e2.NodeID(nodeID)).Unsubscribe(ctx, id)
	if err != nil {
		a.log.Error(err)
		return err
	}
	return nil
}

func (a *App) Start() error {
	a.server = grpc.NewServer()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", controlPort))
	if err != nil {
		a.log.Error(err)
		return err
	}
	RegisterSimServiceServer(a.server, &AppServer{a})
	err = a.server.Serve(lis)
	if err != nil {
		a.log.Error(err)
		return err
	}
	return nil
}

func (a *App) Stop() error {
	a.server.Stop()
	return nil
}

type AppServer struct {
	app *App
}

func (a *AppServer) StartSubscription(ctx context.Context, request *StartSubscriptionRequest) (*StartSubscriptionResponse, error) {
	err := a.app.startSubscription(
		ctx,
		request.SubscriptionId,
		request.NodeId,
		request.CellObjectId,
		request.ReportPeriod,
		request.Granularity)
	if err != nil {
		return nil, err
	}
	return &StartSubscriptionResponse{}, nil
}

func (a *AppServer) StopSubscription(ctx context.Context, request *StopSubscriptionRequest) (*StopSubscriptionResponse, error) {
	err := a.app.stopSubscription(ctx, request.SubscriptionId, request.NodeId)
	if err != nil {
		return nil, err
	}
	return &StopSubscriptionResponse{}, nil
}
