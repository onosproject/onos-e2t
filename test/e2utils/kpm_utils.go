// SPDX-FileCopyrightText: 2022-present Intel Corporation
// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2utils

import (
	"context"
	"github.com/onosproject/onos-api/go/onos/topo"
	"testing"
	"time"

	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"

	sdkclient "github.com/onosproject/onos-ric-sdk-go/pkg/e2/v1beta1"
	"github.com/stretchr/testify/assert"

	"github.com/onosproject/onos-e2t/test/utils"
)

const (
	defaultReportPeriod = uint32(5000)
	defaultGranularity  = uint32(500)
)

// Sub represents the common fields for subscriptions
type Sub struct {
	Name                string
	NodeID              topo.ID
	Node                sdkclient.Node
	SdkClient           sdkclient.Client
	Ch                  chan e2api.Indication
	Timeout             time.Duration
	Actions             []e2api.Action
	EventTriggerBytes   []byte
	ActionID            int32
	ActionType          e2api.ActionType
	ServiceModelName    e2api.ServiceModelName
	ServiceModelVersion e2api.ServiceModelVersion
	EncodingType        sdkclient.Encoding
}

// KPMV2Sub represents a KPM V2 service model subscription
type KPMV2Sub struct {
	Sub          Sub
	CellObjectID string
	ReportPeriod uint32
	Granularity  uint32
}

// RCPreSub represents an RC Pre service model subscription
type RCPreSub struct {
	Sub Sub
}

// Subscribe is the common subscription implementation for all service models
func (sub *Sub) Subscribe(ctx context.Context) (e2api.ChannelID, error) {
	if sub.EncodingType == 0 {
		sub.EncodingType = sdkclient.ProtoEncoding
	}
	if sub.ActionID == 0 {
		sub.ActionID = 100
	}
	if sub.ActionType == 0 {
		sub.ActionType = e2api.ActionType_ACTION_TYPE_REPORT
	}
	subRequest := utils.Subscription{
		NodeID:              string(sub.NodeID),
		EventTrigger:        sub.EventTriggerBytes,
		ServiceModelName:    sub.ServiceModelName,
		ServiceModelVersion: sub.ServiceModelVersion,
		Actions:             sub.Actions,
	}

	subSpec, err := subRequest.CreateWithActionDefinition()
	if err != nil {
		return "", err
	}

	sub.SdkClient = utils.GetE2Client(nil, string(sub.ServiceModelName), string(sub.ServiceModelVersion), sdkclient.ProtoEncoding)
	sub.Node = sub.SdkClient.Node(sdkclient.NodeID(sub.NodeID))
	sub.Ch = make(chan e2api.Indication)
	subscribeOptions := make([]sdkclient.SubscribeOption, 0)
	if sub.Timeout != 0 {
		subscribeOptions = append(subscribeOptions, sdkclient.WithTransactionTimeout(sub.Timeout))
	}

	return sub.Node.Subscribe(ctx, sub.Name, subSpec, sub.Ch, subscribeOptions...)
}

// UnsubscribeOrFail unsubscribes from the service model.
// If an error occurs, the test is failed.
func (sub *Sub) UnsubscribeOrFail(ctx context.Context, t *testing.T) {
	assert.NoError(t, sub.Node.Unsubscribe(ctx, sub.Name))
}

// Unsubscribe from the service model
func (sub *Sub) Unsubscribe(ctx context.Context) error {
	return sub.Node.Unsubscribe(ctx, sub.Name)
}

// Subscribe is the KPM V2 service model specific implementation
func (sub *KPMV2Sub) Subscribe(ctx context.Context) (e2api.ChannelID, error) {
	if sub.ReportPeriod == 0 {
		sub.ReportPeriod = defaultReportPeriod
	}
	if sub.Granularity == 0 {
		sub.Granularity = defaultGranularity
	}
	if sub.Sub.ServiceModelVersion == "" {
		sub.Sub.ServiceModelVersion = utils.Version2
	}
	if sub.Sub.ServiceModelName == "" {
		sub.Sub.ServiceModelName = utils.KpmServiceModelName
	}

	if len(sub.Sub.EventTriggerBytes) == 0 {
		eventTriggerBytes, err := utils.CreateKpmV2EventTrigger(sub.ReportPeriod)
		if err != nil {
			return "", err
		}
		sub.Sub.EventTriggerBytes = eventTriggerBytes
	}

	actionDefinitionBytes, err := utils.CreateKpmV2ActionDefinition(sub.CellObjectID, sub.Granularity)
	if err != nil {
		return "", err
	}

	if len(sub.Sub.Actions) == 0 {
		action := e2api.Action{
			ID:   sub.Sub.ActionID,
			Type: sub.Sub.ActionType,
			SubsequentAction: &e2api.SubsequentAction{
				Type:       e2api.SubsequentActionType_SUBSEQUENT_ACTION_TYPE_CONTINUE,
				TimeToWait: e2api.TimeToWait_TIME_TO_WAIT_ZERO,
			},
			Payload: actionDefinitionBytes,
		}

		sub.Sub.Actions = append(sub.Sub.Actions, action)
	}

	return sub.Sub.Subscribe(ctx)
}

// SubscribeOrFail subscribes to the KPM V2 service model.
// If an error occurs, the test is failed.
func (sub *KPMV2Sub) SubscribeOrFail(ctx context.Context, t *testing.T) e2api.ChannelID {
	channelID, err := sub.Subscribe(ctx)
	assert.NoError(t, err)
	return channelID
}

// Subscribe is the RC Pre service model specific implementation
func (sub *RCPreSub) Subscribe(ctx context.Context) (e2api.ChannelID, error) {
	if sub.Sub.ServiceModelVersion == "" {
		sub.Sub.ServiceModelVersion = utils.Version2
	}
	if sub.Sub.ServiceModelName == "" {
		sub.Sub.ServiceModelName = utils.RcServiceModelName
	}

	if len(sub.Sub.EventTriggerBytes) == 0 {
		eventTriggerBytes, err := utils.CreateRcEventTrigger()
		if err != nil {
			return "", err
		}
		sub.Sub.EventTriggerBytes = eventTriggerBytes
	}

	if len(sub.Sub.Actions) == 0 {
		RCAction := e2api.Action{
			ID:   sub.Sub.ActionID,
			Type: e2api.ActionType_ACTION_TYPE_REPORT,
			SubsequentAction: &e2api.SubsequentAction{
				Type:       e2api.SubsequentActionType_SUBSEQUENT_ACTION_TYPE_CONTINUE,
				TimeToWait: e2api.TimeToWait_TIME_TO_WAIT_ZERO,
			},
		}
		sub.Sub.Actions = append(sub.Sub.Actions, RCAction)
	}

	return sub.Sub.Subscribe(ctx)
}

// SubscribeOrFail subscribes to the RC Pre service model.
// If an error occurs, the test is failed.
func (sub *RCPreSub) SubscribeOrFail(ctx context.Context, t *testing.T) e2api.ChannelID {
	channelID, err := sub.Subscribe(ctx)
	assert.NoError(t, err)
	return channelID
}
