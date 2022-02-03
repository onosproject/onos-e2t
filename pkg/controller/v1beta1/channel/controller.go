// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package channel

import (
	"context"
	topoapi "github.com/onosproject/onos-api/go/onos/topo"
	"github.com/onosproject/onos-e2t/pkg/controller/utils"
	"github.com/onosproject/onos-e2t/pkg/northbound/e2/stream"
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
func NewController(chans chanstore.Store, subs substore.Store, streams stream.Manager, topo rnib.Store) *controller.Controller {
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
	streams stream.Manager
	topo    rnib.Store
}

// Reconcile reconciles the state of a channel
func (r *Reconciler) Reconcile(id controller.ID) (controller.Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	channelID := id.Value.(e2api.ChannelID)
	log.Infof("Reconciling Channel '%s'", channelID)
	channel, err := r.chans.Get(ctx, channelID)
	if err != nil {
		if errors.IsNotFound(err) {
			log.Debugf("Channel '%s' not found", channelID)
			return controller.Result{}, nil
		}
		log.Errorf("Failed to reconcile Channel '%s'", channelID, err)
		return controller.Result{}, err
	}

	log.Debug(channel)

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
			log.Errorf("Failed to reconcile Channel '%s'", channel.ID, err)
			return controller.Result{}, err
		}

		log.Infof("Creating Channel '%s' Subscription '%s'", channel.ID, channel.SubscriptionID)
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
		log.Debug(sub)
		err := r.subs.Create(ctx, sub)
		if err != nil && !errors.IsAlreadyExists(err) {
			log.Errorf("Error creating Channel '%s' Subscription '%s'", channel.ID, sub.ID, err)
			return controller.Result{}, err
		}
		return controller.Result{}, nil
	}

	// If the subscription is being closed, wait for it to be deleted and recreated by this controller
	if sub.Status.Phase != e2api.SubscriptionPhase_SUBSCRIPTION_OPEN {
		log.Infof("Skipping reconciliation for Channel '%s': Subscription '%s' is being closed", channel.ID, sub.ID)
		return controller.Result{}, nil
	}

	// If the subscription is OPEN, add the channel to the subscription if necessary
	channels := make(map[e2api.ChannelID]bool)
	for _, id := range sub.Status.Channels {
		channels[id] = true
	}
	if _, ok := channels[channel.ID]; !ok {
		log.Infof("Binding Channel '%s' to existing Subscription '%s'", channel.ID, sub.ID)
		sub.Status.Channels = append(sub.Status.Channels, channel.ID)
		log.Debug(sub)
		if err := r.subs.Update(ctx, sub); err != nil && !errors.IsNotFound(err) && !errors.IsConflict(err) {
			log.Errorf("Error binding Channel '%s' to existing Subscription '%s'", channel.ID, sub.ID, err)
			return controller.Result{}, err
		}
		return controller.Result{}, nil
	}

	// If the subscription term has changed, close the channel stream and update the channel term and master
	if channel.Status.Term < sub.Status.Term {
		log.Debugf("Fetching master relation for Subscription '%s'", sub.ID)
		e2NodeMasterRelation, err := r.topo.Get(ctx, topoapi.ID(sub.Status.Master))
		if err != nil {
			if errors.IsNotFound(err) {
				log.Warnf("Master relation not found for Subscription '%s'", sub.ID)
				return controller.Result{}, nil
			}
			log.Errorf("Error fetching master relation for Subscription '%s'", sub.ID, err)
			return controller.Result{}, err
		}

		log.Infof("Updating Channel '%s' mastership to term %d", channel.ID, sub.Status.Term)
		e2NodeMasterID := e2api.MasterID(e2NodeMasterRelation.GetRelation().SrcEntityID)
		channel.Status.Term = sub.Status.Term
		channel.Status.Master = e2NodeMasterID
		channel.Status.State = e2api.ChannelState_CHANNEL_PENDING
		channel.Status.Error = nil
		now := time.Now()
		channel.Status.Timestamp = &now
		log.Debug(channel)
		if err := r.chans.Update(ctx, channel); err != nil && !errors.IsNotFound(err) && !errors.IsConflict(err) {
			log.Errorf("Error updating mastership for Channel '%s'", channel.ID, err)
			return controller.Result{}, err
		}
		return controller.Result{}, nil
	}

	// If the channel term is not set, skip reconciliation to wait for mastership election
	if channel.Status.Term == 0 {
		log.Debugf("Master not set for Channel '%s'", channel.ID)
		return controller.Result{}, nil
	}

	// If this is the master for the channel term, update the channel/stream state
	localInstanceID := e2api.E2TInstanceID(utils.GetE2TID())
	masterInstanceID := e2api.E2TInstanceID(channel.Status.Master)
	if localInstanceID == masterInstanceID {
		// If the channel expiration timestamp is set, check if it needs to be unset
		// If the channel expiration timestamp is not set, check if it needs to be set
		if channel.Status.Timestamp != nil {
			if stream, ok := r.streams.Get(channel.ID); ok && len(stream.Output().Streams()) > 0 {
				log.Infof("Unsetting disconnect timestamp for Channel '%s'", channel.ID)
				channel.Status.Timestamp = nil
				log.Debug(channel)
				if err := r.chans.Update(ctx, channel); err != nil && !errors.IsNotFound(err) && !errors.IsConflict(err) {
					log.Errorf("Error setting disconnect timestamp for Channel '%s'", channel.ID, err)
					return controller.Result{}, err
				}
				return controller.Result{}, nil
			}
		} else {
			if stream, ok := r.streams.Get(channel.ID); !ok || len(stream.Output().Streams()) == 0 {
				log.Infof("Setting disconnect timestamp for Channel '%s'", channel.ID)
				now := time.Now()
				channel.Status.Timestamp = &now
				log.Debug(channel)
				if err := r.chans.Update(ctx, channel); err != nil && !errors.IsNotFound(err) && !errors.IsConflict(err) {
					log.Errorf("Error setting disconnect timestamp for Channel '%s'", channel.ID, err)
					return controller.Result{}, err
				}
				return controller.Result{}, nil
			}
		}

		// Reconcile the channel based on its state
		switch channel.Status.State {
		case e2api.ChannelState_CHANNEL_PENDING:
			// If the channel is pending on this node, ensure the channel stream is open
			r.streams.Open(channel.ID, channel.ChannelMeta)

			// If the subscription is in a finished state, update the channel state
			switch sub.Status.State {
			case e2api.SubscriptionState_SUBSCRIPTION_COMPLETE:
				log.Debugf("Completing Channel '%s': Subscription complete", channel.ID)
				channel.Status.State = e2api.ChannelState_CHANNEL_COMPLETE
				log.Debug(channel)
				if err := r.chans.Update(ctx, channel); err != nil && !errors.IsNotFound(err) && !errors.IsConflict(err) {
					log.Warnf("Failed to reconcile Channel '%s'", channel.ID, err)
					return controller.Result{}, err
				}
				return controller.Result{}, nil
			case e2api.SubscriptionState_SUBSCRIPTION_FAILED:
				log.Debugf("Failing Channel '%s': Subscription failed", channel.ID)
				channel.Status.State = e2api.ChannelState_CHANNEL_FAILED
				channel.Status.Error = sub.Status.Error
				log.Debug(channel)
				if err := r.chans.Update(ctx, channel); err != nil && !errors.IsNotFound(err) && !errors.IsConflict(err) {
					log.Warnf("Failed to reconcile Channel '%s'", channel.ID, err)
					return controller.Result{}, err
				}
				return controller.Result{}, nil
			}
		case e2api.ChannelState_CHANNEL_COMPLETE:
			// If the channel state is COMPLETE, acknowledge the channel stream
			if stream, ok := r.streams.Get(channel.ID); ok {
				log.Infof("Acknowledging Channel '%s' stream", channel.ID)
				stream.Input().Open()
			}
		case e2api.ChannelState_CHANNEL_FAILED:
			// If the channel state is FAILED, fail the channel stream
			if stream, ok := r.streams.Get(channel.ID); ok {
				log.Infof("Failing Channel '%s' stream", channel.ID)
				errStat := status.New(codes.Aborted, "an E2AP failure occurred")
				errStat, err := errStat.WithDetails(channel.Status.Error)
				if err != nil {
					log.Errorf("Error failing Channel '%s' stream", channel.ID, err)
					return controller.Result{}, nil
				}
				stream.Input().Close(errStat.Err())
			}
		}
	} else {
		log.Warnf("Not the master for Channel '%s'", channel.ID)
		if stream, ok := r.streams.Get(channel.ID); ok {
			log.Infof("Closing Channel '%s' stream: mastership changed", channel.ID)
			stream.Input().Close(errors.NewUnavailable("mastership term changed"))
			return controller.Result{}, nil
		}
	}

	// If the channel expiration timestamp is set, determine whether the channel has expired
	if channel.Status.Timestamp != nil {
		transactionTimeout := defaultTransactionTimeout
		if channel.Spec.TransactionTimeout != nil {
			transactionTimeout = *channel.Spec.TransactionTimeout
		}

		expireTime := channel.Status.Timestamp.Add(transactionTimeout)
		if time.Now().After(expireTime) {
			log.Infof("Closing Channel '%s': transaction timed out", channel.ID)
			channel.Status.Phase = e2api.ChannelPhase_CHANNEL_CLOSED
			channel.Status.State = e2api.ChannelState_CHANNEL_PENDING
			channel.Status.Error = nil
			log.Debug(channel)
			if err := r.chans.Update(ctx, channel); err != nil && !errors.IsNotFound(err) && !errors.IsConflict(err) {
				log.Errorf("Error closing Channel '%s'", channel.ID, err)
				return controller.Result{}, err
			}
			return controller.Result{}, nil
		}

		log.Debugf("Rescheduling reconciliation for Channel '%s' at %v", channel.ID, expireTime)
		return controller.Result{
			RequeueAt: expireTime,
		}, nil
	}
	return controller.Result{}, nil
}

func (r *Reconciler) reconcileClosedChannel(ctx context.Context, channel *e2api.Channel) (controller.Result, error) {
	// If the close has completed, delete the channel
	if channel.Status.State == e2api.ChannelState_CHANNEL_COMPLETE {
		log.Infof("Deleting closed Channel '%s'", channel.ID)
		err := r.chans.Delete(ctx, channel)
		if err != nil && !errors.IsNotFound(err) {
			log.Errorf("Error deleting closed Channel '%s'", channel.ID, err)
			return controller.Result{}, err
		}
		return controller.Result{}, nil
	}

	// If this node is not the current master for the channel, skip the remainder of reconciliation
	localInstanceID := e2api.E2TInstanceID(utils.GetE2TID())
	masterInstanceID := e2api.E2TInstanceID(channel.Status.Master)
	if channel.Status.Term > 0 && localInstanceID != masterInstanceID {
		log.Debugf("Fetching master entity for Channel '%s'", channel.ID)
		_, err := r.topo.Get(ctx, topoapi.ID(masterInstanceID))
		if err == nil {
			log.Warnf("Not the master for Channel '%s'", channel.ID)
			return controller.Result{}, nil
		} else if errors.IsNotFound(err) {
			log.Warnf("Master entity not found for Channel '%s'", channel.ID)
		} else {
			log.Errorf("Error fetching master entity for Channel '%s'", channel.ID, err)
			return controller.Result{}, err
		}
	}

	// Get the underlying Subscription
	sub, err := r.subs.Get(ctx, channel.SubscriptionID)
	if err != nil {
		if !errors.IsNotFound(err) {
			log.Errorf("Error reconciling closed Channel '%s'", channel.ID, err)
			return controller.Result{}, err
		}

		// If the subscription is not found, complete the channel close
		log.Infof("Completing closed Channel '%s': subscription not found", channel.ID)
		if stream, ok := r.streams.Get(channel.ID); ok {
			log.Infof("Closing Channel '%s' stream", channel.ID)
			stream.Input().Close(nil)
		}
		channel.Status.State = e2api.ChannelState_CHANNEL_COMPLETE
		log.Debug(channel)
		if err := r.chans.Update(ctx, channel); err != nil && !errors.IsNotFound(err) && !errors.IsConflict(err) {
			log.Errorf("Error completing closed Channel '%s'", channel.ID, err)
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
		log.Infof("Unbinding closed Channel '%s' from Subscription '%s'", channel.ID, sub.ID)
		sub.Status.Channels = channels
		log.Debug(sub)
		if err := r.subs.Update(ctx, sub); err != nil && !errors.IsNotFound(err) && !errors.IsConflict(err) {
			log.Errorf("Failed to reconcile Channel %+v: %s", channel.ID, err)
			return controller.Result{}, err
		}
		return controller.Result{}, nil
	}

	// If the subscription is OPEN or it's CLOSED with a finished state, update the channel state
	if len(sub.Status.Channels) > 0 ||
		(sub.Status.Phase == e2api.SubscriptionPhase_SUBSCRIPTION_CLOSED &&
			sub.Status.State == e2api.SubscriptionState_SUBSCRIPTION_COMPLETE) {
		if stream, ok := r.streams.Get(channel.ID); ok {
			log.Infof("Closing closed Channel '%s' stream", channel.ID)
			stream.Input().Close(nil)
		}

		// Complete the closed channel
		log.Infof("Completing closed Channel '%s'", channel.ID)
		channel.Status.State = e2api.ChannelState_CHANNEL_COMPLETE
		log.Debug(channel)
		if err := r.chans.Update(ctx, channel); err != nil && !errors.IsNotFound(err) && !errors.IsConflict(err) {
			log.Errorf("Error completing closed Channel '%s'", channel.ID, err)
			return controller.Result{}, err
		}
		return controller.Result{}, nil
	}
	return controller.Result{}, nil
}
