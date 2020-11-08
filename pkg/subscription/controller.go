// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package subscription

import (
	"context"
	"fmt"
	"github.com/gogo/protobuf/proto"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2ap-commondatatypes"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2apies"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2appducontents"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2appdudescriptions"
	subapi "github.com/onosproject/onos-e2t/api/subscription/v1beta1"
	"github.com/onosproject/onos-e2t/pkg/config"
	"github.com/onosproject/onos-e2t/pkg/controller"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2/channel"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2/channel/codec"
	channelfilter "github.com/onosproject/onos-e2t/pkg/southbound/e2/channel/filter"
	"github.com/onosproject/onos-e2t/pkg/store/subscription"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"time"
)

var log = logging.GetLogger("subscription", "controller")

const defaultTimeout = 30 * time.Second

// NewController returns a new network controller
func NewController(subs subscription.Store, channels *channel.Manager) *controller.Controller {
	c := controller.NewController("Subscription")
	c.Watch(&Watcher{
		subs: subs,
	})
	c.Watch(&ChannelWatcher{
		subs:     subs,
		channels: channels,
	})
	c.Reconcile(&Reconciler{
		subs:     subs,
		channels: channels,
	})
	return c
}

// Reconciler is a device change reconciler
type Reconciler struct {
	subs      subscription.Store
	channels  *channel.Manager
	requestID int32
}

// Reconcile reconciles the state of a device change
func (r *Reconciler) Reconcile(id controller.ID) (controller.Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	sub, err := r.subs.Get(ctx, subapi.ID(id.String()))

	// TODO: Unsubscribe on delete
	if err != nil {
		return controller.Result{}, err
	}

	switch sub.Status.State {
	case subapi.State_INACTIVE:
		channel, err := r.channels.Get(ctx, channel.ID(sub.E2NodeID))
		if err != nil {
			return controller.Result{}, err
		}

		r.requestID++
		requestID := r.requestID

		request := &e2appdudescriptions.E2ApPdu{}
		err = proto.Unmarshal(sub.Payload.Bytes, request)
		if err != nil {
			return controller.Result{}, err
		}

		// Generate a request ID
		ricRequestID := &e2apies.RicrequestId{
			RicRequestorId: requestID,
			RicInstanceId:  config.InstanceID,
		}

		// Update the subscription request with a request ID
		request.GetInitiatingMessage().ProcedureCode.RicSubscription.InitiatingMessage.ProtocolIes.E2ApProtocolIes29 = &e2appducontents.RicsubscriptionRequestIes_RicsubscriptionRequestIes29{
			Id:          int32(v1beta1.ProtocolIeIDRicrequestID),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value:       ricRequestID,
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}

		// Validate the subscribe request
		if err := request.Validate(); err != nil {
			return controller.Result{}, err
		}

		// Send the subscription request and await a response
		response, err := channel.SendRecv(request, channelfilter.RicSubscription(ricRequestID), codec.XER)
		if err != nil {
			return controller.Result{}, err
		}

		switch response.E2ApPdu.(type) {
		case *e2appdudescriptions.E2ApPdu_SuccessfulOutcome:
			sub.Status.State = subapi.State_ACTIVE
			sub.Status.E2TermID = config.InstanceID
			sub.Status.E2ConnID = uint64(channel.ID())
			sub.Status.E2RequestID = requestID
			if err := r.subs.Update(ctx, sub); err != nil {
				return controller.Result{}, err
			}
		case *e2appdudescriptions.E2ApPdu_UnsuccessfulOutcome:
			return controller.Result{}, fmt.Errorf("failed to initialize subscription %v", requestID)
		}
	case subapi.State_ACTIVE:
		c, err := r.channels.Get(ctx, channel.ID(sub.E2NodeID))
		if err != nil || sub.Status.E2TermID != config.InstanceID || channel.ID(sub.Status.E2ConnID) != c.ID() {
			sub.Status.State = subapi.State_INACTIVE
			sub.Status.E2TermID = 0
			sub.Status.E2ConnID = 0
			sub.Status.E2RequestID = 0
			if err := r.subs.Update(ctx, sub); err != nil {
				return controller.Result{}, err
			}
		}
	}
	return controller.Result{}, nil
}
