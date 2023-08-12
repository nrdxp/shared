package types

import (
	"encoding/json"
	"testing"

	"github.com/candiddev/shared/go/assert"
)

type color struct {
	Color Color `json:"color"`
}

func TestColorUnmarshal(t *testing.T) {
	tests := map[string]struct {
		err   bool
		input string
		want  color
	}{
		"less": {
			err:   true,
			input: `{"color":-1}`,
			want:  color{},
		},
		"greater": {
			err:   true,
			input: `{"color":14}`,
			want:  color{},
		},
		"good": {
			input: `{"color":1}`,
			want: color{
				Color: 1,
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			var c color

			assert.Equal(t, json.Unmarshal([]byte(tc.input), &c) != nil, tc.err)
			assert.Equal(t, c, tc.want)
		})
	}
}
