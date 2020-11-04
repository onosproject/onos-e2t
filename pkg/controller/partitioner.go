// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package controller

import (
	"regexp"

	"github.com/onosproject/onos-e2t/api/types"
)

// PartitionKey is a key by which to partition requests
type PartitionKey string

const staticKey PartitionKey = ""

// WorkPartitioner is an interface for partitioning requests among a set of goroutines
// The WorkPartitioner can enable safe concurrency inside controllers. For each request, the
// partitioner will be called to provide a PartitionKey for the request. For each unique
// PartitionKey, a separate channel and goroutine will be created to process requests for
// the partition.
type WorkPartitioner interface {
	// Partition gets a partition key for the given request
	Partition(id types.ID) (PartitionKey, error)
}

// UnaryPartitioner is a WorkPartitioner that assigns all work to a single goroutine
type UnaryPartitioner struct {
}

// Partition returns a static partition key
func (p *UnaryPartitioner) Partition(id types.ID) (PartitionKey, error) {
	return staticKey, nil
}

var _ WorkPartitioner = &UnaryPartitioner{}

// RegexpPartitioner is a WorkPartitioner that assigns work to a gouroutine per regex output
type RegexpPartitioner struct {
	Regexp regexp.Regexp
}

// Partition returns a PartitionKey from the configured regex
func (p *RegexpPartitioner) Partition(id types.ID) (PartitionKey, error) {
	return PartitionKey(p.Regexp.FindString(string(id))), nil
}

var _ WorkPartitioner = &RegexpPartitioner{}
