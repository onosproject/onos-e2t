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

	"github.com/onosproject/onos-lib-go/pkg/errors"

	e2appdudescriptions "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-descriptions"

	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
)

// E2ConnectionUpdate is an E2 connection update procedure
type E2ConnectionUpdate interface {
	E2ConnectionUpdate(ctx context.Context, request *e2appducontents.E2ConnectionUpdate) (response *e2appducontents.E2ConnectionUpdateAcknowledge, failure *e2appducontents.E2ConnectionUpdateFailure, err error)
}

// NewE2ConnectionUpdateInitiator creates a new E2 connection update initiator
func NewE2ConnectionUpdateInitiator(dispatcher Dispatcher) *E2ConnectionUpdateInitiator {
	return &E2ConnectionUpdateInitiator{
		dispatcher:  dispatcher,
		responseChs: make(map[int32]chan e2appdudescriptions.E2ApPdu),
		closeCh:     make(chan struct{}),
	}
}

// E2ConnectionUpdateInitiator initiates the E2 connection update procedure
type E2ConnectionUpdateInitiator struct {
	dispatcher  Dispatcher
	responseChs map[int32]chan e2appdudescriptions.E2ApPdu
	closeCh     chan struct{}
	mu          sync.RWMutex
}

func (p *E2ConnectionUpdateInitiator) Initiate(ctx context.Context, request *e2appducontents.E2ConnectionUpdate) (*e2appducontents.E2ConnectionUpdateAcknowledge, *e2appducontents.E2ConnectionUpdateFailure, error) {
	requestPDU := &e2appdudescriptions.E2ApPdu{
		E2ApPdu: &e2appdudescriptions.E2ApPdu_InitiatingMessage{
			InitiatingMessage: &e2appdudescriptions.InitiatingMessage{
				ProcedureCode: int32(e2api.ProcedureCodeIDE2connectionUpdate),
				Criticality:   e2apcommondatatypes.Criticality_CRITICALITY_REJECT,
				Value: &e2appdudescriptions.InitiatingMessageE2ApElementaryProcedures{
					ImValues: &e2appdudescriptions.InitiatingMessageE2ApElementaryProcedures_E2ConnectionUpdate{
						E2ConnectionUpdate: request,
					},
				},
			},
		},
	}
	if err := requestPDU.Validate(); err != nil {
		return nil, nil, errors.NewInvalid("E2AP PDU validation failed: %v", err)
	}

	responseCh := make(chan e2appdudescriptions.E2ApPdu, 1)
	var transactionID int32 = -1
	for _, v := range request.GetProtocolIes() {
		if v.Id == int32(e2api.ProtocolIeIDTransactionID) {
			transactionID = v.GetValue().GetTrId().GetValue()
			break
		}
	}
	p.mu.Lock()
	p.responseChs[transactionID] = responseCh
	p.mu.Unlock()

	if err := p.dispatcher(requestPDU); err != nil {
		return nil, nil, errors.NewUnavailable("E2 Connection Update initiation failed: %v", err)
	}

	select {
	case responsePDU, ok := <-responseCh:
		if !ok {
			return nil, nil, errors.NewUnavailable("connection closed")
		}

		switch msg := responsePDU.E2ApPdu.(type) {
		case *e2appdudescriptions.E2ApPdu_SuccessfulOutcome:
			switch ret := msg.SuccessfulOutcome.Value.SoValues.(type) {
			case *e2appdudescriptions.SuccessfulOutcomeE2ApElementaryProcedures_E2ConnectionUpdate:
				return ret.E2ConnectionUpdate, nil, nil
			default:
				return nil, nil, errors.NewInternal("received unexpected outcome")
			}
		case *e2appdudescriptions.E2ApPdu_UnsuccessfulOutcome:
			//return nil, msg.UnsuccessfulOutcome.Value.GetE2ConnectionUpdate(), nil
			switch ret := msg.UnsuccessfulOutcome.Value.UoValues.(type) {
			case *e2appdudescriptions.UnsuccessfulOutcomeE2ApElementaryProcedures_E2ConnectionUpdate:
				return nil, ret.E2ConnectionUpdate, nil
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

func (p *E2ConnectionUpdateInitiator) Matches(pdu *e2appdudescriptions.E2ApPdu) bool {
	switch msg := pdu.E2ApPdu.(type) {
	case *e2appdudescriptions.E2ApPdu_SuccessfulOutcome:
		//return msg.SuccessfulOutcome.Value.GetE2ConnectionUpdate() != nil
		switch ret := msg.SuccessfulOutcome.Value.SoValues.(type) {
		case *e2appdudescriptions.SuccessfulOutcomeE2ApElementaryProcedures_E2ConnectionUpdate:
			return ret.E2ConnectionUpdate != nil
		default:
			return false
		}
	case *e2appdudescriptions.E2ApPdu_UnsuccessfulOutcome:
		//return msg.UnsuccessfulOutcome.Value.GetE2ConnectionUpdate() != nil
		switch ret := msg.UnsuccessfulOutcome.Value.UoValues.(type) {
		case *e2appdudescriptions.UnsuccessfulOutcomeE2ApElementaryProcedures_E2ConnectionUpdate:
			return ret.E2ConnectionUpdate != nil
		default:
			return false
		}
	default:
		return false
	}
}

func (p *E2ConnectionUpdateInitiator) Handle(pdu *e2appdudescriptions.E2ApPdu) {
	var transactionID int32
	switch pdu.GetE2ApPdu().(type) {
	case *e2appdudescriptions.E2ApPdu_SuccessfulOutcome:
		for _, v := range pdu.GetSuccessfulOutcome().GetValue().GetE2ConnectionUpdate().GetProtocolIes() {
			if v.Id == int32(e2api.ProtocolIeIDTransactionID) {
				transactionID = v.GetValue().GetTrId().GetValue()
				break
			}
		}
	case *e2appdudescriptions.E2ApPdu_UnsuccessfulOutcome:
		for _, v := range pdu.GetUnsuccessfulOutcome().GetValue().GetE2ConnectionUpdate().GetProtocolIes() {
			if v.Id == int32(e2api.ProtocolIeIDTransactionID) {
				transactionID = v.GetValue().GetTrId().GetValue()
				break
			}
		}
	}

	p.mu.RLock()
	responseCh, ok := p.responseChs[transactionID]
	p.mu.RUnlock()
	if ok {
		responseCh <- *pdu
	} else {
		log.Warnf("Received RIC Connection update response for unknown transaction %d", transactionID)
	}
}

func (p *E2ConnectionUpdateInitiator) Close() error {
	close(p.closeCh)
	return nil
}

var _ ElementaryProcedure = &E2ConnectionUpdateInitiator{}

// NewE2ConnectionUpdateProcedure creates a new E2 connection update procedure
func NewE2ConnectionUpdateProcedure(dispatcher Dispatcher, handler E2ConnectionUpdate) *E2ConnectionUpdateProcedure {
	return &E2ConnectionUpdateProcedure{
		dispatcher: dispatcher,
		handler:    handler,
	}
}

// E2ConnectionUpdateProcedure implements the E2 connection update procedure
type E2ConnectionUpdateProcedure struct {
	dispatcher Dispatcher
	handler    E2ConnectionUpdate
}

func (p *E2ConnectionUpdateProcedure) Matches(pdu *e2appdudescriptions.E2ApPdu) bool {
	switch msg := pdu.E2ApPdu.(type) {
	case *e2appdudescriptions.E2ApPdu_InitiatingMessage:
		switch ret := msg.InitiatingMessage.Value.ImValues.(type) {
		case *e2appdudescriptions.InitiatingMessageE2ApElementaryProcedures_E2ConnectionUpdate:
			return ret.E2ConnectionUpdate != nil
		default:
			return false
		}
	default:
		return false
	}
}

func (p *E2ConnectionUpdateProcedure) Handle(requestPDU *e2appdudescriptions.E2ApPdu) {
	response, failure, err := p.handler.E2ConnectionUpdate(context.Background(), requestPDU.GetInitiatingMessage().GetValue().GetE2ConnectionUpdate())
	if err != nil {
		log.Errorf("E2 Connection Update procedure failed: %v", err)
	} else if response != nil {
		responsePDU := &e2appdudescriptions.E2ApPdu{
			E2ApPdu: &e2appdudescriptions.E2ApPdu_SuccessfulOutcome{
				SuccessfulOutcome: &e2appdudescriptions.SuccessfulOutcome{
					ProcedureCode: int32(e2api.ProcedureCodeIDE2connectionUpdate),
					Criticality:   e2apcommondatatypes.Criticality_CRITICALITY_REJECT,
					Value: &e2appdudescriptions.SuccessfulOutcomeE2ApElementaryProcedures{
						SoValues: &e2appdudescriptions.SuccessfulOutcomeE2ApElementaryProcedures_E2ConnectionUpdate{
							E2ConnectionUpdate: response,
						},
					},
				},
			},
		}
		if err := requestPDU.Validate(); err != nil {
			log.Errorf("E2 Connection Update response validation failed: %v", err)
		} else {
			log.Debugf("Response PDU is following\n%v", responsePDU)
			err := p.dispatcher(responsePDU)
			if err != nil {
				if err == context.Canceled || err == context.DeadlineExceeded || err == syscall.EPIPE || err == syscall.EBADF {
					log.Warnf("E2 Connection Update response failed: %v", err)
					return
				}
				log.Errorf("E2 Connection Update response failed: %v", err)
			}
		}
	} else if failure != nil {
		responsePDU := &e2appdudescriptions.E2ApPdu{
			E2ApPdu: &e2appdudescriptions.E2ApPdu_UnsuccessfulOutcome{
				UnsuccessfulOutcome: &e2appdudescriptions.UnsuccessfulOutcome{
					ProcedureCode: int32(e2api.ProcedureCodeIDE2connectionUpdate),
					Criticality:   e2apcommondatatypes.Criticality_CRITICALITY_REJECT,
					Value: &e2appdudescriptions.UnsuccessfulOutcomeE2ApElementaryProcedures{
						UoValues: &e2appdudescriptions.UnsuccessfulOutcomeE2ApElementaryProcedures_E2ConnectionUpdate{
							E2ConnectionUpdate: failure,
						},
					},
				},
			},
		}
		if err := requestPDU.Validate(); err != nil {
			log.Errorf("E2 Connection Update response validation failed: %v", err)
		} else {
			err := p.dispatcher(responsePDU)
			if err != nil {
				if err == context.Canceled || err == context.DeadlineExceeded || err == syscall.EPIPE || err == syscall.EBADF {
					log.Warnf("E2 Connection Update response failed: %v", err)
					return
				}
				log.Errorf("E2 Connection Update response failed: %v", err)
			}
		}

	} else {
		log.Errorf("E2 Connection Update function returned invalid output: no response message found")
	}
}

func (p *E2ConnectionUpdateProcedure) Close() error {
	return nil
}

var _ ElementaryProcedure = &E2ConnectionUpdateProcedure{}
