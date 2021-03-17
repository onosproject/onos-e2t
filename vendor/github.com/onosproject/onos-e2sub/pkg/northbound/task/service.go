// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package task

import (
	"context"

	taskapi "github.com/onosproject/onos-api/go/onos/e2sub/task"
	store "github.com/onosproject/onos-e2sub/pkg/store/task"
	"github.com/onosproject/onos-lib-go/pkg/errors"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/onosproject/onos-lib-go/pkg/northbound"
	"google.golang.org/grpc"
)

var log = logging.GetLogger("northbound", "task")

// NewService creates a new subscription service
func NewService(store store.Store) northbound.Service {
	return &Service{
		store: store,
	}
}

// Service is a Service implementation for subscription service.
type Service struct {
	store store.Store
}

// Register registers the Service with the gRPC server.
func (s *Service) Register(r *grpc.Server) {
	server := &Server{
		store: s.store,
	}
	taskapi.RegisterE2SubscriptionTaskServiceServer(r, server)
}

var _ northbound.Service = &Service{}

// Server implements the gRPC service for managing of subscriptions
type Server struct {
	store store.Store
}

func (s *Server) GetSubscriptionTask(ctx context.Context, req *taskapi.GetSubscriptionTaskRequest) (*taskapi.GetSubscriptionTaskResponse, error) {
	log.Infof("Received GetSubscriptionTaskRequest %+v", req)
	task, err := s.store.Get(ctx, req.ID)
	if err != nil {
		log.Warnf("GetSubscriptionTaskRequest %+v failed: %v", req, err)
		return nil, errors.Status(err).Err()
	}
	res := &taskapi.GetSubscriptionTaskResponse{
		Task: task,
	}
	log.Infof("Sending GetSubscriptionTaskResponse %+v", res)
	return res, nil
}

func (s *Server) ListSubscriptionTasks(ctx context.Context, req *taskapi.ListSubscriptionTasksRequest) (*taskapi.ListSubscriptionTasksResponse, error) {
	log.Infof("Received ListSubscriptionTasksRequest %+v", req)
	tasks, err := s.store.List(ctx)
	if err != nil {
		log.Warnf("ListSubscriptionTasksRequest %+v failed: %v", req, err)
		return nil, errors.Status(err).Err()
	}

	filtered := make([]taskapi.SubscriptionTask, 0, len(tasks))
	for _, task := range tasks {
		if req.SubscriptionID != "" && task.SubscriptionID != req.SubscriptionID {
			continue
		}
		if req.EndpointID != "" && task.EndpointID != req.EndpointID {
			continue
		}
		filtered = append(filtered, task)
	}

	res := &taskapi.ListSubscriptionTasksResponse{
		Tasks: filtered,
	}
	log.Infof("Sending ListSubscriptionTasksResponse %+v", res)
	return res, nil
}

func (s *Server) WatchSubscriptionTasks(req *taskapi.WatchSubscriptionTasksRequest, server taskapi.E2SubscriptionTaskService_WatchSubscriptionTasksServer) error {
	log.Infof("Received WatchSubscriptionTasksRequest %+v", req)
	var watchOpts []store.WatchOption
	if !req.Noreplay {
		watchOpts = append(watchOpts, store.WithReplay())
	}

	ch := make(chan taskapi.Event)
	if err := s.store.Watch(server.Context(), ch, watchOpts...); err != nil {
		log.Warnf("WatchSubscriptionTasksRequest %+v failed: %v", req, err)
		return errors.Status(err).Err()
	}

	for event := range ch {
		if req.SubscriptionID != "" && event.Task.SubscriptionID != req.SubscriptionID {
			continue
		}
		if req.EndpointID != "" && event.Task.EndpointID != req.EndpointID {
			continue
		}

		res := &taskapi.WatchSubscriptionTasksResponse{
			Event: event,
		}

		log.Infof("Sending WatchSubscriptionTasksResponse %+v", res)
		if err := server.Send(res); err != nil {
			log.Warnf("WatchSubscriptionTasksResponse %+v failed: %v", res, err)
			return err
		}
	}
	return nil
}

func (s *Server) UpdateSubscriptionTask(ctx context.Context, req *taskapi.UpdateSubscriptionTaskRequest) (*taskapi.UpdateSubscriptionTaskResponse, error) {
	log.Infof("Received UpdateSubscriptionTaskRequest %+v", req)
	err := s.store.Update(ctx, req.Task)
	if err != nil {
		log.Warnf("UpdateSubscriptionTaskRequest %+v failed: %v", req, err)
		return nil, errors.Status(err).Err()
	}
	res := &taskapi.UpdateSubscriptionTaskResponse{}
	log.Infof("Sending UpdateSubscriptionTaskResponse %+v", res)
	return res, nil
}
