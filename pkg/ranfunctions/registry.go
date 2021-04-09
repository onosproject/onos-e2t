// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package ranfunctions

import (
	"sync"

	e2smtypes "github.com/onosproject/onos-api/go/onos/e2t/e2sm"
	"github.com/onosproject/onos-lib-go/pkg/errors"

	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/types"
)

// RANFunctions supported RAN functions per service model
type RANFunctions struct {
	ranFunctions map[ID]RANFunction
	mu           sync.RWMutex
}

// RANFunction RAN function information
type RANFunction struct {
	OID e2smtypes.OID
	ID  types.RanFunctionID
	// protobuf encoded description
	Description []byte
}

// ID ID for a RAN function
type ID struct {
	oid    e2smtypes.OID
	nodeID string
}

// NewID creates a key for RANFunction based on OID and node ID
func NewID(oid e2smtypes.OID, nodeID string) ID {
	return ID{
		oid:    oid,
		nodeID: nodeID,
	}
}

// NewRegistry creates a new registry
func NewRegistry() *RANFunctions {
	return &RANFunctions{
		ranFunctions: make(map[ID]RANFunction),
	}

}

// Remove removes a RAN function from registry
func (r *RANFunctions) Remove(id ID, ranFunction RANFunction) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if id.nodeID == "" || id.oid == "" {
		return errors.NewInvalid("node ID or OID cannot be empty")
	}
	delete(r.ranFunctions, id)
	return nil
}

// Add adds a RAN function to the registry
func (r *RANFunctions) Add(id ID, ranFunction RANFunction) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.ranFunctions[id] = ranFunction
	return nil
}

// Get gets a RANFunction per id
func (r *RANFunctions) Get(id ID) (RANFunction, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	ranFunction, ok := r.ranFunctions[id]
	if !ok {
		return RANFunction{}, errors.New(errors.NotFound, "ran function has not been found")
	}
	return ranFunction, nil
}

// GetRANFunctionsPerNode ...
func (r *RANFunctions) GetRANFunctionsByNodeID(nodeID string) []RANFunction {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var ranFunctions []RANFunction
	for ranFunctionID, ranFunction := range r.ranFunctions {
		if nodeID == ranFunctionID.nodeID {
			ranFunctions = append(ranFunctions, ranFunction)
		}
	}

	return ranFunctions
}

// Registry register RAN functions for each service model
type Registry interface {
	// Add adds a RAN function for a service model
	Add(id ID, ranFunction RANFunction) error
	// Remove removes a RAN function ID for a service model
	Remove(id ID, ranFunction RANFunction) error

	// Get gets a RANFunction per id
	Get(id ID) (RANFunction, error)

	// GetRANFunctionsPerNode get RAN functions per node
	GetRANFunctionsByNodeID(nodeID string) []RANFunction
}

var _ Registry = &RANFunctions{}
