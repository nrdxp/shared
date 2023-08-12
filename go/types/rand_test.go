package types

import (
	"testing"

	"github.com/candiddev/shared/go/assert"
)

func TestRandSt(t *testing.T) {
	n1 := RandString(10)
	n2 := RandString(10)

	assert.Equal(t, n1 != n2, true)
	assert.Equal(t, len(n1), 10)
}
