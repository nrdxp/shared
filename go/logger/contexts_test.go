package logger

import (
	"context"
	"testing"

	"github.com/candiddev/shared/go/assert"
)

func TestAttributes(t *testing.T) {
	ctx := SetAttribute(context.Background(), "key", "value")

	assert.Equal(t, "value", GetAttribute(ctx, "key"))

	ctx = SetAttribute(ctx, "key", "value2")
	ctx = SetAttribute(ctx, "key1", "value")

	assert.Equal(t, "value2", GetAttribute(ctx, "key"))
	assert.Equal(t, "value", GetAttribute(ctx, "key1"))
}

func TestDebug(t *testing.T) {
	ctx := SetDebug(context.Background(), true)

	assert.Equal(t, true, GetDebug(ctx))

	ctx = SetDebug(ctx, true)

	assert.Equal(t, true, GetDebug(ctx))

	ctx = SetDebug(ctx, false)

	assert.Equal(t, false, GetDebug(ctx))
}
