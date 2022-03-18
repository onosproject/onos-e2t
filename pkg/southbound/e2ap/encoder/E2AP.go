// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package encoder

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2t/api/e2ap/v2/choiceOptions"

	e2appdudescriptions "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-descriptions"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	"github.com/onosproject/onos-lib-go/pkg/logging"
)

var log = logging.GetLogger()

func PerEncodeE2ApPdu(e2ap *e2appdudescriptions.E2ApPdu) ([]byte, error) {

	log.Debugf("Obtained E2AP-PDU message is\n%v", e2ap)

	per, err := aper.MarshalWithParams(e2ap, "choiceExt", choiceOptions.E2ApPduChoicemap, choiceOptions.E2ApPduCanonicalChoicemap)
	if err != nil {
		return nil, err
	}
	log.Debugf("Encoded E2AP-PDU PER bytes are\n%v", hex.Dump(per))

	return per, nil
}

func PerDecodeE2ApPdu(per []byte) (*e2appdudescriptions.E2ApPdu, error) {

	log.Debugf("Obtained E2AP-PDU PER bytes are\n%v", hex.Dump(per))

	result := e2appdudescriptions.E2ApPdu{}
	err := aper.UnmarshalWithParams(per, &result, "choiceExt", choiceOptions.E2ApPduChoicemap, choiceOptions.E2ApPduCanonicalChoicemap)
	if err != nil {
		return nil, err
	}

	log.Debugf("Decoded E2AP-PDU from PER is\n%v", &result)

	return &result, nil
}
