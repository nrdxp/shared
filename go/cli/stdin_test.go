package cli

import (
	"testing"

	"github.com/candiddev/shared/go/assert"
)

func TestStdin(t *testing.T) {
	SetStdin("hello")
	assert.Equal(t, ReadStdin(), "hello")
}
