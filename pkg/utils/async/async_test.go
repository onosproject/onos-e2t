// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package async

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestApply(t *testing.T) {
	values := []string{
		"one",
		"two",
		"three",
	}
	err := Apply(len(values), func(i int) error {
		return nil
	})
	assert.NoError(t, err)
}

func TestExecute(t *testing.T) {
	values := []string{
		"one",
		"two",
		"three",
	}
	results, err := Execute(len(values), func(i int) (interface{}, error) {
		return values[i], nil
	})
	assert.NoError(t, err)
	assert.NotNil(t, results)
	assert.NotEqual(t, "", results[0].(string))
	assert.NotEqual(t, "", results[1].(string))
	assert.NotEqual(t, "", results[2].(string))
}
