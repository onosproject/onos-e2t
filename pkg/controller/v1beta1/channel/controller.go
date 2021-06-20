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
	if channel.Status.State != e2api.ChannelState_CHANNEL_PENDING {
		return controller.Result{}, nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	sub, err := r.subs.Get(ctx, channel.SubscriptionID)
	if err != nil {
		if !errors.IsNotFound(err) {
			log.Warnf("Failed to reconcile Channel %+v: %s", channel, err)
			return controller.Result{}, err
		}

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
		log.Debugf("Creating Channel %+v Subscription %+v", channel, sub)
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
		log.Debugf("Binding Channel %+v to existing Subscription %+v", channel, sub)
		sub.Status.Channels = append(sub.Status.Channels, channel.ID)
		if err := r.subs.Update(ctx, sub); err != nil {
			log.Warnf("Failed to reconcile Channel %+v: %s", channel, err)
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
		log.Debugf("Completing Channel %+v: Subscription complete", channel)
		channel.Status.State = e2api.ChannelState_CHANNEL_COMPLETE
		if err := r.chans.Update(ctx, channel); err != nil {
			log.Warnf("Failed to reconcile Channel %+v: %s", channel, err)
			return controller.Result{}, err
		}
	case e2api.SubscriptionState_SUBSCRIPTION_FAILED:
		log.Debugf("Failing Channel %+v: Subscription failed", channel)
		channel.Status.State = e2api.ChannelState_CHANNEL_FAILED
		channel.Status.Error = sub.Status.Error
		if err := r.chans.Update(ctx, channel); err != nil {
			log.Warnf("Failed to reconcile Channel %+v: %s", channel, err)
			return controller.Result{}, err
		}
	}
	return controller.Result{}, nil
}

func (r *Reconciler) reconcileClosedChannel(channel *e2api.Channel) (controller.Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	// If the close has completed, delete the channel
	if channel.Status.State == e2api.ChannelState_CHANNEL_COMPLETE {
		log.Debugf("Deleting closed Channel %+v", channel)
		err := r.chans.Delete(ctx, channel)
		if err != nil && !errors.IsNotFound(err) {
			log.Warnf("Failed to reconcile Channel %+v: %s", channel, err)
			return controller.Result{}, err
		}
		return controller.Result{}, nil
	}

	sub, err := r.subs.Get(ctx, channel.SubscriptionID)
	if err != nil {
		if !errors.IsNotFound(err) {
			log.Warnf("Failed to reconcile Channel %+v: %s", channel, err)
			return controller.Result{}, err
		}

		// If the subscription is not found, complete the channel close
		log.Debugf("Completing closed Channel %+v: subscription not found", channel)
		channel.Status.State = e2api.ChannelState_CHANNEL_COMPLETE
		if err := r.chans.Update(ctx, channel); err != nil {
			log.Warnf("Failed to reconcile Channel %+v: %s", channel, err)
			return controller.Result{}, err
		}
		r.streams.CloseReader(channel.SubscriptionID, channel.AppID, channel.AppInstanceID)
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
		if len(channels) == 0 {
			log.Debugf("Closing Subscription %+v", sub)
			sub.Status.Phase = e2api.SubscriptionPhase_SUBSCRIPTION_CLOSED
			sub.Status.State = e2api.SubscriptionState_SUBSCRIPTION_PENDING
		} else {
			log.Debugf("Unbinding Channel %+v from Subscription %+v", channel, sub)
		}
		sub.Status.Channels = channels
		if err := r.subs.Update(ctx, sub); err != nil {
			log.Warnf("Failed to reconcile Channel %+v: %s", channel, err)
			return controller.Result{}, err
		}
		return controller.Result{}, nil
	}

	// If the subscription is OPEN or it's CLOSED phase with a finished state, update the channel state
	if sub.Status.Phase == e2api.SubscriptionPhase_SUBSCRIPTION_OPEN || sub.Status.State == e2api.SubscriptionState_SUBSCRIPTION_COMPLETE {
		log.Debugf("Completing close Channel %+v: subscription closed", channel)
		channel.Status.State = e2api.ChannelState_CHANNEL_COMPLETE
		if err := r.chans.Update(ctx, channel); err != nil {
			log.Warnf("Failed to reconcile Channel %+v: %s", channel, err)
			return controller.Result{}, err
		}
		r.streams.CloseReader(channel.SubscriptionID, channel.AppID, channel.AppInstanceID)
		return controller.Result{}, nil
	}
	return controller.Result{}, nil
}
