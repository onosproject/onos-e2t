// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package subscription

import (
	"fmt"
	"github.com/onosproject/onos-lib-go/pkg/errors"
)

// Key gets the task ID as a string key
func (s *TaskID) Key() string {
	return fmt.Sprintf("%s:%s:%s", s.NodeID, s.RequestID, s.Hash)
}

// Validate verifies that the ID is valid
func (s *TaskID) Validate() error {
	if s.NodeID == "" {
		return errors.NewInvalid("NodeID is required")
	}
	if s.RequestID == "" {
		return errors.NewInvalid("RequestID is required")
	}
	if s.Hash == "" {
		return errors.NewInvalid("Hash is required")
	}
	return nil
}
