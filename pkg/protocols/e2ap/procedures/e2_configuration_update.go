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

// E2ConfigurationUpdate is an E2 configuration procedure
type E2ConfigurationUpdate interface {
	E2ConfigurationUpdate(ctx context.Context, request *e2appducontents.E2NodeConfigurationUpdate) (response *e2appducontents.E2NodeConfigurationUpdateAcknowledge, failure *e2appducontents.E2NodeConfigurationUpdateFailure, err error)
}

// NewE2ConfigurationUpdateInitiator  creates a new E2 configuration update initiator
func NewE2ConfigurationUpdateInitiator(dispatcher Dispatcher) *E2ConfigurationUpdateInitiator {
	return &E2ConfigurationUpdateInitiator{
		dispatcher:  dispatcher,
		responseChs: make(map[int32]chan e2appdudescriptions.E2ApPdu),
		closeCh:     make(chan struct{}),
	}
}

// E2ConfigurationUpdateInitiator initiates the E2 configuration update procedure procedure
type E2ConfigurationUpdateInitiator struct {
	dispatcher  Dispatcher
	responseChs map[int32]chan e2appdudescriptions.E2ApPdu
	closeCh     chan struct{}
	mu          sync.RWMutex
}

func (p *E2ConfigurationUpdateInitiator) Initiate(ctx context.Context, request *e2appducontents.E2NodeConfigurationUpdate) (*e2appducontents.E2NodeConfigurationUpdateAcknowledge, *e2appducontents.E2NodeConfigurationUpdateFailure, error) {
	requestPDU := &e2appdudescriptions.E2ApPdu{
		E2ApPdu: &e2appdudescriptions.E2ApPdu_InitiatingMessage{
			InitiatingMessage: &e2appdudescriptions.InitiatingMessage{
				ProcedureCode: int32(e2api.ProcedureCodeIDE2nodeConfigurationUpdate),
				Criticality:   e2apcommondatatypes.Criticality_CRITICALITY_REJECT,
				Value: &e2appdudescriptions.InitiatingMessageE2ApElementaryProcedures{
					ImValues: &e2appdudescriptions.InitiatingMessageE2ApElementaryProcedures_E2NodeConfigurationUpdate{
						E2NodeConfigurationUpdate: request,
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
		return nil, nil, errors.NewUnavailable("E2 configuration update initiation failed: %v", err)
	}

	select {
	case responsePDU, ok := <-responseCh:
		if !ok {
			err := errors.NewUnavailable("connection closed")
			log.Warn(err)
			return nil, nil, err
		}

		switch msg := responsePDU.E2ApPdu.(type) {
		case *e2appdudescriptions.E2ApPdu_SuccessfulOutcome:
			switch ret := msg.SuccessfulOutcome.Value.SoValues.(type) {
			case *e2appdudescriptions.SuccessfulOutcomeE2ApElementaryProcedures_E2NodeConfigurationUpdate:
				return ret.E2NodeConfigurationUpdate, nil, nil
			default:
				return nil, nil, errors.NewInternal("received unexpected outcome")
			}
		case *e2appdudescriptions.E2ApPdu_UnsuccessfulOutcome:
			switch ret := msg.UnsuccessfulOutcome.Value.UoValues.(type) {
			case *e2appdudescriptions.UnsuccessfulOutcomeE2ApElementaryProcedures_E2NodeConfigurationUpdate:
				return nil, ret.E2NodeConfigurationUpdate, nil
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

func (p *E2ConfigurationUpdateInitiator) Matches(pdu *e2appdudescriptions.E2ApPdu) bool {
	switch msg := pdu.E2ApPdu.(type) {
	case *e2appdudescriptions.E2ApPdu_SuccessfulOutcome:
		switch ret := msg.SuccessfulOutcome.Value.SoValues.(type) {
		case *e2appdudescriptions.SuccessfulOutcomeE2ApElementaryProcedures_E2NodeConfigurationUpdate:
			return ret.E2NodeConfigurationUpdate != nil
		default:
			return false
		}
	case *e2appdudescriptions.E2ApPdu_UnsuccessfulOutcome:
		switch ret := msg.UnsuccessfulOutcome.Value.UoValues.(type) {
		case *e2appdudescriptions.UnsuccessfulOutcomeE2ApElementaryProcedures_E2NodeConfigurationUpdate:
			return ret.E2NodeConfigurationUpdate != nil
		default:
			return false
		}
	default:
		return false
	}
}

func (p *E2ConfigurationUpdateInitiator) Handle(pdu *e2appdudescriptions.E2ApPdu) {
	var transactionID int32
	switch pdu.GetE2ApPdu().(type) {
	case *e2appdudescriptions.E2ApPdu_SuccessfulOutcome:
		for _, v := range pdu.GetSuccessfulOutcome().GetValue().GetE2NodeConfigurationUpdate().GetProtocolIes() {
			if v.Id == int32(e2api.ProtocolIeIDTransactionID) {
				transactionID = v.GetValue().GetTrId().GetValue()
				break
			}
		}
	case *e2appdudescriptions.E2ApPdu_UnsuccessfulOutcome:
		for _, v := range pdu.GetUnsuccessfulOutcome().GetValue().GetE2NodeConfigurationUpdate().GetProtocolIes() {
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
		log.Warnf("Received RIC Configuration update response for unknown transaction %d", transactionID)
	}
}

func (p *E2ConfigurationUpdateInitiator) Close() error {
	close(p.closeCh)
	return nil
}

var _ ElementaryProcedure = &E2ConfigurationUpdateInitiator{}

// NewE2ConfigurationUpdateProcedure creates a new E2 configuration update procedure
func NewE2ConfigurationUpdateProcedure(dispatcher Dispatcher, handler E2ConfigurationUpdate) *E2ConfigurationUpdateProcedure {
	return &E2ConfigurationUpdateProcedure{
		dispatcher: dispatcher,
		handler:    handler,
	}
}

// E2ConfigurationUpdateProcedure  implements the E2 configuration update procedure
type E2ConfigurationUpdateProcedure struct {
	dispatcher Dispatcher
	handler    E2ConfigurationUpdate
}

func (p *E2ConfigurationUpdateProcedure) Matches(pdu *e2appdudescriptions.E2ApPdu) bool {
	switch msg := pdu.E2ApPdu.(type) {
	case *e2appdudescriptions.E2ApPdu_InitiatingMessage:
		switch ret := msg.InitiatingMessage.Value.ImValues.(type) {
		case *e2appdudescriptions.InitiatingMessageE2ApElementaryProcedures_E2NodeConfigurationUpdate:
			return ret.E2NodeConfigurationUpdate != nil
		default:
			return false
		}
	default:
		return false
	}
}

func (p *E2ConfigurationUpdateProcedure) Handle(requestPDU *e2appdudescriptions.E2ApPdu) {
	response, failure, err := p.handler.E2ConfigurationUpdate(context.Background(), requestPDU.GetInitiatingMessage().GetValue().GetE2NodeConfigurationUpdate())
	if err != nil {
		log.Errorf("E2 configuration update procedure failed: %v", err)
	} else if response != nil {
		responsePDU := &e2appdudescriptions.E2ApPdu{
			E2ApPdu: &e2appdudescriptions.E2ApPdu_SuccessfulOutcome{
				SuccessfulOutcome: &e2appdudescriptions.SuccessfulOutcome{
					ProcedureCode: int32(e2api.ProcedureCodeIDE2nodeConfigurationUpdate),
					Criticality:   e2apcommondatatypes.Criticality_CRITICALITY_REJECT,
					Value: &e2appdudescriptions.SuccessfulOutcomeE2ApElementaryProcedures{
						SoValues: &e2appdudescriptions.SuccessfulOutcomeE2ApElementaryProcedures_E2NodeConfigurationUpdate{
							E2NodeConfigurationUpdate: response,
						},
					},
				},
			},
		}

		if err := requestPDU.Validate(); err != nil {
			log.Errorf("E2 configuration update validation failed: %v", err)
		} else {
			err := p.dispatcher(responsePDU)
			if err != nil {
				if err == context.Canceled || err == context.DeadlineExceeded || err == syscall.EPIPE || err == syscall.EBADF {
					log.Warnf("E2 configuration update response failed: %v", err)
					return
				}
				log.Errorf("E2 configuration update response failed: %v", err)
			}
		}

	} else if failure != nil {
		responsePDU := &e2appdudescriptions.E2ApPdu{
			E2ApPdu: &e2appdudescriptions.E2ApPdu_UnsuccessfulOutcome{
				UnsuccessfulOutcome: &e2appdudescriptions.UnsuccessfulOutcome{
					ProcedureCode: int32(e2api.ProcedureCodeIDE2nodeConfigurationUpdate),
					Criticality:   e2apcommondatatypes.Criticality_CRITICALITY_REJECT,
					Value: &e2appdudescriptions.UnsuccessfulOutcomeE2ApElementaryProcedures{
						UoValues: &e2appdudescriptions.UnsuccessfulOutcomeE2ApElementaryProcedures_E2NodeConfigurationUpdate{
							E2NodeConfigurationUpdate: failure,
						},
					},
				},
			},
		}

		if err := requestPDU.Validate(); err != nil {
			log.Errorf("E2 configuration update validation failed: %v", err)
		} else {
			err := p.dispatcher(responsePDU)
			if err != nil {
				if err == context.Canceled || err == context.DeadlineExceeded || err == syscall.EPIPE || err == syscall.EBADF {
					log.Warnf("E2 configuration update response failed: %v", err)
					return
				}
				log.Errorf("E2 configuration update response failed: %v", err)
			}
		}

	} else {
		log.Errorf("E2 configuration update function returned invalid output: no response message found")
	}
}

func (p *E2ConfigurationUpdateProcedure) Close() error {
	return nil
}

var _ ElementaryProcedure = &E2ConfigurationUpdateProcedure{}
