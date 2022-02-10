// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package procedures

import (
	"context"
	v2 "github.com/onosproject/onos-e2t/api/e2ap/v2"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-commondatatypes"
	"syscall"

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
		dispatcher: dispatcher,
		responseCh: make(chan e2appdudescriptions.E2ApPdu),
	}
}

// E2SetupInitiator initiates the E2 setup procedure
type E2SetupInitiator struct {
	dispatcher Dispatcher
	responseCh chan e2appdudescriptions.E2ApPdu
}

func (p *E2SetupInitiator) Initiate(ctx context.Context, request *e2appducontents.E2SetupRequest) (*e2appducontents.E2SetupResponse, *e2appducontents.E2SetupFailure, error) {
	requestPDU := &e2appdudescriptions.E2ApPdu{
		E2ApPdu: &e2appdudescriptions.E2ApPdu_InitiatingMessage{
			InitiatingMessage: &e2appdudescriptions.InitiatingMessage{
				ProcedureCode: int32(v2.ProcedureCodeIDE2setup),
				Criticality:   e2ap_commondatatypes.Criticality_CRITICALITY_REJECT,
				Value: &e2appdudescriptions.InitiatingMessageE2ApElementaryProcedures{
					ImValues: &e2appdudescriptions.InitiatingMessageE2ApElementaryProcedures_E2Setup{
						E2Setup: request,
					},
				},
			},
		},
	}
	// TODO enable it when it is supported
	/*if err := requestPDU.Validate(); err != nil {
		return nil, nil, errors.NewInvalid("E2AP PDU validation failed: %v", err)
	}*/

	if err := p.dispatcher(requestPDU); err != nil {
		return nil, nil, errors.NewUnavailable("E2 Setup initiation failed: %v", err)
	}

	select {
	case responsePDU, ok := <-p.responseCh:
		if !ok {
			return nil, nil, errors.NewUnavailable("connection closed")
		}

		switch msg := responsePDU.E2ApPdu.(type) {
		case *e2appdudescriptions.E2ApPdu_SuccessfulOutcome:
			//return responsePDU.GetSuccessfulOutcome().GetValue().GetE2Setup(), nil, nil
			switch ret := msg.SuccessfulOutcome.Value.SoValues.(type) {
			case *e2appdudescriptions.SuccessfulOutcomeE2ApElementaryProcedures_E2Setup:
				return ret.E2Setup, nil, nil
			default:
				return nil, nil, errors.NewInternal("received unexpected outcome")
			}
		case *e2appdudescriptions.E2ApPdu_UnsuccessfulOutcome:
			//return nil, msg.UnsuccessfulOutcome.Value.GetE2Setup(), nil
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
	}
}

func (p *E2SetupInitiator) Matches(pdu *e2appdudescriptions.E2ApPdu) bool {
	switch msg := pdu.E2ApPdu.(type) {
	case *e2appdudescriptions.E2ApPdu_SuccessfulOutcome:
		//return msg.SuccessfulOutcome.Value.GetE2Setup() != nil
		switch ret := msg.SuccessfulOutcome.Value.SoValues.(type) {
		case *e2appdudescriptions.SuccessfulOutcomeE2ApElementaryProcedures_E2Setup:
			return ret.E2Setup != nil
		default:
			return false
		}
	case *e2appdudescriptions.E2ApPdu_UnsuccessfulOutcome:
		//return msg.UnsuccessfulOutcome.Value.GetE2Setup() != nil
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
	p.responseCh <- *pdu
}

func (p *E2SetupInitiator) Close() error {
	defer func() {
		if err := recover(); err != nil {
			log.Debug("recovering from panic", err)
		}
	}()
	close(p.responseCh)
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
		//return msg.InitiatingMessage.Value.GetE2Setup() != nil
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
					ProcedureCode: int32(v2.ProcedureCodeIDE2setup),
					Criticality:   e2ap_commondatatypes.Criticality_CRITICALITY_REJECT,
					Value: &e2appdudescriptions.SuccessfulOutcomeE2ApElementaryProcedures{
						SoValues: &e2appdudescriptions.SuccessfulOutcomeE2ApElementaryProcedures_E2Setup{
							E2Setup: response,
						},
					},
				},
			},
		}
		// TODO enable validation when it is supported
		/*if err := requestPDU.Validate(); err != nil {
			log.Errorf("E2 Setup response validation failed: %v", err)
		} else {
			err := p.dispatcher(responsePDU)
			if err != nil {
				log.Errorf("E2 Setup response failed: %v", err)
			}
		}*/
		err := p.dispatcher(responsePDU)
		if err != nil {
			if err == context.Canceled || err == context.DeadlineExceeded || err == syscall.EPIPE {
				log.Warnf("E2 Setup response failed: %v", err)
				return
			}
			log.Errorf("E2 Setup response failed: %v", err)
		}

	} else if failure != nil {
		responsePDU := &e2appdudescriptions.E2ApPdu{
			E2ApPdu: &e2appdudescriptions.E2ApPdu_UnsuccessfulOutcome{
				UnsuccessfulOutcome: &e2appdudescriptions.UnsuccessfulOutcome{
					ProcedureCode: int32(v2.ProcedureCodeIDE2setup),
					Criticality:   e2ap_commondatatypes.Criticality_CRITICALITY_REJECT,
					Value: &e2appdudescriptions.UnsuccessfulOutcomeE2ApElementaryProcedures{
						UoValues: &e2appdudescriptions.UnsuccessfulOutcomeE2ApElementaryProcedures_E2Setup{
							E2Setup: failure,
						},
					},
				},
			},
		}
		// TODO enable validation when it is supported
		/*if err := requestPDU.Validate(); err != nil {
			log.Errorf("E2 Setup response validation failed: %v", err)
		} else {
			err := p.dispatcher(responsePDU)
			if err != nil {
				log.Errorf("E2 Setup response failed: %v", err)
			}
		}*/
		err := p.dispatcher(responsePDU)
		if err != nil {
			if err == context.Canceled || err == context.DeadlineExceeded || err == syscall.EPIPE {
				log.Warnf("E2 Setup response failed: %v", err)
				return
			}
			log.Errorf("E2 Setup response failed: %v", err)
		}

	} else {
		log.Errorf("E2 Setup function returned invalid output: no response message found")
	}
}

func (p *E2SetupProcedure) Close() error {
	return nil
}

var _ ElementaryProcedure = &E2SetupProcedure{}
