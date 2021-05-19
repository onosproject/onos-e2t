// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package object

import (
	"context"
	"io"
	"time"

	"github.com/onosproject/onos-lib-go/pkg/errors"

	topoapi "github.com/onosproject/onos-api/go/onos/topo"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/onosproject/onos-lib-go/pkg/southbound"
	"google.golang.org/grpc"
)

var log = logging.GetLogger("store", "object")

const (
	defaultTimeout      = 15
	defaultRetryTimeout = 100
)

// Store topo store client interface
type Store interface {
	// Create creates a topo object
	Create(ctx context.Context, object *topoapi.Object) error

	// Update updates an existing topo object
	Update(ctx context.Context, object *topoapi.Object) error

	// Get gets a topo object
	Get(ctx context.Context, id topoapi.ID) (*topoapi.Object, error)

	// List lists topo objects
	List(ctx context.Context) ([]topoapi.Object, error)

	// Delete deletes a topo object using the given ID
	Delete(ctx context.Context, id topoapi.ID) error

	// Watch watches topology events
	Watch(ctx context.Context, ch chan<- topoapi.Event) error
}

// NewTopoStore returns a new topo device store
func NewTopoStore(topoEndpoint string, opts ...grpc.DialOption) (Store, error) {
	if len(opts) == 0 {
		return nil, errors.New(errors.Invalid, "no opts given when creating topo store")
	}
	opts = append(opts, grpc.WithStreamInterceptor(southbound.RetryingStreamClientInterceptor(defaultRetryTimeout*time.Millisecond)))
	conn, err := getTopoConn(topoEndpoint, opts...)
	if err != nil {
		log.Warn(err)
		return nil, err
	}
	client := topoapi.CreateTopoClient(conn)
	return &topoStore{
		client: client,
	}, nil
}

type topoStore struct {
	client topoapi.TopoClient
}

// Create creates an object in topo store
func (s *topoStore) Create(ctx context.Context, object *topoapi.Object) error {
	log.Debugf("Creating topo object: %v", object)
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout*time.Second)
	defer cancel()
	_, err := s.client.Create(ctx, &topoapi.CreateRequest{
		Object: object,
	})
	if err != nil {
		log.Warn(err)
		return err
	}
	return nil
}

// Update updates the given object in topo store
func (s *topoStore) Update(ctx context.Context, object *topoapi.Object) error {
	log.Debugf("Updating topo object: %v", object)
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout*time.Second)
	defer cancel()
	response, err := s.client.Update(ctx, &topoapi.UpdateRequest{
		Object: object,
	})
	if err != nil {
		return err
	}
	object = response.Object
	return nil
}

// Get gets an object based on a given ID
func (s *topoStore) Get(ctx context.Context, id topoapi.ID) (*topoapi.Object, error) {
	log.Debugf("Getting the topo object with ID: %v", id)
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout*time.Second)
	defer cancel()
	getResponse, err := s.client.Get(ctx, &topoapi.GetRequest{
		ID: id,
	})
	if err != nil {
		return nil, err
	}
	return getResponse.Object, nil
}

// List lists all of the topo objects
func (s *topoStore) List(ctx context.Context) ([]topoapi.Object, error) {
	log.Debugf("Listing topo objects")
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout*time.Second)
	defer cancel()
	listResponse, err := s.client.List(ctx, &topoapi.ListRequest{})
	if err != nil {
		return nil, err
	}

	return listResponse.Objects, nil
}

// Delete deletes a topo object using the given ID
func (s *topoStore) Delete(ctx context.Context, id topoapi.ID) error {
	_, err := s.client.Delete(ctx, &topoapi.DeleteRequest{
		ID: id,
	})
	if err != nil {
		return err
	}
	return nil
}

// Watch watches topology events
func (s *topoStore) Watch(ctx context.Context, ch chan<- topoapi.Event) error {
	stream, err := s.client.Watch(ctx, &topoapi.WatchRequest{
		Noreplay: false,
	})
	if err != nil {
		return err
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

// getTopoConn gets a gRPC connection to the topology service
func getTopoConn(topoEndpoint string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	return grpc.Dial(topoEndpoint, opts...)
}

var _ Store = &topoStore{}
