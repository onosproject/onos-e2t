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

// RICSubscription is a RIC subscription procedure
type RICSubscription interface {
	RICSubscription(ctx context.Context, request *e2appducontents.RicsubscriptionRequest) (response *e2appducontents.RicsubscriptionResponse, failure *e2appducontents.RicsubscriptionFailure, err error)
}

func NewRICSubscriptionInitiator(dispatcher Dispatcher) *RICSubscriptionInitiator {
	return &RICSubscriptionInitiator{
		dispatcher:  dispatcher,
		responseChs: make(map[int32]chan e2appdudescriptions.E2ApPdu),
		closeCh:     make(chan struct{}),
	}
}

type RICSubscriptionInitiator struct {
	dispatcher  Dispatcher
	responseChs map[int32]chan e2appdudescriptions.E2ApPdu
	closeCh     chan struct{}
	mu          sync.RWMutex
}

func (p *RICSubscriptionInitiator) Initiate(ctx context.Context, request *e2appducontents.RicsubscriptionRequest) (*e2appducontents.RicsubscriptionResponse, *e2appducontents.RicsubscriptionFailure, error) {
	requestPDU := &e2appdudescriptions.E2ApPdu{
		E2ApPdu: &e2appdudescriptions.E2ApPdu_InitiatingMessage{
			InitiatingMessage: &e2appdudescriptions.InitiatingMessage{
				ProcedureCode: int32(e2api.ProcedureCodeIDRICsubscription),
				Criticality:   e2apcommondatatypes.Criticality_CRITICALITY_REJECT,
				Value: &e2appdudescriptions.InitiatingMessageE2ApElementaryProcedures{
					ImValues: &e2appdudescriptions.InitiatingMessageE2ApElementaryProcedures_RicSubscription{
						RicSubscription: request,
					},
				},
			},
		},
	}
	if err := requestPDU.Validate(); err != nil {
		return nil, nil, errors.NewInvalid("E2AP PDU validation failed: %v", err)
	}

	responseCh := make(chan e2appdudescriptions.E2ApPdu, 1)
	var requestID int32 = -1
	for _, v := range request.GetProtocolIes() {
		if v.Id == int32(e2api.ProtocolIeIDRicrequestID) {
			requestID = v.GetValue().GetRrId().GetRicRequestorId()
			break
		}
	}
	p.mu.Lock()
	p.responseChs[requestID] = responseCh
	p.mu.Unlock()

	if err := p.dispatcher(requestPDU); err != nil {
		return nil, nil, errors.NewUnavailable("RIC Subscription initiation failed: %v", err)
	}

	select {
	case responsePDU, ok := <-responseCh:
		if !ok {
			return nil, nil, errors.NewUnavailable("connection closed")
		}

		switch msg := responsePDU.E2ApPdu.(type) {
		case *e2appdudescriptions.E2ApPdu_SuccessfulOutcome:
			//return msg.SuccessfulOutcome.Value.GetRicSubscription(), nil, nil
			switch ret := msg.SuccessfulOutcome.Value.SoValues.(type) {
			case *e2appdudescriptions.SuccessfulOutcomeE2ApElementaryProcedures_RicSubscription:
				return ret.RicSubscription, nil, nil
			default:
				return nil, nil, errors.NewInternal("received unexpected outcome")
			}
		case *e2appdudescriptions.E2ApPdu_UnsuccessfulOutcome:
			//return nil, msg.UnsuccessfulOutcome.Value.GetRicSubscription(), nil
			switch ret := msg.UnsuccessfulOutcome.Value.UoValues.(type) {
			case *e2appdudescriptions.UnsuccessfulOutcomeE2ApElementaryProcedures_RicSubscription:
				return nil, ret.RicSubscription, nil
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

func (p *RICSubscriptionInitiator) Matches(pdu *e2appdudescriptions.E2ApPdu) bool {
	switch msg := pdu.E2ApPdu.(type) {
	case *e2appdudescriptions.E2ApPdu_SuccessfulOutcome:
		//return msg.SuccessfulOutcome.Value.GetRicSubscription() != nil
		switch ret := msg.SuccessfulOutcome.Value.SoValues.(type) {
		case *e2appdudescriptions.SuccessfulOutcomeE2ApElementaryProcedures_RicSubscription:
			return ret.RicSubscription != nil
		default:
			return false
		}
	case *e2appdudescriptions.E2ApPdu_UnsuccessfulOutcome:
		//return msg.UnsuccessfulOutcome.Value.GetRicSubscription() != nil
		switch ret := msg.UnsuccessfulOutcome.Value.UoValues.(type) {
		case *e2appdudescriptions.UnsuccessfulOutcomeE2ApElementaryProcedures_RicSubscription:
			return ret.RicSubscription != nil
		default:
			return false
		}
	default:
		return false
	}
}

func (p *RICSubscriptionInitiator) Handle(pdu *e2appdudescriptions.E2ApPdu) {
	var requestID int32
	switch pdu.GetE2ApPdu().(type) {
	case *e2appdudescriptions.E2ApPdu_SuccessfulOutcome:
		for _, v := range pdu.GetSuccessfulOutcome().GetValue().GetRicSubscription().GetProtocolIes() {
			if v.Id == int32(e2api.ProtocolIeIDRicrequestID) {
				requestID = v.GetValue().GetRrId().GetRicRequestorId()
				break
			}
		}
	case *e2appdudescriptions.E2ApPdu_UnsuccessfulOutcome:
		for _, v := range pdu.GetUnsuccessfulOutcome().GetValue().GetRicSubscription().GetProtocolIes() {
			if v.Id == int32(e2api.ProtocolIeIDRicrequestID) {
				requestID = v.GetValue().GetRrId().GetRicRequestorId()
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
		log.Warnf("Received RIC Subscription response for unknown request %d", requestID)
	}
}

func (p *RICSubscriptionInitiator) Close() error {
	close(p.closeCh)
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
	switch pdu.E2ApPdu.(type) {
	case *e2appdudescriptions.E2ApPdu_InitiatingMessage:
		return pdu.GetInitiatingMessage().GetValue().GetRicSubscription() != nil
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
					ProcedureCode: int32(e2api.ProcedureCodeIDRICsubscription),
					Criticality:   e2apcommondatatypes.Criticality_CRITICALITY_REJECT,
					Value: &e2appdudescriptions.SuccessfulOutcomeE2ApElementaryProcedures{
						SoValues: &e2appdudescriptions.SuccessfulOutcomeE2ApElementaryProcedures_RicSubscription{
							RicSubscription: response,
						},
					},
				},
			},
		}
		if err := requestPDU.Validate(); err != nil {
			log.Errorf("RIC Subscription response validation failed: %v", err)
		} else {

			err := p.dispatcher(responsePDU)
			if err != nil {
				if err == context.Canceled || err == context.DeadlineExceeded || err == syscall.EPIPE || err == syscall.EBADF {
					log.Warnf("RIC Subscription response failed: %v", err)
					return
				}
				log.Errorf("RIC Subscription response failed: %v", err)
			}
		}

	} else if failure != nil {
		responsePDU := &e2appdudescriptions.E2ApPdu{
			E2ApPdu: &e2appdudescriptions.E2ApPdu_UnsuccessfulOutcome{
				UnsuccessfulOutcome: &e2appdudescriptions.UnsuccessfulOutcome{
					ProcedureCode: int32(e2api.ProcedureCodeIDRICsubscription),
					Criticality:   e2apcommondatatypes.Criticality_CRITICALITY_REJECT,
					Value: &e2appdudescriptions.UnsuccessfulOutcomeE2ApElementaryProcedures{
						UoValues: &e2appdudescriptions.UnsuccessfulOutcomeE2ApElementaryProcedures_RicSubscription{
							RicSubscription: failure,
						},
					},
				},
			},
		}
		if err := requestPDU.Validate(); err != nil {
			log.Errorf("RIC Subscription response validation failed: %v", err)
		} else {
			err := p.dispatcher(responsePDU)
			if err != nil {
				if err == context.Canceled || err == context.DeadlineExceeded || err == syscall.EPIPE || err == syscall.EBADF {
					log.Warnf("RIC Subscription response failed: %v", err)
					return
				}
				log.Errorf("RIC Subscription response failed: %v", err)
			}
		}

	} else {
		log.Errorf("RIC Subscription function returned invalid output: no response message found")
	}
}

func (p *RICSubscriptionProcedure) Close() error {
	return nil
}

var _ ElementaryProcedure = &RICSubscriptionProcedure{}
