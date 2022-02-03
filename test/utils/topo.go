// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package utils

import (
	"context"
	"errors"
	"testing"
	"time"

	toposdk "github.com/onosproject/onos-ric-sdk-go/pkg/topo"

	"github.com/stretchr/testify/assert"

	topoapi "github.com/onosproject/onos-api/go/onos/topo"
)

// TopoClient R-NIB client interface
type TopoClient interface {
	WatchE2Connections(ctx context.Context, ch chan topoapi.Event) error
	WatchE2TNodes(ctx context.Context, ch chan topoapi.Event) error
	GetCells(ctx context.Context, nodeID topoapi.ID) ([]*topoapi.E2Cell, error)
	GetE2NodeAspects(ctx context.Context, nodeID topoapi.ID) (*topoapi.E2Node, error)
	E2NodeRelationIDs(ctx context.Context) ([]topoapi.ID, error)
	GetControlRelations() ([]topoapi.Object, error)
}

// NewTopoClient creates a new topo SDK client
func NewTopoClient() (Client, error) {
	sdkClient, err := toposdk.NewClient()
	if err != nil {
		return Client{}, err
	}
	cl := Client{
		client: sdkClient,
	}
	return cl, nil
}

// Client topo SDK client
type Client struct {
	client toposdk.Client
}

// E2NodeRelationIDs lists the relation entity IDs of all of connected E2 nodes
func (c *Client) E2NodeRelationIDs(ctx context.Context) ([]topoapi.ID, error) {
	objects, err := c.client.List(ctx, toposdk.WithListFilters(getControlRelationFilter()))
	if err != nil {
		return nil, err
	}

	e2NodeIDs := make([]topoapi.ID, len(objects))
	for _, object := range objects {
		relation := object.GetRelation()
		e2NodeID := relation.TgtEntityID
		e2NodeIDs = append(e2NodeIDs, e2NodeID)
	}

	return e2NodeIDs, nil
}

// E2TNodes lists E2T nodes
func (c *Client) E2TNodes(ctx context.Context) ([]topoapi.Object, error) {
	objects, err := c.client.List(ctx, toposdk.WithListFilters(getE2TFilter()))
	if err != nil {
		return nil, err
	}

	return append(make([]topoapi.Object, 0), objects...), nil
}

// E2Nodes lists E2 nodes
func (c *Client) E2Nodes(ctx context.Context) ([]topoapi.Object, error) {
	objects, err := c.client.List(ctx, toposdk.WithListFilters(getE2Filter()))
	if err != nil {
		return nil, err
	}

	return append(make([]topoapi.Object, 0), objects...), nil
}

// GetE2NodeAspects gets E2 node aspects
func (c *Client) GetE2NodeAspects(ctx context.Context, nodeID topoapi.ID) (*topoapi.E2Node, error) {
	object, err := c.client.Get(ctx, nodeID)
	if err != nil {
		return nil, err
	}
	e2Node := &topoapi.E2Node{}
	err = object.GetAspect(e2Node)
	if err != nil {
		return nil, err
	}

	return e2Node, nil

}

func (c *Client) GetE2NodeMastershipState(ctx context.Context, nodeID topoapi.ID) (*topoapi.MastershipState, error) {
	object, err := c.client.Get(ctx, nodeID)
	if err != nil {
		return nil, err
	}
	mastershipState := &topoapi.MastershipState{}
	err = object.GetAspect(mastershipState)
	if err != nil {
		return nil, err
	}

	return mastershipState, nil
}

// GetCells get list of cells for each E2 node
func (c *Client) GetCells(ctx context.Context, nodeID topoapi.ID) ([]*topoapi.E2Cell, error) {
	filter := &topoapi.Filters{
		RelationFilter: &topoapi.RelationFilter{
			SrcId:        string(nodeID),
			RelationKind: topoapi.CONTAINS,
			TargetKind:   topoapi.E2CELL,
		},
	}

	objects, err := c.client.List(ctx, toposdk.WithListFilters(filter))
	if err != nil {
		return nil, err
	}
	var cells []*topoapi.E2Cell
	for _, obj := range objects {
		cellObject := &topoapi.E2Cell{}
		err = obj.GetAspect(cellObject)
		if err != nil {
			return nil, err
		}
		cells = append(cells, cellObject)
	}

	return cells, nil
}

func getFilter(kind string) *topoapi.Filters {
	controlRelationFilter := &topoapi.Filters{
		KindFilter: &topoapi.Filter{
			Filter: &topoapi.Filter_Equal_{
				Equal_: &topoapi.EqualFilter{
					Value: kind,
				},
			},
		},
	}
	return controlRelationFilter

}

func getControlRelationFilter() *topoapi.Filters {
	return getFilter(topoapi.CONTROLS)
}

func getE2TFilter() *topoapi.Filters {
	return getFilter(topoapi.E2T)
}

func getE2Filter() *topoapi.Filters {
	return getFilter(topoapi.E2NODE)
}

// WatchE2Connections watch e2 node connection changes
func (c *Client) WatchE2Connections(ctx context.Context, ch chan topoapi.Event) error {
	err := c.client.Watch(ctx, ch, toposdk.WithWatchFilters(getControlRelationFilter()))
	if err != nil {
		return err
	}
	return nil
}

// WatchE2TNodes watch e2 node connection changes
func (c *Client) WatchE2TNodes(ctx context.Context, ch chan topoapi.Event) error {
	err := c.client.Watch(ctx, ch, toposdk.WithWatchFilters(getE2TFilter()))
	if err != nil {
		return err
	}
	return nil
}

// WatchE2Nodes watch e2 node connection changes
func (c *Client) WatchE2Nodes(ctx context.Context, ch chan topoapi.Event) error {
	err := c.client.Watch(ctx, ch, toposdk.WithWatchFilters(getE2Filter()))
	if err != nil {
		return err
	}
	return nil
}

var _ TopoClient = &Client{}

// GetTestNodeIDs gets n test node IDs
func GetTestNodeIDs(t *testing.T, n int) []topoapi.ID {
	topoSdkClient, err := NewTopoClient()
	assert.NoError(t, err)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	ch := make(chan topoapi.Event)
	err = topoSdkClient.WatchE2Connections(ctx, ch)
	assert.NoError(t, err)

	var nodeIDs []topoapi.ID
	for i := 0; i < n; i++ {
		event := <-ch
		t.Log(event.String())
		object := event.GetObject()
		assert.NotNil(t, object)
		nodeIDs = append(nodeIDs, object.GetRelation().GetTgtEntityID())
	}

	return nodeIDs
}

// GetTestNodeID gets one test node ID
func GetTestNodeID(t *testing.T) topoapi.ID {
	topoSdkClient, err := NewTopoClient()
	assert.NoError(t, err)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	ch := make(chan topoapi.Event)
	err = topoSdkClient.WatchE2Connections(ctx, ch)
	assert.NoError(t, err)
	event, ok := <-ch
	assert.True(t, ok)
	object := event.GetObject()
	assert.NotNil(t, object)
	return object.GetRelation().GetTgtEntityID()
}

type TopoEventCounters struct {
	Added       int
	Removed     int
	None        int
	Updated     int
	AddedOrNone int
}

func CountTopoRemovedEvent(ch chan topoapi.Event, expectedValue int) {
	eventCounters := TopoEventCounters{}
	for event := range ch {
		if event.Type == topoapi.EventType_REMOVED {
			eventCounters.Removed = eventCounters.Removed + 1
		}
		if eventCounters.Removed == expectedValue {
			break
		}
	}
}

func GetUpdatedEvent(ch chan topoapi.Event) (topoapi.Event, error) {
	for event := range ch {
		if event.Type == topoapi.EventType_UPDATED {
			return event, nil
		}
	}
	return topoapi.Event{}, errors.New("no updated event seen")
}

func CountTopoAddedOrNoneEvent(ch chan topoapi.Event, expectedValue int) {
	eventCounters := TopoEventCounters{}
	for event := range ch {
		if event.Type == topoapi.EventType_NONE || event.Type == topoapi.EventType_ADDED {
			eventCounters.AddedOrNone = eventCounters.AddedOrNone + 1
		}
		if eventCounters.AddedOrNone == expectedValue {
			break
		}
	}
}

func (c *Client) GetControlRelations() ([]topoapi.Object, error) {
	filter := getControlRelationFilter()
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	relationsList, err := c.client.List(ctx, toposdk.WithListFilters(filter))
	cancel()
	return relationsList, err
}
