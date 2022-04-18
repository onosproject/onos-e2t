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

type KPMV2Sub struct {
	Sub          Sub
	CellObjectID string
	ReportPeriod uint32
	Granularity  uint32
}

type RCPreSub struct {
	Sub Sub
}

func (sub *KPMV2Sub) Subscribe(ctx context.Context) (e2api.ChannelID, error) {
	if sub.ReportPeriod == 0 {
		sub.ReportPeriod = defaultReportPeriod
	}
	if sub.Granularity == 0 {
		sub.Granularity = defaultGranularity
	}
	if sub.Sub.ActionID == 0 {
		sub.Sub.ActionID = 100
	}
	if sub.Sub.ActionType == 0 {
		sub.Sub.ActionType = e2api.ActionType_ACTION_TYPE_REPORT
	}
	if sub.Sub.ServiceModelVersion == "" {
		sub.Sub.ServiceModelVersion = utils.Version2
	}
	if sub.Sub.ServiceModelName == "" {
		sub.Sub.ServiceModelName = utils.KpmServiceModelName
	}
	if sub.Sub.EncodingType == 0 {
		sub.Sub.EncodingType = sdkclient.ProtoEncoding
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

func (sub *RCPreSub) Subscribe(ctx context.Context) (e2api.ChannelID, error) {
	if sub.Sub.ActionID == 0 {
		sub.Sub.ActionID = 100
	}
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

func (sub *Sub) Subscribe(ctx context.Context) (e2api.ChannelID, error) {
	if sub.EncodingType == 0 {
		sub.EncodingType = sdkclient.ProtoEncoding
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

func (sub *KPMV2Sub) CreateKPMV2SubscriptionOrFail(ctx context.Context, t *testing.T) {
	_, err := sub.Subscribe(ctx)
	assert.NoError(t, err)
}

func (sub *KPMV2Sub) Unsubscribe(ctx context.Context) error {
	return sub.Sub.Node.Unsubscribe(ctx, sub.Sub.Name)
}

func (sub *KPMV2Sub) SubscribeOrFail(ctx context.Context, t *testing.T) e2api.ChannelID {
	channelID, err := sub.Subscribe(ctx)
	assert.NoError(t, err)
	return channelID
}

func (sub *KPMV2Sub) UnsubscribeOrFail(ctx context.Context, t *testing.T) {
	assert.NoError(t, sub.Sub.Node.Unsubscribe(ctx, sub.Sub.Name))
}

func (sub *RCPreSub) SubscribeOrFail(ctx context.Context, t *testing.T) e2api.ChannelID {
	channelID, err := sub.Subscribe(ctx)
	assert.NoError(t, err)
	return channelID
}

func (sub *RCPreSub) UnsubscribeOrFail(ctx context.Context, t *testing.T) {
	assert.NoError(t, sub.Sub.Node.Unsubscribe(ctx, sub.Sub.Name))
}
