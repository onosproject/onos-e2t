// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package channel

// ID is a connection identifier
type ID string

// PlmnID is a connection identifier
type PlmnID string

// RANFunctionID is a RAN function identifier
type RANFunctionID int32

// RANFunctionDescription is a RAN function description
type RANFunctionDescription string

// RANFunctionRevision is a RAN function revision
type RANFunctionRevision string

// RANFunctionMetadata is the metadata for a RAN function
type RANFunctionMetadata struct {
	Description RANFunctionDescription
	Revision    RANFunctionRevision
}

// Metadata is connection metadata
type Metadata struct {
	// ID is the connection identifier
	ID ID

	// PlmnID is the PLMN identifier
	PlmnID PlmnID

	// RANFunctions is a map of RAN functions
	RANFunctions map[RANFunctionID]RANFunctionMetadata
}

// GetRanFunction gets a RAN function by ID
func (m Metadata) GetRANFunction(id RANFunctionID) RANFunctionMetadata {
	return m.RANFunctions[id]
}
