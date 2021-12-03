// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package encoder

import (
	"encoding/hex"
	"github.com/google/martian/log"
	e2appdudescriptions "github.com/onosproject/onos-e2t/api/e2ap_go/v2/e2ap-pdu-descriptions"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
)

func init() {
	log.SetLevel(log.Info)
}

func PerEncodeE2ApPdu(e2ap *e2appdudescriptions.E2ApPdu) ([]byte, error) {

	log.Debugf("Obtained E2AP-PDU message is\n%v", e2ap)
	aper.ChoiceMap = e2appdudescriptions.E2ApPduChoicemap
	per, err := aper.MarshalWithParams(e2ap, "valueExt")
	if err != nil {
		return nil, err
	}
	log.Debugf("Encoded E2AP-PDU PER bytes are\n%v", hex.Dump(per))

	return per, nil
}

func PerDecodeE2ApPdu(per []byte) (*e2appdudescriptions.E2ApPdu, error) {

	log.Debugf("Obtained E2AP-PDU PER bytes are\n%v", hex.Dump(per))
	aper.ChoiceMap = e2appdudescriptions.E2ApPduChoicemap
	result := e2appdudescriptions.E2ApPdu{}
	err := aper.UnmarshalWithParams(per, &result, "valueExt")
	if err != nil {
		return nil, err
	}

	log.Debugf("Decoded E2AP-PDU from PER is\n%v", &result)

	return &result, nil
}
