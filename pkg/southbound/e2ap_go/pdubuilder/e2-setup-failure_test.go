// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0
package pdubuilder

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap_go/encoder"
	"testing"

	"github.com/onosproject/onos-e2t/api/e2ap_go/v2"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap_go/v2/e2ap-commondatatypes"
	e2apies "github.com/onosproject/onos-e2t/api/e2ap_go/v2/e2ap-ies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap_go/types"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"gotest.tools/assert"
)

func TestE2SetupFailure(t *testing.T) {
	ttw := e2apies.TimeToWait_TIME_TO_WAIT_V10S
	procCode := v2.ProcedureCodeIDRICsubscription
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

	newE2apPdu.GetUnsuccessfulOutcome().GetValue().GetE2Setup().
		SetTimeToWait(ttw).SetTnlInformation(tnlInfo).SetCriticalityDiagnostics(&procCode, &criticality, &ftg,
		&types.RicRequest{
			RequestorID: 10,
			InstanceID:  20,
		}, []*types.CritDiag{
			{
				TypeOfError:   e2apies.TypeOfError_TYPE_OF_ERROR_MISSING,
				IECriticality: e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE,
				IEId:          v2.ProtocolIeIDRicsubscriptionDetails,
			},
		})
	//fmt.Printf("TimeToWait is \n%v\n", newE2apPdu.GetUnsuccessfulOutcome().GetProcedureCode().GetE2Setup().GetUnsuccessfulOutcome().GetProtocolIes().GetE2ApProtocolIes31())

	perNew, err := encoder.PerEncodeE2ApPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("E2SetupFailure E2AP PDU PER with Go APER library\n%v", hex.Dump(perNew))

	//Comparing reference PER bytes with Go APER library produced
	//assert.DeepEqual(t, per, perNew)

	//e2apPdu, err := encoder.PerDecodeE2ApPdu(perNew)
	//assert.NilError(t, err)
	//assert.DeepEqual(t, newE2apPdu.String(), e2apPdu.String())

	//per, err := asn1cgo.PerEncodeE2apPdu(newE2apPdu)
	//assert.NilError(t, err)
	//t.Logf("ErrorIndication E2AP PDU PER\n%v", hex.Dump(per))
	//
	//result1, err := asn1cgo.PerDecodeE2apPdu(per)
	//assert.NilError(t, err)
	//t.Logf("ErrorIndication E2AP PDU PER - decoded\n%v", result1)
	//assert.DeepEqual(t, newE2apPdu.String(), result1.String())
}
