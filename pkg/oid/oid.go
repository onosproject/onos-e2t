// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package oid

import (
	"strconv"
	"sync"

	"github.com/onosproject/onos-lib-go/pkg/errors"

	"github.com/onosproject/onos-lib-go/pkg/logging"
)

var log = logging.GetLogger("oid", "registry")

// ObjectIdentifiers OIDs information
type ObjectIdentifiers struct {
	nodes map[string]Oid
	mu    sync.RWMutex
}

// AddOid adds oid
func (r *ObjectIdentifiers) AddOid(key string, oid Oid) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if key == "" {
		return errors.NewInvalid("oid key cannot be empty")
	}
	r.nodes[key] = oid
	return nil
}

// Registry oid registry interface
type Registry interface {
	// GetOid gets oid based on a given key
	GetOid(key string) Oid
	// AddOid adds oid
	AddOid(key string, oid Oid) error

	// TODO add more functions as needed
}

// NewOidRegistry creates a new oid registry
func NewOidRegistry() *ObjectIdentifiers {
	nodes := defaultNodes

	return &ObjectIdentifiers{
		nodes: nodes,
	}
}

func getOid(r Registry, key string) string {
	oid := r.GetOid(key)
	return strconv.Itoa(int(oid))
}

// GetOid gets oid based on a given key
func (r *ObjectIdentifiers) GetOid(key string) Oid {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.nodes[key]
}

var _ Registry = &ObjectIdentifiers{}
