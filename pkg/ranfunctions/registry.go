// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package ranfunctions

import (
	"sync"

	"github.com/onosproject/onos-lib-go/pkg/logging"

	e2smtypes "github.com/onosproject/onos-api/go/onos/e2t/e2sm"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/types"
)

var log = logging.GetLogger("registry", "ranfunctions")

// RANFunctions supported RAN functions per service model
type RANFunctions struct {
	ranFunctions map[ID][]RANFunction
	mu           sync.RWMutex
}

// RANFunction RAN function information
type RANFunction struct {
	ID          types.RanFunctionID
	Name        string
	Description string
}

// ID ID for a RAN function
type ID struct {
	oid e2smtypes.OID
	// TODO Node ID should be added as well since RAN function IDs are unique inside each E2 node
}

// NewID creates a key for RANFunction based on OID and node ID
func NewID(oid e2smtypes.OID) ID {
	return ID{
		oid: oid,
	}
}

// NewRegistry creates a new registry
func NewRegistry() *RANFunctions {
	return &RANFunctions{
		ranFunctions: make(map[ID][]RANFunction),
	}

}

// Remove removes a RAN function from registry
func (r *RANFunctions) Remove(id ID, ranFunction RANFunction) error {
	panic("implement me")
}

func (r *RANFunctions) Contains(id ID, ranFunctionID types.RanFunctionID) bool {
	r.mu.RLock()
	defer r.mu.RUnlock()

	ranFunctions, ok := r.ranFunctions[id]
	if !ok {
		return false
	}
	log.Info("Test: ranfunctions:", r.ranFunctions)
	for _, ranFunctionValue := range ranFunctions {
		if ranFunctionID == ranFunctionValue.ID {
			return true
		}
	}
	return false
}

// Add adds a RAN function to the registry
func (r *RANFunctions) Add(id ID, ranFunction RANFunction) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	ranFunctions, ok := r.ranFunctions[id]
	if !ok {
		var ranFunctions []RANFunction
		ranFunctions = append(ranFunctions, ranFunction)
		r.ranFunctions[id] = ranFunctions
		return nil

	}
	for _, ranFunctionValue := range ranFunctions {
		if ranFunctionValue.ID == ranFunction.ID {
			return nil
		}
	}

	ranFunctions = append(ranFunctions, ranFunction)
	r.ranFunctions[id] = ranFunctions
	return nil
}

// Registry register RAN functions for each service model
type Registry interface {
	// Add adds a RAN function for a service model
	Add(id ID, ranFunction RANFunction) error
	// Remove removes a RAN function ID for a service model
	Remove(id ID, ranFunction RANFunction) error
	// Contains checks if a RAN function ID does exist for a service model
	Contains(id ID, ranFunctionID types.RanFunctionID) bool
}

var _ Registry = &RANFunctions{}
