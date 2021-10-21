// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package channel

import (
	"context"
	topoapi "github.com/onosproject/onos-api/go/onos/topo"
	"github.com/onosproject/onos-e2t/pkg/broker"
	"github.com/onosproject/onos-e2t/pkg/controller/utils"
	"github.com/onosproject/onos-e2t/pkg/store/rnib"
	"time"

	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	chanstore "github.com/onosproject/onos-e2t/pkg/store/channel"
	substore "github.com/onosproject/onos-e2t/pkg/store/subscription"
	"github.com/onosproject/onos-lib-go/pkg/controller"
	"github.com/onosproject/onos-lib-go/pkg/errors"
	"github.com/onosproject/onos-lib-go/pkg/logging"
)

var log = logging.GetLogger("controller", "channel")

const (
	defaultTimeout            = 30 * time.Second
	defaultTransactionTimeout = 60 * time.Second
)

// NewController returns a new channel controller
func NewController(chans chanstore.Store, subs substore.Store, streams broker.Broker, topo rnib.Store) *controller.Controller {
	c := controller.NewController("Channel")
	c.Watch(&Watcher{
		chans: chans,
	})
	c.Watch(&SubscriptionWatcher{
		chans: chans,
		subs:  subs,
	})
	c.Watch(&TopoWatcher{
		chans: chans,
		topo:  topo,
	})
	c.Reconcile(&Reconciler{
		chans:   chans,
		subs:    subs,
		streams: streams,
		topo:    topo,
	})
	return c
}

// Reconciler is a channel reconciler
type Reconciler struct {
	chans   chanstore.Store
	subs    substore.Store
	streams broker.Broker
	topo    rnib.Store
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

	if ok, err := r.finalizeChannel(ctx, channel); err != nil {
		return controller.Result{}, err
	} else if ok {
		return controller.Result{}, nil
	}

	if ok, err := r.reconcileMastership(ctx, channel); err != nil {
		return controller.Result{}, err
	} else if ok {
		return controller.Result{}, nil
	}

	// Reconcile the channel state according to its phase
	switch channel.Status.Phase {
	case e2api.ChannelPhase_CHANNEL_OPEN:
		return r.reconcileOpenChannel(ctx, channel)
	case e2api.ChannelPhase_CHANNEL_CLOSED:
		return r.reconcileClosedChannel(ctx, channel)
	}
	return controller.Result{}, nil
}

func (r *Reconciler) reconcileOpenChannel(ctx context.Context, channel *e2api.Channel) (controller.Result, error) {
	if channel.Status.Timestamp != nil {
		transactionTimeout := defaultTransactionTimeout
		if channel.Spec.TransactionTimeout != nil {
			transactionTimeout = *channel.Spec.TransactionTimeout
		}

		expireTime := channel.Status.Timestamp.Add(transactionTimeout)
		if time.Now().After(expireTime) {
			log.Infof("Channel timeout, Closing channel  %s", channel.ID)
			channel.Status.Phase = e2api.ChannelPhase_CHANNEL_CLOSED
			channel.Status.State = e2api.ChannelState_CHANNEL_PENDING
			channel.Status.Error = nil
			if err := r.chans.Update(ctx, channel); err != nil && !errors.IsNotFound(err) && !errors.IsConflict(err) {
				log.Warnf("Failed to update channel %s: %s", channel.ID, err)
				return controller.Result{}, err
			}
			return controller.Result{}, nil
		}

		log.Debugf("Reconcile the closing channel at: %v", expireTime)
		return controller.Result{
			RequeueAt: expireTime,
		}, nil
	}

	if channel.Status.State != e2api.ChannelState_CHANNEL_PENDING {
		return controller.Result{}, nil
	}

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
		if err != nil && !errors.IsAlreadyExists(err) {
			log.Warnf("Failed to reconcile Channel %+v: %s", sub, err)
			return controller.Result{}, err
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
		if err := r.subs.Update(ctx, sub); err != nil && !errors.IsNotFound(err) && !errors.IsConflict(err) {
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
		if err := r.chans.Update(ctx, channel); err != nil && !errors.IsNotFound(err) && !errors.IsConflict(err) {
			log.Warnf("Failed to reconcile Channel %+v: %s", channel, err)
			return controller.Result{}, err
		}
	case e2api.SubscriptionState_SUBSCRIPTION_FAILED:
		log.Debugf("Failing Channel %+v: Subscription failed", channel)
		channel.Status.State = e2api.ChannelState_CHANNEL_FAILED
		channel.Status.Error = sub.Status.Error
		if err := r.chans.Update(ctx, channel); err != nil && !errors.IsNotFound(err) && !errors.IsConflict(err) {
			log.Warnf("Failed to reconcile Channel %+v: %s", channel, err)
			return controller.Result{}, err
		}
	}
	return controller.Result{}, nil
}

func (r *Reconciler) reconcileClosedChannel(ctx context.Context, channel *e2api.Channel) (controller.Result, error) {
	// If the close has completed, delete the channel
	if channel.Status.State == e2api.ChannelState_CHANNEL_COMPLETE {
		if len(channel.Finalizers) == 0 {
			log.Debugf("Deleting closed Channel %+v", channel)
			err := r.chans.Delete(ctx, channel)
			if err != nil && !errors.IsNotFound(err) {
				log.Warnf("Failed to reconcile Channel %+v: %s", channel, err)
				return controller.Result{}, err
			}
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
		if err := r.chans.Update(ctx, channel); err != nil && !errors.IsNotFound(err) && !errors.IsConflict(err) {
			log.Warnf("Failed to reconcile Channel %+v: %s", channel, err)
			return controller.Result{}, err
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
		log.Debugf("Unbinding Channel %+v from Subscription %+v", channel, sub)
		sub.Status.Channels = channels
		if err := r.subs.Update(ctx, sub); err != nil && !errors.IsNotFound(err) && !errors.IsConflict(err) {
			log.Warnf("Failed to reconcile Channel %+v: %s", channel, err)
			return controller.Result{}, err
		}
		return controller.Result{}, nil
	}

	// If the subscription is OPEN or it's CLOSED phase with a finished state, update the channel state
	if sub.Status.Phase == e2api.SubscriptionPhase_SUBSCRIPTION_OPEN || sub.Status.State == e2api.SubscriptionState_SUBSCRIPTION_COMPLETE {
		log.Debugf("Completing close Channel %+v: subscription closed", channel)
		channel.Status.State = e2api.ChannelState_CHANNEL_COMPLETE
		if err := r.chans.Update(ctx, channel); err != nil && !errors.IsNotFound(err) && !errors.IsConflict(err) {
			log.Warnf("Failed to reconcile Channel %+v: %s", channel, err)
			return controller.Result{}, err
		}
		return controller.Result{}, nil
	}
	return controller.Result{}, nil
}

func (r *Reconciler) reconcileMastership(ctx context.Context, channel *e2api.Channel) (bool, error) {
	log.Debugf("Fetching mastership state for E2Node '%s'", channel.E2NodeID)

	// Get the E2 node entity from topo
	e2NodeEntity, err := r.topo.Get(ctx, topoapi.ID(channel.E2NodeID))
	if err != nil {
		if !errors.IsNotFound(err) {
			log.Warnf("Fetching mastership state for E2Node '%s' failed: %v", channel.E2NodeID, err)
			return false, err
		}
		log.Warnf("Mastership state not found for E2Node '%s' %v", channel.E2NodeID, err)
		return false, nil
	}

	// Decode the mastership state
	mastership := topoapi.MastershipState{}
	_ = e2NodeEntity.GetAspect(&mastership)

	// Return if the channel status is already up-to-date
	if channel.Status.Term == e2api.TermID(mastership.Term) {
		return false, nil
	}

	// Get the master relation
	e2NodeRelation, err := r.topo.Get(ctx, topoapi.ID(mastership.NodeId))
	if err != nil {
		if !errors.IsNotFound(err) {
			log.Warnf("Fetching mastership state for E2Node '%s' failed: %v", channel.E2NodeID, err)
			return false, err
		}
		log.Warnf("Master relation not found for E2Node '%s'", channel.E2NodeID)
		return false, nil
	}

	// Update the channel status with the new term/master
	currentTime := time.Now()
	channel.Status.Timestamp = &currentTime
	channel.Status.Term = e2api.TermID(mastership.Term)
	channel.Status.Master = e2api.MasterID(e2NodeRelation.GetRelation().SrcEntityID)
	if err := r.chans.Update(ctx, channel); err != nil && !errors.IsNotFound(err) && !errors.IsConflict(err) {
		log.Warnf("Failed to update channel %s: %s", channel.ID, err)
		return false, err
	}
	return true, nil
}

func (r *Reconciler) finalizeChannel(ctx context.Context, channel *e2api.Channel) (bool, error) {
	// If this node is not the master for the channel, clean up the streams
	nodeID := e2api.E2TInstanceID(utils.GetE2TID())
	if utils.ContainsString(channel.Finalizers, string(nodeID)) &&
		(channel.Status.Master != e2api.MasterID(nodeID) ||
			(channel.Status.Phase == e2api.ChannelPhase_CHANNEL_CLOSED &&
				channel.Status.State == e2api.ChannelState_CHANNEL_COMPLETE)) {
		log.Infof("New master elected for channel '%s': closing channel stream", channel.ID)
		var (
			sub  *broker.SubscriptionStream
			app  *broker.AppStream
			tx   *broker.TransactionStream
			inst *broker.AppInstanceStream
		)
		sub, ok := r.streams.Subscriptions().Get(channel.SubscriptionID)
		if !ok {
			goto finalize
		}
		app, ok = sub.Apps().Get(channel.AppID)
		if !ok {
			goto finalize
		}
		tx, ok = app.Transactions().Get(channel.TransactionID)
		if !ok {
			goto finalize
		}
		inst, ok = tx.Instances().Get(channel.AppInstanceID)
		if !ok {
			goto finalize
		}
		inst.Close()
	finalize:
		channel.Finalizers = utils.RemoveString(channel.Finalizers, string(nodeID))
		if err := r.chans.Update(ctx, channel); err != nil && !errors.IsNotFound(err) && !errors.IsConflict(err) {
			log.Warnf("Failed to reconcile Channel %+v: %s", channel, err)
			return false, err
		}
		return true, nil
	}

	for _, nodeID := range channel.Finalizers {
		if _, err := r.topo.Get(ctx, topoapi.ID(nodeID)); errors.IsNotFound(err) {
			channel.Finalizers = utils.RemoveString(channel.Finalizers, nodeID)
			if err := r.chans.Update(ctx, channel); err != nil && !errors.IsNotFound(err) && !errors.IsConflict(err) {
				log.Warnf("Failed to reconcile Channel %+v: %s", channel, err)
				return false, err
			}
			return true, nil
		}
	}
	return false, nil
}
