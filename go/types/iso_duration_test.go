package types

import (
	"testing"

	"github.com/candiddev/shared/go/assert"
)

func TestISODurationMinutes(t *testing.T) {
	tests := map[string]int{
		"":           0,
		"P1DT30M":    1470,
		"PT30M":      30,
		"PT1H30M":    90,
		"PT1H30M45S": 91,
	}

	for input, want := range tests {
		t.Run(input, func(t *testing.T) {
			assert.Equal(t, ISODuration(input).Minutes(), want)
		})
	}
}
