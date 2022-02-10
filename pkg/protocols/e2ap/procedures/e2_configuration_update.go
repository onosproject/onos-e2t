// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package procedures

import (
	"context"
	v2 "github.com/onosproject/onos-e2t/api/e2ap/v2"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-commondatatypes"
	"syscall"

	"github.com/onosproject/onos-lib-go/pkg/errors"

	e2appdudescriptions "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-descriptions"

	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
)

// E2ConfigurationUpdate is an E2 configuration procedure
type E2ConfigurationUpdate interface {
	E2ConfigurationUpdate(ctx context.Context, request *e2appducontents.E2NodeConfigurationUpdate) (response *e2appducontents.E2NodeConfigurationUpdateAcknowledge, failure *e2appducontents.E2NodeConfigurationUpdateFailure, err error)
}

// NewConfigurationUpdateInitiator creates a new E2 configuration update initiator
func NewE2ConfigurationUpdateInitiator(dispatcher Dispatcher) *E2ConfigurationUpdateInitiator {
	return &E2ConfigurationUpdateInitiator{
		dispatcher: dispatcher,
		responseCh: make(chan e2appdudescriptions.E2ApPdu),
	}
}

// E2ConfigurationUpdateInitiator initiates the E2 configuration update procedure procedure
type E2ConfigurationUpdateInitiator struct {
	dispatcher Dispatcher
	responseCh chan e2appdudescriptions.E2ApPdu
}

func (p *E2ConfigurationUpdateInitiator) Initiate(ctx context.Context, request *e2appducontents.E2NodeConfigurationUpdate) (*e2appducontents.E2NodeConfigurationUpdateAcknowledge, *e2appducontents.E2NodeConfigurationUpdateFailure, error) {
	requestPDU := &e2appdudescriptions.E2ApPdu{
		E2ApPdu: &e2appdudescriptions.E2ApPdu_InitiatingMessage{
			InitiatingMessage: &e2appdudescriptions.InitiatingMessage{
				ProcedureCode: int32(v2.ProcedureCodeIDE2nodeConfigurationUpdate),
				Criticality:   e2ap_commondatatypes.Criticality_CRITICALITY_REJECT,
				Value: &e2appdudescriptions.InitiatingMessageE2ApElementaryProcedures{
					ImValues: &e2appdudescriptions.InitiatingMessageE2ApElementaryProcedures_E2NodeConfigurationUpdate{
						E2NodeConfigurationUpdate: request,
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
		return nil, nil, errors.NewUnavailable("E2 configuration update initiation failed: %v", err)
	}

	select {
	case responsePDU, ok := <-p.responseCh:
		if !ok {
			return nil, nil, errors.NewUnavailable("connection closed")
		}

		switch msg := responsePDU.E2ApPdu.(type) {
		case *e2appdudescriptions.E2ApPdu_SuccessfulOutcome:
			//return msg.SuccessfulOutcome.Value.GetE2NodeConfigurationUpdate(), nil, nil
			switch ret := msg.SuccessfulOutcome.Value.SoValues.(type) {
			case *e2appdudescriptions.SuccessfulOutcomeE2ApElementaryProcedures_E2NodeConfigurationUpdate:
				return ret.E2NodeConfigurationUpdate, nil, nil
			default:
				return nil, nil, errors.NewInternal("received unexpected outcome")
			}
		case *e2appdudescriptions.E2ApPdu_UnsuccessfulOutcome:
			//return nil, msg.UnsuccessfulOutcome.Value.GetE2NodeConfigurationUpdate(), nil
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
	}
}

func (p *E2ConfigurationUpdateInitiator) Matches(pdu *e2appdudescriptions.E2ApPdu) bool {
	switch msg := pdu.E2ApPdu.(type) {
	case *e2appdudescriptions.E2ApPdu_SuccessfulOutcome:
		//return msg.SuccessfulOutcome.Value.GetE2NodeConfigurationUpdate() != nil
		switch ret := msg.SuccessfulOutcome.Value.SoValues.(type) {
		case *e2appdudescriptions.SuccessfulOutcomeE2ApElementaryProcedures_E2NodeConfigurationUpdate:
			return ret.E2NodeConfigurationUpdate != nil
		default:
			return false
		}
	case *e2appdudescriptions.E2ApPdu_UnsuccessfulOutcome:
		//return msg.UnsuccessfulOutcome.Value.GetE2NodeConfigurationUpdate() != nil
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
	p.responseCh <- *pdu
}

func (p *E2ConfigurationUpdateInitiator) Close() error {
	defer func() {
		if err := recover(); err != nil {
			log.Debug("recovering from panic", err)
		}
	}()
	close(p.responseCh)
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

// E2ConfigurationUpdate implements the E2 configuration update procedure
type E2ConfigurationUpdateProcedure struct {
	dispatcher Dispatcher
	handler    E2ConfigurationUpdate
}

func (p *E2ConfigurationUpdateProcedure) Matches(pdu *e2appdudescriptions.E2ApPdu) bool {
	switch msg := pdu.E2ApPdu.(type) {
	case *e2appdudescriptions.E2ApPdu_InitiatingMessage:
		//return msg.InitiatingMessage.Value.GetE2NodeConfigurationUpdate() != nil
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
					ProcedureCode: int32(v2.ProcedureCodeIDE2nodeConfigurationUpdate),
					Criticality:   e2ap_commondatatypes.Criticality_CRITICALITY_REJECT,
					Value: &e2appdudescriptions.SuccessfulOutcomeE2ApElementaryProcedures{
						SoValues: &e2appdudescriptions.SuccessfulOutcomeE2ApElementaryProcedures_E2NodeConfigurationUpdate{
							E2NodeConfigurationUpdate: response,
						},
					},
				},
			},
		}
		// TODO enable validation when it is supported
		/*if err := requestPDU.Validate(); err != nil {
			log.Errorf("E2 configuration update validation failed: %v", err)
		} else {
			err := p.dispatcher(responsePDU)
			if err != nil {
				log.Errorf("E2 configuration update response failed: %v", err)
			}
		}*/
		err := p.dispatcher(responsePDU)
		if err != nil {
			if err == context.Canceled || err == context.DeadlineExceeded || err == syscall.EPIPE {
				log.Warnf("E2 configuration update response failed: %v", err)
				return
			}
			log.Errorf("E2 configuration update response failed: %v", err)
		}

	} else if failure != nil {
		responsePDU := &e2appdudescriptions.E2ApPdu{
			E2ApPdu: &e2appdudescriptions.E2ApPdu_UnsuccessfulOutcome{
				UnsuccessfulOutcome: &e2appdudescriptions.UnsuccessfulOutcome{
					ProcedureCode: int32(v2.ProcedureCodeIDE2nodeConfigurationUpdate),
					Criticality:   e2ap_commondatatypes.Criticality_CRITICALITY_REJECT,
					Value: &e2appdudescriptions.UnsuccessfulOutcomeE2ApElementaryProcedures{
						UoValues: &e2appdudescriptions.UnsuccessfulOutcomeE2ApElementaryProcedures_E2NodeConfigurationUpdate{
							E2NodeConfigurationUpdate: failure,
						},
					},
				},
			},
		}
		// TODO enable validation when it is supported
		/*if err := requestPDU.Validate(); err != nil {
			log.Errorf("E2 configuration update validation failed: %v", err)
		} else {
			err := p.dispatcher(responsePDU)
			if err != nil {
				log.Errorf("E2 configuration update failed: %v", err)
			}
		}*/
		err := p.dispatcher(responsePDU)
		if err != nil {
			if err == context.Canceled || err == context.DeadlineExceeded || err == syscall.EPIPE {
				log.Warnf("E2 configuration update response failed: %v", err)
				return
			}
			log.Errorf("E2 configuration update response failed: %v", err)

		}

	} else {
		log.Errorf("E2 configuration update function returned invalid output: no response message found")
	}
}

func (p *E2ConfigurationUpdateProcedure) Close() error {
	return nil
}

var _ ElementaryProcedure = &E2ConfigurationUpdateProcedure{}
