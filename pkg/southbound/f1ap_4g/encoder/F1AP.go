// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package encoder

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2t/api/f1ap_4g/v1/choiceOptions"
	f1appdudescriptionsv1 "github.com/onosproject/onos-e2t/api/f1ap_4g/v1/f1ap_pdu_descriptions"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	"github.com/onosproject/onos-lib-go/pkg/logging"
)

var log = logging.GetLogger()

func PerEncodeF1ApPdu(f1ap *f1appdudescriptionsv1.F1ApPDu) ([]byte, error) {

	log.Debugf("Obtained F1AP-PDU message is\n%v", f1ap)

	per, err := aper.MarshalWithParams(f1ap, "choiceExt", choiceOptions.F1apChoicemap, choiceOptions.F1apCanonicalChoicemap)
	if err != nil {
		return nil, err
	}
	log.Debugf("Encoded F1AP-PDU PER bytes are\n%v", hex.Dump(per))

	return per, nil
}

func PerDecodeF1ApPdu(per []byte) (*f1appdudescriptionsv1.F1ApPDu, error) {

	log.Debugf("Obtained F1AP-PDU PER bytes are\n%v", hex.Dump(per))

	result := f1appdudescriptionsv1.F1ApPDu{}
	err := aper.UnmarshalWithParams(per, &result, "choiceExt", choiceOptions.F1apChoicemap, choiceOptions.F1apCanonicalChoicemap)
	if err != nil {
		return nil, err
	}

	log.Debugf("Decoded F1AP-PDU from PER is\n%v", &result)

	return &result, nil
}
