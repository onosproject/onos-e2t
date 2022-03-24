// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package procedures

import (
	"context"
	"sync"
	"syscall"

	e2api "github.com/onosproject/onos-e2t/api/e2ap/v2"
	e2apcommondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-commondatatypes"

	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
	e2appdudescriptions "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-descriptions"
	"github.com/onosproject/onos-lib-go/pkg/errors"
)

// RICControl is a RIC control procedure
type RICControl interface {
	RICControl(ctx context.Context, request *e2appducontents.RiccontrolRequest) (response *e2appducontents.RiccontrolAcknowledge, failure *e2appducontents.RiccontrolFailure, err error)
}

func NewRICControlInitiator(dispatcher Dispatcher) *RICControlInitiator {
	return &RICControlInitiator{
		dispatcher:  dispatcher,
		responseChs: make(map[int32]chan e2appdudescriptions.E2ApPdu),
		closeCh:     make(chan struct{}),
	}
}

type RICControlInitiator struct {
	dispatcher  Dispatcher
	responseChs map[int32]chan e2appdudescriptions.E2ApPdu
	closeCh     chan struct{}
	mu          sync.RWMutex
}

func (p *RICControlInitiator) Initiate(ctx context.Context, request *e2appducontents.RiccontrolRequest) (*e2appducontents.RiccontrolAcknowledge, *e2appducontents.RiccontrolFailure, error) {
	requestPDU := &e2appdudescriptions.E2ApPdu{
		E2ApPdu: &e2appdudescriptions.E2ApPdu_InitiatingMessage{
			InitiatingMessage: &e2appdudescriptions.InitiatingMessage{
				ProcedureCode: int32(e2api.ProcedureCodeIDRICcontrol),
				Criticality:   e2apcommondatatypes.Criticality_CRITICALITY_REJECT,
				Value: &e2appdudescriptions.InitiatingMessageE2ApElementaryProcedures{
					ImValues: &e2appdudescriptions.InitiatingMessageE2ApElementaryProcedures_RicControl{
						RicControl: request,
					},
				},
			},
		},
	}
	if err := requestPDU.Validate(); err != nil {
		return nil, nil, errors.NewInvalid("E2AP PDU validation failed: %v", err)
	}

	responseCh := make(chan e2appdudescriptions.E2ApPdu, 1)
	var requestID int32
	for _, v := range request.GetProtocolIes() {
		if v.Id == int32(e2api.ProtocolIeIDRicrequestID) {
			requestID = v.GetValue().GetRicrequestId().GetRicRequestorId()
			break
		}
	}

	p.mu.Lock()
	p.responseChs[requestID] = responseCh
	p.mu.Unlock()

	if err := p.dispatcher(requestPDU); err != nil {
		return nil, nil, errors.NewUnavailable("RIC Control initiation failed: %v", err)
	}

	select {
	case responsePDU, ok := <-responseCh:
		if !ok {
			return nil, nil, errors.NewUnavailable("connection closed")
		}

		switch response := responsePDU.E2ApPdu.(type) {
		case *e2appdudescriptions.E2ApPdu_SuccessfulOutcome:
			switch ret := response.SuccessfulOutcome.Value.SoValues.(type) {
			case *e2appdudescriptions.SuccessfulOutcomeE2ApElementaryProcedures_RicControl:
				return ret.RicControl, nil, nil
			default:
				return nil, nil, errors.NewInternal("received unexpected outcome")
			}
		case *e2appdudescriptions.E2ApPdu_UnsuccessfulOutcome:
			switch ret := response.UnsuccessfulOutcome.Value.UoValues.(type) {
			case *e2appdudescriptions.UnsuccessfulOutcomeE2ApElementaryProcedures_RicControl:
				return nil, ret.RicControl, nil
			default:
				return nil, nil, errors.NewInternal("received unexpected outcome")
			}
		default:
			return nil, nil, errors.NewInternal("received unexpected outcome")
		}
	case <-ctx.Done():
		return nil, nil, ctx.Err()
	case _, ok := <-p.closeCh:
		if !ok {
			return nil, nil, errors.NewUnavailable("connection closed")
		}
		return nil, nil, nil
	}
}

func (p *RICControlInitiator) Matches(pdu *e2appdudescriptions.E2ApPdu) bool {
	switch msg := pdu.E2ApPdu.(type) {
	case *e2appdudescriptions.E2ApPdu_SuccessfulOutcome:
		switch ret := msg.SuccessfulOutcome.Value.SoValues.(type) {
		case *e2appdudescriptions.SuccessfulOutcomeE2ApElementaryProcedures_RicControl:
			return ret.RicControl != nil
		default:
			return false
		}
	case *e2appdudescriptions.E2ApPdu_UnsuccessfulOutcome:
		switch ret := msg.UnsuccessfulOutcome.Value.UoValues.(type) {
		case *e2appdudescriptions.UnsuccessfulOutcomeE2ApElementaryProcedures_RicControl:
			return ret.RicControl != nil
		default:
			return false
		}
	default:
		return false
	}
}

func (p *RICControlInitiator) Handle(pdu *e2appdudescriptions.E2ApPdu) {
	var requestID int32
	// Assuming that RequestID is always included in the message
	switch response := pdu.E2ApPdu.(type) {
	case *e2appdudescriptions.E2ApPdu_SuccessfulOutcome:
		for _, v := range response.SuccessfulOutcome.Value.GetRicControl().GetProtocolIes() {
			if v.Id == int32(e2api.ProtocolIeIDRicrequestID) {
				requestID = v.GetValue().GetRicrequestId().GetRicRequestorId()
				break
			}
		}
	case *e2appdudescriptions.E2ApPdu_UnsuccessfulOutcome:
		for _, v := range response.UnsuccessfulOutcome.Value.GetRicControl().GetProtocolIes() {
			if v.Id == int32(e2api.ProtocolIeIDRicrequestID) {
				requestID = v.GetValue().GetRicrequestId().GetRicRequestorId()
				break
			}
		}
	}

	p.mu.RLock()
	responseCh, ok := p.responseChs[requestID]
	p.mu.RUnlock()
	if ok {
		responseCh <- *pdu
	} else {
		log.Warnf("Received RIC Control response for unknown request %d", requestID)
	}
}

func (p *RICControlInitiator) Close() error {
	close(p.closeCh)
	return nil
}

var _ ElementaryProcedure = &RICControlInitiator{}

func NewRICControlProcedure(dispatcher Dispatcher, handler RICControl) *RICControlProcedure {
	return &RICControlProcedure{
		dispatcher: dispatcher,
		handler:    handler,
	}
}

type RICControlProcedure struct {
	dispatcher Dispatcher
	handler    RICControl
}

func (p *RICControlProcedure) Matches(pdu *e2appdudescriptions.E2ApPdu) bool {
	switch msg := pdu.E2ApPdu.(type) {
	case *e2appdudescriptions.E2ApPdu_InitiatingMessage:
		//return msg.InitiatingMessage.Value.GetRicControl() != nil
		switch ret := msg.InitiatingMessage.Value.ImValues.(type) {
		case *e2appdudescriptions.InitiatingMessageE2ApElementaryProcedures_RicControl:
			return ret.RicControl != nil
		default:
			return false
		}
	default:
		return false
	}
}

func (p *RICControlProcedure) Handle(requestPDU *e2appdudescriptions.E2ApPdu) {
	response, failure, err := p.handler.RICControl(context.Background(), requestPDU.GetInitiatingMessage().GetValue().GetRicControl())
	if err != nil {
		log.Errorf("RIC Control procedure failed: %v", err)
	} else if response != nil {
		responsePDU := &e2appdudescriptions.E2ApPdu{
			E2ApPdu: &e2appdudescriptions.E2ApPdu_SuccessfulOutcome{
				SuccessfulOutcome: &e2appdudescriptions.SuccessfulOutcome{
					ProcedureCode: int32(e2api.ProcedureCodeIDRICcontrol),
					Criticality:   e2apcommondatatypes.Criticality_CRITICALITY_REJECT,
					Value: &e2appdudescriptions.SuccessfulOutcomeE2ApElementaryProcedures{
						SoValues: &e2appdudescriptions.SuccessfulOutcomeE2ApElementaryProcedures_RicControl{
							RicControl: response,
						},
					},
				},
			},
		}
		if err := requestPDU.Validate(); err != nil {
			log.Errorf("RIC Control response validation failed: %v", err)
		} else {
			err := p.dispatcher(responsePDU)
			if err != nil {
				if err == context.Canceled || err == context.DeadlineExceeded || err == syscall.EPIPE || err == syscall.EBADF {
					log.Warnf("RIC Control response failed: %v", err)
					return
				}
				log.Errorf("RIC Control response failed: %v", err)
			}
		}

	} else if failure != nil {
		responsePDU := &e2appdudescriptions.E2ApPdu{
			E2ApPdu: &e2appdudescriptions.E2ApPdu_UnsuccessfulOutcome{
				UnsuccessfulOutcome: &e2appdudescriptions.UnsuccessfulOutcome{
					ProcedureCode: int32(e2api.ProcedureCodeIDRICcontrol),
					Criticality:   e2apcommondatatypes.Criticality_CRITICALITY_REJECT,
					Value: &e2appdudescriptions.UnsuccessfulOutcomeE2ApElementaryProcedures{
						UoValues: &e2appdudescriptions.UnsuccessfulOutcomeE2ApElementaryProcedures_RicControl{
							RicControl: failure,
						},
					},
				},
			},
		}
		if err := requestPDU.Validate(); err != nil {
			log.Errorf("RIC Control response validation failed: %v", err)
		} else {
			err := p.dispatcher(responsePDU)
			if err != nil {
				if err == context.Canceled || err == context.DeadlineExceeded || err == syscall.EPIPE || err == syscall.EBADF {
					log.Warnf("RIC Control response failed: %v", err)
					return
				}
				log.Errorf("RIC Control response failed: %v", err)
			}
		}

	} else {
		log.Debugf("RIC Control function does not have a response message")
	}
}

func (p *RICControlProcedure) Close() error {
	return nil
}

var _ ElementaryProcedure = &RICControlProcedure{}
