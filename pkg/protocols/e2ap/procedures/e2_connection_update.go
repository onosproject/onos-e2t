// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package procedures

import (
	"context"
	v2 "github.com/onosproject/onos-e2t/api/e2ap/v2"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-commondatatypes"
	e2apcommondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-commondatatypes"
	"syscall"

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
				ProcedureCode: int32(v2.ProcedureCodeIDE2connectionUpdate),
				Criticality:   e2apcommondatatypes.Criticality_CRITICALITY_REJECT,
				Value: &e2appdudescriptions.InitiatingMessageE2ApElementaryProcedures{
					ImValues: &e2appdudescriptions.InitiatingMessageE2ApElementaryProcedures_E2ConnectionUpdate{
						E2ConnectionUpdate: request,
					},
				},
			},
		},
	}
	// TODO enable validation when it is supported
	/*if err := requestPDU.Validate(); err != nil {
		return nil, nil, errors.NewInvalid("E2AP PDU validation failed: %v", err)
	}*/

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
			//return msg.SuccessfulOutcome.Value.GetE2ConnectionUpdate(), nil, nil
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
		//return msg.InitiatingMessage.Value.GetE2ConnectionUpdate() != nil
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
					ProcedureCode: int32(v2.ProcedureCodeIDE2connectionUpdate),
					Criticality:   e2ap_commondatatypes.Criticality_CRITICALITY_REJECT,
					Value: &e2appdudescriptions.SuccessfulOutcomeE2ApElementaryProcedures{
						SoValues: &e2appdudescriptions.SuccessfulOutcomeE2ApElementaryProcedures_E2ConnectionUpdate{
							E2ConnectionUpdate: response,
						},
					},
				},
			},
		}
		// TODO enable validation when it is supported
		/*if err := requestPDU.Validate(); err != nil {
			log.Errorf("E2 Connection Update response validation failed: %v", err)
		} else {
			err := p.dispatcher(responsePDU)
			if err != nil {
				log.Errorf("E2 Connection Update response failed: %v", err)
			}
		}*/
		log.Debugf("Response PDU is following\n%v", responsePDU)
		err := p.dispatcher(responsePDU)
		if err != nil {
			if err == context.Canceled || err == context.DeadlineExceeded || err == syscall.EPIPE {
				log.Warnf("E2 Connection Update response failed: %v", err)
				return
			}
			log.Errorf("E2 Connection Update response failed: %v", err)
		}
	} else if failure != nil {
		responsePDU := &e2appdudescriptions.E2ApPdu{
			E2ApPdu: &e2appdudescriptions.E2ApPdu_UnsuccessfulOutcome{
				UnsuccessfulOutcome: &e2appdudescriptions.UnsuccessfulOutcome{
					ProcedureCode: int32(v2.ProcedureCodeIDE2connectionUpdate),
					Criticality:   e2ap_commondatatypes.Criticality_CRITICALITY_REJECT,
					Value: &e2appdudescriptions.UnsuccessfulOutcomeE2ApElementaryProcedures{
						UoValues: &e2appdudescriptions.UnsuccessfulOutcomeE2ApElementaryProcedures_E2ConnectionUpdate{
							E2ConnectionUpdate: failure,
						},
					},
				},
			},
		}
		// TODO enable validation when it is supported
		/*if err := requestPDU.Validate(); err != nil {
			log.Errorf("E2 Connection Update response validation failed: %v", err)
		} else {
			err := p.dispatcher(responsePDU)
			if err != nil {
				log.Errorf("E2 Connection Update response failed: %v", err)
			}
		}*/
		err := p.dispatcher(responsePDU)
		if err != nil {
			if err == context.Canceled || err == context.DeadlineExceeded || err == syscall.EPIPE {
				log.Warnf("E2 Connection Update response failed: %v", err)
				return
			}
			log.Errorf("E2 Connection Update response failed: %v", err)
		}

	} else {
		log.Errorf("E2 Connection Update function returned invalid output: no response message found")
	}
}

func (p *E2ConnectionUpdateProcedure) Close() error {
	return nil
}

var _ ElementaryProcedure = &E2ConnectionUpdateProcedure{}
