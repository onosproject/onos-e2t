// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package encoder

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2t/api/xnap/v1/choiceOptions"
	xnappdudescriptionsv1 "github.com/onosproject/onos-e2t/api/xnap/v1/xnap-pdu-descriptions"

	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	"github.com/onosproject/onos-lib-go/pkg/logging"
)

var log = logging.GetLogger()

func PerEncodeXnApPdu(xnap *xnappdudescriptionsv1.XnApPDu) ([]byte, error) {

	log.Debugf("Obtained XnAP-PDU message is\n%v", xnap)

	per, err := aper.MarshalWithParams(xnap, "choiceExt", choiceOptions.XnapChoicemap, choiceOptions.XnapCanonicalChoicemap)
	if err != nil {
		return nil, err
	}
	log.Debugf("Encoded XnAP-PDU PER bytes are\n%v", hex.Dump(per))

	return per, nil
}

func PerDecodexnapPdu(per []byte) (*xnappdudescriptionsv1.XnApPDu, error) {

	log.Debugf("Obtained XnAP-PDU PER bytes are\n%v", hex.Dump(per))

	result := xnappdudescriptionsv1.XnApPDu{}
	err := aper.UnmarshalWithParams(per, &result, "choiceExt", choiceOptions.XnapChoicemap, choiceOptions.XnapCanonicalChoicemap)
	if err != nil {
		return nil, err
	}

	log.Debugf("Decoded XnAP-PDU from PER is\n%v", &result)

	return &result, nil
}
