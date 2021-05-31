// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package subscription

import (
	"fmt"
	"github.com/onosproject/onos-lib-go/pkg/errors"
)

type RequestID string

type NodeID string

type AppID string

type InstanceID string

type ServiceModelName string

type ServiceModelVersion string

// Revision is a subscription object revision
type Revision uint64

// Key gets the subscription ID as a string key
func (s *SubscriptionID) Key() string {
	return fmt.Sprintf("%s:%s:%s:%s:%s", s.NodeID, s.RequestID, s.AppID, s.InstanceID, s.Hash)
}

// TaskID returns the task ID for the subscription
func (s *SubscriptionID) TaskID() TaskID {
	return TaskID{
		NodeID:    s.NodeID,
		RequestID: s.RequestID,
		Hash:      s.Hash,
	}
}

// Validate verifies that the ID is valid
func (s *SubscriptionID) Validate() error {
	if s.NodeID == "" {
		return errors.NewInvalid("NodeID is required")
	}
	if s.RequestID == "" {
		return errors.NewInvalid("RequestID is required")
	}
	if s.AppID == "" {
		return errors.NewInvalid("AppID is required")
	}
	if s.InstanceID == "" {
		return errors.NewInvalid("InstanceID is required")
	}
	if s.Hash == "" {
		return errors.NewInvalid("Hash is required")
	}
	return nil
}
