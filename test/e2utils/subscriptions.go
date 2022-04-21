// SPDX-FileCopyrightText: 2022-present Intel Corporation
// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2utils

import (
	"context"
	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	"github.com/onosproject/onos-api/go/onos/topo"
	e2smkpmv2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/pdubuilder"
	rcpdubuilder "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre_go/pdubuilder"
	"google.golang.org/protobuf/proto"
	"testing"
	"time"

	sdkclient "github.com/onosproject/onos-ric-sdk-go/pkg/e2/v1beta1"
	"github.com/stretchr/testify/assert"

	kpmv2types "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
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

	subSpec := e2api.SubscriptionSpec{
		EventTrigger: e2api.EventTrigger{
			Payload: sub.EventTriggerBytes,
		},

		Actions: sub.Actions,
	}

	if sub.SdkClient == nil {
		sub.SdkClient = sdkclient.NewClient(sdkclient.WithE2TAddress(utils.E2TServiceHost, utils.E2TServicePort),
			sdkclient.WithServiceModel(sdkclient.ServiceModelName(sub.ServiceModelName),
				sdkclient.ServiceModelVersion(sub.ServiceModelVersion)),
			sdkclient.WithEncoding(sub.EncodingType))
	}
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

// createKpmV2EventTrigger creates a kpm v2 service model event trigger
func (sub *KPMV2Sub) createKpmV2EventTrigger() ([]byte, error) {
	e2SmKpmEventTriggerDefinition, err := e2smkpmv2.CreateE2SmKpmEventTriggerDefinition(int64(sub.ReportPeriod))
	if err != nil {
		return []byte{}, err
	}
	err = e2SmKpmEventTriggerDefinition.Validate()
	if err != nil {
		return []byte{}, err
	}
	protoBytes, err := proto.Marshal(e2SmKpmEventTriggerDefinition)
	if err != nil {
		return []byte{}, err
	}
	return protoBytes, nil
}

func (sub *KPMV2Sub) CreateKpmV2ActionDefinition() ([]byte, error) {
	rrcConAvgName, err := e2smkpmv2.CreateMeasurementTypeMeasName("RRC.Conn.Avg")
	if err != nil {
		return nil, err
	}
	rrcConMaxName, err := e2smkpmv2.CreateMeasurementTypeMeasName("RRC.Conn.Max")
	if err != nil {
		return nil, err
	}
	rrcConnEstabAtt, err := e2smkpmv2.CreateMeasurementTypeMeasName("RRC.ConnEstabAtt.Sum")
	if err != nil {
		return nil, err
	}
	measInfoConAvgItem, err := e2smkpmv2.CreateMeasurementInfoItem(rrcConAvgName)
	if err != nil {
		return nil, err
	}
	measInfoConMaxItem, err := e2smkpmv2.CreateMeasurementInfoItem(rrcConMaxName)
	if err != nil {
		return nil, err
	}
	measInfoConnEstabAttItem, err := e2smkpmv2.CreateMeasurementInfoItem(rrcConnEstabAtt)
	if err != nil {
		return nil, err
	}

	measInfoList := &kpmv2types.MeasurementInfoList{
		Value: make([]*kpmv2types.MeasurementInfoItem, 0),
	}
	measInfoList.Value = append(measInfoList.Value, measInfoConAvgItem)
	measInfoList.Value = append(measInfoList.Value, measInfoConMaxItem)
	measInfoList.Value = append(measInfoList.Value, measInfoConnEstabAttItem)

	actionDefinition, err := e2smkpmv2.CreateActionDefinitionFormat1(sub.CellObjectID, measInfoList, int64(sub.Granularity), 1234)
	if err != nil {
		return nil, err
	}

	e2SmKpmActionDefinition, err := e2smkpmv2.CreateE2SmKpmActionDefinitionFormat1(1, actionDefinition)
	if err != nil {
		return nil, err
	}

	protoBytes, err := proto.Marshal(e2SmKpmActionDefinition)
	if err != nil {
		return nil, err
	}
	return protoBytes, nil
}

// init applies default values to the KPM V2 subscription
func (sub *KPMV2Sub) init() error {
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
		eventTriggerBytes, err := sub.createKpmV2EventTrigger()
		if err != nil {
			return err
		}
		sub.Sub.EventTriggerBytes = eventTriggerBytes
	}

	return nil
}

// UseDefaultReportAction sets up a default KPM V2 REPORT action
func (sub *KPMV2Sub) UseDefaultReportAction() error {
	err := sub.init()
	if err != nil {
		return err
	}

	actionDefinitionBytes, err := sub.CreateKpmV2ActionDefinition()
	if err != nil {
		return err
	}

	action := e2api.Action{
		ID:   100,
		Type: e2api.ActionType_ACTION_TYPE_REPORT,
		SubsequentAction: &e2api.SubsequentAction{
			Type:       e2api.SubsequentActionType_SUBSEQUENT_ACTION_TYPE_CONTINUE,
			TimeToWait: e2api.TimeToWait_TIME_TO_WAIT_ZERO,
		},
		Payload: actionDefinitionBytes,
	}

	sub.Sub.Actions = append(sub.Sub.Actions, action)

	return nil
}

// Subscribe is the KPM V2 service model specific implementation
func (sub *KPMV2Sub) Subscribe(ctx context.Context) (e2api.ChannelID, error) {
	err := sub.init()
	if err != nil {
		return "", err
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

// CreateSubscriptionSpec return the subscription specification for this subscription
func (sub *KPMV2Sub) CreateSubscriptionSpec() (e2api.SubscriptionSpec, error) {
	err := sub.init()
	if err != nil {
		return e2api.SubscriptionSpec{}, err
	}
	spec := e2api.SubscriptionSpec{
		EventTrigger: e2api.EventTrigger{
			Payload: sub.Sub.EventTriggerBytes,
		},

		Actions: sub.Sub.Actions,
	}
	return spec, nil
}

// createRcEventTrigger creates a rc service model event trigger
func (sub *RCPreSub) createRcEventTrigger() ([]byte, error) {
	e2SmKpmEventTriggerDefinition, err := rcpdubuilder.CreateE2SmRcPreEventTriggerDefinitionUponChange()
	if err != nil {
		return []byte{}, err
	}
	err = e2SmKpmEventTriggerDefinition.Validate()
	if err != nil {
		return []byte{}, err
	}
	protoBytes, err := proto.Marshal(e2SmKpmEventTriggerDefinition)
	if err != nil {
		return []byte{}, err
	}
	return protoBytes, nil
}

// UseDefaultReportAction sets up a default RC Pre REPORT action
func (sub *RCPreSub) UseDefaultReportAction() error {
	err := sub.init()
	if err != nil {
		return err
	}

	if len(sub.Sub.Actions) == 0 {
		RCAction := e2api.Action{
			ID:   100,
			Type: e2api.ActionType_ACTION_TYPE_REPORT,
			SubsequentAction: &e2api.SubsequentAction{
				Type:       e2api.SubsequentActionType_SUBSEQUENT_ACTION_TYPE_CONTINUE,
				TimeToWait: e2api.TimeToWait_TIME_TO_WAIT_ZERO,
			},
		}
		sub.Sub.Actions = append(sub.Sub.Actions, RCAction)
	}
	return nil
}

// init applies default values to the RC Pre subscription
func (sub *RCPreSub) init() error {
	if sub.Sub.ServiceModelVersion == "" {
		sub.Sub.ServiceModelVersion = utils.Version2
	}
	if sub.Sub.ServiceModelName == "" {
		sub.Sub.ServiceModelName = utils.RcServiceModelName
	}

	if len(sub.Sub.EventTriggerBytes) == 0 {
		eventTriggerBytes, err := sub.createRcEventTrigger()
		if err != nil {
			return err
		}
		sub.Sub.EventTriggerBytes = eventTriggerBytes
	}

	return nil
}

// Subscribe is the RC Pre service model specific implementation
func (sub *RCPreSub) Subscribe(ctx context.Context) (e2api.ChannelID, error) {
	err := sub.init()
	if err != nil {
		return "", err
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
