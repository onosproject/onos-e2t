// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package types

import e2apies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"

type RanFunctionDescription []byte
type RanFunctionRevision int
type RanFunctionOID string
type RanFunctionID uint8

type RanFunctionItem struct {
	Description RanFunctionDescription
	Revision    RanFunctionRevision
	OID         RanFunctionOID
}

type RanFunctions map[RanFunctionID]RanFunctionItem

type RanFunctionRevisions map[RanFunctionID]RanFunctionRevision

type RanFunctionCauses map[RanFunctionID]*e2apies.Cause
