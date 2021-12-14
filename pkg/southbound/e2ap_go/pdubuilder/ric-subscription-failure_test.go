// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0
package pdubuilder

import (
	"encoding/hex"
	v21 "github.com/onosproject/onos-e2t/api/e2ap/v2"
	e2apcommondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-commondatatypes"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/asn1cgo"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/pdubuilder"
	types1 "github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap_go/encoder"
	"testing"

	"github.com/onosproject/onos-e2t/api/e2ap_go/v2"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap_go/v2/e2ap-commondatatypes"
	e2apies "github.com/onosproject/onos-e2t/api/e2ap_go/v2/e2ap-ies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap_go/types"
	"gotest.tools/assert"
)

func TestRicSubscriptionFailure(t *testing.T) {
	ricActionsNotAdmittedList1 := make(map[types1.RicActionID]*e2ap_ies.Cause)
	ricActionsNotAdmittedList1[100] = &e2ap_ies.Cause{
		Cause: &e2ap_ies.Cause_Transport{
			Transport: e2ap_ies.CauseTransport_CAUSE_TRANSPORT_TRANSPORT_RESOURCE_UNAVAILABLE,
		},
	}
	ricActionsNotAdmittedList1[200] = &e2ap_ies.Cause{
		Cause: &e2ap_ies.Cause_Misc{
			Misc: e2ap_ies.CauseMisc_CAUSE_MISC_HARDWARE_FAILURE,
		},
	}

	procCode1 := v21.ProcedureCodeIDRICsubscription
	criticality1 := e2apcommondatatypes.Criticality_CRITICALITY_IGNORE
	ftg1 := e2apcommondatatypes.TriggeringMessage_TRIGGERING_MESSAGE_UNSUCCESSFUL_OUTCOME

	e2apPdu, err := pdubuilder.CreateRicSubscriptionFailureE2apPdu(&types1.RicRequest{
		RequestorID: 22,
		InstanceID:  6,
	}, 9, &e2ap_ies.Cause{
		Cause: &e2ap_ies.Cause_Misc{
			Misc: e2ap_ies.CauseMisc_CAUSE_MISC_CONTROL_PROCESSING_OVERLOAD,
		},
	})
	assert.NilError(t, err)
	assert.Assert(t, e2apPdu != nil)
	e2apPdu.GetUnsuccessfulOutcome().GetProcedureCode().GetRicSubscription().GetUnsuccessfulOutcome().
		SetCriticalityDiagnostics(&procCode1, &criticality1, &ftg1,
			&types1.RicRequest{
				RequestorID: 10,
				InstanceID:  20,
			}, []*types1.CritDiag{
				{
					TypeOfError:   e2ap_ies.TypeOfError_TYPE_OF_ERROR_MISSING,
					IECriticality: e2apcommondatatypes.Criticality_CRITICALITY_IGNORE,
					IEId:          v21.ProtocolIeIDRicsubscriptionDetails,
				},
			})

	per, err := asn1cgo.PerEncodeE2apPdu(e2apPdu)
	assert.NilError(t, err)
	t.Logf("RicSubscriptionFailure E2AP PDU PER\n%v", hex.Dump(per))

	ricActionsNotAdmittedList := make(map[types.RicActionID]*e2apies.Cause)
	ricActionsNotAdmittedList[100] = &e2apies.Cause{
		Cause: &e2apies.Cause_Transport{
			Transport: e2apies.CauseTransport_CAUSE_TRANSPORT_TRANSPORT_RESOURCE_UNAVAILABLE,
		},
	}
	ricActionsNotAdmittedList[200] = &e2apies.Cause{
		Cause: &e2apies.Cause_Misc{
			Misc: e2apies.CauseMisc_CAUSE_MISC_HARDWARE_FAILURE,
		},
	}

	procCode := v2.ProcedureCodeIDRICsubscription
	criticality := e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE
	ftg := e2ap_commondatatypes.TriggeringMessage_TRIGGERING_MESSAGE_UNSUCCESSFUL_OUTCOME

	newE2apPdu, err := CreateRicSubscriptionFailureE2apPdu(&types.RicRequest{
		RequestorID: 22,
		InstanceID:  6,
	}, 9, &e2apies.Cause{
		Cause: &e2apies.Cause_Misc{
			Misc: e2apies.CauseMisc_CAUSE_MISC_CONTROL_PROCESSING_OVERLOAD,
		},
	})
	assert.NilError(t, err)
	assert.Assert(t, newE2apPdu != nil)
	newE2apPdu.GetUnsuccessfulOutcome().GetValue().GetRicSubscription().
		SetCriticalityDiagnostics(&procCode, &criticality, &ftg,
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
	t.Logf("RicSubscriptionFailure E2AP PDU PER with Go APER library\n%v", hex.Dump(perNew))

	//Comparing reference PER bytes with Go APER library produced
	assert.DeepEqual(t, per, perNew)

	//e2apPdu, err := encoder.PerDecodeE2ApPdu(perNew)
	//assert.NilError(t, err)
	//assert.DeepEqual(t, newE2apPdu.String(), e2apPdu.String())

	result1, err := asn1cgo.PerDecodeE2apPdu(perNew)
	assert.NilError(t, err)
	t.Logf("RicSubscriptionFailure E2AP PDU PER - decoded\n%v\n", result1)
	assert.DeepEqual(t, e2apPdu.String(), result1.String())
}

func TestRicSubscriptionFailureExcludeOptionalIE(t *testing.T) {
	ricActionsNotAdmittedList1 := make(map[types1.RicActionID]*e2ap_ies.Cause)
	ricActionsNotAdmittedList1[100] = &e2ap_ies.Cause{
		Cause: &e2ap_ies.Cause_Transport{
			Transport: e2ap_ies.CauseTransport_CAUSE_TRANSPORT_TRANSPORT_RESOURCE_UNAVAILABLE,
		},
	}
	ricActionsNotAdmittedList1[200] = &e2ap_ies.Cause{
		Cause: &e2ap_ies.Cause_Misc{
			Misc: e2ap_ies.CauseMisc_CAUSE_MISC_HARDWARE_FAILURE,
		},
	}

	//procCode := v2.ProcedureCodeIDRICsubscription
	//criticality := e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE
	//ftg := e2ap_commondatatypes.TriggeringMessage_TRIGGERING_MESSAGE_UNSUCCESSFUL_OUTCOME

	e2apPdu, err := pdubuilder.CreateRicSubscriptionFailureE2apPdu(&types1.RicRequest{
		RequestorID: 22,
		InstanceID:  6,
	}, 9,
		&e2ap_ies.Cause{
			Cause: &e2ap_ies.Cause_Misc{
				Misc: e2ap_ies.CauseMisc_CAUSE_MISC_CONTROL_PROCESSING_OVERLOAD,
			},
		})
	assert.NilError(t, err)
	assert.Assert(t, e2apPdu != nil)

	per, err := asn1cgo.PerEncodeE2apPdu(e2apPdu)
	assert.NilError(t, err)
	t.Logf("RicSubscriptionFailure E2AP PDU PER\n%v", hex.Dump(per))

	ricActionsNotAdmittedList := make(map[types.RicActionID]*e2apies.Cause)
	ricActionsNotAdmittedList[100] = &e2apies.Cause{
		Cause: &e2apies.Cause_Transport{
			Transport: e2apies.CauseTransport_CAUSE_TRANSPORT_TRANSPORT_RESOURCE_UNAVAILABLE,
		},
	}
	ricActionsNotAdmittedList[200] = &e2apies.Cause{
		Cause: &e2apies.Cause_Misc{
			Misc: e2apies.CauseMisc_CAUSE_MISC_HARDWARE_FAILURE,
		},
	}

	//procCode := v2.ProcedureCodeIDRICsubscription
	//criticality := e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE
	//ftg := e2ap_commondatatypes.TriggeringMessage_TRIGGERING_MESSAGE_UNSUCCESSFUL_OUTCOME

	newE2apPdu, err := CreateRicSubscriptionFailureE2apPdu(&types.RicRequest{
		RequestorID: 22,
		InstanceID:  6,
	}, 9,
		&e2apies.Cause{
			Cause: &e2apies.Cause_Misc{
				Misc: e2apies.CauseMisc_CAUSE_MISC_CONTROL_PROCESSING_OVERLOAD,
			},
		})
	assert.NilError(t, err)
	assert.Assert(t, newE2apPdu != nil)

	perNew, err := encoder.PerEncodeE2ApPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("RicSubscriptionFailure E2AP PDU PER with Go APER library\n%v", hex.Dump(perNew))

	//Comparing reference PER bytes with Go APER library produced
	assert.DeepEqual(t, per, perNew)

	//e2apPdu, err := encoder.PerDecodeE2ApPdu(perNew)
	//assert.NilError(t, err)
	//assert.DeepEqual(t, newE2apPdu.String(), e2apPdu.String())

	result1, err := asn1cgo.PerDecodeE2apPdu(perNew)
	assert.NilError(t, err)
	t.Logf("RicSubscriptionDeleteFailure E2AP PDU PER - decoded\n%v\n", result1)
	assert.DeepEqual(t, e2apPdu.String(), result1.String())
}
