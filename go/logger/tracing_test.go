package logger

import (
	"context"
	"testing"

	"github.com/candiddev/shared/go/assert"
	"go.opentelemetry.io/otel/trace"
)

func TestSetupTracing(t *testing.T) {
	ctx := context.Background()

	c := TracingConfig{
		Endpoint:     "dev.homechart.app",
		SamplerRatio: 0.5,
	}

	_, err := SetupTracing(ctx, c, "example.com", "1.0")

	assert.HasErr(t, err, nil)

	ctx = Trace(ctx)
	span := trace.SpanFromContext(ctx)

	defer span.End()
}
