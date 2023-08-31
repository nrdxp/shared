package errs

import (
	"errors"
	"testing"

	"github.com/candiddev/shared/go/assert"
)

func TestErrAppend(t *testing.T) {
	err1 := errors.New("hello")
	err2 := errors.New("world")
	e1 := ErrClientBadRequestMissing.Append(err1)
	e2 := ErrClientBadRequestMissing.Append(err2)

	assert.Equal(t, e1.Is(err1), true)
	assert.Equal(t, e1.Is(err2), false)
	assert.Equal(t, e2.Is(err1), false)
	assert.Equal(t, e2.Is(err2), true)
}
