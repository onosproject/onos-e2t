// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package channel

import (
	"context"
	subscription "github.com/onosproject/onos-e2t/pkg/broker/subscription/v1beta1"
	"time"

	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	chanstore "github.com/onosproject/onos-e2t/pkg/store/channel"
	substore "github.com/onosproject/onos-e2t/pkg/store/subscription"
	"github.com/onosproject/onos-lib-go/pkg/controller"
	"github.com/onosproject/onos-lib-go/pkg/errors"
	"github.com/onosproject/onos-lib-go/pkg/logging"
)

var log = logging.GetLogger("controller", "channel")

const defaultTimeout = 30 * time.Second

// NewController returns a new channel controller
func NewController(chans chanstore.Store, subs substore.Store, streams subscription.Broker) *controller.Controller {
	c := controller.NewController("Channel")
	c.Watch(&Watcher{
		chans: chans,
	})
	c.Watch(&TaskWatcher{
		chans: chans,
		subs:  subs,
	})
	c.Reconcile(&Reconciler{
		chans:   chans,
		subs:    subs,
		streams: streams,
	})
	return c
}

// Reconciler is a channel reconciler
type Reconciler struct {
	chans   chanstore.Store
	subs    substore.Store
	streams subscription.Broker
}

// Reconcile reconciles the state of a channel
func (r *Reconciler) Reconcile(id controller.ID) (controller.Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	channel, err := r.chans.Get(ctx, id.Value.(e2api.ChannelID))
	if err != nil {
		if errors.IsNotFound(err) {
			return controller.Result{}, nil
		}
		return controller.Result{}, err
	}

	log.Infof("Reconciling Channel %+v", channel)

	// If the channel state is COMPLETE or FAILED, ignore the request
	switch channel.Status.State {
	case e2api.ChannelState_CHANNEL_COMPLETE, e2api.ChannelState_CHANNEL_FAILED:
		return controller.Result{}, nil
	}

	// Reconcile the channel state according to its phase
	switch channel.Status.Phase {
	case e2api.ChannelPhase_CHANNEL_OPEN:
		return r.reconcileOpenChannel(channel)
	case e2api.ChannelPhase_CHANNEL_CLOSED:
		return r.reconcileClosedChannel(channel)
	}
	return controller.Result{}, nil
}

func (r *Reconciler) reconcileOpenChannel(channel *e2api.Channel) (controller.Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	sub, err := r.subs.Get(ctx, channel.SubscriptionID)
	if err != nil {
		if !errors.IsNotFound(err) {
			log.Warnf("Failed to reconcile Channel %+v: %s", channel, err)
			return controller.Result{}, err
		}

		log.Warnf("Creating Subscription for Channel %+v: %s", channel, err)
		sub := &e2api.Subscription{
			ID: channel.SubscriptionID,
			SubscriptionMeta: e2api.SubscriptionMeta{
				E2NodeID:     channel.E2NodeID,
				ServiceModel: channel.ServiceModel,
				Encoding:     channel.Encoding,
			},
			Spec: channel.Spec.SubscriptionSpec,
			Status: e2api.SubscriptionStatus{
				Phase:    e2api.SubscriptionPhase_SUBSCRIPTION_OPEN,
				Channels: []e2api.ChannelID{channel.ID},
			},
		}
		err := r.subs.Create(ctx, sub)
		if err != nil {
			if !errors.IsAlreadyExists(err) {
				log.Warnf("Failed to reconcile Channel %+v: %s", sub, err)
				return controller.Result{}, err
			}
		}
		return controller.Result{}, nil
	}

	// Create a set of linked northbound channels for the subscription
	channels := make(map[e2api.ChannelID]bool)
	for _, id := range sub.Status.Channels {
		channels[id] = true
	}

	// If the channel is not linked to the subscription, add it
	if _, ok := channels[channel.ID]; !ok {
		sub.Status.Channels = append(sub.Status.Channels, channel.ID)
		if err := r.subs.Update(ctx, sub); err != nil {
			return controller.Result{}, err
		}
		return controller.Result{}, nil
	}

	// If the subscription is not yet in the OPEN phase, skip reconciliation for this request
	if sub.Status.Phase != e2api.SubscriptionPhase_SUBSCRIPTION_OPEN {
		return controller.Result{}, nil
	}

	// If the subscription is in a finished state, update the channel state
	switch sub.Status.State {
	case e2api.SubscriptionState_SUBSCRIPTION_COMPLETE:
		channel.Status.State = e2api.ChannelState_CHANNEL_COMPLETE
		if err := r.chans.Update(ctx, channel); err != nil {
			return controller.Result{}, err
		}
	case e2api.SubscriptionState_SUBSCRIPTION_FAILED:
		channel.Status.State = e2api.ChannelState_CHANNEL_FAILED
		channel.Status.Error = sub.Status.Error
		if err := r.chans.Update(ctx, channel); err != nil {
			return controller.Result{}, err
		}
	}
	return controller.Result{}, nil
}

func (r *Reconciler) reconcileClosedChannel(channel *e2api.Channel) (controller.Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	sub, err := r.subs.Get(ctx, channel.SubscriptionID)
	if err != nil {
		if !errors.IsNotFound(err) {
			log.Warnf("Failed to reconcile Channel %+v: %s", channel, err)
			return controller.Result{}, err
		}
		return controller.Result{}, nil
	}

	// If the subscription is in the CLOSED phase with a finished state, update the channel state
	if sub.Status.Phase == e2api.SubscriptionPhase_SUBSCRIPTION_CLOSED {
		switch sub.Status.State {
		case e2api.SubscriptionState_SUBSCRIPTION_COMPLETE:
			channel.Status.State = e2api.ChannelState_CHANNEL_COMPLETE
			if err := r.chans.Update(ctx, channel); err != nil {
				return controller.Result{}, err
			}
			r.streams.CloseReader(channel.SubscriptionID, channel.AppID, channel.AppInstanceID)
		case e2api.SubscriptionState_SUBSCRIPTION_FAILED:
			channel.Status.State = e2api.ChannelState_CHANNEL_FAILED
			channel.Status.Error = sub.Status.Error
			if err := r.chans.Update(ctx, channel); err != nil {
				return controller.Result{}, err
			}
			r.streams.CloseReader(channel.SubscriptionID, channel.AppID, channel.AppInstanceID)
		}
		return controller.Result{}, nil
	}

	// Create a list of northbound channels linked to the subscription, excluding the request channel
	channels := make([]e2api.ChannelID, 0, len(sub.Status.Channels))
	for _, id := range sub.Status.Channels {
		if id != channel.ID {
			channels = append(channels, id)
		}
	}

	// If the linked channels changed, update the subscription status
	if len(sub.Status.Channels) != len(channels) {
		sub.Status.Channels = channels
		if err := r.subs.Update(ctx, sub); err != nil {
			return controller.Result{}, err
		}
		return controller.Result{}, nil
	}

	// If northbound channels are still linked to the subscription, the subscription will remain open
	// Complete the channel phase
	if len(sub.Status.Channels) > 0 {
		channel.Status.State = e2api.ChannelState_CHANNEL_COMPLETE
		if err := r.chans.Update(ctx, channel); err != nil {
			return controller.Result{}, err
		}
		r.streams.CloseReader(channel.SubscriptionID, channel.AppID, channel.AppInstanceID)
	}
	return controller.Result{}, nil
}
