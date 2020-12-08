// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package channel

import (
	"context"
	"fmt"
	"github.com/onosproject/onos-e2t/pkg/modelregistry"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/asn1cgo"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/pdubuilder"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/pdudecoder"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
	"github.com/onosproject/onos-lib-go/pkg/errors"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"io"
	"net"
	"sync"
)

var log = logging.GetLogger("southbound", "e2", "channel")

// TODO: Change the RIC ID to something appropriate
var ricID = types.RicIdentifier{
	RicIdentifierValue: 0xABCDE,
	RicIdentifierLen:   20,
}

// NewManager creates a new channel manager
func NewManager(modelregistry *modelregistry.ModelRegistry) *Manager {
	mgr := &Manager{
		channels:      make(map[ID]Channel),
		eventCh:       make(chan Channel),
		modelRegistry: modelregistry,
	}
	go mgr.processEvents()
	return mgr
}

// Manager is a stream manager
type Manager struct {
	channels      map[ID]Channel
	channelsMu    sync.RWMutex
	watchers      []chan<- Channel
	watchersMu    sync.RWMutex
	eventCh       chan Channel
	modelRegistry *modelregistry.ModelRegistry
}

func (m *Manager) processEvents() {
	for channel := range m.eventCh {
		m.processEvent(channel)
	}
}

func (m *Manager) processEvent(channel Channel) {
	log.Infof("Notifying channel %v+", channel.ID)
	m.watchersMu.RLock()
	for _, watcher := range m.watchers {
		watcher <- channel
	}
	m.watchersMu.RUnlock()
}

// Open opens a new channel
func (m *Manager) Open(ctx context.Context, conn net.Conn) (Channel, error) {
	log.Infof("Opening channel %s", conn.RemoteAddr())
	c, err := m.setup(ctx, conn)
	if err != nil {
		return nil, err
	}

	m.channelsMu.Lock()
	defer m.channelsMu.Unlock()
	m.channels[c.ID()] = c
	m.eventCh <- c
	go func() {
		<-ctx.Done()
		log.Infof("Closing channel %s", c.ID())
		m.channelsMu.Lock()
		delete(m.channels, c.ID())
		m.channelsMu.Unlock()
	}()
	return c, nil
}

// setup sets up a channel
func (m *Manager) setup(ctx context.Context, conn net.Conn) (Channel, error) {
	buf := make([]byte, readBufSize)
	n, err := conn.Read(buf)
	if err != nil {
		return nil, err
	}

	// Decode the E2 request in PER encoding
	e2PDUReqBytes := buf[:n]
	e2PDUReq, err := asn1cgo.PerDecodeE2apPdu(e2PDUReqBytes)
	if err != nil {
		return nil, err
	}

	nodeID, ranFuncs, err := pdudecoder.DecodeE2SetupRequestPdu(e2PDUReq)
	if err != nil {
		return nil, err
	}
	channelID := ID(fmt.Sprintf("%s:%d", string(nodeID.NodeIdentifier), nodeID.NodeType))
	plmnID := PlmnID([]byte{nodeID.Plmn[0], nodeID.Plmn[1], nodeID.Plmn[2]})

	serviceModelName := modelregistry.ModelFullName("e2sm_kpm-v1beta1") // TODO: Remove hardcoded name
	serviceModel, ok := m.modelRegistry.ModelPlugins[serviceModelName]
	if !ok {
		log.Warnf("No Service Model found for %s", serviceModelName)
	}
	rfAccepted := make(types.RanFunctionRevisions)
	rfRejected := make(types.RanFunctionCauses)
	for id, ranFunc := range *ranFuncs {
		rfAccepted[id] = ranFunc.Revision
		if serviceModel != nil {
			names, triggers, reports, err := serviceModel.DecodeRanFunctionDescription(ranFunc.Description)
			if err != nil {
				return nil, errors.NewInvalid("Error decoding RanFunctionDescription in E2SetupRequest %s", err.Error())
			}
			log.Infof("RanFunctionDescription ShortName: %s, Desc: %s,"+
				"Instance: %d, Oid: %s. #Triggers: %d. #Reports: %d",
				names.RanFunctionShortName,
				names.RanFunctionDescription,
				names.RanFunctionInstance,
				names.RanFunctionE2SmOid,
				len(*triggers), len(*reports))
		}
	}

	// Create an E2 setup response
	e2PDUResp, err := pdubuilder.CreateResponseE2apPdu(nodeID.Plmn, ricID, rfAccepted, rfRejected)
	if err != nil {
		return nil, err
	}

	if err := e2PDUResp.Validate(); err != nil {
		return nil, err
	}

	// Encode the setup response in PER
	e2PDURespBytes, err := asn1cgo.PerEncodeE2apPdu(e2PDUResp)
	if err != nil {
		return nil, err
	}

	_, err = conn.Write(e2PDURespBytes)
	if err != nil {
		return nil, err
	}

	meta := Metadata{
		ID:     channelID,
		PlmnID: plmnID,
	}
	return newChannel(ctx, conn, meta), nil
}

// Get gets a channel by ID
func (m *Manager) Get(ctx context.Context, id ID) (Channel, error) {
	m.channelsMu.RLock()
	defer m.channelsMu.RUnlock()
	channel, ok := m.channels[id]
	if !ok {
		return nil, errors.NewNotFound("channel '%s' not found", id)
	}
	return channel, nil
}

// List lists channels
func (m *Manager) List(ctx context.Context) ([]Channel, error) {
	m.channelsMu.RLock()
	defer m.channelsMu.RUnlock()
	channels := make([]Channel, 0, len(m.channels))
	for _, channel := range m.channels {
		channels = append(channels, channel)
	}
	return channels, nil
}

// Watch watches for new channels
func (m *Manager) Watch(ctx context.Context, ch chan<- Channel) error {
	m.watchersMu.Lock()
	m.channelsMu.Lock()
	m.watchers = append(m.watchers, ch)
	m.watchersMu.Unlock()

	go func() {
		for _, stream := range m.channels {
			ch <- stream
		}
		m.channelsMu.Unlock()

		<-ctx.Done()
		m.watchersMu.Lock()
		watchers := make([]chan<- Channel, 0, len(m.watchers)-1)
		for _, watcher := range watchers {
			if watcher != ch {
				watchers = append(watchers, watcher)
			}
		}
		m.watchers = watchers
		m.watchersMu.Unlock()
	}()
	return nil
}

// Close closes the manager
func (m *Manager) Close() error {
	close(m.eventCh)
	return nil
}

var _ io.Closer = &Manager{}
