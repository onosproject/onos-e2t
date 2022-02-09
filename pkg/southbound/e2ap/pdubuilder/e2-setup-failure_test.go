// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
package pdubuilder

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/encoder"
	"testing"

	"github.com/onosproject/onos-e2t/api/e2ap/v2"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-commondatatypes"
	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"gotest.tools/assert"
)

func TestE2SetupFailure(t *testing.T) {
	//ttw1 := e2ap_ies.TimeToWait_TIME_TO_WAIT_V10S
	//procCode1 := v21.ProcedureCodeIDRICsubscription
	//criticality1 := e2apcommondatatypes.Criticality_CRITICALITY_IGNORE
	//ftg1 := e2apcommondatatypes.TriggeringMessage_TRIGGERING_MESSAGE_UNSUCCESSFUL_OUTCOME
	//
	//tnlInfo1, err := pdubuilder.CreateTnlInformation(&asn1.BitString{
	//	Value: []byte{0x00, 0x00, 0x01},
	//	Len:   24,
	//})
	//assert.NilError(t, err)
	//tnlInfo1.SetTnlPort(&asn1.BitString{
	//	Value: []byte{0x00, 0x01},
	//	Len:   16,
	//})
	//
	//e2apPdu, err := pdubuilder.CreateE2SetupFailurePdu(1,
	//	&e2ap_ies.Cause{
	//		Cause: &e2ap_ies.Cause_Misc{ // Probably, could be any other reason
	//			Misc: e2ap_ies.CauseMisc_CAUSE_MISC_UNSPECIFIED,
	//		},
	//	})
	//assert.NilError(t, err)
	//assert.Assert(t, e2apPdu != nil)
	//
	//e2apPdu.GetUnsuccessfulOutcome().GetProcedureCode().GetE2Setup().GetUnsuccessfulOutcome().
	//	SetTimeToWait(ttw1).SetTnlInformation(tnlInfo1).SetCriticalityDiagnostics(&procCode1, &criticality1, &ftg1,
	//	&types1.RicRequest{
	//		RequestorID: 10,
	//		InstanceID:  20,
	//	}, []*types1.CritDiag{
	//		{
	//			TypeOfError:   e2ap_ies.TypeOfError_TYPE_OF_ERROR_MISSING,
	//			IECriticality: e2apcommondatatypes.Criticality_CRITICALITY_IGNORE,
	//			IEId:          v21.ProtocolIeIDRicsubscriptionDetails,
	//		},
	//	})
	////fmt.Printf("TimeToWait is \n%v\n", newE2apPdu.GetUnsuccessfulOutcome().GetProcedureCode().GetE2Setup().GetUnsuccessfulOutcome().GetProtocolIes().GetE2ApProtocolIes31())
	//
	//per, err := asn1cgo.PerEncodeE2apPdu(e2apPdu)
	//assert.NilError(t, err)
	//t.Logf("E2SetupFailure E2AP PDU PER\n%v", hex.Dump(per))

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

	perNew, err := encoder.PerEncodeE2ApPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("E2SetupFailure E2AP PDU PER with Go APER library\n%v", hex.Dump(perNew))

	//Comparing reference PER bytes with Go APER library produced
	//assert.DeepEqual(t, per, perNew)

	result, err := encoder.PerDecodeE2ApPdu(perNew)
	assert.NilError(t, err)
	assert.DeepEqual(t, newE2apPdu.String(), result.String())

	// Decoding the message from the APER bytes produced by CGo
	//result11, err := encoder.PerDecodeE2ApPdu(per)
	//assert.NilError(t, err)
	//assert.DeepEqual(t, newE2apPdu.String(), result11.String())
	//
	//result1, err := asn1cgo.PerDecodeE2apPdu(perNew)
	//assert.NilError(t, err)
	//t.Logf("E2SetupFailure E2AP PDU PER - decoded\n%v", result1)
	//assert.DeepEqual(t, e2apPdu.String(), result1.String())
}
