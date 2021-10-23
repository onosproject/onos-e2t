// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package channel

import (
	"context"
	topoapi "github.com/onosproject/onos-api/go/onos/topo"
	"github.com/onosproject/onos-e2t/pkg/controller/utils"
	"github.com/onosproject/onos-e2t/pkg/northbound/e2/channel"
	"github.com/onosproject/onos-e2t/pkg/store/rnib"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
func NewController(chans chanstore.Store, subs substore.Store, streams channel.Manager, topo rnib.Store) *controller.Controller {
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
	c.Watch(&StreamWatcher{
		streams: streams,
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
	streams channel.Manager
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
	// Get the subscription or create one if it doesn't already exist
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

	// If the subscription is OPEN, add the channel to the subscription if necessary
	if sub.Status.Phase == e2api.SubscriptionPhase_SUBSCRIPTION_OPEN {
		channels := make(map[e2api.ChannelID]bool)
		for _, id := range sub.Status.Channels {
			channels[id] = true
		}
		if _, ok := channels[channel.ID]; !ok {
			log.Debugf("Binding Channel %+v to existing Subscription %+v", channel, sub)
			sub.Status.Channels = append(sub.Status.Channels, channel.ID)
			if err := r.subs.Update(ctx, sub); err != nil && !errors.IsNotFound(err) && !errors.IsConflict(err) {
				log.Warnf("Failed to reconcile Channel %+v: %s", channel, err)
				return controller.Result{}, err
			}
			return controller.Result{}, nil
		}
	}

	// Update the channel expiration timestamp if necessary
	setTimestamp := func(timestamp *time.Time) error {
		log.Infof("Setting disconnect timestamp for Channel '%s'", channel.ID)
		channel.Status.Timestamp = timestamp
		log.Debug(channel)
		if err := r.chans.Update(ctx, channel); err != nil && !errors.IsNotFound(err) && !errors.IsConflict(err) {
			return err
		}
		return nil
	}

	now := time.Now()

	// If the channel stream does not exist, set the expiration timestamp
	// If the channel stream does exist, unset the expiration timestamp
	_, ok := r.streams.Get(channel.ID)
	if !ok && channel.Status.Timestamp == nil {
		if err := setTimestamp(&now); err != nil {
			log.Errorf("Error setting disconnect timestamp for Channel '%s'", channel.ID, err)
			return controller.Result{}, err
		}
		return controller.Result{}, nil
	} else if ok && channel.Status.Timestamp != nil {
		if err := setTimestamp(nil); err != nil {
			log.Errorf("Error setting disconnect timestamp for Channel '%s'", channel.ID, err)
			return controller.Result{}, err
		}
		return controller.Result{}, nil
	}

	// If the channel expiration timestamp is set, determine whether the channel has expired
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

	localInstanceID := e2api.E2TInstanceID(utils.GetE2TID())

	// If the subscription term has changed, close the channel stream and update the channel term and master
	if channel.Status.Term < sub.Status.Term {
		// Close the local channel stream with an Unavailable error to force the client to try a new master
		if stream, ok := r.streams.Get(channel.ID); ok {
			stream.Writer().Close(errors.NewUnavailable("mastership term changed"))
		}

		log.Debugf("Fetching mastership state for E2Node '%s'", sub.E2NodeID)
		e2NodeMasterRelation, err := r.topo.Get(ctx, topoapi.ID(sub.Status.Master))
		if err != nil {
			if errors.IsNotFound(err) {
				log.Warnf("Master relation not found for E2Node '%s'", sub.E2NodeID)
				return controller.Result{}, nil
			}
			log.Errorf("Error fetching mastership state for E2Node '%s'", sub.E2NodeID, err)
			return controller.Result{}, err
		}

		log.Infof("Updating Channel %s mastership to term %d", channel.ID, sub.Status.Term)
		e2NodeMasterID := e2api.MasterID(e2NodeMasterRelation.GetRelation().SrcEntityID)
		channel.Status.Term = sub.Status.Term
		channel.Status.Master = e2NodeMasterID
		channel.Status.State = e2api.ChannelState_CHANNEL_PENDING
		channel.Status.Error = nil
		log.Debug(channel)
		if err := r.chans.Update(ctx, channel); err != nil && !errors.IsNotFound(err) && !errors.IsConflict(err) {
			log.Errorf("Error updating mastership for Channel '%s'", channel.ID, err)
			return controller.Result{}, err
		}
		return controller.Result{}, nil
	}

	if e2api.E2TInstanceID(channel.Status.Master) != localInstanceID {
		log.Warnf("Not the master for E2Node '%s'", channel.E2NodeID)
		return controller.Result{}, nil
	}

	switch channel.Status.State {
	case e2api.ChannelState_CHANNEL_PENDING:
		// If the subscription is in a finished state, update the channel state
		switch sub.Status.State {
		case e2api.SubscriptionState_SUBSCRIPTION_COMPLETE:
			log.Debugf("Completing Channel %+v: Subscription complete", channel)
			channel.Status.State = e2api.ChannelState_CHANNEL_COMPLETE
			if err := r.chans.Update(ctx, channel); err != nil && !errors.IsNotFound(err) && !errors.IsConflict(err) {
				log.Warnf("Failed to reconcile Channel %+v: %s", channel, err)
				return controller.Result{}, err
			}
			return controller.Result{}, nil
		case e2api.SubscriptionState_SUBSCRIPTION_FAILED:
			log.Debugf("Failing Channel %+v: Subscription failed", channel)
			channel.Status.State = e2api.ChannelState_CHANNEL_FAILED
			channel.Status.Error = sub.Status.Error
			if err := r.chans.Update(ctx, channel); err != nil && !errors.IsNotFound(err) && !errors.IsConflict(err) {
				log.Warnf("Failed to reconcile Channel %+v: %s", channel, err)
				return controller.Result{}, err
			}
			return controller.Result{}, nil
		}
	case e2api.ChannelState_CHANNEL_COMPLETE:
		r.streams.Open(channel).Writer().Ack()
		return controller.Result{}, nil
	case e2api.ChannelState_CHANNEL_FAILED:
		stream, ok := r.streams.Get(channel.ID)
		if ok {
			errStat := status.New(codes.Aborted, "an E2AP failure occurred")
			errStat, err := errStat.WithDetails(channel.Status.Error)
			if err != nil {
				log.Errorf("Error failing Channel '%s'", channel.ID, err)
				return controller.Result{}, nil
			}
			stream.Writer().Fail(errStat.Err())
		}
		return controller.Result{}, nil
	}
	return controller.Result{}, nil
}

func (r *Reconciler) reconcileClosedChannel(ctx context.Context, channel *e2api.Channel) (controller.Result, error) {
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

		// Close the local channel stream gracefully
		if stream, ok := r.streams.Get(channel.ID); ok {
			stream.Writer().Close(nil)
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
