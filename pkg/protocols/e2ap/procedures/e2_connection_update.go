// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package procedures

import (
	"context"

	"github.com/onosproject/onos-lib-go/pkg/errors"

	e2appdudescriptions "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-descriptions"

	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
)

// E2ConnectionUpdate is an E2 connection update procedure
type E2ConnectionUpdate interface {
	E2ConnectionUpdate(ctx context.Context, request *e2appducontents.E2ConnectionUpdate) (response *e2appducontents.E2ConnectionUpdateAcknowledge, failure *e2appducontents.E2ConnectionUpdateFailure, err error)
}

// NewE2ConnectionUpdateInitiator creates a new E2 connection update initiator
func NewE2ConnectionUpdateInitiator(dispatcher Dispatcher) *E2ConnectionUpdateInitiator {
	return &E2ConnectionUpdateInitiator{
		dispatcher: dispatcher,
		responseCh: make(chan e2appdudescriptions.E2ApPdu),
	}
}

// E2ConnectionUpdateInitiator initiates the E2 connection update procedure
type E2ConnectionUpdateInitiator struct {
	dispatcher Dispatcher
	responseCh chan e2appdudescriptions.E2ApPdu
}

func (p *E2ConnectionUpdateInitiator) Initiate(ctx context.Context, request *e2appducontents.E2ConnectionUpdate) (*e2appducontents.E2ConnectionUpdateAcknowledge, *e2appducontents.E2ConnectionUpdateFailure, error) {
	requestPDU := &e2appdudescriptions.E2ApPdu{
		E2ApPdu: &e2appdudescriptions.E2ApPdu_InitiatingMessage{
			InitiatingMessage: &e2appdudescriptions.InitiatingMessage{
				ProcedureCode: &e2appdudescriptions.E2ApElementaryProcedures{
					E2ConnectionUpdate: &e2appdudescriptions.E2ConnectionUpdateEp{
						InitiatingMessage: request,
					},
				},
			},
		},
	}
	if err := requestPDU.Validate(); err != nil {
		return nil, nil, errors.NewInvalid("E2AP PDU validation failed: %v", err)
	}

	if err := p.dispatcher(requestPDU); err != nil {
		return nil, nil, errors.NewUnavailable("E2 Connection Update initiation failed: %v", err)
	}

	select {
	case responsePDU, ok := <-p.responseCh:
		if !ok {
			return nil, nil, errors.NewUnavailable("connection closed")
		}

		switch msg := responsePDU.E2ApPdu.(type) {
		case *e2appdudescriptions.E2ApPdu_SuccessfulOutcome:
			return msg.SuccessfulOutcome.ProcedureCode.E2ConnectionUpdate.SuccessfulOutcome, nil, nil
		case *e2appdudescriptions.E2ApPdu_UnsuccessfulOutcome:
			return nil, msg.UnsuccessfulOutcome.ProcedureCode.E2ConnectionUpdate.UnsuccessfulOutcome, nil
		default:
			return nil, nil, errors.NewInternal("received unexpected outcome")
		}
	case <-ctx.Done():
		return nil, nil, ctx.Err()
	}
}

func (p *E2ConnectionUpdateInitiator) Matches(pdu *e2appdudescriptions.E2ApPdu) bool {
	switch msg := pdu.E2ApPdu.(type) {
	case *e2appdudescriptions.E2ApPdu_SuccessfulOutcome:
		return msg.SuccessfulOutcome.ProcedureCode.E2ConnectionUpdate != nil
	case *e2appdudescriptions.E2ApPdu_UnsuccessfulOutcome:
		return msg.UnsuccessfulOutcome.ProcedureCode.E2ConnectionUpdate != nil
	default:
		return false
	}
}

func (p *E2ConnectionUpdateInitiator) Handle(pdu *e2appdudescriptions.E2ApPdu) {
	p.responseCh <- *pdu
}

func (p *E2ConnectionUpdateInitiator) Close() error {
	defer func() {
		if err := recover(); err != nil {
			log.Debug("recovering from panic", err)
		}
	}()
	close(p.responseCh)
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
		return msg.InitiatingMessage.ProcedureCode.E2ConnectionUpdate != nil
	default:
		return false
	}
}

func (p *E2ConnectionUpdateProcedure) Handle(requestPDU *e2appdudescriptions.E2ApPdu) {
	response, failure, err := p.handler.E2ConnectionUpdate(context.Background(), requestPDU.GetInitiatingMessage().ProcedureCode.E2ConnectionUpdate.InitiatingMessage)
	if err != nil {
		log.Errorf("E2 Connection Update procedure failed: %v", err)
	} else if response != nil {
		responsePDU := &e2appdudescriptions.E2ApPdu{
			E2ApPdu: &e2appdudescriptions.E2ApPdu_SuccessfulOutcome{
				SuccessfulOutcome: &e2appdudescriptions.SuccessfulOutcome{
					ProcedureCode: &e2appdudescriptions.E2ApElementaryProcedures{
						E2ConnectionUpdate: &e2appdudescriptions.E2ConnectionUpdateEp{
							SuccessfulOutcome: response,
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
				log.Errorf("E2 Connection Update response failed: %v", err)
			}
		}
	} else if failure != nil {
		responsePDU := &e2appdudescriptions.E2ApPdu{
			E2ApPdu: &e2appdudescriptions.E2ApPdu_UnsuccessfulOutcome{
				UnsuccessfulOutcome: &e2appdudescriptions.UnsuccessfulOutcome{
					ProcedureCode: &e2appdudescriptions.E2ApElementaryProcedures{
						E2ConnectionUpdate: &e2appdudescriptions.E2ConnectionUpdateEp{
							UnsuccessfulOutcome: failure,
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
