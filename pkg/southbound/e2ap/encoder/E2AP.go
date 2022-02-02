// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package encoder

import (
	"encoding/hex"
	"github.com/google/martian/log"
	e2appdudescriptions "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-descriptions"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/goaperlib"

	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
)

func init() {
	log.SetLevel(log.Debug)
}

func PerEncodeE2ApPdu(e2ap *e2appdudescriptions.E2ApPdu) ([]byte, error) {

	log.Debugf("Obtained E2AP-PDU message is\n%v", e2ap)
	//choicemap = e2appdudescriptions.E2ApPduChoicemap
	//canonicalChoiceMap = e2appdudescriptions.E2ApPduCanonicalChoicemap

	log.Debugf("CHOICE map is following:\n%v", aper.ChoiceMap)
	log.Debugf("Canonical CHOICE map is following:\n%v", aper.CanonicalChoiceMap)

	per, err := goaperlib.MarshalWithParams(e2ap, "choiceExt")
	if err != nil {
		return nil, err
	}
	log.Debugf("Encoded E2AP-PDU PER bytes are\n%v", hex.Dump(per))

	return per, nil
}

func PerDecodeE2ApPdu(per []byte) (*e2appdudescriptions.E2ApPdu, error) {

	log.Debugf("Obtained E2AP-PDU PER bytes are\n%v", hex.Dump(per))
	//choicemap = e2appdudescriptions.E2ApPduChoicemap
	//canonicalChoiceMap = e2appdudescriptions.E2ApPduCanonicalChoicemap

	log.Debugf("CHOICE map is following:\n%v", aper.ChoiceMap)
	log.Debugf("Canonical CHOICE map is following:\n%v", aper.CanonicalChoiceMap)

	result := e2appdudescriptions.E2ApPdu{}
	err := goaperlib.UnmarshalWithParams(per, &result, "choiceExt")
	if err != nil {
		return nil, err
	}

	log.Debugf("Decoded E2AP-PDU from PER is\n%v", &result)

	return &result, nil
}
