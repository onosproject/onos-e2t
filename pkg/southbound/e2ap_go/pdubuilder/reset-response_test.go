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
	"gotest.tools/assert"
)

func TestResetResponse(t *testing.T) {
	procCode := v2.ProcedureCodeIDReset
	criticality := e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE
	ftg := e2ap_commondatatypes.TriggeringMessage_TRIGGERING_MESSAGE_UNSUCCESSFUL_OUTCOME

	newE2apPdu, err := CreateResetResponseE2apPdu(1)
	assert.NilError(t, err)
	assert.Assert(t, newE2apPdu != nil)
	newE2apPdu.GetSuccessfulOutcome().GetValue().GetReset_().
		SetCriticalityDiagnostics(procCode, &criticality, &ftg,
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
	t.Logf("ResetResponse E2AP PDU PER with Go APER library\n%v", hex.Dump(perNew))

	//Comparing reference PER bytes with Go APER library produced
	//assert.DeepEqual(t, per, perNew)

	e2apPdu, err := encoder.PerDecodeE2ApPdu(perNew)
	assert.NilError(t, err)
	assert.DeepEqual(t, newE2apPdu.String(), e2apPdu.String())

	//per, err := asn1cgo.PerEncodeE2apPdu(newE2apPdu)
	//assert.NilError(t, err)
	//t.Logf("ResetResponse E2AP PDU PER\n%v", hex.Dump(per))
	//
	//e2apPdu, err = asn1cgo.PerDecodeE2apPdu(per)
	//assert.NilError(t, err)
	//assert.DeepEqual(t, newE2apPdu.String(), e2apPdu.String())
}
