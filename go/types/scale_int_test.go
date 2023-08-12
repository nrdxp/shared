package types

import (
	"encoding/json"
	"testing"

	"github.com/candiddev/shared/go/assert"
)

type scaleInt struct {
	ScaleInt ScaleInt `json:"scaleInt"`
}

func TestScaleIntUnmarshal(t *testing.T) {
	tests := map[string]struct {
		err   error
		input string
		want  scaleInt
	}{
		"negative": {
			err:   ErrScaleInt,
			input: `{"scaleInt":-1}`,
			want:  scaleInt{},
		},
		"big": {
			err:   ErrScaleInt,
			input: `{"scaleInt":6}`,
			want:  scaleInt{},
		},
		"good": {
			input: `{"scaleInt":5}`,
			want: scaleInt{
				ScaleInt: 5,
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			var c scaleInt

			assert.Equal(t, json.Unmarshal([]byte(tc.input), &c), tc.err)
			assert.Equal(t, c, tc.want)
		})
	}
}
