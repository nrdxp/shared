package types

import (
	"testing"

	"github.com/candiddev/shared/go/assert"
)

func TestTimeZoneUnmarshalJSON(t *testing.T) {
	tests := map[string]struct {
		err   error
		input string
		want  string
	}{
		"invalid": {
			err:   ErrTimeZone,
			input: "1",
			want:  "",
		},
		"not a time zone": {
			err:   ErrTimeZone,
			input: `"somewhere/faraway"`,
			want:  "",
		},
		"good": {
			input: `"America/Chicago"`,
			want:  "America/Chicago",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			var tz TimeZone

			assert.Equal(t, tz.UnmarshalJSON([]byte(tc.input)), tc.err)
			assert.Equal(t, tz.String(), tc.want)
		})
	}
}
