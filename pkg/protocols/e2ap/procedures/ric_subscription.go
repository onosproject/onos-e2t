// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package procedures

import (
	"context"
	v2 "github.com/onosproject/onos-e2t/api/e2ap/v2"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap_go/v2/e2ap-commondatatypes"
	"sync"
	"syscall"

	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap_go/v2/e2ap-pdu-contents"
	e2appdudescriptions "github.com/onosproject/onos-e2t/api/e2ap_go/v2/e2ap-pdu-descriptions"
	"github.com/onosproject/onos-lib-go/pkg/errors"
)

// RICSubscription is a RIC subscription procedure
type RICSubscription interface {
	RICSubscription(ctx context.Context, request *e2appducontents.RicsubscriptionRequest) (response *e2appducontents.RicsubscriptionResponse, failure *e2appducontents.RicsubscriptionFailure, err error)
}

func NewRICSubscriptionInitiator(dispatcher Dispatcher) *RICSubscriptionInitiator {
	return &RICSubscriptionInitiator{
		dispatcher:  dispatcher,
		responseChs: make(map[int32]chan e2appdudescriptions.E2ApPdu),
	}
}

type RICSubscriptionInitiator struct {
	dispatcher  Dispatcher
	responseChs map[int32]chan e2appdudescriptions.E2ApPdu
	mu          sync.RWMutex
}

func (p *RICSubscriptionInitiator) Initiate(ctx context.Context, request *e2appducontents.RicsubscriptionRequest) (*e2appducontents.RicsubscriptionResponse, *e2appducontents.RicsubscriptionFailure, error) {
	requestPDU := &e2appdudescriptions.E2ApPdu{
		E2ApPdu: &e2appdudescriptions.E2ApPdu_InitiatingMessage{
			InitiatingMessage: &e2appdudescriptions.InitiatingMessage{
				ProcedureCode: int32(v2.ProcedureCodeIDRICsubscription),
				Criticality:   e2ap_commondatatypes.Criticality_CRITICALITY_REJECT,
				Value: &e2appdudescriptions.InitiatingMessageE2ApElementaryProcedures{
					ImValues: &e2appdudescriptions.InitiatingMessageE2ApElementaryProcedures_RicSubscription{
						RicSubscription: request,
					},
				},
			},
		},
	}
	// TODO enable it when it is supported
	/*if err := requestPDU.Validate(); err != nil {
		return nil, nil, errors.NewInvalid("E2AP PDU validation failed: %v", err)
	}*/

	responseCh := make(chan e2appdudescriptions.E2ApPdu, 1)
	requestID := request.ProtocolIes.E2ApProtocolIes29.Value.RicRequestorId
	p.mu.Lock()
	p.responseChs[requestID] = responseCh
	p.mu.Unlock()

	defer func() {
		p.mu.Lock()
		delete(p.responseChs, requestID)
		p.mu.Unlock()
	}()

	if err := p.dispatcher(requestPDU); err != nil {
		return nil, nil, errors.NewUnavailable("RIC Subscription initiation failed: %v", err)
	}

	select {
	case responsePDU, ok := <-responseCh:
		if !ok {
			return nil, nil, errors.NewUnavailable("connection closed")
		}

		switch response := responsePDU.E2ApPdu.(type) {
		case *e2appdudescriptions.E2ApPdu_SuccessfulOutcome:
			return response.SuccessfulOutcome.Value.GetRicSubscription(), nil, nil
		case *e2appdudescriptions.E2ApPdu_UnsuccessfulOutcome:
			return nil, response.UnsuccessfulOutcome.Value.GetRicSubscription(), nil
		default:
			return nil, nil, errors.NewInternal("received unexpected outcome")
		}
	case <-ctx.Done():
		return nil, nil, ctx.Err()
	}
}

func (p *RICSubscriptionInitiator) Matches(pdu *e2appdudescriptions.E2ApPdu) bool {
	switch msg := pdu.E2ApPdu.(type) {
	case *e2appdudescriptions.E2ApPdu_SuccessfulOutcome:
		return msg.SuccessfulOutcome.Value.GetRicSubscription() != nil
	case *e2appdudescriptions.E2ApPdu_UnsuccessfulOutcome:
		return msg.UnsuccessfulOutcome.Value.GetRicSubscription() != nil
	default:
		return false
	}
}

func (p *RICSubscriptionInitiator) Handle(pdu *e2appdudescriptions.E2ApPdu) {
	var requestID int32
	switch response := pdu.E2ApPdu.(type) {
	case *e2appdudescriptions.E2ApPdu_SuccessfulOutcome:
		requestID = response.SuccessfulOutcome.Value.GetRicSubscription().ProtocolIes.E2ApProtocolIes29.Value.RicRequestorId
	case *e2appdudescriptions.E2ApPdu_UnsuccessfulOutcome:
		requestID = response.UnsuccessfulOutcome.Value.GetRicSubscription().ProtocolIes.E2ApProtocolIes29.Value.RicRequestorId
	}

	p.mu.RLock()
	responseCh, ok := p.responseChs[requestID]
	p.mu.RUnlock()
	if ok {
		responseCh <- *pdu
		close(responseCh)
	} else {
		log.Errorf("Received RIC Subscription response for unknown request %d", requestID)
	}
}

func (p *RICSubscriptionInitiator) Close() error {
	p.mu.Lock()
	for _, responseCh := range p.responseChs {
		close(responseCh)
	}
	p.mu.Unlock()
	return nil
}

var _ ElementaryProcedure = &RICSubscriptionInitiator{}

func NewRICSubscriptionProcedure(dispatcher Dispatcher, handler RICSubscription) *RICSubscriptionProcedure {
	return &RICSubscriptionProcedure{
		dispatcher: dispatcher,
		handler:    handler,
	}
}

type RICSubscriptionProcedure struct {
	dispatcher Dispatcher
	handler    RICSubscription
}

func (p *RICSubscriptionProcedure) Matches(pdu *e2appdudescriptions.E2ApPdu) bool {
	switch msg := pdu.E2ApPdu.(type) {
	case *e2appdudescriptions.E2ApPdu_InitiatingMessage:
		return msg.InitiatingMessage.Value.GetRicSubscription() != nil
	default:
		return false
	}
}

func (p *RICSubscriptionProcedure) Handle(requestPDU *e2appdudescriptions.E2ApPdu) {
	response, failure, err := p.handler.RICSubscription(context.Background(), requestPDU.GetInitiatingMessage().GetValue().GetRicSubscription())
	if err != nil {
		log.Errorf("RIC Subscription procedure failed: %v", err)
	} else if response != nil {
		responsePDU := &e2appdudescriptions.E2ApPdu{
			E2ApPdu: &e2appdudescriptions.E2ApPdu_SuccessfulOutcome{
				SuccessfulOutcome: &e2appdudescriptions.SuccessfulOutcome{
					ProcedureCode: int32(v2.ProcedureCodeIDRICsubscription),
					Criticality:   e2ap_commondatatypes.Criticality_CRITICALITY_REJECT,
					Value: &e2appdudescriptions.SuccessfulOutcomeE2ApElementaryProcedures{
						SoValues: &e2appdudescriptions.SuccessfulOutcomeE2ApElementaryProcedures_RicSubscription{
							RicSubscription: response,
						},
					},
				},
			},
		}
		// TODO enable validation when it is supported
		/*if err := requestPDU.Validate(); err != nil {
			log.Errorf("RIC Subscription response validation failed: %v", err)
		} else {
			err := p.dispatcher(responsePDU)
			if err != nil {
				log.Errorf("RIC Subscription response failed: %v", err)
			}
		}*/
		err := p.dispatcher(responsePDU)
		if err != nil {
			if err == context.Canceled || err == context.DeadlineExceeded || err == syscall.EPIPE {
				log.Warnf("RIC Subscription response failed: %v", err)
				return
			}
			log.Errorf("RIC Subscription response failed: %v", err)
		}

	} else if failure != nil {
		responsePDU := &e2appdudescriptions.E2ApPdu{
			E2ApPdu: &e2appdudescriptions.E2ApPdu_UnsuccessfulOutcome{
				UnsuccessfulOutcome: &e2appdudescriptions.UnsuccessfulOutcome{
					ProcedureCode: int32(v2.ProcedureCodeIDRICsubscription),
					Criticality:   e2ap_commondatatypes.Criticality_CRITICALITY_REJECT,
					Value: &e2appdudescriptions.UnsuccessfulOutcomeE2ApElementaryProcedures{
						UoValues: &e2appdudescriptions.UnsuccessfulOutcomeE2ApElementaryProcedures_RicSubscription{
							RicSubscription: failure,
						},
					},
				},
			},
		}
		// TODO enable validation when it is supported
		/*if err := requestPDU.Validate(); err != nil {
			log.Errorf("RIC Subscription response validation failed: %v", err)
		} else {
			err := p.dispatcher(responsePDU)
			if err != nil {
				log.Errorf("RIC Subscription response failed: %v", err)
			}
		}*/
		err := p.dispatcher(responsePDU)
		if err != nil {
			if err == context.Canceled || err == context.DeadlineExceeded || err == syscall.EPIPE {
				log.Warnf("RIC Subscription response failed: %v", err)
				return
			}
			log.Errorf("RIC Subscription response failed: %v", err)
		}

	} else {
		log.Errorf("RIC Subscription function returned invalid output: no response message found")
	}
}

func (p *RICSubscriptionProcedure) Close() error {
	return nil
}

var _ ElementaryProcedure = &RICSubscriptionProcedure{}
