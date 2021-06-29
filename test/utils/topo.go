// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package utils

import (
	"context"
	"testing"
	"time"

	toposdk "github.com/onosproject/onos-ric-sdk-go/pkg/topo"

	"github.com/stretchr/testify/assert"

	topoapi "github.com/onosproject/onos-api/go/onos/topo"
)

// TopoClient R-NIB client interface
type TopoClient interface {
	WatchE2Connections(ctx context.Context, ch chan topoapi.Event) error
	GetCells(ctx context.Context, nodeID topoapi.ID) ([]*topoapi.E2Cell, error)
	GetE2NodeAspects(ctx context.Context, nodeID topoapi.ID) (*topoapi.E2Node, error)
	E2NodeIDs(ctx context.Context) ([]topoapi.ID, error)
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

// E2NodeIDs lists all of connected E2 nodes
func (c *Client) E2NodeIDs(ctx context.Context) ([]topoapi.ID, error) {
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

// GetE2NodeAspects gets E2 node aspects
func (c *Client) GetE2NodeAspects(ctx context.Context, nodeID topoapi.ID) (*topoapi.E2Node, error) {
	object, err := c.client.Get(ctx, nodeID)
	if err != nil {
		return nil, err
	}
	e2Node := &topoapi.E2Node{}
	object.GetAspect(e2Node)

	return e2Node, nil

}

// GetCells get list of cells for each E2 node
func (c *Client) GetCells(ctx context.Context, nodeID topoapi.ID) ([]*topoapi.E2Cell, error) {
	filter := &topoapi.Filters{
		RelationFilter: &topoapi.RelationFilter{SrcId: string(nodeID),
			RelationKind: topoapi.CONTAINS,
			TargetKind:   ""}}

	objects, err := c.client.List(ctx, toposdk.WithListFilters(filter))
	if err != nil {
		return nil, err
	}
	var cells []*topoapi.E2Cell
	for _, obj := range objects {
		targetEntity := obj.GetEntity()
		if targetEntity.GetKindID() == topoapi.E2CELL {
			cellObject := &topoapi.E2Cell{}
			obj.GetAspect(cellObject)
			cells = append(cells, cellObject)
		}
	}

	return cells, nil
}

func getControlRelationFilter() *topoapi.Filters {
	controlRelationFilter := &topoapi.Filters{
		KindFilter: &topoapi.Filter{
			Filter: &topoapi.Filter_Equal_{
				Equal_: &topoapi.EqualFilter{
					Value: topoapi.CONTROLS,
				},
			},
		},
	}
	return controlRelationFilter
}

// WatchE2Connections watch e2 node connection changes
func (c *Client) WatchE2Connections(ctx context.Context, ch chan topoapi.Event) error {
	err := c.client.Watch(ctx, ch, toposdk.WithWatchFilters(getControlRelationFilter()))
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
	event := <-ch
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
