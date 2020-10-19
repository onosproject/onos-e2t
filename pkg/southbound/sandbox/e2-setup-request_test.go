package sandbox

import (
	"gotest.tools/assert"
	"testing"
)

func TestE2SetupRequest(t *testing.T) {
	newE2apPdu := createE2apPdu()
	assert.Assert(t, newE2apPdu != nil)
}
