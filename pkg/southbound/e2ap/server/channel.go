// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package server

import (
	"context"
	"sync"

	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2apies"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2appducontents"
	"github.com/onosproject/onos-e2t/pkg/modelregistry"
	"github.com/onosproject/onos-e2t/pkg/protocols/e2"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
)

func newE2Channel(id ChannelID, plmdID string, channel e2.ServerChannel, modelFuncIDs map[modelregistry.ModelFullName]types.RanFunctionID) *E2Channel {
	return &E2Channel{
		ServerChannel: channel,
		ID:            id,
		PlmnID:        plmdID,
		modelFuncIDs:  modelFuncIDs,
		watchers:      make(map[int32]chan<- e2appducontents.Ricindication),
	}
}

type E2Channel struct {
	e2.ServerChannel
	ID           ChannelID
	PlmnID       string
	modelFuncIDs map[modelregistry.ModelFullName]types.RanFunctionID
	watchers     map[int32]chan<- e2appducontents.Ricindication
	watchersMu   sync.RWMutex
}

func (c *E2Channel) GetRANFunctionID(modelName modelregistry.ModelFullName) types.RanFunctionID {
	return c.modelFuncIDs[modelName]
}

func (c *E2Channel) ricIndication(ctx context.Context, request *e2appducontents.Ricindication) error {
	c.watchersMu.RLock()
	watcher, ok := c.watchers[request.ProtocolIes.E2ApProtocolIes29.Value.RicRequestorId]
	c.watchersMu.RUnlock()
	if ok {
		watcher <- *request
	}
	return nil
}

func (c *E2Channel) WatchRICIndications(ctx context.Context, requestID e2apies.RicrequestId, ch chan<- e2appducontents.Ricindication) {
	watchCh := make(chan e2appducontents.Ricindication)
	c.watchersMu.Lock()
	c.watchers[requestID.RicRequestorId] = watchCh
	c.watchersMu.Unlock()
	go func() {
		defer close(ch)
		for {
			select {
			case indication, ok := <-watchCh:
				if !ok {
					return
				}
				ch <- indication
			case <-ctx.Done():
				c.watchersMu.Lock()
				if _, ok := c.watchers[requestID.RicRequestorId]; ok {
					delete(c.watchers, requestID.RicRequestorId)
					close(watchCh)
				}
				c.watchersMu.Unlock()
			}
		}
	}()
}
