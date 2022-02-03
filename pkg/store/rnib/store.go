// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package rnib

import (
	"context"
	"io"

	"github.com/onosproject/onos-lib-go/pkg/grpc/retry"
	"google.golang.org/grpc/codes"

	"github.com/onosproject/onos-lib-go/pkg/errors"

	topoapi "github.com/onosproject/onos-api/go/onos/topo"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"google.golang.org/grpc"
)

var log = logging.GetLogger("store", "rnib")

// Store topo store client interface
type Store interface {
	// Create creates an R-NIB object
	Create(ctx context.Context, object *topoapi.Object) error

	// Update updates an existing R-NIB object
	Update(ctx context.Context, object *topoapi.Object) error

	// Get gets an R-NIB object
	Get(ctx context.Context, id topoapi.ID) (*topoapi.Object, error)

	// List lists R-NIB objects
	List(ctx context.Context, filters *topoapi.Filters) ([]topoapi.Object, error)

	// Delete deletes an R-NIB object using the given ID
	Delete(ctx context.Context, object *topoapi.Object) error

	// Watch watches topology events
	Watch(ctx context.Context, ch chan<- topoapi.Event, filters *topoapi.Filters) error
}

// NewStore creates a new R-NIB store
func NewStore(topoEndpoint string, opts ...grpc.DialOption) (Store, error) {
	if len(opts) == 0 {
		return nil, errors.New(errors.Invalid, "no opts given when creating R-NIB store")
	}
	opts = append(opts,
		grpc.WithStreamInterceptor(retry.RetryingStreamClientInterceptor(retry.WithRetryOn(codes.Unavailable, codes.Unknown))),
		grpc.WithUnaryInterceptor(retry.RetryingUnaryClientInterceptor(retry.WithRetryOn(codes.Unavailable, codes.Unknown))))
	conn, err := grpc.DialContext(context.Background(), topoEndpoint, opts...)
	if err != nil {
		log.Warn(err)
		return nil, err
	}
	client := topoapi.CreateTopoClient(conn)
	return &rnibStore{
		client: client,
	}, nil
}

type rnibStore struct {
	client topoapi.TopoClient
}

// Create creates an R-NIB object in topo store
func (s *rnibStore) Create(ctx context.Context, object *topoapi.Object) error {
	log.Debugf("Creating R-NIB object: %v", object)
	_, err := s.client.Create(ctx, &topoapi.CreateRequest{
		Object: object,
	})
	if err != nil {
		log.Warn(err)
		return errors.FromGRPC(err)
	}
	return nil
}

// Update updates the given R-NIB object in topo store
func (s *rnibStore) Update(ctx context.Context, object *topoapi.Object) error {
	log.Debugf("Updating R-NIB object: %v", object)
	response, err := s.client.Update(ctx, &topoapi.UpdateRequest{
		Object: object,
	})
	if err != nil {
		return errors.FromGRPC(err)
	}
	object = response.Object
	log.Debug("Updated R-NIB object is:", object)
	return nil
}

// Get gets an R-NIB object based on a given ID
func (s *rnibStore) Get(ctx context.Context, id topoapi.ID) (*topoapi.Object, error) {
	log.Debugf("Getting R-NIB object with ID: %v", id)
	getResponse, err := s.client.Get(ctx, &topoapi.GetRequest{
		ID: id,
	})
	if err != nil {
		return nil, errors.FromGRPC(err)
	}
	return getResponse.Object, nil
}

// List lists all of the R-NIB objects
func (s *rnibStore) List(ctx context.Context, filters *topoapi.Filters) ([]topoapi.Object, error) {
	log.Debugf("Listing R-NIB objects")
	listResponse, err := s.client.List(ctx, &topoapi.ListRequest{
		Filters: filters,
	})
	if err != nil {
		return nil, errors.FromGRPC(err)
	}

	return listResponse.Objects, nil
}

// Delete deletes an R-NIB object using the given ID
func (s *rnibStore) Delete(ctx context.Context, object *topoapi.Object) error {
	_, err := s.client.Delete(ctx, &topoapi.DeleteRequest{
		ID:       object.GetID(),
		Revision: object.GetRevision(),
	})
	if err != nil {
		return errors.FromGRPC(err)
	}
	return nil
}

// Watch watches topology events
func (s *rnibStore) Watch(ctx context.Context, ch chan<- topoapi.Event, filters *topoapi.Filters) error {
	stream, err := s.client.Watch(ctx, &topoapi.WatchRequest{
		Noreplay: false,
		Filters:  filters,
	})
	if err != nil {
		return errors.FromGRPC(err)
	}
	go func() {
		defer close(ch)
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Warn(err)
				break
			}
			ch <- resp.Event
		}
	}()
	return nil
}

var _ Store = &rnibStore{}
