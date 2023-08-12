package types

import (
	"encoding/json"
	"testing"

	"github.com/candiddev/shared/go/assert"
)

type positiveInt struct {
	PositiveInt PositiveInt `json:"positiveInt"`
}

func TestPositiveIntUnmarshal(t *testing.T) {
	tests := map[string]struct {
		err   error
		input string
		want  positiveInt
	}{
		"negative": {
			err:   ErrPositiveInt,
			input: `{"positiveInt":-1}`,
			want:  positiveInt{},
		},
		"good": {
			input: `{"positiveInt":13}`,
			want: positiveInt{
				PositiveInt: 13,
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			var c positiveInt

			assert.Equal(t, json.Unmarshal([]byte(tc.input), &c), tc.err)
			assert.Equal(t, c, tc.want)
		})
	}
}
