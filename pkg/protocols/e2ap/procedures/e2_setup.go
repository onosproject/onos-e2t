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

// E2Setup is an E2 setup procedure
type E2Setup interface {
	E2Setup(ctx context.Context, request *e2appducontents.E2SetupRequest) (response *e2appducontents.E2SetupResponse, failure *e2appducontents.E2SetupFailure, err error)
}

// NewE2SetupInitiator creates a new E2 setup initiator
func NewE2SetupInitiator(dispatcher Dispatcher) *E2SetupInitiator {
	return &E2SetupInitiator{
		dispatcher:  dispatcher,
		responseChs: make(map[int32]chan e2appdudescriptions.E2ApPdu),
		closeCh:     make(chan struct{}),
	}
}

// E2SetupInitiator initiates the E2 setup procedure
type E2SetupInitiator struct {
	dispatcher  Dispatcher
	responseChs map[int32]chan e2appdudescriptions.E2ApPdu
	closeCh     chan struct{}
	mu          sync.RWMutex
}

func (p *E2SetupInitiator) Initiate(ctx context.Context, request *e2appducontents.E2SetupRequest) (*e2appducontents.E2SetupResponse, *e2appducontents.E2SetupFailure, error) {
	requestPDU := &e2appdudescriptions.E2ApPdu{
		E2ApPdu: &e2appdudescriptions.E2ApPdu_InitiatingMessage{
			InitiatingMessage: &e2appdudescriptions.InitiatingMessage{
				ProcedureCode: int32(e2api.ProcedureCodeIDE2setup),
				Criticality:   e2apcommondatatypes.Criticality_CRITICALITY_REJECT,
				Value: &e2appdudescriptions.InitiatingMessageE2ApElementaryProcedures{
					ImValues: &e2appdudescriptions.InitiatingMessageE2ApElementaryProcedures_E2Setup{
						E2Setup: request,
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
		return nil, nil, errors.NewUnavailable("E2 Setup initiation failed: %v", err)
	}

	select {
	case responsePDU, ok := <-responseCh:
		if !ok {
			return nil, nil, errors.NewUnavailable("connection closed")
		}

		switch msg := responsePDU.E2ApPdu.(type) {
		case *e2appdudescriptions.E2ApPdu_SuccessfulOutcome:
			switch ret := msg.SuccessfulOutcome.Value.SoValues.(type) {
			case *e2appdudescriptions.SuccessfulOutcomeE2ApElementaryProcedures_E2Setup:
				return ret.E2Setup, nil, nil
			default:
				return nil, nil, errors.NewInternal("received unexpected outcome")
			}
		case *e2appdudescriptions.E2ApPdu_UnsuccessfulOutcome:
			switch ret := msg.UnsuccessfulOutcome.Value.UoValues.(type) {
			case *e2appdudescriptions.UnsuccessfulOutcomeE2ApElementaryProcedures_E2Setup:
				return nil, ret.E2Setup, nil
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

func (p *E2SetupInitiator) Matches(pdu *e2appdudescriptions.E2ApPdu) bool {
	switch msg := pdu.E2ApPdu.(type) {
	case *e2appdudescriptions.E2ApPdu_SuccessfulOutcome:
		switch ret := msg.SuccessfulOutcome.Value.SoValues.(type) {
		case *e2appdudescriptions.SuccessfulOutcomeE2ApElementaryProcedures_E2Setup:
			return ret.E2Setup != nil
		default:
			return false
		}
	case *e2appdudescriptions.E2ApPdu_UnsuccessfulOutcome:
		switch ret := msg.UnsuccessfulOutcome.Value.UoValues.(type) {
		case *e2appdudescriptions.UnsuccessfulOutcomeE2ApElementaryProcedures_E2Setup:
			return ret.E2Setup != nil
		default:
			return false
		}
	default:
		return false
	}
}

func (p *E2SetupInitiator) Handle(pdu *e2appdudescriptions.E2ApPdu) {
	var transactionID int32
	switch pdu.GetE2ApPdu().(type) {
	case *e2appdudescriptions.E2ApPdu_SuccessfulOutcome:
		for _, v := range pdu.GetSuccessfulOutcome().GetValue().GetE2Setup().GetProtocolIes() {
			if v.Id == int32(e2api.ProtocolIeIDTransactionID) {
				transactionID = v.GetValue().GetTrId().GetValue()
				break
			}
		}
	case *e2appdudescriptions.E2ApPdu_UnsuccessfulOutcome:
		for _, v := range pdu.GetUnsuccessfulOutcome().GetValue().GetE2Setup().GetProtocolIes() {
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
		log.Warnf("Received RIC E2 setup response for unknown transaction %d", transactionID)
	}

}

func (p *E2SetupInitiator) Close() error {
	close(p.closeCh)
	return nil
}

var _ ElementaryProcedure = &E2SetupInitiator{}

// NewE2SetupProcedure creates a new E2 setup procedure
func NewE2SetupProcedure(dispatcher Dispatcher, handler E2Setup) *E2SetupProcedure {
	return &E2SetupProcedure{
		dispatcher: dispatcher,
		handler:    handler,
	}
}

// E2SetupProcedure implements the E2 setup procedure
type E2SetupProcedure struct {
	dispatcher Dispatcher
	handler    E2Setup
}

func (p *E2SetupProcedure) Matches(pdu *e2appdudescriptions.E2ApPdu) bool {
	switch msg := pdu.E2ApPdu.(type) {
	case *e2appdudescriptions.E2ApPdu_InitiatingMessage:
		switch ret := msg.InitiatingMessage.Value.ImValues.(type) {
		case *e2appdudescriptions.InitiatingMessageE2ApElementaryProcedures_E2Setup:
			return ret.E2Setup != nil
		default:
			return false
		}
	default:
		return false
	}
}

func (p *E2SetupProcedure) Handle(requestPDU *e2appdudescriptions.E2ApPdu) {
	response, failure, err := p.handler.E2Setup(context.Background(), requestPDU.GetInitiatingMessage().Value.GetE2Setup())
	if err != nil {
		log.Errorf("E2 Setup procedure failed: %v", err)
	} else if response != nil {
		responsePDU := &e2appdudescriptions.E2ApPdu{
			E2ApPdu: &e2appdudescriptions.E2ApPdu_SuccessfulOutcome{
				SuccessfulOutcome: &e2appdudescriptions.SuccessfulOutcome{
					ProcedureCode: int32(e2api.ProcedureCodeIDE2setup),
					Criticality:   e2apcommondatatypes.Criticality_CRITICALITY_REJECT,
					Value: &e2appdudescriptions.SuccessfulOutcomeE2ApElementaryProcedures{
						SoValues: &e2appdudescriptions.SuccessfulOutcomeE2ApElementaryProcedures_E2Setup{
							E2Setup: response,
						},
					},
				},
			},
		}
		if err := requestPDU.Validate(); err != nil {
			log.Errorf("E2 Setup response validation failed: %v", err)
		} else {
			err := p.dispatcher(responsePDU)
			if err != nil {
				if err == context.Canceled || err == context.DeadlineExceeded || err == syscall.EPIPE || err == syscall.EBADF {
					log.Warnf("E2 Setup response failed: %v", err)
					return
				}
				log.Errorf("E2 Setup response failed: %v", err)
			}
		}

	} else if failure != nil {
		responsePDU := &e2appdudescriptions.E2ApPdu{
			E2ApPdu: &e2appdudescriptions.E2ApPdu_UnsuccessfulOutcome{
				UnsuccessfulOutcome: &e2appdudescriptions.UnsuccessfulOutcome{
					ProcedureCode: int32(e2api.ProcedureCodeIDE2setup),
					Criticality:   e2apcommondatatypes.Criticality_CRITICALITY_REJECT,
					Value: &e2appdudescriptions.UnsuccessfulOutcomeE2ApElementaryProcedures{
						UoValues: &e2appdudescriptions.UnsuccessfulOutcomeE2ApElementaryProcedures_E2Setup{
							E2Setup: failure,
						},
					},
				},
			},
		}
		if err := requestPDU.Validate(); err != nil {
			log.Errorf("E2 Setup response validation failed: %v", err)
		} else {
			err := p.dispatcher(responsePDU)
			if err != nil {
				if err == context.Canceled || err == context.DeadlineExceeded || err == syscall.EPIPE || err == syscall.EBADF {
					log.Warnf("E2 Setup response failed: %v", err)
					return
				}
				log.Errorf("E2 Setup response failed: %v", err)
			}
		}

	} else {
		log.Errorf("E2 Setup function returned invalid output: no response message found")
	}
}

func (p *E2SetupProcedure) Close() error {
	return nil
}

var _ ElementaryProcedure = &E2SetupProcedure{}
