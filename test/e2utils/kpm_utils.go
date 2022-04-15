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

	if sub.Sub.EventTriggerBytes == nil {
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

	var actions []e2api.Action
	if sub.Sub.Actions != nil {
		actions = sub.Sub.Actions
	} else {
		action := e2api.Action{
			ID:   sub.Sub.ActionID,
			Type: sub.Sub.ActionType,
			SubsequentAction: &e2api.SubsequentAction{
				Type:       e2api.SubsequentActionType_SUBSEQUENT_ACTION_TYPE_CONTINUE,
				TimeToWait: e2api.TimeToWait_TIME_TO_WAIT_ZERO,
			},
			Payload: actionDefinitionBytes,
		}

		actions = append(actions, action)
	}

	subRequest := utils.Subscription{
		NodeID:              string(sub.Sub.NodeID),
		EventTrigger:        sub.Sub.EventTriggerBytes,
		ServiceModelName:    sub.Sub.ServiceModelName,
		ServiceModelVersion: sub.Sub.ServiceModelVersion,
		Actions:             actions,
	}

	subSpec, err := subRequest.CreateWithActionDefinition()
	if err != nil {
		return "", err
	}

	sub.Sub.SdkClient = utils.GetE2Client(nil, string(sub.Sub.ServiceModelName), string(sub.Sub.ServiceModelVersion), sdkclient.ProtoEncoding)
	sub.Sub.Node = sub.Sub.SdkClient.Node(sdkclient.NodeID(sub.Sub.NodeID))
	sub.Sub.Ch = make(chan e2api.Indication)
	subscribeOptions := make([]sdkclient.SubscribeOption, 0)
	if sub.Sub.Timeout != 0 {
		subscribeOptions = append(subscribeOptions, sdkclient.WithTransactionTimeout(sub.Sub.Timeout))
	}

	return sub.Sub.Node.Subscribe(ctx, sub.Sub.Name, subSpec, sub.Sub.Ch, subscribeOptions...)
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
