// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package utils

import (
	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre_go/pdubuilder"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"google.golang.org/protobuf/proto"
)

// Control control request fields
type Control struct {
	Header  []byte
	Payload []byte
}

// RcControlHeader required fields for creating RC service model control header
type RcControlHeader struct {
	Priority int32
	PlmnID   []byte
	CellID   []byte
}

// RcControlMessage required fields for creating RC service model control message
type RcControlMessage struct {
	RanParameterID    int32
	RanParameterName  string
	RanParameterValue int64
}

// CreateRcControlHeader  creates rc control header
func (ch *RcControlHeader) CreateRcControlHeader() ([]byte, error) {

	cellID := &asn1.BitString{
		Value: ch.CellID,
		Len:   36,
	}
	cgi, err := pdubuilder.CreateCellGlobalIDNrCgi(ch.PlmnID, cellID)
	if err != nil {
		return []byte{}, err
	}

	newE2SmRcPrePdu, err := pdubuilder.CreateE2SmRcPreControlHeader()
	if err != nil {
		return []byte{}, err
	}
	newE2SmRcPrePdu.GetControlHeaderFormat1().SetCGI(cgi)
	newE2SmRcPrePdu.GetControlHeaderFormat1().SetRicControlMessagePriority(ch.Priority)

	err = newE2SmRcPrePdu.Validate()
	if err != nil {
		return []byte{}, err
	}

	protoBytes, err := proto.Marshal(newE2SmRcPrePdu)
	if err != nil {
		return []byte{}, err
	}

	return protoBytes, nil
}

// CreateRcControlMessage creates rc control message
func (cm *RcControlMessage) CreateRcControlMessage() ([]byte, error) {
	ranParameterValue, err := pdubuilder.CreateRanParameterValueInt(cm.RanParameterValue)
	if err != nil {
		return nil, err
	}
	newE2SmRcPrePdu, err := pdubuilder.CreateE2SmRcPreControlMessage(cm.RanParameterID, cm.RanParameterName, ranParameterValue)
	if err != nil {
		return nil, err
	}
	err = newE2SmRcPrePdu.Validate()
	if err != nil {
		return nil, err
	}

	protoBytes, err := proto.Marshal(newE2SmRcPrePdu)
	if err != nil {
		return nil, err
	}

	return protoBytes, nil
}

// Create creates a control request using SDK
func (cr *Control) Create() (*e2api.ControlMessage, error) {
	request := &e2api.ControlMessage{
		Header:  cr.Header,
		Payload: cr.Payload,
	}

	return request, nil

}
