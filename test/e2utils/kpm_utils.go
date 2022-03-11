// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2utils

import (
	"context"
	"github.com/onosproject/onos-api/go/onos/topo"
	"testing"

	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"

	sdkclient "github.com/onosproject/onos-ric-sdk-go/pkg/e2/v1beta1"
	"github.com/stretchr/testify/assert"

	"github.com/onosproject/onos-e2t/test/utils"
)

const (
	defaultReportPeriod = uint32(5000)
	defaultGranularity  = uint32(500)
)

type KPMV2Sub struct {
	Ctx          context.Context
	SubName      string
	NodeID       topo.ID
	Node         sdkclient.Node
	CellObjectID string
	ReportPeriod uint32
	Granularity  uint32
	SdkClient    sdkclient.Client
	Ch           chan e2api.Indication
}

func (sub *KPMV2Sub) Subscribe() (e2api.ChannelID, error) {
	if sub.ReportPeriod == 0 {
		sub.ReportPeriod = defaultReportPeriod
	}

	if sub.Granularity == 0 {
		sub.Granularity = defaultGranularity
	}
	eventTriggerBytes, err := utils.CreateKpmV2EventTrigger(sub.ReportPeriod)
	if err != nil {
		return "", err
	}

	actionDefinitionBytes, err := utils.CreateKpmV2ActionDefinition(sub.CellObjectID, sub.Granularity)
	if err != nil {
		return "", err
	}

	var actions []e2api.Action
	action := e2api.Action{
		ID:   100,
		Type: e2api.ActionType_ACTION_TYPE_REPORT,
		SubsequentAction: &e2api.SubsequentAction{
			Type:       e2api.SubsequentActionType_SUBSEQUENT_ACTION_TYPE_CONTINUE,
			TimeToWait: e2api.TimeToWait_TIME_TO_WAIT_ZERO,
		},
		Payload: actionDefinitionBytes,
	}

	actions = append(actions, action)

	subRequest := utils.Subscription{
		NodeID:              string(sub.NodeID),
		EventTrigger:        eventTriggerBytes,
		ServiceModelName:    utils.KpmServiceModelName,
		ServiceModelVersion: utils.Version2,
		Actions:             actions,
	}

	subSpec, err := subRequest.CreateWithActionDefinition()
	if err != nil {
		return "", err
	}

	sub.SdkClient = utils.GetE2Client(nil, utils.KpmServiceModelName, utils.Version2, sdkclient.ProtoEncoding)
	sub.Node = sub.SdkClient.Node(sdkclient.NodeID(sub.NodeID))
	sub.Ch = make(chan e2api.Indication)
	return sub.Node.Subscribe(sub.Ctx, sub.SubName, subSpec, sub.Ch)
}

func (sub *KPMV2Sub) CreateKPMV2SubscriptionOrFail(t *testing.T) {
	_, err := sub.Subscribe()
	assert.NoError(t, err)
}

func (sub *KPMV2Sub) Unsubscribe() error {
	return sub.Node.Unsubscribe(sub.Ctx, sub.SubName)
}
