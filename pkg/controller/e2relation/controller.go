// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package e2relation

import (
	"context"
	"time"

	topoapi "github.com/onosproject/onos-api/go/onos/topo"
	"github.com/onosproject/onos-e2t/pkg/controller/utils"
	e2server "github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/server"
	"github.com/onosproject/onos-e2t/pkg/store/rnib"
	"github.com/onosproject/onos-lib-go/pkg/controller"
	"github.com/onosproject/onos-lib-go/pkg/errors"
	"github.com/onosproject/onos-lib-go/pkg/logging"
)

const (
	defaultTimeout = 30 * time.Second
)

var log = logging.GetLogger("controller", "e2relation")

// NewController returns a new E2 control relation controller
func NewController(rnib rnib.Store, channels e2server.ChannelManager) *controller.Controller {
	c := controller.NewController("E2ControlRelation")
	c.Watch(&ChannelWatcher{
		channels: channels,
	})

	c.Reconcile(&Reconciler{
		channels: channels,
		rnib:     rnib,
	})

	return c
}

// Reconciler is an E2T/E2 node relations reconciler
type Reconciler struct {
	channels e2server.ChannelManager
	rnib     rnib.Store
}

func (r *Reconciler) createE2ControlRelation(ctx context.Context, channel *e2server.E2Channel) error {
	relationID := utils.GetE2ControlRelationID(channel.ID)
	_, err := r.rnib.Get(ctx, relationID)
	if err == nil {
		return nil
	} else if !errors.IsNotFound(err) {
		log.Warnf("Creating E2Node '%s' control relation '%s' failed: %v", channel.E2NodeID, relationID, err)
		return err
	}

	log.Debugf("Creating E2Node '%s' control relation '%s'", channel.E2NodeID, relationID)
	object := &topoapi.Object{
		ID:   relationID,
		Type: topoapi.Object_RELATION,
		Obj: &topoapi.Object_Relation{
			Relation: &topoapi.Relation{
				KindID:      topoapi.CONTROLS,
				SrcEntityID: utils.GetE2TID(),
				TgtEntityID: channel.E2NodeID,
			},
		},
	}

	err = r.rnib.Create(ctx, object)
	if err != nil {
		if !errors.IsAlreadyExists(err) {
			log.Warnf("Creating E2Node '%s' control relation '%s' failed: %v", channel.E2NodeID, relationID, err)
			return err
		}
		return nil
	}
	return nil
}

func (r *Reconciler) Reconcile(id controller.ID) (controller.Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	channelID := id.Value.(e2server.ChannelID)
	log.Infof("Reconciling E2 node Control relation for channel: %s", channelID)
	channel, err := r.channels.Get(ctx, channelID)
	if err != nil {
		if errors.IsNotFound(err) {
			return r.reconcileDeleteE2ControlRelation(channelID)
		}
		log.Warnf("Failed to reconcile E2 node control relation for channel %s: %s", channelID, err)
		return controller.Result{}, err
	}
	return r.reconcileE2ControlRelation(channel)
}

func (r *Reconciler) reconcileE2ControlRelation(channel *e2server.E2Channel) (controller.Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	if err := r.createE2ControlRelation(ctx, channel); err != nil {
		return controller.Result{}, err
	}
	return controller.Result{}, nil
}

func (r *Reconciler) deleteE2ControlRelation(ctx context.Context, channelID e2server.ChannelID) error {
	relationID := utils.GetE2ControlRelationID(channelID)
	log.Debugf("Deleting E2Node relation '%s' for Channel '%s'", relationID, channelID)
	err := r.rnib.Delete(ctx, relationID)
	if err != nil {
		if !errors.IsNotFound(err) {
			log.Warnf("Deleting E2Node relation '%s' for Channel '%s' failed: %v", relationID, channelID, err)
			return err
		}
		return nil
	}
	return nil
}

func (r *Reconciler) reconcileDeleteE2ControlRelation(channelID e2server.ChannelID) (controller.Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	if err := r.deleteE2ControlRelation(ctx, channelID); err != nil {
		return controller.Result{}, err
	}
	return controller.Result{}, nil
}
