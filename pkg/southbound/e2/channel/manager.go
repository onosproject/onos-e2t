// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package channel

import (
	"context"
	"errors"
	"fmt"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2ap-commondatatypes"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2apies"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2appducontents"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2appdudescriptions"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/asn1cgo"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"io"
	"net"
	"sync"
)

var log = logging.GetLogger("southbound", "e2", "channel")

// TODO: Change the RIC ID to something appropriate
const ricID = 0x01 // µONOS RIC is #1!!!

// NewManager creates a new channel manager
func NewManager() *Manager {
	mgr := &Manager{
		channels: make(map[ID]Channel),
		eventCh:  make(chan Channel),
	}
	go mgr.processEvents()
	return mgr
}

// Manager is a stream manager
type Manager struct {
	channels   map[ID]Channel
	channelsMu sync.RWMutex
	watchers   []chan<- Channel
	watchersMu sync.RWMutex
	eventCh    chan Channel
}

func (m *Manager) processEvents() {
	for channel := range m.eventCh {
		m.processEvent(channel)
	}
}

func (m *Manager) processEvent(channel Channel) {
	log.Infof("Notifying channel %s", channel.ID)
	m.watchersMu.RLock()
	for _, watcher := range m.watchers {
		watcher <- channel
	}
	m.watchersMu.RUnlock()
}

// Open opens a new channel
func (m *Manager) Open(ctx context.Context, conn net.Conn) (Channel, error) {
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

	// Verify this is a setup request
	e2SetupReq := e2PDUReq.GetInitiatingMessage().GetProcedureCode().GetE2Setup()
	if e2SetupReq == nil {
		return nil, errors.New("unexpected message type")
	}

	// Extract the RAN function list
	ranFunctions := make(map[RANFunctionID]RANFunctionMetadata)
	if e2SetupReq.InitiatingMessage.ProtocolIes.E2ApProtocolIes10 != nil {
		for _, ranFunction := range e2SetupReq.InitiatingMessage.ProtocolIes.E2ApProtocolIes10.Value.Value {
			ranFunctions[RANFunctionID(ranFunction.E2ApProtocolIes10.Value.RanFunctionId.Value)] = RANFunctionMetadata{
				Description: RANFunctionDescription(ranFunction.E2ApProtocolIes10.Value.RanFunctionDefinition.Value),
				Revision:    RANFunctionRevision(ranFunction.E2ApProtocolIes10.Value.RanFunctionRevision.Value),
			}
		}
	}

	// Verify an E2 node ID is provided
	e2NodeID := e2SetupReq.InitiatingMessage.ProtocolIes.E2ApProtocolIes3.Value.GlobalE2NodeId
	globalE2NodeID, ok := e2NodeID.(*e2apies.GlobalE2NodeId_GNb)
	if !ok {
		return nil, errors.New("unexpected message format")
	}

	// Verify a gNB ID is provided
	gnbID, ok := globalE2NodeID.GNb.GlobalGNbId.GnbId.GnbIdChoice.(*e2apies.GnbIdChoice_GnbId)
	if !ok {
		return nil, errors.New("unexpected message format")
	}

	// Create a channel ID from the gNB ID and plmn ID
	channelID := ID(fmt.Sprintf("%s:%d", string(globalE2NodeID.GNb.GlobalGNbId.PlmnId.Value), gnbID.GnbId.Value))
	plmnID := PlmnID(globalE2NodeID.GNb.GlobalGNbId.PlmnId.Value)

	// Create an E2 setup response
	e2SetupResp := e2appducontents.E2SetupResponseIes_E2SetupResponseIes4{
		Id:          int32(v1beta1.ProtocolIeIDGlobalRicID),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &e2apies.GlobalRicId{
			PLmnIdentity: &e2ap_commondatatypes.PlmnIdentity{
				Value: []byte(plmnID),
			},
			RicId: &e2ap_commondatatypes.BitString{
				Value: ricID,
				Len:   20,
			},
		},
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
	}

	// Create an E2 response
	e2PDUResp := &e2appdudescriptions.E2ApPdu{
		E2ApPdu: &e2appdudescriptions.E2ApPdu_SuccessfulOutcome{
			SuccessfulOutcome: &e2appdudescriptions.SuccessfulOutcome{
				ProcedureCode: &e2appdudescriptions.E2ApElementaryProcedures{
					E2Setup: &e2appdudescriptions.E2Setup{
						SuccessfulOutcome: &e2appducontents.E2SetupResponse{
							ProtocolIes: &e2appducontents.E2SetupResponseIes{
								E2ApProtocolIes4: &e2SetupResp,
							},
						},
					},
				},
			},
		},
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
		ID:           channelID,
		PlmnID:       plmnID,
		RANFunctions: ranFunctions,
	}
	return newChannel(ctx, conn, meta), nil
}

// Get gets a channel by ID
func (m *Manager) Get(ctx context.Context, id ID) (Channel, error) {
	m.channelsMu.RLock()
	defer m.channelsMu.RUnlock()
	channel, ok := m.channels[id]
	if !ok {
		return nil, fmt.Errorf("unknown channel %s", id)
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
