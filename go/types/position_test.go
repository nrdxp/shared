package types

import (
	"encoding/json"
	"testing"

	"github.com/candiddev/shared/go/assert"
)

type position struct {
	Position Position `json:"position"`
}

func TestPositionUnmarshal(t *testing.T) {
	tests := map[string]struct {
		err   error
		input string
		want  position
	}{
		"empty": {
			input: `{"position":""}`,
			want:  position{},
		},
		"null": {
			input: `{"position":null}`,
			want:  position{},
		},
		"invalid": {
			err:   ErrPosition,
			input: `{"position":"aaacc1"}`,
			want:  position{},
		},
		"good - number": {
			input: `{"position":"1"}`,
			want: position{
				Position: "1",
			},
		},
		"good - string": {
			input: `{"position":"1:aaa"}`,
			want: position{
				Position: "1:aaa",
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			var c position

			assert.Equal(t, json.Unmarshal([]byte(tc.input), &c), tc.err)
			assert.Equal(t, c, tc.want)
		})
	}
}
