// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0
package pdubuilder

import (
	"encoding/hex"
	"testing"

	"github.com/onosproject/onos-e2t/api/e2ap/v2beta1"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-commondatatypes"
	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-ies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/asn1cgo"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"gotest.tools/assert"
)

func TestE2SetupFailure(t *testing.T) {
	ttw := e2apies.TimeToWait_TIME_TO_WAIT_V10S
	procCode := v2beta1.ProcedureCodeIDRICsubscription
	criticality := e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE
	ftg := e2ap_commondatatypes.TriggeringMessage_TRIGGERING_MESSAGE_UNSUCCESSFUL_OUTCOME

	tnlInfo, err := CreateTnlInformation(&asn1.BitString{
		Value: []byte{0x00, 0x00, 0x01},
		Len:   24,
	})
	assert.NilError(t, err)
	tnlInfo.SetTnlPort(&asn1.BitString{
		Value: []byte{0x00, 0x01},
		Len:   16,
	})

	newE2apPdu, err := CreateE2SetupFailurePdu(1,
		&e2apies.Cause{
			Cause: &e2apies.Cause_Misc{ // Probably, could be any other reason
				Misc: e2apies.CauseMisc_CAUSE_MISC_UNSPECIFIED,
			},
		})
	assert.NilError(t, err)
	assert.Assert(t, newE2apPdu != nil)

	newE2apPdu.GetUnsuccessfulOutcome().GetProcedureCode().GetE2Setup().GetUnsuccessfulOutcome().
		SetTimeToWait(ttw).SetTnlInformation(tnlInfo).SetCriticalityDiagnostics(&procCode, &criticality, &ftg,
		&types.RicRequest{
			RequestorID: 10,
			InstanceID:  20,
		}, []*types.CritDiag{
			{
				TypeOfError:   e2apies.TypeOfError_TYPE_OF_ERROR_MISSING,
				IECriticality: e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE,
				IEId:          v2beta1.ProtocolIeIDRicsubscriptionDetails,
			},
		})
	//fmt.Printf("TimeToWait is \n%v\n", newE2apPdu.GetUnsuccessfulOutcome().GetProcedureCode().GetE2Setup().GetUnsuccessfulOutcome().GetProtocolIes().GetE2ApProtocolIes31())

	xer, err := asn1cgo.XerEncodeE2apPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("ErrorIndication E2AP PDU XER\n%s", string(xer))

	result, err := asn1cgo.XerDecodeE2apPdu(xer)
	assert.NilError(t, err)
	t.Logf("ErrorIndication E2AP PDU XER - decoded\n%v", result)
	assert.DeepEqual(t, newE2apPdu.String(), result.String())

	per, err := asn1cgo.PerEncodeE2apPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("ErrorIndication E2AP PDU PER\n%v", hex.Dump(per))

	result1, err := asn1cgo.PerDecodeE2apPdu(per)
	assert.NilError(t, err)
	t.Logf("ErrorIndication E2AP PDU PER - decoded\n%v", result1)
	assert.DeepEqual(t, newE2apPdu.String(), result1.String())
}
