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

// RICSubscriptionDelete is a RIC subscription delete procedure
type RICSubscriptionDelete interface {
	RICSubscriptionDelete(ctx context.Context, request *e2appducontents.RicsubscriptionDeleteRequest) (response *e2appducontents.RicsubscriptionDeleteResponse, failure *e2appducontents.RicsubscriptionDeleteFailure, err error)
}

func NewRICSubscriptionDeleteInitiator(dispatcher Dispatcher) *RICSubscriptionDeleteInitiator {
	return &RICSubscriptionDeleteInitiator{
		dispatcher:  dispatcher,
		responseChs: make(map[int32]chan e2appdudescriptions.E2ApPdu),
		closeCh:     make(chan struct{}),
	}
}

type RICSubscriptionDeleteInitiator struct {
	dispatcher  Dispatcher
	responseChs map[int32]chan e2appdudescriptions.E2ApPdu
	closeCh     chan struct{}
	mu          sync.RWMutex
}

func (p *RICSubscriptionDeleteInitiator) Initiate(ctx context.Context, request *e2appducontents.RicsubscriptionDeleteRequest) (*e2appducontents.RicsubscriptionDeleteResponse, *e2appducontents.RicsubscriptionDeleteFailure, error) {
	requestPDU := &e2appdudescriptions.E2ApPdu{
		E2ApPdu: &e2appdudescriptions.E2ApPdu_InitiatingMessage{
			InitiatingMessage: &e2appdudescriptions.InitiatingMessage{
				ProcedureCode: int32(e2api.ProcedureCodeIDRICsubscriptionDelete),
				Criticality:   e2apcommondatatypes.Criticality_CRITICALITY_REJECT,
				Value: &e2appdudescriptions.InitiatingMessageE2ApElementaryProcedures{
					ImValues: &e2appdudescriptions.InitiatingMessageE2ApElementaryProcedures_RicSubscriptionDelete{
						RicSubscriptionDelete: request,
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
			requestID = v.GetValue().GetRrId().GetRicRequestorId()
			break
		}
	}
	p.mu.Lock()
	p.responseChs[requestID] = responseCh
	p.mu.Unlock()

	if err := p.dispatcher(requestPDU); err != nil {
		return nil, nil, errors.NewUnavailable("RIC Subscription Delete initiation failed: %v", err)
	}

	select {
	case responsePDU, ok := <-responseCh:
		if !ok {
			return nil, nil, errors.NewUnavailable("connection closed")
		}

		switch response := responsePDU.E2ApPdu.(type) {
		case *e2appdudescriptions.E2ApPdu_SuccessfulOutcome:
			//return response.SuccessfulOutcome.Value.GetRicSubscriptionDelete(), nil, nil
			switch ret := response.SuccessfulOutcome.Value.SoValues.(type) {
			case *e2appdudescriptions.SuccessfulOutcomeE2ApElementaryProcedures_RicSubscriptionDelete:
				return ret.RicSubscriptionDelete, nil, nil
			default:
				return nil, nil, errors.NewInternal("received unexpected outcome")
			}
		case *e2appdudescriptions.E2ApPdu_UnsuccessfulOutcome:
			//return nil, response.UnsuccessfulOutcome.Value.GetRicSubscriptionDelete(), nil
			switch ret := response.UnsuccessfulOutcome.Value.UoValues.(type) {
			case *e2appdudescriptions.UnsuccessfulOutcomeE2ApElementaryProcedures_RicSubscriptionDelete:
				return nil, ret.RicSubscriptionDelete, nil
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

func (p *RICSubscriptionDeleteInitiator) Matches(pdu *e2appdudescriptions.E2ApPdu) bool {
	switch msg := pdu.E2ApPdu.(type) {
	case *e2appdudescriptions.E2ApPdu_SuccessfulOutcome:
		switch ret := msg.SuccessfulOutcome.Value.SoValues.(type) {
		case *e2appdudescriptions.SuccessfulOutcomeE2ApElementaryProcedures_RicSubscriptionDelete:
			return ret.RicSubscriptionDelete != nil
		default:
			return false
		}
	case *e2appdudescriptions.E2ApPdu_UnsuccessfulOutcome:
		switch ret := msg.UnsuccessfulOutcome.Value.UoValues.(type) {
		case *e2appdudescriptions.UnsuccessfulOutcomeE2ApElementaryProcedures_RicSubscriptionDelete:
			return ret.RicSubscriptionDelete != nil
		default:
			return false
		}
	default:
		return false
	}
}

func (p *RICSubscriptionDeleteInitiator) Handle(pdu *e2appdudescriptions.E2ApPdu) {
	var requestID int32
	switch response := pdu.E2ApPdu.(type) {
	case *e2appdudescriptions.E2ApPdu_SuccessfulOutcome:
		for _, v := range response.SuccessfulOutcome.Value.GetRicSubscriptionDelete().GetProtocolIes() {
			if v.Id == int32(e2api.ProtocolIeIDRicrequestID) {
				requestID = v.GetValue().GetRrId().GetRicRequestorId()
				break
			}
		}
	case *e2appdudescriptions.E2ApPdu_UnsuccessfulOutcome:
		for _, v := range response.UnsuccessfulOutcome.Value.GetRicSubscriptionDelete().GetProtocolIes() {
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
		log.Warnf("Received RIC Subscription Delete response for unknown request %d", requestID)
	}
}

func (p *RICSubscriptionDeleteInitiator) Close() error {
	close(p.closeCh)
	return nil
}

var _ ElementaryProcedure = &RICSubscriptionDeleteInitiator{}

func NewRICSubscriptionDeleteProcedure(dispatcher Dispatcher, handler RICSubscriptionDelete) *RICSubscriptionDeleteProcedure {
	return &RICSubscriptionDeleteProcedure{
		dispatcher: dispatcher,
		handler:    handler,
	}
}

type RICSubscriptionDeleteProcedure struct {
	dispatcher Dispatcher
	handler    RICSubscriptionDelete
}

func (p *RICSubscriptionDeleteProcedure) Matches(pdu *e2appdudescriptions.E2ApPdu) bool {
	switch msg := pdu.E2ApPdu.(type) {
	case *e2appdudescriptions.E2ApPdu_InitiatingMessage:
		switch ret := msg.InitiatingMessage.Value.ImValues.(type) {
		case *e2appdudescriptions.InitiatingMessageE2ApElementaryProcedures_RicSubscriptionDelete:
			return ret.RicSubscriptionDelete != nil
		default:
			return false
		}
	default:
		return false
	}
}

func (p *RICSubscriptionDeleteProcedure) Handle(requestPDU *e2appdudescriptions.E2ApPdu) {
	response, failure, err := p.handler.RICSubscriptionDelete(context.Background(), requestPDU.GetInitiatingMessage().GetValue().GetRicSubscriptionDelete())
	if err != nil {
		log.Errorf("RIC Subscription Delete procedure failed: %v", err)
	} else if response != nil {
		responsePDU := &e2appdudescriptions.E2ApPdu{
			E2ApPdu: &e2appdudescriptions.E2ApPdu_SuccessfulOutcome{
				SuccessfulOutcome: &e2appdudescriptions.SuccessfulOutcome{
					ProcedureCode: int32(e2api.ProcedureCodeIDRICsubscriptionDelete),
					Criticality:   e2apcommondatatypes.Criticality_CRITICALITY_REJECT,
					Value: &e2appdudescriptions.SuccessfulOutcomeE2ApElementaryProcedures{
						SoValues: &e2appdudescriptions.SuccessfulOutcomeE2ApElementaryProcedures_RicSubscriptionDelete{
							RicSubscriptionDelete: response,
						},
					},
				},
			},
		}
		if err := requestPDU.Validate(); err != nil {
			log.Errorf("RIC Subscription Delete response validation failed: %v", err)
		} else {
			err := p.dispatcher(responsePDU)
			if err != nil {
				if err == context.Canceled || err == context.DeadlineExceeded || err == syscall.EPIPE || err == syscall.EBADF {
					log.Warnf("RIC Subscription Delete response failed: %v", err)
					return
				}
				log.Errorf("RIC Subscription Delete response failed: %v", err)
			}
		}

	} else if failure != nil {
		responsePDU := &e2appdudescriptions.E2ApPdu{
			E2ApPdu: &e2appdudescriptions.E2ApPdu_UnsuccessfulOutcome{
				UnsuccessfulOutcome: &e2appdudescriptions.UnsuccessfulOutcome{
					ProcedureCode: int32(e2api.ProcedureCodeIDRICsubscriptionDelete),
					Criticality:   e2apcommondatatypes.Criticality_CRITICALITY_REJECT,
					Value: &e2appdudescriptions.UnsuccessfulOutcomeE2ApElementaryProcedures{
						UoValues: &e2appdudescriptions.UnsuccessfulOutcomeE2ApElementaryProcedures_RicSubscriptionDelete{
							RicSubscriptionDelete: failure,
						},
					},
				},
			},
		}
		if err := requestPDU.Validate(); err != nil {
			log.Errorf("RIC Subscription Delete response validation failed: %v", err)
		} else {
			err := p.dispatcher(responsePDU)
			if err != nil {
				log.Errorf("RIC Subscription Delete response failed: %v", err)
			}
		}
		err := p.dispatcher(responsePDU)
		if err != nil {
			if err == context.Canceled || err == context.DeadlineExceeded || err == syscall.EPIPE || err == syscall.EBADF {
				log.Warnf("RIC Subscription Delete response failed: %v", err)
				return
			}
			log.Errorf("RIC Subscription Delete response faile: %v", err)
		}

	} else {
		log.Errorf("RIC Subscription Delete function returned invalid output: no response message found")
	}
}

func (p *RICSubscriptionDeleteProcedure) Close() error {
	return nil
}

var _ ElementaryProcedure = &RICSubscriptionDeleteProcedure{}
