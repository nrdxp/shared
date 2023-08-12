package types

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/candiddev/shared/go/assert"
)

type stringLimit struct {
	StringLimit StringLimit `json:"stringLimit"`
}

func TestStringLimitUnmarshal(t *testing.T) {
	tests := map[string]struct {
		err   error
		input string
		want  stringLimit
	}{
		"too much": {
			err:   ErrStringLimit,
			input: fmt.Sprintf(`{"stringLimit":"%1001v"}`, "a"),
			want: stringLimit{
				StringLimit: "",
			},
		},
		"nothing": {
			input: `{"stringLimit":""}`,
			want: stringLimit{
				StringLimit: "",
			},
		},
		"good": {
			input: `{"stringLimit":"Hello"}`,
			want: stringLimit{
				StringLimit: "Hello",
			},
		},
		"quoted": {
			input: `{"stringLimit":"Hello\""}`,
			want: stringLimit{
				StringLimit: `Hello"`,
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			var c stringLimit

			assert.Equal(t, json.Unmarshal([]byte(tc.input), &c), tc.err)
			assert.Equal(t, c, tc.want)
		})
	}
}
