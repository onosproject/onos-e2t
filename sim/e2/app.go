// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package e2

import (
	"context"
	"github.com/atomix/atomix-go-framework/pkg/atomix/logging"
	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	"github.com/onosproject/onos-e2t/test/utils"
	e2 "github.com/onosproject/onos-ric-sdk-go/pkg/e2/v1beta1"
	"google.golang.org/grpc"
	"net"
)

var log = logging.GetLogger("sim", "app")

func NewApp(client e2.Client) *App {
	return &App{
		client: client,
	}
}

type App struct {
	client e2.Client
	server *grpc.Server
}

func (a *App) startSubscription(ctx context.Context, id string, nodeID string, cellObjectID string, reportPeriod uint32, granularity uint32) error {
	eventTriggerBytes, err := utils.CreateKpmV2EventTrigger(reportPeriod)
	if err != nil {
		return err
	}

	// Use one of the cell object IDs for action definition
	actionDefinitionBytes, err := utils.CreateKpmV2ActionDefinition(cellObjectID, granularity)
	if err != nil {
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

	err = a.client.Node(e2.NodeID(nodeID)).Subscribe(context.Background(), id, spec, ch)
	if err != nil {
		return err
	}

	go func() {
		for ind := range ch {
			log.Infof("Received indication for subscription %s: %v", id, ind)
		}
	}()
	return nil
}

func (a *App) stopSubscription(ctx context.Context, id string, nodeID string) error {
	return a.client.Node(e2.NodeID(nodeID)).Unsubscribe(ctx, id)
}

func (a *App) Start() error {
	a.server = grpc.NewServer()
	lis, err := net.Listen("tcp", ":5000")
	if err != nil {
		return err
	}
	RegisterSimServiceServer(a.server, &AppServer{a})
	err = a.server.Serve(lis)
	if err != nil {
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